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
		cache_Control_Monad_Except_Trans_lift = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
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

func Call_Control_Monad_Except_Trans_lift(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_4})})
}))
})
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
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorExceptT1_1_0 -> gopurs_runtime.Value
functorExceptT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_2))
_ = __local_var_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, v_4)
})
}))
_ = functorExceptT1_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_4 -> *Constructor_Control_Bind_Bind
Bind1_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_4
// TAST (Let): pure_4_5 -> gopurs_runtime.Value
pure_4_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_5
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_4.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t6 = gopurs_runtime.Apply(pure_4_5, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_6
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t6 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
})
}))
}))
_ = __local_var_2_3
// TAST (Let): Bind1_3_7 -> *Constructor_Control_Bind_Bind
Bind1_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_7
// TAST (Let): Applicative0_4_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_8
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_1_0
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_7.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_7.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_8.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}

func Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_10 -> gopurs_runtime.Value
__local_var_1_10 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_10
// TAST (Let): __local_var_1_9 -> gopurs_runtime.Value
__local_var_1_9 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_10, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_2})})
})
_ = __local_var_1_9
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): functorExceptT1_2_0 -> gopurs_runtime.Value
functorExceptT1_2_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_3))
_ = __local_var_4_2
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, v_5)
})
}))
_ = functorExceptT1_2_0
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_4 -> *Constructor_Control_Bind_Bind
Bind1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_4
// TAST (Let): pure_5_5 -> gopurs_runtime.Value
pure_5_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_5
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t6 = gopurs_runtime.Apply(pure_5_5, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_6
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t6 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
})
}))
}))
_ = __local_var_3_3
// TAST (Let): Bind1_4_7 -> *Constructor_Control_Bind_Bind
Bind1_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_7
// TAST (Let): Applicative0_5_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_8
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_2_0
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_7.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_7.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_8.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_9, x_2)
}))
}

func Call_Control_Monad_Except_Trans_semigroupExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): functorExceptT1_1_1 -> gopurs_runtime.Value
functorExceptT1_1_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_2))
_ = __local_var_3_3
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, v_4)
})
}))
_ = functorExceptT1_1_1
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_15
// TAST (Let): __local_var_3_14 -> gopurs_runtime.Value
__local_var_3_14 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_15, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_14
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): functorExceptT1_4_5 -> gopurs_runtime.Value
functorExceptT1_4_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_6, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_5))
_ = __local_var_6_7
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_7, v_7)
})
}))
_ = functorExceptT1_4_5
// TAST (Let): __local_var_5_8 -> gopurs_runtime.Value
__local_var_5_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_9 -> *Constructor_Control_Bind_Bind
Bind1_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_9
// TAST (Let): pure_7_10 -> gopurs_runtime.Value
pure_7_10 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t11 = gopurs_runtime.Apply(pure_7_10, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_11
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t11 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
}))
})
}))
}))
_ = __local_var_5_8
// TAST (Let): Bind1_6_12 -> *Constructor_Control_Bind_Bind
Bind1_6_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_12
// TAST (Let): Applicative0_7_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_4_5
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_12.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_12.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_13.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_14, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_16 -> *Constructor_Control_Bind_Bind
Bind1_3_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_16
// TAST (Let): pure_4_17 -> gopurs_runtime.Value
pure_4_17 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_17
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_16.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_4_17, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_18
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
})
}))
}))
_ = __local_var_2_4
// TAST (Let): Bind1_3_19 -> *Constructor_Control_Bind_Bind
Bind1_3_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_19
// TAST (Let): Applicative0_4_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_20
// TAST (Let): applyExceptT1_1_0 -> *Constructor_Control_Apply_Apply
applyExceptT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_1_1
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_19.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_19.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_20.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
})))
_ = applyExceptT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_21 -> *Constructor_Data_Functor_Functor
Functor0_3_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyExceptT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_21
// TAST (Let): __local_var_4_22 -> gopurs_runtime.Value
__local_var_4_22 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_22
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyExceptT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_21.V0), __local_var_4_22, a_5), b_6)
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
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_21
// TAST (Let): __local_var_3_20 -> gopurs_runtime.Value
__local_var_3_20 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_21, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_20
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 -> gopurs_runtime.Value
functorExceptT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_5))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
}))
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_6 -> *Constructor_Control_Bind_Bind
Bind1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_6
// TAST (Let): pure_7_7 -> gopurs_runtime.Value
pure_7_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_7
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_9
// TAST (Let): functorExceptT1_9_8 -> gopurs_runtime.Value
functorExceptT1_9_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_10))
_ = __local_var_11_10
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_10, v_12)
})
}))
_ = functorExceptT1_9_8
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_12 -> *Constructor_Control_Bind_Bind
Bind1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_12
// TAST (Let): pure_12_13 -> gopurs_runtime.Value
pure_12_13 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_13
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_12.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t14 = gopurs_runtime.Apply(pure_12_13, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_14
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t14 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}))
})
}))
}))
_ = __local_var_10_11
// TAST (Let): Bind1_11_15 -> *Constructor_Control_Bind_Bind
Bind1_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_15
// TAST (Let): Applicative0_12_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_9_8
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_16.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_6.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(pure_7_7, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_17
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}))
})
}))
}))
_ = __local_var_5_5
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_4_2
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_20, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_22 -> *Constructor_Control_Bind_Bind
Bind1_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_22
// TAST (Let): pure_4_23 -> gopurs_runtime.Value
pure_4_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_25 -> gopurs_runtime.Value
__local_var_6_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_25
// TAST (Let): functorExceptT1_6_24 -> gopurs_runtime.Value
functorExceptT1_6_24 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_26 -> gopurs_runtime.Value
__local_var_8_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_25, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_7))
_ = __local_var_8_26
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_26, v_9)
})
}))
_ = functorExceptT1_6_24
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_38 -> gopurs_runtime.Value
__local_var_8_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_38
// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
__local_var_8_37 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_38, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_37
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_29
// TAST (Let): functorExceptT1_9_28 -> gopurs_runtime.Value
functorExceptT1_9_28 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_30 -> gopurs_runtime.Value
__local_var_11_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_29, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_10))
_ = __local_var_11_30
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_30, v_12)
})
}))
_ = functorExceptT1_9_28
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_32 -> *Constructor_Control_Bind_Bind
Bind1_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_32
// TAST (Let): pure_12_33 -> gopurs_runtime.Value
pure_12_33 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_33
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_32.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t34 = gopurs_runtime.Apply(pure_12_33, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_34
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t34 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_34
} else {

}
}
{
__t34 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_34:
return __t34
}))
})
}))
}))
_ = __local_var_10_31
// TAST (Let): Bind1_11_35 -> *Constructor_Control_Bind_Bind
Bind1_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_35
// TAST (Let): Applicative0_12_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_36
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_9_28
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_36.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_37, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_39 -> *Constructor_Control_Bind_Bind
Bind1_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_39
// TAST (Let): pure_9_40 -> gopurs_runtime.Value
pure_9_40 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_40
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_39.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t41 = gopurs_runtime.Apply(pure_9_40, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_41
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t41 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
goto end_branch_41
} else {

}
}
{
__t41 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_41:
return __t41
}))
})
}))
}))
_ = __local_var_7_27
// TAST (Let): Bind1_8_42 -> *Constructor_Control_Bind_Bind
Bind1_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_42
// TAST (Let): Applicative0_9_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_43
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_6_24
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_43.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_22.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t44 = gopurs_runtime.Apply(pure_4_23, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_44
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t44 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_44
} else {

}
}
{
__t44 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_44:
return __t44
}))
})
}))
}))
_ = monadExceptT1_1_0
// TAST (Let): __local_var_2_45 -> *Constructor_Control_Monad_Monad
__local_var_2_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_2_45
// TAST (Let): pure_3_46 -> gopurs_runtime.Value
pure_3_46 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_45.V0), gopurs_runtime.Value{}), "pure")
_ = pure_3_46
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_1_0
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_45.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_0, "ask"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_46, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_4})})
})))
}

func Call_Control_Monad_Except_Trans_monadReaderExceptT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): monadExceptT1_2_2 -> gopurs_runtime.Value
monadExceptT1_2_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_83 -> gopurs_runtime.Value
__local_var_4_83 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_83
// TAST (Let): __local_var_4_82 -> gopurs_runtime.Value
__local_var_4_82 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_83, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_82
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorExceptT1_5_4 -> gopurs_runtime.Value
functorExceptT1_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_6))
_ = __local_var_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, v_8)
})
}))
_ = functorExceptT1_5_4
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_27
// TAST (Let): __local_var_7_26 -> gopurs_runtime.Value
__local_var_7_26 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_27, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_8})})
})
_ = __local_var_7_26
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_9
// TAST (Let): functorExceptT1_8_8 -> gopurs_runtime.Value
functorExceptT1_8_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_9))
_ = __local_var_10_10
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_10, v_11)
})
}))
_ = functorExceptT1_8_8
// TAST (Let): __local_var_9_11 -> gopurs_runtime.Value
__local_var_9_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_12 -> *Constructor_Control_Bind_Bind
Bind1_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_12
// TAST (Let): pure_11_13 -> gopurs_runtime.Value
pure_11_13 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_13
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_15 -> gopurs_runtime.Value
__local_var_13_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_15
// TAST (Let): functorExceptT1_13_14 -> gopurs_runtime.Value
functorExceptT1_13_14 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_16 -> gopurs_runtime.Value
__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_15, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_14))
_ = __local_var_15_16
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_16, v_16)
})
}))
_ = functorExceptT1_13_14
// TAST (Let): __local_var_14_17 -> gopurs_runtime.Value
__local_var_14_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_18 -> *Constructor_Control_Bind_Bind
Bind1_15_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_18
// TAST (Let): pure_16_19 -> gopurs_runtime.Value
pure_16_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_19
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_18.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(pure_16_19, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_20
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t20 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
}))
})
}))
}))
_ = __local_var_14_17
// TAST (Let): Bind1_15_21 -> *Constructor_Control_Bind_Bind
Bind1_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_21
// TAST (Let): Applicative0_16_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_22
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_13_14
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_21.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_21.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_22.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_12.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t23 = gopurs_runtime.Apply(pure_11_13, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_23
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t23 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
}))
})
}))
}))
_ = __local_var_9_11
// TAST (Let): Bind1_10_24 -> *Constructor_Control_Bind_Bind
Bind1_10_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_24
// TAST (Let): Applicative0_11_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_8_8
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_24.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_24.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_25.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_26, x_8)
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_28 -> *Constructor_Control_Bind_Bind
Bind1_7_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_28
// TAST (Let): pure_8_29 -> gopurs_runtime.Value
pure_8_29 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_29
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_31
// TAST (Let): functorExceptT1_10_30 -> gopurs_runtime.Value
functorExceptT1_10_30 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_32 -> gopurs_runtime.Value
__local_var_12_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_32
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_32, v_13)
})
}))
_ = functorExceptT1_10_30
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_53 -> gopurs_runtime.Value
__local_var_12_53 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_53
// TAST (Let): __local_var_12_52 -> gopurs_runtime.Value
__local_var_12_52 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_53, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_52
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_35 -> gopurs_runtime.Value
__local_var_13_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_35
// TAST (Let): functorExceptT1_13_34 -> gopurs_runtime.Value
functorExceptT1_13_34 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_36 -> gopurs_runtime.Value
__local_var_15_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_35, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_14))
_ = __local_var_15_36
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_36, v_16)
})
}))
_ = functorExceptT1_13_34
// TAST (Let): __local_var_14_37 -> gopurs_runtime.Value
__local_var_14_37 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_38 -> *Constructor_Control_Bind_Bind
Bind1_15_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_38
// TAST (Let): pure_16_39 -> gopurs_runtime.Value
pure_16_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_39
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_41 -> gopurs_runtime.Value
__local_var_18_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_41
// TAST (Let): functorExceptT1_18_40 -> gopurs_runtime.Value
functorExceptT1_18_40 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_42 -> gopurs_runtime.Value
__local_var_20_42 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_41, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_19))
_ = __local_var_20_42
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_42, v_21)
})
}))
_ = functorExceptT1_18_40
// TAST (Let): __local_var_19_43 -> gopurs_runtime.Value
__local_var_19_43 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_44 -> *Constructor_Control_Bind_Bind
Bind1_20_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_44
// TAST (Let): pure_21_45 -> gopurs_runtime.Value
pure_21_45 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_45
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_44.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t46 = gopurs_runtime.Apply(pure_21_45, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_46
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t46 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_46
} else {

}
}
{
__t46 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_46:
return __t46
}))
})
}))
}))
_ = __local_var_19_43
// TAST (Let): Bind1_20_47 -> *Constructor_Control_Bind_Bind
Bind1_20_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_43, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_47
// TAST (Let): Applicative0_21_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_43, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_48
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_18_40
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_47.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_47.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_48.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_38.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t49 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t49 = gopurs_runtime.Apply(pure_16_39, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_49
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t49 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_49
} else {

}
}
{
__t49 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_49:
return __t49
}))
})
}))
}))
_ = __local_var_14_37
// TAST (Let): Bind1_15_50 -> *Constructor_Control_Bind_Bind
Bind1_15_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_50
// TAST (Let): Applicative0_16_51 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_51
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_13_34
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_50.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_50.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_51.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_52, x_13)
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_54 -> *Constructor_Control_Bind_Bind
Bind1_12_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_54
// TAST (Let): pure_13_55 -> gopurs_runtime.Value
pure_13_55 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_55
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_57 -> gopurs_runtime.Value
__local_var_15_57 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_57
// TAST (Let): functorExceptT1_15_56 -> gopurs_runtime.Value
functorExceptT1_15_56 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_58 -> gopurs_runtime.Value
__local_var_17_58 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_57, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_16))
_ = __local_var_17_58
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_58, v_18)
})
}))
_ = functorExceptT1_15_56
// TAST (Let): __local_var_16_59 -> gopurs_runtime.Value
__local_var_16_59 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_70 -> gopurs_runtime.Value
__local_var_17_70 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_70
// TAST (Let): __local_var_17_69 -> gopurs_runtime.Value
__local_var_17_69 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_70, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_18})})
})
_ = __local_var_17_69
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_61 -> gopurs_runtime.Value
__local_var_18_61 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_61
// TAST (Let): functorExceptT1_18_60 -> gopurs_runtime.Value
functorExceptT1_18_60 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_62 -> gopurs_runtime.Value
__local_var_20_62 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_61, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_19))
_ = __local_var_20_62
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_62, v_21)
})
}))
_ = functorExceptT1_18_60
// TAST (Let): __local_var_19_63 -> gopurs_runtime.Value
__local_var_19_63 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_64 -> *Constructor_Control_Bind_Bind
Bind1_20_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_64
// TAST (Let): pure_21_65 -> gopurs_runtime.Value
pure_21_65 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_65
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_64.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t66 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t66 = gopurs_runtime.Apply(pure_21_65, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_66
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t66 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_66
} else {

}
}
{
__t66 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_66:
return __t66
}))
})
}))
}))
_ = __local_var_19_63
// TAST (Let): Bind1_20_67 -> *Constructor_Control_Bind_Bind
Bind1_20_67 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_63, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_67
// TAST (Let): Applicative0_21_68 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_68 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_63, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_68
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_18_60
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_67.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_67.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_68.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_69, x_18)
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_71 -> *Constructor_Control_Bind_Bind
Bind1_17_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_71
// TAST (Let): pure_18_72 -> gopurs_runtime.Value
pure_18_72 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_72
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_71.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t73 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t73 = gopurs_runtime.Apply(pure_18_72, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_73
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t73 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_73
} else {

}
}
{
__t73 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_73:
return __t73
}))
})
}))
}))
_ = __local_var_16_59
// TAST (Let): Bind1_17_74 -> *Constructor_Control_Bind_Bind
Bind1_17_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_59, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_74
// TAST (Let): Applicative0_18_75 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_59, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_75
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_15_56
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_74.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_74.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_75.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_54.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t76 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t76 = gopurs_runtime.Apply(pure_13_55, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_76
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t76 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_76
} else {

}
}
{
__t76 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_76:
return __t76
}))
})
}))
}))
_ = __local_var_11_33
// TAST (Let): Bind1_12_77 -> *Constructor_Control_Bind_Bind
Bind1_12_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_77
// TAST (Let): Applicative0_13_78 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_78
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_30
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_77.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_77.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_78.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_28.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t79 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t79 = gopurs_runtime.Apply(pure_8_29, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_79
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t79 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_79
} else {

}
}
{
__t79 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_79:
return __t79
}))
})
}))
}))
_ = __local_var_6_7
// TAST (Let): Bind1_7_80 -> *Constructor_Control_Bind_Bind
Bind1_7_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_80
// TAST (Let): Applicative0_8_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_81
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_5_4
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_80.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_80.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_81.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_82, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_84 -> *Constructor_Control_Bind_Bind
Bind1_4_84 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_84
// TAST (Let): pure_5_85 -> gopurs_runtime.Value
pure_5_85 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_85
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_87 -> gopurs_runtime.Value
__local_var_7_87 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_87
// TAST (Let): functorExceptT1_7_86 -> gopurs_runtime.Value
functorExceptT1_7_86 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_88 -> gopurs_runtime.Value
__local_var_9_88 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_87, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_8))
_ = __local_var_9_88
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_88, v_10)
})
}))
_ = functorExceptT1_7_86
// TAST (Let): __local_var_8_89 -> gopurs_runtime.Value
__local_var_8_89 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_140 -> gopurs_runtime.Value
__local_var_9_140 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_140
// TAST (Let): __local_var_9_139 -> gopurs_runtime.Value
__local_var_9_139 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_140, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_139
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_91 -> gopurs_runtime.Value
__local_var_10_91 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_91
// TAST (Let): functorExceptT1_10_90 -> gopurs_runtime.Value
functorExceptT1_10_90 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_92 -> gopurs_runtime.Value
__local_var_12_92 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_91, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_92
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_92, v_13)
})
}))
_ = functorExceptT1_10_90
// TAST (Let): __local_var_11_93 -> gopurs_runtime.Value
__local_var_11_93 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_113 -> gopurs_runtime.Value
__local_var_12_113 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_113
// TAST (Let): __local_var_12_112 -> gopurs_runtime.Value
__local_var_12_112 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_113, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_112
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_95 -> gopurs_runtime.Value
__local_var_13_95 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_95
// TAST (Let): functorExceptT1_13_94 -> gopurs_runtime.Value
functorExceptT1_13_94 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_96 -> gopurs_runtime.Value
__local_var_15_96 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_95, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_14))
_ = __local_var_15_96
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_96, v_16)
})
}))
_ = functorExceptT1_13_94
// TAST (Let): __local_var_14_97 -> gopurs_runtime.Value
__local_var_14_97 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_98 -> *Constructor_Control_Bind_Bind
Bind1_15_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_98
// TAST (Let): pure_16_99 -> gopurs_runtime.Value
pure_16_99 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_99
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_101 -> gopurs_runtime.Value
__local_var_18_101 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_101
// TAST (Let): functorExceptT1_18_100 -> gopurs_runtime.Value
functorExceptT1_18_100 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_102 -> gopurs_runtime.Value
__local_var_20_102 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_101, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_19))
_ = __local_var_20_102
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_102, v_21)
})
}))
_ = functorExceptT1_18_100
// TAST (Let): __local_var_19_103 -> gopurs_runtime.Value
__local_var_19_103 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_104 -> *Constructor_Control_Bind_Bind
Bind1_20_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_104
// TAST (Let): pure_21_105 -> gopurs_runtime.Value
pure_21_105 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_105
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_104.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t106 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t106 = gopurs_runtime.Apply(pure_21_105, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_106
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t106 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_106
} else {

}
}
{
__t106 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_106:
return __t106
}))
})
}))
}))
_ = __local_var_19_103
// TAST (Let): Bind1_20_107 -> *Constructor_Control_Bind_Bind
Bind1_20_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_103, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_107
// TAST (Let): Applicative0_21_108 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_103, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_108
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_18_100
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_107.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_107.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_108.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_98.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t109 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t109 = gopurs_runtime.Apply(pure_16_99, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_109
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t109 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_109
} else {

}
}
{
__t109 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_109:
return __t109
}))
})
}))
}))
_ = __local_var_14_97
// TAST (Let): Bind1_15_110 -> *Constructor_Control_Bind_Bind
Bind1_15_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_97, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_110
// TAST (Let): Applicative0_16_111 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_97, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_111
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_13_94
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_110.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_110.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_111.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_112, x_13)
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_114 -> *Constructor_Control_Bind_Bind
Bind1_12_114 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_114
// TAST (Let): pure_13_115 -> gopurs_runtime.Value
pure_13_115 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_115
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_117 -> gopurs_runtime.Value
__local_var_15_117 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_117
// TAST (Let): functorExceptT1_15_116 -> gopurs_runtime.Value
functorExceptT1_15_116 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_118 -> gopurs_runtime.Value
__local_var_17_118 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_117, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_16))
_ = __local_var_17_118
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_118, v_18)
})
}))
_ = functorExceptT1_15_116
// TAST (Let): __local_var_16_119 -> gopurs_runtime.Value
__local_var_16_119 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_130 -> gopurs_runtime.Value
__local_var_17_130 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_130
// TAST (Let): __local_var_17_129 -> gopurs_runtime.Value
__local_var_17_129 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_130, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_18})})
})
_ = __local_var_17_129
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_121 -> gopurs_runtime.Value
__local_var_18_121 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_121
// TAST (Let): functorExceptT1_18_120 -> gopurs_runtime.Value
functorExceptT1_18_120 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_122 -> gopurs_runtime.Value
__local_var_20_122 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_121, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_19))
_ = __local_var_20_122
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_122, v_21)
})
}))
_ = functorExceptT1_18_120
// TAST (Let): __local_var_19_123 -> gopurs_runtime.Value
__local_var_19_123 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_124 -> *Constructor_Control_Bind_Bind
Bind1_20_124 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_124
// TAST (Let): pure_21_125 -> gopurs_runtime.Value
pure_21_125 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_125
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_124.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t126 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t126 = gopurs_runtime.Apply(pure_21_125, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_126
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t126 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_126
} else {

}
}
{
__t126 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_126:
return __t126
}))
})
}))
}))
_ = __local_var_19_123
// TAST (Let): Bind1_20_127 -> *Constructor_Control_Bind_Bind
Bind1_20_127 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_123, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_127
// TAST (Let): Applicative0_21_128 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_128 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_123, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_128
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_18_120
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_127.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_127.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_128.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_129, x_18)
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_131 -> *Constructor_Control_Bind_Bind
Bind1_17_131 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_131
// TAST (Let): pure_18_132 -> gopurs_runtime.Value
pure_18_132 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_132
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_131.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t133 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t133 = gopurs_runtime.Apply(pure_18_132, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_133
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t133 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_133
} else {

}
}
{
__t133 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_133:
return __t133
}))
})
}))
}))
_ = __local_var_16_119
// TAST (Let): Bind1_17_134 -> *Constructor_Control_Bind_Bind
Bind1_17_134 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_119, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_134
// TAST (Let): Applicative0_18_135 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_135 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_119, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_135
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_15_116
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_134.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_134.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_135.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_114.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t136 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t136 = gopurs_runtime.Apply(pure_13_115, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_136
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t136 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_136
} else {

}
}
{
__t136 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_136:
return __t136
}))
})
}))
}))
_ = __local_var_11_93
// TAST (Let): Bind1_12_137 -> *Constructor_Control_Bind_Bind
Bind1_12_137 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_93, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_137
// TAST (Let): Applicative0_13_138 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_138 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_93, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_138
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_90
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_137.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_137.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_138.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_139, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_141 -> *Constructor_Control_Bind_Bind
Bind1_9_141 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_141
// TAST (Let): pure_10_142 -> gopurs_runtime.Value
pure_10_142 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_142
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_144 -> gopurs_runtime.Value
__local_var_12_144 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_144
// TAST (Let): functorExceptT1_12_143 -> gopurs_runtime.Value
functorExceptT1_12_143 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_145 -> gopurs_runtime.Value
__local_var_14_145 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_144, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_13))
_ = __local_var_14_145
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_145, v_15)
})
}))
_ = functorExceptT1_12_143
// TAST (Let): __local_var_13_146 -> gopurs_runtime.Value
__local_var_13_146 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_157 -> gopurs_runtime.Value
__local_var_14_157 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_14_157
// TAST (Let): __local_var_14_156 -> gopurs_runtime.Value
__local_var_14_156 := gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_157, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_15})})
})
_ = __local_var_14_156
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_148 -> gopurs_runtime.Value
__local_var_15_148 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_148
// TAST (Let): functorExceptT1_15_147 -> gopurs_runtime.Value
functorExceptT1_15_147 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_149 -> gopurs_runtime.Value
__local_var_17_149 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_148, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_16))
_ = __local_var_17_149
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_149, v_18)
})
}))
_ = functorExceptT1_15_147
// TAST (Let): __local_var_16_150 -> gopurs_runtime.Value
__local_var_16_150 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_151 -> *Constructor_Control_Bind_Bind
Bind1_17_151 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_151
// TAST (Let): pure_18_152 -> gopurs_runtime.Value
pure_18_152 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_152
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_151.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t153 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t153 = gopurs_runtime.Apply(pure_18_152, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_153
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t153 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_153
} else {

}
}
{
__t153 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_153:
return __t153
}))
})
}))
}))
_ = __local_var_16_150
// TAST (Let): Bind1_17_154 -> *Constructor_Control_Bind_Bind
Bind1_17_154 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_150, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_154
// TAST (Let): Applicative0_18_155 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_155 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_150, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_155
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_15_147
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_154.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_154.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_155.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_156, x_15)
}))
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_14_158 -> *Constructor_Control_Bind_Bind
Bind1_14_158 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_158
// TAST (Let): pure_15_159 -> gopurs_runtime.Value
pure_15_159 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_15_159
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_158.V1), v_16, gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t160 gopurs_runtime.Value
{
if (v2_18.Type == 9 && v2_18.IntVal == 3711209382) {
__t160 = gopurs_runtime.Apply(pure_15_159, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_18.UnsafePtr).V0})})
goto end_branch_160
} else {

}
}
{
if (v2_18.Type == 9 && v2_18.IntVal == 2465973597) {
__t160 = gopurs_runtime.Apply(k_17, (*Constructor_Data_Either_Right)(v2_18.UnsafePtr).V0)
goto end_branch_160
} else {

}
}
{
__t160 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_160:
return __t160
}))
})
}))
}))
_ = __local_var_13_146
// TAST (Let): Bind1_14_161 -> *Constructor_Control_Bind_Bind
Bind1_14_161 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_146, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_161
// TAST (Let): Applicative0_15_162 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_162 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_146, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_162
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_12_143
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_161.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_161.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_162.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_141.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t163 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t163 = gopurs_runtime.Apply(pure_10_142, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_163
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t163 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
goto end_branch_163
} else {

}
}
{
__t163 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_163:
return __t163
}))
})
}))
}))
_ = __local_var_8_89
// TAST (Let): Bind1_9_164 -> *Constructor_Control_Bind_Bind
Bind1_9_164 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_89, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_164
// TAST (Let): Applicative0_10_165 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_165 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_89, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_165
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_7_86
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_164.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_164.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_165.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_84.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t166 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t166 = gopurs_runtime.Apply(pure_5_85, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_166
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t166 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
goto end_branch_166
} else {

}
}
{
__t166 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_166:
return __t166
}))
})
}))
}))
_ = monadExceptT1_2_2
// TAST (Let): __local_var_3_167 -> *Constructor_Control_Monad_Monad
__local_var_3_167 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_167
// TAST (Let): pure_4_168 -> gopurs_runtime.Value
pure_4_168 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_167.V0), gopurs_runtime.Value{}), "pure")
_ = pure_4_168
// TAST (Let): monadAskExceptT1_1_0 -> gopurs_runtime.Value
monadAskExceptT1_1_0 := gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_167.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(__local_var_1_1, "ask"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_168, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_5})})
})))
_ = monadAskExceptT1_1_0
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskExceptT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_169 -> gopurs_runtime.Value
__local_var_3_169 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_2)
_ = __local_var_3_169
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_169, v_4)
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
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_21
// TAST (Let): __local_var_3_20 -> gopurs_runtime.Value
__local_var_3_20 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_21, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_20
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 -> gopurs_runtime.Value
functorExceptT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_5))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
}))
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_6 -> *Constructor_Control_Bind_Bind
Bind1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_6
// TAST (Let): pure_7_7 -> gopurs_runtime.Value
pure_7_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_7
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_9
// TAST (Let): functorExceptT1_9_8 -> gopurs_runtime.Value
functorExceptT1_9_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_10))
_ = __local_var_11_10
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_10, v_12)
})
}))
_ = functorExceptT1_9_8
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_12 -> *Constructor_Control_Bind_Bind
Bind1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_12
// TAST (Let): pure_12_13 -> gopurs_runtime.Value
pure_12_13 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_13
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_12.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t14 = gopurs_runtime.Apply(pure_12_13, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_14
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t14 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}))
})
}))
}))
_ = __local_var_10_11
// TAST (Let): Bind1_11_15 -> *Constructor_Control_Bind_Bind
Bind1_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_15
// TAST (Let): Applicative0_12_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_9_8
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_16.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_6.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(pure_7_7, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_17
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}))
})
}))
}))
_ = __local_var_5_5
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_4_2
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_20, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_22 -> *Constructor_Control_Bind_Bind
Bind1_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_22
// TAST (Let): pure_4_23 -> gopurs_runtime.Value
pure_4_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_25 -> gopurs_runtime.Value
__local_var_6_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_25
// TAST (Let): functorExceptT1_6_24 -> gopurs_runtime.Value
functorExceptT1_6_24 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_26 -> gopurs_runtime.Value
__local_var_8_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_25, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_7))
_ = __local_var_8_26
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_26, v_9)
})
}))
_ = functorExceptT1_6_24
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_38 -> gopurs_runtime.Value
__local_var_8_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_38
// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
__local_var_8_37 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_38, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_37
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_29
// TAST (Let): functorExceptT1_9_28 -> gopurs_runtime.Value
functorExceptT1_9_28 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_30 -> gopurs_runtime.Value
__local_var_11_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_29, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_10))
_ = __local_var_11_30
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_30, v_12)
})
}))
_ = functorExceptT1_9_28
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_32 -> *Constructor_Control_Bind_Bind
Bind1_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_32
// TAST (Let): pure_12_33 -> gopurs_runtime.Value
pure_12_33 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_33
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_32.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t34 = gopurs_runtime.Apply(pure_12_33, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_34
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t34 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_34
} else {

}
}
{
__t34 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_34:
return __t34
}))
})
}))
}))
_ = __local_var_10_31
// TAST (Let): Bind1_11_35 -> *Constructor_Control_Bind_Bind
Bind1_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_35
// TAST (Let): Applicative0_12_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_36
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_9_28
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_36.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_37, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_39 -> *Constructor_Control_Bind_Bind
Bind1_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_39
// TAST (Let): pure_9_40 -> gopurs_runtime.Value
pure_9_40 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_40
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_39.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t41 = gopurs_runtime.Apply(pure_9_40, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_41
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t41 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
goto end_branch_41
} else {

}
}
{
__t41 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_41:
return __t41
}))
})
}))
}))
_ = __local_var_7_27
// TAST (Let): Bind1_8_42 -> *Constructor_Control_Bind_Bind
Bind1_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_42
// TAST (Let): Applicative0_9_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_43
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_6_24
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_43.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_22.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t44 = gopurs_runtime.Apply(pure_4_23, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_44
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t44 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_44
} else {

}
}
{
__t44 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_44:
return __t44
}))
})
}))
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
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_21
// TAST (Let): __local_var_3_20 -> gopurs_runtime.Value
__local_var_3_20 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_21, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_20
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 -> gopurs_runtime.Value
functorExceptT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_5))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
}))
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_6 -> *Constructor_Control_Bind_Bind
Bind1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_6
// TAST (Let): pure_7_7 -> gopurs_runtime.Value
pure_7_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_7
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_9
// TAST (Let): functorExceptT1_9_8 -> gopurs_runtime.Value
functorExceptT1_9_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_10))
_ = __local_var_11_10
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_10, v_12)
})
}))
_ = functorExceptT1_9_8
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_12 -> *Constructor_Control_Bind_Bind
Bind1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_12
// TAST (Let): pure_12_13 -> gopurs_runtime.Value
pure_12_13 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_13
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_12.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t14 = gopurs_runtime.Apply(pure_12_13, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_14
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t14 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}))
})
}))
}))
_ = __local_var_10_11
// TAST (Let): Bind1_11_15 -> *Constructor_Control_Bind_Bind
Bind1_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_15
// TAST (Let): Applicative0_12_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_9_8
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_16.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_6.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(pure_7_7, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_17
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}))
})
}))
}))
_ = __local_var_5_5
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_4_2
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_20, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_22 -> *Constructor_Control_Bind_Bind
Bind1_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_22
// TAST (Let): pure_4_23 -> gopurs_runtime.Value
pure_4_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_25 -> gopurs_runtime.Value
__local_var_6_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_25
// TAST (Let): functorExceptT1_6_24 -> gopurs_runtime.Value
functorExceptT1_6_24 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_26 -> gopurs_runtime.Value
__local_var_8_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_25, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_7))
_ = __local_var_8_26
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_26, v_9)
})
}))
_ = functorExceptT1_6_24
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_38 -> gopurs_runtime.Value
__local_var_8_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_38
// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
__local_var_8_37 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_38, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_37
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_29
// TAST (Let): functorExceptT1_9_28 -> gopurs_runtime.Value
functorExceptT1_9_28 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_30 -> gopurs_runtime.Value
__local_var_11_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_29, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_10))
_ = __local_var_11_30
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_30, v_12)
})
}))
_ = functorExceptT1_9_28
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_32 -> *Constructor_Control_Bind_Bind
Bind1_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_32
// TAST (Let): pure_12_33 -> gopurs_runtime.Value
pure_12_33 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_33
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_32.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t34 = gopurs_runtime.Apply(pure_12_33, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_34
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t34 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_34
} else {

}
}
{
__t34 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_34:
return __t34
}))
})
}))
}))
_ = __local_var_10_31
// TAST (Let): Bind1_11_35 -> *Constructor_Control_Bind_Bind
Bind1_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_35
// TAST (Let): Applicative0_12_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_36
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_9_28
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_36.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_37, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_39 -> *Constructor_Control_Bind_Bind
Bind1_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_39
// TAST (Let): pure_9_40 -> gopurs_runtime.Value
pure_9_40 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_40
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_39.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t41 = gopurs_runtime.Apply(pure_9_40, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_41
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t41 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
goto end_branch_41
} else {

}
}
{
__t41 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_41:
return __t41
}))
})
}))
}))
_ = __local_var_7_27
// TAST (Let): Bind1_8_42 -> *Constructor_Control_Bind_Bind
Bind1_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_42
// TAST (Let): Applicative0_9_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_43
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_6_24
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_43.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_22.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t44 = gopurs_runtime.Apply(pure_4_23, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_44
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t44 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_44
} else {

}
}
{
__t44 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_44:
return __t44
}))
})
}))
}))
_ = monadExceptT1_2_1
// TAST (Let): Bind1_3_46 -> *Constructor_Control_Bind_Bind
Bind1_3_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_46
// TAST (Let): pure_4_47 -> gopurs_runtime.Value
pure_4_47 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_47
// TAST (Let): __local_var_3_45 -> gopurs_runtime.Value
__local_var_3_45 := gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_46.V1), m_5, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_47, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_6})})
}))
})
_ = __local_var_3_45
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_45, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
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
// TAST (Let): __local_var_5_23 -> gopurs_runtime.Value
__local_var_5_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_23
// TAST (Let): __local_var_5_22 -> gopurs_runtime.Value
__local_var_5_22 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_23, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_6})})
})
_ = __local_var_5_22
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): functorExceptT1_6_4 -> gopurs_runtime.Value
functorExceptT1_6_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_7))
_ = __local_var_8_6
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_6, v_9)
})
}))
_ = functorExceptT1_6_4
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_8 -> *Constructor_Control_Bind_Bind
Bind1_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_8
// TAST (Let): pure_9_9 -> gopurs_runtime.Value
pure_9_9 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_9
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorExceptT1_11_10 -> gopurs_runtime.Value
functorExceptT1_11_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_12 -> gopurs_runtime.Value
__local_var_13_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_12))
_ = __local_var_13_12
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_12, v_14)
})
}))
_ = functorExceptT1_11_10
// TAST (Let): __local_var_12_13 -> gopurs_runtime.Value
__local_var_12_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_14 -> *Constructor_Control_Bind_Bind
Bind1_13_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_14
// TAST (Let): pure_14_15 -> gopurs_runtime.Value
pure_14_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_15
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_14.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(pure_14_15, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_16
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t16 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}))
})
}))
}))
_ = __local_var_12_13
// TAST (Let): Bind1_13_17 -> *Constructor_Control_Bind_Bind
Bind1_13_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_17
// TAST (Let): Applicative0_14_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_11_10
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_17.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_17.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_18.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_8.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(pure_9_9, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_19
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t19 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
}))
})
}))
}))
_ = __local_var_7_7
// TAST (Let): Bind1_8_20 -> *Constructor_Control_Bind_Bind
Bind1_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_20
// TAST (Let): Applicative0_9_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_21
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_6_4
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_20.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_20.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_21.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_22, x_6)
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_24 -> *Constructor_Control_Bind_Bind
Bind1_5_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_24
// TAST (Let): pure_6_25 -> gopurs_runtime.Value
pure_6_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_25
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_27 -> gopurs_runtime.Value
__local_var_8_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_27
// TAST (Let): functorExceptT1_8_26 -> gopurs_runtime.Value
functorExceptT1_8_26 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_28 -> gopurs_runtime.Value
__local_var_10_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_27, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_9))
_ = __local_var_10_28
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_28, v_11)
})
}))
_ = functorExceptT1_8_26
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_40 -> gopurs_runtime.Value
__local_var_10_40 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_40
// TAST (Let): __local_var_10_39 -> gopurs_runtime.Value
__local_var_10_39 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_40, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_11})})
})
_ = __local_var_10_39
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_31 -> gopurs_runtime.Value
__local_var_11_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_31
// TAST (Let): functorExceptT1_11_30 -> gopurs_runtime.Value
functorExceptT1_11_30 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_32 -> gopurs_runtime.Value
__local_var_13_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_31, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_12))
_ = __local_var_13_32
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_32, v_14)
})
}))
_ = functorExceptT1_11_30
// TAST (Let): __local_var_12_33 -> gopurs_runtime.Value
__local_var_12_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_34 -> *Constructor_Control_Bind_Bind
Bind1_13_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_34
// TAST (Let): pure_14_35 -> gopurs_runtime.Value
pure_14_35 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_35
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_34.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t36 = gopurs_runtime.Apply(pure_14_35, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_36
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t36 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
goto end_branch_36
} else {

}
}
{
__t36 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_36:
return __t36
}))
})
}))
}))
_ = __local_var_12_33
// TAST (Let): Bind1_13_37 -> *Constructor_Control_Bind_Bind
Bind1_13_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_37
// TAST (Let): Applicative0_14_38 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_38
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_11_30
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_37.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_37.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_38.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_39, x_11)
}))
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_41 -> *Constructor_Control_Bind_Bind
Bind1_10_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_41
// TAST (Let): pure_11_42 -> gopurs_runtime.Value
pure_11_42 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_42
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_41.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t43 = gopurs_runtime.Apply(pure_11_42, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_43
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t43 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_43
} else {

}
}
{
__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_43:
return __t43
}))
})
}))
}))
_ = __local_var_9_29
// TAST (Let): Bind1_10_44 -> *Constructor_Control_Bind_Bind
Bind1_10_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_44
// TAST (Let): Applicative0_11_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_45
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_8_26
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_44.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_44.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_45.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_24.V1), v_7, gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 3711209382) {
__t46 = gopurs_runtime.Apply(pure_6_25, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_9.UnsafePtr).V0})})
goto end_branch_46
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 2465973597) {
__t46 = gopurs_runtime.Apply(k_8, (*Constructor_Data_Either_Right)(v2_9.UnsafePtr).V0)
goto end_branch_46
} else {

}
}
{
__t46 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_46:
return __t46
}))
})
}))
}))
_ = monadExceptT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_4_3
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_47 -> gopurs_runtime.Value
__local_var_6_47 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(f_5, a_6), gopurs_runtime.Func(func(m_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t51 *Constructor_Control_Monad_Rec_Class_Done
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 3711209382) {
__t51 = &Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_prime_7.UnsafePtr).V0})}}
goto end_branch_51
} else {

}
}
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 2465973597) {
var __t50 gopurs_runtime.Value
{
var __t_tag_48 gopurs_runtime.Value = (*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0
if (__t_tag_48.Type == 9 && __t_tag_48.IntVal == 525585346) {
__t50 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}
goto end_branch_50
} else {

}
}
{
var __t_tag_49 gopurs_runtime.Value = (*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0
if (__t_tag_49.Type == 9 && __t_tag_49.IntVal == 60402430) {
__t50 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_50
} else {

}
}
{
__t50 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_50:
__t51 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](__t50)
goto end_branch_51
} else {

}
}
{
__t51 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_51:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(__t51)})
}))
}))
_ = __local_var_6_47
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_47, x_7)
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
// TAST (Let): __local_var_4_22 -> gopurs_runtime.Value
__local_var_4_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_22
// TAST (Let): __local_var_4_21 -> gopurs_runtime.Value
__local_var_4_21 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_22, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_21
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorExceptT1_5_3 -> gopurs_runtime.Value
functorExceptT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_6))
_ = __local_var_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, v_8)
})
}))
_ = functorExceptT1_5_3
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_7 -> *Constructor_Control_Bind_Bind
Bind1_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_7
// TAST (Let): pure_8_8 -> gopurs_runtime.Value
pure_8_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_8
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_10
// TAST (Let): functorExceptT1_10_9 -> gopurs_runtime.Value
functorExceptT1_10_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_11 -> gopurs_runtime.Value
__local_var_12_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_10, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_11
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_11, v_13)
})
}))
_ = functorExceptT1_10_9
// TAST (Let): __local_var_11_12 -> gopurs_runtime.Value
__local_var_11_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_13 -> *Constructor_Control_Bind_Bind
Bind1_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_13
// TAST (Let): pure_13_14 -> gopurs_runtime.Value
pure_13_14 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_14
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_13.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(pure_13_14, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_15
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t15 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}))
})
}))
}))
_ = __local_var_11_12
// TAST (Let): Bind1_12_16 -> *Constructor_Control_Bind_Bind
Bind1_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_16
// TAST (Let): Applicative0_13_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_9
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_17.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_7.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_8_8, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_18
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
})
}))
}))
_ = __local_var_6_6
// TAST (Let): Bind1_7_19 -> *Constructor_Control_Bind_Bind
Bind1_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_19
// TAST (Let): Applicative0_8_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_20
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_20.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_21, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_23 -> *Constructor_Control_Bind_Bind
Bind1_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_23
// TAST (Let): pure_5_24 -> gopurs_runtime.Value
pure_5_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_24
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_26 -> gopurs_runtime.Value
__local_var_7_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_26
// TAST (Let): functorExceptT1_7_25 -> gopurs_runtime.Value
functorExceptT1_7_25 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_27 -> gopurs_runtime.Value
__local_var_9_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_26, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_8))
_ = __local_var_9_27
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_27, v_10)
})
}))
_ = functorExceptT1_7_25
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_39 -> gopurs_runtime.Value
__local_var_9_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_39
// TAST (Let): __local_var_9_38 -> gopurs_runtime.Value
__local_var_9_38 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_39, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_38
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_30 -> gopurs_runtime.Value
__local_var_10_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_30
// TAST (Let): functorExceptT1_10_29 -> gopurs_runtime.Value
functorExceptT1_10_29 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_31 -> gopurs_runtime.Value
__local_var_12_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_30, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_31
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_31, v_13)
})
}))
_ = functorExceptT1_10_29
// TAST (Let): __local_var_11_32 -> gopurs_runtime.Value
__local_var_11_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_33 -> *Constructor_Control_Bind_Bind
Bind1_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_33
// TAST (Let): pure_13_34 -> gopurs_runtime.Value
pure_13_34 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_34
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_33.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t35 = gopurs_runtime.Apply(pure_13_34, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_35
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t35 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_35
} else {

}
}
{
__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_35:
return __t35
}))
})
}))
}))
_ = __local_var_11_32
// TAST (Let): Bind1_12_36 -> *Constructor_Control_Bind_Bind
Bind1_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_36
// TAST (Let): Applicative0_13_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_37
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_29
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_37.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_38, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_40 -> *Constructor_Control_Bind_Bind
Bind1_9_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_40
// TAST (Let): pure_10_41 -> gopurs_runtime.Value
pure_10_41 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_41
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_40.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t42 = gopurs_runtime.Apply(pure_10_41, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_42
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t42 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
goto end_branch_42
} else {

}
}
{
__t42 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_42:
return __t42
}))
})
}))
}))
_ = __local_var_8_28
// TAST (Let): Bind1_9_43 -> *Constructor_Control_Bind_Bind
Bind1_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_43
// TAST (Let): Applicative0_10_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_44
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_7_25
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_44.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_23.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t45 = gopurs_runtime.Apply(pure_5_24, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_45
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t45 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
goto end_branch_45
} else {

}
}
{
__t45 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_45:
return __t45
}))
})
}))
}))
_ = monadExceptT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_4_46 -> gopurs_runtime.Value
pure_4_46 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_4_46
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), f_3), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_46, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_5})})
}))
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
// TAST (Let): __local_var_4_22 -> gopurs_runtime.Value
__local_var_4_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_22
// TAST (Let): __local_var_4_21 -> gopurs_runtime.Value
__local_var_4_21 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_22, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_21
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorExceptT1_5_3 -> gopurs_runtime.Value
functorExceptT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_6))
_ = __local_var_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, v_8)
})
}))
_ = functorExceptT1_5_3
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_7 -> *Constructor_Control_Bind_Bind
Bind1_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_7
// TAST (Let): pure_8_8 -> gopurs_runtime.Value
pure_8_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_8
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_10
// TAST (Let): functorExceptT1_10_9 -> gopurs_runtime.Value
functorExceptT1_10_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_11 -> gopurs_runtime.Value
__local_var_12_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_10, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_11
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_11, v_13)
})
}))
_ = functorExceptT1_10_9
// TAST (Let): __local_var_11_12 -> gopurs_runtime.Value
__local_var_11_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_13 -> *Constructor_Control_Bind_Bind
Bind1_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_13
// TAST (Let): pure_13_14 -> gopurs_runtime.Value
pure_13_14 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_14
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_1_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_13.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(pure_13_14, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_15
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t15 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}))
})
}))
}))
_ = __local_var_11_12
// TAST (Let): Bind1_12_16 -> *Constructor_Control_Bind_Bind
Bind1_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_16
// TAST (Let): Applicative0_13_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_9
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_17.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_7.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_8_8, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_18
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
})
}))
}))
_ = __local_var_6_6
// TAST (Let): Bind1_7_19 -> *Constructor_Control_Bind_Bind
Bind1_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_19
// TAST (Let): Applicative0_8_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_20
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_20.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_21, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_23 -> *Constructor_Control_Bind_Bind
Bind1_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_23
// TAST (Let): pure_5_24 -> gopurs_runtime.Value
pure_5_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_24
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_26 -> gopurs_runtime.Value
__local_var_7_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_26
// TAST (Let): functorExceptT1_7_25 -> gopurs_runtime.Value
functorExceptT1_7_25 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_27 -> gopurs_runtime.Value
__local_var_9_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_26, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_8))
_ = __local_var_9_27
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_27, v_10)
})
}))
_ = functorExceptT1_7_25
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_39 -> gopurs_runtime.Value
__local_var_9_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_39
// TAST (Let): __local_var_9_38 -> gopurs_runtime.Value
__local_var_9_38 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_39, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_38
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_30 -> gopurs_runtime.Value
__local_var_10_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_30
// TAST (Let): functorExceptT1_10_29 -> gopurs_runtime.Value
functorExceptT1_10_29 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_31 -> gopurs_runtime.Value
__local_var_12_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_30, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_31
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_31, v_13)
})
}))
_ = functorExceptT1_10_29
// TAST (Let): __local_var_11_32 -> gopurs_runtime.Value
__local_var_11_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_33 -> *Constructor_Control_Bind_Bind
Bind1_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_33
// TAST (Let): pure_13_34 -> gopurs_runtime.Value
pure_13_34 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_34
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_1_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_33.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t35 = gopurs_runtime.Apply(pure_13_34, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_35
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t35 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_35
} else {

}
}
{
__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_35:
return __t35
}))
})
}))
}))
_ = __local_var_11_32
// TAST (Let): Bind1_12_36 -> *Constructor_Control_Bind_Bind
Bind1_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_36
// TAST (Let): Applicative0_13_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_37
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_29
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_37.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_38, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_40 -> *Constructor_Control_Bind_Bind
Bind1_9_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_40
// TAST (Let): pure_10_41 -> gopurs_runtime.Value
pure_10_41 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_41
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_1_0)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_40.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t42 = gopurs_runtime.Apply(pure_10_41, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_42
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t42 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
goto end_branch_42
} else {

}
}
{
__t42 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_42:
return __t42
}))
})
}))
}))
_ = __local_var_8_28
// TAST (Let): Bind1_9_43 -> *Constructor_Control_Bind_Bind
Bind1_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_43
// TAST (Let): Applicative0_10_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_44
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_7_25
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_44.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_23.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t45 = gopurs_runtime.Apply(pure_5_24, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_45
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t45 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
goto end_branch_45
} else {

}
}
{
__t45 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_45:
return __t45
}))
})
}))
}))
_ = monadExceptT1_3_2
// TAST (Let): Bind1_4_47 -> *Constructor_Control_Bind_Bind
Bind1_4_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_47
// TAST (Let): pure_5_48 -> gopurs_runtime.Value
pure_5_48 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_48
// TAST (Let): __local_var_4_46 -> gopurs_runtime.Value
__local_var_4_46 := gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_47.V1), m_6, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_48, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_7})})
}))
})
_ = __local_var_4_46
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_46, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
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
// TAST (Let): Monad1_7_7 -> gopurs_runtime.Value
Monad1_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_7_7
// TAST (Let): Semigroup0_8_8 -> gopurs_runtime.Value
Semigroup0_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_8_8
// TAST (Let): monadExceptT1_9_9 -> gopurs_runtime.Value
monadExceptT1_9_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_89 -> gopurs_runtime.Value
__local_var_10_89 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_89
// TAST (Let): __local_var_10_88 -> gopurs_runtime.Value
__local_var_10_88 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_89, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_11})})
})
_ = __local_var_10_88
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorExceptT1_11_10 -> gopurs_runtime.Value
functorExceptT1_11_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_12 -> gopurs_runtime.Value
__local_var_13_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_12))
_ = __local_var_13_12
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_12, v_14)
})
}))
_ = functorExceptT1_11_10
// TAST (Let): __local_var_12_13 -> gopurs_runtime.Value
__local_var_12_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_33 -> gopurs_runtime.Value
__local_var_13_33 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_33
// TAST (Let): __local_var_13_32 -> gopurs_runtime.Value
__local_var_13_32 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_33, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_14})})
})
_ = __local_var_13_32
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_15 -> gopurs_runtime.Value
__local_var_14_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_15
// TAST (Let): functorExceptT1_14_14 -> gopurs_runtime.Value
functorExceptT1_14_14 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_16 -> gopurs_runtime.Value
__local_var_16_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_15, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_15))
_ = __local_var_16_16
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_16, v_17)
})
}))
_ = functorExceptT1_14_14
// TAST (Let): __local_var_15_17 -> gopurs_runtime.Value
__local_var_15_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_18 -> *Constructor_Control_Bind_Bind
Bind1_16_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_18
// TAST (Let): pure_17_19 -> gopurs_runtime.Value
pure_17_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_19
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_21 -> gopurs_runtime.Value
__local_var_19_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_21
// TAST (Let): functorExceptT1_19_20 -> gopurs_runtime.Value
functorExceptT1_19_20 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_22 -> gopurs_runtime.Value
__local_var_21_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_21, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_20))
_ = __local_var_21_22
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_22, v_22)
})
}))
_ = functorExceptT1_19_20
// TAST (Let): __local_var_20_23 -> gopurs_runtime.Value
__local_var_20_23 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_24 -> *Constructor_Control_Bind_Bind
Bind1_21_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_24
// TAST (Let): pure_22_25 -> gopurs_runtime.Value
pure_22_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_25
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_24.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t26 = gopurs_runtime.Apply(pure_22_25, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_26
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t26 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
return __t26
}))
})
}))
}))
_ = __local_var_20_23
// TAST (Let): Bind1_21_27 -> *Constructor_Control_Bind_Bind
Bind1_21_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_23, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_27
// TAST (Let): Applicative0_22_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_23, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_28
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_19_20
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_27.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_27.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_28.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_18.V1), v_18, gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (v2_20.Type == 9 && v2_20.IntVal == 3711209382) {
__t29 = gopurs_runtime.Apply(pure_17_19, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_20.UnsafePtr).V0})})
goto end_branch_29
} else {

}
}
{
if (v2_20.Type == 9 && v2_20.IntVal == 2465973597) {
__t29 = gopurs_runtime.Apply(k_19, (*Constructor_Data_Either_Right)(v2_20.UnsafePtr).V0)
goto end_branch_29
} else {

}
}
{
__t29 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_29:
return __t29
}))
})
}))
}))
_ = __local_var_15_17
// TAST (Let): Bind1_16_30 -> *Constructor_Control_Bind_Bind
Bind1_16_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_30
// TAST (Let): Applicative0_17_31 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_31
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_14_14
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_30.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_30.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_31.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_32, x_14)
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_34 -> *Constructor_Control_Bind_Bind
Bind1_13_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_34
// TAST (Let): pure_14_35 -> gopurs_runtime.Value
pure_14_35 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_35
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_37 -> gopurs_runtime.Value
__local_var_16_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_37
// TAST (Let): functorExceptT1_16_36 -> gopurs_runtime.Value
functorExceptT1_16_36 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_38 -> gopurs_runtime.Value
__local_var_18_38 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_37, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_17))
_ = __local_var_18_38
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_38, v_19)
})
}))
_ = functorExceptT1_16_36
// TAST (Let): __local_var_17_39 -> gopurs_runtime.Value
__local_var_17_39 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_59 -> gopurs_runtime.Value
__local_var_18_59 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_59
// TAST (Let): __local_var_18_58 -> gopurs_runtime.Value
__local_var_18_58 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_59, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_19})})
})
_ = __local_var_18_58
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_41 -> gopurs_runtime.Value
__local_var_19_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_41
// TAST (Let): functorExceptT1_19_40 -> gopurs_runtime.Value
functorExceptT1_19_40 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_42 -> gopurs_runtime.Value
__local_var_21_42 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_41, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_20))
_ = __local_var_21_42
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_42, v_22)
})
}))
_ = functorExceptT1_19_40
// TAST (Let): __local_var_20_43 -> gopurs_runtime.Value
__local_var_20_43 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_44 -> *Constructor_Control_Bind_Bind
Bind1_21_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_44
// TAST (Let): pure_22_45 -> gopurs_runtime.Value
pure_22_45 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_45
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_47 -> gopurs_runtime.Value
__local_var_24_47 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_47
// TAST (Let): functorExceptT1_24_46 -> gopurs_runtime.Value
functorExceptT1_24_46 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_26_48 -> gopurs_runtime.Value
__local_var_26_48 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_47, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_25))
_ = __local_var_26_48
return gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_26_48, v_27)
})
}))
_ = functorExceptT1_24_46
// TAST (Let): __local_var_25_49 -> gopurs_runtime.Value
__local_var_25_49 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_50 -> *Constructor_Control_Bind_Bind
Bind1_26_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_50
// TAST (Let): pure_27_51 -> gopurs_runtime.Value
pure_27_51 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_27_51
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_50.V1), v_28, gopurs_runtime.Func(func(v2_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t52 gopurs_runtime.Value
{
if (v2_30.Type == 9 && v2_30.IntVal == 3711209382) {
__t52 = gopurs_runtime.Apply(pure_27_51, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_30.UnsafePtr).V0})})
goto end_branch_52
} else {

}
}
{
if (v2_30.Type == 9 && v2_30.IntVal == 2465973597) {
__t52 = gopurs_runtime.Apply(k_29, (*Constructor_Data_Either_Right)(v2_30.UnsafePtr).V0)
goto end_branch_52
} else {

}
}
{
__t52 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_52:
return __t52
}))
})
}))
}))
_ = __local_var_25_49
// TAST (Let): Bind1_26_53 -> *Constructor_Control_Bind_Bind
Bind1_26_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_49, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_53
// TAST (Let): Applicative0_27_54 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_49, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_54
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_24_46
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_53.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_53.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_54.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_44.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t55 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t55 = gopurs_runtime.Apply(pure_22_45, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_55
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t55 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
goto end_branch_55
} else {

}
}
{
__t55 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_55:
return __t55
}))
})
}))
}))
_ = __local_var_20_43
// TAST (Let): Bind1_21_56 -> *Constructor_Control_Bind_Bind
Bind1_21_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_43, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_56
// TAST (Let): Applicative0_22_57 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_43, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_57
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_19_40
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_56.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_56.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_57.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_58, x_19)
}))
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_60 -> *Constructor_Control_Bind_Bind
Bind1_18_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_60
// TAST (Let): pure_19_61 -> gopurs_runtime.Value
pure_19_61 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_19_61
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_63 -> gopurs_runtime.Value
__local_var_21_63 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_63
// TAST (Let): functorExceptT1_21_62 -> gopurs_runtime.Value
functorExceptT1_21_62 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_64 -> gopurs_runtime.Value
__local_var_23_64 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_21_63, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_22))
_ = __local_var_23_64
return gopurs_runtime.Func(func(v_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_64, v_24)
})
}))
_ = functorExceptT1_21_62
// TAST (Let): __local_var_22_65 -> gopurs_runtime.Value
__local_var_22_65 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_76 -> gopurs_runtime.Value
__local_var_23_76 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_23_76
// TAST (Let): __local_var_23_75 -> gopurs_runtime.Value
__local_var_23_75 := gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_76, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_24})})
})
_ = __local_var_23_75
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_67 -> gopurs_runtime.Value
__local_var_24_67 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_67
// TAST (Let): functorExceptT1_24_66 -> gopurs_runtime.Value
functorExceptT1_24_66 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_26_68 -> gopurs_runtime.Value
__local_var_26_68 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_67, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_25))
_ = __local_var_26_68
return gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_26_68, v_27)
})
}))
_ = functorExceptT1_24_66
// TAST (Let): __local_var_25_69 -> gopurs_runtime.Value
__local_var_25_69 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_70 -> *Constructor_Control_Bind_Bind
Bind1_26_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_70
// TAST (Let): pure_27_71 -> gopurs_runtime.Value
pure_27_71 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_27_71
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_70.V1), v_28, gopurs_runtime.Func(func(v2_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t72 gopurs_runtime.Value
{
if (v2_30.Type == 9 && v2_30.IntVal == 3711209382) {
__t72 = gopurs_runtime.Apply(pure_27_71, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_30.UnsafePtr).V0})})
goto end_branch_72
} else {

}
}
{
if (v2_30.Type == 9 && v2_30.IntVal == 2465973597) {
__t72 = gopurs_runtime.Apply(k_29, (*Constructor_Data_Either_Right)(v2_30.UnsafePtr).V0)
goto end_branch_72
} else {

}
}
{
__t72 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_72:
return __t72
}))
})
}))
}))
_ = __local_var_25_69
// TAST (Let): Bind1_26_73 -> *Constructor_Control_Bind_Bind
Bind1_26_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_69, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_73
// TAST (Let): Applicative0_27_74 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_69, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_74
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_24_66
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_73.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_73.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_74.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_75, x_24)
}))
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_77 -> *Constructor_Control_Bind_Bind
Bind1_23_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_77
// TAST (Let): pure_24_78 -> gopurs_runtime.Value
pure_24_78 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_24_78
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_77.V1), v_25, gopurs_runtime.Func(func(v2_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t79 gopurs_runtime.Value
{
if (v2_27.Type == 9 && v2_27.IntVal == 3711209382) {
__t79 = gopurs_runtime.Apply(pure_24_78, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_27.UnsafePtr).V0})})
goto end_branch_79
} else {

}
}
{
if (v2_27.Type == 9 && v2_27.IntVal == 2465973597) {
__t79 = gopurs_runtime.Apply(k_26, (*Constructor_Data_Either_Right)(v2_27.UnsafePtr).V0)
goto end_branch_79
} else {

}
}
{
__t79 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_79:
return __t79
}))
})
}))
}))
_ = __local_var_22_65
// TAST (Let): Bind1_23_80 -> *Constructor_Control_Bind_Bind
Bind1_23_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_65, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_80
// TAST (Let): Applicative0_24_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_65, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_81
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_21_62
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_80.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_80.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_81.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_60.V1), v_20, gopurs_runtime.Func(func(v2_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t82 gopurs_runtime.Value
{
if (v2_22.Type == 9 && v2_22.IntVal == 3711209382) {
__t82 = gopurs_runtime.Apply(pure_19_61, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_22.UnsafePtr).V0})})
goto end_branch_82
} else {

}
}
{
if (v2_22.Type == 9 && v2_22.IntVal == 2465973597) {
__t82 = gopurs_runtime.Apply(k_21, (*Constructor_Data_Either_Right)(v2_22.UnsafePtr).V0)
goto end_branch_82
} else {

}
}
{
__t82 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_82:
return __t82
}))
})
}))
}))
_ = __local_var_17_39
// TAST (Let): Bind1_18_83 -> *Constructor_Control_Bind_Bind
Bind1_18_83 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_39, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_83
// TAST (Let): Applicative0_19_84 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_84 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_39, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_84
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_16_36
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_83.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_83.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_84.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_34.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t85 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t85 = gopurs_runtime.Apply(pure_14_35, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_85
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t85 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
goto end_branch_85
} else {

}
}
{
__t85 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_85:
return __t85
}))
})
}))
}))
_ = __local_var_12_13
// TAST (Let): Bind1_13_86 -> *Constructor_Control_Bind_Bind
Bind1_13_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_86
// TAST (Let): Applicative0_14_87 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_87
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_11_10
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_86.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_86.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_87.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_88, x_11)
}))
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_90 -> *Constructor_Control_Bind_Bind
Bind1_10_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_90
// TAST (Let): pure_11_91 -> gopurs_runtime.Value
pure_11_91 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_91
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_93 -> gopurs_runtime.Value
__local_var_13_93 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_93
// TAST (Let): functorExceptT1_13_92 -> gopurs_runtime.Value
functorExceptT1_13_92 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_94 -> gopurs_runtime.Value
__local_var_15_94 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_93, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_14))
_ = __local_var_15_94
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_94, v_16)
})
}))
_ = functorExceptT1_13_92
// TAST (Let): __local_var_14_95 -> gopurs_runtime.Value
__local_var_14_95 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_146 -> gopurs_runtime.Value
__local_var_15_146 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_15_146
// TAST (Let): __local_var_15_145 -> gopurs_runtime.Value
__local_var_15_145 := gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_146, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_16})})
})
_ = __local_var_15_145
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_97 -> gopurs_runtime.Value
__local_var_16_97 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_97
// TAST (Let): functorExceptT1_16_96 -> gopurs_runtime.Value
functorExceptT1_16_96 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_98 -> gopurs_runtime.Value
__local_var_18_98 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_97, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_17))
_ = __local_var_18_98
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_98, v_19)
})
}))
_ = functorExceptT1_16_96
// TAST (Let): __local_var_17_99 -> gopurs_runtime.Value
__local_var_17_99 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_119 -> gopurs_runtime.Value
__local_var_18_119 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_119
// TAST (Let): __local_var_18_118 -> gopurs_runtime.Value
__local_var_18_118 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_119, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_19})})
})
_ = __local_var_18_118
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_101 -> gopurs_runtime.Value
__local_var_19_101 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_101
// TAST (Let): functorExceptT1_19_100 -> gopurs_runtime.Value
functorExceptT1_19_100 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_102 -> gopurs_runtime.Value
__local_var_21_102 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_101, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_20))
_ = __local_var_21_102
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_102, v_22)
})
}))
_ = functorExceptT1_19_100
// TAST (Let): __local_var_20_103 -> gopurs_runtime.Value
__local_var_20_103 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_104 -> *Constructor_Control_Bind_Bind
Bind1_21_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_104
// TAST (Let): pure_22_105 -> gopurs_runtime.Value
pure_22_105 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_105
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_107 -> gopurs_runtime.Value
__local_var_24_107 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_107
// TAST (Let): functorExceptT1_24_106 -> gopurs_runtime.Value
functorExceptT1_24_106 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_26_108 -> gopurs_runtime.Value
__local_var_26_108 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_107, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_25))
_ = __local_var_26_108
return gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_26_108, v_27)
})
}))
_ = functorExceptT1_24_106
// TAST (Let): __local_var_25_109 -> gopurs_runtime.Value
__local_var_25_109 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_110 -> *Constructor_Control_Bind_Bind
Bind1_26_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_110
// TAST (Let): pure_27_111 -> gopurs_runtime.Value
pure_27_111 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_27_111
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_110.V1), v_28, gopurs_runtime.Func(func(v2_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t112 gopurs_runtime.Value
{
if (v2_30.Type == 9 && v2_30.IntVal == 3711209382) {
__t112 = gopurs_runtime.Apply(pure_27_111, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_30.UnsafePtr).V0})})
goto end_branch_112
} else {

}
}
{
if (v2_30.Type == 9 && v2_30.IntVal == 2465973597) {
__t112 = gopurs_runtime.Apply(k_29, (*Constructor_Data_Either_Right)(v2_30.UnsafePtr).V0)
goto end_branch_112
} else {

}
}
{
__t112 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_112:
return __t112
}))
})
}))
}))
_ = __local_var_25_109
// TAST (Let): Bind1_26_113 -> *Constructor_Control_Bind_Bind
Bind1_26_113 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_109, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_113
// TAST (Let): Applicative0_27_114 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_114 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_109, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_114
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_24_106
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_113.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_113.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_114.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_104.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t115 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t115 = gopurs_runtime.Apply(pure_22_105, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_115
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t115 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
goto end_branch_115
} else {

}
}
{
__t115 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_115:
return __t115
}))
})
}))
}))
_ = __local_var_20_103
// TAST (Let): Bind1_21_116 -> *Constructor_Control_Bind_Bind
Bind1_21_116 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_103, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_116
// TAST (Let): Applicative0_22_117 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_117 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_103, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_117
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_19_100
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_116.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_116.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_117.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_118, x_19)
}))
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_120 -> *Constructor_Control_Bind_Bind
Bind1_18_120 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_120
// TAST (Let): pure_19_121 -> gopurs_runtime.Value
pure_19_121 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_19_121
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_123 -> gopurs_runtime.Value
__local_var_21_123 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_123
// TAST (Let): functorExceptT1_21_122 -> gopurs_runtime.Value
functorExceptT1_21_122 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_124 -> gopurs_runtime.Value
__local_var_23_124 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_21_123, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_22))
_ = __local_var_23_124
return gopurs_runtime.Func(func(v_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_124, v_24)
})
}))
_ = functorExceptT1_21_122
// TAST (Let): __local_var_22_125 -> gopurs_runtime.Value
__local_var_22_125 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_136 -> gopurs_runtime.Value
__local_var_23_136 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_23_136
// TAST (Let): __local_var_23_135 -> gopurs_runtime.Value
__local_var_23_135 := gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_136, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_24})})
})
_ = __local_var_23_135
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_127 -> gopurs_runtime.Value
__local_var_24_127 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_127
// TAST (Let): functorExceptT1_24_126 -> gopurs_runtime.Value
functorExceptT1_24_126 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_26_128 -> gopurs_runtime.Value
__local_var_26_128 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_127, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_25))
_ = __local_var_26_128
return gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_26_128, v_27)
})
}))
_ = functorExceptT1_24_126
// TAST (Let): __local_var_25_129 -> gopurs_runtime.Value
__local_var_25_129 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_130 -> *Constructor_Control_Bind_Bind
Bind1_26_130 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_130
// TAST (Let): pure_27_131 -> gopurs_runtime.Value
pure_27_131 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_27_131
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_130.V1), v_28, gopurs_runtime.Func(func(v2_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t132 gopurs_runtime.Value
{
if (v2_30.Type == 9 && v2_30.IntVal == 3711209382) {
__t132 = gopurs_runtime.Apply(pure_27_131, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_30.UnsafePtr).V0})})
goto end_branch_132
} else {

}
}
{
if (v2_30.Type == 9 && v2_30.IntVal == 2465973597) {
__t132 = gopurs_runtime.Apply(k_29, (*Constructor_Data_Either_Right)(v2_30.UnsafePtr).V0)
goto end_branch_132
} else {

}
}
{
__t132 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_132:
return __t132
}))
})
}))
}))
_ = __local_var_25_129
// TAST (Let): Bind1_26_133 -> *Constructor_Control_Bind_Bind
Bind1_26_133 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_129, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_133
// TAST (Let): Applicative0_27_134 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_134 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_129, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_134
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_24_126
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_133.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_133.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_134.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_135, x_24)
}))
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_137 -> *Constructor_Control_Bind_Bind
Bind1_23_137 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_137
// TAST (Let): pure_24_138 -> gopurs_runtime.Value
pure_24_138 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_24_138
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_137.V1), v_25, gopurs_runtime.Func(func(v2_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t139 gopurs_runtime.Value
{
if (v2_27.Type == 9 && v2_27.IntVal == 3711209382) {
__t139 = gopurs_runtime.Apply(pure_24_138, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_27.UnsafePtr).V0})})
goto end_branch_139
} else {

}
}
{
if (v2_27.Type == 9 && v2_27.IntVal == 2465973597) {
__t139 = gopurs_runtime.Apply(k_26, (*Constructor_Data_Either_Right)(v2_27.UnsafePtr).V0)
goto end_branch_139
} else {

}
}
{
__t139 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_139:
return __t139
}))
})
}))
}))
_ = __local_var_22_125
// TAST (Let): Bind1_23_140 -> *Constructor_Control_Bind_Bind
Bind1_23_140 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_125, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_140
// TAST (Let): Applicative0_24_141 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_141 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_125, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_141
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_21_122
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_140.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_140.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_141.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_120.V1), v_20, gopurs_runtime.Func(func(v2_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t142 gopurs_runtime.Value
{
if (v2_22.Type == 9 && v2_22.IntVal == 3711209382) {
__t142 = gopurs_runtime.Apply(pure_19_121, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_22.UnsafePtr).V0})})
goto end_branch_142
} else {

}
}
{
if (v2_22.Type == 9 && v2_22.IntVal == 2465973597) {
__t142 = gopurs_runtime.Apply(k_21, (*Constructor_Data_Either_Right)(v2_22.UnsafePtr).V0)
goto end_branch_142
} else {

}
}
{
__t142 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_142:
return __t142
}))
})
}))
}))
_ = __local_var_17_99
// TAST (Let): Bind1_18_143 -> *Constructor_Control_Bind_Bind
Bind1_18_143 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_99, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_143
// TAST (Let): Applicative0_19_144 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_144 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_99, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_144
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_16_96
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_143.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_143.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_144.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_145, x_16)
}))
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_147 -> *Constructor_Control_Bind_Bind
Bind1_15_147 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_147
// TAST (Let): pure_16_148 -> gopurs_runtime.Value
pure_16_148 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_148
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_150 -> gopurs_runtime.Value
__local_var_18_150 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_150
// TAST (Let): functorExceptT1_18_149 -> gopurs_runtime.Value
functorExceptT1_18_149 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_151 -> gopurs_runtime.Value
__local_var_20_151 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_150, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_19))
_ = __local_var_20_151
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_151, v_21)
})
}))
_ = functorExceptT1_18_149
// TAST (Let): __local_var_19_152 -> gopurs_runtime.Value
__local_var_19_152 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_163 -> gopurs_runtime.Value
__local_var_20_163 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_20_163
// TAST (Let): __local_var_20_162 -> gopurs_runtime.Value
__local_var_20_162 := gopurs_runtime.Func(func(x_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_163, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_21})})
})
_ = __local_var_20_162
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_154 -> gopurs_runtime.Value
__local_var_21_154 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_154
// TAST (Let): functorExceptT1_21_153 -> gopurs_runtime.Value
functorExceptT1_21_153 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_155 -> gopurs_runtime.Value
__local_var_23_155 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_21_154, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_22))
_ = __local_var_23_155
return gopurs_runtime.Func(func(v_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_155, v_24)
})
}))
_ = functorExceptT1_21_153
// TAST (Let): __local_var_22_156 -> gopurs_runtime.Value
__local_var_22_156 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_157 -> *Constructor_Control_Bind_Bind
Bind1_23_157 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_157
// TAST (Let): pure_24_158 -> gopurs_runtime.Value
pure_24_158 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_24_158
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_157.V1), v_25, gopurs_runtime.Func(func(v2_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t159 gopurs_runtime.Value
{
if (v2_27.Type == 9 && v2_27.IntVal == 3711209382) {
__t159 = gopurs_runtime.Apply(pure_24_158, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_27.UnsafePtr).V0})})
goto end_branch_159
} else {

}
}
{
if (v2_27.Type == 9 && v2_27.IntVal == 2465973597) {
__t159 = gopurs_runtime.Apply(k_26, (*Constructor_Data_Either_Right)(v2_27.UnsafePtr).V0)
goto end_branch_159
} else {

}
}
{
__t159 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_159:
return __t159
}))
})
}))
}))
_ = __local_var_22_156
// TAST (Let): Bind1_23_160 -> *Constructor_Control_Bind_Bind
Bind1_23_160 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_156, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_160
// TAST (Let): Applicative0_24_161 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_161 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_156, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_161
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_21_153
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_160.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_160.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_161.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_162, x_21)
}))
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_164 -> *Constructor_Control_Bind_Bind
Bind1_20_164 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_164
// TAST (Let): pure_21_165 -> gopurs_runtime.Value
pure_21_165 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_165
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_164.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t166 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t166 = gopurs_runtime.Apply(pure_21_165, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_166
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t166 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_166
} else {

}
}
{
__t166 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_166:
return __t166
}))
})
}))
}))
_ = __local_var_19_152
// TAST (Let): Bind1_20_167 -> *Constructor_Control_Bind_Bind
Bind1_20_167 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_152, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_167
// TAST (Let): Applicative0_21_168 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_168 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_152, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_168
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_18_149
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_167.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_167.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_168.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_147.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t169 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t169 = gopurs_runtime.Apply(pure_16_148, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_169
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t169 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_169
} else {

}
}
{
__t169 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_169:
return __t169
}))
})
}))
}))
_ = __local_var_14_95
// TAST (Let): Bind1_15_170 -> *Constructor_Control_Bind_Bind
Bind1_15_170 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_95, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_170
// TAST (Let): Applicative0_16_171 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_171 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_95, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_171
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_13_92
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_170.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_170.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_171.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_90.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t172 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t172 = gopurs_runtime.Apply(pure_11_91, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_172
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t172 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_172
} else {

}
}
{
__t172 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_172:
return __t172
}))
})
}))
}))
_ = monadExceptT1_9_9
// TAST (Let): Bind1_10_174 -> *Constructor_Control_Bind_Bind
Bind1_10_174 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_174
// TAST (Let): pure_11_175 -> gopurs_runtime.Value
pure_11_175 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_175
// TAST (Let): __local_var_10_173 -> gopurs_runtime.Value
__local_var_10_173 := gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_174.V1), m_12, gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_11_175, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_13})})
}))
})
_ = __local_var_10_173
// TAST (Let): monadTellExceptT1_7_6 -> gopurs_runtime.Value
monadTellExceptT1_7_6 := gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_9_9
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_8_8
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_173, gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "tell"), x_11))
}))
_ = monadTellExceptT1_7_6
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellExceptT1_7_6
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_6_5
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), v_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_176 -> gopurs_runtime.Value
__local_var_10_176 := (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1
_ = __local_var_10_176
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, r_11, __local_var_10_176})}
}), (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0))
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_8, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t177 *Constructor_Data_Tuple_Tuple
{
if (a_9.Type == 9 && a_9.IntVal == 3711209382) {
__t177 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(a_9.UnsafePtr).V0})}, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_10
})}
goto end_branch_177
} else {

}
}
{
if (a_9.Type == 9 && a_9.IntVal == 2465973597) {
__t177 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Either_Right)(a_9.UnsafePtr).V0.UnsafePtr).V0})}, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Either_Right)(a_9.UnsafePtr).V0.UnsafePtr).V1}
goto end_branch_177
} else {

}
}
{
__t177 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_177:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t177)})
})))
}))
}

func Call_Control_Monad_Except_Trans_monadThrowExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadExceptT1_1_0 -> gopurs_runtime.Value
monadExceptT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_20 -> gopurs_runtime.Value
__local_var_2_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_2_20
// TAST (Let): __local_var_2_19 -> gopurs_runtime.Value
__local_var_2_19 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_20, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_3})})
})
_ = __local_var_2_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorExceptT1_3_1 -> gopurs_runtime.Value
functorExceptT1_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_4))
_ = __local_var_5_3
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, v_6)
})
}))
_ = functorExceptT1_3_1
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_5 -> *Constructor_Control_Bind_Bind
Bind1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_5
// TAST (Let): pure_6_6 -> gopurs_runtime.Value
pure_6_6 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_8
// TAST (Let): functorExceptT1_8_7 -> gopurs_runtime.Value
functorExceptT1_8_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_9))
_ = __local_var_10_9
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_9, v_11)
})
}))
_ = functorExceptT1_8_7
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_11 -> *Constructor_Control_Bind_Bind
Bind1_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_11
// TAST (Let): pure_11_12 -> gopurs_runtime.Value
pure_11_12 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_12
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_11.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(pure_11_12, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_13
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t13 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}))
})
}))
}))
_ = __local_var_9_10
// TAST (Let): Bind1_10_14 -> *Constructor_Control_Bind_Bind
Bind1_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_14
// TAST (Let): Applicative0_11_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_15
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_8_7
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_15.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_5.V1), v_7, gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(pure_6_6, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_9.UnsafePtr).V0})})
goto end_branch_16
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 2465973597) {
__t16 = gopurs_runtime.Apply(k_8, (*Constructor_Data_Either_Right)(v2_9.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}))
})
}))
}))
_ = __local_var_4_4
// TAST (Let): Bind1_5_17 -> *Constructor_Control_Bind_Bind
Bind1_5_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_17
// TAST (Let): Applicative0_6_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_3_1
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_17.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_17.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_18.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_19, x_3)
}))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_21 -> *Constructor_Control_Bind_Bind
Bind1_2_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_21
// TAST (Let): pure_3_22 -> gopurs_runtime.Value
pure_3_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_22
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_24 -> gopurs_runtime.Value
__local_var_5_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_24
// TAST (Let): functorExceptT1_5_23 -> gopurs_runtime.Value
functorExceptT1_5_23 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_25 -> gopurs_runtime.Value
__local_var_7_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_24, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_6))
_ = __local_var_7_25
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_25, v_8)
})
}))
_ = functorExceptT1_5_23
// TAST (Let): __local_var_6_26 -> gopurs_runtime.Value
__local_var_6_26 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_37 -> gopurs_runtime.Value
__local_var_7_37 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_37
// TAST (Let): __local_var_7_36 -> gopurs_runtime.Value
__local_var_7_36 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_37, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_8})})
})
_ = __local_var_7_36
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_28
// TAST (Let): functorExceptT1_8_27 -> gopurs_runtime.Value
functorExceptT1_8_27 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_29 -> gopurs_runtime.Value
__local_var_10_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_9))
_ = __local_var_10_29
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_29, v_11)
})
}))
_ = functorExceptT1_8_27
// TAST (Let): __local_var_9_30 -> gopurs_runtime.Value
__local_var_9_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_31 -> *Constructor_Control_Bind_Bind
Bind1_10_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_31
// TAST (Let): pure_11_32 -> gopurs_runtime.Value
pure_11_32 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_32
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_31.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t33 = gopurs_runtime.Apply(pure_11_32, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_33
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t33 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_33
} else {

}
}
{
__t33 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_33:
return __t33
}))
})
}))
}))
_ = __local_var_9_30
// TAST (Let): Bind1_10_34 -> *Constructor_Control_Bind_Bind
Bind1_10_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_34
// TAST (Let): Applicative0_11_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_35
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_8_27
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_34.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_34.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_35.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_36, x_8)
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_38 -> *Constructor_Control_Bind_Bind
Bind1_7_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_38
// TAST (Let): pure_8_39 -> gopurs_runtime.Value
pure_8_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_39
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_38.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t40 = gopurs_runtime.Apply(pure_8_39, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_40
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t40 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_40
} else {

}
}
{
__t40 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_40:
return __t40
}))
})
}))
}))
_ = __local_var_6_26
// TAST (Let): Bind1_7_41 -> *Constructor_Control_Bind_Bind
Bind1_7_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_26, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_41
// TAST (Let): Applicative0_8_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_26, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_42
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_5_23
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_41.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_41.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_42.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_21.V1), v_4, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t43 = gopurs_runtime.Apply(pure_3_22, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0})})
goto end_branch_43
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t43 = gopurs_runtime.Apply(k_5, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0)
goto end_branch_43
} else {

}
}
{
__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_43:
return __t43
}))
})
}))
}))
_ = monadExceptT1_1_0
// TAST (Let): __local_var_2_45 -> gopurs_runtime.Value
__local_var_2_45 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_2_45
// TAST (Let): __local_var_2_44 -> gopurs_runtime.Value
__local_var_2_44 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_45, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_3})})
})
_ = __local_var_2_44
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_1_0
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_44, x_3)
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
// TAST (Let): monadExceptT1_3_3 -> gopurs_runtime.Value
monadExceptT1_3_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_83 -> gopurs_runtime.Value
__local_var_4_83 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_83
// TAST (Let): __local_var_4_82 -> gopurs_runtime.Value
__local_var_4_82 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_83, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_82
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorExceptT1_5_4 -> gopurs_runtime.Value
functorExceptT1_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_6))
_ = __local_var_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, v_8)
})
}))
_ = functorExceptT1_5_4
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_27
// TAST (Let): __local_var_7_26 -> gopurs_runtime.Value
__local_var_7_26 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_27, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_8})})
})
_ = __local_var_7_26
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_9
// TAST (Let): functorExceptT1_8_8 -> gopurs_runtime.Value
functorExceptT1_8_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_9))
_ = __local_var_10_10
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_10, v_11)
})
}))
_ = functorExceptT1_8_8
// TAST (Let): __local_var_9_11 -> gopurs_runtime.Value
__local_var_9_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_12 -> *Constructor_Control_Bind_Bind
Bind1_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_12
// TAST (Let): pure_11_13 -> gopurs_runtime.Value
pure_11_13 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_13
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_15 -> gopurs_runtime.Value
__local_var_13_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_15
// TAST (Let): functorExceptT1_13_14 -> gopurs_runtime.Value
functorExceptT1_13_14 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_16 -> gopurs_runtime.Value
__local_var_15_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_15, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_14))
_ = __local_var_15_16
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_16, v_16)
})
}))
_ = functorExceptT1_13_14
// TAST (Let): __local_var_14_17 -> gopurs_runtime.Value
__local_var_14_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_18 -> *Constructor_Control_Bind_Bind
Bind1_15_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_18
// TAST (Let): pure_16_19 -> gopurs_runtime.Value
pure_16_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_19
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_18.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(pure_16_19, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_20
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t20 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_20
} else {

}
}
{
__t20 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_20:
return __t20
}))
})
}))
}))
_ = __local_var_14_17
// TAST (Let): Bind1_15_21 -> *Constructor_Control_Bind_Bind
Bind1_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_21
// TAST (Let): Applicative0_16_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_22
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_13_14
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_21.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_21.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_22.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_12.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t23 = gopurs_runtime.Apply(pure_11_13, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_23
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t23 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_23
} else {

}
}
{
__t23 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_23:
return __t23
}))
})
}))
}))
_ = __local_var_9_11
// TAST (Let): Bind1_10_24 -> *Constructor_Control_Bind_Bind
Bind1_10_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_24
// TAST (Let): Applicative0_11_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_8_8
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_24.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_24.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_25.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_26, x_8)
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_28 -> *Constructor_Control_Bind_Bind
Bind1_7_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_28
// TAST (Let): pure_8_29 -> gopurs_runtime.Value
pure_8_29 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_29
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_31
// TAST (Let): functorExceptT1_10_30 -> gopurs_runtime.Value
functorExceptT1_10_30 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_32 -> gopurs_runtime.Value
__local_var_12_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_32
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_32, v_13)
})
}))
_ = functorExceptT1_10_30
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_53 -> gopurs_runtime.Value
__local_var_12_53 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_53
// TAST (Let): __local_var_12_52 -> gopurs_runtime.Value
__local_var_12_52 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_53, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_52
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_35 -> gopurs_runtime.Value
__local_var_13_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_35
// TAST (Let): functorExceptT1_13_34 -> gopurs_runtime.Value
functorExceptT1_13_34 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_36 -> gopurs_runtime.Value
__local_var_15_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_35, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_14))
_ = __local_var_15_36
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_36, v_16)
})
}))
_ = functorExceptT1_13_34
// TAST (Let): __local_var_14_37 -> gopurs_runtime.Value
__local_var_14_37 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_38 -> *Constructor_Control_Bind_Bind
Bind1_15_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_38
// TAST (Let): pure_16_39 -> gopurs_runtime.Value
pure_16_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_39
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_41 -> gopurs_runtime.Value
__local_var_18_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_41
// TAST (Let): functorExceptT1_18_40 -> gopurs_runtime.Value
functorExceptT1_18_40 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_42 -> gopurs_runtime.Value
__local_var_20_42 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_41, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_19))
_ = __local_var_20_42
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_42, v_21)
})
}))
_ = functorExceptT1_18_40
// TAST (Let): __local_var_19_43 -> gopurs_runtime.Value
__local_var_19_43 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_44 -> *Constructor_Control_Bind_Bind
Bind1_20_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_44
// TAST (Let): pure_21_45 -> gopurs_runtime.Value
pure_21_45 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_45
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_44.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t46 = gopurs_runtime.Apply(pure_21_45, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_46
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t46 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_46
} else {

}
}
{
__t46 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_46:
return __t46
}))
})
}))
}))
_ = __local_var_19_43
// TAST (Let): Bind1_20_47 -> *Constructor_Control_Bind_Bind
Bind1_20_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_43, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_47
// TAST (Let): Applicative0_21_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_43, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_48
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_18_40
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_47.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_47.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_48.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_38.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t49 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t49 = gopurs_runtime.Apply(pure_16_39, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_49
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t49 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_49
} else {

}
}
{
__t49 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_49:
return __t49
}))
})
}))
}))
_ = __local_var_14_37
// TAST (Let): Bind1_15_50 -> *Constructor_Control_Bind_Bind
Bind1_15_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_50
// TAST (Let): Applicative0_16_51 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_51
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_13_34
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_50.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_50.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_51.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_52, x_13)
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_54 -> *Constructor_Control_Bind_Bind
Bind1_12_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_54
// TAST (Let): pure_13_55 -> gopurs_runtime.Value
pure_13_55 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_55
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_57 -> gopurs_runtime.Value
__local_var_15_57 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_57
// TAST (Let): functorExceptT1_15_56 -> gopurs_runtime.Value
functorExceptT1_15_56 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_58 -> gopurs_runtime.Value
__local_var_17_58 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_57, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_16))
_ = __local_var_17_58
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_58, v_18)
})
}))
_ = functorExceptT1_15_56
// TAST (Let): __local_var_16_59 -> gopurs_runtime.Value
__local_var_16_59 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_70 -> gopurs_runtime.Value
__local_var_17_70 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_70
// TAST (Let): __local_var_17_69 -> gopurs_runtime.Value
__local_var_17_69 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_70, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_18})})
})
_ = __local_var_17_69
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_61 -> gopurs_runtime.Value
__local_var_18_61 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_61
// TAST (Let): functorExceptT1_18_60 -> gopurs_runtime.Value
functorExceptT1_18_60 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_62 -> gopurs_runtime.Value
__local_var_20_62 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_61, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_19))
_ = __local_var_20_62
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_62, v_21)
})
}))
_ = functorExceptT1_18_60
// TAST (Let): __local_var_19_63 -> gopurs_runtime.Value
__local_var_19_63 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_64 -> *Constructor_Control_Bind_Bind
Bind1_20_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_64
// TAST (Let): pure_21_65 -> gopurs_runtime.Value
pure_21_65 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_65
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_64.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t66 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t66 = gopurs_runtime.Apply(pure_21_65, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_66
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t66 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_66
} else {

}
}
{
__t66 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_66:
return __t66
}))
})
}))
}))
_ = __local_var_19_63
// TAST (Let): Bind1_20_67 -> *Constructor_Control_Bind_Bind
Bind1_20_67 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_63, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_67
// TAST (Let): Applicative0_21_68 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_68 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_63, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_68
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_18_60
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_67.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_67.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_68.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_69, x_18)
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_71 -> *Constructor_Control_Bind_Bind
Bind1_17_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_71
// TAST (Let): pure_18_72 -> gopurs_runtime.Value
pure_18_72 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_72
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_71.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t73 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t73 = gopurs_runtime.Apply(pure_18_72, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_73
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t73 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_73
} else {

}
}
{
__t73 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_73:
return __t73
}))
})
}))
}))
_ = __local_var_16_59
// TAST (Let): Bind1_17_74 -> *Constructor_Control_Bind_Bind
Bind1_17_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_59, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_74
// TAST (Let): Applicative0_18_75 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_59, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_75
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_15_56
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_74.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_74.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_75.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_54.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t76 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t76 = gopurs_runtime.Apply(pure_13_55, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_76
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t76 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_76
} else {

}
}
{
__t76 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_76:
return __t76
}))
})
}))
}))
_ = __local_var_11_33
// TAST (Let): Bind1_12_77 -> *Constructor_Control_Bind_Bind
Bind1_12_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_77
// TAST (Let): Applicative0_13_78 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_78
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_30
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_77.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_77.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_78.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_28.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t79 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t79 = gopurs_runtime.Apply(pure_8_29, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_79
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t79 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_79
} else {

}
}
{
__t79 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_79:
return __t79
}))
})
}))
}))
_ = __local_var_6_7
// TAST (Let): Bind1_7_80 -> *Constructor_Control_Bind_Bind
Bind1_7_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_80
// TAST (Let): Applicative0_8_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_81
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_5_4
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_80.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_80.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_81.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_82, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_84 -> *Constructor_Control_Bind_Bind
Bind1_4_84 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_84
// TAST (Let): pure_5_85 -> gopurs_runtime.Value
pure_5_85 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_85
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_87 -> gopurs_runtime.Value
__local_var_7_87 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_87
// TAST (Let): functorExceptT1_7_86 -> gopurs_runtime.Value
functorExceptT1_7_86 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_88 -> gopurs_runtime.Value
__local_var_9_88 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_87, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_8))
_ = __local_var_9_88
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_88, v_10)
})
}))
_ = functorExceptT1_7_86
// TAST (Let): __local_var_8_89 -> gopurs_runtime.Value
__local_var_8_89 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_140 -> gopurs_runtime.Value
__local_var_9_140 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_140
// TAST (Let): __local_var_9_139 -> gopurs_runtime.Value
__local_var_9_139 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_140, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_139
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_91 -> gopurs_runtime.Value
__local_var_10_91 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_91
// TAST (Let): functorExceptT1_10_90 -> gopurs_runtime.Value
functorExceptT1_10_90 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_92 -> gopurs_runtime.Value
__local_var_12_92 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_91, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_92
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_92, v_13)
})
}))
_ = functorExceptT1_10_90
// TAST (Let): __local_var_11_93 -> gopurs_runtime.Value
__local_var_11_93 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_113 -> gopurs_runtime.Value
__local_var_12_113 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_113
// TAST (Let): __local_var_12_112 -> gopurs_runtime.Value
__local_var_12_112 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_113, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_112
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_95 -> gopurs_runtime.Value
__local_var_13_95 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_95
// TAST (Let): functorExceptT1_13_94 -> gopurs_runtime.Value
functorExceptT1_13_94 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_96 -> gopurs_runtime.Value
__local_var_15_96 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_95, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_14))
_ = __local_var_15_96
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_96, v_16)
})
}))
_ = functorExceptT1_13_94
// TAST (Let): __local_var_14_97 -> gopurs_runtime.Value
__local_var_14_97 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_98 -> *Constructor_Control_Bind_Bind
Bind1_15_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_98
// TAST (Let): pure_16_99 -> gopurs_runtime.Value
pure_16_99 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_99
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_101 -> gopurs_runtime.Value
__local_var_18_101 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_101
// TAST (Let): functorExceptT1_18_100 -> gopurs_runtime.Value
functorExceptT1_18_100 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_102 -> gopurs_runtime.Value
__local_var_20_102 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_101, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_19))
_ = __local_var_20_102
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_102, v_21)
})
}))
_ = functorExceptT1_18_100
// TAST (Let): __local_var_19_103 -> gopurs_runtime.Value
__local_var_19_103 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_104 -> *Constructor_Control_Bind_Bind
Bind1_20_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_104
// TAST (Let): pure_21_105 -> gopurs_runtime.Value
pure_21_105 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_105
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_104.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t106 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t106 = gopurs_runtime.Apply(pure_21_105, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_106
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t106 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_106
} else {

}
}
{
__t106 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_106:
return __t106
}))
})
}))
}))
_ = __local_var_19_103
// TAST (Let): Bind1_20_107 -> *Constructor_Control_Bind_Bind
Bind1_20_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_103, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_107
// TAST (Let): Applicative0_21_108 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_103, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_108
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_18_100
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_107.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_107.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_108.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_98.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t109 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t109 = gopurs_runtime.Apply(pure_16_99, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_109
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t109 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_109
} else {

}
}
{
__t109 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_109:
return __t109
}))
})
}))
}))
_ = __local_var_14_97
// TAST (Let): Bind1_15_110 -> *Constructor_Control_Bind_Bind
Bind1_15_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_97, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_110
// TAST (Let): Applicative0_16_111 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_97, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_111
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_13_94
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_110.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_110.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_111.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_112, x_13)
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_114 -> *Constructor_Control_Bind_Bind
Bind1_12_114 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_114
// TAST (Let): pure_13_115 -> gopurs_runtime.Value
pure_13_115 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_115
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_117 -> gopurs_runtime.Value
__local_var_15_117 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_117
// TAST (Let): functorExceptT1_15_116 -> gopurs_runtime.Value
functorExceptT1_15_116 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_118 -> gopurs_runtime.Value
__local_var_17_118 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_117, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_16))
_ = __local_var_17_118
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_118, v_18)
})
}))
_ = functorExceptT1_15_116
// TAST (Let): __local_var_16_119 -> gopurs_runtime.Value
__local_var_16_119 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_130 -> gopurs_runtime.Value
__local_var_17_130 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_130
// TAST (Let): __local_var_17_129 -> gopurs_runtime.Value
__local_var_17_129 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_130, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_18})})
})
_ = __local_var_17_129
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_121 -> gopurs_runtime.Value
__local_var_18_121 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_121
// TAST (Let): functorExceptT1_18_120 -> gopurs_runtime.Value
functorExceptT1_18_120 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_122 -> gopurs_runtime.Value
__local_var_20_122 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_121, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_19))
_ = __local_var_20_122
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_122, v_21)
})
}))
_ = functorExceptT1_18_120
// TAST (Let): __local_var_19_123 -> gopurs_runtime.Value
__local_var_19_123 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_124 -> *Constructor_Control_Bind_Bind
Bind1_20_124 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_124
// TAST (Let): pure_21_125 -> gopurs_runtime.Value
pure_21_125 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_125
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_124.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t126 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t126 = gopurs_runtime.Apply(pure_21_125, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_126
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t126 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_126
} else {

}
}
{
__t126 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_126:
return __t126
}))
})
}))
}))
_ = __local_var_19_123
// TAST (Let): Bind1_20_127 -> *Constructor_Control_Bind_Bind
Bind1_20_127 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_123, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_127
// TAST (Let): Applicative0_21_128 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_128 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_123, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_128
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_18_120
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_127.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_127.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_128.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_129, x_18)
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_131 -> *Constructor_Control_Bind_Bind
Bind1_17_131 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_131
// TAST (Let): pure_18_132 -> gopurs_runtime.Value
pure_18_132 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_132
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_131.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t133 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t133 = gopurs_runtime.Apply(pure_18_132, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_133
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t133 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_133
} else {

}
}
{
__t133 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_133:
return __t133
}))
})
}))
}))
_ = __local_var_16_119
// TAST (Let): Bind1_17_134 -> *Constructor_Control_Bind_Bind
Bind1_17_134 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_119, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_134
// TAST (Let): Applicative0_18_135 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_135 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_119, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_135
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_15_116
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_134.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_134.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_135.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_114.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t136 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t136 = gopurs_runtime.Apply(pure_13_115, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_136
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t136 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_136
} else {

}
}
{
__t136 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_136:
return __t136
}))
})
}))
}))
_ = __local_var_11_93
// TAST (Let): Bind1_12_137 -> *Constructor_Control_Bind_Bind
Bind1_12_137 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_93, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_137
// TAST (Let): Applicative0_13_138 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_138 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_93, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_138
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_90
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_137.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_137.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_138.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_139, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_141 -> *Constructor_Control_Bind_Bind
Bind1_9_141 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_141
// TAST (Let): pure_10_142 -> gopurs_runtime.Value
pure_10_142 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_142
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_144 -> gopurs_runtime.Value
__local_var_12_144 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_144
// TAST (Let): functorExceptT1_12_143 -> gopurs_runtime.Value
functorExceptT1_12_143 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_145 -> gopurs_runtime.Value
__local_var_14_145 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_144, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_13))
_ = __local_var_14_145
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_145, v_15)
})
}))
_ = functorExceptT1_12_143
// TAST (Let): __local_var_13_146 -> gopurs_runtime.Value
__local_var_13_146 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_157 -> gopurs_runtime.Value
__local_var_14_157 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_14_157
// TAST (Let): __local_var_14_156 -> gopurs_runtime.Value
__local_var_14_156 := gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_157, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_15})})
})
_ = __local_var_14_156
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_148 -> gopurs_runtime.Value
__local_var_15_148 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_148
// TAST (Let): functorExceptT1_15_147 -> gopurs_runtime.Value
functorExceptT1_15_147 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_149 -> gopurs_runtime.Value
__local_var_17_149 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_148, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_16))
_ = __local_var_17_149
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_149, v_18)
})
}))
_ = functorExceptT1_15_147
// TAST (Let): __local_var_16_150 -> gopurs_runtime.Value
__local_var_16_150 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_151 -> *Constructor_Control_Bind_Bind
Bind1_17_151 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_151
// TAST (Let): pure_18_152 -> gopurs_runtime.Value
pure_18_152 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_152
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_151.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t153 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t153 = gopurs_runtime.Apply(pure_18_152, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_153
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t153 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_153
} else {

}
}
{
__t153 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_153:
return __t153
}))
})
}))
}))
_ = __local_var_16_150
// TAST (Let): Bind1_17_154 -> *Constructor_Control_Bind_Bind
Bind1_17_154 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_150, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_154
// TAST (Let): Applicative0_18_155 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_155 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_150, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_155
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_15_147
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_154.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_154.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_155.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_156, x_15)
}))
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_14_158 -> *Constructor_Control_Bind_Bind
Bind1_14_158 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_158
// TAST (Let): pure_15_159 -> gopurs_runtime.Value
pure_15_159 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_15_159
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_158.V1), v_16, gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t160 gopurs_runtime.Value
{
if (v2_18.Type == 9 && v2_18.IntVal == 3711209382) {
__t160 = gopurs_runtime.Apply(pure_15_159, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_18.UnsafePtr).V0})})
goto end_branch_160
} else {

}
}
{
if (v2_18.Type == 9 && v2_18.IntVal == 2465973597) {
__t160 = gopurs_runtime.Apply(k_17, (*Constructor_Data_Either_Right)(v2_18.UnsafePtr).V0)
goto end_branch_160
} else {

}
}
{
__t160 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_160:
return __t160
}))
})
}))
}))
_ = __local_var_13_146
// TAST (Let): Bind1_14_161 -> *Constructor_Control_Bind_Bind
Bind1_14_161 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_146, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_161
// TAST (Let): Applicative0_15_162 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_162 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_146, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_162
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_12_143
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_161.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_161.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_162.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_141.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t163 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t163 = gopurs_runtime.Apply(pure_10_142, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_163
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t163 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
goto end_branch_163
} else {

}
}
{
__t163 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_163:
return __t163
}))
})
}))
}))
_ = __local_var_8_89
// TAST (Let): Bind1_9_164 -> *Constructor_Control_Bind_Bind
Bind1_9_164 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_89, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_164
// TAST (Let): Applicative0_10_165 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_165 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_89, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_165
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_7_86
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_164.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_164.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_165.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_84.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t166 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t166 = gopurs_runtime.Apply(pure_5_85, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_166
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t166 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
goto end_branch_166
} else {

}
}
{
__t166 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_166:
return __t166
}))
})
}))
}))
_ = monadExceptT1_3_3
// TAST (Let): __local_var_4_168 -> gopurs_runtime.Value
__local_var_4_168 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_168
// TAST (Let): __local_var_4_167 -> gopurs_runtime.Value
__local_var_4_167 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_168, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_5})})
})
_ = __local_var_4_167
// TAST (Let): monadThrowExceptT1_3_2 -> gopurs_runtime.Value
monadThrowExceptT1_3_2 := gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_3_3
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_167, x_5)
}))
_ = monadThrowExceptT1_3_2
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowExceptT1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_4, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t169 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t169 = gopurs_runtime.Apply(k_5, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0)
goto end_branch_169
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t169 = gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0})})
goto end_branch_169
} else {

}
}
{
__t169 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_169:
return __t169
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
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_21
// TAST (Let): __local_var_3_20 -> gopurs_runtime.Value
__local_var_3_20 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_21, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_20
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 -> gopurs_runtime.Value
functorExceptT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_5))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
}))
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_6 -> *Constructor_Control_Bind_Bind
Bind1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_6
// TAST (Let): pure_7_7 -> gopurs_runtime.Value
pure_7_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_7
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_9
// TAST (Let): functorExceptT1_9_8 -> gopurs_runtime.Value
functorExceptT1_9_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_10))
_ = __local_var_11_10
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_10, v_12)
})
}))
_ = functorExceptT1_9_8
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_12 -> *Constructor_Control_Bind_Bind
Bind1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_12
// TAST (Let): pure_12_13 -> gopurs_runtime.Value
pure_12_13 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_13
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_12.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t14 = gopurs_runtime.Apply(pure_12_13, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_14
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t14 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
}))
})
}))
}))
_ = __local_var_10_11
// TAST (Let): Bind1_11_15 -> *Constructor_Control_Bind_Bind
Bind1_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_15
// TAST (Let): Applicative0_12_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_9_8
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_16.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_6.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(pure_7_7, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_17
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
goto end_branch_17
} else {

}
}
{
__t17 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_17:
return __t17
}))
})
}))
}))
_ = __local_var_5_5
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_4_2
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_20, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_22 -> *Constructor_Control_Bind_Bind
Bind1_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_22
// TAST (Let): pure_4_23 -> gopurs_runtime.Value
pure_4_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_25 -> gopurs_runtime.Value
__local_var_6_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_25
// TAST (Let): functorExceptT1_6_24 -> gopurs_runtime.Value
functorExceptT1_6_24 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_26 -> gopurs_runtime.Value
__local_var_8_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_25, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_7))
_ = __local_var_8_26
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_26, v_9)
})
}))
_ = functorExceptT1_6_24
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_38 -> gopurs_runtime.Value
__local_var_8_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_38
// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
__local_var_8_37 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_38, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_37
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_29
// TAST (Let): functorExceptT1_9_28 -> gopurs_runtime.Value
functorExceptT1_9_28 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_30 -> gopurs_runtime.Value
__local_var_11_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_29, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_10))
_ = __local_var_11_30
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_30, v_12)
})
}))
_ = functorExceptT1_9_28
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_32 -> *Constructor_Control_Bind_Bind
Bind1_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_32
// TAST (Let): pure_12_33 -> gopurs_runtime.Value
pure_12_33 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_33
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_32.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t34 = gopurs_runtime.Apply(pure_12_33, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_34
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t34 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_34
} else {

}
}
{
__t34 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_34:
return __t34
}))
})
}))
}))
_ = __local_var_10_31
// TAST (Let): Bind1_11_35 -> *Constructor_Control_Bind_Bind
Bind1_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_35
// TAST (Let): Applicative0_12_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_36
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_9_28
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_36.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_37, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_39 -> *Constructor_Control_Bind_Bind
Bind1_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_39
// TAST (Let): pure_9_40 -> gopurs_runtime.Value
pure_9_40 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_40
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_39.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t41 = gopurs_runtime.Apply(pure_9_40, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_41
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t41 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
goto end_branch_41
} else {

}
}
{
__t41 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_41:
return __t41
}))
})
}))
}))
_ = __local_var_7_27
// TAST (Let): Bind1_8_42 -> *Constructor_Control_Bind_Bind
Bind1_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_42
// TAST (Let): Applicative0_9_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_43
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_6_24
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_43.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_22.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t44 = gopurs_runtime.Apply(pure_4_23, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_44
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t44 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_44
} else {

}
}
{
__t44 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_44:
return __t44
}))
})
}))
}))
_ = monadExceptT1_2_1
// TAST (Let): Bind1_3_46 -> *Constructor_Control_Bind_Bind
Bind1_3_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_46
// TAST (Let): pure_4_47 -> gopurs_runtime.Value
pure_4_47 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_47
// TAST (Let): __local_var_3_45 -> gopurs_runtime.Value
__local_var_3_45 := gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_46.V1), m_5, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_47, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_6})})
}))
})
_ = __local_var_3_45
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_45, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_Control_Monad_Except_Trans_monoidExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_20 -> gopurs_runtime.Value
__local_var_1_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_20
// TAST (Let): __local_var_1_19 -> gopurs_runtime.Value
__local_var_1_19 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_20, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_2})})
})
_ = __local_var_1_19
// TAST (Let): applicativeExceptT1_1_0 -> *Constructor_Control_Applicative_Applicative
applicativeExceptT1_1_0 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorExceptT1_2_1 -> gopurs_runtime.Value
functorExceptT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_3))
_ = __local_var_4_3
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, v_5)
})
}))
_ = functorExceptT1_2_1
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_5 -> *Constructor_Control_Bind_Bind
Bind1_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_5
// TAST (Let): pure_5_6 -> gopurs_runtime.Value
pure_5_6 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_8
// TAST (Let): functorExceptT1_7_7 -> gopurs_runtime.Value
functorExceptT1_7_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_8))
_ = __local_var_9_9
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_9, v_10)
})
}))
_ = functorExceptT1_7_7
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_11 -> *Constructor_Control_Bind_Bind
Bind1_9_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_11
// TAST (Let): pure_10_12 -> gopurs_runtime.Value
pure_10_12 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_12
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_11.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(pure_10_12, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_13
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t13 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}))
})
}))
}))
_ = __local_var_8_10
// TAST (Let): Bind1_9_14 -> *Constructor_Control_Bind_Bind
Bind1_9_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_14
// TAST (Let): Applicative0_10_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_15
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_7_7
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_14.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_14.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_15.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_5.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(pure_5_6, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_16
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t16 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}))
})
}))
}))
_ = __local_var_3_4
// TAST (Let): Bind1_4_17 -> *Constructor_Control_Bind_Bind
Bind1_4_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_17
// TAST (Let): Applicative0_5_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_2_1
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_17.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_17.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_18.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_19, x_2)
})}
_ = applicativeExceptT1_1_0
// TAST (Let): __local_var_2_23 -> gopurs_runtime.Value
__local_var_2_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_23
// TAST (Let): functorExceptT1_2_22 -> gopurs_runtime.Value
functorExceptT1_2_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_24 -> gopurs_runtime.Value
__local_var_4_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_23, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_3))
_ = __local_var_4_24
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_24, v_5)
})
}))
_ = functorExceptT1_2_22
// TAST (Let): __local_var_3_25 -> gopurs_runtime.Value
__local_var_3_25 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_76 -> gopurs_runtime.Value
__local_var_4_76 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_76
// TAST (Let): __local_var_4_75 -> gopurs_runtime.Value
__local_var_4_75 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_76, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_75
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_27 -> gopurs_runtime.Value
__local_var_5_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_27
// TAST (Let): functorExceptT1_5_26 -> gopurs_runtime.Value
functorExceptT1_5_26 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_28 -> gopurs_runtime.Value
__local_var_7_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_27, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_6))
_ = __local_var_7_28
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_28, v_8)
})
}))
_ = functorExceptT1_5_26
// TAST (Let): __local_var_6_29 -> gopurs_runtime.Value
__local_var_6_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_49 -> gopurs_runtime.Value
__local_var_7_49 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_49
// TAST (Let): __local_var_7_48 -> gopurs_runtime.Value
__local_var_7_48 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_49, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_8})})
})
_ = __local_var_7_48
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_31 -> gopurs_runtime.Value
__local_var_8_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_31
// TAST (Let): functorExceptT1_8_30 -> gopurs_runtime.Value
functorExceptT1_8_30 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_32 -> gopurs_runtime.Value
__local_var_10_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_31, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_9))
_ = __local_var_10_32
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_32, v_11)
})
}))
_ = functorExceptT1_8_30
// TAST (Let): __local_var_9_33 -> gopurs_runtime.Value
__local_var_9_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_34 -> *Constructor_Control_Bind_Bind
Bind1_10_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_34
// TAST (Let): pure_11_35 -> gopurs_runtime.Value
pure_11_35 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_35
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_37 -> gopurs_runtime.Value
__local_var_13_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_37
// TAST (Let): functorExceptT1_13_36 -> gopurs_runtime.Value
functorExceptT1_13_36 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_38 -> gopurs_runtime.Value
__local_var_15_38 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_37, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_14))
_ = __local_var_15_38
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_38, v_16)
})
}))
_ = functorExceptT1_13_36
// TAST (Let): __local_var_14_39 -> gopurs_runtime.Value
__local_var_14_39 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_40 -> *Constructor_Control_Bind_Bind
Bind1_15_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_40
// TAST (Let): pure_16_41 -> gopurs_runtime.Value
pure_16_41 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_41
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_40.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t42 = gopurs_runtime.Apply(pure_16_41, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_42
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t42 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_42
} else {

}
}
{
__t42 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_42:
return __t42
}))
})
}))
}))
_ = __local_var_14_39
// TAST (Let): Bind1_15_43 -> *Constructor_Control_Bind_Bind
Bind1_15_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_39, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_43
// TAST (Let): Applicative0_16_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_39, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_44
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_13_36
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_43.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_43.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_44.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_34.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t45 = gopurs_runtime.Apply(pure_11_35, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_45
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t45 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_45
} else {

}
}
{
__t45 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_45:
return __t45
}))
})
}))
}))
_ = __local_var_9_33
// TAST (Let): Bind1_10_46 -> *Constructor_Control_Bind_Bind
Bind1_10_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_46
// TAST (Let): Applicative0_11_47 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_47
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_8_30
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_46.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_46.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_47.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_48, x_8)
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_50 -> *Constructor_Control_Bind_Bind
Bind1_7_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_50
// TAST (Let): pure_8_51 -> gopurs_runtime.Value
pure_8_51 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_51
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_53 -> gopurs_runtime.Value
__local_var_10_53 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_53
// TAST (Let): functorExceptT1_10_52 -> gopurs_runtime.Value
functorExceptT1_10_52 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_54 -> gopurs_runtime.Value
__local_var_12_54 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_53, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_54
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_54, v_13)
})
}))
_ = functorExceptT1_10_52
// TAST (Let): __local_var_11_55 -> gopurs_runtime.Value
__local_var_11_55 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_66 -> gopurs_runtime.Value
__local_var_12_66 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_66
// TAST (Let): __local_var_12_65 -> gopurs_runtime.Value
__local_var_12_65 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_66, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_65
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_57 -> gopurs_runtime.Value
__local_var_13_57 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_57
// TAST (Let): functorExceptT1_13_56 -> gopurs_runtime.Value
functorExceptT1_13_56 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_58 -> gopurs_runtime.Value
__local_var_15_58 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_57, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_14))
_ = __local_var_15_58
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_58, v_16)
})
}))
_ = functorExceptT1_13_56
// TAST (Let): __local_var_14_59 -> gopurs_runtime.Value
__local_var_14_59 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_60 -> *Constructor_Control_Bind_Bind
Bind1_15_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_60
// TAST (Let): pure_16_61 -> gopurs_runtime.Value
pure_16_61 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_61
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_60.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t62 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t62 = gopurs_runtime.Apply(pure_16_61, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_62
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t62 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_62
} else {

}
}
{
__t62 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_62:
return __t62
}))
})
}))
}))
_ = __local_var_14_59
// TAST (Let): Bind1_15_63 -> *Constructor_Control_Bind_Bind
Bind1_15_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_59, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_63
// TAST (Let): Applicative0_16_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_59, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_64
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_13_56
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_63.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_63.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_64.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_65, x_13)
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_67 -> *Constructor_Control_Bind_Bind
Bind1_12_67 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_67
// TAST (Let): pure_13_68 -> gopurs_runtime.Value
pure_13_68 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_68
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_67.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t69 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t69 = gopurs_runtime.Apply(pure_13_68, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_69
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t69 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_69
} else {

}
}
{
__t69 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_69:
return __t69
}))
})
}))
}))
_ = __local_var_11_55
// TAST (Let): Bind1_12_70 -> *Constructor_Control_Bind_Bind
Bind1_12_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_55, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_70
// TAST (Let): Applicative0_13_71 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_55, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_71
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_52
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_70.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_70.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_71.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_50.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t72 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t72 = gopurs_runtime.Apply(pure_8_51, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_72
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t72 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_72
} else {

}
}
{
__t72 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_72:
return __t72
}))
})
}))
}))
_ = __local_var_6_29
// TAST (Let): Bind1_7_73 -> *Constructor_Control_Bind_Bind
Bind1_7_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_73
// TAST (Let): Applicative0_8_74 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_74
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_5_26
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_73.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_73.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_74.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_75, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_77 -> *Constructor_Control_Bind_Bind
Bind1_4_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_77
// TAST (Let): pure_5_78 -> gopurs_runtime.Value
pure_5_78 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_78
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_80 -> gopurs_runtime.Value
__local_var_7_80 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_80
// TAST (Let): functorExceptT1_7_79 -> gopurs_runtime.Value
functorExceptT1_7_79 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_81 -> gopurs_runtime.Value
__local_var_9_81 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_80, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_8))
_ = __local_var_9_81
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_81, v_10)
})
}))
_ = functorExceptT1_7_79
// TAST (Let): __local_var_8_82 -> gopurs_runtime.Value
__local_var_8_82 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_93 -> gopurs_runtime.Value
__local_var_9_93 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_93
// TAST (Let): __local_var_9_92 -> gopurs_runtime.Value
__local_var_9_92 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_93, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_92
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_84 -> gopurs_runtime.Value
__local_var_10_84 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_84
// TAST (Let): functorExceptT1_10_83 -> gopurs_runtime.Value
functorExceptT1_10_83 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_85 -> gopurs_runtime.Value
__local_var_12_85 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_84, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_85
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_85, v_13)
})
}))
_ = functorExceptT1_10_83
// TAST (Let): __local_var_11_86 -> gopurs_runtime.Value
__local_var_11_86 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_87 -> *Constructor_Control_Bind_Bind
Bind1_12_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_87
// TAST (Let): pure_13_88 -> gopurs_runtime.Value
pure_13_88 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_88
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_87.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t89 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t89 = gopurs_runtime.Apply(pure_13_88, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_89
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t89 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_89
} else {

}
}
{
__t89 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_89:
return __t89
}))
})
}))
}))
_ = __local_var_11_86
// TAST (Let): Bind1_12_90 -> *Constructor_Control_Bind_Bind
Bind1_12_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_86, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_90
// TAST (Let): Applicative0_13_91 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_86, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_91
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_83
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_90.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_90.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_91.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_92, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_94 -> *Constructor_Control_Bind_Bind
Bind1_9_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_94
// TAST (Let): pure_10_95 -> gopurs_runtime.Value
pure_10_95 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_95
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_94.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t96 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t96 = gopurs_runtime.Apply(pure_10_95, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_96
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t96 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
goto end_branch_96
} else {

}
}
{
__t96 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_96:
return __t96
}))
})
}))
}))
_ = __local_var_8_82
// TAST (Let): Bind1_9_97 -> *Constructor_Control_Bind_Bind
Bind1_9_97 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_82, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_97
// TAST (Let): Applicative0_10_98 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_82, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_98
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_7_79
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_97.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_97.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_98.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_77.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t99 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t99 = gopurs_runtime.Apply(pure_5_78, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_99
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t99 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
goto end_branch_99
} else {

}
}
{
__t99 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_99:
return __t99
}))
})
}))
}))
_ = __local_var_3_25
// TAST (Let): Bind1_4_100 -> *Constructor_Control_Bind_Bind
Bind1_4_100 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_25, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_100
// TAST (Let): Applicative0_5_101 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_101 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_25, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_101
// TAST (Let): applyExceptT1_2_21 -> *Constructor_Control_Apply_Apply
applyExceptT1_2_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_2_22
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_100.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_100.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_101.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})))
_ = applyExceptT1_2_21
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_103 -> *Constructor_Data_Functor_Functor
Functor0_4_103 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyExceptT1_2_21.V0), gopurs_runtime.Value{}))
_ = Functor0_4_103
// TAST (Let): __local_var_5_104 -> gopurs_runtime.Value
__local_var_5_104 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_5_104
// TAST (Let): semigroupExceptT2_4_102 -> gopurs_runtime.Value
semigroupExceptT2_4_102 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyExceptT1_2_21.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_103.V0), __local_var_5_104, a_6), b_7)
})
}))
_ = semigroupExceptT2_4_102
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupExceptT2_4_102
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
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 -> gopurs_runtime.Value
functorExceptT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_5))
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
}))
_ = functorExceptT1_4_2
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_4_2
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), v_5, gopurs_runtime.Func(func(rm_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (rm_7.Type == 9 && rm_7.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rm_7.UnsafePtr).V0})})
goto end_branch_7
} else {

}
}
{
if (rm_7.Type == 9 && rm_7.IntVal == 3711209382) {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := (*Constructor_Data_Either_Left)(rm_7.UnsafePtr).V0
_ = __local_var_8_5
__t7 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), v1_6, gopurs_runtime.Func(func(rn_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (rn_9.Type == 9 && rn_9.IntVal == 2465973597) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rn_9.UnsafePtr).V0})})
goto end_branch_6
} else {

}
}
{
if (rn_9.Type == 9 && rn_9.IntVal == 3711209382) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_8_5, (*Constructor_Data_Either_Left)(rn_9.UnsafePtr).V0)})})
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
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
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_3 -> *Constructor_Control_Bind_Bind
Bind1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_3
// TAST (Let): Applicative0_5_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_4
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_6
// TAST (Let): functorExceptT1_6_5 -> gopurs_runtime.Value
functorExceptT1_6_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_7 -> gopurs_runtime.Value
__local_var_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_7))
_ = __local_var_8_7
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_7, v_9)
})
}))
_ = functorExceptT1_6_5
// TAST (Let): altExceptT2_4_2 -> gopurs_runtime.Value
altExceptT2_4_2 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_6_5
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), v_7, gopurs_runtime.Func(func(rm_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (rm_9.Type == 9 && rm_9.IntVal == 2465973597) {
__t10 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rm_9.UnsafePtr).V0})})
goto end_branch_10
} else {

}
}
{
if (rm_9.Type == 9 && rm_9.IntVal == 3711209382) {
// TAST (Let): __local_var_10_8 -> gopurs_runtime.Value
__local_var_10_8 := (*Constructor_Data_Either_Left)(rm_9.UnsafePtr).V0
_ = __local_var_10_8
__t10 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), v1_8, gopurs_runtime.Func(func(rn_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (rn_11.Type == 9 && rn_11.IntVal == 2465973597) {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rn_11.UnsafePtr).V0})})
goto end_branch_9
} else {

}
}
{
if (rn_11.Type == 9 && rn_11.IntVal == 3711209382) {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_10_8, (*Constructor_Data_Either_Left)(rn_11.UnsafePtr).V0)})})
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
return __t9
}))
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}))
})
}))
_ = altExceptT2_4_2
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return altExceptT2_4_2
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, mempty_1_0})}))
})
}

func Call_Control_Monad_Except_Trans_alternativeExceptT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_22 -> gopurs_runtime.Value
__local_var_4_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_22
// TAST (Let): __local_var_4_21 -> gopurs_runtime.Value
__local_var_4_21 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_22, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_21
// TAST (Let): applicativeExceptT1_4_2 -> gopurs_runtime.Value
applicativeExceptT1_4_2 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorExceptT1_5_3 -> gopurs_runtime.Value
functorExceptT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_6))
_ = __local_var_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, v_8)
})
}))
_ = functorExceptT1_5_3
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_7 -> *Constructor_Control_Bind_Bind
Bind1_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_7
// TAST (Let): pure_8_8 -> gopurs_runtime.Value
pure_8_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_8
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_10
// TAST (Let): functorExceptT1_10_9 -> gopurs_runtime.Value
functorExceptT1_10_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_11 -> gopurs_runtime.Value
__local_var_12_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_10, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_11))
_ = __local_var_12_11
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_11, v_13)
})
}))
_ = functorExceptT1_10_9
// TAST (Let): __local_var_11_12 -> gopurs_runtime.Value
__local_var_11_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_13 -> *Constructor_Control_Bind_Bind
Bind1_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_13
// TAST (Let): pure_13_14 -> gopurs_runtime.Value
pure_13_14 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_14
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_13.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(pure_13_14, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_15
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t15 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}))
})
}))
}))
_ = __local_var_11_12
// TAST (Let): Bind1_12_16 -> *Constructor_Control_Bind_Bind
Bind1_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_16
// TAST (Let): Applicative0_13_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_10_9
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_17.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_7.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_8_8, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_18
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
})
}))
}))
_ = __local_var_6_6
// TAST (Let): Bind1_7_19 -> *Constructor_Control_Bind_Bind
Bind1_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_19
// TAST (Let): Applicative0_8_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_20
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_20.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_21, x_5)
}))
_ = applicativeExceptT1_4_2
// TAST (Let): Bind1_5_25 -> *Constructor_Control_Bind_Bind
Bind1_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_25
// TAST (Let): Applicative0_6_26 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_26
// TAST (Let): __local_var_7_28 -> gopurs_runtime.Value
__local_var_7_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_28
// TAST (Let): functorExceptT1_7_27 -> gopurs_runtime.Value
functorExceptT1_7_27 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_28, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_8))
_ = __local_var_9_29
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_29, v_10)
})
}))
_ = functorExceptT1_7_27
// TAST (Let): altExceptT2_5_24 -> gopurs_runtime.Value
altExceptT2_5_24 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_7_27
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_25.V1), v_8, gopurs_runtime.Func(func(rm_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 gopurs_runtime.Value
{
if (rm_10.Type == 9 && rm_10.IntVal == 2465973597) {
__t32 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_26.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rm_10.UnsafePtr).V0})})
goto end_branch_32
} else {

}
}
{
if (rm_10.Type == 9 && rm_10.IntVal == 3711209382) {
// TAST (Let): __local_var_11_30 -> gopurs_runtime.Value
__local_var_11_30 := (*Constructor_Data_Either_Left)(rm_10.UnsafePtr).V0
_ = __local_var_11_30
__t32 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_25.V1), v1_9, gopurs_runtime.Func(func(rn_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (rn_12.Type == 9 && rn_12.IntVal == 2465973597) {
__t31 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_26.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rn_12.UnsafePtr).V0})})
goto end_branch_31
} else {

}
}
{
if (rn_12.Type == 9 && rn_12.IntVal == 3711209382) {
__t31 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_26.V1), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_11_30, (*Constructor_Data_Either_Left)(rn_12.UnsafePtr).V0)})})
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return __t31
}))
goto end_branch_32
} else {

}
}
{
__t32 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_32:
return __t32
}))
})
}))
_ = altExceptT2_5_24
// TAST (Let): plusExceptT2_5_23 -> gopurs_runtime.Value
plusExceptT2_5_23 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return altExceptT2_5_24
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, mempty_1_0})}))
_ = plusExceptT2_5_23
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeExceptT1_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return plusExceptT2_5_23
}))
})
}

func Call_Control_Monad_Except_Trans_monadPlusExceptT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): mempty_1_0 -> gopurs_runtime.Value
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): monadExceptT1_4_2 -> gopurs_runtime.Value
monadExceptT1_4_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_22 -> gopurs_runtime.Value
__local_var_5_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_22
// TAST (Let): __local_var_5_21 -> gopurs_runtime.Value
__local_var_5_21 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_22, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_6})})
})
_ = __local_var_5_21
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_4
// TAST (Let): functorExceptT1_6_3 -> gopurs_runtime.Value
functorExceptT1_6_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_7))
_ = __local_var_8_5
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, v_9)
})
}))
_ = functorExceptT1_6_3
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_7 -> *Constructor_Control_Bind_Bind
Bind1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_7
// TAST (Let): pure_9_8 -> gopurs_runtime.Value
pure_9_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_8
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_10
// TAST (Let): functorExceptT1_11_9 -> gopurs_runtime.Value
functorExceptT1_11_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_11 -> gopurs_runtime.Value
__local_var_13_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_12))
_ = __local_var_13_11
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_11, v_14)
})
}))
_ = functorExceptT1_11_9
// TAST (Let): __local_var_12_12 -> gopurs_runtime.Value
__local_var_12_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_13 -> *Constructor_Control_Bind_Bind
Bind1_13_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_13
// TAST (Let): pure_14_14 -> gopurs_runtime.Value
pure_14_14 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_14
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_13.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(pure_14_14, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_15
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t15 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
}))
})
}))
}))
_ = __local_var_12_12
// TAST (Let): Bind1_13_16 -> *Constructor_Control_Bind_Bind
Bind1_13_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_16
// TAST (Let): Applicative0_14_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_11_9
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_16.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_16.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_17.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_9_8, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_18
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
goto end_branch_18
} else {

}
}
{
__t18 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_18:
return __t18
}))
})
}))
}))
_ = __local_var_7_6
// TAST (Let): Bind1_8_19 -> *Constructor_Control_Bind_Bind
Bind1_8_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_19
// TAST (Let): Applicative0_9_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_20
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_6_3
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_19.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_19.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_20.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_21, x_6)
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_23 -> *Constructor_Control_Bind_Bind
Bind1_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_23
// TAST (Let): pure_6_24 -> gopurs_runtime.Value
pure_6_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_24
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_26 -> gopurs_runtime.Value
__local_var_8_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_26
// TAST (Let): functorExceptT1_8_25 -> gopurs_runtime.Value
functorExceptT1_8_25 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_27 -> gopurs_runtime.Value
__local_var_10_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_26, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_9))
_ = __local_var_10_27
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_27, v_11)
})
}))
_ = functorExceptT1_8_25
// TAST (Let): __local_var_9_28 -> gopurs_runtime.Value
__local_var_9_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_39 -> gopurs_runtime.Value
__local_var_10_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_39
// TAST (Let): __local_var_10_38 -> gopurs_runtime.Value
__local_var_10_38 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_39, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_11})})
})
_ = __local_var_10_38
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_30 -> gopurs_runtime.Value
__local_var_11_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_30
// TAST (Let): functorExceptT1_11_29 -> gopurs_runtime.Value
functorExceptT1_11_29 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_31 -> gopurs_runtime.Value
__local_var_13_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_30, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_12))
_ = __local_var_13_31
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_31, v_14)
})
}))
_ = functorExceptT1_11_29
// TAST (Let): __local_var_12_32 -> gopurs_runtime.Value
__local_var_12_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_33 -> *Constructor_Control_Bind_Bind
Bind1_13_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_33
// TAST (Let): pure_14_34 -> gopurs_runtime.Value
pure_14_34 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_34
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_33.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t35 = gopurs_runtime.Apply(pure_14_34, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_35
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t35 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
goto end_branch_35
} else {

}
}
{
__t35 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_35:
return __t35
}))
})
}))
}))
_ = __local_var_12_32
// TAST (Let): Bind1_13_36 -> *Constructor_Control_Bind_Bind
Bind1_13_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_36
// TAST (Let): Applicative0_14_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_37
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_11_29
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_36.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_36.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_37.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_38, x_11)
}))
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_40 -> *Constructor_Control_Bind_Bind
Bind1_10_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_40
// TAST (Let): pure_11_41 -> gopurs_runtime.Value
pure_11_41 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_41
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_40.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t42 = gopurs_runtime.Apply(pure_11_41, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_42
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t42 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_42
} else {

}
}
{
__t42 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_42:
return __t42
}))
})
}))
}))
_ = __local_var_9_28
// TAST (Let): Bind1_10_43 -> *Constructor_Control_Bind_Bind
Bind1_10_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_43
// TAST (Let): Applicative0_11_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_44
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_8_25
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_43.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_43.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_44.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_23.V1), v_7, gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 3711209382) {
__t45 = gopurs_runtime.Apply(pure_6_24, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_9.UnsafePtr).V0})})
goto end_branch_45
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 2465973597) {
__t45 = gopurs_runtime.Apply(k_8, (*Constructor_Data_Either_Right)(v2_9.UnsafePtr).V0)
goto end_branch_45
} else {

}
}
{
__t45 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_45:
return __t45
}))
})
}))
}))
_ = monadExceptT1_4_2
// TAST (Let): __local_var_5_127 -> gopurs_runtime.Value
__local_var_5_127 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_127
// TAST (Let): __local_var_5_126 -> gopurs_runtime.Value
__local_var_5_126 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_127, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_6})})
})
_ = __local_var_5_126
// TAST (Let): applicativeExceptT1_5_47 -> gopurs_runtime.Value
applicativeExceptT1_5_47 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_49 -> gopurs_runtime.Value
__local_var_6_49 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_49
// TAST (Let): functorExceptT1_6_48 -> gopurs_runtime.Value
functorExceptT1_6_48 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_50 -> gopurs_runtime.Value
__local_var_8_50 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_49, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_7))
_ = __local_var_8_50
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_50, v_9)
})
}))
_ = functorExceptT1_6_48
// TAST (Let): __local_var_7_51 -> gopurs_runtime.Value
__local_var_7_51 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_71 -> gopurs_runtime.Value
__local_var_8_71 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_71
// TAST (Let): __local_var_8_70 -> gopurs_runtime.Value
__local_var_8_70 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_71, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_70
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_53 -> gopurs_runtime.Value
__local_var_9_53 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_53
// TAST (Let): functorExceptT1_9_52 -> gopurs_runtime.Value
functorExceptT1_9_52 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_54 -> gopurs_runtime.Value
__local_var_11_54 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_53, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_10))
_ = __local_var_11_54
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_54, v_12)
})
}))
_ = functorExceptT1_9_52
// TAST (Let): __local_var_10_55 -> gopurs_runtime.Value
__local_var_10_55 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_56 -> *Constructor_Control_Bind_Bind
Bind1_11_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_56
// TAST (Let): pure_12_57 -> gopurs_runtime.Value
pure_12_57 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_57
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_59 -> gopurs_runtime.Value
__local_var_14_59 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_59
// TAST (Let): functorExceptT1_14_58 -> gopurs_runtime.Value
functorExceptT1_14_58 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_60 -> gopurs_runtime.Value
__local_var_16_60 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_59, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_15))
_ = __local_var_16_60
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_60, v_17)
})
}))
_ = functorExceptT1_14_58
// TAST (Let): __local_var_15_61 -> gopurs_runtime.Value
__local_var_15_61 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_62 -> *Constructor_Control_Bind_Bind
Bind1_16_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_62
// TAST (Let): pure_17_63 -> gopurs_runtime.Value
pure_17_63 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_63
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_62.V1), v_18, gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t64 gopurs_runtime.Value
{
if (v2_20.Type == 9 && v2_20.IntVal == 3711209382) {
__t64 = gopurs_runtime.Apply(pure_17_63, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_20.UnsafePtr).V0})})
goto end_branch_64
} else {

}
}
{
if (v2_20.Type == 9 && v2_20.IntVal == 2465973597) {
__t64 = gopurs_runtime.Apply(k_19, (*Constructor_Data_Either_Right)(v2_20.UnsafePtr).V0)
goto end_branch_64
} else {

}
}
{
__t64 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_64:
return __t64
}))
})
}))
}))
_ = __local_var_15_61
// TAST (Let): Bind1_16_65 -> *Constructor_Control_Bind_Bind
Bind1_16_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_61, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_65
// TAST (Let): Applicative0_17_66 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_66 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_61, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_66
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_14_58
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_65.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_65.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_66.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_56.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t67 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t67 = gopurs_runtime.Apply(pure_12_57, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_67
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t67 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_67
} else {

}
}
{
__t67 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_67:
return __t67
}))
})
}))
}))
_ = __local_var_10_55
// TAST (Let): Bind1_11_68 -> *Constructor_Control_Bind_Bind
Bind1_11_68 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_55, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_68
// TAST (Let): Applicative0_12_69 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_69 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_55, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_69
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_9_52
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_68.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_68.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_69.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_70, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_72 -> *Constructor_Control_Bind_Bind
Bind1_8_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_72
// TAST (Let): pure_9_73 -> gopurs_runtime.Value
pure_9_73 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_73
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_75 -> gopurs_runtime.Value
__local_var_11_75 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_75
// TAST (Let): functorExceptT1_11_74 -> gopurs_runtime.Value
functorExceptT1_11_74 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_76 -> gopurs_runtime.Value
__local_var_13_76 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_75, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_12))
_ = __local_var_13_76
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_76, v_14)
})
}))
_ = functorExceptT1_11_74
// TAST (Let): __local_var_12_77 -> gopurs_runtime.Value
__local_var_12_77 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_97 -> gopurs_runtime.Value
__local_var_13_97 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_97
// TAST (Let): __local_var_13_96 -> gopurs_runtime.Value
__local_var_13_96 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_97, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_14})})
})
_ = __local_var_13_96
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_79 -> gopurs_runtime.Value
__local_var_14_79 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_79
// TAST (Let): functorExceptT1_14_78 -> gopurs_runtime.Value
functorExceptT1_14_78 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_80 -> gopurs_runtime.Value
__local_var_16_80 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_79, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_15))
_ = __local_var_16_80
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_80, v_17)
})
}))
_ = functorExceptT1_14_78
// TAST (Let): __local_var_15_81 -> gopurs_runtime.Value
__local_var_15_81 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_82 -> *Constructor_Control_Bind_Bind
Bind1_16_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_82
// TAST (Let): pure_17_83 -> gopurs_runtime.Value
pure_17_83 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_83
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_85 -> gopurs_runtime.Value
__local_var_19_85 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_85
// TAST (Let): functorExceptT1_19_84 -> gopurs_runtime.Value
functorExceptT1_19_84 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_86 -> gopurs_runtime.Value
__local_var_21_86 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_85, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_20))
_ = __local_var_21_86
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_86, v_22)
})
}))
_ = functorExceptT1_19_84
// TAST (Let): __local_var_20_87 -> gopurs_runtime.Value
__local_var_20_87 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_88 -> *Constructor_Control_Bind_Bind
Bind1_21_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_88
// TAST (Let): pure_22_89 -> gopurs_runtime.Value
pure_22_89 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_89
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_88.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t90 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t90 = gopurs_runtime.Apply(pure_22_89, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_90
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t90 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
goto end_branch_90
} else {

}
}
{
__t90 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_90:
return __t90
}))
})
}))
}))
_ = __local_var_20_87
// TAST (Let): Bind1_21_91 -> *Constructor_Control_Bind_Bind
Bind1_21_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_87, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_91
// TAST (Let): Applicative0_22_92 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_87, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_92
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_19_84
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_91.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_91.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_92.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_82.V1), v_18, gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t93 gopurs_runtime.Value
{
if (v2_20.Type == 9 && v2_20.IntVal == 3711209382) {
__t93 = gopurs_runtime.Apply(pure_17_83, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_20.UnsafePtr).V0})})
goto end_branch_93
} else {

}
}
{
if (v2_20.Type == 9 && v2_20.IntVal == 2465973597) {
__t93 = gopurs_runtime.Apply(k_19, (*Constructor_Data_Either_Right)(v2_20.UnsafePtr).V0)
goto end_branch_93
} else {

}
}
{
__t93 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_93:
return __t93
}))
})
}))
}))
_ = __local_var_15_81
// TAST (Let): Bind1_16_94 -> *Constructor_Control_Bind_Bind
Bind1_16_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_81, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_94
// TAST (Let): Applicative0_17_95 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_95 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_81, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_95
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_14_78
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_94.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_94.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_95.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_96, x_14)
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_98 -> *Constructor_Control_Bind_Bind
Bind1_13_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_98
// TAST (Let): pure_14_99 -> gopurs_runtime.Value
pure_14_99 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_99
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_101 -> gopurs_runtime.Value
__local_var_16_101 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_101
// TAST (Let): functorExceptT1_16_100 -> gopurs_runtime.Value
functorExceptT1_16_100 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_102 -> gopurs_runtime.Value
__local_var_18_102 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_101, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_17))
_ = __local_var_18_102
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_102, v_19)
})
}))
_ = functorExceptT1_16_100
// TAST (Let): __local_var_17_103 -> gopurs_runtime.Value
__local_var_17_103 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_114 -> gopurs_runtime.Value
__local_var_18_114 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_114
// TAST (Let): __local_var_18_113 -> gopurs_runtime.Value
__local_var_18_113 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_114, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_19})})
})
_ = __local_var_18_113
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_105 -> gopurs_runtime.Value
__local_var_19_105 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_105
// TAST (Let): functorExceptT1_19_104 -> gopurs_runtime.Value
functorExceptT1_19_104 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_106 -> gopurs_runtime.Value
__local_var_21_106 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_105, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_20))
_ = __local_var_21_106
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_106, v_22)
})
}))
_ = functorExceptT1_19_104
// TAST (Let): __local_var_20_107 -> gopurs_runtime.Value
__local_var_20_107 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_108 -> *Constructor_Control_Bind_Bind
Bind1_21_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_108
// TAST (Let): pure_22_109 -> gopurs_runtime.Value
pure_22_109 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_109
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_108.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t110 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t110 = gopurs_runtime.Apply(pure_22_109, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_110
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t110 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
goto end_branch_110
} else {

}
}
{
__t110 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_110:
return __t110
}))
})
}))
}))
_ = __local_var_20_107
// TAST (Let): Bind1_21_111 -> *Constructor_Control_Bind_Bind
Bind1_21_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_107, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_111
// TAST (Let): Applicative0_22_112 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_112 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_107, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_112
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_19_104
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_111.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_111.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_112.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_113, x_19)
}))
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_115 -> *Constructor_Control_Bind_Bind
Bind1_18_115 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_115
// TAST (Let): pure_19_116 -> gopurs_runtime.Value
pure_19_116 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_19_116
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_115.V1), v_20, gopurs_runtime.Func(func(v2_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t117 gopurs_runtime.Value
{
if (v2_22.Type == 9 && v2_22.IntVal == 3711209382) {
__t117 = gopurs_runtime.Apply(pure_19_116, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_22.UnsafePtr).V0})})
goto end_branch_117
} else {

}
}
{
if (v2_22.Type == 9 && v2_22.IntVal == 2465973597) {
__t117 = gopurs_runtime.Apply(k_21, (*Constructor_Data_Either_Right)(v2_22.UnsafePtr).V0)
goto end_branch_117
} else {

}
}
{
__t117 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_117:
return __t117
}))
})
}))
}))
_ = __local_var_17_103
// TAST (Let): Bind1_18_118 -> *Constructor_Control_Bind_Bind
Bind1_18_118 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_103, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_118
// TAST (Let): Applicative0_19_119 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_119 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_103, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_119
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_16_100
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_118.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_118.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_119.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_98.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t120 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t120 = gopurs_runtime.Apply(pure_14_99, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_120
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t120 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
goto end_branch_120
} else {

}
}
{
__t120 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_120:
return __t120
}))
})
}))
}))
_ = __local_var_12_77
// TAST (Let): Bind1_13_121 -> *Constructor_Control_Bind_Bind
Bind1_13_121 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_77, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_121
// TAST (Let): Applicative0_14_122 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_122 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_77, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_122
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_11_74
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_121.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_121.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_122.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_72.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t123 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t123 = gopurs_runtime.Apply(pure_9_73, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_123
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t123 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
goto end_branch_123
} else {

}
}
{
__t123 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_123:
return __t123
}))
})
}))
}))
_ = __local_var_7_51
// TAST (Let): Bind1_8_124 -> *Constructor_Control_Bind_Bind
Bind1_8_124 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_51, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_124
// TAST (Let): Applicative0_9_125 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_125 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_51, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_125
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_6_48
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_124.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_124.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_125.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_126, x_6)
}))
_ = applicativeExceptT1_5_47
// TAST (Let): Bind1_6_130 -> *Constructor_Control_Bind_Bind
Bind1_6_130 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_130
// TAST (Let): Applicative0_7_131 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_131 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_131
// TAST (Let): __local_var_8_133 -> gopurs_runtime.Value
__local_var_8_133 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_133
// TAST (Let): functorExceptT1_8_132 -> gopurs_runtime.Value
functorExceptT1_8_132 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_134 -> gopurs_runtime.Value
__local_var_10_134 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_133, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Either_functorEither(), "map"), f_9))
_ = __local_var_10_134
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_134, v_11)
})
}))
_ = functorExceptT1_8_132
// TAST (Let): altExceptT2_6_129 -> gopurs_runtime.Value
altExceptT2_6_129 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorExceptT1_8_132
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_130.V1), v_9, gopurs_runtime.Func(func(rm_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t137 gopurs_runtime.Value
{
if (rm_11.Type == 9 && rm_11.IntVal == 2465973597) {
__t137 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_131.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rm_11.UnsafePtr).V0})})
goto end_branch_137
} else {

}
}
{
if (rm_11.Type == 9 && rm_11.IntVal == 3711209382) {
// TAST (Let): __local_var_12_135 -> gopurs_runtime.Value
__local_var_12_135 := (*Constructor_Data_Either_Left)(rm_11.UnsafePtr).V0
_ = __local_var_12_135
__t137 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_130.V1), v1_10, gopurs_runtime.Func(func(rn_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t136 gopurs_runtime.Value
{
if (rn_13.Type == 9 && rn_13.IntVal == 2465973597) {
__t136 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_131.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rn_13.UnsafePtr).V0})})
goto end_branch_136
} else {

}
}
{
if (rn_13.Type == 9 && rn_13.IntVal == 3711209382) {
__t136 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_131.V1), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_12_135, (*Constructor_Data_Either_Left)(rn_13.UnsafePtr).V0)})})
goto end_branch_136
} else {

}
}
{
__t136 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_136:
return __t136
}))
goto end_branch_137
} else {

}
}
{
__t137 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_137:
return __t137
}))
})
}))
_ = altExceptT2_6_129
// TAST (Let): plusExceptT2_6_128 -> gopurs_runtime.Value
plusExceptT2_6_128 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return altExceptT2_6_129
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, mempty_1_0})}))
_ = plusExceptT2_6_128
// TAST (Let): alternativeExceptT2_5_46 -> gopurs_runtime.Value
alternativeExceptT2_5_46 := gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeExceptT1_5_47
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return plusExceptT2_6_128
}))
_ = alternativeExceptT2_5_46
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeExceptT2_5_46
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadExceptT1_4_2
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



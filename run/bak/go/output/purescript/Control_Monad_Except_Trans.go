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
		cache_Control_Monad_Except_Trans_newtypeExceptT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Control_Monad_Except_Trans_newtypeExceptT
}

var cache_Control_Monad_Except_Trans_monadTransExceptT gopurs_runtime.Value
var once_Control_Monad_Except_Trans_monadTransExceptT sync.Once
func Get_Control_Monad_Except_Trans_monadTransExceptT() gopurs_runtime.Value {
	once_Control_Monad_Except_Trans_monadTransExceptT.Do(func() {
		cache_Control_Monad_Except_Trans_monadTransExceptT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})})}
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
		cache_Control_Monad_Except_Trans_monadTransExceptT__4007330348 = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
})})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (m_2.Type == 9 && m_2.IntVal == 3711209382) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_2.UnsafePtr).V0})}
goto end_branch_1
} else {

}
}
{
if (m_2.Type == 9 && m_2.IntVal == 2465973597) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_Either_Right)(m_2.UnsafePtr).V0)})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
_ = __local_var_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, v_3)
})
})})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Call_Control_Monad_Except_Trans_bindExceptT(dictMonad_0)))}
})})}
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
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
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
})})}
}

func Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorExceptT1_1_0 -> *Constructor_Data_Functor_Functor
functorExceptT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (m_3.Type == 9 && m_3.IntVal == 3711209382) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_3.UnsafePtr).V0})}
goto end_branch_3
} else {

}
}
{
if (m_3.Type == 9 && m_3.IntVal == 2465973597) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Either_Right)(m_3.UnsafePtr).V0)})}
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
_ = __local_var_3_2
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, v_4)
})
})))
_ = functorExceptT1_1_0
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_5 -> *Constructor_Control_Bind_Bind
Bind1_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_5
// TAST (Let): pure_4_6 -> gopurs_runtime.Value
pure_4_6 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_6
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_5.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(pure_4_6, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_7
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_2_4
// TAST (Let): Bind1_3_8 -> *Constructor_Control_Bind_Bind
Bind1_3_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_8
// TAST (Let): Applicative0_4_9 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_9
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_1_0)}
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_8.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_8.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_9.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
})})}
}

func Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_11 -> gopurs_runtime.Value
__local_var_1_11 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_11
// TAST (Let): __local_var_1_10 -> gopurs_runtime.Value
__local_var_1_10 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_11, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_2})})
})
_ = __local_var_1_10
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): functorExceptT1_2_0 -> *Constructor_Data_Functor_Functor
functorExceptT1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (m_4.Type == 9 && m_4.IntVal == 3711209382) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_4.UnsafePtr).V0})}
goto end_branch_3
} else {

}
}
{
if (m_4.Type == 9 && m_4.IntVal == 2465973597) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Either_Right)(m_4.UnsafePtr).V0)})}
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
_ = __local_var_4_2
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, v_5)
})
})))
_ = functorExceptT1_2_0
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_5 -> *Constructor_Control_Bind_Bind
Bind1_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_5
// TAST (Let): pure_5_6 -> gopurs_runtime.Value
pure_5_6 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_6
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_5.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(pure_5_6, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_7
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_3_4
// TAST (Let): Bind1_4_8 -> *Constructor_Control_Bind_Bind
Bind1_4_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_8
// TAST (Let): Applicative0_5_9 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_9
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_2_0)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_8.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_8.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_9.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_10, x_2)
})})}
}

func Call_Control_Monad_Except_Trans_semigroupExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): functorExceptT1_1_1 -> *Constructor_Data_Functor_Functor
functorExceptT1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "map"), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (m_3.Type == 9 && m_3.IntVal == 3711209382) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_3.UnsafePtr).V0})}
goto end_branch_4
} else {

}
}
{
if (m_3.Type == 9 && m_3.IntVal == 2465973597) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Either_Right)(m_3.UnsafePtr).V0)})}
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
_ = __local_var_3_3
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, v_4)
})
})))
_ = functorExceptT1_1_1
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
__local_var_2_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_17 -> gopurs_runtime.Value
__local_var_3_17 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_17
// TAST (Let): __local_var_3_16 -> gopurs_runtime.Value
__local_var_3_16 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_17, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_16
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_7 -> gopurs_runtime.Value
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): functorExceptT1_4_6 -> *Constructor_Data_Functor_Functor
functorExceptT1_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_8 -> gopurs_runtime.Value
__local_var_6_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_7, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_6.UnsafePtr).V0})}
goto end_branch_9
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t9 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right)(m_6.UnsafePtr).V0)})}
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
_ = __local_var_6_8
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_8, v_7)
})
})))
_ = functorExceptT1_4_6
// TAST (Let): __local_var_5_10 -> gopurs_runtime.Value
__local_var_5_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_11 -> *Constructor_Control_Bind_Bind
Bind1_6_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_11
// TAST (Let): pure_7_12 -> gopurs_runtime.Value
pure_7_12 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_12
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_11.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t13 = gopurs_runtime.Apply(pure_7_12, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_13
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t13 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_5_10
// TAST (Let): Bind1_6_14 -> *Constructor_Control_Bind_Bind
Bind1_6_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_14
// TAST (Let): Applicative0_7_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_15
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_6)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_14.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_14.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_15.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_16, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_18 -> *Constructor_Control_Bind_Bind
Bind1_3_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_18
// TAST (Let): pure_4_19 -> gopurs_runtime.Value
pure_4_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_19
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_18.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(pure_4_19, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_20
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t20 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_2_5
// TAST (Let): Bind1_3_21 -> *Constructor_Control_Bind_Bind
Bind1_3_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_21
// TAST (Let): Applicative0_4_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_22
// TAST (Let): applyExceptT1_1_0 -> *Constructor_Control_Apply_Apply
applyExceptT1_1_0 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_1_1)}
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_21.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_21.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_22.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
})}
_ = applyExceptT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_23 -> *Constructor_Data_Functor_Functor
Functor0_3_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyExceptT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_23
// TAST (Let): __local_var_4_24 -> gopurs_runtime.Value
__local_var_4_24 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_24
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyExceptT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_23.V0), __local_var_4_24, a_5), b_6)
})
})})}
})
}

func Call_Control_Monad_Except_Trans_monadAskExceptT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadExceptT1_1_0 -> *Constructor_Control_Monad_Monad
monadExceptT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_23 -> gopurs_runtime.Value
__local_var_3_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_23
// TAST (Let): __local_var_3_22 -> gopurs_runtime.Value
__local_var_3_22 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_23, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_22
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 -> *Constructor_Data_Functor_Functor
functorExceptT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_6.UnsafePtr).V0})}
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right)(m_6.UnsafePtr).V0)})}
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
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})))
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_7 -> *Constructor_Control_Bind_Bind
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): pure_7_8 -> gopurs_runtime.Value
pure_7_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_10
// TAST (Let): functorExceptT1_9_9 -> *Constructor_Data_Functor_Functor
functorExceptT1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_11.UnsafePtr).V0})}
goto end_branch_12
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right)(m_11.UnsafePtr).V0)})}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
_ = __local_var_11_11
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_11, v_12)
})
})))
_ = functorExceptT1_9_9
// TAST (Let): __local_var_10_13 -> gopurs_runtime.Value
__local_var_10_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_14 -> *Constructor_Control_Bind_Bind
Bind1_11_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_14
// TAST (Let): pure_12_15 -> gopurs_runtime.Value
pure_12_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_15
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_14.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(pure_12_15, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_16
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t16 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_13
// TAST (Let): Bind1_11_17 -> *Constructor_Control_Bind_Bind
Bind1_11_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_17
// TAST (Let): Applicative0_12_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_9)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_17.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_17.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_18.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(pure_7_8, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_19
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t19 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_5_6
// TAST (Let): Bind1_6_20 -> *Constructor_Control_Bind_Bind
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): Applicative0_7_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_21
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_21.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_22, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_24 -> *Constructor_Control_Bind_Bind
Bind1_3_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_24
// TAST (Let): pure_4_25 -> gopurs_runtime.Value
pure_4_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_25
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_27 -> gopurs_runtime.Value
__local_var_6_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_27
// TAST (Let): functorExceptT1_6_26 -> *Constructor_Data_Functor_Functor
functorExceptT1_6_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_27, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_8.UnsafePtr).V0})}
goto end_branch_29
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right)(m_8.UnsafePtr).V0)})}
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
_ = __local_var_8_28
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_28, v_9)
})
})))
_ = functorExceptT1_6_26
// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
__local_var_7_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_42 -> gopurs_runtime.Value
__local_var_8_42 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_42
// TAST (Let): __local_var_8_41 -> gopurs_runtime.Value
__local_var_8_41 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_42, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_41
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_32 -> gopurs_runtime.Value
__local_var_9_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_32
// TAST (Let): functorExceptT1_9_31 -> *Constructor_Data_Functor_Functor
functorExceptT1_9_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_32, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_11.UnsafePtr).V0})}
goto end_branch_34
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right)(m_11.UnsafePtr).V0)})}
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
_ = __local_var_11_33
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_33, v_12)
})
})))
_ = functorExceptT1_9_31
// TAST (Let): __local_var_10_35 -> gopurs_runtime.Value
__local_var_10_35 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_36 -> *Constructor_Control_Bind_Bind
Bind1_11_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_36
// TAST (Let): pure_12_37 -> gopurs_runtime.Value
pure_12_37 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_37
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_36.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t38 = gopurs_runtime.Apply(pure_12_37, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_38
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t38 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_38
} else {

}
}
{
__t38 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_38:
return __t38
}))
})
})})}
}))
_ = __local_var_10_35
// TAST (Let): Bind1_11_39 -> *Constructor_Control_Bind_Bind
Bind1_11_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_35, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_39
// TAST (Let): Applicative0_12_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_35, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_40
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_31)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_39.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_39.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_40.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_41, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_43 -> *Constructor_Control_Bind_Bind
Bind1_8_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_43
// TAST (Let): pure_9_44 -> gopurs_runtime.Value
pure_9_44 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_44
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_43.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t45 = gopurs_runtime.Apply(pure_9_44, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_45
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t45 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_7_30
// TAST (Let): Bind1_8_46 -> *Constructor_Control_Bind_Bind
Bind1_8_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_46
// TAST (Let): Applicative0_9_47 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_47
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_26)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_46.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_46.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_47.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_24.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t48 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t48 = gopurs_runtime.Apply(pure_4_25, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_48
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t48 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_48
} else {

}
}
{
__t48 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_48:
return __t48
}))
})
})})}
})))
_ = monadExceptT1_1_0
// TAST (Let): __local_var_2_49 -> *Constructor_Control_Monad_Monad
__local_var_2_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_2_49
// TAST (Let): pure_3_50 -> gopurs_runtime.Value
pure_3_50 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_49.V0), gopurs_runtime.Value{}), "pure")
_ = pure_3_50
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Reader_Class_MonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_1_0)}
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_49.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_0, "ask"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_50, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_4})})
}))})}
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
// TAST (Let): monadExceptT1_2_2 -> *Constructor_Control_Monad_Monad
monadExceptT1_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_91 -> gopurs_runtime.Value
__local_var_4_91 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_91
// TAST (Let): __local_var_4_90 -> gopurs_runtime.Value
__local_var_4_90 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_91, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_90
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorExceptT1_5_4 -> *Constructor_Data_Functor_Functor
functorExceptT1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_7.UnsafePtr).V0})}
goto end_branch_7
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right)(m_7.UnsafePtr).V0)})}
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
_ = __local_var_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, v_8)
})
})))
_ = functorExceptT1_5_4
// TAST (Let): __local_var_6_8 -> gopurs_runtime.Value
__local_var_6_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
__local_var_7_30 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_30
// TAST (Let): __local_var_7_29 -> gopurs_runtime.Value
__local_var_7_29 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_30, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_8})})
})
_ = __local_var_7_29
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_10
// TAST (Let): functorExceptT1_8_9 -> *Constructor_Data_Functor_Functor
functorExceptT1_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_10.UnsafePtr).V0})}
goto end_branch_12
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right)(m_10.UnsafePtr).V0)})}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
_ = __local_var_10_11
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_11, v_11)
})
})))
_ = functorExceptT1_8_9
// TAST (Let): __local_var_9_13 -> gopurs_runtime.Value
__local_var_9_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_14 -> *Constructor_Control_Bind_Bind
Bind1_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_14
// TAST (Let): pure_11_15 -> gopurs_runtime.Value
pure_11_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_15
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_17 -> gopurs_runtime.Value
__local_var_13_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_17
// TAST (Let): functorExceptT1_13_16 -> *Constructor_Data_Functor_Functor
functorExceptT1_13_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_18 -> gopurs_runtime.Value
__local_var_15_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_17, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_15.UnsafePtr).V0})}
goto end_branch_19
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right)(m_15.UnsafePtr).V0)})}
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
_ = __local_var_15_18
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_18, v_16)
})
})))
_ = functorExceptT1_13_16
// TAST (Let): __local_var_14_20 -> gopurs_runtime.Value
__local_var_14_20 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_21 -> *Constructor_Control_Bind_Bind
Bind1_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_21
// TAST (Let): pure_16_22 -> gopurs_runtime.Value
pure_16_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_22
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_21.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t23 = gopurs_runtime.Apply(pure_16_22, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_23
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t23 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_14_20
// TAST (Let): Bind1_15_24 -> *Constructor_Control_Bind_Bind
Bind1_15_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_20, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_24
// TAST (Let): Applicative0_16_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_20, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_25
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_16)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_24.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_24.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_25.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t26 = gopurs_runtime.Apply(pure_11_15, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_26
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t26 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_9_13
// TAST (Let): Bind1_10_27 -> *Constructor_Control_Bind_Bind
Bind1_10_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_27
// TAST (Let): Applicative0_11_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_28
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_9)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_27.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_27.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_28.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_29, x_8)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_31 -> *Constructor_Control_Bind_Bind
Bind1_7_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_31
// TAST (Let): pure_8_32 -> gopurs_runtime.Value
pure_8_32 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_32
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_34 -> gopurs_runtime.Value
__local_var_10_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_34
// TAST (Let): functorExceptT1_10_33 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_33 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_35 -> gopurs_runtime.Value
__local_var_12_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_34, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_36
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
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
_ = __local_var_12_35
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_35, v_13)
})
})))
_ = functorExceptT1_10_33
// TAST (Let): __local_var_11_37 -> gopurs_runtime.Value
__local_var_11_37 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_59 -> gopurs_runtime.Value
__local_var_12_59 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_59
// TAST (Let): __local_var_12_58 -> gopurs_runtime.Value
__local_var_12_58 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_59, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_58
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_39 -> gopurs_runtime.Value
__local_var_13_39 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_39
// TAST (Let): functorExceptT1_13_38 -> *Constructor_Data_Functor_Functor
functorExceptT1_13_38 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_40 -> gopurs_runtime.Value
__local_var_15_40 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_39, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t41 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_15.UnsafePtr).V0})}
goto end_branch_41
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t41 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right)(m_15.UnsafePtr).V0)})}
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
_ = __local_var_15_40
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_40, v_16)
})
})))
_ = functorExceptT1_13_38
// TAST (Let): __local_var_14_42 -> gopurs_runtime.Value
__local_var_14_42 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_43 -> *Constructor_Control_Bind_Bind
Bind1_15_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_43
// TAST (Let): pure_16_44 -> gopurs_runtime.Value
pure_16_44 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_44
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_46 -> gopurs_runtime.Value
__local_var_18_46 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_46
// TAST (Let): functorExceptT1_18_45 -> *Constructor_Data_Functor_Functor
functorExceptT1_18_45 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_47 -> gopurs_runtime.Value
__local_var_20_47 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_46, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t48 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t48 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_20.UnsafePtr).V0})}
goto end_branch_48
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t48 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right)(m_20.UnsafePtr).V0)})}
goto end_branch_48
} else {

}
}
{
__t48 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_48:
return __t48
}))
_ = __local_var_20_47
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_47, v_21)
})
})))
_ = functorExceptT1_18_45
// TAST (Let): __local_var_19_49 -> gopurs_runtime.Value
__local_var_19_49 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_50 -> *Constructor_Control_Bind_Bind
Bind1_20_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_50
// TAST (Let): pure_21_51 -> gopurs_runtime.Value
pure_21_51 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_51
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_50.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t52 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t52 = gopurs_runtime.Apply(pure_21_51, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_52
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t52 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_19_49
// TAST (Let): Bind1_20_53 -> *Constructor_Control_Bind_Bind
Bind1_20_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_49, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_53
// TAST (Let): Applicative0_21_54 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_49, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_54
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_45)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_53.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_53.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_54.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_43.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t55 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t55 = gopurs_runtime.Apply(pure_16_44, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_55
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t55 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_14_42
// TAST (Let): Bind1_15_56 -> *Constructor_Control_Bind_Bind
Bind1_15_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_42, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_56
// TAST (Let): Applicative0_16_57 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_42, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_57
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_38)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_56.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_56.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_57.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_58, x_13)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_60 -> *Constructor_Control_Bind_Bind
Bind1_12_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_60
// TAST (Let): pure_13_61 -> gopurs_runtime.Value
pure_13_61 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_61
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_63 -> gopurs_runtime.Value
__local_var_15_63 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_63
// TAST (Let): functorExceptT1_15_62 -> *Constructor_Data_Functor_Functor
functorExceptT1_15_62 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_64 -> gopurs_runtime.Value
__local_var_17_64 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_63, "map"), gopurs_runtime.Func(func(m_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t65 gopurs_runtime.Value
{
if (m_17.Type == 9 && m_17.IntVal == 3711209382) {
__t65 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_17.UnsafePtr).V0})}
goto end_branch_65
} else {

}
}
{
if (m_17.Type == 9 && m_17.IntVal == 2465973597) {
__t65 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Either_Right)(m_17.UnsafePtr).V0)})}
goto end_branch_65
} else {

}
}
{
__t65 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_65:
return __t65
}))
_ = __local_var_17_64
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_64, v_18)
})
})))
_ = functorExceptT1_15_62
// TAST (Let): __local_var_16_66 -> gopurs_runtime.Value
__local_var_16_66 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_78 -> gopurs_runtime.Value
__local_var_17_78 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_78
// TAST (Let): __local_var_17_77 -> gopurs_runtime.Value
__local_var_17_77 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_78, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_18})})
})
_ = __local_var_17_77
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_68 -> gopurs_runtime.Value
__local_var_18_68 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_68
// TAST (Let): functorExceptT1_18_67 -> *Constructor_Data_Functor_Functor
functorExceptT1_18_67 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_69 -> gopurs_runtime.Value
__local_var_20_69 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_68, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t70 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t70 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_20.UnsafePtr).V0})}
goto end_branch_70
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t70 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right)(m_20.UnsafePtr).V0)})}
goto end_branch_70
} else {

}
}
{
__t70 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_70:
return __t70
}))
_ = __local_var_20_69
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_69, v_21)
})
})))
_ = functorExceptT1_18_67
// TAST (Let): __local_var_19_71 -> gopurs_runtime.Value
__local_var_19_71 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_72 -> *Constructor_Control_Bind_Bind
Bind1_20_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_72
// TAST (Let): pure_21_73 -> gopurs_runtime.Value
pure_21_73 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_73
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_72.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t74 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t74 = gopurs_runtime.Apply(pure_21_73, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_74
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t74 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_74
} else {

}
}
{
__t74 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_74:
return __t74
}))
})
})})}
}))
_ = __local_var_19_71
// TAST (Let): Bind1_20_75 -> *Constructor_Control_Bind_Bind
Bind1_20_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_71, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_75
// TAST (Let): Applicative0_21_76 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_71, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_76
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_67)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_75.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_75.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_76.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_77, x_18)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_79 -> *Constructor_Control_Bind_Bind
Bind1_17_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_79
// TAST (Let): pure_18_80 -> gopurs_runtime.Value
pure_18_80 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_80
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_79.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t81 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t81 = gopurs_runtime.Apply(pure_18_80, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_81
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t81 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_81
} else {

}
}
{
__t81 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_81:
return __t81
}))
})
})})}
}))
_ = __local_var_16_66
// TAST (Let): Bind1_17_82 -> *Constructor_Control_Bind_Bind
Bind1_17_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_66, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_82
// TAST (Let): Applicative0_18_83 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_83 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_66, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_83
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_15_62)}
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_82.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_82.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_83.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_60.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t84 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t84 = gopurs_runtime.Apply(pure_13_61, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_84
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t84 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_84
} else {

}
}
{
__t84 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_84:
return __t84
}))
})
})})}
}))
_ = __local_var_11_37
// TAST (Let): Bind1_12_85 -> *Constructor_Control_Bind_Bind
Bind1_12_85 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_85
// TAST (Let): Applicative0_13_86 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_86
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_33)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_85.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_85.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_86.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_31.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t87 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t87 = gopurs_runtime.Apply(pure_8_32, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_87
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t87 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_87
} else {

}
}
{
__t87 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_87:
return __t87
}))
})
})})}
}))
_ = __local_var_6_8
// TAST (Let): Bind1_7_88 -> *Constructor_Control_Bind_Bind
Bind1_7_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_88
// TAST (Let): Applicative0_8_89 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_89 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_89
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_4)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_88.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_88.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_89.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_90, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_92 -> *Constructor_Control_Bind_Bind
Bind1_4_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_92
// TAST (Let): pure_5_93 -> gopurs_runtime.Value
pure_5_93 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_93
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_95 -> gopurs_runtime.Value
__local_var_7_95 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_95
// TAST (Let): functorExceptT1_7_94 -> *Constructor_Data_Functor_Functor
functorExceptT1_7_94 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_96 -> gopurs_runtime.Value
__local_var_9_96 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_95, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t97 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t97 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_9.UnsafePtr).V0})}
goto end_branch_97
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t97 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right)(m_9.UnsafePtr).V0)})}
goto end_branch_97
} else {

}
}
{
__t97 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_97:
return __t97
}))
_ = __local_var_9_96
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_96, v_10)
})
})))
_ = functorExceptT1_7_94
// TAST (Let): __local_var_8_98 -> gopurs_runtime.Value
__local_var_8_98 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_154 -> gopurs_runtime.Value
__local_var_9_154 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_154
// TAST (Let): __local_var_9_153 -> gopurs_runtime.Value
__local_var_9_153 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_154, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_153
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_100 -> gopurs_runtime.Value
__local_var_10_100 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_100
// TAST (Let): functorExceptT1_10_99 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_99 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_101 -> gopurs_runtime.Value
__local_var_12_101 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_100, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t102 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t102 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_102
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t102 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
goto end_branch_102
} else {

}
}
{
__t102 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_102:
return __t102
}))
_ = __local_var_12_101
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_101, v_13)
})
})))
_ = functorExceptT1_10_99
// TAST (Let): __local_var_11_103 -> gopurs_runtime.Value
__local_var_11_103 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_125 -> gopurs_runtime.Value
__local_var_12_125 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_125
// TAST (Let): __local_var_12_124 -> gopurs_runtime.Value
__local_var_12_124 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_125, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_124
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_105 -> gopurs_runtime.Value
__local_var_13_105 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_105
// TAST (Let): functorExceptT1_13_104 -> *Constructor_Data_Functor_Functor
functorExceptT1_13_104 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_106 -> gopurs_runtime.Value
__local_var_15_106 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_105, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t107 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t107 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_15.UnsafePtr).V0})}
goto end_branch_107
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t107 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right)(m_15.UnsafePtr).V0)})}
goto end_branch_107
} else {

}
}
{
__t107 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_107:
return __t107
}))
_ = __local_var_15_106
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_106, v_16)
})
})))
_ = functorExceptT1_13_104
// TAST (Let): __local_var_14_108 -> gopurs_runtime.Value
__local_var_14_108 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_109 -> *Constructor_Control_Bind_Bind
Bind1_15_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_109
// TAST (Let): pure_16_110 -> gopurs_runtime.Value
pure_16_110 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_110
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_112 -> gopurs_runtime.Value
__local_var_18_112 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_112
// TAST (Let): functorExceptT1_18_111 -> *Constructor_Data_Functor_Functor
functorExceptT1_18_111 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_113 -> gopurs_runtime.Value
__local_var_20_113 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_112, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t114 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t114 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_20.UnsafePtr).V0})}
goto end_branch_114
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t114 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right)(m_20.UnsafePtr).V0)})}
goto end_branch_114
} else {

}
}
{
__t114 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_114:
return __t114
}))
_ = __local_var_20_113
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_113, v_21)
})
})))
_ = functorExceptT1_18_111
// TAST (Let): __local_var_19_115 -> gopurs_runtime.Value
__local_var_19_115 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_116 -> *Constructor_Control_Bind_Bind
Bind1_20_116 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_116
// TAST (Let): pure_21_117 -> gopurs_runtime.Value
pure_21_117 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_117
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_116.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t118 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t118 = gopurs_runtime.Apply(pure_21_117, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_118
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t118 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_118
} else {

}
}
{
__t118 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_118:
return __t118
}))
})
})})}
}))
_ = __local_var_19_115
// TAST (Let): Bind1_20_119 -> *Constructor_Control_Bind_Bind
Bind1_20_119 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_115, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_119
// TAST (Let): Applicative0_21_120 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_120 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_115, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_120
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_111)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_119.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_119.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_120.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_109.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t121 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t121 = gopurs_runtime.Apply(pure_16_110, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_121
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t121 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_121
} else {

}
}
{
__t121 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_121:
return __t121
}))
})
})})}
}))
_ = __local_var_14_108
// TAST (Let): Bind1_15_122 -> *Constructor_Control_Bind_Bind
Bind1_15_122 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_108, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_122
// TAST (Let): Applicative0_16_123 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_123 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_108, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_123
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_104)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_122.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_122.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_123.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_124, x_13)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_126 -> *Constructor_Control_Bind_Bind
Bind1_12_126 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_126
// TAST (Let): pure_13_127 -> gopurs_runtime.Value
pure_13_127 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_127
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_129 -> gopurs_runtime.Value
__local_var_15_129 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_129
// TAST (Let): functorExceptT1_15_128 -> *Constructor_Data_Functor_Functor
functorExceptT1_15_128 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_130 -> gopurs_runtime.Value
__local_var_17_130 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_129, "map"), gopurs_runtime.Func(func(m_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t131 gopurs_runtime.Value
{
if (m_17.Type == 9 && m_17.IntVal == 3711209382) {
__t131 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_17.UnsafePtr).V0})}
goto end_branch_131
} else {

}
}
{
if (m_17.Type == 9 && m_17.IntVal == 2465973597) {
__t131 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Either_Right)(m_17.UnsafePtr).V0)})}
goto end_branch_131
} else {

}
}
{
__t131 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_131:
return __t131
}))
_ = __local_var_17_130
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_130, v_18)
})
})))
_ = functorExceptT1_15_128
// TAST (Let): __local_var_16_132 -> gopurs_runtime.Value
__local_var_16_132 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_144 -> gopurs_runtime.Value
__local_var_17_144 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_144
// TAST (Let): __local_var_17_143 -> gopurs_runtime.Value
__local_var_17_143 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_144, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_18})})
})
_ = __local_var_17_143
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_134 -> gopurs_runtime.Value
__local_var_18_134 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_134
// TAST (Let): functorExceptT1_18_133 -> *Constructor_Data_Functor_Functor
functorExceptT1_18_133 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_135 -> gopurs_runtime.Value
__local_var_20_135 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_134, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t136 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t136 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_20.UnsafePtr).V0})}
goto end_branch_136
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t136 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right)(m_20.UnsafePtr).V0)})}
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
_ = __local_var_20_135
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_135, v_21)
})
})))
_ = functorExceptT1_18_133
// TAST (Let): __local_var_19_137 -> gopurs_runtime.Value
__local_var_19_137 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_138 -> *Constructor_Control_Bind_Bind
Bind1_20_138 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_138
// TAST (Let): pure_21_139 -> gopurs_runtime.Value
pure_21_139 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_139
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_138.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t140 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t140 = gopurs_runtime.Apply(pure_21_139, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_140
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t140 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_140
} else {

}
}
{
__t140 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_140:
return __t140
}))
})
})})}
}))
_ = __local_var_19_137
// TAST (Let): Bind1_20_141 -> *Constructor_Control_Bind_Bind
Bind1_20_141 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_137, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_141
// TAST (Let): Applicative0_21_142 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_142 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_137, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_142
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_133)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_141.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_141.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_142.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_143, x_18)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_145 -> *Constructor_Control_Bind_Bind
Bind1_17_145 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_145
// TAST (Let): pure_18_146 -> gopurs_runtime.Value
pure_18_146 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_146
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_145.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t147 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t147 = gopurs_runtime.Apply(pure_18_146, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_147
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t147 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_147
} else {

}
}
{
__t147 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_147:
return __t147
}))
})
})})}
}))
_ = __local_var_16_132
// TAST (Let): Bind1_17_148 -> *Constructor_Control_Bind_Bind
Bind1_17_148 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_132, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_148
// TAST (Let): Applicative0_18_149 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_149 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_132, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_149
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_15_128)}
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_148.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_148.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_149.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_126.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t150 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t150 = gopurs_runtime.Apply(pure_13_127, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_150
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t150 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_150
} else {

}
}
{
__t150 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_150:
return __t150
}))
})
})})}
}))
_ = __local_var_11_103
// TAST (Let): Bind1_12_151 -> *Constructor_Control_Bind_Bind
Bind1_12_151 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_103, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_151
// TAST (Let): Applicative0_13_152 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_152 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_103, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_152
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_99)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_151.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_151.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_152.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_153, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_155 -> *Constructor_Control_Bind_Bind
Bind1_9_155 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_155
// TAST (Let): pure_10_156 -> gopurs_runtime.Value
pure_10_156 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_156
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_158 -> gopurs_runtime.Value
__local_var_12_158 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_158
// TAST (Let): functorExceptT1_12_157 -> *Constructor_Data_Functor_Functor
functorExceptT1_12_157 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_159 -> gopurs_runtime.Value
__local_var_14_159 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_158, "map"), gopurs_runtime.Func(func(m_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t160 gopurs_runtime.Value
{
if (m_14.Type == 9 && m_14.IntVal == 3711209382) {
__t160 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_14.UnsafePtr).V0})}
goto end_branch_160
} else {

}
}
{
if (m_14.Type == 9 && m_14.IntVal == 2465973597) {
__t160 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_13, (*Constructor_Data_Either_Right)(m_14.UnsafePtr).V0)})}
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
_ = __local_var_14_159
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_159, v_15)
})
})))
_ = functorExceptT1_12_157
// TAST (Let): __local_var_13_161 -> gopurs_runtime.Value
__local_var_13_161 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_173 -> gopurs_runtime.Value
__local_var_14_173 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_14_173
// TAST (Let): __local_var_14_172 -> gopurs_runtime.Value
__local_var_14_172 := gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_173, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_15})})
})
_ = __local_var_14_172
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_163 -> gopurs_runtime.Value
__local_var_15_163 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_163
// TAST (Let): functorExceptT1_15_162 -> *Constructor_Data_Functor_Functor
functorExceptT1_15_162 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_164 -> gopurs_runtime.Value
__local_var_17_164 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_163, "map"), gopurs_runtime.Func(func(m_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t165 gopurs_runtime.Value
{
if (m_17.Type == 9 && m_17.IntVal == 3711209382) {
__t165 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_17.UnsafePtr).V0})}
goto end_branch_165
} else {

}
}
{
if (m_17.Type == 9 && m_17.IntVal == 2465973597) {
__t165 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Either_Right)(m_17.UnsafePtr).V0)})}
goto end_branch_165
} else {

}
}
{
__t165 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_165:
return __t165
}))
_ = __local_var_17_164
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_164, v_18)
})
})))
_ = functorExceptT1_15_162
// TAST (Let): __local_var_16_166 -> gopurs_runtime.Value
__local_var_16_166 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_167 -> *Constructor_Control_Bind_Bind
Bind1_17_167 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_167
// TAST (Let): pure_18_168 -> gopurs_runtime.Value
pure_18_168 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_168
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_167.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t169 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t169 = gopurs_runtime.Apply(pure_18_168, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_169
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t169 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_16_166
// TAST (Let): Bind1_17_170 -> *Constructor_Control_Bind_Bind
Bind1_17_170 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_166, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_170
// TAST (Let): Applicative0_18_171 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_171 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_166, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_171
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_15_162)}
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_170.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_170.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_171.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_172, x_15)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_14_174 -> *Constructor_Control_Bind_Bind
Bind1_14_174 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_174
// TAST (Let): pure_15_175 -> gopurs_runtime.Value
pure_15_175 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_15_175
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_174.V1), v_16, gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t176 gopurs_runtime.Value
{
if (v2_18.Type == 9 && v2_18.IntVal == 3711209382) {
__t176 = gopurs_runtime.Apply(pure_15_175, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_18.UnsafePtr).V0})})
goto end_branch_176
} else {

}
}
{
if (v2_18.Type == 9 && v2_18.IntVal == 2465973597) {
__t176 = gopurs_runtime.Apply(k_17, (*Constructor_Data_Either_Right)(v2_18.UnsafePtr).V0)
goto end_branch_176
} else {

}
}
{
__t176 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_176:
return __t176
}))
})
})})}
}))
_ = __local_var_13_161
// TAST (Let): Bind1_14_177 -> *Constructor_Control_Bind_Bind
Bind1_14_177 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_161, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_177
// TAST (Let): Applicative0_15_178 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_178 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_161, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_178
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_12_157)}
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_177.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_177.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_178.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_155.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t179 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t179 = gopurs_runtime.Apply(pure_10_156, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_179
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t179 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
goto end_branch_179
} else {

}
}
{
__t179 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_179:
return __t179
}))
})
})})}
}))
_ = __local_var_8_98
// TAST (Let): Bind1_9_180 -> *Constructor_Control_Bind_Bind
Bind1_9_180 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_98, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_180
// TAST (Let): Applicative0_10_181 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_181 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_98, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_181
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_94)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_180.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_180.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_181.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_92.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t182 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t182 = gopurs_runtime.Apply(pure_5_93, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_182
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t182 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
goto end_branch_182
} else {

}
}
{
__t182 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_182:
return __t182
}))
})
})})}
})))
_ = monadExceptT1_2_2
// TAST (Let): __local_var_3_183 -> *Constructor_Control_Monad_Monad
__local_var_3_183 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_183
// TAST (Let): pure_4_184 -> gopurs_runtime.Value
pure_4_184 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_183.V0), gopurs_runtime.Value{}), "pure")
_ = pure_4_184
// TAST (Let): monadAskExceptT1_1_0 -> *Constructor_Control_Monad_Reader_Class_MonadAsk
monadAskExceptT1_1_0 := &Constructor_Control_Monad_Reader_Class_MonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_2_2)}
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_183.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(__local_var_1_1, "ask"), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_184, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_5})})
}))}
_ = monadAskExceptT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Reader_Class_MonadReader{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskExceptT1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_185 -> gopurs_runtime.Value
__local_var_3_185 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_2)
_ = __local_var_3_185
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_185, v_4)
})
})})}
}

func Call_Control_Monad_Except_Trans_monadContExceptT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadExceptT1_1_0 -> *Constructor_Control_Monad_Monad
monadExceptT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_23 -> gopurs_runtime.Value
__local_var_3_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_23
// TAST (Let): __local_var_3_22 -> gopurs_runtime.Value
__local_var_3_22 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_23, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_22
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 -> *Constructor_Data_Functor_Functor
functorExceptT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_6.UnsafePtr).V0})}
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right)(m_6.UnsafePtr).V0)})}
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
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})))
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_7 -> *Constructor_Control_Bind_Bind
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): pure_7_8 -> gopurs_runtime.Value
pure_7_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_10
// TAST (Let): functorExceptT1_9_9 -> *Constructor_Data_Functor_Functor
functorExceptT1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_11.UnsafePtr).V0})}
goto end_branch_12
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right)(m_11.UnsafePtr).V0)})}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
_ = __local_var_11_11
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_11, v_12)
})
})))
_ = functorExceptT1_9_9
// TAST (Let): __local_var_10_13 -> gopurs_runtime.Value
__local_var_10_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_14 -> *Constructor_Control_Bind_Bind
Bind1_11_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_14
// TAST (Let): pure_12_15 -> gopurs_runtime.Value
pure_12_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_15
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_14.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(pure_12_15, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_16
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t16 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_13
// TAST (Let): Bind1_11_17 -> *Constructor_Control_Bind_Bind
Bind1_11_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_17
// TAST (Let): Applicative0_12_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_9)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_17.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_17.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_18.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(pure_7_8, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_19
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t19 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_5_6
// TAST (Let): Bind1_6_20 -> *Constructor_Control_Bind_Bind
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): Applicative0_7_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_21
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_21.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_22, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_24 -> *Constructor_Control_Bind_Bind
Bind1_3_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_24
// TAST (Let): pure_4_25 -> gopurs_runtime.Value
pure_4_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_25
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_27 -> gopurs_runtime.Value
__local_var_6_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_27
// TAST (Let): functorExceptT1_6_26 -> *Constructor_Data_Functor_Functor
functorExceptT1_6_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_27, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_8.UnsafePtr).V0})}
goto end_branch_29
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right)(m_8.UnsafePtr).V0)})}
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
_ = __local_var_8_28
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_28, v_9)
})
})))
_ = functorExceptT1_6_26
// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
__local_var_7_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_42 -> gopurs_runtime.Value
__local_var_8_42 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_42
// TAST (Let): __local_var_8_41 -> gopurs_runtime.Value
__local_var_8_41 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_42, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_41
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_32 -> gopurs_runtime.Value
__local_var_9_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_32
// TAST (Let): functorExceptT1_9_31 -> *Constructor_Data_Functor_Functor
functorExceptT1_9_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_32, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_11.UnsafePtr).V0})}
goto end_branch_34
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right)(m_11.UnsafePtr).V0)})}
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
_ = __local_var_11_33
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_33, v_12)
})
})))
_ = functorExceptT1_9_31
// TAST (Let): __local_var_10_35 -> gopurs_runtime.Value
__local_var_10_35 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_36 -> *Constructor_Control_Bind_Bind
Bind1_11_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_36
// TAST (Let): pure_12_37 -> gopurs_runtime.Value
pure_12_37 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_37
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_36.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t38 = gopurs_runtime.Apply(pure_12_37, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_38
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t38 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_38
} else {

}
}
{
__t38 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_38:
return __t38
}))
})
})})}
}))
_ = __local_var_10_35
// TAST (Let): Bind1_11_39 -> *Constructor_Control_Bind_Bind
Bind1_11_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_35, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_39
// TAST (Let): Applicative0_12_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_35, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_40
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_31)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_39.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_39.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_40.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_41, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_43 -> *Constructor_Control_Bind_Bind
Bind1_8_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_43
// TAST (Let): pure_9_44 -> gopurs_runtime.Value
pure_9_44 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_44
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_43.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t45 = gopurs_runtime.Apply(pure_9_44, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_45
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t45 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_7_30
// TAST (Let): Bind1_8_46 -> *Constructor_Control_Bind_Bind
Bind1_8_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_46
// TAST (Let): Applicative0_9_47 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_47
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_26)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_46.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_46.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_47.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_24.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t48 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t48 = gopurs_runtime.Apply(pure_4_25, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_48
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t48 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_48
} else {

}
}
{
__t48 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_48:
return __t48
}))
})
})})}
})))
_ = monadExceptT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Cont_Class_MonadCont{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_3, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_4})})
}))
}))
})})}
}

func Call_Control_Monad_Except_Trans_monadEffectExceptT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadExceptT1_2_1 -> *Constructor_Control_Monad_Monad
monadExceptT1_2_1 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_23 -> gopurs_runtime.Value
__local_var_3_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_23
// TAST (Let): __local_var_3_22 -> gopurs_runtime.Value
__local_var_3_22 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_23, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_22
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 -> *Constructor_Data_Functor_Functor
functorExceptT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_6.UnsafePtr).V0})}
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right)(m_6.UnsafePtr).V0)})}
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
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})))
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_7 -> *Constructor_Control_Bind_Bind
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): pure_7_8 -> gopurs_runtime.Value
pure_7_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_10
// TAST (Let): functorExceptT1_9_9 -> *Constructor_Data_Functor_Functor
functorExceptT1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_11.UnsafePtr).V0})}
goto end_branch_12
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right)(m_11.UnsafePtr).V0)})}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
_ = __local_var_11_11
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_11, v_12)
})
})))
_ = functorExceptT1_9_9
// TAST (Let): __local_var_10_13 -> gopurs_runtime.Value
__local_var_10_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_14 -> *Constructor_Control_Bind_Bind
Bind1_11_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_14
// TAST (Let): pure_12_15 -> gopurs_runtime.Value
pure_12_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_15
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_14.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(pure_12_15, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_16
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t16 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_13
// TAST (Let): Bind1_11_17 -> *Constructor_Control_Bind_Bind
Bind1_11_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_17
// TAST (Let): Applicative0_12_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_9)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_17.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_17.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_18.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(pure_7_8, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_19
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t19 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_5_6
// TAST (Let): Bind1_6_20 -> *Constructor_Control_Bind_Bind
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): Applicative0_7_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_21
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_21.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_22, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_24 -> *Constructor_Control_Bind_Bind
Bind1_3_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_24
// TAST (Let): pure_4_25 -> gopurs_runtime.Value
pure_4_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_25
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_27 -> gopurs_runtime.Value
__local_var_6_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_27
// TAST (Let): functorExceptT1_6_26 -> *Constructor_Data_Functor_Functor
functorExceptT1_6_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_27, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_8.UnsafePtr).V0})}
goto end_branch_29
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right)(m_8.UnsafePtr).V0)})}
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
_ = __local_var_8_28
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_28, v_9)
})
})))
_ = functorExceptT1_6_26
// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
__local_var_7_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_42 -> gopurs_runtime.Value
__local_var_8_42 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_42
// TAST (Let): __local_var_8_41 -> gopurs_runtime.Value
__local_var_8_41 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_42, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_41
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_32 -> gopurs_runtime.Value
__local_var_9_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_32
// TAST (Let): functorExceptT1_9_31 -> *Constructor_Data_Functor_Functor
functorExceptT1_9_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_32, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_11.UnsafePtr).V0})}
goto end_branch_34
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right)(m_11.UnsafePtr).V0)})}
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
_ = __local_var_11_33
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_33, v_12)
})
})))
_ = functorExceptT1_9_31
// TAST (Let): __local_var_10_35 -> gopurs_runtime.Value
__local_var_10_35 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_36 -> *Constructor_Control_Bind_Bind
Bind1_11_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_36
// TAST (Let): pure_12_37 -> gopurs_runtime.Value
pure_12_37 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_37
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_36.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t38 = gopurs_runtime.Apply(pure_12_37, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_38
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t38 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_38
} else {

}
}
{
__t38 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_38:
return __t38
}))
})
})})}
}))
_ = __local_var_10_35
// TAST (Let): Bind1_11_39 -> *Constructor_Control_Bind_Bind
Bind1_11_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_35, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_39
// TAST (Let): Applicative0_12_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_35, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_40
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_31)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_39.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_39.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_40.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_41, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_43 -> *Constructor_Control_Bind_Bind
Bind1_8_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_43
// TAST (Let): pure_9_44 -> gopurs_runtime.Value
pure_9_44 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_44
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_43.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t45 = gopurs_runtime.Apply(pure_9_44, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_45
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t45 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_7_30
// TAST (Let): Bind1_8_46 -> *Constructor_Control_Bind_Bind
Bind1_8_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_46
// TAST (Let): Applicative0_9_47 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_47
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_26)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_46.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_46.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_47.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_24.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t48 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t48 = gopurs_runtime.Apply(pure_4_25, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_48
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t48 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_48
} else {

}
}
{
__t48 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_48:
return __t48
}))
})
})})}
})}
_ = monadExceptT1_2_1
// TAST (Let): Bind1_3_50 -> *Constructor_Control_Bind_Bind
Bind1_3_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_50
// TAST (Let): pure_4_51 -> gopurs_runtime.Value
pure_4_51 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_51
// TAST (Let): __local_var_3_49 -> gopurs_runtime.Value
__local_var_3_49 := gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_50.V1), m_5, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_51, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_6})})
}))
})
_ = __local_var_3_49
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(&Constructor_Effect_Class_MonadEffect{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_2_1)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_49, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
})})}
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
// TAST (Let): monadExceptT1_4_3 -> *Constructor_Control_Monad_Monad
monadExceptT1_4_3 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_25 -> gopurs_runtime.Value
__local_var_5_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_25
// TAST (Let): __local_var_5_24 -> gopurs_runtime.Value
__local_var_5_24 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_25, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_6})})
})
_ = __local_var_5_24
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): functorExceptT1_6_4 -> *Constructor_Data_Functor_Functor
functorExceptT1_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_8.UnsafePtr).V0})}
goto end_branch_7
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right)(m_8.UnsafePtr).V0)})}
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
_ = __local_var_8_6
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_6, v_9)
})
})))
_ = functorExceptT1_6_4
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_9 -> *Constructor_Control_Bind_Bind
Bind1_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_9
// TAST (Let): pure_9_10 -> gopurs_runtime.Value
pure_9_10 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_10
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_12 -> gopurs_runtime.Value
__local_var_11_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_12
// TAST (Let): functorExceptT1_11_11 -> *Constructor_Data_Functor_Functor
functorExceptT1_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_13 -> gopurs_runtime.Value
__local_var_13_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "map"), gopurs_runtime.Func(func(m_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (m_13.Type == 9 && m_13.IntVal == 3711209382) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_13.UnsafePtr).V0})}
goto end_branch_14
} else {

}
}
{
if (m_13.Type == 9 && m_13.IntVal == 2465973597) {
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Either_Right)(m_13.UnsafePtr).V0)})}
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
_ = __local_var_13_13
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_13, v_14)
})
})))
_ = functorExceptT1_11_11
// TAST (Let): __local_var_12_15 -> gopurs_runtime.Value
__local_var_12_15 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_16 -> *Constructor_Control_Bind_Bind
Bind1_13_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_16
// TAST (Let): pure_14_17 -> gopurs_runtime.Value
pure_14_17 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_17
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_16.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_14_17, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_18
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_12_15
// TAST (Let): Bind1_13_19 -> *Constructor_Control_Bind_Bind
Bind1_13_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_19
// TAST (Let): Applicative0_14_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_20
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_11_11)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_19.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_19.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_20.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_9.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t21 = gopurs_runtime.Apply(pure_9_10, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_21
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t21 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
goto end_branch_21
} else {

}
}
{
__t21 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_21:
return __t21
}))
})
})})}
}))
_ = __local_var_7_8
// TAST (Let): Bind1_8_22 -> *Constructor_Control_Bind_Bind
Bind1_8_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_22
// TAST (Let): Applicative0_9_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_23
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_4)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_22.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_22.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_23.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_24, x_6)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_26 -> *Constructor_Control_Bind_Bind
Bind1_5_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_26
// TAST (Let): pure_6_27 -> gopurs_runtime.Value
pure_6_27 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_27
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_29 -> gopurs_runtime.Value
__local_var_8_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_29
// TAST (Let): functorExceptT1_8_28 -> *Constructor_Data_Functor_Functor
functorExceptT1_8_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_30 -> gopurs_runtime.Value
__local_var_10_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_29, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t31 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_10.UnsafePtr).V0})}
goto end_branch_31
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t31 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right)(m_10.UnsafePtr).V0)})}
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
_ = __local_var_10_30
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_30, v_11)
})
})))
_ = functorExceptT1_8_28
// TAST (Let): __local_var_9_32 -> gopurs_runtime.Value
__local_var_9_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_44 -> gopurs_runtime.Value
__local_var_10_44 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_44
// TAST (Let): __local_var_10_43 -> gopurs_runtime.Value
__local_var_10_43 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_44, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_11})})
})
_ = __local_var_10_43
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_34 -> gopurs_runtime.Value
__local_var_11_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_34
// TAST (Let): functorExceptT1_11_33 -> *Constructor_Data_Functor_Functor
functorExceptT1_11_33 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_35 -> gopurs_runtime.Value
__local_var_13_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_34, "map"), gopurs_runtime.Func(func(m_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (m_13.Type == 9 && m_13.IntVal == 3711209382) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_13.UnsafePtr).V0})}
goto end_branch_36
} else {

}
}
{
if (m_13.Type == 9 && m_13.IntVal == 2465973597) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Either_Right)(m_13.UnsafePtr).V0)})}
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
_ = __local_var_13_35
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_35, v_14)
})
})))
_ = functorExceptT1_11_33
// TAST (Let): __local_var_12_37 -> gopurs_runtime.Value
__local_var_12_37 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_38 -> *Constructor_Control_Bind_Bind
Bind1_13_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_38
// TAST (Let): pure_14_39 -> gopurs_runtime.Value
pure_14_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_39
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_38.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t40 = gopurs_runtime.Apply(pure_14_39, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_40
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t40 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_12_37
// TAST (Let): Bind1_13_41 -> *Constructor_Control_Bind_Bind
Bind1_13_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_41
// TAST (Let): Applicative0_14_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_42
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_11_33)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_41.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_41.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_42.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_43, x_11)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_45 -> *Constructor_Control_Bind_Bind
Bind1_10_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_45
// TAST (Let): pure_11_46 -> gopurs_runtime.Value
pure_11_46 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_46
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_45.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t47 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t47 = gopurs_runtime.Apply(pure_11_46, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_47
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t47 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
return __t47
}))
})
})})}
}))
_ = __local_var_9_32
// TAST (Let): Bind1_10_48 -> *Constructor_Control_Bind_Bind
Bind1_10_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_48
// TAST (Let): Applicative0_11_49 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_49
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_28)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_48.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_48.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_49.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_26.V1), v_7, gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t50 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 3711209382) {
__t50 = gopurs_runtime.Apply(pure_6_27, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_9.UnsafePtr).V0})})
goto end_branch_50
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 2465973597) {
__t50 = gopurs_runtime.Apply(k_8, (*Constructor_Data_Either_Right)(v2_9.UnsafePtr).V0)
goto end_branch_50
} else {

}
}
{
__t50 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_50:
return __t50
}))
})
})})}
})}
_ = monadExceptT1_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_MonadRec{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_4_3)}
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_51 -> gopurs_runtime.Value
__local_var_6_51 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(f_5, a_6), gopurs_runtime.Func(func(m_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t55 *Constructor_Control_Monad_Rec_Class_Done
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 3711209382) {
__t55 = &Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_prime_7.UnsafePtr).V0})}}
goto end_branch_55
} else {

}
}
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 2465973597) {
var __t54 gopurs_runtime.Value
{
var __t_tag_52 gopurs_runtime.Value = (*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0
if (__t_tag_52.Type == 9 && __t_tag_52.IntVal == 525585346) {
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}
goto end_branch_54
} else {

}
}
{
var __t_tag_53 gopurs_runtime.Value = (*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0
if (__t_tag_53.Type == 9 && __t_tag_53.IntVal == 60402430) {
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Either_Right)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_54
} else {

}
}
{
__t54 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_54:
__t55 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](__t54)
goto end_branch_55
} else {

}
}
{
__t55 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_55:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(__t55)})
}))
}))
_ = __local_var_6_51
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_51, x_7)
})
})})}
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
// TAST (Let): monadExceptT1_2_1 -> *Constructor_Control_Monad_Monad
monadExceptT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_24 -> gopurs_runtime.Value
__local_var_4_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_24
// TAST (Let): __local_var_4_23 -> gopurs_runtime.Value
__local_var_4_23 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_24, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_23
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorExceptT1_5_3 -> *Constructor_Data_Functor_Functor
functorExceptT1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_7.UnsafePtr).V0})}
goto end_branch_6
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right)(m_7.UnsafePtr).V0)})}
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
_ = __local_var_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, v_8)
})
})))
_ = functorExceptT1_5_3
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_8 -> *Constructor_Control_Bind_Bind
Bind1_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_8
// TAST (Let): pure_8_9 -> gopurs_runtime.Value
pure_8_9 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_9
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_11
// TAST (Let): functorExceptT1_10_10 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_12 -> gopurs_runtime.Value
__local_var_12_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_13
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
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
_ = __local_var_12_12
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_12, v_13)
})
})))
_ = functorExceptT1_10_10
// TAST (Let): __local_var_11_14 -> gopurs_runtime.Value
__local_var_11_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_15 -> *Constructor_Control_Bind_Bind
Bind1_12_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_15
// TAST (Let): pure_13_16 -> gopurs_runtime.Value
pure_13_16 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_16
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_15.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(pure_13_16, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_17
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_14
// TAST (Let): Bind1_12_18 -> *Constructor_Control_Bind_Bind
Bind1_12_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_18
// TAST (Let): Applicative0_13_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_10)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_18.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_18.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_19.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_8.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(pure_8_9, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_20
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t20 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_6_7
// TAST (Let): Bind1_7_21 -> *Constructor_Control_Bind_Bind
Bind1_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_21
// TAST (Let): Applicative0_8_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_22
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_3)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_21.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_21.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_22.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_23, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_25 -> *Constructor_Control_Bind_Bind
Bind1_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_25
// TAST (Let): pure_5_26 -> gopurs_runtime.Value
pure_5_26 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_26
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_28 -> gopurs_runtime.Value
__local_var_7_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_28
// TAST (Let): functorExceptT1_7_27 -> *Constructor_Data_Functor_Functor
functorExceptT1_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_28, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t30 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_9.UnsafePtr).V0})}
goto end_branch_30
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t30 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right)(m_9.UnsafePtr).V0)})}
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
return __t30
}))
_ = __local_var_9_29
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_29, v_10)
})
})))
_ = functorExceptT1_7_27
// TAST (Let): __local_var_8_31 -> gopurs_runtime.Value
__local_var_8_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_43 -> gopurs_runtime.Value
__local_var_9_43 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_43
// TAST (Let): __local_var_9_42 -> gopurs_runtime.Value
__local_var_9_42 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_43, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_42
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_33 -> gopurs_runtime.Value
__local_var_10_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_33
// TAST (Let): functorExceptT1_10_32 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_32 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_34 -> gopurs_runtime.Value
__local_var_12_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_33, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_35
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
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
_ = __local_var_12_34
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_34, v_13)
})
})))
_ = functorExceptT1_10_32
// TAST (Let): __local_var_11_36 -> gopurs_runtime.Value
__local_var_11_36 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_37 -> *Constructor_Control_Bind_Bind
Bind1_12_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_37
// TAST (Let): pure_13_38 -> gopurs_runtime.Value
pure_13_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_38
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_37.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t39 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t39 = gopurs_runtime.Apply(pure_13_38, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_39
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t39 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_39
} else {

}
}
{
__t39 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_39:
return __t39
}))
})
})})}
}))
_ = __local_var_11_36
// TAST (Let): Bind1_12_40 -> *Constructor_Control_Bind_Bind
Bind1_12_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_36, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_40
// TAST (Let): Applicative0_13_41 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_36, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_41
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_32)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_40.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_40.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_41.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_42, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_44 -> *Constructor_Control_Bind_Bind
Bind1_9_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_44
// TAST (Let): pure_10_45 -> gopurs_runtime.Value
pure_10_45 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_45
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_44.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t46 = gopurs_runtime.Apply(pure_10_45, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_46
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t46 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_8_31
// TAST (Let): Bind1_9_47 -> *Constructor_Control_Bind_Bind
Bind1_9_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_47
// TAST (Let): Applicative0_10_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_48
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_27)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_47.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_47.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_48.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_25.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t49 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t49 = gopurs_runtime.Apply(pure_5_26, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_49
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t49 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
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
})})}
})))
_ = monadExceptT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_State_Class_MonadState{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_2_1)}
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): pure_4_50 -> gopurs_runtime.Value
pure_4_50 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V0), gopurs_runtime.Value{}), "pure")
_ = pure_4_50
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), f_3), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_50, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_5})})
}))
})})}
}

func Call_Control_Monad_Except_Trans_monadTellExceptT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
// TAST (Let): Monad1_1_0 -> gopurs_runtime.Value
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
// TAST (Let): monadExceptT1_3_2 -> *Constructor_Control_Monad_Monad
monadExceptT1_3_2 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_24 -> gopurs_runtime.Value
__local_var_4_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_24
// TAST (Let): __local_var_4_23 -> gopurs_runtime.Value
__local_var_4_23 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_24, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_23
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorExceptT1_5_3 -> *Constructor_Data_Functor_Functor
functorExceptT1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_7.UnsafePtr).V0})}
goto end_branch_6
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right)(m_7.UnsafePtr).V0)})}
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
_ = __local_var_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, v_8)
})
})))
_ = functorExceptT1_5_3
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_8 -> *Constructor_Control_Bind_Bind
Bind1_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_8
// TAST (Let): pure_8_9 -> gopurs_runtime.Value
pure_8_9 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_9
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_11
// TAST (Let): functorExceptT1_10_10 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_12 -> gopurs_runtime.Value
__local_var_12_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_13
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
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
_ = __local_var_12_12
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_12, v_13)
})
})))
_ = functorExceptT1_10_10
// TAST (Let): __local_var_11_14 -> gopurs_runtime.Value
__local_var_11_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_15 -> *Constructor_Control_Bind_Bind
Bind1_12_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_15
// TAST (Let): pure_13_16 -> gopurs_runtime.Value
pure_13_16 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_16
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_15.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(pure_13_16, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_17
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_14
// TAST (Let): Bind1_12_18 -> *Constructor_Control_Bind_Bind
Bind1_12_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_18
// TAST (Let): Applicative0_13_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_10)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_18.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_18.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_19.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_8.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(pure_8_9, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_20
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t20 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_6_7
// TAST (Let): Bind1_7_21 -> *Constructor_Control_Bind_Bind
Bind1_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_21
// TAST (Let): Applicative0_8_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_22
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_3)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_21.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_21.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_22.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_23, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_25 -> *Constructor_Control_Bind_Bind
Bind1_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_25
// TAST (Let): pure_5_26 -> gopurs_runtime.Value
pure_5_26 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_26
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_28 -> gopurs_runtime.Value
__local_var_7_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_28
// TAST (Let): functorExceptT1_7_27 -> *Constructor_Data_Functor_Functor
functorExceptT1_7_27 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_28, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t30 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_9.UnsafePtr).V0})}
goto end_branch_30
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t30 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right)(m_9.UnsafePtr).V0)})}
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
return __t30
}))
_ = __local_var_9_29
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_29, v_10)
})
})))
_ = functorExceptT1_7_27
// TAST (Let): __local_var_8_31 -> gopurs_runtime.Value
__local_var_8_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_43 -> gopurs_runtime.Value
__local_var_9_43 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_43
// TAST (Let): __local_var_9_42 -> gopurs_runtime.Value
__local_var_9_42 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_43, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_42
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_33 -> gopurs_runtime.Value
__local_var_10_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_33
// TAST (Let): functorExceptT1_10_32 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_32 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_34 -> gopurs_runtime.Value
__local_var_12_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_33, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_35
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
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
_ = __local_var_12_34
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_34, v_13)
})
})))
_ = functorExceptT1_10_32
// TAST (Let): __local_var_11_36 -> gopurs_runtime.Value
__local_var_11_36 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_37 -> *Constructor_Control_Bind_Bind
Bind1_12_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_37
// TAST (Let): pure_13_38 -> gopurs_runtime.Value
pure_13_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_38
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_37.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t39 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t39 = gopurs_runtime.Apply(pure_13_38, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_39
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t39 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_39
} else {

}
}
{
__t39 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_39:
return __t39
}))
})
})})}
}))
_ = __local_var_11_36
// TAST (Let): Bind1_12_40 -> *Constructor_Control_Bind_Bind
Bind1_12_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_36, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_40
// TAST (Let): Applicative0_13_41 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_36, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_41
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_32)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_40.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_40.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_41.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_42, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_44 -> *Constructor_Control_Bind_Bind
Bind1_9_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_44
// TAST (Let): pure_10_45 -> gopurs_runtime.Value
pure_10_45 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_45
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_44.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t46 = gopurs_runtime.Apply(pure_10_45, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_46
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t46 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_8_31
// TAST (Let): Bind1_9_47 -> *Constructor_Control_Bind_Bind
Bind1_9_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_47
// TAST (Let): Applicative0_10_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_48
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_27)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_47.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_47.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_48.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_25.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t49 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t49 = gopurs_runtime.Apply(pure_5_26, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_49
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t49 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
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
})})}
})}
_ = monadExceptT1_3_2
// TAST (Let): Bind1_4_51 -> *Constructor_Control_Bind_Bind
Bind1_4_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_51
// TAST (Let): pure_5_52 -> gopurs_runtime.Value
pure_5_52 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_52
// TAST (Let): __local_var_4_50 -> gopurs_runtime.Value
__local_var_4_50 := gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_51.V1), m_6, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_52, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_7})})
}))
})
_ = __local_var_4_50
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Writer_Class_MonadTell{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_3_2)}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_2_1)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_50, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
})})}
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
// TAST (Let): Monoid0_6_5 -> *Constructor_Data_Monoid_Monoid
Monoid0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{}))
_ = Monoid0_6_5
// TAST (Let): Monad1_7_7 -> gopurs_runtime.Value
Monad1_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_7_7
// TAST (Let): Semigroup0_8_8 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_8_8
// TAST (Let): monadExceptT1_9_9 -> *Constructor_Control_Monad_Monad
monadExceptT1_9_9 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_97 -> gopurs_runtime.Value
__local_var_10_97 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_97
// TAST (Let): __local_var_10_96 -> gopurs_runtime.Value
__local_var_10_96 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_97, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_11})})
})
_ = __local_var_10_96
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorExceptT1_11_10 -> *Constructor_Data_Functor_Functor
functorExceptT1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_12 -> gopurs_runtime.Value
__local_var_13_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Func(func(m_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (m_13.Type == 9 && m_13.IntVal == 3711209382) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_13.UnsafePtr).V0})}
goto end_branch_13
} else {

}
}
{
if (m_13.Type == 9 && m_13.IntVal == 2465973597) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Either_Right)(m_13.UnsafePtr).V0)})}
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
_ = __local_var_13_12
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_12, v_14)
})
})))
_ = functorExceptT1_11_10
// TAST (Let): __local_var_12_14 -> gopurs_runtime.Value
__local_var_12_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_36 -> gopurs_runtime.Value
__local_var_13_36 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_36
// TAST (Let): __local_var_13_35 -> gopurs_runtime.Value
__local_var_13_35 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_36, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_14})})
})
_ = __local_var_13_35
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_16 -> gopurs_runtime.Value
__local_var_14_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_16
// TAST (Let): functorExceptT1_14_15 -> *Constructor_Data_Functor_Functor
functorExceptT1_14_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_17 -> gopurs_runtime.Value
__local_var_16_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_16, "map"), gopurs_runtime.Func(func(m_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (m_16.Type == 9 && m_16.IntVal == 3711209382) {
__t18 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_16.UnsafePtr).V0})}
goto end_branch_18
} else {

}
}
{
if (m_16.Type == 9 && m_16.IntVal == 2465973597) {
__t18 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Either_Right)(m_16.UnsafePtr).V0)})}
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
_ = __local_var_16_17
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_17, v_17)
})
})))
_ = functorExceptT1_14_15
// TAST (Let): __local_var_15_19 -> gopurs_runtime.Value
__local_var_15_19 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_20 -> *Constructor_Control_Bind_Bind
Bind1_16_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_20
// TAST (Let): pure_17_21 -> gopurs_runtime.Value
pure_17_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_21
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_23 -> gopurs_runtime.Value
__local_var_19_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_23
// TAST (Let): functorExceptT1_19_22 -> *Constructor_Data_Functor_Functor
functorExceptT1_19_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_24 -> gopurs_runtime.Value
__local_var_21_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_23, "map"), gopurs_runtime.Func(func(m_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 gopurs_runtime.Value
{
if (m_21.Type == 9 && m_21.IntVal == 3711209382) {
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_21.UnsafePtr).V0})}
goto end_branch_25
} else {

}
}
{
if (m_21.Type == 9 && m_21.IntVal == 2465973597) {
__t25 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Either_Right)(m_21.UnsafePtr).V0)})}
goto end_branch_25
} else {

}
}
{
__t25 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_25:
return __t25
}))
_ = __local_var_21_24
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_24, v_22)
})
})))
_ = functorExceptT1_19_22
// TAST (Let): __local_var_20_26 -> gopurs_runtime.Value
__local_var_20_26 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_27 -> *Constructor_Control_Bind_Bind
Bind1_21_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_27
// TAST (Let): pure_22_28 -> gopurs_runtime.Value
pure_22_28 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_28
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_27.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t29 = gopurs_runtime.Apply(pure_22_28, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_29
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t29 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_20_26
// TAST (Let): Bind1_21_30 -> *Constructor_Control_Bind_Bind
Bind1_21_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_26, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_30
// TAST (Let): Applicative0_22_31 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_26, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_31
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_19_22)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_30.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_30.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_31.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_20.V1), v_18, gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 gopurs_runtime.Value
{
if (v2_20.Type == 9 && v2_20.IntVal == 3711209382) {
__t32 = gopurs_runtime.Apply(pure_17_21, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_20.UnsafePtr).V0})})
goto end_branch_32
} else {

}
}
{
if (v2_20.Type == 9 && v2_20.IntVal == 2465973597) {
__t32 = gopurs_runtime.Apply(k_19, (*Constructor_Data_Either_Right)(v2_20.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_15_19
// TAST (Let): Bind1_16_33 -> *Constructor_Control_Bind_Bind
Bind1_16_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_19, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_33
// TAST (Let): Applicative0_17_34 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_19, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_34
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_14_15)}
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_33.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_33.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_34.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_35, x_14)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_37 -> *Constructor_Control_Bind_Bind
Bind1_13_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_37
// TAST (Let): pure_14_38 -> gopurs_runtime.Value
pure_14_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_38
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_40 -> gopurs_runtime.Value
__local_var_16_40 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_40
// TAST (Let): functorExceptT1_16_39 -> *Constructor_Data_Functor_Functor
functorExceptT1_16_39 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_41 -> gopurs_runtime.Value
__local_var_18_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_40, "map"), gopurs_runtime.Func(func(m_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (m_18.Type == 9 && m_18.IntVal == 3711209382) {
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_18.UnsafePtr).V0})}
goto end_branch_42
} else {

}
}
{
if (m_18.Type == 9 && m_18.IntVal == 2465973597) {
__t42 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Either_Right)(m_18.UnsafePtr).V0)})}
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
_ = __local_var_18_41
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_41, v_19)
})
})))
_ = functorExceptT1_16_39
// TAST (Let): __local_var_17_43 -> gopurs_runtime.Value
__local_var_17_43 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_65 -> gopurs_runtime.Value
__local_var_18_65 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_65
// TAST (Let): __local_var_18_64 -> gopurs_runtime.Value
__local_var_18_64 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_65, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_19})})
})
_ = __local_var_18_64
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_45 -> gopurs_runtime.Value
__local_var_19_45 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_45
// TAST (Let): functorExceptT1_19_44 -> *Constructor_Data_Functor_Functor
functorExceptT1_19_44 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_46 -> gopurs_runtime.Value
__local_var_21_46 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_45, "map"), gopurs_runtime.Func(func(m_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t47 gopurs_runtime.Value
{
if (m_21.Type == 9 && m_21.IntVal == 3711209382) {
__t47 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_21.UnsafePtr).V0})}
goto end_branch_47
} else {

}
}
{
if (m_21.Type == 9 && m_21.IntVal == 2465973597) {
__t47 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Either_Right)(m_21.UnsafePtr).V0)})}
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
return __t47
}))
_ = __local_var_21_46
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_46, v_22)
})
})))
_ = functorExceptT1_19_44
// TAST (Let): __local_var_20_48 -> gopurs_runtime.Value
__local_var_20_48 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_49 -> *Constructor_Control_Bind_Bind
Bind1_21_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_49
// TAST (Let): pure_22_50 -> gopurs_runtime.Value
pure_22_50 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_50
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_52 -> gopurs_runtime.Value
__local_var_24_52 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_52
// TAST (Let): functorExceptT1_24_51 -> *Constructor_Data_Functor_Functor
functorExceptT1_24_51 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_26_53 -> gopurs_runtime.Value
__local_var_26_53 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_52, "map"), gopurs_runtime.Func(func(m_26 gopurs_runtime.Value) gopurs_runtime.Value {
var __t54 gopurs_runtime.Value
{
if (m_26.Type == 9 && m_26.IntVal == 3711209382) {
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_26.UnsafePtr).V0})}
goto end_branch_54
} else {

}
}
{
if (m_26.Type == 9 && m_26.IntVal == 2465973597) {
__t54 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_25, (*Constructor_Data_Either_Right)(m_26.UnsafePtr).V0)})}
goto end_branch_54
} else {

}
}
{
__t54 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_54:
return __t54
}))
_ = __local_var_26_53
return gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_26_53, v_27)
})
})))
_ = functorExceptT1_24_51
// TAST (Let): __local_var_25_55 -> gopurs_runtime.Value
__local_var_25_55 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_56 -> *Constructor_Control_Bind_Bind
Bind1_26_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_56
// TAST (Let): pure_27_57 -> gopurs_runtime.Value
pure_27_57 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_27_57
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_56.V1), v_28, gopurs_runtime.Func(func(v2_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t58 gopurs_runtime.Value
{
if (v2_30.Type == 9 && v2_30.IntVal == 3711209382) {
__t58 = gopurs_runtime.Apply(pure_27_57, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_30.UnsafePtr).V0})})
goto end_branch_58
} else {

}
}
{
if (v2_30.Type == 9 && v2_30.IntVal == 2465973597) {
__t58 = gopurs_runtime.Apply(k_29, (*Constructor_Data_Either_Right)(v2_30.UnsafePtr).V0)
goto end_branch_58
} else {

}
}
{
__t58 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_58:
return __t58
}))
})
})})}
}))
_ = __local_var_25_55
// TAST (Let): Bind1_26_59 -> *Constructor_Control_Bind_Bind
Bind1_26_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_55, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_59
// TAST (Let): Applicative0_27_60 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_55, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_60
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_24_51)}
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_59.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_59.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_60.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_49.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t61 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t61 = gopurs_runtime.Apply(pure_22_50, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_61
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t61 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
goto end_branch_61
} else {

}
}
{
__t61 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_61:
return __t61
}))
})
})})}
}))
_ = __local_var_20_48
// TAST (Let): Bind1_21_62 -> *Constructor_Control_Bind_Bind
Bind1_21_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_48, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_62
// TAST (Let): Applicative0_22_63 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_48, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_63
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_19_44)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_62.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_62.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_63.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_64, x_19)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_66 -> *Constructor_Control_Bind_Bind
Bind1_18_66 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_66
// TAST (Let): pure_19_67 -> gopurs_runtime.Value
pure_19_67 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_19_67
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_69 -> gopurs_runtime.Value
__local_var_21_69 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_69
// TAST (Let): functorExceptT1_21_68 -> *Constructor_Data_Functor_Functor
functorExceptT1_21_68 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_70 -> gopurs_runtime.Value
__local_var_23_70 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_21_69, "map"), gopurs_runtime.Func(func(m_23 gopurs_runtime.Value) gopurs_runtime.Value {
var __t71 gopurs_runtime.Value
{
if (m_23.Type == 9 && m_23.IntVal == 3711209382) {
__t71 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_23.UnsafePtr).V0})}
goto end_branch_71
} else {

}
}
{
if (m_23.Type == 9 && m_23.IntVal == 2465973597) {
__t71 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_22, (*Constructor_Data_Either_Right)(m_23.UnsafePtr).V0)})}
goto end_branch_71
} else {

}
}
{
__t71 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_71:
return __t71
}))
_ = __local_var_23_70
return gopurs_runtime.Func(func(v_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_70, v_24)
})
})))
_ = functorExceptT1_21_68
// TAST (Let): __local_var_22_72 -> gopurs_runtime.Value
__local_var_22_72 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_84 -> gopurs_runtime.Value
__local_var_23_84 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_23_84
// TAST (Let): __local_var_23_83 -> gopurs_runtime.Value
__local_var_23_83 := gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_84, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_24})})
})
_ = __local_var_23_83
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_74 -> gopurs_runtime.Value
__local_var_24_74 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_74
// TAST (Let): functorExceptT1_24_73 -> *Constructor_Data_Functor_Functor
functorExceptT1_24_73 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_26_75 -> gopurs_runtime.Value
__local_var_26_75 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_74, "map"), gopurs_runtime.Func(func(m_26 gopurs_runtime.Value) gopurs_runtime.Value {
var __t76 gopurs_runtime.Value
{
if (m_26.Type == 9 && m_26.IntVal == 3711209382) {
__t76 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_26.UnsafePtr).V0})}
goto end_branch_76
} else {

}
}
{
if (m_26.Type == 9 && m_26.IntVal == 2465973597) {
__t76 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_25, (*Constructor_Data_Either_Right)(m_26.UnsafePtr).V0)})}
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
_ = __local_var_26_75
return gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_26_75, v_27)
})
})))
_ = functorExceptT1_24_73
// TAST (Let): __local_var_25_77 -> gopurs_runtime.Value
__local_var_25_77 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_78 -> *Constructor_Control_Bind_Bind
Bind1_26_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_78
// TAST (Let): pure_27_79 -> gopurs_runtime.Value
pure_27_79 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_27_79
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_78.V1), v_28, gopurs_runtime.Func(func(v2_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t80 gopurs_runtime.Value
{
if (v2_30.Type == 9 && v2_30.IntVal == 3711209382) {
__t80 = gopurs_runtime.Apply(pure_27_79, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_30.UnsafePtr).V0})})
goto end_branch_80
} else {

}
}
{
if (v2_30.Type == 9 && v2_30.IntVal == 2465973597) {
__t80 = gopurs_runtime.Apply(k_29, (*Constructor_Data_Either_Right)(v2_30.UnsafePtr).V0)
goto end_branch_80
} else {

}
}
{
__t80 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_80:
return __t80
}))
})
})})}
}))
_ = __local_var_25_77
// TAST (Let): Bind1_26_81 -> *Constructor_Control_Bind_Bind
Bind1_26_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_77, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_81
// TAST (Let): Applicative0_27_82 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_77, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_82
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_24_73)}
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_81.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_81.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_82.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_83, x_24)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_85 -> *Constructor_Control_Bind_Bind
Bind1_23_85 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_85
// TAST (Let): pure_24_86 -> gopurs_runtime.Value
pure_24_86 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_24_86
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_85.V1), v_25, gopurs_runtime.Func(func(v2_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t87 gopurs_runtime.Value
{
if (v2_27.Type == 9 && v2_27.IntVal == 3711209382) {
__t87 = gopurs_runtime.Apply(pure_24_86, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_27.UnsafePtr).V0})})
goto end_branch_87
} else {

}
}
{
if (v2_27.Type == 9 && v2_27.IntVal == 2465973597) {
__t87 = gopurs_runtime.Apply(k_26, (*Constructor_Data_Either_Right)(v2_27.UnsafePtr).V0)
goto end_branch_87
} else {

}
}
{
__t87 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_87:
return __t87
}))
})
})})}
}))
_ = __local_var_22_72
// TAST (Let): Bind1_23_88 -> *Constructor_Control_Bind_Bind
Bind1_23_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_72, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_88
// TAST (Let): Applicative0_24_89 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_89 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_72, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_89
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_21_68)}
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_88.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_88.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_89.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_66.V1), v_20, gopurs_runtime.Func(func(v2_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t90 gopurs_runtime.Value
{
if (v2_22.Type == 9 && v2_22.IntVal == 3711209382) {
__t90 = gopurs_runtime.Apply(pure_19_67, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_22.UnsafePtr).V0})})
goto end_branch_90
} else {

}
}
{
if (v2_22.Type == 9 && v2_22.IntVal == 2465973597) {
__t90 = gopurs_runtime.Apply(k_21, (*Constructor_Data_Either_Right)(v2_22.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_17_43
// TAST (Let): Bind1_18_91 -> *Constructor_Control_Bind_Bind
Bind1_18_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_43, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_91
// TAST (Let): Applicative0_19_92 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_43, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_92
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_16_39)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_91.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_91.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_92.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_37.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t93 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t93 = gopurs_runtime.Apply(pure_14_38, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_93
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t93 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_12_14
// TAST (Let): Bind1_13_94 -> *Constructor_Control_Bind_Bind
Bind1_13_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_94
// TAST (Let): Applicative0_14_95 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_95 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_95
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_11_10)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_94.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_94.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_95.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_96, x_11)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_98 -> *Constructor_Control_Bind_Bind
Bind1_10_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_98
// TAST (Let): pure_11_99 -> gopurs_runtime.Value
pure_11_99 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_99
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_101 -> gopurs_runtime.Value
__local_var_13_101 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_101
// TAST (Let): functorExceptT1_13_100 -> *Constructor_Data_Functor_Functor
functorExceptT1_13_100 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_102 -> gopurs_runtime.Value
__local_var_15_102 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_101, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t103 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t103 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_15.UnsafePtr).V0})}
goto end_branch_103
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t103 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right)(m_15.UnsafePtr).V0)})}
goto end_branch_103
} else {

}
}
{
__t103 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_103:
return __t103
}))
_ = __local_var_15_102
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_102, v_16)
})
})))
_ = functorExceptT1_13_100
// TAST (Let): __local_var_14_104 -> gopurs_runtime.Value
__local_var_14_104 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_160 -> gopurs_runtime.Value
__local_var_15_160 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_15_160
// TAST (Let): __local_var_15_159 -> gopurs_runtime.Value
__local_var_15_159 := gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_160, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_16})})
})
_ = __local_var_15_159
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_106 -> gopurs_runtime.Value
__local_var_16_106 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_106
// TAST (Let): functorExceptT1_16_105 -> *Constructor_Data_Functor_Functor
functorExceptT1_16_105 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_107 -> gopurs_runtime.Value
__local_var_18_107 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_106, "map"), gopurs_runtime.Func(func(m_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t108 gopurs_runtime.Value
{
if (m_18.Type == 9 && m_18.IntVal == 3711209382) {
__t108 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_18.UnsafePtr).V0})}
goto end_branch_108
} else {

}
}
{
if (m_18.Type == 9 && m_18.IntVal == 2465973597) {
__t108 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Either_Right)(m_18.UnsafePtr).V0)})}
goto end_branch_108
} else {

}
}
{
__t108 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_108:
return __t108
}))
_ = __local_var_18_107
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_107, v_19)
})
})))
_ = functorExceptT1_16_105
// TAST (Let): __local_var_17_109 -> gopurs_runtime.Value
__local_var_17_109 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_131 -> gopurs_runtime.Value
__local_var_18_131 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_131
// TAST (Let): __local_var_18_130 -> gopurs_runtime.Value
__local_var_18_130 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_131, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_19})})
})
_ = __local_var_18_130
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_111 -> gopurs_runtime.Value
__local_var_19_111 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_111
// TAST (Let): functorExceptT1_19_110 -> *Constructor_Data_Functor_Functor
functorExceptT1_19_110 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_112 -> gopurs_runtime.Value
__local_var_21_112 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_111, "map"), gopurs_runtime.Func(func(m_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t113 gopurs_runtime.Value
{
if (m_21.Type == 9 && m_21.IntVal == 3711209382) {
__t113 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_21.UnsafePtr).V0})}
goto end_branch_113
} else {

}
}
{
if (m_21.Type == 9 && m_21.IntVal == 2465973597) {
__t113 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Either_Right)(m_21.UnsafePtr).V0)})}
goto end_branch_113
} else {

}
}
{
__t113 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_113:
return __t113
}))
_ = __local_var_21_112
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_112, v_22)
})
})))
_ = functorExceptT1_19_110
// TAST (Let): __local_var_20_114 -> gopurs_runtime.Value
__local_var_20_114 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_115 -> *Constructor_Control_Bind_Bind
Bind1_21_115 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_115
// TAST (Let): pure_22_116 -> gopurs_runtime.Value
pure_22_116 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_116
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_118 -> gopurs_runtime.Value
__local_var_24_118 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_118
// TAST (Let): functorExceptT1_24_117 -> *Constructor_Data_Functor_Functor
functorExceptT1_24_117 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_26_119 -> gopurs_runtime.Value
__local_var_26_119 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_118, "map"), gopurs_runtime.Func(func(m_26 gopurs_runtime.Value) gopurs_runtime.Value {
var __t120 gopurs_runtime.Value
{
if (m_26.Type == 9 && m_26.IntVal == 3711209382) {
__t120 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_26.UnsafePtr).V0})}
goto end_branch_120
} else {

}
}
{
if (m_26.Type == 9 && m_26.IntVal == 2465973597) {
__t120 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_25, (*Constructor_Data_Either_Right)(m_26.UnsafePtr).V0)})}
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
_ = __local_var_26_119
return gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_26_119, v_27)
})
})))
_ = functorExceptT1_24_117
// TAST (Let): __local_var_25_121 -> gopurs_runtime.Value
__local_var_25_121 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_122 -> *Constructor_Control_Bind_Bind
Bind1_26_122 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_122
// TAST (Let): pure_27_123 -> gopurs_runtime.Value
pure_27_123 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_27_123
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_122.V1), v_28, gopurs_runtime.Func(func(v2_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t124 gopurs_runtime.Value
{
if (v2_30.Type == 9 && v2_30.IntVal == 3711209382) {
__t124 = gopurs_runtime.Apply(pure_27_123, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_30.UnsafePtr).V0})})
goto end_branch_124
} else {

}
}
{
if (v2_30.Type == 9 && v2_30.IntVal == 2465973597) {
__t124 = gopurs_runtime.Apply(k_29, (*Constructor_Data_Either_Right)(v2_30.UnsafePtr).V0)
goto end_branch_124
} else {

}
}
{
__t124 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_124:
return __t124
}))
})
})})}
}))
_ = __local_var_25_121
// TAST (Let): Bind1_26_125 -> *Constructor_Control_Bind_Bind
Bind1_26_125 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_121, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_125
// TAST (Let): Applicative0_27_126 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_126 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_121, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_126
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_24_117)}
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_125.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_125.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_126.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_115.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t127 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t127 = gopurs_runtime.Apply(pure_22_116, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_127
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t127 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
goto end_branch_127
} else {

}
}
{
__t127 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_127:
return __t127
}))
})
})})}
}))
_ = __local_var_20_114
// TAST (Let): Bind1_21_128 -> *Constructor_Control_Bind_Bind
Bind1_21_128 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_114, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_128
// TAST (Let): Applicative0_22_129 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_129 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_114, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_129
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_19_110)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_128.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_128.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_129.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_130, x_19)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_132 -> *Constructor_Control_Bind_Bind
Bind1_18_132 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_132
// TAST (Let): pure_19_133 -> gopurs_runtime.Value
pure_19_133 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_19_133
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_135 -> gopurs_runtime.Value
__local_var_21_135 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_135
// TAST (Let): functorExceptT1_21_134 -> *Constructor_Data_Functor_Functor
functorExceptT1_21_134 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_136 -> gopurs_runtime.Value
__local_var_23_136 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_21_135, "map"), gopurs_runtime.Func(func(m_23 gopurs_runtime.Value) gopurs_runtime.Value {
var __t137 gopurs_runtime.Value
{
if (m_23.Type == 9 && m_23.IntVal == 3711209382) {
__t137 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_23.UnsafePtr).V0})}
goto end_branch_137
} else {

}
}
{
if (m_23.Type == 9 && m_23.IntVal == 2465973597) {
__t137 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_22, (*Constructor_Data_Either_Right)(m_23.UnsafePtr).V0)})}
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
_ = __local_var_23_136
return gopurs_runtime.Func(func(v_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_136, v_24)
})
})))
_ = functorExceptT1_21_134
// TAST (Let): __local_var_22_138 -> gopurs_runtime.Value
__local_var_22_138 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_150 -> gopurs_runtime.Value
__local_var_23_150 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_23_150
// TAST (Let): __local_var_23_149 -> gopurs_runtime.Value
__local_var_23_149 := gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_150, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_24})})
})
_ = __local_var_23_149
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_140 -> gopurs_runtime.Value
__local_var_24_140 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_140
// TAST (Let): functorExceptT1_24_139 -> *Constructor_Data_Functor_Functor
functorExceptT1_24_139 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_26_141 -> gopurs_runtime.Value
__local_var_26_141 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_24_140, "map"), gopurs_runtime.Func(func(m_26 gopurs_runtime.Value) gopurs_runtime.Value {
var __t142 gopurs_runtime.Value
{
if (m_26.Type == 9 && m_26.IntVal == 3711209382) {
__t142 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_26.UnsafePtr).V0})}
goto end_branch_142
} else {

}
}
{
if (m_26.Type == 9 && m_26.IntVal == 2465973597) {
__t142 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_25, (*Constructor_Data_Either_Right)(m_26.UnsafePtr).V0)})}
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
_ = __local_var_26_141
return gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_26_141, v_27)
})
})))
_ = functorExceptT1_24_139
// TAST (Let): __local_var_25_143 -> gopurs_runtime.Value
__local_var_25_143 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_144 -> *Constructor_Control_Bind_Bind
Bind1_26_144 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_144
// TAST (Let): pure_27_145 -> gopurs_runtime.Value
pure_27_145 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_27_145
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_144.V1), v_28, gopurs_runtime.Func(func(v2_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t146 gopurs_runtime.Value
{
if (v2_30.Type == 9 && v2_30.IntVal == 3711209382) {
__t146 = gopurs_runtime.Apply(pure_27_145, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_30.UnsafePtr).V0})})
goto end_branch_146
} else {

}
}
{
if (v2_30.Type == 9 && v2_30.IntVal == 2465973597) {
__t146 = gopurs_runtime.Apply(k_29, (*Constructor_Data_Either_Right)(v2_30.UnsafePtr).V0)
goto end_branch_146
} else {

}
}
{
__t146 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_146:
return __t146
}))
})
})})}
}))
_ = __local_var_25_143
// TAST (Let): Bind1_26_147 -> *Constructor_Control_Bind_Bind
Bind1_26_147 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_143, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_147
// TAST (Let): Applicative0_27_148 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_148 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_143, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_148
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_24_139)}
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_147.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_147.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_148.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_149, x_24)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_151 -> *Constructor_Control_Bind_Bind
Bind1_23_151 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_151
// TAST (Let): pure_24_152 -> gopurs_runtime.Value
pure_24_152 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_24_152
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_151.V1), v_25, gopurs_runtime.Func(func(v2_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t153 gopurs_runtime.Value
{
if (v2_27.Type == 9 && v2_27.IntVal == 3711209382) {
__t153 = gopurs_runtime.Apply(pure_24_152, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_27.UnsafePtr).V0})})
goto end_branch_153
} else {

}
}
{
if (v2_27.Type == 9 && v2_27.IntVal == 2465973597) {
__t153 = gopurs_runtime.Apply(k_26, (*Constructor_Data_Either_Right)(v2_27.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_22_138
// TAST (Let): Bind1_23_154 -> *Constructor_Control_Bind_Bind
Bind1_23_154 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_138, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_154
// TAST (Let): Applicative0_24_155 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_155 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_138, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_155
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_21_134)}
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_154.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_154.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_155.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_132.V1), v_20, gopurs_runtime.Func(func(v2_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t156 gopurs_runtime.Value
{
if (v2_22.Type == 9 && v2_22.IntVal == 3711209382) {
__t156 = gopurs_runtime.Apply(pure_19_133, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_22.UnsafePtr).V0})})
goto end_branch_156
} else {

}
}
{
if (v2_22.Type == 9 && v2_22.IntVal == 2465973597) {
__t156 = gopurs_runtime.Apply(k_21, (*Constructor_Data_Either_Right)(v2_22.UnsafePtr).V0)
goto end_branch_156
} else {

}
}
{
__t156 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_156:
return __t156
}))
})
})})}
}))
_ = __local_var_17_109
// TAST (Let): Bind1_18_157 -> *Constructor_Control_Bind_Bind
Bind1_18_157 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_109, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_157
// TAST (Let): Applicative0_19_158 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_158 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_109, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_158
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_16_105)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_157.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_157.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_158.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_159, x_16)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_161 -> *Constructor_Control_Bind_Bind
Bind1_15_161 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_161
// TAST (Let): pure_16_162 -> gopurs_runtime.Value
pure_16_162 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_162
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_164 -> gopurs_runtime.Value
__local_var_18_164 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_164
// TAST (Let): functorExceptT1_18_163 -> *Constructor_Data_Functor_Functor
functorExceptT1_18_163 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_165 -> gopurs_runtime.Value
__local_var_20_165 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_164, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t166 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t166 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_20.UnsafePtr).V0})}
goto end_branch_166
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t166 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right)(m_20.UnsafePtr).V0)})}
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
_ = __local_var_20_165
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_165, v_21)
})
})))
_ = functorExceptT1_18_163
// TAST (Let): __local_var_19_167 -> gopurs_runtime.Value
__local_var_19_167 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_179 -> gopurs_runtime.Value
__local_var_20_179 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_20_179
// TAST (Let): __local_var_20_178 -> gopurs_runtime.Value
__local_var_20_178 := gopurs_runtime.Func(func(x_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_179, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_21})})
})
_ = __local_var_20_178
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_169 -> gopurs_runtime.Value
__local_var_21_169 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_169
// TAST (Let): functorExceptT1_21_168 -> *Constructor_Data_Functor_Functor
functorExceptT1_21_168 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_170 -> gopurs_runtime.Value
__local_var_23_170 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_21_169, "map"), gopurs_runtime.Func(func(m_23 gopurs_runtime.Value) gopurs_runtime.Value {
var __t171 gopurs_runtime.Value
{
if (m_23.Type == 9 && m_23.IntVal == 3711209382) {
__t171 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_23.UnsafePtr).V0})}
goto end_branch_171
} else {

}
}
{
if (m_23.Type == 9 && m_23.IntVal == 2465973597) {
__t171 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_22, (*Constructor_Data_Either_Right)(m_23.UnsafePtr).V0)})}
goto end_branch_171
} else {

}
}
{
__t171 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_171:
return __t171
}))
_ = __local_var_23_170
return gopurs_runtime.Func(func(v_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_170, v_24)
})
})))
_ = functorExceptT1_21_168
// TAST (Let): __local_var_22_172 -> gopurs_runtime.Value
__local_var_22_172 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_173 -> *Constructor_Control_Bind_Bind
Bind1_23_173 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_173
// TAST (Let): pure_24_174 -> gopurs_runtime.Value
pure_24_174 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_24_174
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_173.V1), v_25, gopurs_runtime.Func(func(v2_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t175 gopurs_runtime.Value
{
if (v2_27.Type == 9 && v2_27.IntVal == 3711209382) {
__t175 = gopurs_runtime.Apply(pure_24_174, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_27.UnsafePtr).V0})})
goto end_branch_175
} else {

}
}
{
if (v2_27.Type == 9 && v2_27.IntVal == 2465973597) {
__t175 = gopurs_runtime.Apply(k_26, (*Constructor_Data_Either_Right)(v2_27.UnsafePtr).V0)
goto end_branch_175
} else {

}
}
{
__t175 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_175:
return __t175
}))
})
})})}
}))
_ = __local_var_22_172
// TAST (Let): Bind1_23_176 -> *Constructor_Control_Bind_Bind
Bind1_23_176 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_172, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_176
// TAST (Let): Applicative0_24_177 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_177 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_172, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_177
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_21_168)}
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_176.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_176.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_177.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_178, x_21)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_180 -> *Constructor_Control_Bind_Bind
Bind1_20_180 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_180
// TAST (Let): pure_21_181 -> gopurs_runtime.Value
pure_21_181 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_181
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_180.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t182 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t182 = gopurs_runtime.Apply(pure_21_181, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_182
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t182 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_182
} else {

}
}
{
__t182 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_182:
return __t182
}))
})
})})}
}))
_ = __local_var_19_167
// TAST (Let): Bind1_20_183 -> *Constructor_Control_Bind_Bind
Bind1_20_183 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_167, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_183
// TAST (Let): Applicative0_21_184 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_184 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_167, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_184
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_163)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_183.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_183.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_184.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_161.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t185 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t185 = gopurs_runtime.Apply(pure_16_162, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_185
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t185 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_185
} else {

}
}
{
__t185 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_185:
return __t185
}))
})
})})}
}))
_ = __local_var_14_104
// TAST (Let): Bind1_15_186 -> *Constructor_Control_Bind_Bind
Bind1_15_186 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_104, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_186
// TAST (Let): Applicative0_16_187 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_187 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_104, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_187
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_100)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_186.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_186.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_187.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_98.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t188 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t188 = gopurs_runtime.Apply(pure_11_99, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_188
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t188 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_188
} else {

}
}
{
__t188 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_188:
return __t188
}))
})
})})}
})}
_ = monadExceptT1_9_9
// TAST (Let): Bind1_10_190 -> *Constructor_Control_Bind_Bind
Bind1_10_190 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_190
// TAST (Let): pure_11_191 -> gopurs_runtime.Value
pure_11_191 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_191
// TAST (Let): __local_var_10_189 -> gopurs_runtime.Value
__local_var_10_189 := gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_190.V1), m_12, gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_11_191, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_13})})
}))
})
_ = __local_var_10_189
// TAST (Let): monadTellExceptT1_7_6 -> *Constructor_Control_Monad_Writer_Class_MonadTell
monadTellExceptT1_7_6 := &Constructor_Control_Monad_Writer_Class_MonadTell{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_9_9)}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_8_8)}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_189, gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "tell"), x_11))
})}
_ = monadTellExceptT1_7_6
return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Writer_Class_MonadWriter{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellExceptT1_7_6)}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Monoid0_6_5)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), v_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t194 gopurs_runtime.Value
{
var __t_tag_192 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0
if (__t_tag_192.Type == 9 && __t_tag_192.IntVal == 3711209382) {
__t194 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)((*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0.UnsafePtr).V0})}
goto end_branch_194
} else {

}
}
{
var __t_tag_193 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0
if (__t_tag_193.Type == 9 && __t_tag_193.IntVal == 2465973597) {
__t194 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Either_Right)((*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1})}})}
goto end_branch_194
} else {

}
}
{
__t194 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_194:
return gopurs_runtime.Apply(pure_4_3, __t194)
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_8, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t195 *Constructor_Data_Tuple_Tuple
{
if (a_9.Type == 9 && a_9.IntVal == 3711209382) {
__t195 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(a_9.UnsafePtr).V0})}, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_10
})}
goto end_branch_195
} else {

}
}
{
if (a_9.Type == 9 && a_9.IntVal == 2465973597) {
__t195 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Either_Right)(a_9.UnsafePtr).V0.UnsafePtr).V0})}, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Either_Right)(a_9.UnsafePtr).V0.UnsafePtr).V1}
goto end_branch_195
} else {

}
}
{
__t195 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_195:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t195)})
})))
})})}
}

func Call_Control_Monad_Except_Trans_monadThrowExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadExceptT1_1_0 -> *Constructor_Control_Monad_Monad
monadExceptT1_1_0 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_22 -> gopurs_runtime.Value
__local_var_2_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_2_22
// TAST (Let): __local_var_2_21 -> gopurs_runtime.Value
__local_var_2_21 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_22, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_3})})
})
_ = __local_var_2_21
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorExceptT1_3_1 -> *Constructor_Data_Functor_Functor
functorExceptT1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (m_5.Type == 9 && m_5.IntVal == 3711209382) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_5.UnsafePtr).V0})}
goto end_branch_4
} else {

}
}
{
if (m_5.Type == 9 && m_5.IntVal == 2465973597) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Either_Right)(m_5.UnsafePtr).V0)})}
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
_ = __local_var_5_3
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, v_6)
})
})))
_ = functorExceptT1_3_1
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_6 -> *Constructor_Control_Bind_Bind
Bind1_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_6
// TAST (Let): pure_6_7 -> gopurs_runtime.Value
pure_6_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_7
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_9
// TAST (Let): functorExceptT1_8_8 -> *Constructor_Data_Functor_Functor
functorExceptT1_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_9, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_10.UnsafePtr).V0})}
goto end_branch_11
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right)(m_10.UnsafePtr).V0)})}
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
_ = __local_var_10_10
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_10, v_11)
})
})))
_ = functorExceptT1_8_8
// TAST (Let): __local_var_9_12 -> gopurs_runtime.Value
__local_var_9_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_13 -> *Constructor_Control_Bind_Bind
Bind1_10_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_13
// TAST (Let): pure_11_14 -> gopurs_runtime.Value
pure_11_14 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_14
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_13.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(pure_11_14, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_15
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t15 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_9_12
// TAST (Let): Bind1_10_16 -> *Constructor_Control_Bind_Bind
Bind1_10_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_16
// TAST (Let): Applicative0_11_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_17
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_8)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_16.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_16.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_17.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_6.V1), v_7, gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_6_7, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_9.UnsafePtr).V0})})
goto end_branch_18
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_8, (*Constructor_Data_Either_Right)(v2_9.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_4_5
// TAST (Let): Bind1_5_19 -> *Constructor_Control_Bind_Bind
Bind1_5_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_19
// TAST (Let): Applicative0_6_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_20
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_3_1)}
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_19.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_19.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_20.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_21, x_3)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_23 -> *Constructor_Control_Bind_Bind
Bind1_2_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_23
// TAST (Let): pure_3_24 -> gopurs_runtime.Value
pure_3_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_24
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_26 -> gopurs_runtime.Value
__local_var_5_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_26
// TAST (Let): functorExceptT1_5_25 -> *Constructor_Data_Functor_Functor
functorExceptT1_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_26, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t28 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_7.UnsafePtr).V0})}
goto end_branch_28
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t28 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right)(m_7.UnsafePtr).V0)})}
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return __t28
}))
_ = __local_var_7_27
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_27, v_8)
})
})))
_ = functorExceptT1_5_25
// TAST (Let): __local_var_6_29 -> gopurs_runtime.Value
__local_var_6_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_41 -> gopurs_runtime.Value
__local_var_7_41 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_41
// TAST (Let): __local_var_7_40 -> gopurs_runtime.Value
__local_var_7_40 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_41, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_8})})
})
_ = __local_var_7_40
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_31 -> gopurs_runtime.Value
__local_var_8_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_31
// TAST (Let): functorExceptT1_8_30 -> *Constructor_Data_Functor_Functor
functorExceptT1_8_30 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_32 -> gopurs_runtime.Value
__local_var_10_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_31, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t33 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_10.UnsafePtr).V0})}
goto end_branch_33
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t33 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right)(m_10.UnsafePtr).V0)})}
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
_ = __local_var_10_32
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_32, v_11)
})
})))
_ = functorExceptT1_8_30
// TAST (Let): __local_var_9_34 -> gopurs_runtime.Value
__local_var_9_34 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_35 -> *Constructor_Control_Bind_Bind
Bind1_10_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_35
// TAST (Let): pure_11_36 -> gopurs_runtime.Value
pure_11_36 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_36
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_35.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t37 = gopurs_runtime.Apply(pure_11_36, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_37
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t37 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return __t37
}))
})
})})}
}))
_ = __local_var_9_34
// TAST (Let): Bind1_10_38 -> *Constructor_Control_Bind_Bind
Bind1_10_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_34, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_38
// TAST (Let): Applicative0_11_39 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_34, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_39
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_30)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_38.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_38.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_39.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_40, x_8)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_42 -> *Constructor_Control_Bind_Bind
Bind1_7_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_42
// TAST (Let): pure_8_43 -> gopurs_runtime.Value
pure_8_43 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_43
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_42.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t44 = gopurs_runtime.Apply(pure_8_43, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_44
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t44 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_6_29
// TAST (Let): Bind1_7_45 -> *Constructor_Control_Bind_Bind
Bind1_7_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_45
// TAST (Let): Applicative0_8_46 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_46
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_25)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_45.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_45.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_46.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_23.V1), v_4, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t47 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t47 = gopurs_runtime.Apply(pure_3_24, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0})})
goto end_branch_47
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t47 = gopurs_runtime.Apply(k_5, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0)
goto end_branch_47
} else {

}
}
{
__t47 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_47:
return __t47
}))
})
})})}
})}
_ = monadExceptT1_1_0
// TAST (Let): __local_var_2_49 -> gopurs_runtime.Value
__local_var_2_49 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_2_49
// TAST (Let): __local_var_2_48 -> gopurs_runtime.Value
__local_var_2_48 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_49, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_3})})
})
_ = __local_var_2_48
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Error_Class_MonadThrow{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_1_0)}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_48, x_3)
})})}
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
// TAST (Let): monadExceptT1_3_3 -> *Constructor_Control_Monad_Monad
monadExceptT1_3_3 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_91 -> gopurs_runtime.Value
__local_var_4_91 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_91
// TAST (Let): __local_var_4_90 -> gopurs_runtime.Value
__local_var_4_90 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_91, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_90
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorExceptT1_5_4 -> *Constructor_Data_Functor_Functor
functorExceptT1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_7.UnsafePtr).V0})}
goto end_branch_7
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right)(m_7.UnsafePtr).V0)})}
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
_ = __local_var_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_6, v_8)
})
})))
_ = functorExceptT1_5_4
// TAST (Let): __local_var_6_8 -> gopurs_runtime.Value
__local_var_6_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
__local_var_7_30 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_30
// TAST (Let): __local_var_7_29 -> gopurs_runtime.Value
__local_var_7_29 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_30, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_8})})
})
_ = __local_var_7_29
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_10
// TAST (Let): functorExceptT1_8_9 -> *Constructor_Data_Functor_Functor
functorExceptT1_8_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_10.UnsafePtr).V0})}
goto end_branch_12
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right)(m_10.UnsafePtr).V0)})}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
_ = __local_var_10_11
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_11, v_11)
})
})))
_ = functorExceptT1_8_9
// TAST (Let): __local_var_9_13 -> gopurs_runtime.Value
__local_var_9_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_14 -> *Constructor_Control_Bind_Bind
Bind1_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_14
// TAST (Let): pure_11_15 -> gopurs_runtime.Value
pure_11_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_15
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_17 -> gopurs_runtime.Value
__local_var_13_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_17
// TAST (Let): functorExceptT1_13_16 -> *Constructor_Data_Functor_Functor
functorExceptT1_13_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_18 -> gopurs_runtime.Value
__local_var_15_18 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_17, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_15.UnsafePtr).V0})}
goto end_branch_19
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t19 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right)(m_15.UnsafePtr).V0)})}
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
_ = __local_var_15_18
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_18, v_16)
})
})))
_ = functorExceptT1_13_16
// TAST (Let): __local_var_14_20 -> gopurs_runtime.Value
__local_var_14_20 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_21 -> *Constructor_Control_Bind_Bind
Bind1_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_21
// TAST (Let): pure_16_22 -> gopurs_runtime.Value
pure_16_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_22
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_21.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t23 = gopurs_runtime.Apply(pure_16_22, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_23
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t23 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_14_20
// TAST (Let): Bind1_15_24 -> *Constructor_Control_Bind_Bind
Bind1_15_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_20, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_24
// TAST (Let): Applicative0_16_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_20, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_25
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_16)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_24.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_24.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_25.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t26 = gopurs_runtime.Apply(pure_11_15, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_26
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t26 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_9_13
// TAST (Let): Bind1_10_27 -> *Constructor_Control_Bind_Bind
Bind1_10_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_27
// TAST (Let): Applicative0_11_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_28
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_9)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_27.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_27.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_28.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_29, x_8)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_31 -> *Constructor_Control_Bind_Bind
Bind1_7_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_31
// TAST (Let): pure_8_32 -> gopurs_runtime.Value
pure_8_32 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_32
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_34 -> gopurs_runtime.Value
__local_var_10_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_34
// TAST (Let): functorExceptT1_10_33 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_33 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_35 -> gopurs_runtime.Value
__local_var_12_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_34, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_36
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t36 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
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
_ = __local_var_12_35
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_35, v_13)
})
})))
_ = functorExceptT1_10_33
// TAST (Let): __local_var_11_37 -> gopurs_runtime.Value
__local_var_11_37 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_59 -> gopurs_runtime.Value
__local_var_12_59 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_59
// TAST (Let): __local_var_12_58 -> gopurs_runtime.Value
__local_var_12_58 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_59, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_58
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_39 -> gopurs_runtime.Value
__local_var_13_39 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_39
// TAST (Let): functorExceptT1_13_38 -> *Constructor_Data_Functor_Functor
functorExceptT1_13_38 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_40 -> gopurs_runtime.Value
__local_var_15_40 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_39, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t41 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_15.UnsafePtr).V0})}
goto end_branch_41
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t41 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right)(m_15.UnsafePtr).V0)})}
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
_ = __local_var_15_40
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_40, v_16)
})
})))
_ = functorExceptT1_13_38
// TAST (Let): __local_var_14_42 -> gopurs_runtime.Value
__local_var_14_42 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_43 -> *Constructor_Control_Bind_Bind
Bind1_15_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_43
// TAST (Let): pure_16_44 -> gopurs_runtime.Value
pure_16_44 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_44
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_46 -> gopurs_runtime.Value
__local_var_18_46 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_46
// TAST (Let): functorExceptT1_18_45 -> *Constructor_Data_Functor_Functor
functorExceptT1_18_45 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_47 -> gopurs_runtime.Value
__local_var_20_47 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_46, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t48 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t48 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_20.UnsafePtr).V0})}
goto end_branch_48
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t48 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right)(m_20.UnsafePtr).V0)})}
goto end_branch_48
} else {

}
}
{
__t48 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_48:
return __t48
}))
_ = __local_var_20_47
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_47, v_21)
})
})))
_ = functorExceptT1_18_45
// TAST (Let): __local_var_19_49 -> gopurs_runtime.Value
__local_var_19_49 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_50 -> *Constructor_Control_Bind_Bind
Bind1_20_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_50
// TAST (Let): pure_21_51 -> gopurs_runtime.Value
pure_21_51 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_51
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_50.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t52 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t52 = gopurs_runtime.Apply(pure_21_51, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_52
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t52 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_19_49
// TAST (Let): Bind1_20_53 -> *Constructor_Control_Bind_Bind
Bind1_20_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_49, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_53
// TAST (Let): Applicative0_21_54 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_49, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_54
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_45)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_53.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_53.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_54.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_43.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t55 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t55 = gopurs_runtime.Apply(pure_16_44, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_55
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t55 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_14_42
// TAST (Let): Bind1_15_56 -> *Constructor_Control_Bind_Bind
Bind1_15_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_42, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_56
// TAST (Let): Applicative0_16_57 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_42, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_57
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_38)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_56.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_56.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_57.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_58, x_13)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_60 -> *Constructor_Control_Bind_Bind
Bind1_12_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_60
// TAST (Let): pure_13_61 -> gopurs_runtime.Value
pure_13_61 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_61
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_63 -> gopurs_runtime.Value
__local_var_15_63 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_63
// TAST (Let): functorExceptT1_15_62 -> *Constructor_Data_Functor_Functor
functorExceptT1_15_62 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_64 -> gopurs_runtime.Value
__local_var_17_64 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_63, "map"), gopurs_runtime.Func(func(m_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t65 gopurs_runtime.Value
{
if (m_17.Type == 9 && m_17.IntVal == 3711209382) {
__t65 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_17.UnsafePtr).V0})}
goto end_branch_65
} else {

}
}
{
if (m_17.Type == 9 && m_17.IntVal == 2465973597) {
__t65 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Either_Right)(m_17.UnsafePtr).V0)})}
goto end_branch_65
} else {

}
}
{
__t65 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_65:
return __t65
}))
_ = __local_var_17_64
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_64, v_18)
})
})))
_ = functorExceptT1_15_62
// TAST (Let): __local_var_16_66 -> gopurs_runtime.Value
__local_var_16_66 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_78 -> gopurs_runtime.Value
__local_var_17_78 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_78
// TAST (Let): __local_var_17_77 -> gopurs_runtime.Value
__local_var_17_77 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_78, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_18})})
})
_ = __local_var_17_77
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_68 -> gopurs_runtime.Value
__local_var_18_68 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_68
// TAST (Let): functorExceptT1_18_67 -> *Constructor_Data_Functor_Functor
functorExceptT1_18_67 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_69 -> gopurs_runtime.Value
__local_var_20_69 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_68, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t70 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t70 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_20.UnsafePtr).V0})}
goto end_branch_70
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t70 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right)(m_20.UnsafePtr).V0)})}
goto end_branch_70
} else {

}
}
{
__t70 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_70:
return __t70
}))
_ = __local_var_20_69
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_69, v_21)
})
})))
_ = functorExceptT1_18_67
// TAST (Let): __local_var_19_71 -> gopurs_runtime.Value
__local_var_19_71 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_72 -> *Constructor_Control_Bind_Bind
Bind1_20_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_72
// TAST (Let): pure_21_73 -> gopurs_runtime.Value
pure_21_73 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_73
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_72.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t74 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t74 = gopurs_runtime.Apply(pure_21_73, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_74
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t74 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_74
} else {

}
}
{
__t74 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_74:
return __t74
}))
})
})})}
}))
_ = __local_var_19_71
// TAST (Let): Bind1_20_75 -> *Constructor_Control_Bind_Bind
Bind1_20_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_71, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_75
// TAST (Let): Applicative0_21_76 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_71, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_76
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_67)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_75.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_75.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_76.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_77, x_18)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_79 -> *Constructor_Control_Bind_Bind
Bind1_17_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_79
// TAST (Let): pure_18_80 -> gopurs_runtime.Value
pure_18_80 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_80
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_79.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t81 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t81 = gopurs_runtime.Apply(pure_18_80, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_81
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t81 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_81
} else {

}
}
{
__t81 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_81:
return __t81
}))
})
})})}
}))
_ = __local_var_16_66
// TAST (Let): Bind1_17_82 -> *Constructor_Control_Bind_Bind
Bind1_17_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_66, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_82
// TAST (Let): Applicative0_18_83 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_83 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_66, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_83
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_15_62)}
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_82.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_82.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_83.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_60.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t84 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t84 = gopurs_runtime.Apply(pure_13_61, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_84
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t84 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_84
} else {

}
}
{
__t84 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_84:
return __t84
}))
})
})})}
}))
_ = __local_var_11_37
// TAST (Let): Bind1_12_85 -> *Constructor_Control_Bind_Bind
Bind1_12_85 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_85
// TAST (Let): Applicative0_13_86 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_86
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_33)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_85.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_85.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_86.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_31.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t87 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t87 = gopurs_runtime.Apply(pure_8_32, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_87
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t87 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_87
} else {

}
}
{
__t87 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_87:
return __t87
}))
})
})})}
}))
_ = __local_var_6_8
// TAST (Let): Bind1_7_88 -> *Constructor_Control_Bind_Bind
Bind1_7_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_88
// TAST (Let): Applicative0_8_89 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_89 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_89
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_4)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_88.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_88.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_89.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_90, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_92 -> *Constructor_Control_Bind_Bind
Bind1_4_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_92
// TAST (Let): pure_5_93 -> gopurs_runtime.Value
pure_5_93 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_93
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_95 -> gopurs_runtime.Value
__local_var_7_95 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_95
// TAST (Let): functorExceptT1_7_94 -> *Constructor_Data_Functor_Functor
functorExceptT1_7_94 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_96 -> gopurs_runtime.Value
__local_var_9_96 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_95, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t97 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t97 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_9.UnsafePtr).V0})}
goto end_branch_97
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t97 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right)(m_9.UnsafePtr).V0)})}
goto end_branch_97
} else {

}
}
{
__t97 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_97:
return __t97
}))
_ = __local_var_9_96
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_96, v_10)
})
})))
_ = functorExceptT1_7_94
// TAST (Let): __local_var_8_98 -> gopurs_runtime.Value
__local_var_8_98 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_154 -> gopurs_runtime.Value
__local_var_9_154 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_154
// TAST (Let): __local_var_9_153 -> gopurs_runtime.Value
__local_var_9_153 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_154, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_153
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_100 -> gopurs_runtime.Value
__local_var_10_100 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_100
// TAST (Let): functorExceptT1_10_99 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_99 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_101 -> gopurs_runtime.Value
__local_var_12_101 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_100, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t102 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t102 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_102
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t102 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
goto end_branch_102
} else {

}
}
{
__t102 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_102:
return __t102
}))
_ = __local_var_12_101
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_101, v_13)
})
})))
_ = functorExceptT1_10_99
// TAST (Let): __local_var_11_103 -> gopurs_runtime.Value
__local_var_11_103 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_125 -> gopurs_runtime.Value
__local_var_12_125 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_125
// TAST (Let): __local_var_12_124 -> gopurs_runtime.Value
__local_var_12_124 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_125, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_124
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_105 -> gopurs_runtime.Value
__local_var_13_105 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_105
// TAST (Let): functorExceptT1_13_104 -> *Constructor_Data_Functor_Functor
functorExceptT1_13_104 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_106 -> gopurs_runtime.Value
__local_var_15_106 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_105, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t107 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t107 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_15.UnsafePtr).V0})}
goto end_branch_107
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t107 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right)(m_15.UnsafePtr).V0)})}
goto end_branch_107
} else {

}
}
{
__t107 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_107:
return __t107
}))
_ = __local_var_15_106
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_106, v_16)
})
})))
_ = functorExceptT1_13_104
// TAST (Let): __local_var_14_108 -> gopurs_runtime.Value
__local_var_14_108 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_109 -> *Constructor_Control_Bind_Bind
Bind1_15_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_109
// TAST (Let): pure_16_110 -> gopurs_runtime.Value
pure_16_110 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_110
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_112 -> gopurs_runtime.Value
__local_var_18_112 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_112
// TAST (Let): functorExceptT1_18_111 -> *Constructor_Data_Functor_Functor
functorExceptT1_18_111 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_113 -> gopurs_runtime.Value
__local_var_20_113 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_112, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t114 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t114 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_20.UnsafePtr).V0})}
goto end_branch_114
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t114 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right)(m_20.UnsafePtr).V0)})}
goto end_branch_114
} else {

}
}
{
__t114 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_114:
return __t114
}))
_ = __local_var_20_113
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_113, v_21)
})
})))
_ = functorExceptT1_18_111
// TAST (Let): __local_var_19_115 -> gopurs_runtime.Value
__local_var_19_115 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_116 -> *Constructor_Control_Bind_Bind
Bind1_20_116 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_116
// TAST (Let): pure_21_117 -> gopurs_runtime.Value
pure_21_117 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_117
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_116.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t118 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t118 = gopurs_runtime.Apply(pure_21_117, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_118
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t118 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_118
} else {

}
}
{
__t118 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_118:
return __t118
}))
})
})})}
}))
_ = __local_var_19_115
// TAST (Let): Bind1_20_119 -> *Constructor_Control_Bind_Bind
Bind1_20_119 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_115, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_119
// TAST (Let): Applicative0_21_120 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_120 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_115, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_120
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_111)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_119.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_119.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_120.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_109.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t121 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t121 = gopurs_runtime.Apply(pure_16_110, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_121
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t121 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_121
} else {

}
}
{
__t121 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_121:
return __t121
}))
})
})})}
}))
_ = __local_var_14_108
// TAST (Let): Bind1_15_122 -> *Constructor_Control_Bind_Bind
Bind1_15_122 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_108, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_122
// TAST (Let): Applicative0_16_123 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_123 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_108, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_123
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_104)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_122.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_122.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_123.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_124, x_13)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_126 -> *Constructor_Control_Bind_Bind
Bind1_12_126 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_126
// TAST (Let): pure_13_127 -> gopurs_runtime.Value
pure_13_127 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_127
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_129 -> gopurs_runtime.Value
__local_var_15_129 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_129
// TAST (Let): functorExceptT1_15_128 -> *Constructor_Data_Functor_Functor
functorExceptT1_15_128 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_130 -> gopurs_runtime.Value
__local_var_17_130 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_129, "map"), gopurs_runtime.Func(func(m_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t131 gopurs_runtime.Value
{
if (m_17.Type == 9 && m_17.IntVal == 3711209382) {
__t131 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_17.UnsafePtr).V0})}
goto end_branch_131
} else {

}
}
{
if (m_17.Type == 9 && m_17.IntVal == 2465973597) {
__t131 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Either_Right)(m_17.UnsafePtr).V0)})}
goto end_branch_131
} else {

}
}
{
__t131 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_131:
return __t131
}))
_ = __local_var_17_130
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_130, v_18)
})
})))
_ = functorExceptT1_15_128
// TAST (Let): __local_var_16_132 -> gopurs_runtime.Value
__local_var_16_132 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_144 -> gopurs_runtime.Value
__local_var_17_144 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_144
// TAST (Let): __local_var_17_143 -> gopurs_runtime.Value
__local_var_17_143 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_144, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_18})})
})
_ = __local_var_17_143
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_134 -> gopurs_runtime.Value
__local_var_18_134 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_134
// TAST (Let): functorExceptT1_18_133 -> *Constructor_Data_Functor_Functor
functorExceptT1_18_133 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_135 -> gopurs_runtime.Value
__local_var_20_135 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_18_134, "map"), gopurs_runtime.Func(func(m_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t136 gopurs_runtime.Value
{
if (m_20.Type == 9 && m_20.IntVal == 3711209382) {
__t136 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_20.UnsafePtr).V0})}
goto end_branch_136
} else {

}
}
{
if (m_20.Type == 9 && m_20.IntVal == 2465973597) {
__t136 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Either_Right)(m_20.UnsafePtr).V0)})}
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
_ = __local_var_20_135
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_135, v_21)
})
})))
_ = functorExceptT1_18_133
// TAST (Let): __local_var_19_137 -> gopurs_runtime.Value
__local_var_19_137 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_138 -> *Constructor_Control_Bind_Bind
Bind1_20_138 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_138
// TAST (Let): pure_21_139 -> gopurs_runtime.Value
pure_21_139 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_21_139
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_138.V1), v_22, gopurs_runtime.Func(func(v2_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t140 gopurs_runtime.Value
{
if (v2_24.Type == 9 && v2_24.IntVal == 3711209382) {
__t140 = gopurs_runtime.Apply(pure_21_139, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_24.UnsafePtr).V0})})
goto end_branch_140
} else {

}
}
{
if (v2_24.Type == 9 && v2_24.IntVal == 2465973597) {
__t140 = gopurs_runtime.Apply(k_23, (*Constructor_Data_Either_Right)(v2_24.UnsafePtr).V0)
goto end_branch_140
} else {

}
}
{
__t140 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_140:
return __t140
}))
})
})})}
}))
_ = __local_var_19_137
// TAST (Let): Bind1_20_141 -> *Constructor_Control_Bind_Bind
Bind1_20_141 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_137, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_141
// TAST (Let): Applicative0_21_142 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_142 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_137, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_142
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_18_133)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_141.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_141.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_142.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_143, x_18)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_145 -> *Constructor_Control_Bind_Bind
Bind1_17_145 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_145
// TAST (Let): pure_18_146 -> gopurs_runtime.Value
pure_18_146 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_146
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_145.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t147 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t147 = gopurs_runtime.Apply(pure_18_146, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_147
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t147 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
goto end_branch_147
} else {

}
}
{
__t147 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_147:
return __t147
}))
})
})})}
}))
_ = __local_var_16_132
// TAST (Let): Bind1_17_148 -> *Constructor_Control_Bind_Bind
Bind1_17_148 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_132, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_148
// TAST (Let): Applicative0_18_149 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_149 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_132, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_149
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_15_128)}
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_148.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_148.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_149.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_126.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t150 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t150 = gopurs_runtime.Apply(pure_13_127, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_150
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t150 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_150
} else {

}
}
{
__t150 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_150:
return __t150
}))
})
})})}
}))
_ = __local_var_11_103
// TAST (Let): Bind1_12_151 -> *Constructor_Control_Bind_Bind
Bind1_12_151 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_103, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_151
// TAST (Let): Applicative0_13_152 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_152 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_103, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_152
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_99)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_151.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_151.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_152.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_153, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_155 -> *Constructor_Control_Bind_Bind
Bind1_9_155 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_155
// TAST (Let): pure_10_156 -> gopurs_runtime.Value
pure_10_156 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_156
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_158 -> gopurs_runtime.Value
__local_var_12_158 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_158
// TAST (Let): functorExceptT1_12_157 -> *Constructor_Data_Functor_Functor
functorExceptT1_12_157 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_159 -> gopurs_runtime.Value
__local_var_14_159 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_158, "map"), gopurs_runtime.Func(func(m_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t160 gopurs_runtime.Value
{
if (m_14.Type == 9 && m_14.IntVal == 3711209382) {
__t160 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_14.UnsafePtr).V0})}
goto end_branch_160
} else {

}
}
{
if (m_14.Type == 9 && m_14.IntVal == 2465973597) {
__t160 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_13, (*Constructor_Data_Either_Right)(m_14.UnsafePtr).V0)})}
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
_ = __local_var_14_159
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_159, v_15)
})
})))
_ = functorExceptT1_12_157
// TAST (Let): __local_var_13_161 -> gopurs_runtime.Value
__local_var_13_161 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_173 -> gopurs_runtime.Value
__local_var_14_173 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_14_173
// TAST (Let): __local_var_14_172 -> gopurs_runtime.Value
__local_var_14_172 := gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_173, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_15})})
})
_ = __local_var_14_172
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_163 -> gopurs_runtime.Value
__local_var_15_163 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_163
// TAST (Let): functorExceptT1_15_162 -> *Constructor_Data_Functor_Functor
functorExceptT1_15_162 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_164 -> gopurs_runtime.Value
__local_var_17_164 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_163, "map"), gopurs_runtime.Func(func(m_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t165 gopurs_runtime.Value
{
if (m_17.Type == 9 && m_17.IntVal == 3711209382) {
__t165 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_17.UnsafePtr).V0})}
goto end_branch_165
} else {

}
}
{
if (m_17.Type == 9 && m_17.IntVal == 2465973597) {
__t165 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Either_Right)(m_17.UnsafePtr).V0)})}
goto end_branch_165
} else {

}
}
{
__t165 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_165:
return __t165
}))
_ = __local_var_17_164
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_164, v_18)
})
})))
_ = functorExceptT1_15_162
// TAST (Let): __local_var_16_166 -> gopurs_runtime.Value
__local_var_16_166 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_167 -> *Constructor_Control_Bind_Bind
Bind1_17_167 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_167
// TAST (Let): pure_18_168 -> gopurs_runtime.Value
pure_18_168 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_18_168
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_167.V1), v_19, gopurs_runtime.Func(func(v2_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t169 gopurs_runtime.Value
{
if (v2_21.Type == 9 && v2_21.IntVal == 3711209382) {
__t169 = gopurs_runtime.Apply(pure_18_168, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_21.UnsafePtr).V0})})
goto end_branch_169
} else {

}
}
{
if (v2_21.Type == 9 && v2_21.IntVal == 2465973597) {
__t169 = gopurs_runtime.Apply(k_20, (*Constructor_Data_Either_Right)(v2_21.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_16_166
// TAST (Let): Bind1_17_170 -> *Constructor_Control_Bind_Bind
Bind1_17_170 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_166, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_170
// TAST (Let): Applicative0_18_171 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_171 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_166, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_171
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_15_162)}
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_170.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_170.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_171.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_172, x_15)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_14_174 -> *Constructor_Control_Bind_Bind
Bind1_14_174 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_174
// TAST (Let): pure_15_175 -> gopurs_runtime.Value
pure_15_175 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_15_175
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_174.V1), v_16, gopurs_runtime.Func(func(v2_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t176 gopurs_runtime.Value
{
if (v2_18.Type == 9 && v2_18.IntVal == 3711209382) {
__t176 = gopurs_runtime.Apply(pure_15_175, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_18.UnsafePtr).V0})})
goto end_branch_176
} else {

}
}
{
if (v2_18.Type == 9 && v2_18.IntVal == 2465973597) {
__t176 = gopurs_runtime.Apply(k_17, (*Constructor_Data_Either_Right)(v2_18.UnsafePtr).V0)
goto end_branch_176
} else {

}
}
{
__t176 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_176:
return __t176
}))
})
})})}
}))
_ = __local_var_13_161
// TAST (Let): Bind1_14_177 -> *Constructor_Control_Bind_Bind
Bind1_14_177 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_161, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_177
// TAST (Let): Applicative0_15_178 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_178 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_161, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_178
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_12_157)}
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_177.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_177.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_178.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_155.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t179 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t179 = gopurs_runtime.Apply(pure_10_156, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_179
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t179 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
goto end_branch_179
} else {

}
}
{
__t179 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_179:
return __t179
}))
})
})})}
}))
_ = __local_var_8_98
// TAST (Let): Bind1_9_180 -> *Constructor_Control_Bind_Bind
Bind1_9_180 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_98, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_180
// TAST (Let): Applicative0_10_181 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_181 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_98, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_181
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_94)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_180.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_180.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_181.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_92.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t182 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t182 = gopurs_runtime.Apply(pure_5_93, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_182
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t182 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
goto end_branch_182
} else {

}
}
{
__t182 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_182:
return __t182
}))
})
})})}
})}
_ = monadExceptT1_3_3
// TAST (Let): __local_var_4_184 -> gopurs_runtime.Value
__local_var_4_184 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_184
// TAST (Let): __local_var_4_183 -> gopurs_runtime.Value
__local_var_4_183 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_184, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, x_5})})
})
_ = __local_var_4_183
// TAST (Let): monadThrowExceptT1_3_2 -> *Constructor_Control_Monad_Error_Class_MonadThrow
monadThrowExceptT1_3_2 := &Constructor_Control_Monad_Error_Class_MonadThrow{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_3_3)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_183, x_5)
})}
_ = monadThrowExceptT1_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Error_Class_MonadError{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowExceptT1_3_2)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_4, gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t185 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 3711209382) {
__t185 = gopurs_runtime.Apply(k_5, (*Constructor_Data_Either_Left)(v2_6.UnsafePtr).V0)
goto end_branch_185
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 2465973597) {
__t185 = gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(v2_6.UnsafePtr).V0})})
goto end_branch_185
} else {

}
}
{
__t185 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_185:
return __t185
}))
})
})})}
}

func Call_Control_Monad_Except_Trans_monadSTExceptT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadExceptT1_2_1 -> *Constructor_Control_Monad_Monad
monadExceptT1_2_1 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_23 -> gopurs_runtime.Value
__local_var_3_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_23
// TAST (Let): __local_var_3_22 -> gopurs_runtime.Value
__local_var_3_22 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_23, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_4})})
})
_ = __local_var_3_22
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorExceptT1_4_2 -> *Constructor_Data_Functor_Functor
functorExceptT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_6.UnsafePtr).V0})}
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right)(m_6.UnsafePtr).V0)})}
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
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})))
_ = functorExceptT1_4_2
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_7 -> *Constructor_Control_Bind_Bind
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): pure_7_8 -> gopurs_runtime.Value
pure_7_8 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_10
// TAST (Let): functorExceptT1_9_9 -> *Constructor_Data_Functor_Functor
functorExceptT1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_11.UnsafePtr).V0})}
goto end_branch_12
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t12 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right)(m_11.UnsafePtr).V0)})}
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
}))
_ = __local_var_11_11
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_11, v_12)
})
})))
_ = functorExceptT1_9_9
// TAST (Let): __local_var_10_13 -> gopurs_runtime.Value
__local_var_10_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_14 -> *Constructor_Control_Bind_Bind
Bind1_11_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_14
// TAST (Let): pure_12_15 -> gopurs_runtime.Value
pure_12_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_15
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_14.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t16 = gopurs_runtime.Apply(pure_12_15, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_16
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t16 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_13
// TAST (Let): Bind1_11_17 -> *Constructor_Control_Bind_Bind
Bind1_11_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_17
// TAST (Let): Applicative0_12_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_9)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_17.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_17.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_18.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), v_8, gopurs_runtime.Func(func(v2_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v2_10.Type == 9 && v2_10.IntVal == 3711209382) {
__t19 = gopurs_runtime.Apply(pure_7_8, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_10.UnsafePtr).V0})})
goto end_branch_19
} else {

}
}
{
if (v2_10.Type == 9 && v2_10.IntVal == 2465973597) {
__t19 = gopurs_runtime.Apply(k_9, (*Constructor_Data_Either_Right)(v2_10.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_5_6
// TAST (Let): Bind1_6_20 -> *Constructor_Control_Bind_Bind
Bind1_6_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_20
// TAST (Let): Applicative0_7_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_21
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_20.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_21.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_22, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_24 -> *Constructor_Control_Bind_Bind
Bind1_3_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_24
// TAST (Let): pure_4_25 -> gopurs_runtime.Value
pure_4_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_25
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_27 -> gopurs_runtime.Value
__local_var_6_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_27
// TAST (Let): functorExceptT1_6_26 -> *Constructor_Data_Functor_Functor
functorExceptT1_6_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_27, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_8.UnsafePtr).V0})}
goto end_branch_29
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t29 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right)(m_8.UnsafePtr).V0)})}
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
_ = __local_var_8_28
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_28, v_9)
})
})))
_ = functorExceptT1_6_26
// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
__local_var_7_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_42 -> gopurs_runtime.Value
__local_var_8_42 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_42
// TAST (Let): __local_var_8_41 -> gopurs_runtime.Value
__local_var_8_41 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_42, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_41
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_32 -> gopurs_runtime.Value
__local_var_9_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_32
// TAST (Let): functorExceptT1_9_31 -> *Constructor_Data_Functor_Functor
functorExceptT1_9_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_32, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_11.UnsafePtr).V0})}
goto end_branch_34
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t34 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right)(m_11.UnsafePtr).V0)})}
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
_ = __local_var_11_33
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_33, v_12)
})
})))
_ = functorExceptT1_9_31
// TAST (Let): __local_var_10_35 -> gopurs_runtime.Value
__local_var_10_35 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_36 -> *Constructor_Control_Bind_Bind
Bind1_11_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_36
// TAST (Let): pure_12_37 -> gopurs_runtime.Value
pure_12_37 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_37
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_36.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t38 = gopurs_runtime.Apply(pure_12_37, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_38
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t38 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_38
} else {

}
}
{
__t38 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_38:
return __t38
}))
})
})})}
}))
_ = __local_var_10_35
// TAST (Let): Bind1_11_39 -> *Constructor_Control_Bind_Bind
Bind1_11_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_35, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_39
// TAST (Let): Applicative0_12_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_35, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_40
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_31)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_39.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_39.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_40.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_41, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_43 -> *Constructor_Control_Bind_Bind
Bind1_8_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_43
// TAST (Let): pure_9_44 -> gopurs_runtime.Value
pure_9_44 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_44
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_43.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t45 = gopurs_runtime.Apply(pure_9_44, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_45
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t45 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_7_30
// TAST (Let): Bind1_8_46 -> *Constructor_Control_Bind_Bind
Bind1_8_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_46
// TAST (Let): Applicative0_9_47 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_47
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_26)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_46.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_46.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_47.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_24.V1), v_5, gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t48 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 3711209382) {
__t48 = gopurs_runtime.Apply(pure_4_25, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_7.UnsafePtr).V0})})
goto end_branch_48
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 2465973597) {
__t48 = gopurs_runtime.Apply(k_6, (*Constructor_Data_Either_Right)(v2_7.UnsafePtr).V0)
goto end_branch_48
} else {

}
}
{
__t48 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_48:
return __t48
}))
})
})})}
})}
_ = monadExceptT1_2_1
// TAST (Let): Bind1_3_50 -> *Constructor_Control_Bind_Bind
Bind1_3_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_50
// TAST (Let): pure_4_51 -> gopurs_runtime.Value
pure_4_51 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_51
// TAST (Let): __local_var_3_49 -> gopurs_runtime.Value
__local_var_3_49 := gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_50.V1), m_5, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_4_51, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_6})})
}))
})
_ = __local_var_3_49
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_ST_Class_MonadST{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_2_1)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_49, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
})})}
}

func Call_Control_Monad_Except_Trans_monoidExceptT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_22 -> gopurs_runtime.Value
__local_var_1_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_22
// TAST (Let): __local_var_1_21 -> gopurs_runtime.Value
__local_var_1_21 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_22, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_2})})
})
_ = __local_var_1_21
// TAST (Let): applicativeExceptT1_1_0 -> *Constructor_Control_Applicative_Applicative
applicativeExceptT1_1_0 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorExceptT1_2_1 -> *Constructor_Data_Functor_Functor
functorExceptT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (m_4.Type == 9 && m_4.IntVal == 3711209382) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_4.UnsafePtr).V0})}
goto end_branch_4
} else {

}
}
{
if (m_4.Type == 9 && m_4.IntVal == 2465973597) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Either_Right)(m_4.UnsafePtr).V0)})}
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
_ = __local_var_4_3
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, v_5)
})
})))
_ = functorExceptT1_2_1
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_6 -> *Constructor_Control_Bind_Bind
Bind1_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_6
// TAST (Let): pure_5_7 -> gopurs_runtime.Value
pure_5_7 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_7
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_9 -> gopurs_runtime.Value
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_9
// TAST (Let): functorExceptT1_7_8 -> *Constructor_Data_Functor_Functor
functorExceptT1_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_9, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_9.UnsafePtr).V0})}
goto end_branch_11
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right)(m_9.UnsafePtr).V0)})}
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
_ = __local_var_9_10
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_10, v_10)
})
})))
_ = functorExceptT1_7_8
// TAST (Let): __local_var_8_12 -> gopurs_runtime.Value
__local_var_8_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_13 -> *Constructor_Control_Bind_Bind
Bind1_9_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_13
// TAST (Let): pure_10_14 -> gopurs_runtime.Value
pure_10_14 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_14
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_13.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t15 = gopurs_runtime.Apply(pure_10_14, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_15
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t15 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_8_12
// TAST (Let): Bind1_9_16 -> *Constructor_Control_Bind_Bind
Bind1_9_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_16
// TAST (Let): Applicative0_10_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_17
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_8)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_16.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_16.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_17.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_6.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t18 = gopurs_runtime.Apply(pure_5_7, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_18
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t18 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_3_5
// TAST (Let): Bind1_4_19 -> *Constructor_Control_Bind_Bind
Bind1_4_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_19
// TAST (Let): Applicative0_5_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_20
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_2_1)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_19.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_19.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_20.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_21, x_2)
})}
_ = applicativeExceptT1_1_0
// TAST (Let): __local_var_2_25 -> gopurs_runtime.Value
__local_var_2_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_25
// TAST (Let): functorExceptT1_2_24 -> *Constructor_Data_Functor_Functor
functorExceptT1_2_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_26 -> gopurs_runtime.Value
__local_var_4_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_25, "map"), gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 gopurs_runtime.Value
{
if (m_4.Type == 9 && m_4.IntVal == 3711209382) {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_4.UnsafePtr).V0})}
goto end_branch_27
} else {

}
}
{
if (m_4.Type == 9 && m_4.IntVal == 2465973597) {
__t27 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Either_Right)(m_4.UnsafePtr).V0)})}
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
return __t27
}))
_ = __local_var_4_26
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_26, v_5)
})
})))
_ = functorExceptT1_2_24
// TAST (Let): __local_var_3_28 -> gopurs_runtime.Value
__local_var_3_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_84 -> gopurs_runtime.Value
__local_var_4_84 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_84
// TAST (Let): __local_var_4_83 -> gopurs_runtime.Value
__local_var_4_83 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_84, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_83
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_30 -> gopurs_runtime.Value
__local_var_5_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_30
// TAST (Let): functorExceptT1_5_29 -> *Constructor_Data_Functor_Functor
functorExceptT1_5_29 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_31 -> gopurs_runtime.Value
__local_var_7_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_30, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t32 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_7.UnsafePtr).V0})}
goto end_branch_32
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t32 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right)(m_7.UnsafePtr).V0)})}
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
_ = __local_var_7_31
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_31, v_8)
})
})))
_ = functorExceptT1_5_29
// TAST (Let): __local_var_6_33 -> gopurs_runtime.Value
__local_var_6_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_55 -> gopurs_runtime.Value
__local_var_7_55 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_55
// TAST (Let): __local_var_7_54 -> gopurs_runtime.Value
__local_var_7_54 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_55, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_8})})
})
_ = __local_var_7_54
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_35 -> gopurs_runtime.Value
__local_var_8_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_35
// TAST (Let): functorExceptT1_8_34 -> *Constructor_Data_Functor_Functor
functorExceptT1_8_34 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_36 -> gopurs_runtime.Value
__local_var_10_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_35, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_10.UnsafePtr).V0})}
goto end_branch_37
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t37 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right)(m_10.UnsafePtr).V0)})}
goto end_branch_37
} else {

}
}
{
__t37 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_37:
return __t37
}))
_ = __local_var_10_36
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_36, v_11)
})
})))
_ = functorExceptT1_8_34
// TAST (Let): __local_var_9_38 -> gopurs_runtime.Value
__local_var_9_38 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_39 -> *Constructor_Control_Bind_Bind
Bind1_10_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_39
// TAST (Let): pure_11_40 -> gopurs_runtime.Value
pure_11_40 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_40
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_42 -> gopurs_runtime.Value
__local_var_13_42 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_42
// TAST (Let): functorExceptT1_13_41 -> *Constructor_Data_Functor_Functor
functorExceptT1_13_41 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_43 -> gopurs_runtime.Value
__local_var_15_43 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_42, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t44 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_15.UnsafePtr).V0})}
goto end_branch_44
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t44 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right)(m_15.UnsafePtr).V0)})}
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
_ = __local_var_15_43
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_43, v_16)
})
})))
_ = functorExceptT1_13_41
// TAST (Let): __local_var_14_45 -> gopurs_runtime.Value
__local_var_14_45 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_46 -> *Constructor_Control_Bind_Bind
Bind1_15_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_46
// TAST (Let): pure_16_47 -> gopurs_runtime.Value
pure_16_47 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_47
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_46.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t48 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t48 = gopurs_runtime.Apply(pure_16_47, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_48
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t48 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_48
} else {

}
}
{
__t48 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_48:
return __t48
}))
})
})})}
}))
_ = __local_var_14_45
// TAST (Let): Bind1_15_49 -> *Constructor_Control_Bind_Bind
Bind1_15_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_45, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_49
// TAST (Let): Applicative0_16_50 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_45, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_50
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_41)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_49.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_49.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_50.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_39.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t51 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t51 = gopurs_runtime.Apply(pure_11_40, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_51
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t51 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
goto end_branch_51
} else {

}
}
{
__t51 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_51:
return __t51
}))
})
})})}
}))
_ = __local_var_9_38
// TAST (Let): Bind1_10_52 -> *Constructor_Control_Bind_Bind
Bind1_10_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_38, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_52
// TAST (Let): Applicative0_11_53 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_38, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_53
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_34)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_52.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_52.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_53.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_54, x_8)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_56 -> *Constructor_Control_Bind_Bind
Bind1_7_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_56
// TAST (Let): pure_8_57 -> gopurs_runtime.Value
pure_8_57 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_57
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_59 -> gopurs_runtime.Value
__local_var_10_59 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_59
// TAST (Let): functorExceptT1_10_58 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_58 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_60 -> gopurs_runtime.Value
__local_var_12_60 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_59, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t61 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t61 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_61
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t61 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
goto end_branch_61
} else {

}
}
{
__t61 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_61:
return __t61
}))
_ = __local_var_12_60
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_60, v_13)
})
})))
_ = functorExceptT1_10_58
// TAST (Let): __local_var_11_62 -> gopurs_runtime.Value
__local_var_11_62 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_74 -> gopurs_runtime.Value
__local_var_12_74 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_74
// TAST (Let): __local_var_12_73 -> gopurs_runtime.Value
__local_var_12_73 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_74, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_13})})
})
_ = __local_var_12_73
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_64 -> gopurs_runtime.Value
__local_var_13_64 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_64
// TAST (Let): functorExceptT1_13_63 -> *Constructor_Data_Functor_Functor
functorExceptT1_13_63 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_65 -> gopurs_runtime.Value
__local_var_15_65 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_64, "map"), gopurs_runtime.Func(func(m_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t66 gopurs_runtime.Value
{
if (m_15.Type == 9 && m_15.IntVal == 3711209382) {
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_15.UnsafePtr).V0})}
goto end_branch_66
} else {

}
}
{
if (m_15.Type == 9 && m_15.IntVal == 2465973597) {
__t66 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Either_Right)(m_15.UnsafePtr).V0)})}
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
_ = __local_var_15_65
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_65, v_16)
})
})))
_ = functorExceptT1_13_63
// TAST (Let): __local_var_14_67 -> gopurs_runtime.Value
__local_var_14_67 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_68 -> *Constructor_Control_Bind_Bind
Bind1_15_68 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_68
// TAST (Let): pure_16_69 -> gopurs_runtime.Value
pure_16_69 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_16_69
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_68.V1), v_17, gopurs_runtime.Func(func(v2_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t70 gopurs_runtime.Value
{
if (v2_19.Type == 9 && v2_19.IntVal == 3711209382) {
__t70 = gopurs_runtime.Apply(pure_16_69, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_19.UnsafePtr).V0})})
goto end_branch_70
} else {

}
}
{
if (v2_19.Type == 9 && v2_19.IntVal == 2465973597) {
__t70 = gopurs_runtime.Apply(k_18, (*Constructor_Data_Either_Right)(v2_19.UnsafePtr).V0)
goto end_branch_70
} else {

}
}
{
__t70 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_70:
return __t70
}))
})
})})}
}))
_ = __local_var_14_67
// TAST (Let): Bind1_15_71 -> *Constructor_Control_Bind_Bind
Bind1_15_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_67, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_71
// TAST (Let): Applicative0_16_72 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_67, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_72
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_13_63)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_71.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_71.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_72.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_73, x_13)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_75 -> *Constructor_Control_Bind_Bind
Bind1_12_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_75
// TAST (Let): pure_13_76 -> gopurs_runtime.Value
pure_13_76 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_76
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_75.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t77 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t77 = gopurs_runtime.Apply(pure_13_76, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_77
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t77 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
goto end_branch_77
} else {

}
}
{
__t77 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_77:
return __t77
}))
})
})})}
}))
_ = __local_var_11_62
// TAST (Let): Bind1_12_78 -> *Constructor_Control_Bind_Bind
Bind1_12_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_62, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_78
// TAST (Let): Applicative0_13_79 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_62, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_79
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_58)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_78.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_78.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_79.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_56.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t80 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t80 = gopurs_runtime.Apply(pure_8_57, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_80
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t80 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
goto end_branch_80
} else {

}
}
{
__t80 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_80:
return __t80
}))
})
})})}
}))
_ = __local_var_6_33
// TAST (Let): Bind1_7_81 -> *Constructor_Control_Bind_Bind
Bind1_7_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_81
// TAST (Let): Applicative0_8_82 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_82
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_29)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_81.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_81.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_82.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_83, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_85 -> *Constructor_Control_Bind_Bind
Bind1_4_85 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_85
// TAST (Let): pure_5_86 -> gopurs_runtime.Value
pure_5_86 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_86
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_88 -> gopurs_runtime.Value
__local_var_7_88 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_88
// TAST (Let): functorExceptT1_7_87 -> *Constructor_Data_Functor_Functor
functorExceptT1_7_87 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_89 -> gopurs_runtime.Value
__local_var_9_89 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_88, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t90 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t90 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_9.UnsafePtr).V0})}
goto end_branch_90
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t90 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right)(m_9.UnsafePtr).V0)})}
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
_ = __local_var_9_89
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_89, v_10)
})
})))
_ = functorExceptT1_7_87
// TAST (Let): __local_var_8_91 -> gopurs_runtime.Value
__local_var_8_91 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_103 -> gopurs_runtime.Value
__local_var_9_103 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_103
// TAST (Let): __local_var_9_102 -> gopurs_runtime.Value
__local_var_9_102 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_103, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_10})})
})
_ = __local_var_9_102
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_93 -> gopurs_runtime.Value
__local_var_10_93 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_93
// TAST (Let): functorExceptT1_10_92 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_92 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_94 -> gopurs_runtime.Value
__local_var_12_94 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_93, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t95 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t95 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_95
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t95 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
goto end_branch_95
} else {

}
}
{
__t95 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_95:
return __t95
}))
_ = __local_var_12_94
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_94, v_13)
})
})))
_ = functorExceptT1_10_92
// TAST (Let): __local_var_11_96 -> gopurs_runtime.Value
__local_var_11_96 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_97 -> *Constructor_Control_Bind_Bind
Bind1_12_97 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_97
// TAST (Let): pure_13_98 -> gopurs_runtime.Value
pure_13_98 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_98
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_97.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t99 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t99 = gopurs_runtime.Apply(pure_13_98, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_99
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t99 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_96
// TAST (Let): Bind1_12_100 -> *Constructor_Control_Bind_Bind
Bind1_12_100 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_96, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_100
// TAST (Let): Applicative0_13_101 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_101 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_96, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_101
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_92)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_100.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_100.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_101.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_102, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_104 -> *Constructor_Control_Bind_Bind
Bind1_9_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_104
// TAST (Let): pure_10_105 -> gopurs_runtime.Value
pure_10_105 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_10_105
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_104.V1), v_11, gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t106 gopurs_runtime.Value
{
if (v2_13.Type == 9 && v2_13.IntVal == 3711209382) {
__t106 = gopurs_runtime.Apply(pure_10_105, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_13.UnsafePtr).V0})})
goto end_branch_106
} else {

}
}
{
if (v2_13.Type == 9 && v2_13.IntVal == 2465973597) {
__t106 = gopurs_runtime.Apply(k_12, (*Constructor_Data_Either_Right)(v2_13.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_8_91
// TAST (Let): Bind1_9_107 -> *Constructor_Control_Bind_Bind
Bind1_9_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_91, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_107
// TAST (Let): Applicative0_10_108 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_91, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_108
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_87)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_107.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_107.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_108.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_85.V1), v_6, gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t109 gopurs_runtime.Value
{
if (v2_8.Type == 9 && v2_8.IntVal == 3711209382) {
__t109 = gopurs_runtime.Apply(pure_5_86, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_8.UnsafePtr).V0})})
goto end_branch_109
} else {

}
}
{
if (v2_8.Type == 9 && v2_8.IntVal == 2465973597) {
__t109 = gopurs_runtime.Apply(k_7, (*Constructor_Data_Either_Right)(v2_8.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_3_28
// TAST (Let): Bind1_4_110 -> *Constructor_Control_Bind_Bind
Bind1_4_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_110
// TAST (Let): Applicative0_5_111 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_111
// TAST (Let): applyExceptT1_2_23 -> *Constructor_Control_Apply_Apply
applyExceptT1_2_23 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_2_24)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_110.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_110.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_111.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})}
_ = applyExceptT1_2_23
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_113 -> *Constructor_Data_Functor_Functor
Functor0_4_113 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyExceptT1_2_23.V0), gopurs_runtime.Value{}))
_ = Functor0_4_113
// TAST (Let): __local_var_5_114 -> gopurs_runtime.Value
__local_var_5_114 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_5_114
// TAST (Let): semigroupExceptT2_4_112 -> *Constructor_Data_Semigroup_Semigroup
semigroupExceptT2_4_112 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyExceptT1_2_23.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_113.V0), __local_var_5_114, a_6), b_7)
})
})}
_ = semigroupExceptT2_4_112
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupExceptT2_4_112)}
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeExceptT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))})}
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
// TAST (Let): functorExceptT1_4_2 -> *Constructor_Data_Functor_Functor
functorExceptT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 3711209382) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_6.UnsafePtr).V0})}
goto end_branch_5
} else {

}
}
{
if (m_6.Type == 9 && m_6.IntVal == 2465973597) {
__t5 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Either_Right)(m_6.UnsafePtr).V0)})}
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
_ = __local_var_6_4
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_4, v_7)
})
})))
_ = functorExceptT1_4_2
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_4_2)}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), v_5, gopurs_runtime.Func(func(rm_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (rm_7.Type == 9 && rm_7.IntVal == 2465973597) {
__t8 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rm_7.UnsafePtr).V0})})
goto end_branch_8
} else {

}
}
{
if (rm_7.Type == 9 && rm_7.IntVal == 3711209382) {
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := (*Constructor_Data_Either_Left)(rm_7.UnsafePtr).V0
_ = __local_var_8_6
__t8 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_0.V1), v1_6, gopurs_runtime.Func(func(rn_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (rn_9.Type == 9 && rn_9.IntVal == 2465973597) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rn_9.UnsafePtr).V0})})
goto end_branch_7
} else {

}
}
{
if (rn_9.Type == 9 && rn_9.IntVal == 3711209382) {
__t7 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), __local_var_8_6, (*Constructor_Data_Either_Left)(rn_9.UnsafePtr).V0)})})
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
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}))
})
})})}
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
// TAST (Let): functorExceptT1_6_5 -> *Constructor_Data_Functor_Functor
functorExceptT1_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_7 -> gopurs_runtime.Value
__local_var_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_8.UnsafePtr).V0})}
goto end_branch_8
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right)(m_8.UnsafePtr).V0)})}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}))
_ = __local_var_8_7
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_7, v_9)
})
})))
_ = functorExceptT1_6_5
// TAST (Let): altExceptT2_4_2 -> *Constructor_Control_Alt_Alt
altExceptT2_4_2 := &Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_5)}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), v_7, gopurs_runtime.Func(func(rm_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (rm_9.Type == 9 && rm_9.IntVal == 2465973597) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rm_9.UnsafePtr).V0})})
goto end_branch_11
} else {

}
}
{
if (rm_9.Type == 9 && rm_9.IntVal == 3711209382) {
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := (*Constructor_Data_Either_Left)(rm_9.UnsafePtr).V0
_ = __local_var_10_9
__t11 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), v1_8, gopurs_runtime.Func(func(rn_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (rn_11.Type == 9 && rn_11.IntVal == 2465973597) {
__t10 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rn_11.UnsafePtr).V0})})
goto end_branch_10
} else {

}
}
{
if (rn_11.Type == 9 && rn_11.IntVal == 3711209382) {
__t10 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_10_9, (*Constructor_Data_Either_Left)(rn_11.UnsafePtr).V0)})})
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
})}
_ = altExceptT2_4_2
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altExceptT2_4_2)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, mempty_1_0})})})}
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
// TAST (Let): __local_var_4_24 -> gopurs_runtime.Value
__local_var_4_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_24
// TAST (Let): __local_var_4_23 -> gopurs_runtime.Value
__local_var_4_23 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_24, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_5})})
})
_ = __local_var_4_23
// TAST (Let): applicativeExceptT1_4_2 -> *Constructor_Control_Applicative_Applicative
applicativeExceptT1_4_2 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorExceptT1_5_3 -> *Constructor_Data_Functor_Functor
functorExceptT1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 3711209382) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_7.UnsafePtr).V0})}
goto end_branch_6
} else {

}
}
{
if (m_7.Type == 9 && m_7.IntVal == 2465973597) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Either_Right)(m_7.UnsafePtr).V0)})}
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
_ = __local_var_7_5
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, v_8)
})
})))
_ = functorExceptT1_5_3
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_8 -> *Constructor_Control_Bind_Bind
Bind1_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_8
// TAST (Let): pure_8_9 -> gopurs_runtime.Value
pure_8_9 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_8_9
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_11
// TAST (Let): functorExceptT1_10_10 -> *Constructor_Data_Functor_Functor
functorExceptT1_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_12 -> gopurs_runtime.Value
__local_var_12_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "map"), gopurs_runtime.Func(func(m_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (m_12.Type == 9 && m_12.IntVal == 3711209382) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_12.UnsafePtr).V0})}
goto end_branch_13
} else {

}
}
{
if (m_12.Type == 9 && m_12.IntVal == 2465973597) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Either_Right)(m_12.UnsafePtr).V0)})}
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
_ = __local_var_12_12
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_12, v_13)
})
})))
_ = functorExceptT1_10_10
// TAST (Let): __local_var_11_14 -> gopurs_runtime.Value
__local_var_11_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_15 -> *Constructor_Control_Bind_Bind
Bind1_12_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_15
// TAST (Let): pure_13_16 -> gopurs_runtime.Value
pure_13_16 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_13_16
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_15.V1), v_14, gopurs_runtime.Func(func(v2_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_16.Type == 9 && v2_16.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(pure_13_16, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_16.UnsafePtr).V0})})
goto end_branch_17
} else {

}
}
{
if (v2_16.Type == 9 && v2_16.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(k_15, (*Constructor_Data_Either_Right)(v2_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_14
// TAST (Let): Bind1_12_18 -> *Constructor_Control_Bind_Bind
Bind1_12_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_18
// TAST (Let): Applicative0_13_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_10_10)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_18.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_18.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_19.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_8.V1), v_9, gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(pure_8_9, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_11.UnsafePtr).V0})})
goto end_branch_20
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 2465973597) {
__t20 = gopurs_runtime.Apply(k_10, (*Constructor_Data_Either_Right)(v2_11.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_6_7
// TAST (Let): Bind1_7_21 -> *Constructor_Control_Bind_Bind
Bind1_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_21
// TAST (Let): Applicative0_8_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_22
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_5_3)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_21.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_21.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_22.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_23, x_5)
})}
_ = applicativeExceptT1_4_2
// TAST (Let): Bind1_5_27 -> *Constructor_Control_Bind_Bind
Bind1_5_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_27
// TAST (Let): Applicative0_6_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_28
// TAST (Let): __local_var_7_30 -> gopurs_runtime.Value
__local_var_7_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_30
// TAST (Let): functorExceptT1_7_29 -> *Constructor_Data_Functor_Functor
functorExceptT1_7_29 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_31 -> gopurs_runtime.Value
__local_var_9_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_30, "map"), gopurs_runtime.Func(func(m_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 gopurs_runtime.Value
{
if (m_9.Type == 9 && m_9.IntVal == 3711209382) {
__t32 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_9.UnsafePtr).V0})}
goto end_branch_32
} else {

}
}
{
if (m_9.Type == 9 && m_9.IntVal == 2465973597) {
__t32 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Either_Right)(m_9.UnsafePtr).V0)})}
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
_ = __local_var_9_31
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_31, v_10)
})
})))
_ = functorExceptT1_7_29
// TAST (Let): altExceptT2_5_26 -> *Constructor_Control_Alt_Alt
altExceptT2_5_26 := &Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_7_29)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_27.V1), v_8, gopurs_runtime.Func(func(rm_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (rm_10.Type == 9 && rm_10.IntVal == 2465973597) {
__t35 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_28.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rm_10.UnsafePtr).V0})})
goto end_branch_35
} else {

}
}
{
if (rm_10.Type == 9 && rm_10.IntVal == 3711209382) {
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := (*Constructor_Data_Either_Left)(rm_10.UnsafePtr).V0
_ = __local_var_11_33
__t35 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_27.V1), v1_9, gopurs_runtime.Func(func(rn_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (rn_12.Type == 9 && rn_12.IntVal == 2465973597) {
__t34 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_28.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rn_12.UnsafePtr).V0})})
goto end_branch_34
} else {

}
}
{
if (rn_12.Type == 9 && rn_12.IntVal == 3711209382) {
__t34 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_28.V1), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_11_33, (*Constructor_Data_Either_Left)(rn_12.UnsafePtr).V0)})})
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
})}
_ = altExceptT2_5_26
// TAST (Let): plusExceptT2_5_25 -> *Constructor_Control_Plus_Plus
plusExceptT2_5_25 := &Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altExceptT2_5_26)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, mempty_1_0})})}
_ = plusExceptT2_5_25
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeExceptT1_4_2)}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusExceptT2_5_25)}
})})}
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
// TAST (Let): monadExceptT1_4_2 -> *Constructor_Control_Monad_Monad
monadExceptT1_4_2 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_24 -> gopurs_runtime.Value
__local_var_5_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_24
// TAST (Let): __local_var_5_23 -> gopurs_runtime.Value
__local_var_5_23 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_24, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_6})})
})
_ = __local_var_5_23
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_4
// TAST (Let): functorExceptT1_6_3 -> *Constructor_Data_Functor_Functor
functorExceptT1_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_5 -> gopurs_runtime.Value
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_4, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_8.UnsafePtr).V0})}
goto end_branch_6
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right)(m_8.UnsafePtr).V0)})}
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
_ = __local_var_8_5
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, v_9)
})
})))
_ = functorExceptT1_6_3
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_8 -> *Constructor_Control_Bind_Bind
Bind1_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_8
// TAST (Let): pure_9_9 -> gopurs_runtime.Value
pure_9_9 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_9
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorExceptT1_11_10 -> *Constructor_Data_Functor_Functor
functorExceptT1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_12 -> gopurs_runtime.Value
__local_var_13_12 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Func(func(m_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (m_13.Type == 9 && m_13.IntVal == 3711209382) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_13.UnsafePtr).V0})}
goto end_branch_13
} else {

}
}
{
if (m_13.Type == 9 && m_13.IntVal == 2465973597) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Either_Right)(m_13.UnsafePtr).V0)})}
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
_ = __local_var_13_12
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_12, v_14)
})
})))
_ = functorExceptT1_11_10
// TAST (Let): __local_var_12_14 -> gopurs_runtime.Value
__local_var_12_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_15 -> *Constructor_Control_Bind_Bind
Bind1_13_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_15
// TAST (Let): pure_14_16 -> gopurs_runtime.Value
pure_14_16 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_16
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_15.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t17 = gopurs_runtime.Apply(pure_14_16, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_17
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t17 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_12_14
// TAST (Let): Bind1_13_18 -> *Constructor_Control_Bind_Bind
Bind1_13_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_18
// TAST (Let): Applicative0_14_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_11_10)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_18.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_18.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_19.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_8.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t20 = gopurs_runtime.Apply(pure_9_9, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_20
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t20 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_7_7
// TAST (Let): Bind1_8_21 -> *Constructor_Control_Bind_Bind
Bind1_8_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_21
// TAST (Let): Applicative0_9_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_22
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_3)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_21.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_21.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_22.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_23, x_6)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_25 -> *Constructor_Control_Bind_Bind
Bind1_5_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_25
// TAST (Let): pure_6_26 -> gopurs_runtime.Value
pure_6_26 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_6_26
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_28
// TAST (Let): functorExceptT1_8_27 -> *Constructor_Data_Functor_Functor
functorExceptT1_8_27 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_29 -> gopurs_runtime.Value
__local_var_10_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t30 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_10.UnsafePtr).V0})}
goto end_branch_30
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t30 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right)(m_10.UnsafePtr).V0)})}
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
return __t30
}))
_ = __local_var_10_29
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_29, v_11)
})
})))
_ = functorExceptT1_8_27
// TAST (Let): __local_var_9_31 -> gopurs_runtime.Value
__local_var_9_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_43 -> gopurs_runtime.Value
__local_var_10_43 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_43
// TAST (Let): __local_var_10_42 -> gopurs_runtime.Value
__local_var_10_42 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_43, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_11})})
})
_ = __local_var_10_42
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_33
// TAST (Let): functorExceptT1_11_32 -> *Constructor_Data_Functor_Functor
functorExceptT1_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_34 -> gopurs_runtime.Value
__local_var_13_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_33, "map"), gopurs_runtime.Func(func(m_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (m_13.Type == 9 && m_13.IntVal == 3711209382) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_13.UnsafePtr).V0})}
goto end_branch_35
} else {

}
}
{
if (m_13.Type == 9 && m_13.IntVal == 2465973597) {
__t35 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Either_Right)(m_13.UnsafePtr).V0)})}
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
_ = __local_var_13_34
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_34, v_14)
})
})))
_ = functorExceptT1_11_32
// TAST (Let): __local_var_12_36 -> gopurs_runtime.Value
__local_var_12_36 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_37 -> *Constructor_Control_Bind_Bind
Bind1_13_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_37
// TAST (Let): pure_14_38 -> gopurs_runtime.Value
pure_14_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_38
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_37.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t39 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t39 = gopurs_runtime.Apply(pure_14_38, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_39
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t39 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
goto end_branch_39
} else {

}
}
{
__t39 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_39:
return __t39
}))
})
})})}
}))
_ = __local_var_12_36
// TAST (Let): Bind1_13_40 -> *Constructor_Control_Bind_Bind
Bind1_13_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_36, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_40
// TAST (Let): Applicative0_14_41 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_36, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_41
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_11_32)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_40.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_40.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_41.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_42, x_11)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_44 -> *Constructor_Control_Bind_Bind
Bind1_10_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_44
// TAST (Let): pure_11_45 -> gopurs_runtime.Value
pure_11_45 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_11_45
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_44.V1), v_12, gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 3711209382) {
__t46 = gopurs_runtime.Apply(pure_11_45, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_14.UnsafePtr).V0})})
goto end_branch_46
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 2465973597) {
__t46 = gopurs_runtime.Apply(k_13, (*Constructor_Data_Either_Right)(v2_14.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_9_31
// TAST (Let): Bind1_10_47 -> *Constructor_Control_Bind_Bind
Bind1_10_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_47
// TAST (Let): Applicative0_11_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_48
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_27)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_47.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_47.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_48.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_25.V1), v_7, gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t49 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 3711209382) {
__t49 = gopurs_runtime.Apply(pure_6_26, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_9.UnsafePtr).V0})})
goto end_branch_49
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 2465973597) {
__t49 = gopurs_runtime.Apply(k_8, (*Constructor_Data_Either_Right)(v2_9.UnsafePtr).V0)
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
})})}
})}
_ = monadExceptT1_4_2
// TAST (Let): __local_var_5_139 -> gopurs_runtime.Value
__local_var_5_139 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_139
// TAST (Let): __local_var_5_138 -> gopurs_runtime.Value
__local_var_5_138 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_139, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_6})})
})
_ = __local_var_5_138
// TAST (Let): applicativeExceptT1_5_51 -> *Constructor_Control_Applicative_Applicative
applicativeExceptT1_5_51 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_53 -> gopurs_runtime.Value
__local_var_6_53 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_53
// TAST (Let): functorExceptT1_6_52 -> *Constructor_Data_Functor_Functor
functorExceptT1_6_52 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_54 -> gopurs_runtime.Value
__local_var_8_54 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_53, "map"), gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t55 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 3711209382) {
__t55 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_8.UnsafePtr).V0})}
goto end_branch_55
} else {

}
}
{
if (m_8.Type == 9 && m_8.IntVal == 2465973597) {
__t55 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Either_Right)(m_8.UnsafePtr).V0)})}
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
_ = __local_var_8_54
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_54, v_9)
})
})))
_ = functorExceptT1_6_52
// TAST (Let): __local_var_7_56 -> gopurs_runtime.Value
__local_var_7_56 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_78 -> gopurs_runtime.Value
__local_var_8_78 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_78
// TAST (Let): __local_var_8_77 -> gopurs_runtime.Value
__local_var_8_77 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_78, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_9})})
})
_ = __local_var_8_77
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_58 -> gopurs_runtime.Value
__local_var_9_58 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_58
// TAST (Let): functorExceptT1_9_57 -> *Constructor_Data_Functor_Functor
functorExceptT1_9_57 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_59 -> gopurs_runtime.Value
__local_var_11_59 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_58, "map"), gopurs_runtime.Func(func(m_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t60 gopurs_runtime.Value
{
if (m_11.Type == 9 && m_11.IntVal == 3711209382) {
__t60 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_11.UnsafePtr).V0})}
goto end_branch_60
} else {

}
}
{
if (m_11.Type == 9 && m_11.IntVal == 2465973597) {
__t60 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Either_Right)(m_11.UnsafePtr).V0)})}
goto end_branch_60
} else {

}
}
{
__t60 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_60:
return __t60
}))
_ = __local_var_11_59
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_11_59, v_12)
})
})))
_ = functorExceptT1_9_57
// TAST (Let): __local_var_10_61 -> gopurs_runtime.Value
__local_var_10_61 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_62 -> *Constructor_Control_Bind_Bind
Bind1_11_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_62
// TAST (Let): pure_12_63 -> gopurs_runtime.Value
pure_12_63 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_12_63
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_65 -> gopurs_runtime.Value
__local_var_14_65 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_65
// TAST (Let): functorExceptT1_14_64 -> *Constructor_Data_Functor_Functor
functorExceptT1_14_64 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_66 -> gopurs_runtime.Value
__local_var_16_66 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_65, "map"), gopurs_runtime.Func(func(m_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t67 gopurs_runtime.Value
{
if (m_16.Type == 9 && m_16.IntVal == 3711209382) {
__t67 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_16.UnsafePtr).V0})}
goto end_branch_67
} else {

}
}
{
if (m_16.Type == 9 && m_16.IntVal == 2465973597) {
__t67 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Either_Right)(m_16.UnsafePtr).V0)})}
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
_ = __local_var_16_66
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_66, v_17)
})
})))
_ = functorExceptT1_14_64
// TAST (Let): __local_var_15_68 -> gopurs_runtime.Value
__local_var_15_68 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_69 -> *Constructor_Control_Bind_Bind
Bind1_16_69 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_69
// TAST (Let): pure_17_70 -> gopurs_runtime.Value
pure_17_70 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_70
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_69.V1), v_18, gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t71 gopurs_runtime.Value
{
if (v2_20.Type == 9 && v2_20.IntVal == 3711209382) {
__t71 = gopurs_runtime.Apply(pure_17_70, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_20.UnsafePtr).V0})})
goto end_branch_71
} else {

}
}
{
if (v2_20.Type == 9 && v2_20.IntVal == 2465973597) {
__t71 = gopurs_runtime.Apply(k_19, (*Constructor_Data_Either_Right)(v2_20.UnsafePtr).V0)
goto end_branch_71
} else {

}
}
{
__t71 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_71:
return __t71
}))
})
})})}
}))
_ = __local_var_15_68
// TAST (Let): Bind1_16_72 -> *Constructor_Control_Bind_Bind
Bind1_16_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_68, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_72
// TAST (Let): Applicative0_17_73 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_68, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_73
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_14_64)}
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_72.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_72.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_73.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_62.V1), v_13, gopurs_runtime.Func(func(v2_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t74 gopurs_runtime.Value
{
if (v2_15.Type == 9 && v2_15.IntVal == 3711209382) {
__t74 = gopurs_runtime.Apply(pure_12_63, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_15.UnsafePtr).V0})})
goto end_branch_74
} else {

}
}
{
if (v2_15.Type == 9 && v2_15.IntVal == 2465973597) {
__t74 = gopurs_runtime.Apply(k_14, (*Constructor_Data_Either_Right)(v2_15.UnsafePtr).V0)
goto end_branch_74
} else {

}
}
{
__t74 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_74:
return __t74
}))
})
})})}
}))
_ = __local_var_10_61
// TAST (Let): Bind1_11_75 -> *Constructor_Control_Bind_Bind
Bind1_11_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_61, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_75
// TAST (Let): Applicative0_12_76 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_61, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_76
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_9_57)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_75.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_75.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_76.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_77, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_79 -> *Constructor_Control_Bind_Bind
Bind1_8_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_79
// TAST (Let): pure_9_80 -> gopurs_runtime.Value
pure_9_80 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_9_80
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_82 -> gopurs_runtime.Value
__local_var_11_82 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_82
// TAST (Let): functorExceptT1_11_81 -> *Constructor_Data_Functor_Functor
functorExceptT1_11_81 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_83 -> gopurs_runtime.Value
__local_var_13_83 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_82, "map"), gopurs_runtime.Func(func(m_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t84 gopurs_runtime.Value
{
if (m_13.Type == 9 && m_13.IntVal == 3711209382) {
__t84 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_13.UnsafePtr).V0})}
goto end_branch_84
} else {

}
}
{
if (m_13.Type == 9 && m_13.IntVal == 2465973597) {
__t84 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Either_Right)(m_13.UnsafePtr).V0)})}
goto end_branch_84
} else {

}
}
{
__t84 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_84:
return __t84
}))
_ = __local_var_13_83
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_83, v_14)
})
})))
_ = functorExceptT1_11_81
// TAST (Let): __local_var_12_85 -> gopurs_runtime.Value
__local_var_12_85 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_107 -> gopurs_runtime.Value
__local_var_13_107 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_107
// TAST (Let): __local_var_13_106 -> gopurs_runtime.Value
__local_var_13_106 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_107, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_14})})
})
_ = __local_var_13_106
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_87 -> gopurs_runtime.Value
__local_var_14_87 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_87
// TAST (Let): functorExceptT1_14_86 -> *Constructor_Data_Functor_Functor
functorExceptT1_14_86 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_88 -> gopurs_runtime.Value
__local_var_16_88 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_87, "map"), gopurs_runtime.Func(func(m_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t89 gopurs_runtime.Value
{
if (m_16.Type == 9 && m_16.IntVal == 3711209382) {
__t89 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_16.UnsafePtr).V0})}
goto end_branch_89
} else {

}
}
{
if (m_16.Type == 9 && m_16.IntVal == 2465973597) {
__t89 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Either_Right)(m_16.UnsafePtr).V0)})}
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
_ = __local_var_16_88
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_16_88, v_17)
})
})))
_ = functorExceptT1_14_86
// TAST (Let): __local_var_15_90 -> gopurs_runtime.Value
__local_var_15_90 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_91 -> *Constructor_Control_Bind_Bind
Bind1_16_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_91
// TAST (Let): pure_17_92 -> gopurs_runtime.Value
pure_17_92 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_17_92
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_94 -> gopurs_runtime.Value
__local_var_19_94 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_94
// TAST (Let): functorExceptT1_19_93 -> *Constructor_Data_Functor_Functor
functorExceptT1_19_93 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_95 -> gopurs_runtime.Value
__local_var_21_95 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_94, "map"), gopurs_runtime.Func(func(m_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t96 gopurs_runtime.Value
{
if (m_21.Type == 9 && m_21.IntVal == 3711209382) {
__t96 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_21.UnsafePtr).V0})}
goto end_branch_96
} else {

}
}
{
if (m_21.Type == 9 && m_21.IntVal == 2465973597) {
__t96 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Either_Right)(m_21.UnsafePtr).V0)})}
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
_ = __local_var_21_95
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_95, v_22)
})
})))
_ = functorExceptT1_19_93
// TAST (Let): __local_var_20_97 -> gopurs_runtime.Value
__local_var_20_97 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_98 -> *Constructor_Control_Bind_Bind
Bind1_21_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_98
// TAST (Let): pure_22_99 -> gopurs_runtime.Value
pure_22_99 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_99
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_98.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t100 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t100 = gopurs_runtime.Apply(pure_22_99, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_100
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t100 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
goto end_branch_100
} else {

}
}
{
__t100 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_100:
return __t100
}))
})
})})}
}))
_ = __local_var_20_97
// TAST (Let): Bind1_21_101 -> *Constructor_Control_Bind_Bind
Bind1_21_101 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_97, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_101
// TAST (Let): Applicative0_22_102 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_102 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_97, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_102
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_19_93)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_101.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_101.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_102.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_91.V1), v_18, gopurs_runtime.Func(func(v2_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t103 gopurs_runtime.Value
{
if (v2_20.Type == 9 && v2_20.IntVal == 3711209382) {
__t103 = gopurs_runtime.Apply(pure_17_92, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_20.UnsafePtr).V0})})
goto end_branch_103
} else {

}
}
{
if (v2_20.Type == 9 && v2_20.IntVal == 2465973597) {
__t103 = gopurs_runtime.Apply(k_19, (*Constructor_Data_Either_Right)(v2_20.UnsafePtr).V0)
goto end_branch_103
} else {

}
}
{
__t103 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_103:
return __t103
}))
})
})})}
}))
_ = __local_var_15_90
// TAST (Let): Bind1_16_104 -> *Constructor_Control_Bind_Bind
Bind1_16_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_90, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_104
// TAST (Let): Applicative0_17_105 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_105 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_90, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_105
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_14_86)}
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_104.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_104.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_105.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_106, x_14)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_108 -> *Constructor_Control_Bind_Bind
Bind1_13_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_108
// TAST (Let): pure_14_109 -> gopurs_runtime.Value
pure_14_109 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_109
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_111 -> gopurs_runtime.Value
__local_var_16_111 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_111
// TAST (Let): functorExceptT1_16_110 -> *Constructor_Data_Functor_Functor
functorExceptT1_16_110 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_112 -> gopurs_runtime.Value
__local_var_18_112 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_111, "map"), gopurs_runtime.Func(func(m_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t113 gopurs_runtime.Value
{
if (m_18.Type == 9 && m_18.IntVal == 3711209382) {
__t113 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_18.UnsafePtr).V0})}
goto end_branch_113
} else {

}
}
{
if (m_18.Type == 9 && m_18.IntVal == 2465973597) {
__t113 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Either_Right)(m_18.UnsafePtr).V0)})}
goto end_branch_113
} else {

}
}
{
__t113 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_113:
return __t113
}))
_ = __local_var_18_112
return gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_112, v_19)
})
})))
_ = functorExceptT1_16_110
// TAST (Let): __local_var_17_114 -> gopurs_runtime.Value
__local_var_17_114 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_126 -> gopurs_runtime.Value
__local_var_18_126 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_126
// TAST (Let): __local_var_18_125 -> gopurs_runtime.Value
__local_var_18_125 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_126, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, x_19})})
})
_ = __local_var_18_125
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_116 -> gopurs_runtime.Value
__local_var_19_116 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_116
// TAST (Let): functorExceptT1_19_115 -> *Constructor_Data_Functor_Functor
functorExceptT1_19_115 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_117 -> gopurs_runtime.Value
__local_var_21_117 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_116, "map"), gopurs_runtime.Func(func(m_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t118 gopurs_runtime.Value
{
if (m_21.Type == 9 && m_21.IntVal == 3711209382) {
__t118 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_21.UnsafePtr).V0})}
goto end_branch_118
} else {

}
}
{
if (m_21.Type == 9 && m_21.IntVal == 2465973597) {
__t118 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Either_Right)(m_21.UnsafePtr).V0)})}
goto end_branch_118
} else {

}
}
{
__t118 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_118:
return __t118
}))
_ = __local_var_21_117
return gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_21_117, v_22)
})
})))
_ = functorExceptT1_19_115
// TAST (Let): __local_var_20_119 -> gopurs_runtime.Value
__local_var_20_119 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Except_Trans_applicativeExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_120 -> *Constructor_Control_Bind_Bind
Bind1_21_120 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_120
// TAST (Let): pure_22_121 -> gopurs_runtime.Value
pure_22_121 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_22_121
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_120.V1), v_23, gopurs_runtime.Func(func(v2_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t122 gopurs_runtime.Value
{
if (v2_25.Type == 9 && v2_25.IntVal == 3711209382) {
__t122 = gopurs_runtime.Apply(pure_22_121, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_25.UnsafePtr).V0})})
goto end_branch_122
} else {

}
}
{
if (v2_25.Type == 9 && v2_25.IntVal == 2465973597) {
__t122 = gopurs_runtime.Apply(k_24, (*Constructor_Data_Either_Right)(v2_25.UnsafePtr).V0)
goto end_branch_122
} else {

}
}
{
__t122 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_122:
return __t122
}))
})
})})}
}))
_ = __local_var_20_119
// TAST (Let): Bind1_21_123 -> *Constructor_Control_Bind_Bind
Bind1_21_123 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_119, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_123
// TAST (Let): Applicative0_22_124 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_124 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_119, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_124
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_19_115)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_123.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_123.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_124.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_125, x_19)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_127 -> *Constructor_Control_Bind_Bind
Bind1_18_127 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_127
// TAST (Let): pure_19_128 -> gopurs_runtime.Value
pure_19_128 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_19_128
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Except_Trans_applyExceptT(dictMonad_3)))}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_127.V1), v_20, gopurs_runtime.Func(func(v2_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t129 gopurs_runtime.Value
{
if (v2_22.Type == 9 && v2_22.IntVal == 3711209382) {
__t129 = gopurs_runtime.Apply(pure_19_128, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_22.UnsafePtr).V0})})
goto end_branch_129
} else {

}
}
{
if (v2_22.Type == 9 && v2_22.IntVal == 2465973597) {
__t129 = gopurs_runtime.Apply(k_21, (*Constructor_Data_Either_Right)(v2_22.UnsafePtr).V0)
goto end_branch_129
} else {

}
}
{
__t129 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_129:
return __t129
}))
})
})})}
}))
_ = __local_var_17_114
// TAST (Let): Bind1_18_130 -> *Constructor_Control_Bind_Bind
Bind1_18_130 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_114, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_130
// TAST (Let): Applicative0_19_131 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_131 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_114, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_131
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_16_110)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_130.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_130.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_131.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_108.V1), v_15, gopurs_runtime.Func(func(v2_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t132 gopurs_runtime.Value
{
if (v2_17.Type == 9 && v2_17.IntVal == 3711209382) {
__t132 = gopurs_runtime.Apply(pure_14_109, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_17.UnsafePtr).V0})})
goto end_branch_132
} else {

}
}
{
if (v2_17.Type == 9 && v2_17.IntVal == 2465973597) {
__t132 = gopurs_runtime.Apply(k_16, (*Constructor_Data_Either_Right)(v2_17.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_12_85
// TAST (Let): Bind1_13_133 -> *Constructor_Control_Bind_Bind
Bind1_13_133 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_85, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_133
// TAST (Let): Applicative0_14_134 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_134 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_85, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_134
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_11_81)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_133.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_133.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_134.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_79.V1), v_10, gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t135 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 3711209382) {
__t135 = gopurs_runtime.Apply(pure_9_80, gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(v2_12.UnsafePtr).V0})})
goto end_branch_135
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 2465973597) {
__t135 = gopurs_runtime.Apply(k_11, (*Constructor_Data_Either_Right)(v2_12.UnsafePtr).V0)
goto end_branch_135
} else {

}
}
{
__t135 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_135:
return __t135
}))
})
})})}
}))
_ = __local_var_7_56
// TAST (Let): Bind1_8_136 -> *Constructor_Control_Bind_Bind
Bind1_8_136 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_56, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_136
// TAST (Let): Applicative0_9_137 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_137 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_56, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_137
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_6_52)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_136.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_136.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_137.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_138, x_6)
})}
_ = applicativeExceptT1_5_51
// TAST (Let): Bind1_6_142 -> *Constructor_Control_Bind_Bind
Bind1_6_142 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_142
// TAST (Let): Applicative0_7_143 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_143 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_143
// TAST (Let): __local_var_8_145 -> gopurs_runtime.Value
__local_var_8_145 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_145
// TAST (Let): functorExceptT1_8_144 -> *Constructor_Data_Functor_Functor
functorExceptT1_8_144 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_146 -> gopurs_runtime.Value
__local_var_10_146 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_145, "map"), gopurs_runtime.Func(func(m_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t147 gopurs_runtime.Value
{
if (m_10.Type == 9 && m_10.IntVal == 3711209382) {
__t147 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, (*Constructor_Data_Either_Left)(m_10.UnsafePtr).V0})}
goto end_branch_147
} else {

}
}
{
if (m_10.Type == 9 && m_10.IntVal == 2465973597) {
__t147 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Either_Right)(m_10.UnsafePtr).V0)})}
goto end_branch_147
} else {

}
}
{
__t147 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_147:
return __t147
}))
_ = __local_var_10_146
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_146, v_11)
})
})))
_ = functorExceptT1_8_144
// TAST (Let): altExceptT2_6_141 -> *Constructor_Control_Alt_Alt
altExceptT2_6_141 := &Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorExceptT1_8_144)}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_142.V1), v_9, gopurs_runtime.Func(func(rm_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t150 gopurs_runtime.Value
{
if (rm_11.Type == 9 && rm_11.IntVal == 2465973597) {
__t150 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_143.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rm_11.UnsafePtr).V0})})
goto end_branch_150
} else {

}
}
{
if (rm_11.Type == 9 && rm_11.IntVal == 3711209382) {
// TAST (Let): __local_var_12_148 -> gopurs_runtime.Value
__local_var_12_148 := (*Constructor_Data_Either_Left)(rm_11.UnsafePtr).V0
_ = __local_var_12_148
__t150 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_142.V1), v1_10, gopurs_runtime.Func(func(rn_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t149 gopurs_runtime.Value
{
if (rn_13.Type == 9 && rn_13.IntVal == 2465973597) {
__t149 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_143.V1), gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, (*Constructor_Data_Either_Right)(rn_13.UnsafePtr).V0})})
goto end_branch_149
} else {

}
}
{
if (rn_13.Type == 9 && rn_13.IntVal == 3711209382) {
__t149 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_143.V1), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), __local_var_12_148, (*Constructor_Data_Either_Left)(rn_13.UnsafePtr).V0)})})
goto end_branch_149
} else {

}
}
{
__t149 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_149:
return __t149
}))
goto end_branch_150
} else {

}
}
{
__t150 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_150:
return __t150
}))
})
})}
_ = altExceptT2_6_141
// TAST (Let): plusExceptT2_6_140 -> *Constructor_Control_Plus_Plus
plusExceptT2_6_140 := &Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altExceptT2_6_141)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_3, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, mempty_1_0})})}
_ = plusExceptT2_6_140
// TAST (Let): alternativeExceptT2_5_50 -> *Constructor_Control_Alternative_Alternative
alternativeExceptT2_5_50 := &Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeExceptT1_5_51)}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusExceptT2_6_140)}
})}
_ = alternativeExceptT2_5_50
return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(&Constructor_Control_MonadPlus_MonadPlus{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeExceptT2_5_50)}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadExceptT1_4_2)}
})})}
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



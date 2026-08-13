package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Maybe_Trans_MaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_MaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_MaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_MaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_MaybeT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_MaybeT(x_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_MaybeT
}

var cache_Control_Monad_Maybe_Trans_runMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_runMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_runMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_runMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_runMaybeT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_runMaybeT(v_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_runMaybeT
}

var cache_Control_Monad_Maybe_Trans_newtypeMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_newtypeMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_newtypeMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_newtypeMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_newtypeMaybeT = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(&Constructor_Data_Newtype_Newtype{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})}
	})
	return cache_Control_Monad_Maybe_Trans_newtypeMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadTransMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadTransMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadTransMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadTransMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadTransMaybeT = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), a_3, gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_4})})
}))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
})})}
	})
	return cache_Control_Monad_Maybe_Trans_monadTransMaybeT
}

var cache_Control_Monad_Maybe_Trans_lift gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_lift sync.Once
func Get_Control_Monad_Maybe_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_lift.Do(func() {
		cache_Control_Monad_Maybe_Trans_lift = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Control_Monad_Maybe_Trans_lift
}

var cache_Control_Monad_Maybe_Trans_mapMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_mapMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_mapMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_mapMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_mapMaybeT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_mapMaybeT(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_mapMaybeT
}

var cache_Control_Monad_Maybe_Trans_functorMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_functorMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_functorMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_functorMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_functorMaybeT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_functorMaybeT(dictFunctor_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_functorMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadMaybeT
}

var cache_Control_Monad_Maybe_Trans_bindMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_bindMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_bindMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_bindMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_bindMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_bindMaybeT
}

var cache_Control_Monad_Maybe_Trans_applyMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_applyMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_applyMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_applyMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_applyMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_applyMaybeT
}

var cache_Control_Monad_Maybe_Trans_applicativeMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_applicativeMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_applicativeMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_applicativeMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_applicativeMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_applicativeMaybeT
}

var cache_Control_Monad_Maybe_Trans_semigroupMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_semigroupMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_semigroupMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_semigroupMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_semigroupMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_semigroupMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_semigroupMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadAskMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadAskMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadAskMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadAskMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadAskMaybeT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadAskMaybeT(dictMonadAsk_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadAskMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadReaderMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadReaderMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadReaderMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadReaderMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadReaderMaybeT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadReaderMaybeT(dictMonadReader_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadReaderMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadContMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadContMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadContMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadContMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadContMaybeT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadContMaybeT(dictMonadCont_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadContMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadEffectMaybe gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadEffectMaybe sync.Once
func Get_Control_Monad_Maybe_Trans_monadEffectMaybe() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadEffectMaybe.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadEffectMaybe = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadEffectMaybe(dictMonadEffect_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadEffectMaybe
}

var cache_Control_Monad_Maybe_Trans_monadRecMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadRecMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadRecMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadRecMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadRecMaybeT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadRecMaybeT(dictMonadRec_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadRecMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadStateMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadStateMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadStateMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadStateMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadStateMaybeT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadStateMaybeT(dictMonadState_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadStateMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadTellMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadTellMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadTellMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadTellMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadTellMaybeT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadTellMaybeT(dictMonadTell_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadTellMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadWriterMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadWriterMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadWriterMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadWriterMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadWriterMaybeT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadWriterMaybeT(dictMonadWriter_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadWriterMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadThrowMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadThrowMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadThrowMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadThrowMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadThrowMaybeT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadThrowMaybeT(dictMonadThrow_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadThrowMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadErrorMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadErrorMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadErrorMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadErrorMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadErrorMaybeT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadErrorMaybeT(dictMonadError_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadErrorMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadSTMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadSTMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadSTMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadSTMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadSTMaybeT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadSTMaybeT(dictMonadST_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadSTMaybeT
}

var cache_Control_Monad_Maybe_Trans_monoidMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monoidMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monoidMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monoidMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monoidMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monoidMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monoidMaybeT
}

var cache_Control_Monad_Maybe_Trans_altMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_altMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_altMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_altMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_altMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_altMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_altMaybeT
}

var cache_Control_Monad_Maybe_Trans_plusMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_plusMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_plusMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_plusMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_plusMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_plusMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_plusMaybeT
}

var cache_Control_Monad_Maybe_Trans_alternativeMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_alternativeMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_alternativeMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_alternativeMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_alternativeMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_alternativeMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_alternativeMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadPlusMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadPlusMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadPlusMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadPlusMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadPlusMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_monadPlusMaybeT(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_monadPlusMaybeT
}

var cache_Control_Monad_Maybe_Trans_mapMaybeT__4176318367 gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_mapMaybeT__4176318367 sync.Once
func Get_Control_Monad_Maybe_Trans_mapMaybeT__4176318367() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_mapMaybeT__4176318367.Do(func() {
		cache_Control_Monad_Maybe_Trans_mapMaybeT__4176318367 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_mapMaybeT__4176318367(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_mapMaybeT__4176318367
}

var cache_Control_Monad_Maybe_Trans_mapMaybeT__1721503071 gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_mapMaybeT__1721503071 sync.Once
func Get_Control_Monad_Maybe_Trans_mapMaybeT__1721503071() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_mapMaybeT__1721503071.Do(func() {
		cache_Control_Monad_Maybe_Trans_mapMaybeT__1721503071 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_mapMaybeT__1721503071(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_mapMaybeT__1721503071
}

var cache_Control_Monad_Maybe_Trans_mapMaybeT__1878923231 gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_mapMaybeT__1878923231 sync.Once
func Get_Control_Monad_Maybe_Trans_mapMaybeT__1878923231() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_mapMaybeT__1878923231.Do(func() {
		cache_Control_Monad_Maybe_Trans_mapMaybeT__1878923231 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_mapMaybeT__1878923231(f_0_box, v_1_box)
})
	})
	return cache_Control_Monad_Maybe_Trans_mapMaybeT__1878923231
}

var cache_Control_Monad_Maybe_Trans_monadTransMaybeT__3775352453 gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadTransMaybeT__3775352453 sync.Once
func Get_Control_Monad_Maybe_Trans_monadTransMaybeT__3775352453() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadTransMaybeT__3775352453.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadTransMaybeT__3775352453 = gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Trans_Class_MonadTrans{1, gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), a_3, gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_4})})
}))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
})})}
	})
	return cache_Control_Monad_Maybe_Trans_monadTransMaybeT__3775352453
}

func Call_Control_Monad_Maybe_Trans_MaybeT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Maybe_Trans_runMaybeT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Monad_Maybe_Trans_lift(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_2
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), a_3, gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_4})})
}))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}

func Call_Control_Monad_Maybe_Trans_mapMaybeT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Maybe_Trans_functorMaybeT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(&Constructor_Data_Functor_Functor{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *Constructor_Data_Maybe_Just
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_1, (*Constructor_Data_Maybe_Just)(v1_3.UnsafePtr).V0)}
goto end_branch_0
} else {

}
}
{
__t0 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)}
}), v_2)
})
})})}
}

func Call_Control_Monad_Maybe_Trans_monadMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Call_Control_Monad_Maybe_Trans_bindMaybeT(dictMonad_0)))}
})})}
}

func Call_Control_Monad_Maybe_Trans_bindMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_3, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 930809136 && v1_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_2
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 930809136 && v1_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply(f_4, (*Constructor_Data_Maybe_Just)(v1_5.UnsafePtr).V0)
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

func Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorMaybeT1_1_0 -> *Constructor_Data_Functor_Functor
functorMaybeT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Maybe_Just)(v1_4.UnsafePtr).V0)}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), v_3)
})
})))
_ = functorMaybeT1_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_4 -> *Constructor_Control_Bind_Bind
Bind1_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_4
// TAST (Let): Applicative0_4_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_5
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_4.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_5.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_6
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t6 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_2_3
// TAST (Let): Bind1_3_7 -> *Constructor_Control_Bind_Bind
Bind1_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_7
// TAST (Let): Applicative0_4_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_1_0)}
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_7.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_7.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_8.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
})})}
}

func Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_10 -> gopurs_runtime.Value
__local_var_1_10 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_10
// TAST (Let): __local_var_1_9 -> gopurs_runtime.Value
__local_var_1_9 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_10, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_2})})
})
_ = __local_var_1_9
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): functorMaybeT1_2_0 -> *Constructor_Data_Functor_Functor
functorMaybeT1_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
if (v1_5.Type == 9 && v1_5.IntVal == 930809136 && v1_5.UnsafePtr != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Maybe_Just)(v1_5.UnsafePtr).V0)}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), v_4)
})
})))
_ = functorMaybeT1_2_0
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_4 -> *Constructor_Control_Bind_Bind
Bind1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_4
// TAST (Let): Applicative0_5_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_5
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_5.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_6
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t6 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_3_3
// TAST (Let): Bind1_4_7 -> *Constructor_Control_Bind_Bind
Bind1_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_7
// TAST (Let): Applicative0_5_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_8
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_2_0)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_7.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_7.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_8.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_9, x_2)
})})}
}

func Call_Control_Monad_Maybe_Trans_semigroupMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): functorMaybeT1_1_1 -> *Constructor_Data_Functor_Functor
functorMaybeT1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "map"), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_4.Type == 9 && v1_4.IntVal == 930809136 && v1_4.UnsafePtr != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Maybe_Just)(v1_4.UnsafePtr).V0)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}), v_3)
})
})))
_ = functorMaybeT1_1_1
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_15 -> gopurs_runtime.Value
__local_var_3_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_15
// TAST (Let): __local_var_3_14 -> gopurs_runtime.Value
__local_var_3_14 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_15, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_14
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): functorMaybeT1_4_5 -> *Constructor_Data_Functor_Functor
functorMaybeT1_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_6, "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t7 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)}
goto end_branch_7
} else {

}
}
{
__t7 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
}), v_6)
})
})))
_ = functorMaybeT1_4_5
// TAST (Let): __local_var_5_8 -> gopurs_runtime.Value
__local_var_5_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_9 -> *Constructor_Control_Bind_Bind
Bind1_6_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_9
// TAST (Let): Applicative0_7_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_10
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_9.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_10.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_11
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t11 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_5_8
// TAST (Let): Bind1_6_12 -> *Constructor_Control_Bind_Bind
Bind1_6_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_12
// TAST (Let): Applicative0_7_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_13
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_4_5)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_12.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_12.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_13.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_14, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_16 -> *Constructor_Control_Bind_Bind
Bind1_3_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_16
// TAST (Let): Applicative0_4_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_17
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_16.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t18 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_17.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_18
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t18 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
_ = __local_var_2_4
// TAST (Let): Bind1_3_19 -> *Constructor_Control_Bind_Bind
Bind1_3_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_19
// TAST (Let): Applicative0_4_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_20
// TAST (Let): applyMaybeT1_1_0 -> *Constructor_Control_Apply_Apply
applyMaybeT1_1_0 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_1_1)}
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_19.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_19.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_20.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
})}
_ = applyMaybeT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_21 -> *Constructor_Data_Functor_Functor
Functor0_3_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyMaybeT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_21
// TAST (Let): __local_var_4_22 -> gopurs_runtime.Value
__local_var_4_22 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_22
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyMaybeT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_21.V0), __local_var_4_22, a_5), b_6)
})
})})}
})
}

func Call_Control_Monad_Maybe_Trans_monadAskMaybeT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadMaybeT1_1_0 -> *Constructor_Control_Monad_Monad
monadMaybeT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_21
// TAST (Let): __local_var_3_20 -> gopurs_runtime.Value
__local_var_3_20 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_21, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_20
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorMaybeT1_4_2 -> *Constructor_Data_Functor_Functor
functorMaybeT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), v_6)
})
})))
_ = functorMaybeT1_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_6 -> *Constructor_Control_Bind_Bind
Bind1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_6
// TAST (Let): Applicative0_7_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_7
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_9
// TAST (Let): functorMaybeT1_9_8 -> *Constructor_Data_Functor_Functor
functorMaybeT1_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_9, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Maybe_Just
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t10 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)}
goto end_branch_10
} else {

}
}
{
__t10 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}), v_11)
})
})))
_ = functorMaybeT1_9_8
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_12 -> *Constructor_Control_Bind_Bind
Bind1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_12
// TAST (Let): Applicative0_12_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_13
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_12.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_13.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_14
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t14 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_11
// TAST (Let): Bind1_11_15 -> *Constructor_Control_Bind_Bind
Bind1_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_15
// TAST (Let): Applicative0_12_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_16
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_9_8)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_16.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_6.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t17 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_7.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_17
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t17 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
_ = __local_var_5_5
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_4_2)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_20, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_22 -> *Constructor_Control_Bind_Bind
Bind1_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_22
// TAST (Let): Applicative0_4_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_23
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_25 -> gopurs_runtime.Value
__local_var_6_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_25
// TAST (Let): functorMaybeT1_6_24 -> *Constructor_Data_Functor_Functor
functorMaybeT1_6_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_25, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 *Constructor_Data_Maybe_Just
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t26 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)}
goto end_branch_26
} else {

}
}
{
__t26 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_26:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t26)}
}), v_8)
})
})))
_ = functorMaybeT1_6_24
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_38 -> gopurs_runtime.Value
__local_var_8_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_38
// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
__local_var_8_37 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_38, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_37
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_29
// TAST (Let): functorMaybeT1_9_28 -> *Constructor_Data_Functor_Functor
functorMaybeT1_9_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_29, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 *Constructor_Data_Maybe_Just
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t30 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)}
goto end_branch_30
} else {

}
}
{
__t30 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_30:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t30)}
}), v_11)
})
})))
_ = functorMaybeT1_9_28
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_32 -> *Constructor_Control_Bind_Bind
Bind1_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_32
// TAST (Let): Applicative0_12_33 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_33
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_32.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t34 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_33.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_34
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t34 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_31
// TAST (Let): Bind1_11_35 -> *Constructor_Control_Bind_Bind
Bind1_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_35
// TAST (Let): Applicative0_12_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_36
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_9_28)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_36.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_37, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_39 -> *Constructor_Control_Bind_Bind
Bind1_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_39
// TAST (Let): Applicative0_9_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_40
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_39.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t41 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_40.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_41
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t41 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_7_27
// TAST (Let): Bind1_8_42 -> *Constructor_Control_Bind_Bind
Bind1_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_42
// TAST (Let): Applicative0_9_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_43
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_6_24)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_43.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_22.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t44 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_23.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_44
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t44 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
})))
_ = monadMaybeT1_1_0
// TAST (Let): __local_var_2_45 -> *Constructor_Control_Monad_Monad
__local_var_2_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_2_45
// TAST (Let): Applicative0_3_46 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_45.V0), gopurs_runtime.Value{}))
_ = Applicative0_3_46
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Reader_Class_MonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_1_0)}
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_45.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_0, "ask"), gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_46.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_4})})
}))})}
}

func Call_Control_Monad_Maybe_Trans_monadReaderMaybeT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): monadMaybeT1_2_2 -> *Constructor_Control_Monad_Monad
monadMaybeT1_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_83 -> gopurs_runtime.Value
__local_var_4_83 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_83
// TAST (Let): __local_var_4_82 -> gopurs_runtime.Value
__local_var_4_82 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_83, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_82
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorMaybeT1_5_4 -> *Constructor_Data_Functor_Functor
functorMaybeT1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)}
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
}), v_7)
})
})))
_ = functorMaybeT1_5_4
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_27
// TAST (Let): __local_var_7_26 -> gopurs_runtime.Value
__local_var_7_26 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_27, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_8})})
})
_ = __local_var_7_26
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_9
// TAST (Let): functorMaybeT1_8_8 -> *Constructor_Data_Functor_Functor
functorMaybeT1_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_9, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Maybe_Just
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t10 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)}
goto end_branch_10
} else {

}
}
{
__t10 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}), v_10)
})
})))
_ = functorMaybeT1_8_8
// TAST (Let): __local_var_9_11 -> gopurs_runtime.Value
__local_var_9_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_12 -> *Constructor_Control_Bind_Bind
Bind1_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_12
// TAST (Let): Applicative0_11_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_13
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_15 -> gopurs_runtime.Value
__local_var_13_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_15
// TAST (Let): functorMaybeT1_13_14 -> *Constructor_Data_Functor_Functor
functorMaybeT1_13_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_15, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 *Constructor_Data_Maybe_Just
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t16 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)}
goto end_branch_16
} else {

}
}
{
__t16 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t16)}
}), v_15)
})
})))
_ = functorMaybeT1_13_14
// TAST (Let): __local_var_14_17 -> gopurs_runtime.Value
__local_var_14_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_18 -> *Constructor_Control_Bind_Bind
Bind1_15_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_18
// TAST (Let): Applicative0_16_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_19
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_18.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t20 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_19.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_20
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t20 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
_ = __local_var_14_17
// TAST (Let): Bind1_15_21 -> *Constructor_Control_Bind_Bind
Bind1_15_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_21
// TAST (Let): Applicative0_16_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_22
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_13_14)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_21.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_21.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_22.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_12.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t23 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_13.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_23
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t23 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
_ = __local_var_9_11
// TAST (Let): Bind1_10_24 -> *Constructor_Control_Bind_Bind
Bind1_10_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_24
// TAST (Let): Applicative0_11_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_25
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_8_8)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_24.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_24.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_25.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_26, x_8)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_28 -> *Constructor_Control_Bind_Bind
Bind1_7_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_28
// TAST (Let): Applicative0_8_29 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_29
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_31
// TAST (Let): functorMaybeT1_10_30 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_30 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_31, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t32 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_32
} else {

}
}
{
__t32 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_32:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t32)}
}), v_12)
})
})))
_ = functorMaybeT1_10_30
// TAST (Let): __local_var_11_33 -> gopurs_runtime.Value
__local_var_11_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_53 -> gopurs_runtime.Value
__local_var_12_53 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_53
// TAST (Let): __local_var_12_52 -> gopurs_runtime.Value
__local_var_12_52 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_53, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_13})})
})
_ = __local_var_12_52
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_35 -> gopurs_runtime.Value
__local_var_13_35 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_35
// TAST (Let): functorMaybeT1_13_34 -> *Constructor_Data_Functor_Functor
functorMaybeT1_13_34 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_35, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 *Constructor_Data_Maybe_Just
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t36 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)}
goto end_branch_36
} else {

}
}
{
__t36 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_36:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t36)}
}), v_15)
})
})))
_ = functorMaybeT1_13_34
// TAST (Let): __local_var_14_37 -> gopurs_runtime.Value
__local_var_14_37 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_38 -> *Constructor_Control_Bind_Bind
Bind1_15_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_38
// TAST (Let): Applicative0_16_39 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_39
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_41 -> gopurs_runtime.Value
__local_var_18_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_41
// TAST (Let): functorMaybeT1_18_40 -> *Constructor_Data_Functor_Functor
functorMaybeT1_18_40 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_41, "map"), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 *Constructor_Data_Maybe_Just
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t42 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)}
goto end_branch_42
} else {

}
}
{
__t42 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_42:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t42)}
}), v_20)
})
})))
_ = functorMaybeT1_18_40
// TAST (Let): __local_var_19_43 -> gopurs_runtime.Value
__local_var_19_43 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_44 -> *Constructor_Control_Bind_Bind
Bind1_20_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_44
// TAST (Let): Applicative0_21_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_45
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_44.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t46 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_45.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_46
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t46 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
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
_ = __local_var_19_43
// TAST (Let): Bind1_20_47 -> *Constructor_Control_Bind_Bind
Bind1_20_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_43, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_47
// TAST (Let): Applicative0_21_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_43, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_48
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_18_40)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_47.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_47.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_48.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_38.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t49 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t49 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_39.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_49
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t49 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
}))
_ = __local_var_14_37
// TAST (Let): Bind1_15_50 -> *Constructor_Control_Bind_Bind
Bind1_15_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_50
// TAST (Let): Applicative0_16_51 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_51
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_13_34)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_50.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_50.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_51.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_52, x_13)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_54 -> *Constructor_Control_Bind_Bind
Bind1_12_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_54
// TAST (Let): Applicative0_13_55 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_55
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_57 -> gopurs_runtime.Value
__local_var_15_57 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_57
// TAST (Let): functorMaybeT1_15_56 -> *Constructor_Data_Functor_Functor
functorMaybeT1_15_56 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_57, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t58 *Constructor_Data_Maybe_Just
{
if (v1_18.Type == 9 && v1_18.IntVal == 930809136 && v1_18.UnsafePtr != nil) {
__t58 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_18.UnsafePtr).V0)}
goto end_branch_58
} else {

}
}
{
__t58 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_58:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t58)}
}), v_17)
})
})))
_ = functorMaybeT1_15_56
// TAST (Let): __local_var_16_59 -> gopurs_runtime.Value
__local_var_16_59 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_70 -> gopurs_runtime.Value
__local_var_17_70 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_70
// TAST (Let): __local_var_17_69 -> gopurs_runtime.Value
__local_var_17_69 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_70, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_18})})
})
_ = __local_var_17_69
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_61 -> gopurs_runtime.Value
__local_var_18_61 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_61
// TAST (Let): functorMaybeT1_18_60 -> *Constructor_Data_Functor_Functor
functorMaybeT1_18_60 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_61, "map"), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t62 *Constructor_Data_Maybe_Just
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t62 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)}
goto end_branch_62
} else {

}
}
{
__t62 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_62:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t62)}
}), v_20)
})
})))
_ = functorMaybeT1_18_60
// TAST (Let): __local_var_19_63 -> gopurs_runtime.Value
__local_var_19_63 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_64 -> *Constructor_Control_Bind_Bind
Bind1_20_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_64
// TAST (Let): Applicative0_21_65 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_65
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_64.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t66 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t66 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_65.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_66
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t66 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_19_63
// TAST (Let): Bind1_20_67 -> *Constructor_Control_Bind_Bind
Bind1_20_67 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_63, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_67
// TAST (Let): Applicative0_21_68 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_68 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_63, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_68
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_18_60)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_67.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_67.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_68.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_69, x_18)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_71 -> *Constructor_Control_Bind_Bind
Bind1_17_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_71
// TAST (Let): Applicative0_18_72 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_72
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_71.V1), v_19, gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t73 gopurs_runtime.Value
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr == nil) {
__t73 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_72.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_73
} else {

}
}
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t73 = gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_16_59
// TAST (Let): Bind1_17_74 -> *Constructor_Control_Bind_Bind
Bind1_17_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_59, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_74
// TAST (Let): Applicative0_18_75 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_59, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_75
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_15_56)}
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_74.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_74.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_75.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_54.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t76 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t76 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_55.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_76
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t76 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_33
// TAST (Let): Bind1_12_77 -> *Constructor_Control_Bind_Bind
Bind1_12_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_77
// TAST (Let): Applicative0_13_78 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_78
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_30)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_77.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_77.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_78.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_28.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t79 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t79 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_29.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_79
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t79 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_6_7
// TAST (Let): Bind1_7_80 -> *Constructor_Control_Bind_Bind
Bind1_7_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_80
// TAST (Let): Applicative0_8_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_81
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_5_4)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_80.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_80.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_81.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_82, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_84 -> *Constructor_Control_Bind_Bind
Bind1_4_84 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_84
// TAST (Let): Applicative0_5_85 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_85 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_85
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_87 -> gopurs_runtime.Value
__local_var_7_87 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_87
// TAST (Let): functorMaybeT1_7_86 -> *Constructor_Data_Functor_Functor
functorMaybeT1_7_86 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_87, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t88 *Constructor_Data_Maybe_Just
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t88 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)}
goto end_branch_88
} else {

}
}
{
__t88 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_88:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t88)}
}), v_9)
})
})))
_ = functorMaybeT1_7_86
// TAST (Let): __local_var_8_89 -> gopurs_runtime.Value
__local_var_8_89 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_140 -> gopurs_runtime.Value
__local_var_9_140 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_140
// TAST (Let): __local_var_9_139 -> gopurs_runtime.Value
__local_var_9_139 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_140, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_139
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_91 -> gopurs_runtime.Value
__local_var_10_91 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_91
// TAST (Let): functorMaybeT1_10_90 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_90 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_91, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t92 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t92 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_92
} else {

}
}
{
__t92 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_92:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t92)}
}), v_12)
})
})))
_ = functorMaybeT1_10_90
// TAST (Let): __local_var_11_93 -> gopurs_runtime.Value
__local_var_11_93 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_113 -> gopurs_runtime.Value
__local_var_12_113 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_113
// TAST (Let): __local_var_12_112 -> gopurs_runtime.Value
__local_var_12_112 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_113, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_13})})
})
_ = __local_var_12_112
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_95 -> gopurs_runtime.Value
__local_var_13_95 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_95
// TAST (Let): functorMaybeT1_13_94 -> *Constructor_Data_Functor_Functor
functorMaybeT1_13_94 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_95, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t96 *Constructor_Data_Maybe_Just
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t96 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)}
goto end_branch_96
} else {

}
}
{
__t96 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_96:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t96)}
}), v_15)
})
})))
_ = functorMaybeT1_13_94
// TAST (Let): __local_var_14_97 -> gopurs_runtime.Value
__local_var_14_97 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_98 -> *Constructor_Control_Bind_Bind
Bind1_15_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_98
// TAST (Let): Applicative0_16_99 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_99 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_99
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_101 -> gopurs_runtime.Value
__local_var_18_101 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_101
// TAST (Let): functorMaybeT1_18_100 -> *Constructor_Data_Functor_Functor
functorMaybeT1_18_100 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_101, "map"), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t102 *Constructor_Data_Maybe_Just
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t102 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)}
goto end_branch_102
} else {

}
}
{
__t102 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_102:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t102)}
}), v_20)
})
})))
_ = functorMaybeT1_18_100
// TAST (Let): __local_var_19_103 -> gopurs_runtime.Value
__local_var_19_103 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_104 -> *Constructor_Control_Bind_Bind
Bind1_20_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_104
// TAST (Let): Applicative0_21_105 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_105 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_105
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_104.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t106 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t106 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_105.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_106
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t106 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
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
_ = __local_var_19_103
// TAST (Let): Bind1_20_107 -> *Constructor_Control_Bind_Bind
Bind1_20_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_103, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_107
// TAST (Let): Applicative0_21_108 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_103, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_108
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_18_100)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_107.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_107.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_108.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_98.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t109 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t109 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_99.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_109
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t109 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
_ = __local_var_14_97
// TAST (Let): Bind1_15_110 -> *Constructor_Control_Bind_Bind
Bind1_15_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_97, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_110
// TAST (Let): Applicative0_16_111 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_97, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_111
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_13_94)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_110.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_110.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_111.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_112, x_13)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_114 -> *Constructor_Control_Bind_Bind
Bind1_12_114 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_114
// TAST (Let): Applicative0_13_115 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_115 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_115
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_117 -> gopurs_runtime.Value
__local_var_15_117 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_117
// TAST (Let): functorMaybeT1_15_116 -> *Constructor_Data_Functor_Functor
functorMaybeT1_15_116 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_117, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t118 *Constructor_Data_Maybe_Just
{
if (v1_18.Type == 9 && v1_18.IntVal == 930809136 && v1_18.UnsafePtr != nil) {
__t118 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_18.UnsafePtr).V0)}
goto end_branch_118
} else {

}
}
{
__t118 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_118:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t118)}
}), v_17)
})
})))
_ = functorMaybeT1_15_116
// TAST (Let): __local_var_16_119 -> gopurs_runtime.Value
__local_var_16_119 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_130 -> gopurs_runtime.Value
__local_var_17_130 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_130
// TAST (Let): __local_var_17_129 -> gopurs_runtime.Value
__local_var_17_129 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_130, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_18})})
})
_ = __local_var_17_129
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_121 -> gopurs_runtime.Value
__local_var_18_121 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_121
// TAST (Let): functorMaybeT1_18_120 -> *Constructor_Data_Functor_Functor
functorMaybeT1_18_120 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_121, "map"), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t122 *Constructor_Data_Maybe_Just
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t122 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)}
goto end_branch_122
} else {

}
}
{
__t122 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_122:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t122)}
}), v_20)
})
})))
_ = functorMaybeT1_18_120
// TAST (Let): __local_var_19_123 -> gopurs_runtime.Value
__local_var_19_123 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_124 -> *Constructor_Control_Bind_Bind
Bind1_20_124 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_124
// TAST (Let): Applicative0_21_125 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_125 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_125
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_124.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t126 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t126 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_125.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_126
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t126 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_19_123
// TAST (Let): Bind1_20_127 -> *Constructor_Control_Bind_Bind
Bind1_20_127 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_123, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_127
// TAST (Let): Applicative0_21_128 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_128 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_123, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_128
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_18_120)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_127.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_127.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_128.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_129, x_18)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_131 -> *Constructor_Control_Bind_Bind
Bind1_17_131 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_131
// TAST (Let): Applicative0_18_132 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_132 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_132
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_131.V1), v_19, gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t133 gopurs_runtime.Value
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr == nil) {
__t133 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_132.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_133
} else {

}
}
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t133 = gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_16_119
// TAST (Let): Bind1_17_134 -> *Constructor_Control_Bind_Bind
Bind1_17_134 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_119, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_134
// TAST (Let): Applicative0_18_135 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_135 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_119, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_135
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_15_116)}
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_134.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_134.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_135.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_114.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t136 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t136 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_115.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_136
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t136 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_93
// TAST (Let): Bind1_12_137 -> *Constructor_Control_Bind_Bind
Bind1_12_137 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_93, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_137
// TAST (Let): Applicative0_13_138 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_138 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_93, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_138
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_90)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_137.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_137.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_138.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_139, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_141 -> *Constructor_Control_Bind_Bind
Bind1_9_141 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_141
// TAST (Let): Applicative0_10_142 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_142 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_142
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_144 -> gopurs_runtime.Value
__local_var_12_144 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_144
// TAST (Let): functorMaybeT1_12_143 -> *Constructor_Data_Functor_Functor
functorMaybeT1_12_143 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_12_144, "map"), gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t145 *Constructor_Data_Maybe_Just
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t145 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)}
goto end_branch_145
} else {

}
}
{
__t145 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_145:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t145)}
}), v_14)
})
})))
_ = functorMaybeT1_12_143
// TAST (Let): __local_var_13_146 -> gopurs_runtime.Value
__local_var_13_146 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_157 -> gopurs_runtime.Value
__local_var_14_157 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_14_157
// TAST (Let): __local_var_14_156 -> gopurs_runtime.Value
__local_var_14_156 := gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_157, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_15})})
})
_ = __local_var_14_156
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_148 -> gopurs_runtime.Value
__local_var_15_148 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_148
// TAST (Let): functorMaybeT1_15_147 -> *Constructor_Data_Functor_Functor
functorMaybeT1_15_147 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_148, "map"), gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t149 *Constructor_Data_Maybe_Just
{
if (v1_18.Type == 9 && v1_18.IntVal == 930809136 && v1_18.UnsafePtr != nil) {
__t149 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_18.UnsafePtr).V0)}
goto end_branch_149
} else {

}
}
{
__t149 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_149:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t149)}
}), v_17)
})
})))
_ = functorMaybeT1_15_147
// TAST (Let): __local_var_16_150 -> gopurs_runtime.Value
__local_var_16_150 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_151 -> *Constructor_Control_Bind_Bind
Bind1_17_151 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_151
// TAST (Let): Applicative0_18_152 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_152 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_152
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_151.V1), v_19, gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t153 gopurs_runtime.Value
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr == nil) {
__t153 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_152.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_153
} else {

}
}
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t153 = gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)
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
_ = __local_var_16_150
// TAST (Let): Bind1_17_154 -> *Constructor_Control_Bind_Bind
Bind1_17_154 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_150, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_154
// TAST (Let): Applicative0_18_155 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_155 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_150, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_155
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_15_147)}
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_154.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_154.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_155.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_156, x_15)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_14_158 -> *Constructor_Control_Bind_Bind
Bind1_14_158 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_158
// TAST (Let): Applicative0_15_159 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_159 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_159
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)))}
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_158.V1), v_16, gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t160 gopurs_runtime.Value
{
if (v1_18.Type == 9 && v1_18.IntVal == 930809136 && v1_18.UnsafePtr == nil) {
__t160 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_159.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_160
} else {

}
}
{
if (v1_18.Type == 9 && v1_18.IntVal == 930809136 && v1_18.UnsafePtr != nil) {
__t160 = gopurs_runtime.Apply(f_17, (*Constructor_Data_Maybe_Just)(v1_18.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_13_146
// TAST (Let): Bind1_14_161 -> *Constructor_Control_Bind_Bind
Bind1_14_161 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_146, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_161
// TAST (Let): Applicative0_15_162 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_162 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_146, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_162
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_12_143)}
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_161.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_161.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_162.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_141.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t163 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t163 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_142.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_163
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t163 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_8_89
// TAST (Let): Bind1_9_164 -> *Constructor_Control_Bind_Bind
Bind1_9_164 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_89, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_164
// TAST (Let): Applicative0_10_165 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_165 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_89, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_165
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_7_86)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_164.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_164.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_165.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_84.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t166 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t166 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_85.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_166
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t166 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
})})}
})))
_ = monadMaybeT1_2_2
// TAST (Let): __local_var_3_167 -> *Constructor_Control_Monad_Monad
__local_var_3_167 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_167
// TAST (Let): Applicative0_4_168 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_168 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_167.V0), gopurs_runtime.Value{}))
_ = Applicative0_4_168
// TAST (Let): monadAskMaybeT1_1_0 -> *Constructor_Control_Monad_Reader_Class_MonadAsk
monadAskMaybeT1_1_0 := &Constructor_Control_Monad_Reader_Class_MonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_2_2)}
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_167.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(__local_var_1_1, "ask"), gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_168.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_5})})
}))}
_ = monadAskMaybeT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Reader_Class_MonadReader{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(monadAskMaybeT1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_169 -> gopurs_runtime.Value
__local_var_3_169 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_2)
_ = __local_var_3_169
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_169, v_4)
})
})})}
}

func Call_Control_Monad_Maybe_Trans_monadContMaybeT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadMaybeT1_1_0 -> *Constructor_Control_Monad_Monad
monadMaybeT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_21
// TAST (Let): __local_var_3_20 -> gopurs_runtime.Value
__local_var_3_20 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_21, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_20
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorMaybeT1_4_2 -> *Constructor_Data_Functor_Functor
functorMaybeT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), v_6)
})
})))
_ = functorMaybeT1_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_6 -> *Constructor_Control_Bind_Bind
Bind1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_6
// TAST (Let): Applicative0_7_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_7
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_9
// TAST (Let): functorMaybeT1_9_8 -> *Constructor_Data_Functor_Functor
functorMaybeT1_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_9, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Maybe_Just
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t10 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)}
goto end_branch_10
} else {

}
}
{
__t10 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}), v_11)
})
})))
_ = functorMaybeT1_9_8
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_12 -> *Constructor_Control_Bind_Bind
Bind1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_12
// TAST (Let): Applicative0_12_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_13
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_12.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_13.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_14
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t14 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_11
// TAST (Let): Bind1_11_15 -> *Constructor_Control_Bind_Bind
Bind1_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_15
// TAST (Let): Applicative0_12_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_16
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_9_8)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_16.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_6.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t17 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_7.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_17
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t17 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
_ = __local_var_5_5
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_4_2)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_20, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_22 -> *Constructor_Control_Bind_Bind
Bind1_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_22
// TAST (Let): Applicative0_4_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_23
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_25 -> gopurs_runtime.Value
__local_var_6_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_25
// TAST (Let): functorMaybeT1_6_24 -> *Constructor_Data_Functor_Functor
functorMaybeT1_6_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_25, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 *Constructor_Data_Maybe_Just
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t26 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)}
goto end_branch_26
} else {

}
}
{
__t26 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_26:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t26)}
}), v_8)
})
})))
_ = functorMaybeT1_6_24
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_38 -> gopurs_runtime.Value
__local_var_8_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_38
// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
__local_var_8_37 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_38, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_37
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_29
// TAST (Let): functorMaybeT1_9_28 -> *Constructor_Data_Functor_Functor
functorMaybeT1_9_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_29, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 *Constructor_Data_Maybe_Just
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t30 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)}
goto end_branch_30
} else {

}
}
{
__t30 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_30:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t30)}
}), v_11)
})
})))
_ = functorMaybeT1_9_28
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_32 -> *Constructor_Control_Bind_Bind
Bind1_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_32
// TAST (Let): Applicative0_12_33 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_33
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_32.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t34 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_33.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_34
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t34 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_31
// TAST (Let): Bind1_11_35 -> *Constructor_Control_Bind_Bind
Bind1_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_35
// TAST (Let): Applicative0_12_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_36
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_9_28)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_36.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_37, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_39 -> *Constructor_Control_Bind_Bind
Bind1_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_39
// TAST (Let): Applicative0_9_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_40
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)))}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_39.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t41 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_40.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_41
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t41 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_7_27
// TAST (Let): Bind1_8_42 -> *Constructor_Control_Bind_Bind
Bind1_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_42
// TAST (Let): Applicative0_9_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_43
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_6_24)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_43.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_22.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t44 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_23.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_44
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t44 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
})))
_ = monadMaybeT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Cont_Class_MonadCont{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_4})})
}))
}))
})})}
}

func Call_Control_Monad_Maybe_Trans_monadEffectMaybe(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadMaybeT1_2_1 -> *Constructor_Control_Monad_Monad
monadMaybeT1_2_1 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_21
// TAST (Let): __local_var_3_20 -> gopurs_runtime.Value
__local_var_3_20 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_21, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_20
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorMaybeT1_4_2 -> *Constructor_Data_Functor_Functor
functorMaybeT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), v_6)
})
})))
_ = functorMaybeT1_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_6 -> *Constructor_Control_Bind_Bind
Bind1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_6
// TAST (Let): Applicative0_7_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_7
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_9
// TAST (Let): functorMaybeT1_9_8 -> *Constructor_Data_Functor_Functor
functorMaybeT1_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_9, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Maybe_Just
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t10 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)}
goto end_branch_10
} else {

}
}
{
__t10 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}), v_11)
})
})))
_ = functorMaybeT1_9_8
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_12 -> *Constructor_Control_Bind_Bind
Bind1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_12
// TAST (Let): Applicative0_12_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_13
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_12.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_13.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_14
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t14 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_11
// TAST (Let): Bind1_11_15 -> *Constructor_Control_Bind_Bind
Bind1_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_15
// TAST (Let): Applicative0_12_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_16
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_9_8)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_16.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_6.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t17 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_7.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_17
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t17 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
_ = __local_var_5_5
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_4_2)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_20, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_22 -> *Constructor_Control_Bind_Bind
Bind1_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_22
// TAST (Let): Applicative0_4_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_23
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_25 -> gopurs_runtime.Value
__local_var_6_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_25
// TAST (Let): functorMaybeT1_6_24 -> *Constructor_Data_Functor_Functor
functorMaybeT1_6_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_25, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 *Constructor_Data_Maybe_Just
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t26 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)}
goto end_branch_26
} else {

}
}
{
__t26 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_26:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t26)}
}), v_8)
})
})))
_ = functorMaybeT1_6_24
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_38 -> gopurs_runtime.Value
__local_var_8_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_38
// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
__local_var_8_37 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_38, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_37
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_29
// TAST (Let): functorMaybeT1_9_28 -> *Constructor_Data_Functor_Functor
functorMaybeT1_9_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_29, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 *Constructor_Data_Maybe_Just
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t30 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)}
goto end_branch_30
} else {

}
}
{
__t30 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_30:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t30)}
}), v_11)
})
})))
_ = functorMaybeT1_9_28
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_32 -> *Constructor_Control_Bind_Bind
Bind1_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_32
// TAST (Let): Applicative0_12_33 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_33
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_32.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t34 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_33.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_34
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t34 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_31
// TAST (Let): Bind1_11_35 -> *Constructor_Control_Bind_Bind
Bind1_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_35
// TAST (Let): Applicative0_12_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_36
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_9_28)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_36.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_37, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_39 -> *Constructor_Control_Bind_Bind
Bind1_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_39
// TAST (Let): Applicative0_9_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_40
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_39.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t41 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_40.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_41
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t41 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_7_27
// TAST (Let): Bind1_8_42 -> *Constructor_Control_Bind_Bind
Bind1_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_42
// TAST (Let): Applicative0_9_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_43
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_6_24)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_43.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_22.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t44 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_23.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_44
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t44 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
})}
_ = monadMaybeT1_2_1
// TAST (Let): Bind1_3_47 -> *Constructor_Control_Bind_Bind
Bind1_3_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_47
// TAST (Let): Applicative0_4_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_48
// TAST (Let): __local_var_3_46 -> gopurs_runtime.Value
__local_var_3_46 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_47.V1), a_5, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_48.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_6})})
}))
})
_ = __local_var_3_46
// TAST (Let): __local_var_3_45 -> gopurs_runtime.Value
__local_var_3_45 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_46, x_4)
})
_ = __local_var_3_45
return gopurs_runtime.Value{Type: 9, IntVal: 2217729261, UnsafePtr: unsafe.Pointer(&Constructor_Effect_Class_MonadEffect{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_2_1)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_45, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
})})}
}

func Call_Control_Monad_Maybe_Trans_monadRecMaybeT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): monadMaybeT1_4_3 -> *Constructor_Control_Monad_Monad
monadMaybeT1_4_3 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_23 -> gopurs_runtime.Value
__local_var_5_23 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_23
// TAST (Let): __local_var_5_22 -> gopurs_runtime.Value
__local_var_5_22 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_23, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_6})})
})
_ = __local_var_5_22
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): functorMaybeT1_6_4 -> *Constructor_Data_Functor_Functor
functorMaybeT1_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_5, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 *Constructor_Data_Maybe_Just
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t6 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)}
goto end_branch_6
} else {

}
}
{
__t6 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
}), v_8)
})
})))
_ = functorMaybeT1_6_4
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_8 -> *Constructor_Control_Bind_Bind
Bind1_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_8
// TAST (Let): Applicative0_9_9 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_9
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorMaybeT1_11_10 -> *Constructor_Data_Functor_Functor
functorMaybeT1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 *Constructor_Data_Maybe_Just
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t12 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)}
goto end_branch_12
} else {

}
}
{
__t12 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t12)}
}), v_13)
})
})))
_ = functorMaybeT1_11_10
// TAST (Let): __local_var_12_13 -> gopurs_runtime.Value
__local_var_12_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_14 -> *Constructor_Control_Bind_Bind
Bind1_13_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_14
// TAST (Let): Applicative0_14_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_15
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_14.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t16 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_15.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_16
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t16 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
_ = __local_var_12_13
// TAST (Let): Bind1_13_17 -> *Constructor_Control_Bind_Bind
Bind1_13_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_17
// TAST (Let): Applicative0_14_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_11_10)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_17.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_17.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_18.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_8.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t19 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_9.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_19
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t19 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
_ = __local_var_7_7
// TAST (Let): Bind1_8_20 -> *Constructor_Control_Bind_Bind
Bind1_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_20
// TAST (Let): Applicative0_9_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_21
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_6_4)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_20.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_20.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_21.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_22, x_6)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_24 -> *Constructor_Control_Bind_Bind
Bind1_5_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_24
// TAST (Let): Applicative0_6_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_25
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_27 -> gopurs_runtime.Value
__local_var_8_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_27
// TAST (Let): functorMaybeT1_8_26 -> *Constructor_Data_Functor_Functor
functorMaybeT1_8_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_27, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 *Constructor_Data_Maybe_Just
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t28 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)}
goto end_branch_28
} else {

}
}
{
__t28 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_28:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t28)}
}), v_10)
})
})))
_ = functorMaybeT1_8_26
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_40 -> gopurs_runtime.Value
__local_var_10_40 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_40
// TAST (Let): __local_var_10_39 -> gopurs_runtime.Value
__local_var_10_39 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_40, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_11})})
})
_ = __local_var_10_39
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_31 -> gopurs_runtime.Value
__local_var_11_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_31
// TAST (Let): functorMaybeT1_11_30 -> *Constructor_Data_Functor_Functor
functorMaybeT1_11_30 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_31, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 *Constructor_Data_Maybe_Just
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t32 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)}
goto end_branch_32
} else {

}
}
{
__t32 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_32:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t32)}
}), v_13)
})
})))
_ = functorMaybeT1_11_30
// TAST (Let): __local_var_12_33 -> gopurs_runtime.Value
__local_var_12_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_34 -> *Constructor_Control_Bind_Bind
Bind1_13_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_34
// TAST (Let): Applicative0_14_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_35
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_34.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t36 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_35.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_36
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t36 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_12_33
// TAST (Let): Bind1_13_37 -> *Constructor_Control_Bind_Bind
Bind1_13_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_37
// TAST (Let): Applicative0_14_38 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_38
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_11_30)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_37.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_37.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_38.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_39, x_11)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_41 -> *Constructor_Control_Bind_Bind
Bind1_10_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_41
// TAST (Let): Applicative0_11_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_42
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_41.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t43 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_42.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_43
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t43 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_9_29
// TAST (Let): Bind1_10_44 -> *Constructor_Control_Bind_Bind
Bind1_10_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_44
// TAST (Let): Applicative0_11_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_45
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_8_26)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_44.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_44.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_45.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_24.V1), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr == nil) {
__t46 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_25.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_46
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t46 = gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)
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
})}
_ = monadMaybeT1_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_MonadRec{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_4_3)}
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_47 -> gopurs_runtime.Value
__local_var_6_47 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(f_5, a_6), gopurs_runtime.Func(func(m_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t51 *Constructor_Control_Monad_Rec_Class_Done
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 930809136 && m_prime_7.UnsafePtr == nil) {
__t51 = &Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_51
} else {

}
}
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 930809136 && m_prime_7.UnsafePtr != nil) {
var __t50 gopurs_runtime.Value
{
var __t_tag_48 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(m_prime_7.UnsafePtr).V0
if (__t_tag_48.Type == 9 && __t_tag_48.IntVal == 525585346) {
__t50 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Maybe_Just)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}
goto end_branch_50
} else {

}
}
{
var __t_tag_49 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(m_prime_7.UnsafePtr).V0
if (__t_tag_49.Type == 9 && __t_tag_49.IntVal == 60402430) {
__t50 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Maybe_Just)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}})}
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
})})}
}

func Call_Control_Monad_Maybe_Trans_monadStateMaybeT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
// TAST (Let): Monad0_1_0 -> *Constructor_Control_Monad_Monad
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): monadMaybeT1_2_1 -> *Constructor_Control_Monad_Monad
monadMaybeT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_22 -> gopurs_runtime.Value
__local_var_4_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_22
// TAST (Let): __local_var_4_21 -> gopurs_runtime.Value
__local_var_4_21 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_22, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_21
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorMaybeT1_5_3 -> *Constructor_Data_Functor_Functor
functorMaybeT1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)}
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}), v_7)
})
})))
_ = functorMaybeT1_5_3
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_7 -> *Constructor_Control_Bind_Bind
Bind1_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_7
// TAST (Let): Applicative0_8_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_8
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_10
// TAST (Let): functorMaybeT1_10_9 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_10, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t11 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_11
} else {

}
}
{
__t11 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t11)}
}), v_12)
})
})))
_ = functorMaybeT1_10_9
// TAST (Let): __local_var_11_12 -> gopurs_runtime.Value
__local_var_11_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_13 -> *Constructor_Control_Bind_Bind
Bind1_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_13
// TAST (Let): Applicative0_13_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_14
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_13.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_14.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_15
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t15 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
_ = __local_var_11_12
// TAST (Let): Bind1_12_16 -> *Constructor_Control_Bind_Bind
Bind1_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_16
// TAST (Let): Applicative0_13_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_17
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_9)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_17.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_7.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t18 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_8.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_18
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t18 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
_ = __local_var_6_6
// TAST (Let): Bind1_7_19 -> *Constructor_Control_Bind_Bind
Bind1_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_19
// TAST (Let): Applicative0_8_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_20
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_5_3)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_20.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_21, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_23 -> *Constructor_Control_Bind_Bind
Bind1_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_23
// TAST (Let): Applicative0_5_24 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_24
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_26 -> gopurs_runtime.Value
__local_var_7_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_26
// TAST (Let): functorMaybeT1_7_25 -> *Constructor_Data_Functor_Functor
functorMaybeT1_7_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_26, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 *Constructor_Data_Maybe_Just
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t27 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)}
goto end_branch_27
} else {

}
}
{
__t27 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t27)}
}), v_9)
})
})))
_ = functorMaybeT1_7_25
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_39 -> gopurs_runtime.Value
__local_var_9_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_39
// TAST (Let): __local_var_9_38 -> gopurs_runtime.Value
__local_var_9_38 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_39, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_38
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_30 -> gopurs_runtime.Value
__local_var_10_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_30
// TAST (Let): functorMaybeT1_10_29 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_29 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_30, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t31 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_31
} else {

}
}
{
__t31 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_31:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t31)}
}), v_12)
})
})))
_ = functorMaybeT1_10_29
// TAST (Let): __local_var_11_32 -> gopurs_runtime.Value
__local_var_11_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_33 -> *Constructor_Control_Bind_Bind
Bind1_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_33
// TAST (Let): Applicative0_13_34 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_34
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_33.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t35 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_34.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_35
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t35 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_32
// TAST (Let): Bind1_12_36 -> *Constructor_Control_Bind_Bind
Bind1_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_36
// TAST (Let): Applicative0_13_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_37
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_29)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_37.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_38, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_40 -> *Constructor_Control_Bind_Bind
Bind1_9_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_40
// TAST (Let): Applicative0_10_41 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_41
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_40.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t42 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_41.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_42
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t42 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_8_28
// TAST (Let): Bind1_9_43 -> *Constructor_Control_Bind_Bind
Bind1_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_43
// TAST (Let): Applicative0_10_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_44
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_7_25)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_44.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_23.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t45 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_24.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_45
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t45 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
})))
_ = monadMaybeT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 2100320995, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_State_Class_MonadState{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_2_1)}
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Applicative0_4_46 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_4_46
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), f_3), gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_46.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_5})})
}))
})})}
}

func Call_Control_Monad_Maybe_Trans_monadTellMaybeT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
// TAST (Let): Monad1_1_0 -> gopurs_runtime.Value
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
// TAST (Let): monadMaybeT1_3_2 -> *Constructor_Control_Monad_Monad
monadMaybeT1_3_2 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_22 -> gopurs_runtime.Value
__local_var_4_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_22
// TAST (Let): __local_var_4_21 -> gopurs_runtime.Value
__local_var_4_21 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_22, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_21
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorMaybeT1_5_3 -> *Constructor_Data_Functor_Functor
functorMaybeT1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)}
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}), v_7)
})
})))
_ = functorMaybeT1_5_3
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_7 -> *Constructor_Control_Bind_Bind
Bind1_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_7
// TAST (Let): Applicative0_8_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_8
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_10
// TAST (Let): functorMaybeT1_10_9 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_10, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t11 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_11
} else {

}
}
{
__t11 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t11)}
}), v_12)
})
})))
_ = functorMaybeT1_10_9
// TAST (Let): __local_var_11_12 -> gopurs_runtime.Value
__local_var_11_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_13 -> *Constructor_Control_Bind_Bind
Bind1_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_13
// TAST (Let): Applicative0_13_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_14
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_13.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_14.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_15
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t15 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
_ = __local_var_11_12
// TAST (Let): Bind1_12_16 -> *Constructor_Control_Bind_Bind
Bind1_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_16
// TAST (Let): Applicative0_13_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_17
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_9)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_17.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_7.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t18 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_8.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_18
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t18 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
_ = __local_var_6_6
// TAST (Let): Bind1_7_19 -> *Constructor_Control_Bind_Bind
Bind1_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_19
// TAST (Let): Applicative0_8_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_20
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_5_3)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_20.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_21, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_23 -> *Constructor_Control_Bind_Bind
Bind1_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_23
// TAST (Let): Applicative0_5_24 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_24
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_26 -> gopurs_runtime.Value
__local_var_7_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_26
// TAST (Let): functorMaybeT1_7_25 -> *Constructor_Data_Functor_Functor
functorMaybeT1_7_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_26, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 *Constructor_Data_Maybe_Just
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t27 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)}
goto end_branch_27
} else {

}
}
{
__t27 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t27)}
}), v_9)
})
})))
_ = functorMaybeT1_7_25
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_39 -> gopurs_runtime.Value
__local_var_9_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_39
// TAST (Let): __local_var_9_38 -> gopurs_runtime.Value
__local_var_9_38 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_39, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_38
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_30 -> gopurs_runtime.Value
__local_var_10_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_30
// TAST (Let): functorMaybeT1_10_29 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_29 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_30, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t31 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_31
} else {

}
}
{
__t31 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_31:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t31)}
}), v_12)
})
})))
_ = functorMaybeT1_10_29
// TAST (Let): __local_var_11_32 -> gopurs_runtime.Value
__local_var_11_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_33 -> *Constructor_Control_Bind_Bind
Bind1_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_33
// TAST (Let): Applicative0_13_34 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_34
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_33.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t35 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_34.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_35
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t35 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_32
// TAST (Let): Bind1_12_36 -> *Constructor_Control_Bind_Bind
Bind1_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_36
// TAST (Let): Applicative0_13_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_37
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_29)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_37.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_38, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_40 -> *Constructor_Control_Bind_Bind
Bind1_9_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_40
// TAST (Let): Applicative0_10_41 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_41
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_1_0)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_40.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t42 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_41.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_42
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t42 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_8_28
// TAST (Let): Bind1_9_43 -> *Constructor_Control_Bind_Bind
Bind1_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_43
// TAST (Let): Applicative0_10_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_44
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_7_25)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_44.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_23.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t45 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_24.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_45
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t45 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
})}
_ = monadMaybeT1_3_2
// TAST (Let): Bind1_4_48 -> *Constructor_Control_Bind_Bind
Bind1_4_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_48
// TAST (Let): Applicative0_5_49 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_49
// TAST (Let): __local_var_4_47 -> gopurs_runtime.Value
__local_var_4_47 := gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_48.V1), a_6, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_49.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_7})})
}))
})
_ = __local_var_4_47
// TAST (Let): __local_var_4_46 -> gopurs_runtime.Value
__local_var_4_46 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_47, x_5)
})
_ = __local_var_4_46
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Writer_Class_MonadTell{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_3_2)}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_2_1)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_46, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
})})}
}

func Call_Control_Monad_Maybe_Trans_monadWriterMaybeT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): monadMaybeT1_9_9 -> *Constructor_Control_Monad_Monad
monadMaybeT1_9_9 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_89 -> gopurs_runtime.Value
__local_var_10_89 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_89
// TAST (Let): __local_var_10_88 -> gopurs_runtime.Value
__local_var_10_88 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_89, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_11})})
})
_ = __local_var_10_88
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorMaybeT1_11_10 -> *Constructor_Data_Functor_Functor
functorMaybeT1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 *Constructor_Data_Maybe_Just
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t12 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)}
goto end_branch_12
} else {

}
}
{
__t12 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t12)}
}), v_13)
})
})))
_ = functorMaybeT1_11_10
// TAST (Let): __local_var_12_13 -> gopurs_runtime.Value
__local_var_12_13 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_33 -> gopurs_runtime.Value
__local_var_13_33 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_33
// TAST (Let): __local_var_13_32 -> gopurs_runtime.Value
__local_var_13_32 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_33, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_14})})
})
_ = __local_var_13_32
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_15 -> gopurs_runtime.Value
__local_var_14_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_15
// TAST (Let): functorMaybeT1_14_14 -> *Constructor_Data_Functor_Functor
functorMaybeT1_14_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_15, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 *Constructor_Data_Maybe_Just
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t16 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)}
goto end_branch_16
} else {

}
}
{
__t16 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_16:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t16)}
}), v_16)
})
})))
_ = functorMaybeT1_14_14
// TAST (Let): __local_var_15_17 -> gopurs_runtime.Value
__local_var_15_17 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_18 -> *Constructor_Control_Bind_Bind
Bind1_16_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_18
// TAST (Let): Applicative0_17_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_19
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_21 -> gopurs_runtime.Value
__local_var_19_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_21
// TAST (Let): functorMaybeT1_19_20 -> *Constructor_Data_Functor_Functor
functorMaybeT1_19_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_21, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t22 *Constructor_Data_Maybe_Just
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t22 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)}
goto end_branch_22
} else {

}
}
{
__t22 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_22:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t22)}
}), v_21)
})
})))
_ = functorMaybeT1_19_20
// TAST (Let): __local_var_20_23 -> gopurs_runtime.Value
__local_var_20_23 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_24 -> *Constructor_Control_Bind_Bind
Bind1_21_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_24
// TAST (Let): Applicative0_22_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_25
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_24.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t26 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_25.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_26
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t26 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
_ = __local_var_20_23
// TAST (Let): Bind1_21_27 -> *Constructor_Control_Bind_Bind
Bind1_21_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_23, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_27
// TAST (Let): Applicative0_22_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_23, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_28
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_19_20)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_27.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_27.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_28.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_18.V1), v_18, gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr == nil) {
__t29 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_19.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_29
} else {

}
}
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr != nil) {
__t29 = gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_20.UnsafePtr).V0)
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
_ = __local_var_15_17
// TAST (Let): Bind1_16_30 -> *Constructor_Control_Bind_Bind
Bind1_16_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_17, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_30
// TAST (Let): Applicative0_17_31 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_17, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_31
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_14_14)}
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_30.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_30.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_31.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_32, x_14)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_34 -> *Constructor_Control_Bind_Bind
Bind1_13_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_34
// TAST (Let): Applicative0_14_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_35
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_37 -> gopurs_runtime.Value
__local_var_16_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_37
// TAST (Let): functorMaybeT1_16_36 -> *Constructor_Data_Functor_Functor
functorMaybeT1_16_36 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_37, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 *Constructor_Data_Maybe_Just
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t38 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)}
goto end_branch_38
} else {

}
}
{
__t38 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_38:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t38)}
}), v_18)
})
})))
_ = functorMaybeT1_16_36
// TAST (Let): __local_var_17_39 -> gopurs_runtime.Value
__local_var_17_39 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_59 -> gopurs_runtime.Value
__local_var_18_59 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_59
// TAST (Let): __local_var_18_58 -> gopurs_runtime.Value
__local_var_18_58 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_59, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_19})})
})
_ = __local_var_18_58
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_41 -> gopurs_runtime.Value
__local_var_19_41 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_41
// TAST (Let): functorMaybeT1_19_40 -> *Constructor_Data_Functor_Functor
functorMaybeT1_19_40 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_41, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 *Constructor_Data_Maybe_Just
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t42 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)}
goto end_branch_42
} else {

}
}
{
__t42 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_42:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t42)}
}), v_21)
})
})))
_ = functorMaybeT1_19_40
// TAST (Let): __local_var_20_43 -> gopurs_runtime.Value
__local_var_20_43 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_44 -> *Constructor_Control_Bind_Bind
Bind1_21_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_44
// TAST (Let): Applicative0_22_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_45
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_47 -> gopurs_runtime.Value
__local_var_24_47 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_47
// TAST (Let): functorMaybeT1_24_46 -> *Constructor_Data_Functor_Functor
functorMaybeT1_24_46 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_24_47, "map"), gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t48 *Constructor_Data_Maybe_Just
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t48 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_25, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)}
goto end_branch_48
} else {

}
}
{
__t48 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_48:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t48)}
}), v_26)
})
})))
_ = functorMaybeT1_24_46
// TAST (Let): __local_var_25_49 -> gopurs_runtime.Value
__local_var_25_49 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_50 -> *Constructor_Control_Bind_Bind
Bind1_26_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_50
// TAST (Let): Applicative0_27_51 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_51
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_50.V1), v_28, gopurs_runtime.Func(func(v1_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t52 gopurs_runtime.Value
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr == nil) {
__t52 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_51.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_52
} else {

}
}
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr != nil) {
__t52 = gopurs_runtime.Apply(f_29, (*Constructor_Data_Maybe_Just)(v1_30.UnsafePtr).V0)
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
_ = __local_var_25_49
// TAST (Let): Bind1_26_53 -> *Constructor_Control_Bind_Bind
Bind1_26_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_49, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_53
// TAST (Let): Applicative0_27_54 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_49, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_54
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_24_46)}
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_53.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_53.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_54.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_44.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t55 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t55 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_45.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_55
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t55 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
_ = __local_var_20_43
// TAST (Let): Bind1_21_56 -> *Constructor_Control_Bind_Bind
Bind1_21_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_43, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_56
// TAST (Let): Applicative0_22_57 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_43, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_57
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_19_40)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_56.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_56.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_57.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_58, x_19)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_60 -> *Constructor_Control_Bind_Bind
Bind1_18_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_60
// TAST (Let): Applicative0_19_61 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_61
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_63 -> gopurs_runtime.Value
__local_var_21_63 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_63
// TAST (Let): functorMaybeT1_21_62 -> *Constructor_Data_Functor_Functor
functorMaybeT1_21_62 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_21_63, "map"), gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t64 *Constructor_Data_Maybe_Just
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t64 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_22, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)}
goto end_branch_64
} else {

}
}
{
__t64 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_64:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t64)}
}), v_23)
})
})))
_ = functorMaybeT1_21_62
// TAST (Let): __local_var_22_65 -> gopurs_runtime.Value
__local_var_22_65 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_76 -> gopurs_runtime.Value
__local_var_23_76 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_23_76
// TAST (Let): __local_var_23_75 -> gopurs_runtime.Value
__local_var_23_75 := gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_76, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_24})})
})
_ = __local_var_23_75
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_67 -> gopurs_runtime.Value
__local_var_24_67 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_67
// TAST (Let): functorMaybeT1_24_66 -> *Constructor_Data_Functor_Functor
functorMaybeT1_24_66 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_24_67, "map"), gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t68 *Constructor_Data_Maybe_Just
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t68 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_25, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)}
goto end_branch_68
} else {

}
}
{
__t68 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_68:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t68)}
}), v_26)
})
})))
_ = functorMaybeT1_24_66
// TAST (Let): __local_var_25_69 -> gopurs_runtime.Value
__local_var_25_69 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_70 -> *Constructor_Control_Bind_Bind
Bind1_26_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_70
// TAST (Let): Applicative0_27_71 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_71
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_70.V1), v_28, gopurs_runtime.Func(func(v1_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t72 gopurs_runtime.Value
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr == nil) {
__t72 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_71.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_72
} else {

}
}
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr != nil) {
__t72 = gopurs_runtime.Apply(f_29, (*Constructor_Data_Maybe_Just)(v1_30.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_25_69
// TAST (Let): Bind1_26_73 -> *Constructor_Control_Bind_Bind
Bind1_26_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_69, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_73
// TAST (Let): Applicative0_27_74 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_69, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_74
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_24_66)}
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_73.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_73.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_74.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_75, x_24)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_77 -> *Constructor_Control_Bind_Bind
Bind1_23_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_77
// TAST (Let): Applicative0_24_78 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_78
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_77.V1), v_25, gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t79 gopurs_runtime.Value
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr == nil) {
__t79 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_78.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_79
} else {

}
}
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t79 = gopurs_runtime.Apply(f_26, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_22_65
// TAST (Let): Bind1_23_80 -> *Constructor_Control_Bind_Bind
Bind1_23_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_65, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_80
// TAST (Let): Applicative0_24_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_65, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_81
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_21_62)}
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_80.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_80.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_81.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_60.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t82 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t82 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_61.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_82
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t82 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_17_39
// TAST (Let): Bind1_18_83 -> *Constructor_Control_Bind_Bind
Bind1_18_83 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_39, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_83
// TAST (Let): Applicative0_19_84 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_84 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_39, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_84
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_16_36)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_83.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_83.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_84.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_34.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t85 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t85 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_35.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_85
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t85 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_12_13
// TAST (Let): Bind1_13_86 -> *Constructor_Control_Bind_Bind
Bind1_13_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_13, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_86
// TAST (Let): Applicative0_14_87 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_13, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_87
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_11_10)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_86.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_86.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_87.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_88, x_11)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_90 -> *Constructor_Control_Bind_Bind
Bind1_10_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_90
// TAST (Let): Applicative0_11_91 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_91
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_93 -> gopurs_runtime.Value
__local_var_13_93 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_93
// TAST (Let): functorMaybeT1_13_92 -> *Constructor_Data_Functor_Functor
functorMaybeT1_13_92 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_93, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t94 *Constructor_Data_Maybe_Just
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t94 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)}
goto end_branch_94
} else {

}
}
{
__t94 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_94:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t94)}
}), v_15)
})
})))
_ = functorMaybeT1_13_92
// TAST (Let): __local_var_14_95 -> gopurs_runtime.Value
__local_var_14_95 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_146 -> gopurs_runtime.Value
__local_var_15_146 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_15_146
// TAST (Let): __local_var_15_145 -> gopurs_runtime.Value
__local_var_15_145 := gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_146, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_16})})
})
_ = __local_var_15_145
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_97 -> gopurs_runtime.Value
__local_var_16_97 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_97
// TAST (Let): functorMaybeT1_16_96 -> *Constructor_Data_Functor_Functor
functorMaybeT1_16_96 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_97, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t98 *Constructor_Data_Maybe_Just
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t98 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)}
goto end_branch_98
} else {

}
}
{
__t98 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_98:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t98)}
}), v_18)
})
})))
_ = functorMaybeT1_16_96
// TAST (Let): __local_var_17_99 -> gopurs_runtime.Value
__local_var_17_99 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_119 -> gopurs_runtime.Value
__local_var_18_119 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_119
// TAST (Let): __local_var_18_118 -> gopurs_runtime.Value
__local_var_18_118 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_119, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_19})})
})
_ = __local_var_18_118
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_101 -> gopurs_runtime.Value
__local_var_19_101 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_101
// TAST (Let): functorMaybeT1_19_100 -> *Constructor_Data_Functor_Functor
functorMaybeT1_19_100 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_101, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t102 *Constructor_Data_Maybe_Just
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t102 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)}
goto end_branch_102
} else {

}
}
{
__t102 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_102:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t102)}
}), v_21)
})
})))
_ = functorMaybeT1_19_100
// TAST (Let): __local_var_20_103 -> gopurs_runtime.Value
__local_var_20_103 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_104 -> *Constructor_Control_Bind_Bind
Bind1_21_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_104
// TAST (Let): Applicative0_22_105 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_105 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_105
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_107 -> gopurs_runtime.Value
__local_var_24_107 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_107
// TAST (Let): functorMaybeT1_24_106 -> *Constructor_Data_Functor_Functor
functorMaybeT1_24_106 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_24_107, "map"), gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t108 *Constructor_Data_Maybe_Just
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t108 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_25, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)}
goto end_branch_108
} else {

}
}
{
__t108 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_108:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t108)}
}), v_26)
})
})))
_ = functorMaybeT1_24_106
// TAST (Let): __local_var_25_109 -> gopurs_runtime.Value
__local_var_25_109 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_110 -> *Constructor_Control_Bind_Bind
Bind1_26_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_110
// TAST (Let): Applicative0_27_111 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_111
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_110.V1), v_28, gopurs_runtime.Func(func(v1_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t112 gopurs_runtime.Value
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr == nil) {
__t112 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_111.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_112
} else {

}
}
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr != nil) {
__t112 = gopurs_runtime.Apply(f_29, (*Constructor_Data_Maybe_Just)(v1_30.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_25_109
// TAST (Let): Bind1_26_113 -> *Constructor_Control_Bind_Bind
Bind1_26_113 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_109, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_113
// TAST (Let): Applicative0_27_114 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_114 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_109, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_114
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_24_106)}
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_113.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_113.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_114.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_104.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t115 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t115 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_105.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_115
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t115 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_20_103
// TAST (Let): Bind1_21_116 -> *Constructor_Control_Bind_Bind
Bind1_21_116 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_103, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_116
// TAST (Let): Applicative0_22_117 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_117 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_103, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_117
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_19_100)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_116.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_116.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_117.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_118, x_19)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_120 -> *Constructor_Control_Bind_Bind
Bind1_18_120 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_120
// TAST (Let): Applicative0_19_121 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_121 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_121
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_123 -> gopurs_runtime.Value
__local_var_21_123 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_123
// TAST (Let): functorMaybeT1_21_122 -> *Constructor_Data_Functor_Functor
functorMaybeT1_21_122 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_21_123, "map"), gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t124 *Constructor_Data_Maybe_Just
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t124 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_22, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)}
goto end_branch_124
} else {

}
}
{
__t124 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_124:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t124)}
}), v_23)
})
})))
_ = functorMaybeT1_21_122
// TAST (Let): __local_var_22_125 -> gopurs_runtime.Value
__local_var_22_125 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_136 -> gopurs_runtime.Value
__local_var_23_136 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_23_136
// TAST (Let): __local_var_23_135 -> gopurs_runtime.Value
__local_var_23_135 := gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_136, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_24})})
})
_ = __local_var_23_135
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_127 -> gopurs_runtime.Value
__local_var_24_127 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_127
// TAST (Let): functorMaybeT1_24_126 -> *Constructor_Data_Functor_Functor
functorMaybeT1_24_126 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_24_127, "map"), gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t128 *Constructor_Data_Maybe_Just
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t128 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_25, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)}
goto end_branch_128
} else {

}
}
{
__t128 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_128:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t128)}
}), v_26)
})
})))
_ = functorMaybeT1_24_126
// TAST (Let): __local_var_25_129 -> gopurs_runtime.Value
__local_var_25_129 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_130 -> *Constructor_Control_Bind_Bind
Bind1_26_130 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_130
// TAST (Let): Applicative0_27_131 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_131 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_131
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_130.V1), v_28, gopurs_runtime.Func(func(v1_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t132 gopurs_runtime.Value
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr == nil) {
__t132 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_131.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_132
} else {

}
}
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr != nil) {
__t132 = gopurs_runtime.Apply(f_29, (*Constructor_Data_Maybe_Just)(v1_30.UnsafePtr).V0)
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
_ = __local_var_25_129
// TAST (Let): Bind1_26_133 -> *Constructor_Control_Bind_Bind
Bind1_26_133 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_129, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_133
// TAST (Let): Applicative0_27_134 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_134 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_129, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_134
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_24_126)}
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_133.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_133.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_134.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_135, x_24)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_137 -> *Constructor_Control_Bind_Bind
Bind1_23_137 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_137
// TAST (Let): Applicative0_24_138 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_138 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_138
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_137.V1), v_25, gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t139 gopurs_runtime.Value
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr == nil) {
__t139 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_138.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_139
} else {

}
}
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t139 = gopurs_runtime.Apply(f_26, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_22_125
// TAST (Let): Bind1_23_140 -> *Constructor_Control_Bind_Bind
Bind1_23_140 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_125, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_140
// TAST (Let): Applicative0_24_141 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_141 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_125, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_141
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_21_122)}
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_140.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_140.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_141.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_120.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t142 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t142 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_121.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_142
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t142 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_17_99
// TAST (Let): Bind1_18_143 -> *Constructor_Control_Bind_Bind
Bind1_18_143 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_99, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_143
// TAST (Let): Applicative0_19_144 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_144 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_99, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_144
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_16_96)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_143.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_143.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_144.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_145, x_16)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_147 -> *Constructor_Control_Bind_Bind
Bind1_15_147 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_147
// TAST (Let): Applicative0_16_148 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_148 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_148
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_150 -> gopurs_runtime.Value
__local_var_18_150 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_150
// TAST (Let): functorMaybeT1_18_149 -> *Constructor_Data_Functor_Functor
functorMaybeT1_18_149 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_150, "map"), gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t151 *Constructor_Data_Maybe_Just
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t151 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)}
goto end_branch_151
} else {

}
}
{
__t151 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_151:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t151)}
}), v_20)
})
})))
_ = functorMaybeT1_18_149
// TAST (Let): __local_var_19_152 -> gopurs_runtime.Value
__local_var_19_152 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_163 -> gopurs_runtime.Value
__local_var_20_163 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_20_163
// TAST (Let): __local_var_20_162 -> gopurs_runtime.Value
__local_var_20_162 := gopurs_runtime.Func(func(x_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_163, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_21})})
})
_ = __local_var_20_162
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_154 -> gopurs_runtime.Value
__local_var_21_154 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_154
// TAST (Let): functorMaybeT1_21_153 -> *Constructor_Data_Functor_Functor
functorMaybeT1_21_153 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_21_154, "map"), gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t155 *Constructor_Data_Maybe_Just
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t155 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_22, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)}
goto end_branch_155
} else {

}
}
{
__t155 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_155:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t155)}
}), v_23)
})
})))
_ = functorMaybeT1_21_153
// TAST (Let): __local_var_22_156 -> gopurs_runtime.Value
__local_var_22_156 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_157 -> *Constructor_Control_Bind_Bind
Bind1_23_157 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_157
// TAST (Let): Applicative0_24_158 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_158 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_158
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_157.V1), v_25, gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t159 gopurs_runtime.Value
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr == nil) {
__t159 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_158.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_159
} else {

}
}
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t159 = gopurs_runtime.Apply(f_26, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_22_156
// TAST (Let): Bind1_23_160 -> *Constructor_Control_Bind_Bind
Bind1_23_160 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_156, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_160
// TAST (Let): Applicative0_24_161 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_161 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_156, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_161
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_21_153)}
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_160.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_160.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_161.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_162, x_21)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_164 -> *Constructor_Control_Bind_Bind
Bind1_20_164 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_164
// TAST (Let): Applicative0_21_165 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_165 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_165
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)))}
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_164.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t166 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t166 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_165.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_166
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t166 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_19_152
// TAST (Let): Bind1_20_167 -> *Constructor_Control_Bind_Bind
Bind1_20_167 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_152, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_167
// TAST (Let): Applicative0_21_168 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_168 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_152, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_168
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_18_149)}
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_167.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_167.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_168.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_147.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t169 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t169 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_148.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_169
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t169 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
_ = __local_var_14_95
// TAST (Let): Bind1_15_170 -> *Constructor_Control_Bind_Bind
Bind1_15_170 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_95, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_170
// TAST (Let): Applicative0_16_171 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_171 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_95, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_171
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_13_92)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_170.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_170.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_171.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_90.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t172 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t172 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_91.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_172
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t172 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
})})}
})}
_ = monadMaybeT1_9_9
// TAST (Let): Bind1_10_175 -> *Constructor_Control_Bind_Bind
Bind1_10_175 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_175
// TAST (Let): Applicative0_11_176 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_176 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_176
// TAST (Let): __local_var_10_174 -> gopurs_runtime.Value
__local_var_10_174 := gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_175.V1), a_12, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_176.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_13})})
}))
})
_ = __local_var_10_174
// TAST (Let): __local_var_10_173 -> gopurs_runtime.Value
__local_var_10_173 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_174, x_11)
})
_ = __local_var_10_173
// TAST (Let): monadTellMaybeT1_7_6 -> *Constructor_Control_Monad_Writer_Class_MonadTell
monadTellMaybeT1_7_6 := &Constructor_Control_Monad_Writer_Class_MonadTell{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_9_9)}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Semigroup0_8_8)}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_173, gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "tell"), x_11))
})}
_ = monadTellMaybeT1_7_6
return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Writer_Class_MonadWriter{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(monadTellMaybeT1_7_6)}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Monoid0_6_5)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), v_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t178 *Constructor_Data_Maybe_Just
{
var __t_tag_177 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0)
if (__t_tag_177 != nil) {
__t178 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Maybe_Just)((*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1})}}
goto end_branch_178
} else {

}
}
{
__t178 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_178:
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t178)})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_8, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t179 *Constructor_Data_Tuple_Tuple
{
if (a_9.Type == 9 && a_9.IntVal == 930809136 && a_9.UnsafePtr == nil) {
__t179 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_10
})}
goto end_branch_179
} else {

}
}
{
if (a_9.Type == 9 && a_9.IntVal == 930809136 && a_9.UnsafePtr != nil) {
__t179 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(a_9.UnsafePtr).V0.UnsafePtr).V0})}, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(a_9.UnsafePtr).V0.UnsafePtr).V1}
goto end_branch_179
} else {

}
}
{
__t179 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_179:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t179)})
})))
})})}
}

func Call_Control_Monad_Maybe_Trans_monadThrowMaybeT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): Monad0_1_0 -> *Constructor_Control_Monad_Monad
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): monadMaybeT1_2_1 -> *Constructor_Control_Monad_Monad
monadMaybeT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_22 -> gopurs_runtime.Value
__local_var_4_22 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_22
// TAST (Let): __local_var_4_21 -> gopurs_runtime.Value
__local_var_4_21 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_22, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_21
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorMaybeT1_5_3 -> *Constructor_Data_Functor_Functor
functorMaybeT1_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)}
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}), v_7)
})
})))
_ = functorMaybeT1_5_3
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_7 -> *Constructor_Control_Bind_Bind
Bind1_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_7
// TAST (Let): Applicative0_8_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_8
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_10
// TAST (Let): functorMaybeT1_10_9 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_10, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t11 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_11
} else {

}
}
{
__t11 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t11)}
}), v_12)
})
})))
_ = functorMaybeT1_10_9
// TAST (Let): __local_var_11_12 -> gopurs_runtime.Value
__local_var_11_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_13 -> *Constructor_Control_Bind_Bind
Bind1_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_13
// TAST (Let): Applicative0_13_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_14
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_13.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_14.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_15
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t15 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
_ = __local_var_11_12
// TAST (Let): Bind1_12_16 -> *Constructor_Control_Bind_Bind
Bind1_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_16
// TAST (Let): Applicative0_13_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_17
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_9)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_16.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_17.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_7.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t18 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_8.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_18
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t18 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
_ = __local_var_6_6
// TAST (Let): Bind1_7_19 -> *Constructor_Control_Bind_Bind
Bind1_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_19
// TAST (Let): Applicative0_8_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_20
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_5_3)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_19.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_20.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_21, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_23 -> *Constructor_Control_Bind_Bind
Bind1_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_23
// TAST (Let): Applicative0_5_24 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_24
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_26 -> gopurs_runtime.Value
__local_var_7_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_26
// TAST (Let): functorMaybeT1_7_25 -> *Constructor_Data_Functor_Functor
functorMaybeT1_7_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_26, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 *Constructor_Data_Maybe_Just
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t27 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)}
goto end_branch_27
} else {

}
}
{
__t27 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t27)}
}), v_9)
})
})))
_ = functorMaybeT1_7_25
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_39 -> gopurs_runtime.Value
__local_var_9_39 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_39
// TAST (Let): __local_var_9_38 -> gopurs_runtime.Value
__local_var_9_38 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_39, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_38
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_30 -> gopurs_runtime.Value
__local_var_10_30 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_30
// TAST (Let): functorMaybeT1_10_29 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_29 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_30, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t31 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_31
} else {

}
}
{
__t31 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_31:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t31)}
}), v_12)
})
})))
_ = functorMaybeT1_10_29
// TAST (Let): __local_var_11_32 -> gopurs_runtime.Value
__local_var_11_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_33 -> *Constructor_Control_Bind_Bind
Bind1_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_33
// TAST (Let): Applicative0_13_34 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_34
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_33.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t35 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_34.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_35
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t35 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_32
// TAST (Let): Bind1_12_36 -> *Constructor_Control_Bind_Bind
Bind1_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_36
// TAST (Let): Applicative0_13_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_37
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_29)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_36.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_37.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_38, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_40 -> *Constructor_Control_Bind_Bind
Bind1_9_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_40
// TAST (Let): Applicative0_10_41 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_41
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_40.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t42 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_41.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_42
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t42 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_8_28
// TAST (Let): Bind1_9_43 -> *Constructor_Control_Bind_Bind
Bind1_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_43
// TAST (Let): Applicative0_10_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_44
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_7_25)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_43.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_44.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_23.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t45 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_24.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_45
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t45 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
})))
_ = monadMaybeT1_2_1
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Error_Class_MonadThrow{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_2_1)}
}), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Applicative0_4_46 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_4_46
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_3), gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_46.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_5})})
}))
})})}
}

func Call_Control_Monad_Maybe_Trans_monadErrorMaybeT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): monadMaybeT1_3_3 -> *Constructor_Control_Monad_Monad
monadMaybeT1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_84 -> gopurs_runtime.Value
__local_var_5_84 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_84
// TAST (Let): __local_var_5_83 -> gopurs_runtime.Value
__local_var_5_83 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_84, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_6})})
})
_ = __local_var_5_83
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_6
// TAST (Let): functorMaybeT1_6_5 -> *Constructor_Data_Functor_Functor
functorMaybeT1_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_6, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 *Constructor_Data_Maybe_Just
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t7 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)}
goto end_branch_7
} else {

}
}
{
__t7 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t7)}
}), v_8)
})
})))
_ = functorMaybeT1_6_5
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_28
// TAST (Let): __local_var_8_27 -> gopurs_runtime.Value
__local_var_8_27 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_28, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_27
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_10
// TAST (Let): functorMaybeT1_9_9 -> *Constructor_Data_Functor_Functor
functorMaybeT1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_10, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 *Constructor_Data_Maybe_Just
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t11 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)}
goto end_branch_11
} else {

}
}
{
__t11 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_11:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t11)}
}), v_11)
})
})))
_ = functorMaybeT1_9_9
// TAST (Let): __local_var_10_12 -> gopurs_runtime.Value
__local_var_10_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_13 -> *Constructor_Control_Bind_Bind
Bind1_11_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_13
// TAST (Let): Applicative0_12_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_14
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_16 -> gopurs_runtime.Value
__local_var_14_16 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_16
// TAST (Let): functorMaybeT1_14_15 -> *Constructor_Data_Functor_Functor
functorMaybeT1_14_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_16, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 *Constructor_Data_Maybe_Just
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t17 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)}
goto end_branch_17
} else {

}
}
{
__t17 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_17:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t17)}
}), v_16)
})
})))
_ = functorMaybeT1_14_15
// TAST (Let): __local_var_15_18 -> gopurs_runtime.Value
__local_var_15_18 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_19 -> *Constructor_Control_Bind_Bind
Bind1_16_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_19
// TAST (Let): Applicative0_17_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_20
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_19.V1), v_18, gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr == nil) {
__t21 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_20.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_21
} else {

}
}
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr != nil) {
__t21 = gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_20.UnsafePtr).V0)
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
_ = __local_var_15_18
// TAST (Let): Bind1_16_22 -> *Constructor_Control_Bind_Bind
Bind1_16_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_18, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_22
// TAST (Let): Applicative0_17_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_18, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_23
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_14_15)}
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_22.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_22.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_23.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_13.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t24 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_14.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_24
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t24 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
goto end_branch_24
} else {

}
}
{
__t24 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_24:
return __t24
}))
})
})})}
}))
_ = __local_var_10_12
// TAST (Let): Bind1_11_25 -> *Constructor_Control_Bind_Bind
Bind1_11_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_25
// TAST (Let): Applicative0_12_26 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_26
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_9_9)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_25.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_25.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_26.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_27, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_29 -> *Constructor_Control_Bind_Bind
Bind1_8_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_29
// TAST (Let): Applicative0_9_30 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_30
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_32 -> gopurs_runtime.Value
__local_var_11_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_32
// TAST (Let): functorMaybeT1_11_31 -> *Constructor_Data_Functor_Functor
functorMaybeT1_11_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_32, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 *Constructor_Data_Maybe_Just
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t33 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)}
goto end_branch_33
} else {

}
}
{
__t33 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_33:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t33)}
}), v_13)
})
})))
_ = functorMaybeT1_11_31
// TAST (Let): __local_var_12_34 -> gopurs_runtime.Value
__local_var_12_34 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_54 -> gopurs_runtime.Value
__local_var_13_54 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_54
// TAST (Let): __local_var_13_53 -> gopurs_runtime.Value
__local_var_13_53 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_54, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_14})})
})
_ = __local_var_13_53
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_36 -> gopurs_runtime.Value
__local_var_14_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_36
// TAST (Let): functorMaybeT1_14_35 -> *Constructor_Data_Functor_Functor
functorMaybeT1_14_35 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_36, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 *Constructor_Data_Maybe_Just
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t37 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)}
goto end_branch_37
} else {

}
}
{
__t37 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t37)}
}), v_16)
})
})))
_ = functorMaybeT1_14_35
// TAST (Let): __local_var_15_38 -> gopurs_runtime.Value
__local_var_15_38 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_39 -> *Constructor_Control_Bind_Bind
Bind1_16_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_39
// TAST (Let): Applicative0_17_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_40
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_42 -> gopurs_runtime.Value
__local_var_19_42 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_42
// TAST (Let): functorMaybeT1_19_41 -> *Constructor_Data_Functor_Functor
functorMaybeT1_19_41 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_42, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 *Constructor_Data_Maybe_Just
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t43 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)}
goto end_branch_43
} else {

}
}
{
__t43 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_43:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t43)}
}), v_21)
})
})))
_ = functorMaybeT1_19_41
// TAST (Let): __local_var_20_44 -> gopurs_runtime.Value
__local_var_20_44 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_45 -> *Constructor_Control_Bind_Bind
Bind1_21_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_45
// TAST (Let): Applicative0_22_46 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_46
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_45.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t47 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t47 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_46.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_47
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t47 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
_ = __local_var_20_44
// TAST (Let): Bind1_21_48 -> *Constructor_Control_Bind_Bind
Bind1_21_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_44, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_48
// TAST (Let): Applicative0_22_49 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_44, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_49
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_19_41)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_48.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_48.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_49.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_39.V1), v_18, gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t50 gopurs_runtime.Value
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr == nil) {
__t50 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_40.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_50
} else {

}
}
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr != nil) {
__t50 = gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_20.UnsafePtr).V0)
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
}))
_ = __local_var_15_38
// TAST (Let): Bind1_16_51 -> *Constructor_Control_Bind_Bind
Bind1_16_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_38, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_51
// TAST (Let): Applicative0_17_52 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_38, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_52
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_14_35)}
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_51.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_51.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_52.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_53, x_14)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_55 -> *Constructor_Control_Bind_Bind
Bind1_13_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_55
// TAST (Let): Applicative0_14_56 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_56
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_58 -> gopurs_runtime.Value
__local_var_16_58 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_58
// TAST (Let): functorMaybeT1_16_57 -> *Constructor_Data_Functor_Functor
functorMaybeT1_16_57 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_58, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t59 *Constructor_Data_Maybe_Just
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t59 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)}
goto end_branch_59
} else {

}
}
{
__t59 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_59:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t59)}
}), v_18)
})
})))
_ = functorMaybeT1_16_57
// TAST (Let): __local_var_17_60 -> gopurs_runtime.Value
__local_var_17_60 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_71 -> gopurs_runtime.Value
__local_var_18_71 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_71
// TAST (Let): __local_var_18_70 -> gopurs_runtime.Value
__local_var_18_70 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_71, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_19})})
})
_ = __local_var_18_70
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_62 -> gopurs_runtime.Value
__local_var_19_62 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_62
// TAST (Let): functorMaybeT1_19_61 -> *Constructor_Data_Functor_Functor
functorMaybeT1_19_61 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_62, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t63 *Constructor_Data_Maybe_Just
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t63 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)}
goto end_branch_63
} else {

}
}
{
__t63 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_63:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t63)}
}), v_21)
})
})))
_ = functorMaybeT1_19_61
// TAST (Let): __local_var_20_64 -> gopurs_runtime.Value
__local_var_20_64 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_65 -> *Constructor_Control_Bind_Bind
Bind1_21_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_65
// TAST (Let): Applicative0_22_66 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_66 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_66
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_65.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t67 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t67 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_66.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_67
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t67 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_20_64
// TAST (Let): Bind1_21_68 -> *Constructor_Control_Bind_Bind
Bind1_21_68 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_64, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_68
// TAST (Let): Applicative0_22_69 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_69 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_64, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_69
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_19_61)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_68.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_68.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_69.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_70, x_19)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_72 -> *Constructor_Control_Bind_Bind
Bind1_18_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_72
// TAST (Let): Applicative0_19_73 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_73
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_72.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t74 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t74 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_73.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_74
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t74 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
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
_ = __local_var_17_60
// TAST (Let): Bind1_18_75 -> *Constructor_Control_Bind_Bind
Bind1_18_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_60, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_75
// TAST (Let): Applicative0_19_76 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_60, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_76
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_16_57)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_75.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_75.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_76.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_55.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t77 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t77 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_56.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_77
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t77 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
_ = __local_var_12_34
// TAST (Let): Bind1_13_78 -> *Constructor_Control_Bind_Bind
Bind1_13_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_34, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_78
// TAST (Let): Applicative0_14_79 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_34, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_79
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_11_31)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_78.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_78.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_79.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_29.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t80 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t80 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_30.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_80
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t80 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
_ = __local_var_7_8
// TAST (Let): Bind1_8_81 -> *Constructor_Control_Bind_Bind
Bind1_8_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_81
// TAST (Let): Applicative0_9_82 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_82
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_6_5)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_81.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_81.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_82.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_83, x_6)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_85 -> *Constructor_Control_Bind_Bind
Bind1_5_85 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_85
// TAST (Let): Applicative0_6_86 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_86
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_88 -> gopurs_runtime.Value
__local_var_8_88 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_88
// TAST (Let): functorMaybeT1_8_87 -> *Constructor_Data_Functor_Functor
functorMaybeT1_8_87 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_88, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t89 *Constructor_Data_Maybe_Just
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t89 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)}
goto end_branch_89
} else {

}
}
{
__t89 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_89:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t89)}
}), v_10)
})
})))
_ = functorMaybeT1_8_87
// TAST (Let): __local_var_9_90 -> gopurs_runtime.Value
__local_var_9_90 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_141 -> gopurs_runtime.Value
__local_var_10_141 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_141
// TAST (Let): __local_var_10_140 -> gopurs_runtime.Value
__local_var_10_140 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_141, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_11})})
})
_ = __local_var_10_140
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_92 -> gopurs_runtime.Value
__local_var_11_92 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_92
// TAST (Let): functorMaybeT1_11_91 -> *Constructor_Data_Functor_Functor
functorMaybeT1_11_91 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_92, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t93 *Constructor_Data_Maybe_Just
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t93 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)}
goto end_branch_93
} else {

}
}
{
__t93 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_93:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t93)}
}), v_13)
})
})))
_ = functorMaybeT1_11_91
// TAST (Let): __local_var_12_94 -> gopurs_runtime.Value
__local_var_12_94 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_114 -> gopurs_runtime.Value
__local_var_13_114 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_114
// TAST (Let): __local_var_13_113 -> gopurs_runtime.Value
__local_var_13_113 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_114, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_14})})
})
_ = __local_var_13_113
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_96 -> gopurs_runtime.Value
__local_var_14_96 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_96
// TAST (Let): functorMaybeT1_14_95 -> *Constructor_Data_Functor_Functor
functorMaybeT1_14_95 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_96, "map"), gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t97 *Constructor_Data_Maybe_Just
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t97 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)}
goto end_branch_97
} else {

}
}
{
__t97 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_97:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t97)}
}), v_16)
})
})))
_ = functorMaybeT1_14_95
// TAST (Let): __local_var_15_98 -> gopurs_runtime.Value
__local_var_15_98 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_99 -> *Constructor_Control_Bind_Bind
Bind1_16_99 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_99
// TAST (Let): Applicative0_17_100 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_100 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_100
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_102 -> gopurs_runtime.Value
__local_var_19_102 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_102
// TAST (Let): functorMaybeT1_19_101 -> *Constructor_Data_Functor_Functor
functorMaybeT1_19_101 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_102, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t103 *Constructor_Data_Maybe_Just
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t103 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)}
goto end_branch_103
} else {

}
}
{
__t103 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_103:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t103)}
}), v_21)
})
})))
_ = functorMaybeT1_19_101
// TAST (Let): __local_var_20_104 -> gopurs_runtime.Value
__local_var_20_104 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_105 -> *Constructor_Control_Bind_Bind
Bind1_21_105 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_105
// TAST (Let): Applicative0_22_106 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_106 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_106
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_105.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t107 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t107 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_106.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_107
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t107 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
})
})})}
}))
_ = __local_var_20_104
// TAST (Let): Bind1_21_108 -> *Constructor_Control_Bind_Bind
Bind1_21_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_104, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_108
// TAST (Let): Applicative0_22_109 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_104, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_109
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_19_101)}
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_108.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_108.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_109.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_99.V1), v_18, gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t110 gopurs_runtime.Value
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr == nil) {
__t110 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_100.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_110
} else {

}
}
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr != nil) {
__t110 = gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_20.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_15_98
// TAST (Let): Bind1_16_111 -> *Constructor_Control_Bind_Bind
Bind1_16_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_98, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_111
// TAST (Let): Applicative0_17_112 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_112 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_98, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_112
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_14_95)}
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_111.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_111.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_112.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_113, x_14)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_115 -> *Constructor_Control_Bind_Bind
Bind1_13_115 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_115
// TAST (Let): Applicative0_14_116 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_116 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_116
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_118 -> gopurs_runtime.Value
__local_var_16_118 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_118
// TAST (Let): functorMaybeT1_16_117 -> *Constructor_Data_Functor_Functor
functorMaybeT1_16_117 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_118, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t119 *Constructor_Data_Maybe_Just
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t119 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)}
goto end_branch_119
} else {

}
}
{
__t119 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_119:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t119)}
}), v_18)
})
})))
_ = functorMaybeT1_16_117
// TAST (Let): __local_var_17_120 -> gopurs_runtime.Value
__local_var_17_120 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_131 -> gopurs_runtime.Value
__local_var_18_131 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_131
// TAST (Let): __local_var_18_130 -> gopurs_runtime.Value
__local_var_18_130 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_131, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_19})})
})
_ = __local_var_18_130
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_122 -> gopurs_runtime.Value
__local_var_19_122 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_122
// TAST (Let): functorMaybeT1_19_121 -> *Constructor_Data_Functor_Functor
functorMaybeT1_19_121 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_122, "map"), gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t123 *Constructor_Data_Maybe_Just
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t123 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)}
goto end_branch_123
} else {

}
}
{
__t123 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_123:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t123)}
}), v_21)
})
})))
_ = functorMaybeT1_19_121
// TAST (Let): __local_var_20_124 -> gopurs_runtime.Value
__local_var_20_124 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_125 -> *Constructor_Control_Bind_Bind
Bind1_21_125 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_125
// TAST (Let): Applicative0_22_126 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_126 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_126
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_125.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t127 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t127 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_126.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_127
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t127 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
_ = __local_var_20_124
// TAST (Let): Bind1_21_128 -> *Constructor_Control_Bind_Bind
Bind1_21_128 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_124, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_128
// TAST (Let): Applicative0_22_129 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_129 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_124, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_129
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_19_121)}
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
Bind1_18_132 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_132
// TAST (Let): Applicative0_19_133 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_133 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_133
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_132.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t134 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t134 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_133.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_134
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t134 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
goto end_branch_134
} else {

}
}
{
__t134 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_134:
return __t134
}))
})
})})}
}))
_ = __local_var_17_120
// TAST (Let): Bind1_18_135 -> *Constructor_Control_Bind_Bind
Bind1_18_135 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_120, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_135
// TAST (Let): Applicative0_19_136 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_136 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_120, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_136
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_16_117)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_135.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_135.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_136.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_115.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t137 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t137 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_116.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_137
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t137 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_12_94
// TAST (Let): Bind1_13_138 -> *Constructor_Control_Bind_Bind
Bind1_13_138 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_94, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_138
// TAST (Let): Applicative0_14_139 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_139 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_94, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_139
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_11_91)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_138.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_138.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_139.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_140, x_11)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_142 -> *Constructor_Control_Bind_Bind
Bind1_10_142 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_142
// TAST (Let): Applicative0_11_143 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_143 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_143
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_145 -> gopurs_runtime.Value
__local_var_13_145 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_145
// TAST (Let): functorMaybeT1_13_144 -> *Constructor_Data_Functor_Functor
functorMaybeT1_13_144 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_145, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t146 *Constructor_Data_Maybe_Just
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t146 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)}
goto end_branch_146
} else {

}
}
{
__t146 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_146:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t146)}
}), v_15)
})
})))
_ = functorMaybeT1_13_144
// TAST (Let): __local_var_14_147 -> gopurs_runtime.Value
__local_var_14_147 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_158 -> gopurs_runtime.Value
__local_var_15_158 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_15_158
// TAST (Let): __local_var_15_157 -> gopurs_runtime.Value
__local_var_15_157 := gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_158, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_16})})
})
_ = __local_var_15_157
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_149 -> gopurs_runtime.Value
__local_var_16_149 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_149
// TAST (Let): functorMaybeT1_16_148 -> *Constructor_Data_Functor_Functor
functorMaybeT1_16_148 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_149, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t150 *Constructor_Data_Maybe_Just
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t150 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)}
goto end_branch_150
} else {

}
}
{
__t150 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_150:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t150)}
}), v_18)
})
})))
_ = functorMaybeT1_16_148
// TAST (Let): __local_var_17_151 -> gopurs_runtime.Value
__local_var_17_151 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_152 -> *Constructor_Control_Bind_Bind
Bind1_18_152 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_152
// TAST (Let): Applicative0_19_153 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_153 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_153
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_152.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t154 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t154 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_153.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_154
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t154 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
goto end_branch_154
} else {

}
}
{
__t154 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_154:
return __t154
}))
})
})})}
}))
_ = __local_var_17_151
// TAST (Let): Bind1_18_155 -> *Constructor_Control_Bind_Bind
Bind1_18_155 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_151, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_155
// TAST (Let): Applicative0_19_156 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_156 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_151, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_156
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_16_148)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_155.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_155.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_156.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_157, x_16)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_159 -> *Constructor_Control_Bind_Bind
Bind1_15_159 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_159
// TAST (Let): Applicative0_16_160 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_160 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_160
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)))}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_159.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t161 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t161 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_160.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_161
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t161 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
goto end_branch_161
} else {

}
}
{
__t161 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_161:
return __t161
}))
})
})})}
}))
_ = __local_var_14_147
// TAST (Let): Bind1_15_162 -> *Constructor_Control_Bind_Bind
Bind1_15_162 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_147, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_162
// TAST (Let): Applicative0_16_163 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_163 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_147, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_163
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_13_144)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_162.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_162.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_163.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_142.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t164 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t164 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_143.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_164
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t164 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
goto end_branch_164
} else {

}
}
{
__t164 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_164:
return __t164
}))
})
})})}
}))
_ = __local_var_9_90
// TAST (Let): Bind1_10_165 -> *Constructor_Control_Bind_Bind
Bind1_10_165 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_90, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_165
// TAST (Let): Applicative0_11_166 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_166 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_90, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_166
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_8_87)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_165.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_165.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_166.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_85.V1), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t167 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr == nil) {
__t167 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_86.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_167
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t167 = gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)
goto end_branch_167
} else {

}
}
{
__t167 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_167:
return __t167
}))
})
})})}
})))
_ = monadMaybeT1_3_3
// TAST (Let): monadThrowMaybeT1_1_0 -> *Constructor_Control_Monad_Error_Class_MonadThrow
monadThrowMaybeT1_1_0 := &Constructor_Control_Monad_Error_Class_MonadThrow{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_3_3)}
}), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Applicative0_5_168 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_168 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_2_2.V0), gopurs_runtime.Value{}))
_ = Applicative0_5_168
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_2_2.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "throwError"), e_4), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_168.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_6})})
}))
})}
_ = monadThrowMaybeT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1402181699, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Error_Class_MonadError{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 23967309, UnsafePtr: unsafe.Pointer(monadThrowMaybeT1_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_3, a_4)
}))
})
})})}
}

func Call_Control_Monad_Maybe_Trans_monadSTMaybeT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadMaybeT1_2_1 -> *Constructor_Control_Monad_Monad
monadMaybeT1_2_1 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_21 -> gopurs_runtime.Value
__local_var_3_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_21
// TAST (Let): __local_var_3_20 -> gopurs_runtime.Value
__local_var_3_20 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_21, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_20
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorMaybeT1_4_2 -> *Constructor_Data_Functor_Functor
functorMaybeT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), v_6)
})
})))
_ = functorMaybeT1_4_2
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_6 -> *Constructor_Control_Bind_Bind
Bind1_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_6
// TAST (Let): Applicative0_7_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_7
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_9
// TAST (Let): functorMaybeT1_9_8 -> *Constructor_Data_Functor_Functor
functorMaybeT1_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_9, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_Maybe_Just
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t10 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)}
goto end_branch_10
} else {

}
}
{
__t10 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t10)}
}), v_11)
})
})))
_ = functorMaybeT1_9_8
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_12 -> *Constructor_Control_Bind_Bind
Bind1_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_12
// TAST (Let): Applicative0_12_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_13
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_12.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_13.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_14
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t14 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_11
// TAST (Let): Bind1_11_15 -> *Constructor_Control_Bind_Bind
Bind1_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_15
// TAST (Let): Applicative0_12_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_16
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_9_8)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_15.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_16.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_6.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t17 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_7.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_17
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t17 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
_ = __local_var_5_5
// TAST (Let): Bind1_6_18 -> *Constructor_Control_Bind_Bind
Bind1_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_18
// TAST (Let): Applicative0_7_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_19
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_4_2)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_18.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_19.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_20, x_4)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_22 -> *Constructor_Control_Bind_Bind
Bind1_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_22
// TAST (Let): Applicative0_4_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_23
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_25 -> gopurs_runtime.Value
__local_var_6_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_25
// TAST (Let): functorMaybeT1_6_24 -> *Constructor_Data_Functor_Functor
functorMaybeT1_6_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_25, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 *Constructor_Data_Maybe_Just
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t26 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)}
goto end_branch_26
} else {

}
}
{
__t26 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_26:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t26)}
}), v_8)
})
})))
_ = functorMaybeT1_6_24
// TAST (Let): __local_var_7_27 -> gopurs_runtime.Value
__local_var_7_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_38 -> gopurs_runtime.Value
__local_var_8_38 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_38
// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
__local_var_8_37 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_38, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_37
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_29 -> gopurs_runtime.Value
__local_var_9_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_29
// TAST (Let): functorMaybeT1_9_28 -> *Constructor_Data_Functor_Functor
functorMaybeT1_9_28 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_29, "map"), gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 *Constructor_Data_Maybe_Just
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t30 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)}
goto end_branch_30
} else {

}
}
{
__t30 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_30:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t30)}
}), v_11)
})
})))
_ = functorMaybeT1_9_28
// TAST (Let): __local_var_10_31 -> gopurs_runtime.Value
__local_var_10_31 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_32 -> *Constructor_Control_Bind_Bind
Bind1_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_32
// TAST (Let): Applicative0_12_33 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_33
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_32.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t34 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t34 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_33.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_34
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t34 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_10_31
// TAST (Let): Bind1_11_35 -> *Constructor_Control_Bind_Bind
Bind1_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_35
// TAST (Let): Applicative0_12_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_31, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_36
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_9_28)}
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_35.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_36.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_37, x_9)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_39 -> *Constructor_Control_Bind_Bind
Bind1_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_39
// TAST (Let): Applicative0_9_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_40
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)))}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_39.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t41 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_40.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_41
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t41 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_7_27
// TAST (Let): Bind1_8_42 -> *Constructor_Control_Bind_Bind
Bind1_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_42
// TAST (Let): Applicative0_9_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_43
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_6_24)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_42.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_43.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_22.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t44 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_23.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_44
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t44 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
})}
_ = monadMaybeT1_2_1
// TAST (Let): Bind1_3_47 -> *Constructor_Control_Bind_Bind
Bind1_3_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_47
// TAST (Let): Applicative0_4_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_48
// TAST (Let): __local_var_3_46 -> gopurs_runtime.Value
__local_var_3_46 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_47.V1), a_5, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_48.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_6})})
}))
})
_ = __local_var_3_46
// TAST (Let): __local_var_3_45 -> gopurs_runtime.Value
__local_var_3_45 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_46, x_4)
})
_ = __local_var_3_45
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_ST_Class_MonadST{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_2_1)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_45, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
})})}
}

func Call_Control_Monad_Maybe_Trans_monoidMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_20 -> gopurs_runtime.Value
__local_var_1_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_20
// TAST (Let): __local_var_1_19 -> gopurs_runtime.Value
__local_var_1_19 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_20, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_2})})
})
_ = __local_var_1_19
// TAST (Let): applicativeMaybeT1_1_0 -> *Constructor_Control_Applicative_Applicative
applicativeMaybeT1_1_0 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorMaybeT1_2_1 -> *Constructor_Data_Functor_Functor
functorMaybeT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_5.Type == 9 && v1_5.IntVal == 930809136 && v1_5.UnsafePtr != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Maybe_Just)(v1_5.UnsafePtr).V0)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}), v_4)
})
})))
_ = functorMaybeT1_2_1
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_5 -> *Constructor_Control_Bind_Bind
Bind1_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_5
// TAST (Let): Applicative0_5_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_6
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_8
// TAST (Let): functorMaybeT1_7_7 -> *Constructor_Data_Functor_Functor
functorMaybeT1_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Maybe_Just
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)}
goto end_branch_9
} else {

}
}
{
__t9 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
}), v_9)
})
})))
_ = functorMaybeT1_7_7
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_11 -> *Constructor_Control_Bind_Bind
Bind1_9_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_11
// TAST (Let): Applicative0_10_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_12
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_11.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_12.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_13
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t13 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
_ = __local_var_8_10
// TAST (Let): Bind1_9_14 -> *Constructor_Control_Bind_Bind
Bind1_9_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_14
// TAST (Let): Applicative0_10_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_15
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_7_7)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_14.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_14.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_15.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_5.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t16 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_6.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_16
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t16 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
_ = __local_var_3_4
// TAST (Let): Bind1_4_17 -> *Constructor_Control_Bind_Bind
Bind1_4_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_17
// TAST (Let): Applicative0_5_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_2_1)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_17.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_17.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_18.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_19, x_2)
})}
_ = applicativeMaybeT1_1_0
// TAST (Let): __local_var_2_23 -> gopurs_runtime.Value
__local_var_2_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_23
// TAST (Let): functorMaybeT1_2_22 -> *Constructor_Data_Functor_Functor
functorMaybeT1_2_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_23, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t24 *Constructor_Data_Maybe_Just
{
if (v1_5.Type == 9 && v1_5.IntVal == 930809136 && v1_5.UnsafePtr != nil) {
__t24 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Maybe_Just)(v1_5.UnsafePtr).V0)}
goto end_branch_24
} else {

}
}
{
__t24 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_24:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t24)}
}), v_4)
})
})))
_ = functorMaybeT1_2_22
// TAST (Let): __local_var_3_25 -> gopurs_runtime.Value
__local_var_3_25 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_76 -> gopurs_runtime.Value
__local_var_4_76 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_76
// TAST (Let): __local_var_4_75 -> gopurs_runtime.Value
__local_var_4_75 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_76, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_75
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_27 -> gopurs_runtime.Value
__local_var_5_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_27
// TAST (Let): functorMaybeT1_5_26 -> *Constructor_Data_Functor_Functor
functorMaybeT1_5_26 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_27, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 *Constructor_Data_Maybe_Just
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t28 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)}
goto end_branch_28
} else {

}
}
{
__t28 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_28:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t28)}
}), v_7)
})
})))
_ = functorMaybeT1_5_26
// TAST (Let): __local_var_6_29 -> gopurs_runtime.Value
__local_var_6_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_49 -> gopurs_runtime.Value
__local_var_7_49 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_49
// TAST (Let): __local_var_7_48 -> gopurs_runtime.Value
__local_var_7_48 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_49, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_8})})
})
_ = __local_var_7_48
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_31 -> gopurs_runtime.Value
__local_var_8_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_31
// TAST (Let): functorMaybeT1_8_30 -> *Constructor_Data_Functor_Functor
functorMaybeT1_8_30 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_31, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 *Constructor_Data_Maybe_Just
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t32 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)}
goto end_branch_32
} else {

}
}
{
__t32 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_32:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t32)}
}), v_10)
})
})))
_ = functorMaybeT1_8_30
// TAST (Let): __local_var_9_33 -> gopurs_runtime.Value
__local_var_9_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_34 -> *Constructor_Control_Bind_Bind
Bind1_10_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_34
// TAST (Let): Applicative0_11_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_35
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_37 -> gopurs_runtime.Value
__local_var_13_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_37
// TAST (Let): functorMaybeT1_13_36 -> *Constructor_Data_Functor_Functor
functorMaybeT1_13_36 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_37, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 *Constructor_Data_Maybe_Just
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t38 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)}
goto end_branch_38
} else {

}
}
{
__t38 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_38:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t38)}
}), v_15)
})
})))
_ = functorMaybeT1_13_36
// TAST (Let): __local_var_14_39 -> gopurs_runtime.Value
__local_var_14_39 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_40 -> *Constructor_Control_Bind_Bind
Bind1_15_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_40
// TAST (Let): Applicative0_16_41 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_41
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_40.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t42 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_41.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_42
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t42 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_14_39
// TAST (Let): Bind1_15_43 -> *Constructor_Control_Bind_Bind
Bind1_15_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_39, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_43
// TAST (Let): Applicative0_16_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_39, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_44
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_13_36)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_43.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_43.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_44.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_34.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t45 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t45 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_35.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_45
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t45 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
_ = __local_var_9_33
// TAST (Let): Bind1_10_46 -> *Constructor_Control_Bind_Bind
Bind1_10_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_46
// TAST (Let): Applicative0_11_47 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_47
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_8_30)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_46.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_46.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_47.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_48, x_8)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_50 -> *Constructor_Control_Bind_Bind
Bind1_7_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_50
// TAST (Let): Applicative0_8_51 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_51
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_53 -> gopurs_runtime.Value
__local_var_10_53 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_53
// TAST (Let): functorMaybeT1_10_52 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_52 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_53, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t54 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t54 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_54
} else {

}
}
{
__t54 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_54:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t54)}
}), v_12)
})
})))
_ = functorMaybeT1_10_52
// TAST (Let): __local_var_11_55 -> gopurs_runtime.Value
__local_var_11_55 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_66 -> gopurs_runtime.Value
__local_var_12_66 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_66
// TAST (Let): __local_var_12_65 -> gopurs_runtime.Value
__local_var_12_65 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_66, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_13})})
})
_ = __local_var_12_65
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_57 -> gopurs_runtime.Value
__local_var_13_57 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_57
// TAST (Let): functorMaybeT1_13_56 -> *Constructor_Data_Functor_Functor
functorMaybeT1_13_56 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_57, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t58 *Constructor_Data_Maybe_Just
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t58 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)}
goto end_branch_58
} else {

}
}
{
__t58 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_58:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t58)}
}), v_15)
})
})))
_ = functorMaybeT1_13_56
// TAST (Let): __local_var_14_59 -> gopurs_runtime.Value
__local_var_14_59 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_60 -> *Constructor_Control_Bind_Bind
Bind1_15_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_60
// TAST (Let): Applicative0_16_61 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_61
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_60.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t62 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t62 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_61.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_62
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t62 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_14_59
// TAST (Let): Bind1_15_63 -> *Constructor_Control_Bind_Bind
Bind1_15_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_59, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_63
// TAST (Let): Applicative0_16_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_59, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_64
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_13_56)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_63.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_63.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_64.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_65, x_13)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_67 -> *Constructor_Control_Bind_Bind
Bind1_12_67 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_67
// TAST (Let): Applicative0_13_68 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_68 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_68
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_67.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t69 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t69 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_68.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_69
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t69 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_55
// TAST (Let): Bind1_12_70 -> *Constructor_Control_Bind_Bind
Bind1_12_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_55, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_70
// TAST (Let): Applicative0_13_71 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_55, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_71
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_52)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_70.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_70.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_71.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_50.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t72 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t72 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_51.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_72
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t72 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_6_29
// TAST (Let): Bind1_7_73 -> *Constructor_Control_Bind_Bind
Bind1_7_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_73
// TAST (Let): Applicative0_8_74 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_74
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_5_26)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_73.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_73.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_74.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_75, x_5)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_77 -> *Constructor_Control_Bind_Bind
Bind1_4_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_77
// TAST (Let): Applicative0_5_78 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_78
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_80 -> gopurs_runtime.Value
__local_var_7_80 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_80
// TAST (Let): functorMaybeT1_7_79 -> *Constructor_Data_Functor_Functor
functorMaybeT1_7_79 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_80, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t81 *Constructor_Data_Maybe_Just
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t81 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)}
goto end_branch_81
} else {

}
}
{
__t81 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_81:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t81)}
}), v_9)
})
})))
_ = functorMaybeT1_7_79
// TAST (Let): __local_var_8_82 -> gopurs_runtime.Value
__local_var_8_82 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_93 -> gopurs_runtime.Value
__local_var_9_93 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_93
// TAST (Let): __local_var_9_92 -> gopurs_runtime.Value
__local_var_9_92 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_93, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_92
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_84 -> gopurs_runtime.Value
__local_var_10_84 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_84
// TAST (Let): functorMaybeT1_10_83 -> *Constructor_Data_Functor_Functor
functorMaybeT1_10_83 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_84, "map"), gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t85 *Constructor_Data_Maybe_Just
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t85 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)}
goto end_branch_85
} else {

}
}
{
__t85 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_85:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t85)}
}), v_12)
})
})))
_ = functorMaybeT1_10_83
// TAST (Let): __local_var_11_86 -> gopurs_runtime.Value
__local_var_11_86 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_87 -> *Constructor_Control_Bind_Bind
Bind1_12_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_87
// TAST (Let): Applicative0_13_88 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_88
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_87.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t89 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t89 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_88.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_89
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t89 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_11_86
// TAST (Let): Bind1_12_90 -> *Constructor_Control_Bind_Bind
Bind1_12_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_86, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_90
// TAST (Let): Applicative0_13_91 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_86, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_91
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_10_83)}
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_90.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_90.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_91.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_92, x_10)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_94 -> *Constructor_Control_Bind_Bind
Bind1_9_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_94
// TAST (Let): Applicative0_10_95 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_95 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_95
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_94.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t96 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t96 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_95.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_96
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t96 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_8_82
// TAST (Let): Bind1_9_97 -> *Constructor_Control_Bind_Bind
Bind1_9_97 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_82, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_97
// TAST (Let): Applicative0_10_98 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_82, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_98
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_7_79)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_97.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_97.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_98.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_77.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t99 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t99 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_78.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_99
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t99 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
_ = __local_var_3_25
// TAST (Let): Bind1_4_100 -> *Constructor_Control_Bind_Bind
Bind1_4_100 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_25, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_100
// TAST (Let): Applicative0_5_101 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_101 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_25, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_101
// TAST (Let): applyMaybeT1_2_21 -> *Constructor_Control_Apply_Apply
applyMaybeT1_2_21 := &Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_2_22)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_100.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_100.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_101.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})}
_ = applyMaybeT1_2_21
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_103 -> *Constructor_Data_Functor_Functor
Functor0_4_103 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyMaybeT1_2_21.V0), gopurs_runtime.Value{}))
_ = Functor0_4_103
// TAST (Let): __local_var_5_104 -> gopurs_runtime.Value
__local_var_5_104 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_5_104
// TAST (Let): semigroupMaybeT2_4_102 -> *Constructor_Data_Semigroup_Semigroup
semigroupMaybeT2_4_102 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyMaybeT1_2_21.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_103.V0), __local_var_5_104, a_6), b_7)
})
})}
_ = semigroupMaybeT2_4_102
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupMaybeT2_4_102)}
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeMaybeT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))})}
})
}

func Call_Control_Monad_Maybe_Trans_altMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorMaybeT1_3_2 -> *Constructor_Data_Functor_Functor
functorMaybeT1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 *Constructor_Data_Maybe_Just
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil) {
__t4 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Maybe_Just)(v1_6.UnsafePtr).V0)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t4)}
}), v_5)
})
})))
_ = functorMaybeT1_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_3_2)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_4, gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 930809136 && m_6.UnsafePtr == nil) {
__t5 = v1_5
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](m_6))})
}
end_branch_5:
return __t5
}))
})
})})}
}

func Call_Control_Monad_Maybe_Trans_plusMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorMaybeT1_3_3 -> *Constructor_Data_Functor_Functor
functorMaybeT1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 *Constructor_Data_Maybe_Just
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil) {
__t5 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Maybe_Just)(v1_6.UnsafePtr).V0)}
goto end_branch_5
} else {

}
}
{
__t5 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t5)}
}), v_5)
})
})))
_ = functorMaybeT1_3_3
// TAST (Let): altMaybeT1_1_0 -> *Constructor_Control_Alt_Alt
altMaybeT1_1_0 := &Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_3_3)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), v_4, gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 930809136 && m_6.UnsafePtr == nil) {
__t6 = v1_5
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](m_6))})
}
end_branch_6:
return __t6
}))
})
})}
_ = altMaybeT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(&Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altMaybeT1_1_0)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})})}
}

func Call_Control_Monad_Maybe_Trans_alternativeMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_20 -> gopurs_runtime.Value
__local_var_1_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_20
// TAST (Let): __local_var_1_19 -> gopurs_runtime.Value
__local_var_1_19 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_20, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_2})})
})
_ = __local_var_1_19
// TAST (Let): applicativeMaybeT1_1_0 -> *Constructor_Control_Applicative_Applicative
applicativeMaybeT1_1_0 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorMaybeT1_2_1 -> *Constructor_Data_Functor_Functor
functorMaybeT1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_5.Type == 9 && v1_5.IntVal == 930809136 && v1_5.UnsafePtr != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_3, (*Constructor_Data_Maybe_Just)(v1_5.UnsafePtr).V0)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}), v_4)
})
})))
_ = functorMaybeT1_2_1
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_5 -> *Constructor_Control_Bind_Bind
Bind1_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_5
// TAST (Let): Applicative0_5_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_6
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_8
// TAST (Let): functorMaybeT1_7_7 -> *Constructor_Data_Functor_Functor
functorMaybeT1_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_8, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Maybe_Just
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)}
goto end_branch_9
} else {

}
}
{
__t9 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
}), v_9)
})
})))
_ = functorMaybeT1_7_7
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_11 -> *Constructor_Control_Bind_Bind
Bind1_9_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_11
// TAST (Let): Applicative0_10_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_12
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_11.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_12.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_13
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t13 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
_ = __local_var_8_10
// TAST (Let): Bind1_9_14 -> *Constructor_Control_Bind_Bind
Bind1_9_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_14
// TAST (Let): Applicative0_10_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_15
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_7_7)}
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_14.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_14.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_15.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_5.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t16 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_6.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_16
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t16 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
_ = __local_var_3_4
// TAST (Let): Bind1_4_17 -> *Constructor_Control_Bind_Bind
Bind1_4_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_17
// TAST (Let): Applicative0_5_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_2_1)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_17.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_17.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_18.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_19, x_2)
})}
_ = applicativeMaybeT1_1_0
// TAST (Let): Bind1_2_23 -> *Constructor_Control_Bind_Bind
Bind1_2_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_23
// TAST (Let): Applicative0_3_24 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_24
// TAST (Let): __local_var_4_26 -> gopurs_runtime.Value
__local_var_4_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_26
// TAST (Let): functorMaybeT1_4_25 -> *Constructor_Data_Functor_Functor
functorMaybeT1_4_25 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_26, "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 *Constructor_Data_Maybe_Just
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t27 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_5, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)}
goto end_branch_27
} else {

}
}
{
__t27 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_27:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t27)}
}), v_6)
})
})))
_ = functorMaybeT1_4_25
// TAST (Let): altMaybeT1_2_22 -> *Constructor_Control_Alt_Alt
altMaybeT1_2_22 := &Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_4_25)}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_23.V1), v_5, gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 930809136 && m_7.UnsafePtr == nil) {
__t28 = v1_6
goto end_branch_28
} else {

}
}
{
__t28 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_24.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](m_7))})
}
end_branch_28:
return __t28
}))
})
})}
_ = altMaybeT1_2_22
// TAST (Let): plusMaybeT1_2_21 -> *Constructor_Control_Plus_Plus
plusMaybeT1_2_21 := &Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altMaybeT1_2_22)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})}
_ = plusMaybeT1_2_21
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(&Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeMaybeT1_1_0)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusMaybeT1_2_21)}
})})}
}

func Call_Control_Monad_Maybe_Trans_monadPlusMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadMaybeT1_1_0 -> *Constructor_Control_Monad_Monad
monadMaybeT1_1_0 := &Constructor_Control_Monad_Monad{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_20 -> gopurs_runtime.Value
__local_var_2_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_2_20
// TAST (Let): __local_var_2_19 -> gopurs_runtime.Value
__local_var_2_19 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_20, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_3})})
})
_ = __local_var_2_19
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorMaybeT1_3_1 -> *Constructor_Data_Functor_Functor
functorMaybeT1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil) {
__t3 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Maybe_Just)(v1_6.UnsafePtr).V0)}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
}), v_5)
})
})))
_ = functorMaybeT1_3_1
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_5 -> *Constructor_Control_Bind_Bind
Bind1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_5
// TAST (Let): Applicative0_6_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_6
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_8
// TAST (Let): functorMaybeT1_8_7 -> *Constructor_Data_Functor_Functor
functorMaybeT1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_8, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Maybe_Just
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)}
goto end_branch_9
} else {

}
}
{
__t9 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_9:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)}
}), v_10)
})
})))
_ = functorMaybeT1_8_7
// TAST (Let): __local_var_9_10 -> gopurs_runtime.Value
__local_var_9_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_11 -> *Constructor_Control_Bind_Bind
Bind1_10_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_11
// TAST (Let): Applicative0_11_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_12
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_11.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_12.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_13
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t13 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
_ = __local_var_9_10
// TAST (Let): Bind1_10_14 -> *Constructor_Control_Bind_Bind
Bind1_10_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_14
// TAST (Let): Applicative0_11_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_15
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_8_7)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_14.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_15.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_5.V1), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr == nil) {
__t16 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_6.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_16
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t16 = gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)
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
_ = __local_var_4_4
// TAST (Let): Bind1_5_17 -> *Constructor_Control_Bind_Bind
Bind1_5_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_17
// TAST (Let): Applicative0_6_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_18
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_3_1)}
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_17.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_17.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_18.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_19, x_3)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_21 -> *Constructor_Control_Bind_Bind
Bind1_2_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_21
// TAST (Let): Applicative0_3_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_22
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_24 -> gopurs_runtime.Value
__local_var_5_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_24
// TAST (Let): functorMaybeT1_5_23 -> *Constructor_Data_Functor_Functor
functorMaybeT1_5_23 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_24, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 *Constructor_Data_Maybe_Just
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t25 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)}
goto end_branch_25
} else {

}
}
{
__t25 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_25:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t25)}
}), v_7)
})
})))
_ = functorMaybeT1_5_23
// TAST (Let): __local_var_6_26 -> gopurs_runtime.Value
__local_var_6_26 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_37 -> gopurs_runtime.Value
__local_var_7_37 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_37
// TAST (Let): __local_var_7_36 -> gopurs_runtime.Value
__local_var_7_36 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_37, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_8})})
})
_ = __local_var_7_36
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_28 -> gopurs_runtime.Value
__local_var_8_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_28
// TAST (Let): functorMaybeT1_8_27 -> *Constructor_Data_Functor_Functor
functorMaybeT1_8_27 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_28, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 *Constructor_Data_Maybe_Just
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t29 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)}
goto end_branch_29
} else {

}
}
{
__t29 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_29:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t29)}
}), v_10)
})
})))
_ = functorMaybeT1_8_27
// TAST (Let): __local_var_9_30 -> gopurs_runtime.Value
__local_var_9_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_31 -> *Constructor_Control_Bind_Bind
Bind1_10_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_31
// TAST (Let): Applicative0_11_32 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_32
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_31.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t33 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t33 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_32.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_33
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t33 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_9_30
// TAST (Let): Bind1_10_34 -> *Constructor_Control_Bind_Bind
Bind1_10_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_34
// TAST (Let): Applicative0_11_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_35
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_8_27)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_34.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_34.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_35.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_36, x_8)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_38 -> *Constructor_Control_Bind_Bind
Bind1_7_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_38
// TAST (Let): Applicative0_8_39 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_39
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_38.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t40 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_39.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_40
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t40 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
_ = __local_var_6_26
// TAST (Let): Bind1_7_41 -> *Constructor_Control_Bind_Bind
Bind1_7_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_26, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_41
// TAST (Let): Applicative0_8_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_26, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_42
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_5_23)}
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_41.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_41.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_42.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_21.V1), v_4, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr == nil) {
__t43 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_22.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_43
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil) {
__t43 = gopurs_runtime.Apply(f_5, (*Constructor_Data_Maybe_Just)(v1_6.UnsafePtr).V0)
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
})})}
})}
_ = monadMaybeT1_1_0
// TAST (Let): __local_var_2_125 -> gopurs_runtime.Value
__local_var_2_125 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_2_125
// TAST (Let): __local_var_2_124 -> gopurs_runtime.Value
__local_var_2_124 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_125, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_3})})
})
_ = __local_var_2_124
// TAST (Let): applicativeMaybeT1_2_45 -> *Constructor_Control_Applicative_Applicative
applicativeMaybeT1_2_45 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_47 -> gopurs_runtime.Value
__local_var_3_47 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_47
// TAST (Let): functorMaybeT1_3_46 -> *Constructor_Data_Functor_Functor
functorMaybeT1_3_46 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_47, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t48 *Constructor_Data_Maybe_Just
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil) {
__t48 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_4, (*Constructor_Data_Maybe_Just)(v1_6.UnsafePtr).V0)}
goto end_branch_48
} else {

}
}
{
__t48 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_48:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t48)}
}), v_5)
})
})))
_ = functorMaybeT1_3_46
// TAST (Let): __local_var_4_49 -> gopurs_runtime.Value
__local_var_4_49 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_69 -> gopurs_runtime.Value
__local_var_5_69 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_69
// TAST (Let): __local_var_5_68 -> gopurs_runtime.Value
__local_var_5_68 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_69, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_6})})
})
_ = __local_var_5_68
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_51 -> gopurs_runtime.Value
__local_var_6_51 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_51
// TAST (Let): functorMaybeT1_6_50 -> *Constructor_Data_Functor_Functor
functorMaybeT1_6_50 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_51, "map"), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t52 *Constructor_Data_Maybe_Just
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t52 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)}
goto end_branch_52
} else {

}
}
{
__t52 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_52:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t52)}
}), v_8)
})
})))
_ = functorMaybeT1_6_50
// TAST (Let): __local_var_7_53 -> gopurs_runtime.Value
__local_var_7_53 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_54 -> *Constructor_Control_Bind_Bind
Bind1_8_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_54
// TAST (Let): Applicative0_9_55 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_55
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_57 -> gopurs_runtime.Value
__local_var_11_57 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_57
// TAST (Let): functorMaybeT1_11_56 -> *Constructor_Data_Functor_Functor
functorMaybeT1_11_56 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_57, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t58 *Constructor_Data_Maybe_Just
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t58 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)}
goto end_branch_58
} else {

}
}
{
__t58 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_58:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t58)}
}), v_13)
})
})))
_ = functorMaybeT1_11_56
// TAST (Let): __local_var_12_59 -> gopurs_runtime.Value
__local_var_12_59 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_60 -> *Constructor_Control_Bind_Bind
Bind1_13_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_60
// TAST (Let): Applicative0_14_61 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_61
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_60.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t62 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t62 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_61.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_62
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t62 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_12_59
// TAST (Let): Bind1_13_63 -> *Constructor_Control_Bind_Bind
Bind1_13_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_59, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_63
// TAST (Let): Applicative0_14_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_59, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_64
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_11_56)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_63.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_63.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_64.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_54.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t65 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t65 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_55.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_65
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t65 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
})
})})}
}))
_ = __local_var_7_53
// TAST (Let): Bind1_8_66 -> *Constructor_Control_Bind_Bind
Bind1_8_66 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_53, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_66
// TAST (Let): Applicative0_9_67 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_67 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_53, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_67
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_6_50)}
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_66.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_66.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_67.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_68, x_6)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_70 -> *Constructor_Control_Bind_Bind
Bind1_5_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_70
// TAST (Let): Applicative0_6_71 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_71
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_73 -> gopurs_runtime.Value
__local_var_8_73 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_73
// TAST (Let): functorMaybeT1_8_72 -> *Constructor_Data_Functor_Functor
functorMaybeT1_8_72 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_73, "map"), gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t74 *Constructor_Data_Maybe_Just
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t74 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)}
goto end_branch_74
} else {

}
}
{
__t74 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_74:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t74)}
}), v_10)
})
})))
_ = functorMaybeT1_8_72
// TAST (Let): __local_var_9_75 -> gopurs_runtime.Value
__local_var_9_75 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_95 -> gopurs_runtime.Value
__local_var_10_95 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_95
// TAST (Let): __local_var_10_94 -> gopurs_runtime.Value
__local_var_10_94 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_95, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_11})})
})
_ = __local_var_10_94
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_77 -> gopurs_runtime.Value
__local_var_11_77 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_77
// TAST (Let): functorMaybeT1_11_76 -> *Constructor_Data_Functor_Functor
functorMaybeT1_11_76 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_77, "map"), gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t78 *Constructor_Data_Maybe_Just
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t78 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)}
goto end_branch_78
} else {

}
}
{
__t78 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_78:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t78)}
}), v_13)
})
})))
_ = functorMaybeT1_11_76
// TAST (Let): __local_var_12_79 -> gopurs_runtime.Value
__local_var_12_79 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_80 -> *Constructor_Control_Bind_Bind
Bind1_13_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_80
// TAST (Let): Applicative0_14_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_81
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_83 -> gopurs_runtime.Value
__local_var_16_83 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_83
// TAST (Let): functorMaybeT1_16_82 -> *Constructor_Data_Functor_Functor
functorMaybeT1_16_82 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_83, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t84 *Constructor_Data_Maybe_Just
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t84 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)}
goto end_branch_84
} else {

}
}
{
__t84 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_84:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t84)}
}), v_18)
})
})))
_ = functorMaybeT1_16_82
// TAST (Let): __local_var_17_85 -> gopurs_runtime.Value
__local_var_17_85 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_86 -> *Constructor_Control_Bind_Bind
Bind1_18_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_86
// TAST (Let): Applicative0_19_87 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_87
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_86.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t88 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t88 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_87.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_88
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t88 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
goto end_branch_88
} else {

}
}
{
__t88 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_88:
return __t88
}))
})
})})}
}))
_ = __local_var_17_85
// TAST (Let): Bind1_18_89 -> *Constructor_Control_Bind_Bind
Bind1_18_89 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_85, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_89
// TAST (Let): Applicative0_19_90 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_85, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_90
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_16_82)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_89.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_89.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_90.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_80.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t91 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t91 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_81.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_91
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t91 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
goto end_branch_91
} else {

}
}
{
__t91 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_91:
return __t91
}))
})
})})}
}))
_ = __local_var_12_79
// TAST (Let): Bind1_13_92 -> *Constructor_Control_Bind_Bind
Bind1_13_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_79, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_92
// TAST (Let): Applicative0_14_93 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_93 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_79, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_93
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_11_76)}
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_92.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_92.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_93.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_94, x_11)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_96 -> *Constructor_Control_Bind_Bind
Bind1_10_96 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_96
// TAST (Let): Applicative0_11_97 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_97 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_97
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_99 -> gopurs_runtime.Value
__local_var_13_99 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_99
// TAST (Let): functorMaybeT1_13_98 -> *Constructor_Data_Functor_Functor
functorMaybeT1_13_98 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_99, "map"), gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t100 *Constructor_Data_Maybe_Just
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t100 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)}
goto end_branch_100
} else {

}
}
{
__t100 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_100:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t100)}
}), v_15)
})
})))
_ = functorMaybeT1_13_98
// TAST (Let): __local_var_14_101 -> gopurs_runtime.Value
__local_var_14_101 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_112 -> gopurs_runtime.Value
__local_var_15_112 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_15_112
// TAST (Let): __local_var_15_111 -> gopurs_runtime.Value
__local_var_15_111 := gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_112, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_16})})
})
_ = __local_var_15_111
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(&Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_103 -> gopurs_runtime.Value
__local_var_16_103 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_103
// TAST (Let): functorMaybeT1_16_102 -> *Constructor_Data_Functor_Functor
functorMaybeT1_16_102 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_103, "map"), gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t104 *Constructor_Data_Maybe_Just
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t104 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_17, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)}
goto end_branch_104
} else {

}
}
{
__t104 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_104:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t104)}
}), v_18)
})
})))
_ = functorMaybeT1_16_102
// TAST (Let): __local_var_17_105 -> gopurs_runtime.Value
__local_var_17_105 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_106 -> *Constructor_Control_Bind_Bind
Bind1_18_106 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_106
// TAST (Let): Applicative0_19_107 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_107
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_106.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t108 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t108 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_107.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_108
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t108 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
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
})
})})}
}))
_ = __local_var_17_105
// TAST (Let): Bind1_18_109 -> *Constructor_Control_Bind_Bind
Bind1_18_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_105, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_109
// TAST (Let): Applicative0_19_110 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_105, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_110
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_16_102)}
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_109.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_109.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_110.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_111, x_16)
})})}
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_113 -> *Constructor_Control_Bind_Bind
Bind1_15_113 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_113
// TAST (Let): Applicative0_16_114 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_114 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_114
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(&Constructor_Control_Bind_Bind{1, gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)))}
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_113.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t115 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t115 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_114.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_115
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t115 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
})})}
}))
_ = __local_var_14_101
// TAST (Let): Bind1_15_116 -> *Constructor_Control_Bind_Bind
Bind1_15_116 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_101, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_116
// TAST (Let): Applicative0_16_117 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_117 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_101, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_117
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_13_98)}
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_116.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_116.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_117.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_96.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t118 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t118 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_97.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_118
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t118 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
_ = __local_var_9_75
// TAST (Let): Bind1_10_119 -> *Constructor_Control_Bind_Bind
Bind1_10_119 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_75, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_119
// TAST (Let): Applicative0_11_120 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_120 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_75, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_120
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_8_72)}
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_119.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_119.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_120.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_70.V1), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t121 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr == nil) {
__t121 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_71.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_121
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t121 = gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)
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
_ = __local_var_4_49
// TAST (Let): Bind1_5_122 -> *Constructor_Control_Bind_Bind
Bind1_5_122 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_49, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_122
// TAST (Let): Applicative0_6_123 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_123 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_49, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_123
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(&Constructor_Control_Apply_Apply{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_3_46)}
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_122.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_122.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_123.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
})})}
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_124, x_3)
})}
_ = applicativeMaybeT1_2_45
// TAST (Let): Bind1_3_128 -> *Constructor_Control_Bind_Bind
Bind1_3_128 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_128
// TAST (Let): Applicative0_4_129 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_129 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_129
// TAST (Let): __local_var_5_131 -> gopurs_runtime.Value
__local_var_5_131 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_131
// TAST (Let): functorMaybeT1_5_130 -> *Constructor_Data_Functor_Functor
functorMaybeT1_5_130 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_131, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t132 *Constructor_Data_Maybe_Just
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t132 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)}
goto end_branch_132
} else {

}
}
{
__t132 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_132:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t132)}
}), v_7)
})
})))
_ = functorMaybeT1_5_130
// TAST (Let): altMaybeT1_3_127 -> *Constructor_Control_Alt_Alt
altMaybeT1_3_127 := &Constructor_Control_Alt_Alt{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorMaybeT1_5_130)}
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_128.V1), v_6, gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t133 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 930809136 && m_8.UnsafePtr == nil) {
__t133 = v1_7
goto end_branch_133
} else {

}
}
{
__t133 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_129.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](m_8))})
}
end_branch_133:
return __t133
}))
})
})}
_ = altMaybeT1_3_127
// TAST (Let): plusMaybeT1_3_126 -> *Constructor_Control_Plus_Plus
plusMaybeT1_3_126 := &Constructor_Control_Plus_Plus{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(altMaybeT1_3_127)}
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})}
_ = plusMaybeT1_3_126
// TAST (Let): alternativeMaybeT1_2_44 -> *Constructor_Control_Alternative_Alternative
alternativeMaybeT1_2_44 := &Constructor_Control_Alternative_Alternative{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(applicativeMaybeT1_2_45)}
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3709470893, UnsafePtr: unsafe.Pointer(plusMaybeT1_3_126)}
})}
_ = alternativeMaybeT1_2_44
return gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(&Constructor_Control_MonadPlus_MonadPlus{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(alternativeMaybeT1_2_44)}
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(monadMaybeT1_1_0)}
})})}
}

func Call_Control_Monad_Maybe_Trans_mapMaybeT__4176318367(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Maybe_Trans_mapMaybeT__1721503071(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Control_Monad_Maybe_Trans_mapMaybeT__1878923231(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}



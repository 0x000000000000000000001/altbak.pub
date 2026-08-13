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
		cache_Control_Monad_Maybe_Trans_newtypeMaybeT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Control_Monad_Maybe_Trans_newtypeMaybeT
}

var cache_Control_Monad_Maybe_Trans_monadTransMaybeT gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_monadTransMaybeT sync.Once
func Get_Control_Monad_Maybe_Trans_monadTransMaybeT() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_monadTransMaybeT.Do(func() {
		cache_Control_Monad_Maybe_Trans_monadTransMaybeT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), a_3, gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, a_prime_4})})
}))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}))
	})
	return cache_Control_Monad_Maybe_Trans_monadTransMaybeT
}

var cache_Control_Monad_Maybe_Trans_lift gopurs_runtime.Value
var once_Control_Monad_Maybe_Trans_lift sync.Once
func Get_Control_Monad_Maybe_Trans_lift() gopurs_runtime.Value {
	once_Control_Monad_Maybe_Trans_lift.Do(func() {
		cache_Control_Monad_Maybe_Trans_lift = gopurs_runtime.RecordGet(Get_Control_Monad_Maybe_Trans_monadTransMaybeT(), "lift")
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
		cache_Control_Monad_Maybe_Trans_monadTransMaybeT__3775352453 = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), a_3, gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, a_prime_4})})
}))
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
})
}))
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
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_1), v_2)
})
}))
}

func Call_Control_Monad_Maybe_Trans_monadMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(dictMonad_0)
}))
}

func Call_Control_Monad_Maybe_Trans_bindMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_3, gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 930809136 && v1_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
goto end_branch_2
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 930809136 && v1_5.UnsafePtr != nil) {
__t2 = gopurs_runtime.Apply(f_4, (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(v1_5.UnsafePtr).V0)
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

func Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorMaybeT1_1_0 -> gopurs_runtime.Value
functorMaybeT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_2), v_3)
})
}))
_ = functorMaybeT1_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(dictMonad_0)
}))
_ = __local_var_2_2
// TAST (Let): Bind1_3_3 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_3
// TAST (Let): Applicative0_4_4 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_4
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_1_0
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

func Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_1
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, x_2})})
})
_ = __local_var_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, x_2)
}))
}

func Call_Control_Monad_Maybe_Trans_semigroupMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applyMaybeT1_1_0 -> *Constructor_Control_Apply_Apply[gopurs_runtime.Value]
applyMaybeT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0))
_ = applyMaybeT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(applyMaybeT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_2
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyMaybeT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_1.V0), __local_var_4_2, a_5), b_6)
})
}))
})
}

func Call_Control_Monad_Maybe_Trans_monadAskMaybeT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadMaybeT1_1_0 -> gopurs_runtime.Value
monadMaybeT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(__local_var_1_1)
}))
_ = monadMaybeT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_1_0
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_Maybe_Trans_monadTransMaybeT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})))}, gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")))
}

func Call_Control_Monad_Maybe_Trans_monadReaderMaybeT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
// TAST (Let): monadAskMaybeT1_1_0 -> gopurs_runtime.Value
monadAskMaybeT1_1_0 := Call_Control_Monad_Maybe_Trans_monadAskMaybeT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskMaybeT1_1_0
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskMaybeT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, v_4)
})
}))
}

func Call_Control_Monad_Maybe_Trans_monadContMaybeT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): monadMaybeT1_1_0 -> gopurs_runtime.Value
monadMaybeT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(__local_var_1_1)
}))
_ = monadMaybeT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, a_4})})
}))
}))
}))
}

func Call_Control_Monad_Maybe_Trans_monadEffectMaybe(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadMaybeT1_2_1 -> gopurs_runtime.Value
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(Monad0_1_0)
}))
_ = monadMaybeT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Maybe_Trans_monadTransMaybeT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_Control_Monad_Maybe_Trans_monadRecMaybeT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): monadMaybeT1_4_3 -> gopurs_runtime.Value
monadMaybeT1_4_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(Monad0_1_0)
}))
_ = monadMaybeT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_4_3
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_4 -> gopurs_runtime.Value
__local_var_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(f_5, a_6), gopurs_runtime.Func(func(m_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 930809136 && m_prime_7.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}})}
goto end_branch_8
} else {

}
}
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 930809136 && m_prime_7.UnsafePtr != nil) {
var __t7 gopurs_runtime.Value
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(m_prime_7.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 525585346) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}
goto end_branch_7
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(m_prime_7.UnsafePtr).V0
if (__t_tag_6.Type == 9 && __t_tag_6.IntVal == 60402430) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}})}
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

func Call_Control_Monad_Maybe_Trans_monadStateMaybeT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
// TAST (Let): Monad0_1_0 -> *Constructor_Control_Monad_Monad[gopurs_runtime.Value]
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): monadMaybeT1_2_1 -> gopurs_runtime.Value
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(__local_var_2_2)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_Maybe_Trans_monadTransMaybeT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_1_0)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), f_3))
}))
}

func Call_Control_Monad_Maybe_Trans_monadTellMaybeT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
// TAST (Let): Monad1_1_0 -> gopurs_runtime.Value
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Semigroup0_2_1 -> gopurs_runtime.Value
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
// TAST (Let): monadMaybeT1_3_2 -> gopurs_runtime.Value
monadMaybeT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(Monad1_1_0)
}))
_ = monadMaybeT1_3_2
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Maybe_Trans_monadTransMaybeT(), "lift"), Monad1_1_0)
_ = __local_var_4_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
}))
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
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): pure_4_3 -> gopurs_runtime.Value
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
// TAST (Let): Applicative0_5_4 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_4
// TAST (Let): Monoid0_6_5 -> gopurs_runtime.Value
Monoid0_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_6_5
// TAST (Let): monadTellMaybeT1_7_6 -> gopurs_runtime.Value
monadTellMaybeT1_7_6 := Call_Control_Monad_Maybe_Trans_monadTellMaybeT(MonadTell1_1_0)
_ = monadTellMaybeT1_7_6
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellMaybeT1_7_6
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_6_5
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), v_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_7 -> gopurs_runtime.Value
__local_var_10_7 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1
_ = __local_var_10_7
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, r_11, __local_var_10_7})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]]((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0))})))})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_8, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (a_9.Type == 9 && a_9.IntVal == 930809136 && a_9.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_10
})})}
goto end_branch_8
} else {

}
}
{
if (a_9.Type == 9 && a_9.IntVal == 930809136 && a_9.UnsafePtr != nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(a_9.UnsafePtr).V0.UnsafePtr).V0})}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(a_9.UnsafePtr).V0.UnsafePtr).V1})}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[*Constructor_Data_Maybe_Just[gopurs_runtime.Value], gopurs_runtime.Value]](__t8))})
})))
}))
}

func Call_Control_Monad_Maybe_Trans_monadThrowMaybeT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
// TAST (Let): Monad0_1_0 -> *Constructor_Control_Monad_Monad[gopurs_runtime.Value]
Monad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{}))
_ = Monad0_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): monadMaybeT1_2_1 -> gopurs_runtime.Value
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(__local_var_2_2)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_Maybe_Trans_monadTransMaybeT(), "lift"), gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(Monad0_1_0)}, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_3))
}))
}

func Call_Control_Monad_Maybe_Trans_monadErrorMaybeT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
// TAST (Let): monadThrowMaybeT1_1_0 -> gopurs_runtime.Value
monadThrowMaybeT1_1_0 := Call_Control_Monad_Maybe_Trans_monadThrowMaybeT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowMaybeT1_1_0
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowMaybeT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_3, a_4)
}))
})
}))
}

func Call_Control_Monad_Maybe_Trans_monadSTMaybeT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): monadMaybeT1_2_1 -> gopurs_runtime.Value
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(Monad0_1_0)
}))
_ = monadMaybeT1_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_Maybe_Trans_monadTransMaybeT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_Control_Monad_Maybe_Trans_monoidMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeMaybeT1_1_0 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
applicativeMaybeT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0))
_ = applicativeMaybeT1_1_0
// TAST (Let): semigroupMaybeT1_2_1 -> gopurs_runtime.Value
semigroupMaybeT1_2_1 := Call_Control_Monad_Maybe_Trans_semigroupMaybeT(dictMonad_0)
_ = semigroupMaybeT1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): semigroupMaybeT2_4_2 -> gopurs_runtime.Value
semigroupMaybeT2_4_2 := gopurs_runtime.Apply(semigroupMaybeT1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupMaybeT2_4_2
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMaybeT2_4_2
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeMaybeT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty")))
})
}

func Call_Control_Monad_Maybe_Trans_altMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): functorMaybeT1_3_2 -> gopurs_runtime.Value
functorMaybeT1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_4), v_5)
})
}))
_ = functorMaybeT1_3_2
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), v_4, gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 930809136 && m_6.UnsafePtr == nil) {
__t4 = v1_5
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](m_6))})
}
end_branch_4:
return __t4
}))
})
}))
}

func Call_Control_Monad_Maybe_Trans_plusMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): altMaybeT1_1_0 -> gopurs_runtime.Value
altMaybeT1_1_0 := Call_Control_Monad_Maybe_Trans_altMaybeT(dictMonad_0)
_ = altMaybeT1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMaybeT1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
}

func Call_Control_Monad_Maybe_Trans_alternativeMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): applicativeMaybeT1_1_0 -> gopurs_runtime.Value
applicativeMaybeT1_1_0 := Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
_ = applicativeMaybeT1_1_0
// TAST (Let): plusMaybeT1_2_1 -> gopurs_runtime.Value
plusMaybeT1_2_1 := Call_Control_Monad_Maybe_Trans_plusMaybeT(dictMonad_0)
_ = plusMaybeT1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeMaybeT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusMaybeT1_2_1
}))
}

func Call_Control_Monad_Maybe_Trans_monadPlusMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadMaybeT1_1_0 -> gopurs_runtime.Value
monadMaybeT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_bindMaybeT(dictMonad_0)
}))
_ = monadMaybeT1_1_0
// TAST (Let): alternativeMaybeT1_2_1 -> gopurs_runtime.Value
alternativeMaybeT1_2_1 := Call_Control_Monad_Maybe_Trans_alternativeMaybeT(dictMonad_0)
_ = alternativeMaybeT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeMaybeT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_1_0
}))
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



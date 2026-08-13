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
}))
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
		cache_Control_Monad_Maybe_Trans_monadTransMaybeT__3775352453 = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
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
// TAST (Let): Bind1_3_3 -> *Constructor_Control_Bind_Bind
Bind1_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_3
// TAST (Let): Applicative0_4_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_4
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_3.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_5
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
}))
_ = __local_var_2_2
// TAST (Let): Bind1_3_6 -> *Constructor_Control_Bind_Bind
Bind1_3_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_6
// TAST (Let): Applicative0_4_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_7
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_1_0
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_6.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_6.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_7.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
}))
}

func Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_9 -> gopurs_runtime.Value
__local_var_1_9 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_9
// TAST (Let): __local_var_1_8 -> gopurs_runtime.Value
__local_var_1_8 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_9, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_2})})
})
_ = __local_var_1_8
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): functorMaybeT1_2_0 -> gopurs_runtime.Value
functorMaybeT1_2_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_3), v_4)
})
}))
_ = functorMaybeT1_2_0
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_3 -> *Constructor_Control_Bind_Bind
Bind1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_3
// TAST (Let): Applicative0_5_4 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_4
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_3.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_5
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
}))
_ = __local_var_3_2
// TAST (Let): Bind1_4_6 -> *Constructor_Control_Bind_Bind
Bind1_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_6
// TAST (Let): Applicative0_5_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_7
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_2_0
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_6.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_6.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_7.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_8, x_2)
}))
}

func Call_Control_Monad_Maybe_Trans_semigroupMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): functorMaybeT1_1_1 -> gopurs_runtime.Value
functorMaybeT1_1_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_2), v_3)
})
}))
_ = functorMaybeT1_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_13 -> gopurs_runtime.Value
__local_var_3_13 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_13
// TAST (Let): __local_var_3_12 -> gopurs_runtime.Value
__local_var_3_12 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_13, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_12
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_5 -> gopurs_runtime.Value
__local_var_4_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_5
// TAST (Let): functorMaybeT1_4_4 -> gopurs_runtime.Value
functorMaybeT1_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_5), v_6)
})
}))
_ = functorMaybeT1_4_4
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_7 -> *Constructor_Control_Bind_Bind
Bind1_6_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_7
// TAST (Let): Applicative0_7_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_8
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_7.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t9 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_8.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_9
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t9 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_5_6
// TAST (Let): Bind1_6_10 -> *Constructor_Control_Bind_Bind
Bind1_6_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_10
// TAST (Let): Applicative0_7_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_11
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_4_4
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_10.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_10.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_11.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_12, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_14 -> *Constructor_Control_Bind_Bind
Bind1_3_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_14
// TAST (Let): Applicative0_4_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_15
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_14.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t16 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_15.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_16
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t16 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
_ = __local_var_2_3
// TAST (Let): Bind1_3_17 -> *Constructor_Control_Bind_Bind
Bind1_3_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_17
// TAST (Let): Applicative0_4_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_18
// TAST (Let): applyMaybeT1_1_0 -> *Constructor_Control_Apply_Apply
applyMaybeT1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_1_1
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_17.V1), f_5, gopurs_runtime.Func(func(f_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_17.V1), a_6, gopurs_runtime.Func(func(a_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_18.V1), gopurs_runtime.Apply(f_prime_7, a_prime_8))
}))
}))
})
})))
_ = applyMaybeT1_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_3_19 -> *Constructor_Data_Functor_Functor
Functor0_3_19 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyMaybeT1_1_0.V0), gopurs_runtime.Value{}))
_ = Functor0_3_19
// TAST (Let): __local_var_4_20 -> gopurs_runtime.Value
__local_var_4_20 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_4_20
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyMaybeT1_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_19.V0), __local_var_4_20, a_5), b_6)
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
// TAST (Let): __local_var_3_19 -> gopurs_runtime.Value
__local_var_3_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_19
// TAST (Let): __local_var_3_18 -> gopurs_runtime.Value
__local_var_3_18 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_19, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_18
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorMaybeT1_4_2 -> gopurs_runtime.Value
functorMaybeT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_5), v_6)
})
}))
_ = functorMaybeT1_4_2
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_5 -> *Constructor_Control_Bind_Bind
Bind1_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_5
// TAST (Let): Applicative0_7_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorMaybeT1_9_7 -> gopurs_runtime.Value
functorMaybeT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_10), v_11)
})
}))
_ = functorMaybeT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
// TAST (Let): Applicative0_12_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_11
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t12 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_11.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_12
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t12 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_13 -> *Constructor_Control_Bind_Bind
Bind1_11_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_13
// TAST (Let): Applicative0_12_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_13.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_13.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_14.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_5.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_6.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_15
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t15 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
_ = __local_var_5_4
// TAST (Let): Bind1_6_16 -> *Constructor_Control_Bind_Bind
Bind1_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_16
// TAST (Let): Applicative0_7_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_4_2
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_16.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_16.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_17.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_18, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_20 -> *Constructor_Control_Bind_Bind
Bind1_3_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_20
// TAST (Let): Applicative0_4_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_21
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_23 -> gopurs_runtime.Value
__local_var_6_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_23
// TAST (Let): functorMaybeT1_6_22 -> gopurs_runtime.Value
functorMaybeT1_6_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_7), v_8)
})
}))
_ = functorMaybeT1_6_22
// TAST (Let): __local_var_7_24 -> gopurs_runtime.Value
__local_var_7_24 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_34 -> gopurs_runtime.Value
__local_var_8_34 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_34
// TAST (Let): __local_var_8_33 -> gopurs_runtime.Value
__local_var_8_33 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_34, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_33
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_26 -> gopurs_runtime.Value
__local_var_9_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_26
// TAST (Let): functorMaybeT1_9_25 -> gopurs_runtime.Value
functorMaybeT1_9_25 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_26, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_10), v_11)
})
}))
_ = functorMaybeT1_9_25
// TAST (Let): __local_var_10_27 -> gopurs_runtime.Value
__local_var_10_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_28 -> *Constructor_Control_Bind_Bind
Bind1_11_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_28
// TAST (Let): Applicative0_12_29 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_29
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_28.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t30 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_29.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_30
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t30 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_10_27
// TAST (Let): Bind1_11_31 -> *Constructor_Control_Bind_Bind
Bind1_11_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_31
// TAST (Let): Applicative0_12_32 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_32
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_9_25
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_31.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_31.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_32.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_33, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_35 -> *Constructor_Control_Bind_Bind
Bind1_8_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_35
// TAST (Let): Applicative0_9_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_36
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_35.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t37 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_36.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_37
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t37 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
}))
}))
_ = __local_var_7_24
// TAST (Let): Bind1_8_38 -> *Constructor_Control_Bind_Bind
Bind1_8_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_24, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_38
// TAST (Let): Applicative0_9_39 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_24, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_39
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_6_22
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_39.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_20.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t40 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_21.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_40
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t40 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
_ = monadMaybeT1_1_0
// TAST (Let): __local_var_2_41 -> *Constructor_Control_Monad_Monad
__local_var_2_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_2_41
// TAST (Let): Applicative0_3_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_41.V0), gopurs_runtime.Value{}))
_ = Applicative0_3_42
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_1_0
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_2_41.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_0, "ask"), gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_42.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_4})})
})))
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
// TAST (Let): monadMaybeT1_2_2 -> gopurs_runtime.Value
monadMaybeT1_2_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_75 -> gopurs_runtime.Value
__local_var_4_75 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_75
// TAST (Let): __local_var_4_74 -> gopurs_runtime.Value
__local_var_4_74 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_75, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_74
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_5
// TAST (Let): functorMaybeT1_5_4 -> gopurs_runtime.Value
functorMaybeT1_5_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_6), v_7)
})
}))
_ = functorMaybeT1_5_4
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_24 -> gopurs_runtime.Value
__local_var_7_24 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_24
// TAST (Let): __local_var_7_23 -> gopurs_runtime.Value
__local_var_7_23 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_24, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_8})})
})
_ = __local_var_7_23
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_8
// TAST (Let): functorMaybeT1_8_7 -> gopurs_runtime.Value
functorMaybeT1_8_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_8, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_9), v_10)
})
}))
_ = functorMaybeT1_8_7
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_10 -> *Constructor_Control_Bind_Bind
Bind1_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_10
// TAST (Let): Applicative0_11_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_11
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_13 -> gopurs_runtime.Value
__local_var_13_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_13
// TAST (Let): functorMaybeT1_13_12 -> gopurs_runtime.Value
functorMaybeT1_13_12 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_13, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_14), v_15)
})
}))
_ = functorMaybeT1_13_12
// TAST (Let): __local_var_14_14 -> gopurs_runtime.Value
__local_var_14_14 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_15 -> *Constructor_Control_Bind_Bind
Bind1_15_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_15
// TAST (Let): Applicative0_16_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_16
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_15.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t17 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_16.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_17
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t17 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
_ = __local_var_14_14
// TAST (Let): Bind1_15_18 -> *Constructor_Control_Bind_Bind
Bind1_15_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_18
// TAST (Let): Applicative0_16_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_14, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_13_12
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_18.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_18.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_19.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_10.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t20 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t20 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_11.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_20
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t20 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
_ = __local_var_9_9
// TAST (Let): Bind1_10_21 -> *Constructor_Control_Bind_Bind
Bind1_10_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_21
// TAST (Let): Applicative0_11_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_22
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_8_7
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_21.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_21.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_22.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_23, x_8)
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_25 -> *Constructor_Control_Bind_Bind
Bind1_7_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_25
// TAST (Let): Applicative0_8_26 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_26
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_28 -> gopurs_runtime.Value
__local_var_10_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_28
// TAST (Let): functorMaybeT1_10_27 -> gopurs_runtime.Value
functorMaybeT1_10_27 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_28, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_27
// TAST (Let): __local_var_11_29 -> gopurs_runtime.Value
__local_var_11_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_47 -> gopurs_runtime.Value
__local_var_12_47 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_47
// TAST (Let): __local_var_12_46 -> gopurs_runtime.Value
__local_var_12_46 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_47, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_13})})
})
_ = __local_var_12_46
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_31 -> gopurs_runtime.Value
__local_var_13_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_31
// TAST (Let): functorMaybeT1_13_30 -> gopurs_runtime.Value
functorMaybeT1_13_30 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_31, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_14), v_15)
})
}))
_ = functorMaybeT1_13_30
// TAST (Let): __local_var_14_32 -> gopurs_runtime.Value
__local_var_14_32 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_33 -> *Constructor_Control_Bind_Bind
Bind1_15_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_33
// TAST (Let): Applicative0_16_34 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_34
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_36 -> gopurs_runtime.Value
__local_var_18_36 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_36
// TAST (Let): functorMaybeT1_18_35 -> gopurs_runtime.Value
functorMaybeT1_18_35 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_36, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_19), v_20)
})
}))
_ = functorMaybeT1_18_35
// TAST (Let): __local_var_19_37 -> gopurs_runtime.Value
__local_var_19_37 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_38 -> *Constructor_Control_Bind_Bind
Bind1_20_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_38
// TAST (Let): Applicative0_21_39 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_39
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_38.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t40 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_39.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_40
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t40 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
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
_ = __local_var_19_37
// TAST (Let): Bind1_20_41 -> *Constructor_Control_Bind_Bind
Bind1_20_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_37, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_41
// TAST (Let): Applicative0_21_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_37, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_42
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_18_35
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_41.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_41.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_42.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_33.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t43 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t43 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_34.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_43
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t43 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
_ = __local_var_14_32
// TAST (Let): Bind1_15_44 -> *Constructor_Control_Bind_Bind
Bind1_15_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_32, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_44
// TAST (Let): Applicative0_16_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_32, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_45
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_13_30
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_44.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_44.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_45.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_46, x_13)
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_48 -> *Constructor_Control_Bind_Bind
Bind1_12_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_48
// TAST (Let): Applicative0_13_49 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_49
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_51 -> gopurs_runtime.Value
__local_var_15_51 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_51
// TAST (Let): functorMaybeT1_15_50 -> gopurs_runtime.Value
functorMaybeT1_15_50 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_51, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_16), v_17)
})
}))
_ = functorMaybeT1_15_50
// TAST (Let): __local_var_16_52 -> gopurs_runtime.Value
__local_var_16_52 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_62 -> gopurs_runtime.Value
__local_var_17_62 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_62
// TAST (Let): __local_var_17_61 -> gopurs_runtime.Value
__local_var_17_61 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_62, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_18})})
})
_ = __local_var_17_61
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_54 -> gopurs_runtime.Value
__local_var_18_54 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_54
// TAST (Let): functorMaybeT1_18_53 -> gopurs_runtime.Value
functorMaybeT1_18_53 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_54, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_19), v_20)
})
}))
_ = functorMaybeT1_18_53
// TAST (Let): __local_var_19_55 -> gopurs_runtime.Value
__local_var_19_55 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_56 -> *Constructor_Control_Bind_Bind
Bind1_20_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_56
// TAST (Let): Applicative0_21_57 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_57
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_56.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t58 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t58 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_57.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_58
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t58 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
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
}))
}))
_ = __local_var_19_55
// TAST (Let): Bind1_20_59 -> *Constructor_Control_Bind_Bind
Bind1_20_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_55, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_59
// TAST (Let): Applicative0_21_60 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_55, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_60
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_18_53
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_59.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_59.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_60.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_61, x_18)
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_63 -> *Constructor_Control_Bind_Bind
Bind1_17_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_63
// TAST (Let): Applicative0_18_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_64
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_63.V1), v_19, gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t65 gopurs_runtime.Value
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr == nil) {
__t65 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_64.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_65
} else {

}
}
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t65 = gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)
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
}))
}))
_ = __local_var_16_52
// TAST (Let): Bind1_17_66 -> *Constructor_Control_Bind_Bind
Bind1_17_66 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_52, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_66
// TAST (Let): Applicative0_18_67 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_67 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_52, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_67
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_15_50
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_66.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_66.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_67.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_48.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t68 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t68 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_49.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_68
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t68 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
goto end_branch_68
} else {

}
}
{
__t68 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_68:
return __t68
}))
})
}))
}))
_ = __local_var_11_29
// TAST (Let): Bind1_12_69 -> *Constructor_Control_Bind_Bind
Bind1_12_69 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_69
// TAST (Let): Applicative0_13_70 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_70
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_27
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_69.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_69.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_70.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_25.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t71 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t71 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_26.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_71
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t71 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
}))
}))
_ = __local_var_6_6
// TAST (Let): Bind1_7_72 -> *Constructor_Control_Bind_Bind
Bind1_7_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_72
// TAST (Let): Applicative0_8_73 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_73
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_5_4
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_72.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_72.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_73.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_74, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_76 -> *Constructor_Control_Bind_Bind
Bind1_4_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_76
// TAST (Let): Applicative0_5_77 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_77
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_79 -> gopurs_runtime.Value
__local_var_7_79 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_79
// TAST (Let): functorMaybeT1_7_78 -> gopurs_runtime.Value
functorMaybeT1_7_78 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_79, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_8), v_9)
})
}))
_ = functorMaybeT1_7_78
// TAST (Let): __local_var_8_80 -> gopurs_runtime.Value
__local_var_8_80 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_126 -> gopurs_runtime.Value
__local_var_9_126 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_126
// TAST (Let): __local_var_9_125 -> gopurs_runtime.Value
__local_var_9_125 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_126, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_125
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_82 -> gopurs_runtime.Value
__local_var_10_82 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_82
// TAST (Let): functorMaybeT1_10_81 -> gopurs_runtime.Value
functorMaybeT1_10_81 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_82, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_81
// TAST (Let): __local_var_11_83 -> gopurs_runtime.Value
__local_var_11_83 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_101 -> gopurs_runtime.Value
__local_var_12_101 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_101
// TAST (Let): __local_var_12_100 -> gopurs_runtime.Value
__local_var_12_100 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_101, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_13})})
})
_ = __local_var_12_100
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_85 -> gopurs_runtime.Value
__local_var_13_85 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_85
// TAST (Let): functorMaybeT1_13_84 -> gopurs_runtime.Value
functorMaybeT1_13_84 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_85, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_14), v_15)
})
}))
_ = functorMaybeT1_13_84
// TAST (Let): __local_var_14_86 -> gopurs_runtime.Value
__local_var_14_86 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_87 -> *Constructor_Control_Bind_Bind
Bind1_15_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_87
// TAST (Let): Applicative0_16_88 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_88
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_90 -> gopurs_runtime.Value
__local_var_18_90 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_90
// TAST (Let): functorMaybeT1_18_89 -> gopurs_runtime.Value
functorMaybeT1_18_89 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_90, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_19), v_20)
})
}))
_ = functorMaybeT1_18_89
// TAST (Let): __local_var_19_91 -> gopurs_runtime.Value
__local_var_19_91 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_92 -> *Constructor_Control_Bind_Bind
Bind1_20_92 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_92
// TAST (Let): Applicative0_21_93 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_93 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_93
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_92.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t94 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t94 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_93.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_94
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t94 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
goto end_branch_94
} else {

}
}
{
__t94 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_94:
return __t94
}))
})
}))
}))
_ = __local_var_19_91
// TAST (Let): Bind1_20_95 -> *Constructor_Control_Bind_Bind
Bind1_20_95 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_91, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_95
// TAST (Let): Applicative0_21_96 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_96 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_91, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_96
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_18_89
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_95.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_95.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_96.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_87.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t97 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t97 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_88.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_97
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t97 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_14_86
// TAST (Let): Bind1_15_98 -> *Constructor_Control_Bind_Bind
Bind1_15_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_86, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_98
// TAST (Let): Applicative0_16_99 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_99 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_86, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_99
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_13_84
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_98.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_98.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_99.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_100, x_13)
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_102 -> *Constructor_Control_Bind_Bind
Bind1_12_102 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_102
// TAST (Let): Applicative0_13_103 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_103 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_103
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_105 -> gopurs_runtime.Value
__local_var_15_105 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_105
// TAST (Let): functorMaybeT1_15_104 -> gopurs_runtime.Value
functorMaybeT1_15_104 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_105, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_16), v_17)
})
}))
_ = functorMaybeT1_15_104
// TAST (Let): __local_var_16_106 -> gopurs_runtime.Value
__local_var_16_106 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_17_116 -> gopurs_runtime.Value
__local_var_17_116 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_17_116
// TAST (Let): __local_var_17_115 -> gopurs_runtime.Value
__local_var_17_115 := gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_116, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_18})})
})
_ = __local_var_17_115
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_108 -> gopurs_runtime.Value
__local_var_18_108 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_108
// TAST (Let): functorMaybeT1_18_107 -> gopurs_runtime.Value
functorMaybeT1_18_107 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_108, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_19), v_20)
})
}))
_ = functorMaybeT1_18_107
// TAST (Let): __local_var_19_109 -> gopurs_runtime.Value
__local_var_19_109 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_110 -> *Constructor_Control_Bind_Bind
Bind1_20_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_110
// TAST (Let): Applicative0_21_111 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_111
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_110.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t112 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t112 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_111.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_112
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t112 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
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
_ = __local_var_19_109
// TAST (Let): Bind1_20_113 -> *Constructor_Control_Bind_Bind
Bind1_20_113 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_109, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_113
// TAST (Let): Applicative0_21_114 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_114 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_109, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_114
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_18_107
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_113.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_113.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_114.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_115, x_18)
}))
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_117 -> *Constructor_Control_Bind_Bind
Bind1_17_117 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_117
// TAST (Let): Applicative0_18_118 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_118 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_118
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_117.V1), v_19, gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t119 gopurs_runtime.Value
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr == nil) {
__t119 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_118.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_119
} else {

}
}
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t119 = gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)
goto end_branch_119
} else {

}
}
{
__t119 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_119:
return __t119
}))
})
}))
}))
_ = __local_var_16_106
// TAST (Let): Bind1_17_120 -> *Constructor_Control_Bind_Bind
Bind1_17_120 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_106, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_120
// TAST (Let): Applicative0_18_121 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_121 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_106, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_121
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_15_104
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_120.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_120.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_121.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_102.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t122 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t122 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_103.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_122
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t122 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
}))
}))
_ = __local_var_11_83
// TAST (Let): Bind1_12_123 -> *Constructor_Control_Bind_Bind
Bind1_12_123 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_83, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_123
// TAST (Let): Applicative0_13_124 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_124 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_83, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_124
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_81
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_123.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_123.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_124.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_125, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_127 -> *Constructor_Control_Bind_Bind
Bind1_9_127 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_127
// TAST (Let): Applicative0_10_128 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_128 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_128
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_130 -> gopurs_runtime.Value
__local_var_12_130 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_12_130
// TAST (Let): functorMaybeT1_12_129 -> gopurs_runtime.Value
functorMaybeT1_12_129 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_12_130, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_13), v_14)
})
}))
_ = functorMaybeT1_12_129
// TAST (Let): __local_var_13_131 -> gopurs_runtime.Value
__local_var_13_131 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_141 -> gopurs_runtime.Value
__local_var_14_141 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_14_141
// TAST (Let): __local_var_14_140 -> gopurs_runtime.Value
__local_var_14_140 := gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_141, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_15})})
})
_ = __local_var_14_140
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_133 -> gopurs_runtime.Value
__local_var_15_133 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_15_133
// TAST (Let): functorMaybeT1_15_132 -> gopurs_runtime.Value
functorMaybeT1_15_132 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_15_133, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_16), v_17)
})
}))
_ = functorMaybeT1_15_132
// TAST (Let): __local_var_16_134 -> gopurs_runtime.Value
__local_var_16_134 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_17_135 -> *Constructor_Control_Bind_Bind
Bind1_17_135 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_135
// TAST (Let): Applicative0_18_136 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_136 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_136
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_135.V1), v_19, gopurs_runtime.Func(func(v1_21 gopurs_runtime.Value) gopurs_runtime.Value {
var __t137 gopurs_runtime.Value
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr == nil) {
__t137 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_136.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_137
} else {

}
}
{
if (v1_21.Type == 9 && v1_21.IntVal == 930809136 && v1_21.UnsafePtr != nil) {
__t137 = gopurs_runtime.Apply(f_20, (*Constructor_Data_Maybe_Just)(v1_21.UnsafePtr).V0)
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
}))
_ = __local_var_16_134
// TAST (Let): Bind1_17_138 -> *Constructor_Control_Bind_Bind
Bind1_17_138 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_134, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_17_138
// TAST (Let): Applicative0_18_139 -> *Constructor_Control_Applicative_Applicative
Applicative0_18_139 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_16_134, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_18_139
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_15_132
}), gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_138.V1), f_19, gopurs_runtime.Func(func(f_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_17_138.V1), a_20, gopurs_runtime.Func(func(a_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_18_139.V1), gopurs_runtime.Apply(f_prime_21, a_prime_22))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_14_140, x_15)
}))
}), gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_14_142 -> *Constructor_Control_Bind_Bind
Bind1_14_142 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_142
// TAST (Let): Applicative0_15_143 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_143 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_143
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_3)
}), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_142.V1), v_16, gopurs_runtime.Func(func(v1_18 gopurs_runtime.Value) gopurs_runtime.Value {
var __t144 gopurs_runtime.Value
{
if (v1_18.Type == 9 && v1_18.IntVal == 930809136 && v1_18.UnsafePtr == nil) {
__t144 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_143.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_144
} else {

}
}
{
if (v1_18.Type == 9 && v1_18.IntVal == 930809136 && v1_18.UnsafePtr != nil) {
__t144 = gopurs_runtime.Apply(f_17, (*Constructor_Data_Maybe_Just)(v1_18.UnsafePtr).V0)
goto end_branch_144
} else {

}
}
{
__t144 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_144:
return __t144
}))
})
}))
}))
_ = __local_var_13_131
// TAST (Let): Bind1_14_145 -> *Constructor_Control_Bind_Bind
Bind1_14_145 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_131, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_14_145
// TAST (Let): Applicative0_15_146 -> *Constructor_Control_Applicative_Applicative
Applicative0_15_146 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_13_131, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_15_146
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_12_129
}), gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_145.V1), f_16, gopurs_runtime.Func(func(f_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_14_145.V1), a_17, gopurs_runtime.Func(func(a_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_15_146.V1), gopurs_runtime.Apply(f_prime_18, a_prime_19))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_127.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t147 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t147 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_128.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_147
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t147 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
}))
}))
_ = __local_var_8_80
// TAST (Let): Bind1_9_148 -> *Constructor_Control_Bind_Bind
Bind1_9_148 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_80, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_148
// TAST (Let): Applicative0_10_149 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_149 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_80, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_149
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_7_78
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_148.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_148.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_149.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_76.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t150 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t150 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_77.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_150
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t150 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
}))
}))
_ = monadMaybeT1_2_2
// TAST (Let): __local_var_3_151 -> *Constructor_Control_Monad_Monad
__local_var_3_151 := gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Monad0"), gopurs_runtime.Value{}))
_ = __local_var_3_151
// TAST (Let): Applicative0_4_152 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_152 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_151.V0), gopurs_runtime.Value{}))
_ = Applicative0_4_152
// TAST (Let): monadAskMaybeT1_1_0 -> gopurs_runtime.Value
monadAskMaybeT1_1_0 := gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_151.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(__local_var_1_1, "ask"), gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_152.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_5})})
})))
_ = monadAskMaybeT1_1_0
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskMaybeT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_153 -> gopurs_runtime.Value
__local_var_3_153 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_2)
_ = __local_var_3_153
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_153, v_4)
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
// TAST (Let): __local_var_3_19 -> gopurs_runtime.Value
__local_var_3_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_19
// TAST (Let): __local_var_3_18 -> gopurs_runtime.Value
__local_var_3_18 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_19, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_18
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorMaybeT1_4_2 -> gopurs_runtime.Value
functorMaybeT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_5), v_6)
})
}))
_ = functorMaybeT1_4_2
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_5 -> *Constructor_Control_Bind_Bind
Bind1_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_5
// TAST (Let): Applicative0_7_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorMaybeT1_9_7 -> gopurs_runtime.Value
functorMaybeT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_10), v_11)
})
}))
_ = functorMaybeT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
// TAST (Let): Applicative0_12_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_11
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t12 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_11.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_12
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t12 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_13 -> *Constructor_Control_Bind_Bind
Bind1_11_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_13
// TAST (Let): Applicative0_12_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_13.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_13.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_14.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_5.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_6.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_15
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t15 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
_ = __local_var_5_4
// TAST (Let): Bind1_6_16 -> *Constructor_Control_Bind_Bind
Bind1_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_16
// TAST (Let): Applicative0_7_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_4_2
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_16.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_16.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_17.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_18, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_20 -> *Constructor_Control_Bind_Bind
Bind1_3_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_20
// TAST (Let): Applicative0_4_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_21
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_23 -> gopurs_runtime.Value
__local_var_6_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_23
// TAST (Let): functorMaybeT1_6_22 -> gopurs_runtime.Value
functorMaybeT1_6_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_7), v_8)
})
}))
_ = functorMaybeT1_6_22
// TAST (Let): __local_var_7_24 -> gopurs_runtime.Value
__local_var_7_24 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_34 -> gopurs_runtime.Value
__local_var_8_34 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_34
// TAST (Let): __local_var_8_33 -> gopurs_runtime.Value
__local_var_8_33 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_34, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_33
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_26 -> gopurs_runtime.Value
__local_var_9_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_26
// TAST (Let): functorMaybeT1_9_25 -> gopurs_runtime.Value
functorMaybeT1_9_25 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_26, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_10), v_11)
})
}))
_ = functorMaybeT1_9_25
// TAST (Let): __local_var_10_27 -> gopurs_runtime.Value
__local_var_10_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_28 -> *Constructor_Control_Bind_Bind
Bind1_11_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_28
// TAST (Let): Applicative0_12_29 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_29
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_28.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t30 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_29.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_30
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t30 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_10_27
// TAST (Let): Bind1_11_31 -> *Constructor_Control_Bind_Bind
Bind1_11_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_31
// TAST (Let): Applicative0_12_32 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_32
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_9_25
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_31.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_31.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_32.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_33, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_35 -> *Constructor_Control_Bind_Bind
Bind1_8_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_35
// TAST (Let): Applicative0_9_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_36
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_1_1)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_35.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t37 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_36.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_37
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t37 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
}))
}))
_ = __local_var_7_24
// TAST (Let): Bind1_8_38 -> *Constructor_Control_Bind_Bind
Bind1_8_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_24, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_38
// TAST (Let): Applicative0_9_39 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_24, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_39
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_6_22
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_39.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_20.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t40 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_21.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_40
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t40 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
_ = monadMaybeT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_4})})
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
// TAST (Let): __local_var_3_19 -> gopurs_runtime.Value
__local_var_3_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_19
// TAST (Let): __local_var_3_18 -> gopurs_runtime.Value
__local_var_3_18 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_19, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_18
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorMaybeT1_4_2 -> gopurs_runtime.Value
functorMaybeT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_5), v_6)
})
}))
_ = functorMaybeT1_4_2
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_5 -> *Constructor_Control_Bind_Bind
Bind1_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_5
// TAST (Let): Applicative0_7_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorMaybeT1_9_7 -> gopurs_runtime.Value
functorMaybeT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_10), v_11)
})
}))
_ = functorMaybeT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
// TAST (Let): Applicative0_12_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_11
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t12 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_11.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_12
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t12 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_13 -> *Constructor_Control_Bind_Bind
Bind1_11_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_13
// TAST (Let): Applicative0_12_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_13.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_13.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_14.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_5.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_6.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_15
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t15 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
_ = __local_var_5_4
// TAST (Let): Bind1_6_16 -> *Constructor_Control_Bind_Bind
Bind1_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_16
// TAST (Let): Applicative0_7_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_4_2
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_16.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_16.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_17.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_18, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_20 -> *Constructor_Control_Bind_Bind
Bind1_3_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_20
// TAST (Let): Applicative0_4_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_21
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_23 -> gopurs_runtime.Value
__local_var_6_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_23
// TAST (Let): functorMaybeT1_6_22 -> gopurs_runtime.Value
functorMaybeT1_6_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_7), v_8)
})
}))
_ = functorMaybeT1_6_22
// TAST (Let): __local_var_7_24 -> gopurs_runtime.Value
__local_var_7_24 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_34 -> gopurs_runtime.Value
__local_var_8_34 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_34
// TAST (Let): __local_var_8_33 -> gopurs_runtime.Value
__local_var_8_33 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_34, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_33
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_26 -> gopurs_runtime.Value
__local_var_9_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_26
// TAST (Let): functorMaybeT1_9_25 -> gopurs_runtime.Value
functorMaybeT1_9_25 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_26, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_10), v_11)
})
}))
_ = functorMaybeT1_9_25
// TAST (Let): __local_var_10_27 -> gopurs_runtime.Value
__local_var_10_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_28 -> *Constructor_Control_Bind_Bind
Bind1_11_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_28
// TAST (Let): Applicative0_12_29 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_29
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_28.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t30 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_29.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_30
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t30 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_10_27
// TAST (Let): Bind1_11_31 -> *Constructor_Control_Bind_Bind
Bind1_11_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_31
// TAST (Let): Applicative0_12_32 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_32
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_9_25
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_31.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_31.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_32.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_33, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_35 -> *Constructor_Control_Bind_Bind
Bind1_8_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_35
// TAST (Let): Applicative0_9_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_36
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_35.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t37 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_36.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_37
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t37 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
}))
}))
_ = __local_var_7_24
// TAST (Let): Bind1_8_38 -> *Constructor_Control_Bind_Bind
Bind1_8_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_24, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_38
// TAST (Let): Applicative0_9_39 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_24, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_39
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_6_22
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_39.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_20.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t40 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_21.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_40
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t40 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
_ = monadMaybeT1_2_1
// TAST (Let): Bind1_3_43 -> *Constructor_Control_Bind_Bind
Bind1_3_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_43
// TAST (Let): Applicative0_4_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_44
// TAST (Let): __local_var_3_42 -> gopurs_runtime.Value
__local_var_3_42 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_43.V1), a_5, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_44.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_6})})
}))
})
_ = __local_var_3_42
// TAST (Let): __local_var_3_41 -> gopurs_runtime.Value
__local_var_3_41 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_42, x_4)
})
_ = __local_var_3_41
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_41, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
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
// TAST (Let): monadMaybeT1_4_3 -> gopurs_runtime.Value
monadMaybeT1_4_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_21 -> gopurs_runtime.Value
__local_var_5_21 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_21
// TAST (Let): __local_var_5_20 -> gopurs_runtime.Value
__local_var_5_20 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_21, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_6})})
})
_ = __local_var_5_20
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_5
// TAST (Let): functorMaybeT1_6_4 -> gopurs_runtime.Value
functorMaybeT1_6_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_5, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_7), v_8)
})
}))
_ = functorMaybeT1_6_4
// TAST (Let): __local_var_7_6 -> gopurs_runtime.Value
__local_var_7_6 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_7 -> *Constructor_Control_Bind_Bind
Bind1_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_7
// TAST (Let): Applicative0_9_8 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_8
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_10
// TAST (Let): functorMaybeT1_11_9 -> gopurs_runtime.Value
functorMaybeT1_11_9 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_10, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_12), v_13)
})
}))
_ = functorMaybeT1_11_9
// TAST (Let): __local_var_12_11 -> gopurs_runtime.Value
__local_var_12_11 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_12 -> *Constructor_Control_Bind_Bind
Bind1_13_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_12
// TAST (Let): Applicative0_14_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_13
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_12.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_13.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_14
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t14 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
_ = __local_var_12_11
// TAST (Let): Bind1_13_15 -> *Constructor_Control_Bind_Bind
Bind1_13_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_11, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_15
// TAST (Let): Applicative0_14_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_11, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_11_9
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_15.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_15.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_16.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_7.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t17 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t17 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_8.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_17
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t17 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
_ = __local_var_7_6
// TAST (Let): Bind1_8_18 -> *Constructor_Control_Bind_Bind
Bind1_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_18
// TAST (Let): Applicative0_9_19 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_6, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_19
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_6_4
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_18.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_19.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_20, x_6)
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_22 -> *Constructor_Control_Bind_Bind
Bind1_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_22
// TAST (Let): Applicative0_6_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_23
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_25 -> gopurs_runtime.Value
__local_var_8_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_25
// TAST (Let): functorMaybeT1_8_24 -> gopurs_runtime.Value
functorMaybeT1_8_24 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_25, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_9), v_10)
})
}))
_ = functorMaybeT1_8_24
// TAST (Let): __local_var_9_26 -> gopurs_runtime.Value
__local_var_9_26 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_36 -> gopurs_runtime.Value
__local_var_10_36 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_36
// TAST (Let): __local_var_10_35 -> gopurs_runtime.Value
__local_var_10_35 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_36, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_11})})
})
_ = __local_var_10_35
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_28 -> gopurs_runtime.Value
__local_var_11_28 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_28
// TAST (Let): functorMaybeT1_11_27 -> gopurs_runtime.Value
functorMaybeT1_11_27 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_28, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_12), v_13)
})
}))
_ = functorMaybeT1_11_27
// TAST (Let): __local_var_12_29 -> gopurs_runtime.Value
__local_var_12_29 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_30 -> *Constructor_Control_Bind_Bind
Bind1_13_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_30
// TAST (Let): Applicative0_14_31 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_31
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_30.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t32 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t32 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_31.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_32
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t32 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
}))
_ = __local_var_12_29
// TAST (Let): Bind1_13_33 -> *Constructor_Control_Bind_Bind
Bind1_13_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_29, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_33
// TAST (Let): Applicative0_14_34 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_29, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_34
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_11_27
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_33.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_33.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_34.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_35, x_11)
}))
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_37 -> *Constructor_Control_Bind_Bind
Bind1_10_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_37
// TAST (Let): Applicative0_11_38 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_38
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_37.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t39 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t39 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_38.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_39
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t39 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
}))
}))
_ = __local_var_9_26
// TAST (Let): Bind1_10_40 -> *Constructor_Control_Bind_Bind
Bind1_10_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_26, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_40
// TAST (Let): Applicative0_11_41 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_26, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_41
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_8_24
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_40.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_40.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_41.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_22.V1), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t42 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr == nil) {
__t42 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_23.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_42
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t42 = gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)
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
_ = monadMaybeT1_4_3
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_4_3
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_43 -> gopurs_runtime.Value
__local_var_6_43 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(f_5, a_6), gopurs_runtime.Func(func(m_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t47 *Constructor_Control_Monad_Rec_Class_Done
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 930809136 && m_prime_7.UnsafePtr == nil) {
__t47 = &Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_47
} else {

}
}
{
if (m_prime_7.Type == 9 && m_prime_7.IntVal == 930809136 && m_prime_7.UnsafePtr != nil) {
var __t46 gopurs_runtime.Value
{
var __t_tag_44 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(m_prime_7.UnsafePtr).V0
if (__t_tag_44.Type == 9 && __t_tag_44.IntVal == 525585346) {
__t46 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, (*Constructor_Control_Monad_Rec_Class_Loop)((*Constructor_Data_Maybe_Just)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}
goto end_branch_46
} else {

}
}
{
var __t_tag_45 gopurs_runtime.Value = (*Constructor_Data_Maybe_Just)(m_prime_7.UnsafePtr).V0
if (__t_tag_45.Type == 9 && __t_tag_45.IntVal == 60402430) {
__t46 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, (*Constructor_Control_Monad_Rec_Class_Done)((*Constructor_Data_Maybe_Just)(m_prime_7.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_46
} else {

}
}
{
__t46 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_46:
__t47 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](__t46)
goto end_branch_47
} else {

}
}
{
__t47 = gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_Done](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_47:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(__t47)})
}))
}))
_ = __local_var_6_43
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_43, x_7)
})
}))
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
// TAST (Let): monadMaybeT1_2_1 -> gopurs_runtime.Value
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_20 -> gopurs_runtime.Value
__local_var_4_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_20
// TAST (Let): __local_var_4_19 -> gopurs_runtime.Value
__local_var_4_19 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_20, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorMaybeT1_5_3 -> gopurs_runtime.Value
functorMaybeT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_6), v_7)
})
}))
_ = functorMaybeT1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
// TAST (Let): Applicative0_8_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_7
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_9
// TAST (Let): functorMaybeT1_10_8 -> gopurs_runtime.Value
functorMaybeT1_10_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_8
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_11 -> *Constructor_Control_Bind_Bind
Bind1_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_11
// TAST (Let): Applicative0_13_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_12
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_11.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_12.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_13
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t13 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
_ = __local_var_11_10
// TAST (Let): Bind1_12_14 -> *Constructor_Control_Bind_Bind
Bind1_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_14
// TAST (Let): Applicative0_13_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_15
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_8
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_14.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_14.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_15.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t16 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_7.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_16
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t16 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
_ = __local_var_6_5
// TAST (Let): Bind1_7_17 -> *Constructor_Control_Bind_Bind
Bind1_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_17
// TAST (Let): Applicative0_8_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_17.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_17.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_18.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_19, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_21 -> *Constructor_Control_Bind_Bind
Bind1_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_21
// TAST (Let): Applicative0_5_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_22
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_24 -> gopurs_runtime.Value
__local_var_7_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_24
// TAST (Let): functorMaybeT1_7_23 -> gopurs_runtime.Value
functorMaybeT1_7_23 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_24, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_8), v_9)
})
}))
_ = functorMaybeT1_7_23
// TAST (Let): __local_var_8_25 -> gopurs_runtime.Value
__local_var_8_25 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_35 -> gopurs_runtime.Value
__local_var_9_35 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_35
// TAST (Let): __local_var_9_34 -> gopurs_runtime.Value
__local_var_9_34 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_35, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_34
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_27 -> gopurs_runtime.Value
__local_var_10_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_27
// TAST (Let): functorMaybeT1_10_26 -> gopurs_runtime.Value
functorMaybeT1_10_26 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_27, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_26
// TAST (Let): __local_var_11_28 -> gopurs_runtime.Value
__local_var_11_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_29 -> *Constructor_Control_Bind_Bind
Bind1_12_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_29
// TAST (Let): Applicative0_13_30 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_30
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_29.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t31 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_30.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_31
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t31 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_11_28
// TAST (Let): Bind1_12_32 -> *Constructor_Control_Bind_Bind
Bind1_12_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_32
// TAST (Let): Applicative0_13_33 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_33
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_26
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_32.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_32.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_33.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_34, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_36 -> *Constructor_Control_Bind_Bind
Bind1_9_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_36
// TAST (Let): Applicative0_10_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_37
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_36.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t38 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_37.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_38
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t38 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
}))
}))
_ = __local_var_8_25
// TAST (Let): Bind1_9_39 -> *Constructor_Control_Bind_Bind
Bind1_9_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_25, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_39
// TAST (Let): Applicative0_10_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_25, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_40
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_7_23
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_39.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_39.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_40.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_21.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t41 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_22.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_41
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t41 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Applicative0_4_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_4_42
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), f_3), gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_42.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_5})})
}))
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
// TAST (Let): __local_var_4_20 -> gopurs_runtime.Value
__local_var_4_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_20
// TAST (Let): __local_var_4_19 -> gopurs_runtime.Value
__local_var_4_19 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_20, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorMaybeT1_5_3 -> gopurs_runtime.Value
functorMaybeT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_6), v_7)
})
}))
_ = functorMaybeT1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
// TAST (Let): Applicative0_8_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_7
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_9
// TAST (Let): functorMaybeT1_10_8 -> gopurs_runtime.Value
functorMaybeT1_10_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_8
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_11 -> *Constructor_Control_Bind_Bind
Bind1_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_11
// TAST (Let): Applicative0_13_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_12
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_1_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_11.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_12.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_13
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t13 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
_ = __local_var_11_10
// TAST (Let): Bind1_12_14 -> *Constructor_Control_Bind_Bind
Bind1_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_14
// TAST (Let): Applicative0_13_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_15
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_8
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_14.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_14.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_15.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t16 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_7.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_16
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t16 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
_ = __local_var_6_5
// TAST (Let): Bind1_7_17 -> *Constructor_Control_Bind_Bind
Bind1_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_17
// TAST (Let): Applicative0_8_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_17.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_17.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_18.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_19, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_21 -> *Constructor_Control_Bind_Bind
Bind1_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_21
// TAST (Let): Applicative0_5_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_22
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_24 -> gopurs_runtime.Value
__local_var_7_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_24
// TAST (Let): functorMaybeT1_7_23 -> gopurs_runtime.Value
functorMaybeT1_7_23 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_24, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_8), v_9)
})
}))
_ = functorMaybeT1_7_23
// TAST (Let): __local_var_8_25 -> gopurs_runtime.Value
__local_var_8_25 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_35 -> gopurs_runtime.Value
__local_var_9_35 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_35
// TAST (Let): __local_var_9_34 -> gopurs_runtime.Value
__local_var_9_34 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_35, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_34
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_27 -> gopurs_runtime.Value
__local_var_10_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_27
// TAST (Let): functorMaybeT1_10_26 -> gopurs_runtime.Value
functorMaybeT1_10_26 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_27, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_26
// TAST (Let): __local_var_11_28 -> gopurs_runtime.Value
__local_var_11_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_29 -> *Constructor_Control_Bind_Bind
Bind1_12_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_29
// TAST (Let): Applicative0_13_30 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_30
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_1_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_29.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t31 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_30.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_31
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t31 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_11_28
// TAST (Let): Bind1_12_32 -> *Constructor_Control_Bind_Bind
Bind1_12_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_32
// TAST (Let): Applicative0_13_33 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_33
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_26
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_32.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_32.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_33.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_34, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_36 -> *Constructor_Control_Bind_Bind
Bind1_9_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_36
// TAST (Let): Applicative0_10_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_37
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_1_0)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_36.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t38 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_37.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_38
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t38 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
}))
}))
_ = __local_var_8_25
// TAST (Let): Bind1_9_39 -> *Constructor_Control_Bind_Bind
Bind1_9_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_25, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_39
// TAST (Let): Applicative0_10_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_25, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_40
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_7_23
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_39.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_39.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_40.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_21.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t41 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_22.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_41
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t41 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
_ = monadMaybeT1_3_2
// TAST (Let): Bind1_4_44 -> *Constructor_Control_Bind_Bind
Bind1_4_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_44
// TAST (Let): Applicative0_5_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_45
// TAST (Let): __local_var_4_43 -> gopurs_runtime.Value
__local_var_4_43 := gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_44.V1), a_6, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_45.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_7})})
}))
})
_ = __local_var_4_43
// TAST (Let): __local_var_4_42 -> gopurs_runtime.Value
__local_var_4_42 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_43, x_5)
})
_ = __local_var_4_42
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_42, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
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
// TAST (Let): monadMaybeT1_9_9 -> gopurs_runtime.Value
monadMaybeT1_9_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_81 -> gopurs_runtime.Value
__local_var_10_81 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_81
// TAST (Let): __local_var_10_80 -> gopurs_runtime.Value
__local_var_10_80 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_81, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_11})})
})
_ = __local_var_10_80
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_11 -> gopurs_runtime.Value
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_11
// TAST (Let): functorMaybeT1_11_10 -> gopurs_runtime.Value
functorMaybeT1_11_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_11, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_12), v_13)
})
}))
_ = functorMaybeT1_11_10
// TAST (Let): __local_var_12_12 -> gopurs_runtime.Value
__local_var_12_12 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_30 -> gopurs_runtime.Value
__local_var_13_30 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_30
// TAST (Let): __local_var_13_29 -> gopurs_runtime.Value
__local_var_13_29 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_30, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_14})})
})
_ = __local_var_13_29
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_14 -> gopurs_runtime.Value
__local_var_14_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_14
// TAST (Let): functorMaybeT1_14_13 -> gopurs_runtime.Value
functorMaybeT1_14_13 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_14, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_15), v_16)
})
}))
_ = functorMaybeT1_14_13
// TAST (Let): __local_var_15_15 -> gopurs_runtime.Value
__local_var_15_15 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_16 -> *Constructor_Control_Bind_Bind
Bind1_16_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_16
// TAST (Let): Applicative0_17_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_17
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_19 -> gopurs_runtime.Value
__local_var_19_19 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_19
// TAST (Let): functorMaybeT1_19_18 -> gopurs_runtime.Value
functorMaybeT1_19_18 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_19, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_20), v_21)
})
}))
_ = functorMaybeT1_19_18
// TAST (Let): __local_var_20_20 -> gopurs_runtime.Value
__local_var_20_20 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_21 -> *Constructor_Control_Bind_Bind
Bind1_21_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_21
// TAST (Let): Applicative0_22_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_22
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_21.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t23 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t23 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_22.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_23
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t23 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
_ = __local_var_20_20
// TAST (Let): Bind1_21_24 -> *Constructor_Control_Bind_Bind
Bind1_21_24 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_20, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_24
// TAST (Let): Applicative0_22_25 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_25 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_20, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_25
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_19_18
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_24.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_24.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_25.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_16.V1), v_18, gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 gopurs_runtime.Value
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr == nil) {
__t26 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_17.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_26
} else {

}
}
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr != nil) {
__t26 = gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_20.UnsafePtr).V0)
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
_ = __local_var_15_15
// TAST (Let): Bind1_16_27 -> *Constructor_Control_Bind_Bind
Bind1_16_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_27
// TAST (Let): Applicative0_17_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_28
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_14_13
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_27.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_27.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_28.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_29, x_14)
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_31 -> *Constructor_Control_Bind_Bind
Bind1_13_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_31
// TAST (Let): Applicative0_14_32 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_32
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_34 -> gopurs_runtime.Value
__local_var_16_34 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_34
// TAST (Let): functorMaybeT1_16_33 -> gopurs_runtime.Value
functorMaybeT1_16_33 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_34, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_17), v_18)
})
}))
_ = functorMaybeT1_16_33
// TAST (Let): __local_var_17_35 -> gopurs_runtime.Value
__local_var_17_35 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_53 -> gopurs_runtime.Value
__local_var_18_53 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_53
// TAST (Let): __local_var_18_52 -> gopurs_runtime.Value
__local_var_18_52 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_53, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_19})})
})
_ = __local_var_18_52
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_37 -> gopurs_runtime.Value
__local_var_19_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_37
// TAST (Let): functorMaybeT1_19_36 -> gopurs_runtime.Value
functorMaybeT1_19_36 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_37, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_20), v_21)
})
}))
_ = functorMaybeT1_19_36
// TAST (Let): __local_var_20_38 -> gopurs_runtime.Value
__local_var_20_38 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_39 -> *Constructor_Control_Bind_Bind
Bind1_21_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_39
// TAST (Let): Applicative0_22_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_40
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_42 -> gopurs_runtime.Value
__local_var_24_42 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_42
// TAST (Let): functorMaybeT1_24_41 -> gopurs_runtime.Value
functorMaybeT1_24_41 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_24_42, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_25), v_26)
})
}))
_ = functorMaybeT1_24_41
// TAST (Let): __local_var_25_43 -> gopurs_runtime.Value
__local_var_25_43 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_44 -> *Constructor_Control_Bind_Bind
Bind1_26_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_44
// TAST (Let): Applicative0_27_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_45
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_44.V1), v_28, gopurs_runtime.Func(func(v1_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t46 gopurs_runtime.Value
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr == nil) {
__t46 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_45.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_46
} else {

}
}
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr != nil) {
__t46 = gopurs_runtime.Apply(f_29, (*Constructor_Data_Maybe_Just)(v1_30.UnsafePtr).V0)
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
_ = __local_var_25_43
// TAST (Let): Bind1_26_47 -> *Constructor_Control_Bind_Bind
Bind1_26_47 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_43, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_47
// TAST (Let): Applicative0_27_48 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_43, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_48
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_24_41
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_47.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_47.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_48.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_39.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t49 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t49 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_40.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_49
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t49 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
_ = __local_var_20_38
// TAST (Let): Bind1_21_50 -> *Constructor_Control_Bind_Bind
Bind1_21_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_38, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_50
// TAST (Let): Applicative0_22_51 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_51 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_38, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_51
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_19_36
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_50.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_50.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_51.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_52, x_19)
}))
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_54 -> *Constructor_Control_Bind_Bind
Bind1_18_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_54
// TAST (Let): Applicative0_19_55 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_55
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_57 -> gopurs_runtime.Value
__local_var_21_57 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_57
// TAST (Let): functorMaybeT1_21_56 -> gopurs_runtime.Value
functorMaybeT1_21_56 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_21_57, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_22), v_23)
})
}))
_ = functorMaybeT1_21_56
// TAST (Let): __local_var_22_58 -> gopurs_runtime.Value
__local_var_22_58 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_68 -> gopurs_runtime.Value
__local_var_23_68 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_23_68
// TAST (Let): __local_var_23_67 -> gopurs_runtime.Value
__local_var_23_67 := gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_68, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_24})})
})
_ = __local_var_23_67
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_60 -> gopurs_runtime.Value
__local_var_24_60 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_60
// TAST (Let): functorMaybeT1_24_59 -> gopurs_runtime.Value
functorMaybeT1_24_59 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_24_60, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_25), v_26)
})
}))
_ = functorMaybeT1_24_59
// TAST (Let): __local_var_25_61 -> gopurs_runtime.Value
__local_var_25_61 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_62 -> *Constructor_Control_Bind_Bind
Bind1_26_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_62
// TAST (Let): Applicative0_27_63 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_63
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_62.V1), v_28, gopurs_runtime.Func(func(v1_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t64 gopurs_runtime.Value
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr == nil) {
__t64 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_63.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_64
} else {

}
}
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr != nil) {
__t64 = gopurs_runtime.Apply(f_29, (*Constructor_Data_Maybe_Just)(v1_30.UnsafePtr).V0)
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
_ = __local_var_25_61
// TAST (Let): Bind1_26_65 -> *Constructor_Control_Bind_Bind
Bind1_26_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_61, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_65
// TAST (Let): Applicative0_27_66 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_66 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_61, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_66
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_24_59
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_65.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_65.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_66.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_67, x_24)
}))
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_69 -> *Constructor_Control_Bind_Bind
Bind1_23_69 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_69
// TAST (Let): Applicative0_24_70 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_70
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_69.V1), v_25, gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t71 gopurs_runtime.Value
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr == nil) {
__t71 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_70.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_71
} else {

}
}
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t71 = gopurs_runtime.Apply(f_26, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)
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
}))
}))
_ = __local_var_22_58
// TAST (Let): Bind1_23_72 -> *Constructor_Control_Bind_Bind
Bind1_23_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_58, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_72
// TAST (Let): Applicative0_24_73 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_58, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_73
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_21_56
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_72.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_72.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_73.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_54.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t74 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t74 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_55.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
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
}))
}))
_ = __local_var_17_35
// TAST (Let): Bind1_18_75 -> *Constructor_Control_Bind_Bind
Bind1_18_75 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_35, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_75
// TAST (Let): Applicative0_19_76 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_35, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_76
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_16_33
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_75.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_75.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_76.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_31.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t77 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t77 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_32.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
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
}))
}))
_ = __local_var_12_12
// TAST (Let): Bind1_13_78 -> *Constructor_Control_Bind_Bind
Bind1_13_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_12, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_78
// TAST (Let): Applicative0_14_79 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_12, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_79
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_11_10
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_78.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_78.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_79.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_80, x_11)
}))
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_82 -> *Constructor_Control_Bind_Bind
Bind1_10_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_82
// TAST (Let): Applicative0_11_83 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_83 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_83
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_85 -> gopurs_runtime.Value
__local_var_13_85 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_85
// TAST (Let): functorMaybeT1_13_84 -> gopurs_runtime.Value
functorMaybeT1_13_84 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_85, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_14), v_15)
})
}))
_ = functorMaybeT1_13_84
// TAST (Let): __local_var_14_86 -> gopurs_runtime.Value
__local_var_14_86 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_132 -> gopurs_runtime.Value
__local_var_15_132 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_15_132
// TAST (Let): __local_var_15_131 -> gopurs_runtime.Value
__local_var_15_131 := gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_132, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_16})})
})
_ = __local_var_15_131
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_88 -> gopurs_runtime.Value
__local_var_16_88 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_88
// TAST (Let): functorMaybeT1_16_87 -> gopurs_runtime.Value
functorMaybeT1_16_87 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_88, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_17), v_18)
})
}))
_ = functorMaybeT1_16_87
// TAST (Let): __local_var_17_89 -> gopurs_runtime.Value
__local_var_17_89 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_107 -> gopurs_runtime.Value
__local_var_18_107 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_107
// TAST (Let): __local_var_18_106 -> gopurs_runtime.Value
__local_var_18_106 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_107, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_19})})
})
_ = __local_var_18_106
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_91 -> gopurs_runtime.Value
__local_var_19_91 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_91
// TAST (Let): functorMaybeT1_19_90 -> gopurs_runtime.Value
functorMaybeT1_19_90 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_91, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_20), v_21)
})
}))
_ = functorMaybeT1_19_90
// TAST (Let): __local_var_20_92 -> gopurs_runtime.Value
__local_var_20_92 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_93 -> *Constructor_Control_Bind_Bind
Bind1_21_93 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_93
// TAST (Let): Applicative0_22_94 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_94
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_96 -> gopurs_runtime.Value
__local_var_24_96 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_96
// TAST (Let): functorMaybeT1_24_95 -> gopurs_runtime.Value
functorMaybeT1_24_95 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_24_96, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_25), v_26)
})
}))
_ = functorMaybeT1_24_95
// TAST (Let): __local_var_25_97 -> gopurs_runtime.Value
__local_var_25_97 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_98 -> *Constructor_Control_Bind_Bind
Bind1_26_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_98
// TAST (Let): Applicative0_27_99 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_99 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_99
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_98.V1), v_28, gopurs_runtime.Func(func(v1_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t100 gopurs_runtime.Value
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr == nil) {
__t100 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_99.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_100
} else {

}
}
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr != nil) {
__t100 = gopurs_runtime.Apply(f_29, (*Constructor_Data_Maybe_Just)(v1_30.UnsafePtr).V0)
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
}))
}))
_ = __local_var_25_97
// TAST (Let): Bind1_26_101 -> *Constructor_Control_Bind_Bind
Bind1_26_101 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_97, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_101
// TAST (Let): Applicative0_27_102 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_102 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_97, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_102
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_24_95
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_101.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_101.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_102.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_93.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t103 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t103 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_94.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_103
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t103 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
}))
}))
_ = __local_var_20_92
// TAST (Let): Bind1_21_104 -> *Constructor_Control_Bind_Bind
Bind1_21_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_92, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_104
// TAST (Let): Applicative0_22_105 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_105 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_92, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_105
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_19_90
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_104.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_104.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_105.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_106, x_19)
}))
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_108 -> *Constructor_Control_Bind_Bind
Bind1_18_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_108
// TAST (Let): Applicative0_19_109 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_109 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_109
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_111 -> gopurs_runtime.Value
__local_var_21_111 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_111
// TAST (Let): functorMaybeT1_21_110 -> gopurs_runtime.Value
functorMaybeT1_21_110 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_21_111, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_22), v_23)
})
}))
_ = functorMaybeT1_21_110
// TAST (Let): __local_var_22_112 -> gopurs_runtime.Value
__local_var_22_112 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_23_122 -> gopurs_runtime.Value
__local_var_23_122 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_23_122
// TAST (Let): __local_var_23_121 -> gopurs_runtime.Value
__local_var_23_121 := gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_122, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_24})})
})
_ = __local_var_23_121
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_24_114 -> gopurs_runtime.Value
__local_var_24_114 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_24_114
// TAST (Let): functorMaybeT1_24_113 -> gopurs_runtime.Value
functorMaybeT1_24_113 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_24_114, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_25), v_26)
})
}))
_ = functorMaybeT1_24_113
// TAST (Let): __local_var_25_115 -> gopurs_runtime.Value
__local_var_25_115 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_26_116 -> *Constructor_Control_Bind_Bind
Bind1_26_116 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_116
// TAST (Let): Applicative0_27_117 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_117 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_117
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_28 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_116.V1), v_28, gopurs_runtime.Func(func(v1_30 gopurs_runtime.Value) gopurs_runtime.Value {
var __t118 gopurs_runtime.Value
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr == nil) {
__t118 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_117.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_118
} else {

}
}
{
if (v1_30.Type == 9 && v1_30.IntVal == 930809136 && v1_30.UnsafePtr != nil) {
__t118 = gopurs_runtime.Apply(f_29, (*Constructor_Data_Maybe_Just)(v1_30.UnsafePtr).V0)
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
}))
}))
_ = __local_var_25_115
// TAST (Let): Bind1_26_119 -> *Constructor_Control_Bind_Bind
Bind1_26_119 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_115, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_26_119
// TAST (Let): Applicative0_27_120 -> *Constructor_Control_Applicative_Applicative
Applicative0_27_120 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_25_115, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_27_120
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_24_113
}), gopurs_runtime.Func(func(f_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_119.V1), f_28, gopurs_runtime.Func(func(f_prime_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_26_119.V1), a_29, gopurs_runtime.Func(func(a_prime_31 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_27_120.V1), gopurs_runtime.Apply(f_prime_30, a_prime_31))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_23_121, x_24)
}))
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_123 -> *Constructor_Control_Bind_Bind
Bind1_23_123 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_123
// TAST (Let): Applicative0_24_124 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_124 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_124
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_123.V1), v_25, gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t125 gopurs_runtime.Value
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr == nil) {
__t125 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_124.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_125
} else {

}
}
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t125 = gopurs_runtime.Apply(f_26, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)
goto end_branch_125
} else {

}
}
{
__t125 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_125:
return __t125
}))
})
}))
}))
_ = __local_var_22_112
// TAST (Let): Bind1_23_126 -> *Constructor_Control_Bind_Bind
Bind1_23_126 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_112, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_126
// TAST (Let): Applicative0_24_127 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_127 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_112, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_127
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_21_110
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_126.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_126.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_127.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_108.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t128 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t128 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_109.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_128
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t128 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
goto end_branch_128
} else {

}
}
{
__t128 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_128:
return __t128
}))
})
}))
}))
_ = __local_var_17_89
// TAST (Let): Bind1_18_129 -> *Constructor_Control_Bind_Bind
Bind1_18_129 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_89, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_129
// TAST (Let): Applicative0_19_130 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_130 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_89, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_130
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_16_87
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_129.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_129.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_130.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_131, x_16)
}))
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_133 -> *Constructor_Control_Bind_Bind
Bind1_15_133 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_133
// TAST (Let): Applicative0_16_134 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_134 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_134
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_136 -> gopurs_runtime.Value
__local_var_18_136 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_18_136
// TAST (Let): functorMaybeT1_18_135 -> gopurs_runtime.Value
functorMaybeT1_18_135 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_18_136, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_19), v_20)
})
}))
_ = functorMaybeT1_18_135
// TAST (Let): __local_var_19_137 -> gopurs_runtime.Value
__local_var_19_137 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_20_147 -> gopurs_runtime.Value
__local_var_20_147 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_20_147
// TAST (Let): __local_var_20_146 -> gopurs_runtime.Value
__local_var_20_146 := gopurs_runtime.Func(func(x_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_147, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_21})})
})
_ = __local_var_20_146
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_21_139 -> gopurs_runtime.Value
__local_var_21_139 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_21_139
// TAST (Let): functorMaybeT1_21_138 -> gopurs_runtime.Value
functorMaybeT1_21_138 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_21_139, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_22), v_23)
})
}))
_ = functorMaybeT1_21_138
// TAST (Let): __local_var_22_140 -> gopurs_runtime.Value
__local_var_22_140 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_23_141 -> *Constructor_Control_Bind_Bind
Bind1_23_141 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_141
// TAST (Let): Applicative0_24_142 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_142 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_142
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_141.V1), v_25, gopurs_runtime.Func(func(v1_27 gopurs_runtime.Value) gopurs_runtime.Value {
var __t143 gopurs_runtime.Value
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr == nil) {
__t143 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_142.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_143
} else {

}
}
{
if (v1_27.Type == 9 && v1_27.IntVal == 930809136 && v1_27.UnsafePtr != nil) {
__t143 = gopurs_runtime.Apply(f_26, (*Constructor_Data_Maybe_Just)(v1_27.UnsafePtr).V0)
goto end_branch_143
} else {

}
}
{
__t143 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_143:
return __t143
}))
})
}))
}))
_ = __local_var_22_140
// TAST (Let): Bind1_23_144 -> *Constructor_Control_Bind_Bind
Bind1_23_144 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_140, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_23_144
// TAST (Let): Applicative0_24_145 -> *Constructor_Control_Applicative_Applicative
Applicative0_24_145 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_22_140, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_24_145
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_21_138
}), gopurs_runtime.Func(func(f_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_144.V1), f_25, gopurs_runtime.Func(func(f_prime_27 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_23_144.V1), a_26, gopurs_runtime.Func(func(a_prime_28 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_24_145.V1), gopurs_runtime.Apply(f_prime_27, a_prime_28))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_20_146, x_21)
}))
}), gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_20_148 -> *Constructor_Control_Bind_Bind
Bind1_20_148 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_148
// TAST (Let): Applicative0_21_149 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_149 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_149
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad1_7_7)
}), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_148.V1), v_22, gopurs_runtime.Func(func(v1_24 gopurs_runtime.Value) gopurs_runtime.Value {
var __t150 gopurs_runtime.Value
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr == nil) {
__t150 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_149.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_150
} else {

}
}
{
if (v1_24.Type == 9 && v1_24.IntVal == 930809136 && v1_24.UnsafePtr != nil) {
__t150 = gopurs_runtime.Apply(f_23, (*Constructor_Data_Maybe_Just)(v1_24.UnsafePtr).V0)
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
}))
}))
_ = __local_var_19_137
// TAST (Let): Bind1_20_151 -> *Constructor_Control_Bind_Bind
Bind1_20_151 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_137, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_20_151
// TAST (Let): Applicative0_21_152 -> *Constructor_Control_Applicative_Applicative
Applicative0_21_152 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_19_137, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_21_152
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_19 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_18_135
}), gopurs_runtime.Func(func(f_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_151.V1), f_22, gopurs_runtime.Func(func(f_prime_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_20_151.V1), a_23, gopurs_runtime.Func(func(a_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_21_152.V1), gopurs_runtime.Apply(f_prime_24, a_prime_25))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_133.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t153 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t153 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_134.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_153
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t153 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
_ = __local_var_14_86
// TAST (Let): Bind1_15_154 -> *Constructor_Control_Bind_Bind
Bind1_15_154 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_86, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_154
// TAST (Let): Applicative0_16_155 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_155 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_86, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_155
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_13_84
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_154.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_154.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_155.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_82.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t156 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t156 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_83.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_156
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t156 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
}))
}))
_ = monadMaybeT1_9_9
// TAST (Let): Bind1_10_159 -> *Constructor_Control_Bind_Bind
Bind1_10_159 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_159
// TAST (Let): Applicative0_11_160 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_160 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_160
// TAST (Let): __local_var_10_158 -> gopurs_runtime.Value
__local_var_10_158 := gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_159.V1), a_12, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_160.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_13})})
}))
})
_ = __local_var_10_158
// TAST (Let): __local_var_10_157 -> gopurs_runtime.Value
__local_var_10_157 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_158, x_11)
})
_ = __local_var_10_157
// TAST (Let): monadTellMaybeT1_7_6 -> gopurs_runtime.Value
monadTellMaybeT1_7_6 := gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_9_9
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_8_8
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_157, gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "tell"), x_11))
}))
_ = monadTellMaybeT1_7_6
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellMaybeT1_7_6
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_6_5
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), v_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_161 -> gopurs_runtime.Value
__local_var_10_161 := (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1
_ = __local_var_10_161
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(r_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, r_11, __local_var_10_161})}
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just]((*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0))})))})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), v_8, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t162 *Constructor_Data_Tuple_Tuple
{
if (a_9.Type == 9 && a_9.IntVal == 930809136 && a_9.UnsafePtr == nil) {
__t162 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_10
})}
goto end_branch_162
} else {

}
}
{
if (a_9.Type == 9 && a_9.IntVal == 930809136 && a_9.UnsafePtr != nil) {
__t162 = &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(a_9.UnsafePtr).V0.UnsafePtr).V0})}, (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_Maybe_Just)(a_9.UnsafePtr).V0.UnsafePtr).V1}
goto end_branch_162
} else {

}
}
{
__t162 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_162:
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t162)})
})))
}))
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
// TAST (Let): monadMaybeT1_2_1 -> gopurs_runtime.Value
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_20 -> gopurs_runtime.Value
__local_var_4_20 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_20
// TAST (Let): __local_var_4_19 -> gopurs_runtime.Value
__local_var_4_19 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_20, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_19
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
// TAST (Let): functorMaybeT1_5_3 -> gopurs_runtime.Value
functorMaybeT1_5_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_6), v_7)
})
}))
_ = functorMaybeT1_5_3
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_6 -> *Constructor_Control_Bind_Bind
Bind1_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_6
// TAST (Let): Applicative0_8_7 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_7
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_9
// TAST (Let): functorMaybeT1_10_8 -> gopurs_runtime.Value
functorMaybeT1_10_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_8
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_11 -> *Constructor_Control_Bind_Bind
Bind1_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_11
// TAST (Let): Applicative0_13_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_12
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_11.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t13 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_12.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_13
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t13 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
_ = __local_var_11_10
// TAST (Let): Bind1_12_14 -> *Constructor_Control_Bind_Bind
Bind1_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_14
// TAST (Let): Applicative0_13_15 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_15
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_8
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_14.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_14.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_15.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_6.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t16 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_7.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_16
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t16 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
_ = __local_var_6_5
// TAST (Let): Bind1_7_17 -> *Constructor_Control_Bind_Bind
Bind1_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_17
// TAST (Let): Applicative0_8_18 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_18 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_5, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_18
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_5_3
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_17.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_17.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_18.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_19, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_21 -> *Constructor_Control_Bind_Bind
Bind1_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_21
// TAST (Let): Applicative0_5_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_22
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_24 -> gopurs_runtime.Value
__local_var_7_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_24
// TAST (Let): functorMaybeT1_7_23 -> gopurs_runtime.Value
functorMaybeT1_7_23 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_24, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_8), v_9)
})
}))
_ = functorMaybeT1_7_23
// TAST (Let): __local_var_8_25 -> gopurs_runtime.Value
__local_var_8_25 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_35 -> gopurs_runtime.Value
__local_var_9_35 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_35
// TAST (Let): __local_var_9_34 -> gopurs_runtime.Value
__local_var_9_34 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_35, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_34
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_27 -> gopurs_runtime.Value
__local_var_10_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_27
// TAST (Let): functorMaybeT1_10_26 -> gopurs_runtime.Value
functorMaybeT1_10_26 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_27, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_26
// TAST (Let): __local_var_11_28 -> gopurs_runtime.Value
__local_var_11_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_29 -> *Constructor_Control_Bind_Bind
Bind1_12_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_29
// TAST (Let): Applicative0_13_30 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_30
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_29.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t31 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_30.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_31
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t31 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_11_28
// TAST (Let): Bind1_12_32 -> *Constructor_Control_Bind_Bind
Bind1_12_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_32
// TAST (Let): Applicative0_13_33 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_33 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_33
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_26
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_32.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_32.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_33.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_34, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_36 -> *Constructor_Control_Bind_Bind
Bind1_9_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_36
// TAST (Let): Applicative0_10_37 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_2, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_37
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_2_2)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_36.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t38 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t38 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_37.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_38
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t38 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
}))
}))
_ = __local_var_8_25
// TAST (Let): Bind1_9_39 -> *Constructor_Control_Bind_Bind
Bind1_9_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_25, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_39
// TAST (Let): Applicative0_10_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_25, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_40
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_7_23
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_39.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_39.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_40.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_21.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t41 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_22.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_41
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t41 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Applicative0_4_42 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_4_42
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_1_0.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_3), gopurs_runtime.Func(func(a_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_42.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_5})})
}))
}))
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
// TAST (Let): monadMaybeT1_3_3 -> gopurs_runtime.Value
monadMaybeT1_3_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_76 -> gopurs_runtime.Value
__local_var_5_76 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_76
// TAST (Let): __local_var_5_75 -> gopurs_runtime.Value
__local_var_5_75 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_76, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_6})})
})
_ = __local_var_5_75
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_6
// TAST (Let): functorMaybeT1_6_5 -> gopurs_runtime.Value
functorMaybeT1_6_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_6, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_7), v_8)
})
}))
_ = functorMaybeT1_6_5
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_25 -> gopurs_runtime.Value
__local_var_8_25 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_25
// TAST (Let): __local_var_8_24 -> gopurs_runtime.Value
__local_var_8_24 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_25, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_24
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_9 -> gopurs_runtime.Value
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_9
// TAST (Let): functorMaybeT1_9_8 -> gopurs_runtime.Value
functorMaybeT1_9_8 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_9, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_10), v_11)
})
}))
_ = functorMaybeT1_9_8
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_11 -> *Constructor_Control_Bind_Bind
Bind1_11_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_11
// TAST (Let): Applicative0_12_12 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_12
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_14 -> gopurs_runtime.Value
__local_var_14_14 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_14
// TAST (Let): functorMaybeT1_14_13 -> gopurs_runtime.Value
functorMaybeT1_14_13 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_14, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_15), v_16)
})
}))
_ = functorMaybeT1_14_13
// TAST (Let): __local_var_15_15 -> gopurs_runtime.Value
__local_var_15_15 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_16 -> *Constructor_Control_Bind_Bind
Bind1_16_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_16
// TAST (Let): Applicative0_17_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_17
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_16.V1), v_18, gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t18 gopurs_runtime.Value
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr == nil) {
__t18 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_17.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_18
} else {

}
}
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr != nil) {
__t18 = gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_20.UnsafePtr).V0)
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
_ = __local_var_15_15
// TAST (Let): Bind1_16_19 -> *Constructor_Control_Bind_Bind
Bind1_16_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_15, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_19
// TAST (Let): Applicative0_17_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_15, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_20
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_14_13
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_19.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_19.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_20.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_11.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t21 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t21 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_12.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_21
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t21 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
}))
}))
_ = __local_var_10_10
// TAST (Let): Bind1_11_22 -> *Constructor_Control_Bind_Bind
Bind1_11_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_10, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_22
// TAST (Let): Applicative0_12_23 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_23 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_10, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_23
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_9_8
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_22.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_22.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_23.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_24, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_26 -> *Constructor_Control_Bind_Bind
Bind1_8_26 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_26
// TAST (Let): Applicative0_9_27 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_27
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_29 -> gopurs_runtime.Value
__local_var_11_29 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_29
// TAST (Let): functorMaybeT1_11_28 -> gopurs_runtime.Value
functorMaybeT1_11_28 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_29, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_12), v_13)
})
}))
_ = functorMaybeT1_11_28
// TAST (Let): __local_var_12_30 -> gopurs_runtime.Value
__local_var_12_30 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_48 -> gopurs_runtime.Value
__local_var_13_48 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_48
// TAST (Let): __local_var_13_47 -> gopurs_runtime.Value
__local_var_13_47 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_48, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_14})})
})
_ = __local_var_13_47
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_32 -> gopurs_runtime.Value
__local_var_14_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_32
// TAST (Let): functorMaybeT1_14_31 -> gopurs_runtime.Value
functorMaybeT1_14_31 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_32, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_15), v_16)
})
}))
_ = functorMaybeT1_14_31
// TAST (Let): __local_var_15_33 -> gopurs_runtime.Value
__local_var_15_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_34 -> *Constructor_Control_Bind_Bind
Bind1_16_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_34
// TAST (Let): Applicative0_17_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_35
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_37 -> gopurs_runtime.Value
__local_var_19_37 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_37
// TAST (Let): functorMaybeT1_19_36 -> gopurs_runtime.Value
functorMaybeT1_19_36 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_37, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_20), v_21)
})
}))
_ = functorMaybeT1_19_36
// TAST (Let): __local_var_20_38 -> gopurs_runtime.Value
__local_var_20_38 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_39 -> *Constructor_Control_Bind_Bind
Bind1_21_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_39
// TAST (Let): Applicative0_22_40 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_40
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_39.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t41 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_40.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_41
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t41 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
_ = __local_var_20_38
// TAST (Let): Bind1_21_42 -> *Constructor_Control_Bind_Bind
Bind1_21_42 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_38, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_42
// TAST (Let): Applicative0_22_43 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_38, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_43
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_19_36
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_42.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_42.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_43.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_34.V1), v_18, gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr == nil) {
__t44 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_35.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_44
} else {

}
}
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr != nil) {
__t44 = gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_20.UnsafePtr).V0)
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
_ = __local_var_15_33
// TAST (Let): Bind1_16_45 -> *Constructor_Control_Bind_Bind
Bind1_16_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_45
// TAST (Let): Applicative0_17_46 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_46 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_46
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_14_31
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_45.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_45.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_46.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_47, x_14)
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_49 -> *Constructor_Control_Bind_Bind
Bind1_13_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_49
// TAST (Let): Applicative0_14_50 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_50 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_50
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_52 -> gopurs_runtime.Value
__local_var_16_52 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_52
// TAST (Let): functorMaybeT1_16_51 -> gopurs_runtime.Value
functorMaybeT1_16_51 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_52, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_17), v_18)
})
}))
_ = functorMaybeT1_16_51
// TAST (Let): __local_var_17_53 -> gopurs_runtime.Value
__local_var_17_53 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_63 -> gopurs_runtime.Value
__local_var_18_63 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_63
// TAST (Let): __local_var_18_62 -> gopurs_runtime.Value
__local_var_18_62 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_63, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_19})})
})
_ = __local_var_18_62
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_55 -> gopurs_runtime.Value
__local_var_19_55 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_55
// TAST (Let): functorMaybeT1_19_54 -> gopurs_runtime.Value
functorMaybeT1_19_54 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_55, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_20), v_21)
})
}))
_ = functorMaybeT1_19_54
// TAST (Let): __local_var_20_56 -> gopurs_runtime.Value
__local_var_20_56 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_57 -> *Constructor_Control_Bind_Bind
Bind1_21_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_57
// TAST (Let): Applicative0_22_58 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_58 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_58
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_57.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t59 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t59 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_58.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_59
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t59 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
goto end_branch_59
} else {

}
}
{
__t59 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_59:
return __t59
}))
})
}))
}))
_ = __local_var_20_56
// TAST (Let): Bind1_21_60 -> *Constructor_Control_Bind_Bind
Bind1_21_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_56, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_60
// TAST (Let): Applicative0_22_61 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_61 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_56, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_61
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_19_54
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_60.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_60.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_61.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_62, x_19)
}))
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_64 -> *Constructor_Control_Bind_Bind
Bind1_18_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_64
// TAST (Let): Applicative0_19_65 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_65
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_64.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t66 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t66 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_65.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_66
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t66 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
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
_ = __local_var_17_53
// TAST (Let): Bind1_18_67 -> *Constructor_Control_Bind_Bind
Bind1_18_67 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_53, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_67
// TAST (Let): Applicative0_19_68 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_68 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_53, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_68
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_16_51
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_67.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_67.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_68.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_49.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t69 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t69 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_50.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_69
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t69 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
_ = __local_var_12_30
// TAST (Let): Bind1_13_70 -> *Constructor_Control_Bind_Bind
Bind1_13_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_30, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_70
// TAST (Let): Applicative0_14_71 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_30, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_71
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_11_28
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_70.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_70.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_71.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_26.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t72 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t72 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_27.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_72
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t72 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
_ = __local_var_7_7
// TAST (Let): Bind1_8_73 -> *Constructor_Control_Bind_Bind
Bind1_8_73 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_73
// TAST (Let): Applicative0_9_74 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_74 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_7, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_74
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_6_5
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_73.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_73.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_74.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_75, x_6)
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_77 -> *Constructor_Control_Bind_Bind
Bind1_5_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_77
// TAST (Let): Applicative0_6_78 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_78
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_80 -> gopurs_runtime.Value
__local_var_8_80 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_80
// TAST (Let): functorMaybeT1_8_79 -> gopurs_runtime.Value
functorMaybeT1_8_79 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_80, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_9), v_10)
})
}))
_ = functorMaybeT1_8_79
// TAST (Let): __local_var_9_81 -> gopurs_runtime.Value
__local_var_9_81 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_127 -> gopurs_runtime.Value
__local_var_10_127 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_127
// TAST (Let): __local_var_10_126 -> gopurs_runtime.Value
__local_var_10_126 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_127, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_11})})
})
_ = __local_var_10_126
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_83 -> gopurs_runtime.Value
__local_var_11_83 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_83
// TAST (Let): functorMaybeT1_11_82 -> gopurs_runtime.Value
functorMaybeT1_11_82 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_83, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_12), v_13)
})
}))
_ = functorMaybeT1_11_82
// TAST (Let): __local_var_12_84 -> gopurs_runtime.Value
__local_var_12_84 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_102 -> gopurs_runtime.Value
__local_var_13_102 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_13_102
// TAST (Let): __local_var_13_101 -> gopurs_runtime.Value
__local_var_13_101 := gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_102, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_14})})
})
_ = __local_var_13_101
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_14_86 -> gopurs_runtime.Value
__local_var_14_86 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_14_86
// TAST (Let): functorMaybeT1_14_85 -> gopurs_runtime.Value
functorMaybeT1_14_85 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_14_86, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_15), v_16)
})
}))
_ = functorMaybeT1_14_85
// TAST (Let): __local_var_15_87 -> gopurs_runtime.Value
__local_var_15_87 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_16_88 -> *Constructor_Control_Bind_Bind
Bind1_16_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_88
// TAST (Let): Applicative0_17_89 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_89 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_89
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_91 -> gopurs_runtime.Value
__local_var_19_91 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_91
// TAST (Let): functorMaybeT1_19_90 -> gopurs_runtime.Value
functorMaybeT1_19_90 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_91, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_20), v_21)
})
}))
_ = functorMaybeT1_19_90
// TAST (Let): __local_var_20_92 -> gopurs_runtime.Value
__local_var_20_92 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_93 -> *Constructor_Control_Bind_Bind
Bind1_21_93 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_93
// TAST (Let): Applicative0_22_94 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_94
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_93.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t95 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t95 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_94.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_95
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t95 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_20_92
// TAST (Let): Bind1_21_96 -> *Constructor_Control_Bind_Bind
Bind1_21_96 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_92, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_96
// TAST (Let): Applicative0_22_97 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_97 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_92, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_97
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_19_90
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_96.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_96.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_97.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_88.V1), v_18, gopurs_runtime.Func(func(v1_20 gopurs_runtime.Value) gopurs_runtime.Value {
var __t98 gopurs_runtime.Value
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr == nil) {
__t98 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_89.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_98
} else {

}
}
{
if (v1_20.Type == 9 && v1_20.IntVal == 930809136 && v1_20.UnsafePtr != nil) {
__t98 = gopurs_runtime.Apply(f_19, (*Constructor_Data_Maybe_Just)(v1_20.UnsafePtr).V0)
goto end_branch_98
} else {

}
}
{
__t98 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_98:
return __t98
}))
})
}))
}))
_ = __local_var_15_87
// TAST (Let): Bind1_16_99 -> *Constructor_Control_Bind_Bind
Bind1_16_99 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_87, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_16_99
// TAST (Let): Applicative0_17_100 -> *Constructor_Control_Applicative_Applicative
Applicative0_17_100 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_15_87, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_17_100
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_14_85
}), gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_99.V1), f_18, gopurs_runtime.Func(func(f_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_16_99.V1), a_19, gopurs_runtime.Func(func(a_prime_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_17_100.V1), gopurs_runtime.Apply(f_prime_20, a_prime_21))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_13_101, x_14)
}))
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_103 -> *Constructor_Control_Bind_Bind
Bind1_13_103 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_103
// TAST (Let): Applicative0_14_104 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_104
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_106 -> gopurs_runtime.Value
__local_var_16_106 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_106
// TAST (Let): functorMaybeT1_16_105 -> gopurs_runtime.Value
functorMaybeT1_16_105 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_106, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_17), v_18)
})
}))
_ = functorMaybeT1_16_105
// TAST (Let): __local_var_17_107 -> gopurs_runtime.Value
__local_var_17_107 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_18_117 -> gopurs_runtime.Value
__local_var_18_117 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_18_117
// TAST (Let): __local_var_18_116 -> gopurs_runtime.Value
__local_var_18_116 := gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_117, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_19})})
})
_ = __local_var_18_116
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_19_109 -> gopurs_runtime.Value
__local_var_19_109 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_19_109
// TAST (Let): functorMaybeT1_19_108 -> gopurs_runtime.Value
functorMaybeT1_19_108 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_19_109, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_20), v_21)
})
}))
_ = functorMaybeT1_19_108
// TAST (Let): __local_var_20_110 -> gopurs_runtime.Value
__local_var_20_110 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_21_111 -> *Constructor_Control_Bind_Bind
Bind1_21_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_111
// TAST (Let): Applicative0_22_112 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_112 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_112
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_111.V1), v_23, gopurs_runtime.Func(func(v1_25 gopurs_runtime.Value) gopurs_runtime.Value {
var __t113 gopurs_runtime.Value
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr == nil) {
__t113 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_112.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_113
} else {

}
}
{
if (v1_25.Type == 9 && v1_25.IntVal == 930809136 && v1_25.UnsafePtr != nil) {
__t113 = gopurs_runtime.Apply(f_24, (*Constructor_Data_Maybe_Just)(v1_25.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_20_110
// TAST (Let): Bind1_21_114 -> *Constructor_Control_Bind_Bind
Bind1_21_114 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_110, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_21_114
// TAST (Let): Applicative0_22_115 -> *Constructor_Control_Applicative_Applicative
Applicative0_22_115 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_20_110, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_22_115
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_19_108
}), gopurs_runtime.Func(func(f_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_24 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_114.V1), f_23, gopurs_runtime.Func(func(f_prime_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_21_114.V1), a_24, gopurs_runtime.Func(func(a_prime_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_22_115.V1), gopurs_runtime.Apply(f_prime_25, a_prime_26))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_18_116, x_19)
}))
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_118 -> *Constructor_Control_Bind_Bind
Bind1_18_118 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_118
// TAST (Let): Applicative0_19_119 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_119 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_119
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_118.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t120 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t120 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_119.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_120
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t120 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
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
_ = __local_var_17_107
// TAST (Let): Bind1_18_121 -> *Constructor_Control_Bind_Bind
Bind1_18_121 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_107, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_121
// TAST (Let): Applicative0_19_122 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_122 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_107, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_122
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_16_105
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_121.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_121.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_122.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_103.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t123 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t123 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_104.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_123
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t123 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
_ = __local_var_12_84
// TAST (Let): Bind1_13_124 -> *Constructor_Control_Bind_Bind
Bind1_13_124 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_84, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_124
// TAST (Let): Applicative0_14_125 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_125 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_84, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_125
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_11_82
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_124.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_124.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_125.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_126, x_11)
}))
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_128 -> *Constructor_Control_Bind_Bind
Bind1_10_128 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_128
// TAST (Let): Applicative0_11_129 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_129 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_129
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_131 -> gopurs_runtime.Value
__local_var_13_131 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_131
// TAST (Let): functorMaybeT1_13_130 -> gopurs_runtime.Value
functorMaybeT1_13_130 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_131, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_14), v_15)
})
}))
_ = functorMaybeT1_13_130
// TAST (Let): __local_var_14_132 -> gopurs_runtime.Value
__local_var_14_132 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_142 -> gopurs_runtime.Value
__local_var_15_142 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_15_142
// TAST (Let): __local_var_15_141 -> gopurs_runtime.Value
__local_var_15_141 := gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_142, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_16})})
})
_ = __local_var_15_141
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_134 -> gopurs_runtime.Value
__local_var_16_134 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_134
// TAST (Let): functorMaybeT1_16_133 -> gopurs_runtime.Value
functorMaybeT1_16_133 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_134, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_17), v_18)
})
}))
_ = functorMaybeT1_16_133
// TAST (Let): __local_var_17_135 -> gopurs_runtime.Value
__local_var_17_135 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_136 -> *Constructor_Control_Bind_Bind
Bind1_18_136 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_136
// TAST (Let): Applicative0_19_137 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_137 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_137
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_136.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t138 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t138 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_137.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_138
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t138 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
goto end_branch_138
} else {

}
}
{
__t138 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_138:
return __t138
}))
})
}))
}))
_ = __local_var_17_135
// TAST (Let): Bind1_18_139 -> *Constructor_Control_Bind_Bind
Bind1_18_139 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_135, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_139
// TAST (Let): Applicative0_19_140 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_140 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_135, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_140
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_16_133
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_139.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_139.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_140.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_141, x_16)
}))
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_143 -> *Constructor_Control_Bind_Bind
Bind1_15_143 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_143
// TAST (Let): Applicative0_16_144 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_144 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_144
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(__local_var_3_4)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_143.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t145 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t145 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_144.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_145
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t145 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
goto end_branch_145
} else {

}
}
{
__t145 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_145:
return __t145
}))
})
}))
}))
_ = __local_var_14_132
// TAST (Let): Bind1_15_146 -> *Constructor_Control_Bind_Bind
Bind1_15_146 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_132, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_146
// TAST (Let): Applicative0_16_147 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_147 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_132, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_147
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_13_130
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_146.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_146.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_147.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_128.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t148 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t148 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_129.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_148
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t148 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
goto end_branch_148
} else {

}
}
{
__t148 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_148:
return __t148
}))
})
}))
}))
_ = __local_var_9_81
// TAST (Let): Bind1_10_149 -> *Constructor_Control_Bind_Bind
Bind1_10_149 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_81, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_149
// TAST (Let): Applicative0_11_150 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_150 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_81, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_150
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_8_79
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_149.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_149.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_150.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_77.V1), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t151 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr == nil) {
__t151 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_78.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_151
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t151 = gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)
goto end_branch_151
} else {

}
}
{
__t151 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_151:
return __t151
}))
})
}))
}))
_ = monadMaybeT1_3_3
// TAST (Let): monadThrowMaybeT1_1_0 -> gopurs_runtime.Value
monadThrowMaybeT1_1_0 := gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_3_3
}), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Applicative0_5_152 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_152 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_2_2.V0), gopurs_runtime.Value{}))
_ = Applicative0_5_152
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(Monad0_2_2.V1), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "throwError"), e_4), gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_152.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_6})})
}))
}))
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
// TAST (Let): __local_var_3_19 -> gopurs_runtime.Value
__local_var_3_19 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_3_19
// TAST (Let): __local_var_3_18 -> gopurs_runtime.Value
__local_var_3_18 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_19, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_4})})
})
_ = __local_var_3_18
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_3
// TAST (Let): functorMaybeT1_4_2 -> gopurs_runtime.Value
functorMaybeT1_4_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_5), v_6)
})
}))
_ = functorMaybeT1_4_2
// TAST (Let): __local_var_5_4 -> gopurs_runtime.Value
__local_var_5_4 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_6_5 -> *Constructor_Control_Bind_Bind
Bind1_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_5
// TAST (Let): Applicative0_7_6 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_6
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_8
// TAST (Let): functorMaybeT1_9_7 -> gopurs_runtime.Value
functorMaybeT1_9_7 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_8, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_10), v_11)
})
}))
_ = functorMaybeT1_9_7
// TAST (Let): __local_var_10_9 -> gopurs_runtime.Value
__local_var_10_9 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_10 -> *Constructor_Control_Bind_Bind
Bind1_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_10
// TAST (Let): Applicative0_12_11 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_11
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_10.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t12 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_11.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_12
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t12 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_10_9
// TAST (Let): Bind1_11_13 -> *Constructor_Control_Bind_Bind
Bind1_11_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_13
// TAST (Let): Applicative0_12_14 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_14 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_9, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_14
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_9_7
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_13.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_13.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_14.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_5.V1), v_8, gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr == nil) {
__t15 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_6.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_15
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 930809136 && v1_10.UnsafePtr != nil) {
__t15 = gopurs_runtime.Apply(f_9, (*Constructor_Data_Maybe_Just)(v1_10.UnsafePtr).V0)
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
_ = __local_var_5_4
// TAST (Let): Bind1_6_16 -> *Constructor_Control_Bind_Bind
Bind1_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_6_16
// TAST (Let): Applicative0_7_17 -> *Constructor_Control_Applicative_Applicative
Applicative0_7_17 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_4, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_7_17
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_4_2
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_16.V1), f_8, gopurs_runtime.Func(func(f_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_6_16.V1), a_9, gopurs_runtime.Func(func(a_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_7_17.V1), gopurs_runtime.Apply(f_prime_10, a_prime_11))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_18, x_4)
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_3_20 -> *Constructor_Control_Bind_Bind
Bind1_3_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_20
// TAST (Let): Applicative0_4_21 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_21
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_23 -> gopurs_runtime.Value
__local_var_6_23 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_23
// TAST (Let): functorMaybeT1_6_22 -> gopurs_runtime.Value
functorMaybeT1_6_22 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_23, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_7), v_8)
})
}))
_ = functorMaybeT1_6_22
// TAST (Let): __local_var_7_24 -> gopurs_runtime.Value
__local_var_7_24 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_34 -> gopurs_runtime.Value
__local_var_8_34 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_8_34
// TAST (Let): __local_var_8_33 -> gopurs_runtime.Value
__local_var_8_33 := gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_34, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_9})})
})
_ = __local_var_8_33
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_26 -> gopurs_runtime.Value
__local_var_9_26 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_9_26
// TAST (Let): functorMaybeT1_9_25 -> gopurs_runtime.Value
functorMaybeT1_9_25 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_9_26, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_10), v_11)
})
}))
_ = functorMaybeT1_9_25
// TAST (Let): __local_var_10_27 -> gopurs_runtime.Value
__local_var_10_27 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_11_28 -> *Constructor_Control_Bind_Bind
Bind1_11_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_28
// TAST (Let): Applicative0_12_29 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_29
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_28.V1), v_13, gopurs_runtime.Func(func(v1_15 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr == nil) {
__t30 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_29.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_30
} else {

}
}
{
if (v1_15.Type == 9 && v1_15.IntVal == 930809136 && v1_15.UnsafePtr != nil) {
__t30 = gopurs_runtime.Apply(f_14, (*Constructor_Data_Maybe_Just)(v1_15.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_10_27
// TAST (Let): Bind1_11_31 -> *Constructor_Control_Bind_Bind
Bind1_11_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_27, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_11_31
// TAST (Let): Applicative0_12_32 -> *Constructor_Control_Applicative_Applicative
Applicative0_12_32 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_10_27, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_12_32
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_9_25
}), gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_31.V1), f_13, gopurs_runtime.Func(func(f_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_11_31.V1), a_14, gopurs_runtime.Func(func(a_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_12_32.V1), gopurs_runtime.Apply(f_prime_15, a_prime_16))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_33, x_9)
}))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_35 -> *Constructor_Control_Bind_Bind
Bind1_8_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_35
// TAST (Let): Applicative0_9_36 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_36 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_36
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(Monad0_1_0)
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_35.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t37 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t37 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_36.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_37
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t37 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
}))
}))
_ = __local_var_7_24
// TAST (Let): Bind1_8_38 -> *Constructor_Control_Bind_Bind
Bind1_8_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_24, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_38
// TAST (Let): Applicative0_9_39 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_39 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_24, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_39
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_6_22
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_38.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_39.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_20.V1), v_5, gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t40 gopurs_runtime.Value
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr == nil) {
__t40 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_21.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_40
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 930809136 && v1_7.UnsafePtr != nil) {
__t40 = gopurs_runtime.Apply(f_6, (*Constructor_Data_Maybe_Just)(v1_7.UnsafePtr).V0)
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
_ = monadMaybeT1_2_1
// TAST (Let): Bind1_3_43 -> *Constructor_Control_Bind_Bind
Bind1_3_43 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_43
// TAST (Let): Applicative0_4_44 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_44
// TAST (Let): __local_var_3_42 -> gopurs_runtime.Value
__local_var_3_42 := gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_43.V1), a_5, gopurs_runtime.Func(func(a_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_44.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_6})})
}))
})
_ = __local_var_3_42
// TAST (Let): __local_var_3_41 -> gopurs_runtime.Value
__local_var_3_41 := gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_42, x_4)
})
_ = __local_var_3_41
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_41, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_Control_Monad_Maybe_Trans_monoidMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_18 -> gopurs_runtime.Value
__local_var_1_18 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_18
// TAST (Let): __local_var_1_17 -> gopurs_runtime.Value
__local_var_1_17 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_18, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_2})})
})
_ = __local_var_1_17
// TAST (Let): applicativeMaybeT1_1_0 -> *Constructor_Control_Applicative_Applicative
applicativeMaybeT1_1_0 := &Constructor_Control_Applicative_Applicative{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorMaybeT1_2_1 -> gopurs_runtime.Value
functorMaybeT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_3), v_4)
})
}))
_ = functorMaybeT1_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_4 -> *Constructor_Control_Bind_Bind
Bind1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_4
// TAST (Let): Applicative0_5_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_5
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): functorMaybeT1_7_6 -> gopurs_runtime.Value
functorMaybeT1_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_8), v_9)
})
}))
_ = functorMaybeT1_7_6
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_9 -> *Constructor_Control_Bind_Bind
Bind1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_9
// TAST (Let): Applicative0_10_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_9.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_10.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_11
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t11 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
_ = __local_var_8_8
// TAST (Let): Bind1_9_12 -> *Constructor_Control_Bind_Bind
Bind1_9_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_12
// TAST (Let): Applicative0_10_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_7_6
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_12.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_12.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_13.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_5.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_14
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t14 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
_ = __local_var_3_3
// TAST (Let): Bind1_4_15 -> *Constructor_Control_Bind_Bind
Bind1_4_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_15
// TAST (Let): Applicative0_5_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_2_1
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_15.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_15.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_16.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_17, x_2)
})}
_ = applicativeMaybeT1_1_0
// TAST (Let): __local_var_2_21 -> gopurs_runtime.Value
__local_var_2_21 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_21
// TAST (Let): functorMaybeT1_2_20 -> gopurs_runtime.Value
functorMaybeT1_2_20 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_21, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_3), v_4)
})
}))
_ = functorMaybeT1_2_20
// TAST (Let): __local_var_3_22 -> gopurs_runtime.Value
__local_var_3_22 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_68 -> gopurs_runtime.Value
__local_var_4_68 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_4_68
// TAST (Let): __local_var_4_67 -> gopurs_runtime.Value
__local_var_4_67 := gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_68, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_5})})
})
_ = __local_var_4_67
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_24 -> gopurs_runtime.Value
__local_var_5_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_24
// TAST (Let): functorMaybeT1_5_23 -> gopurs_runtime.Value
functorMaybeT1_5_23 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_24, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_6), v_7)
})
}))
_ = functorMaybeT1_5_23
// TAST (Let): __local_var_6_25 -> gopurs_runtime.Value
__local_var_6_25 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_43 -> gopurs_runtime.Value
__local_var_7_43 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_43
// TAST (Let): __local_var_7_42 -> gopurs_runtime.Value
__local_var_7_42 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_43, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_8})})
})
_ = __local_var_7_42
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_27 -> gopurs_runtime.Value
__local_var_8_27 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_27
// TAST (Let): functorMaybeT1_8_26 -> gopurs_runtime.Value
functorMaybeT1_8_26 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_27, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_9), v_10)
})
}))
_ = functorMaybeT1_8_26
// TAST (Let): __local_var_9_28 -> gopurs_runtime.Value
__local_var_9_28 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_29 -> *Constructor_Control_Bind_Bind
Bind1_10_29 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_29
// TAST (Let): Applicative0_11_30 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_30
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_32 -> gopurs_runtime.Value
__local_var_13_32 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_32
// TAST (Let): functorMaybeT1_13_31 -> gopurs_runtime.Value
functorMaybeT1_13_31 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_32, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_14), v_15)
})
}))
_ = functorMaybeT1_13_31
// TAST (Let): __local_var_14_33 -> gopurs_runtime.Value
__local_var_14_33 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_34 -> *Constructor_Control_Bind_Bind
Bind1_15_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_34
// TAST (Let): Applicative0_16_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_35
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_34.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t36 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_35.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_36
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t36 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
_ = __local_var_14_33
// TAST (Let): Bind1_15_37 -> *Constructor_Control_Bind_Bind
Bind1_15_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_33, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_37
// TAST (Let): Applicative0_16_38 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_33, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_38
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_13_31
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_37.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_37.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_38.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_29.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t39 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t39 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_30.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_39
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t39 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
}))
}))
_ = __local_var_9_28
// TAST (Let): Bind1_10_40 -> *Constructor_Control_Bind_Bind
Bind1_10_40 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_28, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_40
// TAST (Let): Applicative0_11_41 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_41 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_28, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_41
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_8_26
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_40.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_40.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_41.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_42, x_8)
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_44 -> *Constructor_Control_Bind_Bind
Bind1_7_44 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_44
// TAST (Let): Applicative0_8_45 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_45 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_45
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_47 -> gopurs_runtime.Value
__local_var_10_47 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_47
// TAST (Let): functorMaybeT1_10_46 -> gopurs_runtime.Value
functorMaybeT1_10_46 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_47, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_46
// TAST (Let): __local_var_11_48 -> gopurs_runtime.Value
__local_var_11_48 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_58 -> gopurs_runtime.Value
__local_var_12_58 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_12_58
// TAST (Let): __local_var_12_57 -> gopurs_runtime.Value
__local_var_12_57 := gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_58, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_13})})
})
_ = __local_var_12_57
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_50 -> gopurs_runtime.Value
__local_var_13_50 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_50
// TAST (Let): functorMaybeT1_13_49 -> gopurs_runtime.Value
functorMaybeT1_13_49 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_50, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_14), v_15)
})
}))
_ = functorMaybeT1_13_49
// TAST (Let): __local_var_14_51 -> gopurs_runtime.Value
__local_var_14_51 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_52 -> *Constructor_Control_Bind_Bind
Bind1_15_52 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_52
// TAST (Let): Applicative0_16_53 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_53
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_52.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t54 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t54 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_53.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_54
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t54 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
})
}))
}))
_ = __local_var_14_51
// TAST (Let): Bind1_15_55 -> *Constructor_Control_Bind_Bind
Bind1_15_55 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_51, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_55
// TAST (Let): Applicative0_16_56 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_51, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_56
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_13_49
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_55.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_55.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_56.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_12_57, x_13)
}))
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_59 -> *Constructor_Control_Bind_Bind
Bind1_12_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_59
// TAST (Let): Applicative0_13_60 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_60
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_59.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t61 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t61 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_60.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_61
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t61 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
}))
}))
_ = __local_var_11_48
// TAST (Let): Bind1_12_62 -> *Constructor_Control_Bind_Bind
Bind1_12_62 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_48, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_62
// TAST (Let): Applicative0_13_63 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_48, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_63
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_46
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_62.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_62.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_63.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_44.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t64 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t64 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_45.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_64
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t64 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
_ = __local_var_6_25
// TAST (Let): Bind1_7_65 -> *Constructor_Control_Bind_Bind
Bind1_7_65 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_25, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_65
// TAST (Let): Applicative0_8_66 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_66 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_25, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_66
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_5_23
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_65.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_65.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_66.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_67, x_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_69 -> *Constructor_Control_Bind_Bind
Bind1_4_69 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_69
// TAST (Let): Applicative0_5_70 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_70 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_70
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_72 -> gopurs_runtime.Value
__local_var_7_72 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_72
// TAST (Let): functorMaybeT1_7_71 -> gopurs_runtime.Value
functorMaybeT1_7_71 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_72, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_8), v_9)
})
}))
_ = functorMaybeT1_7_71
// TAST (Let): __local_var_8_73 -> gopurs_runtime.Value
__local_var_8_73 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_83 -> gopurs_runtime.Value
__local_var_9_83 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_9_83
// TAST (Let): __local_var_9_82 -> gopurs_runtime.Value
__local_var_9_82 := gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_83, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_10})})
})
_ = __local_var_9_82
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_75 -> gopurs_runtime.Value
__local_var_10_75 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_10_75
// TAST (Let): functorMaybeT1_10_74 -> gopurs_runtime.Value
functorMaybeT1_10_74 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_10_75, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_11), v_12)
})
}))
_ = functorMaybeT1_10_74
// TAST (Let): __local_var_11_76 -> gopurs_runtime.Value
__local_var_11_76 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_12_77 -> *Constructor_Control_Bind_Bind
Bind1_12_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_77
// TAST (Let): Applicative0_13_78 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_78 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_78
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_77.V1), v_14, gopurs_runtime.Func(func(v1_16 gopurs_runtime.Value) gopurs_runtime.Value {
var __t79 gopurs_runtime.Value
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr == nil) {
__t79 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_78.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_79
} else {

}
}
{
if (v1_16.Type == 9 && v1_16.IntVal == 930809136 && v1_16.UnsafePtr != nil) {
__t79 = gopurs_runtime.Apply(f_15, (*Constructor_Data_Maybe_Just)(v1_16.UnsafePtr).V0)
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
_ = __local_var_11_76
// TAST (Let): Bind1_12_80 -> *Constructor_Control_Bind_Bind
Bind1_12_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_76, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_12_80
// TAST (Let): Applicative0_13_81 -> *Constructor_Control_Applicative_Applicative
Applicative0_13_81 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_11_76, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_13_81
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_10_74
}), gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_80.V1), f_14, gopurs_runtime.Func(func(f_prime_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_12_80.V1), a_15, gopurs_runtime.Func(func(a_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_13_81.V1), gopurs_runtime.Apply(f_prime_16, a_prime_17))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_9_82, x_10)
}))
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_84 -> *Constructor_Control_Bind_Bind
Bind1_9_84 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_84
// TAST (Let): Applicative0_10_85 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_85 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_85
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_84.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t86 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t86 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_85.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_86
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t86 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
goto end_branch_86
} else {

}
}
{
__t86 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_86:
return __t86
}))
})
}))
}))
_ = __local_var_8_73
// TAST (Let): Bind1_9_87 -> *Constructor_Control_Bind_Bind
Bind1_9_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_73, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_87
// TAST (Let): Applicative0_10_88 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_88 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_73, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_88
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_7_71
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_87.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_87.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_88.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_69.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t89 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t89 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_70.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_89
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t89 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
_ = __local_var_3_22
// TAST (Let): Bind1_4_90 -> *Constructor_Control_Bind_Bind
Bind1_4_90 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_22, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_90
// TAST (Let): Applicative0_5_91 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_91 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_22, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_91
// TAST (Let): applyMaybeT1_2_19 -> *Constructor_Control_Apply_Apply
applyMaybeT1_2_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_2_20
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_90.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_90.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_91.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
})))
_ = applyMaybeT1_2_19
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_93 -> *Constructor_Data_Functor_Functor
Functor0_4_93 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.Box(applyMaybeT1_2_19.V0), gopurs_runtime.Value{}))
_ = Functor0_4_93
// TAST (Let): __local_var_5_94 -> gopurs_runtime.Value
__local_var_5_94 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_5_94
// TAST (Let): semigroupMaybeT2_4_92 -> gopurs_runtime.Value
semigroupMaybeT2_4_92 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(applyMaybeT1_2_19.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_93.V0), __local_var_5_94, a_6), b_7)
})
}))
_ = semigroupMaybeT2_4_92
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMaybeT2_4_92
}), gopurs_runtime.Apply(gopurs_runtime.Box(applicativeMaybeT1_1_0.V1), gopurs_runtime.RecordGet(dictMonoid_3, "mempty")))
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
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](m_6))})
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
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorMaybeT1_3_3 -> gopurs_runtime.Value
functorMaybeT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_4), v_5)
})
}))
_ = functorMaybeT1_3_3
// TAST (Let): altMaybeT1_1_0 -> gopurs_runtime.Value
altMaybeT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_3_3
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), v_4, gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 930809136 && m_6.UnsafePtr == nil) {
__t5 = v1_5
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](m_6))})
}
end_branch_5:
return __t5
}))
})
}))
_ = altMaybeT1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMaybeT1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}))
}

func Call_Control_Monad_Maybe_Trans_alternativeMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_18 -> gopurs_runtime.Value
__local_var_1_18 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_1_18
// TAST (Let): __local_var_1_17 -> gopurs_runtime.Value
__local_var_1_17 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_18, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_2})})
})
_ = __local_var_1_17
// TAST (Let): applicativeMaybeT1_1_0 -> gopurs_runtime.Value
applicativeMaybeT1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorMaybeT1_2_1 -> gopurs_runtime.Value
functorMaybeT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_3), v_4)
})
}))
_ = functorMaybeT1_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_4_4 -> *Constructor_Control_Bind_Bind
Bind1_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_4
// TAST (Let): Applicative0_5_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_5
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_7 -> gopurs_runtime.Value
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_7_7
// TAST (Let): functorMaybeT1_7_6 -> gopurs_runtime.Value
functorMaybeT1_7_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_7, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_8), v_9)
})
}))
_ = functorMaybeT1_7_6
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_9_9 -> *Constructor_Control_Bind_Bind
Bind1_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_9
// TAST (Let): Applicative0_10_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_9.V1), v_11, gopurs_runtime.Func(func(v1_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr == nil) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_10.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_11
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 930809136 && v1_13.UnsafePtr != nil) {
__t11 = gopurs_runtime.Apply(f_12, (*Constructor_Data_Maybe_Just)(v1_13.UnsafePtr).V0)
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
_ = __local_var_8_8
// TAST (Let): Bind1_9_12 -> *Constructor_Control_Bind_Bind
Bind1_9_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_9_12
// TAST (Let): Applicative0_10_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_10_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_8_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_10_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_7_6
}), gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_12.V1), f_11, gopurs_runtime.Func(func(f_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_9_12.V1), a_12, gopurs_runtime.Func(func(a_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_10_13.V1), gopurs_runtime.Apply(f_prime_13, a_prime_14))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_4.V1), v_6, gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr == nil) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_5.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_14
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 930809136 && v1_8.UnsafePtr != nil) {
__t14 = gopurs_runtime.Apply(f_7, (*Constructor_Data_Maybe_Just)(v1_8.UnsafePtr).V0)
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
_ = __local_var_3_3
// TAST (Let): Bind1_4_15 -> *Constructor_Control_Bind_Bind
Bind1_4_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_4_15
// TAST (Let): Applicative0_5_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_5_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_5_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_2_1
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_15.V1), f_6, gopurs_runtime.Func(func(f_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_4_15.V1), a_7, gopurs_runtime.Func(func(a_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_5_16.V1), gopurs_runtime.Apply(f_prime_8, a_prime_9))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_17, x_2)
}))
_ = applicativeMaybeT1_1_0
// TAST (Let): Bind1_2_21 -> *Constructor_Control_Bind_Bind
Bind1_2_21 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_21
// TAST (Let): Applicative0_3_22 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_22 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_22
// TAST (Let): __local_var_4_24 -> gopurs_runtime.Value
__local_var_4_24 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_24
// TAST (Let): functorMaybeT1_4_23 -> gopurs_runtime.Value
functorMaybeT1_4_23 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_24, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_5), v_6)
})
}))
_ = functorMaybeT1_4_23
// TAST (Let): altMaybeT1_2_20 -> gopurs_runtime.Value
altMaybeT1_2_20 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_4_23
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_21.V1), v_5, gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t25 gopurs_runtime.Value
{
if (m_7.Type == 9 && m_7.IntVal == 930809136 && m_7.UnsafePtr == nil) {
__t25 = v1_6
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_22.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](m_7))})
}
end_branch_25:
return __t25
}))
})
}))
_ = altMaybeT1_2_20
// TAST (Let): plusMaybeT1_2_19 -> gopurs_runtime.Value
plusMaybeT1_2_19 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return altMaybeT1_2_20
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}))
_ = plusMaybeT1_2_19
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeMaybeT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusMaybeT1_2_19
}))
}

func Call_Control_Monad_Maybe_Trans_monadPlusMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): monadMaybeT1_1_0 -> gopurs_runtime.Value
monadMaybeT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_18 -> gopurs_runtime.Value
__local_var_2_18 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_2_18
// TAST (Let): __local_var_2_17 -> gopurs_runtime.Value
__local_var_2_17 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_18, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_3})})
})
_ = __local_var_2_17
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
// TAST (Let): functorMaybeT1_3_1 -> gopurs_runtime.Value
functorMaybeT1_3_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_4), v_5)
})
}))
_ = functorMaybeT1_3_1
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_4 -> *Constructor_Control_Bind_Bind
Bind1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_4
// TAST (Let): Applicative0_6_5 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_5
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_7 -> gopurs_runtime.Value
__local_var_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_7
// TAST (Let): functorMaybeT1_8_6 -> gopurs_runtime.Value
functorMaybeT1_8_6 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_7, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_9), v_10)
})
}))
_ = functorMaybeT1_8_6
// TAST (Let): __local_var_9_8 -> gopurs_runtime.Value
__local_var_9_8 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_9 -> *Constructor_Control_Bind_Bind
Bind1_10_9 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_9
// TAST (Let): Applicative0_11_10 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_10
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_9.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t11 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t11 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_10.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_11
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t11 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
_ = __local_var_9_8
// TAST (Let): Bind1_10_12 -> *Constructor_Control_Bind_Bind
Bind1_10_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_8, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_12
// TAST (Let): Applicative0_11_13 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_13 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_8, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_13
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_8_6
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_12.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_12.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_13.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_4.V1), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr == nil) {
__t14 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_5.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_14
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t14 = gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)
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
_ = __local_var_4_3
// TAST (Let): Bind1_5_15 -> *Constructor_Control_Bind_Bind
Bind1_5_15 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_15
// TAST (Let): Applicative0_6_16 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_16
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_3_1
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_15.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_15.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_16.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_17, x_3)
}))
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_2_19 -> *Constructor_Control_Bind_Bind
Bind1_2_19 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_19
// TAST (Let): Applicative0_3_20 -> *Constructor_Control_Applicative_Applicative
Applicative0_3_20 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_3_20
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_22 -> gopurs_runtime.Value
__local_var_5_22 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_22
// TAST (Let): functorMaybeT1_5_21 -> gopurs_runtime.Value
functorMaybeT1_5_21 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_22, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_6), v_7)
})
}))
_ = functorMaybeT1_5_21
// TAST (Let): __local_var_6_23 -> gopurs_runtime.Value
__local_var_6_23 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_33 -> gopurs_runtime.Value
__local_var_7_33 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_7_33
// TAST (Let): __local_var_7_32 -> gopurs_runtime.Value
__local_var_7_32 := gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_33, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_8})})
})
_ = __local_var_7_32
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_25 -> gopurs_runtime.Value
__local_var_8_25 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_25
// TAST (Let): functorMaybeT1_8_24 -> gopurs_runtime.Value
functorMaybeT1_8_24 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_25, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_9), v_10)
})
}))
_ = functorMaybeT1_8_24
// TAST (Let): __local_var_9_26 -> gopurs_runtime.Value
__local_var_9_26 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_27 -> *Constructor_Control_Bind_Bind
Bind1_10_27 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_27
// TAST (Let): Applicative0_11_28 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_28 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_28
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_27.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t29 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t29 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_28.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_29
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t29 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
_ = __local_var_9_26
// TAST (Let): Bind1_10_30 -> *Constructor_Control_Bind_Bind
Bind1_10_30 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_26, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_30
// TAST (Let): Applicative0_11_31 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_31 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_26, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_31
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_8_24
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_30.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_30.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_31.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_32, x_8)
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_7_34 -> *Constructor_Control_Bind_Bind
Bind1_7_34 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_34
// TAST (Let): Applicative0_8_35 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_35 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_35
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_34.V1), v_9, gopurs_runtime.Func(func(v1_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t36 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr == nil) {
__t36 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_35.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_36
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 930809136 && v1_11.UnsafePtr != nil) {
__t36 = gopurs_runtime.Apply(f_10, (*Constructor_Data_Maybe_Just)(v1_11.UnsafePtr).V0)
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
_ = __local_var_6_23
// TAST (Let): Bind1_7_37 -> *Constructor_Control_Bind_Bind
Bind1_7_37 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_23, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_7_37
// TAST (Let): Applicative0_8_38 -> *Constructor_Control_Applicative_Applicative
Applicative0_8_38 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_6_23, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_8_38
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_5_21
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_37.V1), f_9, gopurs_runtime.Func(func(f_prime_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_7_37.V1), a_10, gopurs_runtime.Func(func(a_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_8_38.V1), gopurs_runtime.Apply(f_prime_11, a_prime_12))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_19.V1), v_4, gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t39 gopurs_runtime.Value
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr == nil) {
__t39 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_3_20.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_39
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 930809136 && v1_6.UnsafePtr != nil) {
__t39 = gopurs_runtime.Apply(f_5, (*Constructor_Data_Maybe_Just)(v1_6.UnsafePtr).V0)
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
}))
}))
_ = monadMaybeT1_1_0
// TAST (Let): __local_var_2_113 -> gopurs_runtime.Value
__local_var_2_113 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_2_113
// TAST (Let): __local_var_2_112 -> gopurs_runtime.Value
__local_var_2_112 := gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_113, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_3})})
})
_ = __local_var_2_112
// TAST (Let): applicativeMaybeT1_2_41 -> gopurs_runtime.Value
applicativeMaybeT1_2_41 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_43 -> gopurs_runtime.Value
__local_var_3_43 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_43
// TAST (Let): functorMaybeT1_3_42 -> gopurs_runtime.Value
functorMaybeT1_3_42 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_43, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_4), v_5)
})
}))
_ = functorMaybeT1_3_42
// TAST (Let): __local_var_4_44 -> gopurs_runtime.Value
__local_var_4_44 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_62 -> gopurs_runtime.Value
__local_var_5_62 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_5_62
// TAST (Let): __local_var_5_61 -> gopurs_runtime.Value
__local_var_5_61 := gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_62, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_6})})
})
_ = __local_var_5_61
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_46 -> gopurs_runtime.Value
__local_var_6_46 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_46
// TAST (Let): functorMaybeT1_6_45 -> gopurs_runtime.Value
functorMaybeT1_6_45 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_46, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_7), v_8)
})
}))
_ = functorMaybeT1_6_45
// TAST (Let): __local_var_7_47 -> gopurs_runtime.Value
__local_var_7_47 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_8_48 -> *Constructor_Control_Bind_Bind
Bind1_8_48 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_48
// TAST (Let): Applicative0_9_49 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_49 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_49
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_51 -> gopurs_runtime.Value
__local_var_11_51 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_51
// TAST (Let): functorMaybeT1_11_50 -> gopurs_runtime.Value
functorMaybeT1_11_50 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_51, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_12), v_13)
})
}))
_ = functorMaybeT1_11_50
// TAST (Let): __local_var_12_52 -> gopurs_runtime.Value
__local_var_12_52 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_53 -> *Constructor_Control_Bind_Bind
Bind1_13_53 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_53
// TAST (Let): Applicative0_14_54 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_54 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_54
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_53.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t55 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t55 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_54.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_55
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t55 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
_ = __local_var_12_52
// TAST (Let): Bind1_13_56 -> *Constructor_Control_Bind_Bind
Bind1_13_56 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_52, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_56
// TAST (Let): Applicative0_14_57 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_57 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_52, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_57
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_11_50
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_56.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_56.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_57.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_48.V1), v_10, gopurs_runtime.Func(func(v1_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t58 gopurs_runtime.Value
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr == nil) {
__t58 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_49.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_58
} else {

}
}
{
if (v1_12.Type == 9 && v1_12.IntVal == 930809136 && v1_12.UnsafePtr != nil) {
__t58 = gopurs_runtime.Apply(f_11, (*Constructor_Data_Maybe_Just)(v1_12.UnsafePtr).V0)
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
}))
}))
_ = __local_var_7_47
// TAST (Let): Bind1_8_59 -> *Constructor_Control_Bind_Bind
Bind1_8_59 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_47, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_8_59
// TAST (Let): Applicative0_9_60 -> *Constructor_Control_Applicative_Applicative
Applicative0_9_60 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_7_47, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_9_60
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_6_45
}), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_59.V1), f_10, gopurs_runtime.Func(func(f_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_8_59.V1), a_11, gopurs_runtime.Func(func(a_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_9_60.V1), gopurs_runtime.Apply(f_prime_12, a_prime_13))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_61, x_6)
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_5_63 -> *Constructor_Control_Bind_Bind
Bind1_5_63 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_63
// TAST (Let): Applicative0_6_64 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_64 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_64
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_66 -> gopurs_runtime.Value
__local_var_8_66 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_8_66
// TAST (Let): functorMaybeT1_8_65 -> gopurs_runtime.Value
functorMaybeT1_8_65 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_8_66, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_9), v_10)
})
}))
_ = functorMaybeT1_8_65
// TAST (Let): __local_var_9_67 -> gopurs_runtime.Value
__local_var_9_67 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_85 -> gopurs_runtime.Value
__local_var_10_85 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_10_85
// TAST (Let): __local_var_10_84 -> gopurs_runtime.Value
__local_var_10_84 := gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_85, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_11})})
})
_ = __local_var_10_84
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_69 -> gopurs_runtime.Value
__local_var_11_69 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_11_69
// TAST (Let): functorMaybeT1_11_68 -> gopurs_runtime.Value
functorMaybeT1_11_68 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_11_69, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_12), v_13)
})
}))
_ = functorMaybeT1_11_68
// TAST (Let): __local_var_12_70 -> gopurs_runtime.Value
__local_var_12_70 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_13_71 -> *Constructor_Control_Bind_Bind
Bind1_13_71 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_71
// TAST (Let): Applicative0_14_72 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_72 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_72
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_74 -> gopurs_runtime.Value
__local_var_16_74 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_74
// TAST (Let): functorMaybeT1_16_73 -> gopurs_runtime.Value
functorMaybeT1_16_73 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_74, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_17), v_18)
})
}))
_ = functorMaybeT1_16_73
// TAST (Let): __local_var_17_75 -> gopurs_runtime.Value
__local_var_17_75 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_76 -> *Constructor_Control_Bind_Bind
Bind1_18_76 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_76
// TAST (Let): Applicative0_19_77 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_77 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_77
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_76.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t78 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t78 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_77.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_78
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t78 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
goto end_branch_78
} else {

}
}
{
__t78 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_78:
return __t78
}))
})
}))
}))
_ = __local_var_17_75
// TAST (Let): Bind1_18_79 -> *Constructor_Control_Bind_Bind
Bind1_18_79 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_75, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_79
// TAST (Let): Applicative0_19_80 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_80 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_75, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_80
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_16_73
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_79.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_79.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_80.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_71.V1), v_15, gopurs_runtime.Func(func(v1_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t81 gopurs_runtime.Value
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr == nil) {
__t81 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_72.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_81
} else {

}
}
{
if (v1_17.Type == 9 && v1_17.IntVal == 930809136 && v1_17.UnsafePtr != nil) {
__t81 = gopurs_runtime.Apply(f_16, (*Constructor_Data_Maybe_Just)(v1_17.UnsafePtr).V0)
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
}))
}))
_ = __local_var_12_70
// TAST (Let): Bind1_13_82 -> *Constructor_Control_Bind_Bind
Bind1_13_82 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_70, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_13_82
// TAST (Let): Applicative0_14_83 -> *Constructor_Control_Applicative_Applicative
Applicative0_14_83 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_12_70, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_14_83
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_11_68
}), gopurs_runtime.Func(func(f_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_82.V1), f_15, gopurs_runtime.Func(func(f_prime_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_13_82.V1), a_16, gopurs_runtime.Func(func(a_prime_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_14_83.V1), gopurs_runtime.Apply(f_prime_17, a_prime_18))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_10_84, x_11)
}))
}), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_10_86 -> *Constructor_Control_Bind_Bind
Bind1_10_86 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_86
// TAST (Let): Applicative0_11_87 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_87 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_87
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_89 -> gopurs_runtime.Value
__local_var_13_89 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_13_89
// TAST (Let): functorMaybeT1_13_88 -> gopurs_runtime.Value
functorMaybeT1_13_88 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_13_89, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_14), v_15)
})
}))
_ = functorMaybeT1_13_88
// TAST (Let): __local_var_14_90 -> gopurs_runtime.Value
__local_var_14_90 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_15_100 -> gopurs_runtime.Value
__local_var_15_100 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = __local_var_15_100
// TAST (Let): __local_var_15_99 -> gopurs_runtime.Value
__local_var_15_99 := gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_100, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, x_16})})
})
_ = __local_var_15_99
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_16_92 -> gopurs_runtime.Value
__local_var_16_92 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_16_92
// TAST (Let): functorMaybeT1_16_91 -> gopurs_runtime.Value
functorMaybeT1_16_91 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_16_92, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_17), v_18)
})
}))
_ = functorMaybeT1_16_91
// TAST (Let): __local_var_17_93 -> gopurs_runtime.Value
__local_var_17_93 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applicativeMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_18_94 -> *Constructor_Control_Bind_Bind
Bind1_18_94 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_94
// TAST (Let): Applicative0_19_95 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_95 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_95
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_94.V1), v_20, gopurs_runtime.Func(func(v1_22 gopurs_runtime.Value) gopurs_runtime.Value {
var __t96 gopurs_runtime.Value
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr == nil) {
__t96 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_95.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_96
} else {

}
}
{
if (v1_22.Type == 9 && v1_22.IntVal == 930809136 && v1_22.UnsafePtr != nil) {
__t96 = gopurs_runtime.Apply(f_21, (*Constructor_Data_Maybe_Just)(v1_22.UnsafePtr).V0)
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
_ = __local_var_17_93
// TAST (Let): Bind1_18_97 -> *Constructor_Control_Bind_Bind
Bind1_18_97 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_93, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_18_97
// TAST (Let): Applicative0_19_98 -> *Constructor_Control_Applicative_Applicative
Applicative0_19_98 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_17_93, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_19_98
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_16_91
}), gopurs_runtime.Func(func(f_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_97.V1), f_20, gopurs_runtime.Func(func(f_prime_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_18_97.V1), a_21, gopurs_runtime.Func(func(a_prime_23 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_19_98.V1), gopurs_runtime.Apply(f_prime_22, a_prime_23))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_16 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_15_99, x_16)
}))
}), gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Bind1_15_101 -> *Constructor_Control_Bind_Bind
Bind1_15_101 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_101
// TAST (Let): Applicative0_16_102 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_102 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_102
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Maybe_Trans_applyMaybeT(dictMonad_0)
}), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_101.V1), v_17, gopurs_runtime.Func(func(v1_19 gopurs_runtime.Value) gopurs_runtime.Value {
var __t103 gopurs_runtime.Value
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr == nil) {
__t103 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_102.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_103
} else {

}
}
{
if (v1_19.Type == 9 && v1_19.IntVal == 930809136 && v1_19.UnsafePtr != nil) {
__t103 = gopurs_runtime.Apply(f_18, (*Constructor_Data_Maybe_Just)(v1_19.UnsafePtr).V0)
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
}))
}))
_ = __local_var_14_90
// TAST (Let): Bind1_15_104 -> *Constructor_Control_Bind_Bind
Bind1_15_104 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_90, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_104
// TAST (Let): Applicative0_16_105 -> *Constructor_Control_Applicative_Applicative
Applicative0_16_105 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_14_90, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_16_105
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_13_88
}), gopurs_runtime.Func(func(f_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_104.V1), f_17, gopurs_runtime.Func(func(f_prime_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_104.V1), a_18, gopurs_runtime.Func(func(a_prime_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_16_105.V1), gopurs_runtime.Apply(f_prime_19, a_prime_20))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_86.V1), v_12, gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t106 gopurs_runtime.Value
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr == nil) {
__t106 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_87.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_106
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 930809136 && v1_14.UnsafePtr != nil) {
__t106 = gopurs_runtime.Apply(f_13, (*Constructor_Data_Maybe_Just)(v1_14.UnsafePtr).V0)
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
_ = __local_var_9_67
// TAST (Let): Bind1_10_107 -> *Constructor_Control_Bind_Bind
Bind1_10_107 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_67, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_10_107
// TAST (Let): Applicative0_11_108 -> *Constructor_Control_Applicative_Applicative
Applicative0_11_108 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_9_67, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_11_108
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_8_65
}), gopurs_runtime.Func(func(f_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_107.V1), f_12, gopurs_runtime.Func(func(f_prime_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_10_107.V1), a_13, gopurs_runtime.Func(func(a_prime_15 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_11_108.V1), gopurs_runtime.Apply(f_prime_14, a_prime_15))
}))
}))
})
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_63.V1), v_7, gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t109 gopurs_runtime.Value
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr == nil) {
__t109 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_64.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_109
} else {

}
}
{
if (v1_9.Type == 9 && v1_9.IntVal == 930809136 && v1_9.UnsafePtr != nil) {
__t109 = gopurs_runtime.Apply(f_8, (*Constructor_Data_Maybe_Just)(v1_9.UnsafePtr).V0)
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
_ = __local_var_4_44
// TAST (Let): Bind1_5_110 -> *Constructor_Control_Bind_Bind
Bind1_5_110 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_44, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_5_110
// TAST (Let): Applicative0_6_111 -> *Constructor_Control_Applicative_Applicative
Applicative0_6_111 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_44, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_6_111
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_3_42
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_110.V1), f_7, gopurs_runtime.Func(func(f_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_5_110.V1), a_8, gopurs_runtime.Func(func(a_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_6_111.V1), gopurs_runtime.Apply(f_prime_9, a_prime_10))
}))
}))
})
}))
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_112, x_3)
}))
_ = applicativeMaybeT1_2_41
// TAST (Let): Bind1_3_116 -> *Constructor_Control_Bind_Bind
Bind1_3_116 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_116
// TAST (Let): Applicative0_4_117 -> *Constructor_Control_Applicative_Applicative
Applicative0_4_117 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_4_117
// TAST (Let): __local_var_5_119 -> gopurs_runtime.Value
__local_var_5_119 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_119
// TAST (Let): functorMaybeT1_5_118 -> gopurs_runtime.Value
functorMaybeT1_5_118 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_119, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), f_6), v_7)
})
}))
_ = functorMaybeT1_5_118
// TAST (Let): altMaybeT1_3_115 -> gopurs_runtime.Value
altMaybeT1_3_115 := gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_5_118
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_116.V1), v_6, gopurs_runtime.Func(func(m_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t120 gopurs_runtime.Value
{
if (m_8.Type == 9 && m_8.IntVal == 930809136 && m_8.UnsafePtr == nil) {
__t120 = v1_7
goto end_branch_120
} else {

}
}
{
__t120 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_4_117.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](m_8))})
}
end_branch_120:
return __t120
}))
})
}))
_ = altMaybeT1_3_115
// TAST (Let): plusMaybeT1_3_114 -> gopurs_runtime.Value
plusMaybeT1_3_114 := gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altMaybeT1_3_115
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}))
_ = plusMaybeT1_3_114
// TAST (Let): alternativeMaybeT1_2_40 -> gopurs_runtime.Value
alternativeMaybeT1_2_40 := gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeMaybeT1_2_41
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return plusMaybeT1_3_114
}))
_ = alternativeMaybeT1_2_40
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeMaybeT1_2_40
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



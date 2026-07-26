package Control_Monad_Maybe_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_MaybeT gopurs_runtime.Value
var once_MaybeT sync.Once
func Get_MaybeT() gopurs_runtime.Value {
	once_MaybeT.Do(func() {
		cache_MaybeT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_MaybeT(x_0_box)
})
	})
	return cache_MaybeT
}

var cache_runMaybeT gopurs_runtime.Value
var once_runMaybeT sync.Once
func Get_runMaybeT() gopurs_runtime.Value {
	once_runMaybeT.Do(func() {
		cache_runMaybeT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runMaybeT(v_0_box)
})
	})
	return cache_runMaybeT
}

var cache_newtypeMaybeT gopurs_runtime.Value
var once_newtypeMaybeT sync.Once
func Get_newtypeMaybeT() gopurs_runtime.Value {
	once_newtypeMaybeT.Do(func() {
		cache_newtypeMaybeT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeMaybeT
}

var cache_monadTransMaybeT gopurs_runtime.Value
var once_monadTransMaybeT sync.Once
func Get_monadTransMaybeT() gopurs_runtime.Value {
	once_monadTransMaybeT.Do(func() {
		cache_monadTransMaybeT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func2(func(dictMonad_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), x_1, gopurs_runtime.Func(func(a_prime_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{a_prime_2})})
}))
}))
	})
	return cache_monadTransMaybeT
}

var cache_mapMaybeT gopurs_runtime.Value
var once_mapMaybeT sync.Once
func Get_mapMaybeT() gopurs_runtime.Value {
	once_mapMaybeT.Do(func() {
		cache_mapMaybeT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapMaybeT(f_0_box, v_1_box)
})
	})
	return cache_mapMaybeT
}

var cache_functorMaybeT gopurs_runtime.Value
var once_functorMaybeT sync.Once
func Get_functorMaybeT() gopurs_runtime.Value {
	once_functorMaybeT.Do(func() {
		cache_functorMaybeT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorMaybeT(dictFunctor_0_box)
})
	})
	return cache_functorMaybeT
}

var cache_monadMaybeT gopurs_runtime.Value
var once_monadMaybeT sync.Once
func Get_monadMaybeT() gopurs_runtime.Value {
	once_monadMaybeT.Do(func() {
		cache_monadMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadMaybeT(dictMonad_0_box)
})
	})
	return cache_monadMaybeT
}

var cache_bindMaybeT gopurs_runtime.Value
var once_bindMaybeT sync.Once
func Get_bindMaybeT() gopurs_runtime.Value {
	once_bindMaybeT.Do(func() {
		cache_bindMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindMaybeT(dictMonad_0_box)
})
	})
	return cache_bindMaybeT
}

var cache_applyMaybeT gopurs_runtime.Value
var once_applyMaybeT sync.Once
func Get_applyMaybeT() gopurs_runtime.Value {
	once_applyMaybeT.Do(func() {
		cache_applyMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyMaybeT(dictMonad_0_box)
})
	})
	return cache_applyMaybeT
}

var cache_applicativeMaybeT gopurs_runtime.Value
var once_applicativeMaybeT sync.Once
func Get_applicativeMaybeT() gopurs_runtime.Value {
	once_applicativeMaybeT.Do(func() {
		cache_applicativeMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeMaybeT(dictMonad_0_box)
})
	})
	return cache_applicativeMaybeT
}

var cache_semigroupMaybeT gopurs_runtime.Value
var once_semigroupMaybeT sync.Once
func Get_semigroupMaybeT() gopurs_runtime.Value {
	once_semigroupMaybeT.Do(func() {
		cache_semigroupMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupMaybeT(dictMonad_0_box)
})
	})
	return cache_semigroupMaybeT
}

var cache_monadAskMaybeT gopurs_runtime.Value
var once_monadAskMaybeT sync.Once
func Get_monadAskMaybeT() gopurs_runtime.Value {
	once_monadAskMaybeT.Do(func() {
		cache_monadAskMaybeT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAskMaybeT(dictMonadAsk_0_box)
})
	})
	return cache_monadAskMaybeT
}

var cache_monadReaderMaybeT gopurs_runtime.Value
var once_monadReaderMaybeT sync.Once
func Get_monadReaderMaybeT() gopurs_runtime.Value {
	once_monadReaderMaybeT.Do(func() {
		cache_monadReaderMaybeT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderMaybeT(dictMonadReader_0_box)
})
	})
	return cache_monadReaderMaybeT
}

var cache_monadContMaybeT gopurs_runtime.Value
var once_monadContMaybeT sync.Once
func Get_monadContMaybeT() gopurs_runtime.Value {
	once_monadContMaybeT.Do(func() {
		cache_monadContMaybeT = gopurs_runtime.Func(func(dictMonadCont_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContMaybeT(dictMonadCont_0_box)
})
	})
	return cache_monadContMaybeT
}

var cache_monadEffectMaybe gopurs_runtime.Value
var once_monadEffectMaybe sync.Once
func Get_monadEffectMaybe() gopurs_runtime.Value {
	once_monadEffectMaybe.Do(func() {
		cache_monadEffectMaybe = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectMaybe(dictMonadEffect_0_box)
})
	})
	return cache_monadEffectMaybe
}

var cache_monadRecMaybeT gopurs_runtime.Value
var once_monadRecMaybeT sync.Once
func Get_monadRecMaybeT() gopurs_runtime.Value {
	once_monadRecMaybeT.Do(func() {
		cache_monadRecMaybeT = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadRecMaybeT(dictMonadRec_0_box)
})
	})
	return cache_monadRecMaybeT
}

var cache_monadStateMaybeT gopurs_runtime.Value
var once_monadStateMaybeT sync.Once
func Get_monadStateMaybeT() gopurs_runtime.Value {
	once_monadStateMaybeT.Do(func() {
		cache_monadStateMaybeT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateMaybeT(dictMonadState_0_box)
})
	})
	return cache_monadStateMaybeT
}

var cache_monadTellMaybeT gopurs_runtime.Value
var once_monadTellMaybeT sync.Once
func Get_monadTellMaybeT() gopurs_runtime.Value {
	once_monadTellMaybeT.Do(func() {
		cache_monadTellMaybeT = gopurs_runtime.Func(func(dictMonadTell_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadTellMaybeT(dictMonadTell_0_box)
})
	})
	return cache_monadTellMaybeT
}

var cache_monadWriterMaybeT gopurs_runtime.Value
var once_monadWriterMaybeT sync.Once
func Get_monadWriterMaybeT() gopurs_runtime.Value {
	once_monadWriterMaybeT.Do(func() {
		cache_monadWriterMaybeT = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadWriterMaybeT(dictMonadWriter_0_box)
})
	})
	return cache_monadWriterMaybeT
}

var cache_monadThrowMaybeT gopurs_runtime.Value
var once_monadThrowMaybeT sync.Once
func Get_monadThrowMaybeT() gopurs_runtime.Value {
	once_monadThrowMaybeT.Do(func() {
		cache_monadThrowMaybeT = gopurs_runtime.Func(func(dictMonadThrow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadThrowMaybeT(dictMonadThrow_0_box)
})
	})
	return cache_monadThrowMaybeT
}

var cache_monadErrorMaybeT gopurs_runtime.Value
var once_monadErrorMaybeT sync.Once
func Get_monadErrorMaybeT() gopurs_runtime.Value {
	once_monadErrorMaybeT.Do(func() {
		cache_monadErrorMaybeT = gopurs_runtime.Func(func(dictMonadError_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadErrorMaybeT(dictMonadError_0_box)
})
	})
	return cache_monadErrorMaybeT
}

var cache_monadSTMaybeT gopurs_runtime.Value
var once_monadSTMaybeT sync.Once
func Get_monadSTMaybeT() gopurs_runtime.Value {
	once_monadSTMaybeT.Do(func() {
		cache_monadSTMaybeT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTMaybeT(dictMonadST_0_box)
})
	})
	return cache_monadSTMaybeT
}

var cache_monoidMaybeT gopurs_runtime.Value
var once_monoidMaybeT sync.Once
func Get_monoidMaybeT() gopurs_runtime.Value {
	once_monoidMaybeT.Do(func() {
		cache_monoidMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidMaybeT(dictMonad_0_box)
})
	})
	return cache_monoidMaybeT
}

var cache_altMaybeT gopurs_runtime.Value
var once_altMaybeT sync.Once
func Get_altMaybeT() gopurs_runtime.Value {
	once_altMaybeT.Do(func() {
		cache_altMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_altMaybeT(dictMonad_0_box)
})
	})
	return cache_altMaybeT
}

var cache_plusMaybeT gopurs_runtime.Value
var once_plusMaybeT sync.Once
func Get_plusMaybeT() gopurs_runtime.Value {
	once_plusMaybeT.Do(func() {
		cache_plusMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_plusMaybeT(dictMonad_0_box)
})
	})
	return cache_plusMaybeT
}

var cache_alternativeMaybeT gopurs_runtime.Value
var once_alternativeMaybeT sync.Once
func Get_alternativeMaybeT() gopurs_runtime.Value {
	once_alternativeMaybeT.Do(func() {
		cache_alternativeMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeMaybeT(dictMonad_0_box)
})
	})
	return cache_alternativeMaybeT
}

var cache_monadPlusMaybeT gopurs_runtime.Value
var once_monadPlusMaybeT sync.Once
func Get_monadPlusMaybeT() gopurs_runtime.Value {
	once_monadPlusMaybeT.Do(func() {
		cache_monadPlusMaybeT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadPlusMaybeT(dictMonad_0_box)
})
	})
	return cache_monadPlusMaybeT
}

func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_MaybeT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_runMaybeT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_mapMaybeT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_functorMaybeT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), f_1), v_2)
}))
}

func Call_monadMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), dictMonad_0)
}))
}

func Call_bindMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyMaybeT(), dictMonad_0)
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), v_1, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 930809136 && v1_3.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(f_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_3.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}))
}))
}

func Call_applyMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorMaybeT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), f_2), v_3)
}))
_ = functorMaybeT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(Get_bindMaybeT(), dictMonad_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_2_1
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), f_4, gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), a_5, gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0), "pure"), gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
}))
}

func Call_applicativeMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyMaybeT(), dictMonad_0)
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{x_1})})
}))
}

func Call_semigroupMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(Get_applyMaybeT(), dictMonad_0)
_ = __local_var_1_0
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.RecordGet(dictSemigroup_2, "append")
_ = __local_var_3_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_3_1, a_4), b_5)
}))
})
}

func Call_monadAskMaybeT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_monadTransMaybeT(), "lift"), Monad0_1_0, gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")))
}

func Call_monadReaderMaybeT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
monadAskMaybeT1_1_0 := gopurs_runtime.Apply(Get_monadAskMaybeT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{}))
_ = monadAskMaybeT1_1_0
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskMaybeT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_2)
}))
}

func Call_monadContMaybeT(dictMonadCont_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadCont_0 gopurs_runtime.Value = dictMonadCont_0_loop
_ = dictMonadCont_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "Monad0"), gopurs_runtime.Value{})
_ = __local_var_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), __local_var_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), __local_var_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadCont_0, "callCC"), gopurs_runtime.Func(func(c_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{a_5})})
}))
}))
}))
}

func Call_monadEffectMaybe(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransMaybeT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_monadRecMaybeT(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "tailRecM", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply(f_3, a_4), gopurs_runtime.Func(func(m_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (m_prime_5.Type == 9 && m_prime_5.IntVal == 930809136 && m_prime_5.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}})}
goto end_branch_2
} else {

}
}
{
if (m_prime_5.Type == 9 && m_prime_5.IntVal == 930809136 && m_prime_5.UnsafePtr != nil) {
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(m_prime_5.UnsafePtr).V0
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 525585346) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Control_Monad_Rec_Class.Constructor_Loop[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(m_prime_5.UnsafePtr).V0.UnsafePtr).V0})}
goto end_branch_3
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(m_prime_5.UnsafePtr).V0
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 60402430) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{(*pkg_Control_Monad_Rec_Class.Constructor_Done[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(m_prime_5.UnsafePtr).V0.UnsafePtr).V0})}})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t2 = __t3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t2)
}))
}))
}))
}

func Call_monadStateMaybeT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
lift1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransMaybeT(), "lift"), Monad0_1_0)
_ = lift1_2_1
monadMaybeT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_3_2
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_3_2
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(lift1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), f_4))
}))
}

func Call_monadTellMaybeT(dictMonadTell_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadTell_0 gopurs_runtime.Value = dictMonadTell_0_loop
_ = dictMonadTell_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
Semigroup0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_2_1
monadMaybeT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad1_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad1_1_0)
}))
_ = monadMaybeT1_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransMaybeT(), "lift"), Monad1_1_0)
_ = __local_var_4_3
return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadTell_0, "tell"), x_5))
}))
}

func Call_monadWriterMaybeT(dictMonadWriter_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadWriter_0 gopurs_runtime.Value = dictMonadWriter_0_loop
_ = dictMonadWriter_0
MonadTell1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "MonadTell1"), gopurs_runtime.Value{})
_ = MonadTell1_1_0
Monad1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadTell1_1_0, "Monad1"), gopurs_runtime.Value{})
_ = Monad1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_3_2
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_2_1, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_4_3
Monoid0_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "Monoid0"), gopurs_runtime.Value{})
_ = Monoid0_5_4
monadTellMaybeT1_6_5 := gopurs_runtime.Apply(Get_monadTellMaybeT(), MonadTell1_1_0)
_ = monadTellMaybeT1_6_5
return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellMaybeT1_6_5
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_5_4
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "listen"), v_7), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1
_ = __local_var_9_6
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), gopurs_runtime.Func(func(r_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{r_10, __local_var_9_6})}
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0))
}))
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadWriter_0, "pass"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "bind"), v_7, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (a_8.Type == 9 && a_8.IntVal == 930809136 && a_8.UnsafePtr == nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, Get_identity()})}
goto end_branch_7
} else {

}
}
{
if (a_8.Type == 9 && a_8.IntVal == 930809136 && a_8.UnsafePtr != nil) {
__t7 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(a_8.UnsafePtr).V0.UnsafePtr).V0})}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(a_8.UnsafePtr).V0.UnsafePtr).V1})}
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "pure"), __t7)
})))
}))
}

func Call_monadThrowMaybeT(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
lift1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransMaybeT(), "lift"), Monad0_1_0)
_ = lift1_2_1
monadMaybeT1_3_2 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_3_2
return gopurs_runtime.RecordDict2("Monad0", "throwError", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_3_2
}), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(lift1_2_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_4))
}))
}

func Call_monadErrorMaybeT(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
monadThrowMaybeT1_1_0 := gopurs_runtime.Apply(Get_monadThrowMaybeT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowMaybeT1_1_0
return gopurs_runtime.RecordDict2("MonadThrow0", "catchError", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowMaybeT1_1_0
}), gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_3, a_4)
}))
}))
}

func Call_monadSTMaybeT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadMaybeT1_2_1 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), Monad0_1_0)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), Monad0_1_0)
}))
_ = monadMaybeT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransMaybeT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_monoidMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
semigroupMaybeT1_1_0 := gopurs_runtime.Apply(Get_semigroupMaybeT(), dictMonad_0)
_ = semigroupMaybeT1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupMaybeT2_3_1 := gopurs_runtime.Apply(semigroupMaybeT1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupMaybeT2_3_1
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupMaybeT2_3_1
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0), "pure"), gopurs_runtime.RecordGet(dictMonoid_2, "mempty")))
})
}

func Call_altMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorMaybeT1_3_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), f_3), v_4)
}))
_ = functorMaybeT1_3_2
return gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorMaybeT1_3_2
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_1_0, "bind"), v_4, gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (m_6.Type == 9 && m_6.IntVal == 930809136 && m_6.UnsafePtr == nil) {
__t3 = v1_5
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), m_6)
}
end_branch_3:
return __t3
}))
}))
}

func Call_plusMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
altMaybeT1_1_0 := gopurs_runtime.Apply(Get_altMaybeT(), dictMonad_0)
_ = altMaybeT1_1_0
return gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return altMaybeT1_1_0
}), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))
}

func Call_alternativeMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeMaybeT1_1_0 := gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0)
_ = applicativeMaybeT1_1_0
plusMaybeT1_2_1 := gopurs_runtime.Apply(Get_plusMaybeT(), dictMonad_0)
_ = plusMaybeT1_2_1
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeMaybeT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return plusMaybeT1_2_1
}))
}

func Call_monadPlusMaybeT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadMaybeT1_1_0 := gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeMaybeT(), dictMonad_0)
}), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindMaybeT(), dictMonad_0)
}))
_ = monadMaybeT1_1_0
alternativeMaybeT1_2_1 := gopurs_runtime.Apply(Get_alternativeMaybeT(), dictMonad_0)
_ = alternativeMaybeT1_2_1
return gopurs_runtime.RecordDict2("Alternative1", "Monad0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeMaybeT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadMaybeT1_1_0
}))
}



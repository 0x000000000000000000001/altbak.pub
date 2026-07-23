package Control_Monad_Cont_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var ContT gopurs_runtime.Value
var once_ContT sync.Once
func Get_ContT() gopurs_runtime.Value {
	once_ContT.Do(func() {
		ContT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return ContT
}

var withContT gopurs_runtime.Value
var once_withContT sync.Once
func Get_withContT() gopurs_runtime.Value {
	once_withContT.Do(func() {
		withContT = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, k_2))
})
	})
	return withContT
}

var runContT gopurs_runtime.Value
var once_runContT sync.Once
func Get_runContT() gopurs_runtime.Value {
	once_runContT.Do(func() {
		runContT = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, k_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, k_1)
})
	})
	return runContT
}

var newtypeContT gopurs_runtime.Value
var once_newtypeContT sync.Once
func Get_newtypeContT() gopurs_runtime.Value {
	once_newtypeContT.Do(func() {
		newtypeContT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeContT
}

var monadTransContT gopurs_runtime.Value
var once_monadTransContT sync.Once
func Get_monadTransContT() gopurs_runtime.Value {
	once_monadTransContT.Do(func() {
		monadTransContT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind")
}))
	})
	return monadTransContT
}

var mapContT gopurs_runtime.Value
var once_mapContT sync.Once
func Get_mapContT() gopurs_runtime.Value {
	once_mapContT.Do(func() {
		mapContT = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
})
	})
	return mapContT
}

var functorContT gopurs_runtime.Value
var once_functorContT sync.Once
func Get_functorContT() gopurs_runtime.Value {
	once_functorContT.Do(func() {
		functorContT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
}))
})
	})
	return functorContT
}

var applyContT gopurs_runtime.Value
var once_applyContT sync.Once
func Get_applyContT() gopurs_runtime.Value {
	once_applyContT.Do(func() {
		applyContT = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
}))
_ = functorContT1_1_0
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
}))
})
	})
	return applyContT
}

var bindContT gopurs_runtime.Value
var once_bindContT sync.Once
func Get_bindContT() gopurs_runtime.Value {
	once_bindContT.Do(func() {
		bindContT = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
}))
_ = functorContT1_1_0
applyContT1_2_1 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
}))
_ = applyContT1_2_1
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value, k_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_4, a_6, k_prime_5)
}))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_2_1
}))
})
	})
	return bindContT
}

var semigroupContT gopurs_runtime.Value
var once_semigroupContT sync.Once
func Get_semigroupContT() gopurs_runtime.Value {
	once_semigroupContT.Do(func() {
		semigroupContT = gopurs_runtime.Func2(func(dictApply_0 gopurs_runtime.Value, dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(a_2, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), a_5)
_ = __local_var_6_0
return gopurs_runtime.Apply(b_3, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(__local_var_6_0, a_7))
}))
}))
}))
})
	})
	return semigroupContT
}

var applicativeContT gopurs_runtime.Value
var once_applicativeContT sync.Once
func Get_applicativeContT() gopurs_runtime.Value {
	once_applicativeContT.Do(func() {
		applicativeContT = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
}))
_ = functorContT1_1_0
applyContT1_2_1 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
}))
_ = applyContT1_2_1
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, a_3)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_2_1
}))
})
	})
	return applicativeContT
}

var monadContT gopurs_runtime.Value
var once_monadContT sync.Once
func Get_monadContT() gopurs_runtime.Value {
	once_monadContT.Do(func() {
		monadContT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
}))
_ = functorContT1_1_0
applyContT1_2_2 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
}))
_ = applyContT1_2_2
applicativeContT1_2_1 := gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, a_3)
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_2_2
}))
_ = applicativeContT1_2_1
functorContT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(f_3, a_6))
}))
}))
_ = functorContT1_3_3
applyContT1_4_5 := gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func3(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_4, gopurs_runtime.Func(func(g_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_5, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_6, gopurs_runtime.Apply(g_7, a_8))
}))
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_3_3
}))
_ = applyContT1_4_5
bindContT1_4_4 := gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func3(func(v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value, k_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_6, a_8, k_prime_7)
}))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_4_5
}))
_ = bindContT1_4_4
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeContT1_2_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return bindContT1_4_4
}))
})
	})
	return monadContT
}

var monadAskContT gopurs_runtime.Value
var once_monadAskContT sync.Once
func Get_monadAskContT() gopurs_runtime.Value {
	once_monadAskContT.Do(func() {
		monadAskContT = gopurs_runtime.Func(func(dictMonadAsk_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := gopurs_runtime.Apply(Get_monadContT(), Monad0_1_0)
_ = monadContT1_2_1
return gopurs_runtime.RecordDict2("ask", "Monad0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}))
})
	})
	return monadAskContT
}

var monadReaderContT gopurs_runtime.Value
var once_monadReaderContT sync.Once
func Get_monadReaderContT() gopurs_runtime.Value {
	once_monadReaderContT.Do(func() {
		monadReaderContT = gopurs_runtime.Func(func(dictMonadReader_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadAsk0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})
_ = MonadAsk0_1_0
ask_2_1 := gopurs_runtime.RecordGet(MonadAsk0_1_0, "ask")
_ = ask_2_1
monadAskContT1_3_2 := gopurs_runtime.Apply(Get_monadAskContT(), MonadAsk0_1_0)
_ = monadAskContT1_3_2
return gopurs_runtime.RecordDict2("local", "MonadAsk0", gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadAsk0_1_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), ask_2_1, gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return r_7
}))
_ = __local_var_8_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_4, gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_3, gopurs_runtime.Apply(k_6, x_9))
})))
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskContT1_3_2
}))
})
	})
	return monadReaderContT
}

var monadContContT gopurs_runtime.Value
var once_monadContContT sync.Once
func Get_monadContContT() gopurs_runtime.Value {
	once_monadContContT.Do(func() {
		monadContContT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadContT1_1_0 := gopurs_runtime.Apply(Get_monadContT(), dictMonad_0)
_ = monadContT1_1_0
return gopurs_runtime.RecordDict2("callCC", "Monad0", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func2(func(a_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_4)
}), k_3)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_1_0
}))
})
	})
	return monadContContT
}

var monadEffectContT gopurs_runtime.Value
var once_monadEffectContT sync.Once
func Get_monadEffectContT() gopurs_runtime.Value {
	once_monadEffectContT.Do(func() {
		monadEffectContT = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := gopurs_runtime.Apply(Get_monadContT(), Monad0_1_0)
_ = monadContT1_2_1
__local_var_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind")
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("liftEffect", "Monad0", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}))
})
	})
	return monadEffectContT
}

var monadStateContT gopurs_runtime.Value
var once_monadStateContT sync.Once
func Get_monadStateContT() gopurs_runtime.Value {
	once_monadStateContT.Do(func() {
		monadStateContT = gopurs_runtime.Func(func(dictMonadState_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := gopurs_runtime.Apply(Get_monadContT(), Monad0_1_0)
_ = monadContT1_2_1
__local_var_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind")
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("state", "Monad0", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), x_4))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}))
})
	})
	return monadStateContT
}

var monadSTContT gopurs_runtime.Value
var once_monadSTContT sync.Once
func Get_monadSTContT() gopurs_runtime.Value {
	once_monadSTContT.Do(func() {
		monadSTContT = gopurs_runtime.Func(func(dictMonadST_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := gopurs_runtime.Apply(Get_monadContT(), Monad0_1_0)
_ = monadContT1_2_1
__local_var_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind")
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("liftST", "Monad0", gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}))
})
	})
	return monadSTContT
}

var monoidContT gopurs_runtime.Value
var once_monoidContT sync.Once
func Get_monoidContT() gopurs_runtime.Value {
	once_monoidContT.Do(func() {
		monoidContT = gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_0
semigroupContT2_3_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(a_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "append"), a_6)
_ = __local_var_7_2
return gopurs_runtime.Apply(b_4, gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(__local_var_7_2, a_8))
}))
}))
}))
_ = semigroupContT2_3_1
__local_var_4_3 := gopurs_runtime.RecordGet(dictMonoid_1, "mempty")
_ = __local_var_4_3
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, __local_var_4_3)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupContT2_3_1
}))
})
	})
	return monoidContT
}



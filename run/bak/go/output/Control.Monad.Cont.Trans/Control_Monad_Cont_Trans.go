package Control_Monad_Cont_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_ContT gopurs_runtime.Value
var once_ContT sync.Once
func Get_ContT() gopurs_runtime.Value {
	once_ContT.Do(func() {
		cache_ContT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ContT(x_0_box)
})
	})
	return cache_ContT
}

var cache_withContT gopurs_runtime.Value
var once_withContT sync.Once
func Get_withContT() gopurs_runtime.Value {
	once_withContT.Do(func() {
		cache_withContT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withContT(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_withContT
}

var cache_withContT__gopurs_runtime_Value_927944282 gopurs_runtime.Value
var once_withContT__gopurs_runtime_Value_927944282 sync.Once
func Get_withContT__gopurs_runtime_Value_927944282() gopurs_runtime.Value {
	once_withContT__gopurs_runtime_Value_927944282.Do(func() {
		cache_withContT__gopurs_runtime_Value_927944282 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withContT__gopurs_runtime_Value_927944282(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_withContT__gopurs_runtime_Value_927944282
}

var cache_runContT gopurs_runtime.Value
var once_runContT sync.Once
func Get_runContT() gopurs_runtime.Value {
	once_runContT.Do(func() {
		cache_runContT = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runContT(v_0_box, k_1_box)
})
	})
	return cache_runContT
}

var cache_runContT__gopurs_runtime_Value_1201902778 gopurs_runtime.Value
var once_runContT__gopurs_runtime_Value_1201902778 sync.Once
func Get_runContT__gopurs_runtime_Value_1201902778() gopurs_runtime.Value {
	once_runContT__gopurs_runtime_Value_1201902778.Do(func() {
		cache_runContT__gopurs_runtime_Value_1201902778 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runContT__gopurs_runtime_Value_1201902778(v_0_box, k_1_box)
})
	})
	return cache_runContT__gopurs_runtime_Value_1201902778
}

var cache_runContT__gopurs_runtime_Value_952668154 gopurs_runtime.Value
var once_runContT__gopurs_runtime_Value_952668154 sync.Once
func Get_runContT__gopurs_runtime_Value_952668154() gopurs_runtime.Value {
	once_runContT__gopurs_runtime_Value_952668154.Do(func() {
		cache_runContT__gopurs_runtime_Value_952668154 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runContT__gopurs_runtime_Value_952668154(v_0_box, k_1_box)
})
	})
	return cache_runContT__gopurs_runtime_Value_952668154
}

var cache_newtypeContT gopurs_runtime.Value
var once_newtypeContT sync.Once
func Get_newtypeContT() gopurs_runtime.Value {
	once_newtypeContT.Do(func() {
		cache_newtypeContT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeContT
}

var cache_monadTransContT gopurs_runtime.Value
var once_monadTransContT sync.Once
func Get_monadTransContT() gopurs_runtime.Value {
	once_monadTransContT.Do(func() {
		cache_monadTransContT = gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), m_1, k_2)
})
})
}))
	})
	return cache_monadTransContT
}

var cache_mapContT gopurs_runtime.Value
var once_mapContT sync.Once
func Get_mapContT() gopurs_runtime.Value {
	once_mapContT.Do(func() {
		cache_mapContT = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapContT(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_mapContT
}

var cache_mapContT__gopurs_runtime_Value_3409870043 gopurs_runtime.Value
var once_mapContT__gopurs_runtime_Value_3409870043 sync.Once
func Get_mapContT__gopurs_runtime_Value_3409870043() gopurs_runtime.Value {
	once_mapContT__gopurs_runtime_Value_3409870043.Do(func() {
		cache_mapContT__gopurs_runtime_Value_3409870043 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, k_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapContT__gopurs_runtime_Value_3409870043(f_0_box, v_1_box, k_2_box)
})
	})
	return cache_mapContT__gopurs_runtime_Value_3409870043
}

var cache_functorContT gopurs_runtime.Value
var once_functorContT sync.Once
func Get_functorContT() gopurs_runtime.Value {
	once_functorContT.Do(func() {
		cache_functorContT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorContT(dictFunctor_0_box)
})
	})
	return cache_functorContT
}

var cache_applyContT gopurs_runtime.Value
var once_applyContT sync.Once
func Get_applyContT() gopurs_runtime.Value {
	once_applyContT.Do(func() {
		cache_applyContT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyContT(dictApply_0_box)
})
	})
	return cache_applyContT
}

var cache_bindContT gopurs_runtime.Value
var once_bindContT sync.Once
func Get_bindContT() gopurs_runtime.Value {
	once_bindContT.Do(func() {
		cache_bindContT = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindContT(dictBind_0_box)
})
	})
	return cache_bindContT
}

var cache_semigroupContT gopurs_runtime.Value
var once_semigroupContT sync.Once
func Get_semigroupContT() gopurs_runtime.Value {
	once_semigroupContT.Do(func() {
		cache_semigroupContT = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupContT(dictApply_0_box)
})
	})
	return cache_semigroupContT
}

var cache_applicativeContT gopurs_runtime.Value
var once_applicativeContT sync.Once
func Get_applicativeContT() gopurs_runtime.Value {
	once_applicativeContT.Do(func() {
		cache_applicativeContT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeContT(dictApplicative_0_box)
})
	})
	return cache_applicativeContT
}

var cache_monadContT gopurs_runtime.Value
var once_monadContT sync.Once
func Get_monadContT() gopurs_runtime.Value {
	once_monadContT.Do(func() {
		cache_monadContT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContT(dictMonad_0_box)
})
	})
	return cache_monadContT
}

var cache_monadAskContT gopurs_runtime.Value
var once_monadAskContT sync.Once
func Get_monadAskContT() gopurs_runtime.Value {
	once_monadAskContT.Do(func() {
		cache_monadAskContT = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadAskContT(dictMonadAsk_0_box)
})
	})
	return cache_monadAskContT
}

var cache_monadReaderContT gopurs_runtime.Value
var once_monadReaderContT sync.Once
func Get_monadReaderContT() gopurs_runtime.Value {
	once_monadReaderContT.Do(func() {
		cache_monadReaderContT = gopurs_runtime.Func(func(dictMonadReader_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadReaderContT(dictMonadReader_0_box)
})
	})
	return cache_monadReaderContT
}

var cache_monadContContT gopurs_runtime.Value
var once_monadContContT sync.Once
func Get_monadContContT() gopurs_runtime.Value {
	once_monadContContT.Do(func() {
		cache_monadContContT = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadContContT(dictMonad_0_box)
})
	})
	return cache_monadContContT
}

var cache_monadEffectContT gopurs_runtime.Value
var once_monadEffectContT sync.Once
func Get_monadEffectContT() gopurs_runtime.Value {
	once_monadEffectContT.Do(func() {
		cache_monadEffectContT = gopurs_runtime.Func(func(dictMonadEffect_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadEffectContT(dictMonadEffect_0_box)
})
	})
	return cache_monadEffectContT
}

var cache_monadStateContT gopurs_runtime.Value
var once_monadStateContT sync.Once
func Get_monadStateContT() gopurs_runtime.Value {
	once_monadStateContT.Do(func() {
		cache_monadStateContT = gopurs_runtime.Func(func(dictMonadState_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadStateContT(dictMonadState_0_box)
})
	})
	return cache_monadStateContT
}

var cache_monadSTContT gopurs_runtime.Value
var once_monadSTContT sync.Once
func Get_monadSTContT() gopurs_runtime.Value {
	once_monadSTContT.Do(func() {
		cache_monadSTContT = gopurs_runtime.Func(func(dictMonadST_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadSTContT(dictMonadST_0_box)
})
	})
	return cache_monadSTContT
}

var cache_monoidContT gopurs_runtime.Value
var once_monoidContT sync.Once
func Get_monoidContT() gopurs_runtime.Value {
	once_monoidContT.Do(func() {
		cache_monoidContT = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidContT(dictApplicative_0_box)
})
	})
	return cache_monoidContT
}

func Call_ContT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_withContT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, k_2))
}

func Call_withContT__gopurs_runtime_Value_927944282(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, k_2))
}

func Call_runContT(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(v_0, k_1)
}

func Call_runContT__gopurs_runtime_Value_1201902778(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(v_0, k_1)
}

func Call_runContT__gopurs_runtime_Value_952668154(v_0_loop gopurs_runtime.Value, k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
return gopurs_runtime.Apply(v_0, k_1)
}

func Call_mapContT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
}

func Call_mapContT__gopurs_runtime_Value_3409870043(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, k_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var k_2 gopurs_runtime.Value = k_2_loop
_ = k_2
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, k_2))
}

func Call_functorContT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
}

func Call_applyContT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})
})
}))
}

func Call_bindContT(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
functorContT1_1_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_1
applyContT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})
})
}))
_ = applyContT1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_3, a_5, k_prime_4)
}))
})
})
}))
}

func Call_semigroupContT(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_0
__local_var_2_1 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})
})
}))
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.RecordGet(dictSemigroup_3, "append")
_ = __local_var_4_2
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_4_2, a_5), b_6)
})
}))
})
}

func Call_applicativeContT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
functorContT1_1_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_1
applyContT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})
})
}))
_ = applyContT1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_1_0
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})
}))
}

func Call_monadContT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
functorContT1_1_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_2
applyContT1_1_1 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_2
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})
})
}))
_ = applyContT1_1_1
applicativeContT1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_1_1
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_2)
})
}))
_ = applicativeContT1_1_0
functorContT1_2_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(f_2, a_5))
}))
})
})
}))
_ = functorContT1_2_5
applyContT1_2_4 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_2_5
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_5, gopurs_runtime.Apply(g_6, a_7))
}))
}))
})
})
}))
_ = applyContT1_2_4
bindContT1_2_3 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyContT1_2_4
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(k_4, a_6, k_prime_5)
}))
})
})
}))
_ = bindContT1_2_3
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeContT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindContT1_2_3
}))
}

func Call_monadAskContT(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := Call_monadContT(Monad0_1_0)
_ = monadContT1_2_1
return gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_monadTransContT(), "lift"), Monad0_1_0, gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")))
}

func Call_monadReaderContT(dictMonadReader_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadReader_0 gopurs_runtime.Value = dictMonadReader_0_loop
_ = dictMonadReader_0
MonadAsk0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "MonadAsk0"), gopurs_runtime.Value{})
_ = MonadAsk0_1_0
ask_2_1 := gopurs_runtime.RecordGet(MonadAsk0_1_0, "ask")
_ = ask_2_1
monadAskContT1_3_2 := Call_monadAskContT(MonadAsk0_1_0)
_ = monadAskContT1_3_2
return gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskContT1_3_2
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(MonadAsk0_1_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "bind"), ask_2_1, gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return r_7
}))
_ = __local_var_8_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadReader_0, "local"), f_4, gopurs_runtime.Apply(v_5, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_3, gopurs_runtime.Apply(k_6, x_9))
})))
}))
})
})
}))
}

func Call_monadContContT(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadContT1_1_0 := Call_monadContT(dictMonad_0)
_ = monadContT1_1_0
return gopurs_runtime.RecordDict2("Monad0", "callCC", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, a_4)
})
}), k_3)
})
}))
}

func Call_monadEffectContT(dictMonadEffect_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadEffect_0 gopurs_runtime.Value = dictMonadEffect_0_loop
_ = dictMonadEffect_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := Call_monadContT(Monad0_1_0)
_ = monadContT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransContT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_0, "liftEffect"), x_4))
}))
}

func Call_monadStateContT(dictMonadState_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadState_0 gopurs_runtime.Value = dictMonadState_0_loop
_ = dictMonadState_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := Call_monadContT(Monad0_1_0)
_ = monadContT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransContT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "state", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadState_0, "state"), x_4))
}))
}

func Call_monadSTContT(dictMonadST_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadST_0 gopurs_runtime.Value = dictMonadST_0_loop
_ = dictMonadST_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadContT1_2_1 := Call_monadContT(Monad0_1_0)
_ = monadContT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_monadTransContT(), "lift"), Monad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadContT1_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_0, "liftST"), x_4))
}))
}

func Call_monoidContT(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
functorContT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_3, gopurs_runtime.Apply(f_1, a_4))
}))
})
})
}))
_ = functorContT1_1_0
__local_var_2_1 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorContT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_2, gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_4, gopurs_runtime.Apply(g_5, a_6))
}))
}))
})
})
}))
_ = __local_var_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append")
_ = __local_var_4_3
semigroupContT2_4_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_4_3, a_5), b_6)
})
}))
_ = semigroupContT2_4_2
__local_var_5_4 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
_ = __local_var_5_4
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupContT2_4_2
}), gopurs_runtime.Func(func(k_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(k_6, __local_var_5_4)
}))
})
}



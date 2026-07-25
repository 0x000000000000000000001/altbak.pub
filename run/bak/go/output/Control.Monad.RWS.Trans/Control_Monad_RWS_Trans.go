package Control_Monad_RWS_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Control_Monad_Rec_Class "gopurs/output/Control.Monad.Rec.Class"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	unsafe "unsafe"
)

var cache_RWSResult gopurs_runtime.Value
var once_RWSResult sync.Once
func Get_RWSResult() gopurs_runtime.Value {
	once_RWSResult.Do(func() {
		cache_RWSResult = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{value0, value1, value2})}
})
})
})
	})
	return cache_RWSResult
}

var cache_RWST gopurs_runtime.Value
var once_RWST sync.Once
func Get_RWST() gopurs_runtime.Value {
	once_RWST.Do(func() {
		cache_RWST = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_RWST
}

var cache_withRWST gopurs_runtime.Value
var once_withRWST sync.Once
func Get_withRWST() gopurs_runtime.Value {
	once_withRWST.Do(func() {
		cache_withRWST = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withRWST(f_0_box, m_1_box, r_2_box, s_3_box)
})
	})
	return cache_withRWST
}

var cache_runRWST gopurs_runtime.Value
var once_runRWST sync.Once
func Get_runRWST() gopurs_runtime.Value {
	once_runRWST.Do(func() {
		cache_runRWST = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}()
})
	})
	return cache_runRWST
}

var cache_newtypeRWST gopurs_runtime.Value
var once_newtypeRWST sync.Once
func Get_newtypeRWST() gopurs_runtime.Value {
	once_newtypeRWST.Do(func() {
		cache_newtypeRWST = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeRWST
}

var cache_monadTransRWST gopurs_runtime.Value
var once_monadTransRWST sync.Once
func Get_monadTransRWST() gopurs_runtime.Value {
	once_monadTransRWST.Do(func() {
		cache_monadTransRWST = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.RecordDict1("lift", gopurs_runtime.Func4(func(dictMonad_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Bind1"), gopurs_runtime.Value{}), "bind"), m_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_2, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{s_5, a_6, mempty_1_0})})
}))
}))
}()
})
	})
	return cache_monadTransRWST
}

var cache_mapRWST gopurs_runtime.Value
var once_mapRWST sync.Once
func Get_mapRWST() gopurs_runtime.Value {
	once_mapRWST.Do(func() {
		cache_mapRWST = gopurs_runtime.Func4(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapRWST(f_0_box, v_1_box, r_2_box, s_3_box)
})
	})
	return cache_mapRWST
}

var cache_lazyRWST gopurs_runtime.Value
var once_lazyRWST sync.Once
func Get_lazyRWST() gopurs_runtime.Value {
	once_lazyRWST.Do(func() {
		cache_lazyRWST = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, r_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, pkg_Data_Unit.Get_unit(), r_1, s_2)
}))
	})
	return cache_lazyRWST
}

var cache_functorRWST gopurs_runtime.Value
var once_functorRWST sync.Once
func Get_functorRWST() gopurs_runtime.Value {
	once_functorRWST.Do(func() {
		cache_functorRWST = gopurs_runtime.Func(func(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func4(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value, r_3 gopurs_runtime.Value, s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v1_5.UnsafePtr).V0, gopurs_runtime.Apply(f_1, (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_5.UnsafePtr).V1), (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_5.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_2, r_3, s_4))
}))
}()
})
	})
	return cache_functorRWST
}

var cache_execRWST gopurs_runtime.Value
var once_execRWST sync.Once
func Get_execRWST() gopurs_runtime.Value {
	once_execRWST.Do(func() {
		cache_execRWST = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_execRWST(dictMonad_0_box, v_1_box, r_2_box, s_3_box)
})
	})
	return cache_execRWST
}

var cache_evalRWST gopurs_runtime.Value
var once_evalRWST sync.Once
func Get_evalRWST() gopurs_runtime.Value {
	once_evalRWST.Do(func() {
		cache_evalRWST = gopurs_runtime.Func4(func(dictMonad_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, r_2_box gopurs_runtime.Value, s_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_evalRWST(dictMonad_0_box, v_1_box, r_2_box, s_3_box)
})
	})
	return cache_evalRWST
}

var cache_applyRWST gopurs_runtime.Value
var once_applyRWST sync.Once
func Get_applyRWST() gopurs_runtime.Value {
	once_applyRWST.Do(func() {
		cache_applyRWST = gopurs_runtime.Func(func(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = Functor0_1_0
functorRWST1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V0, gopurs_runtime.Apply(f_2, (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V1), (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_3, r_4, s_5))
}))
_ = functorRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func4(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_4, r_6, s_7), gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_2 := (*Data_Control_Monad_RWS_Trans_RWSResult)(v2_8.UnsafePtr).V2
_ = __local_var_9_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v3_10.UnsafePtr).V0, gopurs_runtime.Apply((*Data_Control_Monad_RWS_Trans_RWSResult)(v2_8.UnsafePtr).V1, (*Data_Control_Monad_RWS_Trans_RWSResult)(v3_10.UnsafePtr).V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append"), __local_var_9_2, (*Data_Control_Monad_RWS_Trans_RWSResult)(v3_10.UnsafePtr).V2)})}
}), gopurs_runtime.Apply2(v1_5, r_6, (*Data_Control_Monad_RWS_Trans_RWSResult)(v2_8.UnsafePtr).V0))
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_2_1
}))
})
}()
})
	})
	return cache_applyRWST
}

var cache_bindRWST gopurs_runtime.Value
var once_bindRWST sync.Once
func Get_bindRWST() gopurs_runtime.Value {
	once_bindRWST.Do(func() {
		cache_bindRWST = gopurs_runtime.Func(func(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
applyRWST1_2_1 := gopurs_runtime.Apply(Get_applyRWST(), dictBind_0)
_ = applyRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
applyRWST2_4_2 := gopurs_runtime.Apply(applyRWST1_2_1, dictMonoid_3)
_ = applyRWST2_4_2
return gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func4(func(v_5 gopurs_runtime.Value, f_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), gopurs_runtime.Apply2(v_5, r_7, s_8), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_10_3 := (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_9.UnsafePtr).V2
_ = __local_var_10_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v3_11.UnsafePtr).V0, (*Data_Control_Monad_RWS_Trans_RWSResult)(v3_11.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append"), __local_var_10_3, (*Data_Control_Monad_RWS_Trans_RWSResult)(v3_11.UnsafePtr).V2)})}
}), gopurs_runtime.Apply3(f_6, (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_9.UnsafePtr).V1, r_7, (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_9.UnsafePtr).V0))
}))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyRWST2_4_2
}))
})
}()
})
	})
	return cache_bindRWST
}

var cache_semigroupRWST gopurs_runtime.Value
var once_semigroupRWST sync.Once
func Get_semigroupRWST() gopurs_runtime.Value {
	once_semigroupRWST.Do(func() {
		cache_semigroupRWST = gopurs_runtime.Func(func(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
applyRWST1_1_0 := gopurs_runtime.Apply(Get_applyRWST(), dictBind_0)
_ = applyRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(applyRWST1_1_0, dictMonoid_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.RecordGet(dictSemigroup_4, "append")
_ = __local_var_5_2
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_1, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_1, "Functor0"), gopurs_runtime.Value{}), "map"), __local_var_5_2, a_6), b_7)
}))
})
})
}()
})
	})
	return cache_semigroupRWST
}

var cache_applicativeRWST gopurs_runtime.Value
var once_applicativeRWST sync.Once
func Get_applicativeRWST() gopurs_runtime.Value {
	once_applicativeRWST.Do(func() {
		cache_applicativeRWST = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applyRWST1_1_0 := gopurs_runtime.Apply(Get_applyRWST(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = applyRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_3_1 := gopurs_runtime.RecordGet(dictMonoid_2, "mempty")
_ = mempty_3_1
applyRWST2_4_2 := gopurs_runtime.Apply(applyRWST1_1_0, dictMonoid_2)
_ = applyRWST2_4_2
return gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func3(func(a_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{s_7, a_5, mempty_3_1})})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyRWST2_4_2
}))
})
}()
})
	})
	return cache_applicativeRWST
}

var cache_monadRWST gopurs_runtime.Value
var once_monadRWST sync.Once
func Get_monadRWST() gopurs_runtime.Value {
	once_monadRWST.Do(func() {
		cache_monadRWST = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeRWST1_1_0 := gopurs_runtime.Apply(Get_applicativeRWST(), dictMonad_0)
_ = applicativeRWST1_1_0
bindRWST1_2_1 := gopurs_runtime.Apply(Get_bindRWST(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = bindRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeRWST2_4_2 := gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3)
_ = applicativeRWST2_4_2
bindRWST2_5_3 := gopurs_runtime.Apply(bindRWST1_2_1, dictMonoid_3)
_ = bindRWST2_5_3
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeRWST2_4_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindRWST2_5_3
}))
})
}()
})
	})
	return cache_monadRWST
}

var cache_monadAskRWST gopurs_runtime.Value
var once_monadAskRWST sync.Once
func Get_monadAskRWST() gopurs_runtime.Value {
	once_monadAskRWST.Do(func() {
		cache_monadAskRWST = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadRWST1_1_0 := gopurs_runtime.Apply(Get_monadRWST(), dictMonad_0)
_ = monadRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_3_1 := gopurs_runtime.RecordGet(dictMonoid_2, "mempty")
_ = mempty_3_1
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_1_0, dictMonoid_2)
_ = monadRWST2_4_2
return gopurs_runtime.RecordDict2("ask", "Monad0", gopurs_runtime.Func2(func(r_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{s_6, r_5, mempty_3_1})})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
}))
})
}()
})
	})
	return cache_monadAskRWST
}

var cache_monadReaderRWST gopurs_runtime.Value
var once_monadReaderRWST sync.Once
func Get_monadReaderRWST() gopurs_runtime.Value {
	once_monadReaderRWST.Do(func() {
		cache_monadReaderRWST = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadAskRWST1_1_0 := gopurs_runtime.Apply(Get_monadAskRWST(), dictMonad_0)
_ = monadAskRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskRWST2_3_1 := gopurs_runtime.Apply(monadAskRWST1_1_0, dictMonoid_2)
_ = monadAskRWST2_3_1
return gopurs_runtime.RecordDict2("local", "MonadAsk0", gopurs_runtime.Func4(func(f_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(m_5, gopurs_runtime.Apply(f_4, r_6), s_7)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskRWST2_3_1
}))
})
}()
})
	})
	return cache_monadReaderRWST
}

var cache_monadEffectRWS gopurs_runtime.Value
var once_monadEffectRWS sync.Once
func Get_monadEffectRWS() gopurs_runtime.Value {
	once_monadEffectRWS.Do(func() {
		cache_monadEffectRWS = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(dictMonadEffect_2 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_2, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_1
monadRWST1_4_2 := gopurs_runtime.Apply2(Get_monadRWST(), Monad0_3_1, dictMonoid_0)
_ = monadRWST1_4_2
return gopurs_runtime.RecordDict2("liftEffect", "Monad0", gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadEffect_2, "liftEffect"), x_5)
_ = __local_var_6_3
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Bind1"), gopurs_runtime.Value{}), "bind"), __local_var_6_3, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{s_8, a_9, mempty_1_0})})
}))
})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST1_4_2
}))
})
}()
})
	})
	return cache_monadEffectRWS
}

var cache_monadRecRWST gopurs_runtime.Value
var once_monadRecRWST sync.Once
func Get_monadRecRWST() gopurs_runtime.Value {
	once_monadRecRWST.Do(func() {
		cache_monadRecRWST = gopurs_runtime.Func(func(dictMonadRec_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadRec_0 gopurs_runtime.Value = dictMonadRec_0_loop
_ = dictMonadRec_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadRec_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadRWST1_2_1 := gopurs_runtime.Apply(Get_monadRWST(), Monad0_1_0)
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_4_2
mempty_5_3 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
_ = mempty_5_3
monadRWST2_6_4 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_6_4
return gopurs_runtime.RecordDict2("tailRecM", "Monad0", gopurs_runtime.Func4(func(k_7 gopurs_runtime.Value, a_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadRec_0, "tailRecM"), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_12_5 := (*Data_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V2
_ = __local_var_12_5
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply3(k_7, (*Data_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V1, r_9, (*Data_Control_Monad_RWS_Trans_RWSResult)(v_11.UnsafePtr).V0), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if ((*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V1.Type == 9 && (*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V1.IntVal == 525585346) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop{gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V0, (*pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Loop)((*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "append"), __local_var_12_5, (*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V2)})}})}
goto end_branch_6
} else {

}
}
{
if ((*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V1.Type == 9 && (*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V1.IntVal == 60402430) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done{gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V0, (*pkg_Control_Monad_Rec_Class.Data_Control_Monad_Rec_Class_Done)((*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_2, "append"), __local_var_12_5, (*Data_Control_Monad_RWS_Trans_RWSResult)(v2_13.UnsafePtr).V2)})}})}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t6)
}))
}), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{s_10, a_8, mempty_5_3})})
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_6_4
}))
})
}()
})
	})
	return cache_monadRecRWST
}

var cache_monadStateRWST gopurs_runtime.Value
var once_monadStateRWST sync.Once
func Get_monadStateRWST() gopurs_runtime.Value {
	once_monadStateRWST.Do(func() {
		cache_monadStateRWST = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadRWST1_1_0 := gopurs_runtime.Apply(Get_monadRWST(), dictMonad_0)
_ = monadRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_3_1 := gopurs_runtime.RecordGet(dictMonoid_2, "mempty")
_ = mempty_3_1
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_1_0, dictMonoid_2)
_ = monadRWST2_4_2
return gopurs_runtime.RecordDict2("state", "Monad0", gopurs_runtime.Func3(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
v1_8_3 := gopurs_runtime.Apply(f_5, s_7)
_ = v1_8_3
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v1_8_3.UnsafePtr).V1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v1_8_3.UnsafePtr).V0, mempty_3_1})})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
}))
})
}()
})
	})
	return cache_monadStateRWST
}

var cache_monadTellRWST gopurs_runtime.Value
var once_monadTellRWST sync.Once
func Get_monadTellRWST() gopurs_runtime.Value {
	once_monadTellRWST.Do(func() {
		cache_monadTellRWST = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
monadRWST1_1_0 := gopurs_runtime.Apply(Get_monadRWST(), dictMonad_0)
_ = monadRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{})
_ = Semigroup0_3_1
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_1_0, dictMonoid_2)
_ = monadRWST2_4_2
return gopurs_runtime.RecordDict3("tell", "Semigroup0", "Monad1", gopurs_runtime.Func3(func(w_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{s_7, pkg_Data_Unit.Get_unit(), w_5})})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_3_1
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
}))
})
}()
})
	})
	return cache_monadTellRWST
}

var cache_monadWriterRWST gopurs_runtime.Value
var once_monadWriterRWST sync.Once
func Get_monadWriterRWST() gopurs_runtime.Value {
	once_monadWriterRWST.Do(func() {
		cache_monadWriterRWST = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_2_1
monadTellRWST1_3_2 := gopurs_runtime.Apply(Get_monadTellRWST(), dictMonad_0)
_ = monadTellRWST1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
monadTellRWST2_5_3 := gopurs_runtime.Apply(monadTellRWST1_3_2, dictMonoid_4)
_ = monadTellRWST2_5_3
return gopurs_runtime.RecordDict4("listen", "pass", "Monoid0", "MonadTell1", gopurs_runtime.Func3(func(m_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply2(m_6, r_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*Data_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V1, (*Data_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V2})}, (*Data_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V2})})
}))
}), gopurs_runtime.Func3(func(m_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "bind"), gopurs_runtime.Apply2(m_6, r_7, s_8), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*Data_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V1.UnsafePtr).V0, gopurs_runtime.Apply((*pkg_Data_Tuple.Data_Data_Tuple_Tuple)((*Data_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V1.UnsafePtr).V1, (*Data_Control_Monad_RWS_Trans_RWSResult)(v_9.UnsafePtr).V2)})})
}))
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellRWST2_5_3
}))
})
}()
})
	})
	return cache_monadWriterRWST
}

var cache_monadThrowRWST gopurs_runtime.Value
var once_monadThrowRWST sync.Once
func Get_monadThrowRWST() gopurs_runtime.Value {
	once_monadThrowRWST.Do(func() {
		cache_monadThrowRWST = gopurs_runtime.Func(func(dictMonadThrow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadThrow_0 gopurs_runtime.Value = dictMonadThrow_0_loop
_ = dictMonadThrow_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
monadRWST1_2_1 := gopurs_runtime.Apply(Get_monadRWST(), Monad0_1_0)
_ = monadRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_4_2 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
_ = mempty_4_2
monadRWST2_5_3 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
_ = monadRWST2_5_3
return gopurs_runtime.RecordDict2("throwError", "Monad0", gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadThrow_0, "throwError"), e_6)
_ = __local_var_7_4
return gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), __local_var_7_4, gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{s_9, a_10, mempty_4_2})})
}))
})
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_5_3
}))
})
}()
})
	})
	return cache_monadThrowRWST
}

var cache_monadErrorRWST gopurs_runtime.Value
var once_monadErrorRWST sync.Once
func Get_monadErrorRWST() gopurs_runtime.Value {
	once_monadErrorRWST.Do(func() {
		cache_monadErrorRWST = gopurs_runtime.Func(func(dictMonadError_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadError_0 gopurs_runtime.Value = dictMonadError_0_loop
_ = dictMonadError_0
monadThrowRWST1_1_0 := gopurs_runtime.Apply(Get_monadThrowRWST(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadError_0, "MonadThrow0"), gopurs_runtime.Value{}))
_ = monadThrowRWST1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowRWST2_3_1 := gopurs_runtime.Apply(monadThrowRWST1_1_0, dictMonoid_2)
_ = monadThrowRWST2_3_1
return gopurs_runtime.RecordDict2("catchError", "MonadThrow0", gopurs_runtime.Func4(func(m_4 gopurs_runtime.Value, h_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadError_0, "catchError"), gopurs_runtime.Apply2(m_4, r_6, s_7), gopurs_runtime.Func(func(e_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(h_5, e_8, r_6, s_7)
}))
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowRWST2_3_1
}))
})
}()
})
	})
	return cache_monadErrorRWST
}

var cache_monadSTRWST gopurs_runtime.Value
var once_monadSTRWST sync.Once
func Get_monadSTRWST() gopurs_runtime.Value {
	once_monadSTRWST.Do(func() {
		cache_monadSTRWST = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
return gopurs_runtime.Func(func(dictMonadST_2 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_2, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_1
monadRWST1_4_2 := gopurs_runtime.Apply2(Get_monadRWST(), Monad0_3_1, dictMonoid_0)
_ = monadRWST1_4_2
return gopurs_runtime.RecordDict2("liftST", "Monad0", gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadST_2, "liftST"), x_5)
_ = __local_var_6_3
return gopurs_runtime.Func2(func(v_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Bind1"), gopurs_runtime.Value{}), "bind"), __local_var_6_3, gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_1, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{s_8, a_9, mempty_1_0})})
}))
})
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST1_4_2
}))
})
}()
})
	})
	return cache_monadSTRWST
}

var cache_monoidRWST gopurs_runtime.Value
var once_monoidRWST sync.Once
func Get_monoidRWST() gopurs_runtime.Value {
	once_monoidRWST.Do(func() {
		cache_monoidRWST = gopurs_runtime.Func(func(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
applicativeRWST1_1_0 := gopurs_runtime.Apply(Get_applicativeRWST(), dictMonad_0)
_ = applicativeRWST1_1_0
semigroupRWST1_2_1 := gopurs_runtime.Apply(Get_semigroupRWST(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = semigroupRWST1_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupRWST2_4_2 := gopurs_runtime.Apply(semigroupRWST1_2_1, dictMonoid_3)
_ = semigroupRWST2_4_2
return gopurs_runtime.Func(func(dictMonoid1_5 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupRWST3_6_3 := gopurs_runtime.Apply(semigroupRWST2_4_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid1_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupRWST3_6_3
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3), "pure"), gopurs_runtime.RecordGet(dictMonoid1_5, "mempty")), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRWST3_6_3
}))
})
})
}()
})
	})
	return cache_monoidRWST
}

var cache_altRWST gopurs_runtime.Value
var once_altRWST sync.Once
func Get_altRWST() gopurs_runtime.Value {
	once_altRWST.Do(func() {
		cache_altRWST = gopurs_runtime.Func(func(dictAlt_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictAlt_0 gopurs_runtime.Value = dictAlt_0_loop
_ = dictAlt_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlt_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorRWST1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func4(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value, r_4 gopurs_runtime.Value, s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V0, gopurs_runtime.Apply(f_2, (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V1), (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_6.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_3, r_4, s_5))
}))
_ = functorRWST1_2_1
return gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func4(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value, r_5 gopurs_runtime.Value, s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictAlt_0, "alt"), gopurs_runtime.Apply2(v_3, r_5, s_6), gopurs_runtime.Apply2(v1_4, r_5, s_6))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_2_1
}))
}()
})
	})
	return cache_altRWST
}

var cache_plusRWST gopurs_runtime.Value
var once_plusRWST sync.Once
func Get_plusRWST() gopurs_runtime.Value {
	once_plusRWST.Do(func() {
		cache_plusRWST = gopurs_runtime.Func(func(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictPlus_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_2
functorRWST1_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func4(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value, r_6 gopurs_runtime.Value, s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_2, "map"), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V1), (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_8.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_5, r_6, s_7))
}))
_ = functorRWST1_4_4
altRWST1_4_3 := gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func4(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value, r_7 gopurs_runtime.Value, s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "alt"), gopurs_runtime.Apply2(v_5, r_7, s_8), gopurs_runtime.Apply2(v1_6, r_7, s_8))
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_4_4
}))
_ = altRWST1_4_3
return gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_1_0
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return altRWST1_4_3
}))
}()
})
	})
	return cache_plusRWST
}

var cache_alternativeRWST gopurs_runtime.Value
var once_alternativeRWST sync.Once
func Get_alternativeRWST() gopurs_runtime.Value {
	once_alternativeRWST.Do(func() {
		cache_alternativeRWST = gopurs_runtime.Func2(func(dictMonoid_0_box gopurs_runtime.Value, dictAlternative_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alternativeRWST(dictMonoid_0_box, dictAlternative_1_box)
})
	})
	return cache_alternativeRWST
}

type Data_Control_Monad_RWS_Trans_RWSResult struct {
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}
func Is_Data_Control_Monad_RWS_Trans_RWSResult(v gopurs_runtime.Value) bool {
	return v.Type == 9 && v.IntVal == 2367475031
}

func Call_withRWST(f_0_loop gopurs_runtime.Value, m_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
__local_var_4_0 := gopurs_runtime.Apply2(f_0, r_2, s_3)
_ = __local_var_4_0
return gopurs_runtime.Apply2(m_1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(__local_var_4_0.UnsafePtr).V0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(__local_var_4_0.UnsafePtr).V1)
}

func Call_mapRWST(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_execRWST(dictMonad_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_1, r_2, s_3), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*Data_Control_Monad_RWS_Trans_RWSResult)(v1_4.UnsafePtr).V0, (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_4.UnsafePtr).V2})})
}))
}

func Call_evalRWST(dictMonad_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, r_2_loop gopurs_runtime.Value, s_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var r_2 gopurs_runtime.Value = r_2_loop
_ = r_2
var s_3 gopurs_runtime.Value = s_3_loop
_ = s_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(v_1, r_2, s_3), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*Data_Control_Monad_RWS_Trans_RWSResult)(v1_4.UnsafePtr).V1, (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_4.UnsafePtr).V2})})
}))
}

func Call_alternativeRWST(dictMonoid_0_loop gopurs_runtime.Value, dictAlternative_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
var dictAlternative_1 gopurs_runtime.Value = dictAlternative_1_loop
_ = dictAlternative_1
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_1, "Plus1"), gopurs_runtime.Value{})
_ = __local_var_2_0
empty_3_1 := gopurs_runtime.RecordGet(__local_var_2_0, "empty")
_ = empty_3_1
__local_var_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_0, "Alt0"), gopurs_runtime.Value{})
_ = __local_var_4_3
__local_var_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_4
functorRWST1_6_5 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func4(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value, r_8 gopurs_runtime.Value, s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_4, "map"), gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2367475031, UnsafePtr: unsafe.Pointer(&Data_Control_Monad_RWS_Trans_RWSResult{(*Data_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V0, gopurs_runtime.Apply(f_6, (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V1), (*Data_Control_Monad_RWS_Trans_RWSResult)(v1_10.UnsafePtr).V2})}
}), gopurs_runtime.Apply2(v_7, r_8, s_9))
}))
_ = functorRWST1_6_5
altRWST1_7_6 := gopurs_runtime.RecordDict2("alt", "Functor0", gopurs_runtime.Func4(func(v_7 gopurs_runtime.Value, v1_8 gopurs_runtime.Value, r_9 gopurs_runtime.Value, s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_3, "alt"), gopurs_runtime.Apply2(v_7, r_9, s_10), gopurs_runtime.Apply2(v1_8, r_9, s_10))
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_6_5
}))
_ = altRWST1_7_6
plusRWST1_4_2 := gopurs_runtime.RecordDict2("empty", "Alt0", gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_3_1
}), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return altRWST1_7_6
}))
_ = plusRWST1_4_2
return gopurs_runtime.Func(func(dictMonad_5 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeRWST1_6_7 := gopurs_runtime.Apply2(Get_applicativeRWST(), dictMonad_5, dictMonoid_0)
_ = applicativeRWST1_6_7
return gopurs_runtime.RecordDict2("Applicative0", "Plus1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeRWST1_6_7
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return plusRWST1_4_2
}))
})
}



package Control_Monad_State_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var StateT gopurs_runtime.Value
var once_StateT sync.Once
func Get_StateT() gopurs_runtime.Value {
	once_StateT.Do(func() {
		StateT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return StateT
}

var withStateT gopurs_runtime.Value
var once_withStateT sync.Once
func Get_withStateT() gopurs_runtime.Value {
	once_withStateT.Do(func() {
		withStateT = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1, gopurs_runtime.Apply(f_0, x_2))
})
})
})
	})
	return withStateT
}

var runStateT gopurs_runtime.Value
var once_runStateT sync.Once
func Get_runStateT() gopurs_runtime.Value {
	once_runStateT.Do(func() {
		runStateT = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runStateT
}

var newtypeStateT gopurs_runtime.Value
var once_newtypeStateT sync.Once
func Get_newtypeStateT() gopurs_runtime.Value {
	once_newtypeStateT.Do(func() {
		newtypeStateT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeStateT
}

var monadTransStateT gopurs_runtime.Value
var once_monadTransStateT sync.Once
func Get_monadTransStateT() gopurs_runtime.Value {
	once_monadTransStateT.Do(func() {
		monadTransStateT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lift": gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], m_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_3, "value1": s_2}))
}))
})
})
})})
	})
	return monadTransStateT
}

var mapStateT gopurs_runtime.Value
var once_mapStateT sync.Once
func Get_mapStateT() gopurs_runtime.Value {
	once_mapStateT.Do(func() {
		mapStateT = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2))
})
})
})
	})
	return mapStateT
}

var lazyStateT gopurs_runtime.Value
var once_lazyStateT sync.Once
func Get_lazyStateT() gopurs_runtime.Value {
	once_lazyStateT.Do(func() {
		lazyStateT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"defer": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()), s_1)
})
})})
	})
	return lazyStateT
}

var functorStateT gopurs_runtime.Value
var once_functorStateT sync.Once
func Get_functorStateT() gopurs_runtime.Value {
	once_functorStateT.Do(func() {
		functorStateT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_1, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})), gopurs_runtime.Apply(v_2, s_3))
})
})
})})
})
	})
	return functorStateT
}

var execStateT gopurs_runtime.Value
var once_execStateT sync.Once
func Get_execStateT() gopurs_runtime.Value {
	once_execStateT.Do(func() {
		execStateT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_snd()), gopurs_runtime.Apply(v_1, s_2))
})
})
})
	})
	return execStateT
}

var evalStateT gopurs_runtime.Value
var once_evalStateT sync.Once
func Get_evalStateT() gopurs_runtime.Value {
	once_evalStateT.Do(func() {
		evalStateT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_fst()), gopurs_runtime.Apply(v_1, s_2))
})
})
})
	})
	return evalStateT
}

var monadStateT gopurs_runtime.Value
var once_monadStateT sync.Once
func Get_monadStateT() gopurs_runtime.Value {
	once_monadStateT.Do(func() {
		monadStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), dictMonad_0)
})})
})
	})
	return monadStateT
}

var bindStateT gopurs_runtime.Value
var once_bindStateT sync.Once
func Get_bindStateT() gopurs_runtime.Value {
	once_bindStateT.Do(func() {
		bindStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(v_1, s_3)), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_2, v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}))
})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyStateT(), dictMonad_0)
})})
})
	})
	return bindStateT
}

var applyStateT gopurs_runtime.Value
var once_applyStateT sync.Once
func Get_applyStateT() gopurs_runtime.Value {
	once_applyStateT.Do(func() {
		applyStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStateT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_2, v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})), gopurs_runtime.Apply(v_3, s_4))
})
})
})})
__local_var_3_2 := gopurs_runtime.Apply(Get_bindStateT(), dictMonad_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["bind"], f_4), gopurs_runtime.Func(func(f_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["bind"], a_5), gopurs_runtime.Func(func(a_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(f_prime_6, a_prime_7))
}))
}))
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_2_1
})})
})
	})
	return applyStateT
}

var applicativeStateT gopurs_runtime.Value
var once_applicativeStateT sync.Once
func Get_applicativeStateT() gopurs_runtime.Value {
	once_applicativeStateT.Do(func() {
		applicativeStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_1, "value1": s_2}))
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applyStateT(), dictMonad_0)
})})
})
	})
	return applicativeStateT
}

var semigroupStateT gopurs_runtime.Value
var once_semigroupStateT sync.Once
func Get_semigroupStateT() gopurs_runtime.Value {
	once_semigroupStateT.Do(func() {
		semigroupStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(Get_applyStateT(), dictMonad_0)
return gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := dictSemigroup_2.PtrVal.(map[string]gopurs_runtime.Value)["append"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], __local_var_3_1), a_4)), b_5)
})
})})
})
})
	})
	return semigroupStateT
}

var monadAskStateT gopurs_runtime.Value
var once_monadAskStateT sync.Once
func Get_monadAskStateT() gopurs_runtime.Value {
	once_monadAskStateT.Do(func() {
		monadAskStateT = gopurs_runtime.Func(func(dictMonadAsk_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadStateT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
})})
__local_var_3_2 := dictMonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["ask"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_3_2), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_5, "value1": s_4}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
})})
})
	})
	return monadAskStateT
}

var monadReaderStateT gopurs_runtime.Value
var once_monadReaderStateT sync.Once
func Get_monadReaderStateT() gopurs_runtime.Value {
	once_monadReaderStateT.Do(func() {
		monadReaderStateT = gopurs_runtime.Func(func(dictMonadReader_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskStateT1_1_0 := gopurs_runtime.Apply(Get_monadAskStateT(), gopurs_runtime.Apply(dictMonadReader_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadAsk0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictMonadReader_0.PtrVal.(map[string]gopurs_runtime.Value)["local"], x_2)
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(v_4, x_5))
})
})
}), "MonadAsk0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskStateT1_1_0
})})
})
	})
	return monadReaderStateT
}

var monadContStateT gopurs_runtime.Value
var once_monadContStateT sync.Once
func Get_monadContStateT() gopurs_runtime.Value {
	once_monadContStateT.Do(func() {
		monadContStateT = gopurs_runtime.Func(func(dictMonadCont_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonadCont_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadStateT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), __local_var_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), __local_var_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"callCC": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadCont_0.PtrVal.(map[string]gopurs_runtime.Value)["callCC"], gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_3, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_5, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_6, "value1": s_prime_7}))
})
})), s_4)
}))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
})})
})
	})
	return monadContStateT
}

var monadEffectState gopurs_runtime.Value
var once_monadEffectState sync.Once
func Get_monadEffectState() gopurs_runtime.Value {
	once_monadEffectState.Do(func() {
		monadEffectState = gopurs_runtime.Func(func(dictMonadEffect_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadStateT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftEffect": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictMonadEffect_0.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], x_3)
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_4_2), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_6, "value1": s_5}))
}))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
})})
})
	})
	return monadEffectState
}

var monadRecStateT gopurs_runtime.Value
var once_monadRecStateT sync.Once
func Get_monadRecStateT() gopurs_runtime.Value {
	once_monadRecStateT.Do(func() {
		monadRecStateT = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadStateT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(f_3, v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"])), gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Bool(v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t2 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v2_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})})
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], __t2)
}))
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_4, "value1": s_5}))
})
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
})})
})
	})
	return monadRecStateT
}

var monadStateStateT gopurs_runtime.Value
var once_monadStateStateT sync.Once
func Get_monadStateStateT() gopurs_runtime.Value {
	once_monadStateStateT.Do(func() {
		monadStateStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadStateT1_1_0 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), dictMonad_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"state": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(f_2, x_3))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_1_0
})})
})
	})
	return monadStateStateT
}

var monadTellStateT gopurs_runtime.Value
var once_monadTellStateT sync.Once
func Get_monadTellStateT() gopurs_runtime.Value {
	once_monadTellStateT.Do(func() {
		monadTellStateT = gopurs_runtime.Func(func(dictMonadTell_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad1_1_0 := gopurs_runtime.Apply(dictMonadTell_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad1"], gopurs_runtime.Value{})
Semigroup0_2_1 := gopurs_runtime.Apply(dictMonadTell_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
monadStateT1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad1_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad1_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tell": gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(dictMonadTell_0.PtrVal.(map[string]gopurs_runtime.Value)["tell"], x_4)
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_5_3), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_7, "value1": s_6}))
}))
})
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_2_1
}), "Monad1": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_3_2
})})
})
	})
	return monadTellStateT
}

var monadWriterStateT gopurs_runtime.Value
var once_monadWriterStateT sync.Once
func Get_monadWriterStateT() gopurs_runtime.Value {
	once_monadWriterStateT.Do(func() {
		monadWriterStateT = gopurs_runtime.Func(func(dictMonadWriter_0 gopurs_runtime.Value) gopurs_runtime.Value {
MonadTell1_1_0 := gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadTell1"], gopurs_runtime.Value{})
Monad1_2_1 := gopurs_runtime.Apply(MonadTell1_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad1"], gopurs_runtime.Value{})
__local_var_3_2 := gopurs_runtime.Apply(Monad1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
__local_var_4_3 := gopurs_runtime.Apply(Monad1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
Monoid0_5_4 := gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["Monoid0"], gopurs_runtime.Value{})
monadTellStateT1_6_5 := gopurs_runtime.Apply(Get_monadTellStateT(), MonadTell1_1_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"listen": gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["listen"], gopurs_runtime.Apply(m_7, s_8))), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value1": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
}))
})
}), "pass": gopurs_runtime.Func(func(m_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadWriter_0.PtrVal.(map[string]gopurs_runtime.Value)["pass"], gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(m_7, s_8)), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value1": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
})))
})
}), "Monoid0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Monoid0_5_4
}), "MonadTell1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellStateT1_6_5
})})
})
	})
	return monadWriterStateT
}

var monadThrowStateT gopurs_runtime.Value
var once_monadThrowStateT sync.Once
func Get_monadThrowStateT() gopurs_runtime.Value {
	once_monadThrowStateT.Do(func() {
		monadThrowStateT = gopurs_runtime.Func(func(dictMonadThrow_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadThrow_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadStateT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"throwError": gopurs_runtime.Func(func(e_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictMonadThrow_0.PtrVal.(map[string]gopurs_runtime.Value)["throwError"], e_3)
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_4_2), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_6, "value1": s_5}))
}))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
})})
})
	})
	return monadThrowStateT
}

var monadErrorStateT gopurs_runtime.Value
var once_monadErrorStateT sync.Once
func Get_monadErrorStateT() gopurs_runtime.Value {
	once_monadErrorStateT.Do(func() {
		monadErrorStateT = gopurs_runtime.Func(func(dictMonadError_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowStateT1_1_0 := gopurs_runtime.Apply(Get_monadThrowStateT(), gopurs_runtime.Apply(dictMonadError_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadThrow0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"catchError": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadError_0.PtrVal.(map[string]gopurs_runtime.Value)["catchError"], gopurs_runtime.Apply(v_2, s_4)), gopurs_runtime.Func(func(e_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(h_3, e_5), s_4)
}))
})
})
}), "MonadThrow0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowStateT1_1_0
})})
})
	})
	return monadErrorStateT
}

var monadSTStateT gopurs_runtime.Value
var once_monadSTStateT sync.Once
func Get_monadSTStateT() gopurs_runtime.Value {
	once_monadSTStateT.Do(func() {
		monadSTStateT = gopurs_runtime.Func(func(dictMonadST_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadST_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadStateT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftST": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictMonadST_0.PtrVal.(map[string]gopurs_runtime.Value)["liftST"], x_3)
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_4_2), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": x_6, "value1": s_5}))
}))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
})})
})
	})
	return monadSTStateT
}

var monoidStateT gopurs_runtime.Value
var once_monoidStateT sync.Once
func Get_monoidStateT() gopurs_runtime.Value {
	once_monoidStateT.Do(func() {
		monoidStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupStateT1_1_0 := gopurs_runtime.Apply(Get_semigroupStateT(), dictMonad_0)
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupStateT2_3_1 := gopurs_runtime.Apply(semigroupStateT1_1_0, gopurs_runtime.Apply(dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0).PtrVal.(map[string]gopurs_runtime.Value)["pure"], dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupStateT2_3_1
})})
})
})
	})
	return monoidStateT
}

var altStateT gopurs_runtime.Value
var once_altStateT sync.Once
func Get_altStateT() gopurs_runtime.Value {
	once_altStateT.Do(func() {
		altStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictAlt_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictAlt_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStateT1_3_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_3, v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})), gopurs_runtime.Apply(v_4, s_5))
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_1.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(v_4, s_6)), gopurs_runtime.Apply(v1_5, s_6))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_3_1
})})
})
})
	})
	return altStateT
}

var plusStateT gopurs_runtime.Value
var once_plusStateT sync.Once
func Get_plusStateT() gopurs_runtime.Value {
	once_plusStateT.Do(func() {
		plusStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictPlus_1 gopurs_runtime.Value) gopurs_runtime.Value {
empty_2_0 := dictPlus_1.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
__local_var_3_1 := gopurs_runtime.Apply(dictPlus_1.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_4_2 := gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStateT1_5_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_5, v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})), gopurs_runtime.Apply(v_6, s_7))
})
})
})})
altStateT2_5_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(v_6, s_8)), gopurs_runtime.Apply(v1_7, s_8))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_5_4
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_2_0
}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return altStateT2_5_3
})})
})
})
	})
	return plusStateT
}

var alternativeStateT gopurs_runtime.Value
var once_alternativeStateT sync.Once
func Get_alternativeStateT() gopurs_runtime.Value {
	once_alternativeStateT.Do(func() {
		alternativeStateT = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeStateT1_1_0 := gopurs_runtime.Apply(Get_applicativeStateT(), dictMonad_0)
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictAlternative_2.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{})
empty_4_2 := __local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
__local_var_5_4 := gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_6_5 := gopurs_runtime.Apply(__local_var_5_4.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorStateT1_7_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_6_5.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_7, v1_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v1_10.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})), gopurs_runtime.Apply(v_8, s_9))
})
})
})})
altStateT2_8_7 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_4.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(v_8, s_10)), gopurs_runtime.Apply(v1_9, s_10))
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorStateT1_7_6
})})
plusStateT2_5_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_4_2
}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altStateT2_8_7
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeStateT1_1_0
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return plusStateT2_5_3
})})
})
})
	})
	return alternativeStateT
}

var monadPlusStateT gopurs_runtime.Value
var once_monadPlusStateT sync.Once
func Get_monadPlusStateT() gopurs_runtime.Value {
	once_monadPlusStateT.Do(func() {
		monadPlusStateT = gopurs_runtime.Func(func(dictMonadPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadStateT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_applicativeStateT(), Monad0_1_0)
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_bindStateT(), Monad0_1_0)
})})
alternativeStateT1_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_alternativeStateT(), Monad0_1_0), gopurs_runtime.Apply(dictMonadPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Alternative1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Monad0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadStateT1_2_1
}), "Alternative1": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeStateT1_3_2
})})
})
	})
	return monadPlusStateT
}



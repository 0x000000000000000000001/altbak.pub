package Control_Monad_RWS_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var RWSResult gopurs_runtime.Value
var once_RWSResult sync.Once
func Get_RWSResult() gopurs_runtime.Value {
	once_RWSResult.Do(func() {
		RWSResult = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": value0, "value1": value1, "value2": value2})
})
})
})
	})
	return RWSResult
}

var RWST gopurs_runtime.Value
var once_RWST sync.Once
func Get_RWST() gopurs_runtime.Value {
	once_RWST.Do(func() {
		RWST = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return RWST
}

var withRWST gopurs_runtime.Value
var once_withRWST sync.Once
func Get_withRWST() gopurs_runtime.Value {
	once_withRWST.Do(func() {
		withRWST = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, r_2), s_3)
return gopurs_runtime.Apply(gopurs_runtime.Apply(m_1, __local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), __local_var_4_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
})
})
})
	})
	return withRWST
}

var runRWST gopurs_runtime.Value
var once_runRWST sync.Once
func Get_runRWST() gopurs_runtime.Value {
	once_runRWST.Do(func() {
		runRWST = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runRWST
}

var newtypeRWST gopurs_runtime.Value
var once_newtypeRWST sync.Once
func Get_newtypeRWST() gopurs_runtime.Value {
	once_newtypeRWST.Do(func() {
		newtypeRWST = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeRWST
}

var monadTransRWST gopurs_runtime.Value
var once_monadTransRWST sync.Once
func Get_monadTransRWST() gopurs_runtime.Value {
	once_monadTransRWST.Do(func() {
		monadTransRWST = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lift": gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], m_3), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": s_5, "value1": a_6, "value2": mempty_1_0}))
}))
})
})
})
})})
})
	})
	return monadTransRWST
}

var mapRWST gopurs_runtime.Value
var once_mapRWST sync.Once
func Get_mapRWST() gopurs_runtime.Value {
	once_mapRWST.Do(func() {
		mapRWST = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, r_2), s_3))
})
})
})
})
	})
	return mapRWST
}

var lazyRWST gopurs_runtime.Value
var once_lazyRWST sync.Once
func Get_lazyRWST() gopurs_runtime.Value {
	once_lazyRWST.Do(func() {
		lazyRWST = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"defer": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()), r_1), s_2)
})
})
})})
	})
	return lazyRWST
}

var functorRWST gopurs_runtime.Value
var once_functorRWST sync.Once
func Get_functorRWST() gopurs_runtime.Value {
	once_functorRWST.Do(func() {
		functorRWST = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_1, v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": v1_5.PtrVal.(map[string]gopurs_runtime.Value)["value2"]})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(v_2, r_3), s_4))
})
})
})
})})
})
	})
	return functorRWST
}

var execRWST gopurs_runtime.Value
var once_execRWST sync.Once
func Get_execRWST() gopurs_runtime.Value {
	once_execRWST.Do(func() {
		execRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, r_2), s_3)), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]}))
}))
})
})
})
})
	})
	return execRWST
}

var evalRWST gopurs_runtime.Value
var once_evalRWST sync.Once
func Get_evalRWST() gopurs_runtime.Value {
	once_evalRWST.Do(func() {
		evalRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, r_2), s_3)), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": v1_4.PtrVal.(map[string]gopurs_runtime.Value)["value2"]}))
}))
})
})
})
})
	})
	return evalRWST
}

var applyRWST gopurs_runtime.Value
var once_applyRWST sync.Once
func Get_applyRWST() gopurs_runtime.Value {
	once_applyRWST.Do(func() {
		applyRWST = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorRWST1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_2, v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value2"]})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(v_3, r_4), s_5))
})
})
})
})})
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(v_4, r_6), s_7)), gopurs_runtime.Func(func(v2_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_2 := v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v3_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"], v3_10.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], __local_var_9_2), v3_10.PtrVal.(map[string]gopurs_runtime.Value)["value2"])})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(v1_5, r_6), v2_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
}))
})
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_2_1
})})
})
})
	})
	return applyRWST
}

var bindRWST gopurs_runtime.Value
var once_bindRWST sync.Once
func Get_bindRWST() gopurs_runtime.Value {
	once_bindRWST.Do(func() {
		bindRWST = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
applyRWST1_2_1 := gopurs_runtime.Apply(Get_applyRWST(), dictBind_0)
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
applyRWST2_4_2 := gopurs_runtime.Apply(applyRWST1_2_1, dictMonoid_3)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(v_5, r_7), s_8)), gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_10_3 := v1_9.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v3_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v3_11.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v3_11.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value2": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["append"], __local_var_10_3), v3_11.PtrVal.(map[string]gopurs_runtime.Value)["value2"])})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(f_6, v1_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), r_7), v1_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
}))
})
})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyRWST2_4_2
})})
})
})
	})
	return bindRWST
}

var semigroupRWST gopurs_runtime.Value
var once_semigroupRWST sync.Once
func Get_semigroupRWST() gopurs_runtime.Value {
	once_semigroupRWST.Do(func() {
		semigroupRWST = gopurs_runtime.Func(func(dictBind_0 gopurs_runtime.Value) gopurs_runtime.Value {
applyRWST1_1_0 := gopurs_runtime.Apply(Get_applyRWST(), dictBind_0)
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(applyRWST1_1_0, dictMonoid_2)
return gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := dictSemigroup_4.PtrVal.(map[string]gopurs_runtime.Value)["append"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], __local_var_5_2), a_6)), b_7)
})
})})
})
})
})
	})
	return semigroupRWST
}

var applicativeRWST gopurs_runtime.Value
var once_applicativeRWST sync.Once
func Get_applicativeRWST() gopurs_runtime.Value {
	once_applicativeRWST.Do(func() {
		applicativeRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
applyRWST1_1_0 := gopurs_runtime.Apply(Get_applyRWST(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_3_1 := dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
applyRWST2_4_2 := gopurs_runtime.Apply(applyRWST1_1_0, dictMonoid_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": s_7, "value1": a_5, "value2": mempty_3_1}))
})
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return applyRWST2_4_2
})})
})
})
	})
	return applicativeRWST
}

var monadRWST gopurs_runtime.Value
var once_monadRWST sync.Once
func Get_monadRWST() gopurs_runtime.Value {
	once_monadRWST.Do(func() {
		monadRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeRWST1_1_0 := gopurs_runtime.Apply(Get_applicativeRWST(), dictMonad_0)
bindRWST1_2_1 := gopurs_runtime.Apply(Get_bindRWST(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeRWST2_4_2 := gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3)
bindRWST2_5_3 := gopurs_runtime.Apply(bindRWST1_2_1, dictMonoid_3)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeRWST2_4_2
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindRWST2_5_3
})})
})
})
	})
	return monadRWST
}

var monadAskRWST gopurs_runtime.Value
var once_monadAskRWST sync.Once
func Get_monadAskRWST() gopurs_runtime.Value {
	once_monadAskRWST.Do(func() {
		monadAskRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadRWST1_1_0 := gopurs_runtime.Apply(Get_monadRWST(), dictMonad_0)
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_3_1 := dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_1_0, dictMonoid_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": s_6, "value1": r_5, "value2": mempty_3_1}))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
})})
})
})
	})
	return monadAskRWST
}

var monadReaderRWST gopurs_runtime.Value
var once_monadReaderRWST sync.Once
func Get_monadReaderRWST() gopurs_runtime.Value {
	once_monadReaderRWST.Do(func() {
		monadReaderRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskRWST1_1_0 := gopurs_runtime.Apply(Get_monadAskRWST(), dictMonad_0)
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskRWST2_3_1 := gopurs_runtime.Apply(monadAskRWST1_1_0, dictMonoid_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(m_5, gopurs_runtime.Apply(f_4, r_6)), s_7)
})
})
})
}), "MonadAsk0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskRWST2_3_1
})})
})
})
	})
	return monadReaderRWST
}

var monadEffectRWS gopurs_runtime.Value
var once_monadEffectRWS sync.Once
func Get_monadEffectRWS() gopurs_runtime.Value {
	once_monadEffectRWS.Do(func() {
		monadEffectRWS = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(dictMonadEffect_2 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_3_1 := gopurs_runtime.Apply(dictMonadEffect_2.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadRWST1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_monadRWST(), Monad0_3_1), dictMonoid_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftEffect": gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(dictMonadEffect_2.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], x_5)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_6_3), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": s_8, "value1": a_9, "value2": mempty_1_0}))
}))
})
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST1_4_2
})})
})
})
	})
	return monadEffectRWS
}

var monadRecRWST gopurs_runtime.Value
var once_monadRecRWST sync.Once
func Get_monadRecRWST() gopurs_runtime.Value {
	once_monadRecRWST.Do(func() {
		monadRecRWST = gopurs_runtime.Func(func(dictMonadRec_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadRWST1_2_1 := gopurs_runtime.Apply(Get_monadRWST(), Monad0_1_0)
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
mempty_5_3 := dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadRWST2_6_4 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_0.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_12_5 := v_11.PtrVal.(map[string]gopurs_runtime.Value)["value2"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(k_7, v_11.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), r_9), v_11.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Func(func(v2_13 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_13.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v2_13.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v2_13.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], __local_var_12_5), v2_13.PtrVal.(map[string]gopurs_runtime.Value)["value2"])})})
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(v2_13.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v2_13.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v2_13.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["append"], __local_var_12_5), v2_13.PtrVal.(map[string]gopurs_runtime.Value)["value2"])})})
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], __t6)
}))
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": s_10, "value1": a_8, "value2": mempty_5_3}))
})
})
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_6_4
})})
})
})
	})
	return monadRecRWST
}

var monadStateRWST gopurs_runtime.Value
var once_monadStateRWST sync.Once
func Get_monadStateRWST() gopurs_runtime.Value {
	once_monadStateRWST.Do(func() {
		monadStateRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadRWST1_1_0 := gopurs_runtime.Apply(Get_monadRWST(), dictMonad_0)
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_3_1 := dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_1_0, dictMonoid_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"state": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
v1_8_3 := gopurs_runtime.Apply(f_5, s_7)
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v1_8_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": v1_8_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": mempty_3_1}))
})
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
})})
})
})
	})
	return monadStateRWST
}

var monadTellRWST gopurs_runtime.Value
var once_monadTellRWST sync.Once
func Get_monadTellRWST() gopurs_runtime.Value {
	once_monadTellRWST.Do(func() {
		monadTellRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadRWST1_1_0 := gopurs_runtime.Apply(Get_monadRWST(), dictMonad_0)
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_3_1 := gopurs_runtime.Apply(dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
monadRWST2_4_2 := gopurs_runtime.Apply(monadRWST1_1_0, dictMonoid_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tell": gopurs_runtime.Func(func(w_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": s_7, "value1": pkg_Data_Unit.Get_unit(), "value2": w_5}))
})
})
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_3_1
}), "Monad1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_4_2
})})
})
})
	})
	return monadTellRWST
}

var monadWriterRWST gopurs_runtime.Value
var once_monadWriterRWST sync.Once
func Get_monadWriterRWST() gopurs_runtime.Value {
	once_monadWriterRWST.Do(func() {
		monadWriterRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
monadTellRWST1_3_2 := gopurs_runtime.Apply(Get_monadTellRWST(), dictMonad_0)
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
monadTellRWST2_5_3 := gopurs_runtime.Apply(monadTellRWST1_3_2, dictMonoid_4)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"listen": gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(m_6, r_7), s_8)), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value2"]}), "value2": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value2"]}))
}))
})
})
}), "pass": gopurs_runtime.Func(func(m_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(m_6, r_7), s_8)), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value2": gopurs_runtime.Apply(v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], v_9.PtrVal.(map[string]gopurs_runtime.Value)["value2"])}))
}))
})
})
}), "Monoid0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_4
}), "MonadTell1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellRWST2_5_3
})})
})
})
	})
	return monadWriterRWST
}

var monadThrowRWST gopurs_runtime.Value
var once_monadThrowRWST sync.Once
func Get_monadThrowRWST() gopurs_runtime.Value {
	once_monadThrowRWST.Do(func() {
		monadThrowRWST = gopurs_runtime.Func(func(dictMonadThrow_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadThrow_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadRWST1_2_1 := gopurs_runtime.Apply(Get_monadRWST(), Monad0_1_0)
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_4_2 := dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadRWST2_5_3 := gopurs_runtime.Apply(monadRWST1_2_1, dictMonoid_3)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"throwError": gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_4 := gopurs_runtime.Apply(dictMonadThrow_0.PtrVal.(map[string]gopurs_runtime.Value)["throwError"], e_6)
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_7_4), gopurs_runtime.Func(func(a_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": s_9, "value1": a_10, "value2": mempty_4_2}))
}))
})
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST2_5_3
})})
})
})
	})
	return monadThrowRWST
}

var monadErrorRWST gopurs_runtime.Value
var once_monadErrorRWST sync.Once
func Get_monadErrorRWST() gopurs_runtime.Value {
	once_monadErrorRWST.Do(func() {
		monadErrorRWST = gopurs_runtime.Func(func(dictMonadError_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowRWST1_1_0 := gopurs_runtime.Apply(Get_monadThrowRWST(), gopurs_runtime.Apply(dictMonadError_0.PtrVal.(map[string]gopurs_runtime.Value)["MonadThrow0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowRWST2_3_1 := gopurs_runtime.Apply(monadThrowRWST1_1_0, dictMonoid_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"catchError": gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadError_0.PtrVal.(map[string]gopurs_runtime.Value)["catchError"], gopurs_runtime.Apply(gopurs_runtime.Apply(m_4, r_6), s_7)), gopurs_runtime.Func(func(e_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(h_5, e_8), r_6), s_7)
}))
})
})
})
}), "MonadThrow0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowRWST2_3_1
})})
})
})
	})
	return monadErrorRWST
}

var monadSTRWST gopurs_runtime.Value
var once_monadSTRWST sync.Once
func Get_monadSTRWST() gopurs_runtime.Value {
	once_monadSTRWST.Do(func() {
		monadSTRWST = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Func(func(dictMonadST_2 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_3_1 := gopurs_runtime.Apply(dictMonadST_2.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadRWST1_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_monadRWST(), Monad0_3_1), dictMonoid_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftST": gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(dictMonadST_2.PtrVal.(map[string]gopurs_runtime.Value)["liftST"], x_5)
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], __local_var_6_3), gopurs_runtime.Func(func(a_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": s_8, "value1": a_9, "value2": mempty_1_0}))
}))
})
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadRWST1_4_2
})})
})
})
	})
	return monadSTRWST
}

var monoidRWST gopurs_runtime.Value
var once_monoidRWST sync.Once
func Get_monoidRWST() gopurs_runtime.Value {
	once_monoidRWST.Do(func() {
		monoidRWST = gopurs_runtime.Func(func(dictMonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeRWST1_1_0 := gopurs_runtime.Apply(Get_applicativeRWST(), dictMonad_0)
semigroupRWST1_2_1 := gopurs_runtime.Apply(Get_semigroupRWST(), gopurs_runtime.Apply(dictMonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupRWST2_4_2 := gopurs_runtime.Apply(semigroupRWST1_2_1, dictMonoid_3)
return gopurs_runtime.Func(func(dictMonoid1_5 gopurs_runtime.Value) gopurs_runtime.Value {
semigroupRWST3_6_3 := gopurs_runtime.Apply(semigroupRWST2_4_2, gopurs_runtime.Apply(dictMonoid1_5.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Apply(gopurs_runtime.Apply(applicativeRWST1_1_0, dictMonoid_3).PtrVal.(map[string]gopurs_runtime.Value)["pure"], dictMonoid1_5.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupRWST3_6_3
})})
})
})
})
	})
	return monoidRWST
}

var altRWST gopurs_runtime.Value
var once_altRWST sync.Once
func Get_altRWST() gopurs_runtime.Value {
	once_altRWST.Do(func() {
		altRWST = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorRWST1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_2, v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": v1_6.PtrVal.(map[string]gopurs_runtime.Value)["value2"]})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(v_3, r_4), s_5))
})
})
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(gopurs_runtime.Apply(v_3, r_5), s_6)), gopurs_runtime.Apply(gopurs_runtime.Apply(v1_4, r_5), s_6))
})
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_2_1
})})
})
	})
	return altRWST
}

var plusRWST gopurs_runtime.Value
var once_plusRWST sync.Once
func Get_plusRWST() gopurs_runtime.Value {
	once_plusRWST.Do(func() {
		plusRWST = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
empty_1_0 := dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
__local_var_2_1 := gopurs_runtime.Apply(dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_3_2 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorRWST1_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_4, v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value2"]})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(v_5, r_6), s_7))
})
})
})
})})
altRWST1_4_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(gopurs_runtime.Apply(v_5, r_7), s_8)), gopurs_runtime.Apply(gopurs_runtime.Apply(v1_6, r_7), s_8))
})
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_4_4
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_1_0
})
}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return altRWST1_4_3
})})
})
	})
	return plusRWST
}

var alternativeRWST gopurs_runtime.Value
var once_alternativeRWST sync.Once
func Get_alternativeRWST() gopurs_runtime.Value {
	once_alternativeRWST.Do(func() {
		alternativeRWST = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictAlternative_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := gopurs_runtime.Apply(dictAlternative_1.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{})
empty_3_1 := __local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"]
__local_var_4_3 := gopurs_runtime.Apply(__local_var_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_5_4 := gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorRWST1_6_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_4.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("RWSResult"), "value0": v1_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_6, v1_10.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), "value2": v1_10.PtrVal.(map[string]gopurs_runtime.Value)["value2"]})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(v_7, r_8), s_9))
})
})
})
})})
altRWST1_7_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_3.PtrVal.(map[string]gopurs_runtime.Value)["alt"], gopurs_runtime.Apply(gopurs_runtime.Apply(v_7, r_9), s_10)), gopurs_runtime.Apply(gopurs_runtime.Apply(v1_8, r_9), s_10))
})
})
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorRWST1_6_5
})})
plusRWST1_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return empty_3_1
})
}), "Alt0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return altRWST1_7_6
})})
return gopurs_runtime.Func(func(dictMonad_5 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeRWST1_6_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applicativeRWST(), dictMonad_5), dictMonoid_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeRWST1_6_7
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return plusRWST1_4_2
})})
})
})
})
	})
	return alternativeRWST
}



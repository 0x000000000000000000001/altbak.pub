package Control_Monad_Writer_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var WriterT gopurs_runtime.Value
var once_WriterT sync.Once
func Get_WriterT() gopurs_runtime.Value {
	once_WriterT.Do(func() {
		WriterT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return WriterT
}

var runWriterT gopurs_runtime.Value
var once_runWriterT sync.Once
func Get_runWriterT() gopurs_runtime.Value {
	once_runWriterT.Do(func() {
		runWriterT = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runWriterT
}

var newtypeWriterT gopurs_runtime.Value
var once_newtypeWriterT sync.Once
func Get_newtypeWriterT() gopurs_runtime.Value {
	once_newtypeWriterT.Do(func() {
		newtypeWriterT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeWriterT
}

var monadTransWriterT gopurs_runtime.Value
var once_monadTransWriterT sync.Once
func Get_monadTransWriterT() gopurs_runtime.Value {
	once_monadTransWriterT.Do(func() {
		monadTransWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lift": gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], m_3), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_4, "value1": mempty_1_0}))
}))
})
})})
})
	})
	return monadTransWriterT
}

var mapWriterT gopurs_runtime.Value
var once_mapWriterT sync.Once
func Get_mapWriterT() gopurs_runtime.Value {
	once_mapWriterT.Do(func() {
		mapWriterT = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, v_1)
})
})
	})
	return mapWriterT
}

var functorWriterT gopurs_runtime.Value
var once_functorWriterT sync.Once
func Get_functorWriterT() gopurs_runtime.Value {
	once_functorWriterT.Do(func() {
		functorWriterT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_1, v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}))
})})
})
	})
	return functorWriterT
}

var execWriterT gopurs_runtime.Value
var once_execWriterT sync.Once
func Get_execWriterT() gopurs_runtime.Value {
	once_execWriterT.Do(func() {
		execWriterT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Tuple.Get_snd()), v_1)
})
})
	})
	return execWriterT
}

var applyWriterT gopurs_runtime.Value
var once_applyWriterT sync.Once
func Get_applyWriterT() gopurs_runtime.Value {
	once_applyWriterT.Do(func() {
		applyWriterT = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictApply_1 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_2_0 := gopurs_runtime.Apply(dictApply_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorWriterT1_3_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Functor0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_3, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}))
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_1.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(v3_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v4_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v3_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v4_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})), v_4)), v1_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_3_1
})})
})
})
	})
	return applyWriterT
}

var bindWriterT gopurs_runtime.Value
var once_bindWriterT sync.Once
func Get_bindWriterT() gopurs_runtime.Value {
	once_bindWriterT.Do(func() {
		bindWriterT = gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictBind_1 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_2_0 := gopurs_runtime.Apply(dictBind_1.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
Functor0_3_1 := gopurs_runtime.Apply(Apply0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorWriterT1_4_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Functor0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_4, v_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}))
})})
applyWriterT2_5_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_3_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v3_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(v3_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v4_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], v3_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v4_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})), v_5)), v1_6)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_4_2
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"bind": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictBind_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_6), gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_4 := v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v3_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v3_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], __local_var_9_4), v3_10.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})), gopurs_runtime.Apply(k_7, v1_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
}))
})
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_5_3
})})
})
})
	})
	return bindWriterT
}

var semigroupWriterT gopurs_runtime.Value
var once_semigroupWriterT sync.Once
func Get_semigroupWriterT() gopurs_runtime.Value {
	once_semigroupWriterT.Do(func() {
		semigroupWriterT = gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_2_0 := gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictSemigroup1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictApply_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v3_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(v3_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v4_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], v3_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v4_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(dictSemigroup1_3.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_6.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_6.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})), a_4))), b_5)
})
})})
})
})
})
	})
	return semigroupWriterT
}

var applicativeWriterT gopurs_runtime.Value
var once_applicativeWriterT sync.Once
func Get_applicativeWriterT() gopurs_runtime.Value {
	once_applicativeWriterT.Do(func() {
		applicativeWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
__local_var_2_1 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictApplicative_3.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
Functor0_5_3 := gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorWriterT1_6_5 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Functor0_5_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_6, v_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}))
})})
applyWriterT2_6_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"apply": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_5_3.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(v3_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v4_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], v3_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v4_10.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})), v_7)), v1_8)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_6_5
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pure": gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictApplicative_3.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_7, "value1": mempty_1_0}))
}), "Apply0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applyWriterT2_6_4
})})
})
})
	})
	return applicativeWriterT
}

var monadWriterT gopurs_runtime.Value
var once_monadWriterT sync.Once
func Get_monadWriterT() gopurs_runtime.Value {
	once_monadWriterT.Do(func() {
		monadWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeWriterT1_1_0 := gopurs_runtime.Apply(Get_applicativeWriterT(), dictMonoid_0)
bindWriterT1_2_1 := gopurs_runtime.Apply(Get_bindWriterT(), gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeWriterT2_4_2 := gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(dictMonad_3.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}))
bindWriterT2_5_3 := gopurs_runtime.Apply(bindWriterT1_2_1, gopurs_runtime.Apply(dictMonad_3.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_4_2
}), "Bind1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return bindWriterT2_5_3
})})
})
})
	})
	return monadWriterT
}

var monadAskWriterT gopurs_runtime.Value
var once_monadAskWriterT sync.Once
func Get_monadAskWriterT() gopurs_runtime.Value {
	once_monadAskWriterT.Do(func() {
		monadAskWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadAsk_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(dictMonadAsk_3.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], dictMonadAsk_3.PtrVal.(map[string]gopurs_runtime.Value)["ask"]), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_6, "value1": mempty_1_0}))
})), "Monad0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
})})
})
})
	})
	return monadAskWriterT
}

var monadReaderWriterT gopurs_runtime.Value
var once_monadReaderWriterT sync.Once
func Get_monadReaderWriterT() gopurs_runtime.Value {
	once_monadReaderWriterT.Do(func() {
		monadReaderWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskWriterT1_1_0 := gopurs_runtime.Apply(Get_monadAskWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadReader_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadAskWriterT2_3_1 := gopurs_runtime.Apply(monadAskWriterT1_1_0, gopurs_runtime.Apply(dictMonadReader_2.PtrVal.(map[string]gopurs_runtime.Value)["MonadAsk0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadReader_2.PtrVal.(map[string]gopurs_runtime.Value)["local"], f_4)
}), "MonadAsk0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadAskWriterT2_3_1
})})
})
})
	})
	return monadReaderWriterT
}

var monadContWriterT gopurs_runtime.Value
var once_monadContWriterT sync.Once
func Get_monadContWriterT() gopurs_runtime.Value {
	once_monadContWriterT.Do(func() {
		monadContWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadCont_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_2_1, gopurs_runtime.Apply(dictMonadCont_3.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"callCC": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadCont_3.PtrVal.(map[string]gopurs_runtime.Value)["callCC"], gopurs_runtime.Func(func(c_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(c_6, gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_7, "value1": mempty_1_0}))
}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
})})
})
})
	})
	return monadContWriterT
}

var monadEffectWriter gopurs_runtime.Value
var once_monadEffectWriter sync.Once
func Get_monadEffectWriter() gopurs_runtime.Value {
	once_monadEffectWriter.Do(func() {
		monadEffectWriter = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadEffect_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(dictMonadEffect_3.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftEffect": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadEffect_3.PtrVal.(map[string]gopurs_runtime.Value)["liftEffect"], x_6)), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_7, "value1": mempty_1_0}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
})})
})
})
	})
	return monadEffectWriter
}

var monadRecWriterT gopurs_runtime.Value
var once_monadRecWriterT sync.Once
func Get_monadRecWriterT() gopurs_runtime.Value {
	once_monadRecWriterT.Do(func() {
		monadRecWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
mempty_2_1 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadWriterT1_3_2 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadRec_4 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_5_3 := gopurs_runtime.Apply(dictMonadRec_4.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadWriterT2_6_4 := gopurs_runtime.Apply(monadWriterT1_3_2, Monad0_5_3)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tailRecM": gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadRec_4.PtrVal.(map[string]gopurs_runtime.Value)["tailRecM"], gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_10_5 := v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_5_3.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(f_7, v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"])), gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v2_11.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Loop")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Loop"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v2_11.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], __local_var_10_5), v2_11.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})})
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(v2_11.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Done")).IntVal != 0 {
__t6 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Done"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v2_11.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["append"], __local_var_10_5), v2_11.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})})
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_5_3.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], __t6)
}))
})), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_8, "value1": mempty_2_1}))
})
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_6_4
})})
})
})
	})
	return monadRecWriterT
}

var monadStateWriterT gopurs_runtime.Value
var once_monadStateWriterT sync.Once
func Get_monadStateWriterT() gopurs_runtime.Value {
	once_monadStateWriterT.Do(func() {
		monadStateWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadState_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(dictMonadState_3.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"state": gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadState_3.PtrVal.(map[string]gopurs_runtime.Value)["state"], f_6)), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_7, "value1": mempty_1_0}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
})})
})
})
	})
	return monadStateWriterT
}

var monadTellWriterT gopurs_runtime.Value
var once_monadTellWriterT sync.Once
func Get_monadTellWriterT() gopurs_runtime.Value {
	once_monadTellWriterT.Do(func() {
		monadTellWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.Apply(dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonad_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_2_1, dictMonad_3)
__local_var_5_3 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), pkg_Data_Unit.Get_unit())
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"tell": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonad_3.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(__local_var_5_3, x_6))
}), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Semigroup0_1_0
}), "Monad1": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
})})
})
})
	})
	return monadTellWriterT
}

var monadWriterWriterT gopurs_runtime.Value
var once_monadWriterWriterT sync.Once
func Get_monadWriterWriterT() gopurs_runtime.Value {
	once_monadWriterWriterT.Do(func() {
		monadWriterWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadTellWriterT1_1_0 := gopurs_runtime.Apply(Get_monadTellWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictMonad_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
__local_var_4_2 := gopurs_runtime.Apply(dictMonad_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{})
monadTellWriterT2_5_3 := gopurs_runtime.Apply(monadTellWriterT1_1_0, dictMonad_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"listen": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}), "value1": v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"]}))
}))
}), "pass": gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], v_6), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value0"].PtrVal.(map[string]gopurs_runtime.Value)["value1"], v1_7.PtrVal.(map[string]gopurs_runtime.Value)["value1"])}))
}))
}), "Monoid0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return dictMonoid_0
}), "MonadTell1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadTellWriterT2_5_3
})})
})
})
	})
	return monadWriterWriterT
}

var monadThrowWriterT gopurs_runtime.Value
var once_monadThrowWriterT sync.Once
func Get_monadThrowWriterT() gopurs_runtime.Value {
	once_monadThrowWriterT.Do(func() {
		monadThrowWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadThrow_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(dictMonadThrow_3.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"throwError": gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadThrow_3.PtrVal.(map[string]gopurs_runtime.Value)["throwError"], e_6)), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_7, "value1": mempty_1_0}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
})})
})
})
	})
	return monadThrowWriterT
}

var monadErrorWriterT gopurs_runtime.Value
var once_monadErrorWriterT sync.Once
func Get_monadErrorWriterT() gopurs_runtime.Value {
	once_monadErrorWriterT.Do(func() {
		monadErrorWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowWriterT1_1_0 := gopurs_runtime.Apply(Get_monadThrowWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadError_2 gopurs_runtime.Value) gopurs_runtime.Value {
monadThrowWriterT2_3_1 := gopurs_runtime.Apply(monadThrowWriterT1_1_0, gopurs_runtime.Apply(dictMonadError_2.PtrVal.(map[string]gopurs_runtime.Value)["MonadThrow0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"catchError": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(h_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadError_2.PtrVal.(map[string]gopurs_runtime.Value)["catchError"], v_4), gopurs_runtime.Func(func(e_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_5, e_6)
}))
})
}), "MonadThrow0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return monadThrowWriterT2_3_1
})})
})
})
	})
	return monadErrorWriterT
}

var monadSTWriterT gopurs_runtime.Value
var once_monadSTWriterT sync.Once
func Get_monadSTWriterT() gopurs_runtime.Value {
	once_monadSTWriterT.Do(func() {
		monadSTWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
monadWriterT1_2_1 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadST_3 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_4_2 := gopurs_runtime.Apply(dictMonadST_3.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
monadWriterT2_5_3 := gopurs_runtime.Apply(monadWriterT1_2_1, Monad0_4_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"liftST": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(dictMonadST_3.PtrVal.(map[string]gopurs_runtime.Value)["liftST"], x_6)), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": a_7, "value1": mempty_1_0}))
}))
}), "Monad0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_5_3
})})
})
})
	})
	return monadSTWriterT
}

var monoidWriterT gopurs_runtime.Value
var once_monoidWriterT sync.Once
func Get_monoidWriterT() gopurs_runtime.Value {
	once_monoidWriterT.Do(func() {
		monoidWriterT = gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictApplicative_0.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictMonoid_2.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
Functor0_4_2 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
return gopurs_runtime.Func(func(dictMonoid1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(dictMonoid1_5.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{})
semigroupWriterT3_7_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"append": gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v3_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v4_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(v3_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"], v4_10.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_3_1.PtrVal.(map[string]gopurs_runtime.Value)["append"], v3_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"]), v4_10.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})), gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_4_2.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(__local_var_6_3.PtrVal.(map[string]gopurs_runtime.Value)["append"], v_9.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_9.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})), a_7))), b_8)
})
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"mempty": gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Get_applicativeWriterT(), dictMonoid_2), dictApplicative_0).PtrVal.(map[string]gopurs_runtime.Value)["pure"], dictMonoid1_5.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]), "Semigroup0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupWriterT3_7_4
})})
})
})
})
	})
	return monoidWriterT
}

var altWriterT gopurs_runtime.Value
var once_altWriterT sync.Once
func Get_altWriterT() gopurs_runtime.Value {
	once_altWriterT.Do(func() {
		altWriterT = gopurs_runtime.Func(func(dictAlt_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorWriterT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_2, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}))
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictAlt_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], v_3), v1_4)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_2_1
})})
})
	})
	return altWriterT
}

var plusWriterT gopurs_runtime.Value
var once_plusWriterT sync.Once
func Get_plusWriterT() gopurs_runtime.Value {
	once_plusWriterT.Do(func() {
		plusWriterT = gopurs_runtime.Func(func(dictPlus_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorWriterT1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_3, v_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}))
})})
altWriterT1_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["alt"], v_4), v1_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_3_3
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": dictPlus_0.PtrVal.(map[string]gopurs_runtime.Value)["empty"], "Alt0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_3_2
})})
})
	})
	return plusWriterT
}

var alternativeWriterT gopurs_runtime.Value
var once_alternativeWriterT sync.Once
func Get_alternativeWriterT() gopurs_runtime.Value {
	once_alternativeWriterT.Do(func() {
		alternativeWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeWriterT1_1_0 := gopurs_runtime.Apply(Get_applicativeWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictAlternative_2 gopurs_runtime.Value) gopurs_runtime.Value {
applicativeWriterT2_3_1 := gopurs_runtime.Apply(applicativeWriterT1_1_0, gopurs_runtime.Apply(dictAlternative_2.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}))
__local_var_4_2 := gopurs_runtime.Apply(dictAlternative_2.PtrVal.(map[string]gopurs_runtime.Value)["Plus1"], gopurs_runtime.Value{})
__local_var_5_3 := gopurs_runtime.Apply(__local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["Alt0"], gopurs_runtime.Value{})
__local_var_6_5 := gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorWriterT1_7_6 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_7, v_8.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_8.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
}))
})})
altWriterT1_8_7 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"alt": gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_5_3.PtrVal.(map[string]gopurs_runtime.Value)["alt"], v_8), v1_9)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWriterT1_7_6
})})
plusWriterT1_6_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"empty": __local_var_4_2.PtrVal.(map[string]gopurs_runtime.Value)["empty"], "Alt0": gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return altWriterT1_8_7
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Applicative0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeWriterT2_3_1
}), "Plus1": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return plusWriterT1_6_4
})})
})
})
	})
	return alternativeWriterT
}

var monadPlusWriterT gopurs_runtime.Value
var once_monadPlusWriterT sync.Once
func Get_monadPlusWriterT() gopurs_runtime.Value {
	once_monadPlusWriterT.Do(func() {
		monadPlusWriterT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT1_1_0 := gopurs_runtime.Apply(Get_monadWriterT(), dictMonoid_0)
alternativeWriterT1_2_1 := gopurs_runtime.Apply(Get_alternativeWriterT(), dictMonoid_0)
return gopurs_runtime.Func(func(dictMonadPlus_3 gopurs_runtime.Value) gopurs_runtime.Value {
monadWriterT2_4_2 := gopurs_runtime.Apply(monadWriterT1_1_0, gopurs_runtime.Apply(dictMonadPlus_3.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}))
alternativeWriterT2_5_3 := gopurs_runtime.Apply(alternativeWriterT1_2_1, gopurs_runtime.Apply(dictMonadPlus_3.PtrVal.(map[string]gopurs_runtime.Value)["Alternative1"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Monad0": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return monadWriterT2_4_2
}), "Alternative1": gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return alternativeWriterT2_5_3
})})
})
})
	})
	return monadPlusWriterT
}



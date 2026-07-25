package Data_Functor_Joker

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Either "gopurs/output/Data.Either"
)

var cache_Joker gopurs_runtime.Value
var once_Joker sync.Once
func Get_Joker() gopurs_runtime.Value {
	once_Joker.Do(func() {
		cache_Joker = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Joker(x_0_box)
})
	})
	return cache_Joker
}

var cache_showJoker gopurs_runtime.Value
var once_showJoker sync.Once
func Get_showJoker() gopurs_runtime.Value {
	once_showJoker.Do(func() {
		cache_showJoker = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showJoker(dictShow_0_box)
})
	})
	return cache_showJoker
}

var cache_profunctorJoker gopurs_runtime.Value
var once_profunctorJoker sync.Once
func Get_profunctorJoker() gopurs_runtime.Value {
	once_profunctorJoker.Do(func() {
		cache_profunctorJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_profunctorJoker(dictFunctor_0_box)
})
	})
	return cache_profunctorJoker
}

var cache_ordJoker gopurs_runtime.Value
var once_ordJoker sync.Once
func Get_ordJoker() gopurs_runtime.Value {
	once_ordJoker.Do(func() {
		cache_ordJoker = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordJoker(dictOrd_0_box)
})
	})
	return cache_ordJoker
}

var cache_newtypeJoker gopurs_runtime.Value
var once_newtypeJoker sync.Once
func Get_newtypeJoker() gopurs_runtime.Value {
	once_newtypeJoker.Do(func() {
		cache_newtypeJoker = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeJoker
}

var cache_hoistJoker gopurs_runtime.Value
var once_hoistJoker sync.Once
func Get_hoistJoker() gopurs_runtime.Value {
	once_hoistJoker.Do(func() {
		cache_hoistJoker = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hoistJoker(f_0_box, v_1_box)
})
	})
	return cache_hoistJoker
}

var cache_functorJoker gopurs_runtime.Value
var once_functorJoker sync.Once
func Get_functorJoker() gopurs_runtime.Value {
	once_functorJoker.Do(func() {
		cache_functorJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorJoker(dictFunctor_0_box)
})
	})
	return cache_functorJoker
}

var cache_eqJoker gopurs_runtime.Value
var once_eqJoker sync.Once
func Get_eqJoker() gopurs_runtime.Value {
	once_eqJoker.Do(func() {
		cache_eqJoker = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqJoker(dictEq_0_box)
})
	})
	return cache_eqJoker
}

var cache_choiceJoker gopurs_runtime.Value
var once_choiceJoker sync.Once
func Get_choiceJoker() gopurs_runtime.Value {
	once_choiceJoker.Do(func() {
		cache_choiceJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_choiceJoker(dictFunctor_0_box)
})
	})
	return cache_choiceJoker
}

var cache_bifunctorJoker gopurs_runtime.Value
var once_bifunctorJoker sync.Once
func Get_bifunctorJoker() gopurs_runtime.Value {
	once_bifunctorJoker.Do(func() {
		cache_bifunctorJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bifunctorJoker(dictFunctor_0_box)
})
	})
	return cache_bifunctorJoker
}

var cache_biapplyJoker gopurs_runtime.Value
var once_biapplyJoker sync.Once
func Get_biapplyJoker() gopurs_runtime.Value {
	once_biapplyJoker.Do(func() {
		cache_biapplyJoker = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplyJoker(dictApply_0_box)
})
	})
	return cache_biapplyJoker
}

var cache_biapplicativeJoker gopurs_runtime.Value
var once_biapplicativeJoker sync.Once
func Get_biapplicativeJoker() gopurs_runtime.Value {
	once_biapplicativeJoker.Do(func() {
		cache_biapplicativeJoker = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_biapplicativeJoker(dictApplicative_0_box)
})
	})
	return cache_biapplicativeJoker
}

var cache_applyJoker gopurs_runtime.Value
var once_applyJoker sync.Once
func Get_applyJoker() gopurs_runtime.Value {
	once_applyJoker.Do(func() {
		cache_applyJoker = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applyJoker(dictApply_0_box)
})
	})
	return cache_applyJoker
}

var cache_bindJoker gopurs_runtime.Value
var once_bindJoker sync.Once
func Get_bindJoker() gopurs_runtime.Value {
	once_bindJoker.Do(func() {
		cache_bindJoker = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindJoker(dictBind_0_box)
})
	})
	return cache_bindJoker
}

var cache_applicativeJoker gopurs_runtime.Value
var once_applicativeJoker sync.Once
func Get_applicativeJoker() gopurs_runtime.Value {
	once_applicativeJoker.Do(func() {
		cache_applicativeJoker = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applicativeJoker(dictApplicative_0_box)
})
	})
	return cache_applicativeJoker
}

var cache_monadJoker gopurs_runtime.Value
var once_monadJoker sync.Once
func Get_monadJoker() gopurs_runtime.Value {
	once_monadJoker.Do(func() {
		cache_monadJoker = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monadJoker(dictMonad_0_box)
})
	})
	return cache_monadJoker
}

func Call_Joker(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showJoker(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Joker "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictShow_0.UnsafePtr)).V0, v_1), gopurs_runtime.Str(")")))
}))
}

func Call_profunctorJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, g_2, v1_3)
}))
}

func Call_ordJoker(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_hoistJoker(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_functorJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, f_1, v_2)
}))
}

func Call_eqJoker(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_choiceJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
profunctorJoker1_1_0 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, g_2, v1_3)
}))
_ = profunctorJoker1_1_0
return gopurs_runtime.RecordDict3("Profunctor0", "left", "right", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorJoker1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, pkg_Data_Either.Get_Left(), v_2)
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, pkg_Data_Either.Get_Right(), v_2)
}))
}

func Call_bifunctorJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(v_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, g_2, v1_3)
}))
}

func Call_biapplyJoker(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
bifunctorJoker1_2_1 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(v_2 gopurs_runtime.Value, g_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), g_3, v1_4)
}))
_ = bifunctorJoker1_2_1
return gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker1_2_1
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictApply_0.UnsafePtr)).V0, v_3, v1_4)
}))
}

func Call_biapplicativeJoker(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
bifunctorJoker1_3_3 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func3(func(v_3 gopurs_runtime.Value, g_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), g_4, v1_5)
}))
_ = bifunctorJoker1_3_3
biapplyJoker1_3_2 := gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker1_3_3
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), v_4, v1_5)
}))
_ = biapplyJoker1_3_2
return gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyJoker1_3_2
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, b_5)
}))
}

func Call_applyJoker(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorJoker1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, v_3)
}))
_ = functorJoker1_2_1
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_2_1
}), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictApply_0.UnsafePtr)).V0, v_3, v1_4)
}))
}

func Call_bindJoker(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorJoker1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_3, v_4)
}))
_ = functorJoker1_3_3
applyJoker1_3_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_3_3
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), v_4, v1_5)
}))
_ = applyJoker1_3_2
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_3_2
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, amb_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictBind_0.UnsafePtr)).V0, v_4, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(amb_5, x_6)
}))
}))
}

func Call_applicativeJoker(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorJoker1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_3, v_4)
}))
_ = functorJoker1_3_3
applyJoker1_3_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_3_3
}), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "apply"), v_4, v1_5)
}))
_ = applyJoker1_3_2
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_3_2
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictApplicative_0.UnsafePtr)).V0, x_4)
}))
}

func Call_monadJoker(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
functorJoker1_4_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "map"), f_4, v_5)
}))
_ = functorJoker1_4_4
applyJoker1_5_5 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_4_4
}), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "apply"), v_5, v1_6)
}))
_ = applyJoker1_5_5
applicativeJoker1_3_2 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_5_5
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "pure"), x_6)
}))
_ = applicativeJoker1_3_2
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_4_6
__local_var_5_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_4_6, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_5_7
__local_var_6_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_5_7, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_6_9
functorJoker1_7_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_9, "map"), f_7, v_8)
}))
_ = functorJoker1_7_10
applyJoker1_8_11 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_7_10
}), gopurs_runtime.Func2(func(v_8 gopurs_runtime.Value, v1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_7, "apply"), v_8, v1_9)
}))
_ = applyJoker1_8_11
bindJoker1_6_8 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_8_11
}), gopurs_runtime.Func2(func(v_9 gopurs_runtime.Value, amb_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_6, "bind"), v_9, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(amb_10, x_11)
}))
}))
_ = bindJoker1_6_8
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeJoker1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return bindJoker1_6_8
}))
}



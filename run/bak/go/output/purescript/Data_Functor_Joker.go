package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Functor_Joker_Joker gopurs_runtime.Value
var once_Data_Functor_Joker_Joker sync.Once
func Get_Data_Functor_Joker_Joker() gopurs_runtime.Value {
	once_Data_Functor_Joker_Joker.Do(func() {
		cache_Data_Functor_Joker_Joker = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_Joker(x_0_box)
})
	})
	return cache_Data_Functor_Joker_Joker
}

var cache_Data_Functor_Joker_showJoker gopurs_runtime.Value
var once_Data_Functor_Joker_showJoker sync.Once
func Get_Data_Functor_Joker_showJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_showJoker.Do(func() {
		cache_Data_Functor_Joker_showJoker = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_showJoker(dictShow_0_box)
})
	})
	return cache_Data_Functor_Joker_showJoker
}

var cache_Data_Functor_Joker_profunctorJoker gopurs_runtime.Value
var once_Data_Functor_Joker_profunctorJoker sync.Once
func Get_Data_Functor_Joker_profunctorJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_profunctorJoker.Do(func() {
		cache_Data_Functor_Joker_profunctorJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_profunctorJoker(dictFunctor_0_box)
})
	})
	return cache_Data_Functor_Joker_profunctorJoker
}

var cache_Data_Functor_Joker_ordJoker gopurs_runtime.Value
var once_Data_Functor_Joker_ordJoker sync.Once
func Get_Data_Functor_Joker_ordJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_ordJoker.Do(func() {
		cache_Data_Functor_Joker_ordJoker = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_ordJoker(dictOrd_0_box)
})
	})
	return cache_Data_Functor_Joker_ordJoker
}

var cache_Data_Functor_Joker_newtypeJoker gopurs_runtime.Value
var once_Data_Functor_Joker_newtypeJoker sync.Once
func Get_Data_Functor_Joker_newtypeJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_newtypeJoker.Do(func() {
		cache_Data_Functor_Joker_newtypeJoker = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_Functor_Joker_newtypeJoker
}

var cache_Data_Functor_Joker_hoistJoker gopurs_runtime.Value
var once_Data_Functor_Joker_hoistJoker sync.Once
func Get_Data_Functor_Joker_hoistJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_hoistJoker.Do(func() {
		cache_Data_Functor_Joker_hoistJoker = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_hoistJoker(f_0_box, v_1_box)
})
	})
	return cache_Data_Functor_Joker_hoistJoker
}

var cache_Data_Functor_Joker_functorJoker gopurs_runtime.Value
var once_Data_Functor_Joker_functorJoker sync.Once
func Get_Data_Functor_Joker_functorJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_functorJoker.Do(func() {
		cache_Data_Functor_Joker_functorJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_functorJoker(dictFunctor_0_box)
})
	})
	return cache_Data_Functor_Joker_functorJoker
}

var cache_Data_Functor_Joker_eqJoker gopurs_runtime.Value
var once_Data_Functor_Joker_eqJoker sync.Once
func Get_Data_Functor_Joker_eqJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_eqJoker.Do(func() {
		cache_Data_Functor_Joker_eqJoker = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_eqJoker(dictEq_0_box)
})
	})
	return cache_Data_Functor_Joker_eqJoker
}

var cache_Data_Functor_Joker_choiceJoker gopurs_runtime.Value
var once_Data_Functor_Joker_choiceJoker sync.Once
func Get_Data_Functor_Joker_choiceJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_choiceJoker.Do(func() {
		cache_Data_Functor_Joker_choiceJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_choiceJoker(dictFunctor_0_box)
})
	})
	return cache_Data_Functor_Joker_choiceJoker
}

var cache_Data_Functor_Joker_bifunctorJoker gopurs_runtime.Value
var once_Data_Functor_Joker_bifunctorJoker sync.Once
func Get_Data_Functor_Joker_bifunctorJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_bifunctorJoker.Do(func() {
		cache_Data_Functor_Joker_bifunctorJoker = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_bifunctorJoker(dictFunctor_0_box)
})
	})
	return cache_Data_Functor_Joker_bifunctorJoker
}

var cache_Data_Functor_Joker_biapplyJoker gopurs_runtime.Value
var once_Data_Functor_Joker_biapplyJoker sync.Once
func Get_Data_Functor_Joker_biapplyJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_biapplyJoker.Do(func() {
		cache_Data_Functor_Joker_biapplyJoker = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_biapplyJoker(dictApply_0_box)
})
	})
	return cache_Data_Functor_Joker_biapplyJoker
}

var cache_Data_Functor_Joker_biapplicativeJoker gopurs_runtime.Value
var once_Data_Functor_Joker_biapplicativeJoker sync.Once
func Get_Data_Functor_Joker_biapplicativeJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_biapplicativeJoker.Do(func() {
		cache_Data_Functor_Joker_biapplicativeJoker = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_biapplicativeJoker(dictApplicative_0_box)
})
	})
	return cache_Data_Functor_Joker_biapplicativeJoker
}

var cache_Data_Functor_Joker_applyJoker gopurs_runtime.Value
var once_Data_Functor_Joker_applyJoker sync.Once
func Get_Data_Functor_Joker_applyJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_applyJoker.Do(func() {
		cache_Data_Functor_Joker_applyJoker = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_applyJoker(dictApply_0_box)
})
	})
	return cache_Data_Functor_Joker_applyJoker
}

var cache_Data_Functor_Joker_bindJoker gopurs_runtime.Value
var once_Data_Functor_Joker_bindJoker sync.Once
func Get_Data_Functor_Joker_bindJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_bindJoker.Do(func() {
		cache_Data_Functor_Joker_bindJoker = gopurs_runtime.Func(func(dictBind_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_bindJoker(dictBind_0_box)
})
	})
	return cache_Data_Functor_Joker_bindJoker
}

var cache_Data_Functor_Joker_applicativeJoker gopurs_runtime.Value
var once_Data_Functor_Joker_applicativeJoker sync.Once
func Get_Data_Functor_Joker_applicativeJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_applicativeJoker.Do(func() {
		cache_Data_Functor_Joker_applicativeJoker = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_applicativeJoker(dictApplicative_0_box)
})
	})
	return cache_Data_Functor_Joker_applicativeJoker
}

var cache_Data_Functor_Joker_monadJoker gopurs_runtime.Value
var once_Data_Functor_Joker_monadJoker sync.Once
func Get_Data_Functor_Joker_monadJoker() gopurs_runtime.Value {
	once_Data_Functor_Joker_monadJoker.Do(func() {
		cache_Data_Functor_Joker_monadJoker = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Functor_Joker_monadJoker(dictMonad_0_box)
})
	})
	return cache_Data_Functor_Joker_monadJoker
}

func Call_Data_Functor_Joker_Joker(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Functor_Joker_showJoker(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Joker ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), v_1).StrVal())) + (")"))
}))
}

func Call_Data_Functor_Joker_profunctorJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2, v1_3)
})
})
}))
}

func Call_Data_Functor_Joker_ordJoker(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return dictOrd_0
}

func Call_Data_Functor_Joker_hoistJoker(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_Data_Functor_Joker_functorJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, v_2)
})
}))
}

func Call_Data_Functor_Joker_eqJoker(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return dictEq_0
}

func Call_Data_Functor_Joker_choiceJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
// TAST (Let): profunctorJoker1_1_0 -> gopurs_runtime.Value
profunctorJoker1_1_0 := gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2, v1_3)
})
})
}))
_ = profunctorJoker1_1_0
return gopurs_runtime.RecordDict3("Profunctor0", "left", "right", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return profunctorJoker1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), Get_Data_Either_Left(), v_2)
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), Get_Data_Either_Right(), v_2)
}))
}

func Call_Data_Functor_Joker_bifunctorJoker(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), g_2, v1_3)
})
})
}))
}

func Call_Data_Functor_Joker_biapplyJoker(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): bifunctorJoker1_1_0 -> gopurs_runtime.Value
bifunctorJoker1_1_0 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), g_3, v1_4)
})
})
}))
_ = bifunctorJoker1_1_0
return gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), v_2, v1_3)
})
}))
}

func Call_Data_Functor_Joker_biapplicativeJoker(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): bifunctorJoker1_2_2 -> gopurs_runtime.Value
bifunctorJoker1_2_2 := gopurs_runtime.RecordDict1("bimap", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), g_4, v1_5)
})
})
}))
_ = bifunctorJoker1_2_2
// TAST (Let): biapplyJoker1_1_0 -> gopurs_runtime.Value
biapplyJoker1_1_0 := gopurs_runtime.RecordDict2("Bifunctor0", "biapply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bifunctorJoker1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), v_3, v1_4)
})
}))
_ = biapplyJoker1_1_0
return gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return biapplyJoker1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), b_3)
})
}))
}

func Call_Data_Functor_Joker_applyJoker(dictApply_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApply_0 gopurs_runtime.Value = dictApply_0_loop
_ = dictApply_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorJoker1_1_0 -> gopurs_runtime.Value
functorJoker1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, v_3)
})
}))
_ = functorJoker1_1_0
return gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictApply_0, "apply"), v_2, v1_3)
})
}))
}

func Call_Data_Functor_Joker_bindJoker(dictBind_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 gopurs_runtime.Value = dictBind_0_loop
_ = dictBind_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBind_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorJoker1_2_2 -> gopurs_runtime.Value
functorJoker1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3, v_4)
})
}))
_ = functorJoker1_2_2
// TAST (Let): applyJoker1_1_0 -> gopurs_runtime.Value
applyJoker1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), v_3, v1_4)
})
}))
_ = applyJoker1_1_0
return gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(amb_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictBind_0, "bind"), v_2, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(amb_3, x_4)
}))
})
}))
}

func Call_Data_Functor_Joker_applicativeJoker(dictApplicative_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorJoker1_2_2 -> gopurs_runtime.Value
functorJoker1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3, v_4)
})
}))
_ = functorJoker1_2_2
// TAST (Let): applyJoker1_1_0 -> gopurs_runtime.Value
applyJoker1_1_0 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_2_2
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "apply"), v_3, v1_4)
})
}))
_ = applyJoker1_1_0
return gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), x_2)
}))
}

func Call_Data_Functor_Joker_monadJoker(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorJoker1_3_4 -> gopurs_runtime.Value
functorJoker1_3_4 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "map"), f_4, v_5)
})
}))
_ = functorJoker1_3_4
// TAST (Let): applyJoker1_2_2 -> gopurs_runtime.Value
applyJoker1_2_2 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_3_4
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "apply"), v_4, v1_5)
})
}))
_ = applyJoker1_2_2
// TAST (Let): applicativeJoker1_1_0 -> gopurs_runtime.Value
applicativeJoker1_1_0 := gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_2_2
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "pure"), x_3)
}))
_ = applicativeJoker1_1_0
// TAST (Let): __local_var_2_7 -> gopurs_runtime.Value
__local_var_2_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{})
_ = __local_var_2_7
// TAST (Let): __local_var_3_9 -> gopurs_runtime.Value
__local_var_3_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_7, "Apply0"), gopurs_runtime.Value{})
_ = __local_var_3_9
// TAST (Let): __local_var_4_11 -> gopurs_runtime.Value
__local_var_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_9, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_11
// TAST (Let): functorJoker1_4_10 -> gopurs_runtime.Value
functorJoker1_4_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_11, "map"), f_5, v_6)
})
}))
_ = functorJoker1_4_10
// TAST (Let): applyJoker1_3_8 -> gopurs_runtime.Value
applyJoker1_3_8 := gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorJoker1_4_10
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_9, "apply"), v_5, v1_6)
})
}))
_ = applyJoker1_3_8
// TAST (Let): bindJoker1_2_6 -> gopurs_runtime.Value
bindJoker1_2_6 := gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return applyJoker1_3_8
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(amb_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_7, "bind"), v_4, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(amb_5, x_6)
}))
})
}))
_ = bindJoker1_2_6
return gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return applicativeJoker1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return bindJoker1_2_6
}))
}



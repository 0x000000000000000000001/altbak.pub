package Control_Comonad_Traced_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var TracedT gopurs_runtime.Value
var once_TracedT sync.Once
func Get_TracedT() gopurs_runtime.Value {
	once_TracedT.Do(func() {
		TracedT = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return TracedT
}

var runTracedT gopurs_runtime.Value
var once_runTracedT sync.Once
func Get_runTracedT() gopurs_runtime.Value {
	once_runTracedT.Do(func() {
		runTracedT = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runTracedT
}

var newtypeTracedT gopurs_runtime.Value
var once_newtypeTracedT sync.Once
func Get_newtypeTracedT() gopurs_runtime.Value {
	once_newtypeTracedT.Do(func() {
		newtypeTracedT = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"Coercible0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})
	})
	return newtypeTracedT
}

var functorTracedT gopurs_runtime.Value
var once_functorTracedT sync.Once
func Get_functorTracedT() gopurs_runtime.Value {
	once_functorTracedT.Do(func() {
		functorTracedT = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(g_3, t_4))
})
})), v_2)
})
})})
})
	})
	return functorTracedT
}

var extendTracedT gopurs_runtime.Value
var once_extendTracedT sync.Once
func Get_extendTracedT() gopurs_runtime.Value {
	once_extendTracedT.Do(func() {
		extendTracedT = gopurs_runtime.Func(func(dictExtend_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorTracedT1_2_1 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(g_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(g_4, t_5))
})
})), v_3)
})
})})
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictExtend_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], gopurs_runtime.Func(func(w_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(h_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_8, gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_3.PtrVal.(map[string]gopurs_runtime.Value)["append"], t_7), t_prime_9))
})
})), w_prime_6))
})
})), v_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorTracedT1_2_1
})})
})
})
	})
	return extendTracedT
}

var comonadTransTracedT gopurs_runtime.Value
var once_comonadTransTracedT sync.Once
func Get_comonadTransTracedT() gopurs_runtime.Value {
	once_comonadTransTracedT.Do(func() {
		comonadTransTracedT = gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_1_0 := dictMonoid_0.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"lower": gopurs_runtime.Func(func(dictComonad_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonad_2.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, mempty_1_0)
})), v_3)
})
})})
})
	})
	return comonadTransTracedT
}

var comonadTracedT gopurs_runtime.Value
var once_comonadTracedT sync.Once
func Get_comonadTracedT() gopurs_runtime.Value {
	once_comonadTracedT.Do(func() {
		comonadTracedT = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{})
Functor0_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
functorTracedT1_3_3 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"map": gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Apply(g_5, t_6))
})
})), v_4)
})
})})
extendTracedT1_3_2 := gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["extend"], gopurs_runtime.Func(func(w_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Apply(gopurs_runtime.Apply(Functor0_2_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(h_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_9, gopurs_runtime.Apply(gopurs_runtime.Apply(dictSemigroup_4.PtrVal.(map[string]gopurs_runtime.Value)["append"], t_8), t_prime_10))
})
})), w_prime_7))
})
})), v_6)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorTracedT1_3_3
})})
})
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_5_4 := dictMonoid_4.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
extendTracedT2_6_5 := gopurs_runtime.Apply(extendTracedT1_3_2, gopurs_runtime.Apply(dictMonoid_4.PtrVal.(map[string]gopurs_runtime.Value)["Semigroup0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extract": gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["extract"], v_7), mempty_5_4)
}), "Extend0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return extendTracedT2_6_5
})})
})
})
	})
	return comonadTracedT
}



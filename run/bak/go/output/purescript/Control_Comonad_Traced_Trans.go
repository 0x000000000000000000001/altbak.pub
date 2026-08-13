package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Control_Comonad_Traced_Trans_TracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_TracedT sync.Once
func Get_Control_Comonad_Traced_Trans_TracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_TracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_TracedT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_TracedT(x_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_TracedT
}

var cache_Control_Comonad_Traced_Trans_runTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_runTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_runTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_runTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_runTracedT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_runTracedT(v_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_runTracedT
}

var cache_Control_Comonad_Traced_Trans_newtypeTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_newtypeTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_newtypeTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_newtypeTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_newtypeTracedT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Control_Comonad_Traced_Trans_newtypeTracedT
}

var cache_Control_Comonad_Traced_Trans_functorTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_functorTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_functorTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_functorTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_functorTracedT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_functorTracedT(dictFunctor_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_functorTracedT
}

var cache_Control_Comonad_Traced_Trans_extendTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_extendTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_extendTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_extendTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_extendTracedT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_extendTracedT(dictExtend_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_extendTracedT
}

var cache_Control_Comonad_Traced_Trans_comonadTransTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_comonadTransTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_comonadTransTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_comonadTransTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_comonadTransTracedT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_comonadTransTracedT(dictMonoid_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_comonadTransTracedT
}

var cache_Control_Comonad_Traced_Trans_comonadTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Trans_comonadTracedT sync.Once
func Get_Control_Comonad_Traced_Trans_comonadTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Trans_comonadTracedT.Do(func() {
		cache_Control_Comonad_Traced_Trans_comonadTracedT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Trans_comonadTracedT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Traced_Trans_comonadTracedT
}

func Call_Control_Comonad_Traced_Trans_TracedT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Traced_Trans_runTracedT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Comonad_Traced_Trans_functorTracedT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(g_3, t_4))
})
}), v_2)
})
}))
}

func Call_Control_Comonad_Traced_Trans_extendTracedT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorTracedT1_2_1 -> gopurs_runtime.Value
functorTracedT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), gopurs_runtime.Func(func(g_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Apply(g_5, t_6))
})
}), v_4)
})
}))
_ = functorTracedT1_2_1
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorTracedT1_2_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(w_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(h_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_8, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_3, "append"), t_7, t_prime_9))
})
}), w_prime_6))
})
}), v_5)
})
}))
})
}

func Call_Control_Comonad_Traced_Trans_comonadTransTracedT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func(func(dictComonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_2_0 -> *Constructor_Data_Functor_Functor
Functor0_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_1, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_0.V0), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
}), v_3)
})
}))
}

func Call_Control_Comonad_Traced_Trans_comonadTracedT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorTracedT1_3_3 -> gopurs_runtime.Value
functorTracedT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), gopurs_runtime.Func(func(g_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Apply(g_6, t_7))
})
}), v_5)
})
}))
_ = functorTracedT1_3_3
// TAST (Let): extendTracedT1_3_2 -> gopurs_runtime.Value
extendTracedT1_3_2 := gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorTracedT1_3_3
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "extend"), gopurs_runtime.Func(func(w_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), gopurs_runtime.Func(func(h_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_9, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_4, "append"), t_8, t_prime_10))
})
}), w_prime_7))
})
}), v_6)
})
}))
})
_ = extendTracedT1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): extendTracedT2_5_5 -> gopurs_runtime.Value
extendTracedT2_5_5 := gopurs_runtime.Apply(extendTracedT1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = extendTracedT2_5_5
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return extendTracedT2_5_5
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), v_6, gopurs_runtime.RecordGet(dictMonoid_4, "mempty"))
}))
})
}



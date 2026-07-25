package Control_Comonad_Traced_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_TracedT gopurs_runtime.Value
var once_TracedT sync.Once
func Get_TracedT() gopurs_runtime.Value {
	once_TracedT.Do(func() {
		cache_TracedT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_TracedT(x_0_box)
})
	})
	return cache_TracedT
}

var cache_runTracedT gopurs_runtime.Value
var once_runTracedT sync.Once
func Get_runTracedT() gopurs_runtime.Value {
	once_runTracedT.Do(func() {
		cache_runTracedT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runTracedT(v_0_box)
})
	})
	return cache_runTracedT
}

var cache_newtypeTracedT gopurs_runtime.Value
var once_newtypeTracedT sync.Once
func Get_newtypeTracedT() gopurs_runtime.Value {
	once_newtypeTracedT.Do(func() {
		cache_newtypeTracedT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeTracedT
}

var cache_functorTracedT gopurs_runtime.Value
var once_functorTracedT sync.Once
func Get_functorTracedT() gopurs_runtime.Value {
	once_functorTracedT.Do(func() {
		cache_functorTracedT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorTracedT(dictFunctor_0_box)
})
	})
	return cache_functorTracedT
}

var cache_extendTracedT gopurs_runtime.Value
var once_extendTracedT sync.Once
func Get_extendTracedT() gopurs_runtime.Value {
	once_extendTracedT.Do(func() {
		cache_extendTracedT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extendTracedT(dictExtend_0_box)
})
	})
	return cache_extendTracedT
}

var cache_comonadTransTracedT gopurs_runtime.Value
var once_comonadTransTracedT sync.Once
func Get_comonadTransTracedT() gopurs_runtime.Value {
	once_comonadTransTracedT.Do(func() {
		cache_comonadTransTracedT = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadTransTracedT(dictMonoid_0_box)
})
	})
	return cache_comonadTransTracedT
}

var cache_comonadTracedT gopurs_runtime.Value
var once_comonadTracedT sync.Once
func Get_comonadTracedT() gopurs_runtime.Value {
	once_comonadTracedT.Do(func() {
		cache_comonadTracedT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadTracedT(dictComonad_0_box)
})
	})
	return cache_comonadTracedT
}

func Call_TracedT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_runTracedT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_functorTracedT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func2(func(g_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(g_3, t_4))
}), v_2)
}))
}

func Call_extendTracedT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0_NOT_FOUND"), gopurs_runtime.Value{})
_ = Functor0_1_0
functorTracedT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.Func2(func(g_4 gopurs_runtime.Value, t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(g_4, t_5))
}), v_3)
}))
_ = functorTracedT1_2_1
return gopurs_runtime.Func(func(dictSemigroup_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorTracedT1_2_1
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictExtend_0.UnsafePtr)).V0, gopurs_runtime.Func2(func(w_prime_6 gopurs_runtime.Value, t_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), gopurs_runtime.Func2(func(h_8 gopurs_runtime.Value, t_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_8, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_3, "append"), t_7, t_prime_9))
}), w_prime_6))
}), v_5)
}))
})
}

func Call_comonadTransTracedT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := ((*gopurs_runtime.RecordData1)(dictMonoid_0.UnsafePtr)).V0
_ = mempty_1_0
return gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func2(func(dictComonad_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_2, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, mempty_1_0)
}), v_3)
}))
}

func Call_comonadTracedT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
Functor0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_2_1
functorTracedT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Func2(func(g_5 gopurs_runtime.Value, t_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Apply(g_5, t_6))
}), v_4)
}))
_ = functorTracedT1_3_3
extendTracedT1_3_2 := gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorTracedT1_3_3
}), gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "extend"), gopurs_runtime.Func2(func(w_prime_7 gopurs_runtime.Value, t_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_2_1, "map"), gopurs_runtime.Func2(func(h_9 gopurs_runtime.Value, t_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_9, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_4, "append"), t_8, t_prime_10))
}), w_prime_7))
}), v_6)
}))
})
_ = extendTracedT1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_5_4 := gopurs_runtime.RecordGet(dictMonoid_4, "mempty")
_ = mempty_5_4
extendTracedT2_6_5 := gopurs_runtime.Apply(extendTracedT1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = extendTracedT2_6_5
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return extendTracedT2_6_5
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictComonad_0.UnsafePtr)).V0, v_7, mempty_5_4)
}))
})
}



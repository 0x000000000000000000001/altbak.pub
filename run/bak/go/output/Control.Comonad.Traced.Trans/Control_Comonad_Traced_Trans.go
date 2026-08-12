package Control_Comonad_Traced_Trans

import (
	pkg_Control_Comonad "gopurs/output/Control.Comonad"
	pkg_Control_Extend "gopurs/output/Control.Extend"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
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

var cache_extract__1031647521 gopurs_runtime.Value
var once_extract__1031647521 sync.Once
func Get_extract__1031647521() gopurs_runtime.Value {
	once_extract__1031647521.Do(func() {
		cache_extract__1031647521 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extract__1031647521(gopurs_runtime.CoerceToStruct[pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extract__1031647521
}

var cache_extend__1264481661 gopurs_runtime.Value
var once_extend__1264481661 sync.Once
func Get_extend__1264481661() gopurs_runtime.Value {
	once_extend__1264481661.Do(func() {
		cache_extend__1264481661 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend__1264481661(gopurs_runtime.CoerceToStruct[pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extend__1264481661
}

var cache_extend__1965081501 gopurs_runtime.Value
var once_extend__1965081501 sync.Once
func Get_extend__1965081501() gopurs_runtime.Value {
	once_extend__1965081501.Do(func() {
		cache_extend__1965081501 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend__1965081501(gopurs_runtime.CoerceToStruct[pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extend__1965081501
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__1938733460 gopurs_runtime.Value
var once_map__1938733460 sync.Once
func Get_map__1938733460() gopurs_runtime.Value {
	once_map__1938733460.Do(func() {
		cache_map__1938733460 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1938733460(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1938733460
}

var cache_map__3897763604 gopurs_runtime.Value
var once_map__3897763604 sync.Once
func Get_map__3897763604() gopurs_runtime.Value {
	once_map__3897763604.Do(func() {
		cache_map__3897763604 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3897763604(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3897763604
}

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
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

func Call_extendTracedT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
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
return gopurs_runtime.Apply(f_4, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(h_8 gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_comonadTransTracedT(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func(func(dictComonad_1 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_1, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_0
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_0.V0, gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.RecordGet(dictMonoid_0, "mempty"))
}), v_3)
})
}))
}

func Call_comonadTracedT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_0
Functor0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
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
extendTracedT1_3_2 := gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorTracedT1_3_3
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "extend"), gopurs_runtime.Func(func(w_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Apply2(Functor0_2_1.V0, gopurs_runtime.Func(func(h_9 gopurs_runtime.Value) gopurs_runtime.Value {
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
extendTracedT2_5_5 := gopurs_runtime.Apply(extendTracedT1_3_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}))
_ = extendTracedT2_5_5
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return extendTracedT2_5_5
}), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), v_6, gopurs_runtime.RecordGet(dictMonoid_4, "mempty"))
}))
})
}

func Call_extract__1031647521(dict_0_loop *pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_extend__1264481661(dict_0_loop *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_extend__1965081501(dict_0_loop *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1938733460(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3897763604(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}



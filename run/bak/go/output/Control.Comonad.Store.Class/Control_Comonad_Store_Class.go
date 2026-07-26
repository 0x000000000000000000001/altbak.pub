package Control_Comonad_Store_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Extend "gopurs/output/Control.Extend"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
	pkg_Control_Comonad_Store_Trans "gopurs/output/Control.Comonad.Store.Trans"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
)

var cache_pos gopurs_runtime.Value
var once_pos sync.Once
func Get_pos() gopurs_runtime.Value {
	once_pos.Do(func() {
		cache_pos = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pos(dict_0_box)
})
	})
	return cache_pos
}

var cache_peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		cache_peek = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peek(dict_0_box)
})
	})
	return cache_peek
}

var cache_peeks gopurs_runtime.Value
var once_peeks sync.Once
func Get_peeks() gopurs_runtime.Value {
	once_peeks.Do(func() {
		cache_peeks = gopurs_runtime.Func3(func(dictComonadStore_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peeks(dictComonadStore_0_box, f_1_box, x_2_box)
})
	})
	return cache_peeks
}

var cache_seeks gopurs_runtime.Value
var once_seeks sync.Once
func Get_seeks() gopurs_runtime.Value {
	once_seeks.Do(func() {
		cache_seeks = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_seeks(dictComonadStore_0_box)
})
	})
	return cache_seeks
}

var cache_seek gopurs_runtime.Value
var once_seek sync.Once
func Get_seek() gopurs_runtime.Value {
	once_seek.Do(func() {
		cache_seek = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_seek(dictComonadStore_0_box)
})
	})
	return cache_seek
}

var cache_experiment gopurs_runtime.Value
var once_experiment sync.Once
func Get_experiment() gopurs_runtime.Value {
	once_experiment.Do(func() {
		cache_experiment = gopurs_runtime.Func4(func(dictComonadStore_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_experiment(dictComonadStore_0_box, dictFunctor_1_box, f_2_box, x_3_box)
})
	})
	return cache_experiment
}

var cache_comonadStoreTracedT gopurs_runtime.Value
var once_comonadStoreTracedT sync.Once
func Get_comonadStoreTracedT() gopurs_runtime.Value {
	once_comonadStoreTracedT.Do(func() {
		cache_comonadStoreTracedT = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadStoreTracedT(dictComonadStore_0_box)
})
	})
	return cache_comonadStoreTracedT
}

var cache_comonadStoreStoreT gopurs_runtime.Value
var once_comonadStoreStoreT sync.Once
func Get_comonadStoreStoreT() gopurs_runtime.Value {
	once_comonadStoreStoreT.Do(func() {
		cache_comonadStoreStoreT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadStoreStoreT(dictComonad_0_box)
})
	})
	return cache_comonadStoreStoreT
}

var cache_comonadStoreEnvT gopurs_runtime.Value
var once_comonadStoreEnvT sync.Once
func Get_comonadStoreEnvT() gopurs_runtime.Value {
	once_comonadStoreEnvT.Do(func() {
		cache_comonadStoreEnvT = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadStoreEnvT(dictComonadStore_0_box)
})
	})
	return cache_comonadStoreEnvT
}

func Call_pos(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData2)(dict_0.UnsafePtr)).V1
}

func Call_peek(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData2)(dict_0.UnsafePtr)).V0
}

func Call_peeks(dictComonadStore_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V0, gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V1, x_2)), x_2)
}

func Call_seeks(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), pkg_Control_Extend.Get_identity())
_ = duplicate_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(duplicate_1_0, x_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V0, gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V1, __local_var_4_1)), __local_var_4_1)
})
}

func Call_seek(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), pkg_Control_Extend.Get_identity())
_ = duplicate_1_0
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V0, s_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(duplicate_1_0, x_4))
})
})
}

func Call_experiment(dictComonadStore_0_loop gopurs_runtime.Value, dictFunctor_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
var dictFunctor_1 gopurs_runtime.Value = dictFunctor_1_loop
_ = dictFunctor_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_1.UnsafePtr)).V0, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V0, a_4, x_3)
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V1, x_3)))
}

func Call_comonadStoreTracedT(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
comonadTracedT_2_1 := gopurs_runtime.Apply(pkg_Control_Comonad_Traced_Trans.Get_comonadTracedT(), Comonad0_1_0)
_ = comonadTracedT_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_4_2 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
_ = mempty_4_2
lower1_5_3 := gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, mempty_4_2)
}), v_5)
})
_ = lower1_5_3
comonadTracedT1_6_4 := gopurs_runtime.Apply(comonadTracedT_2_1, dictMonoid_3)
_ = comonadTracedT1_6_4
return gopurs_runtime.RecordDict3("Comonad0", "peek", "pos", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_6_4
}), gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_5 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V0, s_7)
_ = __local_var_8_5
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.UncurriedApp(lower1_5_3, x_9))
})
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V1, gopurs_runtime.UncurriedApp(lower1_5_3, x_7))
}))
})
}

func Call_comonadStoreStoreT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
comonadStoreT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Store_Trans.Get_comonadStoreT(), dictComonad_0)
_ = comonadStoreT_1_0
return gopurs_runtime.RecordDict3("Comonad0", "peek", "pos", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_1_0
}), gopurs_runtime.Func2(func(s_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictComonad_0.UnsafePtr)).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_2)
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
}))
}

func Call_comonadStoreEnvT(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
lower1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Comonad_Env_Trans.Get_comonadTransEnvT(), "lower"), Comonad0_1_0)
_ = lower1_2_1
comonadEnvT_3_2 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), Comonad0_1_0)
_ = comonadEnvT_3_2
return gopurs_runtime.RecordDict3("Comonad0", "peek", "pos", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_3_2
}), gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V0, s_4)
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(lower1_2_1, x_6))
})
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData2)(dictComonadStore_0.UnsafePtr)).V1, gopurs_runtime.Apply(lower1_2_1, x_4))
}))
}



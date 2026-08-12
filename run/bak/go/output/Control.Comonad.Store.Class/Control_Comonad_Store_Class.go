package Control_Comonad_Store_Class

import (
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
	pkg_Control_Comonad_Store_Trans "gopurs/output/Control.Comonad.Store.Trans"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
	pkg_Control_Extend "gopurs/output/Control.Extend"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_pos gopurs_runtime.Value
var once_pos sync.Once
func Get_pos() gopurs_runtime.Value {
	once_pos.Do(func() {
		cache_pos = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pos(gopurs_runtime.CoerceToStruct[Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pos
}

var cache_pos__gopurs_runtime_Value_639385354 gopurs_runtime.Value
var once_pos__gopurs_runtime_Value_639385354 sync.Once
func Get_pos__gopurs_runtime_Value_639385354() gopurs_runtime.Value {
	once_pos__gopurs_runtime_Value_639385354.Do(func() {
		cache_pos__gopurs_runtime_Value_639385354 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pos__gopurs_runtime_Value_639385354(gopurs_runtime.CoerceToStruct[Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pos__gopurs_runtime_Value_639385354
}

var cache_peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		cache_peek = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peek(gopurs_runtime.CoerceToStruct[Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_peek
}

var cache_peek__gopurs_runtime_Value_413153475 gopurs_runtime.Value
var once_peek__gopurs_runtime_Value_413153475 sync.Once
func Get_peek__gopurs_runtime_Value_413153475() gopurs_runtime.Value {
	once_peek__gopurs_runtime_Value_413153475.Do(func() {
		cache_peek__gopurs_runtime_Value_413153475 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peek__gopurs_runtime_Value_413153475(gopurs_runtime.CoerceToStruct[Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_peek__gopurs_runtime_Value_413153475
}

var cache_peeks gopurs_runtime.Value
var once_peeks sync.Once
func Get_peeks() gopurs_runtime.Value {
	once_peeks.Do(func() {
		cache_peeks = gopurs_runtime.Func3(func(dictComonadStore_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peeks(gopurs_runtime.CoerceToStruct[Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box), f_1_box, x_2_box)
})
	})
	return cache_peeks
}

var cache_peeks__gopurs_runtime_Value_2913274603 gopurs_runtime.Value
var once_peeks__gopurs_runtime_Value_2913274603 sync.Once
func Get_peeks__gopurs_runtime_Value_2913274603() gopurs_runtime.Value {
	once_peeks__gopurs_runtime_Value_2913274603.Do(func() {
		cache_peeks__gopurs_runtime_Value_2913274603 = gopurs_runtime.Func3(func(dictComonadStore_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_peeks__gopurs_runtime_Value_2913274603(gopurs_runtime.CoerceToStruct[Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box), f_1_box, x_2_box)
})
	})
	return cache_peeks__gopurs_runtime_Value_2913274603
}

var cache_seeks gopurs_runtime.Value
var once_seeks sync.Once
func Get_seeks() gopurs_runtime.Value {
	once_seeks.Do(func() {
		cache_seeks = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_seeks(gopurs_runtime.CoerceToStruct[Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box))
})
	})
	return cache_seeks
}

var cache_seek gopurs_runtime.Value
var once_seek sync.Once
func Get_seek() gopurs_runtime.Value {
	once_seek.Do(func() {
		cache_seek = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_seek(gopurs_runtime.CoerceToStruct[Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box))
})
	})
	return cache_seek
}

var cache_experiment gopurs_runtime.Value
var once_experiment sync.Once
func Get_experiment() gopurs_runtime.Value {
	once_experiment.Do(func() {
		cache_experiment = gopurs_runtime.Func4(func(dictComonadStore_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_experiment(gopurs_runtime.CoerceToStruct[Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), f_2_box, x_3_box)
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

type Constructor_ComonadStore[T_s any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4264314723] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "Comonad0": return c.V0
		case "peek": return c.V1
		case "pos": return c.V2
		default: panic("Key not found in dictionary Constructor_ComonadStore: " + key)
		}
	}
}


func Call_pos(dict_0_loop *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_pos__gopurs_runtime_Value_639385354(dict_0_loop *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_peek(dict_0_loop *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_peek__gopurs_runtime_Value_413153475(dict_0_loop *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_peeks(dictComonadStore_0_loop *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(dictComonadStore_0.V1, gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(dictComonadStore_0.V2, x_2)), x_2)
}

func Call_peeks__gopurs_runtime_Value_2913274603(dictComonadStore_0_loop *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(dictComonadStore_0.V1, gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(dictComonadStore_0.V2, x_2)), x_2)
}

func Call_seeks(dictComonadStore_0_loop *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictComonadStore_0.V0, gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), pkg_Control_Extend.Get_identity())
_ = duplicate_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(duplicate_1_0, x_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2(dictComonadStore_0.V1, gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(dictComonadStore_0.V2, __local_var_4_1)), __local_var_4_1)
})
})
}

func Call_seek(dictComonadStore_0_loop *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictComonadStore_0.V0, gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), pkg_Control_Extend.Get_identity())
_ = duplicate_1_0
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictComonadStore_0.V1, s_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(duplicate_1_0, x_4))
})
})
}

func Call_experiment(dictComonadStore_0_loop *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value], dictFunctor_1_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
var dictFunctor_1 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
return gopurs_runtime.Apply2(dictFunctor_1.V0, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictComonadStore_0.V1, a_4, x_3)
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(dictComonadStore_0.V2, x_3)))
}

func Call_comonadStoreTracedT(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
comonadTracedT_2_1 := gopurs_runtime.Apply(pkg_Control_Comonad_Traced_Trans.Get_comonadTracedT(), Comonad0_1_0)
_ = comonadTracedT_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_3 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
lower1_4_2 := gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_3.V0, gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))
}), v_5)
})
_ = lower1_4_2
comonadTracedT1_5_4 := gopurs_runtime.Apply(comonadTracedT_2_1, dictMonoid_3)
_ = comonadTracedT1_5_4
return gopurs_runtime.RecordDict3("Comonad0", "peek", "pos", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_5_4
}), gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_6)
_ = __local_var_7_5
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Apply(lower1_4_2, x_8))
})
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), gopurs_runtime.Apply(lower1_4_2, x_6))
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
}), gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_2)
})
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
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_4)
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(lower1_2_1, x_6))
})
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), gopurs_runtime.Apply(lower1_2_1, x_4))
}))
}



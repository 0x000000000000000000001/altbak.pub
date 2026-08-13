package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Store_Class_lower gopurs_runtime.Value
var once_Control_Comonad_Store_Class_lower sync.Once
func Get_Control_Comonad_Store_Class_lower() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_lower.Do(func() {
		cache_Control_Comonad_Store_Class_lower = gopurs_runtime.RecordGet(Get_Control_Comonad_Env_Trans_comonadTransEnvT(), "lower")
	})
	return cache_Control_Comonad_Store_Class_lower
}

var cache_Control_Comonad_Store_Class_ComonadStore_dollarDict gopurs_runtime.Value
var once_Control_Comonad_Store_Class_ComonadStore_dollarDict sync.Once
func Get_Control_Comonad_Store_Class_ComonadStore_dollarDict() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_ComonadStore_dollarDict.Do(func() {
		cache_Control_Comonad_Store_Class_ComonadStore_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_ComonadStore_dollarDict(x_0_box)
})
	})
	return cache_Control_Comonad_Store_Class_ComonadStore_dollarDict
}

var cache_Control_Comonad_Store_Class_pos gopurs_runtime.Value
var once_Control_Comonad_Store_Class_pos sync.Once
func Get_Control_Comonad_Store_Class_pos() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_pos.Do(func() {
		cache_Control_Comonad_Store_Class_pos = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_pos(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Store_Class_pos
}

var cache_Control_Comonad_Store_Class_peek gopurs_runtime.Value
var once_Control_Comonad_Store_Class_peek sync.Once
func Get_Control_Comonad_Store_Class_peek() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_peek.Do(func() {
		cache_Control_Comonad_Store_Class_peek = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_peek(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Store_Class_peek
}

var cache_Control_Comonad_Store_Class_peeks gopurs_runtime.Value
var once_Control_Comonad_Store_Class_peeks sync.Once
func Get_Control_Comonad_Store_Class_peeks() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_peeks.Do(func() {
		cache_Control_Comonad_Store_Class_peeks = gopurs_runtime.Func3(func(dictComonadStore_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_peeks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box), f_1_box, x_2_box)
})
	})
	return cache_Control_Comonad_Store_Class_peeks
}

var cache_Control_Comonad_Store_Class_seeks gopurs_runtime.Value
var once_Control_Comonad_Store_Class_seeks sync.Once
func Get_Control_Comonad_Store_Class_seeks() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_seeks.Do(func() {
		cache_Control_Comonad_Store_Class_seeks = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_seeks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box))
})
	})
	return cache_Control_Comonad_Store_Class_seeks
}

var cache_Control_Comonad_Store_Class_seek gopurs_runtime.Value
var once_Control_Comonad_Store_Class_seek sync.Once
func Get_Control_Comonad_Store_Class_seek() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_seek.Do(func() {
		cache_Control_Comonad_Store_Class_seek = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_seek(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box))
})
	})
	return cache_Control_Comonad_Store_Class_seek
}

var cache_Control_Comonad_Store_Class_experiment gopurs_runtime.Value
var once_Control_Comonad_Store_Class_experiment sync.Once
func Get_Control_Comonad_Store_Class_experiment() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_experiment.Do(func() {
		cache_Control_Comonad_Store_Class_experiment = gopurs_runtime.Func4(func(dictComonadStore_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, x_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_experiment(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box), gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), f_2_box, x_3_box)
})
	})
	return cache_Control_Comonad_Store_Class_experiment
}

var cache_Control_Comonad_Store_Class_comonadStoreTracedT gopurs_runtime.Value
var once_Control_Comonad_Store_Class_comonadStoreTracedT sync.Once
func Get_Control_Comonad_Store_Class_comonadStoreTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_comonadStoreTracedT.Do(func() {
		cache_Control_Comonad_Store_Class_comonadStoreTracedT = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_comonadStoreTracedT(dictComonadStore_0_box)
})
	})
	return cache_Control_Comonad_Store_Class_comonadStoreTracedT
}

var cache_Control_Comonad_Store_Class_comonadStoreStoreT gopurs_runtime.Value
var once_Control_Comonad_Store_Class_comonadStoreStoreT sync.Once
func Get_Control_Comonad_Store_Class_comonadStoreStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_comonadStoreStoreT.Do(func() {
		cache_Control_Comonad_Store_Class_comonadStoreStoreT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_comonadStoreStoreT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Store_Class_comonadStoreStoreT
}

var cache_Control_Comonad_Store_Class_comonadStoreEnvT gopurs_runtime.Value
var once_Control_Comonad_Store_Class_comonadStoreEnvT sync.Once
func Get_Control_Comonad_Store_Class_comonadStoreEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_comonadStoreEnvT.Do(func() {
		cache_Control_Comonad_Store_Class_comonadStoreEnvT = gopurs_runtime.Func(func(dictComonadStore_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_comonadStoreEnvT(dictComonadStore_0_box)
})
	})
	return cache_Control_Comonad_Store_Class_comonadStoreEnvT
}

var cache_Control_Comonad_Store_Class_peek__413153475 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_peek__413153475 sync.Once
func Get_Control_Comonad_Store_Class_peek__413153475() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_peek__413153475.Do(func() {
		cache_Control_Comonad_Store_Class_peek__413153475 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_peek__413153475(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Store_Class_peek__413153475
}

var cache_Control_Comonad_Store_Class_peek__3957251203 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_peek__3957251203 sync.Once
func Get_Control_Comonad_Store_Class_peek__3957251203() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_peek__3957251203.Do(func() {
		cache_Control_Comonad_Store_Class_peek__3957251203 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_peek__3957251203(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Store_Class_peek__3957251203
}

var cache_Control_Comonad_Store_Class_peeks__2913274603 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_peeks__2913274603 sync.Once
func Get_Control_Comonad_Store_Class_peeks__2913274603() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_peeks__2913274603.Do(func() {
		cache_Control_Comonad_Store_Class_peeks__2913274603 = gopurs_runtime.Func3(func(dictComonadStore_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_peeks__2913274603(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box), f_1_box, x_2_box)
})
	})
	return cache_Control_Comonad_Store_Class_peeks__2913274603
}

var cache_Control_Comonad_Store_Class_peeks__110217771 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_peeks__110217771 sync.Once
func Get_Control_Comonad_Store_Class_peeks__110217771() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_peeks__110217771.Do(func() {
		cache_Control_Comonad_Store_Class_peeks__110217771 = gopurs_runtime.Func3(func(dictComonadStore_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_peeks__110217771(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadStore_0_box), f_1_box, x_2_box)
})
	})
	return cache_Control_Comonad_Store_Class_peeks__110217771
}

var cache_Control_Comonad_Store_Class_pos__639385354 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_pos__639385354 sync.Once
func Get_Control_Comonad_Store_Class_pos__639385354() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_pos__639385354.Do(func() {
		cache_Control_Comonad_Store_Class_pos__639385354 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_pos__639385354(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Store_Class_pos__639385354
}

var cache_Control_Comonad_Store_Class_pos__3752461637 gopurs_runtime.Value
var once_Control_Comonad_Store_Class_pos__3752461637 sync.Once
func Get_Control_Comonad_Store_Class_pos__3752461637() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_pos__3752461637.Do(func() {
		cache_Control_Comonad_Store_Class_pos__3752461637 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Store_Class_pos__3752461637(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Store_Class_pos__3752461637
}

type Constructor_Control_Comonad_Store_Class_ComonadStore[T_s any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4264314723] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Comonad0": return gopurs_runtime.Box(c.V0)
		case "peek": return gopurs_runtime.Box(c.V1)
		case "pos": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Store_Class_ComonadStore: " + key)
		}
	}
}


func Call_Control_Comonad_Store_Class_ComonadStore_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Store_Class_pos(dict_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Comonad_Store_Class_peek(dict_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Store_Class_peeks(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictComonadStore_0.V1), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V2), x_2)), x_2)
}

func Call_Control_Comonad_Store_Class_seeks(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
// TAST (Let): duplicate_1_0 -> gopurs_runtime.Value
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V0), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), Get_Control_Extend_identity())
_ = duplicate_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(duplicate_1_0, x_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictComonadStore_0.V1), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V2), __local_var_4_1)), __local_var_4_1)
})
})
}

func Call_Control_Comonad_Store_Class_seek(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
// TAST (Let): duplicate_1_0 -> gopurs_runtime.Value
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V0), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), Get_Control_Extend_identity())
_ = duplicate_1_0
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V1), s_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(duplicate_1_0, x_4))
})
})
}

func Call_Control_Comonad_Store_Class_experiment(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value], dictFunctor_1_loop *Constructor_Data_Functor_Functor[gopurs_runtime.Value], f_2_loop gopurs_runtime.Value, x_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
var dictFunctor_1 *Constructor_Data_Functor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var x_3 gopurs_runtime.Value = x_3_loop
_ = x_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_1.V0), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictComonadStore_0.V1), a_4, x_3)
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V2), x_3)))
}

func Call_Control_Comonad_Store_Class_comonadStoreTracedT(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
// TAST (Let): Comonad0_1_0 -> gopurs_runtime.Value
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
// TAST (Let): comonadTracedT_2_1 -> gopurs_runtime.Value
comonadTracedT_2_1 := gopurs_runtime.Apply(Get_Control_Comonad_Traced_Trans_comonadTracedT(), Comonad0_1_0)
_ = comonadTracedT_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
// TAST (Let): lower1_4_2 -> gopurs_runtime.Value
lower1_4_2 := gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))
}), v_5)
})
_ = lower1_4_2
// TAST (Let): comonadTracedT1_5_4 -> gopurs_runtime.Value
comonadTracedT1_5_4 := gopurs_runtime.Apply(comonadTracedT_2_1, dictMonoid_3)
_ = comonadTracedT1_5_4
return gopurs_runtime.RecordDict3("Comonad0", "peek", "pos", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_5_4
}), gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 -> gopurs_runtime.Value
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

func Call_Control_Comonad_Store_Class_comonadStoreStoreT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): comonadStoreT_1_0 -> gopurs_runtime.Value
comonadStoreT_1_0 := gopurs_runtime.Apply(Get_Control_Comonad_Store_Trans_comonadStoreT(), dictComonad_0)
_ = comonadStoreT_1_0
return gopurs_runtime.RecordDict3("Comonad0", "peek", "pos", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_1_0
}), gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_2)
})
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
}))
}

func Call_Control_Comonad_Store_Class_comonadStoreEnvT(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
// TAST (Let): Comonad0_1_0 -> gopurs_runtime.Value
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
// TAST (Let): lower1_2_1 -> gopurs_runtime.Value
lower1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Comonad_Env_Trans_comonadTransEnvT(), "lower"), Comonad0_1_0)
_ = lower1_2_1
// TAST (Let): comonadEnvT_3_2 -> gopurs_runtime.Value
comonadEnvT_3_2 := gopurs_runtime.Apply(Get_Control_Comonad_Env_Trans_comonadEnvT(), Comonad0_1_0)
_ = comonadEnvT_3_2
return gopurs_runtime.RecordDict3("Comonad0", "peek", "pos", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_3_2
}), gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_4)
_ = __local_var_5_3
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Apply(lower1_2_1, x_6))
})
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), gopurs_runtime.Apply(lower1_2_1, x_4))
}))
}

func Call_Control_Comonad_Store_Class_peek__413153475(dict_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Store_Class_peek__3957251203(dict_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Store_Class_peeks__2913274603(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictComonadStore_0.V1), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V2), x_2)), x_2)
}

func Call_Control_Comonad_Store_Class_peeks__110217771(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictComonadStore_0.V1), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V2), x_2)), x_2)
}

func Call_Control_Comonad_Store_Class_pos__639385354(dict_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Comonad_Store_Class_pos__3752461637(dict_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}



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

var cache_Control_Comonad_Store_Class_ComonadStore_dollar_Dict gopurs_runtime.Value
var once_Control_Comonad_Store_Class_ComonadStore_dollar_Dict sync.Once
func Get_Control_Comonad_Store_Class_ComonadStore_dollar_Dict() gopurs_runtime.Value {
	once_Control_Comonad_Store_Class_ComonadStore_dollar_Dict.Do(func() {
		cache_Control_Comonad_Store_Class_ComonadStore_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4264314723, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Store_Class_ComonadStore_dollar_Dict(func() struct{
	Comonad0 gopurs_runtime.Value
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Comonad0 gopurs_runtime.Value
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
}{}
					clone.Comonad0 = gopurs_runtime.RecordGet(orig, "Comonad0")
					clone.peek = gopurs_runtime.RecordGet(orig, "peek")
					clone.pos = gopurs_runtime.RecordGet(orig, "pos")
					return clone
				}()))}
})
	})
	return cache_Control_Comonad_Store_Class_ComonadStore_dollar_Dict
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

type Constructor_Control_Comonad_Store_Class_ComonadStore[T_s any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4264314723] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Store_Class_ComonadStore[any, any])(ptr)
		_ = c
		switch key {
		case "Comonad0": return gopurs_runtime.Box(c.V0)
		case "peek": return gopurs_runtime.Box(c.V1)
		case "pos": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Store_Class_ComonadStore: " + key)
		}
	}
}


func Call_Control_Comonad_Store_Class_ComonadStore_dollar_Dict(x_0_loop struct{
	Comonad0 gopurs_runtime.Value
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
}) *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	Comonad0 gopurs_runtime.Value
	peek gopurs_runtime.Value
	pos gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Comonad0", "peek", "pos"}, []gopurs_runtime.Value{orig.Comonad0, orig.peek, orig.pos})
				}())
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
// TAST (Let): duplicate_1_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar w) [(TypeVar a)])] (TypeApp (TypeVar w) [(TypeApp (TypeVar w) [(TypeVar a)])]))
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V0), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
_ = duplicate_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 shape=App(Other) bindingType=(TypeVar c)
__local_var_4_1 := gopurs_runtime.Apply(duplicate_1_0, x_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictComonadStore_0.V1), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V2), __local_var_4_1)), __local_var_4_1)
})
}

func Call_Control_Comonad_Store_Class_seek(dictComonadStore_0_loop *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictComonadStore_0 *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadStore_0_loop
_ = dictComonadStore_0
// TAST (Let): duplicate_1_0 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar w) [(TypeVar a)])] (TypeApp (TypeVar w) [(TypeApp (TypeVar w) [(TypeVar a)])]))
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadStore_0.V0), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
_ = duplicate_1_0
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
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
// TAST (Let): Comonad0_1_0 shape=App(Other) bindingType=Any
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
// TAST (Let): __local_var_2_3 shape=App(Other) bindingType=Any
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): Functor0_3_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar w)])
Functor0_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_4
// TAST (Let): __local_var_4_6 shape=App(Other) bindingType=Any
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_6
// TAST (Let): functorTracedT1_4_5 shape=Let(LitRecord) bindingType=(ADT ["Data","Functor","Functor"] [(TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])])
functorTracedT1_4_5 := (&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_6, "map"), gopurs_runtime.Func2(func(g_7 gopurs_runtime.Value, t_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Apply(g_7, t_8))
}), v_6)
})})
_ = functorTracedT1_4_5
// TAST (Let): extendTracedT1__193435443_2_2 shape=Let(Let(Let(Abs(LitRecord)))) bindingType=Any
extendTracedT1__193435443_2_2 := gopurs_runtime.Func(func(dictSemigroup_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorTracedT1_4_5)}
}), gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "extend"), gopurs_runtime.Func2(func(w_prime__8 gopurs_runtime.Value, t_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_4.V0), gopurs_runtime.Func2(func(h_10 gopurs_runtime.Value, t_prime__11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_10, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_5, "append"), t_9, t_prime__11))
}), w_prime__8))
}), v_7)
})}))}
})
_ = extendTracedT1__193435443_2_2
// TAST (Let): comonadTracedT__193435443_2_1 shape=Let(Abs(Let(LitRecord))) bindingType=Any
comonadTracedT__193435443_2_1 := gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): extendTracedT2_4_7 shape=App(Other) bindingType=(ADT ["Control","Extend","Extend"] [(TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])])
extendTracedT2_4_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](gopurs_runtime.Apply(extendTracedT1__193435443_2_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{})))
_ = extendTracedT2_4_7
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendTracedT2_4_7)}
}), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Comonad0_1_0, "extract"), v_5, gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))
})}))}
})
_ = comonadTracedT__193435443_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_9 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar w)])
Functor0_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_9
// TAST (Let): lower1_4_8 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeApp (TypeVar w) [(Func [(TypeVar m)] (TypeVar a))])] (TypeApp (TypeVar w) [(TypeVar a)]))
lower1_4_8 := gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_9.V0), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))
}), v_5)
})
_ = lower1_4_8
// TAST (Let): comonadTracedT1_5_10 shape=App(Other) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (TypeVar w) [(Func [(TypeVar m)] (TypeVar a))])])
comonadTracedT1_5_10 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(comonadTracedT__193435443_2_1, dictMonoid_3))
_ = comonadTracedT1_5_10
return gopurs_runtime.Value{Type: 9, IntVal: 4264314723, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadTracedT1_5_10)}
}), gopurs_runtime.Func(func(s_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_11 shape=App(Other) bindingType=Any
__local_var_7_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_6)
_ = __local_var_7_11
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_7_11, gopurs_runtime.Apply(lower1_4_8, x_8))
})
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), gopurs_runtime.Apply(lower1_4_8, x_6))
})}))}
})
}

func Call_Control_Comonad_Store_Class_comonadStoreStoreT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): comonadStoreT_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w) [(Func [(TypeVar s)] (TypeVar a))]), (TypeVar s)])])
comonadStoreT_1_0 := Rebox_Control_Comonad_Store_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Comonad_Store_Trans_comonadStoreT(), dictComonad_0)))
_ = comonadStoreT_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4264314723, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Class_229239260_3433217233((&Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Class_3056445460_2550391993(comonadStoreT_1_0))}
}), gopurs_runtime.Func2(func(s_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, s_2)
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
})})))}
}

func Call_Control_Comonad_Store_Class_comonadStoreEnvT(dictComonadStore_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadStore_0 gopurs_runtime.Value = dictComonadStore_0_loop
_ = dictComonadStore_0
// TAST (Let): comonadEnvT_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e), (TypeVar w)])])
comonadEnvT_1_0 := Rebox_Control_Comonad_Store_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Comonad_Env_Trans_comonadEnvT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{}))))
_ = comonadEnvT_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 4264314723, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Class_229239260_3433217233((&Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Store_Class_3056445460_2550391993(comonadEnvT_1_0))}
}), gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=Any
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1)
})})))}
}

func Rebox_Control_Comonad_Store_Class_229239260_3433217233(in *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Comonad_Store_Class_ComonadStore[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
	return out
}

func Rebox_Control_Comonad_Store_Class_2550391993_3056445460(in *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]) *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Comonad_Store_Class_3056445460_2550391993(in *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}



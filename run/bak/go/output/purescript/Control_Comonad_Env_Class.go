package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Env_Class_lower gopurs_runtime.Value
var once_Control_Comonad_Env_Class_lower sync.Once
func Get_Control_Comonad_Env_Class_lower() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_lower.Do(func() {
		cache_Control_Comonad_Env_Class_lower = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_lower(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad](dictComonad_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_lower
}

var cache_Control_Comonad_Env_Class_ComonadAsk_dollarDict gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ComonadAsk_dollarDict sync.Once
func Get_Control_Comonad_Env_Class_ComonadAsk_dollarDict() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ComonadAsk_dollarDict.Do(func() {
		cache_Control_Comonad_Env_Class_ComonadAsk_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_ComonadAsk_dollarDict(x_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_ComonadAsk_dollarDict
}

var cache_Control_Comonad_Env_Class_ComonadEnv_dollarDict gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ComonadEnv_dollarDict sync.Once
func Get_Control_Comonad_Env_Class_ComonadEnv_dollarDict() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ComonadEnv_dollarDict.Do(func() {
		cache_Control_Comonad_Env_Class_ComonadEnv_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_ComonadEnv_dollarDict(x_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_ComonadEnv_dollarDict
}

var cache_Control_Comonad_Env_Class_local gopurs_runtime.Value
var once_Control_Comonad_Env_Class_local sync.Once
func Get_Control_Comonad_Env_Class_local() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_local.Do(func() {
		cache_Control_Comonad_Env_Class_local = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_local(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_local
}

var cache_Control_Comonad_Env_Class_comonadAskTuple gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskTuple sync.Once
func Get_Control_Comonad_Env_Class_comonadAskTuple() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskTuple.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskTuple = gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad](Get_Data_Tuple_comonadTuple()))}
}), Get_Data_Tuple_fst()})}
	})
	return cache_Control_Comonad_Env_Class_comonadAskTuple
}

var cache_Control_Comonad_Env_Class_comonadEnvTuple gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvTuple sync.Once
func Get_Control_Comonad_Env_Class_comonadEnvTuple() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvTuple.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvTuple = gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadEnv{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk](Get_Control_Comonad_Env_Class_comonadAskTuple()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V1})}
})
})})}
	})
	return cache_Control_Comonad_Env_Class_comonadEnvTuple
}

var cache_Control_Comonad_Env_Class_comonadAskEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskEnvT sync.Once
func Get_Control_Comonad_Env_Class_comonadAskEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskEnvT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadAskEnvT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadAskEnvT
}

var cache_Control_Comonad_Env_Class_comonadEnvEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvEnvT sync.Once
func Get_Control_Comonad_Env_Class_comonadEnvEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvEnvT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadEnvEnvT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadEnvEnvT
}

var cache_Control_Comonad_Env_Class_ask gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ask sync.Once
func Get_Control_Comonad_Env_Class_ask() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ask.Do(func() {
		cache_Control_Comonad_Env_Class_ask = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_ask(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_ask
}

var cache_Control_Comonad_Env_Class_asks gopurs_runtime.Value
var once_Control_Comonad_Env_Class_asks sync.Once
func Get_Control_Comonad_Env_Class_asks() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_asks.Do(func() {
		cache_Control_Comonad_Env_Class_asks = gopurs_runtime.Func3(func(dictComonadAsk_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_asks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk](dictComonadAsk_0_box), f_1_box, x_2_box)
})
	})
	return cache_Control_Comonad_Env_Class_asks
}

var cache_Control_Comonad_Env_Class_comonadAskStoreT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskStoreT sync.Once
func Get_Control_Comonad_Env_Class_comonadAskStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskStoreT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskStoreT = gopurs_runtime.Func(func(dictComonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadAskStoreT(dictComonadAsk_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadAskStoreT
}

var cache_Control_Comonad_Env_Class_comonadEnvStoreT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvStoreT sync.Once
func Get_Control_Comonad_Env_Class_comonadEnvStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvStoreT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvStoreT = gopurs_runtime.Func(func(dictComonadEnv_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadEnvStoreT(dictComonadEnv_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadEnvStoreT
}

var cache_Control_Comonad_Env_Class_comonadAskTracedT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskTracedT sync.Once
func Get_Control_Comonad_Env_Class_comonadAskTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskTracedT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskTracedT = gopurs_runtime.Func(func(dictComonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadAskTracedT(dictComonadAsk_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadAskTracedT
}

var cache_Control_Comonad_Env_Class_comonadEnvTracedT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvTracedT sync.Once
func Get_Control_Comonad_Env_Class_comonadEnvTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvTracedT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvTracedT = gopurs_runtime.Func(func(dictComonadEnv_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadEnvTracedT(dictComonadEnv_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadEnvTracedT
}

var cache_Control_Comonad_Env_Class_ask__2345252920 gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ask__2345252920 sync.Once
func Get_Control_Comonad_Env_Class_ask__2345252920() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ask__2345252920.Do(func() {
		cache_Control_Comonad_Env_Class_ask__2345252920 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_ask__2345252920(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_ask__2345252920
}

var cache_Control_Comonad_Env_Class_comonadAskTuple__132564679 gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskTuple__132564679 sync.Once
func Get_Control_Comonad_Env_Class_comonadAskTuple__132564679() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskTuple__132564679.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskTuple__132564679 = gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad](Get_Data_Tuple_comonadTuple()))}
}), Get_Data_Tuple_fst()})}
	})
	return cache_Control_Comonad_Env_Class_comonadAskTuple__132564679
}

var cache_Control_Comonad_Env_Class_local__994074898 gopurs_runtime.Value
var once_Control_Comonad_Env_Class_local__994074898 sync.Once
func Get_Control_Comonad_Env_Class_local__994074898() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_local__994074898.Do(func() {
		cache_Control_Comonad_Env_Class_local__994074898 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_local__994074898(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_local__994074898
}

var cache_Control_Comonad_Env_Class_local__1731941714 gopurs_runtime.Value
var once_Control_Comonad_Env_Class_local__1731941714 sync.Once
func Get_Control_Comonad_Env_Class_local__1731941714() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_local__1731941714.Do(func() {
		cache_Control_Comonad_Env_Class_local__1731941714 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_local__1731941714(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_local__1731941714
}

type Constructor_Control_Comonad_Env_Class_ComonadAsk struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1424092807] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Env_Class_ComonadAsk)(ptr)
		_ = c
		switch key {
		case "Comonad0": return gopurs_runtime.Box(c.V0)
		case "ask": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Env_Class_ComonadAsk: " + key)
		}
	}
}


type Constructor_Control_Comonad_Env_Class_ComonadEnv struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3863290147] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Env_Class_ComonadEnv)(ptr)
		_ = c
		switch key {
		case "ComonadAsk0": return gopurs_runtime.Box(c.V0)
		case "local": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Env_Class_ComonadEnv: " + key)
		}
	}
}


func Call_Control_Comonad_Env_Class_lower(dictComonad_0_loop *Constructor_Control_Comonad_Comonad) gopurs_runtime.Value {
var dictComonad_0 *Constructor_Control_Comonad_Comonad = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictComonad_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1
_ = __local_var_3_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, __local_var_3_1)
}), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0)
})
}

func Call_Control_Comonad_Env_Class_ComonadAsk_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Env_Class_ComonadEnv_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Env_Class_local(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadEnv) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadEnv = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_comonadAskEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_2
// TAST (Let): Functor0_2_3 -> *Constructor_Data_Functor_Functor
Functor0_2_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_3
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_2, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorEnvT1_3_4 -> *Constructor_Data_Functor_Functor
functorEnvT1_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "map"), f_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1)})}
})
})))
_ = functorEnvT1_3_4
// TAST (Let): extendEnvT1_1_1 -> *Constructor_Control_Extend_Extend
extendEnvT1_1_1 := &Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorEnvT1_3_4)}
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0)
_ = __local_var_6_6
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_3.V0), f_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_2, "extend"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_6, x_7)
}), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1))})}
})
})}
_ = extendEnvT1_1_1
// TAST (Let): comonadEnvT_1_0 -> *Constructor_Control_Comonad_Comonad
comonadEnvT_1_0 := &Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendEnvT1_1_1)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})}
_ = comonadEnvT_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadEnvT_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0
})})}
}

func Call_Control_Comonad_Env_Class_comonadEnvEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): __local_var_1_3 -> gopurs_runtime.Value
__local_var_1_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_3
// TAST (Let): Functor0_2_4 -> *Constructor_Data_Functor_Functor
Functor0_2_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_4
// TAST (Let): __local_var_3_6 -> gopurs_runtime.Value
__local_var_3_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_6
// TAST (Let): functorEnvT1_3_5 -> *Constructor_Data_Functor_Functor
functorEnvT1_3_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_6, "map"), f_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1)})}
})
})))
_ = functorEnvT1_3_5
// TAST (Let): extendEnvT1_1_2 -> *Constructor_Control_Extend_Extend
extendEnvT1_1_2 := &Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorEnvT1_3_5)}
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0)
_ = __local_var_6_7
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_4.V0), f_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_3, "extend"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_7, x_7)
}), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1))})}
})
})}
_ = extendEnvT1_1_2
// TAST (Let): comonadEnvT_1_1 -> *Constructor_Control_Comonad_Comonad
comonadEnvT_1_1 := &Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendEnvT1_1_2)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
})}
_ = comonadEnvT_1_1
// TAST (Let): comonadAskEnvT1_1_0 -> *Constructor_Control_Comonad_Env_Class_ComonadAsk
comonadAskEnvT1_1_0 := &Constructor_Control_Comonad_Env_Class_ComonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadEnvT_1_1)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0
})}
_ = comonadAskEnvT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadEnv{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(comonadAskEnvT1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1})}
})
})})}
}

func Call_Control_Comonad_Env_Class_ask(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_asks(dictComonadAsk_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk = dictComonadAsk_0_loop
_ = dictComonadAsk_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadAsk_0.V1), x_2))
}

func Call_Control_Comonad_Env_Class_comonadAskStoreT(dictComonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
_ = dictComonadAsk_0
// TAST (Let): Comonad0_1_0 -> gopurs_runtime.Value
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): functorStoreT1_3_4 -> *Constructor_Data_Functor_Functor
functorStoreT1_3_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "map"), gopurs_runtime.Func(func(h_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Apply(h_6, x_7))
})
}), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1})}
})
})))
_ = functorStoreT1_3_4
// TAST (Let): extendStoreT1_2_2 -> *Constructor_Control_Extend_Extend
extendStoreT1_2_2 := &Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStoreT1_3_4)}
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "extend"), gopurs_runtime.Func(func(w_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, w_prime_6, s_prime_7})})
})
}), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1})}
})
})}
_ = extendStoreT1_2_2
// TAST (Let): comonadStoreT_2_1 -> *Constructor_Control_Comonad_Comonad
comonadStoreT_2_1 := &Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendStoreT1_2_2)}
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Comonad0_1_0, "extract"), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)
})}
_ = comonadStoreT_2_1
// TAST (Let): Functor0_3_7 -> *Constructor_Data_Functor_Functor
Functor0_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_7
// TAST (Let): __local_var_3_6 -> gopurs_runtime.Value
__local_var_3_6 := gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_8 -> gopurs_runtime.Value
__local_var_5_8 := (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1
_ = __local_var_5_8
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_7.V0), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, __local_var_5_8)
}), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0)
})
_ = __local_var_3_6
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadStoreT_2_1)}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "ask"), gopurs_runtime.Apply(__local_var_3_6, x_4))
})})}
}

func Call_Control_Comonad_Env_Class_comonadEnvStoreT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
_ = dictComonadEnv_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Comonad0_2_2 -> gopurs_runtime.Value
Comonad0_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_2_2
// TAST (Let): __local_var_3_5 -> gopurs_runtime.Value
__local_var_3_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_2_2, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_3_5
// TAST (Let): __local_var_4_7 -> gopurs_runtime.Value
__local_var_4_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_5, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_7
// TAST (Let): functorStoreT1_4_6 -> *Constructor_Data_Functor_Functor
functorStoreT1_4_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_7, "map"), gopurs_runtime.Func(func(h_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Apply(h_7, x_8))
})
}), (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1})}
})
})))
_ = functorStoreT1_4_6
// TAST (Let): extendStoreT1_3_4 -> *Constructor_Control_Extend_Extend
extendStoreT1_3_4 := &Constructor_Control_Extend_Extend{1, gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorStoreT1_4_6)}
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_5, "extend"), gopurs_runtime.Func(func(w_prime_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, w_prime_7, s_prime_8})})
})
}), (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1})}
})
})}
_ = extendStoreT1_3_4
// TAST (Let): comonadStoreT_3_3 -> *Constructor_Control_Comonad_Comonad
comonadStoreT_3_3 := &Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendStoreT1_3_4)}
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Comonad0_2_2, "extract"), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)
})}
_ = comonadStoreT_3_3
// TAST (Let): Functor0_4_9 -> *Constructor_Data_Functor_Functor
Functor0_4_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_2_2, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_9
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_10 -> gopurs_runtime.Value
__local_var_6_10 := (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1
_ = __local_var_6_10
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_9.V0), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_7, __local_var_6_10)
}), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0)
})
_ = __local_var_4_8
// TAST (Let): comonadAskStoreT1_1_0 -> *Constructor_Control_Comonad_Env_Class_ComonadAsk
comonadAskStoreT1_1_0 := &Constructor_Control_Comonad_Env_Class_ComonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadStoreT_3_3)}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "ask"), gopurs_runtime.Apply(__local_var_4_8, x_5))
})}
_ = comonadAskStoreT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadEnv{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(comonadAskStoreT1_1_0)}
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1})}
})
})})}
}

func Call_Control_Comonad_Env_Class_comonadAskTracedT(dictComonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
_ = dictComonadAsk_0
// TAST (Let): Comonad0_1_0 -> gopurs_runtime.Value
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
// TAST (Let): __local_var_2_1 -> gopurs_runtime.Value
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_2_1
// TAST (Let): Functor0_3_2 -> *Constructor_Data_Functor_Functor
Functor0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_4
// TAST (Let): functorTracedT1_4_3 -> *Constructor_Data_Functor_Functor
functorTracedT1_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_4, "map"), gopurs_runtime.Func(func(g_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Apply(g_7, t_8))
})
}), v_6)
})
})))
_ = functorTracedT1_4_3
return gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_7 -> gopurs_runtime.Value
__local_var_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_6_7
// TAST (Let): extendTracedT2_6_6 -> *Constructor_Control_Extend_Extend
extendTracedT2_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorTracedT1_4_3)}
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "extend"), gopurs_runtime.Func(func(w_prime_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_2.V0), gopurs_runtime.Func(func(h_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_prime_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_11, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_6_7, "append"), t_10, t_prime_12))
})
}), w_prime_9))
})
}), v_8)
})
})))
_ = extendTracedT2_6_6
// TAST (Let): comonadTracedT1_6_5 -> *Constructor_Control_Comonad_Comonad
comonadTracedT1_6_5 := &Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendTracedT2_6_6)}
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Comonad0_1_0, "extract"), v_7, gopurs_runtime.RecordGet(dictMonoid_5, "mempty"))
})}
_ = comonadTracedT1_6_5
// TAST (Let): Functor0_7_9 -> *Constructor_Data_Functor_Functor
Functor0_7_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_9
// TAST (Let): __local_var_7_8 -> gopurs_runtime.Value
__local_var_7_8 := gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_9.V0), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_9, gopurs_runtime.RecordGet(dictMonoid_5, "mempty"))
}), v_8)
})
_ = __local_var_7_8
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadTracedT1_6_5)}
}), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "ask"), gopurs_runtime.Apply(__local_var_7_8, x_8))
})})}
})
}

func Call_Control_Comonad_Env_Class_comonadEnvTracedT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
_ = dictComonadEnv_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{})
_ = __local_var_1_0
// TAST (Let): Comonad0_2_1 -> gopurs_runtime.Value
Comonad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_2_1
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_2_1, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_3_3
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_3, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_5_6
// TAST (Let): functorTracedT1_5_5 -> *Constructor_Data_Functor_Functor
functorTracedT1_5_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_5_6, "map"), gopurs_runtime.Func(func(g_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, gopurs_runtime.Apply(g_8, t_9))
})
}), v_7)
})
})))
_ = functorTracedT1_5_5
// TAST (Let): comonadAskTracedT1_3_2 -> gopurs_runtime.Value
comonadAskTracedT1_3_2 := gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_9 -> gopurs_runtime.Value
__local_var_7_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_7_9
// TAST (Let): extendTracedT2_7_8 -> *Constructor_Control_Extend_Extend
extendTracedT2_7_8 := gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend](gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(functorTracedT1_5_5)}
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_3, "extend"), gopurs_runtime.Func(func(w_prime_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_8, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), gopurs_runtime.Func(func(h_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_prime_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(h_12, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_7_9, "append"), t_11, t_prime_13))
})
}), w_prime_10))
})
}), v_9)
})
})))
_ = extendTracedT2_7_8
// TAST (Let): comonadTracedT1_7_7 -> *Constructor_Control_Comonad_Comonad
comonadTracedT1_7_7 := &Constructor_Control_Comonad_Comonad{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(extendTracedT2_7_8)}
}), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Comonad0_2_1, "extract"), v_8, gopurs_runtime.RecordGet(dictMonoid_6, "mempty"))
})}
_ = comonadTracedT1_7_7
// TAST (Let): Functor0_8_11 -> *Constructor_Data_Functor_Functor
Functor0_8_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_2_1, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_8_11
// TAST (Let): __local_var_8_10 -> gopurs_runtime.Value
__local_var_8_10 := gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_8_11.V0), gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_10, gopurs_runtime.RecordGet(dictMonoid_6, "mempty"))
}), v_9)
})
_ = __local_var_8_10
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadAsk{1, gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadTracedT1_7_7)}
}), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "ask"), gopurs_runtime.Apply(__local_var_8_10, x_9))
})})}
})
_ = comonadAskTracedT1_3_2
return gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): comonadAskTracedT2_5_12 -> *Constructor_Control_Comonad_Env_Class_ComonadAsk
comonadAskTracedT2_5_12 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk](gopurs_runtime.Apply(comonadAskTracedT1_3_2, dictMonoid_4))
_ = comonadAskTracedT2_5_12
return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Env_Class_ComonadEnv{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(comonadAskTracedT2_5_12)}
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_6, v_7)
})
})})}
})
}

func Call_Control_Comonad_Env_Class_ask__2345252920(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_local__994074898(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadEnv) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadEnv = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_local__1731941714(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadEnv) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadEnv = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



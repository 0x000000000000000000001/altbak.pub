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
		cache_Control_Comonad_Env_Class_lower = gopurs_runtime.RecordGet(Get_Control_Comonad_Store_Trans_comonadTransStoreT(), "lower")
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
return Call_Control_Comonad_Env_Class_local(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_local
}

var cache_Control_Comonad_Env_Class_comonadAskTuple gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskTuple sync.Once
func Get_Control_Comonad_Env_Class_comonadAskTuple() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskTuple.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskTuple = gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_comonadTuple()
}), Get_Data_Tuple_fst())
	})
	return cache_Control_Comonad_Env_Class_comonadAskTuple
}

var cache_Control_Comonad_Env_Class_comonadEnvTuple gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvTuple sync.Once
func Get_Control_Comonad_Env_Class_comonadEnvTuple() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvTuple.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvTuple = gopurs_runtime.RecordDict2("ComonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Comonad_Env_Class_comonadAskTuple()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1})}
})
}))
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
return Call_Control_Comonad_Env_Class_ask(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_ask
}

var cache_Control_Comonad_Env_Class_asks gopurs_runtime.Value
var once_Control_Comonad_Env_Class_asks sync.Once
func Get_Control_Comonad_Env_Class_asks() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_asks.Do(func() {
		cache_Control_Comonad_Env_Class_asks = gopurs_runtime.Func3(func(dictComonadAsk_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_asks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadAsk_0_box), f_1_box, x_2_box)
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
return Call_Control_Comonad_Env_Class_ask__2345252920(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_ask__2345252920
}

var cache_Control_Comonad_Env_Class_comonadAskTuple__1515377815 gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskTuple__1515377815 sync.Once
func Get_Control_Comonad_Env_Class_comonadAskTuple__1515377815() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskTuple__1515377815.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskTuple__1515377815 = gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Tuple_comonadTuple()
}), Get_Data_Tuple_fst())
	})
	return cache_Control_Comonad_Env_Class_comonadAskTuple__1515377815
}

var cache_Control_Comonad_Env_Class_local__994074898 gopurs_runtime.Value
var once_Control_Comonad_Env_Class_local__994074898 sync.Once
func Get_Control_Comonad_Env_Class_local__994074898() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_local__994074898.Do(func() {
		cache_Control_Comonad_Env_Class_local__994074898 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_local__994074898(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_local__994074898
}

var cache_Control_Comonad_Env_Class_local__1731941714 gopurs_runtime.Value
var once_Control_Comonad_Env_Class_local__1731941714 sync.Once
func Get_Control_Comonad_Env_Class_local__1731941714() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_local__1731941714.Do(func() {
		cache_Control_Comonad_Env_Class_local__1731941714 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_local__1731941714(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_local__1731941714
}

type Constructor_Control_Comonad_Env_Class_ComonadAsk[T_e any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1424092807] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Comonad0": return gopurs_runtime.Box(c.V0)
		case "ask": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Env_Class_ComonadAsk: " + key)
		}
	}
}


type Constructor_Control_Comonad_Env_Class_ComonadEnv[T_e any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3863290147] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "ComonadAsk0": return gopurs_runtime.Box(c.V0)
		case "local": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Env_Class_ComonadEnv: " + key)
		}
	}
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

func Call_Control_Comonad_Env_Class_local(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_comonadAskEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): comonadEnvT_1_0 -> gopurs_runtime.Value
comonadEnvT_1_0 := gopurs_runtime.Apply(Get_Control_Comonad_Env_Trans_comonadEnvT(), dictComonad_0)
_ = comonadEnvT_1_0
return gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
}))
}

func Call_Control_Comonad_Env_Class_comonadEnvEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): comonadEnvT_1_1 -> gopurs_runtime.Value
comonadEnvT_1_1 := gopurs_runtime.Apply(Get_Control_Comonad_Env_Trans_comonadEnvT(), dictComonad_0)
_ = comonadEnvT_1_1
// TAST (Let): comonadAskEnvT1_1_0 -> gopurs_runtime.Value
comonadAskEnvT1_1_0 := gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
}))
_ = comonadAskEnvT1_1_0
return gopurs_runtime.RecordDict2("ComonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadAskEnvT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1})}
})
}))
}

func Call_Control_Comonad_Env_Class_ask(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_asks(dictComonadAsk_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadAsk_0_loop
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
// TAST (Let): comonadStoreT_2_1 -> gopurs_runtime.Value
comonadStoreT_2_1 := gopurs_runtime.Apply(Get_Control_Comonad_Store_Trans_comonadStoreT(), Comonad0_1_0)
_ = comonadStoreT_2_1
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Comonad_Store_Trans_comonadTransStoreT(), "lower"), Comonad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "ask"), gopurs_runtime.Apply(__local_var_3_2, x_4))
}))
}

func Call_Control_Comonad_Env_Class_comonadEnvStoreT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
_ = dictComonadEnv_0
// TAST (Let): comonadAskStoreT1_1_0 -> gopurs_runtime.Value
comonadAskStoreT1_1_0 := Call_Control_Comonad_Env_Class_comonadAskStoreT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{}))
_ = comonadAskStoreT1_1_0
return gopurs_runtime.RecordDict2("ComonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadAskStoreT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1})}
})
}))
}

func Call_Control_Comonad_Env_Class_comonadAskTracedT(dictComonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
_ = dictComonadAsk_0
// TAST (Let): Comonad0_1_0 -> gopurs_runtime.Value
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
// TAST (Let): comonadTracedT_2_1 -> gopurs_runtime.Value
comonadTracedT_2_1 := gopurs_runtime.Apply(Get_Control_Comonad_Traced_Trans_comonadTracedT(), Comonad0_1_0)
_ = comonadTracedT_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): comonadTracedT1_4_2 -> gopurs_runtime.Value
comonadTracedT1_4_2 := gopurs_runtime.Apply(comonadTracedT_2_1, dictMonoid_3)
_ = comonadTracedT1_4_2
// TAST (Let): Functor0_5_4 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_4
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_4.V0), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))
}), v_6)
})
_ = __local_var_5_3
return gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_4_2
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "ask"), gopurs_runtime.Apply(__local_var_5_3, x_6))
}))
})
}

func Call_Control_Comonad_Env_Class_comonadEnvTracedT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
_ = dictComonadEnv_0
// TAST (Let): comonadAskTracedT1_1_0 -> gopurs_runtime.Value
comonadAskTracedT1_1_0 := Call_Control_Comonad_Env_Class_comonadAskTracedT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{}))
_ = comonadAskTracedT1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): comonadAskTracedT2_3_1 -> gopurs_runtime.Value
comonadAskTracedT2_3_1 := gopurs_runtime.Apply(comonadAskTracedT1_1_0, dictMonoid_2)
_ = comonadAskTracedT2_3_1
return gopurs_runtime.RecordDict2("ComonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadAskTracedT2_3_1
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_4, v_5)
})
}))
})
}

func Call_Control_Comonad_Env_Class_ask__2345252920(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_local__994074898(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_local__1731941714(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



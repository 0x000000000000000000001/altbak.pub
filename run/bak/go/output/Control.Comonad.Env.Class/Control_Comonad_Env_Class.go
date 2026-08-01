package Control_Comonad_Env_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
	pkg_Control_Comonad_Store_Trans "gopurs/output/Control.Comonad.Store.Trans"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
	unsafe "unsafe"
)

var cache_local gopurs_runtime.Value
var once_local sync.Once
func Get_local() gopurs_runtime.Value {
	once_local.Do(func() {
		cache_local = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local(dict_0_box)
})
	})
	return cache_local
}

var cache_local__gopurs_runtime_Value_1862408154 gopurs_runtime.Value
var once_local__gopurs_runtime_Value_1862408154 sync.Once
func Get_local__gopurs_runtime_Value_1862408154() gopurs_runtime.Value {
	once_local__gopurs_runtime_Value_1862408154.Do(func() {
		cache_local__gopurs_runtime_Value_1862408154 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local__gopurs_runtime_Value_1862408154(dict_0_box)
})
	})
	return cache_local__gopurs_runtime_Value_1862408154
}

var cache_comonadAskTuple gopurs_runtime.Value
var once_comonadAskTuple sync.Once
func Get_comonadAskTuple() gopurs_runtime.Value {
	once_comonadAskTuple.Do(func() {
		cache_comonadAskTuple = gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_comonadTuple()
}), pkg_Data_Tuple.Get_fst())
	})
	return cache_comonadAskTuple
}

var cache_comonadEnvTuple gopurs_runtime.Value
var once_comonadEnvTuple sync.Once
func Get_comonadEnvTuple() gopurs_runtime.Value {
	once_comonadEnvTuple.Do(func() {
		cache_comonadEnvTuple = gopurs_runtime.RecordDict2("ComonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_comonadAskTuple()
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1})}
}))
	})
	return cache_comonadEnvTuple
}

var cache_comonadAskEnvT gopurs_runtime.Value
var once_comonadAskEnvT sync.Once
func Get_comonadAskEnvT() gopurs_runtime.Value {
	once_comonadAskEnvT.Do(func() {
		cache_comonadAskEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadAskEnvT(dictComonad_0_box)
})
	})
	return cache_comonadAskEnvT
}

var cache_comonadEnvEnvT gopurs_runtime.Value
var once_comonadEnvEnvT sync.Once
func Get_comonadEnvEnvT() gopurs_runtime.Value {
	once_comonadEnvEnvT.Do(func() {
		cache_comonadEnvEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadEnvEnvT(dictComonad_0_box)
})
	})
	return cache_comonadEnvEnvT
}

var cache_ask gopurs_runtime.Value
var once_ask sync.Once
func Get_ask() gopurs_runtime.Value {
	once_ask.Do(func() {
		cache_ask = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ask(dict_0_box)
})
	})
	return cache_ask
}

var cache_ask__gopurs_runtime_Value_2839064398 gopurs_runtime.Value
var once_ask__gopurs_runtime_Value_2839064398 sync.Once
func Get_ask__gopurs_runtime_Value_2839064398() gopurs_runtime.Value {
	once_ask__gopurs_runtime_Value_2839064398.Do(func() {
		cache_ask__gopurs_runtime_Value_2839064398 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ask__gopurs_runtime_Value_2839064398(dict_0_box)
})
	})
	return cache_ask__gopurs_runtime_Value_2839064398
}

var cache_asks gopurs_runtime.Value
var once_asks sync.Once
func Get_asks() gopurs_runtime.Value {
	once_asks.Do(func() {
		cache_asks = gopurs_runtime.Func3(func(dictComonadAsk_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_asks(dictComonadAsk_0_box, f_1_box, x_2_box)
})
	})
	return cache_asks
}

var cache_comonadAskStoreT gopurs_runtime.Value
var once_comonadAskStoreT sync.Once
func Get_comonadAskStoreT() gopurs_runtime.Value {
	once_comonadAskStoreT.Do(func() {
		cache_comonadAskStoreT = gopurs_runtime.Func(func(dictComonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadAskStoreT(dictComonadAsk_0_box)
})
	})
	return cache_comonadAskStoreT
}

var cache_comonadEnvStoreT gopurs_runtime.Value
var once_comonadEnvStoreT sync.Once
func Get_comonadEnvStoreT() gopurs_runtime.Value {
	once_comonadEnvStoreT.Do(func() {
		cache_comonadEnvStoreT = gopurs_runtime.Func(func(dictComonadEnv_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadEnvStoreT(dictComonadEnv_0_box)
})
	})
	return cache_comonadEnvStoreT
}

var cache_comonadAskTracedT gopurs_runtime.Value
var once_comonadAskTracedT sync.Once
func Get_comonadAskTracedT() gopurs_runtime.Value {
	once_comonadAskTracedT.Do(func() {
		cache_comonadAskTracedT = gopurs_runtime.Func(func(dictComonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadAskTracedT(dictComonadAsk_0_box)
})
	})
	return cache_comonadAskTracedT
}

var cache_comonadEnvTracedT gopurs_runtime.Value
var once_comonadEnvTracedT sync.Once
func Get_comonadEnvTracedT() gopurs_runtime.Value {
	once_comonadEnvTracedT.Do(func() {
		cache_comonadEnvTracedT = gopurs_runtime.Func(func(dictComonadEnv_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadEnvTracedT(dictComonadEnv_0_box)
})
	})
	return cache_comonadEnvTracedT
}

func Call_local(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "local")
}

func Call_local__gopurs_runtime_Value_1862408154(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "local")
}

func Call_comonadAskEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), dictComonad_0)
_ = comonadEnvT_1_0
return gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
}))
}

func Call_comonadEnvEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), dictComonad_0)
_ = comonadEnvT_1_0
return gopurs_runtime.RecordDict2("ComonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
}))
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1})}
}))
}

func Call_ask(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ask")
}

func Call_ask__gopurs_runtime_Value_2839064398(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ask")
}

func Call_asks(dictComonadAsk_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
_ = dictComonadAsk_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "ask"), x_2))
}

func Call_comonadAskStoreT(dictComonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
_ = dictComonadAsk_0
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
comonadStoreT_2_1 := gopurs_runtime.Apply(pkg_Control_Comonad_Store_Trans.Get_comonadStoreT(), Comonad0_1_0)
_ = comonadStoreT_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Comonad_Store_Trans.Get_comonadTransStoreT(), "lower"), Comonad0_1_0)
_ = __local_var_3_2
return gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_2_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "ask"), gopurs_runtime.Apply(__local_var_3_2, x_4))
}))
}

func Call_comonadEnvStoreT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
_ = dictComonadEnv_0
comonadAskStoreT1_1_0 := Call_comonadAskStoreT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{}))
_ = comonadAskStoreT1_1_0
return gopurs_runtime.RecordDict2("ComonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadAskStoreT1_1_0
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1})}
}))
}

func Call_comonadAskTracedT(dictComonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
_ = dictComonadAsk_0
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
comonadTracedT_2_1 := gopurs_runtime.Apply(pkg_Control_Comonad_Traced_Trans.Get_comonadTracedT(), Comonad0_1_0)
_ = comonadTracedT_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
comonadTracedT1_4_2 := gopurs_runtime.Apply(comonadTracedT_2_1, dictMonoid_3)
_ = comonadTracedT1_4_2
mempty_5_3 := gopurs_runtime.RecordGet(dictMonoid_3, "mempty")
_ = mempty_5_3
return gopurs_runtime.RecordDict2("Comonad0", "ask", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_4_2
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "ask"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, mempty_5_3)
}), x_6))
}))
})
}

func Call_comonadEnvTracedT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
_ = dictComonadEnv_0
comonadAskTracedT1_1_0 := Call_comonadAskTracedT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{}))
_ = comonadAskTracedT1_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
comonadAskTracedT2_3_1 := gopurs_runtime.Apply(comonadAskTracedT1_1_0, dictMonoid_2)
_ = comonadAskTracedT2_3_1
return gopurs_runtime.RecordDict2("ComonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadAskTracedT2_3_1
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_4, v_5)
}))
})
}



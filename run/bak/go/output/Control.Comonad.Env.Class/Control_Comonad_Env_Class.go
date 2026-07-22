package Control_Comonad_Env_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
	pkg_Control_Comonad_Store_Trans "gopurs/output/Control.Comonad.Store.Trans"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
)

var local gopurs_runtime.Value
var once_local sync.Once
func Get_local() gopurs_runtime.Value {
	once_local.Do(func() {
		local = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["local"]
})
	})
	return local
}

var comonadAskTuple gopurs_runtime.Value
var once_comonadAskTuple sync.Once
func Get_comonadAskTuple() gopurs_runtime.Value {
	once_comonadAskTuple.Do(func() {
		comonadAskTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": pkg_Data_Tuple.Get_fst(), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Tuple.Get_comonadTuple()
})})
	})
	return comonadAskTuple
}

var comonadEnvTuple gopurs_runtime.Value
var once_comonadEnvTuple sync.Once
func Get_comonadEnvTuple() gopurs_runtime.Value {
	once_comonadEnvTuple.Do(func() {
		comonadEnvTuple = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
}), "ComonadAsk0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_comonadAskTuple()
})})
	})
	return comonadEnvTuple
}

var comonadAskEnvT gopurs_runtime.Value
var once_comonadAskEnvT sync.Once
func Get_comonadAskEnvT() gopurs_runtime.Value {
	once_comonadAskEnvT.Do(func() {
		comonadAskEnvT = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), dictComonad_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return v_2.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
})})
})
	})
	return comonadAskEnvT
}

var comonadEnvEnvT gopurs_runtime.Value
var once_comonadEnvEnvT sync.Once
func Get_comonadEnvEnvT() gopurs_runtime.Value {
	once_comonadEnvEnvT.Do(func() {
		comonadEnvEnvT = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), dictComonad_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_2, v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
}), "ComonadAsk0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
})})
})})
})
	})
	return comonadEnvEnvT
}

var ask gopurs_runtime.Value
var once_ask sync.Once
func Get_ask() gopurs_runtime.Value {
	once_ask.Do(func() {
		ask = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["ask"]
})
	})
	return ask
}

var asks gopurs_runtime.Value
var once_asks sync.Once
func Get_asks() gopurs_runtime.Value {
	once_asks.Do(func() {
		asks = gopurs_runtime.Func(func(dictComonadAsk_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(dictComonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["ask"], x_2))
})
})
})
	})
	return asks
}

var comonadAskStoreT gopurs_runtime.Value
var once_comonadAskStoreT sync.Once
func Get_comonadAskStoreT() gopurs_runtime.Value {
	once_comonadAskStoreT.Do(func() {
		comonadAskStoreT = gopurs_runtime.Func(func(dictComonadAsk_0 gopurs_runtime.Value) gopurs_runtime.Value {
Comonad0_1_0 := gopurs_runtime.Apply(dictComonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{})
comonadStoreT_2_1 := gopurs_runtime.Apply(pkg_Control_Comonad_Store_Trans.Get_comonadStoreT(), Comonad0_1_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := x_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(dictComonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["ask"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Comonad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_5, __local_var_4_2)
})), x_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_2_1
})})
})
	})
	return comonadAskStoreT
}

var comonadEnvStoreT gopurs_runtime.Value
var once_comonadEnvStoreT sync.Once
func Get_comonadEnvStoreT() gopurs_runtime.Value {
	once_comonadEnvStoreT.Do(func() {
		comonadEnvStoreT = gopurs_runtime.Func(func(dictComonadEnv_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadAskStoreT1_1_0 := gopurs_runtime.Apply(Get_comonadAskStoreT(), gopurs_runtime.Apply(dictComonadEnv_0.PtrVal.(map[string]gopurs_runtime.Value)["ComonadAsk0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonadEnv_0.PtrVal.(map[string]gopurs_runtime.Value)["local"], f_2), v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), "value1": v_3.PtrVal.(map[string]gopurs_runtime.Value)["value1"]})
})
}), "ComonadAsk0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadAskStoreT1_1_0
})})
})
	})
	return comonadEnvStoreT
}

var comonadAskTracedT gopurs_runtime.Value
var once_comonadAskTracedT sync.Once
func Get_comonadAskTracedT() gopurs_runtime.Value {
	once_comonadAskTracedT.Do(func() {
		comonadAskTracedT = gopurs_runtime.Func(func(dictComonadAsk_0 gopurs_runtime.Value) gopurs_runtime.Value {
Comonad0_1_0 := gopurs_runtime.Apply(dictComonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{})
comonadTracedT_2_1 := gopurs_runtime.Apply(pkg_Control_Comonad_Traced_Trans.Get_comonadTracedT(), Comonad0_1_0)
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
comonadTracedT1_4_2 := gopurs_runtime.Apply(comonadTracedT_2_1, dictMonoid_3)
mempty_5_3 := dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictComonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["ask"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Comonad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, mempty_5_3)
})), x_6))
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_4_2
})})
})
})
	})
	return comonadAskTracedT
}

var comonadEnvTracedT gopurs_runtime.Value
var once_comonadEnvTracedT sync.Once
func Get_comonadEnvTracedT() gopurs_runtime.Value {
	once_comonadEnvTracedT.Do(func() {
		comonadEnvTracedT = gopurs_runtime.Func(func(dictComonadEnv_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadAskTracedT1_1_0 := gopurs_runtime.Apply(Get_comonadAskTracedT(), gopurs_runtime.Apply(dictComonadEnv_0.PtrVal.(map[string]gopurs_runtime.Value)["ComonadAsk0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
comonadAskTracedT2_3_1 := gopurs_runtime.Apply(comonadAskTracedT1_1_0, dictMonoid_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonadEnv_0.PtrVal.(map[string]gopurs_runtime.Value)["local"], f_4), v_5)
})
}), "ComonadAsk0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadAskTracedT2_3_1
})})
})
})
	})
	return comonadEnvTracedT
}



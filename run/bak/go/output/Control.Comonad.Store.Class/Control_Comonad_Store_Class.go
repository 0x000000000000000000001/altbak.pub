package Control_Comonad_Store_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
	pkg_Control_Comonad_Store_Trans "gopurs/output/Control.Comonad.Store.Trans"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
)

var pos gopurs_runtime.Value
var once_pos sync.Once
func Get_pos() gopurs_runtime.Value {
	once_pos.Do(func() {
		pos = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["pos"]
})
	})
	return pos
}

var peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		peek = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["peek"]
})
	})
	return peek
}

var peeks gopurs_runtime.Value
var once_peeks sync.Once
func Get_peeks() gopurs_runtime.Value {
	once_peeks.Do(func() {
		peeks = gopurs_runtime.Func(func(dictComonadStore_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["peek"], gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["pos"], x_2))), x_2)
})
})
})
	})
	return peeks
}

var seeks gopurs_runtime.Value
var once_seeks sync.Once
func Get_seeks() gopurs_runtime.Value {
	once_seeks.Do(func() {
		seeks = gopurs_runtime.Func(func(dictComonadStore_0 gopurs_runtime.Value) gopurs_runtime.Value {
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["extend"], pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(duplicate_1_0, x_3)
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["peek"], gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["pos"], __local_var_4_1))), __local_var_4_1)
})
})
})
	})
	return seeks
}

var seek gopurs_runtime.Value
var once_seek sync.Once
func Get_seek() gopurs_runtime.Value {
	once_seek.Do(func() {
		seek = gopurs_runtime.Func(func(dictComonadStore_0 gopurs_runtime.Value) gopurs_runtime.Value {
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["extend"], pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"])
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["peek"], s_2)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(duplicate_1_0, x_4))
})
})
})
	})
	return seek
}

var experiment gopurs_runtime.Value
var once_experiment sync.Once
func Get_experiment() gopurs_runtime.Value {
	once_experiment.Do(func() {
		experiment = gopurs_runtime.Func(func(dictComonadStore_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(dictFunctor_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_1.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["peek"], a_4), x_3)
})), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["pos"], x_3)))
})
})
})
})
	})
	return experiment
}

var comonadStoreTracedT gopurs_runtime.Value
var once_comonadStoreTracedT sync.Once
func Get_comonadStoreTracedT() gopurs_runtime.Value {
	once_comonadStoreTracedT.Do(func() {
		comonadStoreTracedT = gopurs_runtime.Func(func(dictComonadStore_0 gopurs_runtime.Value) gopurs_runtime.Value {
Comonad0_1_0 := gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{})
comonadTracedT_2_1 := gopurs_runtime.Apply(pkg_Control_Comonad_Traced_Trans.Get_comonadTracedT(), Comonad0_1_0)
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
mempty_4_2 := dictMonoid_3.PtrVal.(map[string]gopurs_runtime.Value)["mempty"]
lower1_5_3 := gopurs_runtime.Value{PtrVal: func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Comonad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_6, mempty_4_2)
})), v_5)
}}
comonadTracedT1_6_4 := gopurs_runtime.Apply(comonadTracedT_2_1, dictMonoid_3)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pos": gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["pos"], lower1_5_3.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(x_7))
}), "peek": gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_5 := gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["peek"], s_7)
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, lower1_5_3.PtrVal.(func(gopurs_runtime.Value) gopurs_runtime.Value)(x_9))
})
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_6_4
})})
})
})
	})
	return comonadStoreTracedT
}

var comonadStoreStoreT gopurs_runtime.Value
var once_comonadStoreStoreT sync.Once
func Get_comonadStoreStoreT() gopurs_runtime.Value {
	once_comonadStoreStoreT.Do(func() {
		comonadStoreStoreT = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadStoreT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Store_Trans.Get_comonadStoreT(), dictComonad_0)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pos": gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return v_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
}), "peek": gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["extract"], v_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"]), s_2)
})
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_1_0
})})
})
	})
	return comonadStoreStoreT
}

var comonadStoreEnvT gopurs_runtime.Value
var once_comonadStoreEnvT sync.Once
func Get_comonadStoreEnvT() gopurs_runtime.Value {
	once_comonadStoreEnvT.Do(func() {
		comonadStoreEnvT = gopurs_runtime.Func(func(dictComonadStore_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"pos": gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["pos"], x_2.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
}), "peek": gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictComonadStore_0.PtrVal.(map[string]gopurs_runtime.Value)["peek"], s_2)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
})})
})
	})
	return comonadStoreEnvT
}



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
return gopurs_runtime.RecordGet(dict_0, "pos")
})
	})
	return pos
}

var peek gopurs_runtime.Value
var once_peek sync.Once
func Get_peek() gopurs_runtime.Value {
	once_peek.Do(func() {
		peek = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "peek")
})
	})
	return peek
}

var peeks gopurs_runtime.Value
var once_peeks sync.Once
func Get_peeks() gopurs_runtime.Value {
	once_peeks.Do(func() {
		peeks = gopurs_runtime.Func3(func(dictComonadStore_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), x_2)), x_2)
})
	})
	return peeks
}

var seeks gopurs_runtime.Value
var once_seeks sync.Once
func Get_seeks() gopurs_runtime.Value {
	once_seeks.Do(func() {
		seeks = gopurs_runtime.Func(func(dictComonadStore_0 gopurs_runtime.Value) gopurs_runtime.Value {
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
_ = duplicate_1_0
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(duplicate_1_0, x_3)
_ = __local_var_4_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), __local_var_4_1)), __local_var_4_1)
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
duplicate_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "extend"), gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"))
_ = duplicate_1_0
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_2)
_ = __local_var_3_1
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
		experiment = gopurs_runtime.Func4(func(dictComonadStore_0 gopurs_runtime.Value, dictFunctor_1 gopurs_runtime.Value, f_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_1, "map"), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), a_4, x_3)
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), x_3)))
})
	})
	return experiment
}

var comonadStoreTracedT gopurs_runtime.Value
var once_comonadStoreTracedT sync.Once
func Get_comonadStoreTracedT() gopurs_runtime.Value {
	once_comonadStoreTracedT.Do(func() {
		comonadStoreTracedT = gopurs_runtime.Func(func(dictComonadStore_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
return gopurs_runtime.RecordDict3("pos", "peek", "Comonad0", gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), gopurs_runtime.UncurriedApp(lower1_5_3, x_7))
}), gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_8_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_7)
_ = __local_var_8_5
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_8_5, gopurs_runtime.UncurriedApp(lower1_5_3, x_9))
})
}), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_6_4
}))
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
_ = comonadStoreT_1_0
return gopurs_runtime.RecordDict3("pos", "peek", "Comonad0", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(v_2, "value1")
}), gopurs_runtime.Func2(func(s_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), gopurs_runtime.RecordGet(v_3, "value0"), s_2)
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_1_0
}))
})
	})
	return comonadStoreStoreT
}

var comonadStoreEnvT gopurs_runtime.Value
var once_comonadStoreEnvT sync.Once
func Get_comonadStoreEnvT() gopurs_runtime.Value {
	once_comonadStoreEnvT.Do(func() {
		comonadStoreEnvT = gopurs_runtime.Func(func(dictComonadStore_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "Comonad0"), gopurs_runtime.Value{}))
_ = comonadEnvT_1_0
return gopurs_runtime.RecordDict3("pos", "peek", "Comonad0", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "pos"), gopurs_runtime.RecordGet(x_2, "value1"))
}), gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadStore_0, "peek"), s_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.RecordGet(x_4, "value1"))
})
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
}))
})
	})
	return comonadStoreEnvT
}



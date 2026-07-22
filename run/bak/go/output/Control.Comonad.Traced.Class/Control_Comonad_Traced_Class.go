package Control_Comonad_Traced_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Comonad_Trans_Class "gopurs/output/Control.Comonad.Trans.Class"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
	pkg_Control_Comonad_Store_Trans "gopurs/output/Control.Comonad.Store.Trans"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
)

var track gopurs_runtime.Value
var once_track sync.Once
func Get_track() gopurs_runtime.Value {
	once_track.Do(func() {
		track = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["track"]
})
	})
	return track
}

var tracks gopurs_runtime.Value
var once_tracks sync.Once
func Get_tracks() gopurs_runtime.Value {
	once_tracks.Do(func() {
		tracks = gopurs_runtime.Func(func(dictComonadTraced_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["track"], gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["extract"], w_2))), w_2)
})
})
})
	})
	return tracks
}

var lowerTrack2 gopurs_runtime.Value
var once_lowerTrack2 sync.Once
func Get_lowerTrack2() gopurs_runtime.Value {
	once_lowerTrack2.Do(func() {
		lowerTrack2 = gopurs_runtime.Func(func(dictComonadTraced_0 gopurs_runtime.Value) gopurs_runtime.Value {
lower1_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Trans_Class.Get_comonadTransIdentityT().PtrVal.(map[string]gopurs_runtime.Value)["lower"], gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{}))
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["track"], m_2)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(lower1_1_0, x_4))
})
})
})
	})
	return lowerTrack2
}

var listens gopurs_runtime.Value
var once_listens sync.Once
func Get_listens() gopurs_runtime.Value {
	once_listens.Do(func() {
		listens = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(g_3, t_4), "value1": gopurs_runtime.Apply(f_1, t_4)})
})
})), v_2)
})
})
})
	})
	return listens
}

var listen gopurs_runtime.Value
var once_listen sync.Once
func Get_listen() gopurs_runtime.Value {
	once_listen.Do(func() {
		listen = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": gopurs_runtime.Apply(f_2, t_3), "value1": t_3})
})
})), v_1)
})
})
	})
	return listen
}

var comonadTracedTracedT gopurs_runtime.Value
var once_comonadTracedTracedT sync.Once
func Get_comonadTracedTracedT() gopurs_runtime.Value {
	once_comonadTracedTracedT.Do(func() {
		comonadTracedTracedT = gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadTracedT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Traced_Trans.Get_comonadTracedT(), dictComonad_0)
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
comonadTracedT1_3_1 := gopurs_runtime.Apply(comonadTracedT_1_0, dictMonoid_2)
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"track": gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictComonad_0.PtrVal.(map[string]gopurs_runtime.Value)["extract"], v_5), t_4)
})
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_3_1
})})
})
})
	})
	return comonadTracedTracedT
}

var comonadTracedStoreT gopurs_runtime.Value
var once_comonadTracedStoreT sync.Once
func Get_comonadTracedStoreT() gopurs_runtime.Value {
	once_comonadTracedStoreT.Do(func() {
		comonadTracedStoreT = gopurs_runtime.Func(func(dictComonadTraced_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadStoreT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Store_Trans.Get_comonadStoreT(), gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{}))
__local_var_2_1 := gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"track": gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["track"], m_3)
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := x_5.PtrVal.(map[string]gopurs_runtime.Value)["value1"]
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_7, __local_var_6_3)
})), x_5.PtrVal.(map[string]gopurs_runtime.Value)["value0"]))
})
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_1_0
})})
})
	})
	return comonadTracedStoreT
}

var comonadTracedIdentityT gopurs_runtime.Value
var once_comonadTracedIdentityT sync.Once
func Get_comonadTracedIdentityT() gopurs_runtime.Value {
	once_comonadTracedIdentityT.Do(func() {
		comonadTracedIdentityT = gopurs_runtime.Func(func(dictComonadTraced_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{})
__local_var_2_1 := gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Extend0"], gopurs_runtime.Value{})
__local_var_3_3 := gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{})
extendIdentityI1_4_4 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extend": gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(__local_var_2_1.PtrVal.(map[string]gopurs_runtime.Value)["extend"], gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
})), v_5)
})
}), "Functor0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_3
})})
comonadIdentityT_3_2 := gopurs_runtime.Record(map[string]gopurs_runtime.Value{"extract": gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0.PtrVal.(map[string]gopurs_runtime.Value)["extract"], x_5)
}), "Extend0": gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return extendIdentityI1_4_4
})})
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"track": gopurs_runtime.Apply(Get_lowerTrack2(), dictComonadTraced_0), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadIdentityT_3_2
})})
})
	})
	return comonadTracedIdentityT
}

var comonadTracedEnvT gopurs_runtime.Value
var once_comonadTracedEnvT sync.Once
func Get_comonadTracedEnvT() gopurs_runtime.Value {
	once_comonadTracedEnvT.Do(func() {
		comonadTracedEnvT = gopurs_runtime.Func(func(dictComonadTraced_0 gopurs_runtime.Value) gopurs_runtime.Value {
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["Comonad0"], gopurs_runtime.Value{}))
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"track": gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(dictComonadTraced_0.PtrVal.(map[string]gopurs_runtime.Value)["track"], m_2)
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4.PtrVal.(map[string]gopurs_runtime.Value)["value1"])
})
}), "Comonad0": gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
})})
})
	})
	return comonadTracedEnvT
}

var censor gopurs_runtime.Value
var once_censor sync.Once
func Get_censor() gopurs_runtime.Value {
	once_censor.Do(func() {
		censor = gopurs_runtime.Func(func(dictFunctor_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(dictFunctor_0.PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Apply(f_1, x_4))
})
})), v_2)
})
})
})
	})
	return censor
}



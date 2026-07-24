package Control_Comonad_Traced_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Comonad_Trans_Class "gopurs/output/Control.Comonad.Trans.Class"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
	pkg_Control_Comonad_Store_Trans "gopurs/output/Control.Comonad.Store.Trans"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
	unsafe "unsafe"
)

var track gopurs_runtime.Value
var once_track sync.Once
func Get_track() gopurs_runtime.Value {
	once_track.Do(func() {
		track = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "track")
}()
})
	})
	return track
}

var tracks gopurs_runtime.Value
var once_tracks sync.Once
func Get_tracks() gopurs_runtime.Value {
	once_tracks.Do(func() {
		tracks = gopurs_runtime.Func3(func(dictComonadTraced_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, w_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tracks(dictComonadTraced_0_box, f_1_box, w_2_box)
})
	})
	return tracks
}

var lowerTrack2 gopurs_runtime.Value
var once_lowerTrack2 sync.Once
func Get_lowerTrack2() gopurs_runtime.Value {
	once_lowerTrack2.Do(func() {
		lowerTrack2 = gopurs_runtime.Func(func(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
lower1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Comonad_Trans_Class.Get_comonadTransIdentityT(), "lower"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = lower1_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "track"), m_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(lower1_1_0, x_4))
})
})
}()
})
	})
	return lowerTrack2
}

var listens gopurs_runtime.Value
var once_listens sync.Once
func Get_listens() gopurs_runtime.Value {
	once_listens.Do(func() {
		listens = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listens(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return listens
}

var listen gopurs_runtime.Value
var once_listen sync.Once
func Get_listen() gopurs_runtime.Value {
	once_listen.Do(func() {
		listen = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listen(dictFunctor_0_box, v_1_box)
})
	})
	return listen
}

var comonadTracedTracedT gopurs_runtime.Value
var once_comonadTracedTracedT sync.Once
func Get_comonadTracedTracedT() gopurs_runtime.Value {
	once_comonadTracedTracedT.Do(func() {
		comonadTracedTracedT = gopurs_runtime.Func(func(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
comonadTracedT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Traced_Trans.Get_comonadTracedT(), dictComonad_0)
_ = comonadTracedT_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
comonadTracedT1_3_1 := gopurs_runtime.Apply(comonadTracedT_1_0, dictMonoid_2)
_ = comonadTracedT1_3_1
return gopurs_runtime.RecordDict2("track", "Comonad0", gopurs_runtime.Func2(func(t_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), v_5, t_4)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_3_1
}))
})
}()
})
	})
	return comonadTracedTracedT
}

var comonadTracedStoreT gopurs_runtime.Value
var once_comonadTracedStoreT sync.Once
func Get_comonadTracedStoreT() gopurs_runtime.Value {
	once_comonadTracedStoreT.Do(func() {
		comonadTracedStoreT = gopurs_runtime.Func(func(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
comonadStoreT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Store_Trans.Get_comonadStoreT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = comonadStoreT_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.RecordDict2("track", "Comonad0", gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "track"), m_3)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(x_5.UnsafePtr).V1
_ = __local_var_6_3
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_7, __local_var_6_3)
}), (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(x_5.UnsafePtr).V0))
})
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_1_0
}))
}()
})
	})
	return comonadTracedStoreT
}

var comonadTracedIdentityT gopurs_runtime.Value
var once_comonadTracedIdentityT sync.Once
func Get_comonadTracedIdentityT() gopurs_runtime.Value {
	once_comonadTracedIdentityT.Do(func() {
		comonadTracedIdentityT = gopurs_runtime.Func(func(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_2_1
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_3
extendIdentityI1_4_4 := gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), v_5)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_3_3
}))
_ = extendIdentityI1_4_4
comonadIdentityT_3_2 := gopurs_runtime.RecordDict2("extract", "Extend0", gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "extract"), x_5)
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return extendIdentityI1_4_4
}))
_ = comonadIdentityT_3_2
return gopurs_runtime.RecordDict2("track", "Comonad0", gopurs_runtime.Apply(Get_lowerTrack2(), dictComonadTraced_0), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadIdentityT_3_2
}))
}()
})
	})
	return comonadTracedIdentityT
}

var comonadTracedEnvT gopurs_runtime.Value
var once_comonadTracedEnvT sync.Once
func Get_comonadTracedEnvT() gopurs_runtime.Value {
	once_comonadTracedEnvT.Do(func() {
		comonadTracedEnvT = gopurs_runtime.Func(func(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = comonadEnvT_1_0
return gopurs_runtime.RecordDict2("track", "Comonad0", gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "track"), m_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(x_4.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
}))
}()
})
	})
	return comonadTracedEnvT
}

var censor gopurs_runtime.Value
var once_censor sync.Once
func Get_censor() gopurs_runtime.Value {
	once_censor.Do(func() {
		censor = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_censor(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return censor
}

func Call_tracks(dictComonadTraced_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, w_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var w_2 gopurs_runtime.Value = w_2_loop
_ = w_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadTraced_0, "track"), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}), "extract"), w_2)), w_2)
}

func Call_listens(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func2(func(g_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply(g_3, t_4), gopurs_runtime.Apply(f_1, t_4)})}
}), v_2)
}

func Call_listen(dictFunctor_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1102100576, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{gopurs_runtime.Apply(f_2, t_3), t_3})}
}), v_1)
}

func Call_censor(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func2(func(v1_3 gopurs_runtime.Value, x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Apply(f_1, x_4))
}), v_2)
}



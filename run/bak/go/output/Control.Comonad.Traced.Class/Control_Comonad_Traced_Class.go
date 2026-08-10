package Control_Comonad_Traced_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Comonad_Store_Trans "gopurs/output/Control.Comonad.Store.Trans"
	pkg_Control_Comonad_Trans_Class "gopurs/output/Control.Comonad.Trans.Class"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
	unsafe "unsafe"
)

var cache_track gopurs_runtime.Value
var once_track sync.Once
func Get_track() gopurs_runtime.Value {
	once_track.Do(func() {
		cache_track = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_track(dict_0_box)
})
	})
	return cache_track
}

var cache_track__gopurs_runtime_Value_250573732 gopurs_runtime.Value
var once_track__gopurs_runtime_Value_250573732 sync.Once
func Get_track__gopurs_runtime_Value_250573732() gopurs_runtime.Value {
	once_track__gopurs_runtime_Value_250573732.Do(func() {
		cache_track__gopurs_runtime_Value_250573732 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_track__gopurs_runtime_Value_250573732(dict_0_box)
})
	})
	return cache_track__gopurs_runtime_Value_250573732
}

var cache_tracks gopurs_runtime.Value
var once_tracks sync.Once
func Get_tracks() gopurs_runtime.Value {
	once_tracks.Do(func() {
		cache_tracks = gopurs_runtime.Func3(func(dictComonadTraced_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, w_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tracks(dictComonadTraced_0_box, f_1_box, w_2_box)
})
	})
	return cache_tracks
}

var cache_lowerTrack1 gopurs_runtime.Value
var once_lowerTrack1 sync.Once
func Get_lowerTrack1() gopurs_runtime.Value {
	once_lowerTrack1.Do(func() {
		cache_lowerTrack1 = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lowerTrack1(dictComonadTraced_0_box)
})
	})
	return cache_lowerTrack1
}

var cache_lowerTrack2 gopurs_runtime.Value
var once_lowerTrack2 sync.Once
func Get_lowerTrack2() gopurs_runtime.Value {
	once_lowerTrack2.Do(func() {
		cache_lowerTrack2 = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lowerTrack2(dictComonadTraced_0_box)
})
	})
	return cache_lowerTrack2
}

var cache_lowerTrack3 gopurs_runtime.Value
var once_lowerTrack3 sync.Once
func Get_lowerTrack3() gopurs_runtime.Value {
	once_lowerTrack3.Do(func() {
		cache_lowerTrack3 = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lowerTrack3(dictComonadTraced_0_box)
})
	})
	return cache_lowerTrack3
}

var cache_listens gopurs_runtime.Value
var once_listens sync.Once
func Get_listens() gopurs_runtime.Value {
	once_listens.Do(func() {
		cache_listens = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listens(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return cache_listens
}

var cache_listen gopurs_runtime.Value
var once_listen sync.Once
func Get_listen() gopurs_runtime.Value {
	once_listen.Do(func() {
		cache_listen = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listen(dictFunctor_0_box, v_1_box)
})
	})
	return cache_listen
}

var cache_comonadTracedTracedT gopurs_runtime.Value
var once_comonadTracedTracedT sync.Once
func Get_comonadTracedTracedT() gopurs_runtime.Value {
	once_comonadTracedTracedT.Do(func() {
		cache_comonadTracedTracedT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadTracedTracedT(dictComonad_0_box)
})
	})
	return cache_comonadTracedTracedT
}

var cache_comonadTracedStoreT gopurs_runtime.Value
var once_comonadTracedStoreT sync.Once
func Get_comonadTracedStoreT() gopurs_runtime.Value {
	once_comonadTracedStoreT.Do(func() {
		cache_comonadTracedStoreT = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadTracedStoreT(dictComonadTraced_0_box)
})
	})
	return cache_comonadTracedStoreT
}

var cache_comonadTracedIdentityT gopurs_runtime.Value
var once_comonadTracedIdentityT sync.Once
func Get_comonadTracedIdentityT() gopurs_runtime.Value {
	once_comonadTracedIdentityT.Do(func() {
		cache_comonadTracedIdentityT = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadTracedIdentityT(dictComonadTraced_0_box)
})
	})
	return cache_comonadTracedIdentityT
}

var cache_comonadTracedEnvT gopurs_runtime.Value
var once_comonadTracedEnvT sync.Once
func Get_comonadTracedEnvT() gopurs_runtime.Value {
	once_comonadTracedEnvT.Do(func() {
		cache_comonadTracedEnvT = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadTracedEnvT(dictComonadTraced_0_box)
})
	})
	return cache_comonadTracedEnvT
}

var cache_censor gopurs_runtime.Value
var once_censor sync.Once
func Get_censor() gopurs_runtime.Value {
	once_censor.Do(func() {
		cache_censor = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_censor(dictFunctor_0_box, f_1_box, v_2_box)
})
	})
	return cache_censor
}

func Call_track(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "track")
}

func Call_track__gopurs_runtime_Value_250573732(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "track")
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

func Call_lowerTrack1(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
lower1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Comonad_Store_Trans.Get_comonadTransStoreT(), "lower"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = lower1_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "track"), m_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(lower1_1_0, x_4))
})
})
}

func Call_lowerTrack2(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
}

func Call_lowerTrack3(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
lower1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Comonad_Env_Trans.Get_comonadTransEnvT(), "lower"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = lower1_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "track"), m_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(lower1_1_0, x_4))
})
})
}

func Call_listens(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(g_3, t_4), gopurs_runtime.Apply(f_1, t_4)})}
})
}), v_2)
}

func Call_listen(dictFunctor_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, t_3), t_3})}
})
}), v_1)
}

func Call_comonadTracedTracedT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
comonadTracedT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Traced_Trans.Get_comonadTracedT(), dictComonad_0)
_ = comonadTracedT_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
comonadTracedT1_3_1 := gopurs_runtime.Apply(comonadTracedT_1_0, dictMonoid_2)
_ = comonadTracedT1_3_1
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadTracedT1_3_1
}), gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonad_0, "extract"), v_5, t_4)
})
}))
})
}

func Call_comonadTracedStoreT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
comonadStoreT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Store_Trans.Get_comonadStoreT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = comonadStoreT_1_0
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_1_0
}), Call_lowerTrack1(dictComonadTraced_0))
}

func Call_comonadTracedIdentityT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorIdentityT1_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = functorIdentityT1_3_4
extendIdentityI1_2_2 := gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorIdentityT1_3_4
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), v_5)
})
}))
_ = extendIdentityI1_2_2
comonadIdentityT_1_0 := gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return extendIdentityI1_2_2
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "extract"), x_3)
}))
_ = comonadIdentityT_1_0
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadIdentityT_1_0
}), Call_lowerTrack2(dictComonadTraced_0))
}

func Call_comonadTracedEnvT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = comonadEnvT_1_0
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
}), Call_lowerTrack3(dictComonadTraced_0))
}

func Call_censor(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Apply(f_1, x_4))
})
}), v_2)
}



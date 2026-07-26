package Control_Comonad_Traced_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Comonad_Traced_Trans "gopurs/output/Control.Comonad.Traced.Trans"
	pkg_Control_Comonad_Store_Trans "gopurs/output/Control.Comonad.Store.Trans"
	pkg_Control_Monad_Identity_Trans "gopurs/output/Control.Monad.Identity.Trans"
	pkg_Control_Comonad_Trans_Class "gopurs/output/Control.Comonad.Trans.Class"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
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

var cache_lowerTrack gopurs_runtime.Value
var once_lowerTrack sync.Once
func Get_lowerTrack() gopurs_runtime.Value {
	once_lowerTrack.Do(func() {
		cache_lowerTrack = gopurs_runtime.Func2(func(dictComonadTrans_0_box gopurs_runtime.Value, dictComonadTraced_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lowerTrack(dictComonadTrans_0_box, dictComonadTraced_1_box)
})
	})
	return cache_lowerTrack
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
return ((*gopurs_runtime.RecordData1)(dict_0.UnsafePtr)).V0
}

func Call_tracks(dictComonadTraced_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, w_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var w_2 gopurs_runtime.Value = w_2_loop
_ = w_2
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictComonadTraced_0.UnsafePtr)).V0, gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0_NOT_FOUND"), gopurs_runtime.Value{}), "extract"), w_2)), w_2)
}

func Call_lowerTrack(dictComonadTrans_0_loop gopurs_runtime.Value, dictComonadTraced_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTrans_0 gopurs_runtime.Value = dictComonadTrans_0_loop
_ = dictComonadTrans_0
var dictComonadTraced_1 gopurs_runtime.Value = dictComonadTraced_1_loop
_ = dictComonadTraced_1
lower1_2_0 := gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictComonadTrans_0.UnsafePtr)).V0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_1, "Comonad0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = lower1_2_0
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictComonadTraced_1.UnsafePtr)).V0, m_3), lower1_2_0)
})
}

func Call_listens(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func2(func(g_3 gopurs_runtime.Value, t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(g_3, t_4), gopurs_runtime.Apply(f_1, t_4)})}
}), v_2)
}

func Call_listen(dictFunctor_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_2, t_3), t_3})}
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
}), gopurs_runtime.Func2(func(t_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictComonad_0.UnsafePtr)).V0, v_5, t_4)
}))
})
}

func Call_comonadTracedStoreT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
comonadStoreT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Store_Trans.Get_comonadStoreT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = comonadStoreT_1_0
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_1_0
}), Call_lowerTrack(pkg_Control_Comonad_Store_Trans.Get_comonadTransStoreT(), dictComonadTraced_0))
}

func Call_comonadTracedIdentityT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
comonadIdentityT_1_0 := gopurs_runtime.Apply(pkg_Control_Monad_Identity_Trans.Get_comonadIdentityT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = comonadIdentityT_1_0
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadIdentityT_1_0
}), Call_lowerTrack(pkg_Control_Comonad_Trans_Class.Get_comonadTransIdentityT(), dictComonadTraced_0))
}

func Call_comonadTracedEnvT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
comonadEnvT_1_0 := gopurs_runtime.Apply(pkg_Control_Comonad_Env_Trans.Get_comonadEnvT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0_NOT_FOUND"), gopurs_runtime.Value{}))
_ = comonadEnvT_1_0
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
}), Call_lowerTrack(pkg_Control_Comonad_Env_Trans.Get_comonadTransEnvT(), dictComonadTraced_0))
}

func Call_censor(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(((*gopurs_runtime.RecordData1)(dictFunctor_0.UnsafePtr)).V0, gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), v1_3, f_1)
}), v_2)
}



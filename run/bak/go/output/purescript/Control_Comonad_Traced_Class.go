package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Traced_Class_ComonadTraced_dollarDict gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_ComonadTraced_dollarDict sync.Once
func Get_Control_Comonad_Traced_Class_ComonadTraced_dollarDict() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_ComonadTraced_dollarDict.Do(func() {
		cache_Control_Comonad_Traced_Class_ComonadTraced_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_ComonadTraced_dollarDict(x_0_box)
})
	})
	return cache_Control_Comonad_Traced_Class_ComonadTraced_dollarDict
}

var cache_Control_Comonad_Traced_Class_track gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_track sync.Once
func Get_Control_Comonad_Traced_Class_track() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_track.Do(func() {
		cache_Control_Comonad_Traced_Class_track = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_track(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced](dict_0_box))
})
	})
	return cache_Control_Comonad_Traced_Class_track
}

var cache_Control_Comonad_Traced_Class_tracks gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_tracks sync.Once
func Get_Control_Comonad_Traced_Class_tracks() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_tracks.Do(func() {
		cache_Control_Comonad_Traced_Class_tracks = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_tracks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced](dictComonadTraced_0_box))
})
	})
	return cache_Control_Comonad_Traced_Class_tracks
}

var cache_Control_Comonad_Traced_Class_lowerTrack gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_lowerTrack sync.Once
func Get_Control_Comonad_Traced_Class_lowerTrack() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_lowerTrack.Do(func() {
		cache_Control_Comonad_Traced_Class_lowerTrack = gopurs_runtime.Func2(func(dictComonadTrans_0_box gopurs_runtime.Value, dictComonadTraced_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_lowerTrack(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans](dictComonadTrans_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced](dictComonadTraced_1_box))
})
	})
	return cache_Control_Comonad_Traced_Class_lowerTrack
}

var cache_Control_Comonad_Traced_Class_lowerTrack1 gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_lowerTrack1 sync.Once
func Get_Control_Comonad_Traced_Class_lowerTrack1() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_lowerTrack1.Do(func() {
		cache_Control_Comonad_Traced_Class_lowerTrack1 = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_lowerTrack1(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced](dictComonadTraced_0_box))
})
	})
	return cache_Control_Comonad_Traced_Class_lowerTrack1
}

var cache_Control_Comonad_Traced_Class_lowerTrack2 gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_lowerTrack2 sync.Once
func Get_Control_Comonad_Traced_Class_lowerTrack2() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_lowerTrack2.Do(func() {
		cache_Control_Comonad_Traced_Class_lowerTrack2 = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_lowerTrack2(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced](dictComonadTraced_0_box))
})
	})
	return cache_Control_Comonad_Traced_Class_lowerTrack2
}

var cache_Control_Comonad_Traced_Class_lowerTrack3 gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_lowerTrack3 sync.Once
func Get_Control_Comonad_Traced_Class_lowerTrack3() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_lowerTrack3.Do(func() {
		cache_Control_Comonad_Traced_Class_lowerTrack3 = gopurs_runtime.Func2(func(dictComonadTraced_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_lowerTrack3(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced](dictComonadTraced_0_box), m_1_box)
})
	})
	return cache_Control_Comonad_Traced_Class_lowerTrack3
}

var cache_Control_Comonad_Traced_Class_listens gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_listens sync.Once
func Get_Control_Comonad_Traced_Class_listens() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_listens.Do(func() {
		cache_Control_Comonad_Traced_Class_listens = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_listens(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Comonad_Traced_Class_listens
}

var cache_Control_Comonad_Traced_Class_listen gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_listen sync.Once
func Get_Control_Comonad_Traced_Class_listen() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_listen.Do(func() {
		cache_Control_Comonad_Traced_Class_listen = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_listen(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), v_1_box)
})
	})
	return cache_Control_Comonad_Traced_Class_listen
}

var cache_Control_Comonad_Traced_Class_comonadTracedTracedT gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_comonadTracedTracedT sync.Once
func Get_Control_Comonad_Traced_Class_comonadTracedTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_comonadTracedTracedT.Do(func() {
		cache_Control_Comonad_Traced_Class_comonadTracedTracedT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_comonadTracedTracedT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Traced_Class_comonadTracedTracedT
}

var cache_Control_Comonad_Traced_Class_comonadTracedStoreT gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_comonadTracedStoreT sync.Once
func Get_Control_Comonad_Traced_Class_comonadTracedStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_comonadTracedStoreT.Do(func() {
		cache_Control_Comonad_Traced_Class_comonadTracedStoreT = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_comonadTracedStoreT(dictComonadTraced_0_box)
})
	})
	return cache_Control_Comonad_Traced_Class_comonadTracedStoreT
}

var cache_Control_Comonad_Traced_Class_comonadTracedIdentityT gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_comonadTracedIdentityT sync.Once
func Get_Control_Comonad_Traced_Class_comonadTracedIdentityT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_comonadTracedIdentityT.Do(func() {
		cache_Control_Comonad_Traced_Class_comonadTracedIdentityT = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_comonadTracedIdentityT(dictComonadTraced_0_box)
})
	})
	return cache_Control_Comonad_Traced_Class_comonadTracedIdentityT
}

var cache_Control_Comonad_Traced_Class_comonadTracedEnvT gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_comonadTracedEnvT sync.Once
func Get_Control_Comonad_Traced_Class_comonadTracedEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_comonadTracedEnvT.Do(func() {
		cache_Control_Comonad_Traced_Class_comonadTracedEnvT = gopurs_runtime.Func(func(dictComonadTraced_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_comonadTracedEnvT(dictComonadTraced_0_box)
})
	})
	return cache_Control_Comonad_Traced_Class_comonadTracedEnvT
}

var cache_Control_Comonad_Traced_Class_censor gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_censor sync.Once
func Get_Control_Comonad_Traced_Class_censor() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_censor.Do(func() {
		cache_Control_Comonad_Traced_Class_censor = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, v_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_censor(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](dictFunctor_0_box), f_1_box, v_2_box)
})
	})
	return cache_Control_Comonad_Traced_Class_censor
}

var cache_Control_Comonad_Traced_Class_track__250573732 gopurs_runtime.Value
var once_Control_Comonad_Traced_Class_track__250573732 sync.Once
func Get_Control_Comonad_Traced_Class_track__250573732() gopurs_runtime.Value {
	once_Control_Comonad_Traced_Class_track__250573732.Do(func() {
		cache_Control_Comonad_Traced_Class_track__250573732 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_Class_track__250573732(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Traced_Class_ComonadTraced](dict_0_box))
})
	})
	return cache_Control_Comonad_Traced_Class_track__250573732
}

type Constructor_Control_Comonad_Traced_Class_ComonadTraced struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[433046755] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Traced_Class_ComonadTraced)(ptr)
		_ = c
		switch key {
		case "Comonad0": return gopurs_runtime.Box(c.V0)
		case "track": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Traced_Class_ComonadTraced: " + key)
		}
	}
}


func Call_Control_Comonad_Traced_Class_ComonadTraced_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Traced_Class_track(dict_0_loop *Constructor_Control_Comonad_Traced_Class_ComonadTraced) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Traced_Class_ComonadTraced = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Traced_Class_tracks(dictComonadTraced_0_loop *Constructor_Control_Comonad_Traced_Class_ComonadTraced) gopurs_runtime.Value {
var dictComonadTraced_0 *Constructor_Control_Comonad_Traced_Class_ComonadTraced = dictComonadTraced_0_loop
_ = dictComonadTraced_0
// TAST (Let): Comonad0_1_0 -> *Constructor_Control_Comonad_Comonad
Comonad0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad](gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadTraced_0.V0), gopurs_runtime.Value{}))
_ = Comonad0_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(w_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictComonadTraced_0.V1), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(gopurs_runtime.Box(Comonad0_1_0.V1), w_3)), w_3)
})
})
}

func Call_Control_Comonad_Traced_Class_lowerTrack(dictComonadTrans_0_loop *Constructor_Control_Comonad_Trans_Class_ComonadTrans, dictComonadTraced_1_loop *Constructor_Control_Comonad_Traced_Class_ComonadTraced) gopurs_runtime.Value {
var dictComonadTrans_0 *Constructor_Control_Comonad_Trans_Class_ComonadTrans = dictComonadTrans_0_loop
_ = dictComonadTrans_0
var dictComonadTraced_1 *Constructor_Control_Comonad_Traced_Class_ComonadTraced = dictComonadTraced_1_loop
_ = dictComonadTraced_1
// TAST (Let): lower1_2_0 -> gopurs_runtime.Value
lower1_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadTrans_0.V0), gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadTraced_1.V0), gopurs_runtime.Value{}))
_ = lower1_2_0
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadTraced_1.V1), m_3)
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(lower1_2_0, x_5))
})
})
}

func Call_Control_Comonad_Traced_Class_lowerTrack1(dictComonadTraced_0_loop *Constructor_Control_Comonad_Traced_Class_ComonadTraced) gopurs_runtime.Value {
var dictComonadTraced_0 *Constructor_Control_Comonad_Traced_Class_ComonadTraced = dictComonadTraced_0_loop
_ = dictComonadTraced_0
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadTraced_0.V0), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
// TAST (Let): lower1_1_0 -> gopurs_runtime.Value
lower1_1_0 := gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_2 -> gopurs_runtime.Value
__local_var_3_2 := (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1
_ = __local_var_3_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_4, __local_var_3_2)
}), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0)
})
_ = lower1_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadTraced_0.V1), m_2)
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_3, gopurs_runtime.Apply(lower1_1_0, x_4))
})
})
}

func Call_Control_Comonad_Traced_Class_lowerTrack2(dictComonadTraced_0_loop *Constructor_Control_Comonad_Traced_Class_ComonadTraced) gopurs_runtime.Value {
var dictComonadTraced_0 *Constructor_Control_Comonad_Traced_Class_ComonadTraced = dictComonadTraced_0_loop
_ = dictComonadTraced_0
// TAST (Let): lower1_1_0 -> gopurs_runtime.Value
lower1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Comonad_Trans_Class_comonadTransIdentityT(), "lower"), gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadTraced_0.V0), gopurs_runtime.Value{}))
_ = lower1_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadTraced_0.V1), m_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, gopurs_runtime.Apply(lower1_1_0, x_4))
})
})
}

func Call_Control_Comonad_Traced_Class_lowerTrack3(dictComonadTraced_0_loop *Constructor_Control_Comonad_Traced_Class_ComonadTraced, m_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 *Constructor_Control_Comonad_Traced_Class_ComonadTraced = dictComonadTraced_0_loop
_ = dictComonadTraced_0
var m_1 gopurs_runtime.Value = m_1_loop
_ = m_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadTraced_0.V1), m_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, (*Constructor_Data_Tuple_Tuple)(x_3.UnsafePtr).V1)
})
}

func Call_Control_Comonad_Traced_Class_listens(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(g_3, t_4), gopurs_runtime.Apply(f_1, t_4)})}
})
}), v_2)
}

func Call_Control_Comonad_Traced_Class_listen(dictFunctor_0_loop *Constructor_Data_Functor_Functor, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_2, t_3), t_3})}
})
}), v_1)
}

func Call_Control_Comonad_Traced_Class_comonadTracedTracedT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): comonadTracedT_1_0 -> gopurs_runtime.Value
comonadTracedT_1_0 := gopurs_runtime.Apply(Get_Control_Comonad_Traced_Trans_comonadTracedT(), dictComonad_0)
_ = comonadTracedT_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): comonadTracedT1_3_1 -> gopurs_runtime.Value
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

func Call_Control_Comonad_Traced_Class_comonadTracedStoreT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
// TAST (Let): comonadStoreT_1_0 -> gopurs_runtime.Value
comonadStoreT_1_0 := gopurs_runtime.Apply(Get_Control_Comonad_Store_Trans_comonadStoreT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = comonadStoreT_1_0
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}), "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): lower1_2_1 -> gopurs_runtime.Value
lower1_2_1 := gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_3 -> gopurs_runtime.Value
__local_var_4_3 := (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1
_ = __local_var_4_3
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_5, __local_var_4_3)
}), (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0)
})
_ = lower1_2_1
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadStoreT_1_0
}), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_4 -> gopurs_runtime.Value
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "track"), m_3)
_ = __local_var_4_4
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_4, gopurs_runtime.Apply(lower1_2_1, x_5))
})
}))
}

func Call_Control_Comonad_Traced_Class_comonadTracedIdentityT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorIdentityT1_3_4 -> gopurs_runtime.Value
functorIdentityT1_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_3, "Functor0"), gopurs_runtime.Value{})
_ = functorIdentityT1_3_4
// TAST (Let): extendIdentityI1_2_2 -> gopurs_runtime.Value
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
// TAST (Let): comonadIdentityT_1_0 -> gopurs_runtime.Value
comonadIdentityT_1_0 := gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return extendIdentityI1_2_2
}), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "extract"), x_3)
}))
_ = comonadIdentityT_1_0
// TAST (Let): lower1_2_5 -> gopurs_runtime.Value
lower1_2_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Comonad_Trans_Class_comonadTransIdentityT(), "lower"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = lower1_2_5
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadIdentityT_1_0
}), gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_6 -> gopurs_runtime.Value
__local_var_4_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "track"), m_3)
_ = __local_var_4_6
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_6, gopurs_runtime.Apply(lower1_2_5, x_5))
})
}))
}

func Call_Control_Comonad_Traced_Class_comonadTracedEnvT(dictComonadTraced_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadTraced_0 gopurs_runtime.Value = dictComonadTraced_0_loop
_ = dictComonadTraced_0
// TAST (Let): comonadEnvT_1_0 -> gopurs_runtime.Value
comonadEnvT_1_0 := gopurs_runtime.Apply(Get_Control_Comonad_Env_Trans_comonadEnvT(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "Comonad0"), gopurs_runtime.Value{}))
_ = comonadEnvT_1_0
return gopurs_runtime.RecordDict2("Comonad0", "track", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return comonadEnvT_1_0
}), gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadTraced_0, "track"), m_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, (*Constructor_Data_Tuple_Tuple)(x_4.UnsafePtr).V1)
})
}))
}

func Call_Control_Comonad_Traced_Class_censor(dictFunctor_0_loop *Constructor_Data_Functor_Functor, f_1_loop gopurs_runtime.Value, v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *Constructor_Data_Functor_Functor = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictFunctor_0.V0), gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_3, gopurs_runtime.Apply(f_1, x_4))
})
}), v_2)
}

func Call_Control_Comonad_Traced_Class_track__250573732(dict_0_loop *Constructor_Control_Comonad_Traced_Class_ComonadTraced) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Traced_Class_ComonadTraced = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



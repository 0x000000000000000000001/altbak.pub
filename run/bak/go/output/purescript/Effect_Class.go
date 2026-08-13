package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Class_MonadEffect_dollarDict gopurs_runtime.Value
var once_Effect_Class_MonadEffect_dollarDict sync.Once
func Get_Effect_Class_MonadEffect_dollarDict() gopurs_runtime.Value {
	once_Effect_Class_MonadEffect_dollarDict.Do(func() {
		cache_Effect_Class_MonadEffect_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_MonadEffect_dollarDict(x_0_box)
})
	})
	return cache_Effect_Class_MonadEffect_dollarDict
}

var cache_Effect_Class_monadEffectEffect gopurs_runtime.Value
var once_Effect_Class_monadEffectEffect sync.Once
func Get_Effect_Class_monadEffectEffect() gopurs_runtime.Value {
	once_Effect_Class_monadEffectEffect.Do(func() {
		cache_Effect_Class_monadEffectEffect = gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Effect_monadEffect()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Effect_Class_monadEffectEffect
}

var cache_Effect_Class_liftEffect gopurs_runtime.Value
var once_Effect_Class_liftEffect sync.Once
func Get_Effect_Class_liftEffect() gopurs_runtime.Value {
	once_Effect_Class_liftEffect.Do(func() {
		cache_Effect_Class_liftEffect = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect
}

var cache_Effect_Class_liftEffect__2407462165 gopurs_runtime.Value
var once_Effect_Class_liftEffect__2407462165 sync.Once
func Get_Effect_Class_liftEffect__2407462165() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__2407462165.Do(func() {
		cache_Effect_Class_liftEffect__2407462165 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__2407462165(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect__2407462165
}

var cache_Effect_Class_liftEffect__3456588885 gopurs_runtime.Value
var once_Effect_Class_liftEffect__3456588885 sync.Once
func Get_Effect_Class_liftEffect__3456588885() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__3456588885.Do(func() {
		cache_Effect_Class_liftEffect__3456588885 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__3456588885(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect__3456588885
}

var cache_Effect_Class_liftEffect__1892566677 gopurs_runtime.Value
var once_Effect_Class_liftEffect__1892566677 sync.Once
func Get_Effect_Class_liftEffect__1892566677() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__1892566677.Do(func() {
		cache_Effect_Class_liftEffect__1892566677 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__1892566677(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect__1892566677
}

var cache_Effect_Class_liftEffect__2322711157 gopurs_runtime.Value
var once_Effect_Class_liftEffect__2322711157 sync.Once
func Get_Effect_Class_liftEffect__2322711157() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__2322711157.Do(func() {
		cache_Effect_Class_liftEffect__2322711157 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__2322711157(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect__2322711157
}

var cache_Effect_Class_liftEffect__735761941 gopurs_runtime.Value
var once_Effect_Class_liftEffect__735761941 sync.Once
func Get_Effect_Class_liftEffect__735761941() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__735761941.Do(func() {
		cache_Effect_Class_liftEffect__735761941 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__735761941(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect__735761941
}

var cache_Effect_Class_liftEffect__3357942741 gopurs_runtime.Value
var once_Effect_Class_liftEffect__3357942741 sync.Once
func Get_Effect_Class_liftEffect__3357942741() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__3357942741.Do(func() {
		cache_Effect_Class_liftEffect__3357942741 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__3357942741(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect__3357942741
}

var cache_Effect_Class_liftEffect__226852501 gopurs_runtime.Value
var once_Effect_Class_liftEffect__226852501 sync.Once
func Get_Effect_Class_liftEffect__226852501() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__226852501.Do(func() {
		cache_Effect_Class_liftEffect__226852501 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__226852501(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect__226852501
}

var cache_Effect_Class_liftEffect__2550292213 gopurs_runtime.Value
var once_Effect_Class_liftEffect__2550292213 sync.Once
func Get_Effect_Class_liftEffect__2550292213() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__2550292213.Do(func() {
		cache_Effect_Class_liftEffect__2550292213 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__2550292213(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect](dict_0_box))
})
	})
	return cache_Effect_Class_liftEffect__2550292213
}

var cache_Effect_Class_liftEffect__273534483 gopurs_runtime.Value
var once_Effect_Class_liftEffect__273534483 sync.Once
func Get_Effect_Class_liftEffect__273534483() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__273534483.Do(func() {
		cache_Effect_Class_liftEffect__273534483 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__273534483(__eta0_0_box)
})
	})
	return cache_Effect_Class_liftEffect__273534483
}

var cache_Effect_Class_liftEffect__574228595 gopurs_runtime.Value
var once_Effect_Class_liftEffect__574228595 sync.Once
func Get_Effect_Class_liftEffect__574228595() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__574228595.Do(func() {
		cache_Effect_Class_liftEffect__574228595 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__574228595(__eta0_0_box)
})
	})
	return cache_Effect_Class_liftEffect__574228595
}

var cache_Effect_Class_liftEffect__3226494803 gopurs_runtime.Value
var once_Effect_Class_liftEffect__3226494803 sync.Once
func Get_Effect_Class_liftEffect__3226494803() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__3226494803.Do(func() {
		cache_Effect_Class_liftEffect__3226494803 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__3226494803(__eta0_0_box)
})
	})
	return cache_Effect_Class_liftEffect__3226494803
}

var cache_Effect_Class_liftEffect__1442411827 gopurs_runtime.Value
var once_Effect_Class_liftEffect__1442411827 sync.Once
func Get_Effect_Class_liftEffect__1442411827() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__1442411827.Do(func() {
		cache_Effect_Class_liftEffect__1442411827 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__1442411827(__eta0_0_box)
})
	})
	return cache_Effect_Class_liftEffect__1442411827
}

var cache_Effect_Class_liftEffect__88347923 gopurs_runtime.Value
var once_Effect_Class_liftEffect__88347923 sync.Once
func Get_Effect_Class_liftEffect__88347923() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__88347923.Do(func() {
		cache_Effect_Class_liftEffect__88347923 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__88347923(__eta0_0_box)
})
	})
	return cache_Effect_Class_liftEffect__88347923
}

var cache_Effect_Class_liftEffect__2769380243 gopurs_runtime.Value
var once_Effect_Class_liftEffect__2769380243 sync.Once
func Get_Effect_Class_liftEffect__2769380243() gopurs_runtime.Value {
	once_Effect_Class_liftEffect__2769380243.Do(func() {
		cache_Effect_Class_liftEffect__2769380243 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Effect_Class_liftEffect__2769380243(__eta0_0_box)
})
	})
	return cache_Effect_Class_liftEffect__2769380243
}

type Constructor_Effect_Class_MonadEffect struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2217729261] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Effect_Class_MonadEffect)(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "liftEffect": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Effect_Class_MonadEffect: " + key)
		}
	}
}


func Call_Effect_Class_MonadEffect_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Effect_Class_liftEffect(dict_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__2407462165(dict_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__3456588885(dict_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__1892566677(dict_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__2322711157(dict_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__735761941(dict_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__3357942741(dict_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__226852501(dict_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__2550292213(dict_0_loop *Constructor_Effect_Class_MonadEffect) gopurs_runtime.Value {
var dict_0 *Constructor_Effect_Class_MonadEffect = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Effect_Class_liftEffect__273534483(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), __eta0_0)
}

func Call_Effect_Class_liftEffect__574228595(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), __eta0_0)
}

func Call_Effect_Class_liftEffect__3226494803(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), __eta0_0)
}

func Call_Effect_Class_liftEffect__1442411827(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), __eta0_0)
}

func Call_Effect_Class_liftEffect__88347923(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), __eta0_0)
}

func Call_Effect_Class_liftEffect__2769380243(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), __eta0_0)
}



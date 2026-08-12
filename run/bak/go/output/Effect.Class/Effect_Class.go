package Effect_Class

import (
	pkg_Effect "gopurs/output/Effect"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_monadEffectEffect gopurs_runtime.Value
var once_monadEffectEffect sync.Once
func Get_monadEffectEffect() gopurs_runtime.Value {
	once_monadEffectEffect.Do(func() {
		cache_monadEffectEffect = gopurs_runtime.RecordDict2("Monad0", "liftEffect", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_monadEffectEffect
}

var cache_liftEffect gopurs_runtime.Value
var once_liftEffect sync.Once
func Get_liftEffect() gopurs_runtime.Value {
	once_liftEffect.Do(func() {
		cache_liftEffect = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect
}

var cache_liftEffect__gopurs_runtime_Value_1892566677 gopurs_runtime.Value
var once_liftEffect__gopurs_runtime_Value_1892566677 sync.Once
func Get_liftEffect__gopurs_runtime_Value_1892566677() gopurs_runtime.Value {
	once_liftEffect__gopurs_runtime_Value_1892566677.Do(func() {
		cache_liftEffect__gopurs_runtime_Value_1892566677 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__gopurs_runtime_Value_1892566677(gopurs_runtime.CoerceToStruct[Constructor_MonadEffect[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftEffect__gopurs_runtime_Value_1892566677
}

type Constructor_MonadEffect[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2217729261] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadEffect[gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad0": return c.V0
		case "liftEffect": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadEffect: " + key)
		}
	}
}


func Call_liftEffect(dict_0_loop *Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_liftEffect__gopurs_runtime_Value_1892566677(dict_0_loop *Constructor_MonadEffect[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadEffect[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



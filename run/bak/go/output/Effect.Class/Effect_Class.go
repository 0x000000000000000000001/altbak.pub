package Effect_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect "gopurs/output/Effect"
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
return Call_liftEffect(dict_0_box)
})
	})
	return cache_liftEffect
}

var cache_liftEffect__gopurs_runtime_Value_2747932855 gopurs_runtime.Value
var once_liftEffect__gopurs_runtime_Value_2747932855 sync.Once
func Get_liftEffect__gopurs_runtime_Value_2747932855() gopurs_runtime.Value {
	once_liftEffect__gopurs_runtime_Value_2747932855.Do(func() {
		cache_liftEffect__gopurs_runtime_Value_2747932855 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftEffect__gopurs_runtime_Value_2747932855(dict_0_box)
})
	})
	return cache_liftEffect__gopurs_runtime_Value_2747932855
}

func Call_liftEffect(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "liftEffect")
}

func Call_liftEffect__gopurs_runtime_Value_2747932855(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "liftEffect")
}



package Effect_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Effect "gopurs/output/Effect"
)

var monadEffectEffect gopurs_runtime.Value
var once_monadEffectEffect sync.Once
func Get_monadEffectEffect() gopurs_runtime.Value {
	once_monadEffectEffect.Do(func() {
		monadEffectEffect = gopurs_runtime.RecordDict2("liftEffect", "Monad0", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}))
	})
	return monadEffectEffect
}

var liftEffect gopurs_runtime.Value
var once_liftEffect sync.Once
func Get_liftEffect() gopurs_runtime.Value {
	once_liftEffect.Do(func() {
		liftEffect = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "liftEffect")
}()
})
	})
	return liftEffect
}





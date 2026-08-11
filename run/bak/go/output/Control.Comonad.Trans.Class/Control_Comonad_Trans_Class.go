package Control_Comonad_Trans_Class

import (
	pkg_Control_Monad_Identity_Trans "gopurs/output/Control.Monad.Identity.Trans"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_lower gopurs_runtime.Value
var once_lower sync.Once
func Get_lower() gopurs_runtime.Value {
	once_lower.Do(func() {
		cache_lower = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lower(dict_0_box)
})
	})
	return cache_lower
}

var cache_comonadTransIdentityT gopurs_runtime.Value
var once_comonadTransIdentityT sync.Once
func Get_comonadTransIdentityT() gopurs_runtime.Value {
	once_comonadTransIdentityT.Do(func() {
		cache_comonadTransIdentityT = gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_Identity_Trans.Get_runIdentityT()
}))
	})
	return cache_comonadTransIdentityT
}

func Call_lower(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "lower")
}



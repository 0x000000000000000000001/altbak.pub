package Control_Comonad_Trans_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_Identity_Trans "gopurs/output/Control.Monad.Identity.Trans"
)

var lower gopurs_runtime.Value
var once_lower sync.Once
func Get_lower() gopurs_runtime.Value {
	once_lower.Do(func() {
		lower = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "lower")
})
	})
	return lower
}

var comonadTransIdentityT gopurs_runtime.Value
var once_comonadTransIdentityT sync.Once
func Get_comonadTransIdentityT() gopurs_runtime.Value {
	once_comonadTransIdentityT.Do(func() {
		comonadTransIdentityT = gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_Identity_Trans.Get_runIdentityT()
}))
	})
	return comonadTransIdentityT
}





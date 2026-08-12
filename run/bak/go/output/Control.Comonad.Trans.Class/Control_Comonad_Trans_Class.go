package Control_Comonad_Trans_Class

import (
	pkg_Control_Monad_Identity_Trans "gopurs/output/Control.Monad.Identity.Trans"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_lower gopurs_runtime.Value
var once_lower sync.Once
func Get_lower() gopurs_runtime.Value {
	once_lower.Do(func() {
		cache_lower = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lower(gopurs_runtime.CoerceToStruct[Constructor_ComonadTrans[gopurs_runtime.Value]](dict_0_box))
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

type Constructor_ComonadTrans[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3399197123] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_ComonadTrans[gopurs_runtime.Value])(ptr)
		switch key {
		case "lower": return c.V0
		default: panic("Key not found in dictionary Constructor_ComonadTrans: " + key)
		}
	}
}


func Call_lower(dict_0_loop *Constructor_ComonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_ComonadTrans[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}



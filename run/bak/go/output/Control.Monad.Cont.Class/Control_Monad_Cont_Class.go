package Control_Monad_Cont_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_callCC gopurs_runtime.Value
var once_callCC sync.Once
func Get_callCC() gopurs_runtime.Value {
	once_callCC.Do(func() {
		cache_callCC = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_callCC(gopurs_runtime.CoerceToStruct[Constructor_MonadCont[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_callCC
}

type Constructor_MonadCont[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1800060259] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadCont[gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad0": return c.V0
		case "callCC": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadCont: " + key)
		}
	}
}


func Call_callCC(dict_0_loop *Constructor_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadCont[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



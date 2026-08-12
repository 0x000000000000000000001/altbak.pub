package Data_Profunctor_Costrong

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_unsecond gopurs_runtime.Value
var once_unsecond sync.Once
func Get_unsecond() gopurs_runtime.Value {
	once_unsecond.Do(func() {
		cache_unsecond = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsecond(gopurs_runtime.CoerceToStruct[Constructor_Costrong[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unsecond
}

var cache_unfirst gopurs_runtime.Value
var once_unfirst sync.Once
func Get_unfirst() gopurs_runtime.Value {
	once_unfirst.Do(func() {
		cache_unfirst = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfirst(gopurs_runtime.CoerceToStruct[Constructor_Costrong[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfirst
}

type Constructor_Costrong[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2173123103] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Costrong[gopurs_runtime.Value])(ptr)
		switch key {
		case "Profunctor0": return c.V0
		case "unfirst": return c.V1
		case "unsecond": return c.V2
		default: panic("Key not found in dictionary Constructor_Costrong: " + key)
		}
	}
}


func Call_unsecond(dict_0_loop *Constructor_Costrong[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Costrong[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_unfirst(dict_0_loop *Constructor_Costrong[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Costrong[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



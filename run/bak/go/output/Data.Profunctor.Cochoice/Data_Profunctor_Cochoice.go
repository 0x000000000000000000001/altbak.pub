package Data_Profunctor_Cochoice

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_unright gopurs_runtime.Value
var once_unright sync.Once
func Get_unright() gopurs_runtime.Value {
	once_unright.Do(func() {
		cache_unright = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unright(gopurs_runtime.CoerceToStruct[Constructor_Cochoice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unright
}

var cache_unleft gopurs_runtime.Value
var once_unleft sync.Once
func Get_unleft() gopurs_runtime.Value {
	once_unleft.Do(func() {
		cache_unleft = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unleft(gopurs_runtime.CoerceToStruct[Constructor_Cochoice[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unleft
}

type Constructor_Cochoice[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1827340575] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Cochoice[gopurs_runtime.Value])(ptr)
		switch key {
		case "Profunctor0": return c.V0
		case "unleft": return c.V1
		case "unright": return c.V2
		default: panic("Key not found in dictionary Constructor_Cochoice: " + key)
		}
	}
}


func Call_unright(dict_0_loop *Constructor_Cochoice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Cochoice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_unleft(dict_0_loop *Constructor_Cochoice[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Cochoice[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



package Control_Comonad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_extract gopurs_runtime.Value
var once_extract sync.Once
func Get_extract() gopurs_runtime.Value {
	once_extract.Do(func() {
		cache_extract = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extract(gopurs_runtime.CoerceToStruct[Constructor_Comonad[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extract
}

var cache_extract__gopurs_runtime_Value_1031647521 gopurs_runtime.Value
var once_extract__gopurs_runtime_Value_1031647521 sync.Once
func Get_extract__gopurs_runtime_Value_1031647521() gopurs_runtime.Value {
	once_extract__gopurs_runtime_Value_1031647521.Do(func() {
		cache_extract__gopurs_runtime_Value_1031647521 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extract__gopurs_runtime_Value_1031647521(gopurs_runtime.CoerceToStruct[Constructor_Comonad[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extract__gopurs_runtime_Value_1031647521
}

type Constructor_Comonad[T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2886863693] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Comonad[gopurs_runtime.Value])(ptr)
		switch key {
		case "Extend0": return c.V0
		case "extract": return c.V1
		default: panic("Key not found in dictionary Constructor_Comonad: " + key)
		}
	}
}


func Call_extract(dict_0_loop *Constructor_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_extract__gopurs_runtime_Value_1031647521(dict_0_loop *Constructor_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



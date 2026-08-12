package Data_Symbol

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_reifySymbol gopurs_runtime.Value
var once_reifySymbol sync.Once
func Get_reifySymbol() gopurs_runtime.Value {
	once_reifySymbol.Do(func() {
		cache_reifySymbol = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reifySymbol(s_0_box.StrVal(), f_1_box)
})
	})
	return cache_reifySymbol
}

var cache_reflectSymbol gopurs_runtime.Value
var once_reflectSymbol sync.Once
func Get_reflectSymbol() gopurs_runtime.Value {
	once_reflectSymbol.Do(func() {
		cache_reflectSymbol = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol(gopurs_runtime.CoerceToStruct[Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol
}

type Constructor_IsSymbol[T_sym any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2134024384] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_IsSymbol[gopurs_runtime.Value])(ptr)
		switch key {
		case "reflectSymbol": return c.V0
		default: panic("Key not found in dictionary Constructor_IsSymbol: " + key)
		}
	}
}


func Call_reifySymbol(s_0_loop string, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply3(Get_unsafeCoerce(), gopurs_runtime.Func(func(dictIsSymbol_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, dictIsSymbol_2)
}), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(s_0)
})), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
}

func Call_reflectSymbol(dict_0_loop *Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_unsafeCoerce() gopurs_runtime.Value {
	return _Gopurs_UnsafeCoerce
}

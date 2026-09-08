package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Comonad_dollar_Dict gopurs_runtime.Value
var once_Control_Comonad_Comonad_dollar_Dict sync.Once
func Get_Control_Comonad_Comonad_dollar_Dict() gopurs_runtime.Value {
	once_Control_Comonad_Comonad_dollar_Dict.Do(func() {
		cache_Control_Comonad_Comonad_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Comonad_dollar_Dict(func() struct{
	Extend0 gopurs_runtime.Value
	extract gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Extend0 gopurs_runtime.Value
	extract gopurs_runtime.Value
}{}
					clone.Extend0 = gopurs_runtime.RecordGet(orig, "Extend0")
					clone.extract = gopurs_runtime.RecordGet(orig, "extract")
					return clone
				}()))}
})
	})
	return cache_Control_Comonad_Comonad_dollar_Dict
}

var cache_Control_Comonad_extract gopurs_runtime.Value
var once_Control_Comonad_extract sync.Once
func Get_Control_Comonad_extract() gopurs_runtime.Value {
	once_Control_Comonad_extract.Do(func() {
		cache_Control_Comonad_extract = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_extract(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_extract
}

type Constructor_Control_Comonad_Comonad[T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2886863693] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Comonad[any])(ptr)
		_ = c
		switch key {
		case "Extend0": return gopurs_runtime.Box(c.V0)
		case "extract": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Comonad: " + key)
		}
	}
}


func Call_Control_Comonad_Comonad_dollar_Dict(x_0_loop struct{
	Extend0 gopurs_runtime.Value
	extract gopurs_runtime.Value
}) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
var x_0 struct{
	Extend0 gopurs_runtime.Value
	extract gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Extend0", "extract"}, []gopurs_runtime.Value{orig.Extend0, orig.extract})
				}())
}

func Call_Control_Comonad_extract(dict_0_loop *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



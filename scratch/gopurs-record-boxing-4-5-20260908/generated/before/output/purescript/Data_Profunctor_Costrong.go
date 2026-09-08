package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Costrong_Costrong_dollar_Dict gopurs_runtime.Value
var once_Data_Profunctor_Costrong_Costrong_dollar_Dict sync.Once
func Get_Data_Profunctor_Costrong_Costrong_dollar_Dict() gopurs_runtime.Value {
	once_Data_Profunctor_Costrong_Costrong_dollar_Dict.Do(func() {
		cache_Data_Profunctor_Costrong_Costrong_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2173123103, UnsafePtr: unsafe.Pointer(Call_Data_Profunctor_Costrong_Costrong_dollar_Dict(func() struct{
	Profunctor0 gopurs_runtime.Value
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Profunctor0 gopurs_runtime.Value
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
}{}
					clone.Profunctor0 = gopurs_runtime.RecordGet(orig, "Profunctor0")
					clone.unfirst = gopurs_runtime.RecordGet(orig, "unfirst")
					clone.unsecond = gopurs_runtime.RecordGet(orig, "unsecond")
					return clone
				}()))}
})
	})
	return cache_Data_Profunctor_Costrong_Costrong_dollar_Dict
}

var cache_Data_Profunctor_Costrong_unsecond gopurs_runtime.Value
var once_Data_Profunctor_Costrong_unsecond sync.Once
func Get_Data_Profunctor_Costrong_unsecond() gopurs_runtime.Value {
	once_Data_Profunctor_Costrong_unsecond.Do(func() {
		cache_Data_Profunctor_Costrong_unsecond = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Costrong_unsecond(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Costrong_Costrong[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_Costrong_unsecond
}

var cache_Data_Profunctor_Costrong_unfirst gopurs_runtime.Value
var once_Data_Profunctor_Costrong_unfirst sync.Once
func Get_Data_Profunctor_Costrong_unfirst() gopurs_runtime.Value {
	once_Data_Profunctor_Costrong_unfirst.Do(func() {
		cache_Data_Profunctor_Costrong_unfirst = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Costrong_unfirst(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Costrong_Costrong[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Profunctor_Costrong_unfirst
}

type Constructor_Data_Profunctor_Costrong_Costrong[T_p any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2173123103] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Costrong_Costrong[any])(ptr)
		_ = c
		switch key {
		case "Profunctor0": return gopurs_runtime.Box(c.V0)
		case "unfirst": return gopurs_runtime.Box(c.V1)
		case "unsecond": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Profunctor_Costrong_Costrong: " + key)
		}
	}
}


func Call_Data_Profunctor_Costrong_Costrong_dollar_Dict(x_0_loop struct{
	Profunctor0 gopurs_runtime.Value
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
}) *Constructor_Data_Profunctor_Costrong_Costrong[gopurs_runtime.Value] {
var x_0 struct{
	Profunctor0 gopurs_runtime.Value
	unfirst gopurs_runtime.Value
	unsecond gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Costrong_Costrong[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Profunctor0", "unfirst", "unsecond"}, []gopurs_runtime.Value{orig.Profunctor0, orig.unfirst, orig.unsecond})
				}())
}

func Call_Data_Profunctor_Costrong_unsecond(dict_0_loop *Constructor_Data_Profunctor_Costrong_Costrong[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Costrong_Costrong[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Profunctor_Costrong_unfirst(dict_0_loop *Constructor_Data_Profunctor_Costrong_Costrong[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Costrong_Costrong[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



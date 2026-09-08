package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Biapplicative_Biapplicative_dollar_Dict gopurs_runtime.Value
var once_Control_Biapplicative_Biapplicative_dollar_Dict sync.Once
func Get_Control_Biapplicative_Biapplicative_dollar_Dict() gopurs_runtime.Value {
	once_Control_Biapplicative_Biapplicative_dollar_Dict.Do(func() {
		cache_Control_Biapplicative_Biapplicative_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3949191309, UnsafePtr: unsafe.Pointer(Call_Control_Biapplicative_Biapplicative_dollar_Dict(func() struct{
	Biapply0 gopurs_runtime.Value
	bipure gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Biapply0 gopurs_runtime.Value
	bipure gopurs_runtime.Value
}{}
					clone.Biapply0 = gopurs_runtime.RecordGet(orig, "Biapply0")
					clone.bipure = gopurs_runtime.RecordGet(orig, "bipure")
					return clone
				}()))}
})
	})
	return cache_Control_Biapplicative_Biapplicative_dollar_Dict
}

var cache_Control_Biapplicative_bipure gopurs_runtime.Value
var once_Control_Biapplicative_bipure sync.Once
func Get_Control_Biapplicative_bipure() gopurs_runtime.Value {
	once_Control_Biapplicative_bipure.Do(func() {
		cache_Control_Biapplicative_bipure = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapplicative_bipure(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Biapplicative_bipure
}

var cache_Control_Biapplicative_biapplicativeTuple gopurs_runtime.Value
var once_Control_Biapplicative_biapplicativeTuple sync.Once
func Get_Control_Biapplicative_biapplicativeTuple() gopurs_runtime.Value {
	once_Control_Biapplicative_biapplicativeTuple.Do(func() {
		cache_Control_Biapplicative_biapplicativeTuple = gopurs_runtime.Value{Type: 9, IntVal: 3949191309, UnsafePtr: unsafe.Pointer(Rebox_Control_Biapplicative_1794813908_3565318777((&Constructor_Control_Biapplicative_Biapplicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3774602829, UnsafePtr: unsafe.Pointer(Rebox_Control_Biapplicative_4265826836_2448899513(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapply_Biapply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](Get_Control_Biapply_biapplyTuple())))}
}), Get_Data_Tuple_Tuple()})))}
	})
	return cache_Control_Biapplicative_biapplicativeTuple
}

type Constructor_Control_Biapplicative_Biapplicative[T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3949191309] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Biapplicative_Biapplicative[any])(ptr)
		_ = c
		switch key {
		case "Biapply0": return gopurs_runtime.Box(c.V0)
		case "bipure": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Biapplicative_Biapplicative: " + key)
		}
	}
}


func Call_Control_Biapplicative_Biapplicative_dollar_Dict(x_0_loop struct{
	Biapply0 gopurs_runtime.Value
	bipure gopurs_runtime.Value
}) *Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value] {
var x_0 struct{
	Biapply0 gopurs_runtime.Value
	bipure gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Biapply0", "bipure", orig.Biapply0, orig.bipure)
				}())
}

func Call_Control_Biapplicative_bipure(dict_0_loop *Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Rebox_Control_Biapplicative_1794813908_3565318777(in *Constructor_Control_Biapplicative_Biapplicative[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Biapplicative_Biapplicative[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Biapplicative_4265826836_2448899513(in *Constructor_Control_Biapply_Biapply[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Biapply_Biapply[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Biapply_Biapply[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}



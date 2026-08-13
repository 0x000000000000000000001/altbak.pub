package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Symbol_IsSymbol_dollarDict gopurs_runtime.Value
var once_Data_Symbol_IsSymbol_dollarDict sync.Once
func Get_Data_Symbol_IsSymbol_dollarDict() gopurs_runtime.Value {
	once_Data_Symbol_IsSymbol_dollarDict.Do(func() {
		cache_Data_Symbol_IsSymbol_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Symbol_IsSymbol_dollarDict(x_0_box)
})
	})
	return cache_Data_Symbol_IsSymbol_dollarDict
}

var cache_Data_Symbol_reifySymbol gopurs_runtime.Value
var once_Data_Symbol_reifySymbol sync.Once
func Get_Data_Symbol_reifySymbol() gopurs_runtime.Value {
	once_Data_Symbol_reifySymbol.Do(func() {
		cache_Data_Symbol_reifySymbol = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Symbol_reifySymbol(s_0_box.StrVal(), f_1_box)
})
	})
	return cache_Data_Symbol_reifySymbol
}

var cache_Data_Symbol_reflectSymbol gopurs_runtime.Value
var once_Data_Symbol_reflectSymbol sync.Once
func Get_Data_Symbol_reflectSymbol() gopurs_runtime.Value {
	once_Data_Symbol_reflectSymbol.Do(func() {
		cache_Data_Symbol_reflectSymbol = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Symbol_reflectSymbol(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Symbol_reflectSymbol
}

var cache_Data_Symbol_reflectSymbol__3416619207 gopurs_runtime.Value
var once_Data_Symbol_reflectSymbol__3416619207 sync.Once
func Get_Data_Symbol_reflectSymbol__3416619207() gopurs_runtime.Value {
	once_Data_Symbol_reflectSymbol__3416619207.Do(func() {
		cache_Data_Symbol_reflectSymbol__3416619207 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Symbol_reflectSymbol__3416619207(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Symbol_reflectSymbol__3416619207
}

var cache_Data_Symbol_reflectSymbol__1166932993 gopurs_runtime.Value
var once_Data_Symbol_reflectSymbol__1166932993 sync.Once
func Get_Data_Symbol_reflectSymbol__1166932993() gopurs_runtime.Value {
	once_Data_Symbol_reflectSymbol__1166932993.Do(func() {
		cache_Data_Symbol_reflectSymbol__1166932993 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Symbol_reflectSymbol__1166932993(gopurs_runtime.CoerceToStruct[Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_Symbol_reflectSymbol__1166932993
}

type Constructor_Data_Symbol_IsSymbol[T_sym any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2134024384] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "reflectSymbol": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_Symbol_IsSymbol: " + key)
		}
	}
}


func Call_Data_Symbol_IsSymbol_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Symbol_reifySymbol(s_0_loop string, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply3(Get_Data_Symbol_unsafeCoerce(), gopurs_runtime.Func(func(dictIsSymbol_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, dictIsSymbol_2)
}), gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(s_0)
})), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil})
}

func Call_Data_Symbol_reflectSymbol(dict_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Symbol_reflectSymbol__3416619207(dict_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_Symbol_reflectSymbol__1166932993(dict_0_loop *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Symbol_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Get_Data_Symbol_unsafeCoerce() gopurs_runtime.Value {
	return _Gopurs_Data_Symbol_UnsafeCoerce
}

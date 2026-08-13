package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Comonad_dollarDict gopurs_runtime.Value
var once_Control_Comonad_Comonad_dollarDict sync.Once
func Get_Control_Comonad_Comonad_dollarDict() gopurs_runtime.Value {
	once_Control_Comonad_Comonad_dollarDict.Do(func() {
		cache_Control_Comonad_Comonad_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Comonad_dollarDict(x_0_box)
})
	})
	return cache_Control_Comonad_Comonad_dollarDict
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

var cache_Control_Comonad_extract__1031647521 gopurs_runtime.Value
var once_Control_Comonad_extract__1031647521 sync.Once
func Get_Control_Comonad_extract__1031647521() gopurs_runtime.Value {
	once_Control_Comonad_extract__1031647521.Do(func() {
		cache_Control_Comonad_extract__1031647521 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_extract__1031647521(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_extract__1031647521
}

var cache_Control_Comonad_extract__3319904577 gopurs_runtime.Value
var once_Control_Comonad_extract__3319904577 sync.Once
func Get_Control_Comonad_extract__3319904577() gopurs_runtime.Value {
	once_Control_Comonad_extract__3319904577.Do(func() {
		cache_Control_Comonad_extract__3319904577 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_extract__3319904577(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_extract__3319904577
}

type Constructor_Control_Comonad_Comonad[T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2886863693] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Comonad[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Extend0": return gopurs_runtime.Box(c.V0)
		case "extract": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Comonad: " + key)
		}
	}
}


func Call_Control_Comonad_Comonad_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_extract(dict_0_loop *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_extract__1031647521(dict_0_loop *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_extract__3319904577(dict_0_loop *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Biapplicative_Biapplicative_dollarDict gopurs_runtime.Value
var once_Control_Biapplicative_Biapplicative_dollarDict sync.Once
func Get_Control_Biapplicative_Biapplicative_dollarDict() gopurs_runtime.Value {
	once_Control_Biapplicative_Biapplicative_dollarDict.Do(func() {
		cache_Control_Biapplicative_Biapplicative_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapplicative_Biapplicative_dollarDict(x_0_box)
})
	})
	return cache_Control_Biapplicative_Biapplicative_dollarDict
}

var cache_Control_Biapplicative_bipure gopurs_runtime.Value
var once_Control_Biapplicative_bipure sync.Once
func Get_Control_Biapplicative_bipure() gopurs_runtime.Value {
	once_Control_Biapplicative_bipure.Do(func() {
		cache_Control_Biapplicative_bipure = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapplicative_bipure(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapplicative_Biapplicative](dict_0_box))
})
	})
	return cache_Control_Biapplicative_bipure
}

var cache_Control_Biapplicative_biapplicativeTuple gopurs_runtime.Value
var once_Control_Biapplicative_biapplicativeTuple sync.Once
func Get_Control_Biapplicative_biapplicativeTuple() gopurs_runtime.Value {
	once_Control_Biapplicative_biapplicativeTuple.Do(func() {
		cache_Control_Biapplicative_biapplicativeTuple = gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Biapply_biapplyTuple()
}), Get_Data_Tuple_Tuple())
	})
	return cache_Control_Biapplicative_biapplicativeTuple
}

var cache_Control_Biapplicative_bipure__1449949980 gopurs_runtime.Value
var once_Control_Biapplicative_bipure__1449949980 sync.Once
func Get_Control_Biapplicative_bipure__1449949980() gopurs_runtime.Value {
	once_Control_Biapplicative_bipure__1449949980.Do(func() {
		cache_Control_Biapplicative_bipure__1449949980 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Biapplicative_bipure__1449949980(gopurs_runtime.CoerceToStruct[Constructor_Control_Biapplicative_Biapplicative](dict_0_box))
})
	})
	return cache_Control_Biapplicative_bipure__1449949980
}

type Constructor_Control_Biapplicative_Biapplicative struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3949191309] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Biapplicative_Biapplicative)(ptr)
		_ = c
		switch key {
		case "Biapply0": return gopurs_runtime.Box(c.V0)
		case "bipure": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Biapplicative_Biapplicative: " + key)
		}
	}
}


func Call_Control_Biapplicative_Biapplicative_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Biapplicative_bipure(dict_0_loop *Constructor_Control_Biapplicative_Biapplicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Biapplicative_Biapplicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Biapplicative_bipure__1449949980(dict_0_loop *Constructor_Control_Biapplicative_Biapplicative) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Biapplicative_Biapplicative = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



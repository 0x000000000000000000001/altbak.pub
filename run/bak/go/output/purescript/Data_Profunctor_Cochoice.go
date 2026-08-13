package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Profunctor_Cochoice_Cochoice_dollarDict gopurs_runtime.Value
var once_Data_Profunctor_Cochoice_Cochoice_dollarDict sync.Once
func Get_Data_Profunctor_Cochoice_Cochoice_dollarDict() gopurs_runtime.Value {
	once_Data_Profunctor_Cochoice_Cochoice_dollarDict.Do(func() {
		cache_Data_Profunctor_Cochoice_Cochoice_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Cochoice_Cochoice_dollarDict(x_0_box)
})
	})
	return cache_Data_Profunctor_Cochoice_Cochoice_dollarDict
}

var cache_Data_Profunctor_Cochoice_unright gopurs_runtime.Value
var once_Data_Profunctor_Cochoice_unright sync.Once
func Get_Data_Profunctor_Cochoice_unright() gopurs_runtime.Value {
	once_Data_Profunctor_Cochoice_unright.Do(func() {
		cache_Data_Profunctor_Cochoice_unright = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Cochoice_unright(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Cochoice_Cochoice](dict_0_box))
})
	})
	return cache_Data_Profunctor_Cochoice_unright
}

var cache_Data_Profunctor_Cochoice_unleft gopurs_runtime.Value
var once_Data_Profunctor_Cochoice_unleft sync.Once
func Get_Data_Profunctor_Cochoice_unleft() gopurs_runtime.Value {
	once_Data_Profunctor_Cochoice_unleft.Do(func() {
		cache_Data_Profunctor_Cochoice_unleft = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Profunctor_Cochoice_unleft(gopurs_runtime.CoerceToStruct[Constructor_Data_Profunctor_Cochoice_Cochoice](dict_0_box))
})
	})
	return cache_Data_Profunctor_Cochoice_unleft
}

type Constructor_Data_Profunctor_Cochoice_Cochoice struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1827340575] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Profunctor_Cochoice_Cochoice)(ptr)
		_ = c
		switch key {
		case "Profunctor0": return gopurs_runtime.Box(c.V0)
		case "unleft": return gopurs_runtime.Box(c.V1)
		case "unright": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Data_Profunctor_Cochoice_Cochoice: " + key)
		}
	}
}


func Call_Data_Profunctor_Cochoice_Cochoice_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Profunctor_Cochoice_unright(dict_0_loop *Constructor_Data_Profunctor_Cochoice_Cochoice) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Cochoice_Cochoice = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Data_Profunctor_Cochoice_unleft(dict_0_loop *Constructor_Data_Profunctor_Cochoice_Cochoice) gopurs_runtime.Value {
var dict_0 *Constructor_Data_Profunctor_Cochoice_Cochoice = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



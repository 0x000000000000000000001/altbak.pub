package Control_Monad_Gen_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_sized gopurs_runtime.Value
var once_sized sync.Once
func Get_sized() gopurs_runtime.Value {
	once_sized.Do(func() {
		cache_sized = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sized(gopurs_runtime.CoerceToStruct[Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sized
}

var cache_sized__gopurs_runtime_Value_3147117623 gopurs_runtime.Value
var once_sized__gopurs_runtime_Value_3147117623 sync.Once
func Get_sized__gopurs_runtime_Value_3147117623() gopurs_runtime.Value {
	once_sized__gopurs_runtime_Value_3147117623.Do(func() {
		cache_sized__gopurs_runtime_Value_3147117623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sized__gopurs_runtime_Value_3147117623(gopurs_runtime.CoerceToStruct[Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sized__gopurs_runtime_Value_3147117623
}

var cache_resize gopurs_runtime.Value
var once_resize sync.Once
func Get_resize() gopurs_runtime.Value {
	once_resize.Do(func() {
		cache_resize = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_resize(gopurs_runtime.CoerceToStruct[Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_resize
}

var cache_resize__gopurs_runtime_Value_1050945947 gopurs_runtime.Value
var once_resize__gopurs_runtime_Value_1050945947 sync.Once
func Get_resize__gopurs_runtime_Value_1050945947() gopurs_runtime.Value {
	once_resize__gopurs_runtime_Value_1050945947.Do(func() {
		cache_resize__gopurs_runtime_Value_1050945947 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_resize__gopurs_runtime_Value_1050945947(gopurs_runtime.CoerceToStruct[Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_resize__gopurs_runtime_Value_1050945947
}

var cache_chooseInt gopurs_runtime.Value
var once_chooseInt sync.Once
func Get_chooseInt() gopurs_runtime.Value {
	once_chooseInt.Do(func() {
		cache_chooseInt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseInt(gopurs_runtime.CoerceToStruct[Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_chooseInt
}

var cache_chooseInt__gopurs_runtime_Value_1063828903 gopurs_runtime.Value
var once_chooseInt__gopurs_runtime_Value_1063828903 sync.Once
func Get_chooseInt__gopurs_runtime_Value_1063828903() gopurs_runtime.Value {
	once_chooseInt__gopurs_runtime_Value_1063828903.Do(func() {
		cache_chooseInt__gopurs_runtime_Value_1063828903 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseInt__gopurs_runtime_Value_1063828903(gopurs_runtime.CoerceToStruct[Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_chooseInt__gopurs_runtime_Value_1063828903
}

var cache_chooseFloat gopurs_runtime.Value
var once_chooseFloat sync.Once
func Get_chooseFloat() gopurs_runtime.Value {
	once_chooseFloat.Do(func() {
		cache_chooseFloat = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseFloat(gopurs_runtime.CoerceToStruct[Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_chooseFloat
}

var cache_chooseFloat__gopurs_runtime_Value_1964853975 gopurs_runtime.Value
var once_chooseFloat__gopurs_runtime_Value_1964853975 sync.Once
func Get_chooseFloat__gopurs_runtime_Value_1964853975() gopurs_runtime.Value {
	once_chooseFloat__gopurs_runtime_Value_1964853975.Do(func() {
		cache_chooseFloat__gopurs_runtime_Value_1964853975 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseFloat__gopurs_runtime_Value_1964853975(gopurs_runtime.CoerceToStruct[Constructor_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_chooseFloat__gopurs_runtime_Value_1964853975
}

var cache_chooseBool gopurs_runtime.Value
var once_chooseBool sync.Once
func Get_chooseBool() gopurs_runtime.Value {
	once_chooseBool.Do(func() {
		cache_chooseBool = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chooseBool(dict_0_box)
})
	})
	return cache_chooseBool
}

type Constructor_MonadGen[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
	V4 gopurs_runtime.Value
	V5 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2254593219] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadGen[gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad0": return c.V0
		case "chooseBool": return c.V1
		case "chooseFloat": return c.V2
		case "chooseInt": return c.V3
		case "resize": return c.V4
		case "sized": return c.V5
		default: panic("Key not found in dictionary Constructor_MonadGen: " + key)
		}
	}
}


func Call_sized(dict_0_loop *Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
}

func Call_sized__gopurs_runtime_Value_3147117623(dict_0_loop *Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V5
}

func Call_resize(dict_0_loop *Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_resize__gopurs_runtime_Value_1050945947(dict_0_loop *Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_chooseInt(dict_0_loop *Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_chooseInt__gopurs_runtime_Value_1063828903(dict_0_loop *Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_chooseFloat(dict_0_loop *Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_chooseFloat__gopurs_runtime_Value_1964853975(dict_0_loop *Constructor_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_chooseBool(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "chooseBool")
}



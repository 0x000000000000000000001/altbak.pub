package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Gen_Class_MonadGen_dollarDict gopurs_runtime.Value
var once_Control_Monad_Gen_Class_MonadGen_dollarDict sync.Once
func Get_Control_Monad_Gen_Class_MonadGen_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_MonadGen_dollarDict.Do(func() {
		cache_Control_Monad_Gen_Class_MonadGen_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_MonadGen_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Gen_Class_MonadGen_dollarDict
}

var cache_Control_Monad_Gen_Class_sized gopurs_runtime.Value
var once_Control_Monad_Gen_Class_sized sync.Once
func Get_Control_Monad_Gen_Class_sized() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_sized.Do(func() {
		cache_Control_Monad_Gen_Class_sized = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_sized(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_sized
}

var cache_Control_Monad_Gen_Class_resize gopurs_runtime.Value
var once_Control_Monad_Gen_Class_resize sync.Once
func Get_Control_Monad_Gen_Class_resize() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_resize.Do(func() {
		cache_Control_Monad_Gen_Class_resize = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_resize(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_resize
}

var cache_Control_Monad_Gen_Class_chooseInt gopurs_runtime.Value
var once_Control_Monad_Gen_Class_chooseInt sync.Once
func Get_Control_Monad_Gen_Class_chooseInt() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_chooseInt.Do(func() {
		cache_Control_Monad_Gen_Class_chooseInt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_chooseInt(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_chooseInt
}

var cache_Control_Monad_Gen_Class_chooseFloat gopurs_runtime.Value
var once_Control_Monad_Gen_Class_chooseFloat sync.Once
func Get_Control_Monad_Gen_Class_chooseFloat() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_chooseFloat.Do(func() {
		cache_Control_Monad_Gen_Class_chooseFloat = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_chooseFloat(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_chooseFloat
}

var cache_Control_Monad_Gen_Class_chooseBool gopurs_runtime.Value
var once_Control_Monad_Gen_Class_chooseBool sync.Once
func Get_Control_Monad_Gen_Class_chooseBool() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_chooseBool.Do(func() {
		cache_Control_Monad_Gen_Class_chooseBool = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_chooseBool(dict_0_box)
})
	})
	return cache_Control_Monad_Gen_Class_chooseBool
}

var cache_Control_Monad_Gen_Class_chooseFloat__1964853975 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_chooseFloat__1964853975 sync.Once
func Get_Control_Monad_Gen_Class_chooseFloat__1964853975() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_chooseFloat__1964853975.Do(func() {
		cache_Control_Monad_Gen_Class_chooseFloat__1964853975 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_chooseFloat__1964853975(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_chooseFloat__1964853975
}

var cache_Control_Monad_Gen_Class_chooseInt__1063828903 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_chooseInt__1063828903 sync.Once
func Get_Control_Monad_Gen_Class_chooseInt__1063828903() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_chooseInt__1063828903.Do(func() {
		cache_Control_Monad_Gen_Class_chooseInt__1063828903 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_chooseInt__1063828903(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_chooseInt__1063828903
}

var cache_Control_Monad_Gen_Class_resize__4113973243 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_resize__4113973243 sync.Once
func Get_Control_Monad_Gen_Class_resize__4113973243() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_resize__4113973243.Do(func() {
		cache_Control_Monad_Gen_Class_resize__4113973243 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_resize__4113973243(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_resize__4113973243
}

var cache_Control_Monad_Gen_Class_resize__1050945947 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_resize__1050945947 sync.Once
func Get_Control_Monad_Gen_Class_resize__1050945947() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_resize__1050945947.Do(func() {
		cache_Control_Monad_Gen_Class_resize__1050945947 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_resize__1050945947(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_resize__1050945947
}

var cache_Control_Monad_Gen_Class_resize__1313223195 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_resize__1313223195 sync.Once
func Get_Control_Monad_Gen_Class_resize__1313223195() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_resize__1313223195.Do(func() {
		cache_Control_Monad_Gen_Class_resize__1313223195 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_resize__1313223195(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_resize__1313223195
}

var cache_Control_Monad_Gen_Class_resize__2904398683 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_resize__2904398683 sync.Once
func Get_Control_Monad_Gen_Class_resize__2904398683() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_resize__2904398683.Do(func() {
		cache_Control_Monad_Gen_Class_resize__2904398683 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_resize__2904398683(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_resize__2904398683
}

var cache_Control_Monad_Gen_Class_sized__120035991 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_sized__120035991 sync.Once
func Get_Control_Monad_Gen_Class_sized__120035991() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_sized__120035991.Do(func() {
		cache_Control_Monad_Gen_Class_sized__120035991 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_sized__120035991(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_sized__120035991
}

var cache_Control_Monad_Gen_Class_sized__2830838711 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_sized__2830838711 sync.Once
func Get_Control_Monad_Gen_Class_sized__2830838711() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_sized__2830838711.Do(func() {
		cache_Control_Monad_Gen_Class_sized__2830838711 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_sized__2830838711(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_sized__2830838711
}

var cache_Control_Monad_Gen_Class_sized__3147117623 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_sized__3147117623 sync.Once
func Get_Control_Monad_Gen_Class_sized__3147117623() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_sized__3147117623.Do(func() {
		cache_Control_Monad_Gen_Class_sized__3147117623 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_sized__3147117623(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_sized__3147117623
}

var cache_Control_Monad_Gen_Class_sized__4206899191 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_sized__4206899191 sync.Once
func Get_Control_Monad_Gen_Class_sized__4206899191() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_sized__4206899191.Do(func() {
		cache_Control_Monad_Gen_Class_sized__4206899191 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_sized__4206899191(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_sized__4206899191
}

var cache_Control_Monad_Gen_Class_sized__2391211191 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_sized__2391211191 sync.Once
func Get_Control_Monad_Gen_Class_sized__2391211191() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_sized__2391211191.Do(func() {
		cache_Control_Monad_Gen_Class_sized__2391211191 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_sized__2391211191(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_sized__2391211191
}

var cache_Control_Monad_Gen_Class_sized__2241633463 gopurs_runtime.Value
var once_Control_Monad_Gen_Class_sized__2241633463 sync.Once
func Get_Control_Monad_Gen_Class_sized__2241633463() gopurs_runtime.Value {
	once_Control_Monad_Gen_Class_sized__2241633463.Do(func() {
		cache_Control_Monad_Gen_Class_sized__2241633463 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Gen_Class_sized__2241633463(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Gen_Class_sized__2241633463
}

type Constructor_Control_Monad_Gen_Class_MonadGen[T_m any] struct {
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
		c := (*Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "chooseBool": return gopurs_runtime.Box(c.V1)
		case "chooseFloat": return gopurs_runtime.Box(c.V2)
		case "chooseInt": return gopurs_runtime.Box(c.V3)
		case "resize": return gopurs_runtime.Box(c.V4)
		case "sized": return gopurs_runtime.Box(c.V5)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Gen_Class_MonadGen: " + key)
		}
	}
}


func Call_Control_Monad_Gen_Class_MonadGen_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Gen_Class_sized(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}

func Call_Control_Monad_Gen_Class_resize(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Control_Monad_Gen_Class_chooseInt(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Monad_Gen_Class_chooseFloat(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Monad_Gen_Class_chooseBool(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "chooseBool")
}

func Call_Control_Monad_Gen_Class_chooseFloat__1964853975(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Monad_Gen_Class_chooseInt__1063828903(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Monad_Gen_Class_resize__4113973243(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Control_Monad_Gen_Class_resize__1050945947(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Control_Monad_Gen_Class_resize__1313223195(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Control_Monad_Gen_Class_resize__2904398683(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V4)
}

func Call_Control_Monad_Gen_Class_sized__120035991(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}

func Call_Control_Monad_Gen_Class_sized__2830838711(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}

func Call_Control_Monad_Gen_Class_sized__3147117623(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}

func Call_Control_Monad_Gen_Class_sized__4206899191(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}

func Call_Control_Monad_Gen_Class_sized__2391211191(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}

func Call_Control_Monad_Gen_Class_sized__2241633463(dict_0_loop *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Gen_Class_MonadGen[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V5)
}



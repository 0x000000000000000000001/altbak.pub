package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Lazy_Lazy_dollarDict gopurs_runtime.Value
var once_Control_Lazy_Lazy_dollarDict sync.Once
func Get_Control_Lazy_Lazy_dollarDict() gopurs_runtime.Value {
	once_Control_Lazy_Lazy_dollarDict.Do(func() {
		cache_Control_Lazy_Lazy_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_Lazy_dollarDict(x_0_box)
})
	})
	return cache_Control_Lazy_Lazy_dollarDict
}

var cache_Control_Lazy_lazyUnit gopurs_runtime.Value
var once_Control_Lazy_lazyUnit sync.Once
func Get_Control_Lazy_lazyUnit() gopurs_runtime.Value {
	once_Control_Lazy_lazyUnit.Do(func() {
		cache_Control_Lazy_lazyUnit = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
	})
	return cache_Control_Lazy_lazyUnit
}

var cache_Control_Lazy_lazyFn gopurs_runtime.Value
var once_Control_Lazy_lazyFn sync.Once
func Get_Control_Lazy_lazyFn() gopurs_runtime.Value {
	once_Control_Lazy_lazyFn.Do(func() {
		cache_Control_Lazy_lazyFn = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, Get_Data_Unit_unit(), x_1)
})
}))
	})
	return cache_Control_Lazy_lazyFn
}

var cache_Control_Lazy_go__defer gopurs_runtime.Value
var once_Control_Lazy_go__defer sync.Once
func Get_Control_Lazy_go__defer() gopurs_runtime.Value {
	once_Control_Lazy_go__defer.Do(func() {
		cache_Control_Lazy_go__defer = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_go__defer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Lazy_go__defer
}

var cache_Control_Lazy_fix gopurs_runtime.Value
var once_Control_Lazy_fix sync.Once
func Get_Control_Lazy_fix() gopurs_runtime.Value {
	once_Control_Lazy_fix.Do(func() {
		cache_Control_Lazy_fix = gopurs_runtime.Func2(func(dictLazy_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_fix(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_0_box), f_1_box)
})
	})
	return cache_Control_Lazy_fix
}

var cache_Control_Lazy_defer__3258767445 gopurs_runtime.Value
var once_Control_Lazy_defer__3258767445 sync.Once
func Get_Control_Lazy_defer__3258767445() gopurs_runtime.Value {
	once_Control_Lazy_defer__3258767445.Do(func() {
		cache_Control_Lazy_defer__3258767445 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_defer__3258767445(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Lazy_defer__3258767445
}

var cache_Control_Lazy_defer__3734051733 gopurs_runtime.Value
var once_Control_Lazy_defer__3734051733 sync.Once
func Get_Control_Lazy_defer__3734051733() gopurs_runtime.Value {
	once_Control_Lazy_defer__3734051733.Do(func() {
		cache_Control_Lazy_defer__3734051733 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_defer__3734051733(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Lazy_defer__3734051733
}

var cache_Control_Lazy_defer__2812710261 gopurs_runtime.Value
var once_Control_Lazy_defer__2812710261 sync.Once
func Get_Control_Lazy_defer__2812710261() gopurs_runtime.Value {
	once_Control_Lazy_defer__2812710261.Do(func() {
		cache_Control_Lazy_defer__2812710261 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_defer__2812710261(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Lazy_defer__2812710261
}

var cache_Control_Lazy_defer__3967925939 gopurs_runtime.Value
var once_Control_Lazy_defer__3967925939 sync.Once
func Get_Control_Lazy_defer__3967925939() gopurs_runtime.Value {
	once_Control_Lazy_defer__3967925939.Do(func() {
		cache_Control_Lazy_defer__3967925939 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_defer__3967925939(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Lazy_defer__3967925939
}

var cache_Control_Lazy_defer__2590380358 gopurs_runtime.Value
var once_Control_Lazy_defer__2590380358 sync.Once
func Get_Control_Lazy_defer__2590380358() gopurs_runtime.Value {
	once_Control_Lazy_defer__2590380358.Do(func() {
		cache_Control_Lazy_defer__2590380358 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_defer__2590380358(f_0_box)
})
	})
	return cache_Control_Lazy_defer__2590380358
}

var cache_Control_Lazy_defer__308916730 gopurs_runtime.Value
var once_Control_Lazy_defer__308916730 sync.Once
func Get_Control_Lazy_defer__308916730() gopurs_runtime.Value {
	once_Control_Lazy_defer__308916730.Do(func() {
		cache_Control_Lazy_defer__308916730 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_defer__308916730(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Lazy_defer__308916730
}

var cache_Control_Lazy_defer__774977193 gopurs_runtime.Value
var once_Control_Lazy_defer__774977193 sync.Once
func Get_Control_Lazy_defer__774977193() gopurs_runtime.Value {
	once_Control_Lazy_defer__774977193.Do(func() {
		cache_Control_Lazy_defer__774977193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_defer__774977193(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Lazy_defer__774977193
}

var cache_Control_Lazy_defer__449752361 gopurs_runtime.Value
var once_Control_Lazy_defer__449752361 sync.Once
func Get_Control_Lazy_defer__449752361() gopurs_runtime.Value {
	once_Control_Lazy_defer__449752361.Do(func() {
		cache_Control_Lazy_defer__449752361 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_defer__449752361(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Lazy_defer__449752361
}

var cache_Control_Lazy_fix__1475205859 gopurs_runtime.Value
var once_Control_Lazy_fix__1475205859 sync.Once
func Get_Control_Lazy_fix__1475205859() gopurs_runtime.Value {
	once_Control_Lazy_fix__1475205859.Do(func() {
		cache_Control_Lazy_fix__1475205859 = gopurs_runtime.Func2(func(dictLazy_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_fix__1475205859(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_0_box), f_1_box)
})
	})
	return cache_Control_Lazy_fix__1475205859
}

var cache_Control_Lazy_fix__3570066147 gopurs_runtime.Value
var once_Control_Lazy_fix__3570066147 sync.Once
func Get_Control_Lazy_fix__3570066147() gopurs_runtime.Value {
	once_Control_Lazy_fix__3570066147.Do(func() {
		cache_Control_Lazy_fix__3570066147 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Lazy_fix__3570066147(f_0_box)
})
	})
	return cache_Control_Lazy_fix__3570066147
}

type Constructor_Control_Lazy_Lazy[T_l any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1860244333] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Lazy_Lazy[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "defer": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Control_Lazy_Lazy: " + key)
		}
	}
}


func Call_Control_Lazy_Lazy_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Lazy_go__defer(dict_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Lazy_fix(dictLazy_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
_ = dictLazy_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_0 gopurs_runtime.Value
_ = go__go_2_0_0
go__go_2_0_0 = gopurs_runtime.Apply(gopurs_runtime.Box(dictLazy_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, go__go_2_0_0)
}))
return go__go_2_0_0
}

func Call_Control_Lazy_defer__3258767445(dict_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Lazy_defer__3734051733(dict_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Lazy_defer__2812710261(dict_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Lazy_defer__3967925939(dict_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Lazy_defer__2590380358(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_List_Lazy_Types_step(), gopurs_runtime.Apply(f_0, x_1))
}))
}

func Call_Control_Lazy_defer__308916730(dict_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Lazy_defer__774977193(dict_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Lazy_defer__449752361(dict_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Lazy_fix__1475205859(dictLazy_0_loop *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 *Constructor_Control_Lazy_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
_ = dictLazy_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_1 gopurs_runtime.Value
_ = go__go_2_0_1
go__go_2_0_1 = gopurs_runtime.Apply(gopurs_runtime.Box(dictLazy_0.V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, go__go_2_0_1)
}))
return go__go_2_0_1
}

func Call_Control_Lazy_fix__3570066147(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_2 gopurs_runtime.Value
_ = go__go_1_0_2
go__go_1_0_2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Lazy_Types_lazyList(), "defer"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, go__go_1_0_2)
}))
return go__go_1_0_2
}



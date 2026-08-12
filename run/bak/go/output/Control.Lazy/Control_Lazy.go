package Control_Lazy

import (
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_lazyUnit gopurs_runtime.Value
var once_lazyUnit sync.Once
func Get_lazyUnit() gopurs_runtime.Value {
	once_lazyUnit.Do(func() {
		cache_lazyUnit = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return cache_lazyUnit
}

var cache_lazyFn gopurs_runtime.Value
var once_lazyFn sync.Once
func Get_lazyFn() gopurs_runtime.Value {
	once_lazyFn.Do(func() {
		cache_lazyFn = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), x_1)
})
}))
	})
	return cache_lazyFn
}

var cache_go__defer gopurs_runtime.Value
var once_go__defer sync.Once
func Get_go__defer() gopurs_runtime.Value {
	once_go__defer.Do(func() {
		cache_go__defer = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__defer(gopurs_runtime.CoerceToStruct[Constructor_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_go__defer
}

var cache_fix gopurs_runtime.Value
var once_fix sync.Once
func Get_fix() gopurs_runtime.Value {
	once_fix.Do(func() {
		cache_fix = gopurs_runtime.Func2(func(dictLazy_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fix(gopurs_runtime.CoerceToStruct[Constructor_Lazy[gopurs_runtime.Value]](dictLazy_0_box), f_1_box)
})
	})
	return cache_fix
}

var cache_defer__3967925939 gopurs_runtime.Value
var once_defer__3967925939 sync.Once
func Get_defer__3967925939() gopurs_runtime.Value {
	once_defer__3967925939.Do(func() {
		cache_defer__3967925939 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__3967925939(gopurs_runtime.CoerceToStruct[Constructor_Lazy[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_defer__3967925939
}

type Constructor_Lazy[T_l any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1860244333] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Lazy[gopurs_runtime.Value])(ptr)
		switch key {
		case "defer": return c.V0
		default: panic("Key not found in dictionary Constructor_Lazy: " + key)
		}
	}
}


func Call_go__defer(dict_0_loop *Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fix(dictLazy_0_loop *Constructor_Lazy[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 *Constructor_Lazy[gopurs_runtime.Value] = dictLazy_0_loop
_ = dictLazy_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__go_2_0_0 gopurs_runtime.Value
_ = go__go_2_0_0
go__go_2_0_0 = gopurs_runtime.Apply(dictLazy_0.V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, go__go_2_0_0)
}))
return go__go_2_0_0
}

func Call_defer__3967925939(dict_0_loop *Constructor_Lazy[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Lazy[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}



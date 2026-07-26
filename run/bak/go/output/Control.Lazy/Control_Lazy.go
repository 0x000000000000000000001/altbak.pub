package Control_Lazy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
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
		cache_lazyFn = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), x_1)
}))
	})
	return cache_lazyFn
}

var cache_defer_ gopurs_runtime.Value
var once_defer_ sync.Once
func Get_defer_() gopurs_runtime.Value {
	once_defer_.Do(func() {
		cache_defer_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer_(dict_0_box)
})
	})
	return cache_defer_
}

var cache_fix gopurs_runtime.Value
var once_fix sync.Once
func Get_fix() gopurs_runtime.Value {
	once_fix.Do(func() {
		cache_fix = gopurs_runtime.Func2(func(dictLazy_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fix(dictLazy_0_box, f_1_box)
})
	})
	return cache_fix
}

func Call_defer_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "defer")
}

func Call_fix(dictLazy_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictLazy_0 gopurs_runtime.Value = dictLazy_0_loop
_ = dictLazy_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_0, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, go__2_0)
}))
return go__2_0
}



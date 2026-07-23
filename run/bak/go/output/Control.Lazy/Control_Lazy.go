package Control_Lazy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var lazyUnit gopurs_runtime.Value
var once_lazyUnit sync.Once
func Get_lazyUnit() gopurs_runtime.Value {
	once_lazyUnit.Do(func() {
		lazyUnit = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
	})
	return lazyUnit
}

var lazyFn gopurs_runtime.Value
var once_lazyFn sync.Once
func Get_lazyFn() gopurs_runtime.Value {
	once_lazyFn.Do(func() {
		lazyFn = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), x_1)
}))
	})
	return lazyFn
}

var defer_ gopurs_runtime.Value
var once_defer_ sync.Once
func Get_defer_() gopurs_runtime.Value {
	once_defer_.Do(func() {
		defer_ = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "defer")
})
	})
	return defer_
}

var fix gopurs_runtime.Value
var once_fix sync.Once
func Get_fix() gopurs_runtime.Value {
	once_fix.Do(func() {
		fix = gopurs_runtime.Func2(func(dictLazy_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__2_0 gopurs_runtime.Value
_ = go__2_0
go__2_0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_0, "defer"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, go__2_0)
}))
return go__2_0
})
	})
	return fix
}



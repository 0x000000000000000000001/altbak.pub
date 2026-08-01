package Control_Alt

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_altArray gopurs_runtime.Value
var once_altArray sync.Once
func Get_altArray() gopurs_runtime.Value {
	once_altArray.Do(func() {
		cache_altArray = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"))
	})
	return cache_altArray
}

var cache_alt gopurs_runtime.Value
var once_alt sync.Once
func Get_alt() gopurs_runtime.Value {
	once_alt.Do(func() {
		cache_alt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt(dict_0_box)
})
	})
	return cache_alt
}

var cache_alt__gopurs_runtime_Value_3221591854 gopurs_runtime.Value
var once_alt__gopurs_runtime_Value_3221591854 sync.Once
func Get_alt__gopurs_runtime_Value_3221591854() gopurs_runtime.Value {
	once_alt__gopurs_runtime_Value_3221591854.Do(func() {
		cache_alt__gopurs_runtime_Value_3221591854 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__gopurs_runtime_Value_3221591854(dict_0_box)
})
	})
	return cache_alt__gopurs_runtime_Value_3221591854
}

func Call_alt(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "alt")
}

func Call_alt__gopurs_runtime_Value_3221591854(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "alt")
}



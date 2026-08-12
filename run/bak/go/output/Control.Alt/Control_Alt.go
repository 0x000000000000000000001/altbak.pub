package Control_Alt

import (
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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

var cache_altArray__gopurs_runtime_Value_2010533188 gopurs_runtime.Value
var once_altArray__gopurs_runtime_Value_2010533188 sync.Once
func Get_altArray__gopurs_runtime_Value_2010533188() gopurs_runtime.Value {
	once_altArray__gopurs_runtime_Value_2010533188.Do(func() {
		cache_altArray__gopurs_runtime_Value_2010533188 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray__gopurs_runtime_Value_361387505()
}), gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"))
	})
	return cache_altArray__gopurs_runtime_Value_2010533188
}

var cache_alt gopurs_runtime.Value
var once_alt sync.Once
func Get_alt() gopurs_runtime.Value {
	once_alt.Do(func() {
		cache_alt = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt(gopurs_runtime.CoerceToStruct[Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt
}

var cache_alt__gopurs_runtime_Value_267341625 gopurs_runtime.Value
var once_alt__gopurs_runtime_Value_267341625 sync.Once
func Get_alt__gopurs_runtime_Value_267341625() gopurs_runtime.Value {
	once_alt__gopurs_runtime_Value_267341625.Do(func() {
		cache_alt__gopurs_runtime_Value_267341625 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__gopurs_runtime_Value_267341625(gopurs_runtime.CoerceToStruct[Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__gopurs_runtime_Value_267341625
}

type Constructor_Alt[T_f any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[4060500237] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Alt[gopurs_runtime.Value])(ptr)
		switch key {
		case "Functor0": return c.V0
		case "alt": return c.V1
		default: panic("Key not found in dictionary Constructor_Alt: " + key)
		}
	}
}


func Call_alt(dict_0_loop *Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_alt__gopurs_runtime_Value_267341625(dict_0_loop *Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



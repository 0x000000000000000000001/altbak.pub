package Control_Category

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(dict_0_box)
})
	})
	return cache_identity
}

var cache_categoryFn gopurs_runtime.Value
var once_categoryFn sync.Once
func Get_categoryFn() gopurs_runtime.Value {
	once_categoryFn.Do(func() {
		cache_categoryFn = gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Semigroupoid.Get_semigroupoidFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_categoryFn
}

var cache_semigroupoidFn__3002128382 gopurs_runtime.Value
var once_semigroupoidFn__3002128382 sync.Once
func Get_semigroupoidFn__3002128382() gopurs_runtime.Value {
	once_semigroupoidFn__3002128382.Do(func() {
		cache_semigroupoidFn__3002128382 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__3002128382
}

type Constructor_Category[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[784524589] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Category[gopurs_runtime.Value])(ptr)
		switch key {
		case "Semigroupoid0": return c.V0
		case "identity": return c.V1
		default: panic("Key not found in dictionary Constructor_Category: " + key)
		}
	}
}


func Call_identity(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "identity")
}



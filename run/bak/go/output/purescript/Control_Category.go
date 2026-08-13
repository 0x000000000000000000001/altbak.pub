package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Category_Category_dollarDict gopurs_runtime.Value
var once_Control_Category_Category_dollarDict sync.Once
func Get_Control_Category_Category_dollarDict() gopurs_runtime.Value {
	once_Control_Category_Category_dollarDict.Do(func() {
		cache_Control_Category_Category_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Category_Category_dollarDict(x_0_box)
})
	})
	return cache_Control_Category_Category_dollarDict
}

var cache_Control_Category_identity gopurs_runtime.Value
var once_Control_Category_identity sync.Once
func Get_Control_Category_identity() gopurs_runtime.Value {
	once_Control_Category_identity.Do(func() {
		cache_Control_Category_identity = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Category_identity(dict_0_box)
})
	})
	return cache_Control_Category_identity
}

var cache_Control_Category_categoryFn gopurs_runtime.Value
var once_Control_Category_categoryFn sync.Once
func Get_Control_Category_categoryFn() gopurs_runtime.Value {
	once_Control_Category_categoryFn.Do(func() {
		cache_Control_Category_categoryFn = gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Semigroupoid_semigroupoidFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Control_Category_categoryFn
}

var cache_Control_Category_categoryFn__3492036198 gopurs_runtime.Value
var once_Control_Category_categoryFn__3492036198 sync.Once
func Get_Control_Category_categoryFn__3492036198() gopurs_runtime.Value {
	once_Control_Category_categoryFn__3492036198.Do(func() {
		cache_Control_Category_categoryFn__3492036198 = gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Semigroupoid_semigroupoidFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Control_Category_categoryFn__3492036198
}

var cache_Control_Category_identity__2527656589 gopurs_runtime.Value
var once_Control_Category_identity__2527656589 sync.Once
func Get_Control_Category_identity__2527656589() gopurs_runtime.Value {
	once_Control_Category_identity__2527656589.Do(func() {
		cache_Control_Category_identity__2527656589 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Category_identity__2527656589(dict_0_box)
})
	})
	return cache_Control_Category_identity__2527656589
}

type Constructor_Control_Category_Category struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[784524589] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Category_Category)(ptr)
		_ = c
		switch key {
		case "Semigroupoid0": return gopurs_runtime.Box(c.V0)
		case "identity": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Category_Category: " + key)
		}
	}
}


func Call_Control_Category_Category_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Category_identity(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "identity")
}

func Call_Control_Category_identity__2527656589(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "identity")
}



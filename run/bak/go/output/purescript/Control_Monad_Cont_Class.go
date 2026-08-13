package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Cont_Class_MonadCont_dollarDict gopurs_runtime.Value
var once_Control_Monad_Cont_Class_MonadCont_dollarDict sync.Once
func Get_Control_Monad_Cont_Class_MonadCont_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_MonadCont_dollarDict.Do(func() {
		cache_Control_Monad_Cont_Class_MonadCont_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Class_MonadCont_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Cont_Class_MonadCont_dollarDict
}

var cache_Control_Monad_Cont_Class_callCC gopurs_runtime.Value
var once_Control_Monad_Cont_Class_callCC sync.Once
func Get_Control_Monad_Cont_Class_callCC() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_callCC.Do(func() {
		cache_Control_Monad_Cont_Class_callCC = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Class_callCC(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Cont_Class_callCC
}

var cache_Control_Monad_Cont_Class_callCC__1888484333 gopurs_runtime.Value
var once_Control_Monad_Cont_Class_callCC__1888484333 sync.Once
func Get_Control_Monad_Cont_Class_callCC__1888484333() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_callCC__1888484333.Do(func() {
		cache_Control_Monad_Cont_Class_callCC__1888484333 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Class_callCC__1888484333(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Cont_Class_callCC__1888484333
}

var cache_Control_Monad_Cont_Class_callCC__1963329157 gopurs_runtime.Value
var once_Control_Monad_Cont_Class_callCC__1963329157 sync.Once
func Get_Control_Monad_Cont_Class_callCC__1963329157() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_callCC__1963329157.Do(func() {
		cache_Control_Monad_Cont_Class_callCC__1963329157 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Class_callCC__1963329157(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Cont_Class_callCC__1963329157
}

var cache_Control_Monad_Cont_Class_callCC__2474776556 gopurs_runtime.Value
var once_Control_Monad_Cont_Class_callCC__2474776556 sync.Once
func Get_Control_Monad_Cont_Class_callCC__2474776556() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_callCC__2474776556.Do(func() {
		cache_Control_Monad_Cont_Class_callCC__2474776556 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Class_callCC__2474776556(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Cont_Class_callCC__2474776556
}

var cache_Control_Monad_Cont_Class_callCC__2318135621 gopurs_runtime.Value
var once_Control_Monad_Cont_Class_callCC__2318135621 sync.Once
func Get_Control_Monad_Cont_Class_callCC__2318135621() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_callCC__2318135621.Do(func() {
		cache_Control_Monad_Cont_Class_callCC__2318135621 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Class_callCC__2318135621(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Cont_Class_callCC__2318135621
}

type Constructor_Control_Monad_Cont_Class_MonadCont[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1800060259] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "callCC": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Cont_Class_MonadCont: " + key)
		}
	}
}


func Call_Control_Monad_Cont_Class_MonadCont_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Cont_Class_callCC(dict_0_loop *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Cont_Class_callCC__1888484333(dict_0_loop *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Cont_Class_callCC__1963329157(dict_0_loop *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Cont_Class_callCC__2474776556(dict_0_loop *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Cont_Class_callCC__2318135621(dict_0_loop *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



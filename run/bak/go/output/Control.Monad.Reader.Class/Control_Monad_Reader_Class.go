package Control_Monad_Reader_Class

import (
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_monadAskFun gopurs_runtime.Value
var once_monadAskFun sync.Once
func Get_monadAskFun() gopurs_runtime.Value {
	once_monadAskFun.Do(func() {
		cache_monadAskFun = gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_monadAskFun
}

var cache_monadAskFun__gopurs_runtime_Value_466477709 gopurs_runtime.Value
var once_monadAskFun__gopurs_runtime_Value_466477709 sync.Once
func Get_monadAskFun__gopurs_runtime_Value_466477709() gopurs_runtime.Value {
	once_monadAskFun__gopurs_runtime_Value_466477709.Do(func() {
		cache_monadAskFun__gopurs_runtime_Value_466477709 = gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadFn__gopurs_runtime_Value_1938941618()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_monadAskFun__gopurs_runtime_Value_466477709
}

var cache_monadReaderFun gopurs_runtime.Value
var once_monadReaderFun sync.Once
func Get_monadReaderFun() gopurs_runtime.Value {
	once_monadReaderFun.Do(func() {
		cache_monadReaderFun = gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAskFun()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(f_0, x_2))
})
})
}))
	})
	return cache_monadReaderFun
}

var cache_local gopurs_runtime.Value
var once_local sync.Once
func Get_local() gopurs_runtime.Value {
	once_local.Do(func() {
		cache_local = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local(gopurs_runtime.CoerceToStruct[Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_local
}

var cache_local__gopurs_runtime_Value_1299460031 gopurs_runtime.Value
var once_local__gopurs_runtime_Value_1299460031 sync.Once
func Get_local__gopurs_runtime_Value_1299460031() gopurs_runtime.Value {
	once_local__gopurs_runtime_Value_1299460031.Do(func() {
		cache_local__gopurs_runtime_Value_1299460031 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local__gopurs_runtime_Value_1299460031(gopurs_runtime.CoerceToStruct[Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_local__gopurs_runtime_Value_1299460031
}

var cache_ask gopurs_runtime.Value
var once_ask sync.Once
func Get_ask() gopurs_runtime.Value {
	once_ask.Do(func() {
		cache_ask = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ask(dict_0_box)
})
	})
	return cache_ask
}

var cache_asks gopurs_runtime.Value
var once_asks sync.Once
func Get_asks() gopurs_runtime.Value {
	once_asks.Do(func() {
		cache_asks = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_asks(gopurs_runtime.CoerceToStruct[Constructor_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadAsk_0_box))
})
	})
	return cache_asks
}

type Constructor_MonadAsk[T_r any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1229730751] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad0": return c.V0
		case "ask": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadAsk: " + key)
		}
	}
}


type Constructor_MonadReader[T_r any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2457234979] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "MonadAsk0": return c.V0
		case "local": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadReader: " + key)
		}
	}
}


func Call_local(dict_0_loop *Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_local__gopurs_runtime_Value_1299460031(dict_0_loop *Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_ask(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ask")
}

func Call_asks(dictMonadAsk_0_loop *Constructor_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadAsk_0 *Constructor_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadAsk_0_loop
_ = dictMonadAsk_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadAsk_0.V0, gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
ask1_2_1 := dictMonadAsk_0.V1
_ = ask1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, f_3, ask1_2_1)
})
}



package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Reader_Class_MonadAsk_dollarDict gopurs_runtime.Value
var once_Control_Monad_Reader_Class_MonadAsk_dollarDict sync.Once
func Get_Control_Monad_Reader_Class_MonadAsk_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_MonadAsk_dollarDict.Do(func() {
		cache_Control_Monad_Reader_Class_MonadAsk_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_MonadAsk_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Reader_Class_MonadAsk_dollarDict
}

var cache_Control_Monad_Reader_Class_MonadReader_dollarDict gopurs_runtime.Value
var once_Control_Monad_Reader_Class_MonadReader_dollarDict sync.Once
func Get_Control_Monad_Reader_Class_MonadReader_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_MonadReader_dollarDict.Do(func() {
		cache_Control_Monad_Reader_Class_MonadReader_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_MonadReader_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Reader_Class_MonadReader_dollarDict
}

var cache_Control_Monad_Reader_Class_monadAskFun gopurs_runtime.Value
var once_Control_Monad_Reader_Class_monadAskFun sync.Once
func Get_Control_Monad_Reader_Class_monadAskFun() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_monadAskFun.Do(func() {
		cache_Control_Monad_Reader_Class_monadAskFun = gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_monadFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Control_Monad_Reader_Class_monadAskFun
}

var cache_Control_Monad_Reader_Class_monadReaderFun gopurs_runtime.Value
var once_Control_Monad_Reader_Class_monadReaderFun sync.Once
func Get_Control_Monad_Reader_Class_monadReaderFun() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_monadReaderFun.Do(func() {
		cache_Control_Monad_Reader_Class_monadReaderFun = gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_Reader_Class_monadAskFun()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(f_0, x_2))
})
})
}))
	})
	return cache_Control_Monad_Reader_Class_monadReaderFun
}

var cache_Control_Monad_Reader_Class_local gopurs_runtime.Value
var once_Control_Monad_Reader_Class_local sync.Once
func Get_Control_Monad_Reader_Class_local() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_local.Do(func() {
		cache_Control_Monad_Reader_Class_local = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_local(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader](dict_0_box))
})
	})
	return cache_Control_Monad_Reader_Class_local
}

var cache_Control_Monad_Reader_Class_ask gopurs_runtime.Value
var once_Control_Monad_Reader_Class_ask sync.Once
func Get_Control_Monad_Reader_Class_ask() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_ask.Do(func() {
		cache_Control_Monad_Reader_Class_ask = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_ask(dict_0_box)
})
	})
	return cache_Control_Monad_Reader_Class_ask
}

var cache_Control_Monad_Reader_Class_asks gopurs_runtime.Value
var once_Control_Monad_Reader_Class_asks sync.Once
func Get_Control_Monad_Reader_Class_asks() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_asks.Do(func() {
		cache_Control_Monad_Reader_Class_asks = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_asks(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk](dictMonadAsk_0_box))
})
	})
	return cache_Control_Monad_Reader_Class_asks
}

var cache_Control_Monad_Reader_Class_local__1299460031 gopurs_runtime.Value
var once_Control_Monad_Reader_Class_local__1299460031 sync.Once
func Get_Control_Monad_Reader_Class_local__1299460031() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_local__1299460031.Do(func() {
		cache_Control_Monad_Reader_Class_local__1299460031 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_local__1299460031(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader](dict_0_box))
})
	})
	return cache_Control_Monad_Reader_Class_local__1299460031
}

var cache_Control_Monad_Reader_Class_local__190530239 gopurs_runtime.Value
var once_Control_Monad_Reader_Class_local__190530239 sync.Once
func Get_Control_Monad_Reader_Class_local__190530239() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_local__190530239.Do(func() {
		cache_Control_Monad_Reader_Class_local__190530239 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_local__190530239(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader](dict_0_box))
})
	})
	return cache_Control_Monad_Reader_Class_local__190530239
}

var cache_Control_Monad_Reader_Class_local__4056952415 gopurs_runtime.Value
var once_Control_Monad_Reader_Class_local__4056952415 sync.Once
func Get_Control_Monad_Reader_Class_local__4056952415() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_local__4056952415.Do(func() {
		cache_Control_Monad_Reader_Class_local__4056952415 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_local__4056952415(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader](dict_0_box))
})
	})
	return cache_Control_Monad_Reader_Class_local__4056952415
}

var cache_Control_Monad_Reader_Class_local__909940799 gopurs_runtime.Value
var once_Control_Monad_Reader_Class_local__909940799 sync.Once
func Get_Control_Monad_Reader_Class_local__909940799() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_local__909940799.Do(func() {
		cache_Control_Monad_Reader_Class_local__909940799 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_local__909940799(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader](dict_0_box))
})
	})
	return cache_Control_Monad_Reader_Class_local__909940799
}

var cache_Control_Monad_Reader_Class_monadAskFun__466477709 gopurs_runtime.Value
var once_Control_Monad_Reader_Class_monadAskFun__466477709 sync.Once
func Get_Control_Monad_Reader_Class_monadAskFun__466477709() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_monadAskFun__466477709.Do(func() {
		cache_Control_Monad_Reader_Class_monadAskFun__466477709 = gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_monadFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Control_Monad_Reader_Class_monadAskFun__466477709
}

type Constructor_Control_Monad_Reader_Class_MonadAsk struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1229730751] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Reader_Class_MonadAsk)(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "ask": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Reader_Class_MonadAsk: " + key)
		}
	}
}


type Constructor_Control_Monad_Reader_Class_MonadReader struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2457234979] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Reader_Class_MonadReader)(ptr)
		_ = c
		switch key {
		case "MonadAsk0": return gopurs_runtime.Box(c.V0)
		case "local": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Reader_Class_MonadReader: " + key)
		}
	}
}


func Call_Control_Monad_Reader_Class_MonadAsk_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Reader_Class_MonadReader_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Reader_Class_local(dict_0_loop *Constructor_Control_Monad_Reader_Class_MonadReader) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Reader_Class_MonadReader = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Reader_Class_ask(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ask")
}

func Call_Control_Monad_Reader_Class_asks(dictMonadAsk_0_loop *Constructor_Control_Monad_Reader_Class_MonadAsk) gopurs_runtime.Value {
var dictMonadAsk_0 *Constructor_Control_Monad_Reader_Class_MonadAsk = dictMonadAsk_0_loop
_ = dictMonadAsk_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadAsk_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): ask1_2_1 -> gopurs_runtime.Value
ask1_2_1 := gopurs_runtime.Box(dictMonadAsk_0.V1)
_ = ask1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_3, ask1_2_1)
})
}

func Call_Control_Monad_Reader_Class_local__1299460031(dict_0_loop *Constructor_Control_Monad_Reader_Class_MonadReader) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Reader_Class_MonadReader = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Reader_Class_local__190530239(dict_0_loop *Constructor_Control_Monad_Reader_Class_MonadReader) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Reader_Class_MonadReader = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Reader_Class_local__4056952415(dict_0_loop *Constructor_Control_Monad_Reader_Class_MonadReader) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Reader_Class_MonadReader = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Reader_Class_local__909940799(dict_0_loop *Constructor_Control_Monad_Reader_Class_MonadReader) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Reader_Class_MonadReader = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



package Control_Monad_Reader_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad "gopurs/output/Control.Monad"
)

var cache_monadAskFun gopurs_runtime.Value
var once_monadAskFun sync.Once
func Get_monadAskFun() gopurs_runtime.Value {
	once_monadAskFun.Do(func() {
		cache_monadAskFun = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Monad0", "ask", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))))
	})
	return cache_monadAskFun
}

var cache_monadReaderFun gopurs_runtime.Value
var once_monadReaderFun sync.Once
func Get_monadReaderFun() gopurs_runtime.Value {
	once_monadReaderFun.Do(func() {
		cache_monadReaderFun = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("MonadAsk0", "local", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAskFun()
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(f_0, x_2))
}))))
	})
	return cache_monadReaderFun
}

var cache_local gopurs_runtime.Value
var once_local sync.Once
func Get_local() gopurs_runtime.Value {
	once_local.Do(func() {
		cache_local = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local(dict_0_box)
})
	})
	return cache_local
}

var cache_local__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2762410074 gopurs_runtime.Value
var once_local__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2762410074 sync.Once
func Get_local__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2762410074() gopurs_runtime.Value {
	once_local__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2762410074.Do(func() {
		cache_local__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2762410074 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2762410074(dict_0_box)
})
	})
	return cache_local__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2762410074
}

var cache_ask gopurs_runtime.Value
var once_ask sync.Once
func Get_ask() gopurs_runtime.Value {
	once_ask.Do(func() {
		cache_ask = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_ask(dict_0_box))
})
	})
	return cache_ask
}

var cache_asks gopurs_runtime.Value
var once_asks sync.Once
func Get_asks() gopurs_runtime.Value {
	once_asks.Do(func() {
		cache_asks = gopurs_runtime.Func(func(dictMonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_asks(dictMonadAsk_0_box)
})
	})
	return cache_asks
}

func Call_local(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "local")
}

func Call_local__func_gopurs_runtime_Value__func_interface____interface____interface____interface___2762410074(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "local")
}

func Call_ask(dict_0_loop gopurs_runtime.Value) interface{} {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(dict_0, "ask"))
}

func Call_asks(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
ask1_1_0 := gopurs_runtime.RecordGet(dictMonadAsk_0, "ask")
_ = ask1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), f_2, ask1_1_0)
})
}

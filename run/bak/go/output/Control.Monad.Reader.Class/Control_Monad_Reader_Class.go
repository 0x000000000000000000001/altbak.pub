package Control_Monad_Reader_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Category "gopurs/output/Control.Category"
	pkg_Control_Monad "gopurs/output/Control.Monad"
)

var monadAskFun gopurs_runtime.Value
var once_monadAskFun sync.Once
func Get_monadAskFun() gopurs_runtime.Value {
	once_monadAskFun.Do(func() {
		monadAskFun = gopurs_runtime.RecordDict2("ask", "Monad0", gopurs_runtime.RecordGet(pkg_Control_Category.Get_categoryFn(), "identity"), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadFn()
}))
	})
	return monadAskFun
}

var monadReaderFun gopurs_runtime.Value
var once_monadReaderFun sync.Once
func Get_monadReaderFun() gopurs_runtime.Value {
	once_monadReaderFun.Do(func() {
		monadReaderFun = gopurs_runtime.RecordDict2("local", "MonadAsk0", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(f_0, x_2))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAskFun()
}))
	})
	return monadReaderFun
}

var local gopurs_runtime.Value
var once_local sync.Once
func Get_local() gopurs_runtime.Value {
	once_local.Do(func() {
		local = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "local")
}()
})
	})
	return local
}

var ask gopurs_runtime.Value
var once_ask sync.Once
func Get_ask() gopurs_runtime.Value {
	once_ask.Do(func() {
		ask = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "ask")
}()
})
	})
	return ask
}

var asks gopurs_runtime.Value
var once_asks sync.Once
func Get_asks() gopurs_runtime.Value {
	once_asks.Do(func() {
		asks = gopurs_runtime.Func(func(dictMonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadAsk_0 gopurs_runtime.Value = dictMonadAsk_0_loop
_ = dictMonadAsk_0
ask1_1_0 := gopurs_runtime.RecordGet(dictMonadAsk_0_loop, "ask")
_ = ask1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadAsk_0_loop, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), f_2, ask1_1_0)
})
}()
})
	})
	return asks
}





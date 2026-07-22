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
		monadAskFun = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"ask": pkg_Control_Category.Get_categoryFn().PtrVal.(map[string]gopurs_runtime.Value)["identity"], "Monad0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad.Get_monadFn()
})})
	})
	return monadAskFun
}

var monadReaderFun gopurs_runtime.Value
var once_monadReaderFun sync.Once
func Get_monadReaderFun() gopurs_runtime.Value {
	once_monadReaderFun.Do(func() {
		monadReaderFun = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"local": gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(f_0, x_2))
})
})
}), "MonadAsk0": gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_monadAskFun()
})})
	})
	return monadReaderFun
}

var local gopurs_runtime.Value
var once_local sync.Once
func Get_local() gopurs_runtime.Value {
	once_local.Do(func() {
		local = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["local"]
})
	})
	return local
}

var ask gopurs_runtime.Value
var once_ask sync.Once
func Get_ask() gopurs_runtime.Value {
	once_ask.Do(func() {
		ask = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["ask"]
})
	})
	return ask
}

var asks gopurs_runtime.Value
var once_asks sync.Once
func Get_asks() gopurs_runtime.Value {
	once_asks.Do(func() {
		asks = gopurs_runtime.Func(func(dictMonadAsk_0 gopurs_runtime.Value) gopurs_runtime.Value {
ask1_1_0 := dictMonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["ask"]
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadAsk_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], f_2), ask1_1_0)
})
})
	})
	return asks
}



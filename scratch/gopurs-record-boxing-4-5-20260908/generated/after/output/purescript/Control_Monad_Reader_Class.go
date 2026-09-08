package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Reader_Class_MonadAsk_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Reader_Class_MonadAsk_dollar_Dict sync.Once
func Get_Control_Monad_Reader_Class_MonadAsk_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_MonadAsk_dollar_Dict.Do(func() {
		cache_Control_Monad_Reader_Class_MonadAsk_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Reader_Class_MonadAsk_dollar_Dict(func() struct{
	Monad0 gopurs_runtime.Value
	ask gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Monad0 gopurs_runtime.Value
	ask gopurs_runtime.Value
}{}
					clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
					clone.ask = gopurs_runtime.RecordGet(orig, "ask")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Reader_Class_MonadAsk_dollar_Dict
}

var cache_Control_Monad_Reader_Class_MonadReader_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Reader_Class_MonadReader_dollar_Dict sync.Once
func Get_Control_Monad_Reader_Class_MonadReader_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_MonadReader_dollar_Dict.Do(func() {
		cache_Control_Monad_Reader_Class_MonadReader_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Reader_Class_MonadReader_dollar_Dict(func() struct{
	MonadAsk0 gopurs_runtime.Value
	local gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	MonadAsk0 gopurs_runtime.Value
	local gopurs_runtime.Value
}{}
					clone.MonadAsk0 = gopurs_runtime.RecordGet(orig, "MonadAsk0")
					clone.local = gopurs_runtime.RecordGet(orig, "local")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Reader_Class_MonadReader_dollar_Dict
}

var cache_Control_Monad_Reader_Class_monadAskFun gopurs_runtime.Value
var once_Control_Monad_Reader_Class_monadAskFun sync.Once
func Get_Control_Monad_Reader_Class_monadAskFun() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_monadAskFun.Do(func() {
		cache_Control_Monad_Reader_Class_monadAskFun = gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_monadFn()))}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)}))}
	})
	return cache_Control_Monad_Reader_Class_monadAskFun
}

var cache_Control_Monad_Reader_Class_monadReaderFun gopurs_runtime.Value
var once_Control_Monad_Reader_Class_monadReaderFun sync.Once
func Get_Control_Monad_Reader_Class_monadReaderFun() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_monadReaderFun.Do(func() {
		cache_Control_Monad_Reader_Class_monadReaderFun = gopurs_runtime.Value{Type: 9, IntVal: 2457234979, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1229730751, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Control_Monad_Reader_Class_monadAskFun()))}
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(f_0, x_2))
})}))}
	})
	return cache_Control_Monad_Reader_Class_monadReaderFun
}

var cache_Control_Monad_Reader_Class_local gopurs_runtime.Value
var once_Control_Monad_Reader_Class_local sync.Once
func Get_Control_Monad_Reader_Class_local() gopurs_runtime.Value {
	once_Control_Monad_Reader_Class_local.Do(func() {
		cache_Control_Monad_Reader_Class_local = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Reader_Class_local(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
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
return Call_Control_Monad_Reader_Class_asks(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadAsk_0_box))
})
	})
	return cache_Control_Monad_Reader_Class_asks
}

type Constructor_Control_Monad_Reader_Class_MonadAsk[T_r any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1229730751] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Reader_Class_MonadAsk[any, any])(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "ask": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Reader_Class_MonadAsk: " + key)
		}
	}
}


type Constructor_Control_Monad_Reader_Class_MonadReader[T_r any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2457234979] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Reader_Class_MonadReader[any, any])(ptr)
		_ = c
		switch key {
		case "MonadAsk0": return gopurs_runtime.Box(c.V0)
		case "local": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Reader_Class_MonadReader: " + key)
		}
	}
}


func Call_Control_Monad_Reader_Class_MonadAsk_dollar_Dict(x_0_loop struct{
	Monad0 gopurs_runtime.Value
	ask gopurs_runtime.Value
}) *Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	Monad0 gopurs_runtime.Value
	ask gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Monad0", "ask", orig.Monad0, orig.ask)
				}())
}

func Call_Control_Monad_Reader_Class_MonadReader_dollar_Dict(x_0_loop struct{
	MonadAsk0 gopurs_runtime.Value
	local gopurs_runtime.Value
}) *Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	MonadAsk0 gopurs_runtime.Value
	local gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("MonadAsk0", "local", orig.MonadAsk0, orig.local)
				}())
}

func Call_Control_Monad_Reader_Class_local(dict_0_loop *Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Reader_Class_MonadReader[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Monad_Reader_Class_ask(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ask")
}

func Call_Control_Monad_Reader_Class_asks(dictMonadAsk_0_loop *Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadAsk_0 *Constructor_Control_Monad_Reader_Class_MonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadAsk_0_loop
_ = dictMonadAsk_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadAsk_0.V0), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): ask1_2_1 shape=Other bindingType=(TypeApp (TypeVar m) [(TypeVar r)])
ask1_2_1 := gopurs_runtime.Box(dictMonadAsk_0.V1)
_ = ask1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_3, ask1_2_1)
})
}



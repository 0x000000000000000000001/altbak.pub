package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Env_Class_lower gopurs_runtime.Value
var once_Control_Comonad_Env_Class_lower sync.Once
func Get_Control_Comonad_Env_Class_lower() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_lower.Do(func() {
		cache_Control_Comonad_Env_Class_lower = gopurs_runtime.RecordGet(Get_Control_Comonad_Store_Trans_comonadTransStoreT(), "lower")
	})
	return cache_Control_Comonad_Env_Class_lower
}

var cache_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict sync.Once
func Get_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict.Do(func() {
		cache_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict(func() struct{
	Comonad0 gopurs_runtime.Value
	ask gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Comonad0 gopurs_runtime.Value
	ask gopurs_runtime.Value
}{}
					clone.Comonad0 = gopurs_runtime.RecordGet(orig, "Comonad0")
					clone.ask = gopurs_runtime.RecordGet(orig, "ask")
					return clone
				}()))}
})
	})
	return cache_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict
}

var cache_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict sync.Once
func Get_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict.Do(func() {
		cache_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict(func() struct{
	ComonadAsk0 gopurs_runtime.Value
	local gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	ComonadAsk0 gopurs_runtime.Value
	local gopurs_runtime.Value
}{}
					clone.ComonadAsk0 = gopurs_runtime.RecordGet(orig, "ComonadAsk0")
					clone.local = gopurs_runtime.RecordGet(orig, "local")
					return clone
				}()))}
})
	})
	return cache_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict
}

var cache_Control_Comonad_Env_Class_local gopurs_runtime.Value
var once_Control_Comonad_Env_Class_local sync.Once
func Get_Control_Comonad_Env_Class_local() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_local.Do(func() {
		cache_Control_Comonad_Env_Class_local = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_local(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_local
}

var cache_Control_Comonad_Env_Class_comonadAskTuple gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskTuple sync.Once
func Get_Control_Comonad_Env_Class_comonadAskTuple() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskTuple.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskTuple = gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037((&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_3056445460_2550391993(Rebox_Control_Comonad_Env_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](Get_Data_Tuple_comonadTuple()))))}
}), Get_Data_Tuple_fst()})))}
	})
	return cache_Control_Comonad_Env_Class_comonadAskTuple
}

var cache_Control_Comonad_Env_Class_comonadEnvTuple gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvTuple sync.Once
func Get_Control_Comonad_Env_Class_comonadEnvTuple() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvTuple.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvTuple = gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_144036252_155172497((&Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037(Rebox_Control_Comonad_Env_Class_2829962037_1649752824(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Control_Comonad_Env_Class_comonadAskTuple()))))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
})})))}
	})
	return cache_Control_Comonad_Env_Class_comonadEnvTuple
}

var cache_Control_Comonad_Env_Class_comonadAskEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskEnvT sync.Once
func Get_Control_Comonad_Env_Class_comonadAskEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskEnvT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadAskEnvT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadAskEnvT
}

var cache_Control_Comonad_Env_Class_comonadEnvEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvEnvT sync.Once
func Get_Control_Comonad_Env_Class_comonadEnvEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvEnvT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadEnvEnvT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadEnvEnvT
}

var cache_Control_Comonad_Env_Class_ask gopurs_runtime.Value
var once_Control_Comonad_Env_Class_ask sync.Once
func Get_Control_Comonad_Env_Class_ask() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_ask.Do(func() {
		cache_Control_Comonad_Env_Class_ask = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_ask(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Comonad_Env_Class_ask
}

var cache_Control_Comonad_Env_Class_asks gopurs_runtime.Value
var once_Control_Comonad_Env_Class_asks sync.Once
func Get_Control_Comonad_Env_Class_asks() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_asks.Do(func() {
		cache_Control_Comonad_Env_Class_asks = gopurs_runtime.Func3(func(dictComonadAsk_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_asks(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](dictComonadAsk_0_box), f_1_box, x_2_box)
})
	})
	return cache_Control_Comonad_Env_Class_asks
}

var cache_Control_Comonad_Env_Class_comonadAskStoreT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskStoreT sync.Once
func Get_Control_Comonad_Env_Class_comonadAskStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskStoreT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskStoreT = gopurs_runtime.Func(func(dictComonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadAskStoreT(dictComonadAsk_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadAskStoreT
}

var cache_Control_Comonad_Env_Class_comonadEnvStoreT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvStoreT sync.Once
func Get_Control_Comonad_Env_Class_comonadEnvStoreT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvStoreT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvStoreT = gopurs_runtime.Func(func(dictComonadEnv_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadEnvStoreT(dictComonadEnv_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadEnvStoreT
}

var cache_Control_Comonad_Env_Class_comonadAskTracedT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadAskTracedT sync.Once
func Get_Control_Comonad_Env_Class_comonadAskTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadAskTracedT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadAskTracedT = gopurs_runtime.Func(func(dictComonadAsk_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadAskTracedT(dictComonadAsk_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadAskTracedT
}

var cache_Control_Comonad_Env_Class_comonadEnvTracedT gopurs_runtime.Value
var once_Control_Comonad_Env_Class_comonadEnvTracedT sync.Once
func Get_Control_Comonad_Env_Class_comonadEnvTracedT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Class_comonadEnvTracedT.Do(func() {
		cache_Control_Comonad_Env_Class_comonadEnvTracedT = gopurs_runtime.Func(func(dictComonadEnv_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Class_comonadEnvTracedT(dictComonadEnv_0_box)
})
	})
	return cache_Control_Comonad_Env_Class_comonadEnvTracedT
}

type Constructor_Control_Comonad_Env_Class_ComonadAsk[T_e any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1424092807] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Env_Class_ComonadAsk[any, any])(ptr)
		_ = c
		switch key {
		case "Comonad0": return gopurs_runtime.Box(c.V0)
		case "ask": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Env_Class_ComonadAsk: " + key)
		}
	}
}


type Constructor_Control_Comonad_Env_Class_ComonadEnv[T_e any, T_w any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3863290147] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Env_Class_ComonadEnv[any, any])(ptr)
		_ = c
		switch key {
		case "ComonadAsk0": return gopurs_runtime.Box(c.V0)
		case "local": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Env_Class_ComonadEnv: " + key)
		}
	}
}


func Call_Control_Comonad_Env_Class_ComonadAsk_dollar_Dict(x_0_loop struct{
	Comonad0 gopurs_runtime.Value
	ask gopurs_runtime.Value
}) *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	Comonad0 gopurs_runtime.Value
	ask gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("Comonad0", "ask", orig.Comonad0, orig.ask)
				}())
}

func Call_Control_Comonad_Env_Class_ComonadEnv_dollar_Dict(x_0_loop struct{
	ComonadAsk0 gopurs_runtime.Value
	local gopurs_runtime.Value
}) *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	ComonadAsk0 gopurs_runtime.Value
	local gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict2("ComonadAsk0", "local", orig.ComonadAsk0, orig.local)
				}())
}

func Call_Control_Comonad_Env_Class_local(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_comonadAskEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): comonadEnvT_1_0 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e), (TypeVar w)])])
comonadEnvT_1_0 := Rebox_Control_Comonad_Env_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Comonad_Env_Trans_comonadEnvT(), dictComonad_0)))
_ = comonadEnvT_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037((&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_3056445460_2550391993(comonadEnvT_1_0))}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
})})))}
}

func Call_Control_Comonad_Env_Class_comonadEnvEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): comonadEnvT_1_1 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e), (TypeVar w)])])
comonadEnvT_1_1 := Rebox_Control_Comonad_Env_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Comonad_Env_Trans_comonadEnvT(), dictComonad_0)))
_ = comonadEnvT_1_1
// TAST (Let): comonadAskEnvT1_1_0 shape=Let(LitRecord) bindingType=(ADT ["Control","Comonad","Env","Class","ComonadAsk"] [(TypeVar e), (TypeApp (ADT ["Data","Tuple","Tuple"] [(TypeVar e), (TypeApp (TypeVar w) [(TypeVar a)])]) [(TypeVar e), (TypeVar w)])])
comonadAskEnvT1_1_0 := (&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_3056445460_2550391993(comonadEnvT_1_1))}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
})})
_ = comonadAskEnvT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_144036252_155172497((&Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037(comonadAskEnvT1_1_0))}
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply(f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
})})))}
}

func Call_Control_Comonad_Env_Class_ask(dict_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}

func Call_Control_Comonad_Env_Class_asks(dictComonadAsk_0_loop *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] = dictComonadAsk_0_loop
_ = dictComonadAsk_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(gopurs_runtime.Box(dictComonadAsk_0.V1), x_2))
}

func Call_Control_Comonad_Env_Class_comonadAskStoreT(dictComonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
_ = dictComonadAsk_0
// TAST (Let): Comonad0_1_0 shape=App(Other) bindingType=Any
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
// TAST (Let): comonadStoreT_2_1 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w) [(Func [(TypeVar s)] (TypeVar a))]), (TypeVar s)])])
comonadStoreT_2_1 := Rebox_Control_Comonad_Env_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Comonad_Store_Trans_comonadStoreT(), Comonad0_1_0)))
_ = comonadStoreT_2_1
// TAST (Let): Functor0_3_3 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar w)])
Functor0_3_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_3
// TAST (Let): __local_var_3_2 shape=Let(Abs(Let(App(Other)))) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w) [(Func [(TypeVar s)] (TypeVar a))]), (TypeVar s)])] (TypeApp (TypeVar w) [(TypeVar a)]))
__local_var_3_2 := gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 shape=Other bindingType=Any
__local_var_5_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_5_4
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_3_3.V0), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_6, __local_var_5_4)
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
})
_ = __local_var_3_2
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037((&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_3056445460_2550391993(comonadStoreT_2_1))}
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "ask"), gopurs_runtime.Apply(__local_var_3_2, x_4))
})})))}
}

func Call_Control_Comonad_Env_Class_comonadEnvStoreT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
_ = dictComonadEnv_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Comonad0_2_2 shape=App(Other) bindingType=Any
Comonad0_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_2_2
// TAST (Let): comonadStoreT_3_3 shape=App(Var) bindingType=(ADT ["Control","Comonad","Comonad"] [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w) [(Func [(TypeVar s)] (TypeVar a))]), (TypeVar s)])])
comonadStoreT_3_3 := Rebox_Control_Comonad_Env_Class_2550391993_3056445460(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Comonad_Store_Trans_comonadStoreT(), Comonad0_2_2)))
_ = comonadStoreT_3_3
// TAST (Let): Functor0_4_5 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar w)])
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_2_2, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
// TAST (Let): __local_var_4_4 shape=Let(Abs(Let(App(Other)))) bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w) [(Func [(TypeVar s)] (TypeVar a))]), (TypeVar s)])] (TypeApp (TypeVar w) [(TypeVar a)]))
__local_var_4_4 := gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 shape=Other bindingType=Any
__local_var_6_6 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1
_ = __local_var_6_6
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v1_7, __local_var_6_6)
}), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)
})
_ = __local_var_4_4
// TAST (Let): comonadAskStoreT1_1_0 shape=Let(Let(Let(LitRecord))) bindingType=(ADT ["Control","Comonad","Env","Class","ComonadAsk"] [(TypeVar e), (ADT ["Data","Tuple","Tuple"] [(TypeApp (TypeVar w) [(Func [(TypeVar s)] (TypeVar a))]), (TypeVar s)])])
comonadAskStoreT1_1_0 := (&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_3056445460_2550391993(comonadStoreT_3_3))}
}), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "ask"), gopurs_runtime.Apply(__local_var_4_4, x_5))
})})
_ = comonadAskStoreT1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_144036252_155172497((&Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(Rebox_Control_Comonad_Env_Class_1649752824_2829962037(comonadAskStoreT1_1_0))}
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_2, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
})})))}
}

func Call_Control_Comonad_Env_Class_comonadAskTracedT(dictComonadAsk_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadAsk_0 gopurs_runtime.Value = dictComonadAsk_0_loop
_ = dictComonadAsk_0
// TAST (Let): Comonad0_1_0 shape=App(Other) bindingType=Any
Comonad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_1_0
// TAST (Let): comonadTracedT__193435443_2_1 shape=App(Var) bindingType=Any
comonadTracedT__193435443_2_1 := gopurs_runtime.Apply(Get_Control_Comonad_Traced_Trans_comonadTracedT(), Comonad0_1_0)
_ = comonadTracedT__193435443_2_1
return gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): comonadTracedT1_4_2 shape=App(Other) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])])
comonadTracedT1_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(comonadTracedT__193435443_2_1, dictMonoid_3))
_ = comonadTracedT1_4_2
// TAST (Let): Functor0_5_4 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar w)])
Functor0_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_1_0, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_4
// TAST (Let): __local_var_5_3 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])] (TypeApp (TypeVar w) [(TypeVar a)]))
__local_var_5_3 := gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_4.V0), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.RecordGet(dictMonoid_3, "mempty"))
}), v_6)
})
_ = __local_var_5_3
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadTracedT1_4_2)}
}), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadAsk_0, "ask"), gopurs_runtime.Apply(__local_var_5_3, x_6))
})}))}
})
}

func Call_Control_Comonad_Env_Class_comonadEnvTracedT(dictComonadEnv_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonadEnv_0 gopurs_runtime.Value = dictComonadEnv_0_loop
_ = dictComonadEnv_0
// TAST (Let): __local_var_1_1 shape=App(Other) bindingType=Any
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonadEnv_0, "ComonadAsk0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Comonad0_2_2 shape=App(Other) bindingType=Any
Comonad0_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Comonad0"), gopurs_runtime.Value{})
_ = Comonad0_2_2
// TAST (Let): comonadTracedT__193435443_3_3 shape=App(Var) bindingType=Any
comonadTracedT__193435443_3_3 := gopurs_runtime.Apply(Get_Control_Comonad_Traced_Trans_comonadTracedT(), Comonad0_2_2)
_ = comonadTracedT__193435443_3_3
// TAST (Let): comonadAskTracedT1__193435443_1_0 shape=Let(Let(Let(Abs(Let(LitRecord))))) bindingType=Any
comonadAskTracedT1__193435443_1_0 := gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): comonadTracedT1_5_4 shape=App(Other) bindingType=(ADT ["Control","Comonad","Comonad"] [(TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])])
comonadTracedT1_5_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]](gopurs_runtime.Apply(comonadTracedT__193435443_3_3, dictMonoid_4))
_ = comonadTracedT1_5_4
// TAST (Let): Functor0_6_6 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar w)])
Functor0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Comonad0_2_2, "Extend0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_6
// TAST (Let): __local_var_6_5 shape=Let(Abs(App(Other))) bindingType=(Func [(TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])] (TypeApp (TypeVar w) [(TypeVar a)]))
__local_var_6_5 := gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_6_6.V0), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_8, gopurs_runtime.RecordGet(dictMonoid_4, "mempty"))
}), v_7)
})
_ = __local_var_6_5
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer(comonadTracedT1_5_4)}
}), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "ask"), gopurs_runtime.Apply(__local_var_6_5, x_7))
})}))}
})
_ = comonadAskTracedT1__193435443_1_0
return gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): comonadAskTracedT2_3_7 shape=App(Other) bindingType=(ADT ["Control","Comonad","Env","Class","ComonadAsk"] [(TypeVar e), (TypeApp (TypeVar w) [(Func [(TypeVar t)] (TypeVar a))])])
comonadAskTracedT2_3_7 := gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(comonadAskTracedT1__193435443_1_0, dictMonoid_2))
_ = comonadAskTracedT2_3_7
return gopurs_runtime.Value{Type: 9, IntVal: 3863290147, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1424092807, UnsafePtr: unsafe.Pointer(comonadAskTracedT2_3_7)}
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictComonadEnv_0, "local"), f_4, v_5)
})}))}
})
}

func Rebox_Control_Comonad_Env_Class_144036252_155172497(in *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Comonad_Env_Class_ComonadEnv[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Comonad_Env_Class_1649752824_2829962037(in *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Comonad_Env_Class_2550391993_3056445460(in *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]) *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Comonad_Env_Class_2829962037_1649752824(in *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Control_Comonad_Env_Class_ComonadAsk[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Control_Comonad_Env_Class_3056445460_2550391993(in *Constructor_Control_Comonad_Comonad[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Control_Comonad_Comonad[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}



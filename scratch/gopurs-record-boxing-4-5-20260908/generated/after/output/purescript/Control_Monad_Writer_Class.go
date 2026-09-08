package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Writer_Class_MonadTell_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Writer_Class_MonadTell_dollar_Dict sync.Once
func Get_Control_Monad_Writer_Class_MonadTell_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_MonadTell_dollar_Dict.Do(func() {
		cache_Control_Monad_Writer_Class_MonadTell_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 551781469, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_Class_MonadTell_dollar_Dict(func() struct{
	Monad1 gopurs_runtime.Value
	Semigroup0 gopurs_runtime.Value
	tell gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Monad1 gopurs_runtime.Value
	Semigroup0 gopurs_runtime.Value
	tell gopurs_runtime.Value
}{}
					clone.Monad1 = gopurs_runtime.RecordGet(orig, "Monad1")
					clone.Semigroup0 = gopurs_runtime.RecordGet(orig, "Semigroup0")
					clone.tell = gopurs_runtime.RecordGet(orig, "tell")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Writer_Class_MonadTell_dollar_Dict
}

var cache_Control_Monad_Writer_Class_MonadWriter_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Writer_Class_MonadWriter_dollar_Dict sync.Once
func Get_Control_Monad_Writer_Class_MonadWriter_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_MonadWriter_dollar_Dict.Do(func() {
		cache_Control_Monad_Writer_Class_MonadWriter_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 784743459, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Writer_Class_MonadWriter_dollar_Dict(func() struct{
	MonadTell1 gopurs_runtime.Value
	Monoid0 gopurs_runtime.Value
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	MonadTell1 gopurs_runtime.Value
	Monoid0 gopurs_runtime.Value
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
}{}
					clone.MonadTell1 = gopurs_runtime.RecordGet(orig, "MonadTell1")
					clone.Monoid0 = gopurs_runtime.RecordGet(orig, "Monoid0")
					clone.listen = gopurs_runtime.RecordGet(orig, "listen")
					clone.pass = gopurs_runtime.RecordGet(orig, "pass")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Writer_Class_MonadWriter_dollar_Dict
}

var cache_Control_Monad_Writer_Class_tell gopurs_runtime.Value
var once_Control_Monad_Writer_Class_tell sync.Once
func Get_Control_Monad_Writer_Class_tell() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_tell.Do(func() {
		cache_Control_Monad_Writer_Class_tell = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_tell(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_tell
}

var cache_Control_Monad_Writer_Class_pass gopurs_runtime.Value
var once_Control_Monad_Writer_Class_pass sync.Once
func Get_Control_Monad_Writer_Class_pass() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_pass.Do(func() {
		cache_Control_Monad_Writer_Class_pass = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_pass(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_pass
}

var cache_Control_Monad_Writer_Class_listen gopurs_runtime.Value
var once_Control_Monad_Writer_Class_listen sync.Once
func Get_Control_Monad_Writer_Class_listen() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_listen.Do(func() {
		cache_Control_Monad_Writer_Class_listen = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_listen(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_listen
}

var cache_Control_Monad_Writer_Class_listens gopurs_runtime.Value
var once_Control_Monad_Writer_Class_listens sync.Once
func Get_Control_Monad_Writer_Class_listens() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_listens.Do(func() {
		cache_Control_Monad_Writer_Class_listens = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_listens(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadWriter_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_listens
}

var cache_Control_Monad_Writer_Class_censor gopurs_runtime.Value
var once_Control_Monad_Writer_Class_censor sync.Once
func Get_Control_Monad_Writer_Class_censor() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_censor.Do(func() {
		cache_Control_Monad_Writer_Class_censor = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_censor(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadWriter_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_censor
}

type Constructor_Control_Monad_Writer_Class_MonadTell[T_w any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[551781469] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Writer_Class_MonadTell[any, any])(ptr)
		_ = c
		switch key {
		case "Monad1": return gopurs_runtime.Box(c.V0)
		case "Semigroup0": return gopurs_runtime.Box(c.V1)
		case "tell": return gopurs_runtime.Box(c.V2)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Writer_Class_MonadTell: " + key)
		}
	}
}


type Constructor_Control_Monad_Writer_Class_MonadWriter[T_w any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[784743459] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Writer_Class_MonadWriter[any, any])(ptr)
		_ = c
		switch key {
		case "MonadTell1": return gopurs_runtime.Box(c.V0)
		case "Monoid0": return gopurs_runtime.Box(c.V1)
		case "listen": return gopurs_runtime.Box(c.V2)
		case "pass": return gopurs_runtime.Box(c.V3)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Writer_Class_MonadWriter: " + key)
		}
	}
}


func Call_Control_Monad_Writer_Class_MonadTell_dollar_Dict(x_0_loop struct{
	Monad1 gopurs_runtime.Value
	Semigroup0 gopurs_runtime.Value
	tell gopurs_runtime.Value
}) *Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	Monad1 gopurs_runtime.Value
	Semigroup0 gopurs_runtime.Value
	tell gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict3("Monad1", "Semigroup0", "tell", orig.Monad1, orig.Semigroup0, orig.tell)
				}())
}

func Call_Control_Monad_Writer_Class_MonadWriter_dollar_Dict(x_0_loop struct{
	MonadTell1 gopurs_runtime.Value
	Monoid0 gopurs_runtime.Value
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
}) *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	MonadTell1 gopurs_runtime.Value
	Monoid0 gopurs_runtime.Value
	listen gopurs_runtime.Value
	pass gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict4("MonadTell1", "Monoid0", "listen", "pass", orig.MonadTell1, orig.Monoid0, orig.listen, orig.pass)
				}())
}

func Call_Control_Monad_Writer_Class_tell(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Monad_Writer_Class_pass(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Monad_Writer_Class_listen(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Monad_Writer_Class_listens(dictMonadWriter_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadWriter_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadWriter_0_loop
_ = dictMonadWriter_0
// TAST (Let): Monad1_1_0 shape=App(Other) bindingType=Any
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadWriter_0.V0), gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])]))
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadWriter_0.V2), m_5), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
}))
})
}

func Call_Control_Monad_Writer_Class_censor(dictMonadWriter_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadWriter_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadWriter_0_loop
_ = dictMonadWriter_0
// TAST (Let): Monad1_1_0 shape=App(Other) bindingType=Any
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadWriter_0.V0), gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 shape=Other bindingType=(Func [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (Func [(TypeVar w)] (TypeVar w))])] (TypeApp (TypeVar m) [(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (Func [(TypeVar w)] (TypeVar w))])]))
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadWriter_0.V3), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_5, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{a_6, f_4}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))})
})))
})
}



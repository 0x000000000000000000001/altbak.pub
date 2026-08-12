package Control_Monad_Writer_Class

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_tell gopurs_runtime.Value
var once_tell sync.Once
func Get_tell() gopurs_runtime.Value {
	once_tell.Do(func() {
		cache_tell = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tell(gopurs_runtime.CoerceToStruct[Constructor_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_tell
}

var cache_pass gopurs_runtime.Value
var once_pass sync.Once
func Get_pass() gopurs_runtime.Value {
	once_pass.Do(func() {
		cache_pass = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pass(gopurs_runtime.CoerceToStruct[Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pass
}

var cache_pass__gopurs_runtime_Value_1553691451 gopurs_runtime.Value
var once_pass__gopurs_runtime_Value_1553691451 sync.Once
func Get_pass__gopurs_runtime_Value_1553691451() gopurs_runtime.Value {
	once_pass__gopurs_runtime_Value_1553691451.Do(func() {
		cache_pass__gopurs_runtime_Value_1553691451 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pass__gopurs_runtime_Value_1553691451(gopurs_runtime.CoerceToStruct[Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pass__gopurs_runtime_Value_1553691451
}

var cache_listen gopurs_runtime.Value
var once_listen sync.Once
func Get_listen() gopurs_runtime.Value {
	once_listen.Do(func() {
		cache_listen = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listen(gopurs_runtime.CoerceToStruct[Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_listen
}

var cache_listen__gopurs_runtime_Value_1604579315 gopurs_runtime.Value
var once_listen__gopurs_runtime_Value_1604579315 sync.Once
func Get_listen__gopurs_runtime_Value_1604579315() gopurs_runtime.Value {
	once_listen__gopurs_runtime_Value_1604579315.Do(func() {
		cache_listen__gopurs_runtime_Value_1604579315 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listen__gopurs_runtime_Value_1604579315(gopurs_runtime.CoerceToStruct[Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_listen__gopurs_runtime_Value_1604579315
}

var cache_listens gopurs_runtime.Value
var once_listens sync.Once
func Get_listens() gopurs_runtime.Value {
	once_listens.Do(func() {
		cache_listens = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listens(gopurs_runtime.CoerceToStruct[Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadWriter_0_box))
})
	})
	return cache_listens
}

var cache_censor gopurs_runtime.Value
var once_censor sync.Once
func Get_censor() gopurs_runtime.Value {
	once_censor.Do(func() {
		cache_censor = gopurs_runtime.Func(func(dictMonadWriter_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_censor(gopurs_runtime.CoerceToStruct[Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dictMonadWriter_0_box))
})
	})
	return cache_censor
}

type Constructor_MonadTell[T_w any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[551781469] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad1": return c.V0
		case "Semigroup0": return c.V1
		case "tell": return c.V2
		default: panic("Key not found in dictionary Constructor_MonadTell: " + key)
		}
	}
}


type Constructor_MonadWriter[T_w any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[784743459] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "MonadTell1": return c.V0
		case "Monoid0": return c.V1
		case "listen": return c.V2
		case "pass": return c.V3
		default: panic("Key not found in dictionary Constructor_MonadWriter: " + key)
		}
	}
}


func Call_tell(dict_0_loop *Constructor_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_pass(dict_0_loop *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_pass__gopurs_runtime_Value_1553691451(dict_0_loop *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_listen(dict_0_loop *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_listen__gopurs_runtime_Value_1604579315(dict_0_loop *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_listens(dictMonadWriter_0_loop *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadWriter_0 *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadWriter_0_loop
_ = dictMonadWriter_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadWriter_0.V0, gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(dictMonadWriter_0.V2, m_5), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)})})
}))
})
})
}

func Call_censor(dictMonadWriter_0_loop *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadWriter_0 *Constructor_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadWriter_0_loop
_ = dictMonadWriter_0
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictMonadWriter_0.V0, gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(dictMonadWriter_0.V3, gopurs_runtime.Apply2(Bind1_2_1.V1, m_5, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, f_4})})
})))
})
})
}



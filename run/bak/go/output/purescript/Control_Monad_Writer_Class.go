package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Writer_Class_MonadTell_dollarDict gopurs_runtime.Value
var once_Control_Monad_Writer_Class_MonadTell_dollarDict sync.Once
func Get_Control_Monad_Writer_Class_MonadTell_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_MonadTell_dollarDict.Do(func() {
		cache_Control_Monad_Writer_Class_MonadTell_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_MonadTell_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Writer_Class_MonadTell_dollarDict
}

var cache_Control_Monad_Writer_Class_MonadWriter_dollarDict gopurs_runtime.Value
var once_Control_Monad_Writer_Class_MonadWriter_dollarDict sync.Once
func Get_Control_Monad_Writer_Class_MonadWriter_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_MonadWriter_dollarDict.Do(func() {
		cache_Control_Monad_Writer_Class_MonadWriter_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_MonadWriter_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Writer_Class_MonadWriter_dollarDict
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

var cache_Control_Monad_Writer_Class_listen__1604579315 gopurs_runtime.Value
var once_Control_Monad_Writer_Class_listen__1604579315 sync.Once
func Get_Control_Monad_Writer_Class_listen__1604579315() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_listen__1604579315.Do(func() {
		cache_Control_Monad_Writer_Class_listen__1604579315 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_listen__1604579315(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_listen__1604579315
}

var cache_Control_Monad_Writer_Class_listen__1880450835 gopurs_runtime.Value
var once_Control_Monad_Writer_Class_listen__1880450835 sync.Once
func Get_Control_Monad_Writer_Class_listen__1880450835() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_listen__1880450835.Do(func() {
		cache_Control_Monad_Writer_Class_listen__1880450835 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_listen__1880450835(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_listen__1880450835
}

var cache_Control_Monad_Writer_Class_listen__3817169875 gopurs_runtime.Value
var once_Control_Monad_Writer_Class_listen__3817169875 sync.Once
func Get_Control_Monad_Writer_Class_listen__3817169875() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_listen__3817169875.Do(func() {
		cache_Control_Monad_Writer_Class_listen__3817169875 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_listen__3817169875(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_listen__3817169875
}

var cache_Control_Monad_Writer_Class_listen__2047803155 gopurs_runtime.Value
var once_Control_Monad_Writer_Class_listen__2047803155 sync.Once
func Get_Control_Monad_Writer_Class_listen__2047803155() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_listen__2047803155.Do(func() {
		cache_Control_Monad_Writer_Class_listen__2047803155 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_listen__2047803155(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_listen__2047803155
}

var cache_Control_Monad_Writer_Class_pass__1553691451 gopurs_runtime.Value
var once_Control_Monad_Writer_Class_pass__1553691451 sync.Once
func Get_Control_Monad_Writer_Class_pass__1553691451() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_pass__1553691451.Do(func() {
		cache_Control_Monad_Writer_Class_pass__1553691451 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_pass__1553691451(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_pass__1553691451
}

var cache_Control_Monad_Writer_Class_pass__1787406491 gopurs_runtime.Value
var once_Control_Monad_Writer_Class_pass__1787406491 sync.Once
func Get_Control_Monad_Writer_Class_pass__1787406491() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_pass__1787406491.Do(func() {
		cache_Control_Monad_Writer_Class_pass__1787406491 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_pass__1787406491(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_pass__1787406491
}

var cache_Control_Monad_Writer_Class_pass__2416360603 gopurs_runtime.Value
var once_Control_Monad_Writer_Class_pass__2416360603 sync.Once
func Get_Control_Monad_Writer_Class_pass__2416360603() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_pass__2416360603.Do(func() {
		cache_Control_Monad_Writer_Class_pass__2416360603 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_pass__2416360603(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_pass__2416360603
}

var cache_Control_Monad_Writer_Class_pass__261986203 gopurs_runtime.Value
var once_Control_Monad_Writer_Class_pass__261986203 sync.Once
func Get_Control_Monad_Writer_Class_pass__261986203() gopurs_runtime.Value {
	once_Control_Monad_Writer_Class_pass__261986203.Do(func() {
		cache_Control_Monad_Writer_Class_pass__261986203 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Writer_Class_pass__261986203(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Writer_Class_pass__261986203
}

type Constructor_Control_Monad_Writer_Class_MonadTell[T_w any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[551781469] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Writer_Class_MonadTell[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
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
		c := (*Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
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


func Call_Control_Monad_Writer_Class_MonadTell_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Writer_Class_MonadWriter_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
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
// TAST (Let): Monad1_1_0 -> gopurs_runtime.Value
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadWriter_0.V0), gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadWriter_0.V2), m_5), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, gopurs_runtime.Apply(f_4, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)})})
}))
})
})
}

func Call_Control_Monad_Writer_Class_censor(dictMonadWriter_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadWriter_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dictMonadWriter_0_loop
_ = dictMonadWriter_0
// TAST (Let): Monad1_1_0 -> gopurs_runtime.Value
Monad1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadWriter_0.V0), gopurs_runtime.Value{}), "Monad1"), gopurs_runtime.Value{})
_ = Monad1_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind[gopurs_runtime.Value]
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
// TAST (Let): pure_3_2 -> gopurs_runtime.Value
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad1_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadWriter_0.V3), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), m_5, gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_3_2, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_6, f_4})})
})))
})
})
}

func Call_Control_Monad_Writer_Class_listen__1604579315(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Monad_Writer_Class_listen__1880450835(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Monad_Writer_Class_listen__3817169875(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Monad_Writer_Class_listen__2047803155(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V2)
}

func Call_Control_Monad_Writer_Class_pass__1553691451(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Monad_Writer_Class_pass__1787406491(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Monad_Writer_Class_pass__2416360603(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}

func Call_Control_Monad_Writer_Class_pass__261986203(dict_0_loop *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Writer_Class_MonadWriter[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V3)
}



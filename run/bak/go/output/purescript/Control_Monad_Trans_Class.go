package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Trans_Class_MonadTrans_dollarDict gopurs_runtime.Value
var once_Control_Monad_Trans_Class_MonadTrans_dollarDict sync.Once
func Get_Control_Monad_Trans_Class_MonadTrans_dollarDict() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_MonadTrans_dollarDict.Do(func() {
		cache_Control_Monad_Trans_Class_MonadTrans_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Trans_Class_MonadTrans_dollarDict(x_0_box)
})
	})
	return cache_Control_Monad_Trans_Class_MonadTrans_dollarDict
}

var cache_Control_Monad_Trans_Class_lift gopurs_runtime.Value
var once_Control_Monad_Trans_Class_lift sync.Once
func Get_Control_Monad_Trans_Class_lift() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_lift.Do(func() {
		cache_Control_Monad_Trans_Class_lift = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans](dict_0_box))
})
	})
	return cache_Control_Monad_Trans_Class_lift
}

var cache_Control_Monad_Trans_Class_lift__3816229929 gopurs_runtime.Value
var once_Control_Monad_Trans_Class_lift__3816229929 sync.Once
func Get_Control_Monad_Trans_Class_lift__3816229929() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_lift__3816229929.Do(func() {
		cache_Control_Monad_Trans_Class_lift__3816229929 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Trans_Class_lift__3816229929(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans](dict_0_box))
})
	})
	return cache_Control_Monad_Trans_Class_lift__3816229929
}

var cache_Control_Monad_Trans_Class_lift__487910375 gopurs_runtime.Value
var once_Control_Monad_Trans_Class_lift__487910375 sync.Once
func Get_Control_Monad_Trans_Class_lift__487910375() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_lift__487910375.Do(func() {
		cache_Control_Monad_Trans_Class_lift__487910375 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Trans_Class_lift__487910375(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Trans_Class_lift__487910375
}

var cache_Control_Monad_Trans_Class_lift__1866483406 gopurs_runtime.Value
var once_Control_Monad_Trans_Class_lift__1866483406 sync.Once
func Get_Control_Monad_Trans_Class_lift__1866483406() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_lift__1866483406.Do(func() {
		cache_Control_Monad_Trans_Class_lift__1866483406 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Trans_Class_lift__1866483406(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](dictMonad_0_box)))}
})
	})
	return cache_Control_Monad_Trans_Class_lift__1866483406
}

var cache_Control_Monad_Trans_Class_lift__115114023 gopurs_runtime.Value
var once_Control_Monad_Trans_Class_lift__115114023 sync.Once
func Get_Control_Monad_Trans_Class_lift__115114023() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_lift__115114023.Do(func() {
		cache_Control_Monad_Trans_Class_lift__115114023 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Trans_Class_lift__115114023(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans](dict_0_box))
})
	})
	return cache_Control_Monad_Trans_Class_lift__115114023
}

var cache_Control_Monad_Trans_Class_lift__1331755881 gopurs_runtime.Value
var once_Control_Monad_Trans_Class_lift__1331755881 sync.Once
func Get_Control_Monad_Trans_Class_lift__1331755881() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_lift__1331755881.Do(func() {
		cache_Control_Monad_Trans_Class_lift__1331755881 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Trans_Class_lift__1331755881(dictMonad_0_box)
})
	})
	return cache_Control_Monad_Trans_Class_lift__1331755881
}

type Constructor_Control_Monad_Trans_Class_MonadTrans struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2835982595] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Trans_Class_MonadTrans)(ptr)
		_ = c
		switch key {
		case "lift": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Trans_Class_MonadTrans: " + key)
		}
	}
}


func Call_Control_Monad_Trans_Class_MonadTrans_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Monad_Trans_Class_lift(dict_0_loop *Constructor_Control_Monad_Trans_Class_MonadTrans) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Trans_Class_MonadTrans = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Monad_Trans_Class_lift__3816229929(dict_0_loop *Constructor_Control_Monad_Trans_Class_MonadTrans) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Trans_Class_MonadTrans = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Monad_Trans_Class_lift__487910375(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_3, gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, a_4})})
}))
})
}

func Call_Control_Monad_Trans_Class_lift__1866483406(dictMonad_0_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Maybe_Just {
var dictMonad_0 *Constructor_Data_Maybe_Just = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_1 -> *Constructor_Control_Bind_Bind
Bind1_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_1
// TAST (Let): Applicative0_2_2 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(dictMonad_0)}, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_2
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_1.V1), a_3, gopurs_runtime.Func(func(a_prime_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, a_prime_4})})
}))
})
_ = __local_var_1_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Control_Monad_Maybe_Trans_MaybeT(), gopurs_runtime.Apply(__local_var_1_0, x_2))
}))
}

func Call_Control_Monad_Trans_Class_lift__115114023(dict_0_loop *Constructor_Control_Monad_Trans_Class_MonadTrans) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Trans_Class_MonadTrans = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Control_Monad_Trans_Class_lift__1331755881(dictMonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonad_0 gopurs_runtime.Value = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Bind1_1_0 -> *Constructor_Control_Bind_Bind
Bind1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonad_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_1_0
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(k_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_1_0.V1), m_2, k_3)
})
})
}



package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Trans_Class_ComonadTrans_dollarDict gopurs_runtime.Value
var once_Control_Comonad_Trans_Class_ComonadTrans_dollarDict sync.Once
func Get_Control_Comonad_Trans_Class_ComonadTrans_dollarDict() gopurs_runtime.Value {
	once_Control_Comonad_Trans_Class_ComonadTrans_dollarDict.Do(func() {
		cache_Control_Comonad_Trans_Class_ComonadTrans_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Trans_Class_ComonadTrans_dollarDict(x_0_box)
})
	})
	return cache_Control_Comonad_Trans_Class_ComonadTrans_dollarDict
}

var cache_Control_Comonad_Trans_Class_lower gopurs_runtime.Value
var once_Control_Comonad_Trans_Class_lower sync.Once
func Get_Control_Comonad_Trans_Class_lower() gopurs_runtime.Value {
	once_Control_Comonad_Trans_Class_lower.Do(func() {
		cache_Control_Comonad_Trans_Class_lower = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Trans_Class_lower(gopurs_runtime.CoerceToStruct[Constructor_Control_Comonad_Trans_Class_ComonadTrans](dict_0_box))
})
	})
	return cache_Control_Comonad_Trans_Class_lower
}

var cache_Control_Comonad_Trans_Class_comonadTransIdentityT gopurs_runtime.Value
var once_Control_Comonad_Trans_Class_comonadTransIdentityT sync.Once
func Get_Control_Comonad_Trans_Class_comonadTransIdentityT() gopurs_runtime.Value {
	once_Control_Comonad_Trans_Class_comonadTransIdentityT.Do(func() {
		cache_Control_Comonad_Trans_Class_comonadTransIdentityT = gopurs_runtime.Value{Type: 9, IntVal: 3399197123, UnsafePtr: unsafe.Pointer(&Constructor_Control_Comonad_Trans_Class_ComonadTrans{1, gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Control_Monad_Identity_Trans_runIdentityT()
})})}
	})
	return cache_Control_Comonad_Trans_Class_comonadTransIdentityT
}

type Constructor_Control_Comonad_Trans_Class_ComonadTrans struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3399197123] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Comonad_Trans_Class_ComonadTrans)(ptr)
		_ = c
		switch key {
		case "lower": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Control_Comonad_Trans_Class_ComonadTrans: " + key)
		}
	}
}


func Call_Control_Comonad_Trans_Class_ComonadTrans_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Trans_Class_lower(dict_0_loop *Constructor_Control_Comonad_Trans_Class_ComonadTrans) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Comonad_Trans_Class_ComonadTrans = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}



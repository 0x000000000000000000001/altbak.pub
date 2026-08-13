package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_MonadPlus_MonadPlus_dollarDict gopurs_runtime.Value
var once_Control_MonadPlus_MonadPlus_dollarDict sync.Once
func Get_Control_MonadPlus_MonadPlus_dollarDict() gopurs_runtime.Value {
	once_Control_MonadPlus_MonadPlus_dollarDict.Do(func() {
		cache_Control_MonadPlus_MonadPlus_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_MonadPlus_MonadPlus_dollarDict(x_0_box)
})
	})
	return cache_Control_MonadPlus_MonadPlus_dollarDict
}

var cache_Control_MonadPlus_monadPlusArray gopurs_runtime.Value
var once_Control_MonadPlus_monadPlusArray sync.Once
func Get_Control_MonadPlus_monadPlusArray() gopurs_runtime.Value {
	once_Control_MonadPlus_monadPlusArray.Do(func() {
		cache_Control_MonadPlus_monadPlusArray = gopurs_runtime.Value{Type: 9, IntVal: 3236234573, UnsafePtr: unsafe.Pointer(&Constructor_Control_MonadPlus_MonadPlus{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](Get_Control_Alternative_alternativeArray()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Control_Monad_monadArray()))}
})})}
	})
	return cache_Control_MonadPlus_monadPlusArray
}

type Constructor_Control_MonadPlus_MonadPlus struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3236234573] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_MonadPlus_MonadPlus)(ptr)
		_ = c
		switch key {
		case "Alternative1": return gopurs_runtime.Box(c.V0)
		case "Monad0": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_MonadPlus_MonadPlus: " + key)
		}
	}
}


func Call_Control_MonadPlus_MonadPlus_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}



package Control_Monad_Trans_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_lift gopurs_runtime.Value
var once_lift sync.Once
func Get_lift() gopurs_runtime.Value {
	once_lift.Do(func() {
		cache_lift = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift(gopurs_runtime.CoerceToStruct[Constructor_MonadTrans[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_lift
}

var cache_lift__gopurs_runtime_Value_3816229929 gopurs_runtime.Value
var once_lift__gopurs_runtime_Value_3816229929 sync.Once
func Get_lift__gopurs_runtime_Value_3816229929() gopurs_runtime.Value {
	once_lift__gopurs_runtime_Value_3816229929.Do(func() {
		cache_lift__gopurs_runtime_Value_3816229929 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lift__gopurs_runtime_Value_3816229929(gopurs_runtime.CoerceToStruct[Constructor_MonadTrans[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_lift__gopurs_runtime_Value_3816229929
}

type Constructor_MonadTrans[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2835982595] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadTrans[gopurs_runtime.Value])(ptr)
		switch key {
		case "lift": return c.V0
		default: panic("Key not found in dictionary Constructor_MonadTrans: " + key)
		}
	}
}


func Call_lift(dict_0_loop *Constructor_MonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadTrans[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_lift__gopurs_runtime_Value_3816229929(dict_0_loop *Constructor_MonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadTrans[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}



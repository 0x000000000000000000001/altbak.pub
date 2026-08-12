package Control_Monad_ST_Class

import (
	pkg_Control_Monad_ST_Internal "gopurs/output/Control.Monad.ST.Internal"
	pkg_Effect "gopurs/output/Effect"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_monadSTST gopurs_runtime.Value
var once_monadSTST sync.Once
func Get_monadSTST() gopurs_runtime.Value {
	once_monadSTST.Do(func() {
		cache_monadSTST = gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Monad_ST_Internal.Get_monadST()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_monadSTST
}

var cache_monadSTEffect gopurs_runtime.Value
var once_monadSTEffect sync.Once
func Get_monadSTEffect() gopurs_runtime.Value {
	once_monadSTEffect.Do(func() {
		cache_monadSTEffect = gopurs_runtime.RecordDict2("Monad0", "liftST", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Effect.Get_monadEffect()
}), pkg_Unsafe_Coerce.Get_unsafeCoerce())
	})
	return cache_monadSTEffect
}

var cache_liftST gopurs_runtime.Value
var once_liftST sync.Once
func Get_liftST() gopurs_runtime.Value {
	once_liftST.Do(func() {
		cache_liftST = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftST(gopurs_runtime.CoerceToStruct[Constructor_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_liftST
}

type Constructor_MonadST[T_s any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2155655715] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MonadST[gopurs_runtime.Value, gopurs_runtime.Value])(ptr)
		switch key {
		case "Monad0": return c.V0
		case "liftST": return c.V1
		default: panic("Key not found in dictionary Constructor_MonadST: " + key)
		}
	}
}


func Call_liftST(dict_0_loop *Constructor_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MonadST[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



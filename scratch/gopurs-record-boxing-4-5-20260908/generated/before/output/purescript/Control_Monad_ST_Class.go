package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_ST_Class_MonadST_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_ST_Class_MonadST_dollar_Dict sync.Once
func Get_Control_Monad_ST_Class_MonadST_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_ST_Class_MonadST_dollar_Dict.Do(func() {
		cache_Control_Monad_ST_Class_MonadST_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer(Call_Control_Monad_ST_Class_MonadST_dollar_Dict(func() struct{
	Monad0 gopurs_runtime.Value
	liftST gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Monad0 gopurs_runtime.Value
	liftST gopurs_runtime.Value
}{}
					clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
					clone.liftST = gopurs_runtime.RecordGet(orig, "liftST")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_ST_Class_MonadST_dollar_Dict
}

var cache_Control_Monad_ST_Class_monadSTST gopurs_runtime.Value
var once_Control_Monad_ST_Class_monadSTST sync.Once
func Get_Control_Monad_ST_Class_monadSTST() gopurs_runtime.Value {
	once_Control_Monad_ST_Class_monadSTST.Do(func() {
		cache_Control_Monad_ST_Class_monadSTST = gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_monadST()))}
}), gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1)}))}
	})
	return cache_Control_Monad_ST_Class_monadSTST
}

var cache_Control_Monad_ST_Class_monadSTEffect gopurs_runtime.Value
var once_Control_Monad_ST_Class_monadSTEffect sync.Once
func Get_Control_Monad_ST_Class_monadSTEffect() gopurs_runtime.Value {
	once_Control_Monad_ST_Class_monadSTEffect.Do(func() {
		cache_Control_Monad_ST_Class_monadSTEffect = gopurs_runtime.Value{Type: 9, IntVal: 2155655715, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Effect_monadEffect()))}
}), Get_Unsafe_Coerce_unsafeCoerce()}))}
	})
	return cache_Control_Monad_ST_Class_monadSTEffect
}

var cache_Control_Monad_ST_Class_liftST gopurs_runtime.Value
var once_Control_Monad_ST_Class_liftST sync.Once
func Get_Control_Monad_ST_Class_liftST() gopurs_runtime.Value {
	once_Control_Monad_ST_Class_liftST.Do(func() {
		cache_Control_Monad_ST_Class_liftST = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_ST_Class_liftST(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_ST_Class_liftST
}

type Constructor_Control_Monad_ST_Class_MonadST[T_s any, T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2155655715] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_ST_Class_MonadST[any, any])(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "liftST": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_ST_Class_MonadST: " + key)
		}
	}
}


func Call_Control_Monad_ST_Class_MonadST_dollar_Dict(x_0_loop struct{
	Monad0 gopurs_runtime.Value
	liftST gopurs_runtime.Value
}) *Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value] {
var x_0 struct{
	Monad0 gopurs_runtime.Value
	liftST gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Monad0", "liftST"}, []gopurs_runtime.Value{orig.Monad0, orig.liftST})
				}())
}

func Call_Control_Monad_ST_Class_liftST(dict_0_loop *Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_ST_Class_MonadST[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



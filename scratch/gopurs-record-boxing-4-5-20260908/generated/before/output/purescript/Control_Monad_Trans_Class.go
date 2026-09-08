package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Trans_Class_MonadTrans_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Trans_Class_MonadTrans_dollar_Dict sync.Once
func Get_Control_Monad_Trans_Class_MonadTrans_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_MonadTrans_dollar_Dict.Do(func() {
		cache_Control_Monad_Trans_Class_MonadTrans_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2835982595, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Trans_Class_MonadTrans_dollar_Dict(func() struct{
	lift gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	lift gopurs_runtime.Value
}{}
					clone.lift = gopurs_runtime.RecordGet(orig, "lift")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Trans_Class_MonadTrans_dollar_Dict
}

var cache_Control_Monad_Trans_Class_lift gopurs_runtime.Value
var once_Control_Monad_Trans_Class_lift sync.Once
func Get_Control_Monad_Trans_Class_lift() gopurs_runtime.Value {
	once_Control_Monad_Trans_Class_lift.Do(func() {
		cache_Control_Monad_Trans_Class_lift = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Trans_Class_lift(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Trans_Class_lift
}

type Constructor_Control_Monad_Trans_Class_MonadTrans[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2835982595] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Trans_Class_MonadTrans[any])(ptr)
		_ = c
		switch key {
		case "lift": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Trans_Class_MonadTrans: " + key)
		}
	}
}


func Call_Control_Monad_Trans_Class_MonadTrans_dollar_Dict(x_0_loop struct{
	lift gopurs_runtime.Value
}) *Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value] {
var x_0 struct{
	lift gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"lift"}, []gopurs_runtime.Value{orig.lift})
				}())
}

func Call_Control_Monad_Trans_Class_lift(dict_0_loop *Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Trans_Class_MonadTrans[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}



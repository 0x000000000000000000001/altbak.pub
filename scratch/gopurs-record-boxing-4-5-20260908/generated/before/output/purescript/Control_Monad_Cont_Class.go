package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_Cont_Class_MonadCont_dollar_Dict gopurs_runtime.Value
var once_Control_Monad_Cont_Class_MonadCont_dollar_Dict sync.Once
func Get_Control_Monad_Cont_Class_MonadCont_dollar_Dict() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_MonadCont_dollar_Dict.Do(func() {
		cache_Control_Monad_Cont_Class_MonadCont_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1800060259, UnsafePtr: unsafe.Pointer(Call_Control_Monad_Cont_Class_MonadCont_dollar_Dict(func() struct{
	Monad0 gopurs_runtime.Value
	callCC gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	Monad0 gopurs_runtime.Value
	callCC gopurs_runtime.Value
}{}
					clone.Monad0 = gopurs_runtime.RecordGet(orig, "Monad0")
					clone.callCC = gopurs_runtime.RecordGet(orig, "callCC")
					return clone
				}()))}
})
	})
	return cache_Control_Monad_Cont_Class_MonadCont_dollar_Dict
}

var cache_Control_Monad_Cont_Class_callCC gopurs_runtime.Value
var once_Control_Monad_Cont_Class_callCC sync.Once
func Get_Control_Monad_Cont_Class_callCC() gopurs_runtime.Value {
	once_Control_Monad_Cont_Class_callCC.Do(func() {
		cache_Control_Monad_Cont_Class_callCC = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_Cont_Class_callCC(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Control_Monad_Cont_Class_callCC
}

type Constructor_Control_Monad_Cont_Class_MonadCont[T_m any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1800060259] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Control_Monad_Cont_Class_MonadCont[any])(ptr)
		_ = c
		switch key {
		case "Monad0": return gopurs_runtime.Box(c.V0)
		case "callCC": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Control_Monad_Cont_Class_MonadCont: " + key)
		}
	}
}


func Call_Control_Monad_Cont_Class_MonadCont_dollar_Dict(x_0_loop struct{
	Monad0 gopurs_runtime.Value
	callCC gopurs_runtime.Value
}) *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value] {
var x_0 struct{
	Monad0 gopurs_runtime.Value
	callCC gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict([]string{"Monad0", "callCC"}, []gopurs_runtime.Value{orig.Monad0, orig.callCC})
				}())
}

func Call_Control_Monad_Cont_Class_callCC(dict_0_loop *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Control_Monad_Cont_Class_MonadCont[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V1)
}



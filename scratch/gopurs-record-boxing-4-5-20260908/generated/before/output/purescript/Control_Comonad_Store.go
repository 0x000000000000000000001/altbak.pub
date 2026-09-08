package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Store_unwrap gopurs_runtime.Value
var once_Control_Comonad_Store_unwrap sync.Once
func Get_Control_Comonad_Store_unwrap() gopurs_runtime.Value {
	once_Control_Comonad_Store_unwrap.Do(func() {
		cache_Control_Comonad_Store_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Comonad_Store_unwrap
}

var cache_Control_Comonad_Store_store gopurs_runtime.Value
var once_Control_Comonad_Store_store sync.Once
func Get_Control_Comonad_Store_store() gopurs_runtime.Value {
	once_Control_Comonad_Store_store.Do(func() {
		cache_Control_Comonad_Store_store = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Store_store(f_0_box, x_1_box)
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Comonad_Store_store
}

var cache_Control_Comonad_Store_runStore gopurs_runtime.Value
var once_Control_Comonad_Store_runStore sync.Once
func Get_Control_Comonad_Store_runStore() gopurs_runtime.Value {
	once_Control_Comonad_Store_runStore.Do(func() {
		cache_Control_Comonad_Store_runStore = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Store_runStore(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Comonad_Store_runStore
}

func Call_Control_Comonad_Store_store(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{f_0, x_1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Control_Comonad_Store_runStore(v_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var v_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 shape=Other bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar s), (Func [(TypeVar s)] (TypeVar a))])
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(v_0).V1, (v_0).V0}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
_ = __local_var_1_0
// TAST (Let): __local_var_2_1 shape=Other bindingType=(ADT ["Data","Tuple","Tuple"] [(TypeVar a), (TypeVar b)])
__local_var_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(__local_var_1_0).V0, (__local_var_1_0).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())
_ = __local_var_2_1
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(__local_var_2_1).V1, (__local_var_2_1).V0}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}



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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Store_store(f_0_box, x_1_box))}
})
	})
	return cache_Control_Comonad_Store_store
}

var cache_Control_Comonad_Store_runStore gopurs_runtime.Value
var once_Control_Comonad_Store_runStore sync.Once
func Get_Control_Comonad_Store_runStore() gopurs_runtime.Value {
	once_Control_Comonad_Store_runStore.Do(func() {
		cache_Control_Comonad_Store_runStore = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Store_runStore(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))}
})
	})
	return cache_Control_Comonad_Store_runStore
}

func Call_Control_Comonad_Store_store(f_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, f_0, x_1})})
}

func Call_Control_Comonad_Store_runStore(v_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Tuple_Tuple
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Tuple_functorTuple(), "map"), Get_Unsafe_Coerce_unsafeCoerce(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (v_0).V1, (v_0).V0})}))
_ = __local_var_1_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (__local_var_1_0).V1, (__local_var_1_0).V0})})
}



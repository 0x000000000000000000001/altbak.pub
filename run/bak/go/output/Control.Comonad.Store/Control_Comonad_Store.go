package Control_Comonad_Store

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	unsafe "unsafe"
)

var cache_store gopurs_runtime.Value
var once_store sync.Once
func Get_store() gopurs_runtime.Value {
	once_store.Do(func() {
		cache_store = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_store(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(x_1_box))
})
	})
	return cache_store
}

var cache_runStore gopurs_runtime.Value
var once_runStore sync.Once
func Get_runStore() gopurs_runtime.Value {
	once_runStore.Do(func() {
		cache_runStore = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runStore(v_0_box)
})
	})
	return cache_runStore
}

func Call_store(f_0_loop func(interface{}) interface{}, x_1_loop interface{}) gopurs_runtime.Value {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var x_1 interface{} = x_1_loop
_ = x_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(gopurs_runtime.UnboxAny(arg0)))
})), x_1})})
}

func Call_runStore(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0))})}))
_ = __local_var_1_0
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V1)), gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0))})})
}

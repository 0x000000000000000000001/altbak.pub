package Control_Comonad_Env

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	pkg_Data_Identity "gopurs/output/Data.Identity"
	unsafe "unsafe"
)

var cache_withEnv gopurs_runtime.Value
var once_withEnv sync.Once
func Get_withEnv() gopurs_runtime.Value {
	once_withEnv.Do(func() {
		cache_withEnv = pkg_Control_Comonad_Env_Trans.Get_withEnvT__gopurs_runtime_Value()
	})
	return cache_withEnv
}

var cache_runEnv gopurs_runtime.Value
var once_runEnv sync.Once
func Get_runEnv() gopurs_runtime.Value {
	once_runEnv.Do(func() {
		cache_runEnv = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runEnv(v_0_box)
})
	})
	return cache_runEnv
}

var cache_mapEnv gopurs_runtime.Value
var once_mapEnv sync.Once
func Get_mapEnv() gopurs_runtime.Value {
	once_mapEnv.Do(func() {
		cache_mapEnv = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapEnv(f_0_box, v_1_box)
})
	})
	return cache_mapEnv
}

var cache_env gopurs_runtime.Value
var once_env sync.Once
func Get_env() gopurs_runtime.Value {
	once_env.Do(func() {
		cache_env = gopurs_runtime.Func2(func(e_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_env(e_0_box, a_1_box)
})
	})
	return cache_env
}

func Call_runEnv(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), pkg_Unsafe_Coerce.Get_unsafeCoerce(), v_0)
}

func Call_mapEnv(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{(*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_1.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Identity.Get_functorIdentity(), "map"), f_0, (*pkg_Data_Tuple.Data_Data_Tuple_Tuple)(v_1.UnsafePtr).V1)})}
}

func Call_env(e_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var e_0 gopurs_runtime.Value = e_0_loop
_ = e_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Data_Data_Tuple_Tuple{e_0, a_1})}
}



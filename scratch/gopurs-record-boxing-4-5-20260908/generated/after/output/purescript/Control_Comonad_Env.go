package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Env_unwrap gopurs_runtime.Value
var once_Control_Comonad_Env_unwrap sync.Once
func Get_Control_Comonad_Env_unwrap() gopurs_runtime.Value {
	once_Control_Comonad_Env_unwrap.Do(func() {
		cache_Control_Comonad_Env_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Comonad_Env_unwrap
}

var cache_Control_Comonad_Env_withEnv gopurs_runtime.Value
var once_Control_Comonad_Env_withEnv sync.Once
func Get_Control_Comonad_Env_withEnv() gopurs_runtime.Value {
	once_Control_Comonad_Env_withEnv.Do(func() {
		cache_Control_Comonad_Env_withEnv = Get_Control_Comonad_Env_Trans_withEnvT()
	})
	return cache_Control_Comonad_Env_withEnv
}

var cache_Control_Comonad_Env_runEnv gopurs_runtime.Value
var once_Control_Comonad_Env_runEnv sync.Once
func Get_Control_Comonad_Env_runEnv() gopurs_runtime.Value {
	once_Control_Comonad_Env_runEnv.Do(func() {
		cache_Control_Comonad_Env_runEnv = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Env_runEnv(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Comonad_Env_runEnv
}

var cache_Control_Comonad_Env_mapEnv gopurs_runtime.Value
var once_Control_Comonad_Env_mapEnv sync.Once
func Get_Control_Comonad_Env_mapEnv() gopurs_runtime.Value {
	once_Control_Comonad_Env_mapEnv.Do(func() {
		cache_Control_Comonad_Env_mapEnv = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Env_mapEnv(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Comonad_Env_mapEnv
}

var cache_Control_Comonad_Env_env gopurs_runtime.Value
var once_Control_Comonad_Env_env sync.Once
func Get_Control_Comonad_Env_env() gopurs_runtime.Value {
	once_Control_Comonad_Env_env.Do(func() {
		cache_Control_Comonad_Env_env = gopurs_runtime.Func2(func(e_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Control_Comonad_Env_env(e_0_box, a_1_box)
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Control_Comonad_Env_env
}

func Call_Control_Comonad_Env_runEnv(v_0_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var v_0 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(v_0).V0, (v_0).V1}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Control_Comonad_Env_mapEnv(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{(v_1).V0, gopurs_runtime.Apply(f_0, (v_1).V1)}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Control_Comonad_Env_env(e_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var e_0 gopurs_runtime.Value = e_0_loop
_ = e_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{e_0, a_1}
}



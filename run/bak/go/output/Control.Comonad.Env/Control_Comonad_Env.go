package Control_Comonad_Env

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var withEnv gopurs_runtime.Value
var once_withEnv sync.Once
func Get_withEnv() gopurs_runtime.Value {
	once_withEnv.Do(func() {
		withEnv = pkg_Control_Comonad_Env_Trans.Get_withEnvT()
	})
	return withEnv
}

var runEnv gopurs_runtime.Value
var once_runEnv sync.Once
func Get_runEnv() gopurs_runtime.Value {
	once_runEnv.Do(func() {
		runEnv = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), v_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
	})
	return runEnv
}

var mapEnv gopurs_runtime.Value
var once_mapEnv sync.Once
func Get_mapEnv() gopurs_runtime.Value {
	once_mapEnv.Do(func() {
		mapEnv = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": v_1.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": gopurs_runtime.Apply(f_0, v_1.PtrVal.(map[string]gopurs_runtime.Value)["value1"])})
})
})
	})
	return mapEnv
}

var env gopurs_runtime.Value
var once_env sync.Once
func Get_env() gopurs_runtime.Value {
	once_env.Do(func() {
		env = gopurs_runtime.Func(func(e_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": e_0, "value1": a_1})
})
})
	})
	return env
}



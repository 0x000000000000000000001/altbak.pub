package Control_Comonad_Env

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Comonad_Env_Trans "gopurs/output/Control.Comonad.Env.Trans"
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
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[0], (*[1024]gopurs_runtime.Value)(v_0.UnsafePtr)[1])
})
	})
	return runEnv
}

var mapEnv gopurs_runtime.Value
var once_mapEnv sync.Once
func Get_mapEnv() gopurs_runtime.Value {
	once_mapEnv.Do(func() {
		mapEnv = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[0], gopurs_runtime.Apply(f_0, (*[1024]gopurs_runtime.Value)(v_1.UnsafePtr)[1]))
})
	})
	return mapEnv
}

var env gopurs_runtime.Value
var once_env sync.Once
func Get_env() gopurs_runtime.Value {
	once_env.Do(func() {
		env = gopurs_runtime.Func2(func(e_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Constructor2("Tuple", e_0, a_1)
})
	})
	return env
}





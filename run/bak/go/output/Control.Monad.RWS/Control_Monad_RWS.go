package Control_Monad_RWS

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_RWS_Trans "gopurs/output/Control.Monad.RWS.Trans"
)

var withRWS gopurs_runtime.Value
var once_withRWS sync.Once
func Get_withRWS() gopurs_runtime.Value {
	once_withRWS.Do(func() {
		withRWS = pkg_Control_Monad_RWS_Trans.Get_withRWST()
	})
	return withRWS
}

var rws gopurs_runtime.Value
var once_rws sync.Once
func Get_rws() gopurs_runtime.Value {
	once_rws.Do(func() {
		rws = gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, r_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, r_1, s_2)
})
	})
	return rws
}

var runRWS gopurs_runtime.Value
var once_runRWS sync.Once
func Get_runRWS() gopurs_runtime.Value {
	once_runRWS.Do(func() {
		runRWS = gopurs_runtime.Func3(func(m_0 gopurs_runtime.Value, r_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(m_0, r_1, s_2)
})
	})
	return runRWS
}

var mapRWS gopurs_runtime.Value
var once_mapRWS sync.Once
func Get_mapRWS() gopurs_runtime.Value {
	once_mapRWS.Do(func() {
		mapRWS = gopurs_runtime.Func4(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value, r_2 gopurs_runtime.Value, s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply2(v_1, r_2, s_3))
})
	})
	return mapRWS
}

var execRWS gopurs_runtime.Value
var once_execRWS sync.Once
func Get_execRWS() gopurs_runtime.Value {
	once_execRWS.Do(func() {
		execRWS = gopurs_runtime.Func3(func(m_0 gopurs_runtime.Value, r_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(m_0, r_1, s_2)
_ = __local_var_3_0
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.ConstructorGet(__local_var_3_0, 0), gopurs_runtime.ConstructorGet(__local_var_3_0, 2))
})
	})
	return execRWS
}

var evalRWS gopurs_runtime.Value
var once_evalRWS sync.Once
func Get_evalRWS() gopurs_runtime.Value {
	once_evalRWS.Do(func() {
		evalRWS = gopurs_runtime.Func3(func(m_0 gopurs_runtime.Value, r_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(m_0, r_1, s_2)
_ = __local_var_3_0
return gopurs_runtime.Constructor2("Tuple", gopurs_runtime.ConstructorGet(__local_var_3_0, 1), gopurs_runtime.ConstructorGet(__local_var_3_0, 2))
})
	})
	return evalRWS
}



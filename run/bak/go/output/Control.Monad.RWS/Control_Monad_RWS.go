package Control_Monad_RWS

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_RWS_Trans "gopurs/output/Control.Monad.RWS.Trans"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
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
		rws = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(f_0, r_1), s_2)
})
})
})
	})
	return rws
}

var runRWS gopurs_runtime.Value
var once_runRWS sync.Once
func Get_runRWS() gopurs_runtime.Value {
	once_runRWS.Do(func() {
		runRWS = gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(m_0, r_1), s_2)
})
})
})
	})
	return runRWS
}

var mapRWS gopurs_runtime.Value
var once_mapRWS sync.Once
func Get_mapRWS() gopurs_runtime.Value {
	once_mapRWS.Do(func() {
		mapRWS = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(gopurs_runtime.Apply(v_1, r_2), s_3)))
})
})
})
})
	})
	return mapRWS
}

var execRWS gopurs_runtime.Value
var once_execRWS sync.Once
func Get_execRWS() gopurs_runtime.Value {
	once_execRWS.Do(func() {
		execRWS = gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(m_0, r_1), s_2)
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"], "value1": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]}))
})
})
})
	})
	return execRWS
}

var evalRWS gopurs_runtime.Value
var once_evalRWS sync.Once
func Get_evalRWS() gopurs_runtime.Value {
	once_evalRWS.Do(func() {
		evalRWS = gopurs_runtime.Func(func(m_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(r_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(m_0, r_1), s_2)
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Tuple"), "value0": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value1"], "value1": __local_var_3_0.PtrVal.(map[string]gopurs_runtime.Value)["value2"]}))
})
})
})
	})
	return evalRWS
}



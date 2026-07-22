package Control_Monad_State

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_State_Trans "gopurs/output/Control.Monad.State.Trans"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
)

var withState gopurs_runtime.Value
var once_withState sync.Once
func Get_withState() gopurs_runtime.Value {
	once_withState.Do(func() {
		withState = pkg_Control_Monad_State_Trans.Get_withStateT()
	})
	return withState
}

var runState gopurs_runtime.Value
var once_runState sync.Once
func Get_runState() gopurs_runtime.Value {
	once_runState.Do(func() {
		runState = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(v_0, x_1))
})
})
	})
	return runState
}

var mapState gopurs_runtime.Value
var once_mapState sync.Once
func Get_mapState() gopurs_runtime.Value {
	once_mapState.Do(func() {
		mapState = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(pkg_Unsafe_Coerce.Get_unsafeCoerce(), gopurs_runtime.Apply(v_1, x_2)))
})
})
})
	})
	return mapState
}

var execState gopurs_runtime.Value
var once_execState sync.Once
func Get_execState() gopurs_runtime.Value {
	once_execState.Do(func() {
		execState = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, s_1).PtrVal.(map[string]gopurs_runtime.Value)["value1"]
})
})
	})
	return execState
}

var evalState gopurs_runtime.Value
var once_evalState sync.Once
func Get_evalState() gopurs_runtime.Value {
	once_evalState.Do(func() {
		evalState = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, s_1).PtrVal.(map[string]gopurs_runtime.Value)["value0"]
})
})
	})
	return evalState
}



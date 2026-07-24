package Control_Monad_State

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_State_Trans "gopurs/output/Control.Monad.State.Trans"
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
		runState = gopurs_runtime.Func2(Call_runState)
	})
	return runState
}

var mapState gopurs_runtime.Value
var once_mapState sync.Once
func Get_mapState() gopurs_runtime.Value {
	once_mapState.Do(func() {
		mapState = gopurs_runtime.Func3(Call_mapState)
	})
	return mapState
}

var execState gopurs_runtime.Value
var once_execState sync.Once
func Get_execState() gopurs_runtime.Value {
	once_execState.Do(func() {
		execState = gopurs_runtime.Func2(Call_execState)
	})
	return execState
}

var evalState gopurs_runtime.Value
var once_evalState sync.Once
func Get_evalState() gopurs_runtime.Value {
	once_evalState.Do(func() {
		evalState = gopurs_runtime.Func2(Call_evalState)
	})
	return evalState
}

func Call_runState(v_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(v_0_loop, x_1_loop)
}

func Call_mapState(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(f_0_loop, gopurs_runtime.Apply(v_1_loop, x_2_loop))
}

func Call_execState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(v_0_loop, s_1_loop).UnsafePtr)[1]
}

func Call_evalState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return (*[1024]gopurs_runtime.Value)(gopurs_runtime.Apply(v_0_loop, s_1_loop).UnsafePtr)[0]
}



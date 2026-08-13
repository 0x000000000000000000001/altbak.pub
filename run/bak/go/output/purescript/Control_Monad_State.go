package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Monad_State_unwrap gopurs_runtime.Value
var once_Control_Monad_State_unwrap sync.Once
func Get_Control_Monad_State_unwrap() gopurs_runtime.Value {
	once_Control_Monad_State_unwrap.Do(func() {
		cache_Control_Monad_State_unwrap = Get_Unsafe_Coerce_unsafeCoerce()
	})
	return cache_Control_Monad_State_unwrap
}

var cache_Control_Monad_State_withState gopurs_runtime.Value
var once_Control_Monad_State_withState sync.Once
func Get_Control_Monad_State_withState() gopurs_runtime.Value {
	once_Control_Monad_State_withState.Do(func() {
		cache_Control_Monad_State_withState = Get_Control_Monad_State_Trans_withStateT()
	})
	return cache_Control_Monad_State_withState
}

var cache_Control_Monad_State_runState gopurs_runtime.Value
var once_Control_Monad_State_runState sync.Once
func Get_Control_Monad_State_runState() gopurs_runtime.Value {
	once_Control_Monad_State_runState.Do(func() {
		cache_Control_Monad_State_runState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_State_runState(v_0_box, x_1_box))}
})
	})
	return cache_Control_Monad_State_runState
}

var cache_Control_Monad_State_mapState gopurs_runtime.Value
var once_Control_Monad_State_mapState sync.Once
func Get_Control_Monad_State_mapState() gopurs_runtime.Value {
	once_Control_Monad_State_mapState.Do(func() {
		cache_Control_Monad_State_mapState = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Monad_State_mapState(f_0_box, v_1_box, x_2_box))}
})
	})
	return cache_Control_Monad_State_mapState
}

var cache_Control_Monad_State_execState gopurs_runtime.Value
var once_Control_Monad_State_execState sync.Once
func Get_Control_Monad_State_execState() gopurs_runtime.Value {
	once_Control_Monad_State_execState.Do(func() {
		cache_Control_Monad_State_execState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_execState(v_0_box, s_1_box)
})
	})
	return cache_Control_Monad_State_execState
}

var cache_Control_Monad_State_evalState gopurs_runtime.Value
var once_Control_Monad_State_evalState sync.Once
func Get_Control_Monad_State_evalState() gopurs_runtime.Value {
	once_Control_Monad_State_evalState.Do(func() {
		cache_Control_Monad_State_evalState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Monad_State_evalState(v_0_box, s_1_box)
})
	})
	return cache_Control_Monad_State_evalState
}

func Call_Control_Monad_State_runState(v_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(v_0, x_1))
}

func Call_Control_Monad_State_mapState(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(v_1, x_2)))
}

func Call_Control_Monad_State_execState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(v_0, s_1).UnsafePtr).V1
}

func Call_Control_Monad_State_evalState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Apply(v_0, s_1).UnsafePtr).V0
}



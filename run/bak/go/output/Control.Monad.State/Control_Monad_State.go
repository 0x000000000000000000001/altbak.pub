package Control_Monad_State

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Monad_State_Trans "gopurs/output/Control.Monad.State.Trans"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var cache_withState gopurs_runtime.Value
var once_withState sync.Once
func Get_withState() gopurs_runtime.Value {
	once_withState.Do(func() {
		cache_withState = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return func(inner_arg0 func(interface{}) interface{}, inner_arg1 func(interface{}) gopurs_runtime.Value, inner_arg2 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply3(pkg_Control_Monad_State_Trans.Get_withStateT__func_func_interface____interface____func_interface____gopurs_runtime_Value__interface____gopurs_runtime_Value_1992341243(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return inner_arg1(gopurs_runtime.UnboxAny(arg0))
}), gopurs_runtime.Any(inner_arg2))
}(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply(arg1, gopurs_runtime.Any(inner_arg0))
}, gopurs_runtime.UnboxAny(arg2))
})
	})
	return cache_withState
}

var cache_runState gopurs_runtime.Value
var once_runState sync.Once
func Get_runState() gopurs_runtime.Value {
	once_runState.Do(func() {
		cache_runState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runState(func(inner_arg0 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0))
}, gopurs_runtime.UnboxAny(x_1_box))
})
	})
	return cache_runState
}

var cache_mapState gopurs_runtime.Value
var once_mapState sync.Once
func Get_mapState() gopurs_runtime.Value {
	once_mapState.Do(func() {
		cache_mapState = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapState(func(inner_arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0_box, inner_arg0)
}, func(inner_arg0 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_1_box, gopurs_runtime.Any(inner_arg0))
}, gopurs_runtime.UnboxAny(x_2_box))
})
	})
	return cache_mapState
}

var cache_execState gopurs_runtime.Value
var once_execState sync.Once
func Get_execState() gopurs_runtime.Value {
	once_execState.Do(func() {
		cache_execState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_execState(func(inner_arg0 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0))
}, gopurs_runtime.UnboxAny(s_1_box)))
})
	})
	return cache_execState
}

var cache_evalState gopurs_runtime.Value
var once_evalState sync.Once
func Get_evalState() gopurs_runtime.Value {
	once_evalState.Do(func() {
		cache_evalState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_evalState(func(inner_arg0 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0))
}, gopurs_runtime.UnboxAny(s_1_box)))
})
	})
	return cache_evalState
}

func Call_runState(v_0_loop func(interface{}) gopurs_runtime.Value, x_1_loop interface{}) gopurs_runtime.Value {
var v_0 func(interface{}) gopurs_runtime.Value = v_0_loop
_ = v_0
var x_1 interface{} = x_1_loop
_ = x_1
return v_0(x_1)
}

func Call_mapState(f_0_loop func(gopurs_runtime.Value) gopurs_runtime.Value, v_1_loop func(interface{}) gopurs_runtime.Value, x_2_loop interface{}) gopurs_runtime.Value {
var f_0 func(gopurs_runtime.Value) gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 func(interface{}) gopurs_runtime.Value = v_1_loop
_ = v_1
var x_2 interface{} = x_2_loop
_ = x_2
return f_0(v_1(x_2))
}

func Call_execState(v_0_loop func(interface{}) gopurs_runtime.Value, s_1_loop interface{}) interface{} {
var v_0 func(interface{}) gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 interface{} = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0(s_1).UnsafePtr).V1))
}

func Call_evalState(v_0_loop func(interface{}) gopurs_runtime.Value, s_1_loop interface{}) interface{} {
var v_0 func(interface{}) gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 interface{} = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_0(s_1).UnsafePtr).V0))
}

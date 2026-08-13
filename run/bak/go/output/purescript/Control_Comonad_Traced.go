package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Control_Comonad_Traced_traced gopurs_runtime.Value
var once_Control_Comonad_Traced_traced sync.Once
func Get_Control_Comonad_Traced_traced() gopurs_runtime.Value {
	once_Control_Comonad_Traced_traced.Do(func() {
		cache_Control_Comonad_Traced_traced = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_traced(x_0_box)
})
	})
	return cache_Control_Comonad_Traced_traced
}

var cache_Control_Comonad_Traced_runTraced gopurs_runtime.Value
var once_Control_Comonad_Traced_runTraced sync.Once
func Get_Control_Comonad_Traced_runTraced() gopurs_runtime.Value {
	once_Control_Comonad_Traced_runTraced.Do(func() {
		cache_Control_Comonad_Traced_runTraced = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Traced_runTraced(v_0_box)
})
	})
	return cache_Control_Comonad_Traced_runTraced
}

func Call_Control_Comonad_Traced_traced(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Traced_runTraced(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}



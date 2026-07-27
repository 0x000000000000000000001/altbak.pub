package Control_Comonad_Traced

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_traced gopurs_runtime.Value
var once_traced sync.Once
func Get_traced() gopurs_runtime.Value {
	once_traced.Do(func() {
		cache_traced = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traced(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(x_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_traced
}

var cache_runTraced gopurs_runtime.Value
var once_runTraced sync.Once
func Get_runTraced() gopurs_runtime.Value {
	once_runTraced.Do(func() {
		cache_runTraced = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runTraced(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
})
})
	})
	return cache_runTraced
}

func Call_traced(x_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var x_0 func(interface{}) interface{} = x_0_loop
_ = x_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(x_0(gopurs_runtime.UnboxAny(arg0)))
})
}

func Call_runTraced(v_0_loop func(interface{}) interface{}) gopurs_runtime.Value {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(v_0(gopurs_runtime.UnboxAny(arg0)))
})
}

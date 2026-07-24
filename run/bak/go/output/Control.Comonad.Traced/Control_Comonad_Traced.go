package Control_Comonad_Traced

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var traced gopurs_runtime.Value
var once_traced sync.Once
func Get_traced() gopurs_runtime.Value {
	once_traced.Do(func() {
		traced = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return traced
}

var runTraced gopurs_runtime.Value
var once_runTraced sync.Once
func Get_runTraced() gopurs_runtime.Value {
	once_runTraced.Do(func() {
		runTraced = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0_loop
}()
})
	})
	return runTraced
}





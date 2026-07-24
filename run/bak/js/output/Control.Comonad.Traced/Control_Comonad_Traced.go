package Control_Comonad_Traced

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var traced gopurs_runtime.Value
var once_traced sync.Once
func Get_traced() gopurs_runtime.Value {
	once_traced.Do(func() {
		traced = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return traced
}

var runTraced gopurs_runtime.Value
var once_runTraced sync.Once
func Get_runTraced() gopurs_runtime.Value {
	once_runTraced.Do(func() {
		runTraced = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return v_0
})
	})
	return runTraced
}





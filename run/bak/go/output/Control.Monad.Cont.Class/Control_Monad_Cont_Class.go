package Control_Monad_Cont_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var callCC gopurs_runtime.Value
var once_callCC sync.Once
func Get_callCC() gopurs_runtime.Value {
	once_callCC.Do(func() {
		callCC = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0_loop, "callCC")
}()
})
	})
	return callCC
}





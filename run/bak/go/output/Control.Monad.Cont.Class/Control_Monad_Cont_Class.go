package Control_Monad_Cont_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var callCC gopurs_runtime.Value
var once_callCC sync.Once
func Get_callCC() gopurs_runtime.Value {
	once_callCC.Do(func() {
		callCC = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return dict_0.PtrVal.(map[string]gopurs_runtime.Value)["callCC"]
})
	})
	return callCC
}



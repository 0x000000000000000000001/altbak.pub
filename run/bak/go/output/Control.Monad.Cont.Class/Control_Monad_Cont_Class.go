package Control_Monad_Cont_Class

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_callCC gopurs_runtime.Value
var once_callCC sync.Once
func Get_callCC() gopurs_runtime.Value {
	once_callCC.Do(func() {
		cache_callCC = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_callCC(dict_0_box)
})
	})
	return cache_callCC
}

var cache_callCC__gopurs_runtime_Value_2517782421 gopurs_runtime.Value
var once_callCC__gopurs_runtime_Value_2517782421 sync.Once
func Get_callCC__gopurs_runtime_Value_2517782421() gopurs_runtime.Value {
	once_callCC__gopurs_runtime_Value_2517782421.Do(func() {
		cache_callCC__gopurs_runtime_Value_2517782421 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_callCC__gopurs_runtime_Value_2517782421(dict_0_box)
})
	})
	return cache_callCC__gopurs_runtime_Value_2517782421
}

func Call_callCC(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "callCC")
}

func Call_callCC__gopurs_runtime_Value_2517782421(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "callCC")
}



package Control_Comonad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_extract gopurs_runtime.Value
var once_extract sync.Once
func Get_extract() gopurs_runtime.Value {
	once_extract.Do(func() {
		cache_extract = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extract(dict_0_box)
})
	})
	return cache_extract
}

var cache_extract__gopurs_runtime_Value_3853395930 gopurs_runtime.Value
var once_extract__gopurs_runtime_Value_3853395930 sync.Once
func Get_extract__gopurs_runtime_Value_3853395930() gopurs_runtime.Value {
	once_extract__gopurs_runtime_Value_3853395930.Do(func() {
		cache_extract__gopurs_runtime_Value_3853395930 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extract__gopurs_runtime_Value_3853395930(dict_0_box)
})
	})
	return cache_extract__gopurs_runtime_Value_3853395930
}

func Call_extract(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "extract")
}

func Call_extract__gopurs_runtime_Value_3853395930(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "extract")
}



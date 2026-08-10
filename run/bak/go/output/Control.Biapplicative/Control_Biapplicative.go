package Control_Biapplicative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Biapply "gopurs/output/Control.Biapply"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
)

var cache_bipure gopurs_runtime.Value
var once_bipure sync.Once
func Get_bipure() gopurs_runtime.Value {
	once_bipure.Do(func() {
		cache_bipure = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bipure(dict_0_box)
})
	})
	return cache_bipure
}

var cache_bipure__gopurs_runtime_Value_1449949980 gopurs_runtime.Value
var once_bipure__gopurs_runtime_Value_1449949980 sync.Once
func Get_bipure__gopurs_runtime_Value_1449949980() gopurs_runtime.Value {
	once_bipure__gopurs_runtime_Value_1449949980.Do(func() {
		cache_bipure__gopurs_runtime_Value_1449949980 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bipure__gopurs_runtime_Value_1449949980(dict_0_box)
})
	})
	return cache_bipure__gopurs_runtime_Value_1449949980
}

var cache_biapplicativeTuple gopurs_runtime.Value
var once_biapplicativeTuple sync.Once
func Get_biapplicativeTuple() gopurs_runtime.Value {
	once_biapplicativeTuple.Do(func() {
		cache_biapplicativeTuple = gopurs_runtime.RecordDict2("Biapply0", "bipure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Biapply.Get_biapplyTuple()
}), pkg_Data_Tuple.Get_Tuple())
	})
	return cache_biapplicativeTuple
}

func Call_bipure(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bipure")
}

func Call_bipure__gopurs_runtime_Value_1449949980(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bipure")
}



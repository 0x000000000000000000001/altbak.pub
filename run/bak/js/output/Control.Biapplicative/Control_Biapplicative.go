package Control_Biapplicative

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Control_Biapply "gopurs/output/Control.Biapply"
)

var bipure gopurs_runtime.Value
var once_bipure sync.Once
func Get_bipure() gopurs_runtime.Value {
	once_bipure.Do(func() {
		bipure = gopurs_runtime.Func(func(dict_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dict_0, "bipure")
})
	})
	return bipure
}

var biapplicativeTuple gopurs_runtime.Value
var once_biapplicativeTuple sync.Once
func Get_biapplicativeTuple() gopurs_runtime.Value {
	once_biapplicativeTuple.Do(func() {
		biapplicativeTuple = gopurs_runtime.RecordDict2("bipure", "Biapply0", pkg_Data_Tuple.Get_Tuple(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Biapply.Get_biapplyTuple()
}))
	})
	return biapplicativeTuple
}





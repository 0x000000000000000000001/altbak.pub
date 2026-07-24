package Data_Field

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var field gopurs_runtime.Value
var once_field sync.Once
func Get_field() gopurs_runtime.Value {
	once_field.Do(func() {
		field = gopurs_runtime.Func2(func(dictEuclideanRing_0_box gopurs_runtime.Value, dictDivisionRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_field(dictEuclideanRing_0_box, dictDivisionRing_1_box)
})
	})
	return field
}

func Call_field(dictEuclideanRing_0_loop gopurs_runtime.Value, dictDivisionRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
var dictDivisionRing_1 gopurs_runtime.Value = dictDivisionRing_1_loop
_ = dictDivisionRing_1
return gopurs_runtime.RecordDict2("EuclideanRing0", "DivisionRing1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEuclideanRing_0
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return dictDivisionRing_1
}))
}



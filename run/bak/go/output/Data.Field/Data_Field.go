package Data_Field

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_field gopurs_runtime.Value
var once_field sync.Once
func Get_field() gopurs_runtime.Value {
	once_field.Do(func() {
		cache_field = gopurs_runtime.Func2(func(dictEuclideanRing_0_box gopurs_runtime.Value, dictDivisionRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_field(dictEuclideanRing_0_box, dictDivisionRing_1_box))
})
	})
	return cache_field
}

func Call_field(dictEuclideanRing_0_loop gopurs_runtime.Value, dictDivisionRing_1_loop gopurs_runtime.Value) interface{} {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
var dictDivisionRing_1 gopurs_runtime.Value = dictDivisionRing_1_loop
_ = dictDivisionRing_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("DivisionRing1", "EuclideanRing0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return dictDivisionRing_1
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEuclideanRing_0
})))
}

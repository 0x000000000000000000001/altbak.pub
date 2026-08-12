package Data_Field

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_field gopurs_runtime.Value
var once_field sync.Once
func Get_field() gopurs_runtime.Value {
	once_field.Do(func() {
		cache_field = gopurs_runtime.Func2(func(dictEuclideanRing_0_box gopurs_runtime.Value, dictDivisionRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_field(dictEuclideanRing_0_box, dictDivisionRing_1_box)
})
	})
	return cache_field
}

type Constructor_Field[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3997783546] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Field[gopurs_runtime.Value])(ptr)
		switch key {
		case "DivisionRing1": return c.V0
		case "EuclideanRing0": return c.V1
		default: panic("Key not found in dictionary Constructor_Field: " + key)
		}
	}
}


func Call_field(dictEuclideanRing_0_loop gopurs_runtime.Value, dictDivisionRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
var dictDivisionRing_1 gopurs_runtime.Value = dictDivisionRing_1_loop
_ = dictDivisionRing_1
return gopurs_runtime.RecordDict2("DivisionRing1", "EuclideanRing0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return dictDivisionRing_1
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return dictEuclideanRing_0
}))
}



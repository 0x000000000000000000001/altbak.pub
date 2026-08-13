package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Field_Field_dollarDict gopurs_runtime.Value
var once_Data_Field_Field_dollarDict sync.Once
func Get_Data_Field_Field_dollarDict() gopurs_runtime.Value {
	once_Data_Field_Field_dollarDict.Do(func() {
		cache_Data_Field_Field_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Field_Field_dollarDict(x_0_box)
})
	})
	return cache_Data_Field_Field_dollarDict
}

var cache_Data_Field_field gopurs_runtime.Value
var once_Data_Field_field sync.Once
func Get_Data_Field_field() gopurs_runtime.Value {
	once_Data_Field_field.Do(func() {
		cache_Data_Field_field = gopurs_runtime.Func2(func(dictEuclideanRing_0_box gopurs_runtime.Value, dictDivisionRing_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Field_field(dictEuclideanRing_0_box, dictDivisionRing_1_box)
})
	})
	return cache_Data_Field_field
}

type Constructor_Data_Field_Field struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[3997783546] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_Field_Field)(ptr)
		_ = c
		switch key {
		case "DivisionRing1": return gopurs_runtime.Box(c.V0)
		case "EuclideanRing0": return gopurs_runtime.Box(c.V1)
		default: panic("Key not found in dictionary Constructor_Data_Field_Field: " + key)
		}
	}
}


func Call_Data_Field_Field_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Field_field(dictEuclideanRing_0_loop gopurs_runtime.Value, dictDivisionRing_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
var dictDivisionRing_1 gopurs_runtime.Value = dictDivisionRing_1_loop
_ = dictDivisionRing_1
return gopurs_runtime.Value{Type: 9, IntVal: 3997783546, UnsafePtr: unsafe.Pointer(&Constructor_Data_Field_Field{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2548491258, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_DivisionRing_DivisionRing](dictDivisionRing_1))}
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing](dictEuclideanRing_0))}
})})}
}



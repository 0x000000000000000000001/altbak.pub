package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Char_toCharCode gopurs_runtime.Value
var once_Data_Char_toCharCode sync.Once
func Get_Data_Char_toCharCode() gopurs_runtime.Value {
	once_Data_Char_toCharCode.Do(func() {
		cache_Data_Char_toCharCode = Call_Data_Enum_fromEnum(Rebox_Data_Char_3569500834_123048125(Rebox_Data_Char_123048125_3569500834(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Enum_boundedEnumChar()))))
	})
	return cache_Data_Char_toCharCode
}

var cache_Data_Char_fromCharCode gopurs_runtime.Value
var once_Data_Char_fromCharCode sync.Once
func Get_Data_Char_fromCharCode() gopurs_runtime.Value {
	once_Data_Char_fromCharCode.Do(func() {
		cache_Data_Char_fromCharCode = Call_Data_Enum_toEnum(Rebox_Data_Char_3569500834_123048125(Rebox_Data_Char_123048125_3569500834(gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]](Get_Data_Enum_boundedEnumChar()))))
	})
	return cache_Data_Char_fromCharCode
}

func Rebox_Data_Char_123048125_3569500834(in *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]) *Constructor_Data_Enum_BoundedEnum[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[string]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}

func Rebox_Data_Char_3569500834_123048125(in *Constructor_Data_Enum_BoundedEnum[string]) *Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Enum_BoundedEnum[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
		out.V4 = in.V4
	return out
}



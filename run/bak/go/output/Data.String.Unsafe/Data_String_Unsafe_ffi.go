package Data_String_Unsafe

import "gopurs/output/gopurs_runtime"

func CharAt(i any) any {
	return func(s any) any {
		str := s.(string)
		idx := int(i.(int))
		if idx >= 0 && idx < len(str) {
			return (string(str[idx]))
		}
		panic("Data.String.Unsafe.charAt: Invalid index.")
	}
}

func Char(s any) any {
	str := s.(string)
	if len(str) == 1 {
		return (string(str[0]))
	}
	panic("Data.String.Unsafe.char: Expected string of length 1.")
}


// --- Auto-generated FFI wrappers ---
func Call_charAt(arg0 any) any {
	return CharAt(arg0)
}
var _Gopurs_CharAt = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := CharAt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_char(arg0 any) any {
	return Char(arg0)
}
var _Gopurs_Char = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Char(go_arg0)
	return gopurs_runtime.Box(go_res)
})

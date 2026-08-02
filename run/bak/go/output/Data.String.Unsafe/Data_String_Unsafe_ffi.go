package Data_String_Unsafe

import "gopurs/output/gopurs_runtime"

func CharAt(i interface{}) interface{} {
	return func(s interface{}) interface{} {
		str := s.(string)
		idx := int(i.(int))
		if idx >= 0 && idx < len(str) {
			return (string(str[idx]))
		}
		panic("Data.String.Unsafe.charAt: Invalid index.")
	}
}

func Char(s interface{}) interface{} {
	str := s.(string)
	if len(str) == 1 {
		return (string(str[0]))
	}
	panic("Data.String.Unsafe.char: Expected string of length 1.")
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Char = // TAST: (Func [String] Char)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Char(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_CharAt = // TAST: (Func [Int, String] Char)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := CharAt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
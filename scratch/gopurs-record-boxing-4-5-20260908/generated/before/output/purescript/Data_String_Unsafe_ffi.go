package purescript

import "gopurs/output/gopurs_runtime"

func Data_String_Unsafe_CharAt(i interface{}) interface{} {
	return func(s interface{}) interface{} {
		str := gopurs_runtime.Unbox[string](s)
		idx := gopurs_runtime.Unbox[int](i)
		if idx >= 0 && idx < len(str) {
			return string(str[idx])
		}
		panic("Data.String.Unsafe.charAt: Invalid index.")
	}
}

func Data_String_Unsafe_Char(s interface{}) interface{} {
	str := gopurs_runtime.Unbox[string](s)
	if len(str) == 1 {
		return string(str[0])
	}
	panic("Data.String.Unsafe.char: Expected string of length 1.")
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Data_String_Unsafe_Char = // TAST: (Func [String] Char)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_String_Unsafe_Char(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Data_String_Unsafe_CharAt = // TAST: (Func [Int, String] Char)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Data_String_Unsafe_CharAt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
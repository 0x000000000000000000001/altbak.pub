package Data_String_Unsafe

import "gopurs/output/gopurs_runtime"

var CharAt = gopurs_runtime.Func(func(i gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
		str := s.StrVal
		idx := int(i.IntVal)
		if idx >= 0 && idx < len(str) {
			return gopurs_runtime.Str(string(str[idx]))
		}
		panic("Data.String.Unsafe.charAt: Invalid index.")
	})
})

var Char = gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
	str := s.StrVal
	if len(str) == 1 {
		return gopurs_runtime.Str(string(str[0]))
	}
	panic("Data.String.Unsafe.char: Expected string of length 1.")
})


// --- Auto-generated FFI wrappers ---
var _Gopurs_CharAt = gopurs_runtime.Box(CharAt)
var _Gopurs_Char = gopurs_runtime.Box(Char)

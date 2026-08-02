package Data_Enum

import "gopurs/output/gopurs_runtime"

func FromCharCode(c int) string { return string(rune(c)) }
func ToCharCode(c string) int { return int([]rune(c)[0]) }


// --- Auto-generated FFI wrappers ---
var _Gopurs_FromCharCode = // TAST: (Func [Int] Char)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := FromCharCode(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ToCharCode = // TAST: (Func [Char] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ToCharCode(go_arg0)
	return gopurs_runtime.Box(go_res)
})
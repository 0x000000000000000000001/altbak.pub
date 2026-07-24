package Data_Enum

import "gopurs/output/gopurs_runtime"

func FromCharCode(c int) string { return string(rune(c)) }
func ToCharCode(c string) int { return int([]rune(c)[0]) }


// --- Auto-generated FFI wrappers ---
func Call_fromCharCode(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := FromCharCode(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_FromCharCode = gopurs_runtime.Func(Call_fromCharCode)
func Call_toCharCode(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ToCharCode(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_ToCharCode = gopurs_runtime.Func(Call_toCharCode)

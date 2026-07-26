package Data_Enum

import "gopurs/output/gopurs_runtime"

func FromCharCode(c int) string { return string(rune(c)) }
func ToCharCode(c string) int { return int([]rune(c)[0]) }


// --- Auto-generated FFI wrappers ---
func Call_fromCharCode(arg0 int) string {
	return FromCharCode(arg0)
}
var _Gopurs_FromCharCode = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := FromCharCode(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_toCharCode(arg0 string) int {
	return ToCharCode(arg0)
}
var _Gopurs_ToCharCode = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := ToCharCode(go_arg0)
	return gopurs_runtime.Box(go_res)
})

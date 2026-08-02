package Data_Show_Generic

import "gopurs/output/gopurs_runtime"

import "strings"
func Intercalate(separator string, arr []string) string {
	return strings.Join(arr, separator)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Intercalate = // TAST: (Func [String, (Array String)] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]string, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = gopurs_runtime.Unbox[string](v) }
	go_res := Intercalate(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
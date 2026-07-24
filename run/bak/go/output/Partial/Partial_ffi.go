package Partial

import "gopurs/output/gopurs_runtime"

func _CrashWith(msg string) any {
	panic(msg)
}


// --- Auto-generated FFI wrappers ---
func Call__CrashWith(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := _CrashWith(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs__CrashWith = gopurs_runtime.Func(Call__CrashWith)

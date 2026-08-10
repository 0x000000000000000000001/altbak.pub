package Partial

import "gopurs/output/gopurs_runtime"

func _CrashWith(msg string) interface{} {
	panic(msg)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs__CrashWith = // TAST: (ForAll [a] (Func [String] (TypeVar a)))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_res := _CrashWith(go_arg0)
	return gopurs_runtime.Box(go_res)
})
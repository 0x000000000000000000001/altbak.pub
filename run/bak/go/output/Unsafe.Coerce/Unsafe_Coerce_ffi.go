package Unsafe_Coerce

import "gopurs/output/gopurs_runtime"

func UnsafeCoerce(x interface{}) interface{} {
	return x
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_UnsafeCoerce = // TAST: (Func [(TypeVar a)] (TypeVar b))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := UnsafeCoerce(go_arg0)
	return gopurs_runtime.Box(go_res)
})
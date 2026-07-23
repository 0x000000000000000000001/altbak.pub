package Data_Symbol

import "gopurs/output/gopurs_runtime"

func UnsafeCoerce(x any) any { return x }


// --- Auto-generated FFI wrappers ---
var _Gopurs_UnsafeCoerce = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_res := UnsafeCoerce(go_arg0)
	return gopurs_runtime.Box(go_res)
})

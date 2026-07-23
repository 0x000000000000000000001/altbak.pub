package Unsafe_Coerce

import "gopurs/output/gopurs_runtime"

var UnsafeCoerce = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return x
})


// --- Auto-generated FFI wrappers ---
var _Gopurs_UnsafeCoerce = gopurs_runtime.Box(UnsafeCoerce)

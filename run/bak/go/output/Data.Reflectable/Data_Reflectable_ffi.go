package Data_Reflectable

import "gopurs/output/gopurs_runtime"

func UnsafeCoerce(x any) any { return x }


// --- Auto-generated FFI wrappers ---
func Call_unsafeCoerce(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := UnsafeCoerce(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_UnsafeCoerce = gopurs_runtime.Func(Call_unsafeCoerce)

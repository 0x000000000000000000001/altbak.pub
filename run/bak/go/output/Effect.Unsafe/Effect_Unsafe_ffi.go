package Effect_Unsafe

import "gopurs/output/gopurs_runtime"

func UnsafePerformEffect(f func() any) any { return f() }


// --- Auto-generated FFI wrappers ---
func Call_unsafePerformEffect(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func() any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Value{})
		}
	go_res := UnsafePerformEffect(go_arg0)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_UnsafePerformEffect = gopurs_runtime.Func(Call_unsafePerformEffect)

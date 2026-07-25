package Partial_Unsafe

import "gopurs/output/gopurs_runtime"

func _UnsafePartial(f func(interface{}) interface{}) interface{} {
	return f(nil)
}


// --- Auto-generated FFI wrappers ---
func Call__UnsafePartial(arg0 func(interface{}) interface{}) interface{} {
	return _UnsafePartial(arg0)
}
var _Gopurs__UnsafePartial = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_res := _UnsafePartial(go_arg0)
	return gopurs_runtime.Box(go_res)
})

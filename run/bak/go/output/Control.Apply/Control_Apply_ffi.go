package Control_Apply

import "gopurs/output/gopurs_runtime"

func ArrayApply(fs []func(interface{}) interface{}, xs []interface{}) []interface{} {
	result := make([]interface{}, 0, len(fs)*len(xs))
	for _, f := range fs {
		for _, x := range xs {
			result = append(result, f(x))
		}
	}
	return result
}


// --- Auto-generated FFI wrappers ---
func Call_arrayApply(arg0 []func(interface{}) interface{}, arg1 []interface{}) []interface{} {
	return ArrayApply(arg0, arg1)
}
var _Gopurs_ArrayApply = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]func(interface{}) interface{}, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[func(interface{}) interface{}](v) }
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := ArrayApply(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})

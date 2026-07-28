package Data_Functor

import "gopurs/output/gopurs_runtime"

func ArrayMap(f func(interface{}) interface{}, arr []interface{}) []interface{} {
	result := make([]interface{}, len(arr))
	for i, v := range arr {
		result[i] = f(v)
	}
	return result
}


// --- Auto-generated FFI wrappers ---
func Call_arrayMap(arg0 func(interface{}) interface{}, arg1 []interface{}) []interface{} {
	return ArrayMap(arg0, arg1)
}
var _Gopurs_ArrayMap = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) interface{} {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := ArrayMap(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})

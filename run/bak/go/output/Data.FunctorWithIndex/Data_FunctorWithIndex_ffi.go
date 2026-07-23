package Data_FunctorWithIndex

import "gopurs/output/gopurs_runtime"

func MapWithIndexArray(f func(int) func(interface{}) interface{}, xs []interface{}) []interface{} {
	result := make([]interface{}, len(xs))
	for i, x := range xs {
		result[i] = f(i)(x)
	}
	return result
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_MapWithIndexArray = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0 int) func(interface{}) interface{} {
		res := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0))
		return gopurs_runtime.Unbox[func(interface{}) interface{}](res)
	}
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v.PtrVal }
	go_res := MapWithIndexArray(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})

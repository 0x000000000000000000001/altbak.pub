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
func Call_mapWithIndexArray(arg0 func(int) func(interface{}) interface{}, arg1 []interface{}) []interface{} {
	return MapWithIndexArray(arg0, arg1)
}
var _Gopurs_MapWithIndexArray = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	arg1_arr := arg1.PtrVal().([]gopurs_runtime.Value)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := MapWithIndexArray(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})

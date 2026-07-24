package Control_Bind

import "gopurs/output/gopurs_runtime"

func ArrayBind(arr []any, f func(any) []any) []any {
	var result []any
	for _, v := range arr {
		result = append(result, f(v)...)
	}
	return result
}


// --- Auto-generated FFI wrappers ---
func Call_arrayBind(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := arg0.PtrVal.([]gopurs_runtime.Value)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := func(p0_0 any) []any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			res_arr0 := inner_res0.PtrVal.([]gopurs_runtime.Value)
			res_go0 := make([]any, len(res_arr0))
			for i, v := range res_arr0 { res_go0[i] = gopurs_runtime.Unbox[any](v) }
			return res_go0
		}
	go_res := ArrayBind(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_ArrayBind = gopurs_runtime.Func2(Call_arrayBind)

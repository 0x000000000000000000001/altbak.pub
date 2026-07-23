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
var _Gopurs_ArrayBind = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := arg0.PtrVal.([]gopurs_runtime.Value)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v.PtrVal }
	go_arg1 := func(p0 any) []any {
		res := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0))
		res_arr := res.PtrVal.([]gopurs_runtime.Value)
		res_go := make([]any, len(res_arr))
		for i, v := range res_arr { res_go[i] = gopurs_runtime.Unbox[any](v) }
		return res_go
	}
	go_res := ArrayBind(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
})

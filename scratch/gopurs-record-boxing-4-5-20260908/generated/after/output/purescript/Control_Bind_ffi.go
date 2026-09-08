package purescript

import "gopurs/output/gopurs_runtime"

func Control_Bind_ArrayBind(arr []interface{}, f func(interface{}) []interface{}) []interface{} {
	var result []interface{}
	for _, v := range arr {
		result = append(result, f(v)...)
	}
	return result
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Control_Bind_ArrayBind = // TAST: (ForAll [a, b] (Func [(Array (TypeVar a)), (Func [(TypeVar a)] (Array (TypeVar b)))] (Array (TypeVar b))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	go_arg1 := func(p0_0 any) []any {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			res_arr0 := *(*[]gopurs_runtime.Value)(inner_res0.UnsafePtr)
			res_go0 := make([]any, len(res_arr0))
			for i, v := range res_arr0 { res_go0[i] = gopurs_runtime.Unbox[any](v) }
			return res_go0
		}
	go_res := Control_Bind_ArrayBind(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
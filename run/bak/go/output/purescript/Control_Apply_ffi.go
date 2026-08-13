package purescript

import "gopurs/output/gopurs_runtime"

func Control_Apply_ArrayApply(fs []func(interface{}) interface{}, xs []interface{}) []interface{} {
	result := make([]interface{}, 0, len(fs)*len(xs))
	for _, f := range fs {
		for _, x := range xs {
			result = append(result, f(x))
		}
	}
	return result
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Control_Apply_ArrayApply = // TAST: (ForAll [a, b] (Func [(Array (Func [(TypeVar a)] (TypeVar b))), (Array (TypeVar a))] (Array (TypeVar b))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]func(any) any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = gopurs_runtime.Unbox[func(any) any](v) }
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := Control_Apply_ArrayApply(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
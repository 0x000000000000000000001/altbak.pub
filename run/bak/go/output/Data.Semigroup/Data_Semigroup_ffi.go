package Data_Semigroup

import "gopurs/output/gopurs_runtime"

func ConcatString(s1 string, s2 string) string {
	return s1 + s2
}
func ConcatArray(xs []interface{}, ys []interface{}) []interface{} {
	if len(xs) == 0 {
		return ys
	}
	if len(ys) == 0 {
		return xs
	}
	res := make([]interface{}, 0, len(xs)+len(ys))
	res = append(res, xs...)
	res = append(res, ys...)
	return res
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_ConcatArray = // TAST: (ForAll [a] (Func [(Array (TypeVar a)), (Array (TypeVar a))] (Array (TypeVar a))))
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := *(*[]gopurs_runtime.Value)(arg0.UnsafePtr)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := ConcatArray(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
				res_arr := make([]gopurs_runtime.Value, len(go_res))
				for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
				return gopurs_runtime.Array(res_arr)
			}()
})
var _Gopurs_ConcatString = // TAST: (Func [String, String] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := ConcatString(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
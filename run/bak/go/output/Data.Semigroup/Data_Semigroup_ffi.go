package Data_Semigroup

import "gopurs/output/gopurs_runtime"

func ConcatString(s1 string, s2 string) string {
	return s1 + s2
}
func ConcatArray(xs []any, ys []any) []any {
	if len(xs) == 0 {
		return ys
	}
	if len(ys) == 0 {
		return xs
	}
	res := make([]any, 0, len(xs)+len(ys))
	res = append(res, xs...)
	res = append(res, ys...)
	return res
}


// --- Auto-generated FFI wrappers ---
func Call_concatString(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := gopurs_runtime.Unbox[string](arg1)
	go_res := ConcatString(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_ConcatString = gopurs_runtime.Func2(Call_concatString)
func Call_concatArray(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	arg0_arr := arg0.PtrVal.([]gopurs_runtime.Value)
	go_arg0 := make([]any, len(arg0_arr))
	for i, v := range arg0_arr { go_arg0[i] = v }
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := ConcatArray(go_arg0, go_arg1)
	return func() gopurs_runtime.Value {
			res_arr := make([]gopurs_runtime.Value, len(go_res))
			for i, v := range go_res { res_arr[i] = gopurs_runtime.Box(v) }
			return gopurs_runtime.Array(res_arr)
		}()
}
var _Gopurs_ConcatArray = gopurs_runtime.Func2(Call_concatArray)

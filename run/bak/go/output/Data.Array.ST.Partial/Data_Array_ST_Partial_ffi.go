package Data_Array_ST_Partial

import "gopurs/output/gopurs_runtime"

func PeekImpl(i int, xs []any) func() any { return func() any { return xs[i] } }
func PokeImpl(i int, a any, xs []any) func() bool { return func() bool { xs[i] = a; return true } }


// --- Auto-generated FFI wrappers ---
var _Gopurs_PeekImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v.PtrVal }
	go_res := PeekImpl(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_PokeImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := arg1.PtrVal
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v.PtrVal }
	go_res := PokeImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})

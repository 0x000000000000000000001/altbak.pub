package Data_Array_ST_Partial

import "gopurs/output/gopurs_runtime"

func PeekImpl(i int, xs []interface{}) func() interface{} { return func() interface{} { return xs[i] } }
func PokeImpl(i int, a interface{}, xs []interface{}) func() bool { return func() bool { xs[i] = a; return true } }


// --- Auto-generated FFI wrappers ---
func Call_peekImpl(arg0 int, arg1 []interface{}) func() interface{} { return func() interface{} { return xs[i] } } {
	return PeekImpl(arg0, arg1)
}
var _Gopurs_PeekImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	arg1_arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
	go_arg1 := make([]interface{}, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	go_res := PeekImpl(go_arg0, go_arg1)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
func Call_pokeImpl(arg0 int, arg1 interface{}, arg2 []interface{}) func() bool { return func() bool { xs[i] = a; return true } } {
	return PokeImpl(arg0, arg1, arg2)
}
var _Gopurs_PokeImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := arg1
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := PokeImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})

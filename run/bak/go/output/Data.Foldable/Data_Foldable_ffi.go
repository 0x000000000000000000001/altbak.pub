package Data_Foldable

import "gopurs/output/gopurs_runtime"

func FoldrArray(f func(interface{}) func(interface{}) interface{}, init interface{}, xs []interface{}) interface{} {
	acc := init
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i])(acc)
	}
	return acc
}

func FoldlArray(f func(interface{}) func(interface{}) interface{}, init interface{}, xs []interface{}) interface{} {
	acc := init
	for i := 0; i < len(xs); i++ {
		acc = f(acc)(xs[i])
	}
	return acc
}


// --- Auto-generated FFI wrappers ---
func Call_foldrArray(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg1 := arg1
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := FoldrArray(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_FoldrArray = gopurs_runtime.Func3(Call_foldrArray)
func Call_foldlArray(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 interface{}) func(interface{}) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 interface{}) interface{} {
			return gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
		}
		}
	go_arg1 := arg1
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := FoldlArray(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_FoldlArray = gopurs_runtime.Func3(Call_foldlArray)

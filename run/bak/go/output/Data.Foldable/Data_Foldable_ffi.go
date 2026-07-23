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
var _Gopurs_FoldrArray = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0 interface{}) func(interface{}) interface{} {
		res := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0))
		return gopurs_runtime.Unbox[func(interface{}) interface{}](res)
	}
	go_arg1 := arg1.PtrVal
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v.PtrVal }
	go_res := FoldrArray(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FoldlArray = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0 interface{}) func(interface{}) interface{} {
		res := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0))
		return gopurs_runtime.Unbox[func(interface{}) interface{}](res)
	}
	go_arg1 := arg1.PtrVal
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]interface{}, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v.PtrVal }
	go_res := FoldlArray(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})

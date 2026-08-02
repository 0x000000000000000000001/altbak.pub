package Data_Foldable

import "gopurs/output/gopurs_runtime"

func FoldrArray(f func(interface{}, interface{}) interface{}, init interface{}, xs []interface{}) interface{} {
	acc := init
	for i := len(xs) - 1; i >= 0; i-- {
		acc = f(xs[i], acc)
	}
	return acc
}

func FoldlArray(f func(interface{}, interface{}) interface{}, init interface{}, xs []interface{}) interface{} {
	acc := init
	for i := 0; i < len(xs); i++ {
		acc = f(acc, xs[i])
	}
	return acc
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_FoldlArray = // TAST: (Func [(Func [(TypeVar b), (TypeVar a)] (TypeVar b)), (TypeVar b), (Array (TypeVar a))] (TypeVar b))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	go_arg1 := arg1
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := FoldlArray(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FoldrArray = // TAST: (Func [(Func [(TypeVar a), (TypeVar b)] (TypeVar b)), (TypeVar b), (Array (TypeVar a))] (TypeVar b))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any, p0_1 any) any {
			return gopurs_runtime.Apply2(arg0, gopurs_runtime.Box(p0_0), gopurs_runtime.Box(p0_1))
		}
	go_arg1 := arg1
	arg2_arr := *(*[]gopurs_runtime.Value)(arg2.UnsafePtr)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := FoldrArray(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
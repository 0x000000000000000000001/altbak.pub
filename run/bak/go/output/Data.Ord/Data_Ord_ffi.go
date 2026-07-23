package Data_Ord

import "gopurs/output/gopurs_runtime"

func OrdBooleanImpl(lt any, eq any, gt any, x bool, y bool) any {
	if !x && y {
		return lt
	} else if x == y {
		return eq
	}
	return gt
}
func OrdIntImpl(lt any, eq any, gt any, x int, y int) any {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdCharImpl(lt any, eq any, gt any, x string, y string) any {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdStringImpl(lt any, eq any, gt any, x string, y string) any {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdNumberImpl(lt any, eq any, gt any, x float64, y float64) any {
	if x < y { return lt }
	if x == y { return eq }
	return gt
}
func OrdArrayImpl(f func(any) func(any) int, xs []any, ys []any) int {
	xlen := len(xs)
	ylen := len(ys)
	for i := 0; i < xlen && i < ylen; i++ {
		o := f(xs[i])(ys[i])
		if o != 0 {
			return o
		}
	}
	if xlen == ylen {
		return 0
	} else if xlen > ylen {
		return 1
	} else {
		return -1
	}
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_OrdBooleanImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_arg3 := gopurs_runtime.Unbox[bool](arg3)
	go_arg4 := gopurs_runtime.Unbox[bool](arg4)
	go_res := OrdBooleanImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdIntImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[int](arg4)
	go_res := OrdIntImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdCharImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := OrdCharImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdStringImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := OrdStringImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdNumberImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0.PtrVal
	go_arg1 := arg1.PtrVal
	go_arg2 := arg2.PtrVal
	go_arg3 := gopurs_runtime.Unbox[float64](arg3)
	go_arg4 := gopurs_runtime.Unbox[float64](arg4)
	go_res := OrdNumberImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_OrdArrayImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0 any) func(any) int {
		res := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0))
		return gopurs_runtime.Unbox[func(any) int](res)
	}
	arg1_arr := arg1.PtrVal.([]gopurs_runtime.Value)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v.PtrVal }
	arg2_arr := arg2.PtrVal.([]gopurs_runtime.Value)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v.PtrVal }
	go_res := OrdArrayImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})

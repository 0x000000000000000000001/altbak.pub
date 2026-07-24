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
func Call_ordBooleanImpl(arg0 any, arg1 any, arg2 any, arg3 bool, arg4 bool) any {
	return OrdBooleanImpl(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs_OrdBooleanImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[bool](arg3)
	go_arg4 := gopurs_runtime.Unbox[bool](arg4)
	go_res := OrdBooleanImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call_ordIntImpl(arg0 any, arg1 any, arg2 any, arg3 int, arg4 int) any {
	return OrdIntImpl(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs_OrdIntImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[int](arg4)
	go_res := OrdIntImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call_ordCharImpl(arg0 any, arg1 any, arg2 any, arg3 string, arg4 string) any {
	return OrdCharImpl(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs_OrdCharImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := OrdCharImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call_ordStringImpl(arg0 any, arg1 any, arg2 any, arg3 string, arg4 string) any {
	return OrdStringImpl(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs_OrdStringImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_arg4 := gopurs_runtime.Unbox[string](arg4)
	go_res := OrdStringImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call_ordNumberImpl(arg0 any, arg1 any, arg2 any, arg3 float64, arg4 float64) any {
	return OrdNumberImpl(arg0, arg1, arg2, arg3, arg4)
}
var _Gopurs_OrdNumberImpl = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := arg1
	go_arg2 := arg2
	go_arg3 := gopurs_runtime.Unbox[float64](arg3)
	go_arg4 := gopurs_runtime.Unbox[float64](arg4)
	go_res := OrdNumberImpl(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4)
	return gopurs_runtime.Box(go_res)
})
func Call_ordArrayImpl(arg0 func(any) func(any) int, arg1 []any, arg2 []any) int {
	return OrdArrayImpl(arg0, arg1, arg2)
}
var _Gopurs_OrdArrayImpl = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 any) func(any) int {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 any) int {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return gopurs_runtime.Unbox[int](inner_res1)
		}
		}
	arg1_arr := arg1.PtrVal().([]gopurs_runtime.Value)
	go_arg1 := make([]any, len(arg1_arr))
	for i, v := range arg1_arr { go_arg1[i] = v }
	arg2_arr := arg2.PtrVal().([]gopurs_runtime.Value)
	go_arg2 := make([]any, len(arg2_arr))
	for i, v := range arg2_arr { go_arg2[i] = v }
	go_res := OrdArrayImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})

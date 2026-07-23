package Data_Semiring

import "gopurs/output/gopurs_runtime"

func IntAdd(x int, y int) int {
	return x + y
}
func IntMul(x int, y int) int {
	return x * y
}
func NumAdd(n1 float64, n2 float64) float64 {
	return n1 + n2
}
func NumMul(n1 float64, n2 float64) float64 {
	return n1 * n2
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_IntAdd = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := IntAdd(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_IntMul = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := IntMul(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_NumAdd = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := NumAdd(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_NumMul = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := NumMul(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

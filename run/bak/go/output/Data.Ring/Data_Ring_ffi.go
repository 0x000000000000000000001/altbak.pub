package Data_Ring

import "gopurs/output/gopurs_runtime"

func IntSub(x int, y int) int {
	return x - y
}
func NumSub(x float64, y float64) float64 {
	return x - y
}


// --- Auto-generated FFI wrappers ---
func Call_intSub(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := IntSub(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_IntSub = gopurs_runtime.Func2(Call_intSub)
func Call_numSub(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := NumSub(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
}
var _Gopurs_NumSub = gopurs_runtime.Func2(Call_numSub)

package Data_EuclideanRing

import "gopurs/output/gopurs_runtime"

import "math"
func IntDegree(x int64) int64 {
	if x < 0 {
		x = -x
	}
	if x > 2147483647 {
		x = 2147483647
	}
	return x
}
func IntDiv(x int64, y int64) int64 {
	if y == 0 {
		return 0
	}
	if y > 0 {
		return int64(math.Floor(float64(x) / float64(y)))
	}
	return int64(-math.Floor(float64(x) / float64(-y)))
}
func IntMod(x int64, y int64) int64 {
	if y == 0 {
		return 0
	}
	if y < 0 {
		y = -y
	}
	return ((x % y) + y) % y
}
func NumDiv(n1 float64, n2 float64) float64 {
	return n1 / n2
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_IntDegree = // TAST: (Func [Int] Int)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_res := IntDegree(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_IntDiv = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := IntDiv(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_IntMod = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int64](arg0)
	go_arg1 := gopurs_runtime.Unbox[int64](arg1)
	go_res := IntMod(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_NumDiv = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := NumDiv(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
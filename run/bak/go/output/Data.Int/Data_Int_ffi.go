package Data_Int

import "gopurs/output/gopurs_runtime"

import (
	"math"
	"strconv"
)

func FromNumberImpl(just func(int) interface{}, nothing interface{}, n float64) interface{} {
	if math.IsNaN(n) || math.IsInf(n, 0) || math.Trunc(n) != n {
		return nothing
	}
	return just(int(n))
}

func ToNumber(n int) float64 {
	return float64(n)
}

func FromStringAsImpl(just func(int) interface{}, nothing interface{}, radix int, s string) interface{} {
	val, err := strconv.ParseInt(s, radix, 64)
	if err != nil {
		return nothing
	}
	return just(int(val))
}

func ToStringAs(radix int, i int) string {
	return strconv.FormatInt(int64(i), radix)
}

func Quot(x int, y int) int {
	if y == 0 {
		return 0
	}
	return x / y
}

func Rem(x int, y int) int {
	if y == 0 {
		return 0
	}
	return x % y
}

func Pow(x int, y int) int {
	if y < 0 {
		return 0
	}
	res := 1
	for i := 0; i < y; i++ {
		res *= x
	}
	return res
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_FromNumberImpl = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), Number] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[float64](arg2)
	go_res := FromNumberImpl(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FromStringAsImpl = // TAST: (Func [(Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]), Int, String] (ADT ["Data","Maybe","Maybe"] [Int]))
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) any {
			return gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
		}
	go_arg1 := arg1
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := gopurs_runtime.Unbox[string](arg3)
	go_res := FromStringAsImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Pow = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Pow(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Quot = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Quot(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Rem = // TAST: (Func [Int, Int] Int)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := Rem(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ToNumber = // TAST: (Func [Int] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_res := ToNumber(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_ToStringAs = // TAST: (Func [Int, Int] String)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_res := ToStringAs(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
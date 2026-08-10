package Data_Number

import "gopurs/output/gopurs_runtime"

import (
	"math"
	"strconv"
)

func IsNaN(n float64) bool {
	return math.IsNaN(n)
}

func IsFinite(v float64) bool {
	return !math.IsNaN(v) && !math.IsInf(v, 0)
}

func FromStringImpl(str string, isFinite func(float64) bool, just func(float64) interface{}, nothing interface{}) interface{} {
	val, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return nothing
	}
	if math.IsNaN(val) || math.IsInf(val, 0) {
		return nothing
	}
	return just(val)
}

func Abs(n float64) float64 {
	return math.Abs(n)
}

func Acos(n float64) float64 {
	return math.Acos(n)
}

func Asin(n float64) float64 {
	return math.Asin(n)
}

func Atan(n float64) float64 {
	return math.Atan(n)
}

func Atan2(y float64, x float64) float64 {
	return math.Atan2(y, x)
}

func Ceil(n float64) float64 {
	return math.Ceil(n)
}

func Cos(n float64) float64 {
	return math.Cos(n)
}

func Exp(n float64) float64 {
	return math.Exp(n)
}

func Floor(n float64) float64 {
	return math.Floor(n)
}

func Log(n float64) float64 {
	return math.Log(n)
}

func Max(n1 float64, n2 float64) float64 {
	return math.Max(n1, n2)
}

func Min(n1 float64, n2 float64) float64 {
	return math.Min(n1, n2)
}

func Pow(n float64, p float64) float64 {
	return math.Pow(n, p)
}

func Remainder(n float64, m float64) float64 {
	return math.Mod(n, m)
}

func Round(n float64) float64 {
	return math.Round(n)
}

func Sign(v float64) float64 {
	if math.IsNaN(v) || v == 0 {
		return v
	}
	if v < 0 {
		return -1
	}
	return 1
}

func Sin(n float64) float64 {
	return math.Sin(n)
}

func Sqrt(n float64) float64 {
	return math.Sqrt(n)
}

func Tan(n float64) float64 {
	return math.Tan(n)
}

func Trunc(n float64) float64 {
	return math.Trunc(n)
}

var Infinity = math.Inf(1)
var Nan = math.NaN()


// --- Auto-generated FFI wrappers ---
var _Gopurs_Abs = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Abs(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Acos = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Acos(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Asin = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Asin(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Atan = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Atan(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Atan2 = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Atan2(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Ceil = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Ceil(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Cos = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Cos(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Exp = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Exp(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Floor = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Floor(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FromStringImpl = // TAST: (ADT ["Data","Function","Uncurried","Fn4"] [String, (Func [Number] Boolean), (ForAll [a] (Func [(TypeVar a)] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)]))), (ForAll [a] (ADT ["Data","Maybe","Maybe"] [(TypeVar a)])), (ADT ["Data","Maybe","Maybe"] [Number])])
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[string](arg0)
	go_arg1 := func(p0_0 float64) bool {
			inner_res0 := gopurs_runtime.Apply(arg1, gopurs_runtime.Box(p0_0))
			return gopurs_runtime.Unbox[bool](inner_res0)
		}
	go_arg2 := func(p0_0 float64) any {
			return gopurs_runtime.Apply(arg2, gopurs_runtime.Box(p0_0))
		}
	go_arg3 := arg3
	go_res := FromStringImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Infinity = // TAST: Number
gopurs_runtime.Box(Infinity)
var _Gopurs_IsFinite = // TAST: (Func [Number] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := IsFinite(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_IsNaN = // TAST: (Func [Number] Boolean)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := IsNaN(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Log = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Log(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Max = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Max(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Min = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Min(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Nan = // TAST: Number
gopurs_runtime.Box(Nan)
var _Gopurs_Pow = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Pow(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Remainder = // TAST: (Func [Number, Number] Number)
gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Remainder(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Round = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Round(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Sign = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Sign(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Sin = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Sin(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Sqrt = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Sqrt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Tan = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Tan(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Trunc = // TAST: (Func [Number] Number)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Trunc(go_arg0)
	return gopurs_runtime.Box(go_res)
})
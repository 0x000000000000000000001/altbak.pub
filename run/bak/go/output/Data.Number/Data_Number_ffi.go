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

func FromStringImpl(str string, isFinite func(float64) bool, just func(float64) any, nothing any) any {
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
func Call_isNaN(arg0 float64) bool {
	return IsNaN(arg0)
}
var _Gopurs_IsNaN = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := IsNaN(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_isFinite(arg0 float64) bool {
	return IsFinite(arg0)
}
var _Gopurs_IsFinite = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := IsFinite(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_fromStringImpl(arg0 string, arg1 func(float64) bool, arg2 func(float64) any, arg3 any) any {
	return FromStringImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_FromStringImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
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
func Call_abs(arg0 float64) float64 {
	return Abs(arg0)
}
var _Gopurs_Abs = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Abs(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_acos(arg0 float64) float64 {
	return Acos(arg0)
}
var _Gopurs_Acos = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Acos(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_asin(arg0 float64) float64 {
	return Asin(arg0)
}
var _Gopurs_Asin = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Asin(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_atan(arg0 float64) float64 {
	return Atan(arg0)
}
var _Gopurs_Atan = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Atan(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_atan2(arg0 float64, arg1 float64) float64 {
	return Atan2(arg0, arg1)
}
var _Gopurs_Atan2 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Atan2(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_ceil(arg0 float64) float64 {
	return Ceil(arg0)
}
var _Gopurs_Ceil = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Ceil(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_cos(arg0 float64) float64 {
	return Cos(arg0)
}
var _Gopurs_Cos = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Cos(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_exp(arg0 float64) float64 {
	return Exp(arg0)
}
var _Gopurs_Exp = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Exp(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_floor(arg0 float64) float64 {
	return Floor(arg0)
}
var _Gopurs_Floor = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Floor(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_log(arg0 float64) float64 {
	return Log(arg0)
}
var _Gopurs_Log = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Log(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_max(arg0 float64, arg1 float64) float64 {
	return Max(arg0, arg1)
}
var _Gopurs_Max = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Max(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_min(arg0 float64, arg1 float64) float64 {
	return Min(arg0, arg1)
}
var _Gopurs_Min = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Min(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_pow(arg0 float64, arg1 float64) float64 {
	return Pow(arg0, arg1)
}
var _Gopurs_Pow = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Pow(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_remainder(arg0 float64, arg1 float64) float64 {
	return Remainder(arg0, arg1)
}
var _Gopurs_Remainder = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := Remainder(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})
func Call_round(arg0 float64) float64 {
	return Round(arg0)
}
var _Gopurs_Round = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Round(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_sign(arg0 float64) float64 {
	return Sign(arg0)
}
var _Gopurs_Sign = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Sign(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_sin(arg0 float64) float64 {
	return Sin(arg0)
}
var _Gopurs_Sin = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Sin(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_sqrt(arg0 float64) float64 {
	return Sqrt(arg0)
}
var _Gopurs_Sqrt = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Sqrt(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_tan(arg0 float64) float64 {
	return Tan(arg0)
}
var _Gopurs_Tan = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Tan(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_trunc(arg0 float64) float64 {
	return Trunc(arg0)
}
var _Gopurs_Trunc = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := Trunc(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Infinity = gopurs_runtime.Box(Infinity)
var _Gopurs_Nan = gopurs_runtime.Box(Nan)

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

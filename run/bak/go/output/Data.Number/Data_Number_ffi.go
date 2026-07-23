package Data_Number

import (
	"gopurs/output/gopurs_runtime"
	"math"
	"strconv"
)

func getFloat(n gopurs_runtime.Value) float64 {
	return math.Float64frombits(uint64(n.IntVal))
}

var Nan = gopurs_runtime.Float(math.NaN())

var IsNaN = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Bool(math.IsNaN(getFloat(n)))
})

var Infinity = gopurs_runtime.Float(math.Inf(1))

var IsFinite = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	v := getFloat(n)
	return gopurs_runtime.Bool(!math.IsNaN(v) && !math.IsInf(v, 0))
})

var FromStringImpl = gopurs_runtime.Func(func(str gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(isFinite gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
				val, err := strconv.ParseFloat(str.StrVal, 64)
				if err != nil {
					return nothing
				}
				if math.IsNaN(val) || math.IsInf(val, 0) {
					return nothing
				}
				return gopurs_runtime.Apply(just, gopurs_runtime.Float(val))
			})
		})
	})
})

var Abs = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Abs(getFloat(n)))
})

var Acos = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Acos(getFloat(n)))
})

var Asin = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Asin(getFloat(n)))
})

var Atan = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Atan(getFloat(n)))
})

var Atan2 = gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float(math.Atan2(getFloat(y), getFloat(x)))
	})
})

var Ceil = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Ceil(getFloat(n)))
})

var Cos = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Cos(getFloat(n)))
})

var Exp = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Exp(getFloat(n)))
})

var Floor = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Floor(getFloat(n)))
})

var Log = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Log(getFloat(n)))
})

var Max = gopurs_runtime.Func(func(n1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(n2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float(math.Max(getFloat(n1), getFloat(n2)))
	})
})

var Min = gopurs_runtime.Func(func(n1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(n2 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float(math.Min(getFloat(n1), getFloat(n2)))
	})
})

var Pow = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(p gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float(math.Pow(getFloat(n), getFloat(p)))
	})
})

var Remainder = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(m gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Float(math.Mod(getFloat(n), getFloat(m)))
	})
})

var Round = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Round(getFloat(n)))
})

var Sign = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	v := getFloat(n)
	if math.IsNaN(v) || v == 0 {
		return gopurs_runtime.Float(v)
	}
	if v < 0 {
		return gopurs_runtime.Float(-1)
	}
	return gopurs_runtime.Float(1)
})

var Sin = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Sin(getFloat(n)))
})

var Sqrt = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Sqrt(getFloat(n)))
})

var Tan = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Tan(getFloat(n)))
})

var Trunc = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(math.Trunc(getFloat(n)))
})

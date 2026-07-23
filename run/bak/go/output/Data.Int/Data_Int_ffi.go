package Data_Int

import (
	"gopurs/output/gopurs_runtime"
	"math"
	"strconv"
)

func getFloat(n gopurs_runtime.Value) float64 {
	return math.Float64frombits(uint64(n.IntVal))
}

var FromNumberImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
			v := getFloat(n)
			if math.Trunc(v) == v && v >= math.MinInt32 && v <= math.MaxInt32 {
				return gopurs_runtime.Apply(just, gopurs_runtime.Int(int64(int32(v))))
			}
			return nothing
		})
	})
})

var ToNumber = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(float64(int32(n.IntVal)))
})

var FromStringAsImpl = gopurs_runtime.Func(func(just gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(nothing gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(radix gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
				val, err := strconv.ParseInt(s.StrVal, int(radix.IntVal), 32)
				if err != nil {
					return nothing
				}
				return gopurs_runtime.Apply(just, gopurs_runtime.Int(int64(val)))
			})
		})
	})
})

var ToStringAs = gopurs_runtime.Func(func(radix gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(i gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Str(strconv.FormatInt(int64(int32(i.IntVal)), int(radix.IntVal)))
	})
})

var Quot = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		if y.IntVal == 0 {
			return gopurs_runtime.Int(0)
		}
		return gopurs_runtime.Int(int64(int32(x.IntVal / y.IntVal)))
	})
})

var Rem = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		if y.IntVal == 0 {
			return gopurs_runtime.Float(math.NaN())
		}
		return gopurs_runtime.Int(int64(int32(x.IntVal % y.IntVal)))
	})
})

var Pow = gopurs_runtime.Func(func(x gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		res := math.Pow(float64(x.IntVal), float64(y.IntVal))
		return gopurs_runtime.Int(int64(int32(res)))
	})
})

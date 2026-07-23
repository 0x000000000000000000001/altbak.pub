package Data_DateTime_Instant

import (
	"gopurs/output/gopurs_runtime"
	"math"
	"time"
)

func createDateTime(y, m, d, h, mi, s, ms int) time.Time {
	return time.Date(y, time.Month(m+1), d, h, mi, s, ms*1000000, time.UTC)
}

var FromDateTimeImpl = gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(mo gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(h gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(mi gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(s gopurs_runtime.Value) gopurs_runtime.Value {
						return gopurs_runtime.Func(func(ms gopurs_runtime.Value) gopurs_runtime.Value {
							dt := createDateTime(int(y.IntVal), int(mo.IntVal)-1, int(d.IntVal), int(h.IntVal), int(mi.IntVal), int(s.IntVal), int(ms.IntVal))
							return gopurs_runtime.Float(float64(dt.UnixMilli()))
						})
					})
				})
			})
		})
	})
})

var ToDateTimeImpl = gopurs_runtime.Func(func(ctor gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(instant gopurs_runtime.Value) gopurs_runtime.Value {
		ms := math.Float64frombits(uint64(instant.IntVal))
		dt := time.UnixMilli(int64(ms)).UTC()
		
		res := gopurs_runtime.Apply(ctor, gopurs_runtime.Int(int64(dt.Year())))
		res = gopurs_runtime.Apply(res, gopurs_runtime.Int(int64(dt.Month())))
		res = gopurs_runtime.Apply(res, gopurs_runtime.Int(int64(dt.Day())))
		res = gopurs_runtime.Apply(res, gopurs_runtime.Int(int64(dt.Hour())))
		res = gopurs_runtime.Apply(res, gopurs_runtime.Int(int64(dt.Minute())))
		res = gopurs_runtime.Apply(res, gopurs_runtime.Int(int64(dt.Second())))
		res = gopurs_runtime.Apply(res, gopurs_runtime.Int(int64(dt.Nanosecond()/1000000)))
		return res
	})
})

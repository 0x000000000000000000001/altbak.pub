package Data_Date

import (
	"gopurs/output/gopurs_runtime"
	"time"
)

func createDate(y, m, d int) time.Time {
	// m is 0-indexed in JS, but 1-indexed in Go time.Month
	return time.Date(y, time.Month(m+1), d, 0, 0, 0, 0, time.UTC)
}

var CanonicalDateImpl = gopurs_runtime.Func(func(ctor gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(m gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
				date := createDate(int(y.IntVal), int(m.IntVal)-1, int(d.IntVal))
				res := gopurs_runtime.Apply(ctor, gopurs_runtime.Int(int64(date.Year())))
				res = gopurs_runtime.Apply(res, gopurs_runtime.Int(int64(date.Month())))
				res = gopurs_runtime.Apply(res, gopurs_runtime.Int(int64(date.Day())))
				return res
			})
		})
	})
})

var CalcWeekday = gopurs_runtime.Func(func(y gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(m gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(d gopurs_runtime.Value) gopurs_runtime.Value {
			date := createDate(int(y.IntVal), int(m.IntVal)-1, int(d.IntVal))
			return gopurs_runtime.Int(int64(date.Weekday()))
		})
	})
})

var CalcDiff = gopurs_runtime.Func(func(y1 gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(m1 gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(d1 gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(y2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Func(func(m2 gopurs_runtime.Value) gopurs_runtime.Value {
					return gopurs_runtime.Func(func(d2 gopurs_runtime.Value) gopurs_runtime.Value {
						dt1 := createDate(int(y1.IntVal), int(m1.IntVal)-1, int(d1.IntVal))
						dt2 := createDate(int(y2.IntVal), int(m2.IntVal)-1, int(d2.IntVal))
						return gopurs_runtime.Float(float64(dt1.UnixMilli() - dt2.UnixMilli()))
					})
				})
			})
		})
	})
})

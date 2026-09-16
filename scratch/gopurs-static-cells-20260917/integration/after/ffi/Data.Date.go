

import (
	"time"
)

func createDate(y, m, d int) time.Time {
	// m is 0-indexed in JS, but 1-indexed in Go time.Month
	return time.Date(y, time.Month(m+1), d, 0, 0, 0, 0, time.UTC)
}

func CanonicalDateImpl(ctor interface{}, y int, m int, d int) interface{} {
	date := createDate(y, m-1, d)
	
	// ctor is a gopurs_runtime.Value (passed as interface{}) representing a curried function Date -> Int -> Int -> Date
	valCtor := ctor.(gopurs_runtime.Value)
	
	res1 := gopurs_runtime.Apply(valCtor, gopurs_runtime.Int(int64(date.Year())))
	res2 := gopurs_runtime.Apply(res1, gopurs_runtime.Int(int64(date.Month())))
	res3 := gopurs_runtime.Apply(res2, gopurs_runtime.Int(int64(date.Day())))
	
	return res3
}

func CalcWeekday(y int, m int, d int) int {
	date := createDate(y, m-1, d)
	return int(date.Weekday())
}

func CalcDiff(y1 int, m1 int, d1 int, y2 int, m2 int, d2 int) float64 {
	dt1 := createDate(y1, m1-1, d1)
	dt2 := createDate(y2, m2-1, d2)
	return float64(dt1.UnixMilli() - dt2.UnixMilli())
}

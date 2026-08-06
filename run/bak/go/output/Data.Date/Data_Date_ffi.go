package Data_Date

import "gopurs/output/gopurs_runtime"



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


// --- Auto-generated FFI wrappers ---
var _Gopurs_CalcDiff = // TAST: (ADT ["Data","Function","Uncurried","Fn6"] [Int, Int, Int, Int, Int, Int, Number])
gopurs_runtime.Func6(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value, arg5 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_arg4 := gopurs_runtime.Unbox[int](arg4)
	go_arg5 := gopurs_runtime.Unbox[int](arg5)
	go_res := CalcDiff(go_arg0, go_arg1, go_arg2, go_arg3, go_arg4, go_arg5)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_CalcWeekday = // TAST: (ADT ["Data","Function","Uncurried","Fn3"] [Int, Int, Int, Int])
gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_res := CalcWeekday(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_CanonicalDateImpl = // TAST: (ADT ["Data","Function","Uncurried","Fn4"] [(Func [Int, Int, Int] (ADT ["Data","Date","Date"] [])), Int, Int, Int, (ADT ["Data","Date","Date"] [])])
gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_res := CanonicalDateImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
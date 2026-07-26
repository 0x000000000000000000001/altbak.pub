package Data_Date

import "gopurs/output/gopurs_runtime"



import (
	"time"
)

func createDate(y, m, d int) time.Time {
	// m is 0-indexed in JS, but 1-indexed in Go time.Month
	return time.Date(y, time.Month(m+1), d, 0, 0, 0, 0, time.UTC)
}

func CanonicalDateImpl(ctor func(int) func(int) func(int) interface{}, y int, m int, d int) interface{} {
	date := createDate(y, m-1, d)
	return ctor(date.Year())(int(date.Month()))(date.Day())
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

func CalcDiff(y int, m int, d int) int { return 0 }


// --- Auto-generated FFI wrappers ---
func Call_canonicalDateImpl(arg0 func(int) func(int) func(int) interface{}, arg1 int, arg2 int, arg3 int) interface{} {
	return CanonicalDateImpl(arg0, arg1, arg2, arg3)
}
var _Gopurs_CanonicalDateImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) func(int) func(int) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 int) func(int) interface{} {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 int) interface{} {
			return gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
		}
		}
		}
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_arg3 := gopurs_runtime.Unbox[int](arg3)
	go_res := CanonicalDateImpl(go_arg0, go_arg1, go_arg2, go_arg3)
	return gopurs_runtime.Box(go_res)
})
func Call_calcWeekday(arg0 int, arg1 int, arg2 int) int {
	return CalcWeekday(arg0, arg1, arg2)
}
var _Gopurs_CalcWeekday = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_res := CalcWeekday(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})
func Call_calcDiff(arg0 int, arg1 int, arg2 int) int {
	return CalcDiff(arg0, arg1, arg2)
}
var _Gopurs_CalcDiff = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[int](arg0)
	go_arg1 := gopurs_runtime.Unbox[int](arg1)
	go_arg2 := gopurs_runtime.Unbox[int](arg2)
	go_res := CalcDiff(go_arg0, go_arg1, go_arg2)
	return gopurs_runtime.Box(go_res)
})

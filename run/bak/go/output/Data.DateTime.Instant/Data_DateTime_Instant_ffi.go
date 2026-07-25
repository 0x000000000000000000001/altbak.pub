package Data_DateTime_Instant

import "gopurs/output/gopurs_runtime"



import (
	"math"
	"time"
)

func createDateTime(y, m, d, h, mi, s, ms int) time.Time {
	return time.Date(y, time.Month(m+1), d, h, mi, s, ms*1000000, time.UTC)
}

func FromDateTimeImpl(y int, mo int, d int, h int, mi int, s int, ms int) float64 {
	dt := createDateTime(y, mo-1, d, h, mi, s, ms)
	return float64(dt.UnixMilli())
}

func ToDateTimeImpl(ctor func(int) func(int) func(int) func(int) func(int) func(int) func(int) interface{}, instant float64) interface{} {
	dt := time.UnixMilli(int64(instant)).UTC()
	
	return ctor(dt.Year())(int(dt.Month()))(dt.Day())(dt.Hour())(dt.Minute())(dt.Second())(dt.Nanosecond() / 1000000)
}


// --- Auto-generated FFI wrappers ---
func Call_toDateTimeImpl(arg0 func(int) func(int) func(int) func(int) func(int) func(int) func(int) interface{}, arg1 float64) interface{} {
	return ToDateTimeImpl(arg0, arg1)
}
var _Gopurs_ToDateTimeImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0_0 int) func(int) func(int) func(int) func(int) func(int) func(int) interface{} {
			inner_res0 := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0_0))
			return func(p1_0 int) func(int) func(int) func(int) func(int) func(int) interface{} {
			inner_res1 := gopurs_runtime.Apply(inner_res0, gopurs_runtime.Box(p1_0))
			return func(p2_0 int) func(int) func(int) func(int) func(int) interface{} {
			inner_res2 := gopurs_runtime.Apply(inner_res1, gopurs_runtime.Box(p2_0))
			return func(p3_0 int) func(int) func(int) func(int) interface{} {
			inner_res3 := gopurs_runtime.Apply(inner_res2, gopurs_runtime.Box(p3_0))
			return func(p4_0 int) func(int) func(int) interface{} {
			inner_res4 := gopurs_runtime.Apply(inner_res3, gopurs_runtime.Box(p4_0))
			return func(p5_0 int) func(int) interface{} {
			inner_res5 := gopurs_runtime.Apply(inner_res4, gopurs_runtime.Box(p5_0))
			return func(p6_0 int) interface{} {
			return gopurs_runtime.Apply(inner_res5, gopurs_runtime.Box(p6_0))
		}
		}
		}
		}
		}
		}
		}
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := ToDateTimeImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

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

func ToDateTimeImpl(ctor func(int) func(int) func(int) func(int) func(int) func(int) func(int) any, instant float64) any {
	dt := time.UnixMilli(int64(instant)).UTC()
	
	return ctor(dt.Year())(int(dt.Month()))(dt.Day())(dt.Hour())(dt.Minute())(dt.Second())(dt.Nanosecond() / 1000000)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_ToDateTimeImpl = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := func(p0 int) func(int) func(int) func(int) func(int) func(int) func(int) any {
		res := gopurs_runtime.Apply(arg0, gopurs_runtime.Box(p0))
		return gopurs_runtime.Unbox[func(int) func(int) func(int) func(int) func(int) func(int) any](res)
	}
	go_arg1 := gopurs_runtime.Unbox[float64](arg1)
	go_res := ToDateTimeImpl(go_arg0, go_arg1)
	return gopurs_runtime.Box(go_res)
})

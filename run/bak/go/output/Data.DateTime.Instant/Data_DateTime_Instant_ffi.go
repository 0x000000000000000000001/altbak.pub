package Data_DateTime_Instant



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

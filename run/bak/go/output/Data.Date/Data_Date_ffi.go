package Data_Date



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

import (
	"time"
)

func Now(_ interface{}) float64 {
	return float64(time.Now().UnixNano()) / 1e6
}

func GetTimezoneOffset(_ interface{}) float64 {
	_, offset := time.Now().Zone()
	// In JavaScript, getTimezoneOffset() returns the difference in minutes
	// between UTC and local time (e.g. UTC+1 is -60)
	return float64(-offset / 60)
}

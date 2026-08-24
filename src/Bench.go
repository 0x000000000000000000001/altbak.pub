package Bench

import (
	"fmt"
	"time"
)

func BenchNow() float64 {
	return float64(time.Now().UnixNano()) / 1e3
}

func FormatNumber(n float64) string {
	return fmt.Sprintf("%.2f", n)
}

func Opaque(a interface{}) func() interface{} {
	return func() interface{} {
		return a
	}
}

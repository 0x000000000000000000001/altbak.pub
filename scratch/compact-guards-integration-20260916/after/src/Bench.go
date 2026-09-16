package Bench

import (
	"fmt"
	"time"
)

var benchmarkEpoch = time.Now()

func BenchNow() float64 {
	return float64(time.Since(benchmarkEpoch).Nanoseconds()) / 1e3
}

func FormatNumber(n float64) string {
	return fmt.Sprintf("%.2f", n)
}

func Opaque(a interface{}) func() interface{} {
	return func() interface{} {
		return opaqueValue(a)
	}
}

// Keep benchmark inputs opaque to the Go compiler without constraining kernels.
//go:noinline
func opaqueValue(a interface{}) interface{} {
	return a
}

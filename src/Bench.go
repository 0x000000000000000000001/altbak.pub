package Bench

import (
	"fmt"
	"sync/atomic"
	"time"
)

var benchmarkEpoch = time.Now()

func BenchNow() float64 {
	return float64(time.Since(benchmarkEpoch).Nanoseconds()) / 1e3
}

func FormatNumber(n float64) string {
	return fmt.Sprintf("%.6f", n)
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

var benchmarkResult int64

func MeasureBatch(iterations int, expected int, act func() int) float64 {
	result := 0
	start := time.Now()
	for i := 0; i < iterations; i++ {
		result = act()
		atomic.StoreInt64(&benchmarkResult, int64(result))
	}
	elapsed := float64(time.Since(start).Nanoseconds()) / 1e3
	if result != expected {
		panic("Unstable benchmark result")
	}
	return elapsed
}

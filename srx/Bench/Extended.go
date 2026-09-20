package Bench_Extended

import "sync/atomic"

var extendedResult atomic.Value

func ConsumeResult(expected string, result string) func() {
	return func() {
		extendedResult.Store(result)
		if result != expected {
			panic("Unstable extended benchmark result")
		}
	}
}

package Bench_Extended

import "sync/atomic"

var extendedResult atomic.Value

func ConsumeResult(result string) func() {
	return func() {
		extendedResult.Store(result)
	}
}

package ffi_loader

import (
    "fmt"
    "sync/atomic"
    "time"
    . "github.com/purescript-native/go-runtime"
)

var benchmarkEpoch = time.Now()
var benchmarkResult int64

//go:noinline
func benchmarkOpaque(a Any) Any { return a }

func init() {
    exports := Foreign("Bench")
    exports["benchNow"] = func() Any {
        return float64(time.Since(benchmarkEpoch).Nanoseconds()) / 1e3
    }
    exports["formatNumber"] = func(n Any) Any {
        return fmt.Sprintf("%.6f", n.(float64))
    }
    exports["opaque"] = func(a Any) Any {
        return func() Any { return benchmarkOpaque(a) }
    }
    exports["measureBatch"] = func(count Any) Any {
        return func(expected Any) Any {
            return func(action Any) Any {
                return func() Any {
                    act := action.(func() Any)
                    result := 0
                    start := time.Now()
                    for i := 0; i < count.(int); i++ {
                        result = act().(int)
                        atomic.StoreInt64(&benchmarkResult, int64(result))
                    }
                    elapsed := float64(time.Since(start).Nanoseconds()) / 1e3
                    if result != expected.(int) { panic("Unstable benchmark result") }
                    return elapsed / float64(count.(int))
                }
            }
        }
    }
}

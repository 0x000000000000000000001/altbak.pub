package ffi_loader

import (
    "fmt"
    "time"
    . "github.com/purescript-native/go-runtime"
)

var benchmarkEpoch = time.Now()

//go:noinline
func benchmarkOpaque(a Any) Any { return a }

func init() {
    exports := Foreign("Bench")
    exports["benchNow"] = func() Any {
        return float64(time.Since(benchmarkEpoch).Nanoseconds()) / 1e3
    }
    exports["formatNumber"] = func(n Any) Any {
        return fmt.Sprintf("%.2f", n.(float64))
    }
    exports["opaque"] = func(a Any) Any {
        return func() Any { return benchmarkOpaque(a) }
    }
}

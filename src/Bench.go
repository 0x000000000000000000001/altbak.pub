package Bench

import (
	"gopurs/output/gopurs_runtime"
	"time"
	"fmt"
	"math"
)

var BenchNow = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(float64(time.Now().UnixNano()) / 1e3)
})

var Opaque = gopurs_runtime.Func(func(a gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		return a
	})
})

var FormatNumber = gopurs_runtime.Func(func(n gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Str(fmt.Sprintf("%.2f", math.Float64frombits(uint64(n.IntVal))))
})

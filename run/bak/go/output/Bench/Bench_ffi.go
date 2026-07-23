package Bench
import (
	"gopurs/output/gopurs_runtime"
	"time"
)
var BenchNow = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	return gopurs_runtime.Float(float64(time.Now().UnixNano()) / 1e3)
})

package Bench

import "gopurs/output/gopurs_runtime"


import (
	"time"
	"fmt"
)

func BenchNow() float64 {
	return float64(time.Now().UnixNano()) / 1e3
}

func Opaque(a any) func() any {
	return func() any {
		return a
	}
}

func FormatNumber(n float64) string {
	return fmt.Sprintf("%.2f", n)
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_BenchNow = // TAST: (ADT ["Effect","Effect"] [Number])
gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {

	go_res := BenchNow()
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_FormatNumber = // TAST: (Func [Number] String)
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := gopurs_runtime.Unbox[float64](arg0)
	go_res := FormatNumber(go_arg0)
	return gopurs_runtime.Box(go_res)
})
var _Gopurs_Opaque = // TAST: (Func [(TypeVar a)] (ADT ["Effect","Effect"] [(TypeVar a)]))
gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Opaque(go_arg0)
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				inner_res := go_res()
				return gopurs_runtime.Box(inner_res)
			})
})
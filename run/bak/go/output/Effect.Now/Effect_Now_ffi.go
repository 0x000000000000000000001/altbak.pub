package Effect_Now

import "gopurs/output/gopurs_runtime"

import (
	"time"
)

func Now() func() float64 {
	return func() float64 {
		return float64(time.Now().UnixNano()) / 1e6
	}
}

func GetTimezoneOffset() func() float64 {
	return func() float64 {
		_, offset := time.Now().Zone()
		// In JavaScript, getTimezoneOffset() returns the difference in minutes
		// between UTC and local time (e.g. UTC+1 is -60)
		return float64(-offset / 60)
	}
}


// --- Auto-generated FFI wrappers ---
var _Gopurs_Now = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	go_res := Now()
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})
var _Gopurs_GetTimezoneOffset = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
	go_res := GetTimezoneOffset()
	return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			inner_res := go_res()
			return gopurs_runtime.Box(inner_res)
		})
})

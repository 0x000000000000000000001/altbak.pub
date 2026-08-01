package Effect_Now

import "gopurs/output/gopurs_runtime"

import (
	"time"
)

func Now(_ interface{}) float64 {
	return float64(time.Now().UnixNano()) / 1e6
}

func GetTimezoneOffset(_ interface{}) float64 {
	_, offset := time.Now().Zone()
	// In JavaScript, getTimezoneOffset() returns the difference in minutes
	// between UTC and local time (e.g. UTC+1 is -60)
	return float64(-offset / 60)
}


// --- Auto-generated FFI wrappers ---
func Call_now(arg0 interface{}) float64 {
	return Now(arg0)
}
var _Gopurs_Now = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := Now(go_arg0)
	return gopurs_runtime.Box(go_res)
})
func Call_getTimezoneOffset(arg0 interface{}) float64 {
	return GetTimezoneOffset(arg0)
}
var _Gopurs_GetTimezoneOffset = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
	go_arg0 := arg0
	go_res := GetTimezoneOffset(go_arg0)
	return gopurs_runtime.Box(go_res)
})

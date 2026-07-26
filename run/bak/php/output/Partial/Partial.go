package Partial

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_crashWith gopurs_runtime.Value
var once_crashWith sync.Once
func Get_crashWith() gopurs_runtime.Value {
	once_crashWith.Do(func() {
		cache_crashWith = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_crashWith(_dollar__unused_0_box)
})
	})
	return cache_crashWith
}

var cache_crashWith__gopurs_runtime_Value gopurs_runtime.Value
var once_crashWith__gopurs_runtime_Value sync.Once
func Get_crashWith__gopurs_runtime_Value() gopurs_runtime.Value {
	once_crashWith__gopurs_runtime_Value.Do(func() {
		cache_crashWith__gopurs_runtime_Value = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_crashWith__gopurs_runtime_Value(_dollar__unused_0_box)
})
	})
	return cache_crashWith__gopurs_runtime_Value
}

var cache_crash gopurs_runtime.Value
var once_crash sync.Once
func Get_crash() gopurs_runtime.Value {
	once_crash.Do(func() {
		cache_crash = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_crash(_dollar__unused_0_box)
})
	})
	return cache_crash
}

func Call_crashWith(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get__crashWith()
}

func Call_crashWith__gopurs_runtime_Value(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get__crashWith()
}

func Call_crash(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.Apply(Get__crashWith(), gopurs_runtime.Str("Partial.crash: partial function"))
}

func Get__crashWith() gopurs_runtime.Value {
	return _Gopurs__CrashWith
}

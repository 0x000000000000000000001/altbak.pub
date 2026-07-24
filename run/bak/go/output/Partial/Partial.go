package Partial

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var crashWith gopurs_runtime.Value
var once_crashWith sync.Once
func Get_crashWith() gopurs_runtime.Value {
	once_crashWith.Do(func() {
		crashWith = gopurs_runtime.Func(func(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get__crashWith()
}()
})
	})
	return crashWith
}

var crash gopurs_runtime.Value
var once_crash sync.Once
func Get_crash() gopurs_runtime.Value {
	once_crash.Do(func() {
		crash = gopurs_runtime.Func(func(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.Apply(Get__crashWith(), gopurs_runtime.Str("Partial.crash: partial function"))
}()
})
	})
	return crash
}



func Get__crashWith() gopurs_runtime.Value {
	return _Gopurs__CrashWith
}

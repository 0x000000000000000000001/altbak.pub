package Partial_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Partial "gopurs/output/Partial"
)

var cache_unsafePartial gopurs_runtime.Value
var once_unsafePartial sync.Once
func Get_unsafePartial() gopurs_runtime.Value {
	once_unsafePartial.Do(func() {
		cache_unsafePartial = Get__unsafePartial()
	})
	return cache_unsafePartial
}

var cache_unsafeCrashWith gopurs_runtime.Value
var once_unsafeCrashWith sync.Once
func Get_unsafeCrashWith() gopurs_runtime.Value {
	once_unsafeCrashWith.Do(func() {
		cache_unsafeCrashWith = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith
}

func Call_unsafeCrashWith(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}

func Get__unsafePartial() gopurs_runtime.Value {
	return _Gopurs__UnsafePartial
}

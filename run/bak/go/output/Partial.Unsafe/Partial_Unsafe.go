package Partial_Unsafe

import (
	pkg_Partial "gopurs/output/Partial"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

var cache_crashWith__1894115486 gopurs_runtime.Value
var once_crashWith__1894115486 sync.Once
func Get_crashWith__1894115486() gopurs_runtime.Value {
	once_crashWith__1894115486.Do(func() {
		cache_crashWith__1894115486 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_crashWith__1894115486(_dollar__unused_0_box)
})
	})
	return cache_crashWith__1894115486
}

func Call_unsafeCrashWith(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}))
}

func Call_crashWith__1894115486(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Partial.Get__crashWith()
}

func Get__unsafePartial() gopurs_runtime.Value {
	return _Gopurs__UnsafePartial
}

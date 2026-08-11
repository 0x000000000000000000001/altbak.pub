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

var cache_unsafePartial__gopurs_runtime_Value_3178441094 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_3178441094 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_3178441094() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_3178441094.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_3178441094 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_3178441094
}

var cache_unsafePartial__gopurs_runtime_Value_3861213094 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_3861213094 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_3861213094() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_3861213094.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_3861213094 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_3861213094
}

var cache_unsafePartial__gopurs_runtime_Value_3574557895 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_3574557895 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_3574557895() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_3574557895.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_3574557895 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_3574557895
}

var cache_unsafePartial__gopurs_runtime_Value_1306634845 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_1306634845 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_1306634845() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_1306634845.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_1306634845 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_1306634845
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

var cache_unsafeCrashWith__gopurs_runtime_Value_3091512314 gopurs_runtime.Value
var once_unsafeCrashWith__gopurs_runtime_Value_3091512314 sync.Once
func Get_unsafeCrashWith__gopurs_runtime_Value_3091512314() gopurs_runtime.Value {
	once_unsafeCrashWith__gopurs_runtime_Value_3091512314.Do(func() {
		cache_unsafeCrashWith__gopurs_runtime_Value_3091512314 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__gopurs_runtime_Value_3091512314(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__gopurs_runtime_Value_3091512314
}

var cache_unsafeCrashWith__gopurs_runtime_Value_69763299 gopurs_runtime.Value
var once_unsafeCrashWith__gopurs_runtime_Value_69763299 sync.Once
func Get_unsafeCrashWith__gopurs_runtime_Value_69763299() gopurs_runtime.Value {
	once_unsafeCrashWith__gopurs_runtime_Value_69763299.Do(func() {
		cache_unsafeCrashWith__gopurs_runtime_Value_69763299 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__gopurs_runtime_Value_69763299(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__gopurs_runtime_Value_69763299
}

func Call_unsafeCrashWith(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}

func Call_unsafeCrashWith__gopurs_runtime_Value_3091512314(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}

func Call_unsafeCrashWith__gopurs_runtime_Value_69763299(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}

func Get__unsafePartial() gopurs_runtime.Value {
	return _Gopurs__UnsafePartial
}

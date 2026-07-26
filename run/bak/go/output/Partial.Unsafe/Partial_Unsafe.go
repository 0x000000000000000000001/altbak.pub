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

var cache_unsafePartial__gopurs_runtime_Value_3467518839 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_3467518839 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_3467518839() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_3467518839.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_3467518839 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_3467518839
}

var cache_unsafePartial__gopurs_runtime_Value_2932342263 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_2932342263 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_2932342263() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_2932342263.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_2932342263 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_2932342263
}

var cache_unsafePartial__gopurs_runtime_Value_3074941623 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_3074941623 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_3074941623() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_3074941623.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_3074941623 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_3074941623
}

var cache_unsafePartial__gopurs_runtime_Value_1768608695 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_1768608695 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_1768608695() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_1768608695.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_1768608695 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_1768608695
}

var cache_unsafePartial__gopurs_runtime_Value_219008183 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_219008183 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_219008183() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_219008183.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_219008183 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_219008183
}

var cache_unsafePartial__gopurs_runtime_Value_2422275383 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_2422275383 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_2422275383() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_2422275383.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_2422275383 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_2422275383
}

var cache_unsafePartial__gopurs_runtime_Value_3936922103 gopurs_runtime.Value
var once_unsafePartial__gopurs_runtime_Value_3936922103 sync.Once
func Get_unsafePartial__gopurs_runtime_Value_3936922103() gopurs_runtime.Value {
	once_unsafePartial__gopurs_runtime_Value_3936922103.Do(func() {
		cache_unsafePartial__gopurs_runtime_Value_3936922103 = Get__unsafePartial()
	})
	return cache_unsafePartial__gopurs_runtime_Value_3936922103
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

var cache_unsafeCrashWith__gopurs_runtime_Value_3212869717 gopurs_runtime.Value
var once_unsafeCrashWith__gopurs_runtime_Value_3212869717 sync.Once
func Get_unsafeCrashWith__gopurs_runtime_Value_3212869717() gopurs_runtime.Value {
	once_unsafeCrashWith__gopurs_runtime_Value_3212869717.Do(func() {
		cache_unsafeCrashWith__gopurs_runtime_Value_3212869717 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__gopurs_runtime_Value_3212869717(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__gopurs_runtime_Value_3212869717
}

var cache_unsafeCrashWith__gopurs_runtime_Value_1416987000 gopurs_runtime.Value
var once_unsafeCrashWith__gopurs_runtime_Value_1416987000 sync.Once
func Get_unsafeCrashWith__gopurs_runtime_Value_1416987000() gopurs_runtime.Value {
	once_unsafeCrashWith__gopurs_runtime_Value_1416987000.Do(func() {
		cache_unsafeCrashWith__gopurs_runtime_Value_1416987000 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeCrashWith__gopurs_runtime_Value_1416987000(msg_0_box.StrVal())
})
	})
	return cache_unsafeCrashWith__gopurs_runtime_Value_1416987000
}

func Call_unsafeCrashWith(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}

func Call_unsafeCrashWith__gopurs_runtime_Value_3212869717(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}

func Call_unsafeCrashWith__gopurs_runtime_Value_1416987000(msg_0_loop string) gopurs_runtime.Value {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0))
}

func Get__unsafePartial() gopurs_runtime.Value {
	return _Gopurs__UnsafePartial
}

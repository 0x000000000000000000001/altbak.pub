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

var cache_crashWith__func_gopurs_runtime_Value__string__interface___3537417528 gopurs_runtime.Value
var once_crashWith__func_gopurs_runtime_Value__string__interface___3537417528 sync.Once
func Get_crashWith__func_gopurs_runtime_Value__string__interface___3537417528() gopurs_runtime.Value {
	once_crashWith__func_gopurs_runtime_Value__string__interface___3537417528.Do(func() {
		cache_crashWith__func_gopurs_runtime_Value__string__interface___3537417528 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_crashWith__func_gopurs_runtime_Value__string__interface___3537417528(_dollar__unused_0_box)
})
	})
	return cache_crashWith__func_gopurs_runtime_Value__string__interface___3537417528
}

var cache_crash gopurs_runtime.Value
var once_crash sync.Once
func Get_crash() gopurs_runtime.Value {
	once_crash.Do(func() {
		cache_crash = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_crash(_dollar__unused_0_box))
})
	})
	return cache_crash
}

var cache__crashWith gopurs_runtime.Value
var once__crashWith sync.Once
func Get__crashWith() gopurs_runtime.Value {
	once__crashWith.Do(func() {
		cache__crashWith = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(_CrashWith(arg0.StrVal()))
})
	})
	return cache__crashWith
}

func Call_crashWith(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get__crashWith()
}

func Call_crashWith__func_gopurs_runtime_Value__string__interface___3537417528(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return Get__crashWith()
}

func Call_crash(_dollar__unused_0_loop gopurs_runtime.Value) interface{} {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(Get__crashWith(), gopurs_runtime.Str("Partial.crash: partial function")))
}

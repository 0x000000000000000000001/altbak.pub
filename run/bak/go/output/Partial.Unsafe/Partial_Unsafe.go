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
		cache_unsafePartial = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(gopurs_runtime.Value) interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(arg0))
})))
}(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, inner_arg0))
}))
})
	})
	return cache_unsafePartial
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__int64__int64_1048749239 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__int64__int64_1048749239 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__int64__int64_1048749239() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__int64__int64_1048749239.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__int64__int64_1048749239 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 func(gopurs_runtime.Value) int64) int64 {
return gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(inner_arg0(arg0))
})).IntVal
}(func(inner_arg0 gopurs_runtime.Value) int64 {
return gopurs_runtime.Apply(arg0, inner_arg0).IntVal
}))
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__int64__int64_1048749239
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__string__string_2969595063 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__string__string_2969595063 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__string__string_2969595063() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__string__string_2969595063.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__string__string_2969595063 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(func(inner_arg0 func(gopurs_runtime.Value) string) string {
return gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(inner_arg0(arg0))
})).StrVal()
}(func(inner_arg0 gopurs_runtime.Value) string {
return gopurs_runtime.Apply(arg0, inner_arg0).StrVal()
}))
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__string__string_2969595063
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__arrint64__arrint64_3100370487 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__arrint64__arrint64_3100370487 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__arrint64__arrint64_3100370487() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__arrint64__arrint64_3100370487.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__arrint64__arrint64_3100370487 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func(inner_arg0 func(gopurs_runtime.Value) []int64) []int64 {
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := inner_arg0(arg0)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}(func(inner_arg0 gopurs_runtime.Value) []int64 {
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(arg0, inner_arg0).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
})
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__arrint64__arrint64_3100370487
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__arrstring__arrstring_160877783 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__arrstring__arrstring_160877783 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__arrstring__arrstring_160877783() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__arrstring__arrstring_160877783.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__arrstring__arrstring_160877783 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func(inner_arg0 func(gopurs_runtime.Value) []string) []string {
return func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := inner_arg0(arg0)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
}(func(inner_arg0 gopurs_runtime.Value) []string {
return func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(arg0, inner_arg0).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
})
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__arrstring__arrstring_160877783
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____arrinterface___3467518839 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____arrinterface___3467518839 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____arrinterface___3467518839() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____arrinterface___3467518839.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____arrinterface___3467518839 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := func(inner_arg0 func(gopurs_runtime.Value) []interface{}) []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := inner_arg0(arg0)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
}(func(inner_arg0 gopurs_runtime.Value) []interface{} {
return func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(arg0, inner_arg0).UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()
})
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____arrinterface___3467518839
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___2932342263 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___2932342263 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___2932342263() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___2932342263.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___2932342263 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(gopurs_runtime.Value) interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(arg0))
})))
}(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, inner_arg0))
}))
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___2932342263
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___3074941623 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___3074941623 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___3074941623() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___3074941623.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___3074941623 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(gopurs_runtime.Value) interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(arg0))
})))
}(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, inner_arg0))
}))
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___3074941623
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___1768608695 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___1768608695 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___1768608695() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___1768608695.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___1768608695 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(gopurs_runtime.Value) interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(arg0))
})))
}(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, inner_arg0))
}))
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___1768608695
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___219008183 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___219008183 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___219008183() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___219008183.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___219008183 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(gopurs_runtime.Value) interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(Get__unsafePartial(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(arg0))
})))
}(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, inner_arg0))
}))
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__interface____interface___219008183
}

var cache_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____interface____arrinterface____interface___1004292151 gopurs_runtime.Value
var once_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____interface____arrinterface____interface___1004292151 sync.Once
func Get_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____interface____arrinterface____interface___1004292151() gopurs_runtime.Value {
	once_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____interface____arrinterface____interface___1004292151.Do(func() {
		cache_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____interface____arrinterface____interface___1004292151 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(gopurs_runtime.Value, []interface{}) interface{}, inner_arg1 []interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(Get__unsafePartial(), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(arg0, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
}), func() gopurs_runtime.Value {
					arr := inner_arg1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}(func(inner_arg0 gopurs_runtime.Value, inner_arg1 []interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(arg0, inner_arg0, func() gopurs_runtime.Value {
					arr := inner_arg1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}()))
})
	})
	return cache_unsafePartial__func_func_gopurs_runtime_Value__arrinterface____interface____arrinterface____interface___1004292151
}

var cache_unsafeCrashWith gopurs_runtime.Value
var once_unsafeCrashWith sync.Once
func Get_unsafeCrashWith() gopurs_runtime.Value {
	once_unsafeCrashWith.Do(func() {
		cache_unsafeCrashWith = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unsafeCrashWith(msg_0_box.StrVal()))
})
	})
	return cache_unsafeCrashWith
}

var cache_unsafeCrashWith__func_string__interface___3212869717 gopurs_runtime.Value
var once_unsafeCrashWith__func_string__interface___3212869717 sync.Once
func Get_unsafeCrashWith__func_string__interface___3212869717() gopurs_runtime.Value {
	once_unsafeCrashWith__func_string__interface___3212869717.Do(func() {
		cache_unsafeCrashWith__func_string__interface___3212869717 = gopurs_runtime.Func(func(msg_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unsafeCrashWith__func_string__interface___3212869717(msg_0_box.StrVal()))
})
	})
	return cache_unsafeCrashWith__func_string__interface___3212869717
}

var cache__unsafePartial gopurs_runtime.Value
var once__unsafePartial sync.Once
func Get__unsafePartial() gopurs_runtime.Value {
	once__unsafePartial.Do(func() {
		cache__unsafePartial = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(_UnsafePartial(gopurs_runtime.UnboxAny(arg0)))
})
	})
	return cache__unsafePartial
}

func Call_unsafeCrashWith(msg_0_loop string) interface{} {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0)))
}

func Call_unsafeCrashWith__func_string__interface___3212869717(msg_0_loop string) interface{} {
var msg_0 string = msg_0_loop
_ = msg_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), gopurs_runtime.Str(msg_0)))
}

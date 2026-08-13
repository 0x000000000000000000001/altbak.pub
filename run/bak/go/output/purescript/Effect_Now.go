package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Effect_Now_nowTime gopurs_runtime.Value
var once_Effect_Now_nowTime sync.Once
func Get_Effect_Now_nowTime() gopurs_runtime.Value {
	once_Effect_Now_nowTime.Do(func() {
		cache_Effect_Now_nowTime = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 922918650, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(gopurs_runtime.Apply(Get_Data_DateTime_Instant_toDateTime(), x_0).UnsafePtr).V1)}
})
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(Get_Effect_Now_now(), gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Apply(__local_var_0_0, __local_var_1_1)
})
	})
	return cache_Effect_Now_nowTime
}

var cache_Effect_Now_nowDateTime gopurs_runtime.Value
var once_Effect_Now_nowDateTime sync.Once
func Get_Effect_Now_nowDateTime() gopurs_runtime.Value {
	once_Effect_Now_nowDateTime.Do(func() {
		cache_Effect_Now_nowDateTime = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_0_0 := Get_Data_DateTime_Instant_toDateTime()
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(Get_Effect_Now_now(), gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Apply(__local_var_0_0, __local_var_1_1)
})
	})
	return cache_Effect_Now_nowDateTime
}

var cache_Effect_Now_nowDate gopurs_runtime.Value
var once_Effect_Now_nowDate sync.Once
func Get_Effect_Now_nowDate() gopurs_runtime.Value {
	once_Effect_Now_nowDate.Do(func() {
		cache_Effect_Now_nowDate = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer((*Constructor_Data_DateTime_DateTime)(gopurs_runtime.Apply(Get_Data_DateTime_Instant_toDateTime(), x_0).UnsafePtr).V0)}
})
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(Get_Effect_Now_now(), gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Apply(__local_var_0_0, __local_var_1_1)
})
	})
	return cache_Effect_Now_nowDate
}



func Get_Effect_Now_getTimezoneOffset() gopurs_runtime.Value {
	return _Gopurs_Effect_Now_GetTimezoneOffset
}

func Get_Effect_Now_now() gopurs_runtime.Value {
	return _Gopurs_Effect_Now_Now
}

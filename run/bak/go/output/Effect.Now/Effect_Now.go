package Effect_Now

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect "gopurs/output/Effect"
	pkg_Data_DateTime "gopurs/output/Data.DateTime"
	pkg_Data_DateTime_Instant "gopurs/output/Data.DateTime.Instant"
)

var cache_nowTime gopurs_runtime.Value
var once_nowTime sync.Once
func Get_nowTime() gopurs_runtime.Value {
	once_nowTime.Do(func() {
		cache_nowTime = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(gopurs_runtime.Apply(pkg_Data_DateTime_Instant.Get_toDateTime(), x_0).UnsafePtr).V1
}), Get_now())
	})
	return cache_nowTime
}

var cache_nowDateTime gopurs_runtime.Value
var once_nowDateTime sync.Once
func Get_nowDateTime() gopurs_runtime.Value {
	once_nowDateTime.Do(func() {
		cache_nowDateTime = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), pkg_Data_DateTime_Instant.Get_toDateTime(), Get_now())
	})
	return cache_nowDateTime
}

var cache_nowDate gopurs_runtime.Value
var once_nowDate sync.Once
func Get_nowDate() gopurs_runtime.Value {
	once_nowDate.Do(func() {
		cache_nowDate = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_functorEffect(), "map"), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_DateTime.Data_Data_DateTime_DateTime)(gopurs_runtime.Apply(pkg_Data_DateTime_Instant.Get_toDateTime(), x_0).UnsafePtr).V0
}), Get_now())
	})
	return cache_nowDate
}



func Get_getTimezoneOffset() gopurs_runtime.Value {
	return _Gopurs_GetTimezoneOffset
}

func Get_now() gopurs_runtime.Value {
	return _Gopurs_Now
}

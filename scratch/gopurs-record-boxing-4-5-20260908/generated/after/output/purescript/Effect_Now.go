package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Effect_Now_nowTime gopurs_runtime.Value
var once_Effect_Now_nowTime sync.Once
func Get_Effect_Now_nowTime() gopurs_runtime.Value {
	once_Effect_Now_nowTime.Do(func() {
		cache_Effect_Now_nowTime = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), Get_Effect_Now_now(), Get_Effect_Now_now())
	})
	return cache_Effect_Now_nowTime
}

var cache_Effect_Now_nowDateTime gopurs_runtime.Value
var once_Effect_Now_nowDateTime sync.Once
func Get_Effect_Now_nowDateTime() gopurs_runtime.Value {
	once_Effect_Now_nowDateTime.Do(func() {
		cache_Effect_Now_nowDateTime = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), Get_Data_DateTime_Instant_toDateTime(), Get_Effect_Now_now())
	})
	return cache_Effect_Now_nowDateTime
}

var cache_Effect_Now_nowDate gopurs_runtime.Value
var once_Effect_Now_nowDate sync.Once
func Get_Effect_Now_nowDate() gopurs_runtime.Value {
	once_Effect_Now_nowDate.Do(func() {
		cache_Effect_Now_nowDate = gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), Get_Effect_Now_now(), Get_Effect_Now_now())
	})
	return cache_Effect_Now_nowDate
}



func Get_Effect_Now_getTimezoneOffset() gopurs_runtime.Value {
	return _Gopurs_Effect_Now_GetTimezoneOffset
}

func Get_Effect_Now_now() gopurs_runtime.Value {
	return _Gopurs_Effect_Now_Now
}

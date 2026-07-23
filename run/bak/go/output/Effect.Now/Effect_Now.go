package Effect_Now

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_DateTime_Instant "gopurs/output/Data.DateTime.Instant"
)

var nowTime gopurs_runtime.Value
var once_nowTime sync.Once
func Get_nowTime() gopurs_runtime.Value {
	once_nowTime.Do(func() {
		nowTime = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_0_0 := gopurs_runtime.Apply(Get_now(), gopurs_runtime.Value{})
_ = a_prime_0_0
return gopurs_runtime.ConstructorGet(gopurs_runtime.Apply(pkg_Data_DateTime_Instant.Get_toDateTime(), a_prime_0_0), 1)
})
	})
	return nowTime
}

var nowDateTime gopurs_runtime.Value
var once_nowDateTime sync.Once
func Get_nowDateTime() gopurs_runtime.Value {
	once_nowDateTime.Do(func() {
		nowDateTime = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_0_0 := gopurs_runtime.Apply(Get_now(), gopurs_runtime.Value{})
_ = a_prime_0_0
return gopurs_runtime.Apply(pkg_Data_DateTime_Instant.Get_toDateTime(), a_prime_0_0)
})
	})
	return nowDateTime
}

var nowDate gopurs_runtime.Value
var once_nowDate sync.Once
func Get_nowDate() gopurs_runtime.Value {
	once_nowDate.Do(func() {
		nowDate = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_0_0 := gopurs_runtime.Apply(Get_now(), gopurs_runtime.Value{})
_ = a_prime_0_0
return gopurs_runtime.ConstructorGet(gopurs_runtime.Apply(pkg_Data_DateTime_Instant.Get_toDateTime(), a_prime_0_0), 0)
})
	})
	return nowDate
}

func Get_getTimezoneOffset() gopurs_runtime.Value {
	return _Gopurs_GetTimezoneOffset
}

func Get_now() gopurs_runtime.Value {
	return _Gopurs_Now
}

package Data_Time_Component_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Enum_Gen "gopurs/output/Data.Enum.Gen"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
)

var genSecond gopurs_runtime.Value
var once_genSecond sync.Once
func Get_genSecond() gopurs_runtime.Value {
	once_genSecond.Do(func() {
		genSecond = gopurs_runtime.Func(func(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumSecond())
}()
})
	})
	return genSecond
}

var genMinute gopurs_runtime.Value
var once_genMinute sync.Once
func Get_genMinute() gopurs_runtime.Value {
	once_genMinute.Do(func() {
		genMinute = gopurs_runtime.Func(func(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumMinute())
}()
})
	})
	return genMinute
}

var genMillisecond gopurs_runtime.Value
var once_genMillisecond sync.Once
func Get_genMillisecond() gopurs_runtime.Value {
	once_genMillisecond.Do(func() {
		genMillisecond = gopurs_runtime.Func(func(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumMillisecond())
}()
})
	})
	return genMillisecond
}

var genHour gopurs_runtime.Value
var once_genHour sync.Once
func Get_genHour() gopurs_runtime.Value {
	once_genHour.Do(func() {
		genHour = gopurs_runtime.Func(func(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumHour())
}()
})
	})
	return genHour
}





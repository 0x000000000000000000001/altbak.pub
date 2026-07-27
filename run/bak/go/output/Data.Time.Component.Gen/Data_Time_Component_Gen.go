package Data_Time_Component_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Enum_Gen "gopurs/output/Data.Enum.Gen"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
)

var cache_genSecond gopurs_runtime.Value
var once_genSecond sync.Once
func Get_genSecond() gopurs_runtime.Value {
	once_genSecond.Do(func() {
		cache_genSecond = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genSecond(dictMonadGen_0_box))
})
	})
	return cache_genSecond
}

var cache_genMinute gopurs_runtime.Value
var once_genMinute sync.Once
func Get_genMinute() gopurs_runtime.Value {
	once_genMinute.Do(func() {
		cache_genMinute = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genMinute(dictMonadGen_0_box))
})
	})
	return cache_genMinute
}

var cache_genMillisecond gopurs_runtime.Value
var once_genMillisecond sync.Once
func Get_genMillisecond() gopurs_runtime.Value {
	once_genMillisecond.Do(func() {
		cache_genMillisecond = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genMillisecond(dictMonadGen_0_box))
})
	})
	return cache_genMillisecond
}

var cache_genHour gopurs_runtime.Value
var once_genHour sync.Once
func Get_genHour() gopurs_runtime.Value {
	once_genHour.Do(func() {
		cache_genHour = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_genHour(dictMonadGen_0_box))
})
	})
	return cache_genHour
}

func Call_genSecond(dictMonadGen_0_loop gopurs_runtime.Value) interface{} {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumSecond()))
}

func Call_genMinute(dictMonadGen_0_loop gopurs_runtime.Value) interface{} {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumMinute()))
}

func Call_genMillisecond(dictMonadGen_0_loop gopurs_runtime.Value) interface{} {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumMillisecond()))
}

func Call_genHour(dictMonadGen_0_loop gopurs_runtime.Value) interface{} {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumHour()))
}

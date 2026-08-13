package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Time_Component_Gen_genSecond gopurs_runtime.Value
var once_Data_Time_Component_Gen_genSecond sync.Once
func Get_Data_Time_Component_Gen_genSecond() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genSecond.Do(func() {
		cache_Data_Time_Component_Gen_genSecond = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Gen_genSecond(dictMonadGen_0_box)
})
	})
	return cache_Data_Time_Component_Gen_genSecond
}

var cache_Data_Time_Component_Gen_genMinute gopurs_runtime.Value
var once_Data_Time_Component_Gen_genMinute sync.Once
func Get_Data_Time_Component_Gen_genMinute() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genMinute.Do(func() {
		cache_Data_Time_Component_Gen_genMinute = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Gen_genMinute(dictMonadGen_0_box)
})
	})
	return cache_Data_Time_Component_Gen_genMinute
}

var cache_Data_Time_Component_Gen_genMillisecond gopurs_runtime.Value
var once_Data_Time_Component_Gen_genMillisecond sync.Once
func Get_Data_Time_Component_Gen_genMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genMillisecond.Do(func() {
		cache_Data_Time_Component_Gen_genMillisecond = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Gen_genMillisecond(dictMonadGen_0_box)
})
	})
	return cache_Data_Time_Component_Gen_genMillisecond
}

var cache_Data_Time_Component_Gen_genHour gopurs_runtime.Value
var once_Data_Time_Component_Gen_genHour sync.Once
func Get_Data_Time_Component_Gen_genHour() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genHour.Do(func() {
		cache_Data_Time_Component_Gen_genHour = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Gen_genHour(dictMonadGen_0_box)
})
	})
	return cache_Data_Time_Component_Gen_genHour
}

func Call_Data_Time_Component_Gen_genSecond(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, Get_Data_Time_Component_boundedEnumSecond())
}

func Call_Data_Time_Component_Gen_genMinute(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, Get_Data_Time_Component_boundedEnumMinute())
}

func Call_Data_Time_Component_Gen_genMillisecond(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, Get_Data_Time_Component_boundedEnumMillisecond())
}

func Call_Data_Time_Component_Gen_genHour(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, Get_Data_Time_Component_boundedEnumHour())
}



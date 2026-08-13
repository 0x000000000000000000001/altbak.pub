package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Date_Component_Gen_toEnum gopurs_runtime.Value
var once_Data_Date_Component_Gen_toEnum sync.Once
func Get_Data_Date_Component_Gen_toEnum() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_toEnum.Do(func() {
		cache_Data_Date_Component_Gen_toEnum = gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumYear(), "toEnum")
	})
	return cache_Data_Date_Component_Gen_toEnum
}

var cache_Data_Date_Component_Gen_genYear gopurs_runtime.Value
var once_Data_Date_Component_Gen_genYear sync.Once
func Get_Data_Date_Component_Gen_genYear() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genYear.Do(func() {
		cache_Data_Date_Component_Gen_genYear = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_Component_Gen_genYear(dictMonadGen_0_box)
})
	})
	return cache_Data_Date_Component_Gen_genYear
}

var cache_Data_Date_Component_Gen_genWeekday gopurs_runtime.Value
var once_Data_Date_Component_Gen_genWeekday sync.Once
func Get_Data_Date_Component_Gen_genWeekday() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genWeekday.Do(func() {
		cache_Data_Date_Component_Gen_genWeekday = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_Component_Gen_genWeekday(dictMonadGen_0_box)
})
	})
	return cache_Data_Date_Component_Gen_genWeekday
}

var cache_Data_Date_Component_Gen_genMonth gopurs_runtime.Value
var once_Data_Date_Component_Gen_genMonth sync.Once
func Get_Data_Date_Component_Gen_genMonth() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genMonth.Do(func() {
		cache_Data_Date_Component_Gen_genMonth = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_Component_Gen_genMonth(dictMonadGen_0_box)
})
	})
	return cache_Data_Date_Component_Gen_genMonth
}

var cache_Data_Date_Component_Gen_genDay gopurs_runtime.Value
var once_Data_Date_Component_Gen_genDay sync.Once
func Get_Data_Date_Component_Gen_genDay() gopurs_runtime.Value {
	once_Data_Date_Component_Gen_genDay.Do(func() {
		cache_Data_Date_Component_Gen_genDay = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Date_Component_Gen_genDay(dictMonadGen_0_box)
})
	})
	return cache_Data_Date_Component_Gen_genDay
}

func Call_Data_Date_Component_Gen_genYear(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__1577979644()
}))
_ = __local_var_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Date_Component_boundedEnumYear(), "toEnum"), x_2))
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(1900), gopurs_runtime.Int(2100)))
}

func Call_Data_Date_Component_Gen_genWeekday(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, Get_Data_Date_Component_boundedEnumWeekday())
}

func Call_Data_Date_Component_Gen_genMonth(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, Get_Data_Date_Component_boundedEnumMonth())
}

func Call_Data_Date_Component_Gen_genDay(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(Get_Data_Enum_Gen_genBoundedEnum(), dictMonadGen_0, Get_Data_Date_Component_boundedEnumDay())
}



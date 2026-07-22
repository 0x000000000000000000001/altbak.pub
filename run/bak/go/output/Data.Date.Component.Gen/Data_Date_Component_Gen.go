package Data_Date_Component_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Enum_Gen "gopurs/output/Data.Enum.Gen"
)

var fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		fromJust = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t0 = v_0.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
	})
	return fromJust
}

var genYear gopurs_runtime.Value
var once_genYear sync.Once
func Get_genYear() gopurs_runtime.Value {
	once_genYear.Do(func() {
		genYear = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_fromJust()
}))
return gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Apply(pkg_Data_Date_Component.Get_boundedEnumYear().PtrVal.(map[string]gopurs_runtime.Value)["toEnum"], x_2))
})), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(1900)), gopurs_runtime.Int(2100)))
})
	})
	return genYear
}

var genWeekday gopurs_runtime.Value
var once_genWeekday sync.Once
func Get_genWeekday() gopurs_runtime.Value {
	once_genWeekday.Do(func() {
		genWeekday = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0), pkg_Data_Date_Component.Get_boundedEnumWeekday())
})
	})
	return genWeekday
}

var genMonth gopurs_runtime.Value
var once_genMonth sync.Once
func Get_genMonth() gopurs_runtime.Value {
	once_genMonth.Do(func() {
		genMonth = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0), pkg_Data_Date_Component.Get_boundedEnumMonth())
})
	})
	return genMonth
}

var genDay gopurs_runtime.Value
var once_genDay sync.Once
func Get_genDay() gopurs_runtime.Value {
	once_genDay.Do(func() {
		genDay = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0), pkg_Data_Date_Component.Get_boundedEnumDay())
})
	})
	return genDay
}



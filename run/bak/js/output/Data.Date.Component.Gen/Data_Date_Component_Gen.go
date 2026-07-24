package Data_Date_Component_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Enum_Gen "gopurs/output/Data.Enum.Gen"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
)

var genYear gopurs_runtime.Value
var once_genYear sync.Once
func Get_genYear() gopurs_runtime.Value {
	once_genYear.Do(func() {
		genYear = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if x_1.IntVal >= -271820 && x_1.IntVal <= 275759 {
__t0 = x_1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(1900), gopurs_runtime.Int(2100)))
})
	})
	return genYear
}

var genWeekday gopurs_runtime.Value
var once_genWeekday sync.Once
func Get_genWeekday() gopurs_runtime.Value {
	once_genWeekday.Do(func() {
		genWeekday = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Date_Component.Get_boundedEnumWeekday())
})
	})
	return genWeekday
}

var genMonth gopurs_runtime.Value
var once_genMonth sync.Once
func Get_genMonth() gopurs_runtime.Value {
	once_genMonth.Do(func() {
		genMonth = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Date_Component.Get_boundedEnumMonth())
})
	})
	return genMonth
}

var genDay gopurs_runtime.Value
var once_genDay sync.Once
func Get_genDay() gopurs_runtime.Value {
	once_genDay.Do(func() {
		genDay = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Date_Component.Get_boundedEnumDay())
})
	})
	return genDay
}





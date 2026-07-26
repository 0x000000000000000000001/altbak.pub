package Data_Date_Component_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Enum_Gen "gopurs/output/Data.Enum.Gen"
)

var cache_fromJust gopurs_runtime.Value
var once_fromJust sync.Once
func Get_fromJust() gopurs_runtime.Value {
	once_fromJust.Do(func() {
		cache_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust(v_0_box)
})
	})
	return cache_fromJust
}

var cache_genYear gopurs_runtime.Value
var once_genYear sync.Once
func Get_genYear() gopurs_runtime.Value {
	once_genYear.Do(func() {
		cache_genYear = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genYear(dictMonadGen_0_box)
})
	})
	return cache_genYear
}

var cache_genWeekday gopurs_runtime.Value
var once_genWeekday sync.Once
func Get_genWeekday() gopurs_runtime.Value {
	once_genWeekday.Do(func() {
		cache_genWeekday = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genWeekday(dictMonadGen_0_box)
})
	})
	return cache_genWeekday
}

var cache_genMonth gopurs_runtime.Value
var once_genMonth sync.Once
func Get_genMonth() gopurs_runtime.Value {
	once_genMonth.Do(func() {
		cache_genMonth = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genMonth(dictMonadGen_0_box)
})
	})
	return cache_genMonth
}

var cache_genDay gopurs_runtime.Value
var once_genDay sync.Once
func Get_genDay() gopurs_runtime.Value {
	once_genDay.Do(func() {
		cache_genDay = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDay(dictMonadGen_0_box)
})
	})
	return cache_genDay
}

func Call_fromJust(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_genYear(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0_NOT_FOUND"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(pkg_Control_Semigroupoid.Get_composeImpl(), Get_fromJust(), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedEnumYear(), "toEnum")), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData5)(dictMonadGen_0.UnsafePtr)).V2, gopurs_runtime.Int(1900), gopurs_runtime.Int(2100)))
}

func Call_genWeekday(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Date_Component.Get_boundedEnumWeekday())
}

func Call_genMonth(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Date_Component.Get_boundedEnumMonth())
}

func Call_genDay(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
return gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Date_Component.Get_boundedEnumDay())
}



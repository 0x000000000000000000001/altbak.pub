package Data_Date_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Date_Component_Gen "gopurs/output/Data.Date.Component.Gen"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
)

var cache_genDate gopurs_runtime.Value
var once_genDate sync.Once
func Get_genDate() gopurs_runtime.Value {
	once_genDate.Do(func() {
		cache_genDate = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDate(dictMonadGen_0_box)
})
	})
	return cache_genDate
}

func Call_genDate(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(pkg_Data_Date_Component_Gen.Get_genYear(), dictMonadGen_0), gopurs_runtime.Func(func(year_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Date.Get_isLeapYear(), year_3).IntVal) != (0) {
__t2 = gopurs_runtime.Int(365)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Int(364)
}
end_branch_2:
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), x_4)
}), gopurs_runtime.Apply2(((*gopurs_runtime.RecordData5)(dictMonadGen_0.UnsafePtr)).V2, gopurs_runtime.Int(0), __t2)), gopurs_runtime.Func(func(days_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply3(pkg_Data_Date.Get_exactDate(), year_3, gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "bottom")), gopurs_runtime.Func(func(janFirst_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(pkg_Data_Date.Get_adjust(), days_4, janFirst_5)
}))
_ = __local_var_5_3
var __t4 gopurs_runtime.Value
{
if (__local_var_5_3.Type == 9 && __local_var_5_3.IntVal == 930809136) {
__t4 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_5_3.UnsafePtr).V0
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t4)
}))
}))
}



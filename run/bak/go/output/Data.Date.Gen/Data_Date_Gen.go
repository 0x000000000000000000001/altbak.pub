package Data_Date_Gen

import (
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Date_Component_Gen "gopurs/output/Data.Date.Component.Gen"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
pure_3_2 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_3_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply(pkg_Data_Date_Component_Gen.Get_genYear(), dictMonadGen_0), gopurs_runtime.Func(func(year_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Date.Get_isLeapYear(), year_4).IntVal) != (0) {
__t3 = gopurs_runtime.Int(365)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Int(364)
}
end_branch_3:
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), x_5)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), __t3)), gopurs_runtime.Func(func(days_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply3(pkg_Data_Date.Get_exactDate(), year_4, gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "bottom")), gopurs_runtime.Func(func(janFirst_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Apply2(pkg_Data_Date.Get_adjust(), days_5, janFirst_6).UnsafePtr))}
}))
_ = __local_var_6_4
var __t5 gopurs_runtime.Value
{
if (__local_var_6_4.Type == 9 && __local_var_6_4.IntVal == 930809136 && __local_var_6_4.UnsafePtr != nil) {
__t5 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_6_4.UnsafePtr).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Apply(pure_3_2, __t5)
}))
}))
}



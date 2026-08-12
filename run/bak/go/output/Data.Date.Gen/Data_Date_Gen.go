package Data_Date_Gen

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Date_Component "gopurs/output/Data.Date.Component"
	pkg_Data_Date_Component_Gen "gopurs/output/Data.Date.Component.Gen"
	pkg_Data_Functor "gopurs/output/Data.Functor"
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
Bind1_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_2_1
Functor0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_3_2
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply(pkg_Data_Date_Component_Gen.Get_genYear(), dictMonadGen_0), gopurs_runtime.Func(func(year_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Date.Get_isLeapYear(), year_5).IntVal) != (0) {
__t4 = gopurs_runtime.Int(365)
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Int(364)
}
end_branch_4:
return gopurs_runtime.Apply2(Bind1_2_1.V1, gopurs_runtime.Apply2(Functor0_3_2.V0, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), x_6)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), __t4)), gopurs_runtime.Func(func(days_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply3(pkg_Data_Date.Get_exactDate(), year_5, gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedMonth(), "bottom"), gopurs_runtime.RecordGet(pkg_Data_Date_Component.Get_boundedDay(), "bottom")), gopurs_runtime.Func(func(janFirst_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Date.Constructor_Date]](gopurs_runtime.Apply2(pkg_Data_Date.Get_adjust(), days_6, janFirst_7)))}
}))
_ = __local_var_7_5
var __t6 gopurs_runtime.Value
{
if (__local_var_7_5.Type == 9 && __local_var_7_5.IntVal == 930809136 && __local_var_7_5.UnsafePtr != nil) {
__t6 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_7_5.UnsafePtr).V0
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Apply(pure_4_3, __t6)
}))
}))
}



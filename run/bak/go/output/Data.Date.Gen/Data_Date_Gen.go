package Data_Date_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Int "gopurs/output/Data.Int"
)

var genDate gopurs_runtime.Value
var once_genDate sync.Once
func Get_genDate() gopurs_runtime.Value {
	once_genDate.Do(func() {
		genDate = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
Bind1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.Bool(x_3.IntVal >= gopurs_runtime.Int(-271820).IntVal).IntVal != 0 && gopurs_runtime.Bool(x_3.IntVal <= gopurs_runtime.Int(275759).IntVal).IntVal != 0)).IntVal != 0 {
__t2 = x_3
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(1900), gopurs_runtime.Int(2100))), gopurs_runtime.Func(func(year_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Date.Get_isLeapYear(), year_3)).IntVal != 0 {
__t3 = gopurs_runtime.Int(365)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Int(364)
}
end_branch_3:
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Bind1_2_1, "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_1, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), x_4)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), __t3)), gopurs_runtime.Func(func(days_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply3(pkg_Data_Date.Get_exactDate(), year_3, gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("January")), gopurs_runtime.Int(1))
_ = __local_var_5_4
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_5_4, "_tag").StrVal == "Just")).IntVal != 0 {
__t6 = gopurs_runtime.Apply2(pkg_Data_Date.Get_adjust(), days_4, gopurs_runtime.RecordGet(__local_var_5_4, "value0"))
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_5_4, "_tag").StrVal == "Nothing")).IntVal != 0 {
__t6 = gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing"))
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
__local_var_6_5 := __t6
_ = __local_var_6_5
var __t7 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(gopurs_runtime.RecordGet(__local_var_6_5, "_tag").StrVal == "Just")).IntVal != 0 {
__t7 = gopurs_runtime.RecordGet(__local_var_6_5, "value0")
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), __t7)
}))
}))
})
	})
	return genDate
}



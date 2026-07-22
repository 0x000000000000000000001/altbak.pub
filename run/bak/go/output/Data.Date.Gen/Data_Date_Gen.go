package Data_Date_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Date_Component_Gen "gopurs/output/Data.Date.Component.Gen"
	pkg_Data_Date "gopurs/output/Data.Date"
	pkg_Data_Int "gopurs/output/Data.Int"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
)

var genDate gopurs_runtime.Value
var once_genDate sync.Once
func Get_genDate() gopurs_runtime.Value {
	once_genDate.Do(func() {
		genDate = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
Monad0_1_0 := gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{})
Bind1_2_1 := gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{})
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(pkg_Data_Date_Component_Gen.Get_genYear(), dictMonadGen_0)), gopurs_runtime.Func(func(year_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Apply(pkg_Data_Date.Get_isLeapYear(), year_3)).IntVal != 0 {
__t2 = gopurs_runtime.Int(365)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Int(364)
}
end_branch_2:
return gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["bind"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Bind1_2_1.PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Int.Get_toNumber(), x_4)
})), gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["chooseInt"], gopurs_runtime.Int(0)), __t2))), gopurs_runtime.Func(func(days_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(Monad0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Applicative0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["pure"], gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Date.Get_exactDate(), year_3), gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("January")})), gopurs_runtime.Int(1))
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_6_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t5 = gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Date.Get_adjust(), days_4), __local_var_6_3.PtrVal.(map[string]gopurs_runtime.Value)["value0"])
goto end_branch_5
} else {

}
}
{
if (gopurs_runtime.Bool(__local_var_6_3.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Nothing")).IntVal != 0 {
__t5 = gopurs_runtime.Record(map[string]gopurs_runtime.Value{"_tag": gopurs_runtime.Str("Nothing")})
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__local_var_7_4 := __t5
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(__local_var_7_4.PtrVal.(map[string]gopurs_runtime.Value)["_tag"].StrVal == "Just")).IntVal != 0 {
__t6 = __local_var_7_4.PtrVal.(map[string]gopurs_runtime.Value)["value0"]
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})))
}))
}))
})
	})
	return genDate
}



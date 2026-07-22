package Data_Time_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_Enum_Gen "gopurs/output/Data.Enum.Gen"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
)

var genTime gopurs_runtime.Value
var once_genTime sync.Once
func Get_genTime() gopurs_runtime.Value {
	once_genTime.Do(func() {
		genTime = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_Time.Get_Time()), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0), pkg_Data_Time_Component.Get_boundedEnumHour()))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0), pkg_Data_Time_Component.Get_boundedEnumMinute()))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0), pkg_Data_Time_Component.Get_boundedEnumSecond()))), gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0), pkg_Data_Time_Component.Get_boundedEnumMillisecond()))
})
	})
	return genTime
}



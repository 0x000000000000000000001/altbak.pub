package Data_Time_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Time "gopurs/output/Data.Time"
	pkg_Data_Enum_Gen "gopurs/output/Data.Enum.Gen"
	pkg_Data_Time_Component "gopurs/output/Data.Time.Component"
)

var cache_genTime gopurs_runtime.Value
var once_genTime sync.Once
func Get_genTime() gopurs_runtime.Value {
	once_genTime.Do(func() {
		cache_genTime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genTime(dictMonadGen_0_box)
})
	})
	return cache_genTime
}

func Call_genTime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_Time.Get_Time(), gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumHour())), gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumMinute())), gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumSecond())), gopurs_runtime.Apply2(pkg_Data_Enum_Gen.Get_genBoundedEnum(), dictMonadGen_0, pkg_Data_Time_Component.Get_boundedEnumMillisecond()))
}



package Data_DateTime_Gen

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_DateTime "gopurs/output/Data.DateTime"
	pkg_Data_Date_Gen "gopurs/output/Data.Date.Gen"
	pkg_Data_Time_Gen "gopurs/output/Data.Time.Gen"
)

var genDateTime gopurs_runtime.Value
var once_genDateTime sync.Once
func Get_genDateTime() gopurs_runtime.Value {
	once_genDateTime.Do(func() {
		genDateTime = gopurs_runtime.Func(func(dictMonadGen_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(dictMonadGen_0.PtrVal.(map[string]gopurs_runtime.Value)["Monad0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Bind1"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["Apply0"], gopurs_runtime.Value{})
return gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["apply"], gopurs_runtime.Apply(gopurs_runtime.Apply(gopurs_runtime.Apply(Apply0_1_0.PtrVal.(map[string]gopurs_runtime.Value)["Functor0"], gopurs_runtime.Value{}).PtrVal.(map[string]gopurs_runtime.Value)["map"], pkg_Data_DateTime.Get_DateTime()), gopurs_runtime.Apply(pkg_Data_Date_Gen.Get_genDate(), dictMonadGen_0))), gopurs_runtime.Apply(pkg_Data_Time_Gen.Get_genTime(), dictMonadGen_0))
})
	})
	return genDateTime
}



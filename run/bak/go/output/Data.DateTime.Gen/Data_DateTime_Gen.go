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
		genDateTime = gopurs_runtime.Func(func(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_DateTime.Get_DateTime(), gopurs_runtime.Apply(pkg_Data_Date_Gen.Get_genDate(), dictMonadGen_0)), gopurs_runtime.Apply(pkg_Data_Time_Gen.Get_genTime(), dictMonadGen_0))
}()
})
	})
	return genDateTime
}





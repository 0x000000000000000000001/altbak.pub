package Data_DateTime_Gen

import (
	pkg_Data_Date_Gen "gopurs/output/Data.Date.Gen"
	pkg_Data_DateTime "gopurs/output/Data.DateTime"
	pkg_Data_Time_Gen "gopurs/output/Data.Time.Gen"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_genDateTime gopurs_runtime.Value
var once_genDateTime sync.Once
func Get_genDateTime() gopurs_runtime.Value {
	once_genDateTime.Do(func() {
		cache_genDateTime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_genDateTime(dictMonadGen_0_box)
})
	})
	return cache_genDateTime
}

func Call_genDateTime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), pkg_Data_DateTime.Get_DateTime(), gopurs_runtime.Apply(pkg_Data_Date_Gen.Get_genDate(), dictMonadGen_0)), gopurs_runtime.Apply(pkg_Data_Time_Gen.Get_genTime(), dictMonadGen_0))
}



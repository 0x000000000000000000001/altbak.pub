package Test_BenchCheck

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ring "gopurs/output/Data.Ring"
)

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), pkg_Bench.Get_benchNow(), gopurs_runtime.Func(func(t1_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), pkg_Bench.Get_benchNow(), gopurs_runtime.Func(func(t2_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("Delta: "), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showNumber(), "show"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), t2_1, t1_0))))
}))
})), nil)
}()
})
	})
	return cache_act
}



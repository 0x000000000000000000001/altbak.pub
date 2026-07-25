package Test_BenchCheck

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Bench "gopurs/output/Bench"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
t1_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_benchNow(), gopurs_runtime.Value{})
_ = t1_0_0
t2_1_1 := gopurs_runtime.Apply(pkg_Bench.Get_benchNow(), gopurs_runtime.Value{})
_ = t2_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(("Delta: ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), gopurs_runtime.Float((t2_1_1.FloatVal()) - (t1_0_0.FloatVal()))).StrVal()))), gopurs_runtime.Value{})
})
	})
	return cache_act
}





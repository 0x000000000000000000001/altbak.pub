package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_BenchCheck_act gopurs_runtime.Value
var once_Test_BenchCheck_act sync.Once

func Get_Test_BenchCheck_act() gopurs_runtime.Value {
	once_Test_BenchCheck_act.Do(func() {
		cache_Test_BenchCheck_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			t1_0_0 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
			_ = t1_0_0
			t2_1_1 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
			_ = t2_1_1
			return gopurs_runtime.Str(("Delta: ") + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), gopurs_runtime.Float((t2_1_1.FloatVal())-(t1_0_0.FloatVal()))).StrVal()))
		})
	})
	return cache_Test_BenchCheck_act
}

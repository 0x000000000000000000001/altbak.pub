package AppX

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Test_Church "gopurs/output/Test.Church"
	pkg_Bench "gopurs/output/Bench"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_0_0 := gopurs_runtime.Apply(pkg_Test_Church.Get_describe(), gopurs_runtime.Value{})
_ = _dollar__unused_0_0
t1_1_1 := gopurs_runtime.Apply(pkg_Bench.Get_benchNow(), gopurs_runtime.Value{})
_ = t1_1_1
_dollar__unused_2_2 := gopurs_runtime.Apply(pkg_Test_Church.Get_act(), gopurs_runtime.Value{})
_ = _dollar__unused_2_2
t2_3_3 := gopurs_runtime.Apply(pkg_Bench.Get_benchNow(), gopurs_runtime.Value{})
_ = t2_3_3
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("\n[BENCHMARK] Vitesse interne pure : ").StrVal + gopurs_runtime.Apply(pkg_Data_Show.Get_showNumberImpl(), gopurs_runtime.FloatSub(t2_3_3, t1_1_1)).StrVal).StrVal + gopurs_runtime.Str(" μs").StrVal)), gopurs_runtime.Value{})
})
	})
	return main
}



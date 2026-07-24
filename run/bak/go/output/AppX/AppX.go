package AppX

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Bench "gopurs/output/Bench"
	pkg_Test_TCO "gopurs/output/Test.TCO"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		main = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_TCO.Get_describe(), pkg_Test_TCO.Get_act())
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
a_prime_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = a_prime_1_1
return pkg_Data_Unit.Get_unit()
})
}()
	})
	return main
}





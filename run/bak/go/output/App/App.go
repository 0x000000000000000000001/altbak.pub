package App

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Bench "gopurs/output/Bench"
	pkg_Test_AstTree "gopurs/output/Test.AstTree"
	pkg_Test_Fib "gopurs/output/Test.Fib"
	pkg_Test_ListOps "gopurs/output/Test.ListOps"
	pkg_Test_TCO "gopurs/output/Test.TCO"
	pkg_Test_Records "gopurs/output/Test.Records"
	pkg_Test_Ackermann "gopurs/output/Test.Ackermann"
	pkg_Test_Church "gopurs/output/Test.Church"
	pkg_Test_Primes "gopurs/output/Test.Primes"
	pkg_Test_RBTree "gopurs/output/Test.RBTree"
	pkg_Test_Polymorphism "gopurs/output/Test.Polymorphism"
	pkg_Test_StateMonad "gopurs/output/Test.StateMonad"
	pkg_Test_LazyEvaluation "gopurs/output/Test.LazyEvaluation"
	pkg_Test_ArrayOps "gopurs/output/Test.ArrayOps"
	pkg_Effect_Console "gopurs/output/Effect.Console"
)

var cache_main gopurs_runtime.Value
var once_main sync.Once
func Get_main() gopurs_runtime.Value {
	once_main.Do(func() {
		cache_main = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_AstTree.Get_describe(), pkg_Test_AstTree.Get_act())
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
t1_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = t1_1_1
t2_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Fib.Get_describe(), pkg_Test_Fib.Get_act()), gopurs_runtime.Value{})
_ = t2_2_2
t3_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_ListOps.Get_describe(), pkg_Test_ListOps.Get_act()), gopurs_runtime.Value{})
_ = t3_3_3
t4_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_TCO.Get_describe(), pkg_Test_TCO.Get_act()), gopurs_runtime.Value{})
_ = t4_4_4
t5_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Records.Get_describe(), pkg_Test_Records.Get_act()), gopurs_runtime.Value{})
_ = t5_5_5
t6_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Ackermann.Get_describe(), pkg_Test_Ackermann.Get_act()), gopurs_runtime.Value{})
_ = t6_6_6
t7_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Church.Get_describe(), pkg_Test_Church.Get_act()), gopurs_runtime.Value{})
_ = t7_7_7
t8_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Primes.Get_describe(), pkg_Test_Primes.Get_act()), gopurs_runtime.Value{})
_ = t8_8_8
t9_9_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_RBTree.Get_describe(), pkg_Test_RBTree.Get_act()), gopurs_runtime.Value{})
_ = t9_9_9
t10_10_10 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_Polymorphism.Get_describe(), pkg_Test_Polymorphism.Get_act()), gopurs_runtime.Value{})
_ = t10_10_10
t11_11_11 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_StateMonad.Get_describe(), pkg_Test_StateMonad.Get_act()), gopurs_runtime.Value{})
_ = t11_11_11
t12_12_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_LazyEvaluation.Get_describe(), pkg_Test_LazyEvaluation.Get_act()), gopurs_runtime.Value{})
_ = t12_12_12
t13_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply2(pkg_Bench.Get_runBench(), pkg_Test_ArrayOps.Get_describe(), pkg_Test_ArrayOps.Get_act()), gopurs_runtime.Value{})
_ = t13_13_13
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str((("Total exec time: ") + (gopurs_runtime.Apply(pkg_Bench.Get_formatNumber(), gopurs_runtime.Float((((((((((((((t1_1_1.FloatVal()) / (1000.0)) + ((t2_2_2.FloatVal()) / (1000.0))) + ((t3_3_3.FloatVal()) / (1000.0))) + ((t4_4_4.FloatVal()) / (1000.0))) + ((t5_5_5.FloatVal()) / (1000.0))) + ((t6_6_6.FloatVal()) / (1000.0))) + ((t7_7_7.FloatVal()) / (1000.0))) + ((t8_8_8.FloatVal()) / (1000.0))) + ((t9_9_9.FloatVal()) / (1000.0))) + ((t10_10_10.FloatVal()) / (1000.0))) + ((t11_11_11.FloatVal()) / (1000.0))) + ((t12_12_12.FloatVal()) / (1000.0))) + ((t13_13_13.FloatVal()) / (1000.0)))).StrVal())) + (" ms\n"))), gopurs_runtime.Value{})
})
}()
	})
	return cache_main
}





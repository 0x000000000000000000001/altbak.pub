package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_App_main gopurs_runtime.Value
var once_App_main sync.Once
func Get_App_main() gopurs_runtime.Value {
	once_App_main.Do(func() {
		cache_App_main = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTree_describe(), Get_Test_AstTree_act())
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Fib_describe(), Get_Test_Fib_act()), gopurs_runtime.Value{})
_ = __local_var_2_2
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOps_describe(), Get_Test_ListOps_act()), gopurs_runtime.Value{})
_ = __local_var_3_3
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCO_describe(), Get_Test_TCO_act()), gopurs_runtime.Value{})
_ = __local_var_4_4
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Records_describe(), Get_Test_Records_act()), gopurs_runtime.Value{})
_ = __local_var_5_5
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Ackermann_describe(), Get_Test_Ackermann_act()), gopurs_runtime.Value{})
_ = __local_var_6_6
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Church_describe(), Get_Test_Church_act()), gopurs_runtime.Value{})
_ = __local_var_7_7
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Primes_describe(), Get_Test_Primes_act()), gopurs_runtime.Value{})
_ = __local_var_8_8
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTree_describe(), Get_Test_RBTree_act()), gopurs_runtime.Value{})
_ = __local_var_9_9
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_Polymorphism_describe(), Get_Test_Polymorphism_act()), gopurs_runtime.Value{})
_ = __local_var_10_10
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonad_describe(), Get_Test_StateMonad_act()), gopurs_runtime.Value{})
_ = __local_var_11_11
__local_var_12_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluation_describe(), Get_Test_LazyEvaluation_act()), gopurs_runtime.Value{})
_ = __local_var_12_12
__local_var_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOps_describe(), Get_Test_ArrayOps_act()), gopurs_runtime.Value{})
_ = __local_var_13_13
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("\x0a==================================================\x0a\x0aTotal exec time: ") + (gopurs_runtime.Apply(Get_Bench_formatNumber(), gopurs_runtime.Float((((((((((((((__local_var_1_1.FloatVal()) / (1000.0)) + ((__local_var_2_2.FloatVal()) / (1000.0))) + ((__local_var_3_3.FloatVal()) / (1000.0))) + ((__local_var_4_4.FloatVal()) / (1000.0))) + ((__local_var_5_5.FloatVal()) / (1000.0))) + ((__local_var_6_6.FloatVal()) / (1000.0))) + ((__local_var_7_7.FloatVal()) / (1000.0))) + ((__local_var_8_8.FloatVal()) / (1000.0))) + ((__local_var_9_9.FloatVal()) / (1000.0))) + ((__local_var_10_10.FloatVal()) / (1000.0))) + ((__local_var_11_11.FloatVal()) / (1000.0))) + ((__local_var_12_12.FloatVal()) / (1000.0))) + ((__local_var_13_13.FloatVal()) / (1000.0)))).StrVal())) + (" ms\x0a"))), gopurs_runtime.Value{})
})
}()
	})
	return cache_App_main
}





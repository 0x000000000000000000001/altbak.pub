package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_AppFFICheatcode_main gopurs_runtime.Value
var once_AppFFICheatcode_main sync.Once
func Get_AppFFICheatcode_main() gopurs_runtime.Value {
	once_AppFFICheatcode_main.Do(func() {
		cache_AppFFICheatcode_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
__local_var_0_0 := gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTreeFFICheatcode_describe(), Get_Test_AstTreeFFICheatcode_act())
_ = __local_var_0_0
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_FibFFICheatcode_describe(), Get_Test_FibFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_2_2
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOpsFFICheatcode_describe(), Get_Test_ListOpsFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_3_3
__local_var_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCOFFICheatcode_describe(), Get_Test_TCOFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_4_4
__local_var_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RecordsFFICheatcode_describe(), Get_Test_RecordsFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_5_5
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AckermannFFICheatcode_describe(), Get_Test_AckermannFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_6_6
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ChurchFFICheatcode_describe(), Get_Test_ChurchFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_7_7
__local_var_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PrimesFFICheatcode_describe(), Get_Test_PrimesFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_8_8
__local_var_9_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTreeFFICheatcode_describe(), Get_Test_RBTreeFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_9_9
__local_var_10_10 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PolymorphismFFICheatcode_describe(), Get_Test_PolymorphismFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_10_10
__local_var_11_11 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonadFFICheatcode_describe(), Get_Test_StateMonadFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_11_11
__local_var_12_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluationFFICheatcode_describe(), Get_Test_LazyEvaluationFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_12_12
__local_var_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOpsFFICheatcode_describe(), Get_Test_ArrayOpsFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_13_13
__local_var_14_14 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToListFFICheatcode_describe(), Get_Test_RowToListFFICheatcode_act()), gopurs_runtime.Value{})
_ = __local_var_14_14
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("\x0a==================================================\x0a\x0aTotal exec time: ") + (gopurs_runtime.Apply(Get_Bench_formatNumber(), gopurs_runtime.Float(((((((((((((((__local_var_1_1.FloatVal()) / (1000.0)) + ((__local_var_2_2.FloatVal()) / (1000.0))) + ((__local_var_3_3.FloatVal()) / (1000.0))) + ((__local_var_4_4.FloatVal()) / (1000.0))) + ((__local_var_5_5.FloatVal()) / (1000.0))) + ((__local_var_6_6.FloatVal()) / (1000.0))) + ((__local_var_7_7.FloatVal()) / (1000.0))) + ((__local_var_8_8.FloatVal()) / (1000.0))) + ((__local_var_9_9.FloatVal()) / (1000.0))) + ((__local_var_10_10.FloatVal()) / (1000.0))) + ((__local_var_11_11.FloatVal()) / (1000.0))) + ((__local_var_12_12.FloatVal()) / (1000.0))) + ((__local_var_13_13.FloatVal()) / (1000.0))) + ((__local_var_14_14.FloatVal()) / (1000.0)))).StrVal())) + (" ms\x0a"))), gopurs_runtime.Value{})
})
	})
	return cache_AppFFICheatcode_main
}





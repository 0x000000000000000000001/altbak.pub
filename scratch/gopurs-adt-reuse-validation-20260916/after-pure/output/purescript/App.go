package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_App_warmup gopurs_runtime.Value
var once_App_warmup sync.Once
func Get_App_warmup() gopurs_runtime.Value {
	once_App_warmup.Do(func() {
		cache_App_warmup = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_AstTree_act()), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Fib_act()), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ListOps_act()), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_TCO_act()), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Records_act()), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Ackermann_act()), gopurs_runtime.Func(func(_dollar___unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Church_act()), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Primes_act()), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RBTree_act()), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_Polymorphism_act()), gopurs_runtime.Func(func(_dollar___unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_StateMonad_act()), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_LazyEvaluation_act()), gopurs_runtime.Func(func(_dollar___unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_ArrayOps_act()), gopurs_runtime.Func(func(_dollar___unused_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), Get_Test_RowToList_act())
}))
}))
}))
}))
}))
}))
}))
}))
}))
}))
}))
}))
}))
	})
	return cache_App_warmup
}

var cache_App_main gopurs_runtime.Value
var once_App_main sync.Once
func Get_App_main() gopurs_runtime.Value {
	once_App_main.Do(func() {
		cache_App_main = gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Global warm-up in progress (this may take a moment)...\x0a")), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_App_warmup(), gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_App_warmup(), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), Get_App_warmup(), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Global warm-up complete. Starting benchmarks...\x0a")), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_0 shape=App(Var) bindingType=Any
__local_var_5_0 := Call_Bench_runBench(Get_Test_AstTree_describe(), Get_Test_AstTree_act())
_ = __local_var_5_0
__local_var_6_1 := gopurs_runtime.Apply(__local_var_5_0, gopurs_runtime.Value{})
_ = __local_var_6_1
__local_var_7_2 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_Fib_describe(), Get_Test_Fib_act()), gopurs_runtime.Value{})
_ = __local_var_7_2
__local_var_8_3 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_ListOps_describe(), Get_Test_ListOps_act()), gopurs_runtime.Value{})
_ = __local_var_8_3
__local_var_9_4 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_TCO_describe(), Get_Test_TCO_act()), gopurs_runtime.Value{})
_ = __local_var_9_4
__local_var_10_5 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_Records_describe(), Get_Test_Records_act()), gopurs_runtime.Value{})
_ = __local_var_10_5
__local_var_11_6 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_Ackermann_describe(), Get_Test_Ackermann_act()), gopurs_runtime.Value{})
_ = __local_var_11_6
__local_var_12_7 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_Church_describe(), Get_Test_Church_act()), gopurs_runtime.Value{})
_ = __local_var_12_7
__local_var_13_8 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_Primes_describe(), Get_Test_Primes_act()), gopurs_runtime.Value{})
_ = __local_var_13_8
__local_var_14_9 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_RBTree_describe(), Get_Test_RBTree_act()), gopurs_runtime.Value{})
_ = __local_var_14_9
__local_var_15_10 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_Polymorphism_describe(), Get_Test_Polymorphism_act()), gopurs_runtime.Value{})
_ = __local_var_15_10
__local_var_16_11 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_StateMonad_describe(), Get_Test_StateMonad_act()), gopurs_runtime.Value{})
_ = __local_var_16_11
__local_var_17_12 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_LazyEvaluation_describe(), Get_Test_LazyEvaluation_act()), gopurs_runtime.Value{})
_ = __local_var_17_12
__local_var_18_13 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_ArrayOps_describe(), Get_Test_ArrayOps_act()), gopurs_runtime.Value{})
_ = __local_var_18_13
__local_var_19_14 := gopurs_runtime.Apply(Call_Bench_runBench(Get_Test_RowToList_describe(), Get_Test_RowToList_act()), gopurs_runtime.Value{})
_ = __local_var_19_14
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("\x0a==================================================\x0a\x0aTotal exec time: ") + (gopurs_runtime.Apply(Get_Bench_formatNumber(), gopurs_runtime.Float(((((((((((((((__local_var_6_1.FloatVal()) / (1000.0)) + ((__local_var_7_2.FloatVal()) / (1000.0))) + ((__local_var_8_3.FloatVal()) / (1000.0))) + ((__local_var_9_4.FloatVal()) / (1000.0))) + ((__local_var_10_5.FloatVal()) / (1000.0))) + ((__local_var_11_6.FloatVal()) / (1000.0))) + ((__local_var_12_7.FloatVal()) / (1000.0))) + ((__local_var_13_8.FloatVal()) / (1000.0))) + ((__local_var_14_9.FloatVal()) / (1000.0))) + ((__local_var_15_10.FloatVal()) / (1000.0))) + ((__local_var_16_11.FloatVal()) / (1000.0))) + ((__local_var_17_12.FloatVal()) / (1000.0))) + ((__local_var_18_13.FloatVal()) / (1000.0))) + ((__local_var_19_14.FloatVal()) / (1000.0)))).StrVal())) + (" ms\x0a"))), gopurs_runtime.Value{})
})
}))
}))
}))
}))
}))
	})
	return cache_App_main
}





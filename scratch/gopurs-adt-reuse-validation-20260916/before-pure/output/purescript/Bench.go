package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Bench_runBenchSync gopurs_runtime.Value
var once_Bench_runBenchSync sync.Once
func Get_Bench_runBenchSync() gopurs_runtime.Value {
	once_Bench_runBenchSync.Do(func() {
		cache_Bench_runBenchSync = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Bench_runBenchSync(describe_0_box, act_1_box)
})
	})
	return cache_Bench_runBenchSync
}

var cache_Bench_runBench gopurs_runtime.Value
var once_Bench_runBench sync.Once
func Get_Bench_runBench() gopurs_runtime.Value {
	once_Bench_runBench.Do(func() {
		cache_Bench_runBench = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Bench_runBench(describe_0_box, act_1_box)
})
	})
	return cache_Bench_runBench
}

func Call_Bench_runBenchSync(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a")), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), describe_0, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("\x0a(Output & Warm-up)\x0a")), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := gopurs_runtime.Apply(act_1, gopurs_runtime.Value{})
_ = __local_var_5_0
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__local_var_5_0.StrVal())), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_1 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_9_1
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_11_2 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_11_2
// TAST (Let): d1_12_3 shape=Other bindingType=Number
d1_12_3 := (__local_var_11_2.FloatVal()) - (__local_var_9_1.FloatVal())
_ = d1_12_3
__local_var_13_4 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_13_4
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_15_5 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_15_5
// TAST (Let): d2_16_6 shape=Other bindingType=Number
d2_16_6 := (__local_var_15_5.FloatVal()) - (__local_var_13_4.FloatVal())
_ = d2_16_6
__local_var_17_7 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_17_7
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_19_8 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_19_8
// TAST (Let): d3_20_9 shape=Other bindingType=Number
d3_20_9 := (__local_var_19_8.FloatVal()) - (__local_var_17_7.FloatVal())
_ = d3_20_9
__local_var_21_10 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_21_10
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_23_11 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_23_11
// TAST (Let): d4_24_12 shape=Other bindingType=Number
d4_24_12 := (__local_var_23_11.FloatVal()) - (__local_var_21_10.FloatVal())
_ = d4_24_12
__local_var_25_13 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_25_13
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_27_14 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_27_14
// TAST (Let): d5_28_15 shape=Other bindingType=Number
d5_28_15 := (__local_var_27_14.FloatVal()) - (__local_var_25_13.FloatVal())
_ = d5_28_15
__local_var_29_16 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_29_16
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_30 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_31_17 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_31_17
// TAST (Let): d6_32_18 shape=Other bindingType=Number
d6_32_18 := (__local_var_31_17.FloatVal()) - (__local_var_29_16.FloatVal())
_ = d6_32_18
__local_var_33_19 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_33_19
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_34 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_34 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_35_20 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_35_20
// TAST (Let): d7_36_21 shape=Other bindingType=Number
d7_36_21 := (__local_var_35_20.FloatVal()) - (__local_var_33_19.FloatVal())
_ = d7_36_21
__local_var_37_22 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_37_22
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_38 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_38 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_39_23 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_39_23
// TAST (Let): d8_40_24 shape=Other bindingType=Number
d8_40_24 := (__local_var_39_23.FloatVal()) - (__local_var_37_22.FloatVal())
_ = d8_40_24
__local_var_41_25 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_41_25
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_42 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_42 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_43_26 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_43_26
// TAST (Let): d9_44_27 shape=Other bindingType=Number
d9_44_27 := (__local_var_43_26.FloatVal()) - (__local_var_41_25.FloatVal())
_ = d9_44_27
__local_var_45_28 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_45_28
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_46 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_46 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_47_29 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_47_29
// TAST (Let): best_48_30 shape=App(Var) bindingType=Number
best_48_30 := Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(d1_12_3, d2_16_6), Call_Data_Ord_min__1229405204(d3_20_9, d4_24_12)), Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(d5_28_15, d6_32_18), Call_Data_Ord_min__1229405204(d7_36_21, d8_40_24))), Call_Data_Ord_min__1229405204(d9_44_27, (__local_var_47_29.FloatVal()) - (__local_var_45_28.FloatVal())))
_ = best_48_30
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("\x0a(Execution time - best of 10)\x0a\x0a") + (gopurs_runtime.Apply(Get_Bench_formatNumber(), gopurs_runtime.Float(best_48_30)).StrVal())) + (" μs\x0a"))), gopurs_runtime.Func(func(_dollar___unused_49 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(best_48_30)
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
}))
}))
})), gopurs_runtime.Value{})
})
}))
}))
}))
}

func Call_Bench_runBench(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a")), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), describe_0, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("\x0a(Output & Warm-up)\x0a")), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_0 := gopurs_runtime.Apply(act_1, gopurs_runtime.Value{})
_ = __local_var_5_0
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__local_var_5_0.StrVal())), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_9_1 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_9_1
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_11_2 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_11_2
// TAST (Let): d1_12_3 shape=Other bindingType=Number
d1_12_3 := (__local_var_11_2.FloatVal()) - (__local_var_9_1.FloatVal())
_ = d1_12_3
__local_var_13_4 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_13_4
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_15_5 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_15_5
// TAST (Let): d2_16_6 shape=Other bindingType=Number
d2_16_6 := (__local_var_15_5.FloatVal()) - (__local_var_13_4.FloatVal())
_ = d2_16_6
__local_var_17_7 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_17_7
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_19_8 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_19_8
// TAST (Let): d3_20_9 shape=Other bindingType=Number
d3_20_9 := (__local_var_19_8.FloatVal()) - (__local_var_17_7.FloatVal())
_ = d3_20_9
__local_var_21_10 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_21_10
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_23_11 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_23_11
// TAST (Let): d4_24_12 shape=Other bindingType=Number
d4_24_12 := (__local_var_23_11.FloatVal()) - (__local_var_21_10.FloatVal())
_ = d4_24_12
__local_var_25_13 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_25_13
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_27_14 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_27_14
// TAST (Let): d5_28_15 shape=Other bindingType=Number
d5_28_15 := (__local_var_27_14.FloatVal()) - (__local_var_25_13.FloatVal())
_ = d5_28_15
__local_var_29_16 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_29_16
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_30 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_31_17 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_31_17
// TAST (Let): d6_32_18 shape=Other bindingType=Number
d6_32_18 := (__local_var_31_17.FloatVal()) - (__local_var_29_16.FloatVal())
_ = d6_32_18
__local_var_33_19 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_33_19
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_34 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_34 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_35_20 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_35_20
// TAST (Let): d7_36_21 shape=Other bindingType=Number
d7_36_21 := (__local_var_35_20.FloatVal()) - (__local_var_33_19.FloatVal())
_ = d7_36_21
__local_var_37_22 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_37_22
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_38 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_38 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_39_23 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_39_23
// TAST (Let): d8_40_24 shape=Other bindingType=Number
d8_40_24 := (__local_var_39_23.FloatVal()) - (__local_var_37_22.FloatVal())
_ = d8_40_24
__local_var_41_25 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_41_25
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_42 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_42 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_43_26 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_43_26
// TAST (Let): d9_44_27 shape=Other bindingType=Number
d9_44_27 := (__local_var_43_26.FloatVal()) - (__local_var_41_25.FloatVal())
_ = d9_44_27
__local_var_45_28 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_45_28
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0, gopurs_runtime.Func(func(v_46 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_46 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_47_29 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_47_29
// TAST (Let): best_48_30 shape=App(Var) bindingType=Number
best_48_30 := Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(d1_12_3, d2_16_6), Call_Data_Ord_min__1229405204(d3_20_9, d4_24_12)), Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(d5_28_15, d6_32_18), Call_Data_Ord_min__1229405204(d7_36_21, d8_40_24))), Call_Data_Ord_min__1229405204(d9_44_27, (__local_var_47_29.FloatVal()) - (__local_var_45_28.FloatVal())))
_ = best_48_30
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_bindEffect())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("\x0a(Execution time - best of 10)\x0a\x0a") + (gopurs_runtime.Apply(Get_Bench_formatNumber(), gopurs_runtime.Float(best_48_30)).StrVal())) + (" μs\x0a"))), gopurs_runtime.Func(func(_dollar___unused_49 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(best_48_30)
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
})
}))
}))
})), gopurs_runtime.Value{})
})
}))
}))
}))
}

func Get_Bench_benchNow() gopurs_runtime.Value {
	return _Gopurs_Bench_BenchNow
}

func Get_Bench_formatNumber() gopurs_runtime.Value {
	return _Gopurs_Bench_FormatNumber
}

func Get_Bench_opaque() gopurs_runtime.Value {
	return _Gopurs_Bench_Opaque
}

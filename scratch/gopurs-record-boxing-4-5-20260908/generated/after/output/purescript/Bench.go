package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

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

func Call_Bench_runBench(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Var) bindingType=Any
__local_var_2_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a"))
_ = __local_var_2_0
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(describe_0, gopurs_runtime.Value{})
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("\x0a(Output & Warm-up)\x0a")), gopurs_runtime.Value{})
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.Apply(act_1, gopurs_runtime.Value{})
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(__local_var_6_4.StrVal())), gopurs_runtime.Value{})
_ = __local_var_7_5
__local_var_8_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_8_6
__local_var_9_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_9_7
__local_var_10_8 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_10_8
__local_var_11_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_11_9
__local_var_12_10 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_12_10
__local_var_13_11 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_13_11
__local_var_14_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_14_12
__local_var_15_13 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_15_13
__local_var_16_14 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_16_14
__local_var_17_15 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_17_15
__local_var_18_16 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_18_16
__local_var_19_17 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_19_17
__local_var_20_18 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_20_18
__local_var_21_19 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_21_19
__local_var_22_20 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_22_20
__local_var_23_21 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_23_21
__local_var_24_22 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_24_22
__local_var_25_23 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_25_23
__local_var_26_24 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_26_24
__local_var_27_25 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_27_25
__local_var_28_26 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_28_26
__local_var_29_27 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_29 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_29_27
__local_var_30_28 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_30_28
__local_var_31_29 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_31_29
__local_var_32_30 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_32 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_32_30
__local_var_33_31 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_33_31
__local_var_34_32 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_34_32
__local_var_35_33 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_35 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_35_33
__local_var_36_34 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_36_34
__local_var_37_35 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_37_35
__local_var_38_36 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_38 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Value{})
_ = __local_var_38_36
__local_var_39_37 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_39_37
// TAST (Let): best_40_38 shape=App(Var) bindingType=Number
best_40_38 := gopurs_runtime.Apply2(Get_Data_Ord_min__294757560(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Data_Ord_min__294757560(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Data_Ord_min__294757560(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Data_Ord_min__294757560(), gopurs_runtime.Float((__local_var_12_10.FloatVal()) - (__local_var_10_8.FloatVal())), gopurs_runtime.Float((__local_var_15_13.FloatVal()) - (__local_var_13_11.FloatVal()))).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Data_Ord_min__294757560(), gopurs_runtime.Float((__local_var_18_16.FloatVal()) - (__local_var_16_14.FloatVal())), gopurs_runtime.Float((__local_var_21_19.FloatVal()) - (__local_var_19_17.FloatVal()))).FloatVal())).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Data_Ord_min__294757560(), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Data_Ord_min__294757560(), gopurs_runtime.Float((__local_var_24_22.FloatVal()) - (__local_var_22_20.FloatVal())), gopurs_runtime.Float((__local_var_27_25.FloatVal()) - (__local_var_25_23.FloatVal()))).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Data_Ord_min__294757560(), gopurs_runtime.Float((__local_var_30_28.FloatVal()) - (__local_var_28_26.FloatVal())), gopurs_runtime.Float((__local_var_33_31.FloatVal()) - (__local_var_31_29.FloatVal()))).FloatVal())).FloatVal())).FloatVal()), gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Data_Ord_min__294757560(), gopurs_runtime.Float((__local_var_36_34.FloatVal()) - (__local_var_34_32.FloatVal())), gopurs_runtime.Float((__local_var_39_37.FloatVal()) - (__local_var_37_35.FloatVal()))).FloatVal())).FloatVal()
_ = best_40_38
__local_var_41_39 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("\x0a(Execution time - best of 10)\x0a\x0a") + (gopurs_runtime.Apply(Get_Bench_formatNumber(), gopurs_runtime.Float(best_40_38)).StrVal())) + (" μs\x0a"))), gopurs_runtime.Value{})
_ = __local_var_41_39
return gopurs_runtime.Float(best_40_38)
})
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

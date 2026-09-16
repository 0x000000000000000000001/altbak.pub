package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Bench_Aff_liftEffect gopurs_runtime.Value
var once_Bench_Aff_liftEffect sync.Once
func Get_Bench_Aff_liftEffect() gopurs_runtime.Value {
	once_Bench_Aff_liftEffect.Do(func() {
		cache_Bench_Aff_liftEffect = Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff()))
	})
	return cache_Bench_Aff_liftEffect
}

var cache_Bench_Aff_runBenchAff gopurs_runtime.Value
var once_Bench_Aff_runBenchAff sync.Once
func Get_Bench_Aff_runBenchAff() gopurs_runtime.Value {
	once_Bench_Aff_runBenchAff.Do(func() {
		cache_Bench_Aff_runBenchAff = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Bench_Aff_runBenchAff(describe_0_box, act_1_box)
})
	})
	return cache_Bench_Aff_runBenchAff
}

func Call_Bench_Aff_runBenchAff(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a"))), gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), describe_0), gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("\x0a(Output & Warm-up)\x0a"))), gopurs_runtime.Func(func(_dollar___unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), act_1, gopurs_runtime.Func(func(out_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(out_5.StrVal()))), gopurs_runtime.Func(func(_dollar___unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t1_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t2_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): d1_12_0 shape=Other bindingType=Number
d1_12_0 := (t2_11.FloatVal()) - (t1_9.FloatVal())
_ = d1_12_0
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t3_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t4_15 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): d2_16_1 shape=Other bindingType=Number
d2_16_1 := (t4_15.FloatVal()) - (t3_13.FloatVal())
_ = d2_16_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t5_17 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_18 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t6_19 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): d3_20_2 shape=Other bindingType=Number
d3_20_2 := (t6_19.FloatVal()) - (t5_17.FloatVal())
_ = d3_20_2
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t7_21 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_22 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t8_23 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): d4_24_3 shape=Other bindingType=Number
d4_24_3 := (t8_23.FloatVal()) - (t7_21.FloatVal())
_ = d4_24_3
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t9_25 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_26 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t10_27 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): d5_28_4 shape=Other bindingType=Number
d5_28_4 := (t10_27.FloatVal()) - (t9_25.FloatVal())
_ = d5_28_4
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t11_29 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_30 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_30 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t12_31 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): d6_32_5 shape=Other bindingType=Number
d6_32_5 := (t12_31.FloatVal()) - (t11_29.FloatVal())
_ = d6_32_5
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t13_33 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_34 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_34 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t14_35 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): d7_36_6 shape=Other bindingType=Number
d7_36_6 := (t14_35.FloatVal()) - (t13_33.FloatVal())
_ = d7_36_6
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t15_37 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_38 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_38 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t16_39 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): d8_40_7 shape=Other bindingType=Number
d8_40_7 := (t16_39.FloatVal()) - (t15_37.FloatVal())
_ = d8_40_7
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t17_41 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_42 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_42 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t18_43 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): d9_44_8 shape=Other bindingType=Number
d9_44_8 := (t18_43.FloatVal()) - (t17_41.FloatVal())
_ = d9_44_8
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t19_45 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply2(Get_Effect_Aff__map(), gopurs_runtime.Func(func(v_46 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), act_1), gopurs_runtime.Func(func(_dollar___unused_46 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t20_47 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): best_48_9 shape=App(Var) bindingType=Number
best_48_9 := Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(d1_12_0, d2_16_1), Call_Data_Ord_min__1229405204(d3_20_2, d4_24_3)), Call_Data_Ord_min__1229405204(Call_Data_Ord_min__1229405204(d5_28_4, d6_32_5), Call_Data_Ord_min__1229405204(d7_36_6, d8_40_7))), Call_Data_Ord_min__1229405204(d9_44_8, (t20_47.FloatVal()) - (t19_45.FloatVal())))
_ = best_48_9
return gopurs_runtime.Apply2(Call_Control_Bind_bind(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Effect_Aff_bindAff())), gopurs_runtime.Apply(Call_Effect_Class_liftEffect(gopurs_runtime.CoerceToStruct[Constructor_Effect_Class_MonadEffect[gopurs_runtime.Value]](Get_Effect_Aff_monadEffectAff())), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("\x0a(Execution time - best of 10)\x0a\x0a") + (gopurs_runtime.Apply(Get_Bench_formatNumber(), gopurs_runtime.Float(best_48_9)).StrVal())) + (" μs\x0a")))), gopurs_runtime.Func(func(_dollar___unused_49 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Effect_Aff__pure(), gopurs_runtime.Float(best_48_9))
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
}



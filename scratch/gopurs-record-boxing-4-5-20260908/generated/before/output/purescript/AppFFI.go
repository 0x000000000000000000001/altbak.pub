package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_AppFFI_void gopurs_runtime.Value
var once_AppFFI_void sync.Once
func Get_AppFFI_void() gopurs_runtime.Value {
	once_AppFFI_void.Do(func() {
		cache_AppFFI_void = gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
	})
	return cache_AppFFI_void
}

var cache_AppFFI_main gopurs_runtime.Value
var once_AppFFI_main sync.Once
func Get_AppFFI_main() gopurs_runtime.Value {
	once_AppFFI_main.Do(func() {
		cache_AppFFI_main = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar___unused_0_0 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTreeFFI_describe(), Get_Test_AstTreeFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_0_0
_dollar___unused_1_1 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_FibFFI_describe(), Get_Test_FibFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_1_1
_dollar___unused_2_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOpsFFI_describe(), Get_Test_ListOpsFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_2_2
_dollar___unused_3_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCOFFI_describe(), Get_Test_TCOFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_3_3
_dollar___unused_4_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RecordsFFI_describe(), Get_Test_RecordsFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_4_4
_dollar___unused_5_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AckermannFFI_describe(), Get_Test_AckermannFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_5_5
_dollar___unused_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ChurchFFI_describe(), Get_Test_ChurchFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_6_6
_dollar___unused_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PrimesFFI_describe(), Get_Test_PrimesFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_7_7
_dollar___unused_8_8 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTreeFFI_describe(), Get_Test_RBTreeFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_8_8
_dollar___unused_9_9 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PolymorphismFFI_describe(), Get_Test_PolymorphismFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_9_9
_dollar___unused_10_10 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonadFFI_describe(), Get_Test_StateMonadFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_10_10
_dollar___unused_11_11 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluationFFI_describe(), Get_Test_LazyEvaluationFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_11_11
_dollar___unused_12_12 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOpsFFI_describe(), Get_Test_ArrayOpsFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_12_12
_dollar___unused_13_13 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToListFFI_describe(), Get_Test_RowToListFFI_act())), gopurs_runtime.Value{})
_ = _dollar___unused_13_13
_dollar___unused_14_14 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_14 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AstTreeFFICheatcode_describe(), Get_Test_AstTreeFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_14_14
_dollar___unused_15_15 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_15 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_FibFFICheatcode_describe(), Get_Test_FibFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_15_15
_dollar___unused_16_16 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_16 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ListOpsFFICheatcode_describe(), Get_Test_ListOpsFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_16_16
_dollar___unused_17_17 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_TCOFFICheatcode_describe(), Get_Test_TCOFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_17_17
_dollar___unused_18_18 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_18 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RecordsFFICheatcode_describe(), Get_Test_RecordsFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_18_18
_dollar___unused_19_19 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_19 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_AckermannFFICheatcode_describe(), Get_Test_AckermannFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_19_19
_dollar___unused_20_20 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_20 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ChurchFFICheatcode_describe(), Get_Test_ChurchFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_20_20
_dollar___unused_21_21 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_21 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PrimesFFICheatcode_describe(), Get_Test_PrimesFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_21_21
_dollar___unused_22_22 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_22 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RBTreeFFICheatcode_describe(), Get_Test_RBTreeFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_22_22
_dollar___unused_23_23 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_23 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_PolymorphismFFICheatcode_describe(), Get_Test_PolymorphismFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_23_23
_dollar___unused_24_24 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_24 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_StateMonadFFICheatcode_describe(), Get_Test_StateMonadFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_24_24
_dollar___unused_25_25 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_25 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_LazyEvaluationFFICheatcode_describe(), Get_Test_LazyEvaluationFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_25_25
_dollar___unused_26_26 := gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_26 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_ArrayOpsFFICheatcode_describe(), Get_Test_ArrayOpsFFICheatcode_act())), gopurs_runtime.Value{})
_ = _dollar___unused_26_26
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Effect_functorEffect()).V0), gopurs_runtime.Func(func(v_27 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(Get_Bench_runBench(), Get_Test_RowToListFFICheatcode_describe(), Get_Test_RowToListFFICheatcode_act())), gopurs_runtime.Value{})
})
	})
	return cache_AppFFI_main
}





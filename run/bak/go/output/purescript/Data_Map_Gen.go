package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Map_Gen_genMap gopurs_runtime.Value
var once_Data_Map_Gen_genMap sync.Once
func Get_Data_Map_Gen_genMap() gopurs_runtime.Value {
	once_Data_Map_Gen_genMap.Do(func() {
		cache_Data_Map_Gen_genMap = gopurs_runtime.Func2(func(dictMonadRec_0_box gopurs_runtime.Value, dictMonadGen_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Map_Gen_genMap(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box), gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Class_MonadGen](dictMonadGen_1_box))
})
	})
	return cache_Data_Map_Gen_genMap
}

func Call_Data_Map_Gen_genMap(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec, dictMonadGen_1_loop *Constructor_Control_Monad_Gen_Class_MonadGen) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
var dictMonadGen_1 *Constructor_Control_Monad_Gen_Class_MonadGen = dictMonadGen_1_loop
_ = dictMonadGen_1
// TAST (Let): Monad0_2_0 -> gopurs_runtime.Value
Monad0_2_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_2_0
// TAST (Let): Bind1_3_1 -> *Constructor_Control_Bind_Bind
Bind1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_1
// TAST (Let): Bind11_4_2 -> gopurs_runtime.Value
Bind11_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_0, "Bind1"), gopurs_runtime.Value{})
_ = Bind11_4_2
// TAST (Let): Functor0_5_3 -> *Constructor_Data_Functor_Functor
Functor0_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind11_4_2, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_3
// TAST (Let): Apply0_6_4 -> *Constructor_Control_Apply_Apply
Apply0_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind11_4_2, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_6_4
return gopurs_runtime.Func(func(dictOrd_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): fromFoldable_8_5 -> gopurs_runtime.Value
fromFoldable_8_5 := gopurs_runtime.Apply2(Get_Data_Map_Internal_fromFoldable(), dictOrd_7, Get_Data_List_Types_foldableList())
_ = fromFoldable_8_5
return gopurs_runtime.Func(func(genKey_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genValue_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(size_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V3), gopurs_runtime.Int(0), gopurs_runtime.Int(size_11.IntVal)), gopurs_runtime.Func(func(newSize_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V4), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_12.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), fromFoldable_8_5, gopurs_runtime.Apply4(Get_Control_Monad_Gen_unfoldable(), gopurs_runtime.Value{Type: 9, IntVal: 3709389635, UnsafePtr: unsafe.Pointer(dictMonadRec_0)}, gopurs_runtime.Value{Type: 9, IntVal: 2254593219, UnsafePtr: unsafe.Pointer(dictMonadGen_1)}, gopurs_runtime.Value{Type: 9, IntVal: 2670894170, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](Get_Data_List_Types_unfoldableList()))}, gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_4.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), Get_Data_Tuple_Tuple(), genKey_9), genValue_10))))
}))
}))
})
})
})
}



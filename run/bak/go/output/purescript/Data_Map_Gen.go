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
var go__go_8_6_0 gopurs_runtime.Value
go__go_8_6_0 = gopurs_runtime.Func(func(b_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_9_loop gopurs_runtime.Value = b_9_loop_val
var v_10_loop gopurs_runtime.Value = v_10_loop_val
go__go_8_6_0:
for {
if false { continue go__go_8_6_0 }
var b_9 gopurs_runtime.Value = b_9_loop
_ = b_9
var v_10 gopurs_runtime.Value = v_10_loop
_ = v_10
var __t13 gopurs_runtime.Value
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr == nil) {
__t13 = b_9
goto end_branch_13
} else {

}
}
{
if (v_10.Type == 9 && v_10.IntVal == 1358893437 && v_10.UnsafePtr != nil) {
// TAST (Let): __local_var_11_7 -> gopurs_runtime.Value
__local_var_11_7 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0.UnsafePtr).V0
_ = __local_var_11_7
// TAST (Let): __local_var_12_8 -> gopurs_runtime.Value
__local_var_12_8 := (*Constructor_Data_Tuple_Tuple)((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V0.UnsafePtr).V1
_ = __local_var_12_8
var go__go_13_9_1 gopurs_runtime.Value
_ = go__go_13_9_1
go__go_13_9_1 = gopurs_runtime.Func(func(v1_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 *Constructor_Data_Map_Internal_Node
{
if (v1_14.Type == 9 && v1_14.IntVal == 324739070 && v1_14.UnsafePtr == nil) {
__t12 = &Constructor_Data_Map_Internal_Node{1, 1, 1, __local_var_11_7, __local_var_12_8, (*Constructor_Data_Map_Internal_Node)(nil), (*Constructor_Data_Map_Internal_Node)(nil)}
goto end_branch_12
} else {

}
}
{
if (v1_14.Type == 9 && v1_14.IntVal == 324739070 && v1_14.UnsafePtr != nil) {
// TAST (Let): v2_15_10 -> gopurs_runtime.Value
v2_15_10 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_7, "compare"), __local_var_11_7, (*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V2)
_ = v2_15_10
var __t11 *Constructor_Data_Map_Internal_Node
{
if (uint32(v2_15_10.IntVal) == 1527465420) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_13_9_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V4)})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V5)}))
goto end_branch_11
} else {

}
}
{
if (uint32(v2_15_10.IntVal) == 380165415) {
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.UncurriedApp4(Get_Data_Map_Internal_unsafeBalancedNode(), (*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V2, (*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V4)}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_13_9_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V5)})))}))
goto end_branch_11
} else {

}
}
{
if (uint32(v2_15_10.IntVal) == 902936544) {
__t11 = &Constructor_Data_Map_Internal_Node{1, (*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V0, (*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V1, __local_var_11_7, __local_var_12_8, (*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V4, (*Constructor_Data_Map_Internal_Node)(v1_14.UnsafePtr).V5}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_11:
__t12 = __t11
goto end_branch_12
} else {

}
}
{
__t12 = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_12:
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(__t12)}
})
b_9_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Apply(go__go_13_9_1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](b_9))})))}
v_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_10.UnsafePtr).V1)}
continue go__go_8_6_0
__t13 = gopurs_runtime.Value{}
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
}
}()
})
})
// TAST (Let): fromFoldable_8_5 -> gopurs_runtime.Value
fromFoldable_8_5 := gopurs_runtime.Apply(go__go_8_6_0, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*Constructor_Data_Map_Internal_Node)(nil))})
_ = fromFoldable_8_5
return gopurs_runtime.Func(func(genKey_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(genValue_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(size_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V3), gopurs_runtime.Int(0), gopurs_runtime.Int(size_11.IntVal)), gopurs_runtime.Func(func(newSize_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Monad0_13_14 -> gopurs_runtime.Value
Monad0_13_14 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V0), gopurs_runtime.Value{})
_ = Monad0_13_14
// TAST (Let): pure_14_15 -> gopurs_runtime.Value
pure_14_15 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_13_14, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_14_15
// TAST (Let): Bind1_15_16 -> *Constructor_Control_Bind_Bind
Bind1_15_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_13_14, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_15_16
// TAST (Let): __local_var_16_17 -> gopurs_runtime.Value
__local_var_16_17 := gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_6_4.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), Get_Data_Tuple_Tuple(), genKey_9), genValue_10)
_ = __local_var_16_17
// TAST (Let): __local_var_17_24 -> gopurs_runtime.Value
__local_var_17_24 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(v_17 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 gopurs_runtime.Value
{
var __t27 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v_17.UnsafePtr).V1.IntVal) > (0) {
__t27 = false
goto end_branch_27
} else {

}
}
{
__t27 = true
}
end_branch_27:
if __t27 {
__t28 = gopurs_runtime.Apply(pure_14_15, gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons]((*Constructor_Data_Tuple_Tuple)(v_17.UnsafePtr).V0))}})})
goto end_branch_28
} else {

}
}
{
// TAST (Let): __local_var_18_25 -> gopurs_runtime.Value
__local_var_18_25 := (*Constructor_Data_Tuple_Tuple)(v_17.UnsafePtr).V0
_ = __local_var_18_25
// TAST (Let): __local_var_19_26 -> gopurs_runtime.Value
__local_var_19_26 := (*Constructor_Data_Tuple_Tuple)(v_17.UnsafePtr).V1
_ = __local_var_19_26
__t28 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_15_16.V1), __local_var_16_17, gopurs_runtime.Func(func(x_20 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_14_15, gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Gen_Cons{1, x_20, gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](__local_var_18_25)})}, gopurs_runtime.Int((__local_var_19_26.IntVal) - (1))})}})})
}))
}
end_branch_28:
return __t28
}))
_ = __local_var_17_24
// TAST (Let): __local_var_18_29 -> gopurs_runtime.Value
__local_var_18_29 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))})
_ = __local_var_18_29
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadGen_1.V4), gopurs_runtime.Func(func(v_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(newSize_12.IntVal)
}), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_3.V0), fromFoldable_8_5, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_13_14, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(b_17 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_18_18_2 gopurs_runtime.Value
go__go_18_18_2 = gopurs_runtime.Func(func(source_19_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_20_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_19_loop gopurs_runtime.Value = source_19_loop_val
var memo_20_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](memo_20_loop_val)
go__go_18_18_2:
for {
if false { continue go__go_18_18_2 }
var source_19 gopurs_runtime.Value = source_19_loop
_ = source_19
var memo_20 *Constructor_Data_List_Types_Cons = memo_20_loop
_ = memo_20
var __t20 *Constructor_Data_Maybe_Just
{
if (source_19.Type == 9 && source_19.IntVal == 759514854 && source_19.UnsafePtr == nil) {
__t20 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_20
} else {

}
}
{
if (source_19.Type == 9 && source_19.IntVal == 759514854 && source_19.UnsafePtr != nil) {
__t20 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Control_Monad_Gen_Cons)(source_19.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(source_19.UnsafePtr).V1)}})}}
goto end_branch_20
} else {

}
}
{
__t20 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_20:
// TAST (Let): v_21_19 -> *Constructor_Data_Maybe_Just
var v_21_19 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t20)})
var __t23 *Constructor_Data_List_Types_Cons
{
if (v_21_19 == nil) {
var go__go_22_21_3 gopurs_runtime.Value
go__go_22_21_3 = gopurs_runtime.Func(func(b_23_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_24_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_23_loop gopurs_runtime.Value = b_23_loop_val
var v_24_loop gopurs_runtime.Value = v_24_loop_val
go__go_22_21_3:
for {
if false { continue go__go_22_21_3 }
var b_23 gopurs_runtime.Value = b_23_loop
_ = b_23
var v_24 gopurs_runtime.Value = v_24_loop
_ = v_24
var __t22 gopurs_runtime.Value
{
if (v_24.Type == 9 && v_24.IntVal == 1358893437 && v_24.UnsafePtr == nil) {
__t22 = b_23
goto end_branch_22
} else {

}
}
{
if (v_24.Type == 9 && v_24.IntVal == 1358893437 && v_24.UnsafePtr != nil) {
b_23_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_24.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_23)})}
v_24_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_24.UnsafePtr).V1)}
continue go__go_22_21_3
__t22 = gopurs_runtime.Value{}
goto end_branch_22
} else {

}
}
{
__t22 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_22:
return __t22
}
}()
})
})
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_22_21_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(memo_20)}))
goto end_branch_23
} else {

}
}
{
if (v_21_19 != nil) {
source_19_loop = (*Constructor_Data_Tuple_Tuple)((v_21_19).V0.UnsafePtr).V1
memo_20_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_Tuple_Tuple)((v_21_19).V0.UnsafePtr).V0, memo_20})})
continue go__go_18_18_2
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t23)}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_18_18_2, b_17, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})))}
}), gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadGen_1.V5), gopurs_runtime.Func(func(x_19 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_17_24, gopurs_runtime.Apply(__local_var_18_29, x_19))
})))))
}))
}))
})
})
})
}



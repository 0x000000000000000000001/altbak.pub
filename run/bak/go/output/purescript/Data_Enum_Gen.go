package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Enum_Gen_foldable1NonEmpty gopurs_runtime.Value
var once_Data_Enum_Gen_foldable1NonEmpty sync.Once
func Get_Data_Enum_Gen_foldable1NonEmpty() gopurs_runtime.Value {
	once_Data_Enum_Gen_foldable1NonEmpty.Do(func() {
		cache_Data_Enum_Gen_foldable1NonEmpty = func() gopurs_runtime.Value {
// TAST (Let): foldableNonEmpty1_0_0 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_0_0 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_1_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_2
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_1.V0), gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_2.V0), gopurs_runtime.Apply(f_2, x_5), acc_6)
})
}), gopurs_runtime.RecordGet(dictMonoid_0, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray5 := (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1
_ = arr_val_foldlArray5
res_go_foldlArray5 := gopurs_runtime.Apply2(f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0)
_ = res_go_foldlArray5
arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
_ = arr_go_foldlArray5
for _, v_foldlArray5 := range *arr_go_foldlArray5 {
res_go_foldlArray5 = gopurs_runtime.Apply2(f_0, res_go_foldlArray5, v_foldlArray5)
}
return res_go_foldlArray5
}()
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_0, b_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_0_0)}
}), gopurs_runtime.Func(func(dictSemigroup_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray5 := (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V1
_ = arr_val_foldlArray5
res_go_foldlArray5 := gopurs_runtime.Apply(f_2, (*Constructor_Data_NonEmpty_NonEmpty)(v_3.UnsafePtr).V0)
_ = res_go_foldlArray5
arr_go_foldlArray5 := (*[]gopurs_runtime.Value)(arr_val_foldlArray5.UnsafePtr)
_ = arr_go_foldlArray5
for _, v_foldlArray5 := range *arr_go_foldlArray5 {
res_go_foldlArray5 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), s_4, gopurs_runtime.Apply(f_2, a1_5))
})
}), res_go_foldlArray5, v_foldlArray5)
}
return res_go_foldlArray5
}()
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray4 := (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1
_ = arr_val_foldlArray4
res_go_foldlArray4 := (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
_ = res_go_foldlArray4
arr_go_foldlArray4 := (*[]gopurs_runtime.Value)(arr_val_foldlArray4.UnsafePtr)
_ = arr_go_foldlArray4
for _, v_foldlArray4 := range *arr_go_foldlArray4 {
res_go_foldlArray4 = gopurs_runtime.Apply2(f_1, res_go_foldlArray4, v_foldlArray4)
}
return res_go_foldlArray4
}()
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_3 -> gopurs_runtime.Value
__local_var_3_3 := gopurs_runtime.Apply(f_1, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0)
_ = __local_var_3_3
// TAST (Let): __local_var_4_4 -> *Constructor_Data_Maybe_Just
__local_var_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_6 -> gopurs_runtime.Value
__local_var_5_6 := gopurs_runtime.Apply(f_1, a1_4)
_ = __local_var_5_6
// TAST (Let): __local_var_5_5 -> gopurs_runtime.Value
__local_var_5_5 := gopurs_runtime.Func(func(v2_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 gopurs_runtime.Value
{
if (v2_6.Type == 9 && v2_6.IntVal == 930809136 && v2_6.UnsafePtr == nil) {
__t7 = a1_4
goto end_branch_7
} else {

}
}
{
if (v2_6.Type == 9 && v2_6.IntVal == 930809136 && v2_6.UnsafePtr != nil) {
__t7 = gopurs_runtime.Apply(__local_var_5_6, (*Constructor_Data_Maybe_Just)(v2_6.UnsafePtr).V0)
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
})
_ = __local_var_5_5
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_5_5, x_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V1))
_ = __local_var_4_4
var __t8 gopurs_runtime.Value
{
if (__local_var_4_4 == nil) {
__t8 = (*Constructor_Data_NonEmpty_NonEmpty)(v_2.UnsafePtr).V0
goto end_branch_8
} else {

}
}
{
if (__local_var_4_4 != nil) {
__t8 = gopurs_runtime.Apply(__local_var_3_3, (__local_var_4_4).V0)
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
})
})})}
}()
	})
	return cache_Data_Enum_Gen_foldable1NonEmpty
}

var cache_Data_Enum_Gen_genBoundedEnum gopurs_runtime.Value
var once_Data_Enum_Gen_genBoundedEnum sync.Once
func Get_Data_Enum_Gen_genBoundedEnum() gopurs_runtime.Value {
	once_Data_Enum_Gen_genBoundedEnum.Do(func() {
		cache_Data_Enum_Gen_genBoundedEnum = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0_box)
})
	})
	return cache_Data_Enum_Gen_genBoundedEnum
}

func Call_Data_Enum_Gen_genBoundedEnum(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_1_0
return gopurs_runtime.Func(func(dictBoundedEnum_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Enum1_3_1 -> *Constructor_Data_Enum_Enum
Enum1_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Enum_Enum](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Enum1"), gopurs_runtime.Value{}))
_ = Enum1_3_1
// TAST (Let): Bounded0_4_2 -> *Constructor_Data_Bounded_Bounded
Bounded0_4_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Bounded_Bounded](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBoundedEnum_2, "Bounded0"), gopurs_runtime.Value{}))
_ = Bounded0_4_2
// TAST (Let): v_5_3 -> gopurs_runtime.Value
v_5_3 := gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_3_1.V2), gopurs_runtime.Box(Bounded0_4_2.V1))
_ = v_5_3
var __t44 gopurs_runtime.Value
{
if (v_5_3.Type == 9 && v_5_3.IntVal == 930809136 && v_5_3.UnsafePtr != nil) {
// TAST (Let): Monad0_6_4 -> gopurs_runtime.Value
Monad0_6_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_6_4
// TAST (Let): pure_7_5 -> gopurs_runtime.Value
pure_7_5 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_6_4, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_7_5
// TAST (Let): foldableNonEmpty1_8_7 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_8_7 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_8 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_8
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_12_9 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_12_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_8, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_12_9
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_8.V0), gopurs_runtime.Apply(f_10, (*Constructor_Data_NonEmpty_NonEmpty)(v_11.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_12_9.V0), gopurs_runtime.Apply(f_10, x_13), acc_14)
})
}), gopurs_runtime.RecordGet(dictMonoid_8, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_11.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray14 := (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V1
_ = arr_val_foldlArray14
res_go_foldlArray14 := gopurs_runtime.Apply2(f_8, b_9, (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V0)
_ = res_go_foldlArray14
arr_go_foldlArray14 := (*[]gopurs_runtime.Value)(arr_val_foldlArray14.UnsafePtr)
_ = arr_go_foldlArray14
for _, v_foldlArray14 := range *arr_go_foldlArray14 {
res_go_foldlArray14 = gopurs_runtime.Apply2(f_8, res_go_foldlArray14, v_foldlArray14)
}
return res_go_foldlArray14
}()
})
})
}), gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_8, b_9, (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_8_7
// TAST (Let): __local_var_8_6 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_8_6 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_8_7)}
}), gopurs_runtime.Func(func(dictSemigroup_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray14 := (*Constructor_Data_NonEmpty_NonEmpty)(v_11.UnsafePtr).V1
_ = arr_val_foldlArray14
res_go_foldlArray14 := gopurs_runtime.Apply(f_10, (*Constructor_Data_NonEmpty_NonEmpty)(v_11.UnsafePtr).V0)
_ = res_go_foldlArray14
arr_go_foldlArray14 := (*[]gopurs_runtime.Value)(arr_val_foldlArray14.UnsafePtr)
_ = arr_go_foldlArray14
for _, v_foldlArray14 := range *arr_go_foldlArray14 {
res_go_foldlArray14 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_13 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_9, "append"), s_12, gopurs_runtime.Apply(f_10, a1_13))
})
}), res_go_foldlArray14, v_foldlArray14)
}
return res_go_foldlArray14
}()
})
})
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V1
_ = arr_val_foldlArray13
res_go_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V0
_ = res_go_foldlArray13
arr_go_foldlArray13 := (*[]gopurs_runtime.Value)(arr_val_foldlArray13.UnsafePtr)
_ = arr_go_foldlArray13
for _, v_foldlArray13 := range *arr_go_foldlArray13 {
res_go_foldlArray13 = gopurs_runtime.Apply2(f_9, res_go_foldlArray13, v_foldlArray13)
}
return res_go_foldlArray13
}()
})
}), gopurs_runtime.Func(func(f_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_10 -> gopurs_runtime.Value
__local_var_11_10 := gopurs_runtime.Apply(f_9, (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V0)
_ = __local_var_11_10
// TAST (Let): __local_var_12_11 -> *Constructor_Data_Maybe_Just
__local_var_12_11 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_12 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_13_13 -> gopurs_runtime.Value
__local_var_13_13 := gopurs_runtime.Apply(f_9, a1_12)
_ = __local_var_13_13
// TAST (Let): __local_var_13_12 -> gopurs_runtime.Value
__local_var_13_12 := gopurs_runtime.Func(func(v2_14 gopurs_runtime.Value) gopurs_runtime.Value {
var __t14 gopurs_runtime.Value
{
if (v2_14.Type == 9 && v2_14.IntVal == 930809136 && v2_14.UnsafePtr == nil) {
__t14 = a1_12
goto end_branch_14
} else {

}
}
{
if (v2_14.Type == 9 && v2_14.IntVal == 930809136 && v2_14.UnsafePtr != nil) {
__t14 = gopurs_runtime.Apply(__local_var_13_13, (*Constructor_Data_Maybe_Just)(v2_14.UnsafePtr).V0)
goto end_branch_14
} else {

}
}
{
__t14 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_14:
return __t14
})
_ = __local_var_13_12
return gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_13_12, x_14)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V1))
_ = __local_var_12_11
var __t15 gopurs_runtime.Value
{
if (__local_var_12_11 == nil) {
__t15 = (*Constructor_Data_NonEmpty_NonEmpty)(v_10.UnsafePtr).V0
goto end_branch_15
} else {

}
}
{
if (__local_var_12_11 != nil) {
__t15 = gopurs_runtime.Apply(__local_var_11_10, (__local_var_12_11).V0)
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_15:
return __t15
})
})}
_ = __local_var_8_6
// TAST (Let): Ord0_9_17 -> gopurs_runtime.Value
Ord0_9_17 := gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_3_1.V0), gopurs_runtime.Value{})
_ = Ord0_9_17
// TAST (Let): Ord01_10_18 -> *Constructor_Data_Ord_Ord
Ord01_10_18 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_3_1.V0), gopurs_runtime.Value{}))
_ = Ord01_10_18
var __t38 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Ord0_9_17, "Eq0"), gopurs_runtime.Value{}), "eq"), (*Constructor_Data_Maybe_Just)(v_5_3.UnsafePtr).V0, gopurs_runtime.Box(Bounded0_4_2.V2)).IntVal) != (0) {
__t38 = gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t26 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t26 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
goto end_branch_26
} else {

}
}
{
__t26 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_26:
return __t26
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(i_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t28 *Constructor_Data_Tuple_Tuple
{
var __t27 bool
{
if (i_11.IntVal) > (0) {
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
__t28 = &Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Maybe_Just)(v_5_3.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}}
goto end_branch_28
} else {

}
}
{
__t28 = &Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Maybe_Just)(v_5_3.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Int((i_11.IntVal) - (1))})}}
}
end_branch_28:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__t28)}
}), gopurs_runtime.Int(0))
goto end_branch_38
} else {

}
}
{
var __t30 bool
{
var __t_tag_29 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.Box(Ord01_10_18.V1), (*Constructor_Data_Maybe_Just)(v_5_3.UnsafePtr).V0, gopurs_runtime.Box(Bounded0_4_2.V2))
if (uint32(__t_tag_29.IntVal) == 1527465420) {
__t30 = true
goto end_branch_30
} else {

}
}
{
__t30 = false
}
end_branch_30:
if __t30 {
__t38 = gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t31 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t31 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
goto end_branch_31
} else {

}
}
{
__t31 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_31:
return __t31
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_32 -> *Constructor_Data_Maybe_Just
__local_var_12_32 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_3_1.V2), a_11))
_ = __local_var_12_32
var __t37 *Constructor_Data_Maybe_Just
{
if (__local_var_12_32 != nil) {
var __t36 *Constructor_Data_Maybe_Just
{
var __t35 *Constructor_Data_Maybe_Just
{
var __t34 bool
{
var __t_tag_33 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_9_17, "compare"), (__local_var_12_32).V0, gopurs_runtime.Box(Bounded0_4_2.V2))
if (uint32(__t_tag_33.IntVal) == 380165415) {
__t34 = false
goto end_branch_34
} else {

}
}
{
__t34 = true
}
end_branch_34:
if __t34 {
__t35 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_35
} else {

}
}
{
__t35 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_35:
if (__t35 != nil) {
__t36 = &Constructor_Data_Maybe_Just{1, (__local_var_12_32).V0}
goto end_branch_36
} else {

}
}
{
__t36 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_36:
__t37 = __t36
goto end_branch_37
} else {

}
}
{
if (__local_var_12_32 == nil) {
__t37 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_37
} else {

}
}
{
__t37 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_37:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_11, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t37)}})}
}), (*Constructor_Data_Maybe_Just)(v_5_3.UnsafePtr).V0)
goto end_branch_38
} else {

}
}
{
__t38 = gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t19 gopurs_runtime.Value
{
if (v_12.Type == 9 && v_12.IntVal == 930809136 && v_12.UnsafePtr != nil) {
__t19 = (*Constructor_Data_Maybe_Just)(v_12.UnsafePtr).V0
goto end_branch_19
} else {

}
}
{
__t19 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_19:
return __t19
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_11 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_12_20 -> *Constructor_Data_Maybe_Just
__local_var_12_20 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(gopurs_runtime.Box(Enum1_3_1.V1), a_11))
_ = __local_var_12_20
var __t25 *Constructor_Data_Maybe_Just
{
if (__local_var_12_20 != nil) {
var __t24 *Constructor_Data_Maybe_Just
{
var __t23 *Constructor_Data_Maybe_Just
{
var __t22 bool
{
var __t_tag_21 gopurs_runtime.Value = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Ord0_9_17, "compare"), (__local_var_12_20).V0, gopurs_runtime.Box(Bounded0_4_2.V2))
if (uint32(__t_tag_21.IntVal) == 1527465420) {
__t22 = false
goto end_branch_22
} else {

}
}
{
__t22 = true
}
end_branch_22:
if __t22 {
__t23 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_23
} else {

}
}
{
__t23 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_23:
if (__t23 != nil) {
__t24 = &Constructor_Data_Maybe_Just{1, (__local_var_12_20).V0}
goto end_branch_24
} else {

}
}
{
__t24 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_24:
__t25 = __t24
goto end_branch_25
} else {

}
}
{
if (__local_var_12_20 == nil) {
__t25 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_25:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_11, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t25)}})}
}), (*Constructor_Data_Maybe_Just)(v_5_3.UnsafePtr).V0)
}
end_branch_38:
// TAST (Let): __local_var_9_16 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_9_16 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Box(Bounded0_4_2.V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t38.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
__t44 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_6_4, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_6.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_10.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_9_16)}).IntVal) - (1))), gopurs_runtime.Func(func(n_10 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_11_39_0 gopurs_runtime.Value
go__go_11_39_0 = gopurs_runtime.Func(func(v_12_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_13_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_12_loop int64 = v_12_loop_val.IntVal
var v1_13_loop gopurs_runtime.Value = v1_13_loop_val
go__go_11_39_0:
for {
if false { continue go__go_11_39_0 }
var v_12 int64 = v_12_loop
_ = v_12
var v1_13 gopurs_runtime.Value = v1_13_loop
_ = v1_13
var __t43 gopurs_runtime.Value
{
if (v1_13.Type == 9 && v1_13.IntVal == 759514854 && v1_13.UnsafePtr != nil) {
var __t42 gopurs_runtime.Value
{
var __t_tag_40 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_13.UnsafePtr).V1
if (__t_tag_40 == nil) {
__t42 = (*Constructor_Control_Monad_Gen_Cons)(v1_13.UnsafePtr).V0
goto end_branch_42
} else {

}
}
{
var __t41 bool
{
if (v_12) > (0) {
__t41 = false
goto end_branch_41
} else {

}
}
{
__t41 = true
}
end_branch_41:
if __t41 {
__t42 = (*Constructor_Control_Monad_Gen_Cons)(v1_13.UnsafePtr).V0
goto end_branch_42
} else {

}
}
{
v_12_loop = (v_12) - (1)
v1_13_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_13.UnsafePtr).V1)}
continue go__go_11_39_0
__t42 = gopurs_runtime.Value{}
}
end_branch_42:
__t43 = __t42
goto end_branch_43
} else {

}
}
{
if (v1_13.Type == 9 && v1_13.IntVal == 759514854 && v1_13.UnsafePtr == nil) {
__t43 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_8_6.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_14 gopurs_runtime.Value) gopurs_runtime.Value {
return x_14
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_9_16)})
goto end_branch_43
} else {

}
}
{
__t43 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_43:
return __t43
}
}()
})
})
return gopurs_runtime.Apply(pure_7_5, gopurs_runtime.Apply2(go__go_11_39_0, gopurs_runtime.Int(n_10.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_8_6.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_9_16)})))}))
}))
goto end_branch_44
} else {

}
}
{
if (v_5_3.Type == 9 && v_5_3.IntVal == 930809136 && v_5_3.UnsafePtr == nil) {
__t44 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), gopurs_runtime.Box(Bounded0_4_2.V1))
goto end_branch_44
} else {

}
}
{
__t44 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_44:
return __t44
})
}



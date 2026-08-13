package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Time_Gen_genTime gopurs_runtime.Value
var once_Data_Time_Gen_genTime sync.Once
func Get_Data_Time_Gen_genTime() gopurs_runtime.Value {
	once_Data_Time_Gen_genTime.Do(func() {
		cache_Data_Time_Gen_genTime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Gen_genTime(dictMonadGen_0_box)
})
	})
	return cache_Data_Time_Gen_genTime
}

func Call_Data_Time_Gen_genTime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 -> gopurs_runtime.Value
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
// TAST (Let): Apply0_2_1 -> *Constructor_Control_Apply_Apply
Apply0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_2_1
// TAST (Let): Monad0_3_2 -> gopurs_runtime.Value
Monad0_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_2
// TAST (Let): pure_4_3 -> gopurs_runtime.Value
pure_4_3 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_2, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_3
// TAST (Let): foldableNonEmpty1_5_5 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_5_5 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_6 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_6
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_7 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_9_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_7
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_6.V0), gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_7.V0), gopurs_runtime.Apply(f_7, x_10), acc_11)
})
}), gopurs_runtime.RecordGet(dictMonoid_5, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray14 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray14
res_go_foldlArray14 := gopurs_runtime.Apply2(f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = res_go_foldlArray14
arr_go_foldlArray14 := (*[]gopurs_runtime.Value)(arr_val_foldlArray14.UnsafePtr)
_ = arr_go_foldlArray14
for _, v_foldlArray14 := range *arr_go_foldlArray14 {
res_go_foldlArray14 = gopurs_runtime.Apply2(f_5, res_go_foldlArray14, v_foldlArray14)
}
return res_go_foldlArray14
}()
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_5_5
// TAST (Let): __local_var_5_4 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_5_4 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_5_5)}
}), gopurs_runtime.Func(func(dictSemigroup_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray14 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray14
res_go_foldlArray14 := gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = res_go_foldlArray14
arr_go_foldlArray14 := (*[]gopurs_runtime.Value)(arr_val_foldlArray14.UnsafePtr)
_ = arr_go_foldlArray14
for _, v_foldlArray14 := range *arr_go_foldlArray14 {
res_go_foldlArray14 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_6, "append"), s_9, gopurs_runtime.Apply(f_7, a1_10))
})
}), res_go_foldlArray14, v_foldlArray14)
}
return res_go_foldlArray14
}()
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray13
res_go_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
_ = res_go_foldlArray13
arr_go_foldlArray13 := (*[]gopurs_runtime.Value)(arr_val_foldlArray13.UnsafePtr)
_ = arr_go_foldlArray13
for _, v_foldlArray13 := range *arr_go_foldlArray13 {
res_go_foldlArray13 = gopurs_runtime.Apply2(f_6, res_go_foldlArray13, v_foldlArray13)
}
return res_go_foldlArray13
}()
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = __local_var_8_8
// TAST (Let): __local_var_9_9 -> *Constructor_Data_Maybe_Just
__local_var_9_9 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_11 -> gopurs_runtime.Value
__local_var_10_11 := gopurs_runtime.Apply(f_6, a1_9)
_ = __local_var_10_11
// TAST (Let): __local_var_10_10 -> gopurs_runtime.Value
__local_var_10_10 := gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t12 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr == nil) {
__t12 = a1_9
goto end_branch_12
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr != nil) {
__t12 = gopurs_runtime.Apply(__local_var_10_11, (*Constructor_Data_Maybe_Just)(v2_11.UnsafePtr).V0)
goto end_branch_12
} else {

}
}
{
__t12 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_12:
return __t12
})
_ = __local_var_10_10
return gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_10_10, x_11)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
_ = __local_var_9_9
var __t13 gopurs_runtime.Value
{
if (__local_var_9_9 == nil) {
__t13 = (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
goto end_branch_13
} else {

}
}
{
if (__local_var_9_9 != nil) {
__t13 = gopurs_runtime.Apply(__local_var_8_8, (__local_var_9_9).V0)
goto end_branch_13
} else {

}
}
{
__t13 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_13:
return __t13
})
})}
_ = __local_var_5_4
// TAST (Let): __local_var_6_14 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_6_14 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t15 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__t15 = (*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_17 -> gopurs_runtime.Value
__local_var_7_17 := gopurs_runtime.Int((a_6.IntVal) + (1))
_ = __local_var_7_17
var __t21 *Constructor_Data_Maybe_Just
{
var __t18 bool
{
if (__local_var_7_17.IntVal) < (0) {
__t18 = false
goto end_branch_18
} else {

}
}
{
__t18 = true
}
end_branch_18:
var __t_and_20 bool = false
if __t18 {

var __t19 bool
{
if (__local_var_7_17.IntVal) > (23) {
__t19 = false
goto end_branch_19
} else {

}
}
{
__t19 = true
}
end_branch_19:
__t_and_20 = __t19
}
if __t_and_20 {
__t21 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_7_17.IntVal)}
goto end_branch_21
} else {

}
}
{
__t21 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_21:
// TAST (Let): __local_var_7_16 -> *Constructor_Data_Maybe_Just
var __local_var_7_16 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t21)})
var __t25 *Constructor_Data_Maybe_Just
{
if (__local_var_7_16 != nil) {
var __t24 *Constructor_Data_Maybe_Just
{
var __t23 *Constructor_Data_Maybe_Just
{
var __t22 bool
{
if ((__local_var_7_16).V0.IntVal) > (gopurs_runtime.Int(23).IntVal) {
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
__t24 = &Constructor_Data_Maybe_Just{1, (__local_var_7_16).V0}
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
if (__local_var_7_16 == nil) {
__t25 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_25
} else {

}
}
{
__t25 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_25:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t25)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
// TAST (Let): Monad0_3_31 -> gopurs_runtime.Value
Monad0_3_31 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_31
// TAST (Let): pure_4_32 -> gopurs_runtime.Value
pure_4_32 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_31, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_32
// TAST (Let): foldableNonEmpty1_5_34 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_5_34 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_35 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_35 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_35
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_36 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_9_36 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_36
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_35.V0), gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_36.V0), gopurs_runtime.Apply(f_7, x_10), acc_11)
})
}), gopurs_runtime.RecordGet(dictMonoid_5, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray13
res_go_foldlArray13 := gopurs_runtime.Apply2(f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = res_go_foldlArray13
arr_go_foldlArray13 := (*[]gopurs_runtime.Value)(arr_val_foldlArray13.UnsafePtr)
_ = arr_go_foldlArray13
for _, v_foldlArray13 := range *arr_go_foldlArray13 {
res_go_foldlArray13 = gopurs_runtime.Apply2(f_5, res_go_foldlArray13, v_foldlArray13)
}
return res_go_foldlArray13
}()
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_5_34
// TAST (Let): __local_var_5_33 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_5_33 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_5_34)}
}), gopurs_runtime.Func(func(dictSemigroup_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray13
res_go_foldlArray13 := gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = res_go_foldlArray13
arr_go_foldlArray13 := (*[]gopurs_runtime.Value)(arr_val_foldlArray13.UnsafePtr)
_ = arr_go_foldlArray13
for _, v_foldlArray13 := range *arr_go_foldlArray13 {
res_go_foldlArray13 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_6, "append"), s_9, gopurs_runtime.Apply(f_7, a1_10))
})
}), res_go_foldlArray13, v_foldlArray13)
}
return res_go_foldlArray13
}()
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray12 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray12
res_go_foldlArray12 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
_ = res_go_foldlArray12
arr_go_foldlArray12 := (*[]gopurs_runtime.Value)(arr_val_foldlArray12.UnsafePtr)
_ = arr_go_foldlArray12
for _, v_foldlArray12 := range *arr_go_foldlArray12 {
res_go_foldlArray12 = gopurs_runtime.Apply2(f_6, res_go_foldlArray12, v_foldlArray12)
}
return res_go_foldlArray12
}()
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_37 -> gopurs_runtime.Value
__local_var_8_37 := gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = __local_var_8_37
// TAST (Let): __local_var_9_38 -> *Constructor_Data_Maybe_Just
__local_var_9_38 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_40 -> gopurs_runtime.Value
__local_var_10_40 := gopurs_runtime.Apply(f_6, a1_9)
_ = __local_var_10_40
// TAST (Let): __local_var_10_39 -> gopurs_runtime.Value
__local_var_10_39 := gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t41 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr == nil) {
__t41 = a1_9
goto end_branch_41
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr != nil) {
__t41 = gopurs_runtime.Apply(__local_var_10_40, (*Constructor_Data_Maybe_Just)(v2_11.UnsafePtr).V0)
goto end_branch_41
} else {

}
}
{
__t41 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_41:
return __t41
})
_ = __local_var_10_39
return gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_10_39, x_11)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
_ = __local_var_9_38
var __t42 gopurs_runtime.Value
{
if (__local_var_9_38 == nil) {
__t42 = (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
goto end_branch_42
} else {

}
}
{
if (__local_var_9_38 != nil) {
__t42 = gopurs_runtime.Apply(__local_var_8_37, (__local_var_9_38).V0)
goto end_branch_42
} else {

}
}
{
__t42 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_42:
return __t42
})
})}
_ = __local_var_5_33
// TAST (Let): __local_var_6_43 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_6_43 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t44 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__t44 = (*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_46 -> gopurs_runtime.Value
__local_var_7_46 := gopurs_runtime.Int((a_6.IntVal) + (1))
_ = __local_var_7_46
var __t50 *Constructor_Data_Maybe_Just
{
var __t47 bool
{
if (__local_var_7_46.IntVal) < (0) {
__t47 = false
goto end_branch_47
} else {

}
}
{
__t47 = true
}
end_branch_47:
var __t_and_49 bool = false
if __t47 {

var __t48 bool
{
if (__local_var_7_46.IntVal) > (59) {
__t48 = false
goto end_branch_48
} else {

}
}
{
__t48 = true
}
end_branch_48:
__t_and_49 = __t48
}
if __t_and_49 {
__t50 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_7_46.IntVal)}
goto end_branch_50
} else {

}
}
{
__t50 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_50:
// TAST (Let): __local_var_7_45 -> *Constructor_Data_Maybe_Just
var __local_var_7_45 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t50)})
var __t54 *Constructor_Data_Maybe_Just
{
if (__local_var_7_45 != nil) {
var __t53 *Constructor_Data_Maybe_Just
{
var __t52 *Constructor_Data_Maybe_Just
{
var __t51 bool
{
if ((__local_var_7_45).V0.IntVal) > (gopurs_runtime.Int(59).IntVal) {
__t51 = false
goto end_branch_51
} else {

}
}
{
__t51 = true
}
end_branch_51:
if __t51 {
__t52 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_52
} else {

}
}
{
__t52 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_52:
if (__t52 != nil) {
__t53 = &Constructor_Data_Maybe_Just{1, (__local_var_7_45).V0}
goto end_branch_53
} else {

}
}
{
__t53 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_53:
__t54 = __t53
goto end_branch_54
} else {

}
}
{
if (__local_var_7_45 == nil) {
__t54 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_54
} else {

}
}
{
__t54 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_54:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t54)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
// TAST (Let): Monad0_3_60 -> gopurs_runtime.Value
Monad0_3_60 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_60
// TAST (Let): pure_4_61 -> gopurs_runtime.Value
pure_4_61 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_60, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_61
// TAST (Let): foldableNonEmpty1_5_63 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_5_63 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_64 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_64 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_64
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_65 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_9_65 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_65
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_64.V0), gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_65.V0), gopurs_runtime.Apply(f_7, x_10), acc_11)
})
}), gopurs_runtime.RecordGet(dictMonoid_5, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray12 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray12
res_go_foldlArray12 := gopurs_runtime.Apply2(f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = res_go_foldlArray12
arr_go_foldlArray12 := (*[]gopurs_runtime.Value)(arr_val_foldlArray12.UnsafePtr)
_ = arr_go_foldlArray12
for _, v_foldlArray12 := range *arr_go_foldlArray12 {
res_go_foldlArray12 = gopurs_runtime.Apply2(f_5, res_go_foldlArray12, v_foldlArray12)
}
return res_go_foldlArray12
}()
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_5_63
// TAST (Let): __local_var_5_62 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_5_62 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_5_63)}
}), gopurs_runtime.Func(func(dictSemigroup_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray12 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray12
res_go_foldlArray12 := gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = res_go_foldlArray12
arr_go_foldlArray12 := (*[]gopurs_runtime.Value)(arr_val_foldlArray12.UnsafePtr)
_ = arr_go_foldlArray12
for _, v_foldlArray12 := range *arr_go_foldlArray12 {
res_go_foldlArray12 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_6, "append"), s_9, gopurs_runtime.Apply(f_7, a1_10))
})
}), res_go_foldlArray12, v_foldlArray12)
}
return res_go_foldlArray12
}()
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray11 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray11
res_go_foldlArray11 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
_ = res_go_foldlArray11
arr_go_foldlArray11 := (*[]gopurs_runtime.Value)(arr_val_foldlArray11.UnsafePtr)
_ = arr_go_foldlArray11
for _, v_foldlArray11 := range *arr_go_foldlArray11 {
res_go_foldlArray11 = gopurs_runtime.Apply2(f_6, res_go_foldlArray11, v_foldlArray11)
}
return res_go_foldlArray11
}()
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_66 -> gopurs_runtime.Value
__local_var_8_66 := gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = __local_var_8_66
// TAST (Let): __local_var_9_67 -> *Constructor_Data_Maybe_Just
__local_var_9_67 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_69 -> gopurs_runtime.Value
__local_var_10_69 := gopurs_runtime.Apply(f_6, a1_9)
_ = __local_var_10_69
// TAST (Let): __local_var_10_68 -> gopurs_runtime.Value
__local_var_10_68 := gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t70 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr == nil) {
__t70 = a1_9
goto end_branch_70
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr != nil) {
__t70 = gopurs_runtime.Apply(__local_var_10_69, (*Constructor_Data_Maybe_Just)(v2_11.UnsafePtr).V0)
goto end_branch_70
} else {

}
}
{
__t70 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_70:
return __t70
})
_ = __local_var_10_68
return gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_10_68, x_11)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
_ = __local_var_9_67
var __t71 gopurs_runtime.Value
{
if (__local_var_9_67 == nil) {
__t71 = (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
goto end_branch_71
} else {

}
}
{
if (__local_var_9_67 != nil) {
__t71 = gopurs_runtime.Apply(__local_var_8_66, (__local_var_9_67).V0)
goto end_branch_71
} else {

}
}
{
__t71 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_71:
return __t71
})
})}
_ = __local_var_5_62
// TAST (Let): __local_var_6_72 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_6_72 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t73 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__t73 = (*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0
goto end_branch_73
} else {

}
}
{
__t73 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_73:
return __t73
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_75 -> gopurs_runtime.Value
__local_var_7_75 := gopurs_runtime.Int((a_6.IntVal) + (1))
_ = __local_var_7_75
var __t79 *Constructor_Data_Maybe_Just
{
var __t76 bool
{
if (__local_var_7_75.IntVal) < (0) {
__t76 = false
goto end_branch_76
} else {

}
}
{
__t76 = true
}
end_branch_76:
var __t_and_78 bool = false
if __t76 {

var __t77 bool
{
if (__local_var_7_75.IntVal) > (59) {
__t77 = false
goto end_branch_77
} else {

}
}
{
__t77 = true
}
end_branch_77:
__t_and_78 = __t77
}
if __t_and_78 {
__t79 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_7_75.IntVal)}
goto end_branch_79
} else {

}
}
{
__t79 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_79:
// TAST (Let): __local_var_7_74 -> *Constructor_Data_Maybe_Just
var __local_var_7_74 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t79)})
var __t83 *Constructor_Data_Maybe_Just
{
if (__local_var_7_74 != nil) {
var __t82 *Constructor_Data_Maybe_Just
{
var __t81 *Constructor_Data_Maybe_Just
{
var __t80 bool
{
if ((__local_var_7_74).V0.IntVal) > (gopurs_runtime.Int(59).IntVal) {
__t80 = false
goto end_branch_80
} else {

}
}
{
__t80 = true
}
end_branch_80:
if __t80 {
__t81 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_81
} else {

}
}
{
__t81 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_81:
if (__t81 != nil) {
__t82 = &Constructor_Data_Maybe_Just{1, (__local_var_7_74).V0}
goto end_branch_82
} else {

}
}
{
__t82 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_82:
__t83 = __t82
goto end_branch_83
} else {

}
}
{
if (__local_var_7_74 == nil) {
__t83 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_83
} else {

}
}
{
__t83 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_83:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t83)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
// TAST (Let): Monad0_3_89 -> gopurs_runtime.Value
Monad0_3_89 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_3_89
// TAST (Let): pure_4_90 -> gopurs_runtime.Value
pure_4_90 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_89, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_4_90
// TAST (Let): foldableNonEmpty1_5_92 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_5_92 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_6_93 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_6_93 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_6_93
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_9_94 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_9_94 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_5, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_9_94
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_6_93.V0), gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_9_94.V0), gopurs_runtime.Apply(f_7, x_10), acc_11)
})
}), gopurs_runtime.RecordGet(dictMonoid_5, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray11 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray11
res_go_foldlArray11 := gopurs_runtime.Apply2(f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = res_go_foldlArray11
arr_go_foldlArray11 := (*[]gopurs_runtime.Value)(arr_val_foldlArray11.UnsafePtr)
_ = arr_go_foldlArray11
for _, v_foldlArray11 := range *arr_go_foldlArray11 {
res_go_foldlArray11 = gopurs_runtime.Apply2(f_5, res_go_foldlArray11, v_foldlArray11)
}
return res_go_foldlArray11
}()
})
})
}), gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_5, b_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_5_92
// TAST (Let): __local_var_5_91 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_5_91 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_5_92)}
}), gopurs_runtime.Func(func(dictSemigroup_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray11 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray11
res_go_foldlArray11 := gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = res_go_foldlArray11
arr_go_foldlArray11 := (*[]gopurs_runtime.Value)(arr_val_foldlArray11.UnsafePtr)
_ = arr_go_foldlArray11
for _, v_foldlArray11 := range *arr_go_foldlArray11 {
res_go_foldlArray11 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_6, "append"), s_9, gopurs_runtime.Apply(f_7, a1_10))
})
}), res_go_foldlArray11, v_foldlArray11)
}
return res_go_foldlArray11
}()
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray10 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1
_ = arr_val_foldlArray10
res_go_foldlArray10 := (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
_ = res_go_foldlArray10
arr_go_foldlArray10 := (*[]gopurs_runtime.Value)(arr_val_foldlArray10.UnsafePtr)
_ = arr_go_foldlArray10
for _, v_foldlArray10 := range *arr_go_foldlArray10 {
res_go_foldlArray10 = gopurs_runtime.Apply2(f_6, res_go_foldlArray10, v_foldlArray10)
}
return res_go_foldlArray10
}()
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_95 -> gopurs_runtime.Value
__local_var_8_95 := gopurs_runtime.Apply(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0)
_ = __local_var_8_95
// TAST (Let): __local_var_9_96 -> *Constructor_Data_Maybe_Just
__local_var_9_96 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_10_98 -> gopurs_runtime.Value
__local_var_10_98 := gopurs_runtime.Apply(f_6, a1_9)
_ = __local_var_10_98
// TAST (Let): __local_var_10_97 -> gopurs_runtime.Value
__local_var_10_97 := gopurs_runtime.Func(func(v2_11 gopurs_runtime.Value) gopurs_runtime.Value {
var __t99 gopurs_runtime.Value
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr == nil) {
__t99 = a1_9
goto end_branch_99
} else {

}
}
{
if (v2_11.Type == 9 && v2_11.IntVal == 930809136 && v2_11.UnsafePtr != nil) {
__t99 = gopurs_runtime.Apply(__local_var_10_98, (*Constructor_Data_Maybe_Just)(v2_11.UnsafePtr).V0)
goto end_branch_99
} else {

}
}
{
__t99 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_99:
return __t99
})
_ = __local_var_10_97
return gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_10_97, x_11)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V1))
_ = __local_var_9_96
var __t100 gopurs_runtime.Value
{
if (__local_var_9_96 == nil) {
__t100 = (*Constructor_Data_NonEmpty_NonEmpty)(v_7.UnsafePtr).V0
goto end_branch_100
} else {

}
}
{
if (__local_var_9_96 != nil) {
__t100 = gopurs_runtime.Apply(__local_var_8_95, (__local_var_9_96).V0)
goto end_branch_100
} else {

}
}
{
__t100 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_100:
return __t100
})
})}
_ = __local_var_5_91
// TAST (Let): __local_var_6_101 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_6_101 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t102 gopurs_runtime.Value
{
if (v_7.Type == 9 && v_7.IntVal == 930809136 && v_7.UnsafePtr != nil) {
__t102 = (*Constructor_Data_Maybe_Just)(v_7.UnsafePtr).V0
goto end_branch_102
} else {

}
}
{
__t102 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_102:
return __t102
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_104 -> gopurs_runtime.Value
__local_var_7_104 := gopurs_runtime.Int((a_6.IntVal) + (1))
_ = __local_var_7_104
var __t108 *Constructor_Data_Maybe_Just
{
var __t105 bool
{
if (__local_var_7_104.IntVal) < (0) {
__t105 = false
goto end_branch_105
} else {

}
}
{
__t105 = true
}
end_branch_105:
var __t_and_107 bool = false
if __t105 {

var __t106 bool
{
if (__local_var_7_104.IntVal) > (999) {
__t106 = false
goto end_branch_106
} else {

}
}
{
__t106 = true
}
end_branch_106:
__t_and_107 = __t106
}
if __t_and_107 {
__t108 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_7_104.IntVal)}
goto end_branch_108
} else {

}
}
{
__t108 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_108:
// TAST (Let): __local_var_7_103 -> *Constructor_Data_Maybe_Just
var __local_var_7_103 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t108)})
var __t112 *Constructor_Data_Maybe_Just
{
if (__local_var_7_103 != nil) {
var __t111 *Constructor_Data_Maybe_Just
{
var __t110 *Constructor_Data_Maybe_Just
{
var __t109 bool
{
if ((__local_var_7_103).V0.IntVal) > (gopurs_runtime.Int(999).IntVal) {
__t109 = false
goto end_branch_109
} else {

}
}
{
__t109 = true
}
end_branch_109:
if __t109 {
__t110 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_110
} else {

}
}
{
__t110 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_110:
if (__t110 != nil) {
__t111 = &Constructor_Data_Maybe_Just{1, (__local_var_7_103).V0}
goto end_branch_111
} else {

}
}
{
__t111 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_111:
__t112 = __t111
goto end_branch_112
} else {

}
}
{
if (__local_var_7_103 == nil) {
__t112 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_112
} else {

}
}
{
__t112 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_112:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_6, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t112)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_2_1.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Time_Time(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_2, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_4.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_7.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_14)}).IntVal) - (1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_8_26_0 gopurs_runtime.Value
go__go_8_26_0 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_9_loop int64 = v_9_loop_val.IntVal
var v1_10_loop gopurs_runtime.Value = v1_10_loop_val
go__go_8_26_0:
for {
if false { continue go__go_8_26_0 }
var v_9 int64 = v_9_loop
_ = v_9
var v1_10 gopurs_runtime.Value = v1_10_loop
_ = v1_10
var __t30 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr != nil) {
var __t29 gopurs_runtime.Value
{
var __t_tag_27 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1
if (__t_tag_27 == nil) {
__t29 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_29
} else {

}
}
{
var __t28 bool
{
if (v_9) > (0) {
__t28 = false
goto end_branch_28
} else {

}
}
{
__t28 = true
}
end_branch_28:
if __t28 {
__t29 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_29
} else {

}
}
{
v_9_loop = (v_9) - (1)
v1_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1)}
continue go__go_8_26_0
__t29 = gopurs_runtime.Value{}
}
end_branch_29:
__t30 = __t29
goto end_branch_30
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr == nil) {
__t30 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_5_4.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return x_11
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_14)})
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
return __t30
}
}()
})
})
return gopurs_runtime.Apply(pure_4_3, gopurs_runtime.Apply2(go__go_8_26_0, gopurs_runtime.Int(n_7.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_4.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_14)})))}))
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_31, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_33.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_7.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_43)}).IntVal) - (1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_8_55_1 gopurs_runtime.Value
go__go_8_55_1 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_9_loop int64 = v_9_loop_val.IntVal
var v1_10_loop gopurs_runtime.Value = v1_10_loop_val
go__go_8_55_1:
for {
if false { continue go__go_8_55_1 }
var v_9 int64 = v_9_loop
_ = v_9
var v1_10 gopurs_runtime.Value = v1_10_loop
_ = v1_10
var __t59 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr != nil) {
var __t58 gopurs_runtime.Value
{
var __t_tag_56 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1
if (__t_tag_56 == nil) {
__t58 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_58
} else {

}
}
{
var __t57 bool
{
if (v_9) > (0) {
__t57 = false
goto end_branch_57
} else {

}
}
{
__t57 = true
}
end_branch_57:
if __t57 {
__t58 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_58
} else {

}
}
{
v_9_loop = (v_9) - (1)
v1_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1)}
continue go__go_8_55_1
__t58 = gopurs_runtime.Value{}
}
end_branch_58:
__t59 = __t58
goto end_branch_59
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr == nil) {
__t59 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_5_33.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return x_11
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_43)})
goto end_branch_59
} else {

}
}
{
__t59 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_59:
return __t59
}
}()
})
})
return gopurs_runtime.Apply(pure_4_32, gopurs_runtime.Apply2(go__go_8_55_1, gopurs_runtime.Int(n_7.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_33.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_43)})))}))
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_60, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_62.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_7.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_72)}).IntVal) - (1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_8_84_2 gopurs_runtime.Value
go__go_8_84_2 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_9_loop int64 = v_9_loop_val.IntVal
var v1_10_loop gopurs_runtime.Value = v1_10_loop_val
go__go_8_84_2:
for {
if false { continue go__go_8_84_2 }
var v_9 int64 = v_9_loop
_ = v_9
var v1_10 gopurs_runtime.Value = v1_10_loop
_ = v1_10
var __t88 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr != nil) {
var __t87 gopurs_runtime.Value
{
var __t_tag_85 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1
if (__t_tag_85 == nil) {
__t87 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_87
} else {

}
}
{
var __t86 bool
{
if (v_9) > (0) {
__t86 = false
goto end_branch_86
} else {

}
}
{
__t86 = true
}
end_branch_86:
if __t86 {
__t87 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_87
} else {

}
}
{
v_9_loop = (v_9) - (1)
v1_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1)}
continue go__go_8_84_2
__t87 = gopurs_runtime.Value{}
}
end_branch_87:
__t88 = __t87
goto end_branch_88
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr == nil) {
__t88 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_5_62.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return x_11
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_72)})
goto end_branch_88
} else {

}
}
{
__t88 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_88:
return __t88
}
}()
})
})
return gopurs_runtime.Apply(pure_4_61, gopurs_runtime.Apply2(go__go_8_84_2, gopurs_runtime.Int(n_7.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_62.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_72)})))}))
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_3_89, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_91.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_7.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_101)}).IntVal) - (1))), gopurs_runtime.Func(func(n_7 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_8_113_3 gopurs_runtime.Value
go__go_8_113_3 = gopurs_runtime.Func(func(v_9_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_9_loop int64 = v_9_loop_val.IntVal
var v1_10_loop gopurs_runtime.Value = v1_10_loop_val
go__go_8_113_3:
for {
if false { continue go__go_8_113_3 }
var v_9 int64 = v_9_loop
_ = v_9
var v1_10 gopurs_runtime.Value = v1_10_loop
_ = v1_10
var __t117 gopurs_runtime.Value
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr != nil) {
var __t116 gopurs_runtime.Value
{
var __t_tag_114 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1
if (__t_tag_114 == nil) {
__t116 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_116
} else {

}
}
{
var __t115 bool
{
if (v_9) > (0) {
__t115 = false
goto end_branch_115
} else {

}
}
{
__t115 = true
}
end_branch_115:
if __t115 {
__t116 = (*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V0
goto end_branch_116
} else {

}
}
{
v_9_loop = (v_9) - (1)
v1_10_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_10.UnsafePtr).V1)}
continue go__go_8_113_3
__t116 = gopurs_runtime.Value{}
}
end_branch_116:
__t117 = __t116
goto end_branch_117
} else {

}
}
{
if (v1_10.Type == 9 && v1_10.IntVal == 759514854 && v1_10.UnsafePtr == nil) {
__t117 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_5_91.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return x_11
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_101)})
goto end_branch_117
} else {

}
}
{
__t117 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_117:
return __t117
}
}()
})
})
return gopurs_runtime.Apply(pure_4_90, gopurs_runtime.Apply2(go__go_8_113_3, gopurs_runtime.Int(n_7.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_5_91.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_6_101)})))}))
})))
}



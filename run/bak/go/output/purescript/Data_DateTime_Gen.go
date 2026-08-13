package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_DateTime_Gen_genDateTime gopurs_runtime.Value
var once_Data_DateTime_Gen_genDateTime sync.Once
func Get_Data_DateTime_Gen_genDateTime() gopurs_runtime.Value {
	once_Data_DateTime_Gen_genDateTime.Do(func() {
		cache_Data_DateTime_Gen_genDateTime = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_DateTime_Gen_genDateTime(dictMonadGen_0_box)
})
	})
	return cache_Data_DateTime_Gen_genDateTime
}

func Call_Data_DateTime_Gen_genDateTime(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Bind1_1_0 -> gopurs_runtime.Value
Bind1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_1_0
// TAST (Let): Monad0_2_1 -> gopurs_runtime.Value
Monad0_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
// TAST (Let): Functor0_4_3 -> *Constructor_Data_Functor_Functor
Functor0_4_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_3
// TAST (Let): pure_5_4 -> gopurs_runtime.Value
pure_5_4 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_2_1, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_4
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Maybe_fromJust__1577979644()
}))
_ = __local_var_6_5
// TAST (Let): Bind1_2_15 -> gopurs_runtime.Value
Bind1_2_15 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{})
_ = Bind1_2_15
// TAST (Let): Apply0_3_16 -> *Constructor_Control_Apply_Apply
Apply0_3_16 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_15, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_3_16
// TAST (Let): Monad0_4_17 -> gopurs_runtime.Value
Monad0_4_17 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_17
// TAST (Let): pure_5_18 -> gopurs_runtime.Value
pure_5_18 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_17, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_18
// TAST (Let): foldableNonEmpty1_6_20 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_6_20 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_21 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_21 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_21
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_10_22 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_10_22 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_22
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_21.V0), gopurs_runtime.Apply(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_22.V0), gopurs_runtime.Apply(f_8, x_11), acc_12)
})
}), gopurs_runtime.RecordGet(dictMonoid_6, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray16 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray16
res_go_foldlArray16 := gopurs_runtime.Apply2(f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = res_go_foldlArray16
arr_go_foldlArray16 := (*[]gopurs_runtime.Value)(arr_val_foldlArray16.UnsafePtr)
_ = arr_go_foldlArray16
for _, v_foldlArray16 := range *arr_go_foldlArray16 {
res_go_foldlArray16 = gopurs_runtime.Apply2(f_6, res_go_foldlArray16, v_foldlArray16)
}
return res_go_foldlArray16
}()
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_6_20
// TAST (Let): __local_var_6_19 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_6_19 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_6_20)}
}), gopurs_runtime.Func(func(dictSemigroup_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray16 := (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1
_ = arr_val_foldlArray16
res_go_foldlArray16 := gopurs_runtime.Apply(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0)
_ = res_go_foldlArray16
arr_go_foldlArray16 := (*[]gopurs_runtime.Value)(arr_val_foldlArray16.UnsafePtr)
_ = arr_go_foldlArray16
for _, v_foldlArray16 := range *arr_go_foldlArray16 {
res_go_foldlArray16 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_7, "append"), s_10, gopurs_runtime.Apply(f_8, a1_11))
})
}), res_go_foldlArray16, v_foldlArray16)
}
return res_go_foldlArray16
}()
})
})
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray15 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray15
res_go_foldlArray15 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0
_ = res_go_foldlArray15
arr_go_foldlArray15 := (*[]gopurs_runtime.Value)(arr_val_foldlArray15.UnsafePtr)
_ = arr_go_foldlArray15
for _, v_foldlArray15 := range *arr_go_foldlArray15 {
res_go_foldlArray15 = gopurs_runtime.Apply2(f_7, res_go_foldlArray15, v_foldlArray15)
}
return res_go_foldlArray15
}()
})
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_23 -> gopurs_runtime.Value
__local_var_9_23 := gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = __local_var_9_23
// TAST (Let): __local_var_10_24 -> *Constructor_Data_Maybe_Just
__local_var_10_24 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_26 -> gopurs_runtime.Value
__local_var_11_26 := gopurs_runtime.Apply(f_7, a1_10)
_ = __local_var_11_26
// TAST (Let): __local_var_11_25 -> gopurs_runtime.Value
__local_var_11_25 := gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t27 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 930809136 && v2_12.UnsafePtr == nil) {
__t27 = a1_10
goto end_branch_27
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 930809136 && v2_12.UnsafePtr != nil) {
__t27 = gopurs_runtime.Apply(__local_var_11_26, (*Constructor_Data_Maybe_Just)(v2_12.UnsafePtr).V0)
goto end_branch_27
} else {

}
}
{
__t27 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_27:
return __t27
})
_ = __local_var_11_25
return gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_11_25, x_12)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
_ = __local_var_10_24
var __t28 gopurs_runtime.Value
{
if (__local_var_10_24 == nil) {
__t28 = (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0
goto end_branch_28
} else {

}
}
{
if (__local_var_10_24 != nil) {
__t28 = gopurs_runtime.Apply(__local_var_9_23, (__local_var_10_24).V0)
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return __t28
})
})}
_ = __local_var_6_19
// TAST (Let): __local_var_7_29 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_7_29 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t30 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 930809136 && v_8.UnsafePtr != nil) {
__t30 = (*Constructor_Data_Maybe_Just)(v_8.UnsafePtr).V0
goto end_branch_30
} else {

}
}
{
__t30 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_30:
return __t30
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_32 -> gopurs_runtime.Value
__local_var_8_32 := gopurs_runtime.Int((a_7.IntVal) + (1))
_ = __local_var_8_32
var __t36 *Constructor_Data_Maybe_Just
{
var __t33 bool
{
if (__local_var_8_32.IntVal) < (0) {
__t33 = false
goto end_branch_33
} else {

}
}
{
__t33 = true
}
end_branch_33:
var __t_and_35 bool = false
if __t33 {

var __t34 bool
{
if (__local_var_8_32.IntVal) > (23) {
__t34 = false
goto end_branch_34
} else {

}
}
{
__t34 = true
}
end_branch_34:
__t_and_35 = __t34
}
if __t_and_35 {
__t36 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_8_32.IntVal)}
goto end_branch_36
} else {

}
}
{
__t36 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_36:
// TAST (Let): __local_var_8_31 -> *Constructor_Data_Maybe_Just
var __local_var_8_31 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t36)})
var __t40 *Constructor_Data_Maybe_Just
{
if (__local_var_8_31 != nil) {
var __t39 *Constructor_Data_Maybe_Just
{
var __t38 *Constructor_Data_Maybe_Just
{
var __t37 bool
{
if ((__local_var_8_31).V0.IntVal) > (gopurs_runtime.Int(23).IntVal) {
__t37 = false
goto end_branch_37
} else {

}
}
{
__t37 = true
}
end_branch_37:
if __t37 {
__t38 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_38
} else {

}
}
{
__t38 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_38:
if (__t38 != nil) {
__t39 = &Constructor_Data_Maybe_Just{1, (__local_var_8_31).V0}
goto end_branch_39
} else {

}
}
{
__t39 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_39:
__t40 = __t39
goto end_branch_40
} else {

}
}
{
if (__local_var_8_31 == nil) {
__t40 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_40
} else {

}
}
{
__t40 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_40:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t40)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
// TAST (Let): Monad0_4_46 -> gopurs_runtime.Value
Monad0_4_46 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_46
// TAST (Let): pure_5_47 -> gopurs_runtime.Value
pure_5_47 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_46, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_47
// TAST (Let): foldableNonEmpty1_6_49 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_6_49 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_50 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_50 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_50
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_10_51 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_10_51 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_51
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_50.V0), gopurs_runtime.Apply(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_51.V0), gopurs_runtime.Apply(f_8, x_11), acc_12)
})
}), gopurs_runtime.RecordGet(dictMonoid_6, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray15 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray15
res_go_foldlArray15 := gopurs_runtime.Apply2(f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = res_go_foldlArray15
arr_go_foldlArray15 := (*[]gopurs_runtime.Value)(arr_val_foldlArray15.UnsafePtr)
_ = arr_go_foldlArray15
for _, v_foldlArray15 := range *arr_go_foldlArray15 {
res_go_foldlArray15 = gopurs_runtime.Apply2(f_6, res_go_foldlArray15, v_foldlArray15)
}
return res_go_foldlArray15
}()
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_6_49
// TAST (Let): __local_var_6_48 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_6_48 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_6_49)}
}), gopurs_runtime.Func(func(dictSemigroup_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray15 := (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1
_ = arr_val_foldlArray15
res_go_foldlArray15 := gopurs_runtime.Apply(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0)
_ = res_go_foldlArray15
arr_go_foldlArray15 := (*[]gopurs_runtime.Value)(arr_val_foldlArray15.UnsafePtr)
_ = arr_go_foldlArray15
for _, v_foldlArray15 := range *arr_go_foldlArray15 {
res_go_foldlArray15 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_7, "append"), s_10, gopurs_runtime.Apply(f_8, a1_11))
})
}), res_go_foldlArray15, v_foldlArray15)
}
return res_go_foldlArray15
}()
})
})
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray14 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray14
res_go_foldlArray14 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0
_ = res_go_foldlArray14
arr_go_foldlArray14 := (*[]gopurs_runtime.Value)(arr_val_foldlArray14.UnsafePtr)
_ = arr_go_foldlArray14
for _, v_foldlArray14 := range *arr_go_foldlArray14 {
res_go_foldlArray14 = gopurs_runtime.Apply2(f_7, res_go_foldlArray14, v_foldlArray14)
}
return res_go_foldlArray14
}()
})
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_52 -> gopurs_runtime.Value
__local_var_9_52 := gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = __local_var_9_52
// TAST (Let): __local_var_10_53 -> *Constructor_Data_Maybe_Just
__local_var_10_53 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_55 -> gopurs_runtime.Value
__local_var_11_55 := gopurs_runtime.Apply(f_7, a1_10)
_ = __local_var_11_55
// TAST (Let): __local_var_11_54 -> gopurs_runtime.Value
__local_var_11_54 := gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t56 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 930809136 && v2_12.UnsafePtr == nil) {
__t56 = a1_10
goto end_branch_56
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 930809136 && v2_12.UnsafePtr != nil) {
__t56 = gopurs_runtime.Apply(__local_var_11_55, (*Constructor_Data_Maybe_Just)(v2_12.UnsafePtr).V0)
goto end_branch_56
} else {

}
}
{
__t56 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_56:
return __t56
})
_ = __local_var_11_54
return gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_11_54, x_12)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
_ = __local_var_10_53
var __t57 gopurs_runtime.Value
{
if (__local_var_10_53 == nil) {
__t57 = (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0
goto end_branch_57
} else {

}
}
{
if (__local_var_10_53 != nil) {
__t57 = gopurs_runtime.Apply(__local_var_9_52, (__local_var_10_53).V0)
goto end_branch_57
} else {

}
}
{
__t57 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_57:
return __t57
})
})}
_ = __local_var_6_48
// TAST (Let): __local_var_7_58 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_7_58 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t59 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 930809136 && v_8.UnsafePtr != nil) {
__t59 = (*Constructor_Data_Maybe_Just)(v_8.UnsafePtr).V0
goto end_branch_59
} else {

}
}
{
__t59 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_59:
return __t59
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_61 -> gopurs_runtime.Value
__local_var_8_61 := gopurs_runtime.Int((a_7.IntVal) + (1))
_ = __local_var_8_61
var __t65 *Constructor_Data_Maybe_Just
{
var __t62 bool
{
if (__local_var_8_61.IntVal) < (0) {
__t62 = false
goto end_branch_62
} else {

}
}
{
__t62 = true
}
end_branch_62:
var __t_and_64 bool = false
if __t62 {

var __t63 bool
{
if (__local_var_8_61.IntVal) > (59) {
__t63 = false
goto end_branch_63
} else {

}
}
{
__t63 = true
}
end_branch_63:
__t_and_64 = __t63
}
if __t_and_64 {
__t65 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_8_61.IntVal)}
goto end_branch_65
} else {

}
}
{
__t65 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_65:
// TAST (Let): __local_var_8_60 -> *Constructor_Data_Maybe_Just
var __local_var_8_60 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t65)})
var __t69 *Constructor_Data_Maybe_Just
{
if (__local_var_8_60 != nil) {
var __t68 *Constructor_Data_Maybe_Just
{
var __t67 *Constructor_Data_Maybe_Just
{
var __t66 bool
{
if ((__local_var_8_60).V0.IntVal) > (gopurs_runtime.Int(59).IntVal) {
__t66 = false
goto end_branch_66
} else {

}
}
{
__t66 = true
}
end_branch_66:
if __t66 {
__t67 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_67
} else {

}
}
{
__t67 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_67:
if (__t67 != nil) {
__t68 = &Constructor_Data_Maybe_Just{1, (__local_var_8_60).V0}
goto end_branch_68
} else {

}
}
{
__t68 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_68:
__t69 = __t68
goto end_branch_69
} else {

}
}
{
if (__local_var_8_60 == nil) {
__t69 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_69
} else {

}
}
{
__t69 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_69:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t69)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
// TAST (Let): Monad0_4_75 -> gopurs_runtime.Value
Monad0_4_75 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_75
// TAST (Let): pure_5_76 -> gopurs_runtime.Value
pure_5_76 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_75, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_76
// TAST (Let): foldableNonEmpty1_6_78 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_6_78 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_79 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_79 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_79
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_10_80 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_10_80 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_80
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_79.V0), gopurs_runtime.Apply(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_80.V0), gopurs_runtime.Apply(f_8, x_11), acc_12)
})
}), gopurs_runtime.RecordGet(dictMonoid_6, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray14 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray14
res_go_foldlArray14 := gopurs_runtime.Apply2(f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = res_go_foldlArray14
arr_go_foldlArray14 := (*[]gopurs_runtime.Value)(arr_val_foldlArray14.UnsafePtr)
_ = arr_go_foldlArray14
for _, v_foldlArray14 := range *arr_go_foldlArray14 {
res_go_foldlArray14 = gopurs_runtime.Apply2(f_6, res_go_foldlArray14, v_foldlArray14)
}
return res_go_foldlArray14
}()
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_6_78
// TAST (Let): __local_var_6_77 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_6_77 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_6_78)}
}), gopurs_runtime.Func(func(dictSemigroup_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray14 := (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1
_ = arr_val_foldlArray14
res_go_foldlArray14 := gopurs_runtime.Apply(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0)
_ = res_go_foldlArray14
arr_go_foldlArray14 := (*[]gopurs_runtime.Value)(arr_val_foldlArray14.UnsafePtr)
_ = arr_go_foldlArray14
for _, v_foldlArray14 := range *arr_go_foldlArray14 {
res_go_foldlArray14 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_7, "append"), s_10, gopurs_runtime.Apply(f_8, a1_11))
})
}), res_go_foldlArray14, v_foldlArray14)
}
return res_go_foldlArray14
}()
})
})
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray13
res_go_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0
_ = res_go_foldlArray13
arr_go_foldlArray13 := (*[]gopurs_runtime.Value)(arr_val_foldlArray13.UnsafePtr)
_ = arr_go_foldlArray13
for _, v_foldlArray13 := range *arr_go_foldlArray13 {
res_go_foldlArray13 = gopurs_runtime.Apply2(f_7, res_go_foldlArray13, v_foldlArray13)
}
return res_go_foldlArray13
}()
})
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_81 -> gopurs_runtime.Value
__local_var_9_81 := gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = __local_var_9_81
// TAST (Let): __local_var_10_82 -> *Constructor_Data_Maybe_Just
__local_var_10_82 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_84 -> gopurs_runtime.Value
__local_var_11_84 := gopurs_runtime.Apply(f_7, a1_10)
_ = __local_var_11_84
// TAST (Let): __local_var_11_83 -> gopurs_runtime.Value
__local_var_11_83 := gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t85 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 930809136 && v2_12.UnsafePtr == nil) {
__t85 = a1_10
goto end_branch_85
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 930809136 && v2_12.UnsafePtr != nil) {
__t85 = gopurs_runtime.Apply(__local_var_11_84, (*Constructor_Data_Maybe_Just)(v2_12.UnsafePtr).V0)
goto end_branch_85
} else {

}
}
{
__t85 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_85:
return __t85
})
_ = __local_var_11_83
return gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_11_83, x_12)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
_ = __local_var_10_82
var __t86 gopurs_runtime.Value
{
if (__local_var_10_82 == nil) {
__t86 = (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0
goto end_branch_86
} else {

}
}
{
if (__local_var_10_82 != nil) {
__t86 = gopurs_runtime.Apply(__local_var_9_81, (__local_var_10_82).V0)
goto end_branch_86
} else {

}
}
{
__t86 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_86:
return __t86
})
})}
_ = __local_var_6_77
// TAST (Let): __local_var_7_87 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_7_87 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t88 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 930809136 && v_8.UnsafePtr != nil) {
__t88 = (*Constructor_Data_Maybe_Just)(v_8.UnsafePtr).V0
goto end_branch_88
} else {

}
}
{
__t88 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_88:
return __t88
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_90 -> gopurs_runtime.Value
__local_var_8_90 := gopurs_runtime.Int((a_7.IntVal) + (1))
_ = __local_var_8_90
var __t94 *Constructor_Data_Maybe_Just
{
var __t91 bool
{
if (__local_var_8_90.IntVal) < (0) {
__t91 = false
goto end_branch_91
} else {

}
}
{
__t91 = true
}
end_branch_91:
var __t_and_93 bool = false
if __t91 {

var __t92 bool
{
if (__local_var_8_90.IntVal) > (59) {
__t92 = false
goto end_branch_92
} else {

}
}
{
__t92 = true
}
end_branch_92:
__t_and_93 = __t92
}
if __t_and_93 {
__t94 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_8_90.IntVal)}
goto end_branch_94
} else {

}
}
{
__t94 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_94:
// TAST (Let): __local_var_8_89 -> *Constructor_Data_Maybe_Just
var __local_var_8_89 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t94)})
var __t98 *Constructor_Data_Maybe_Just
{
if (__local_var_8_89 != nil) {
var __t97 *Constructor_Data_Maybe_Just
{
var __t96 *Constructor_Data_Maybe_Just
{
var __t95 bool
{
if ((__local_var_8_89).V0.IntVal) > (gopurs_runtime.Int(59).IntVal) {
__t95 = false
goto end_branch_95
} else {

}
}
{
__t95 = true
}
end_branch_95:
if __t95 {
__t96 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_96
} else {

}
}
{
__t96 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_96:
if (__t96 != nil) {
__t97 = &Constructor_Data_Maybe_Just{1, (__local_var_8_89).V0}
goto end_branch_97
} else {

}
}
{
__t97 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_97:
__t98 = __t97
goto end_branch_98
} else {

}
}
{
if (__local_var_8_89 == nil) {
__t98 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_98
} else {

}
}
{
__t98 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_98:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t98)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
// TAST (Let): Monad0_4_104 -> gopurs_runtime.Value
Monad0_4_104 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_4_104
// TAST (Let): pure_5_105 -> gopurs_runtime.Value
pure_5_105 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_104, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_5_105
// TAST (Let): foldableNonEmpty1_6_107 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_6_107 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_108 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_108 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_108
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_10_109 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_10_109 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_6, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_10_109
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_108.V0), gopurs_runtime.Apply(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_10_109.V0), gopurs_runtime.Apply(f_8, x_11), acc_12)
})
}), gopurs_runtime.RecordGet(dictMonoid_6, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray13
res_go_foldlArray13 := gopurs_runtime.Apply2(f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = res_go_foldlArray13
arr_go_foldlArray13 := (*[]gopurs_runtime.Value)(arr_val_foldlArray13.UnsafePtr)
_ = arr_go_foldlArray13
for _, v_foldlArray13 := range *arr_go_foldlArray13 {
res_go_foldlArray13 = gopurs_runtime.Apply2(f_6, res_go_foldlArray13, v_foldlArray13)
}
return res_go_foldlArray13
}()
})
})
}), gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_6, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_6, b_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_6_107
// TAST (Let): __local_var_6_106 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_6_106 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_6_107)}
}), gopurs_runtime.Func(func(dictSemigroup_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray13 := (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V1
_ = arr_val_foldlArray13
res_go_foldlArray13 := gopurs_runtime.Apply(f_8, (*Constructor_Data_NonEmpty_NonEmpty)(v_9.UnsafePtr).V0)
_ = res_go_foldlArray13
arr_go_foldlArray13 := (*[]gopurs_runtime.Value)(arr_val_foldlArray13.UnsafePtr)
_ = arr_go_foldlArray13
for _, v_foldlArray13 := range *arr_go_foldlArray13 {
res_go_foldlArray13 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_7, "append"), s_10, gopurs_runtime.Apply(f_8, a1_11))
})
}), res_go_foldlArray13, v_foldlArray13)
}
return res_go_foldlArray13
}()
})
})
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray12 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1
_ = arr_val_foldlArray12
res_go_foldlArray12 := (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0
_ = res_go_foldlArray12
arr_go_foldlArray12 := (*[]gopurs_runtime.Value)(arr_val_foldlArray12.UnsafePtr)
_ = arr_go_foldlArray12
for _, v_foldlArray12 := range *arr_go_foldlArray12 {
res_go_foldlArray12 = gopurs_runtime.Apply2(f_7, res_go_foldlArray12, v_foldlArray12)
}
return res_go_foldlArray12
}()
})
}), gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_110 -> gopurs_runtime.Value
__local_var_9_110 := gopurs_runtime.Apply(f_7, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0)
_ = __local_var_9_110
// TAST (Let): __local_var_10_111 -> *Constructor_Data_Maybe_Just
__local_var_10_111 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_10 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_11_113 -> gopurs_runtime.Value
__local_var_11_113 := gopurs_runtime.Apply(f_7, a1_10)
_ = __local_var_11_113
// TAST (Let): __local_var_11_112 -> gopurs_runtime.Value
__local_var_11_112 := gopurs_runtime.Func(func(v2_12 gopurs_runtime.Value) gopurs_runtime.Value {
var __t114 gopurs_runtime.Value
{
if (v2_12.Type == 9 && v2_12.IntVal == 930809136 && v2_12.UnsafePtr == nil) {
__t114 = a1_10
goto end_branch_114
} else {

}
}
{
if (v2_12.Type == 9 && v2_12.IntVal == 930809136 && v2_12.UnsafePtr != nil) {
__t114 = gopurs_runtime.Apply(__local_var_11_113, (*Constructor_Data_Maybe_Just)(v2_12.UnsafePtr).V0)
goto end_branch_114
} else {

}
}
{
__t114 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_114:
return __t114
})
_ = __local_var_11_112
return gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_11_112, x_12)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V1))
_ = __local_var_10_111
var __t115 gopurs_runtime.Value
{
if (__local_var_10_111 == nil) {
__t115 = (*Constructor_Data_NonEmpty_NonEmpty)(v_8.UnsafePtr).V0
goto end_branch_115
} else {

}
}
{
if (__local_var_10_111 != nil) {
__t115 = gopurs_runtime.Apply(__local_var_9_110, (__local_var_10_111).V0)
goto end_branch_115
} else {

}
}
{
__t115 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_115:
return __t115
})
})}
_ = __local_var_6_106
// TAST (Let): __local_var_7_116 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_7_116 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
var __t117 gopurs_runtime.Value
{
if (v_8.Type == 9 && v_8.IntVal == 930809136 && v_8.UnsafePtr != nil) {
__t117 = (*Constructor_Data_Maybe_Just)(v_8.UnsafePtr).V0
goto end_branch_117
} else {

}
}
{
__t117 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_117:
return __t117
})
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_119 -> gopurs_runtime.Value
__local_var_8_119 := gopurs_runtime.Int((a_7.IntVal) + (1))
_ = __local_var_8_119
var __t123 *Constructor_Data_Maybe_Just
{
var __t120 bool
{
if (__local_var_8_119.IntVal) < (0) {
__t120 = false
goto end_branch_120
} else {

}
}
{
__t120 = true
}
end_branch_120:
var __t_and_122 bool = false
if __t120 {

var __t121 bool
{
if (__local_var_8_119.IntVal) > (999) {
__t121 = false
goto end_branch_121
} else {

}
}
{
__t121 = true
}
end_branch_121:
__t_and_122 = __t121
}
if __t_and_122 {
__t123 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_8_119.IntVal)}
goto end_branch_123
} else {

}
}
{
__t123 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_123:
// TAST (Let): __local_var_8_118 -> *Constructor_Data_Maybe_Just
var __local_var_8_118 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t123)})
var __t127 *Constructor_Data_Maybe_Just
{
if (__local_var_8_118 != nil) {
var __t126 *Constructor_Data_Maybe_Just
{
var __t125 *Constructor_Data_Maybe_Just
{
var __t124 bool
{
if ((__local_var_8_118).V0.IntVal) > (gopurs_runtime.Int(999).IntVal) {
__t124 = false
goto end_branch_124
} else {

}
}
{
__t124 = true
}
end_branch_124:
if __t124 {
__t125 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_125
} else {

}
}
{
__t125 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_125:
if (__t125 != nil) {
__t126 = &Constructor_Data_Maybe_Just{1, (__local_var_8_118).V0}
goto end_branch_126
} else {

}
}
{
__t126 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_126:
__t127 = __t126
goto end_branch_127
} else {

}
}
{
if (__local_var_8_118 == nil) {
__t127 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_127
} else {

}
}
{
__t127 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_127:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t127)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_1_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_DateTime_DateTime(), gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{}), "Bind1"), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t9 *Constructor_Data_Maybe_Just
{
var __t6 bool
{
if (x_7.IntVal) < (-271820) {
__t6 = false
goto end_branch_6
} else {

}
}
{
__t6 = true
}
end_branch_6:
var __t_and_8 bool = false
if __t6 {

var __t7 bool
{
if (x_7.IntVal) > (275759) {
__t7 = false
goto end_branch_7
} else {

}
}
{
__t7 = true
}
end_branch_7:
__t_and_8 = __t7
}
if __t_and_8 {
__t9 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(x_7.IntVal)}
goto end_branch_9
} else {

}
}
{
__t9 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_9:
return gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t9)})
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(1900), gopurs_runtime.Int(2100))), gopurs_runtime.Func(func(year_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 int64
{
if (gopurs_runtime.Apply(Get_Data_Date_isLeapYear(), gopurs_runtime.Int(year_6.IntVal)).IntVal) != (0) {
__t10 = 365
goto end_branch_10
} else {

}
}
{
__t10 = 364
}
end_branch_10:
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_3.V0), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Int_toNumber(), x_7)
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int(__t10))), gopurs_runtime.Func(func(days_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pure_5_4, gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_12 -> *Constructor_Data_Maybe_Just
__local_var_9_12 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Date_exactDate(), gopurs_runtime.Int(year_6.IntVal), gopurs_runtime.Value{Type: 9, IntVal: int64(1908470532), UnsafePtr: nil}, gopurs_runtime.Int(1)))
_ = __local_var_9_12
var __t13 *Constructor_Data_Maybe_Just
{
if (__local_var_9_12 != nil) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_Date_adjust(), gopurs_runtime.Float(days_7.FloatVal()), gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_9_12).V0))}))
goto end_branch_13
} else {

}
}
{
if (__local_var_9_12 == nil) {
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))})
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_13:
// TAST (Let): __local_var_9_11 -> *Constructor_Data_Maybe_Just
__local_var_9_11 := __t13
_ = __local_var_9_11
var __t14 *Constructor_Data_Date_Date
{
if (__local_var_9_11 != nil) {
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date]((__local_var_9_11).V0)
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.CoerceToStruct[Constructor_Data_Date_Date](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_14:
return gopurs_runtime.Value{Type: 9, IntVal: 745776346, UnsafePtr: unsafe.Pointer(__t14)}
}))))})
}))
}))), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_16.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_16.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_3_16.V1), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Bind1_2_15, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), Get_Data_Time_Time(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_17, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_19.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_8.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_29)}).IntVal) - (1))), gopurs_runtime.Func(func(n_8 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_9_41_0 gopurs_runtime.Value
go__go_9_41_0 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_10_loop int64 = v_10_loop_val.IntVal
var v1_11_loop gopurs_runtime.Value = v1_11_loop_val
go__go_9_41_0:
for {
if false { continue go__go_9_41_0 }
var v_10 int64 = v_10_loop
_ = v_10
var v1_11 gopurs_runtime.Value = v1_11_loop
_ = v1_11
var __t45 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 759514854 && v1_11.UnsafePtr != nil) {
var __t44 gopurs_runtime.Value
{
var __t_tag_42 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V1
if (__t_tag_42 == nil) {
__t44 = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V0
goto end_branch_44
} else {

}
}
{
var __t43 bool
{
if (v_10) > (0) {
__t43 = false
goto end_branch_43
} else {

}
}
{
__t43 = true
}
end_branch_43:
if __t43 {
__t44 = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V0
goto end_branch_44
} else {

}
}
{
v_10_loop = (v_10) - (1)
v1_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V1)}
continue go__go_9_41_0
__t44 = gopurs_runtime.Value{}
}
end_branch_44:
__t45 = __t44
goto end_branch_45
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 759514854 && v1_11.UnsafePtr == nil) {
__t45 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_6_19.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return x_12
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_29)})
goto end_branch_45
} else {

}
}
{
__t45 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_45:
return __t45
}
}()
})
})
return gopurs_runtime.Apply(pure_5_18, gopurs_runtime.Apply2(go__go_9_41_0, gopurs_runtime.Int(n_8.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_19.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_29)})))}))
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_46, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_48.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_8.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_58)}).IntVal) - (1))), gopurs_runtime.Func(func(n_8 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_9_70_1 gopurs_runtime.Value
go__go_9_70_1 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_10_loop int64 = v_10_loop_val.IntVal
var v1_11_loop gopurs_runtime.Value = v1_11_loop_val
go__go_9_70_1:
for {
if false { continue go__go_9_70_1 }
var v_10 int64 = v_10_loop
_ = v_10
var v1_11 gopurs_runtime.Value = v1_11_loop
_ = v1_11
var __t74 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 759514854 && v1_11.UnsafePtr != nil) {
var __t73 gopurs_runtime.Value
{
var __t_tag_71 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V1
if (__t_tag_71 == nil) {
__t73 = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V0
goto end_branch_73
} else {

}
}
{
var __t72 bool
{
if (v_10) > (0) {
__t72 = false
goto end_branch_72
} else {

}
}
{
__t72 = true
}
end_branch_72:
if __t72 {
__t73 = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V0
goto end_branch_73
} else {

}
}
{
v_10_loop = (v_10) - (1)
v1_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V1)}
continue go__go_9_70_1
__t73 = gopurs_runtime.Value{}
}
end_branch_73:
__t74 = __t73
goto end_branch_74
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 759514854 && v1_11.UnsafePtr == nil) {
__t74 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_6_48.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return x_12
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_58)})
goto end_branch_74
} else {

}
}
{
__t74 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_74:
return __t74
}
}()
})
})
return gopurs_runtime.Apply(pure_5_47, gopurs_runtime.Apply2(go__go_9_70_1, gopurs_runtime.Int(n_8.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_48.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_58)})))}))
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_75, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_77.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_8.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_87)}).IntVal) - (1))), gopurs_runtime.Func(func(n_8 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_9_99_2 gopurs_runtime.Value
go__go_9_99_2 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_10_loop int64 = v_10_loop_val.IntVal
var v1_11_loop gopurs_runtime.Value = v1_11_loop_val
go__go_9_99_2:
for {
if false { continue go__go_9_99_2 }
var v_10 int64 = v_10_loop
_ = v_10
var v1_11 gopurs_runtime.Value = v1_11_loop
_ = v1_11
var __t103 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 759514854 && v1_11.UnsafePtr != nil) {
var __t102 gopurs_runtime.Value
{
var __t_tag_100 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V1
if (__t_tag_100 == nil) {
__t102 = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V0
goto end_branch_102
} else {

}
}
{
var __t101 bool
{
if (v_10) > (0) {
__t101 = false
goto end_branch_101
} else {

}
}
{
__t101 = true
}
end_branch_101:
if __t101 {
__t102 = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V0
goto end_branch_102
} else {

}
}
{
v_10_loop = (v_10) - (1)
v1_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V1)}
continue go__go_9_99_2
__t102 = gopurs_runtime.Value{}
}
end_branch_102:
__t103 = __t102
goto end_branch_103
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 759514854 && v1_11.UnsafePtr == nil) {
__t103 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_6_77.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return x_12
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_87)})
goto end_branch_103
} else {

}
}
{
__t103 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_103:
return __t103
}
}()
})
})
return gopurs_runtime.Apply(pure_5_76, gopurs_runtime.Apply2(go__go_9_99_2, gopurs_runtime.Int(n_8.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_77.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_87)})))}))
}))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_4_104, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_106.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_8.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_116)}).IntVal) - (1))), gopurs_runtime.Func(func(n_8 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_9_128_3 gopurs_runtime.Value
go__go_9_128_3 = gopurs_runtime.Func(func(v_10_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_11_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_10_loop int64 = v_10_loop_val.IntVal
var v1_11_loop gopurs_runtime.Value = v1_11_loop_val
go__go_9_128_3:
for {
if false { continue go__go_9_128_3 }
var v_10 int64 = v_10_loop
_ = v_10
var v1_11 gopurs_runtime.Value = v1_11_loop
_ = v1_11
var __t132 gopurs_runtime.Value
{
if (v1_11.Type == 9 && v1_11.IntVal == 759514854 && v1_11.UnsafePtr != nil) {
var __t131 gopurs_runtime.Value
{
var __t_tag_129 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V1
if (__t_tag_129 == nil) {
__t131 = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V0
goto end_branch_131
} else {

}
}
{
var __t130 bool
{
if (v_10) > (0) {
__t130 = false
goto end_branch_130
} else {

}
}
{
__t130 = true
}
end_branch_130:
if __t130 {
__t131 = (*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V0
goto end_branch_131
} else {

}
}
{
v_10_loop = (v_10) - (1)
v1_11_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_11.UnsafePtr).V1)}
continue go__go_9_128_3
__t131 = gopurs_runtime.Value{}
}
end_branch_131:
__t132 = __t131
goto end_branch_132
} else {

}
}
{
if (v1_11.Type == 9 && v1_11.IntVal == 759514854 && v1_11.UnsafePtr == nil) {
__t132 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_6_106.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_12 gopurs_runtime.Value) gopurs_runtime.Value {
return x_12
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_116)})
goto end_branch_132
} else {

}
}
{
__t132 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_132:
return __t132
}
}()
})
})
return gopurs_runtime.Apply(pure_5_105, gopurs_runtime.Apply2(go__go_9_128_3, gopurs_runtime.Int(n_8.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_6_106.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_7_116)})))}))
}))))
}



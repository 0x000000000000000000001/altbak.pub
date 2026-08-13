package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Time_Component_Gen_genSecond gopurs_runtime.Value
var once_Data_Time_Component_Gen_genSecond sync.Once
func Get_Data_Time_Component_Gen_genSecond() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genSecond.Do(func() {
		cache_Data_Time_Component_Gen_genSecond = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Gen_genSecond(dictMonadGen_0_box)
})
	})
	return cache_Data_Time_Component_Gen_genSecond
}

var cache_Data_Time_Component_Gen_genMinute gopurs_runtime.Value
var once_Data_Time_Component_Gen_genMinute sync.Once
func Get_Data_Time_Component_Gen_genMinute() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genMinute.Do(func() {
		cache_Data_Time_Component_Gen_genMinute = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Gen_genMinute(dictMonadGen_0_box)
})
	})
	return cache_Data_Time_Component_Gen_genMinute
}

var cache_Data_Time_Component_Gen_genMillisecond gopurs_runtime.Value
var once_Data_Time_Component_Gen_genMillisecond sync.Once
func Get_Data_Time_Component_Gen_genMillisecond() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genMillisecond.Do(func() {
		cache_Data_Time_Component_Gen_genMillisecond = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Gen_genMillisecond(dictMonadGen_0_box)
})
	})
	return cache_Data_Time_Component_Gen_genMillisecond
}

var cache_Data_Time_Component_Gen_genHour gopurs_runtime.Value
var once_Data_Time_Component_Gen_genHour sync.Once
func Get_Data_Time_Component_Gen_genHour() gopurs_runtime.Value {
	once_Data_Time_Component_Gen_genHour.Do(func() {
		cache_Data_Time_Component_Gen_genHour = gopurs_runtime.Func(func(dictMonadGen_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Time_Component_Gen_genHour(dictMonadGen_0_box)
})
	})
	return cache_Data_Time_Component_Gen_genHour
}

func Call_Data_Time_Component_Gen_genSecond(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): foldableNonEmpty1_3_3 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_3_3 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_4 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_5 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_5
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_4.V0), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_5.V0), gopurs_runtime.Apply(f_5, x_8), acc_9)
})
}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray8 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1
_ = arr_val_foldlArray8
res_go_foldlArray8 := gopurs_runtime.Apply2(f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0)
_ = res_go_foldlArray8
arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
_ = arr_go_foldlArray8
for _, v_foldlArray8 := range *arr_go_foldlArray8 {
res_go_foldlArray8 = gopurs_runtime.Apply2(f_3, res_go_foldlArray8, v_foldlArray8)
}
return res_go_foldlArray8
}()
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_3_3
// TAST (Let): __local_var_3_2 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_3_2 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_3_3)}
}), gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray8 := (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1
_ = arr_val_foldlArray8
res_go_foldlArray8 := gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0)
_ = res_go_foldlArray8
arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
_ = arr_go_foldlArray8
for _, v_foldlArray8 := range *arr_go_foldlArray8 {
res_go_foldlArray8 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_4, "append"), s_7, gopurs_runtime.Apply(f_5, a1_8))
})
}), res_go_foldlArray8, v_foldlArray8)
}
return res_go_foldlArray8
}()
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray7 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1
_ = arr_val_foldlArray7
res_go_foldlArray7 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0
_ = res_go_foldlArray7
arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
_ = arr_go_foldlArray7
for _, v_foldlArray7 := range *arr_go_foldlArray7 {
res_go_foldlArray7 = gopurs_runtime.Apply2(f_4, res_go_foldlArray7, v_foldlArray7)
}
return res_go_foldlArray7
}()
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0)
_ = __local_var_6_6
// TAST (Let): __local_var_7_7 -> *Constructor_Data_Maybe_Just
__local_var_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(f_4, a1_7)
_ = __local_var_8_9
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 930809136 && v2_9.UnsafePtr == nil) {
__t10 = a1_7
goto end_branch_10
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 930809136 && v2_9.UnsafePtr != nil) {
__t10 = gopurs_runtime.Apply(__local_var_8_9, (*Constructor_Data_Maybe_Just)(v2_9.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
_ = __local_var_8_8
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_8_8, x_9)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
_ = __local_var_7_7
var __t11 gopurs_runtime.Value
{
if (__local_var_7_7 == nil) {
__t11 = (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0
goto end_branch_11
} else {

}
}
{
if (__local_var_7_7 != nil) {
__t11 = gopurs_runtime.Apply(__local_var_6_6, (__local_var_7_7).V0)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
})
})}
_ = __local_var_3_2
// TAST (Let): __local_var_4_12 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_4_12 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 930809136 && v_5.UnsafePtr != nil) {
__t13 = (*Constructor_Data_Maybe_Just)(v_5.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_15 -> gopurs_runtime.Value
__local_var_5_15 := gopurs_runtime.Int((a_4.IntVal) + (1))
_ = __local_var_5_15
var __t19 *Constructor_Data_Maybe_Just
{
var __t16 bool
{
if (__local_var_5_15.IntVal) < (0) {
__t16 = false
goto end_branch_16
} else {

}
}
{
__t16 = true
}
end_branch_16:
var __t_and_18 bool = false
if __t16 {

var __t17 bool
{
if (__local_var_5_15.IntVal) > (59) {
__t17 = false
goto end_branch_17
} else {

}
}
{
__t17 = true
}
end_branch_17:
__t_and_18 = __t17
}
if __t_and_18 {
__t19 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_5_15.IntVal)}
goto end_branch_19
} else {

}
}
{
__t19 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_19:
// TAST (Let): __local_var_5_14 -> *Constructor_Data_Maybe_Just
var __local_var_5_14 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t19)})
var __t23 *Constructor_Data_Maybe_Just
{
if (__local_var_5_14 != nil) {
var __t22 *Constructor_Data_Maybe_Just
{
var __t21 *Constructor_Data_Maybe_Just
{
var __t20 bool
{
if ((__local_var_5_14).V0.IntVal) > (gopurs_runtime.Int(59).IntVal) {
__t20 = false
goto end_branch_20
} else {

}
}
{
__t20 = true
}
end_branch_20:
if __t20 {
__t21 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_21
} else {

}
}
{
__t21 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_21:
if (__t21 != nil) {
__t22 = &Constructor_Data_Maybe_Just{1, (__local_var_5_14).V0}
goto end_branch_22
} else {

}
}
{
__t22 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_22:
__t23 = __t22
goto end_branch_23
} else {

}
}
{
if (__local_var_5_14 == nil) {
__t23 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t23)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_2.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_5.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)}).IntVal) - (1))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_24_0 gopurs_runtime.Value
go__go_6_24_0 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop int64 = v_7_loop_val.IntVal
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_24_0:
for {
if false { continue go__go_6_24_0 }
var v_7 int64 = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t28 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 759514854 && v1_8.UnsafePtr != nil) {
var __t27 gopurs_runtime.Value
{
var __t_tag_25 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V1
if (__t_tag_25 == nil) {
__t27 = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
var __t26 bool
{
if (v_7) > (0) {
__t26 = false
goto end_branch_26
} else {

}
}
{
__t26 = true
}
end_branch_26:
if __t26 {
__t27 = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
v_7_loop = (v_7) - (1)
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_24_0
__t27 = gopurs_runtime.Value{}
}
end_branch_27:
__t28 = __t27
goto end_branch_28
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 759514854 && v1_8.UnsafePtr == nil) {
__t28 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_3_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_9
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)})
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return __t28
}
}()
})
})
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Apply2(go__go_6_24_0, gopurs_runtime.Int(n_5.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_2.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)})))}))
}))
}

func Call_Data_Time_Component_Gen_genMinute(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): foldableNonEmpty1_3_3 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_3_3 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_4 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_5 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_5
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_4.V0), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_5.V0), gopurs_runtime.Apply(f_5, x_8), acc_9)
})
}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray8 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1
_ = arr_val_foldlArray8
res_go_foldlArray8 := gopurs_runtime.Apply2(f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0)
_ = res_go_foldlArray8
arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
_ = arr_go_foldlArray8
for _, v_foldlArray8 := range *arr_go_foldlArray8 {
res_go_foldlArray8 = gopurs_runtime.Apply2(f_3, res_go_foldlArray8, v_foldlArray8)
}
return res_go_foldlArray8
}()
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_3_3
// TAST (Let): __local_var_3_2 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_3_2 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_3_3)}
}), gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray8 := (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1
_ = arr_val_foldlArray8
res_go_foldlArray8 := gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0)
_ = res_go_foldlArray8
arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
_ = arr_go_foldlArray8
for _, v_foldlArray8 := range *arr_go_foldlArray8 {
res_go_foldlArray8 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_4, "append"), s_7, gopurs_runtime.Apply(f_5, a1_8))
})
}), res_go_foldlArray8, v_foldlArray8)
}
return res_go_foldlArray8
}()
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray7 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1
_ = arr_val_foldlArray7
res_go_foldlArray7 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0
_ = res_go_foldlArray7
arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
_ = arr_go_foldlArray7
for _, v_foldlArray7 := range *arr_go_foldlArray7 {
res_go_foldlArray7 = gopurs_runtime.Apply2(f_4, res_go_foldlArray7, v_foldlArray7)
}
return res_go_foldlArray7
}()
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0)
_ = __local_var_6_6
// TAST (Let): __local_var_7_7 -> *Constructor_Data_Maybe_Just
__local_var_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(f_4, a1_7)
_ = __local_var_8_9
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 930809136 && v2_9.UnsafePtr == nil) {
__t10 = a1_7
goto end_branch_10
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 930809136 && v2_9.UnsafePtr != nil) {
__t10 = gopurs_runtime.Apply(__local_var_8_9, (*Constructor_Data_Maybe_Just)(v2_9.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
_ = __local_var_8_8
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_8_8, x_9)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
_ = __local_var_7_7
var __t11 gopurs_runtime.Value
{
if (__local_var_7_7 == nil) {
__t11 = (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0
goto end_branch_11
} else {

}
}
{
if (__local_var_7_7 != nil) {
__t11 = gopurs_runtime.Apply(__local_var_6_6, (__local_var_7_7).V0)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
})
})}
_ = __local_var_3_2
// TAST (Let): __local_var_4_12 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_4_12 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 930809136 && v_5.UnsafePtr != nil) {
__t13 = (*Constructor_Data_Maybe_Just)(v_5.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_15 -> gopurs_runtime.Value
__local_var_5_15 := gopurs_runtime.Int((a_4.IntVal) + (1))
_ = __local_var_5_15
var __t19 *Constructor_Data_Maybe_Just
{
var __t16 bool
{
if (__local_var_5_15.IntVal) < (0) {
__t16 = false
goto end_branch_16
} else {

}
}
{
__t16 = true
}
end_branch_16:
var __t_and_18 bool = false
if __t16 {

var __t17 bool
{
if (__local_var_5_15.IntVal) > (59) {
__t17 = false
goto end_branch_17
} else {

}
}
{
__t17 = true
}
end_branch_17:
__t_and_18 = __t17
}
if __t_and_18 {
__t19 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_5_15.IntVal)}
goto end_branch_19
} else {

}
}
{
__t19 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_19:
// TAST (Let): __local_var_5_14 -> *Constructor_Data_Maybe_Just
var __local_var_5_14 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t19)})
var __t23 *Constructor_Data_Maybe_Just
{
if (__local_var_5_14 != nil) {
var __t22 *Constructor_Data_Maybe_Just
{
var __t21 *Constructor_Data_Maybe_Just
{
var __t20 bool
{
if ((__local_var_5_14).V0.IntVal) > (gopurs_runtime.Int(59).IntVal) {
__t20 = false
goto end_branch_20
} else {

}
}
{
__t20 = true
}
end_branch_20:
if __t20 {
__t21 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_21
} else {

}
}
{
__t21 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_21:
if (__t21 != nil) {
__t22 = &Constructor_Data_Maybe_Just{1, (__local_var_5_14).V0}
goto end_branch_22
} else {

}
}
{
__t22 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_22:
__t23 = __t22
goto end_branch_23
} else {

}
}
{
if (__local_var_5_14 == nil) {
__t23 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t23)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_2.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_5.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)}).IntVal) - (1))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_24_1 gopurs_runtime.Value
go__go_6_24_1 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop int64 = v_7_loop_val.IntVal
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_24_1:
for {
if false { continue go__go_6_24_1 }
var v_7 int64 = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t28 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 759514854 && v1_8.UnsafePtr != nil) {
var __t27 gopurs_runtime.Value
{
var __t_tag_25 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V1
if (__t_tag_25 == nil) {
__t27 = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
var __t26 bool
{
if (v_7) > (0) {
__t26 = false
goto end_branch_26
} else {

}
}
{
__t26 = true
}
end_branch_26:
if __t26 {
__t27 = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
v_7_loop = (v_7) - (1)
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_24_1
__t27 = gopurs_runtime.Value{}
}
end_branch_27:
__t28 = __t27
goto end_branch_28
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 759514854 && v1_8.UnsafePtr == nil) {
__t28 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_3_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_9
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)})
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return __t28
}
}()
})
})
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Apply2(go__go_6_24_1, gopurs_runtime.Int(n_5.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_2.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)})))}))
}))
}

func Call_Data_Time_Component_Gen_genMillisecond(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): foldableNonEmpty1_3_3 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_3_3 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_4 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_5 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_5
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_4.V0), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_5.V0), gopurs_runtime.Apply(f_5, x_8), acc_9)
})
}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray8 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1
_ = arr_val_foldlArray8
res_go_foldlArray8 := gopurs_runtime.Apply2(f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0)
_ = res_go_foldlArray8
arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
_ = arr_go_foldlArray8
for _, v_foldlArray8 := range *arr_go_foldlArray8 {
res_go_foldlArray8 = gopurs_runtime.Apply2(f_3, res_go_foldlArray8, v_foldlArray8)
}
return res_go_foldlArray8
}()
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_3_3
// TAST (Let): __local_var_3_2 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_3_2 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_3_3)}
}), gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray8 := (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1
_ = arr_val_foldlArray8
res_go_foldlArray8 := gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0)
_ = res_go_foldlArray8
arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
_ = arr_go_foldlArray8
for _, v_foldlArray8 := range *arr_go_foldlArray8 {
res_go_foldlArray8 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_4, "append"), s_7, gopurs_runtime.Apply(f_5, a1_8))
})
}), res_go_foldlArray8, v_foldlArray8)
}
return res_go_foldlArray8
}()
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray7 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1
_ = arr_val_foldlArray7
res_go_foldlArray7 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0
_ = res_go_foldlArray7
arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
_ = arr_go_foldlArray7
for _, v_foldlArray7 := range *arr_go_foldlArray7 {
res_go_foldlArray7 = gopurs_runtime.Apply2(f_4, res_go_foldlArray7, v_foldlArray7)
}
return res_go_foldlArray7
}()
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0)
_ = __local_var_6_6
// TAST (Let): __local_var_7_7 -> *Constructor_Data_Maybe_Just
__local_var_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(f_4, a1_7)
_ = __local_var_8_9
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 930809136 && v2_9.UnsafePtr == nil) {
__t10 = a1_7
goto end_branch_10
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 930809136 && v2_9.UnsafePtr != nil) {
__t10 = gopurs_runtime.Apply(__local_var_8_9, (*Constructor_Data_Maybe_Just)(v2_9.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
_ = __local_var_8_8
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_8_8, x_9)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
_ = __local_var_7_7
var __t11 gopurs_runtime.Value
{
if (__local_var_7_7 == nil) {
__t11 = (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0
goto end_branch_11
} else {

}
}
{
if (__local_var_7_7 != nil) {
__t11 = gopurs_runtime.Apply(__local_var_6_6, (__local_var_7_7).V0)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
})
})}
_ = __local_var_3_2
// TAST (Let): __local_var_4_12 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_4_12 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 930809136 && v_5.UnsafePtr != nil) {
__t13 = (*Constructor_Data_Maybe_Just)(v_5.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_15 -> gopurs_runtime.Value
__local_var_5_15 := gopurs_runtime.Int((a_4.IntVal) + (1))
_ = __local_var_5_15
var __t19 *Constructor_Data_Maybe_Just
{
var __t16 bool
{
if (__local_var_5_15.IntVal) < (0) {
__t16 = false
goto end_branch_16
} else {

}
}
{
__t16 = true
}
end_branch_16:
var __t_and_18 bool = false
if __t16 {

var __t17 bool
{
if (__local_var_5_15.IntVal) > (999) {
__t17 = false
goto end_branch_17
} else {

}
}
{
__t17 = true
}
end_branch_17:
__t_and_18 = __t17
}
if __t_and_18 {
__t19 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_5_15.IntVal)}
goto end_branch_19
} else {

}
}
{
__t19 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_19:
// TAST (Let): __local_var_5_14 -> *Constructor_Data_Maybe_Just
var __local_var_5_14 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t19)})
var __t23 *Constructor_Data_Maybe_Just
{
if (__local_var_5_14 != nil) {
var __t22 *Constructor_Data_Maybe_Just
{
var __t21 *Constructor_Data_Maybe_Just
{
var __t20 bool
{
if ((__local_var_5_14).V0.IntVal) > (gopurs_runtime.Int(999).IntVal) {
__t20 = false
goto end_branch_20
} else {

}
}
{
__t20 = true
}
end_branch_20:
if __t20 {
__t21 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_21
} else {

}
}
{
__t21 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_21:
if (__t21 != nil) {
__t22 = &Constructor_Data_Maybe_Just{1, (__local_var_5_14).V0}
goto end_branch_22
} else {

}
}
{
__t22 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_22:
__t23 = __t22
goto end_branch_23
} else {

}
}
{
if (__local_var_5_14 == nil) {
__t23 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t23)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_2.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_5.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)}).IntVal) - (1))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_24_2 gopurs_runtime.Value
go__go_6_24_2 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop int64 = v_7_loop_val.IntVal
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_24_2:
for {
if false { continue go__go_6_24_2 }
var v_7 int64 = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t28 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 759514854 && v1_8.UnsafePtr != nil) {
var __t27 gopurs_runtime.Value
{
var __t_tag_25 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V1
if (__t_tag_25 == nil) {
__t27 = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
var __t26 bool
{
if (v_7) > (0) {
__t26 = false
goto end_branch_26
} else {

}
}
{
__t26 = true
}
end_branch_26:
if __t26 {
__t27 = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
v_7_loop = (v_7) - (1)
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_24_2
__t27 = gopurs_runtime.Value{}
}
end_branch_27:
__t28 = __t27
goto end_branch_28
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 759514854 && v1_8.UnsafePtr == nil) {
__t28 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_3_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_9
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)})
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return __t28
}
}()
})
})
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Apply2(go__go_6_24_2, gopurs_runtime.Int(n_5.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_2.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)})))}))
}))
}

func Call_Data_Time_Component_Gen_genHour(dictMonadGen_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonadGen_0 gopurs_runtime.Value = dictMonadGen_0_loop
_ = dictMonadGen_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonadGen_0, "Monad0"), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): pure_2_1 -> gopurs_runtime.Value
pure_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}), "pure")
_ = pure_2_1
// TAST (Let): foldableNonEmpty1_3_3 -> *Constructor_Data_Foldable_Foldable
foldableNonEmpty1_3_3 := &Constructor_Data_Foldable_Foldable{1, gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_4_4 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_4_4
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_7_5 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_7_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_7_5
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_4_4.V0), gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0), gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_7_5.V0), gopurs_runtime.Apply(f_5, x_8), acc_9)
})
}), gopurs_runtime.RecordGet(dictMonoid_3, "mempty"), (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray8 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1
_ = arr_val_foldlArray8
res_go_foldlArray8 := gopurs_runtime.Apply2(f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0)
_ = res_go_foldlArray8
arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
_ = arr_go_foldlArray8
for _, v_foldlArray8 := range *arr_go_foldlArray8 {
res_go_foldlArray8 = gopurs_runtime.Apply2(f_3, res_go_foldlArray8, v_foldlArray8)
}
return res_go_foldlArray8
}()
})
})
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_3, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0, gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), f_3, b_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
})
})
})}
_ = foldableNonEmpty1_3_3
// TAST (Let): __local_var_3_2 -> *Constructor_Data_Semigroup_Foldable_Foldable1
__local_var_3_2 := &Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(foldableNonEmpty1_3_3)}
}), gopurs_runtime.Func(func(dictSemigroup_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray8 := (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V1
_ = arr_val_foldlArray8
res_go_foldlArray8 := gopurs_runtime.Apply(f_5, (*Constructor_Data_NonEmpty_NonEmpty)(v_6.UnsafePtr).V0)
_ = res_go_foldlArray8
arr_go_foldlArray8 := (*[]gopurs_runtime.Value)(arr_val_foldlArray8.UnsafePtr)
_ = arr_go_foldlArray8
for _, v_foldlArray8 := range *arr_go_foldlArray8 {
res_go_foldlArray8 = gopurs_runtime.Apply2(gopurs_runtime.Func(func(s_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_4, "append"), s_7, gopurs_runtime.Apply(f_5, a1_8))
})
}), res_go_foldlArray8, v_foldlArray8)
}
return res_go_foldlArray8
}()
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
arr_val_foldlArray7 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1
_ = arr_val_foldlArray7
res_go_foldlArray7 := (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0
_ = res_go_foldlArray7
arr_go_foldlArray7 := (*[]gopurs_runtime.Value)(arr_val_foldlArray7.UnsafePtr)
_ = arr_go_foldlArray7
for _, v_foldlArray7 := range *arr_go_foldlArray7 {
res_go_foldlArray7 = gopurs_runtime.Apply2(f_4, res_go_foldlArray7, v_foldlArray7)
}
return res_go_foldlArray7
}()
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_6 -> gopurs_runtime.Value
__local_var_6_6 := gopurs_runtime.Apply(f_4, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0)
_ = __local_var_6_6
// TAST (Let): __local_var_7_7 -> *Constructor_Data_Maybe_Just
__local_var_7_7 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(a1_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_9 -> gopurs_runtime.Value
__local_var_8_9 := gopurs_runtime.Apply(f_4, a1_7)
_ = __local_var_8_9
// TAST (Let): __local_var_8_8 -> gopurs_runtime.Value
__local_var_8_8 := gopurs_runtime.Func(func(v2_9 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 gopurs_runtime.Value
{
if (v2_9.Type == 9 && v2_9.IntVal == 930809136 && v2_9.UnsafePtr == nil) {
__t10 = a1_7
goto end_branch_10
} else {

}
}
{
if (v2_9.Type == 9 && v2_9.IntVal == 930809136 && v2_9.UnsafePtr != nil) {
__t10 = gopurs_runtime.Apply(__local_var_8_9, (*Constructor_Data_Maybe_Just)(v2_9.UnsafePtr).V0)
goto end_branch_10
} else {

}
}
{
__t10 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
})
_ = __local_var_8_8
return gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Apply(__local_var_8_8, x_9)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V1))
_ = __local_var_7_7
var __t11 gopurs_runtime.Value
{
if (__local_var_7_7 == nil) {
__t11 = (*Constructor_Data_NonEmpty_NonEmpty)(v_5.UnsafePtr).V0
goto end_branch_11
} else {

}
}
{
if (__local_var_7_7 != nil) {
__t11 = gopurs_runtime.Apply(__local_var_6_6, (__local_var_7_7).V0)
goto end_branch_11
} else {

}
}
{
__t11 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_11:
return __t11
})
})}
_ = __local_var_3_2
// TAST (Let): __local_var_4_12 -> *Constructor_Data_NonEmpty_NonEmpty
var __local_var_4_12 *Constructor_Data_NonEmpty_NonEmpty = gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty](gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, gopurs_runtime.Int(0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable1_unfoldr1ArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t13 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 930809136 && v_5.UnsafePtr != nil) {
__t13 = (*Constructor_Data_Maybe_Just)(v_5.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_15 -> gopurs_runtime.Value
__local_var_5_15 := gopurs_runtime.Int((a_4.IntVal) + (1))
_ = __local_var_5_15
var __t19 *Constructor_Data_Maybe_Just
{
var __t16 bool
{
if (__local_var_5_15.IntVal) < (0) {
__t16 = false
goto end_branch_16
} else {

}
}
{
__t16 = true
}
end_branch_16:
var __t_and_18 bool = false
if __t16 {

var __t17 bool
{
if (__local_var_5_15.IntVal) > (23) {
__t17 = false
goto end_branch_17
} else {

}
}
{
__t17 = true
}
end_branch_17:
__t_and_18 = __t17
}
if __t_and_18 {
__t19 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Int(__local_var_5_15.IntVal)}
goto end_branch_19
} else {

}
}
{
__t19 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_19:
// TAST (Let): __local_var_5_14 -> *Constructor_Data_Maybe_Just
var __local_var_5_14 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t19)})
var __t23 *Constructor_Data_Maybe_Just
{
if (__local_var_5_14 != nil) {
var __t22 *Constructor_Data_Maybe_Just
{
var __t21 *Constructor_Data_Maybe_Just
{
var __t20 bool
{
if ((__local_var_5_14).V0.IntVal) > (gopurs_runtime.Int(23).IntVal) {
__t20 = false
goto end_branch_20
} else {

}
}
{
__t20 = true
}
end_branch_20:
if __t20 {
__t21 = &Constructor_Data_Maybe_Just{1, Get_Data_Unit_unit()}
goto end_branch_21
} else {

}
}
{
__t21 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_21:
if (__t21 != nil) {
__t22 = &Constructor_Data_Maybe_Just{1, (__local_var_5_14).V0}
goto end_branch_22
} else {

}
}
{
__t22 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_22:
__t23 = __t22
goto end_branch_23
} else {

}
}
{
if (__local_var_5_14 == nil) {
__t23 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_23
} else {

}
}
{
__t23 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_23:
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, a_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t23)}})}
}), gopurs_runtime.Int(gopurs_runtime.Int(1).IntVal)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictMonadGen_0, "chooseInt"), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_2.V0), gopurs_runtime.Value{}), "foldl"), gopurs_runtime.Func(func(c_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((1) + (c_5.IntVal))
})
}), gopurs_runtime.Int(0), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)}).IntVal) - (1))), gopurs_runtime.Func(func(n_5 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_6_24_3 gopurs_runtime.Value
go__go_6_24_3 = gopurs_runtime.Func(func(v_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_8_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_7_loop int64 = v_7_loop_val.IntVal
var v1_8_loop gopurs_runtime.Value = v1_8_loop_val
go__go_6_24_3:
for {
if false { continue go__go_6_24_3 }
var v_7 int64 = v_7_loop
_ = v_7
var v1_8 gopurs_runtime.Value = v1_8_loop
_ = v1_8
var __t28 gopurs_runtime.Value
{
if (v1_8.Type == 9 && v1_8.IntVal == 759514854 && v1_8.UnsafePtr != nil) {
var __t27 gopurs_runtime.Value
{
var __t_tag_25 *Constructor_Control_Monad_Gen_Cons = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V1
if (__t_tag_25 == nil) {
__t27 = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
var __t26 bool
{
if (v_7) > (0) {
__t26 = false
goto end_branch_26
} else {

}
}
{
__t26 = true
}
end_branch_26:
if __t26 {
__t27 = (*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V0
goto end_branch_27
} else {

}
}
{
v_7_loop = (v_7) - (1)
v1_8_loop = gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(v1_8.UnsafePtr).V1)}
continue go__go_6_24_3
__t27 = gopurs_runtime.Value{}
}
end_branch_27:
__t28 = __t27
goto end_branch_28
} else {

}
}
{
if (v1_8.Type == 9 && v1_8.IntVal == 759514854 && v1_8.UnsafePtr == nil) {
__t28 = gopurs_runtime.Apply3(gopurs_runtime.Box(__local_var_3_2.V1), gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_Last_semigroupLast()))}, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_9
}), gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)})
goto end_branch_28
} else {

}
}
{
__t28 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_28:
return __t28
}
}()
})
})
return gopurs_runtime.Apply(pure_2_1, gopurs_runtime.Apply2(go__go_6_24_3, gopurs_runtime.Int(n_5.IntVal), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Gen_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_3_2.V0), gopurs_runtime.Value{}), "foldr"), Get_Control_Monad_Gen_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 759514854, UnsafePtr: unsafe.Pointer((*Constructor_Control_Monad_Gen_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(__local_var_4_12)})))}))
}))
}



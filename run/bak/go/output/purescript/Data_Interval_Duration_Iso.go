package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Interval_Duration_Iso_empty gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_empty sync.Once
func Get_Data_Interval_Duration_Iso_empty() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_empty.Do(func() {
		cache_Data_Interval_Duration_Iso_empty = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}
	})
	return cache_Data_Interval_Duration_Iso_empty
}

var cache_Data_Interval_Duration_Iso_foldMap gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_foldMap sync.Once
func Get_Data_Interval_Duration_Iso_foldMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_foldMap.Do(func() {
		cache_Data_Interval_Duration_Iso_foldMap = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_Duration_Iso_foldMap(f_0_box)
})
	})
	return cache_Data_Interval_Duration_Iso_foldMap
}

var cache_Data_Interval_Duration_Iso_monoidAdditive gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_monoidAdditive sync.Once
func Get_Data_Interval_Duration_Iso_monoidAdditive() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_monoidAdditive.Do(func() {
		cache_Data_Interval_Duration_Iso_monoidAdditive = func() gopurs_runtime.Value {
// TAST (Let): semigroupAdditive1_0_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupAdditive1_0_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})
})}
_ = semigroupAdditive1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAdditive1_0_0)}
}), gopurs_runtime.Float(0.0)})}
}()
	})
	return cache_Data_Interval_Duration_Iso_monoidAdditive
}

var cache_Data_Interval_Duration_Iso_heytingAlgebraFunction gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_heytingAlgebraFunction sync.Once
func Get_Data_Interval_Duration_Iso_heytingAlgebraFunction() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_heytingAlgebraFunction.Do(func() {
		cache_Data_Interval_Duration_Iso_heytingAlgebraFunction = gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(&Constructor_Data_HeytingAlgebra_HeytingAlgebra{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) && ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) || ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) != (true)) || ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0, a_1).IntVal) != (0)) != (true))
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})})}
	})
	return cache_Data_Interval_Duration_Iso_heytingAlgebraFunction
}

var cache_Data_Interval_Duration_Iso_monoidFn gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_monoidFn sync.Once
func Get_Data_Interval_Duration_Iso_monoidFn() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_monoidFn.Do(func() {
		cache_Data_Interval_Duration_Iso_monoidFn = func() gopurs_runtime.Value {
// TAST (Let): semigroupFn_0_0 -> *Constructor_Data_Semigroup_Semigroup
semigroupFn_0_0 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_3_1_3 gopurs_runtime.Value
go__go_3_1_3 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_4_loop gopurs_runtime.Value = b_4_loop_val
var v_5_loop gopurs_runtime.Value = v_5_loop_val
go__go_3_1_3:
for {
if false { continue go__go_3_1_3 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 gopurs_runtime.Value = v_5_loop
_ = v_5
var __t2 gopurs_runtime.Value
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr == nil) {
__t2 = b_4
goto end_branch_2
} else {

}
}
{
if (v_5.Type == 9 && v_5.IntVal == 1358893437 && v_5.UnsafePtr != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_4)})}
v_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_5.UnsafePtr).V1)}
continue go__go_3_1_3
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
var go__go_4_3_4 gopurs_runtime.Value
go__go_4_3_4 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_5_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_5_loop_val)
var v1_6_loop gopurs_runtime.Value = v1_6_loop_val
go__go_4_3_4:
for {
if false { continue go__go_4_3_4 }
var v_5 *Constructor_Data_List_Types_Cons = v_5_loop
_ = v_5
var v1_6 gopurs_runtime.Value = v1_6_loop
_ = v1_6
var __t4 *Constructor_Data_List_Types_Cons
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr == nil) {
__t4 = v_5
goto end_branch_4
} else {

}
}
{
if (v1_6.Type == 9 && v1_6.IntVal == 1358893437 && v1_6.UnsafePtr != nil) {
v_5_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V0, v_5})})
v1_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_6.UnsafePtr).V1)}
continue go__go_4_3_4
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
return gopurs_runtime.Apply2(go__go_3_1_3, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(g_1, x_2)))}, gopurs_runtime.Apply2(go__go_4_3_4, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(f_0, x_2)))}))
})
})
})}
_ = semigroupFn_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(&Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupFn_0_0)}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}
})})}
}()
	})
	return cache_Data_Interval_Duration_Iso_monoidFn
}

var cache_Data_Interval_Duration_Iso_IsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_IsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_IsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_IsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_IsoDuration = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_Duration_Iso_IsoDuration(x_0_box)
})
	})
	return cache_Data_Interval_Duration_Iso_IsoDuration
}

var cache_Data_Interval_Duration_Iso_IsEmpty gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_IsEmpty sync.Once
func Get_Data_Interval_Duration_Iso_IsEmpty() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_IsEmpty.Do(func() {
		cache_Data_Interval_Duration_Iso_IsEmpty = gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Data_Interval_Duration_Iso_IsEmpty
}

var cache_Data_Interval_Duration_Iso_InvalidWeekComponentUsage gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_InvalidWeekComponentUsage sync.Once
func Get_Data_Interval_Duration_Iso_InvalidWeekComponentUsage() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_InvalidWeekComponentUsage.Do(func() {
		cache_Data_Interval_Duration_Iso_InvalidWeekComponentUsage = gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_Data_Interval_Duration_Iso_InvalidWeekComponentUsage
}

var cache_Data_Interval_Duration_Iso_ContainsNegativeValue gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_ContainsNegativeValue sync.Once
func Get_Data_Interval_Duration_Iso_ContainsNegativeValue() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_ContainsNegativeValue.Do(func() {
		cache_Data_Interval_Duration_Iso_ContainsNegativeValue = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue{1, uint32(value0.IntVal)})}
})
	})
	return cache_Data_Interval_Duration_Iso_ContainsNegativeValue
}

var cache_Data_Interval_Duration_Iso_InvalidFractionalUse gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_InvalidFractionalUse sync.Once
func Get_Data_Interval_Duration_Iso_InvalidFractionalUse() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_InvalidFractionalUse.Do(func() {
		cache_Data_Interval_Duration_Iso_InvalidFractionalUse = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse{1, uint32(value0.IntVal)})}
})
	})
	return cache_Data_Interval_Duration_Iso_InvalidFractionalUse
}

var cache_Data_Interval_Duration_Iso_unIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_unIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_unIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_unIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_unIsoDuration = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_unIsoDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Interval_Duration_Iso_unIsoDuration
}

var cache_Data_Interval_Duration_Iso_showIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_showIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_showIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_showIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_showIsoDuration = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(IsoDuration (fromFoldable ") + (gopurs_runtime.Apply2(Get_Data_Show_showArrayImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t7 string
{
var __t_tag_0 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_0.IntVal) == 217821258) {
__t7 = "(Tuple Minute "
goto end_branch_7
} else {

}
}
{
var __t_tag_1 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_1.IntVal) == 3908053364) {
__t7 = "(Tuple Second "
goto end_branch_7
} else {

}
}
{
var __t_tag_2 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_2.IntVal) == 1292308612) {
__t7 = "(Tuple Hour "
goto end_branch_7
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_3.IntVal) == 2311060696) {
__t7 = "(Tuple Day "
goto end_branch_7
} else {

}
}
{
var __t_tag_4 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_4.IntVal) == 401302776) {
__t7 = "(Tuple Week "
goto end_branch_7
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_5.IntVal) == 3327533908) {
__t7 = "(Tuple Month "
goto end_branch_7
} else {

}
}
{
var __t_tag_6 gopurs_runtime.Value = (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V0
if (uint32(__t_tag_6.IntVal) == 3631736139) {
__t7 = "(Tuple Year "
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_7:
return gopurs_runtime.Str(((__t7) + (gopurs_runtime.Apply(Get_Data_Show_showNumberImpl(), (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V1).StrVal())) + (")"))
}), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply6(Get_Data_Unfoldable_unfoldrArrayImpl(), Get_Data_Maybe_isNothing(), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t8 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 930809136 && v_2.UnsafePtr != nil) {
__t8 = (*Constructor_Data_Maybe_Just)(v_2.UnsafePtr).V0
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
})), Get_Data_Tuple_fst(), Get_Data_Tuple_snd(), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())) + ("))"))
})})}
	})
	return cache_Data_Interval_Duration_Iso_showIsoDuration
}

var cache_Data_Interval_Duration_Iso_showError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_showError sync.Once
func Get_Data_Interval_Duration_Iso_showError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_showError.Do(func() {
		cache_Data_Interval_Duration_Iso_showError = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t16 string
{
if (v_0.Type == 9 && v_0.IntVal == 1422140417) {
__t16 = "(IsEmpty)"
goto end_branch_16
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1775501833) {
__t16 = "(InvalidWeekComponentUsage)"
goto end_branch_16
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3224543173) {
var __t7 string
{
var __t_tag_0 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_0) == 217821258) {
__t7 = "(ContainsNegativeValue Minute)"
goto end_branch_7
} else {

}
}
{
var __t_tag_1 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_1) == 3908053364) {
__t7 = "(ContainsNegativeValue Second)"
goto end_branch_7
} else {

}
}
{
var __t_tag_2 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_2) == 1292308612) {
__t7 = "(ContainsNegativeValue Hour)"
goto end_branch_7
} else {

}
}
{
var __t_tag_3 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_3) == 2311060696) {
__t7 = "(ContainsNegativeValue Day)"
goto end_branch_7
} else {

}
}
{
var __t_tag_4 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_4) == 401302776) {
__t7 = "(ContainsNegativeValue Week)"
goto end_branch_7
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_5) == 3327533908) {
__t7 = "(ContainsNegativeValue Month)"
goto end_branch_7
} else {

}
}
{
var __t_tag_6 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_6) == 3631736139) {
__t7 = "(ContainsNegativeValue Year)"
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_7:
__t16 = __t7
goto end_branch_16
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
var __t15 string
{
var __t_tag_8 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_8) == 217821258) {
__t15 = "(InvalidFractionalUse Minute)"
goto end_branch_15
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_9) == 3908053364) {
__t15 = "(InvalidFractionalUse Second)"
goto end_branch_15
} else {

}
}
{
var __t_tag_10 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_10) == 1292308612) {
__t15 = "(InvalidFractionalUse Hour)"
goto end_branch_15
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_11) == 2311060696) {
__t15 = "(InvalidFractionalUse Day)"
goto end_branch_15
} else {

}
}
{
var __t_tag_12 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_12) == 401302776) {
__t15 = "(InvalidFractionalUse Week)"
goto end_branch_15
} else {

}
}
{
var __t_tag_13 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_13) == 3327533908) {
__t15 = "(InvalidFractionalUse Month)"
goto end_branch_15
} else {

}
}
{
var __t_tag_14 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_14) == 3631736139) {
__t15 = "(InvalidFractionalUse Year)"
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_16:
return gopurs_runtime.Str(__t16)
})})}
	})
	return cache_Data_Interval_Duration_Iso_showError
}

var cache_Data_Interval_Duration_Iso_prettyError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_prettyError sync.Once
func Get_Data_Interval_Duration_Iso_prettyError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_prettyError.Do(func() {
		cache_Data_Interval_Duration_Iso_prettyError = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_Interval_Duration_Iso_prettyError(v_0_box))
})
	})
	return cache_Data_Interval_Duration_Iso_prettyError
}

var cache_Data_Interval_Duration_Iso_eqIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_eqIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_eqIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_eqIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_eqIsoDuration = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_1_5 gopurs_runtime.Value
go__go_2_1_5 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_1_5:
for {
if false { continue go__go_2_1_5 }
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
// TAST (Let): v_5_2 -> *Constructor_Data_Map_Internal_IterNext
v_5_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_2
var __t28 bool
{
if (v_5_2 != nil) {
// TAST (Let): v2_6_3 -> *Constructor_Data_Map_Internal_IterNext
v2_6_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_3
var __t27 bool
{
var __t_and_26 bool = false
if (v2_6_3 != nil) {

var __t25 bool
{
var __t_tag_4 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_4.IntVal) == 3908053364) {
var __t6 bool
{
var __t_tag_5 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_5.IntVal) == 3908053364) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
__t25 = __t6
goto end_branch_25
} else {

}
}
{
var __t_tag_7 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_7.IntVal) == 217821258) {
var __t9 bool
{
var __t_tag_8 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_8.IntVal) == 217821258) {
__t9 = true
goto end_branch_9
} else {

}
}
{
__t9 = false
}
end_branch_9:
__t25 = __t9
goto end_branch_25
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_10.IntVal) == 1292308612) {
var __t12 bool
{
var __t_tag_11 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_11.IntVal) == 1292308612) {
__t12 = true
goto end_branch_12
} else {

}
}
{
__t12 = false
}
end_branch_12:
__t25 = __t12
goto end_branch_25
} else {

}
}
{
var __t_tag_13 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_13.IntVal) == 2311060696) {
var __t15 bool
{
var __t_tag_14 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_14.IntVal) == 2311060696) {
__t15 = true
goto end_branch_15
} else {

}
}
{
__t15 = false
}
end_branch_15:
__t25 = __t15
goto end_branch_25
} else {

}
}
{
var __t_tag_16 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_16.IntVal) == 401302776) {
var __t18 bool
{
var __t_tag_17 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_17.IntVal) == 401302776) {
__t18 = true
goto end_branch_18
} else {

}
}
{
__t18 = false
}
end_branch_18:
__t25 = __t18
goto end_branch_25
} else {

}
}
{
var __t_tag_19 gopurs_runtime.Value = (v_5_2).V0
if (uint32(__t_tag_19.IntVal) == 3327533908) {
var __t21 bool
{
var __t_tag_20 gopurs_runtime.Value = (v2_6_3).V0
if (uint32(__t_tag_20.IntVal) == 3327533908) {
__t21 = true
goto end_branch_21
} else {

}
}
{
__t21 = false
}
end_branch_21:
__t25 = __t21
goto end_branch_25
} else {

}
}
{
var __t_tag_22 gopurs_runtime.Value = (v_5_2).V0
var __t_and_24 bool = false
if (uint32(__t_tag_22.IntVal) == 3631736139) {

var __t_tag_23 gopurs_runtime.Value = (v2_6_3).V0
__t_and_24 = (uint32(__t_tag_23.IntVal) == 3631736139)
}
if __t_and_24 {
__t25 = true
goto end_branch_25
} else {

}
}
{
__t25 = false
}
end_branch_25:
__t_and_26 = (__t25) && (((v_5_2).V1.FloatVal()) == ((v2_6_3).V1.FloatVal()))
}
if __t_and_26 {
a_3_loop = (v_5_2).V2
b_4_loop = (v2_6_3).V2
continue go__go_2_1_5
__t27 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t28 = __t27
goto end_branch_28
} else {

}
}
{
if (v_5_2 == nil) {
__t28 = true
goto end_branch_28
} else {

}
}
{
__t28 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_28:
return gopurs_runtime.Bool(__t28)
}
}()
})
})
// TAST (Let): eqMapIter2_2_0 -> *Constructor_Data_Eq_Eq
eqMapIter2_2_0 := &Constructor_Data_Eq_Eq{1, go__go_2_1_5}
_ = eqMapIter2_2_0
var __t35 bool
{
var __t_tag_29 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0)
if (__t_tag_29 == nil) {
var __t31 bool
{
var __t_tag_30 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1)
if (__t_tag_30 == nil) {
__t31 = true
goto end_branch_31
} else {

}
}
{
__t31 = false
}
end_branch_31:
__t35 = __t31
goto end_branch_35
} else {

}
}
{
var __t_tag_32 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0)
if (__t_tag_32 != nil) {
var __t34 bool
{
var __t_tag_33 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1)
if ((__t_tag_33 != nil)) && (((*Constructor_Data_Map_Internal_Node)(x_0.UnsafePtr).V1) == ((*Constructor_Data_Map_Internal_Node)(y_1.UnsafePtr).V1)) {
__t34 = (gopurs_runtime.Apply2(gopurs_runtime.Box(eqMapIter2_2_0.V0), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal) != (0)
goto end_branch_34
} else {

}
}
{
__t34 = false
}
end_branch_34:
__t35 = __t34
goto end_branch_35
} else {

}
}
{
__t35 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_35:
return gopurs_runtime.Bool(__t35)
})
})})}
	})
	return cache_Data_Interval_Duration_Iso_eqIsoDuration
}

var cache_Data_Interval_Duration_Iso_ordIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_ordIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_ordIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_ordIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_ordIsoDuration = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Interval_Duration_Iso_eqIsoDuration()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_2_6 gopurs_runtime.Value
go__go_2_2_6 = gopurs_runtime.Func(func(a_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_3_loop gopurs_runtime.Value = a_3_loop_val
var b_4_loop gopurs_runtime.Value = b_4_loop_val
go__go_2_2_6:
for {
if false { continue go__go_2_2_6 }
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
// TAST (Let): v_5_3 -> *Constructor_Data_Map_Internal_IterNext
v_5_3 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_3))
_ = v_5_3
var __t29 bool
{
if (v_5_3 != nil) {
// TAST (Let): v2_6_4 -> *Constructor_Data_Map_Internal_IterNext
v2_6_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_4))
_ = v2_6_4
var __t28 bool
{
var __t_and_27 bool = false
if (v2_6_4 != nil) {

var __t26 bool
{
var __t_tag_5 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_5.IntVal) == 3908053364) {
var __t7 bool
{
var __t_tag_6 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_6.IntVal) == 3908053364) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t26 = __t7
goto end_branch_26
} else {

}
}
{
var __t_tag_8 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_8.IntVal) == 217821258) {
var __t10 bool
{
var __t_tag_9 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_9.IntVal) == 217821258) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
__t26 = __t10
goto end_branch_26
} else {

}
}
{
var __t_tag_11 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_11.IntVal) == 1292308612) {
var __t13 bool
{
var __t_tag_12 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_12.IntVal) == 1292308612) {
__t13 = true
goto end_branch_13
} else {

}
}
{
__t13 = false
}
end_branch_13:
__t26 = __t13
goto end_branch_26
} else {

}
}
{
var __t_tag_14 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_14.IntVal) == 2311060696) {
var __t16 bool
{
var __t_tag_15 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_15.IntVal) == 2311060696) {
__t16 = true
goto end_branch_16
} else {

}
}
{
__t16 = false
}
end_branch_16:
__t26 = __t16
goto end_branch_26
} else {

}
}
{
var __t_tag_17 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_17.IntVal) == 401302776) {
var __t19 bool
{
var __t_tag_18 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_18.IntVal) == 401302776) {
__t19 = true
goto end_branch_19
} else {

}
}
{
__t19 = false
}
end_branch_19:
__t26 = __t19
goto end_branch_26
} else {

}
}
{
var __t_tag_20 gopurs_runtime.Value = (v_5_3).V0
if (uint32(__t_tag_20.IntVal) == 3327533908) {
var __t22 bool
{
var __t_tag_21 gopurs_runtime.Value = (v2_6_4).V0
if (uint32(__t_tag_21.IntVal) == 3327533908) {
__t22 = true
goto end_branch_22
} else {

}
}
{
__t22 = false
}
end_branch_22:
__t26 = __t22
goto end_branch_26
} else {

}
}
{
var __t_tag_23 gopurs_runtime.Value = (v_5_3).V0
var __t_and_25 bool = false
if (uint32(__t_tag_23.IntVal) == 3631736139) {

var __t_tag_24 gopurs_runtime.Value = (v2_6_4).V0
__t_and_25 = (uint32(__t_tag_24.IntVal) == 3631736139)
}
if __t_and_25 {
__t26 = true
goto end_branch_26
} else {

}
}
{
__t26 = false
}
end_branch_26:
__t_and_27 = (__t26) && (((v_5_3).V1.FloatVal()) == ((v2_6_4).V1.FloatVal()))
}
if __t_and_27 {
a_3_loop = (v_5_3).V2
b_4_loop = (v2_6_4).V2
continue go__go_2_2_6
__t28 = (gopurs_runtime.Value{}.IntVal) != (0)
goto end_branch_28
} else {

}
}
{
__t28 = false
}
end_branch_28:
__t29 = __t28
goto end_branch_29
} else {

}
}
{
if (v_5_3 == nil) {
__t29 = true
goto end_branch_29
} else {

}
}
{
__t29 = (func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal) != (0)
}
end_branch_29:
return gopurs_runtime.Bool(__t29)
}
}()
})
})
// TAST (Let): eqMapIter2_2_1 -> *Constructor_Data_Eq_Eq
eqMapIter2_2_1 := &Constructor_Data_Eq_Eq{1, go__go_2_2_6}
_ = eqMapIter2_2_1
var go__go_3_30_7 gopurs_runtime.Value
go__go_3_30_7 = gopurs_runtime.Func(func(a_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var a_4_loop gopurs_runtime.Value = a_4_loop_val
var b_5_loop gopurs_runtime.Value = b_5_loop_val
go__go_3_30_7:
for {
if false { continue go__go_3_30_7 }
var a_4 gopurs_runtime.Value = a_4_loop
_ = a_4
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
// TAST (Let): v_6_31 -> *Constructor_Data_Map_Internal_IterNext
v_6_31 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), b_5))
_ = v_6_31
// TAST (Let): v1_7_32 -> *Constructor_Data_Map_Internal_IterNext
v1_7_32 := gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_IterNext](gopurs_runtime.Apply(Get_Data_Map_Internal_stepAsc(), a_4))
_ = v1_7_32
var __t67 uint32
{
if (v1_7_32 != nil) {
var __t65 uint32
{
if (v_6_31 != nil) {
var __t61 uint32
{
var __t_tag_34 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_34.IntVal) == 3908053364) {
var __t36 uint32
{
var __t_tag_35 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_35.IntVal) == 3908053364) {
__t36 = 902936544
goto end_branch_36
} else {

}
}
{
__t36 = 1527465420
}
end_branch_36:
__t61 = __t36
goto end_branch_61
} else {

}
}
{
var __t_tag_37 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_37.IntVal) == 3908053364) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_38 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_38.IntVal) == 217821258) {
var __t40 uint32
{
var __t_tag_39 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_39.IntVal) == 217821258) {
__t40 = 902936544
goto end_branch_40
} else {

}
}
{
__t40 = 1527465420
}
end_branch_40:
__t61 = __t40
goto end_branch_61
} else {

}
}
{
var __t_tag_41 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_41.IntVal) == 217821258) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_42 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_42.IntVal) == 1292308612) {
var __t44 uint32
{
var __t_tag_43 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_43.IntVal) == 1292308612) {
__t44 = 902936544
goto end_branch_44
} else {

}
}
{
__t44 = 1527465420
}
end_branch_44:
__t61 = __t44
goto end_branch_61
} else {

}
}
{
var __t_tag_45 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_45.IntVal) == 1292308612) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_46 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_46.IntVal) == 2311060696) {
var __t48 uint32
{
var __t_tag_47 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_47.IntVal) == 2311060696) {
__t48 = 902936544
goto end_branch_48
} else {

}
}
{
__t48 = 1527465420
}
end_branch_48:
__t61 = __t48
goto end_branch_61
} else {

}
}
{
var __t_tag_49 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_49.IntVal) == 2311060696) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_50 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_50.IntVal) == 401302776) {
var __t52 uint32
{
var __t_tag_51 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_51.IntVal) == 401302776) {
__t52 = 902936544
goto end_branch_52
} else {

}
}
{
__t52 = 1527465420
}
end_branch_52:
__t61 = __t52
goto end_branch_61
} else {

}
}
{
var __t_tag_53 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_53.IntVal) == 401302776) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_54 gopurs_runtime.Value = (v1_7_32).V0
if (uint32(__t_tag_54.IntVal) == 3327533908) {
var __t56 uint32
{
var __t_tag_55 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_55.IntVal) == 3327533908) {
__t56 = 902936544
goto end_branch_56
} else {

}
}
{
__t56 = 1527465420
}
end_branch_56:
__t61 = __t56
goto end_branch_61
} else {

}
}
{
var __t_tag_57 gopurs_runtime.Value = (v_6_31).V0
if (uint32(__t_tag_57.IntVal) == 3327533908) {
__t61 = 380165415
goto end_branch_61
} else {

}
}
{
var __t_tag_58 gopurs_runtime.Value = (v1_7_32).V0
var __t_and_60 bool = false
if (uint32(__t_tag_58.IntVal) == 3631736139) {

var __t_tag_59 gopurs_runtime.Value = (v_6_31).V0
__t_and_60 = (uint32(__t_tag_59.IntVal) == 3631736139)
}
if __t_and_60 {
__t61 = 902936544
goto end_branch_61
} else {

}
}
{
__t61 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_61:
// TAST (Let): v3_8_33 -> uint32
v3_8_33 := __t61
_ = v3_8_33
var __t64 uint32
{
if (v3_8_33 == 902936544) {
// TAST (Let): v4_9_62 -> gopurs_runtime.Value
v4_9_62 := gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, (v1_7_32).V1, (v_6_31).V1)
_ = v4_9_62
var __t63 uint32
{
if (uint32(v4_9_62.IntVal) == 902936544) {
a_4_loop = (v1_7_32).V2
b_5_loop = (v_6_31).V2
continue go__go_3_30_7
__t63 = uint32(gopurs_runtime.Value{}.IntVal)
goto end_branch_63
} else {

}
}
{
__t63 = uint32(v4_9_62.IntVal)
}
end_branch_63:
__t64 = __t63
goto end_branch_64
} else {

}
}
{
__t64 = v3_8_33
}
end_branch_64:
__t65 = __t64
goto end_branch_65
} else {

}
}
{
if (v_6_31 == nil) {
__t65 = 380165415
goto end_branch_65
} else {

}
}
{
__t65 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_65:
__t67 = __t65
goto end_branch_67
} else {

}
}
{
if (v1_7_32 == nil) {
var __t66 uint32
{
if (v_6_31 == nil) {
__t66 = 902936544
goto end_branch_66
} else {

}
}
{
__t66 = 1527465420
}
end_branch_66:
__t67 = __t66
goto end_branch_67
} else {

}
}
{
if (v_6_31 == nil) {
__t67 = 380165415
goto end_branch_67
} else {

}
}
{
__t67 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_67:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t67), UnsafePtr: nil}
}
}()
})
})
// TAST (Let): ordMapIter2_2_0 -> *Constructor_Data_Ord_Ord
ordMapIter2_2_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqMapIter2_2_1)}
}), go__go_3_30_7}
_ = ordMapIter2_2_0
var __t73 uint32
{
var __t_tag_70 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0)
if (__t_tag_70 == nil) {
var __t72 uint32
{
var __t_tag_71 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1)
if (__t_tag_71 == nil) {
__t72 = 902936544
goto end_branch_72
} else {

}
}
{
__t72 = 1527465420
}
end_branch_72:
__t73 = __t72
goto end_branch_73
} else {

}
}
{
var __t69 uint32
{
var __t_tag_68 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1)
if (__t_tag_68 == nil) {
__t69 = 380165415
goto end_branch_69
} else {

}
}
{
__t69 = uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(ordMapIter2_2_0.V1), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1))}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}).IntVal)
}
end_branch_69:
__t73 = __t69
}
end_branch_73:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t73), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Interval_Duration_Iso_ordIsoDuration
}

var cache_Data_Interval_Duration_Iso_eqError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_eqError sync.Once
func Get_Data_Interval_Duration_Iso_eqError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_eqError.Do(func() {
		cache_Data_Interval_Duration_Iso_eqError = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t47 bool
{
if (x_0.Type == 9 && x_0.IntVal == 1422140417) {
var __t0 bool
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t47 = __t0
goto end_branch_47
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1775501833) {
var __t1 bool
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t47 = __t1
goto end_branch_47
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t24 bool
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
var __t23 bool
{
var __t_tag_2 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_2) == 3908053364) {
var __t4 bool
{
var __t_tag_3 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_3) == 3908053364) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t23 = __t4
goto end_branch_23
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_5) == 217821258) {
var __t7 bool
{
var __t_tag_6 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_6) == 217821258) {
__t7 = true
goto end_branch_7
} else {

}
}
{
__t7 = false
}
end_branch_7:
__t23 = __t7
goto end_branch_23
} else {

}
}
{
var __t_tag_8 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_8) == 1292308612) {
var __t10 bool
{
var __t_tag_9 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_9) == 1292308612) {
__t10 = true
goto end_branch_10
} else {

}
}
{
__t10 = false
}
end_branch_10:
__t23 = __t10
goto end_branch_23
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_11) == 2311060696) {
var __t13 bool
{
var __t_tag_12 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_12) == 2311060696) {
__t13 = true
goto end_branch_13
} else {

}
}
{
__t13 = false
}
end_branch_13:
__t23 = __t13
goto end_branch_23
} else {

}
}
{
var __t_tag_14 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_14) == 401302776) {
var __t16 bool
{
var __t_tag_15 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_15) == 401302776) {
__t16 = true
goto end_branch_16
} else {

}
}
{
__t16 = false
}
end_branch_16:
__t23 = __t16
goto end_branch_23
} else {

}
}
{
var __t_tag_17 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_17) == 3327533908) {
var __t19 bool
{
var __t_tag_18 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_18) == 3327533908) {
__t19 = true
goto end_branch_19
} else {

}
}
{
__t19 = false
}
end_branch_19:
__t23 = __t19
goto end_branch_23
} else {

}
}
{
var __t_tag_20 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
var __t_and_22 bool = false
if (uint32(__t_tag_20) == 3631736139) {

var __t_tag_21 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
__t_and_22 = (uint32(__t_tag_21) == 3631736139)
}
if __t_and_22 {
__t23 = true
goto end_branch_23
} else {

}
}
{
__t23 = false
}
end_branch_23:
__t24 = __t23
goto end_branch_24
} else {

}
}
{
__t24 = false
}
end_branch_24:
__t47 = __t24
goto end_branch_47
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 574232667)) && ((y_1.Type == 9 && y_1.IntVal == 574232667)) {
var __t46 bool
{
var __t_tag_25 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_25) == 3908053364) {
var __t27 bool
{
var __t_tag_26 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_26) == 3908053364) {
__t27 = true
goto end_branch_27
} else {

}
}
{
__t27 = false
}
end_branch_27:
__t46 = __t27
goto end_branch_46
} else {

}
}
{
var __t_tag_28 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_28) == 217821258) {
var __t30 bool
{
var __t_tag_29 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_29) == 217821258) {
__t30 = true
goto end_branch_30
} else {

}
}
{
__t30 = false
}
end_branch_30:
__t46 = __t30
goto end_branch_46
} else {

}
}
{
var __t_tag_31 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_31) == 1292308612) {
var __t33 bool
{
var __t_tag_32 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_32) == 1292308612) {
__t33 = true
goto end_branch_33
} else {

}
}
{
__t33 = false
}
end_branch_33:
__t46 = __t33
goto end_branch_46
} else {

}
}
{
var __t_tag_34 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_34) == 2311060696) {
var __t36 bool
{
var __t_tag_35 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_35) == 2311060696) {
__t36 = true
goto end_branch_36
} else {

}
}
{
__t36 = false
}
end_branch_36:
__t46 = __t36
goto end_branch_46
} else {

}
}
{
var __t_tag_37 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_37) == 401302776) {
var __t39 bool
{
var __t_tag_38 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_38) == 401302776) {
__t39 = true
goto end_branch_39
} else {

}
}
{
__t39 = false
}
end_branch_39:
__t46 = __t39
goto end_branch_46
} else {

}
}
{
var __t_tag_40 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_40) == 3327533908) {
var __t42 bool
{
var __t_tag_41 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_41) == 3327533908) {
__t42 = true
goto end_branch_42
} else {

}
}
{
__t42 = false
}
end_branch_42:
__t46 = __t42
goto end_branch_46
} else {

}
}
{
var __t_tag_43 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
var __t_and_45 bool = false
if (uint32(__t_tag_43) == 3631736139) {

var __t_tag_44 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
__t_and_45 = (uint32(__t_tag_44) == 3631736139)
}
if __t_and_45 {
__t46 = true
goto end_branch_46
} else {

}
}
{
__t46 = false
}
end_branch_46:
__t47 = __t46
goto end_branch_47
} else {

}
}
{
__t47 = false
}
end_branch_47:
return gopurs_runtime.Bool(__t47)
})
})})}
	})
	return cache_Data_Interval_Duration_Iso_eqError
}

var cache_Data_Interval_Duration_Iso_ordError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_ordError sync.Once
func Get_Data_Interval_Duration_Iso_ordError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_ordError.Do(func() {
		cache_Data_Interval_Duration_Iso_ordError = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Interval_Duration_Iso_eqError()))}
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t59 uint32
{
if (x_0.Type == 9 && x_0.IntVal == 1422140417) {
var __t0 uint32
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t59 = __t0
goto end_branch_59
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t59 = 380165415
goto end_branch_59
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1775501833) {
var __t1 uint32
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t59 = __t1
goto end_branch_59
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t59 = 380165415
goto end_branch_59
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t30 uint32
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
var __t29 uint32
{
var __t_tag_2 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_2) == 3908053364) {
var __t4 uint32
{
var __t_tag_3 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_3) == 3908053364) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t29 = __t4
goto end_branch_29
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_5) == 3908053364) {
__t29 = 380165415
goto end_branch_29
} else {

}
}
{
var __t_tag_6 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_6) == 217821258) {
var __t8 uint32
{
var __t_tag_7 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_7) == 217821258) {
__t8 = 902936544
goto end_branch_8
} else {

}
}
{
__t8 = 1527465420
}
end_branch_8:
__t29 = __t8
goto end_branch_29
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_9) == 217821258) {
__t29 = 380165415
goto end_branch_29
} else {

}
}
{
var __t_tag_10 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_10) == 1292308612) {
var __t12 uint32
{
var __t_tag_11 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_11) == 1292308612) {
__t12 = 902936544
goto end_branch_12
} else {

}
}
{
__t12 = 1527465420
}
end_branch_12:
__t29 = __t12
goto end_branch_29
} else {

}
}
{
var __t_tag_13 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_13) == 1292308612) {
__t29 = 380165415
goto end_branch_29
} else {

}
}
{
var __t_tag_14 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_14) == 2311060696) {
var __t16 uint32
{
var __t_tag_15 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_15) == 2311060696) {
__t16 = 902936544
goto end_branch_16
} else {

}
}
{
__t16 = 1527465420
}
end_branch_16:
__t29 = __t16
goto end_branch_29
} else {

}
}
{
var __t_tag_17 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_17) == 2311060696) {
__t29 = 380165415
goto end_branch_29
} else {

}
}
{
var __t_tag_18 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_18) == 401302776) {
var __t20 uint32
{
var __t_tag_19 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_19) == 401302776) {
__t20 = 902936544
goto end_branch_20
} else {

}
}
{
__t20 = 1527465420
}
end_branch_20:
__t29 = __t20
goto end_branch_29
} else {

}
}
{
var __t_tag_21 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_21) == 401302776) {
__t29 = 380165415
goto end_branch_29
} else {

}
}
{
var __t_tag_22 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_22) == 3327533908) {
var __t24 uint32
{
var __t_tag_23 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_23) == 3327533908) {
__t24 = 902936544
goto end_branch_24
} else {

}
}
{
__t24 = 1527465420
}
end_branch_24:
__t29 = __t24
goto end_branch_29
} else {

}
}
{
var __t_tag_25 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
if (uint32(__t_tag_25) == 3327533908) {
__t29 = 380165415
goto end_branch_29
} else {

}
}
{
var __t_tag_26 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
var __t_and_28 bool = false
if (uint32(__t_tag_26) == 3631736139) {

var __t_tag_27 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
__t_and_28 = (uint32(__t_tag_27) == 3631736139)
}
if __t_and_28 {
__t29 = 902936544
goto end_branch_29
} else {

}
}
{
__t29 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_29:
__t30 = __t29
goto end_branch_30
} else {

}
}
{
__t30 = 1527465420
}
end_branch_30:
__t59 = __t30
goto end_branch_59
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t59 = 380165415
goto end_branch_59
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 574232667)) && ((y_1.Type == 9 && y_1.IntVal == 574232667)) {
var __t58 uint32
{
var __t_tag_31 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_31) == 3908053364) {
var __t33 uint32
{
var __t_tag_32 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_32) == 3908053364) {
__t33 = 902936544
goto end_branch_33
} else {

}
}
{
__t33 = 1527465420
}
end_branch_33:
__t58 = __t33
goto end_branch_58
} else {

}
}
{
var __t_tag_34 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_34) == 3908053364) {
__t58 = 380165415
goto end_branch_58
} else {

}
}
{
var __t_tag_35 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_35) == 217821258) {
var __t37 uint32
{
var __t_tag_36 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_36) == 217821258) {
__t37 = 902936544
goto end_branch_37
} else {

}
}
{
__t37 = 1527465420
}
end_branch_37:
__t58 = __t37
goto end_branch_58
} else {

}
}
{
var __t_tag_38 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_38) == 217821258) {
__t58 = 380165415
goto end_branch_58
} else {

}
}
{
var __t_tag_39 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_39) == 1292308612) {
var __t41 uint32
{
var __t_tag_40 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_40) == 1292308612) {
__t41 = 902936544
goto end_branch_41
} else {

}
}
{
__t41 = 1527465420
}
end_branch_41:
__t58 = __t41
goto end_branch_58
} else {

}
}
{
var __t_tag_42 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_42) == 1292308612) {
__t58 = 380165415
goto end_branch_58
} else {

}
}
{
var __t_tag_43 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_43) == 2311060696) {
var __t45 uint32
{
var __t_tag_44 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_44) == 2311060696) {
__t45 = 902936544
goto end_branch_45
} else {

}
}
{
__t45 = 1527465420
}
end_branch_45:
__t58 = __t45
goto end_branch_58
} else {

}
}
{
var __t_tag_46 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_46) == 2311060696) {
__t58 = 380165415
goto end_branch_58
} else {

}
}
{
var __t_tag_47 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_47) == 401302776) {
var __t49 uint32
{
var __t_tag_48 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_48) == 401302776) {
__t49 = 902936544
goto end_branch_49
} else {

}
}
{
__t49 = 1527465420
}
end_branch_49:
__t58 = __t49
goto end_branch_58
} else {

}
}
{
var __t_tag_50 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_50) == 401302776) {
__t58 = 380165415
goto end_branch_58
} else {

}
}
{
var __t_tag_51 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_51) == 3327533908) {
var __t53 uint32
{
var __t_tag_52 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_52) == 3327533908) {
__t53 = 902936544
goto end_branch_53
} else {

}
}
{
__t53 = 1527465420
}
end_branch_53:
__t58 = __t53
goto end_branch_58
} else {

}
}
{
var __t_tag_54 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
if (uint32(__t_tag_54) == 3327533908) {
__t58 = 380165415
goto end_branch_58
} else {

}
}
{
var __t_tag_55 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
var __t_and_57 bool = false
if (uint32(__t_tag_55) == 3631736139) {

var __t_tag_56 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
__t_and_57 = (uint32(__t_tag_56) == 3631736139)
}
if __t_and_57 {
__t58 = 902936544
goto end_branch_58
} else {

}
}
{
__t58 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_58:
__t59 = __t58
goto end_branch_59
} else {

}
}
{
__t59 = uint32(func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal)
}
end_branch_59:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t59), UnsafePtr: nil}
})
})})}
	})
	return cache_Data_Interval_Duration_Iso_ordError
}

var cache_Data_Interval_Duration_Iso_checkWeekUsage gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkWeekUsage sync.Once
func Get_Data_Interval_Duration_Iso_checkWeekUsage() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkWeekUsage.Do(func() {
		cache_Data_Interval_Duration_Iso_checkWeekUsage = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkWeekUsage(v_0_box))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkWeekUsage
}

var cache_Data_Interval_Duration_Iso_checkNegativeValues gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkNegativeValues sync.Once
func Get_Data_Interval_Duration_Iso_checkNegativeValues() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkNegativeValues.Do(func() {
		cache_Data_Interval_Duration_Iso_checkNegativeValues = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkNegativeValues(v_0_box))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkNegativeValues
}

var cache_Data_Interval_Duration_Iso_checkFractionalUse gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkFractionalUse sync.Once
func Get_Data_Interval_Duration_Iso_checkFractionalUse() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkFractionalUse.Do(func() {
		cache_Data_Interval_Duration_Iso_checkFractionalUse = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkFractionalUse(v_0_box))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkFractionalUse
}

var cache_Data_Interval_Duration_Iso_checkEmptiness gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkEmptiness sync.Once
func Get_Data_Interval_Duration_Iso_checkEmptiness() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkEmptiness.Do(func() {
		cache_Data_Interval_Duration_Iso_checkEmptiness = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkEmptiness(v_0_box))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkEmptiness
}

var cache_Data_Interval_Duration_Iso_checkValidIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkValidIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_checkValidIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkValidIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_checkValidIsoDuration = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkValidIsoDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0_box)))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkValidIsoDuration
}

var cache_Data_Interval_Duration_Iso_mkIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_mkIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_mkIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_mkIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_mkIsoDuration = gopurs_runtime.Func(func(d_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Interval_Duration_Iso_mkIsoDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](d_0_box))
})
	})
	return cache_Data_Interval_Duration_Iso_mkIsoDuration
}

type Constructor_Data_Interval_Duration_Iso_IsEmpty struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Iso_InvalidWeekComponentUsage struct {
	Rc uint32
}


type Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue struct {
	Rc uint32
	V0 uint32
}


type Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse struct {
	Rc uint32
	V0 uint32
}


func Call_Data_Interval_Duration_Iso_foldMap(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_0 gopurs_runtime.Value
go__go_1_0_0 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t5 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t5 = b_2
goto end_branch_5
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
var go__go_4_1_1 gopurs_runtime.Value
go__go_4_1_1 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_1_1:
for {
if false { continue go__go_4_1_1 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t2 = b_5
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_1_1
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
var go__go_5_3_2 gopurs_runtime.Value
go__go_5_3_2 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_3_2:
for {
if false { continue go__go_5_3_2 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t4 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t4 = v_6
goto end_branch_4
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_3_2
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t4)}
}
}()
})
})
b_2_loop = gopurs_runtime.Apply2(go__go_4_1_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(f_0, (*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0)))}, gopurs_runtime.Apply2(go__go_5_3_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_2))}))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_0
__t5 = gopurs_runtime.Value{}
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
}
}()
})
})
return gopurs_runtime.Apply(go__go_1_0_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))})
}

func Call_Data_Interval_Duration_Iso_IsoDuration(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Interval_Duration_Iso_unIsoDuration(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_Map_Internal_Node {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
return v_0
}

func Call_Data_Interval_Duration_Iso_prettyError(v_0_loop gopurs_runtime.Value) string {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t16 string
{
if (v_0.Type == 9 && v_0.IntVal == 1422140417) {
__t16 = "Duration is empty (has no components)"
goto end_branch_16
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1775501833) {
__t16 = "Week component of Duration is used with other components"
goto end_branch_16
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3224543173) {
var __t7 string
{
var __t_tag_0 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_0) == 217821258) {
__t7 = "Component `Minute` contains negative value"
goto end_branch_7
} else {

}
}
{
var __t_tag_1 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_1) == 3908053364) {
__t7 = "Component `Second` contains negative value"
goto end_branch_7
} else {

}
}
{
var __t_tag_2 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_2) == 1292308612) {
__t7 = "Component `Hour` contains negative value"
goto end_branch_7
} else {

}
}
{
var __t_tag_3 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_3) == 2311060696) {
__t7 = "Component `Day` contains negative value"
goto end_branch_7
} else {

}
}
{
var __t_tag_4 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_4) == 401302776) {
__t7 = "Component `Week` contains negative value"
goto end_branch_7
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_5) == 3327533908) {
__t7 = "Component `Month` contains negative value"
goto end_branch_7
} else {

}
}
{
var __t_tag_6 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(v_0.UnsafePtr).V0
if (uint32(__t_tag_6) == 3631736139) {
__t7 = "Component `Year` contains negative value"
goto end_branch_7
} else {

}
}
{
__t7 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_7:
__t16 = __t7
goto end_branch_16
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
var __t15 string
{
var __t_tag_8 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_8) == 217821258) {
__t15 = "Invalid usage of Fractional value at component `Minute`"
goto end_branch_15
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_9) == 3908053364) {
__t15 = "Invalid usage of Fractional value at component `Second`"
goto end_branch_15
} else {

}
}
{
var __t_tag_10 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_10) == 1292308612) {
__t15 = "Invalid usage of Fractional value at component `Hour`"
goto end_branch_15
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_11) == 2311060696) {
__t15 = "Invalid usage of Fractional value at component `Day`"
goto end_branch_15
} else {

}
}
{
var __t_tag_12 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_12) == 401302776) {
__t15 = "Invalid usage of Fractional value at component `Week`"
goto end_branch_15
} else {

}
}
{
var __t_tag_13 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_13) == 3327533908) {
__t15 = "Invalid usage of Fractional value at component `Month`"
goto end_branch_15
} else {

}
}
{
var __t_tag_14 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(v_0.UnsafePtr).V0
if (uint32(__t_tag_14) == 3631736139) {
__t15 = "Invalid usage of Fractional value at component `Year`"
goto end_branch_15
} else {

}
}
{
__t15 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
__t16 = func() gopurs_runtime.Value { panic("Failed pattern match") }().StrVal()
}
end_branch_16:
return __t16
}

func Call_Data_Interval_Duration_Iso_checkWeekUsage(v_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t7 *Constructor_Data_List_Types_Cons
{
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Maybe_Just
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(Get_Data_Map_Internal_lookup__1040249709(), gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.RecordGet(v_0, "asMap")))}))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0 == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
var __t_and_6 bool = false
if (__t1.IntVal) != (0) {

var __t5 bool
{
var __t4 int64
{
var __t_tag_2 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.RecordGet(v_0, "asMap"))
if (__t_tag_2 == nil) {
__t4 = 0
goto end_branch_4
} else {

}
}
{
var __t_tag_3 *Constructor_Data_Map_Internal_Node = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.RecordGet(v_0, "asMap"))
if (__t_tag_3 != nil) {
__t4 = (*Constructor_Data_Map_Internal_Node)(gopurs_runtime.RecordGet(v_0, "asMap").UnsafePtr).V1
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_4:
if (__t4) > (1) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t_and_6 = __t5
}
if __t_and_6 {
__t7 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: unsafe.Pointer(nil)}, (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_7
} else {

}
}
{
__t7 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_7:
return __t7
}

func Call_Data_Interval_Duration_Iso_checkNegativeValues(v_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var go__go_1_0_8 gopurs_runtime.Value
go__go_1_0_8 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_8:
for {
if false { continue go__go_1_0_8 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t8 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t8 = b_2
goto end_branch_8
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
var go__go_4_1_9 gopurs_runtime.Value
go__go_4_1_9 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_1_9:
for {
if false { continue go__go_4_1_9 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t2 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t2 = b_5
goto end_branch_2
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_1_9
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}
}()
})
})
var __t5 *Constructor_Data_List_Types_Cons
{
var __t4 bool
{
var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float((*Constructor_Data_Tuple_Tuple)((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0.UnsafePtr).V1.FloatVal()), gopurs_runtime.Float(0.0))
if (uint32(__t_tag_3.IntVal) == 1527465420) {
__t4 = false
goto end_branch_4
} else {

}
}
{
__t4 = true
}
end_branch_4:
if __t4 {
__t5 = (*Constructor_Data_List_Types_Cons)(nil)
goto end_branch_5
} else {

}
}
{
__t5 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue{1, uint32((*Constructor_Data_Tuple_Tuple)((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V0.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_List_Types_Cons)(nil)}
}
end_branch_5:
var go__go_5_6_10 gopurs_runtime.Value
go__go_5_6_10 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_6_10:
for {
if false { continue go__go_5_6_10 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t7 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t7 = v_6
goto end_branch_7
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_6_10
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_7:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t7)}
}
}()
})
})
b_2_loop = gopurs_runtime.Apply2(go__go_4_1_9, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}, gopurs_runtime.Apply2(go__go_5_6_10, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_2))}))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_3.UnsafePtr).V1)}
continue go__go_1_0_8
__t8 = gopurs_runtime.Value{}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return __t8
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_1_0_8, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_0, "asList")))}))
}

func Call_Data_Interval_Duration_Iso_checkFractionalUse(v_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 -> *Constructor_Data_List_Types_Cons
v1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_List_span__2133741451(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Bool(((gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float((*Constructor_Data_Tuple_Tuple)(x_1.UnsafePtr).V1.FloatVal())).FloatVal()) == ((*Constructor_Data_Tuple_Tuple)(x_1.UnsafePtr).V1.FloatVal())) != (true)).IntVal) != (0)) != (true))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_0, "asList")))}), "rest"))
_ = v1_1_0
var __t4 *Constructor_Data_List_Types_Cons
{
var __t_and_3 bool = false
if (v1_1_0 != nil) {

var __t2 bool
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Float(gopurs_runtime.Apply2(Get_Data_Foldable_foldMap__193737345(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Number_abs(), (*Constructor_Data_Tuple_Tuple)(x_2.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_1_0).V1)}).FloatVal()), gopurs_runtime.Float(0.0))
if (uint32(__t_tag_1.IntVal) == 380165415) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t_and_3 = __t2
}
if __t_and_3 {
__t4 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse{1, uint32((*Constructor_Data_Tuple_Tuple)((v1_1_0).V0.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_4
} else {

}
}
{
__t4 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_4:
return __t4
}

func Call_Data_Interval_Duration_Iso_checkEmptiness(v_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t2 *Constructor_Data_List_Types_Cons
{
var __t1 bool
{
var __t_tag_0 *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_0, "asList"))
if (__t_tag_0 == nil) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: unsafe.Pointer(nil)}, (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_List_Types_Cons)(nil)
}
end_branch_2:
return __t2
}

func Call_Data_Interval_Duration_Iso_checkValidIsoDuration(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
// TAST (Let): semigroupFn_1_1 -> *Constructor_Data_Semigroup_Semigroup
semigroupFn_1_1 := &Constructor_Data_Semigroup_Semigroup{1, gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_4_2_11 gopurs_runtime.Value
go__go_4_2_11 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_5_loop gopurs_runtime.Value = b_5_loop_val
var v_6_loop gopurs_runtime.Value = v_6_loop_val
go__go_4_2_11:
for {
if false { continue go__go_4_2_11 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 gopurs_runtime.Value = v_6_loop
_ = v_6
var __t3 gopurs_runtime.Value
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr == nil) {
__t3 = b_5
goto end_branch_3
} else {

}
}
{
if (v_6.Type == 9 && v_6.IntVal == 1358893437 && v_6.UnsafePtr != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](b_5)})}
v_6_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v_6.UnsafePtr).V1)}
continue go__go_4_2_11
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}
}()
})
})
var go__go_5_4_12 gopurs_runtime.Value
go__go_5_4_12 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_6_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_6_loop_val)
var v1_7_loop gopurs_runtime.Value = v1_7_loop_val
go__go_5_4_12:
for {
if false { continue go__go_5_4_12 }
var v_6 *Constructor_Data_List_Types_Cons = v_6_loop
_ = v_6
var v1_7 gopurs_runtime.Value = v1_7_loop
_ = v1_7
var __t5 *Constructor_Data_List_Types_Cons
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr == nil) {
__t5 = v_6
goto end_branch_5
} else {

}
}
{
if (v1_7.Type == 9 && v1_7.IntVal == 1358893437 && v1_7.UnsafePtr != nil) {
v_6_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, (*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V0, v_6})})
v1_7_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_7.UnsafePtr).V1)}
continue go__go_5_4_12
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t5)}
}
}()
})
})
return gopurs_runtime.Apply2(go__go_4_2_11, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(g_2, x_3)))}, gopurs_runtime.Apply2(go__go_5_4_12, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(f_1, x_3)))}))
})
})
})}
_ = semigroupFn_1_1
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Monoid_Monoid
__local_var_1_0 := &Constructor_Data_Monoid_Monoid{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupFn_1_1)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}
})}
_ = __local_var_1_0
// TAST (Let): Semigroup0_2_6 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_6 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(__local_var_1_0.V0), gopurs_runtime.Value{}))
_ = Semigroup0_2_6
var go__go_3_7_13 gopurs_runtime.Value
go__go_3_7_13 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_13:
for {
if false { continue go__go_3_7_13 }
var v_4 *Constructor_Data_List_Types_Cons = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 *Constructor_Data_List_Types_Cons
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = v_4
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)})})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_5.UnsafePtr).V1)}
continue go__go_3_7_13
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply4(Get_Data_Foldable_foldrArray(), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_6.V0), x_3, acc_4)
})
}), gopurs_runtime.Box(__local_var_1_0.V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Interval_Duration_Iso_checkWeekUsage(), Get_Data_Interval_Duration_Iso_checkEmptiness(), Get_Data_Interval_Duration_Iso_checkFractionalUse(), Get_Data_Interval_Duration_Iso_checkNegativeValues()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.RecordDict2("asList", "asMap", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_3_7_13, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply(Get_Data_Map_Internal_toUnfoldable__2567957978(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
}

func Call_Data_Interval_Duration_Iso_mkIsoDuration(d_0_loop *Constructor_Data_Map_Internal_Node) gopurs_runtime.Value {
var d_0 *Constructor_Data_Map_Internal_Node = d_0_loop
_ = d_0
// TAST (Let): __local_var_1_1 -> *Constructor_Data_List_Types_Cons
__local_var_1_1 := Call_Data_Interval_Duration_Iso_checkValidIsoDuration(d_0)
_ = __local_var_1_1
var __t2 *Constructor_Data_Maybe_Just
{
if (__local_var_1_1 == nil) {
__t2 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_2
} else {

}
}
{
if (__local_var_1_1 != nil) {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_Data_NonEmpty_NonEmpty{1, (__local_var_1_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_1_1).V1)}})}}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
// TAST (Let): v_1_0 -> *Constructor_Data_Maybe_Just
var v_1_0 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
var __t3 gopurs_runtime.Value
{
if (v_1_0 != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_NonEmpty_NonEmpty]((v_1_0).V0))}})}
goto end_branch_3
} else {

}
}
{
if (v_1_0 == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(d_0)}})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return __t3
}



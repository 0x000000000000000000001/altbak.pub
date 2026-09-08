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
		cache_Data_Interval_Duration_Iso_empty = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}
	})
	return cache_Data_Interval_Duration_Iso_empty
}

var cache_Data_Interval_Duration_Iso_foldMap gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_foldMap sync.Once
func Get_Data_Interval_Duration_Iso_foldMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_foldMap.Do(func() {
		cache_Data_Interval_Duration_Iso_foldMap = gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V0), Get_Data_List_Types_monoidList())
	})
	return cache_Data_Interval_Duration_Iso_foldMap
}

var cache_Data_Interval_Duration_Iso_monoidAdditive gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_monoidAdditive sync.Once
func Get_Data_Interval_Duration_Iso_monoidAdditive() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_monoidAdditive.Do(func() {
		cache_Data_Interval_Duration_Iso_monoidAdditive = func() gopurs_runtime.Value {
// TAST (Let): semigroupAdditive1_0_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupAdditive1_0_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})})
_ = semigroupAdditive1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_321927638_1201789390((&Constructor_Data_Monoid_Monoid[float64]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAdditive1_0_0)}
}), gopurs_runtime.Float(0.0).FloatVal()})))}
}()
	})
	return cache_Data_Interval_Duration_Iso_monoidAdditive
}

var cache_Data_Interval_Duration_Iso_heytingAlgebraFunction gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_heytingAlgebraFunction sync.Once
func Get_Data_Interval_Duration_Iso_heytingAlgebraFunction() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_heytingAlgebraFunction.Do(func() {
		cache_Data_Interval_Duration_Iso_heytingAlgebraFunction = gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer((&Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) && ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) || ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(false)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) != (true)) || ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(f_0, a_1).IntVal) != (0)) != (true))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
})}))}
	})
	return cache_Data_Interval_Duration_Iso_heytingAlgebraFunction
}

var cache_Data_Interval_Duration_Iso_monoidFn gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_monoidFn sync.Once
func Get_Data_Interval_Duration_Iso_monoidFn() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_monoidFn.Do(func() {
		cache_Data_Interval_Duration_Iso_monoidFn = func() gopurs_runtime.Value {
// TAST (Let): semigroupFn_0_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar a)] (TypeVar b))])
semigroupFn_0_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Interval_Duration_Iso_go__go_3_1_0 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Interval_Duration_Iso_go__go_3_1_0
var go__go_3_1_0 gopurs_runtime.Value
_ = go__go_3_1_0
Call_local_Data_Interval_Duration_Iso_go__go_3_1_0 = func(b_4_loop gopurs_runtime.Value, v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_3_1_0:
for {
if false { continue go__go_3_1_0 }
var b_4 gopurs_runtime.Value = b_4_loop
_ = b_4
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var __t2 gopurs_runtime.Value
{
if (v_5 == nil) {
__t2 = b_4
goto end_branch_2
} else {

}
}
{
if (v_5 != nil) {
b_4_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_5).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_4)}))}
v_5_loop = (v_5).V1
continue go__go_3_1_0
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_3_1_0 = gopurs_runtime.Func(func(b_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Interval_Duration_Iso_go__go_3_1_0(b_4_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val))
})
})
var Call_local_Data_Interval_Duration_Iso_go__go_4_3_1 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_Interval_Duration_Iso_go__go_4_3_1
var go__go_4_3_1 gopurs_runtime.Value
_ = go__go_4_3_1
Call_local_Data_Interval_Duration_Iso_go__go_4_3_1 = func(v_5_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_4_3_1:
for {
if false { continue go__go_4_3_1 }
var v_5 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_5_loop
_ = v_5
var v1_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_6_loop
_ = v1_6
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_6 == nil) {
__t4 = v_5
goto end_branch_4
} else {

}
}
{
if (v1_6 != nil) {
v_5_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_6).V0, v_5})
v1_6_loop = (v1_6).V1
continue go__go_4_3_1
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_4_3_1 = gopurs_runtime.Func(func(v_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_Interval_Duration_Iso_go__go_4_3_1(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_5_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_6_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_Interval_Duration_Iso_go__go_3_1_0(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(g_1, x_2)))}, Call_local_Data_Interval_Duration_Iso_go__go_4_3_1((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, x_2))))))}
})})
_ = semigroupFn_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupFn_0_0)}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))}
})}))}
}()
	})
	return cache_Data_Interval_Duration_Iso_monoidFn
}

var cache_Data_Interval_Duration_Iso_IsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_IsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_IsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_IsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_IsoDuration = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(Call_Data_Interval_Duration_Iso_IsoDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](x_0_box))))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue{1, uint32(value0.IntVal)}))}
})
	})
	return cache_Data_Interval_Duration_Iso_ContainsNegativeValue
}

var cache_Data_Interval_Duration_Iso_InvalidFractionalUse gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_InvalidFractionalUse sync.Once
func Get_Data_Interval_Duration_Iso_InvalidFractionalUse() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_InvalidFractionalUse.Do(func() {
		cache_Data_Interval_Duration_Iso_InvalidFractionalUse = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse{1, uint32(value0.IntVal)}))}
})
	})
	return cache_Data_Interval_Duration_Iso_InvalidFractionalUse
}

var cache_Data_Interval_Duration_Iso_unIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_unIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_unIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_unIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_unIsoDuration = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(Call_Data_Interval_Duration_Iso_unIsoDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](v_0_box))))}
})
	})
	return cache_Data_Interval_Duration_Iso_unIsoDuration
}

var cache_Data_Interval_Duration_Iso_showIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_showIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_showIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_showIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_showIsoDuration = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_983382226_1386611502((&Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[uint32, float64]]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(IsoDuration ") + (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_showMap()).V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(Rebox_Data_Interval_Duration_Iso_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0))))}).StrVal())) + (")"))
})})))}
	})
	return cache_Data_Interval_Duration_Iso_showIsoDuration
}

var cache_Data_Interval_Duration_Iso_showError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_showError sync.Once
func Get_Data_Interval_Duration_Iso_showError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_showError.Do(func() {
		cache_Data_Interval_Duration_Iso_showError = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t7 = func() string { panic("Failed pattern match") }()
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
__t15 = func() string { panic("Failed pattern match") }()
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
__t16 = func() string { panic("Failed pattern match") }()
}
end_branch_16:
return gopurs_runtime.Str(__t16)
})}))}
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
		cache_Data_Interval_Duration_Iso_eqIsoDuration = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_2958538738_3790796878((&Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_eqMap()).V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(Rebox_Data_Interval_Duration_Iso_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_0))))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(Rebox_Data_Interval_Duration_Iso_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](y_1))))}).IntVal) != (0))
})})))}
	})
	return cache_Data_Interval_Duration_Iso_eqIsoDuration
}

var cache_Data_Interval_Duration_Iso_ordIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_ordIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_ordIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_ordIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_ordIsoDuration = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_3498527378_4177771502((&Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[uint32, float64]]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_2958538738_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_Iso_eqIsoDuration())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[uint32, float64]]](Get_Data_Interval_Duration_ordMap()).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(Rebox_Data_Interval_Duration_Iso_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_0))))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(Rebox_Data_Interval_Duration_Iso_2487766124_261879545(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]](y_1))))}).IntVal)), UnsafePtr: nil}
})})))}
	})
	return cache_Data_Interval_Duration_Iso_ordIsoDuration
}

var cache_Data_Interval_Duration_Iso_eqError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_eqError sync.Once
func Get_Data_Interval_Duration_Iso_eqError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_eqError.Do(func() {
		cache_Data_Interval_Duration_Iso_eqError = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t35 bool
{
if (x_0.Type == 9 && x_0.IntVal == 1422140417) {
__t35 = (y_1.Type == 9 && y_1.IntVal == 1422140417)
goto end_branch_35
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1775501833) {
__t35 = (y_1.Type == 9 && y_1.IntVal == 1775501833)
goto end_branch_35
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t_and_34 bool = false
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {

var __t33 bool
{
var __t_tag_21 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_21) == 3908053364) {
var __t_tag_22 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
__t33 = (uint32(__t_tag_22) == 3908053364)
goto end_branch_33
} else {

}
}
{
var __t_tag_23 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_23) == 217821258) {
var __t_tag_24 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
__t33 = (uint32(__t_tag_24) == 217821258)
goto end_branch_33
} else {

}
}
{
var __t_tag_25 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_25) == 1292308612) {
var __t_tag_26 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
__t33 = (uint32(__t_tag_26) == 1292308612)
goto end_branch_33
} else {

}
}
{
var __t_tag_27 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_27) == 2311060696) {
var __t_tag_28 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
__t33 = (uint32(__t_tag_28) == 2311060696)
goto end_branch_33
} else {

}
}
{
var __t_tag_29 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_29) == 401302776) {
var __t_tag_30 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
__t33 = (uint32(__t_tag_30) == 401302776)
goto end_branch_33
} else {

}
}
{
var __t_tag_31 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
if (uint32(__t_tag_31) == 3327533908) {
var __t_tag_32 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
__t33 = (uint32(__t_tag_32) == 3327533908)
goto end_branch_33
} else {

}
}
{
var __t_tag_18 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(x_0.UnsafePtr).V0
var __t_and_20 bool = false
if (uint32(__t_tag_18) == 3631736139) {

var __t_tag_19 uint32 = (*Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue)(y_1.UnsafePtr).V0
__t_and_20 = (uint32(__t_tag_19) == 3631736139)
}
__t33 = __t_and_20
}
end_branch_33:
__t_and_34 = __t33
}
__t35 = __t_and_34
goto end_branch_35
} else {

}
}
{
var __t_and_17 bool = false
if (x_0.Type == 9 && x_0.IntVal == 574232667) {

var __t_and_16 bool = false
if (y_1.Type == 9 && y_1.IntVal == 574232667) {

var __t15 bool
{
var __t_tag_3 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_3) == 3908053364) {
var __t_tag_4 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
__t15 = (uint32(__t_tag_4) == 3908053364)
goto end_branch_15
} else {

}
}
{
var __t_tag_5 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_5) == 217821258) {
var __t_tag_6 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
__t15 = (uint32(__t_tag_6) == 217821258)
goto end_branch_15
} else {

}
}
{
var __t_tag_7 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_7) == 1292308612) {
var __t_tag_8 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
__t15 = (uint32(__t_tag_8) == 1292308612)
goto end_branch_15
} else {

}
}
{
var __t_tag_9 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_9) == 2311060696) {
var __t_tag_10 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
__t15 = (uint32(__t_tag_10) == 2311060696)
goto end_branch_15
} else {

}
}
{
var __t_tag_11 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_11) == 401302776) {
var __t_tag_12 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
__t15 = (uint32(__t_tag_12) == 401302776)
goto end_branch_15
} else {

}
}
{
var __t_tag_13 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
if (uint32(__t_tag_13) == 3327533908) {
var __t_tag_14 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
__t15 = (uint32(__t_tag_14) == 3327533908)
goto end_branch_15
} else {

}
}
{
var __t_tag_0 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(x_0.UnsafePtr).V0
var __t_and_2 bool = false
if (uint32(__t_tag_0) == 3631736139) {

var __t_tag_1 uint32 = (*Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse)(y_1.UnsafePtr).V0
__t_and_2 = (uint32(__t_tag_1) == 3631736139)
}
__t15 = __t_and_2
}
end_branch_15:
__t_and_16 = __t15
}
__t_and_17 = __t_and_16
}
__t35 = __t_and_17
}
end_branch_35:
return gopurs_runtime.Bool(__t35)
})}))}
	})
	return cache_Data_Interval_Duration_Iso_eqError
}

var cache_Data_Interval_Duration_Iso_ordError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_ordError sync.Once
func Get_Data_Interval_Duration_Iso_ordError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_ordError.Do(func() {
		cache_Data_Interval_Duration_Iso_ordError = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Interval_Duration_Iso_eqError()))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t29 = func() uint32 { panic("Failed pattern match") }()
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
__t58 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_58:
__t59 = __t58
goto end_branch_59
} else {

}
}
{
__t59 = func() uint32 { panic("Failed pattern match") }()
}
end_branch_59:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t59), UnsafePtr: nil}
})}))}
	})
	return cache_Data_Interval_Duration_Iso_ordError
}

var cache_Data_Interval_Duration_Iso_checkWeekUsage gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkWeekUsage sync.Once
func Get_Data_Interval_Duration_Iso_checkWeekUsage() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkWeekUsage.Do(func() {
		cache_Data_Interval_Duration_Iso_checkWeekUsage = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkWeekUsage(func() struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
} {
					orig := v_0_box
					_ = orig
					clone := struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
}{}
					clone.asList = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](gopurs_runtime.RecordGet(orig, "asList"))
					clone.asMap = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.RecordGet(orig, "asMap"))
					return clone
				}()))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkWeekUsage
}

var cache_Data_Interval_Duration_Iso_checkNegativeValues gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkNegativeValues sync.Once
func Get_Data_Interval_Duration_Iso_checkNegativeValues() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkNegativeValues.Do(func() {
		cache_Data_Interval_Duration_Iso_checkNegativeValues = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkNegativeValues(func() struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
} {
					orig := v_0_box
					_ = orig
					clone := struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
}{}
					clone.asList = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](gopurs_runtime.RecordGet(orig, "asList"))
					clone.asMap = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.RecordGet(orig, "asMap"))
					return clone
				}()))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkNegativeValues
}

var cache_Data_Interval_Duration_Iso_checkFractionalUse gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkFractionalUse sync.Once
func Get_Data_Interval_Duration_Iso_checkFractionalUse() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkFractionalUse.Do(func() {
		cache_Data_Interval_Duration_Iso_checkFractionalUse = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkFractionalUse(func() struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
} {
					orig := v_0_box
					_ = orig
					clone := struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
}{}
					clone.asList = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](gopurs_runtime.RecordGet(orig, "asList"))
					clone.asMap = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.RecordGet(orig, "asMap"))
					return clone
				}()))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkFractionalUse
}

var cache_Data_Interval_Duration_Iso_checkEmptiness gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkEmptiness sync.Once
func Get_Data_Interval_Duration_Iso_checkEmptiness() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkEmptiness.Do(func() {
		cache_Data_Interval_Duration_Iso_checkEmptiness = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkEmptiness(func() struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
} {
					orig := v_0_box
					_ = orig
					clone := struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
}{}
					clone.asList = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](gopurs_runtime.RecordGet(orig, "asList"))
					clone.asMap = gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](gopurs_runtime.RecordGet(orig, "asMap"))
					return clone
				}()))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkEmptiness
}

var cache_Data_Interval_Duration_Iso_checkValidIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_checkValidIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_checkValidIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_checkValidIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_checkValidIsoDuration = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_Data_Interval_Duration_Iso_checkValidIsoDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](v_0_box)))}
})
	})
	return cache_Data_Interval_Duration_Iso_checkValidIsoDuration
}

var cache_Data_Interval_Duration_Iso_mkIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_mkIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_mkIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_mkIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_mkIsoDuration = gopurs_runtime.Func(func(d_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Interval_Duration_Iso_mkIsoDuration(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](d_0_box))
				if _v.V2 {
					return gopurs_runtime.Box(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})
				}
				return gopurs_runtime.Box(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})
			}()
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


func Call_Data_Interval_Duration_Iso_IsoDuration(x_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var x_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Interval_Duration_Iso_unIsoDuration(v_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Map_Internal_Node[uint32, float64] {
var v_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = v_0_loop
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
__t7 = func() string { panic("Failed pattern match") }()
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
__t15 = func() string { panic("Failed pattern match") }()
}
end_branch_15:
__t16 = __t15
goto end_branch_16
} else {

}
}
{
__t16 = func() string { panic("Failed pattern match") }()
}
end_branch_16:
return __t16
}

func Call_Data_Interval_Duration_Iso_checkWeekUsage(v_0_loop struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
}) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
} = v_0_loop
_ = v_0
var __t14 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var Call_local_Data_Interval_Duration_Iso_go__go_1_0_2 func(*Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Maybe_Just[float64]
_ = Call_local_Data_Interval_Duration_Iso_go__go_1_0_2
var go__go_1_0_2 gopurs_runtime.Value
_ = go__go_1_0_2
Call_local_Data_Interval_Duration_Iso_go__go_1_0_2 = func(v_2_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Maybe_Just[float64] {
go__go_1_0_2:
for {
if false { continue go__go_1_0_2 }
var v_2 *Constructor_Data_Map_Internal_Node[uint32, float64] = v_2_loop
_ = v_2
var __t7 *Constructor_Data_Maybe_Just[float64]
{
if (v_2 == nil) {
__t7 = Rebox_Data_Interval_Duration_Iso_3094389156_3240988860(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
goto end_branch_7
} else {

}
}
{
if (v_2 != nil) {
var __t6 *Constructor_Data_Maybe_Just[float64]
{
var __t_tag_1 uint32 = (v_2).V2
if (uint32(__t_tag_1) == 3908053364) {
v_2_loop = (v_2).V5
continue go__go_1_0_2
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_2 uint32 = (v_2).V2
if (uint32(__t_tag_2) == 217821258) {
v_2_loop = (v_2).V5
continue go__go_1_0_2
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_3 uint32 = (v_2).V2
if (uint32(__t_tag_3) == 1292308612) {
v_2_loop = (v_2).V5
continue go__go_1_0_2
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_4 uint32 = (v_2).V2
if (uint32(__t_tag_4) == 2311060696) {
v_2_loop = (v_2).V5
continue go__go_1_0_2
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
var __t_tag_5 uint32 = (v_2).V2
if (uint32(__t_tag_5) == 401302776) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[float64]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Float((v_2).V3), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_6
} else {

}
}
{
v_2_loop = (v_2).V4
continue go__go_1_0_2
__t6 = func() *Constructor_Data_Maybe_Just[float64] { panic("unreachable") }()
}
end_branch_6:
__t7 = __t6
goto end_branch_7
} else {

}
}
{
__t7 = func() *Constructor_Data_Maybe_Just[float64] { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__go_1_0_2 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_3240988860_3094389156(Call_local_Data_Interval_Duration_Iso_go__go_1_0_2(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node[uint32, float64]](v_2_loop_val))))}
})
// TAST (Let): __local_var_2_8 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [Number])
__local_var_2_8 := Call_local_Data_Interval_Duration_Iso_go__go_1_0_2(v_0.asMap)
_ = __local_var_2_8
var __t9 gopurs_runtime.Value
{
if (__local_var_2_8 == nil) {
__t9 = gopurs_runtime.Bool(false)
goto end_branch_9
} else {

}
}
{
if (__local_var_2_8 != nil) {
__t9 = gopurs_runtime.Bool(true)
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
var __t_and_13 bool = false
if (__t9.IntVal) != (0) {

var __t12 bool
{
var __t_tag_10 *Constructor_Data_Map_Internal_Node[uint32, float64] = v_0.asMap
if (__t_tag_10 == nil) {
__t12 = false
goto end_branch_12
} else {

}
}
{
var __t_tag_11 *Constructor_Data_Map_Internal_Node[uint32, float64] = v_0.asMap
if (__t_tag_11 != nil) {
__t12 = ((v_0.asMap).V1) > (int64(1))
goto end_branch_12
} else {

}
}
{
__t12 = func() bool { panic("Failed pattern match") }()
}
end_branch_12:
__t_and_13 = __t12
}
if __t_and_13 {
__t14 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: unsafe.Pointer(nil)}, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})
goto end_branch_14
} else {

}
}
{
__t14 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
}
end_branch_14:
return __t14
}

func Call_Data_Interval_Duration_Iso_checkNegativeValues(v_0_loop struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
}) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
} = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V0), Get_Data_List_Types_monoidList(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply5(Get_Data_Ord_ordNumberImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1, gopurs_runtime.Float(0.0))
if ((uint32(__t_tag_0.IntVal) == 1527465420)) != (true) {
__t1 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
goto end_branch_1
} else {

}
}
{
__t1 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue{1, uint32((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0.IntVal)}))}, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_2442833393_849153993(v_0.asList))}))
}

func Call_Data_Interval_Duration_Iso_checkFractionalUse(v_0_loop struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
}) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
} = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 shape=Other bindingType=(ADT ["Data","List","Types","List"] [(ADT ["Data","Tuple","Tuple"] [(ADT ["Data","Interval","Duration","DurationComponent"] []), Number])])
v1_1_0 := Rebox_Data_Interval_Duration_Iso_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_List_span(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1.FloatVal())).FloatVal()) == ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1.FloatVal())).FloatVal()) == (gopurs_runtime.Bool(false).FloatVal())) != (true))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_2442833393_849153993(v_0.asList))}), "rest")))
_ = v1_1_0
var __t3 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_and_2 bool = false
if (v1_1_0 != nil) {

// TAST (Let): semigroupAdditive1_2_1 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar a)])
semigroupAdditive1_2_1 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_2.FloatVal()) + (v1_3.FloatVal()))
})})
_ = semigroupAdditive1_2_1
__t_and_2 = (gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value]]](Get_Data_List_Types_foldableList()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_321927638_1201789390((&Constructor_Data_Monoid_Monoid[float64]{1, gopurs_runtime.Func(func(_dollar___unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAdditive1_2_1)}
}), gopurs_runtime.Float(0.0).FloatVal()})))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Number_abs(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1).FloatVal())
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_2442833393_849153993((v1_1_0).V1))}).FloatVal()) > (0.0)
}
if __t_and_2 {
__t3 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer((&Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse{1, ((v1_1_0).V0).V0}))}, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
}
end_branch_3:
return __t3
}

func Call_Data_Interval_Duration_Iso_checkEmptiness(v_0_loop struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
}) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
} = v_0_loop
_ = v_0
var __t1 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
var __t_tag_0 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v_0.asList
if (__t_tag_0 == nil) {
__t1 = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: unsafe.Pointer(nil)}, (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)})
goto end_branch_1
} else {

}
}
{
__t1 = (*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil)
}
end_branch_1:
return __t1
}

func Call_Data_Interval_Duration_Iso_checkValidIsoDuration(v_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
var v_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = v_0_loop
_ = v_0
// TAST (Let): semigroupFn_1_0 shape=LitRecord bindingType=(ADT ["Data","Semigroup","Semigroup"] [(Func [(TypeVar a)] (TypeVar b))])
semigroupFn_1_0 := (&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
var Call_local_Data_Interval_Duration_Iso_go__go_4_1_3 func(gopurs_runtime.Value, *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value
_ = Call_local_Data_Interval_Duration_Iso_go__go_4_1_3
var go__go_4_1_3 gopurs_runtime.Value
_ = go__go_4_1_3
Call_local_Data_Interval_Duration_Iso_go__go_4_1_3 = func(b_5_loop gopurs_runtime.Value, v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
go__go_4_1_3:
for {
if false { continue go__go_4_1_3 }
var b_5 gopurs_runtime.Value = b_5_loop
_ = b_5
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var __t2 gopurs_runtime.Value
{
if (v_6 == nil) {
__t2 = b_5
goto end_branch_2
} else {

}
}
{
if (v_6 != nil) {
b_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v_6).V0, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](b_5)}))}
v_6_loop = (v_6).V1
continue go__go_4_1_3
__t2 = func() gopurs_runtime.Value { panic("unreachable") }()
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
}
go__go_4_1_3 = gopurs_runtime.Func(func(b_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return Call_local_Data_Interval_Duration_Iso_go__go_4_1_3(b_5_loop_val, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val))
})
})
var Call_local_Data_Interval_Duration_Iso_go__go_5_3_4 func(*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
_ = Call_local_Data_Interval_Duration_Iso_go__go_5_3_4
var go__go_5_3_4 gopurs_runtime.Value
_ = go__go_5_3_4
Call_local_Data_Interval_Duration_Iso_go__go_5_3_4 = func(v_6_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value], v1_7_loop *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
go__go_5_3_4:
for {
if false { continue go__go_5_3_4 }
var v_6 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v_6_loop
_ = v_6
var v1_7 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] = v1_7_loop
_ = v1_7
var __t4 *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]
{
if (v1_7 == nil) {
__t4 = v_6
goto end_branch_4
} else {

}
}
{
if (v1_7 != nil) {
v_6_loop = (&Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{1, (v1_7).V0, v_6})
v1_7_loop = (v1_7).V1
continue go__go_5_3_4
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("unreachable") }()
goto end_branch_4
} else {

}
}
{
__t4 = func() *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}
go__go_5_3_4 = gopurs_runtime.Func(func(v_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_local_Data_Interval_Duration_Iso_go__go_5_3_4(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_6_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_7_loop_val)))}
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](Call_local_Data_Interval_Duration_Iso_go__go_4_1_3(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(g_2, x_3)))}, Call_local_Data_Interval_Duration_Iso_go__go_5_3_4((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, x_3))))))}
})})
_ = semigroupFn_1_0
var Call_local_Data_Interval_Duration_Iso_go__952457181_2_5_5 func(*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
_ = Call_local_Data_Interval_Duration_Iso_go__952457181_2_5_5
var go__952457181_2_5_5 gopurs_runtime.Value
_ = go__952457181_2_5_5
Call_local_Data_Interval_Duration_Iso_go__952457181_2_5_5 = func(v_3_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]], v1_4_loop *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
go__952457181_2_5_5:
for {
if false { continue go__952457181_2_5_5 }
var v_3 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v_3_loop
_ = v_3
var v1_4 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = v1_4_loop
_ = v1_4
var __t6 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
{
if (v1_4 == nil) {
__t6 = v_3
goto end_branch_6
} else {

}
}
{
if (v1_4 != nil) {
v_3_loop = (&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{1, (v1_4).V0, v_3})
v1_4_loop = (v1_4).V1
continue go__952457181_2_5_5
__t6 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] { panic("unreachable") }()
goto end_branch_6
} else {

}
}
{
__t6 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}
}
go__952457181_2_5_5 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_2442833393_849153993(Call_local_Data_Interval_Duration_Iso_go__952457181_2_5_5(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](v_3_loop_val), gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]](v1_4_loop_val))))}
})
})
var go__go_3_7_6 gopurs_runtime.Value
_ = go__go_3_7_6
// FALLBACK TCO: isLoop=false len=1
go__go_3_7_6 = gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t10 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
{
var __t_tag_8 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = Rebox_Data_Interval_Duration_Iso_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5))
if (__t_tag_8 == nil) {
__t10 = Rebox_Data_Interval_Duration_Iso_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4))
goto end_branch_10
} else {

}
}
{
var __t_tag_9 *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] = Rebox_Data_Interval_Duration_Iso_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v1_5))
if (__t_tag_9 != nil) {
__t10 = Call_local_Data_Interval_Duration_Iso_go__952457181_2_5_5((&Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{1, Rebox_Data_Interval_Duration_Iso_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0)), Rebox_Data_Interval_Duration_Iso_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](v_4))}), Rebox_Data_Interval_Duration_Iso_849153993_2442833393((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1))
goto end_branch_10
} else {

}
}
{
__t10 = func() *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] { panic("Failed pattern match") }()
}
end_branch_10:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_2442833393_849153993(__t10))}
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply4(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupFn_1_0)}
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons[gopurs_runtime.Value])(nil))}))}
})}))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Interval_Duration_Iso_checkWeekUsage(), Get_Data_Interval_Duration_Iso_checkEmptiness(), Get_Data_Interval_Duration_Iso_checkFractionalUse(), Get_Data_Interval_Duration_Iso_checkNegativeValues()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), func() gopurs_runtime.Value {
				orig := struct{
	asList *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]
	asMap *Constructor_Data_Map_Internal_Node[uint32, float64]
}{Call_local_Data_Interval_Duration_Iso_go__952457181_2_5_5((*Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]])(nil), Rebox_Data_Interval_Duration_Iso_849153993_2442833393(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Map_Internal_toUnfoldable__994050345(), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(v_0))})))), v_0}
				_ = orig
				return gopurs_runtime.RecordDict([]string{"asList", "asMap"}, []gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_2442833393_849153993(orig.asList))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(orig.asMap))}})
				}()))
}

func Call_Data_Interval_Duration_Iso_mkIsoDuration(d_0_loop *Constructor_Data_Map_Internal_Node[uint32, float64]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
var d_0 *Constructor_Data_Map_Internal_Node[uint32, float64] = d_0_loop
_ = d_0
// TAST (Let): __local_var_1_1 shape=App(Var) bindingType=Any
__local_var_1_1 := Call_Data_Interval_Duration_Iso_checkValidIsoDuration(d_0)
_ = __local_var_1_1
var __t2 *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]
{
if (__local_var_1_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
if (__local_var_1_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer((&Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (__local_var_1_1).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((__local_var_1_1).V1)}}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[*Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] { panic("Failed pattern match") }()
}
end_branch_2:
// TAST (Let): v_1_0 shape=Let(Branch(Other, Other, def=Other)) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","NonEmpty","NonEmpty"] [(ADT ["Data","List","Types","List"] []), (ADT ["Data","Interval","Duration","Iso","Error"] [])])])
v_1_0 := __t2
_ = v_1_0
var __t3 gopurs_runtime.Value
{
if (v_1_0 != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_3123684004_1293498952((v_1_0).V0))}}))}
goto end_branch_3
} else {

}
}
{
if (v_1_0 == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer((&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_261879545_2487766124(d_0))}}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool} {
				_v := __t3
				if _v.Type == 9 && _v.IntVal == 3234899973 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: gopurs_runtime.Value{}, V1: (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V2: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value; V2 bool}{V0: (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: gopurs_runtime.Value{}, V2: false}
			}()
}

func Rebox_Data_Interval_Duration_Iso_138441832_3132786365(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[uint32, float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[uint32, float64]{}
		out.V0 = uint32(in.V0.IntVal)
		out.V1 = in.V1.FloatVal()
	return out
}

func Rebox_Data_Interval_Duration_Iso_2442833393_849153993(in *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]) *Constructor_Data_List_Types_Cons[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Interval_Duration_Iso_3132786365_138441832(in.V0))}
		out.V1 = Rebox_Data_Interval_Duration_Iso_2442833393_849153993(in.V1)
	return out
}

func Rebox_Data_Interval_Duration_Iso_2487766124_261879545(in *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Map_Internal_Node[uint32, float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Map_Internal_Node[uint32, float64]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = uint32(in.V2.IntVal)
		out.V3 = in.V3.FloatVal()
		out.V4 = Rebox_Data_Interval_Duration_Iso_2487766124_261879545(in.V4)
		out.V5 = Rebox_Data_Interval_Duration_Iso_2487766124_261879545(in.V5)
	return out
}

func Rebox_Data_Interval_Duration_Iso_261879545_2487766124(in *Constructor_Data_Map_Internal_Node[uint32, float64]) *Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Map_Internal_Node[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V2), UnsafePtr: nil}
		out.V3 = gopurs_runtime.Float(in.V3)
		out.V4 = Rebox_Data_Interval_Duration_Iso_261879545_2487766124(in.V4)
		out.V5 = Rebox_Data_Interval_Duration_Iso_261879545_2487766124(in.V5)
	return out
}

func Rebox_Data_Interval_Duration_Iso_2958538738_3790796878(in *Constructor_Data_Eq_Eq[*Constructor_Data_Map_Internal_Node[uint32, float64]]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Interval_Duration_Iso_3094389156_3240988860(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[float64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[float64]{}
		out.V0 = in.V0.FloatVal()
	return out
}

func Rebox_Data_Interval_Duration_Iso_3123684004_1293498952(in *Constructor_Data_NonEmpty_NonEmpty[*Constructor_Data_List_Types_Cons[gopurs_runtime.Value], gopurs_runtime.Value]) *Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_NonEmpty_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Interval_Duration_Iso_3132786365_138441832(in *Constructor_Data_Tuple_Tuple[uint32, float64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: int64(in.V0), UnsafePtr: nil}
		out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Interval_Duration_Iso_321927638_1201789390(in *Constructor_Data_Monoid_Monoid[float64]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Float(in.V1)
	return out
}

func Rebox_Data_Interval_Duration_Iso_3240988860_3094389156(in *Constructor_Data_Maybe_Just[float64]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Float(in.V0)
	return out
}

func Rebox_Data_Interval_Duration_Iso_3498527378_4177771502(in *Constructor_Data_Ord_Ord[*Constructor_Data_Map_Internal_Node[uint32, float64]]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Interval_Duration_Iso_849153993_2442833393(in *Constructor_Data_List_Types_Cons[gopurs_runtime.Value]) *Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]] {
	if in == nil { return nil }
	out := &Constructor_Data_List_Types_Cons[*Constructor_Data_Tuple_Tuple[uint32, float64]]{}
		out.V0 = Rebox_Data_Interval_Duration_Iso_138441832_3132786365(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
		out.V1 = Rebox_Data_Interval_Duration_Iso_849153993_2442833393(in.V1)
	return out
}

func Rebox_Data_Interval_Duration_Iso_983382226_1386611502(in *Constructor_Data_Show_Show[*Constructor_Data_Map_Internal_Node[uint32, float64]]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}



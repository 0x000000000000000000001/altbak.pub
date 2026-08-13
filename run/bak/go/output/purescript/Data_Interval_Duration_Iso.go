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
		cache_Data_Interval_Duration_Iso_empty = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty")))}
	})
	return cache_Data_Interval_Duration_Iso_empty
}

var cache_Data_Interval_Duration_Iso_foldMap gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_foldMap sync.Once
func Get_Data_Interval_Duration_Iso_foldMap() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_foldMap.Do(func() {
		cache_Data_Interval_Duration_Iso_foldMap = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), Get_Data_List_Types_monoidList())
	})
	return cache_Data_Interval_Duration_Iso_foldMap
}

var cache_Data_Interval_Duration_Iso_monoidAdditive gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_monoidAdditive sync.Once
func Get_Data_Interval_Duration_Iso_monoidAdditive() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_monoidAdditive.Do(func() {
		cache_Data_Interval_Duration_Iso_monoidAdditive = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", Get_Data_Semiring_numAdd(), Get_Data_Semiring_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
_ = __local_var_0_0
// TAST (Let): semigroupAdditive1_1_1 -> gopurs_runtime.Value
semigroupAdditive1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "add"), v_1, v1_2)
})
}))
_ = semigroupAdditive1_1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_1_1
}), gopurs_runtime.RecordGet(__local_var_0_0, "zero"))))}
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply(f_0, a_2), gopurs_runtime.Apply(g_1, a_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(f_0, a_2), gopurs_runtime.Apply(g_1, a_2))
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "ff")
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) != (true)) || ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(f_0, a_1))
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "tt")
})})}
	})
	return cache_Data_Interval_Duration_Iso_heytingAlgebraFunction
}

var cache_Data_Interval_Duration_Iso_monoidFn gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_monoidFn sync.Once
func Get_Data_Interval_Duration_Iso_monoidFn() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_monoidFn.Do(func() {
		cache_Data_Interval_Duration_Iso_monoidFn = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_1 -> gopurs_runtime.Value
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_monoidList(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_0_1
// TAST (Let): semigroupFn_0_0 -> gopurs_runtime.Value
semigroupFn_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})
})
}))
_ = semigroupFn_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_0_0
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Get_Data_List_Types_monoidList(), "mempty")
}))))}
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
		cache_Data_Interval_Duration_Iso_showIsoDuration = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(IsoDuration ") + (gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Show_Show](Get_Data_Interval_Duration_showMap()).V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](v_0))}).StrVal())) + (")"))
}))
	})
	return cache_Data_Interval_Duration_Iso_showIsoDuration
}

var cache_Data_Interval_Duration_Iso_showError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_showError sync.Once
func Get_Data_Interval_Duration_Iso_showError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_showError.Do(func() {
		cache_Data_Interval_Duration_Iso_showError = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
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
		cache_Data_Interval_Duration_Iso_eqIsoDuration = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](Get_Data_Interval_Duration_eqMap()).V0), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1))}).IntVal) != (0))
})
}))
	})
	return cache_Data_Interval_Duration_Iso_eqIsoDuration
}

var cache_Data_Interval_Duration_Iso_ordIsoDuration gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_ordIsoDuration sync.Once
func Get_Data_Interval_Duration_Iso_ordIsoDuration() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_ordIsoDuration.Do(func() {
		cache_Data_Interval_Duration_Iso_ordIsoDuration = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_Duration_Iso_eqIsoDuration()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](Get_Data_Interval_Duration_ordMap()).V1), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](x_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](y_1))}).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_Data_Interval_Duration_Iso_ordIsoDuration
}

var cache_Data_Interval_Duration_Iso_eqError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_eqError sync.Once
func Get_Data_Interval_Duration_Iso_eqError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_eqError.Do(func() {
		cache_Data_Interval_Duration_Iso_eqError = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
	})
	return cache_Data_Interval_Duration_Iso_eqError
}

var cache_Data_Interval_Duration_Iso_ordError gopurs_runtime.Value
var once_Data_Interval_Duration_Iso_ordError sync.Once
func Get_Data_Interval_Duration_Iso_ordError() gopurs_runtime.Value {
	once_Data_Interval_Duration_Iso_ordError.Do(func() {
		cache_Data_Interval_Duration_Iso_ordError = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Interval_Duration_Iso_eqError()
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
}))
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
__t7 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty"))
}
end_branch_7:
return __t7
}

func Call_Data_Interval_Duration_Iso_checkNegativeValues(v_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_List_Types_foldableList(), "foldMap"), Get_Data_List_Types_monoidList(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_List_Types_Cons
{
var __t0 bool
{
if ((*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V1.FloatVal()) < (0.0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty"))
goto end_branch_1
} else {

}
}
{
__t1 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_Duration_Iso_ContainsNegativeValue{1, uint32((*Constructor_Data_Tuple_Tuple)(v1_1.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_List_Types_Cons)(nil)}
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t1)}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_0, "asList")))}))
}

func Call_Data_Interval_Duration_Iso_checkFractionalUse(v_0_loop gopurs_runtime.Value) *Constructor_Data_List_Types_Cons {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
// TAST (Let): v1_1_0 -> *Constructor_Data_List_Types_Cons
v1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_Data_List_span__2133741451(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not"), gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Eq_eqNumber(), "eq"), gopurs_runtime.Float(gopurs_runtime.Apply(Get_Data_Number_floor(), gopurs_runtime.Float((*Constructor_Data_Tuple_Tuple)(x_1.UnsafePtr).V1.FloatVal())).FloatVal()), gopurs_runtime.Float((*Constructor_Data_Tuple_Tuple)(x_1.UnsafePtr).V1.FloatVal())).IntVal) != (0)) != (true)))
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(v_0, "asList")))}), "rest"))
_ = v1_1_0
var __t3 *Constructor_Data_List_Types_Cons
{
var __t_and_2 bool = false
if (v1_1_0 != nil) {

var __t1 bool
{
if (gopurs_runtime.Apply2(Get_Data_Foldable_foldMap__193737345(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Number_abs(), (*Constructor_Data_Tuple_Tuple)(x_2.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((v1_1_0).V1)}).FloatVal()) > (0.0) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t_and_2 = __t1
}
if __t_and_2 {
__t3 = &Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Constructor_Data_Interval_Duration_Iso_InvalidFractionalUse{1, uint32((*Constructor_Data_Tuple_Tuple)((v1_1_0).V0.UnsafePtr).V0.IntVal)})}, (*Constructor_Data_List_Types_Cons)(nil)}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty"))
}
end_branch_3:
return __t3
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
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.RecordGet(Get_Data_List_Types_plusList(), "empty"))
}
end_branch_2:
return __t2
}

func Call_Data_Interval_Duration_Iso_checkValidIsoDuration(v_0_loop *Constructor_Data_Map_Internal_Node) *Constructor_Data_List_Types_Cons {
var v_0 *Constructor_Data_Map_Internal_Node = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_List_Types_monoidList(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): semigroupFn_1_0 -> gopurs_runtime.Value
semigroupFn_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "append"), gopurs_runtime.Apply(f_2, x_4), gopurs_runtime.Apply(g_3, x_4))
})
})
}))
_ = semigroupFn_1_0
var go__go_1_2_0 gopurs_runtime.Value
go__go_1_2_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *Constructor_Data_List_Types_Cons = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_2_0:
for {
if false { continue go__go_1_2_0 }
var v_2 *Constructor_Data_List_Types_Cons = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t3 *Constructor_Data_List_Types_Cons
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t3 = v_2
goto end_branch_3
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&Constructor_Data_List_Types_Cons{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V0))}, gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(v1_3.UnsafePtr).V1)}
continue go__go_1_2_0
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Value{})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t3)}
}
}()
})
})
return gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply4(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Get_Data_List_Types_monoidList(), "mempty")
}))))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_Data_Interval_Duration_Iso_checkWeekUsage(), Get_Data_Interval_Duration_Iso_checkEmptiness(), Get_Data_Interval_Duration_Iso_checkFractionalUse(), Get_Data_Interval_Duration_Iso_checkNegativeValues()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.RecordDict2("asList", "asMap", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(go__go_1_2_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*Constructor_Data_List_Types_Cons)(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_List_Types_Cons](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_List_Types_unfoldableList(), "unfoldr"), Get_Data_Map_Internal_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&Constructor_Data_Map_Internal_IterNode{1, gopurs_runtime.CoerceToStruct[Constructor_Data_Map_Internal_Node](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})))})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
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



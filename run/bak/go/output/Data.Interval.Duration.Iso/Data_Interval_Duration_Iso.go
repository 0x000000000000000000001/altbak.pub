package Data_Interval_Duration_Iso

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Interval_Duration "gopurs/output/Data.Interval.Duration"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_foldMap gopurs_runtime.Value
var once_foldMap sync.Once
func Get_foldMap() gopurs_runtime.Value {
	once_foldMap.Do(func() {
		cache_foldMap = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), pkg_Data_List_Types.Get_monoidList())
	})
	return cache_foldMap
}

var cache_monoidAdditive gopurs_runtime.Value
var once_monoidAdditive sync.Once
func Get_monoidAdditive() gopurs_runtime.Value {
	once_monoidAdditive.Do(func() {
		cache_monoidAdditive = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.RecordDict4("add", "mul", "one", "zero", pkg_Data_Semiring.Get_numAdd(), pkg_Data_Semiring.Get_numMul(), gopurs_runtime.Float(1.0), gopurs_runtime.Float(0.0))
_ = __local_var_0_0
semigroupAdditive1_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "add"), v_1, v1_2)
})
}))
_ = semigroupAdditive1_1_1
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[float64]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_1_1
}), gopurs_runtime.RecordGet(__local_var_0_0, "zero"))))}
}()
	})
	return cache_monoidAdditive
}

var cache_heytingAlgebraFunction gopurs_runtime.Value
var once_heytingAlgebraFunction sync.Once
func Get_heytingAlgebraFunction() gopurs_runtime.Value {
	once_heytingAlgebraFunction.Do(func() {
		cache_heytingAlgebraFunction = gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(&pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply(f_0, a_2), gopurs_runtime.Apply(g_1, a_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(f_0, a_2), gopurs_runtime.Apply(g_1, a_2))
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "ff")
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((((gopurs_runtime.Apply(f_0, a_2).IntVal) != (0)) != (true)) || ((gopurs_runtime.Apply(g_1, a_2).IntVal) != (0)))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(f_0, a_1))
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "tt")
})})}
	})
	return cache_heytingAlgebraFunction
}

var cache_monoidFn gopurs_runtime.Value
var once_monoidFn sync.Once
func Get_monoidFn() gopurs_runtime.Value {
	once_monoidFn.Do(func() {
		cache_monoidFn = func() gopurs_runtime.Value {
__local_var_0_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_monoidList(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_0_1
semigroupFn_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_1, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
})
})
}))
_ = semigroupFn_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_0_0
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_monoidList(), "mempty")
}))))}
}()
	})
	return cache_monoidFn
}

var cache_IsEmpty gopurs_runtime.Value
var once_IsEmpty sync.Once
func Get_IsEmpty() gopurs_runtime.Value {
	once_IsEmpty.Do(func() {
		cache_IsEmpty = gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_IsEmpty
}

var cache_InvalidWeekComponentUsage gopurs_runtime.Value
var once_InvalidWeekComponentUsage sync.Once
func Get_InvalidWeekComponentUsage() gopurs_runtime.Value {
	once_InvalidWeekComponentUsage.Do(func() {
		cache_InvalidWeekComponentUsage = gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: unsafe.Pointer(nil)}
	})
	return cache_InvalidWeekComponentUsage
}

var cache_ContainsNegativeValue gopurs_runtime.Value
var once_ContainsNegativeValue sync.Once
func Get_ContainsNegativeValue() gopurs_runtime.Value {
	once_ContainsNegativeValue.Do(func() {
		cache_ContainsNegativeValue = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Constructor_ContainsNegativeValue{1, uint32(value0.IntVal)})}
})
	})
	return cache_ContainsNegativeValue
}

var cache_InvalidFractionalUse gopurs_runtime.Value
var once_InvalidFractionalUse sync.Once
func Get_InvalidFractionalUse() gopurs_runtime.Value {
	once_InvalidFractionalUse.Do(func() {
		cache_InvalidFractionalUse = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Constructor_InvalidFractionalUse{1, uint32(value0.IntVal)})}
})
	})
	return cache_InvalidFractionalUse
}

var cache_unIsoDuration gopurs_runtime.Value
var once_unIsoDuration sync.Once
func Get_unIsoDuration() gopurs_runtime.Value {
	once_unIsoDuration.Do(func() {
		cache_unIsoDuration = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(Call_unIsoDuration(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](v_0_box)))}
})
	})
	return cache_unIsoDuration
}

var cache_showIsoDuration gopurs_runtime.Value
var once_showIsoDuration sync.Once
func Get_showIsoDuration() gopurs_runtime.Value {
	once_showIsoDuration.Do(func() {
		cache_showIsoDuration = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(IsoDuration "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_show__2896747026(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](v_0))}).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
	})
	return cache_showIsoDuration
}

var cache_showError gopurs_runtime.Value
var once_showError sync.Once
func Get_showError() gopurs_runtime.Value {
	once_showError.Do(func() {
		cache_showError = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1422140417) {
__t0 = gopurs_runtime.Str("(IsEmpty)")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1775501833) {
__t0 = gopurs_runtime.Str("(InvalidWeekComponentUsage)")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3224543173) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(ContainsNegativeValue "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Str(Call_show__1261750354((*Constructor_ContainsNegativeValue)(v_0.UnsafePtr).V0)).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(InvalidFractionalUse "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Str(Call_show__1261750354((*Constructor_InvalidFractionalUse)(v_0.UnsafePtr).V0)).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Str(__t0.StrVal())
}))
	})
	return cache_showError
}

var cache_prettyError gopurs_runtime.Value
var once_prettyError sync.Once
func Get_prettyError() gopurs_runtime.Value {
	once_prettyError.Do(func() {
		cache_prettyError = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_prettyError(v_0_box))
})
	})
	return cache_prettyError
}

var cache_eqIsoDuration gopurs_runtime.Value
var once_eqIsoDuration sync.Once
func Get_eqIsoDuration() gopurs_runtime.Value {
	once_eqIsoDuration.Do(func() {
		cache_eqIsoDuration = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_eq__2224314568(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](x_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](y_1))}).IntVal) != (0))
})
}))
	})
	return cache_eqIsoDuration
}

var cache_ordIsoDuration gopurs_runtime.Value
var once_ordIsoDuration sync.Once
func Get_ordIsoDuration() gopurs_runtime.Value {
	once_ordIsoDuration.Do(func() {
		cache_ordIsoDuration = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqIsoDuration()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(Call_compare__231252914(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](x_0))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](y_1))}).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordIsoDuration
}

var cache_eqError gopurs_runtime.Value
var once_eqError sync.Once
func Get_eqError() gopurs_runtime.Value {
	once_eqError.Do(func() {
		cache_eqError = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 bool
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
__t3 = __t0
goto end_branch_3
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
__t3 = __t1
goto end_branch_3
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t2 bool
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t2 = (gopurs_runtime.Bool(Call_eq__1241439021((*Constructor_ContainsNegativeValue)(x_0.UnsafePtr).V0, (*Constructor_ContainsNegativeValue)(y_1.UnsafePtr).V0)).IntVal) != (0)
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t3 = __t2
goto end_branch_3
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 574232667)) && ((y_1.Type == 9 && y_1.IntVal == 574232667)) {
__t3 = (gopurs_runtime.Bool(Call_eq__1241439021((*Constructor_InvalidFractionalUse)(x_0.UnsafePtr).V0, (*Constructor_InvalidFractionalUse)(y_1.UnsafePtr).V0)).IntVal) != (0)
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
return gopurs_runtime.Bool(__t3)
})
}))
	})
	return cache_eqError
}

var cache_ordError gopurs_runtime.Value
var once_ordError sync.Once
func Get_ordError() gopurs_runtime.Value {
	once_ordError.Do(func() {
		cache_ordError = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqError()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_3
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
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t2 uint32
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t2 = uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(Call_compare__3077449111((*Constructor_ContainsNegativeValue)(x_0.UnsafePtr).V0, (*Constructor_ContainsNegativeValue)(y_1.UnsafePtr).V0)), UnsafePtr: nil}.IntVal)
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t2), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 574232667)) && ((y_1.Type == 9 && y_1.IntVal == 574232667)) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(Call_compare__3077449111((*Constructor_InvalidFractionalUse)(x_0.UnsafePtr).V0, (*Constructor_InvalidFractionalUse)(y_1.UnsafePtr).V0)), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t3.IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_ordError
}

var cache_checkWeekUsage gopurs_runtime.Value
var once_checkWeekUsage sync.Once
func Get_checkWeekUsage() gopurs_runtime.Value {
	once_checkWeekUsage.Do(func() {
		cache_checkWeekUsage = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_checkWeekUsage(v_0_box))}
})
	})
	return cache_checkWeekUsage
}

var cache_checkNegativeValues gopurs_runtime.Value
var once_checkNegativeValues sync.Once
func Get_checkNegativeValues() gopurs_runtime.Value {
	once_checkNegativeValues.Do(func() {
		cache_checkNegativeValues = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_checkNegativeValues(v_0_box))}
})
	})
	return cache_checkNegativeValues
}

var cache_checkFractionalUse gopurs_runtime.Value
var once_checkFractionalUse sync.Once
func Get_checkFractionalUse() gopurs_runtime.Value {
	once_checkFractionalUse.Do(func() {
		cache_checkFractionalUse = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_checkFractionalUse(v_0_box))}
})
	})
	return cache_checkFractionalUse
}

var cache_checkEmptiness gopurs_runtime.Value
var once_checkEmptiness sync.Once
func Get_checkEmptiness() gopurs_runtime.Value {
	once_checkEmptiness.Do(func() {
		cache_checkEmptiness = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_checkEmptiness(v_0_box))}
})
	})
	return cache_checkEmptiness
}

var cache_checkValidIsoDuration gopurs_runtime.Value
var once_checkValidIsoDuration sync.Once
func Get_checkValidIsoDuration() gopurs_runtime.Value {
	once_checkValidIsoDuration.Do(func() {
		cache_checkValidIsoDuration = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_checkValidIsoDuration(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](v_0_box)))}
})
	})
	return cache_checkValidIsoDuration
}

var cache_mkIsoDuration gopurs_runtime.Value
var once_mkIsoDuration sync.Once
func Get_mkIsoDuration() gopurs_runtime.Value {
	once_mkIsoDuration.Do(func() {
		cache_mkIsoDuration = gopurs_runtime.Func(func(d_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mkIsoDuration(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](d_0_box))
})
	})
	return cache_mkIsoDuration
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_pure__2154981942 gopurs_runtime.Value
var once_pure__2154981942 sync.Once
func Get_pure__2154981942() gopurs_runtime.Value {
	once_pure__2154981942.Do(func() {
		cache_pure__2154981942 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_pure__2154981942(a_0_box))}
})
	})
	return cache_pure__2154981942
}

var cache_apply__1030762512 gopurs_runtime.Value
var once_apply__1030762512 sync.Once
func Get_apply__1030762512() gopurs_runtime.Value {
	once_apply__1030762512.Do(func() {
		cache_apply__1030762512 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__1030762512(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_apply__1030762512
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__2169384906 gopurs_runtime.Value
var once_apply__2169384906 sync.Once
func Get_apply__2169384906() gopurs_runtime.Value {
	once_apply__2169384906.Do(func() {
		cache_apply__2169384906 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_apply__2169384906(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_apply__2169384906
}

var cache_empty__932402776 gopurs_runtime.Value
var once_empty__932402776 sync.Once
func Get_empty__932402776() gopurs_runtime.Value {
	once_empty__932402776.Do(func() {
		cache_empty__932402776 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_empty__932402776(dict_0_box)
})
	})
	return cache_empty__932402776
}

var cache_compose__1987728071 gopurs_runtime.Value
var once_compose__1987728071 sync.Once
func Get_compose__1987728071() gopurs_runtime.Value {
	once_compose__1987728071.Do(func() {
		cache_compose__1987728071 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1987728071(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compose__1987728071
}

var cache_compose__3249109794 gopurs_runtime.Value
var once_compose__3249109794 sync.Once
func Get_compose__3249109794() gopurs_runtime.Value {
	once_compose__3249109794.Do(func() {
		cache_compose__3249109794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__3249109794(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_compose__3249109794
}

var cache_compose__1555187646 gopurs_runtime.Value
var once_compose__1555187646 sync.Once
func Get_compose__1555187646() gopurs_runtime.Value {
	once_compose__1555187646.Do(func() {
		cache_compose__1555187646 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__1555187646(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__1555187646
}

var cache_compose__2527254334 gopurs_runtime.Value
var once_compose__2527254334 sync.Once
func Get_compose__2527254334() gopurs_runtime.Value {
	once_compose__2527254334.Do(func() {
		cache_compose__2527254334 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__2527254334(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__2527254334
}

var cache_compose__2995688990 gopurs_runtime.Value
var once_compose__2995688990 sync.Once
func Get_compose__2995688990() gopurs_runtime.Value {
	once_compose__2995688990.Do(func() {
		cache_compose__2995688990 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__2995688990(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__2995688990
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_composeFlipped__2583068543 gopurs_runtime.Value
var once_composeFlipped__2583068543 sync.Once
func Get_composeFlipped__2583068543() gopurs_runtime.Value {
	once_composeFlipped__2583068543.Do(func() {
		cache_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__2583068543
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_eq__789642299 gopurs_runtime.Value
var once_eq__789642299 sync.Once
func Get_eq__789642299() gopurs_runtime.Value {
	once_eq__789642299.Do(func() {
		cache_eq__789642299 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__789642299(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[bool]](dict_0_box))
})
	})
	return cache_eq__789642299
}

var cache_eq__1697837627 gopurs_runtime.Value
var once_eq__1697837627 sync.Once
func Get_eq__1697837627() gopurs_runtime.Value {
	once_eq__1697837627.Do(func() {
		cache_eq__1697837627 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__1697837627(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__1697837627
}

var cache_eq__2276491096 gopurs_runtime.Value
var once_eq__2276491096 sync.Once
func Get_eq__2276491096() gopurs_runtime.Value {
	once_eq__2276491096.Do(func() {
		cache_eq__2276491096 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2276491096(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2276491096
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq__1241439021 gopurs_runtime.Value
var once_eq__1241439021 sync.Once
func Get_eq__1241439021() gopurs_runtime.Value {
	once_eq__1241439021.Do(func() {
		cache_eq__1241439021 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_eq__1241439021(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal)))
})
	})
	return cache_eq__1241439021
}

var cache_eq__2224314568 gopurs_runtime.Value
var once_eq__2224314568 sync.Once
func Get_eq__2224314568() gopurs_runtime.Value {
	once_eq__2224314568.Do(func() {
		cache_eq__2224314568 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2224314568(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__2224314568
}

var cache_notEq__2334967935 gopurs_runtime.Value
var once_notEq__2334967935 sync.Once
func Get_notEq__2334967935() gopurs_runtime.Value {
	once_notEq__2334967935.Do(func() {
		cache_notEq__2334967935 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__2334967935(x_0_box, y_1_box))
})
	})
	return cache_notEq__2334967935
}

var cache_notEq__2384498378 gopurs_runtime.Value
var once_notEq__2384498378 sync.Once
func Get_notEq__2384498378() gopurs_runtime.Value {
	once_notEq__2384498378.Do(func() {
		cache_notEq__2384498378 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, y_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_notEq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box, y_2_box))
})
	})
	return cache_notEq__2384498378
}

var cache_fold__910331789 gopurs_runtime.Value
var once_fold__910331789 sync.Once
func Get_fold__910331789() gopurs_runtime.Value {
	once_fold__910331789.Do(func() {
		cache_fold__910331789 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold__910331789(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_fold__910331789
}

var cache_foldMap__4098395794 gopurs_runtime.Value
var once_foldMap__4098395794 sync.Once
func Get_foldMap__4098395794() gopurs_runtime.Value {
	once_foldMap__4098395794.Do(func() {
		cache_foldMap__4098395794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__4098395794(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__4098395794
}

var cache_foldMap__193737345 gopurs_runtime.Value
var once_foldMap__193737345 sync.Once
func Get_foldMap__193737345() gopurs_runtime.Value {
	once_foldMap__193737345.Do(func() {
		cache_foldMap__193737345 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__193737345(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_foldMap__193737345
}

var cache_foldableArray__2950015754 gopurs_runtime.Value
var once_foldableArray__2950015754 sync.Once
func Get_foldableArray__2950015754() gopurs_runtime.Value {
	once_foldableArray__2950015754.Do(func() {
		cache_foldableArray__2950015754 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), pkg_Data_Foldable.Get_foldlArray(), pkg_Data_Foldable.Get_foldrArray())
	})
	return cache_foldableArray__2950015754
}

var cache_foldl__2699291984 gopurs_runtime.Value
var once_foldl__2699291984 sync.Once
func Get_foldl__2699291984() gopurs_runtime.Value {
	once_foldl__2699291984.Do(func() {
		cache_foldl__2699291984 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2699291984(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__2699291984
}

var cache_foldl__2602334544 gopurs_runtime.Value
var once_foldl__2602334544 sync.Once
func Get_foldl__2602334544() gopurs_runtime.Value {
	once_foldl__2602334544.Do(func() {
		cache_foldl__2602334544 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2602334544(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__2602334544
}

var cache_foldl__1656262032 gopurs_runtime.Value
var once_foldl__1656262032 sync.Once
func Get_foldl__1656262032() gopurs_runtime.Value {
	once_foldl__1656262032.Do(func() {
		cache_foldl__1656262032 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__1656262032(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldl__1656262032
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldl__3943124669 gopurs_runtime.Value
var once_foldl__3943124669 sync.Once
func Get_foldl__3943124669() gopurs_runtime.Value {
	once_foldl__3943124669.Do(func() {
		cache_foldl__3943124669 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3943124669(f_0_box)
})
	})
	return cache_foldl__3943124669
}

var cache_foldl__396932925 gopurs_runtime.Value
var once_foldl__396932925 sync.Once
func Get_foldl__396932925() gopurs_runtime.Value {
	once_foldl__396932925.Do(func() {
		cache_foldl__396932925 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__396932925(f_0_box)
})
	})
	return cache_foldl__396932925
}

var cache_foldr__2512763050 gopurs_runtime.Value
var once_foldr__2512763050 sync.Once
func Get_foldr__2512763050() gopurs_runtime.Value {
	once_foldr__2512763050.Do(func() {
		cache_foldr__2512763050 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2512763050(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2512763050
}

var cache_foldr__3673994608 gopurs_runtime.Value
var once_foldr__3673994608 sync.Once
func Get_foldr__3673994608() gopurs_runtime.Value {
	once_foldr__3673994608.Do(func() {
		cache_foldr__3673994608 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__3673994608(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_foldr__3673994608
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldr__2979608669 gopurs_runtime.Value
var once_foldr__2979608669 sync.Once
func Get_foldr__2979608669() gopurs_runtime.Value {
	once_foldr__2979608669.Do(func() {
		cache_foldr__2979608669 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2979608669(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_1_box))
})
	})
	return cache_foldr__2979608669
}

var cache_const__220790420 gopurs_runtime.Value
var once_const__220790420 sync.Once
func Get_const__220790420() gopurs_runtime.Value {
	once_const__220790420.Do(func() {
		cache_const__220790420 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__220790420(a_0_box, v_1_box)
})
	})
	return cache_const__220790420
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__2974723072 gopurs_runtime.Value
var once_flip__2974723072 sync.Once
func Get_flip__2974723072() gopurs_runtime.Value {
	once_flip__2974723072.Do(func() {
		cache_flip__2974723072 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2974723072(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2974723072
}

var cache_flip__3709724320 gopurs_runtime.Value
var once_flip__3709724320 sync.Once
func Get_flip__3709724320() gopurs_runtime.Value {
	once_flip__3709724320.Do(func() {
		cache_flip__3709724320 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3709724320(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3709724320
}

var cache_flip__3563101792 gopurs_runtime.Value
var once_flip__3563101792 sync.Once
func Get_flip__3563101792() gopurs_runtime.Value {
	once_flip__3563101792.Do(func() {
		cache_flip__3563101792 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3563101792(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3563101792
}

var cache_flip__1295484160 gopurs_runtime.Value
var once_flip__1295484160 sync.Once
func Get_flip__1295484160() gopurs_runtime.Value {
	once_flip__1295484160.Do(func() {
		cache_flip__1295484160 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1295484160(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1295484160
}

var cache_flip__1172723328 gopurs_runtime.Value
var once_flip__1172723328 sync.Once
func Get_flip__1172723328() gopurs_runtime.Value {
	once_flip__1172723328.Do(func() {
		cache_flip__1172723328 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1172723328(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1172723328
}

var cache_map__3116241637 gopurs_runtime.Value
var once_map__3116241637 sync.Once
func Get_map__3116241637() gopurs_runtime.Value {
	once_map__3116241637.Do(func() {
		cache_map__3116241637 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3116241637(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](dict_0_box))
})
	})
	return cache_map__3116241637
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__438443400 gopurs_runtime.Value
var once_map__438443400 sync.Once
func Get_map__438443400() gopurs_runtime.Value {
	once_map__438443400.Do(func() {
		cache_map__438443400 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__438443400(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_map__438443400
}

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_not__2235433470 gopurs_runtime.Value
var once_not__2235433470 sync.Once
func Get_not__2235433470() gopurs_runtime.Value {
	once_not__2235433470.Do(func() {
		cache_not__2235433470 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__2235433470(__eta0_0_box)
})
	})
	return cache_not__2235433470
}

var cache_fromList__970809024 gopurs_runtime.Value
var once_fromList__970809024 sync.Once
func Get_fromList__970809024() gopurs_runtime.Value {
	once_fromList__970809024.Do(func() {
		cache_fromList__970809024 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromList__970809024(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_fromList__970809024
}

var cache_fromList__1312353984 gopurs_runtime.Value
var once_fromList__1312353984 sync.Once
func Get_fromList__1312353984() gopurs_runtime.Value {
	once_fromList__1312353984.Do(func() {
		cache_fromList__1312353984 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromList__1312353984(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))}
})
	})
	return cache_fromList__1312353984
}

var cache_altList__614667287 gopurs_runtime.Value
var once_altList__614667287 sync.Once
func Get_altList__614667287() gopurs_runtime.Value {
	once_altList__614667287.Do(func() {
		cache_altList__614667287 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_functorList()
}), gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"))
	})
	return cache_altList__614667287
}

var cache_applicativeList__615687001 gopurs_runtime.Value
var once_applicativeList__615687001 sync.Once
func Get_applicativeList__615687001() gopurs_runtime.Value {
	once_applicativeList__615687001.Do(func() {
		cache_applicativeList__615687001 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_applyList()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
}))
	})
	return cache_applicativeList__615687001
}

var cache_applyList__3072763993 gopurs_runtime.Value
var once_applyList__3072763993 sync.Once
func Get_applyList__3072763993() gopurs_runtime.Value {
	once_applyList__3072763993.Do(func() {
		cache_applyList__3072763993 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyList__3072763993
}

var cache_applyList__1109325167 gopurs_runtime.Value
var once_applyList__1109325167 sync.Once
func Get_applyList__1109325167() gopurs_runtime.Value {
	once_applyList__1109325167.Do(func() {
		cache_applyList__1109325167 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_functorList()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1358893437 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_0.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1))})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyList__1109325167
}

var cache_foldableList__1753400174 gopurs_runtime.Value
var once_foldableList__1753400174 sync.Once
func Get_foldableList__1753400174() gopurs_runtime.Value {
	once_foldableList__1753400174.Do(func() {
		cache_foldableList__1753400174 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Semigroup0_1_0.V0, acc_4)
_ = __local_var_5_2
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_2, gopurs_runtime.Apply(f_3, x_6))
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_1_3_3 gopurs_runtime.Value
go__go_1_3_3 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop gopurs_runtime.Value = b_2_loop_val
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_3_3:
for {
if false { continue go__go_1_3_3 }
var b_2 gopurs_runtime.Value = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t4 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t4 = b_2
goto end_branch_4
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_3_3
__t4 = gopurs_runtime.Value{}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return __t4
}
}()
})
})
return go__go_1_3_3
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_5 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), b_1)
_ = __local_var_2_5
var go__go_3_7_4 gopurs_runtime.Value
go__go_3_7_4 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_7_4:
for {
if false { continue go__go_3_7_4 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t8 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_8
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_7_4
__t8 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t8))}
}
}()
})
})
__local_var_3_6 := gopurs_runtime.Apply(go__go_3_7_4, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_6
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_5, gopurs_runtime.Apply(__local_var_3_6, x_4))
})
})
}))
	})
	return cache_foldableList__1753400174
}

var cache_functorList__4121998062 gopurs_runtime.Value
var once_functorList__4121998062 sync.Once
func Get_functorList__4121998062() gopurs_runtime.Value {
	once_functorList__4121998062.Do(func() {
		cache_functorList__4121998062 = gopurs_runtime.RecordDict1("map", pkg_Data_List_Types.Get_listMap())
	})
	return cache_functorList__4121998062
}

var cache_functorList__1783129585 gopurs_runtime.Value
var once_functorList__1783129585 sync.Once
func Get_functorList__1783129585() gopurs_runtime.Value {
	once_functorList__1783129585.Do(func() {
		cache_functorList__1783129585 = gopurs_runtime.RecordDict1("map", pkg_Data_List_Types.Get_listMap())
	})
	return cache_functorList__1783129585
}

var cache_listMap__4135416762 gopurs_runtime.Value
var once_listMap__4135416762 sync.Once
func Get_listMap__4135416762() gopurs_runtime.Value {
	once_listMap__4135416762.Do(func() {
		cache_listMap__4135416762 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_listMap__4135416762(f_0_box)
})
	})
	return cache_listMap__4135416762
}

var cache_plusList__598824825 gopurs_runtime.Value
var once_plusList__598824825 sync.Once
func Get_plusList__598824825() gopurs_runtime.Value {
	once_plusList__598824825.Do(func() {
		cache_plusList__598824825 = gopurs_runtime.RecordDict2("Alt0", "empty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_altList()
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
	})
	return cache_plusList__598824825
}

var cache_semigroupList__2766094215 gopurs_runtime.Value
var once_semigroupList__2766094215 sync.Once
func Get_semigroupList__2766094215() gopurs_runtime.Value {
	once_semigroupList__2766094215.Do(func() {
		cache_semigroupList__2766094215 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(xs_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](ys_1))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](xs_0))})))}
})
}))
	})
	return cache_semigroupList__2766094215
}

var cache_unfoldable1List__3672302568 gopurs_runtime.Value
var once_unfoldable1List__3672302568 sync.Once
func Get_unfoldable1List__3672302568() gopurs_runtime.Value {
	once_unfoldable1List__3672302568.Do(func() {
		cache_unfoldable1List__3672302568 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_7 gopurs_runtime.Value
go__go_2_0_7 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_7:
for {
if false { continue go__go_2_0_7 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t4 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 930809136 && __t_tag_2.UnsafePtr != nil) {
source_3_loop = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1.UnsafePtr).V0
memo_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})
continue go__go_2_0_7
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
var __t_tag_3 gopurs_runtime.Value = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V1
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 930809136 && __t_tag_3.UnsafePtr == nil) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0, memo_4})})))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t4))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_7, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_unfoldable1List__3672302568
}

var cache_unfoldableList__2633941518 gopurs_runtime.Value
var once_unfoldableList__2633941518 sync.Once
func Get_unfoldableList__2633941518() gopurs_runtime.Value {
	once_unfoldableList__2633941518.Do(func() {
		cache_unfoldableList__2633941518 = gopurs_runtime.RecordDict2("Unfoldable10", "unfoldr", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_List_Types.Get_unfoldable1List()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_1 gopurs_runtime.Value) gopurs_runtime.Value {
var go__go_2_0_8 gopurs_runtime.Value
go__go_2_0_8 = gopurs_runtime.Func(func(source_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(memo_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var source_3_loop gopurs_runtime.Value = source_3_loop_val
var memo_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](memo_4_loop_val)
go__go_2_0_8:
for {
if false { continue go__go_2_0_8 }
var source_3 gopurs_runtime.Value = source_3_loop
_ = source_3
var memo_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = memo_4_loop
_ = memo_4
v_5_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]](gopurs_runtime.Apply(f_0, source_3))
_ = v_5_1
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_7, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_6)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(memo_4)})))}
goto end_branch_2
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr != nil) {
source_3_loop = (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V1
memo_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_5_1)}.UnsafePtr).V0.UnsafePtr).V0, memo_4})})
continue go__go_2_0_8
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2))}
}
}()
})
})
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_2_0_8, b_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})))}
})
}))
	})
	return cache_unfoldableList__2633941518
}

var cache_null__74357383 gopurs_runtime.Value
var once_null__74357383 sync.Once
func Get_null__74357383() gopurs_runtime.Value {
	once_null__74357383.Do(func() {
		cache_null__74357383 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null__74357383(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_null__74357383
}

var cache_null__2437342685 gopurs_runtime.Value
var once_null__2437342685 sync.Once
func Get_null__2437342685() gopurs_runtime.Value {
	once_null__2437342685.Do(func() {
		cache_null__2437342685 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null__2437342685(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_null__2437342685
}

var cache_reverse__4230102656 gopurs_runtime.Value
var once_reverse__4230102656 sync.Once
func Get_reverse__4230102656() gopurs_runtime.Value {
	once_reverse__4230102656.Do(func() {
		cache_reverse__4230102656 = func() gopurs_runtime.Value {
var go__go_0_0_9 gopurs_runtime.Value
go__go_0_0_9 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_9:
for {
if false { continue go__go_0_0_9 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
continue go__go_0_0_9
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_9, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_reverse__4230102656
}

var cache_reverse__1758336384 gopurs_runtime.Value
var once_reverse__1758336384 sync.Once
func Get_reverse__1758336384() gopurs_runtime.Value {
	once_reverse__1758336384.Do(func() {
		cache_reverse__1758336384 = func() gopurs_runtime.Value {
var go__go_0_0_10 gopurs_runtime.Value
go__go_0_0_10 = gopurs_runtime.Func(func(v_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_1_loop_val)
var v1_2_loop gopurs_runtime.Value = v1_2_loop_val
go__go_0_0_10:
for {
if false { continue go__go_0_0_10 }
var v_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
var __t1 gopurs_runtime.Value
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_1)}
goto end_branch_1
} else {

}
}
{
if (v1_2.Type == 9 && v1_2.IntVal == 1358893437 && v1_2.UnsafePtr != nil) {
v_1_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V0, v_1})})
v1_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_2.UnsafePtr).V1)}
continue go__go_0_0_10
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return gopurs_runtime.Apply(go__go_0_0_10, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}()
	})
	return cache_reverse__1758336384
}

var cache_span__799093643 gopurs_runtime.Value
var once_span__799093643 sync.Once
func Get_span__799093643() gopurs_runtime.Value {
	once_span__799093643.Do(func() {
		cache_span__799093643 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span__799093643(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box))
})
	})
	return cache_span__799093643
}

var cache_span__2133741451 gopurs_runtime.Value
var once_span__2133741451 sync.Once
func Get_span__2133741451() gopurs_runtime.Value {
	once_span__2133741451.Do(func() {
		cache_span__2133741451 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_span__2133741451(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_1_box))
})
	})
	return cache_span__2133741451
}

var cache_iterMapL__878452066 gopurs_runtime.Value
var once_iterMapL__878452066 sync.Once
func Get_iterMapL__878452066() gopurs_runtime.Value {
	once_iterMapL__878452066.Do(func() {
		cache_iterMapL__878452066 = func() gopurs_runtime.Value {
var go__go_0_0_11 gopurs_runtime.Value
go__go_0_0_11 = gopurs_runtime.Func(func(iter_1_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var iter_1_loop gopurs_runtime.Value = iter_1_loop_val
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_0_0_11:
for {
if false { continue go__go_0_0_11 }
var iter_1 gopurs_runtime.Value = iter_1_loop
_ = iter_1
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = iter_1
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
var __t2 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 324739070 && __t_tag_1.UnsafePtr == nil) {
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, iter_1})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_11
__t2 = gopurs_runtime.Value{}
goto end_branch_2
} else {

}
}
{
iter_1_loop = gopurs_runtime.Value{Type: 9, IntVal: 1343415489, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5, iter_1})}})}
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_0_0_11
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
__t3 = __t2
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
return go__go_0_0_11
}()
	})
	return cache_iterMapL__878452066
}

var cache_lookup__3378638282 gopurs_runtime.Value
var once_lookup__3378638282 sync.Once
func Get_lookup__3378638282() gopurs_runtime.Value {
	once_lookup__3378638282.Do(func() {
		cache_lookup__3378638282 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, k_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookup__3378638282(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), k_1_box)
})
	})
	return cache_lookup__3378638282
}

var cache_lookup__1040249709 gopurs_runtime.Value
var once_lookup__1040249709 sync.Once
func Get_lookup__1040249709() gopurs_runtime.Value {
	once_lookup__1040249709.Do(func() {
		cache_lookup__1040249709 = gopurs_runtime.Func(func(k_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookup__1040249709(k_0_box)
})
	})
	return cache_lookup__1040249709
}

var cache_size__1374028086 gopurs_runtime.Value
var once_size__1374028086 sync.Once
func Get_size__1374028086() gopurs_runtime.Value {
	once_size__1374028086.Do(func() {
		cache_size__1374028086 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_size__1374028086(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_size__1374028086
}

var cache_size__2382154916 gopurs_runtime.Value
var once_size__2382154916 sync.Once
func Get_size__2382154916() gopurs_runtime.Value {
	once_size__2382154916.Do(func() {
		cache_size__2382154916 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_size__2382154916(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))
})
	})
	return cache_size__2382154916
}

var cache_stepAscCps__3090303421 gopurs_runtime.Value
var once_stepAscCps__3090303421 sync.Once
func Get_stepAscCps__3090303421() gopurs_runtime.Value {
	once_stepAscCps__3090303421.Do(func() {
		cache_stepAscCps__3090303421 = gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL())
	})
	return cache_stepAscCps__3090303421
}

var cache_stepAscCps__1323290822 gopurs_runtime.Value
var once_stepAscCps__1323290822 sync.Once
func Get_stepAscCps__1323290822() gopurs_runtime.Value {
	once_stepAscCps__1323290822.Do(func() {
		cache_stepAscCps__1323290822 = gopurs_runtime.Apply(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL())
	})
	return cache_stepAscCps__1323290822
}

var cache_stepUnfoldr__966001626 gopurs_runtime.Value
var once_stepUnfoldr__966001626 sync.Once
func Get_stepUnfoldr__966001626() gopurs_runtime.Value {
	once_stepUnfoldr__966001626.Do(func() {
		cache_stepUnfoldr__966001626 = gopurs_runtime.Apply3(pkg_Data_Map_Internal.Get_stepWith(), pkg_Data_Map_Internal.Get_iterMapL(), gopurs_runtime.Func3(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value, __local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, __local_var_0, __local_var_1})}, __local_var_2})}})}
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}))
	})
	return cache_stepUnfoldr__966001626
}

var cache_stepWith__3186376421 gopurs_runtime.Value
var once_stepWith__3186376421 sync.Once
func Get_stepWith__3186376421() gopurs_runtime.Value {
	once_stepWith__3186376421.Do(func() {
		cache_stepWith__3186376421 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, next_1_box gopurs_runtime.Value, done_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stepWith__3186376421(f_0_box, next_1_box, done_2_box)
})
	})
	return cache_stepWith__3186376421
}

var cache_toMapIter__2014410513 gopurs_runtime.Value
var once_toMapIter__2014410513 sync.Once
func Get_toMapIter__2014410513() gopurs_runtime.Value {
	once_toMapIter__2014410513.Do(func() {
		cache_toMapIter__2014410513 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toMapIter__2014410513(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](a_0_box))
})
	})
	return cache_toMapIter__2014410513
}

var cache_toUnfoldable__2183602684 gopurs_runtime.Value
var once_toUnfoldable__2183602684 sync.Once
func Get_toUnfoldable__2183602684() gopurs_runtime.Value {
	once_toUnfoldable__2183602684.Do(func() {
		cache_toUnfoldable__2183602684 = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable__2183602684(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box))
})
	})
	return cache_toUnfoldable__2183602684
}

var cache_toUnfoldable__2567957978 gopurs_runtime.Value
var once_toUnfoldable__2567957978 sync.Once
func Get_toUnfoldable__2567957978() gopurs_runtime.Value {
	once_toUnfoldable__2567957978.Do(func() {
		cache_toUnfoldable__2567957978 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUnfoldable__2567957978(__eta0_0_box)
})
	})
	return cache_toUnfoldable__2567957978
}

var cache_isJust__4165351782 gopurs_runtime.Value
var once_isJust__4165351782 sync.Once
func Get_isJust__4165351782() gopurs_runtime.Value {
	once_isJust__4165351782.Do(func() {
		cache_isJust__4165351782 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__4165351782(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isJust__4165351782
}

var cache_isJust__4206805139 gopurs_runtime.Value
var once_isJust__4206805139 sync.Once
func Get_isJust__4206805139() gopurs_runtime.Value {
	once_isJust__4206805139.Do(func() {
		cache_isJust__4206805139 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__4206805139(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isJust__4206805139
}

var cache_maybe__1594528518 gopurs_runtime.Value
var once_maybe__1594528518 sync.Once
func Get_maybe__1594528518() gopurs_runtime.Value {
	once_maybe__1594528518.Do(func() {
		cache_maybe__1594528518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1594528518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1594528518
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_unwrap__777744115 gopurs_runtime.Value
var once_unwrap__777744115 sync.Once
func Get_unwrap__777744115() gopurs_runtime.Value {
	once_unwrap__777744115.Do(func() {
		cache_unwrap__777744115 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__777744115(__eta0_0_box)
})
	})
	return cache_unwrap__777744115
}

var cache_unwrap__3267718003 gopurs_runtime.Value
var once_unwrap__3267718003 sync.Once
func Get_unwrap__3267718003() gopurs_runtime.Value {
	once_unwrap__3267718003.Do(func() {
		cache_unwrap__3267718003 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unwrap__3267718003(gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_0_box))
})
	})
	return cache_unwrap__3267718003
}

var cache_compare__669572705 gopurs_runtime.Value
var once_compare__669572705 sync.Once
func Get_compare__669572705() gopurs_runtime.Value {
	once_compare__669572705.Do(func() {
		cache_compare__669572705 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__669572705(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__669572705
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_compare__3077449111 gopurs_runtime.Value
var once_compare__3077449111 sync.Once
func Get_compare__3077449111() gopurs_runtime.Value {
	once_compare__3077449111.Do(func() {
		cache_compare__3077449111 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, y_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_compare__3077449111(uint32(x_0_box.IntVal), uint32(y_1_box.IntVal))), UnsafePtr: nil}
})
	})
	return cache_compare__3077449111
}

var cache_compare__231252914 gopurs_runtime.Value
var once_compare__231252914 sync.Once
func Get_compare__231252914() gopurs_runtime.Value {
	once_compare__231252914.Do(func() {
		cache_compare__231252914 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__231252914(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_compare__231252914
}

var cache_greaterThan__4087042607 gopurs_runtime.Value
var once_greaterThan__4087042607 sync.Once
func Get_greaterThan__4087042607() gopurs_runtime.Value {
	once_greaterThan__4087042607.Do(func() {
		cache_greaterThan__4087042607 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__4087042607(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__4087042607
}

var cache_greaterThan__1061005983 gopurs_runtime.Value
var once_greaterThan__1061005983 sync.Once
func Get_greaterThan__1061005983() gopurs_runtime.Value {
	once_greaterThan__1061005983.Do(func() {
		cache_greaterThan__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThan__1061005983
}

var cache_greaterThan__1409282474 gopurs_runtime.Value
var once_greaterThan__1409282474 sync.Once
func Get_greaterThan__1409282474() gopurs_runtime.Value {
	once_greaterThan__1409282474.Do(func() {
		cache_greaterThan__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThan__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThan__1409282474
}

var cache_greaterThanOrEq__1061005983 gopurs_runtime.Value
var once_greaterThanOrEq__1061005983 sync.Once
func Get_greaterThanOrEq__1061005983() gopurs_runtime.Value {
	once_greaterThanOrEq__1061005983.Do(func() {
		cache_greaterThanOrEq__1061005983 = gopurs_runtime.Func2(func(a1_0_box gopurs_runtime.Value, a2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(a1_0_box, a2_1_box))
})
	})
	return cache_greaterThanOrEq__1061005983
}

var cache_greaterThanOrEq__1409282474 gopurs_runtime.Value
var once_greaterThanOrEq__1409282474 sync.Once
func Get_greaterThanOrEq__1409282474() gopurs_runtime.Value {
	once_greaterThanOrEq__1409282474.Do(func() {
		cache_greaterThanOrEq__1409282474 = gopurs_runtime.Func3(func(dictOrd_0_box gopurs_runtime.Value, a1_1_box gopurs_runtime.Value, a2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_greaterThanOrEq__1409282474(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box), a1_1_box, a2_2_box))
})
	})
	return cache_greaterThanOrEq__1409282474
}

var cache_append__1124926121 gopurs_runtime.Value
var once_append__1124926121 sync.Once
func Get_append__1124926121() gopurs_runtime.Value {
	once_append__1124926121.Do(func() {
		cache_append__1124926121 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1124926121(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1124926121
}

var cache_append__3641242355 gopurs_runtime.Value
var once_append__3641242355 sync.Once
func Get_append__3641242355() gopurs_runtime.Value {
	once_append__3641242355.Do(func() {
		cache_append__3641242355 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__3641242355(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]]](dict_0_box))
})
	})
	return cache_append__3641242355
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_append__2013893496 gopurs_runtime.Value
var once_append__2013893496 sync.Once
func Get_append__2013893496() gopurs_runtime.Value {
	once_append__2013893496.Do(func() {
		cache_append__2013893496 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, ys_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_append__2013893496(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](xs_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](ys_1_box)))}
})
	})
	return cache_append__2013893496
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_show__1261750354 gopurs_runtime.Value
var once_show__1261750354 sync.Once
func Get_show__1261750354() gopurs_runtime.Value {
	once_show__1261750354.Do(func() {
		cache_show__1261750354 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_show__1261750354(uint32(v_0_box.IntVal)))
})
	})
	return cache_show__1261750354
}

var cache_show__2896747026 gopurs_runtime.Value
var once_show__2896747026 sync.Once
func Get_show__2896747026() gopurs_runtime.Value {
	once_show__2896747026.Do(func() {
		cache_show__2896747026 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2896747026(__eta0_0_box)
})
	})
	return cache_show__2896747026
}

var cache_snd__2019004820 gopurs_runtime.Value
var once_snd__2019004820 sync.Once
func Get_snd__2019004820() gopurs_runtime.Value {
	once_snd__2019004820.Do(func() {
		cache_snd__2019004820 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float(Call_snd__2019004820(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[uint32, float64]](v_0_box)))
})
	})
	return cache_snd__2019004820
}

var cache_unfoldr__2235715281 gopurs_runtime.Value
var once_unfoldr__2235715281 sync.Once
func Get_unfoldr__2235715281() gopurs_runtime.Value {
	once_unfoldr__2235715281.Do(func() {
		cache_unfoldr__2235715281 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__2235715281(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__2235715281
}

var cache_unfoldr__3990862552 gopurs_runtime.Value
var once_unfoldr__3990862552 sync.Once
func Get_unfoldr__3990862552() gopurs_runtime.Value {
	once_unfoldr__3990862552.Do(func() {
		cache_unfoldr__3990862552 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__3990862552(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__3990862552
}

var cache_unfoldr__1128708256 gopurs_runtime.Value
var once_unfoldr__1128708256 sync.Once
func Get_unfoldr__1128708256() gopurs_runtime.Value {
	once_unfoldr__1128708256.Do(func() {
		cache_unfoldr__1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1128708256(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1128708256
}

var cache_unfoldr__1519733018 gopurs_runtime.Value
var once_unfoldr__1519733018 sync.Once
func Get_unfoldr__1519733018() gopurs_runtime.Value {
	once_unfoldr__1519733018.Do(func() {
		cache_unfoldr__1519733018 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1519733018(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1519733018
}

type Constructor_IsEmpty struct {
	Rc uint32
}


type Constructor_InvalidWeekComponentUsage struct {
	Rc uint32
}


type Constructor_ContainsNegativeValue struct {
	Rc uint32
	V0 uint32
}


type Constructor_InvalidFractionalUse struct {
	Rc uint32
	V0 uint32
}


func Call_unIsoDuration(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[uint32, float64]) *pkg_Data_Map_Internal.Constructor_Node[uint32, float64] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[uint32, float64] = v_0_loop
_ = v_0
return v_0
}

func Call_prettyError(v_0_loop gopurs_runtime.Value) string {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 1422140417) {
__t0 = gopurs_runtime.Str("Duration is empty (has no components)")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 1775501833) {
__t0 = gopurs_runtime.Str("Week component of Duration is used with other components")
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 3224543173) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("Component `"), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Str(Call_show__1261750354((*Constructor_ContainsNegativeValue)(v_0.UnsafePtr).V0)).StrVal()), gopurs_runtime.Str("` contains negative value")).StrVal())).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
__t0 = gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("Invalid usage of Fractional value at component `"), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Str(Call_show__1261750354((*Constructor_InvalidFractionalUse)(v_0.UnsafePtr).V0)).StrVal()), gopurs_runtime.Str("`")).StrVal())).StrVal())
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.StrVal()
}

func Call_checkWeekUsage(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Bool(Call_isJust__4165351782(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[float64]](gopurs_runtime.Apply(Call_lookup__1040249709(gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}), gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](gopurs_runtime.RecordGet(v_0, "asMap")))})))}))).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Bool(Call_greaterThan__4087042607(gopurs_runtime.Int(gopurs_runtime.Int(Call_size__2382154916(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[uint32, float64]](gopurs_runtime.RecordGet(v_0, "asMap")))}))).IntVal), gopurs_runtime.Int(1))).IntVal) != (0))).IntVal) != (0) {
__t0 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_pure__2154981942(gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: unsafe.Pointer(nil)}))})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty"))
}
end_branch_0:
return __t0
}

func Call_checkNegativeValues(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_foldMap(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Bool(Call_greaterThanOrEq__1061005983(gopurs_runtime.Float((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1.FloatVal()), gopurs_runtime.Float(0.0))).IntVal) != (0) {
__t0 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty"))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_pure__2154981942(gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Constructor_ContainsNegativeValue{1, uint32((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0.IntVal)})}))})
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t0)}
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]]](gopurs_runtime.RecordGet(v_0, "asList")))}))
}

func Call_checkFractionalUse(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_1 := Call_not__2235433470(gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Bool(Call_notEq__2334967935(gopurs_runtime.Float(gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), gopurs_runtime.Float(a_1.FloatVal())).FloatVal()), gopurs_runtime.Float(a_1.FloatVal()))).IntVal) != (0))
}))
_ = __local_var_1_1
v1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]]](gopurs_runtime.RecordGet(Call_span__2133741451(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1)
}), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]]](gopurs_runtime.RecordGet(v_0, "asList")))})), "rest"))
_ = v1_1_0
var __t2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr != nil)) && ((gopurs_runtime.Bool(Call_greaterThan__1061005983(gopurs_runtime.Float(Call_unwrap__777744115(gopurs_runtime.Float(Call_foldMap__193737345(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr).V1)}).FloatVal())).FloatVal()), gopurs_runtime.Float(0.0))).IntVal) != (0)) {
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_pure__2154981942(gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Constructor_InvalidFractionalUse{1, uint32((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr).V0.UnsafePtr).V0.IntVal)})}))})
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty"))
}
end_branch_2:
return __t2
}

func Call_checkEmptiness(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
if (gopurs_runtime.Bool(Call_null__2437342685(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]]](gopurs_runtime.RecordGet(v_0, "asList")))}))).IntVal) != (0) {
__t0 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_pure__2154981942(gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: unsafe.Pointer(nil)}))})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty"))
}
end_branch_0:
return __t0
}

func Call_checkValidIsoDuration(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[uint32, float64]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[uint32, float64] = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply4(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidFn()))}, pkg_Data_Foldable.Get_identity1(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{Get_checkWeekUsage(), Get_checkEmptiness(), Get_checkFractionalUse(), Get_checkNegativeValues()}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.RecordDict2("asList", "asMap", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]]](gopurs_runtime.Apply(Get_reverse__1758336384(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]]](Call_toUnfoldable__2567957978(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))})))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
}

func Call_mkIsoDuration(d_0_loop *pkg_Data_Map_Internal.Constructor_Node[uint32, float64]) gopurs_runtime.Value {
var d_0 *pkg_Data_Map_Internal.Constructor_Node[uint32, float64] = d_0_loop
_ = d_0
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromList__1312353984(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_checkValidIsoDuration(d_0))})))})
_ = v_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0))}})}
goto end_branch_1
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(d_0)}})}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__2154981942(a_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, a_0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})
}

func Call_apply__1030762512(dict_0_loop *pkg_Control_Apply.Constructor_Apply[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2169384906(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_semigroupList(), "append"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_functorList(), "map"), (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applyList(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})))})))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_empty__932402776(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "empty")
}

func Call_compose__1987728071(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__3249109794(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__1555187646(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__2527254334(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__2995688990(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_composeFlipped__2583068543(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
}

func Call_eq__789642299(dict_0_loop *pkg_Data_Eq.Constructor_Eq[bool]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[bool] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__1697837627(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2276491096(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) == ((__eta1_1.IntVal) != (0)))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__1241439021(x_0_loop uint32, y_1_loop uint32) bool {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
var __t6 bool
{
if (x_0 == 3908053364) {
var __t0 bool
{
if (y_1 == 3908053364) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
__t6 = __t0
goto end_branch_6
} else {

}
}
{
if (x_0 == 217821258) {
var __t1 bool
{
if (y_1 == 217821258) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
__t6 = __t1
goto end_branch_6
} else {

}
}
{
if (x_0 == 1292308612) {
var __t2 bool
{
if (y_1 == 1292308612) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
__t6 = __t2
goto end_branch_6
} else {

}
}
{
if (x_0 == 2311060696) {
var __t3 bool
{
if (y_1 == 2311060696) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
__t6 = __t3
goto end_branch_6
} else {

}
}
{
if (x_0 == 401302776) {
var __t4 bool
{
if (y_1 == 401302776) {
__t4 = true
goto end_branch_4
} else {

}
}
{
__t4 = false
}
end_branch_4:
__t6 = __t4
goto end_branch_6
} else {

}
}
{
if (x_0 == 3327533908) {
var __t5 bool
{
if (y_1 == 3327533908) {
__t5 = true
goto end_branch_5
} else {

}
}
{
__t5 = false
}
end_branch_5:
__t6 = __t5
goto end_branch_6
} else {

}
}
{
if ((x_0 == 3631736139)) && ((y_1 == 3631736139)) {
__t6 = true
goto end_branch_6
} else {

}
}
{
__t6 = false
}
end_branch_6:
return __t6
}

func Call_eq__2224314568(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[*pkg_Data_Map_Internal.Constructor_Node[uint32, float64]]](pkg_Data_Interval_Duration.Get_eqMap()).V0, __eta0_0, __eta1_1)
}

func Call_notEq__2334967935(x_0_loop gopurs_runtime.Value, y_1_loop gopurs_runtime.Value) bool {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var y_1 gopurs_runtime.Value = y_1_loop
_ = y_1
return ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqNumber(), "eq"), x_0, y_1).IntVal) != (0)) != (true)
}

func Call_notEq__2384498378(dictEq_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value, y_2_loop gopurs_runtime.Value) bool {
var dictEq_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var y_2 gopurs_runtime.Value = y_2_loop
_ = y_2
return ((gopurs_runtime.Apply2(dictEq_0.V0, x_1, y_2).IntVal) != (0)) != (true)
}

func Call_fold__910331789(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
return gopurs_runtime.Apply2(dictFoldable_0.V0, gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_1)}, pkg_Data_Foldable.Get_identity1())
}

func Call_foldMap__4098395794(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldMap__193737345(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[float64]](gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[float64]](Get_monoidAdditive()).V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(acc_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(Semigroup0_2_0.V0, gopurs_runtime.Float(acc_3.FloatVal()))
_ = __local_var_4_1
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_1, gopurs_runtime.Apply(__eta0_0, x_5))
})
}), gopurs_runtime.Float(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[float64]](Get_monoidAdditive()).V1.FloatVal()), __eta1_1)
}

func Call_foldl__2699291984(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__2602334544(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__1656262032(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__3943124669(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
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
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = b_2
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.Apply2(f_0, b_2, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0)
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_0
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
})
return go__go_1_0_0
}

func Call_foldl__396932925(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var go__go_1_0_1 gopurs_runtime.Value
go__go_1_0_1 = gopurs_runtime.Func(func(b_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var b_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](b_2_loop_val)
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_1_0_1:
for {
if false { continue go__go_1_0_1 }
var b_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = b_2_loop
_ = b_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t1 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}
goto end_branch_1
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 1358893437 && v_3.UnsafePtr != nil) {
b_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(f_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_2)}, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V0))
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_3.UnsafePtr).V1)}
continue go__go_1_0_1
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t1))}
}
}()
})
})
return go__go_1_0_1
}

func Call_foldr__2512763050(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__3673994608(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr__2979608669(f_0_loop gopurs_runtime.Value, b_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = b_1_loop
_ = b_1
__local_var_2_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldl"), gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, a_3, b_2)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(b_1)})
_ = __local_var_2_0
var go__go_3_2_2 gopurs_runtime.Value
go__go_3_2_2 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_4_loop_val)
var v1_5_loop gopurs_runtime.Value = v1_5_loop_val
go__go_3_2_2:
for {
if false { continue go__go_3_2_2 }
var v_4 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_4_loop
_ = v_4
var v1_5 gopurs_runtime.Value = v1_5_loop
_ = v1_5
var __t3 gopurs_runtime.Value
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_4)}
goto end_branch_3
} else {

}
}
{
if (v1_5.Type == 9 && v1_5.IntVal == 1358893437 && v1_5.UnsafePtr != nil) {
v_4_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V0, v_4})})
v1_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_5.UnsafePtr).V1)}
continue go__go_3_2_2
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t3))}
}
}()
})
})
__local_var_3_1 := gopurs_runtime.Apply(go__go_3_2_2, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
}

func Call_const__220790420(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__2974723072(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3709724320(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3563101792(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1295484160(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1172723328(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__3116241637(dict_0_loop *pkg_Data_Functor.Constructor_Functor[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__438443400(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(pkg_Data_List_Types.Get_listMap(), __eta0_0, __eta1_1)
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_not__2235433470(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](Get_heytingAlgebraFunction()).V4, __eta0_0)
}

func Call_fromList__970809024(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}})}})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t0)
}

func Call_fromList__1312353984(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]] {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)}})}})}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t0)
}

func Call_listMap__4135416762(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var chunkedRevMap_1_0_5 gopurs_runtime.Value
chunkedRevMap_1_0_5 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](v_2_loop_val)
var v1_3_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v1_3_loop_val)
chunkedRevMap_1_0_5:
for {
if false { continue chunkedRevMap_1_0_5 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = v_2_loop
_ = v_2
var v1_3 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_3_loop
_ = v1_3
var __t19 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_18 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {

var __t_tag_15 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
var __t_and_17 bool = false
if (__t_tag_15.Type == 9 && __t_tag_15.IntVal == 1358893437 && __t_tag_15.UnsafePtr != nil) {

var __t_tag_16 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_17 = (__t_tag_16.Type == 9 && __t_tag_16.IntVal == 1358893437 && __t_tag_16.UnsafePtr != nil)
}
__t_and_18 = __t_and_17
}
if __t_and_18 {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)})})})
v1_3_loop = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V1
continue chunkedRevMap_1_0_5
__t19 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_19
} else {

}
}
{
var reverseUnrolledMap_4_1_6 gopurs_runtime.Value
reverseUnrolledMap_4_1_6 = gopurs_runtime.Func(func(v2_5_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v3_6_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v2_5_loop gopurs_runtime.Value = v2_5_loop_val
var v3_6_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v3_6_loop_val)
reverseUnrolledMap_4_1_6:
for {
if false { continue reverseUnrolledMap_4_1_6 }
var v2_5 gopurs_runtime.Value = v2_5_loop
_ = v2_5
var v3_6 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v3_6_loop
_ = v3_6
var __t8 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]
{
var __t_and_7 bool = false
if (v2_5.Type == 9 && v2_5.IntVal == 1358893437 && v2_5.UnsafePtr != nil) {

var __t_tag_2 gopurs_runtime.Value = (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0
var __t_and_6 bool = false
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil) {

var __t_tag_3 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}
var __t_and_5 bool = false
if (__t_tag_3.Type == 9 && __t_tag_3.IntVal == 1358893437 && __t_tag_3.UnsafePtr != nil) {

var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}
__t_and_5 = (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 1358893437 && __t_tag_4.UnsafePtr != nil)
}
__t_and_6 = __t_and_5
}
__t_and_7 = __t_and_6
}
if __t_and_7 {
v2_5_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V1)}
v3_6_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v2_5.UnsafePtr).V0.UnsafePtr).V1)}.UnsafePtr).V1)}.UnsafePtr).V0), v3_6})})})})})})
continue reverseUnrolledMap_4_1_6
__t8 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{})
goto end_branch_8
} else {

}
}
{
__t8 = v3_6
}
end_branch_8:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t8)}
}
}()
})
})
var __t14 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr != nil) {
var __t13 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 1358893437 && __t_tag_9.UnsafePtr != nil) {
var __t11 gopurs_runtime.Value
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V1)}
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 1358893437 && __t_tag_10.UnsafePtr == nil) {
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})})})}
goto end_branch_11
} else {

}
}
{
__t11 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_11:
__t13 = __t11
goto end_branch_13
} else {

}
}
{
var __t_tag_12 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V1)}
if (__t_tag_12.Type == 9 && __t_tag_12.IntVal == 1358893437 && __t_tag_12.UnsafePtr == nil) {
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_3)}.UnsafePtr).V0), gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})})}
goto end_branch_13
} else {

}
}
{
__t13 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_13:
__t14 = __t13
goto end_branch_14
} else {

}
}
{
__t14 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}
}
end_branch_14:
__t19 = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(reverseUnrolledMap_4_1_6, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t14))}))
}
end_branch_19:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(__t19)}
}
}()
})
})
return gopurs_runtime.Apply(chunkedRevMap_1_0_5, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))})
}

func Call_null__74357383(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) bool {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 bool
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_null__2437342685(v_0_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) bool {
var v_0 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 bool
{
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_span__799093643(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0).IntVal) != (0)) {
v2_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_span(), v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})
_ = v2_2_0
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}

func Call_span__2133741451(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
if ((gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil)) && ((gopurs_runtime.Apply(v_0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0).IntVal) != (0)) {
v2_2_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_span(), v_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V1)})
_ = v2_2_0
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0, gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "init"))})}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(v2_2_0, "rest")))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(nil))}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1)})
}
end_branch_1:
return __t1
}

func Call_lookup__3378638282(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], k_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var k_1 gopurs_runtime.Value = k_1_loop
_ = k_1
var go__go_2_0_12 gopurs_runtime.Value
go__go_2_0_12 = gopurs_runtime.Func(func(v_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_3_loop gopurs_runtime.Value = v_3_loop_val
go__go_2_0_12:
for {
if false { continue go__go_2_0_12 }
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var __t3 gopurs_runtime.Value
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_3.Type == 9 && v_3.IntVal == 324739070 && v_3.UnsafePtr != nil) {
v1_4_1 := gopurs_runtime.Apply2(dictOrd_0.V1, k_1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V2)
_ = v1_4_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_4_1.IntVal) == 1527465420) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V4)}
continue go__go_2_0_12
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 380165415) {
v_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V5)}
continue go__go_2_0_12
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_4_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V3})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}
}()
})
return go__go_2_0_12
}

func Call_lookup__1040249709(k_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var go__go_1_0_13 gopurs_runtime.Value
go__go_1_0_13 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_0_13:
for {
if false { continue go__go_1_0_13 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t3 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
v1_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_ordDurationComponent(), "compare"), k_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
_ = v1_3_1
var __t2 gopurs_runtime.Value
{
if (uint32(v1_3_1.IntVal) == 1527465420) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_1_0_13
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 380165415) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
continue go__go_1_0_13
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_2
} else {

}
}
{
if (uint32(v1_3_1.IntVal) == 902936544) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3})}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t3))}
}
}()
})
return go__go_1_0_13
}

func Call_size__1374028086(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) int64 {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}

func Call_size__2382154916(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) int64 {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.IntVal == 324739070 && gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.IntVal
}

func Call_stepWith__3186376421(f_0_loop gopurs_runtime.Value, next_1_loop gopurs_runtime.Value, done_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var next_1 gopurs_runtime.Value = next_1_loop
_ = next_1
var done_2 gopurs_runtime.Value = done_2_loop
_ = done_2
var go__go_3_0_14 gopurs_runtime.Value
go__go_3_0_14 = gopurs_runtime.Func(func(v_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_4_loop gopurs_runtime.Value = v_4_loop_val
go__go_3_0_14:
for {
if false { continue go__go_3_0_14 }
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
var __t1 gopurs_runtime.Value
{
if (v_4.Type == 9 && v_4.IntVal == 2509360378) {
__t1 = gopurs_runtime.Apply(done_2, pkg_Data_Unit.Get_unit())
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 1343415489) {
__t1 = gopurs_runtime.UncurriedApp3(next_1, (*pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, (*pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, (*pkg_Data_Map_Internal.Constructor_IterEmit[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V2)
goto end_branch_1
} else {

}
}
{
if (v_4.Type == 9 && v_4.IntVal == 2861335956) {
v_4_loop = gopurs_runtime.Apply2(f_0, (*pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)})
continue go__go_3_0_14
__t1 = gopurs_runtime.Value{}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}
}()
})
return go__go_3_0_14
}

func Call_toMapIter__2014410513(a_0_loop *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var a_0 *pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value] = a_0_loop
_ = a_0
return gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_0, gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}
}

func Call_toUnfoldable__2183602684(dictUnfoldable_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictUnfoldable_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
__local_var_1_0 := gopurs_runtime.Apply(dictUnfoldable_0.V1, pkg_Data_Map_Internal.Get_stepUnfoldr())
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](x_2), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
})
}

func Call_toUnfoldable__2567957978(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_unfoldableList(), "unfoldr"), pkg_Data_Map_Internal.Get_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](__eta0_0), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})})
}

func Call_isJust__4165351782(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_isJust__4206805139(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__1594528518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_unwrap__777744115(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return __eta0_0
}

func Call_unwrap__3267718003(_dollar__unused_0_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_compare__669572705(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare__3077449111(x_0_loop uint32, y_1_loop uint32) uint32 {
var x_0 uint32 = x_0_loop
_ = x_0
var y_1 uint32 = y_1_loop
_ = y_1
var __t6 gopurs_runtime.Value
{
if (x_0 == 3908053364) {
var __t0 uint32
{
if (y_1 == 3908053364) {
__t0 = 902936544
goto end_branch_0
} else {

}
}
{
__t0 = 1527465420
}
end_branch_0:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t0), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (y_1 == 3908053364) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (x_0 == 217821258) {
var __t1 uint32
{
if (y_1 == 217821258) {
__t1 = 902936544
goto end_branch_1
} else {

}
}
{
__t1 = 1527465420
}
end_branch_1:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t1), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (y_1 == 217821258) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (x_0 == 1292308612) {
var __t2 uint32
{
if (y_1 == 1292308612) {
__t2 = 902936544
goto end_branch_2
} else {

}
}
{
__t2 = 1527465420
}
end_branch_2:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t2), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (y_1 == 1292308612) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (x_0 == 2311060696) {
var __t3 uint32
{
if (y_1 == 2311060696) {
__t3 = 902936544
goto end_branch_3
} else {

}
}
{
__t3 = 1527465420
}
end_branch_3:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (y_1 == 2311060696) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (x_0 == 401302776) {
var __t4 uint32
{
if (y_1 == 401302776) {
__t4 = 902936544
goto end_branch_4
} else {

}
}
{
__t4 = 1527465420
}
end_branch_4:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t4), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (y_1 == 401302776) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (x_0 == 3327533908) {
var __t5 uint32
{
if (y_1 == 3327533908) {
__t5 = 902936544
goto end_branch_5
} else {

}
}
{
__t5 = 1527465420
}
end_branch_5:
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(__t5), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if (y_1 == 3327533908) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
if ((x_0 == 3631736139)) && ((y_1 == 3631736139)) {
__t6 = gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return uint32(__t6.IntVal)
}

func Call_compare__231252914(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Apply2(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[*pkg_Data_Map_Internal.Constructor_Node[uint32, float64]]](pkg_Data_Interval_Duration.Get_ordMap()).V1, __eta0_0, __eta1_1)
}

func Call_greaterThan__4087042607(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.IntVal) > (a2_1.IntVal) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) > (a2_1.FloatVal()) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
return __t0
}

func Call_greaterThan__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 380165415) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
return __t1
}

func Call_greaterThanOrEq__1061005983(a1_0_loop gopurs_runtime.Value, a2_1_loop gopurs_runtime.Value) bool {
var a1_0 gopurs_runtime.Value = a1_0_loop
_ = a1_0
var a2_1 gopurs_runtime.Value = a2_1_loop
_ = a2_1
var __t0 bool
{
if (a1_0.FloatVal()) < (a2_1.FloatVal()) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
return __t0
}

func Call_greaterThanOrEq__1409282474(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value], a1_1_loop gopurs_runtime.Value, a2_2_loop gopurs_runtime.Value) bool {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var a1_1 gopurs_runtime.Value = a1_1_loop
_ = a1_1
var a2_2 gopurs_runtime.Value = a2_2_loop
_ = a2_2
var __t1 bool
{
var __t_tag_0 gopurs_runtime.Value = gopurs_runtime.Apply2(dictOrd_0.V1, a1_1, a2_2)
if (uint32(__t_tag_0.IntVal) == 1527465420) {
__t1 = false
goto end_branch_1
} else {

}
}
{
__t1 = true
}
end_branch_1:
return __t1
}

func Call_append__1124926121(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__3641242355(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__2013893496(xs_0_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]], ys_1_loop *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]) *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] {
var xs_0 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = xs_0_loop
_ = xs_0
var ys_1 *pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]] = ys_1_loop
_ = ys_1
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldr"), pkg_Data_List_Types.Get_Cons(), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(ys_1)}, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(xs_0)}))
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__1261750354(v_0_loop uint32) string {
var v_0 uint32 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 == 217821258) {
__t0 = gopurs_runtime.Str("Minute")
goto end_branch_0
} else {

}
}
{
if (v_0 == 3908053364) {
__t0 = gopurs_runtime.Str("Second")
goto end_branch_0
} else {

}
}
{
if (v_0 == 1292308612) {
__t0 = gopurs_runtime.Str("Hour")
goto end_branch_0
} else {

}
}
{
if (v_0 == 2311060696) {
__t0 = gopurs_runtime.Str("Day")
goto end_branch_0
} else {

}
}
{
if (v_0 == 401302776) {
__t0 = gopurs_runtime.Str("Week")
goto end_branch_0
} else {

}
}
{
if (v_0 == 3327533908) {
__t0 = gopurs_runtime.Str("Month")
goto end_branch_0
} else {

}
}
{
if (v_0 == 3631736139) {
__t0 = gopurs_runtime.Str("Year")
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0.StrVal()
}

func Call_show__2896747026(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[*pkg_Data_Map_Internal.Constructor_Node[uint32, float64]]](pkg_Data_Interval_Duration.Get_showMap()).V0, __eta0_0)
}

func Call_snd__2019004820(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[uint32, float64]) float64 {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[uint32, float64] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1.FloatVal()
}

func Call_unfoldr__2235715281(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__3990862552(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__1519733018(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



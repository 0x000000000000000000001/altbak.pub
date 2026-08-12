package Data_Interval_Duration_Iso

import (
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Interval_Duration "gopurs/output/Data.Interval.Duration"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
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
semigroupAdditive1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Float((v_0.FloatVal()) + (v1_1.FloatVal()))
})
}))
_ = semigroupAdditive1_0_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[float64]](gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_0_0
}), gopurs_runtime.Float(0.0))))}
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(f_0, a_2)), gopurs_runtime.Apply(g_1, a_2))
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
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(IsoDuration "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDuration(), "show"), v_0), gopurs_runtime.Str(")"))).StrVal())
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
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(ContainsNegativeValue "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDurationComponent(), "show"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_ContainsNegativeValue)(v_0.UnsafePtr).V0), UnsafePtr: nil}), gopurs_runtime.Str(")"))).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(InvalidFractionalUse "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDurationComponent(), "show"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_InvalidFractionalUse)(v_0.UnsafePtr).V0), UnsafePtr: nil}), gopurs_runtime.Str(")"))).StrVal())
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
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_eqDuration(), "eq"), x_0, y_1).IntVal) != (0))
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_ordDuration(), "compare"), x_0, y_1).IntVal)), UnsafePtr: nil}
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
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1422140417) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1775501833) {
var __t2 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(false)
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t3 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t3 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_eqDurationComponent(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_ContainsNegativeValue)(x_0.UnsafePtr).V0), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_ContainsNegativeValue)(y_1.UnsafePtr).V0), UnsafePtr: nil}).IntVal) != (0))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(false)
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 574232667)) && ((y_1.Type == 9 && y_1.IntVal == 574232667)) {
__t0 = gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_eqDurationComponent(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_InvalidFractionalUse)(x_0.UnsafePtr).V0), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_InvalidFractionalUse)(y_1.UnsafePtr).V0), UnsafePtr: nil}).IntVal) != (0))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(false)
}
end_branch_0:
return gopurs_runtime.Bool((__t0.IntVal) != (0))
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
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1422140417) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1775501833) {
var __t2 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t3 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_ordDurationComponent(), "compare"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_ContainsNegativeValue)(x_0.UnsafePtr).V0), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_ContainsNegativeValue)(y_1.UnsafePtr).V0), UnsafePtr: nil}).IntVal)), UnsafePtr: nil}
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}.IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 574232667)) && ((y_1.Type == 9 && y_1.IntVal == 574232667)) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_ordDurationComponent(), "compare"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_InvalidFractionalUse)(x_0.UnsafePtr).V0), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_InvalidFractionalUse)(y_1.UnsafePtr).V0), UnsafePtr: nil}).IntVal)), UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(__t0.IntVal)), UnsafePtr: nil}
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
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("Component `"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDurationComponent(), "show"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_ContainsNegativeValue)(v_0.UnsafePtr).V0), UnsafePtr: nil}), gopurs_runtime.Str("` contains negative value"))).StrVal())
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
__t0 = gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("Invalid usage of Fractional value at component `"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDurationComponent(), "show"), gopurs_runtime.Value{Type: 9, IntVal: int64((*Constructor_InvalidFractionalUse)(v_0.UnsafePtr).V0), UnsafePtr: nil}), gopurs_runtime.Str("`"))).StrVal())
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
var __t0 gopurs_runtime.Value
{
var go__go_1_1_0 gopurs_runtime.Value
go__go_1_1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__go_1_1_0:
for {
if false { continue go__go_1_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t2 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
v1_3_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_ordDurationComponent(), "compare"), gopurs_runtime.Value{Type: 9, IntVal: int64(401302776), UnsafePtr: nil}, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
_ = v1_3_3
var __t4 gopurs_runtime.Value
{
if (uint32(v1_3_3.IntVal) == 1527465420) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__go_1_1_0
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_3_3.IntVal) == 380165415) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
continue go__go_1_1_0
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{}))}
goto end_branch_4
} else {

}
}
{
if (uint32(v1_3_3.IntVal) == 902936544) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3})}))}
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t4))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t2))}
}
}()
})
__local_var_2_5 := gopurs_runtime.Apply(go__go_1_1_0, gopurs_runtime.RecordGet(v_0, "asMap"))
_ = __local_var_2_5
var __t6 gopurs_runtime.Value
{
if (__local_var_2_5.Type == 9 && __local_var_2_5.IntVal == 930809136 && __local_var_2_5.UnsafePtr == nil) {
__t6 = gopurs_runtime.Bool(false)
goto end_branch_6
} else {

}
}
{
if (__local_var_2_5.Type == 9 && __local_var_2_5.IntVal == 930809136 && __local_var_2_5.UnsafePtr != nil) {
__t6 = gopurs_runtime.Bool(true)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
var __t7 gopurs_runtime.Value
{
var __t8 gopurs_runtime.Value
{
var __t_tag_9 gopurs_runtime.Value = gopurs_runtime.RecordGet(v_0, "asMap")
if (__t_tag_9.Type == 9 && __t_tag_9.IntVal == 324739070 && __t_tag_9.UnsafePtr == nil) {
__t8 = gopurs_runtime.Bool(false)
goto end_branch_8
} else {

}
}
{
var __t_tag_10 gopurs_runtime.Value = gopurs_runtime.RecordGet(v_0, "asMap")
if (__t_tag_10.Type == 9 && __t_tag_10.IntVal == 324739070 && __t_tag_10.UnsafePtr != nil) {
__t8 = gopurs_runtime.Bool((gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.RecordGet(v_0, "asMap").UnsafePtr).V1).IntVal) > (1))
goto end_branch_8
} else {

}
}
{
__t8 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_8:
if (__t8.IntVal) != (0) {
__t7 = gopurs_runtime.Bool(true)
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Bool(false)
}
end_branch_7:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t6, __t7).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: unsafe.Pointer(nil)})))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_checkNegativeValues(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(Get_foldMap(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if ((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1.FloatVal()) < (0.0) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
if (__t1.IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Constructor_ContainsNegativeValue{1, uint32((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0.IntVal)})})))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0))}
}), gopurs_runtime.RecordGet(v_0, "asList")))
}

func Call_checkFractionalUse(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](Get_heytingAlgebraFunction()).V4, gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqNumber(), "eq"), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), a_1), a_1), gopurs_runtime.Bool(false)).IntVal) != (0))
}))
_ = __local_var_1_1
v1_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[*pkg_Data_Tuple.Constructor_Tuple[uint32, float64]]](gopurs_runtime.RecordGet(gopurs_runtime.Apply2(pkg_Data_List.Get_span(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1)
}), gopurs_runtime.RecordGet(v_0, "asList")), "rest"))
_ = v1_1_0
var __t2 gopurs_runtime.Value
{
var __t_and_4 bool = false
if (gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.IntVal == 1358893437 && gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr != nil) {

var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[float64]](Get_monoidAdditive()))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr).V1)}).FloatVal()) > (0.0) {
__t3 = gopurs_runtime.Bool(true)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Bool(false)
}
end_branch_3:
__t_and_4 = (__t3.IntVal) != (0)
}
if __t_and_4 {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Constructor_InvalidFractionalUse{1, uint32((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v1_1_0)}.UnsafePtr).V0.UnsafePtr).V0.IntVal)})})))}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")))}
}
end_branch_2:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t2)
}

func Call_checkEmptiness(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.RecordGet(v_0, "asList")
if (__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(false)
}
end_branch_1:
if (__t1.IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: unsafe.Pointer(nil)})))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](__t0)
}

func Call_checkValidIsoDuration(v_0_loop *pkg_Data_Map_Internal.Constructor_Node[uint32, float64]) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 *pkg_Data_Map_Internal.Constructor_Node[uint32, float64] = v_0_loop
_ = v_0
var go__go_1_0_1 gopurs_runtime.Value
go__go_1_0_1 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](v_2_loop_val)
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__go_1_0_1:
for {
if false { continue go__go_1_0_1 }
var v_2 *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(v_2)}
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, v_2})})
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
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
return gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply4(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](Get_monoidFn()))}, pkg_Data_Foldable.Get_identity1(), gopurs_runtime.Array([]gopurs_runtime.Value{Get_checkWeekUsage(), Get_checkEmptiness(), Get_checkFractionalUse(), Get_checkNegativeValues()}), gopurs_runtime.RecordDict2("asList", "asMap", gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]](gopurs_runtime.Apply2(go__go_1_0_1, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_unfoldableList(), "unfoldr"), pkg_Data_Map_Internal.Get_stepUnfoldr(), gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.CoerceToStruct[pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)}), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: unsafe.Pointer(nil)}})}))))}, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(v_0)})))
}

func Call_mkIsoDuration(d_0_loop *pkg_Data_Map_Internal.Constructor_Node[uint32, float64]) gopurs_runtime.Value {
var d_0 *pkg_Data_Map_Internal.Constructor_Node[uint32, float64] = d_0_loop
_ = d_0
__local_var_1_1 := gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_checkValidIsoDuration(d_0))}
_ = __local_var_1_1
var __t2 gopurs_runtime.Value
{
if (__local_var_1_1.Type == 9 && __local_var_1_1.IntVal == 1358893437 && __local_var_1_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
goto end_branch_2
} else {

}
}
{
if (__local_var_1_1.Type == 9 && __local_var_1_1.IntVal == 1358893437 && __local_var_1_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(__local_var_1_1.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(__local_var_1_1.UnsafePtr).V1)}})}})}))}
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
v_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_NonEmpty.Constructor_NonEmpty[*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value], gopurs_runtime.Value]]](__t2)
_ = v_1_0
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr != nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr).V0})}
goto end_branch_3
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1_0)}.UnsafePtr == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer(d_0)}})}
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



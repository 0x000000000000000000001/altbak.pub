package Data_Interval_Duration_Iso

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_List_Types "gopurs/output/Data.List.Types"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Map_Internal "gopurs/output/Data.Map.Internal"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Interval_Duration "gopurs/output/Data.Interval.Duration"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_List "gopurs/output/Data.List"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Number "gopurs/output/Data.Number"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_NonEmpty "gopurs/output/Data.NonEmpty"
	unsafe "unsafe"
)

var cache_lookup gopurs_runtime.Value
var once_lookup sync.Once
func Get_lookup() gopurs_runtime.Value {
	once_lookup.Do(func() {
		cache_lookup = gopurs_runtime.Func(func(k_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_lookup(k_0_box)
})
	})
	return cache_lookup
}

var cache_greaterThan gopurs_runtime.Value
var once_greaterThan sync.Once
func Get_greaterThan() gopurs_runtime.Value {
	once_greaterThan.Do(func() {
		cache_greaterThan = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
_ = __local_var_0_0
return gopurs_runtime.Func2(func(a1_1 gopurs_runtime.Value, a2_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.Apply2(__local_var_0_0, a1_1, a2_2)
return gopurs_runtime.Bool((__t_tag_1.Type == 9 && __t_tag_1.IntVal == 380165415))
})
}()
	})
	return cache_greaterThan
}

var cache_foldMap1 gopurs_runtime.Value
var once_foldMap1 sync.Once
func Get_foldMap1() gopurs_runtime.Value {
	once_foldMap1.Do(func() {
		cache_foldMap1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), pkg_Data_List_Types.Get_monoidList())
	})
	return cache_foldMap1
}

var cache_foldMap2 gopurs_runtime.Value
var once_foldMap2 sync.Once
func Get_foldMap2() gopurs_runtime.Value {
	once_foldMap2.Do(func() {
		cache_foldMap2 = func() gopurs_runtime.Value {
semigroupAdditive1_0_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "add"), v_0, v1_1)
}))
_ = semigroupAdditive1_0_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_foldableList(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupAdditive1_0_0
}), gopurs_runtime.RecordGet(pkg_Data_Semiring.Get_semiringNumber(), "zero")))
}()
	})
	return cache_foldMap2
}

var cache_fold gopurs_runtime.Value
var once_fold sync.Once
func Get_fold() gopurs_runtime.Value {
	once_fold.Do(func() {
		cache_fold = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_monoidList(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_0_0
semigroupFn_1_1 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, g_2 gopurs_runtime.Value, x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_0_0, "append"), gopurs_runtime.Apply(f_1, x_3), gopurs_runtime.Apply(g_2, x_3))
}))
_ = semigroupFn_1_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldMap"), gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupFn_1_1
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_monoidList(), "mempty")
})), pkg_Data_Foldable.Get_identity())
}()
	})
	return cache_fold
}

var cache_toUnfoldable gopurs_runtime.Value
var once_toUnfoldable sync.Once
func Get_toUnfoldable() gopurs_runtime.Value {
	once_toUnfoldable.Do(func() {
		cache_toUnfoldable = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_unfoldableList(), "unfoldr"), pkg_Data_Map_Internal.Get_stepUnfoldr())
_ = __local_var_0_0
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{Type: 9, IntVal: 2861335956, UnsafePtr: unsafe.Pointer(&pkg_Data_Map_Internal.Constructor_IterNode[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr), gopurs_runtime.Value{Type: 9, IntVal: 2509360378, UnsafePtr: nil}})})
})
}()
	})
	return cache_toUnfoldable
}

var cache_IsEmpty gopurs_runtime.Value
var once_IsEmpty sync.Once
func Get_IsEmpty() gopurs_runtime.Value {
	once_IsEmpty.Do(func() {
		cache_IsEmpty = gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: nil}
	})
	return cache_IsEmpty
}

var cache_InvalidWeekComponentUsage gopurs_runtime.Value
var once_InvalidWeekComponentUsage sync.Once
func Get_InvalidWeekComponentUsage() gopurs_runtime.Value {
	once_InvalidWeekComponentUsage.Do(func() {
		cache_InvalidWeekComponentUsage = gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: nil}
	})
	return cache_InvalidWeekComponentUsage
}

var cache_ContainsNegativeValue gopurs_runtime.Value
var once_ContainsNegativeValue sync.Once
func Get_ContainsNegativeValue() gopurs_runtime.Value {
	once_ContainsNegativeValue.Do(func() {
		cache_ContainsNegativeValue = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Constructor_ContainsNegativeValue{value0})}
})
	})
	return cache_ContainsNegativeValue
}

var cache_InvalidFractionalUse gopurs_runtime.Value
var once_InvalidFractionalUse sync.Once
func Get_InvalidFractionalUse() gopurs_runtime.Value {
	once_InvalidFractionalUse.Do(func() {
		cache_InvalidFractionalUse = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Constructor_InvalidFractionalUse{value0})}
})
	})
	return cache_InvalidFractionalUse
}

var cache_unIsoDuration gopurs_runtime.Value
var once_unIsoDuration sync.Once
func Get_unIsoDuration() gopurs_runtime.Value {
	once_unIsoDuration.Do(func() {
		cache_unIsoDuration = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unIsoDuration(v_0_box)
})
	})
	return cache_unIsoDuration
}

var cache_showIsoDuration gopurs_runtime.Value
var once_showIsoDuration sync.Once
func Get_showIsoDuration() gopurs_runtime.Value {
	once_showIsoDuration.Do(func() {
		cache_showIsoDuration = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(IsoDuration "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDuration(), "show"), v_0), gopurs_runtime.Str(")")))
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
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(ContainsNegativeValue "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDurationComponent(), "show"), (*Constructor_ContainsNegativeValue)(v_0.UnsafePtr).V0), gopurs_runtime.Str(")")))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(InvalidFractionalUse "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDurationComponent(), "show"), (*Constructor_InvalidFractionalUse)(v_0.UnsafePtr).V0), gopurs_runtime.Str(")")))
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
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
		cache_eqIsoDuration = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_eqDuration(), "eq"), x_0, y_1)
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
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_ordDuration(), "compare"), x_0, y_1)
}))
	})
	return cache_ordIsoDuration
}

var cache_eqError gopurs_runtime.Value
var once_eqError sync.Once
func Get_eqError() gopurs_runtime.Value {
	once_eqError.Do(func() {
		cache_eqError = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1422140417) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 1422140417))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1775501833) {
__t0 = gopurs_runtime.Bool((y_1.Type == 9 && y_1.IntVal == 1775501833))
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
__t0 = gopurs_runtime.Bool(((y_1.Type == 9 && y_1.IntVal == 3224543173)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_eqDurationComponent(), "eq"), (*Constructor_ContainsNegativeValue)(x_0.UnsafePtr).V0, (*Constructor_ContainsNegativeValue)(y_1.UnsafePtr).V0).IntVal) != (0)))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Bool(((x_0.Type == 9 && x_0.IntVal == 574232667)) && (((y_1.Type == 9 && y_1.IntVal == 574232667)) && ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_eqDurationComponent(), "eq"), (*Constructor_InvalidFractionalUse)(x_0.UnsafePtr).V0, (*Constructor_InvalidFractionalUse)(y_1.UnsafePtr).V0).IntVal) != (0))))
}
end_branch_0:
return __t0
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
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (x_0.Type == 9 && x_0.IntVal == 1422140417) {
var __t1 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_1:
__t0 = __t1
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1422140417) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 1775501833) {
var __t2 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_2:
__t0 = __t2
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 1775501833) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if (x_0.Type == 9 && x_0.IntVal == 3224543173) {
var __t3 gopurs_runtime.Value
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t3 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_ordDurationComponent(), "compare"), (*Constructor_ContainsNegativeValue)(x_0.UnsafePtr).V0, (*Constructor_ContainsNegativeValue)(y_1.UnsafePtr).V0)
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
}
end_branch_3:
__t0 = __t3
goto end_branch_0
} else {

}
}
{
if (y_1.Type == 9 && y_1.IntVal == 3224543173) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
if ((x_0.Type == 9 && x_0.IntVal == 574232667)) && ((y_1.Type == 9 && y_1.IntVal == 574232667)) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_ordDurationComponent(), "compare"), (*Constructor_InvalidFractionalUse)(x_0.UnsafePtr).V0, (*Constructor_InvalidFractionalUse)(y_1.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
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
return gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_checkValidIsoDuration(v_0_box))}
})
	})
	return cache_checkValidIsoDuration
}

var cache_mkIsoDuration gopurs_runtime.Value
var once_mkIsoDuration sync.Once
func Get_mkIsoDuration() gopurs_runtime.Value {
	once_mkIsoDuration.Do(func() {
		cache_mkIsoDuration = gopurs_runtime.Func(func(d_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mkIsoDuration(d_0_box)
})
	})
	return cache_mkIsoDuration
}

type Constructor_IsEmpty struct {
	
}


type Constructor_InvalidWeekComponentUsage struct {
	
}


type Constructor_ContainsNegativeValue struct {
	V0 gopurs_runtime.Value
}


type Constructor_InvalidFractionalUse struct {
	V0 gopurs_runtime.Value
}


func Call_lookup(k_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var k_0 gopurs_runtime.Value = k_0_loop
_ = k_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var __t1 gopurs_runtime.Value
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
if (v_2.Type == 9 && v_2.IntVal == 324739070 && v_2.UnsafePtr != nil) {
v1_3_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_ordDurationComponent(), "compare"), k_0, (*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2)
_ = v1_3_2
var __t3 gopurs_runtime.Value
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 1527465420) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V4)}
continue go__1_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 380165415) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 324739070, UnsafePtr: unsafe.Pointer((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V5)}
continue go__1_0
__t3 = gopurs_runtime.Value{}
goto end_branch_3
} else {

}
}
{
if (v1_3_2.Type == 9 && v1_3_2.IntVal == 902936544) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{(*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V3})}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t1 = __t3
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
return go__1_0
}

func Call_unIsoDuration(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
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
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("Component `"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDurationComponent(), "show"), (*Constructor_ContainsNegativeValue)(v_0.UnsafePtr).V0), gopurs_runtime.Str("` contains negative value")))
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 574232667) {
__t0 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("Invalid usage of Fractional value at component `"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Interval_Duration.Get_showDurationComponent(), "show"), (*Constructor_InvalidFractionalUse)(v_0.UnsafePtr).V0), gopurs_runtime.Str("`")))
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
__local_var_1_1 := gopurs_runtime.Apply(Call_lookup(gopurs_runtime.Value{Type: 9, IntVal: 401302776, UnsafePtr: nil}), gopurs_runtime.RecordGet(v_0, "asMap"))
_ = __local_var_1_1
var __t2 gopurs_runtime.Value
{
if (__local_var_1_1.Type == 9 && __local_var_1_1.IntVal == 930809136 && __local_var_1_1.UnsafePtr == nil) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
if (__local_var_1_1.Type == 9 && __local_var_1_1.IntVal == 930809136 && __local_var_1_1.UnsafePtr != nil) {
__t2 = gopurs_runtime.Bool(true)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
var __t3 gopurs_runtime.Value
{
var __t_tag_4 gopurs_runtime.Value = gopurs_runtime.RecordGet(v_0, "asMap")
if (__t_tag_4.Type == 9 && __t_tag_4.IntVal == 324739070 && __t_tag_4.UnsafePtr == nil) {
__t3 = gopurs_runtime.Int(0)
goto end_branch_3
} else {

}
}
{
var __t_tag_5 gopurs_runtime.Value = gopurs_runtime.RecordGet(v_0, "asMap")
if (__t_tag_5.Type == 9 && __t_tag_5.IntVal == 324739070 && __t_tag_5.UnsafePtr != nil) {
__t3 = gopurs_runtime.Int((*pkg_Data_Map_Internal.Constructor_Node[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.RecordGet(v_0, "asMap").UnsafePtr).V1)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t2, gopurs_runtime.Apply2(Get_greaterThan(), __t3, gopurs_runtime.Int(1))).IntVal) != (0) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1775501833, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")
}
end_branch_0:
return (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr)
}

func Call_checkNegativeValues(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply2(Get_foldMap1(), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if ((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1.FloatVal()) >= (0.0) {
__t0 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 3224543173, UnsafePtr: unsafe.Pointer(&Constructor_ContainsNegativeValue{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V0})})
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordGet(v_0, "asList")).UnsafePtr)
}

func Call_checkFractionalUse(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply2(pkg_Data_List.Get_span(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqNumber(), "eq"), gopurs_runtime.Apply(pkg_Data_Number.Get_floor(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_1.UnsafePtr).V1), gopurs_runtime.Bool(false)))
}), gopurs_runtime.RecordGet(v_0, "asList"))
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
var __t_tag_2 gopurs_runtime.Value = gopurs_runtime.RecordGet(__local_var_1_0, "rest")
if ((__t_tag_2.Type == 9 && __t_tag_2.IntVal == 1358893437 && __t_tag_2.UnsafePtr != nil)) && ((gopurs_runtime.Apply2(Get_foldMap2(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Number.Get_abs(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1)
}), gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(__local_var_1_0, "rest").UnsafePtr).V1)}).FloatVal()) > (0.0)) {
__t1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 574232667, UnsafePtr: unsafe.Pointer(&Constructor_InvalidFractionalUse{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.RecordGet(__local_var_1_0, "rest").UnsafePtr).V0.UnsafePtr).V0})})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")
}
end_branch_1:
return (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(__t1.UnsafePtr)
}

func Call_checkEmptiness(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
var __t_tag_1 gopurs_runtime.Value = gopurs_runtime.RecordGet(v_0, "asList")
if (__t_tag_1.Type == 9 && __t_tag_1.IntVal == 1358893437 && __t_tag_1.UnsafePtr == nil) {
__t0 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_applicativeList(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 1422140417, UnsafePtr: nil})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordGet(pkg_Data_List_Types.Get_plusList(), "empty")
}
end_branch_0:
return (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(__t0.UnsafePtr)
}

func Call_checkValidIsoDuration(v_0_loop gopurs_runtime.Value) *pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var go__1_0 gopurs_runtime.Value
go__1_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_2_loop gopurs_runtime.Value = v_2_loop_val
var v1_3_loop gopurs_runtime.Value = v1_3_loop_val
go__1_0:
for {
if false { continue go__1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var __t1 gopurs_runtime.Value
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr == nil) {
__t1 = v_2
goto end_branch_1
} else {

}
}
{
if (v1_3.Type == 9 && v1_3.IntVal == 1358893437 && v1_3.UnsafePtr != nil) {
v_2_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(&pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value]{(*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V0, (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v_2.UnsafePtr)})}
v1_3_loop = gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(v1_3.UnsafePtr).V1)}
continue go__1_0
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
return (*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(gopurs_runtime.Apply2(Get_fold(), gopurs_runtime.Array([]gopurs_runtime.Value{Get_checkWeekUsage(), Get_checkEmptiness(), Get_checkFractionalUse(), Get_checkNegativeValues()}), gopurs_runtime.RecordDict2("asList", "asMap", gopurs_runtime.Apply2(go__1_0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: nil}, gopurs_runtime.Apply(Get_toUnfoldable(), v_0)), v_0)).UnsafePtr)
}

func Call_mkIsoDuration(d_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var d_0 gopurs_runtime.Value = d_0_loop
_ = d_0
__local_var_1_0 := gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer(Call_checkValidIsoDuration(d_0))}
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 1358893437 && __local_var_1_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value]{d_0})}
goto end_branch_1
} else {

}
}
{
if (__local_var_1_0.Type == 9 && __local_var_1_0.IntVal == 1358893437 && __local_var_1_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&pkg_Data_NonEmpty.Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 1358893437, UnsafePtr: unsafe.Pointer((*pkg_Data_List_Types.Constructor_Cons[gopurs_runtime.Value])(__local_var_1_0.UnsafePtr).V1)}})}})}
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



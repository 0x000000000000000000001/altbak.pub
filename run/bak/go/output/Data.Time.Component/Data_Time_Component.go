package Data_Time_Component

import (
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_showSecond gopurs_runtime.Value
var once_showSecond sync.Once
func Get_showSecond() gopurs_runtime.Value {
	once_showSecond.Do(func() {
		cache_showSecond = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Second "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), v_0), gopurs_runtime.Str(")"))).StrVal())
}))
	})
	return cache_showSecond
}

var cache_showMinute gopurs_runtime.Value
var once_showMinute sync.Once
func Get_showMinute() gopurs_runtime.Value {
	once_showMinute.Do(func() {
		cache_showMinute = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Minute "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), v_0), gopurs_runtime.Str(")"))).StrVal())
}))
	})
	return cache_showMinute
}

var cache_showMillisecond gopurs_runtime.Value
var once_showMillisecond sync.Once
func Get_showMillisecond() gopurs_runtime.Value {
	once_showMillisecond.Do(func() {
		cache_showMillisecond = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Millisecond "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), v_0), gopurs_runtime.Str(")"))).StrVal())
}))
	})
	return cache_showMillisecond
}

var cache_showHour gopurs_runtime.Value
var once_showHour sync.Once
func Get_showHour() gopurs_runtime.Value {
	once_showHour.Do(func() {
		cache_showHour = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(Hour "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), v_0), gopurs_runtime.Str(")"))).StrVal())
}))
	})
	return cache_showHour
}

var cache_ordSecond gopurs_runtime.Value
var once_ordSecond sync.Once
func Get_ordSecond() gopurs_runtime.Value {
	once_ordSecond.Do(func() {
		cache_ordSecond = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_ordSecond
}

var cache_ordMinute gopurs_runtime.Value
var once_ordMinute sync.Once
func Get_ordMinute() gopurs_runtime.Value {
	once_ordMinute.Do(func() {
		cache_ordMinute = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_ordMinute
}

var cache_ordMillisecond gopurs_runtime.Value
var once_ordMillisecond sync.Once
func Get_ordMillisecond() gopurs_runtime.Value {
	once_ordMillisecond.Do(func() {
		cache_ordMillisecond = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_ordMillisecond
}

var cache_ordHour gopurs_runtime.Value
var once_ordHour sync.Once
func Get_ordHour() gopurs_runtime.Value {
	once_ordHour.Do(func() {
		cache_ordHour = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
}), gopurs_runtime.Apply3(pkg_Data_Ord.Get_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}))
	})
	return cache_ordHour
}

var cache_eqSecond gopurs_runtime.Value
var once_eqSecond sync.Once
func Get_eqSecond() gopurs_runtime.Value {
	once_eqSecond.Do(func() {
		cache_eqSecond = gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
	})
	return cache_eqSecond
}

var cache_eqMinute gopurs_runtime.Value
var once_eqMinute sync.Once
func Get_eqMinute() gopurs_runtime.Value {
	once_eqMinute.Do(func() {
		cache_eqMinute = gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
	})
	return cache_eqMinute
}

var cache_eqMillisecond gopurs_runtime.Value
var once_eqMillisecond sync.Once
func Get_eqMillisecond() gopurs_runtime.Value {
	once_eqMillisecond.Do(func() {
		cache_eqMillisecond = gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
	})
	return cache_eqMillisecond
}

var cache_eqHour gopurs_runtime.Value
var once_eqHour sync.Once
func Get_eqHour() gopurs_runtime.Value {
	once_eqHour.Do(func() {
		cache_eqHour = gopurs_runtime.RecordDict1("eq", pkg_Data_Eq.Get_eqIntImpl())
	})
	return cache_eqHour
}

var cache_boundedSecond gopurs_runtime.Value
var once_boundedSecond sync.Once
func Get_boundedSecond() gopurs_runtime.Value {
	once_boundedSecond.Do(func() {
		cache_boundedSecond = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordSecond()
}), gopurs_runtime.Int(0), gopurs_runtime.Int(59))
	})
	return cache_boundedSecond
}

var cache_boundedMinute gopurs_runtime.Value
var once_boundedMinute sync.Once
func Get_boundedMinute() gopurs_runtime.Value {
	once_boundedMinute.Do(func() {
		cache_boundedMinute = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordMinute()
}), gopurs_runtime.Int(0), gopurs_runtime.Int(59))
	})
	return cache_boundedMinute
}

var cache_boundedMillisecond gopurs_runtime.Value
var once_boundedMillisecond sync.Once
func Get_boundedMillisecond() gopurs_runtime.Value {
	once_boundedMillisecond.Do(func() {
		cache_boundedMillisecond = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordMillisecond()
}), gopurs_runtime.Int(0), gopurs_runtime.Int(999))
	})
	return cache_boundedMillisecond
}

var cache_boundedHour gopurs_runtime.Value
var once_boundedHour sync.Once
func Get_boundedHour() gopurs_runtime.Value {
	once_boundedHour.Do(func() {
		cache_boundedHour = gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordHour()
}), gopurs_runtime.Int(0), gopurs_runtime.Int(23))
	})
	return cache_boundedHour
}

var cache_boundedEnumSecond gopurs_runtime.Value
var once_boundedEnumSecond sync.Once
func Get_boundedEnumSecond() gopurs_runtime.Value {
	once_boundedEnumSecond.Do(func() {
		cache_boundedEnumSecond = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedSecond()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumSecond()
}), gopurs_runtime.Int(60), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (n_0.IntVal) < (0) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (n_0.IntVal) > (59) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(true)
}
end_branch_2:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t1, __t2).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, n_0})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0))}
}))
	})
	return cache_boundedEnumSecond
}

var cache_enumSecond gopurs_runtime.Value
var once_enumSecond sync.Once
func Get_enumSecond() gopurs_runtime.Value {
	once_enumSecond.Do(func() {
		cache_enumSecond = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordSecond()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumSecond(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumSecond(), "fromEnum"), x_0).IntVal) - (1)))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumSecond(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumSecond(), "fromEnum"), x_0).IntVal) + (1)))
}))
	})
	return cache_enumSecond
}

var cache_boundedEnumMinute gopurs_runtime.Value
var once_boundedEnumMinute sync.Once
func Get_boundedEnumMinute() gopurs_runtime.Value {
	once_boundedEnumMinute.Do(func() {
		cache_boundedEnumMinute = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedMinute()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumMinute()
}), gopurs_runtime.Int(60), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (n_0.IntVal) < (0) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (n_0.IntVal) > (59) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(true)
}
end_branch_2:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t1, __t2).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, n_0})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0))}
}))
	})
	return cache_boundedEnumMinute
}

var cache_enumMinute gopurs_runtime.Value
var once_enumMinute sync.Once
func Get_enumMinute() gopurs_runtime.Value {
	once_enumMinute.Do(func() {
		cache_enumMinute = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordMinute()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumMinute(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumMinute(), "fromEnum"), x_0).IntVal) - (1)))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumMinute(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumMinute(), "fromEnum"), x_0).IntVal) + (1)))
}))
	})
	return cache_enumMinute
}

var cache_boundedEnumMillisecond gopurs_runtime.Value
var once_boundedEnumMillisecond sync.Once
func Get_boundedEnumMillisecond() gopurs_runtime.Value {
	once_boundedEnumMillisecond.Do(func() {
		cache_boundedEnumMillisecond = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedMillisecond()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumMillisecond()
}), gopurs_runtime.Int(1000), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (n_0.IntVal) < (0) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (n_0.IntVal) > (999) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(true)
}
end_branch_2:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t1, __t2).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, n_0})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0))}
}))
	})
	return cache_boundedEnumMillisecond
}

var cache_enumMillisecond gopurs_runtime.Value
var once_enumMillisecond sync.Once
func Get_enumMillisecond() gopurs_runtime.Value {
	once_enumMillisecond.Do(func() {
		cache_enumMillisecond = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordMillisecond()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumMillisecond(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumMillisecond(), "fromEnum"), x_0).IntVal) - (1)))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumMillisecond(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumMillisecond(), "fromEnum"), x_0).IntVal) + (1)))
}))
	})
	return cache_enumMillisecond
}

var cache_boundedEnumHour gopurs_runtime.Value
var once_boundedEnumHour sync.Once
func Get_boundedEnumHour() gopurs_runtime.Value {
	once_boundedEnumHour.Do(func() {
		cache_boundedEnumHour = gopurs_runtime.RecordDict5("Bounded0", "Enum1", "cardinality", "fromEnum", "toEnum", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_boundedHour()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_enumHour()
}), gopurs_runtime.Int(24), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v_0.IntVal)
}), gopurs_runtime.Func(func(n_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
var __t1 gopurs_runtime.Value
{
if (n_0.IntVal) < (0) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Bool(true)
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (n_0.IntVal) > (23) {
__t2 = gopurs_runtime.Bool(false)
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Bool(true)
}
end_branch_2:
if (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), __t1, __t2).IntVal) != (0) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, n_0})}))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](__t0))}
}))
	})
	return cache_boundedEnumHour
}

var cache_enumHour gopurs_runtime.Value
var once_enumHour sync.Once
func Get_enumHour() gopurs_runtime.Value {
	once_enumHour.Do(func() {
		cache_enumHour = gopurs_runtime.RecordDict3("Ord0", "pred", "succ", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_ordHour()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumHour(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumHour(), "fromEnum"), x_0).IntVal) - (1)))
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumHour(), "toEnum"), gopurs_runtime.Int((gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_boundedEnumHour(), "fromEnum"), x_0).IntVal) + (1)))
}))
	})
	return cache_enumHour
}





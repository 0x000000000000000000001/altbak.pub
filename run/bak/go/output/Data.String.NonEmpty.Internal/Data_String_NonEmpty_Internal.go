package Data_String_NonEmpty_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	unsafe "unsafe"
)

var cache_NonEmptyString gopurs_runtime.Value
var once_NonEmptyString sync.Once
func Get_NonEmptyString() gopurs_runtime.Value {
	once_NonEmptyString.Do(func() {
		cache_NonEmptyString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_NonEmptyString(x_0_box)
})
	})
	return cache_NonEmptyString
}

var cache_NonEmptyReplacement gopurs_runtime.Value
var once_NonEmptyReplacement sync.Once
func Get_NonEmptyReplacement() gopurs_runtime.Value {
	once_NonEmptyReplacement.Do(func() {
		cache_NonEmptyReplacement = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_NonEmptyReplacement(x_0_box)
})
	})
	return cache_NonEmptyReplacement
}

var cache_toUpper gopurs_runtime.Value
var once_toUpper sync.Once
func Get_toUpper() gopurs_runtime.Value {
	once_toUpper.Do(func() {
		cache_toUpper = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toUpper(v_0_box)
})
	})
	return cache_toUpper
}

var cache_toString gopurs_runtime.Value
var once_toString sync.Once
func Get_toString() gopurs_runtime.Value {
	once_toString.Do(func() {
		cache_toString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toString(v_0_box)
})
	})
	return cache_toString
}

var cache_toLower gopurs_runtime.Value
var once_toLower sync.Once
func Get_toLower() gopurs_runtime.Value {
	once_toLower.Do(func() {
		cache_toLower = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_toLower(v_0_box)
})
	})
	return cache_toLower
}

var cache_showNonEmptyString gopurs_runtime.Value
var once_showNonEmptyString sync.Once
func Get_showNonEmptyString() gopurs_runtime.Value {
	once_showNonEmptyString.Do(func() {
		cache_showNonEmptyString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(NonEmptyString.unsafeFromString "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showString(), "show"), v_0), gopurs_runtime.Str(")")))
}))
	})
	return cache_showNonEmptyString
}

var cache_showNonEmptyReplacement gopurs_runtime.Value
var once_showNonEmptyReplacement sync.Once
func Get_showNonEmptyReplacement() gopurs_runtime.Value {
	once_showNonEmptyReplacement.Do(func() {
		cache_showNonEmptyReplacement = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(NonEmptyReplacement "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_showNonEmptyString(), "show"), v_0), gopurs_runtime.Str(")")))
}))
	})
	return cache_showNonEmptyReplacement
}

var cache_semigroupNonEmptyString gopurs_runtime.Value
var once_semigroupNonEmptyString sync.Once
func Get_semigroupNonEmptyString() gopurs_runtime.Value {
	once_semigroupNonEmptyString.Do(func() {
		cache_semigroupNonEmptyString = pkg_Data_Semigroup.Get_semigroupString()
	})
	return cache_semigroupNonEmptyString
}

var cache_semigroupNonEmptyReplacement gopurs_runtime.Value
var once_semigroupNonEmptyReplacement sync.Once
func Get_semigroupNonEmptyReplacement() gopurs_runtime.Value {
	once_semigroupNonEmptyReplacement.Do(func() {
		cache_semigroupNonEmptyReplacement = pkg_Data_Semigroup.Get_semigroupString()
	})
	return cache_semigroupNonEmptyReplacement
}

var cache_replaceAll gopurs_runtime.Value
var once_replaceAll sync.Once
func Get_replaceAll() gopurs_runtime.Value {
	once_replaceAll.Do(func() {
		cache_replaceAll = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replaceAll(pat_0_box, v_1_box, v1_2_box)
})
	})
	return cache_replaceAll
}

var cache_replace gopurs_runtime.Value
var once_replace sync.Once
func Get_replace() gopurs_runtime.Value {
	once_replace.Do(func() {
		cache_replace = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replace(pat_0_box, v_1_box, v1_2_box)
})
	})
	return cache_replace
}

var cache_prependString gopurs_runtime.Value
var once_prependString sync.Once
func Get_prependString() gopurs_runtime.Value {
	once_prependString.Do(func() {
		cache_prependString = gopurs_runtime.Func2(func(s1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_prependString(s1_0_box.StrVal(), v_1_box)
})
	})
	return cache_prependString
}

var cache_ordNonEmptyString gopurs_runtime.Value
var once_ordNonEmptyString sync.Once
func Get_ordNonEmptyString() gopurs_runtime.Value {
	once_ordNonEmptyString.Do(func() {
		cache_ordNonEmptyString = pkg_Data_Ord.Get_ordString()
	})
	return cache_ordNonEmptyString
}

var cache_ordNonEmptyReplacement gopurs_runtime.Value
var once_ordNonEmptyReplacement sync.Once
func Get_ordNonEmptyReplacement() gopurs_runtime.Value {
	once_ordNonEmptyReplacement.Do(func() {
		cache_ordNonEmptyReplacement = pkg_Data_Ord.Get_ordString()
	})
	return cache_ordNonEmptyReplacement
}

var cache_nonEmptyNonEmpty gopurs_runtime.Value
var once_nonEmptyNonEmpty sync.Once
func Get_nonEmptyNonEmpty() gopurs_runtime.Value {
	once_nonEmptyNonEmpty.Do(func() {
		cache_nonEmptyNonEmpty = gopurs_runtime.Func(func(dictIsSymbol_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nonEmptyNonEmpty(dictIsSymbol_0_box)
})
	})
	return cache_nonEmptyNonEmpty
}

var cache_nes gopurs_runtime.Value
var once_nes sync.Once
func Get_nes() gopurs_runtime.Value {
	once_nes.Do(func() {
		cache_nes = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_nes(dict_0_box)
})
	})
	return cache_nes
}

var cache_makeNonEmptyBad gopurs_runtime.Value
var once_makeNonEmptyBad sync.Once
func Get_makeNonEmptyBad() gopurs_runtime.Value {
	once_makeNonEmptyBad.Do(func() {
		cache_makeNonEmptyBad = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_makeNonEmptyBad(_dollar__unused_0_box)
})
	})
	return cache_makeNonEmptyBad
}

var cache_localeCompare gopurs_runtime.Value
var once_localeCompare sync.Once
func Get_localeCompare() gopurs_runtime.Value {
	once_localeCompare.Do(func() {
		cache_localeCompare = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_localeCompare(v_0_box, v1_1_box)
})
	})
	return cache_localeCompare
}

var cache_liftS gopurs_runtime.Value
var once_liftS sync.Once
func Get_liftS() gopurs_runtime.Value {
	once_liftS.Do(func() {
		cache_liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftS(f_0_box, v_1_box)
})
	})
	return cache_liftS
}

var cache_startsWith gopurs_runtime.Value
var once_startsWith sync.Once
func Get_startsWith() gopurs_runtime.Value {
	once_startsWith.Do(func() {
		cache_startsWith = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_startsWith(x_0_box, v_1_box)
})
	})
	return cache_startsWith
}

var cache_joinWith1 gopurs_runtime.Value
var once_joinWith1 sync.Once
func Get_joinWith1() gopurs_runtime.Value {
	once_joinWith1.Do(func() {
		cache_joinWith1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinWith1(dictFoldable1_0_box)
})
	})
	return cache_joinWith1
}

var cache_joinWith gopurs_runtime.Value
var once_joinWith sync.Once
func Get_joinWith() gopurs_runtime.Value {
	once_joinWith.Do(func() {
		cache_joinWith = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinWith(dictFoldable_0_box)
})
	})
	return cache_joinWith
}

var cache_join1With gopurs_runtime.Value
var once_join1With sync.Once
func Get_join1With() gopurs_runtime.Value {
	once_join1With.Do(func() {
		cache_join1With = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_join1With(dictFoldable1_0_box)
})
	})
	return cache_join1With
}

var cache_fromString gopurs_runtime.Value
var once_fromString sync.Once
func Get_fromString() gopurs_runtime.Value {
	once_fromString.Do(func() {
		cache_fromString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromString(v_0_box.StrVal())
})
	})
	return cache_fromString
}

var cache_stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		cache_stripPrefix = gopurs_runtime.Func2(func(pat_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripPrefix(pat_0_box, a_1_box)
})
	})
	return cache_stripPrefix
}

var cache_stripSuffix gopurs_runtime.Value
var once_stripSuffix sync.Once
func Get_stripSuffix() gopurs_runtime.Value {
	once_stripSuffix.Do(func() {
		cache_stripSuffix = gopurs_runtime.Func2(func(pat_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripSuffix(pat_0_box, a_1_box)
})
	})
	return cache_stripSuffix
}

var cache_trim gopurs_runtime.Value
var once_trim sync.Once
func Get_trim() gopurs_runtime.Value {
	once_trim.Do(func() {
		cache_trim = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_trim(v_0_box)
})
	})
	return cache_trim
}

var cache_unsafeFromString gopurs_runtime.Value
var once_unsafeFromString sync.Once
func Get_unsafeFromString() gopurs_runtime.Value {
	once_unsafeFromString.Do(func() {
		cache_unsafeFromString = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeFromString(_dollar__unused_0_box, x_1_box.StrVal())
})
	})
	return cache_unsafeFromString
}

var cache_eqNonEmptyString gopurs_runtime.Value
var once_eqNonEmptyString sync.Once
func Get_eqNonEmptyString() gopurs_runtime.Value {
	once_eqNonEmptyString.Do(func() {
		cache_eqNonEmptyString = pkg_Data_Eq.Get_eqString()
	})
	return cache_eqNonEmptyString
}

var cache_eqNonEmptyReplacement gopurs_runtime.Value
var once_eqNonEmptyReplacement sync.Once
func Get_eqNonEmptyReplacement() gopurs_runtime.Value {
	once_eqNonEmptyReplacement.Do(func() {
		cache_eqNonEmptyReplacement = pkg_Data_Eq.Get_eqString()
	})
	return cache_eqNonEmptyReplacement
}

var cache_endsWith gopurs_runtime.Value
var once_endsWith sync.Once
func Get_endsWith() gopurs_runtime.Value {
	once_endsWith.Do(func() {
		cache_endsWith = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_endsWith(x_0_box, v_1_box)
})
	})
	return cache_endsWith
}

var cache_contains gopurs_runtime.Value
var once_contains sync.Once
func Get_contains() gopurs_runtime.Value {
	once_contains.Do(func() {
		cache_contains = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_contains(x_0_box)
})
	})
	return cache_contains
}

var cache_appendString gopurs_runtime.Value
var once_appendString sync.Once
func Get_appendString() gopurs_runtime.Value {
	once_appendString.Do(func() {
		cache_appendString = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_appendString(v_0_box, s2_1_box.StrVal())
})
	})
	return cache_appendString
}

func Call_NonEmptyString(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_NonEmptyReplacement(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_toUpper(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(pkg_Data_String_Common.Get_toUpper(), v_0)
}

func Call_toString(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_toLower(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0)
}

func Call_replaceAll(pat_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
return gopurs_runtime.Apply3(pkg_Data_String_Common.Get_replaceAll(), pat_0, v_1, v1_2)
}

func Call_replace(pat_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value, v1_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
var v1_2 gopurs_runtime.Value = v1_2_loop
_ = v1_2
return gopurs_runtime.Apply3(pkg_Data_String_Common.Get_replace(), pat_0, v_1, v1_2)
}

func Call_prependString(s1_0_loop string, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s1_0 string = s1_0_loop
_ = s1_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(s1_0), v_1)
}

func Call_nonEmptyNonEmpty(dictIsSymbol_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
return gopurs_runtime.RecordDict1("nes", gopurs_runtime.Func(func(p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(((*gopurs_runtime.RecordData1)(dictIsSymbol_0.UnsafePtr)).V0, p_1)
}))
}

func Call_nes(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return ((*gopurs_runtime.RecordData1)(dict_0.UnsafePtr)).V0
}

func Call_makeNonEmptyBad(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.RecordDict1("nes", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
}))
}

func Call_localeCompare(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
return gopurs_runtime.Apply2(pkg_Data_String_Common.Get_localeCompare(), v_0, v1_1)
}

func Call_liftS(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, v_1)
}

func Call_startsWith(x_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_stripPrefix(), x_0, v_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3589588149) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136) {
__t1 = gopurs_runtime.Bool(true)
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

func Call_joinWith1(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "Foldable0_NOT_FOUND"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
return gopurs_runtime.Func2(func(sep_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldl"), gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_5, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_6, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), gopurs_runtime.RecordGet(v_5, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), sep_3, v1_6)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_4), "acc")
})
}

func Call_joinWith(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(splice_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(((*gopurs_runtime.RecordData3)(dictFoldable_0.UnsafePtr)).V1, gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_4, "init").IntVal) != (0) {
__t1 = gopurs_runtime.RecordDict2("acc", "init", v1_5, gopurs_runtime.Bool(false))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), gopurs_runtime.RecordGet(v_4, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "append"), splice_2, v1_5)), gopurs_runtime.Bool(false))
}
end_branch_1:
return __t1
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_3), "acc")
})
}

func Call_join1With(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
return gopurs_runtime.Apply(Get_joinWith(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "Foldable0_NOT_FOUND"), gopurs_runtime.Value{}))
}

func Call_fromString(v_0_loop string) gopurs_runtime.Value {
var v_0 string = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == ("") {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{gopurs_runtime.Str(v_0)})}
}
end_branch_0:
return __t0
}

func Call_stripPrefix(pat_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_stripPrefix(), pat_0, a_1), Get_fromString())
}

func Call_stripSuffix(pat_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_stripSuffix(), pat_0, a_1), Get_fromString())
}

func Call_trim(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_Common.Get_trim(), v_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0.StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 3589588149, UnsafePtr: nil}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{__local_var_1_0})}
}
end_branch_1:
return __t1
}

func Call_unsafeFromString(_dollar__unused_0_loop gopurs_runtime.Value, x_1_loop string) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var x_1 string = x_1_loop
_ = x_1
var __t0 gopurs_runtime.Value
{
if (x_1) == ("") {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Str(x_1)
}
end_branch_0:
return __t0
}

func Call_endsWith(x_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_stripSuffix(), x_0, v_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3589588149) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136) {
__t1 = gopurs_runtime.Bool(true)
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

func Call_contains(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_contains(), x_0)
}

func Call_appendString(v_0_loop gopurs_runtime.Value, s2_1_loop string) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s2_1 string = s2_1_loop
_ = s2_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), v_0, gopurs_runtime.Str(s2_1))
}



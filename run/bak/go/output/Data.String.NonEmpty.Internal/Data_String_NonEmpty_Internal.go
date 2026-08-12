package Data_String_NonEmpty_Internal

import (
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_Symbol "gopurs/output/Data.Symbol"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
return gopurs_runtime.Str(Call_toUpper(v_0_box.StrVal()))
})
	})
	return cache_toUpper
}

var cache_toString gopurs_runtime.Value
var once_toString sync.Once
func Get_toString() gopurs_runtime.Value {
	once_toString.Do(func() {
		cache_toString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_toString(v_0_box.StrVal()))
})
	})
	return cache_toString
}

var cache_toLower gopurs_runtime.Value
var once_toLower sync.Once
func Get_toLower() gopurs_runtime.Value {
	once_toLower.Do(func() {
		cache_toLower = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_toLower(v_0_box.StrVal()))
})
	})
	return cache_toLower
}

var cache_showNonEmptyString gopurs_runtime.Value
var once_showNonEmptyString sync.Once
func Get_showNonEmptyString() gopurs_runtime.Value {
	once_showNonEmptyString.Do(func() {
		cache_showNonEmptyString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(NonEmptyString.unsafeFromString "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_show__3756561682(gopurs_runtime.Str(v_0.StrVal())).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
	})
	return cache_showNonEmptyString
}

var cache_showNonEmptyReplacement gopurs_runtime.Value
var once_showNonEmptyReplacement sync.Once
func Get_showNonEmptyReplacement() gopurs_runtime.Value {
	once_showNonEmptyReplacement.Do(func() {
		cache_showNonEmptyReplacement = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(NonEmptyReplacement "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(Call_show__3756561682(gopurs_runtime.Str(v_0.StrVal())).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
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
return gopurs_runtime.Str(Call_replaceAll(pat_0_box.StrVal(), v_1_box.StrVal(), v1_2_box.StrVal()))
})
	})
	return cache_replaceAll
}

var cache_replace gopurs_runtime.Value
var once_replace sync.Once
func Get_replace() gopurs_runtime.Value {
	once_replace.Do(func() {
		cache_replace = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_replace(pat_0_box.StrVal(), v_1_box.StrVal(), v1_2_box.StrVal()))
})
	})
	return cache_replace
}

var cache_prependString gopurs_runtime.Value
var once_prependString sync.Once
func Get_prependString() gopurs_runtime.Value {
	once_prependString.Do(func() {
		cache_prependString = gopurs_runtime.Func2(func(s1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_prependString(s1_0_box.StrVal(), v_1_box.StrVal()))
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
return Call_nes(gopurs_runtime.CoerceToStruct[Constructor_MakeNonEmpty[gopurs_runtime.Value]](dict_0_box))
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
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_localeCompare(v_0_box.StrVal(), v1_1_box.StrVal())), UnsafePtr: nil}
})
	})
	return cache_localeCompare
}

var cache_liftS gopurs_runtime.Value
var once_liftS sync.Once
func Get_liftS() gopurs_runtime.Value {
	once_liftS.Do(func() {
		cache_liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftS(f_0_box, v_1_box.StrVal())
})
	})
	return cache_liftS
}

var cache_startsWith gopurs_runtime.Value
var once_startsWith sync.Once
func Get_startsWith() gopurs_runtime.Value {
	once_startsWith.Do(func() {
		cache_startsWith = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_startsWith(x_0_box.StrVal(), v_1_box.StrVal()))
})
	})
	return cache_startsWith
}

var cache_joinWith1 gopurs_runtime.Value
var once_joinWith1 sync.Once
func Get_joinWith1() gopurs_runtime.Value {
	once_joinWith1.Do(func() {
		cache_joinWith1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinWith1(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_joinWith1
}

var cache_joinWith gopurs_runtime.Value
var once_joinWith sync.Once
func Get_joinWith() gopurs_runtime.Value {
	once_joinWith.Do(func() {
		cache_joinWith = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, splice_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinWith(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), splice_1_box.StrVal())
})
	})
	return cache_joinWith
}

var cache_join1With gopurs_runtime.Value
var once_join1With sync.Once
func Get_join1With() gopurs_runtime.Value {
	once_join1With.Do(func() {
		cache_join1With = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_join1With(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_join1With
}

var cache_fromString gopurs_runtime.Value
var once_fromString sync.Once
func Get_fromString() gopurs_runtime.Value {
	once_fromString.Do(func() {
		cache_fromString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_fromString(v_0_box.StrVal()))}
})
	})
	return cache_fromString
}

var cache_stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		cache_stripPrefix = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripPrefix(pat_0_box.StrVal())
})
	})
	return cache_stripPrefix
}

var cache_stripSuffix gopurs_runtime.Value
var once_stripSuffix sync.Once
func Get_stripSuffix() gopurs_runtime.Value {
	once_stripSuffix.Do(func() {
		cache_stripSuffix = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripSuffix(pat_0_box.StrVal())
})
	})
	return cache_stripSuffix
}

var cache_trim gopurs_runtime.Value
var once_trim sync.Once
func Get_trim() gopurs_runtime.Value {
	once_trim.Do(func() {
		cache_trim = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_trim(v_0_box.StrVal()))}
})
	})
	return cache_trim
}

var cache_unsafeFromString gopurs_runtime.Value
var once_unsafeFromString sync.Once
func Get_unsafeFromString() gopurs_runtime.Value {
	once_unsafeFromString.Do(func() {
		cache_unsafeFromString = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_unsafeFromString(_dollar__unused_0_box, x_1_box.StrVal()))
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
return gopurs_runtime.Bool(Call_endsWith(x_0_box.StrVal(), v_1_box.StrVal()))
})
	})
	return cache_endsWith
}

var cache_contains gopurs_runtime.Value
var once_contains sync.Once
func Get_contains() gopurs_runtime.Value {
	once_contains.Do(func() {
		cache_contains = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_contains(x_0_box.StrVal())
})
	})
	return cache_contains
}

var cache_appendString gopurs_runtime.Value
var once_appendString sync.Once
func Get_appendString() gopurs_runtime.Value {
	once_appendString.Do(func() {
		cache_appendString = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_appendString(v_0_box.StrVal(), s2_1_box.StrVal()))
})
	})
	return cache_appendString
}

var cache_bindFlipped__2564545729 gopurs_runtime.Value
var once_bindFlipped__2564545729 sync.Once
func Get_bindFlipped__2564545729() gopurs_runtime.Value {
	once_bindFlipped__2564545729.Do(func() {
		cache_bindFlipped__2564545729 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__2564545729(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_bindFlipped__2564545729
}

var cache_bindFlipped__1485397639 gopurs_runtime.Value
var once_bindFlipped__1485397639 sync.Once
func Get_bindFlipped__1485397639() gopurs_runtime.Value {
	once_bindFlipped__1485397639.Do(func() {
		cache_bindFlipped__1485397639 = gopurs_runtime.Func3(func(dictBind_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindFlipped__1485397639(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), b_1_box, a_2_box)
})
	})
	return cache_bindFlipped__1485397639
}

var cache_composeKleisliFlipped__3637617434 gopurs_runtime.Value
var once_composeKleisliFlipped__3637617434 sync.Once
func Get_composeKleisliFlipped__3637617434() gopurs_runtime.Value {
	once_composeKleisliFlipped__3637617434.Do(func() {
		cache_composeKleisliFlipped__3637617434 = gopurs_runtime.Func4(func(dictBind_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value, a_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeKleisliFlipped__3637617434(gopurs_runtime.CoerceToStruct[pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value]](dictBind_0_box), f_1_box, g_2_box, a_3_box)
})
	})
	return cache_composeKleisliFlipped__3637617434
}

var cache_composeKleisliFlipped__2781497852 gopurs_runtime.Value
var once_composeKleisliFlipped__2781497852 sync.Once
func Get_composeKleisliFlipped__2781497852() gopurs_runtime.Value {
	once_composeKleisliFlipped__2781497852.Do(func() {
		cache_composeKleisliFlipped__2781497852 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeKleisliFlipped__2781497852(f_0_box, g_1_box, a_2_box)
})
	})
	return cache_composeKleisliFlipped__2781497852
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

var cache_eq__472317769 gopurs_runtime.Value
var once_eq__472317769 sync.Once
func Get_eq__472317769() gopurs_runtime.Value {
	once_eq__472317769.Do(func() {
		cache_eq__472317769 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__472317769(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__472317769
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

var cache_foldl__3785384859 gopurs_runtime.Value
var once_foldl__3785384859 sync.Once
func Get_foldl__3785384859() gopurs_runtime.Value {
	once_foldl__3785384859.Do(func() {
		cache_foldl__3785384859 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__3785384859(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__3785384859
}

var cache_intercalate__3813868388 gopurs_runtime.Value
var once_intercalate__3813868388 sync.Once
func Get_intercalate__3813868388() gopurs_runtime.Value {
	once_intercalate__3813868388.Do(func() {
		cache_intercalate__3813868388 = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, sep_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate__3813868388(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), sep_1_box, xs_2_box)
})
	})
	return cache_intercalate__3813868388
}

var cache_intercalate__3939234276 gopurs_runtime.Value
var once_intercalate__3939234276 sync.Once
func Get_intercalate__3939234276() gopurs_runtime.Value {
	once_intercalate__3939234276.Do(func() {
		cache_intercalate__3939234276 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, dictMonoid_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate__3939234276(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1_box))
})
	})
	return cache_intercalate__3939234276
}

var cache_const__1243414737 gopurs_runtime.Value
var once_const__1243414737 sync.Once
func Get_const__1243414737() gopurs_runtime.Value {
	once_const__1243414737.Do(func() {
		cache_const__1243414737 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1243414737(a_0_box, v_1_box)
})
	})
	return cache_const__1243414737
}

var cache_const__2082174484 gopurs_runtime.Value
var once_const__2082174484 sync.Once
func Get_const__2082174484() gopurs_runtime.Value {
	once_const__2082174484.Do(func() {
		cache_const__2082174484 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__2082174484(a_0_box, v_1_box)
})
	})
	return cache_const__2082174484
}

var cache_const__4157258135 gopurs_runtime.Value
var once_const__4157258135 sync.Once
func Get_const__4157258135() gopurs_runtime.Value {
	once_const__4157258135.Do(func() {
		cache_const__4157258135 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__4157258135(a_0_box, v_1_box)
})
	})
	return cache_const__4157258135
}

var cache_const__1562253172 gopurs_runtime.Value
var once_const__1562253172 sync.Once
func Get_const__1562253172() gopurs_runtime.Value {
	once_const__1562253172.Do(func() {
		cache_const__1562253172 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1562253172(a_0_box, v_1_box)
})
	})
	return cache_const__1562253172
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

var cache_flip__2253242624 gopurs_runtime.Value
var once_flip__2253242624 sync.Once
func Get_flip__2253242624() gopurs_runtime.Value {
	once_flip__2253242624.Do(func() {
		cache_flip__2253242624 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2253242624(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2253242624
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

var cache_map__901270812 gopurs_runtime.Value
var once_map__901270812 sync.Once
func Get_map__901270812() gopurs_runtime.Value {
	once_map__901270812.Do(func() {
		cache_map__901270812 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__901270812(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__901270812
}

var cache_applyMaybe__3698865467 gopurs_runtime.Value
var once_applyMaybe__3698865467 sync.Once
func Get_applyMaybe__3698865467() gopurs_runtime.Value {
	once_applyMaybe__3698865467.Do(func() {
		cache_applyMaybe__3698865467 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_functorMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1))})))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_applyMaybe__3698865467
}

var cache_bindMaybe__1910292045 gopurs_runtime.Value
var once_bindMaybe__1910292045 sync.Once
func Get_bindMaybe__1910292045() gopurs_runtime.Value {
	once_bindMaybe__1910292045.Do(func() {
		cache_bindMaybe__1910292045 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Maybe.Get_applyMaybe()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_0.UnsafePtr).V0)))}
goto end_branch_0
} else {

}
}
{
if (v_0.Type == 9 && v_0.IntVal == 930809136 && v_0.UnsafePtr == nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_bindMaybe__1910292045
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_functorMaybe__2097654001 gopurs_runtime.Value
var once_functorMaybe__2097654001 sync.Once
func Get_functorMaybe__2097654001() gopurs_runtime.Value {
	once_functorMaybe__2097654001.Do(func() {
		cache_functorMaybe__2097654001 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2097654001
}

var cache_isJust__2514352589 gopurs_runtime.Value
var once_isJust__2514352589 sync.Once
func Get_isJust__2514352589() gopurs_runtime.Value {
	once_isJust__2514352589.Do(func() {
		cache_isJust__2514352589 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__2514352589(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[int64]](v2_0_box)))
})
	})
	return cache_isJust__2514352589
}

var cache_isJust__2475527019 gopurs_runtime.Value
var once_isJust__2475527019 sync.Once
func Get_isJust__2475527019() gopurs_runtime.Value {
	once_isJust__2475527019.Do(func() {
		cache_isJust__2475527019 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isJust__2475527019(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](v2_0_box)))
})
	})
	return cache_isJust__2475527019
}

var cache_maybe__3078346790 gopurs_runtime.Value
var once_maybe__3078346790 sync.Once
func Get_maybe__3078346790() gopurs_runtime.Value {
	once_maybe__3078346790.Do(func() {
		cache_maybe__3078346790 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3078346790(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3078346790
}

var cache_maybe__1510464358 gopurs_runtime.Value
var once_maybe__1510464358 sync.Once
func Get_maybe__1510464358() gopurs_runtime.Value {
	once_maybe__1510464358.Do(func() {
		cache_maybe__1510464358 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1510464358(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1510464358
}

var cache_maybe__3718989812 gopurs_runtime.Value
var once_maybe__3718989812 sync.Once
func Get_maybe__3718989812() gopurs_runtime.Value {
	once_maybe__3718989812.Do(func() {
		cache_maybe__3718989812 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3718989812(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3718989812
}

var cache_maybe__1647364852 gopurs_runtime.Value
var once_maybe__1647364852 sync.Once
func Get_maybe__1647364852() gopurs_runtime.Value {
	once_maybe__1647364852.Do(func() {
		cache_maybe__1647364852 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1647364852(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1647364852
}

var cache_sub__1043827704 gopurs_runtime.Value
var once_sub__1043827704 sync.Once
func Get_sub__1043827704() gopurs_runtime.Value {
	once_sub__1043827704.Do(func() {
		cache_sub__1043827704 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__1043827704(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_sub__1043827704
}

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
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

var cache_show__3756561682 gopurs_runtime.Value
var once_show__3756561682 sync.Once
func Get_show__3756561682() gopurs_runtime.Value {
	once_show__3756561682.Do(func() {
		cache_show__3756561682 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3756561682(__eta0_0_box)
})
	})
	return cache_show__3756561682
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

var cache_joinWith__632268499 gopurs_runtime.Value
var once_joinWith__632268499 sync.Once
func Get_joinWith__632268499() gopurs_runtime.Value {
	once_joinWith__632268499.Do(func() {
		cache_joinWith__632268499 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, splice_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinWith__632268499(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), splice_1_box.StrVal())
})
	})
	return cache_joinWith__632268499
}

var cache_liftS__895676186 gopurs_runtime.Value
var once_liftS__895676186 sync.Once
func Get_liftS__895676186() gopurs_runtime.Value {
	once_liftS__895676186.Do(func() {
		cache_liftS__895676186 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_liftS__895676186(f_0_box, v_1_box.StrVal()))
})
	})
	return cache_liftS__895676186
}

var cache_liftS__3230749042 gopurs_runtime.Value
var once_liftS__3230749042 sync.Once
func Get_liftS__3230749042() gopurs_runtime.Value {
	once_liftS__3230749042.Do(func() {
		cache_liftS__3230749042 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftS__3230749042(f_0_box, v_1_box.StrVal())
})
	})
	return cache_liftS__3230749042
}

var cache_liftS__3241548146 gopurs_runtime.Value
var once_liftS__3241548146 sync.Once
func Get_liftS__3241548146() gopurs_runtime.Value {
	once_liftS__3241548146.Do(func() {
		cache_liftS__3241548146 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftS__3241548146(f_0_box, v_1_box.StrVal())
})
	})
	return cache_liftS__3241548146
}

var cache_reflectSymbol__3416619207 gopurs_runtime.Value
var once_reflectSymbol__3416619207 sync.Once
func Get_reflectSymbol__3416619207() gopurs_runtime.Value {
	once_reflectSymbol__3416619207.Do(func() {
		cache_reflectSymbol__3416619207 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__3416619207(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__3416619207
}

var cache_reflectSymbol__1166932993 gopurs_runtime.Value
var once_reflectSymbol__1166932993 sync.Once
func Get_reflectSymbol__1166932993() gopurs_runtime.Value {
	once_reflectSymbol__1166932993.Do(func() {
		cache_reflectSymbol__1166932993 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_reflectSymbol__1166932993(gopurs_runtime.CoerceToStruct[pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_reflectSymbol__1166932993
}

type Constructor_MakeNonEmpty[T_s any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1987403114] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_MakeNonEmpty[gopurs_runtime.Value])(ptr)
		switch key {
		case "nes": return c.V0
		default: panic("Key not found in dictionary Constructor_MakeNonEmpty: " + key)
		}
	}
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

func Call_toUpper(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return gopurs_runtime.Apply(pkg_Data_String_Common.Get_toUpper(), gopurs_runtime.Str(v_0)).StrVal()
}

func Call_toString(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return v_0
}

func Call_toLower(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), gopurs_runtime.Str(v_0)).StrVal()
}

func Call_replaceAll(pat_0_loop string, v_1_loop string, v1_2_loop string) string {
var pat_0 string = pat_0_loop
_ = pat_0
var v_1 string = v_1_loop
_ = v_1
var v1_2 string = v1_2_loop
_ = v1_2
return gopurs_runtime.Apply3(pkg_Data_String_Common.Get_replaceAll(), gopurs_runtime.Str(pat_0), gopurs_runtime.Str(v_1), gopurs_runtime.Str(v1_2)).StrVal()
}

func Call_replace(pat_0_loop string, v_1_loop string, v1_2_loop string) string {
var pat_0 string = pat_0_loop
_ = pat_0
var v_1 string = v_1_loop
_ = v_1
var v1_2 string = v1_2_loop
_ = v1_2
return gopurs_runtime.Apply3(pkg_Data_String_Common.Get_replace(), gopurs_runtime.Str(pat_0), gopurs_runtime.Str(v_1), gopurs_runtime.Str(v1_2)).StrVal()
}

func Call_prependString(s1_0_loop string, v_1_loop string) string {
var s1_0 string = s1_0_loop
_ = s1_0
var v_1 string = v_1_loop
_ = v_1
return Call_append__493084344(gopurs_runtime.Str(s1_0), gopurs_runtime.Str(v_1)).StrVal()
}

func Call_nonEmptyNonEmpty(dictIsSymbol_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
return gopurs_runtime.RecordDict1("nes", gopurs_runtime.Func(func(p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(p_1.IntVal)), UnsafePtr: nil}).StrVal())
}))
}

func Call_nes(dict_0_loop *Constructor_MakeNonEmpty[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_MakeNonEmpty[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_makeNonEmptyBad(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.RecordDict1("nes", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
}))
}

func Call_localeCompare(v_0_loop string, v1_1_loop string) uint32 {
var v_0 string = v_0_loop
_ = v_0
var v1_1 string = v1_1_loop
_ = v1_1
return uint32(gopurs_runtime.Apply2(pkg_Data_String_Common.Get_localeCompare(), gopurs_runtime.Str(v_0), gopurs_runtime.Str(v1_1)).IntVal)
}

func Call_liftS(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_startsWith(x_0_loop string, v_1_loop string) bool {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_stripPrefix(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal()), gopurs_runtime.Str(v_1))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_joinWith1(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_3_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_2
__local_var_3_1 := gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(Foldable0_1_0.V1, gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_5, "init").IntVal) != (0) {
__t3 = gopurs_runtime.RecordDict2("acc", "init", v1_6, gopurs_runtime.Bool(false))
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(Semigroup0_3_2.V0, gopurs_runtime.RecordGet(v_5, "acc"), gopurs_runtime.Apply2(Semigroup0_3_2.V0, gopurs_runtime.Str(v_2.StrVal()), v1_6)), gopurs_runtime.Bool(false))
}
end_branch_3:
return __t3
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_4), "acc")
})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4)
})
})
}

func Call_joinWith(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], splice_1_loop string) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var splice_1 string = splice_1_loop
_ = splice_1
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
__local_var_2_0 := gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(dictFoldable_0.V1, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_4, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_5, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(Semigroup0_2_1.V0, gopurs_runtime.RecordGet(v_4, "acc"), gopurs_runtime.Apply2(Semigroup0_2_1.V0, gopurs_runtime.Str(splice_1), v1_5)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_3), "acc")
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, x_3)
})
}

func Call_join1With(dictFoldable1_0_loop *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *pkg_Data_Semigroup_Foldable.Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(splice_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := Call_joinWith(Foldable0_1_0, splice_2.StrVal())
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4)
})
})
}

func Call_fromString(v_0_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var v_0 string = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == ("") {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(v_0)})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t0)
}

func Call_stripPrefix(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
return gopurs_runtime.Apply2(Get_composeKleisliFlipped__2781497852(), Get_fromString(), gopurs_runtime.Apply(Get_liftS__3241548146(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_stripPrefix(), gopurs_runtime.Str(pat_0))))
}

func Call_stripSuffix(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
return gopurs_runtime.Apply2(Get_composeKleisliFlipped__2781497852(), Get_fromString(), gopurs_runtime.Apply(Get_liftS__3241548146(), gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_stripSuffix(), gopurs_runtime.Str(pat_0))))
}

func Call_trim(v_0_loop string) *pkg_Data_Maybe.Constructor_Just[string] {
var v_0 string = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_Common.Get_trim(), gopurs_runtime.Str(v_0)).StrVal()
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(__local_var_1_0)})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t1)
}

func Call_unsafeFromString(_dollar__unused_0_loop gopurs_runtime.Value, x_1_loop string) string {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var x_1 string = x_1_loop
_ = x_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Str(x_1).StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Str(gopurs_runtime.Str(x_1).StrVal())})}
}
end_branch_1:
__local_var_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[string]](__t1)
_ = __local_var_2_0
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_2_0)}.UnsafePtr != nil) {
__t2 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_2_0)}.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2.StrVal()
}

func Call_endsWith(x_0_loop string, v_1_loop string) bool {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_stripSuffix(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal()), gopurs_runtime.Str(v_1))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_contains(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_appendString(v_0_loop string, s2_1_loop string) string {
var v_0 string = v_0_loop
_ = v_0
var s2_1 string = s2_1_loop
_ = s2_1
return Call_append__493084344(gopurs_runtime.Str(v_0), gopurs_runtime.Str(s2_1)).StrVal()
}

func Call_bindFlipped__2564545729(dictBind_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictBind_0.V1, a_2, b_1)
}

func Call_bindFlipped__1485397639(dictBind_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value], b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(dictBind_0.V1, a_2, b_1)
}

func Call_composeKleisliFlipped__3637617434(dictBind_0_loop *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value, a_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBind_0 *pkg_Control_Bind.Constructor_Bind[gopurs_runtime.Value] = dictBind_0_loop
_ = dictBind_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
var a_3 gopurs_runtime.Value = a_3_loop
_ = a_3
return gopurs_runtime.Apply2(dictBind_0.V1, gopurs_runtime.Apply(g_2, a_3), f_1)
}

func Call_composeKleisliFlipped__2781497852(f_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_bindMaybe(), "bind"), gopurs_runtime.Apply(g_1, a_2), f_0)
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__472317769(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.StrVal()) == (__eta1_1.StrVal()))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl__3785384859(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_intercalate__3813868388(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], sep_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var sep_1 gopurs_runtime.Value = sep_1_loop
_ = sep_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(dictFoldable_0.V1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_3, "init").IntVal) != (0) {
__t0 = gopurs_runtime.RecordDict2("acc", "init", v1_4, gopurs_runtime.Bool(false))
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.RecordGet(v_3, "acc"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), sep_1, v1_4)), gopurs_runtime.Bool(false))
}
end_branch_0:
return __t0
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_2), "acc")
}

func Call_intercalate__3939234276(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], dictMonoid_1_loop *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var dictMonoid_1 *pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value] = dictMonoid_1_loop
_ = dictMonoid_1
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(dictMonoid_1.V0, gopurs_runtime.Value{}))
_ = Semigroup0_2_0
mempty_3_1 := dictMonoid_1.V1
_ = mempty_3_1
return gopurs_runtime.Func(func(sep_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(dictFoldable_0.V1, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_6, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_7, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(Semigroup0_2_0.V0, gopurs_runtime.RecordGet(v_6, "acc"), gopurs_runtime.Apply2(Semigroup0_2_0.V0, sep_4, v1_7)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", mempty_3_1, gopurs_runtime.Bool(true)), xs_5), "acc")
})
})
}

func Call_const__1243414737(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__2082174484(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__4157258135(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__1562253172(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
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

func Call_flip__2253242624(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__901270812(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0)
}

func Call_isJust__2514352589(v2_0_loop *pkg_Data_Maybe.Constructor_Just[int64]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[int64] = v2_0_loop
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

func Call_isJust__2475527019(v2_0_loop *pkg_Data_Maybe.Constructor_Just[string]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[string] = v2_0_loop
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

func Call_maybe__3078346790(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__1510464358(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__3718989812(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_maybe__1647364852(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_sub__1043827704(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Int((__eta0_0.IntVal) - (__eta1_1.IntVal))
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
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

func Call_show__3756561682(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Str(gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), __eta0_0).StrVal())
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_joinWith__632268499(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value], splice_1_loop string) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var splice_1 string = splice_1_loop
_ = splice_1
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
__local_var_2_0 := gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(dictFoldable_0.V1, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_4, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_5, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(Semigroup0_2_1.V0, gopurs_runtime.RecordGet(v_4, "acc"), gopurs_runtime.Apply2(Semigroup0_2_1.V0, gopurs_runtime.Str(splice_1), v1_5)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(pkg_Data_Monoid.Get_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_3), "acc")
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, x_3)
})
}

func Call_liftS__895676186(f_0_loop gopurs_runtime.Value, v_1_loop string) bool {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return (gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)).IntVal) != (0)
}

func Call_liftS__3230749042(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_liftS__3241548146(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_reflectSymbol__3416619207(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_reflectSymbol__1166932993(dict_0_loop *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Symbol.Constructor_IsSymbol[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}



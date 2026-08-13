package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_NonEmpty_Internal_fromJust gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_fromJust sync.Once
func Get_Data_String_NonEmpty_Internal_fromJust() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_fromJust.Do(func() {
		cache_Data_String_NonEmpty_Internal_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box)))
})
	})
	return cache_Data_String_NonEmpty_Internal_fromJust
}

var cache_Data_String_NonEmpty_Internal_NonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_NonEmptyString sync.Once
func Get_Data_String_NonEmpty_Internal_NonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_NonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_Internal_NonEmptyString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_NonEmptyString(x_0_box)
})
	})
	return cache_Data_String_NonEmpty_Internal_NonEmptyString
}

var cache_Data_String_NonEmpty_Internal_NonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_NonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_NonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_NonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_NonEmptyReplacement = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_NonEmptyReplacement(x_0_box)
})
	})
	return cache_Data_String_NonEmpty_Internal_NonEmptyReplacement
}

var cache_Data_String_NonEmpty_Internal_MakeNonEmpty_dollarDict gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_MakeNonEmpty_dollarDict sync.Once
func Get_Data_String_NonEmpty_Internal_MakeNonEmpty_dollarDict() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_MakeNonEmpty_dollarDict.Do(func() {
		cache_Data_String_NonEmpty_Internal_MakeNonEmpty_dollarDict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_MakeNonEmpty_dollarDict(x_0_box)
})
	})
	return cache_Data_String_NonEmpty_Internal_MakeNonEmpty_dollarDict
}

var cache_Data_String_NonEmpty_Internal_toUpper gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_toUpper sync.Once
func Get_Data_String_NonEmpty_Internal_toUpper() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_toUpper.Do(func() {
		cache_Data_String_NonEmpty_Internal_toUpper = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_toUpper(v_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_toUpper
}

var cache_Data_String_NonEmpty_Internal_toString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_toString sync.Once
func Get_Data_String_NonEmpty_Internal_toString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_toString.Do(func() {
		cache_Data_String_NonEmpty_Internal_toString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_toString(v_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_toString
}

var cache_Data_String_NonEmpty_Internal_toLower gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_toLower sync.Once
func Get_Data_String_NonEmpty_Internal_toLower() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_toLower.Do(func() {
		cache_Data_String_NonEmpty_Internal_toLower = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_toLower(v_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_toLower
}

var cache_Data_String_NonEmpty_Internal_showNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_showNonEmptyString sync.Once
func Get_Data_String_NonEmpty_Internal_showNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_showNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_Internal_showNonEmptyString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyString.unsafeFromString ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_String_NonEmpty_Internal_showNonEmptyString
}

var cache_Data_String_NonEmpty_Internal_showNonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_showNonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_showNonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_showNonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_showNonEmptyReplacement = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyReplacement ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_String_NonEmpty_Internal_showNonEmptyReplacement
}

var cache_Data_String_NonEmpty_Internal_semigroupNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_semigroupNonEmptyString sync.Once
func Get_Data_String_NonEmpty_Internal_semigroupNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_semigroupNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_Internal_semigroupNonEmptyString = Get_Data_Semigroup_semigroupString()
	})
	return cache_Data_String_NonEmpty_Internal_semigroupNonEmptyString
}

var cache_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement = Get_Data_Semigroup_semigroupString()
	})
	return cache_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement
}

var cache_Data_String_NonEmpty_Internal_replaceAll gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_replaceAll sync.Once
func Get_Data_String_NonEmpty_Internal_replaceAll() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_replaceAll.Do(func() {
		cache_Data_String_NonEmpty_Internal_replaceAll = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_replaceAll(pat_0_box.StrVal(), v_1_box.StrVal(), v1_2_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_replaceAll
}

var cache_Data_String_NonEmpty_Internal_replace gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_replace sync.Once
func Get_Data_String_NonEmpty_Internal_replace() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_replace.Do(func() {
		cache_Data_String_NonEmpty_Internal_replace = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_replace(pat_0_box.StrVal(), v_1_box.StrVal(), v1_2_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_replace
}

var cache_Data_String_NonEmpty_Internal_prependString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_prependString sync.Once
func Get_Data_String_NonEmpty_Internal_prependString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_prependString.Do(func() {
		cache_Data_String_NonEmpty_Internal_prependString = gopurs_runtime.Func2(func(s1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_prependString(s1_0_box.StrVal(), v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_prependString
}

var cache_Data_String_NonEmpty_Internal_ordNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_ordNonEmptyString sync.Once
func Get_Data_String_NonEmpty_Internal_ordNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_ordNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_Internal_ordNonEmptyString = Get_Data_Ord_ordString()
	})
	return cache_Data_String_NonEmpty_Internal_ordNonEmptyString
}

var cache_Data_String_NonEmpty_Internal_ordNonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_ordNonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_ordNonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_ordNonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_ordNonEmptyReplacement = Get_Data_Ord_ordString()
	})
	return cache_Data_String_NonEmpty_Internal_ordNonEmptyReplacement
}

var cache_Data_String_NonEmpty_Internal_nonEmptyNonEmpty gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_nonEmptyNonEmpty sync.Once
func Get_Data_String_NonEmpty_Internal_nonEmptyNonEmpty() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_nonEmptyNonEmpty.Do(func() {
		cache_Data_String_NonEmpty_Internal_nonEmptyNonEmpty = gopurs_runtime.Func(func(dictIsSymbol_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_nonEmptyNonEmpty(dictIsSymbol_0_box)
})
	})
	return cache_Data_String_NonEmpty_Internal_nonEmptyNonEmpty
}

var cache_Data_String_NonEmpty_Internal_nes gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_nes sync.Once
func Get_Data_String_NonEmpty_Internal_nes() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_nes.Do(func() {
		cache_Data_String_NonEmpty_Internal_nes = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_nes(gopurs_runtime.CoerceToStruct[Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty](dict_0_box))
})
	})
	return cache_Data_String_NonEmpty_Internal_nes
}

var cache_Data_String_NonEmpty_Internal_makeNonEmptyBad gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_makeNonEmptyBad sync.Once
func Get_Data_String_NonEmpty_Internal_makeNonEmptyBad() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_makeNonEmptyBad.Do(func() {
		cache_Data_String_NonEmpty_Internal_makeNonEmptyBad = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_makeNonEmptyBad(_dollar__unused_0_box)
})
	})
	return cache_Data_String_NonEmpty_Internal_makeNonEmptyBad
}

var cache_Data_String_NonEmpty_Internal_localeCompare gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_localeCompare sync.Once
func Get_Data_String_NonEmpty_Internal_localeCompare() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_localeCompare.Do(func() {
		cache_Data_String_NonEmpty_Internal_localeCompare = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(Call_Data_String_NonEmpty_Internal_localeCompare(v_0_box.StrVal(), v1_1_box.StrVal())), UnsafePtr: nil}
})
	})
	return cache_Data_String_NonEmpty_Internal_localeCompare
}

var cache_Data_String_NonEmpty_Internal_liftS gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_liftS sync.Once
func Get_Data_String_NonEmpty_Internal_liftS() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_liftS.Do(func() {
		cache_Data_String_NonEmpty_Internal_liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_liftS(f_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_Internal_liftS
}

var cache_Data_String_NonEmpty_Internal_startsWith gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_startsWith sync.Once
func Get_Data_String_NonEmpty_Internal_startsWith() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_startsWith.Do(func() {
		cache_Data_String_NonEmpty_Internal_startsWith = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_String_NonEmpty_Internal_startsWith(x_0_box.StrVal(), v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_startsWith
}

var cache_Data_String_NonEmpty_Internal_joinWith1 gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_joinWith1 sync.Once
func Get_Data_String_NonEmpty_Internal_joinWith1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_joinWith1.Do(func() {
		cache_Data_String_NonEmpty_Internal_joinWith1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_joinWith1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Data_String_NonEmpty_Internal_joinWith1
}

var cache_Data_String_NonEmpty_Internal_joinWith gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_joinWith sync.Once
func Get_Data_String_NonEmpty_Internal_joinWith() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_joinWith.Do(func() {
		cache_Data_String_NonEmpty_Internal_joinWith = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, splice_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_joinWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), splice_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_Internal_joinWith
}

var cache_Data_String_NonEmpty_Internal_join1With gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_join1With sync.Once
func Get_Data_String_NonEmpty_Internal_join1With() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_join1With.Do(func() {
		cache_Data_String_NonEmpty_Internal_join1With = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_join1With(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](dictFoldable1_0_box))
})
	})
	return cache_Data_String_NonEmpty_Internal_join1With
}

var cache_Data_String_NonEmpty_Internal_fromString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_fromString sync.Once
func Get_Data_String_NonEmpty_Internal_fromString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_fromString.Do(func() {
		cache_Data_String_NonEmpty_Internal_fromString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_Internal_fromString(v_0_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_Internal_fromString
}

var cache_Data_String_NonEmpty_Internal_stripPrefix gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_stripPrefix sync.Once
func Get_Data_String_NonEmpty_Internal_stripPrefix() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_stripPrefix.Do(func() {
		cache_Data_String_NonEmpty_Internal_stripPrefix = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_stripPrefix(pat_0_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_Internal_stripPrefix
}

var cache_Data_String_NonEmpty_Internal_stripSuffix gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_stripSuffix sync.Once
func Get_Data_String_NonEmpty_Internal_stripSuffix() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_stripSuffix.Do(func() {
		cache_Data_String_NonEmpty_Internal_stripSuffix = gopurs_runtime.Func(func(pat_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_stripSuffix(pat_0_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_Internal_stripSuffix
}

var cache_Data_String_NonEmpty_Internal_trim gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_trim sync.Once
func Get_Data_String_NonEmpty_Internal_trim() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_trim.Do(func() {
		cache_Data_String_NonEmpty_Internal_trim = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_Internal_trim(v_0_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_Internal_trim
}

var cache_Data_String_NonEmpty_Internal_unsafeFromString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_unsafeFromString sync.Once
func Get_Data_String_NonEmpty_Internal_unsafeFromString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_unsafeFromString.Do(func() {
		cache_Data_String_NonEmpty_Internal_unsafeFromString = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_unsafeFromString(_dollar__unused_0_box, x_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_unsafeFromString
}

var cache_Data_String_NonEmpty_Internal_eqNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_eqNonEmptyString sync.Once
func Get_Data_String_NonEmpty_Internal_eqNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_eqNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_Internal_eqNonEmptyString = Get_Data_Eq_eqString()
	})
	return cache_Data_String_NonEmpty_Internal_eqNonEmptyString
}

var cache_Data_String_NonEmpty_Internal_eqNonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_eqNonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_eqNonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_eqNonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_eqNonEmptyReplacement = Get_Data_Eq_eqString()
	})
	return cache_Data_String_NonEmpty_Internal_eqNonEmptyReplacement
}

var cache_Data_String_NonEmpty_Internal_endsWith gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_endsWith sync.Once
func Get_Data_String_NonEmpty_Internal_endsWith() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_endsWith.Do(func() {
		cache_Data_String_NonEmpty_Internal_endsWith = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_String_NonEmpty_Internal_endsWith(x_0_box.StrVal(), v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_endsWith
}

var cache_Data_String_NonEmpty_Internal_contains gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_contains sync.Once
func Get_Data_String_NonEmpty_Internal_contains() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_contains.Do(func() {
		cache_Data_String_NonEmpty_Internal_contains = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_contains(x_0_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_Internal_contains
}

var cache_Data_String_NonEmpty_Internal_appendString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_appendString sync.Once
func Get_Data_String_NonEmpty_Internal_appendString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_appendString.Do(func() {
		cache_Data_String_NonEmpty_Internal_appendString = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_appendString(v_0_box.StrVal(), s2_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_appendString
}

var cache_Data_String_NonEmpty_Internal_joinWith__632268499 gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_joinWith__632268499 sync.Once
func Get_Data_String_NonEmpty_Internal_joinWith__632268499() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_joinWith__632268499.Do(func() {
		cache_Data_String_NonEmpty_Internal_joinWith__632268499 = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, splice_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_joinWith__632268499(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), splice_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_Internal_joinWith__632268499
}

var cache_Data_String_NonEmpty_Internal_liftS__895676186 gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_liftS__895676186 sync.Once
func Get_Data_String_NonEmpty_Internal_liftS__895676186() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_liftS__895676186.Do(func() {
		cache_Data_String_NonEmpty_Internal_liftS__895676186 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_String_NonEmpty_Internal_liftS__895676186(f_0_box, v_1_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_liftS__895676186
}

var cache_Data_String_NonEmpty_Internal_liftS__3230749042 gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_liftS__3230749042 sync.Once
func Get_Data_String_NonEmpty_Internal_liftS__3230749042() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_liftS__3230749042.Do(func() {
		cache_Data_String_NonEmpty_Internal_liftS__3230749042 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_liftS__3230749042(f_0_box, v_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_Internal_liftS__3230749042
}

var cache_Data_String_NonEmpty_Internal_liftS__3241548146 gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_liftS__3241548146 sync.Once
func Get_Data_String_NonEmpty_Internal_liftS__3241548146() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_liftS__3241548146.Do(func() {
		cache_Data_String_NonEmpty_Internal_liftS__3241548146 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_Internal_liftS__3241548146(f_0_box, v_1_box.StrVal()))}
})
	})
	return cache_Data_String_NonEmpty_Internal_liftS__3241548146
}

type Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1987403114] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty)(ptr)
		_ = c
		switch key {
		case "nes": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty: " + key)
		}
	}
}


func Call_Data_String_NonEmpty_Internal_fromJust(v_0_loop *Constructor_Data_Maybe_Just) string {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 != nil) {
__t0 = (v_0).V0
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

func Call_Data_String_NonEmpty_Internal_NonEmptyString(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_NonEmpty_Internal_NonEmptyReplacement(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_NonEmpty_Internal_MakeNonEmpty_dollarDict(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_NonEmpty_Internal_toUpper(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Data_String_Common_toUpper(), gopurs_runtime.Str(v_0)).StrVal()
}

func Call_Data_String_NonEmpty_Internal_toString(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return v_0
}

func Call_Data_String_NonEmpty_Internal_toLower(v_0_loop string) string {
var v_0 string = v_0_loop
_ = v_0
return gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v_0)).StrVal()
}

func Call_Data_String_NonEmpty_Internal_replaceAll(pat_0_loop string, v_1_loop string, v1_2_loop string) string {
var pat_0 string = pat_0_loop
_ = pat_0
var v_1 string = v_1_loop
_ = v_1
var v1_2 string = v1_2_loop
_ = v1_2
return gopurs_runtime.Apply3(Get_Data_String_Common_replaceAll(), gopurs_runtime.Str(pat_0), gopurs_runtime.Str(v_1), gopurs_runtime.Str(v1_2)).StrVal()
}

func Call_Data_String_NonEmpty_Internal_replace(pat_0_loop string, v_1_loop string, v1_2_loop string) string {
var pat_0 string = pat_0_loop
_ = pat_0
var v_1 string = v_1_loop
_ = v_1
var v1_2 string = v1_2_loop
_ = v1_2
return gopurs_runtime.Apply3(Get_Data_String_Common_replace(), gopurs_runtime.Str(pat_0), gopurs_runtime.Str(v_1), gopurs_runtime.Str(v1_2)).StrVal()
}

func Call_Data_String_NonEmpty_Internal_prependString(s1_0_loop string, v_1_loop string) string {
var s1_0 string = s1_0_loop
_ = s1_0
var v_1 string = v_1_loop
_ = v_1
return (s1_0) + (v_1)
}

func Call_Data_String_NonEmpty_Internal_nonEmptyNonEmpty(dictIsSymbol_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
return gopurs_runtime.RecordDict1("nes", gopurs_runtime.Func(func(p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(p_1.IntVal)), UnsafePtr: nil}).StrVal())
}))
}

func Call_Data_String_NonEmpty_Internal_nes(dict_0_loop *Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty) gopurs_runtime.Value {
var dict_0 *Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty = dict_0_loop
_ = dict_0
return gopurs_runtime.Box(dict_0.V0)
}

func Call_Data_String_NonEmpty_Internal_makeNonEmptyBad(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.RecordDict1("nes", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
}))
}

func Call_Data_String_NonEmpty_Internal_localeCompare(v_0_loop string, v1_1_loop string) uint32 {
var v_0 string = v_0_loop
_ = v_0
var v1_1 string = v1_1_loop
_ = v1_1
return uint32(gopurs_runtime.Apply2(Get_Data_String_Common_localeCompare(), gopurs_runtime.Str(v_0), gopurs_runtime.Str(v1_1)).IntVal)
}

func Call_Data_String_NonEmpty_Internal_liftS(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_Data_String_NonEmpty_Internal_startsWith(x_0_loop string, v_1_loop string) bool {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_stripPrefix(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal()), gopurs_runtime.Str(v_1))
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

func Call_Data_String_NonEmpty_Internal_joinWith1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 -> *Constructor_Data_Foldable_Foldable
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Semigroup0_3_2 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_monoidString(), "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_2
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.Box(Foldable0_1_0.V1), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t3 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_2.V0), gopurs_runtime.RecordGet(v_5, "acc"), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_3_2.V0), gopurs_runtime.Str(v_2.StrVal()), v1_6)), gopurs_runtime.Bool(false))
}
end_branch_3:
return __t3
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(Get_Data_Monoid_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_4), "acc")
})
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4)
})
})
}

func Call_Data_String_NonEmpty_Internal_joinWith(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, splice_1_loop string) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var splice_1 string = splice_1_loop
_ = splice_1
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_monoidString(), "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.RecordGet(v_4, "acc"), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Str(splice_1), v1_5)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(Get_Data_Monoid_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_3), "acc")
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, x_3)
})
}

func Call_Data_String_NonEmpty_Internal_join1With(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1 = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 -> *Constructor_Data_Foldable_Foldable
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](gopurs_runtime.Apply(gopurs_runtime.Box(dictFoldable1_0.V0), gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(splice_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := Call_Data_String_NonEmpty_Internal_joinWith(Foldable0_1_0, splice_2.StrVal())
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_3_1, x_4)
})
})
}

func Call_Data_String_NonEmpty_Internal_fromString(v_0_loop string) *Constructor_Data_Maybe_Just {
var v_0 string = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == ("") {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(v_0)})}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t0)
}

func Call_Data_String_NonEmpty_Internal_stripPrefix(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
return gopurs_runtime.Apply2(Get_Control_Bind_composeKleisliFlipped__2781497852(), Get_Data_String_NonEmpty_Internal_fromString(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_Internal_liftS__3241548146(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_stripPrefix(), gopurs_runtime.Str(pat_0))))
}

func Call_Data_String_NonEmpty_Internal_stripSuffix(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
return gopurs_runtime.Apply2(Get_Control_Bind_composeKleisliFlipped__2781497852(), Get_Data_String_NonEmpty_Internal_fromString(), gopurs_runtime.Apply(Get_Data_String_NonEmpty_Internal_liftS__3241548146(), gopurs_runtime.Apply(Get_Data_String_CodeUnits_stripSuffix(), gopurs_runtime.Str(pat_0))))
}

func Call_Data_String_NonEmpty_Internal_trim(v_0_loop string) *Constructor_Data_Maybe_Just {
var v_0 string = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 -> string
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_Common_trim(), gopurs_runtime.Str(v_0)).StrVal()
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if (__local_var_1_0) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(__local_var_1_0)})}
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1)
}

func Call_Data_String_NonEmpty_Internal_unsafeFromString(_dollar__unused_0_loop gopurs_runtime.Value, x_1_loop string) string {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var x_1 string = x_1_loop
_ = x_1
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Str(x_1).StrVal()) == ("") {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Str(gopurs_runtime.Str(x_1).StrVal())})}
}
end_branch_1:
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
var __local_var_2_0 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](__t1))}
var __t2 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t2 = (*Constructor_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0
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

func Call_Data_String_NonEmpty_Internal_endsWith(x_0_loop string, v_1_loop string) bool {
var x_0 string = x_0_loop
_ = x_0
var v_1 string = v_1_loop
_ = v_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_String_CodeUnits_stripSuffix(), gopurs_runtime.Str(gopurs_runtime.Str(x_0).StrVal()), gopurs_runtime.Str(v_1))
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

func Call_Data_String_NonEmpty_Internal_contains(x_0_loop string) gopurs_runtime.Value {
var x_0 string = x_0_loop
_ = x_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str(x_0))
_ = __local_var_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Str(v_2.StrVal()))
})
}

func Call_Data_String_NonEmpty_Internal_appendString(v_0_loop string, s2_1_loop string) string {
var v_0 string = v_0_loop
_ = v_0
var s2_1 string = s2_1_loop
_ = s2_1
return (v_0) + (s2_1)
}

func Call_Data_String_NonEmpty_Internal_joinWith__632268499(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, splice_1_loop string) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var splice_1 string = splice_1_loop
_ = splice_1
// TAST (Let): Semigroup0_2_1 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Monoid_monoidString(), "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Func(func(xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.Box(dictFoldable_0.V1), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
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
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.RecordGet(v_4, "acc"), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_2_1.V0), gopurs_runtime.Str(splice_1), v1_5)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.RecordGet(Get_Data_Monoid_monoidString(), "mempty"), gopurs_runtime.Bool(true)), xs_3), "acc")
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_0, x_3)
})
}

func Call_Data_String_NonEmpty_Internal_liftS__895676186(f_0_loop gopurs_runtime.Value, v_1_loop string) bool {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return (gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)).IntVal) != (0)
}

func Call_Data_String_NonEmpty_Internal_liftS__3230749042(f_0_loop gopurs_runtime.Value, v_1_loop string) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))
}

func Call_Data_String_NonEmpty_Internal_liftS__3241548146(f_0_loop gopurs_runtime.Value, v_1_loop string) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1)))
}



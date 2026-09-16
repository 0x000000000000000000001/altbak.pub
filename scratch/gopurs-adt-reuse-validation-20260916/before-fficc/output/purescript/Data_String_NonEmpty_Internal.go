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
		cache_Data_String_NonEmpty_Internal_fromJust = gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{})
	})
	return cache_Data_String_NonEmpty_Internal_fromJust
}

var cache_Data_String_NonEmpty_Internal_NonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_NonEmptyString sync.Once
func Get_Data_String_NonEmpty_Internal_NonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_NonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_Internal_NonEmptyString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_NonEmptyString(x_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_NonEmptyString
}

var cache_Data_String_NonEmpty_Internal_NonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_NonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_NonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_NonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_NonEmptyReplacement = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_NonEmpty_Internal_NonEmptyReplacement(x_0_box.StrVal()))
})
	})
	return cache_Data_String_NonEmpty_Internal_NonEmptyReplacement
}

var cache_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict sync.Once
func Get_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict.Do(func() {
		cache_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1987403114, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict(func() struct{
	nes gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	nes gopurs_runtime.Value
}{}
					clone.nes = gopurs_runtime.RecordGet(orig, "nes")
					return clone
				}()))}
})
	})
	return cache_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict
}

var cache_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict__2734224675 gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict__2734224675 sync.Once
func Get_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict__2734224675() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict__2734224675.Do(func() {
		cache_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict__2734224675 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1987403114, UnsafePtr: unsafe.Pointer(Call_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict__2734224675(func() struct{
	nes gopurs_runtime.Value
} {
					orig := x_0_box
					_ = orig
					clone := struct{
	nes gopurs_runtime.Value
}{}
					clone.nes = gopurs_runtime.RecordGet(orig, "nes")
					return clone
				}()))}
})
	})
	return cache_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict__2734224675
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
		cache_Data_String_NonEmpty_Internal_showNonEmptyString = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyString.unsafeFromString ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
})})))}
	})
	return cache_Data_String_NonEmpty_Internal_showNonEmptyString
}

var cache_Data_String_NonEmpty_Internal_showNonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_showNonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_showNonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_showNonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_showNonEmptyReplacement = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyReplacement (NonEmptyString.unsafeFromString ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + ("))"))
})})))}
	})
	return cache_Data_String_NonEmpty_Internal_showNonEmptyReplacement
}

var cache_Data_String_NonEmpty_Internal_semigroupNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_semigroupNonEmptyString sync.Once
func Get_Data_String_NonEmpty_Internal_semigroupNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_semigroupNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_Internal_semigroupNonEmptyString = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_443971153_4179793454(Rebox_Data_String_NonEmpty_Internal_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))}
	})
	return cache_Data_String_NonEmpty_Internal_semigroupNonEmptyString
}

var cache_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_semigroupNonEmptyReplacement = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_443971153_4179793454(Rebox_Data_String_NonEmpty_Internal_4179793454_443971153(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupString()))))}
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
		cache_Data_String_NonEmpty_Internal_ordNonEmptyString = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_2406510097_4177771502(Rebox_Data_String_NonEmpty_Internal_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString()))))}
	})
	return cache_Data_String_NonEmpty_Internal_ordNonEmptyString
}

var cache_Data_String_NonEmpty_Internal_ordNonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_ordNonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_ordNonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_ordNonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_ordNonEmptyReplacement = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_2406510097_4177771502(Rebox_Data_String_NonEmpty_Internal_4177771502_2406510097(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Get_Data_Ord_ordString()))))}
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
return Call_Data_String_NonEmpty_Internal_nes(gopurs_runtime.CoerceToStruct[Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_Data_String_NonEmpty_Internal_nes
}

var cache_Data_String_NonEmpty_Internal_makeNonEmptyBad gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_makeNonEmptyBad sync.Once
func Get_Data_String_NonEmpty_Internal_makeNonEmptyBad() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_makeNonEmptyBad.Do(func() {
		cache_Data_String_NonEmpty_Internal_makeNonEmptyBad = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_makeNonEmptyBad(_dollar___unused_0_box)
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

var cache_Data_String_NonEmpty_Internal_liftS__1684961947 gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_liftS__1684961947 sync.Once
func Get_Data_String_NonEmpty_Internal_liftS__1684961947() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_liftS__1684961947.Do(func() {
		cache_Data_String_NonEmpty_Internal_liftS__1684961947 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_Internal_liftS__1684961947(f_0_box, v_1_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_String_NonEmpty_Internal_liftS__1684961947
}

var cache_Data_String_NonEmpty_Internal_startsWith gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_startsWith sync.Once
func Get_Data_String_NonEmpty_Internal_startsWith() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_startsWith.Do(func() {
		cache_Data_String_NonEmpty_Internal_startsWith = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_Internal_liftS(), Get_Data_String_CodeUnits_startsWith())
	})
	return cache_Data_String_NonEmpty_Internal_startsWith
}

var cache_Data_String_NonEmpty_Internal_joinWith1 gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_joinWith1 sync.Once
func Get_Data_String_NonEmpty_Internal_joinWith1() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_joinWith1.Do(func() {
		cache_Data_String_NonEmpty_Internal_joinWith1 = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_joinWith1(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_Data_String_NonEmpty_Internal_joinWith1
}

var cache_Data_String_NonEmpty_Internal_joinWith gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_joinWith sync.Once
func Get_Data_String_NonEmpty_Internal_joinWith() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_joinWith.Do(func() {
		cache_Data_String_NonEmpty_Internal_joinWith = gopurs_runtime.Func2(func(dictFoldable_0_box gopurs_runtime.Value, splice_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_joinWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), splice_1_box.StrVal())
})
	})
	return cache_Data_String_NonEmpty_Internal_joinWith
}

var cache_Data_String_NonEmpty_Internal_join1With gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_join1With sync.Once
func Get_Data_String_NonEmpty_Internal_join1With() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_join1With.Do(func() {
		cache_Data_String_NonEmpty_Internal_join1With = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_join1With(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_Data_String_NonEmpty_Internal_join1With
}

var cache_Data_String_NonEmpty_Internal_fromString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_fromString sync.Once
func Get_Data_String_NonEmpty_Internal_fromString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_fromString.Do(func() {
		cache_Data_String_NonEmpty_Internal_fromString = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_Internal_fromString(v_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
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
return func() gopurs_runtime.Value {
				_v := Call_Data_String_NonEmpty_Internal_trim(v_0_box.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()
})
	})
	return cache_Data_String_NonEmpty_Internal_trim
}

var cache_Data_String_NonEmpty_Internal_unsafeFromString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_unsafeFromString sync.Once
func Get_Data_String_NonEmpty_Internal_unsafeFromString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_unsafeFromString.Do(func() {
		cache_Data_String_NonEmpty_Internal_unsafeFromString = gopurs_runtime.Func(func(_dollar___unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_Internal_unsafeFromString(_dollar___unused_0_box)
})
	})
	return cache_Data_String_NonEmpty_Internal_unsafeFromString
}

var cache_Data_String_NonEmpty_Internal_eqNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_eqNonEmptyString sync.Once
func Get_Data_String_NonEmpty_Internal_eqNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_eqNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_Internal_eqNonEmptyString = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_1140313009_3790796878(Rebox_Data_String_NonEmpty_Internal_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}
	})
	return cache_Data_String_NonEmpty_Internal_eqNonEmptyString
}

var cache_Data_String_NonEmpty_Internal_eqNonEmptyReplacement gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_eqNonEmptyReplacement sync.Once
func Get_Data_String_NonEmpty_Internal_eqNonEmptyReplacement() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_eqNonEmptyReplacement.Do(func() {
		cache_Data_String_NonEmpty_Internal_eqNonEmptyReplacement = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_1140313009_3790796878(Rebox_Data_String_NonEmpty_Internal_3790796878_1140313009(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Get_Data_Eq_eqString()))))}
	})
	return cache_Data_String_NonEmpty_Internal_eqNonEmptyReplacement
}

var cache_Data_String_NonEmpty_Internal_endsWith gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_endsWith sync.Once
func Get_Data_String_NonEmpty_Internal_endsWith() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_endsWith.Do(func() {
		cache_Data_String_NonEmpty_Internal_endsWith = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_Internal_liftS(), Get_Data_String_CodeUnits_endsWith())
	})
	return cache_Data_String_NonEmpty_Internal_endsWith
}

var cache_Data_String_NonEmpty_Internal_contains gopurs_runtime.Value
var once_Data_String_NonEmpty_Internal_contains sync.Once
func Get_Data_String_NonEmpty_Internal_contains() gopurs_runtime.Value {
	once_Data_String_NonEmpty_Internal_contains.Do(func() {
		cache_Data_String_NonEmpty_Internal_contains = gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_String_NonEmpty_Internal_liftS(), Get_Data_String_CodeUnits_contains())
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

type Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[T_s any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[1987403114] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value])(ptr)
		_ = c
		switch key {
		case "nes": return gopurs_runtime.Box(c.V0)
		default: panic("Key not found in dictionary Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty: " + key)
		}
	}
}


func Call_Data_String_NonEmpty_Internal_NonEmptyString(x_0_loop string) string {
var x_0 string = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_NonEmpty_Internal_NonEmptyReplacement(x_0_loop string) string {
var x_0 string = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict(x_0_loop struct{
	nes gopurs_runtime.Value
}) *Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value] {
var x_0 struct{
	nes gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("nes", orig.nes)
				}())
}

func Call_Data_String_NonEmpty_Internal_MakeNonEmpty_dollar_Dict__2734224675(x_0_loop struct{
	nes gopurs_runtime.Value
}) *Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value] {
MakeNonEmpty_dollar_Dict__2734224675:
for {
if false { continue MakeNonEmpty_dollar_Dict__2734224675 }
var x_0 struct{
	nes gopurs_runtime.Value
} = x_0_loop
_ = x_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				orig := x_0
				_ = orig
				return gopurs_runtime.RecordDict1("nes", orig.nes)
				}())
}
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
return gopurs_runtime.Value{Type: 9, IntVal: 1987403114, UnsafePtr: unsafe.Pointer((&Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(p_1.IntVal)), UnsafePtr: nil}).StrVal())
})}))}
}

func Call_Data_String_NonEmpty_Internal_nes(dict_0_loop *Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_Data_String_NonEmpty_Internal_makeNonEmptyBad(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
return gopurs_runtime.Value{Type: 9, IntVal: 1987403114, UnsafePtr: unsafe.Pointer((&Constructor_Data_String_NonEmpty_Internal_MakeNonEmpty[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
})}))}
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

func Call_Data_String_NonEmpty_Internal_liftS__1684961947(f_0_loop gopurs_runtime.Value, v_1_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
liftS__1684961947:
for {
if false { continue liftS__1684961947 }
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 string = v_1_loop
_ = v_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_742090555_3094389156(Rebox_Data_String_NonEmpty_Internal_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, gopurs_runtime.Str(v_1))))))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}
}

func Call_Data_String_NonEmpty_Internal_joinWith1(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeVar f$scope11)])
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}), gopurs_runtime.Apply(Call_Data_Foldable_intercalate(Foldable0_1_0, Rebox_Data_String_NonEmpty_Internal_1950344881_1201789390(Rebox_Data_String_NonEmpty_Internal_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString())))), gopurs_runtime.Str(v_2.StrVal())))
})
}

func Call_Data_String_NonEmpty_Internal_joinWith(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], splice_1_loop string) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var splice_1 string = splice_1_loop
_ = splice_1
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Data_Foldable_intercalate(dictFoldable_0, Rebox_Data_String_NonEmpty_Internal_1950344881_1201789390(Rebox_Data_String_NonEmpty_Internal_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString())))), gopurs_runtime.Str(splice_1)), Get_Unsafe_Coerce_unsafeCoerce())
}

func Call_Data_String_NonEmpty_Internal_join1With(dictFoldable1_0_loop *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
// TAST (Let): Foldable0_1_0 shape=App(Other) bindingType=(ADT ["Data","Foldable","Foldable"] [(TypeVar f$scope19)])
Foldable0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](gopurs_runtime.Apply(dictFoldable1_0.V0, gopurs_runtime.Value{}))
_ = Foldable0_1_0
return gopurs_runtime.Func(func(splice_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Call_Data_Foldable_intercalate(Foldable0_1_0, Rebox_Data_String_NonEmpty_Internal_1950344881_1201789390(Rebox_Data_String_NonEmpty_Internal_1201789390_1950344881(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](Get_Data_Monoid_monoidString())))), gopurs_runtime.Str(splice_2.StrVal())), Get_Unsafe_Coerce_unsafeCoerce()))
})
}

func Call_Data_String_NonEmpty_Internal_fromString(v_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 string = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Maybe_Just[string]
{
if (v_0) == ("") {
__t0 = Rebox_Data_String_NonEmpty_Internal_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_0
} else {

}
}
{
__t0 = Rebox_Data_String_NonEmpty_Internal_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Str(v_0), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_742090555_3094389156(__t0))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_Internal_stripPrefix(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
return gopurs_runtime.Apply2(Get_Control_Bind_composeKleisliFlipped__1603318047(), Get_Data_String_NonEmpty_Internal_fromString(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_742090555_3094389156(Rebox_Data_String_NonEmpty_Internal_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_String_CodeUnits_stripPrefix(pat_0, v_1.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
}))
}

func Call_Data_String_NonEmpty_Internal_stripSuffix(pat_0_loop string) gopurs_runtime.Value {
var pat_0 string = pat_0_loop
_ = pat_0
return gopurs_runtime.Apply2(Get_Control_Bind_composeKleisliFlipped__1603318047(), Get_Data_String_NonEmpty_Internal_fromString(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_742090555_3094389156(Rebox_Data_String_NonEmpty_Internal_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := Call_Data_String_CodeUnits_stripSuffix(pat_0, v_1.StrVal())
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))))}
}))
}

func Call_Data_String_NonEmpty_Internal_trim(v_0_loop string) struct{V0 gopurs_runtime.Value; V1 bool} {
var v_0 string = v_0_loop
_ = v_0
// TAST (Let): __local_var_1_0 shape=App(Var) bindingType=Any
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_String_Common_trim(), gopurs_runtime.Str(v_0))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[string]
{
if (__local_var_1_0.StrVal()) == ("") {
__t1 = Rebox_Data_String_NonEmpty_Internal_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
goto end_branch_1
} else {

}
}
{
__t1 = Rebox_Data_String_NonEmpty_Internal_3094389156_742090555(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Str(__local_var_1_0.StrVal()), true}
				if _v.V1 {
					return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{Rc: 1, V0: _v.V0})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 930809136}
			}()))
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_String_NonEmpty_Internal_742090555_3094389156(__t1))}
				if _v.Type == 9 && _v.IntVal == 930809136 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_String_NonEmpty_Internal_unsafeFromString(_dollar___unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
return gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), gopurs_runtime.Apply(Get_Data_Maybe_fromJust(), gopurs_runtime.Value{}), Get_Data_String_NonEmpty_Internal_fromString())
}

func Call_Data_String_NonEmpty_Internal_appendString(v_0_loop string, s2_1_loop string) string {
var v_0 string = v_0_loop
_ = v_0
var s2_1 string = s2_1_loop
_ = s2_1
return (v_0) + (s2_1)
}

func Rebox_Data_String_NonEmpty_Internal_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_NonEmpty_Internal_1201789390_1950344881(in *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) *Constructor_Data_Monoid_Monoid[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[string]{}
		out.V0 = in.V0
		out.V1 = in.V1.StrVal()
	return out
}

func Rebox_Data_String_NonEmpty_Internal_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_NonEmpty_Internal_1950344881_1201789390(in *Constructor_Data_Monoid_Monoid[string]) *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Str(in.V1)
	return out
}

func Rebox_Data_String_NonEmpty_Internal_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_String_NonEmpty_Internal_3094389156_742090555(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[string]{}
		out.V0 = in.V0.StrVal()
	return out
}

func Rebox_Data_String_NonEmpty_Internal_3790796878_1140313009(in *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[string]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_NonEmpty_Internal_4177771502_2406510097(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[string]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_String_NonEmpty_Internal_4179793454_443971153(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[string] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[string]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_NonEmpty_Internal_443971153_4179793454(in *Constructor_Data_Semigroup_Semigroup[string]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_NonEmpty_Internal_742090555_3094389156(in *Constructor_Data_Maybe_Just[string]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Str(in.V0)
	return out
}



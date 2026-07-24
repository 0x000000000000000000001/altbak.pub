package Data_String_NonEmpty_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	unsafe "unsafe"
)

var NonEmptyString gopurs_runtime.Value
var once_NonEmptyString sync.Once
func Get_NonEmptyString() gopurs_runtime.Value {
	once_NonEmptyString.Do(func() {
		NonEmptyString = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return NonEmptyString
}

var NonEmptyReplacement gopurs_runtime.Value
var once_NonEmptyReplacement sync.Once
func Get_NonEmptyReplacement() gopurs_runtime.Value {
	once_NonEmptyReplacement.Do(func() {
		NonEmptyReplacement = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return NonEmptyReplacement
}

var toUpper gopurs_runtime.Value
var once_toUpper sync.Once
func Get_toUpper() gopurs_runtime.Value {
	once_toUpper.Do(func() {
		toUpper = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(pkg_Data_String_Common.Get_toUpper(), v_0)
}()
})
	})
	return toUpper
}

var toString gopurs_runtime.Value
var once_toString sync.Once
func Get_toString() gopurs_runtime.Value {
	once_toString.Do(func() {
		toString = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}()
})
	})
	return toString
}

var toLower gopurs_runtime.Value
var once_toLower sync.Once
func Get_toLower() gopurs_runtime.Value {
	once_toLower.Do(func() {
		toLower = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0)
}()
})
	})
	return toLower
}

var showNonEmptyString gopurs_runtime.Value
var once_showNonEmptyString sync.Once
func Get_showNonEmptyString() gopurs_runtime.Value {
	once_showNonEmptyString.Do(func() {
		showNonEmptyString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(NonEmptyString.unsafeFromString " + gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0).StrVal() + ")")
}))
	})
	return showNonEmptyString
}

var showNonEmptyReplacement gopurs_runtime.Value
var once_showNonEmptyReplacement sync.Once
func Get_showNonEmptyReplacement() gopurs_runtime.Value {
	once_showNonEmptyReplacement.Do(func() {
		showNonEmptyReplacement = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(NonEmptyReplacement (NonEmptyString.unsafeFromString " + gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0).StrVal() + "))")
}))
	})
	return showNonEmptyReplacement
}

var semigroupNonEmptyString gopurs_runtime.Value
var once_semigroupNonEmptyString sync.Once
func Get_semigroupNonEmptyString() gopurs_runtime.Value {
	once_semigroupNonEmptyString.Do(func() {
		semigroupNonEmptyString = pkg_Data_Semigroup.Get_semigroupString()
	})
	return semigroupNonEmptyString
}

var semigroupNonEmptyReplacement gopurs_runtime.Value
var once_semigroupNonEmptyReplacement sync.Once
func Get_semigroupNonEmptyReplacement() gopurs_runtime.Value {
	once_semigroupNonEmptyReplacement.Do(func() {
		semigroupNonEmptyReplacement = pkg_Data_Semigroup.Get_semigroupString()
	})
	return semigroupNonEmptyReplacement
}

var replaceAll gopurs_runtime.Value
var once_replaceAll sync.Once
func Get_replaceAll() gopurs_runtime.Value {
	once_replaceAll.Do(func() {
		replaceAll = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replaceAll(pat_0_box, v_1_box, v1_2_box)
})
	})
	return replaceAll
}

var replace gopurs_runtime.Value
var once_replace sync.Once
func Get_replace() gopurs_runtime.Value {
	once_replace.Do(func() {
		replace = gopurs_runtime.Func3(func(pat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value, v1_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_replace(pat_0_box, v_1_box, v1_2_box)
})
	})
	return replace
}

var prependString gopurs_runtime.Value
var once_prependString sync.Once
func Get_prependString() gopurs_runtime.Value {
	once_prependString.Do(func() {
		prependString = gopurs_runtime.Func2(func(s1_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_prependString(s1_0_box, v_1_box))
})
	})
	return prependString
}

var ordNonEmptyString gopurs_runtime.Value
var once_ordNonEmptyString sync.Once
func Get_ordNonEmptyString() gopurs_runtime.Value {
	once_ordNonEmptyString.Do(func() {
		ordNonEmptyString = pkg_Data_Ord.Get_ordString()
	})
	return ordNonEmptyString
}

var ordNonEmptyReplacement gopurs_runtime.Value
var once_ordNonEmptyReplacement sync.Once
func Get_ordNonEmptyReplacement() gopurs_runtime.Value {
	once_ordNonEmptyReplacement.Do(func() {
		ordNonEmptyReplacement = pkg_Data_Ord.Get_ordString()
	})
	return ordNonEmptyReplacement
}

var nonEmptyNonEmpty gopurs_runtime.Value
var once_nonEmptyNonEmpty sync.Once
func Get_nonEmptyNonEmpty() gopurs_runtime.Value {
	once_nonEmptyNonEmpty.Do(func() {
		nonEmptyNonEmpty = gopurs_runtime.Func(func(dictIsSymbol_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictIsSymbol_0 gopurs_runtime.Value = dictIsSymbol_0_loop
_ = dictIsSymbol_0
return gopurs_runtime.RecordDict1("nes", gopurs_runtime.Func(func(p_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictIsSymbol_0, "reflectSymbol"), p_1)
}))
}()
})
	})
	return nonEmptyNonEmpty
}

var nes gopurs_runtime.Value
var once_nes sync.Once
func Get_nes() gopurs_runtime.Value {
	once_nes.Do(func() {
		nes = gopurs_runtime.Func(func(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "nes")
}()
})
	})
	return nes
}

var makeNonEmptyBad gopurs_runtime.Value
var once_makeNonEmptyBad sync.Once
func Get_makeNonEmptyBad() gopurs_runtime.Value {
	once_makeNonEmptyBad.Do(func() {
		makeNonEmptyBad = gopurs_runtime.Func(func(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return gopurs_runtime.RecordDict1("nes", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("")
}))
}()
})
	})
	return makeNonEmptyBad
}

var localeCompare gopurs_runtime.Value
var once_localeCompare sync.Once
func Get_localeCompare() gopurs_runtime.Value {
	once_localeCompare.Do(func() {
		localeCompare = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_localeCompare(v_0_box, v1_1_box)
})
	})
	return localeCompare
}

var liftS gopurs_runtime.Value
var once_liftS sync.Once
func Get_liftS() gopurs_runtime.Value {
	once_liftS.Do(func() {
		liftS = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftS(f_0_box, v_1_box)
})
	})
	return liftS
}

var startsWith gopurs_runtime.Value
var once_startsWith sync.Once
func Get_startsWith() gopurs_runtime.Value {
	once_startsWith.Do(func() {
		startsWith = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_startsWith(x_0_box, v_1_box)
})
	})
	return startsWith
}

var joinWith1 gopurs_runtime.Value
var once_joinWith1 sync.Once
func Get_joinWith1() gopurs_runtime.Value {
	once_joinWith1.Do(func() {
		joinWith1 = gopurs_runtime.Func(func(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(sep_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldl"), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet(v_4, "init").IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), v1_5)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), gopurs_runtime.Str(gopurs_runtime.RecordGet(v_4, "acc").StrVal() + sep_2.StrVal() + v1_5.StrVal()))
}
end_branch_1:
return __t1
}), gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(true), gopurs_runtime.Str("")), xs_3), "acc")
})
}()
})
	})
	return joinWith1
}

var joinWith gopurs_runtime.Value
var once_joinWith sync.Once
func Get_joinWith() gopurs_runtime.Value {
	once_joinWith.Do(func() {
		joinWith = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, splice_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinWith(dictFoldable_0_box, splice_1_box, xs_2_box)
})
	})
	return joinWith
}

var join1With gopurs_runtime.Value
var once_join1With sync.Once
func Get_join1With() gopurs_runtime.Value {
	once_join1With.Do(func() {
		join1With = gopurs_runtime.Func(func(dictFoldable1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictFoldable1_0 gopurs_runtime.Value = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable1_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_0
return gopurs_runtime.Func2(func(splice_2 gopurs_runtime.Value, xs_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldl"), gopurs_runtime.Func2(func(v_4 gopurs_runtime.Value, v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet(v_4, "init").IntVal != 0 {
__t1 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), v1_5)
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), gopurs_runtime.Str(gopurs_runtime.RecordGet(v_4, "acc").StrVal() + splice_2.StrVal() + v1_5.StrVal()))
}
end_branch_1:
return __t1
}), gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(true), gopurs_runtime.Str("")), xs_3), "acc")
})
}()
})
	})
	return join1With
}

var fromString gopurs_runtime.Value
var once_fromString sync.Once
func Get_fromString() gopurs_runtime.Value {
	once_fromString.Do(func() {
		fromString = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if v_0.StrVal() == "" {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{v_0})}
}
end_branch_0:
return __t0
}()
})
	})
	return fromString
}

var stripPrefix gopurs_runtime.Value
var once_stripPrefix sync.Once
func Get_stripPrefix() gopurs_runtime.Value {
	once_stripPrefix.Do(func() {
		stripPrefix = gopurs_runtime.Func2(func(pat_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripPrefix(pat_0_box, a_1_box)
})
	})
	return stripPrefix
}

var stripSuffix gopurs_runtime.Value
var once_stripSuffix sync.Once
func Get_stripSuffix() gopurs_runtime.Value {
	once_stripSuffix.Do(func() {
		stripSuffix = gopurs_runtime.Func2(func(pat_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_stripSuffix(pat_0_box, a_1_box)
})
	})
	return stripSuffix
}

var trim gopurs_runtime.Value
var once_trim sync.Once
func Get_trim() gopurs_runtime.Value {
	once_trim.Do(func() {
		trim = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
__local_var_1_0 := gopurs_runtime.Apply(pkg_Data_String_Common.Get_trim(), v_0)
_ = __local_var_1_0
var __t1 gopurs_runtime.Value
{
if __local_var_1_0.StrVal() == "" {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{__local_var_1_0})}
}
end_branch_1:
return __t1
}()
})
	})
	return trim
}

var unsafeFromString gopurs_runtime.Value
var once_unsafeFromString sync.Once
func Get_unsafeFromString() gopurs_runtime.Value {
	once_unsafeFromString.Do(func() {
		unsafeFromString = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeFromString(_dollar__unused_0_box, x_1_box)
})
	})
	return unsafeFromString
}

var eqNonEmptyString gopurs_runtime.Value
var once_eqNonEmptyString sync.Once
func Get_eqNonEmptyString() gopurs_runtime.Value {
	once_eqNonEmptyString.Do(func() {
		eqNonEmptyString = pkg_Data_Eq.Get_eqString()
	})
	return eqNonEmptyString
}

var eqNonEmptyReplacement gopurs_runtime.Value
var once_eqNonEmptyReplacement sync.Once
func Get_eqNonEmptyReplacement() gopurs_runtime.Value {
	once_eqNonEmptyReplacement.Do(func() {
		eqNonEmptyReplacement = pkg_Data_Eq.Get_eqString()
	})
	return eqNonEmptyReplacement
}

var endsWith gopurs_runtime.Value
var once_endsWith sync.Once
func Get_endsWith() gopurs_runtime.Value {
	once_endsWith.Do(func() {
		endsWith = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_endsWith(x_0_box, v_1_box)
})
	})
	return endsWith
}

var contains gopurs_runtime.Value
var once_contains sync.Once
func Get_contains() gopurs_runtime.Value {
	once_contains.Do(func() {
		contains = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(pkg_Data_String_CodeUnits.Get_contains(), x_0)
}()
})
	})
	return contains
}

var appendString gopurs_runtime.Value
var once_appendString sync.Once
func Get_appendString() gopurs_runtime.Value {
	once_appendString.Do(func() {
		appendString = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s2_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_appendString(v_0_box, s2_1_box))
})
	})
	return appendString
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

func Call_prependString(s1_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) string {
var s1_0 gopurs_runtime.Value = s1_0_loop
_ = s1_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return s1_0.StrVal() + v_1.StrVal()
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
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 42808261) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 1354639136) {
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

func Call_joinWith(dictFoldable_0_loop gopurs_runtime.Value, splice_1_loop gopurs_runtime.Value, xs_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
var splice_1 gopurs_runtime.Value = splice_1_loop
_ = splice_1
var xs_2 gopurs_runtime.Value = xs_2_loop
_ = xs_2
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(v_3 gopurs_runtime.Value, v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if gopurs_runtime.RecordGet(v_3, "init").IntVal != 0 {
__t0 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), v1_4)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(false), gopurs_runtime.Str(gopurs_runtime.RecordGet(v_3, "acc").StrVal() + splice_1.StrVal() + v1_4.StrVal()))
}
end_branch_0:
return __t0
}), gopurs_runtime.RecordDict2("init", "acc", gopurs_runtime.Bool(true), gopurs_runtime.Str("")), xs_2), "acc")
}

func Call_stripPrefix(pat_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_stripPrefix(), pat_0, a_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 1354639136) {
var __t2 gopurs_runtime.Value
{
if (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0.StrVal() == "" {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0})}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 42808261) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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

func Call_stripSuffix(pat_0_loop gopurs_runtime.Value, a_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var pat_0 gopurs_runtime.Value = pat_0_loop
_ = pat_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_stripSuffix(), pat_0, a_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 1354639136) {
var __t2 gopurs_runtime.Value
{
if (*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0.StrVal() == "" {
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Value{Type: 9, IntVal: 1354639136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Just{(*pkg_Data_Maybe.Data_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0})}
}
end_branch_2:
__t1 = __t2
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 42808261) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 42808261, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Data_Data_Maybe_Nothing{})}
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

func Call_unsafeFromString(_dollar__unused_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var __t0 gopurs_runtime.Value
{
if x_1.StrVal() == "" {
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
goto end_branch_0
} else {

}
}
{
__t0 = x_1
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
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 42808261) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 1354639136) {
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

func Call_appendString(v_0_loop gopurs_runtime.Value, s2_1_loop gopurs_runtime.Value) string {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s2_1 gopurs_runtime.Value = s2_1_loop
_ = s2_1
return v_0.StrVal() + s2_1.StrVal()
}



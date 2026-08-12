package Data_String_Regex

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_showRegex gopurs_runtime.Value
var once_showRegex sync.Once
func Get_showRegex() gopurs_runtime.Value {
	once_showRegex.Do(func() {
		cache_showRegex = gopurs_runtime.RecordDict1("show", Get_showRegexImpl())
	})
	return cache_showRegex
}

var cache_search gopurs_runtime.Value
var once_search sync.Once
func Get_search() gopurs_runtime.Value {
	once_search.Do(func() {
		cache_search = gopurs_runtime.Apply2(Get__search(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_search
}

var cache_replace_prime gopurs_runtime.Value
var once_replace_prime sync.Once
func Get_replace_prime() gopurs_runtime.Value {
	once_replace_prime.Do(func() {
		cache_replace_prime = gopurs_runtime.Apply2(Get__replaceBy(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_replace_prime
}

var cache_renderFlags gopurs_runtime.Value
var once_renderFlags sync.Once
func Get_renderFlags() gopurs_runtime.Value {
	once_renderFlags.Do(func() {
		cache_renderFlags = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_renderFlags(v_0_box))
})
	})
	return cache_renderFlags
}

var cache_regex gopurs_runtime.Value
var once_regex sync.Once
func Get_regex() gopurs_runtime.Value {
	once_regex.Do(func() {
		cache_regex = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_regex(s_0_box.StrVal(), f_1_box)
})
	})
	return cache_regex
}

var cache_parseFlags gopurs_runtime.Value
var once_parseFlags sync.Once
func Get_parseFlags() gopurs_runtime.Value {
	once_parseFlags.Do(func() {
		cache_parseFlags = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_parseFlags(s_0_box.StrVal())
})
	})
	return cache_parseFlags
}

var cache_match gopurs_runtime.Value
var once_match sync.Once
func Get_match() gopurs_runtime.Value {
	once_match.Do(func() {
		cache_match = gopurs_runtime.Apply2(Get__match(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_match
}

var cache_flags gopurs_runtime.Value
var once_flags sync.Once
func Get_flags() gopurs_runtime.Value {
	once_flags.Do(func() {
		cache_flags = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flags(x_0_box)
})
	})
	return cache_flags
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

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
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

func Call_renderFlags(v_0_loop gopurs_runtime.Value) string {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 string
{
if (gopurs_runtime.RecordGet(v_0, "global").IntVal) != (0) {
__t0 = "g"
goto end_branch_0
} else {

}
}
{
__t0 = ""
}
end_branch_0:
var __t1 string
{
if (gopurs_runtime.RecordGet(v_0, "ignoreCase").IntVal) != (0) {
__t1 = "i"
goto end_branch_1
} else {

}
}
{
__t1 = ""
}
end_branch_1:
var __t2 string
{
if (gopurs_runtime.RecordGet(v_0, "multiline").IntVal) != (0) {
__t2 = "m"
goto end_branch_2
} else {

}
}
{
__t2 = ""
}
end_branch_2:
var __t3 string
{
if (gopurs_runtime.RecordGet(v_0, "dotAll").IntVal) != (0) {
__t3 = "s"
goto end_branch_3
} else {

}
}
{
__t3 = ""
}
end_branch_3:
var __t4 string
{
if (gopurs_runtime.RecordGet(v_0, "sticky").IntVal) != (0) {
__t4 = "y"
goto end_branch_4
} else {

}
}
{
__t4 = ""
}
end_branch_4:
var __t5 string
{
if (gopurs_runtime.RecordGet(v_0, "unicode").IntVal) != (0) {
__t5 = "u"
goto end_branch_5
} else {

}
}
{
__t5 = ""
}
end_branch_5:
return gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(__t0), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(__t1), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(__t2), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(__t3), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(__t4), gopurs_runtime.Str(__t5)))))).StrVal()
}

func Call_regex(s_0_loop string, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply4(Get_regexImpl(), pkg_Data_Either.Get_Left(), pkg_Data_Either.Get_Right(), gopurs_runtime.Str(s_0), gopurs_runtime.Str(Call_renderFlags(f_1)))
}

func Call_parseFlags(s_0_loop string) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool((gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("s"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("g"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("i"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("m"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("y"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("u"), gopurs_runtime.Str(s_0)).IntVal) != (0))})
}

func Call_flags(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_flagsImpl(), x_0)
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get__match() gopurs_runtime.Value {
	return _Gopurs__Match
}

func Get__replaceBy() gopurs_runtime.Value {
	return _Gopurs__ReplaceBy
}

func Get__search() gopurs_runtime.Value {
	return _Gopurs__Search
}

func Get_flagsImpl() gopurs_runtime.Value {
	return _Gopurs_FlagsImpl
}

func Get_regexImpl() gopurs_runtime.Value {
	return _Gopurs_RegexImpl
}

func Get_replace() gopurs_runtime.Value {
	return _Gopurs_Replace
}

func Get_showRegexImpl() gopurs_runtime.Value {
	return _Gopurs_ShowRegexImpl
}

func Get_source() gopurs_runtime.Value {
	return _Gopurs_Source
}

func Get_split() gopurs_runtime.Value {
	return _Gopurs_Split
}

func Get_test() gopurs_runtime.Value {
	return _Gopurs_Test
}

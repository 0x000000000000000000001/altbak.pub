package Data_String_Regex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
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
		cache_search = gopurs_runtime.Apply2(Get__search(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
	})
	return cache_search
}

var cache_replace_prime gopurs_runtime.Value
var once_replace_prime sync.Once
func Get_replace_prime() gopurs_runtime.Value {
	once_replace_prime.Do(func() {
		cache_replace_prime = gopurs_runtime.Apply2(Get__replaceBy(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
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
		cache_match = gopurs_runtime.Apply2(Get__match(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})
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

func Call_renderFlags(v_0_loop gopurs_runtime.Value) string {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_0, "global").IntVal) != (0) {
__t0 = gopurs_runtime.Str("g")
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Str("")
}
end_branch_0:
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_0, "ignoreCase").IntVal) != (0) {
__t1 = gopurs_runtime.Str("i")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Str("")
}
end_branch_1:
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_0, "multiline").IntVal) != (0) {
__t2 = gopurs_runtime.Str("m")
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Str("")
}
end_branch_2:
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_0, "dotAll").IntVal) != (0) {
__t3 = gopurs_runtime.Str("s")
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Str("")
}
end_branch_3:
var __t4 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_0, "sticky").IntVal) != (0) {
__t4 = gopurs_runtime.Str("y")
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Str("")
}
end_branch_4:
var __t5 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_0, "unicode").IntVal) != (0) {
__t5 = gopurs_runtime.Str("u")
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Str("")
}
end_branch_5:
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), __t0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), __t1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), __t2, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), __t3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), __t4, __t5))))).StrVal()
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

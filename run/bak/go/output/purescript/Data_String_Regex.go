package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Regex_showRegex gopurs_runtime.Value
var once_Data_String_Regex_showRegex sync.Once
func Get_Data_String_Regex_showRegex() gopurs_runtime.Value {
	once_Data_String_Regex_showRegex.Do(func() {
		cache_Data_String_Regex_showRegex = gopurs_runtime.RecordDict1("show", Get_Data_String_Regex_showRegexImpl())
	})
	return cache_Data_String_Regex_showRegex
}

var cache_Data_String_Regex_search gopurs_runtime.Value
var once_Data_String_Regex_search sync.Once
func Get_Data_String_Regex_search() gopurs_runtime.Value {
	once_Data_String_Regex_search.Do(func() {
		cache_Data_String_Regex_search = gopurs_runtime.Apply2(Get_Data_String_Regex__search(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_String_Regex_search
}

var cache_Data_String_Regex_replace_prime gopurs_runtime.Value
var once_Data_String_Regex_replace_prime sync.Once
func Get_Data_String_Regex_replace_prime() gopurs_runtime.Value {
	once_Data_String_Regex_replace_prime.Do(func() {
		cache_Data_String_Regex_replace_prime = gopurs_runtime.Apply2(Get_Data_String_Regex__replaceBy(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_String_Regex_replace_prime
}

var cache_Data_String_Regex_renderFlags gopurs_runtime.Value
var once_Data_String_Regex_renderFlags sync.Once
func Get_Data_String_Regex_renderFlags() gopurs_runtime.Value {
	once_Data_String_Regex_renderFlags.Do(func() {
		cache_Data_String_Regex_renderFlags = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_Regex_renderFlags(v_0_box))
})
	})
	return cache_Data_String_Regex_renderFlags
}

var cache_Data_String_Regex_regex gopurs_runtime.Value
var once_Data_String_Regex_regex sync.Once
func Get_Data_String_Regex_regex() gopurs_runtime.Value {
	once_Data_String_Regex_regex.Do(func() {
		cache_Data_String_Regex_regex = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Regex_regex(s_0_box.StrVal(), f_1_box)
})
	})
	return cache_Data_String_Regex_regex
}

var cache_Data_String_Regex_parseFlags gopurs_runtime.Value
var once_Data_String_Regex_parseFlags sync.Once
func Get_Data_String_Regex_parseFlags() gopurs_runtime.Value {
	once_Data_String_Regex_parseFlags.Do(func() {
		cache_Data_String_Regex_parseFlags = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Regex_parseFlags(s_0_box.StrVal())
})
	})
	return cache_Data_String_Regex_parseFlags
}

var cache_Data_String_Regex_match gopurs_runtime.Value
var once_Data_String_Regex_match sync.Once
func Get_Data_String_Regex_match() gopurs_runtime.Value {
	once_Data_String_Regex_match.Do(func() {
		cache_Data_String_Regex_match = gopurs_runtime.Apply2(Get_Data_String_Regex__match(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))})
	})
	return cache_Data_String_Regex_match
}

var cache_Data_String_Regex_flags gopurs_runtime.Value
var once_Data_String_Regex_flags sync.Once
func Get_Data_String_Regex_flags() gopurs_runtime.Value {
	once_Data_String_Regex_flags.Do(func() {
		cache_Data_String_Regex_flags = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Regex_flags(x_0_box)
})
	})
	return cache_Data_String_Regex_flags
}

func Call_Data_String_Regex_renderFlags(v_0_loop gopurs_runtime.Value) string {
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
return (((((__t0) + (__t1)) + (__t2)) + (__t3)) + (__t4)) + (__t5)
}

func Call_Data_String_Regex_regex(s_0_loop string, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var __t0 string
{
if (gopurs_runtime.RecordGet(f_1, "global").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(f_1, "ignoreCase").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(f_1, "multiline").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(f_1, "dotAll").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(f_1, "sticky").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(f_1, "unicode").IntVal) != (0) {
__t5 = "u"
goto end_branch_5
} else {

}
}
{
__t5 = ""
}
end_branch_5:
return gopurs_runtime.Apply4(Get_Data_String_Regex_regexImpl(), Get_Data_Either_Left(), Get_Data_Either_Right(), gopurs_runtime.Str(s_0), gopurs_runtime.Str((((((__t0) + (__t1)) + (__t2)) + (__t3)) + (__t4)) + (__t5)))
}

func Call_Data_String_Regex_parseFlags(s_0_loop string) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("s"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("g"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("i"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("m"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("y"), gopurs_runtime.Str(s_0)).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply2(Get_Data_String_CodeUnits_contains(), gopurs_runtime.Str("u"), gopurs_runtime.Str(s_0)).IntVal) != (0))})
}

func Call_Data_String_Regex_flags(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.Apply(Get_Data_String_Regex_flagsImpl(), x_0)
}

func Get_Data_String_Regex__match() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex__Match
}

func Get_Data_String_Regex__replaceBy() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex__ReplaceBy
}

func Get_Data_String_Regex__search() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex__Search
}

func Get_Data_String_Regex_flagsImpl() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_FlagsImpl
}

func Get_Data_String_Regex_regexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_RegexImpl
}

func Get_Data_String_Regex_replace() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_Replace
}

func Get_Data_String_Regex_showRegexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_ShowRegexImpl
}

func Get_Data_String_Regex_source() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_Source
}

func Get_Data_String_Regex_split() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_Split
}

func Get_Data_String_Regex_test() gopurs_runtime.Value {
	return _Gopurs_Data_String_Regex_Test
}

package Data_String_Regex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
)

var showRegex gopurs_runtime.Value
var once_showRegex sync.Once
func Get_showRegex() gopurs_runtime.Value {
	once_showRegex.Do(func() {
		showRegex = gopurs_runtime.RecordDict1("show", Get_showRegexImpl())
	})
	return showRegex
}

var search gopurs_runtime.Value
var once_search sync.Once
func Get_search() gopurs_runtime.Value {
	once_search.Do(func() {
		search = gopurs_runtime.Apply2(Get__search(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
	})
	return search
}

var replace_prime gopurs_runtime.Value
var once_replace_prime sync.Once
func Get_replace_prime() gopurs_runtime.Value {
	once_replace_prime.Do(func() {
		replace_prime = gopurs_runtime.Apply2(Get__replaceBy(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
	})
	return replace_prime
}

var renderFlags gopurs_runtime.Value
var once_renderFlags sync.Once
func Get_renderFlags() gopurs_runtime.Value {
	once_renderFlags.Do(func() {
		renderFlags = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_0, "global")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(v_0, "ignoreCase")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(v_0, "multiline")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(v_0, "dotAll")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(v_0, "sticky")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(v_0, "unicode")).IntVal != 0 {
__t5 = gopurs_runtime.Str("u")
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Str("")
}
end_branch_5:
return gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(__t0.StrVal + __t1.StrVal).StrVal + __t2.StrVal).StrVal + __t3.StrVal).StrVal + __t4.StrVal).StrVal + __t5.StrVal)
})
	})
	return renderFlags
}

var regex gopurs_runtime.Value
var once_regex sync.Once
func Get_regex() gopurs_runtime.Value {
	once_regex.Do(func() {
		regex = gopurs_runtime.Func2(func(s_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(f_1, "global")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(f_1, "ignoreCase")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(f_1, "multiline")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(f_1, "dotAll")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(f_1, "sticky")).IntVal != 0 {
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
if (gopurs_runtime.RecordGet(f_1, "unicode")).IntVal != 0 {
__t5 = gopurs_runtime.Str("u")
goto end_branch_5
} else {

}
}
{
__t5 = gopurs_runtime.Str("")
}
end_branch_5:
return gopurs_runtime.Apply4(Get_regexImpl(), pkg_Data_Either.Get_Left(), pkg_Data_Either.Get_Right(), s_0, gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str(__t0.StrVal + __t1.StrVal).StrVal + __t2.StrVal).StrVal + __t3.StrVal).StrVal + __t4.StrVal).StrVal + __t5.StrVal))
})
	})
	return regex
}

var parseFlags gopurs_runtime.Value
var once_parseFlags sync.Once
func Get_parseFlags() gopurs_runtime.Value {
	once_parseFlags.Do(func() {
		parseFlags = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"global", "ignoreCase", "multiline", "dotAll", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("g"), s_0), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("i"), s_0), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("m"), s_0), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("s"), s_0), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("y"), s_0), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("u"), s_0)})
})
	})
	return parseFlags
}

var match gopurs_runtime.Value
var once_match sync.Once
func Get_match() gopurs_runtime.Value {
	once_match.Do(func() {
		match = gopurs_runtime.Apply2(Get__match(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.RecordDict1("_tag", gopurs_runtime.Str("Nothing")))
	})
	return match
}

var flags gopurs_runtime.Value
var once_flags sync.Once
func Get_flags() gopurs_runtime.Value {
	once_flags.Do(func() {
		flags = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_flagsImpl(), x_0)
})
	})
	return flags
}

func Get__match() gopurs_runtime.Value {
	return X_Match
}

func Get__replaceBy() gopurs_runtime.Value {
	return X_ReplaceBy
}

func Get__search() gopurs_runtime.Value {
	return X_Search
}

func Get_flagsImpl() gopurs_runtime.Value {
	return FlagsImpl
}

func Get_regexImpl() gopurs_runtime.Value {
	return RegexImpl
}

func Get_replace() gopurs_runtime.Value {
	return Replace
}

func Get_showRegexImpl() gopurs_runtime.Value {
	return ShowRegexImpl
}

func Get_source() gopurs_runtime.Value {
	return Source
}

func Get_split() gopurs_runtime.Value {
	return Split
}

func Get_test() gopurs_runtime.Value {
	return Test
}

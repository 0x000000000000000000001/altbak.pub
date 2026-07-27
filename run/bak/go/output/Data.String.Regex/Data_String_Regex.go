package Data_String_Regex

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_String_CodeUnits "gopurs/output/Data.String.CodeUnits"
	unsafe "unsafe"
)

var cache_showRegex gopurs_runtime.Value
var once_showRegex sync.Once
func Get_showRegex() gopurs_runtime.Value {
	once_showRegex.Do(func() {
		cache_showRegex = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", Get_showRegexImpl())))
	})
	return cache_showRegex
}

var cache_search gopurs_runtime.Value
var once_search sync.Once
func Get_search() gopurs_runtime.Value {
	once_search.Do(func() {
		cache_search = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 gopurs_runtime.Value, inner_arg1 string) *pkg_Data_Maybe.Constructor_Just[int64] {
return (*pkg_Data_Maybe.Constructor_Just[int64])(gopurs_runtime.Apply2(gopurs_runtime.Apply2(Get__search(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), inner_arg0, gopurs_runtime.Str(inner_arg1)).UnsafePtr)
}(arg0, arg1.StrVal()))}
})
	})
	return cache_search
}

var cache_replace_prime gopurs_runtime.Value
var once_replace_prime sync.Once
func Get_replace_prime() gopurs_runtime.Value {
	once_replace_prime.Do(func() {
		cache_replace_prime = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(func(inner_arg0 gopurs_runtime.Value, inner_arg1 func(string, []*pkg_Data_Maybe.Constructor_Just[string]) string, inner_arg2 string) string {
return gopurs_runtime.Apply3(gopurs_runtime.Apply2(Get__replaceBy(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), inner_arg0, gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(inner_arg1(arg0.StrVal(), func() []*pkg_Data_Maybe.Constructor_Just[string] {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]*pkg_Data_Maybe.Constructor_Just[string], len(arr))
					for i, v := range arr { unboxed[i] = (*pkg_Data_Maybe.Constructor_Just[string])(v.UnsafePtr) }
					return unboxed
				}()))
}), gopurs_runtime.Str(inner_arg2)).StrVal()
}(arg0, func(inner_arg0 string, inner_arg1 []*pkg_Data_Maybe.Constructor_Just[string]) string {
return gopurs_runtime.Apply2(arg1, gopurs_runtime.Str(inner_arg0), func() gopurs_runtime.Value {
					arr := inner_arg1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal()
}, arg2.StrVal()))
})
	})
	return cache_replace_prime
}

var cache_renderFlags gopurs_runtime.Value
var once_renderFlags sync.Once
func Get_renderFlags() gopurs_runtime.Value {
	once_renderFlags.Do(func() {
		cache_renderFlags = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_renderFlags(gopurs_runtime.UnboxAny(v_0_box)))
})
	})
	return cache_renderFlags
}

var cache_regex gopurs_runtime.Value
var once_regex sync.Once
func Get_regex() gopurs_runtime.Value {
	once_regex.Do(func() {
		cache_regex = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_regex(s_0_box.StrVal(), gopurs_runtime.UnboxAny(f_1_box))
})
	})
	return cache_regex
}

var cache_parseFlags gopurs_runtime.Value
var once_parseFlags sync.Once
func Get_parseFlags() gopurs_runtime.Value {
	once_parseFlags.Do(func() {
		cache_parseFlags = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_parseFlags(s_0_box.StrVal()))
})
	})
	return cache_parseFlags
}

var cache_match gopurs_runtime.Value
var once_match sync.Once
func Get_match() gopurs_runtime.Value {
	once_match.Do(func() {
		cache_match = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(func(inner_arg0 gopurs_runtime.Value, inner_arg1 string) *pkg_Data_Maybe.Constructor_Just[[]*pkg_Data_Maybe.Constructor_Just[string]] {
return (*pkg_Data_Maybe.Constructor_Just[[]*pkg_Data_Maybe.Constructor_Just[string]])(gopurs_runtime.Apply2(gopurs_runtime.Apply2(Get__match(), pkg_Data_Maybe.Get_Just(), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil})), inner_arg0, gopurs_runtime.Str(inner_arg1)).UnsafePtr)
}(arg0, arg1.StrVal()))}
})
	})
	return cache_match
}

var cache_flags gopurs_runtime.Value
var once_flags sync.Once
func Get_flags() gopurs_runtime.Value {
	once_flags.Do(func() {
		cache_flags = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flags(x_0_box))
})
	})
	return cache_flags
}

var cache__match gopurs_runtime.Value
var once__match sync.Once
func Get__match() gopurs_runtime.Value {
	once__match.Do(func() {
		cache__match = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(_Match(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2, arg3.StrVal()))}
})
	})
	return cache__match
}

var cache__replaceBy gopurs_runtime.Value
var once__replaceBy sync.Once
func Get__replaceBy() gopurs_runtime.Value {
	once__replaceBy.Do(func() {
		cache__replaceBy = gopurs_runtime.Func5(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value, arg4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(_ReplaceBy(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2, func(inner_arg0 string, inner_arg1 []*pkg_Data_Maybe.Constructor_Just[string]) string {
return gopurs_runtime.Apply2(arg3, gopurs_runtime.Str(inner_arg0), func() gopurs_runtime.Value {
					arr := inner_arg1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).StrVal()
}, arg4.StrVal()))
})
	})
	return cache__replaceBy
}

var cache__search gopurs_runtime.Value
var once__search sync.Once
func Get__search() gopurs_runtime.Value {
	once__search.Do(func() {
		cache__search = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(_Search(func(inner_arg0 interface{}) *pkg_Data_Maybe.Constructor_Just[interface{}] {
return (*pkg_Data_Maybe.Constructor_Just[interface{}])(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)).UnsafePtr)
}, (*pkg_Data_Maybe.Constructor_Just[interface{}])(arg1.UnsafePtr), arg2, arg3.StrVal()))}
})
	})
	return cache__search
}

var cache_flagsImpl gopurs_runtime.Value
var once_flagsImpl sync.Once
func Get_flagsImpl() gopurs_runtime.Value {
	once_flagsImpl.Do(func() {
		cache_flagsImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(FlagsImpl(arg0))
})
	})
	return cache_flagsImpl
}

var cache_regexImpl gopurs_runtime.Value
var once_regexImpl sync.Once
func Get_regexImpl() gopurs_runtime.Value {
	once_regexImpl.Do(func() {
		cache_regexImpl = gopurs_runtime.Func4(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value, arg3 gopurs_runtime.Value) gopurs_runtime.Value {
return RegexImpl(func(inner_arg0 string) gopurs_runtime.Value {
return gopurs_runtime.Apply(arg0, gopurs_runtime.Str(inner_arg0))
}, func(inner_arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(arg1, inner_arg0)
}, arg2.StrVal(), arg3.StrVal())
})
	})
	return cache_regexImpl
}

var cache_replace gopurs_runtime.Value
var once_replace sync.Once
func Get_replace() gopurs_runtime.Value {
	once_replace.Do(func() {
		cache_replace = gopurs_runtime.Func3(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value, arg2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Replace(arg0, arg1.StrVal(), arg2.StrVal()))
})
	})
	return cache_replace
}

var cache_showRegexImpl gopurs_runtime.Value
var once_showRegexImpl sync.Once
func Get_showRegexImpl() gopurs_runtime.Value {
	once_showRegexImpl.Do(func() {
		cache_showRegexImpl = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(ShowRegexImpl(arg0))
})
	})
	return cache_showRegexImpl
}

var cache_source gopurs_runtime.Value
var once_source sync.Once
func Get_source() gopurs_runtime.Value {
	once_source.Do(func() {
		cache_source = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Source(arg0))
})
	})
	return cache_source
}

var cache_split gopurs_runtime.Value
var once_split sync.Once
func Get_split() gopurs_runtime.Value {
	once_split.Do(func() {
		cache_split = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Split(arg0, arg1.StrVal())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_split
}

var cache_test gopurs_runtime.Value
var once_test sync.Once
func Get_test() gopurs_runtime.Value {
	once_test.Do(func() {
		cache_test = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Test(arg0, arg1.StrVal()))
})
	})
	return cache_test
}

func Call_renderFlags(v_0_loop interface{}) string {
var v_0 interface{} = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(gopurs_runtime.Any(v_0), "global").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(gopurs_runtime.Any(v_0), "ignoreCase").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(gopurs_runtime.Any(v_0), "multiline").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(gopurs_runtime.Any(v_0), "dotAll").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(gopurs_runtime.Any(v_0), "sticky").IntVal) != (0) {
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
if (gopurs_runtime.RecordGet(gopurs_runtime.Any(v_0), "unicode").IntVal) != (0) {
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

func Call_regex(s_0_loop string, f_1_loop interface{}) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
var f_1 interface{} = f_1_loop
_ = f_1
return gopurs_runtime.Apply4(Get_regexImpl(), pkg_Data_Either.Get_Left(), pkg_Data_Either.Get_Right(), gopurs_runtime.Str(s_0), gopurs_runtime.Str(Call_renderFlags(f_1)))
}

func Call_parseFlags(s_0_loop string) interface{} {
var s_0 string = s_0_loop
_ = s_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("s"), gopurs_runtime.Str(s_0)), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("g"), gopurs_runtime.Str(s_0)), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("i"), gopurs_runtime.Str(s_0)), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("m"), gopurs_runtime.Str(s_0)), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("y"), gopurs_runtime.Str(s_0)), gopurs_runtime.Apply2(pkg_Data_String_CodeUnits.Get_contains(), gopurs_runtime.Str("u"), gopurs_runtime.Str(s_0))}))
}

func Call_flags(x_0_loop gopurs_runtime.Value) interface{} {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(Get_flagsImpl(), x_0))
}

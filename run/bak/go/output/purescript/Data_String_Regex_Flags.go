package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Regex_Flags_eqArray gopurs_runtime.Value
var once_Data_String_Regex_Flags_eqArray sync.Once
func Get_Data_String_Regex_Flags_eqArray() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_eqArray.Do(func() {
		cache_Data_String_Regex_Flags_eqArray = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq[[]string]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(Get_Data_Eq_eqString(), "eq"))})}
	})
	return cache_Data_String_Regex_Flags_eqArray
}

var cache_Data_String_Regex_Flags_RegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_RegexFlags sync.Once
func Get_Data_String_Regex_Flags_RegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_RegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_RegexFlags = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Regex_Flags_RegexFlags(x_0_box)
})
	})
	return cache_Data_String_Regex_Flags_RegexFlags
}

var cache_Data_String_Regex_Flags_unicode gopurs_runtime.Value
var once_Data_String_Regex_Flags_unicode sync.Once
func Get_Data_String_Regex_Flags_unicode() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_unicode.Do(func() {
		cache_Data_String_Regex_Flags_unicode = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true)})
	})
	return cache_Data_String_Regex_Flags_unicode
}

var cache_Data_String_Regex_Flags_sticky gopurs_runtime.Value
var once_Data_String_Regex_Flags_sticky sync.Once
func Get_Data_String_Regex_Flags_sticky() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_sticky.Do(func() {
		cache_Data_String_Regex_Flags_sticky = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false)})
	})
	return cache_Data_String_Regex_Flags_sticky
}

var cache_Data_String_Regex_Flags_showRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_showRegexFlags sync.Once
func Get_Data_String_Regex_Flags_showRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_showRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_showRegexFlags = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): usedFlags_1_0 -> gopurs_runtime.Value
usedFlags_1_0 := gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("global")
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Alternative_guard(), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Control_Alternative_alternativeArray()))}, gopurs_runtime.Bool((gopurs_runtime.RecordGet(v_0, "global").IntVal) != (0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("ignoreCase")
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Alternative_guard(), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Control_Alternative_alternativeArray()))}, gopurs_runtime.Bool((gopurs_runtime.RecordGet(v_0, "ignoreCase").IntVal) != (0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("multiline")
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Alternative_guard(), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Control_Alternative_alternativeArray()))}, gopurs_runtime.Bool((gopurs_runtime.RecordGet(v_0, "multiline").IntVal) != (0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("dotAll")
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Alternative_guard(), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Control_Alternative_alternativeArray()))}, gopurs_runtime.Bool((gopurs_runtime.RecordGet(v_0, "dotAll").IntVal) != (0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("sticky")
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Alternative_guard(), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Control_Alternative_alternativeArray()))}, gopurs_runtime.Bool((gopurs_runtime.RecordGet(v_0, "sticky").IntVal) != (0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unicode")
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Alternative_guard(), gopurs_runtime.Value{Type: 9, IntVal: 397869517, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](Get_Control_Alternative_alternativeArray()))}, gopurs_runtime.Bool((gopurs_runtime.RecordGet(v_0, "unicode").IntVal) != (0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
_ = usedFlags_1_0
var __t1 string
{
if (gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[[]string]](Get_Data_String_Regex_Flags_eqArray()).V0), func() gopurs_runtime.Value {
					arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(usedFlags_1_0.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}(), func() gopurs_runtime.Value {
					arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()).IntVal) != (0) {
__t1 = "noFlags"
goto end_branch_1
} else {

}
}
{
__t1 = (("(") + (gopurs_runtime.Apply2(Get_Data_String_Common_joinWith(), gopurs_runtime.Str(" <> "), func() gopurs_runtime.Value {
					arr := func() []string {
					arr := *(*[]gopurs_runtime.Value)(usedFlags_1_0.UnsafePtr)
					unboxed := make([]string, len(arr))
					for i, v := range arr { unboxed[i] = v.StrVal() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Str(v) }
					return gopurs_runtime.Array(boxed)
				}()).StrVal())) + (")")
}
end_branch_1:
return gopurs_runtime.Str(__t1)
}))
	})
	return cache_Data_String_Regex_Flags_showRegexFlags
}

var cache_Data_String_Regex_Flags_semigroupRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_semigroupRegexFlags sync.Once
func Get_Data_String_Regex_Flags_semigroupRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_semigroupRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_semigroupRegexFlags = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "dotAll").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "dotAll").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "global").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "global").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "ignoreCase").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "ignoreCase").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "multiline").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "multiline").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "sticky").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "sticky").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "unicode").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "unicode").IntVal) != (0)))})
})
}))
	})
	return cache_Data_String_Regex_Flags_semigroupRegexFlags
}

var cache_Data_String_Regex_Flags_noFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_noFlags sync.Once
func Get_Data_String_Regex_Flags_noFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_noFlags.Do(func() {
		cache_Data_String_Regex_Flags_noFlags = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_Data_String_Regex_Flags_noFlags
}

var cache_Data_String_Regex_Flags_newtypeRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_newtypeRegexFlags sync.Once
func Get_Data_String_Regex_Flags_newtypeRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_newtypeRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_newtypeRegexFlags = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_String_Regex_Flags_newtypeRegexFlags
}

var cache_Data_String_Regex_Flags_multiline gopurs_runtime.Value
var once_Data_String_Regex_Flags_multiline sync.Once
func Get_Data_String_Regex_Flags_multiline() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_multiline.Do(func() {
		cache_Data_String_Regex_Flags_multiline = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_Data_String_Regex_Flags_multiline
}

var cache_Data_String_Regex_Flags_monoidRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_monoidRegexFlags sync.Once
func Get_Data_String_Regex_Flags_monoidRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_monoidRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_monoidRegexFlags = gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_String_Regex_Flags_semigroupRegexFlags()
}), Get_Data_String_Regex_Flags_noFlags())
	})
	return cache_Data_String_Regex_Flags_monoidRegexFlags
}

var cache_Data_String_Regex_Flags_ignoreCase gopurs_runtime.Value
var once_Data_String_Regex_Flags_ignoreCase sync.Once
func Get_Data_String_Regex_Flags_ignoreCase() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_ignoreCase.Do(func() {
		cache_Data_String_Regex_Flags_ignoreCase = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_Data_String_Regex_Flags_ignoreCase
}

var cache_Data_String_Regex_Flags_global gopurs_runtime.Value
var once_Data_String_Regex_Flags_global sync.Once
func Get_Data_String_Regex_Flags_global() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_global.Do(func() {
		cache_Data_String_Regex_Flags_global = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_Data_String_Regex_Flags_global
}

var cache_Data_String_Regex_Flags_eqRegexFlags gopurs_runtime.Value
var once_Data_String_Regex_Flags_eqRegexFlags sync.Once
func Get_Data_String_Regex_Flags_eqRegexFlags() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_eqRegexFlags.Do(func() {
		cache_Data_String_Regex_Flags_eqRegexFlags = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): get_3_1 -> gopurs_runtime.Value
get_3_1 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str("unicode"))
_ = get_3_1
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Eq_eqBoolean(), "eq"), gopurs_runtime.Apply(get_3_1, ra_1), gopurs_runtime.Apply(get_3_1, rb_2)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Eq_eqRowNil(), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_1, rb_2).IntVal) != (0)))
})
})
}))
_ = __local_var_0_0
// TAST (Let): __local_var_1_2 -> gopurs_runtime.Value
__local_var_1_2 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): get_4_3 -> gopurs_runtime.Value
get_4_3 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str("sticky"))
_ = get_4_3
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Eq_eqBoolean(), "eq"), gopurs_runtime.Apply(get_4_3, ra_2), gopurs_runtime.Apply(get_4_3, rb_3)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_0_0, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_2, rb_3).IntVal) != (0)))
})
})
}))
_ = __local_var_1_2
// TAST (Let): __local_var_2_4 -> gopurs_runtime.Value
__local_var_2_4 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): get_5_5 -> gopurs_runtime.Value
get_5_5 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str("multiline"))
_ = get_5_5
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Eq_eqBoolean(), "eq"), gopurs_runtime.Apply(get_5_5, ra_3), gopurs_runtime.Apply(get_5_5, rb_4)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_2, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_3, rb_4).IntVal) != (0)))
})
})
}))
_ = __local_var_2_4
// TAST (Let): __local_var_3_6 -> gopurs_runtime.Value
__local_var_3_6 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): get_6_7 -> gopurs_runtime.Value
get_6_7 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str("ignoreCase"))
_ = get_6_7
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Eq_eqBoolean(), "eq"), gopurs_runtime.Apply(get_6_7, ra_4), gopurs_runtime.Apply(get_6_7, rb_5)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_4, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_4, rb_5).IntVal) != (0)))
})
})
}))
_ = __local_var_3_6
// TAST (Let): __local_var_4_8 -> gopurs_runtime.Value
__local_var_4_8 := gopurs_runtime.RecordDict1("eqRecord", gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): get_7_9 -> gopurs_runtime.Value
get_7_9 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str("global"))
_ = get_7_9
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Eq_eqBoolean(), "eq"), gopurs_runtime.Apply(get_7_9, ra_5), gopurs_runtime.Apply(get_7_9, rb_6)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_6, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
})
})
}))
_ = __local_var_4_8
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(ra_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(rb_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): get_7_10 -> gopurs_runtime.Value
get_7_10 := gopurs_runtime.Apply(Get_Record_Unsafe_unsafeGet(), gopurs_runtime.Str("dotAll"))
_ = get_7_10
return gopurs_runtime.Bool(((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Eq_eqBoolean(), "eq"), gopurs_runtime.Apply(get_7_10, ra_5), gopurs_runtime.Apply(get_7_10, rb_6)).IntVal) != (0)) && ((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_4_8, "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: int64(513803634), UnsafePtr: nil}, ra_5, rb_6).IntVal) != (0)))
})
}))
}()
	})
	return cache_Data_String_Regex_Flags_eqRegexFlags
}

var cache_Data_String_Regex_Flags_dotAll gopurs_runtime.Value
var once_Data_String_Regex_Flags_dotAll sync.Once
func Get_Data_String_Regex_Flags_dotAll() gopurs_runtime.Value {
	once_Data_String_Regex_Flags_dotAll.Do(func() {
		cache_Data_String_Regex_Flags_dotAll = gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_Data_String_Regex_Flags_dotAll
}

func Call_Data_String_Regex_Flags_RegexFlags(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}



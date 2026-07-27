package Data_String_Regex_Flags

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
)

var cache_guard gopurs_runtime.Value
var once_guard sync.Once
func Get_guard() gopurs_runtime.Value {
	once_guard.Do(func() {
		cache_guard = gopurs_runtime.Apply(pkg_Control_Alternative.Get_guard(), pkg_Control_Alternative.Get_alternativeArray())
	})
	return cache_guard
}

var cache_eq gopurs_runtime.Value
var once_eq sync.Once
func Get_eq() gopurs_runtime.Value {
	once_eq.Do(func() {
		cache_eq = gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq"))
	})
	return cache_eq
}

var cache_RegexFlags gopurs_runtime.Value
var once_RegexFlags sync.Once
func Get_RegexFlags() gopurs_runtime.Value {
	once_RegexFlags.Do(func() {
		cache_RegexFlags = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_RegexFlags(x_0_box)
})
	})
	return cache_RegexFlags
}

var cache_unicode gopurs_runtime.Value
var once_unicode sync.Once
func Get_unicode() gopurs_runtime.Value {
	once_unicode.Do(func() {
		cache_unicode = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true)})))
	})
	return cache_unicode
}

var cache_sticky gopurs_runtime.Value
var once_sticky sync.Once
func Get_sticky() gopurs_runtime.Value {
	once_sticky.Do(func() {
		cache_sticky = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false)})))
	})
	return cache_sticky
}

var cache_showRegexFlags gopurs_runtime.Value
var once_showRegexFlags sync.Once
func Get_showRegexFlags() gopurs_runtime.Value {
	once_showRegexFlags.Do(func() {
		cache_showRegexFlags = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
usedFlags_1_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Array([]gopurs_runtime.Value{}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("global")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "global"))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("ignoreCase")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "ignoreCase"))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("multiline")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "multiline"))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("dotAll")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "dotAll"))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("sticky")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "sticky"))), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unicode")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "unicode")))))))))
_ = usedFlags_1_0
var __t1 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq(), usedFlags_1_0, gopurs_runtime.Array([]gopurs_runtime.Value{})).IntVal) != (0) {
__t1 = gopurs_runtime.Str("noFlags")
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("("), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(pkg_Data_String_Common.Get_joinWith(), gopurs_runtime.Str(" <> "), usedFlags_1_0), gopurs_runtime.Str(")")))
}
end_branch_1:
return __t1
}))))
	})
	return cache_showRegexFlags
}

var cache_semigroupRegexFlags gopurs_runtime.Value
var once_semigroupRegexFlags sync.Once
func Get_semigroupRegexFlags() gopurs_runtime.Value {
	once_semigroupRegexFlags.Do(func() {
		cache_semigroupRegexFlags = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.RecordGet(v_0, "dotAll"), gopurs_runtime.RecordGet(v1_1, "dotAll")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.RecordGet(v_0, "global"), gopurs_runtime.RecordGet(v1_1, "global")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.RecordGet(v_0, "ignoreCase"), gopurs_runtime.RecordGet(v1_1, "ignoreCase")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.RecordGet(v_0, "multiline"), gopurs_runtime.RecordGet(v1_1, "multiline")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.RecordGet(v_0, "sticky"), gopurs_runtime.RecordGet(v1_1, "sticky")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "disj"), gopurs_runtime.RecordGet(v_0, "unicode"), gopurs_runtime.RecordGet(v1_1, "unicode"))})
}))))
	})
	return cache_semigroupRegexFlags
}

var cache_noFlags gopurs_runtime.Value
var once_noFlags sync.Once
func Get_noFlags() gopurs_runtime.Value {
	once_noFlags.Do(func() {
		cache_noFlags = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})))
	})
	return cache_noFlags
}

var cache_newtypeRegexFlags gopurs_runtime.Value
var once_newtypeRegexFlags sync.Once
func Get_newtypeRegexFlags() gopurs_runtime.Value {
	once_newtypeRegexFlags.Do(func() {
		cache_newtypeRegexFlags = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))))
	})
	return cache_newtypeRegexFlags
}

var cache_multiline gopurs_runtime.Value
var once_multiline sync.Once
func Get_multiline() gopurs_runtime.Value {
	once_multiline.Do(func() {
		cache_multiline = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})))
	})
	return cache_multiline
}

var cache_monoidRegexFlags gopurs_runtime.Value
var once_monoidRegexFlags sync.Once
func Get_monoidRegexFlags() gopurs_runtime.Value {
	once_monoidRegexFlags.Do(func() {
		cache_monoidRegexFlags = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupRegexFlags()
}), Get_noFlags())))
	})
	return cache_monoidRegexFlags
}

var cache_ignoreCase gopurs_runtime.Value
var once_ignoreCase sync.Once
func Get_ignoreCase() gopurs_runtime.Value {
	once_ignoreCase.Do(func() {
		cache_ignoreCase = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})))
	})
	return cache_ignoreCase
}

var cache_global gopurs_runtime.Value
var once_global sync.Once
func Get_global() gopurs_runtime.Value {
	once_global.Do(func() {
		cache_global = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})))
	})
	return cache_global
}

var cache_eqRegexFlags gopurs_runtime.Value
var once_eqRegexFlags sync.Once
func Get_eqRegexFlags() gopurs_runtime.Value {
	once_eqRegexFlags.Do(func() {
		cache_eqRegexFlags = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(ra_0 gopurs_runtime.Value, rb_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "dotAll"), gopurs_runtime.RecordGet(rb_1, "dotAll")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "global"), gopurs_runtime.RecordGet(rb_1, "global")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "ignoreCase"), gopurs_runtime.RecordGet(rb_1, "ignoreCase")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "multiline"), gopurs_runtime.RecordGet(rb_1, "multiline")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "sticky"), gopurs_runtime.RecordGet(rb_1, "sticky")), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqBoolean(), "eq"), gopurs_runtime.RecordGet(ra_0, "unicode"), gopurs_runtime.RecordGet(rb_1, "unicode")), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqRowNil(), "eqRecord"), gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}), ra_0, rb_1)))))))
}))))
	})
	return cache_eqRegexFlags
}

var cache_dotAll gopurs_runtime.Value
var once_dotAll sync.Once
func Get_dotAll() gopurs_runtime.Value {
	once_dotAll.Do(func() {
		cache_dotAll = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict([]string{"dotAll", "global", "ignoreCase", "multiline", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})))
	})
	return cache_dotAll
}

func Call_RegexFlags(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

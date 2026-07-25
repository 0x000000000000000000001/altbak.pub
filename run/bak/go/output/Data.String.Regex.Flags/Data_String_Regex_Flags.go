package Data_String_Regex_Flags

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Type_Proxy "gopurs/output/Type.Proxy"
	unsafe "unsafe"
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
		cache_eq = gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), pkg_Data_Eq.Get_eqStringImpl())
	})
	return cache_eq
}

var cache_RegexFlags gopurs_runtime.Value
var once_RegexFlags sync.Once
func Get_RegexFlags() gopurs_runtime.Value {
	once_RegexFlags.Do(func() {
		cache_RegexFlags = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_RegexFlags
}

var cache_unicode gopurs_runtime.Value
var once_unicode sync.Once
func Get_unicode() gopurs_runtime.Value {
	once_unicode.Do(func() {
		cache_unicode = gopurs_runtime.RecordDict([]string{"global", "ignoreCase", "multiline", "dotAll", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true)})
	})
	return cache_unicode
}

var cache_sticky gopurs_runtime.Value
var once_sticky sync.Once
func Get_sticky() gopurs_runtime.Value {
	once_sticky.Do(func() {
		cache_sticky = gopurs_runtime.RecordDict([]string{"global", "ignoreCase", "multiline", "dotAll", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false)})
	})
	return cache_sticky
}

var cache_showRegexFlags gopurs_runtime.Value
var once_showRegexFlags sync.Once
func Get_showRegexFlags() gopurs_runtime.Value {
	once_showRegexFlags.Do(func() {
		cache_showRegexFlags = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
usedFlags_1_0 := gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Apply2(pkg_Data_Semigroup.Get_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{}), gopurs_runtime.Apply2(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("global")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "global")))), gopurs_runtime.Apply2(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("ignoreCase")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "ignoreCase")))), gopurs_runtime.Apply2(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("multiline")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "multiline")))), gopurs_runtime.Apply2(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("dotAll")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "dotAll")))), gopurs_runtime.Apply2(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("sticky")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "sticky")))), gopurs_runtime.Apply2(pkg_Data_Functor.Get_arrayMap(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unicode")
}), gopurs_runtime.Apply(Get_guard(), gopurs_runtime.RecordGet(v_0, "unicode"))))
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
__t1 = gopurs_runtime.Str((("(") + (gopurs_runtime.Apply2(pkg_Data_String_Common.Get_joinWith(), gopurs_runtime.Str(" <> "), usedFlags_1_0).StrVal())) + (")"))
}
end_branch_1:
return __t1
}))
	})
	return cache_showRegexFlags
}

var cache_semigroupRegexFlags gopurs_runtime.Value
var once_semigroupRegexFlags sync.Once
func Get_semigroupRegexFlags() gopurs_runtime.Value {
	once_semigroupRegexFlags.Do(func() {
		cache_semigroupRegexFlags = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict([]string{"global", "ignoreCase", "multiline", "dotAll", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "global").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "global").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "ignoreCase").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "ignoreCase").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "multiline").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "multiline").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "dotAll").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "dotAll").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "sticky").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "sticky").IntVal) != (0))), gopurs_runtime.Bool(((gopurs_runtime.RecordGet(v_0, "unicode").IntVal) != (0)) || ((gopurs_runtime.RecordGet(v1_1, "unicode").IntVal) != (0)))})
}))
	})
	return cache_semigroupRegexFlags
}

var cache_noFlags gopurs_runtime.Value
var once_noFlags sync.Once
func Get_noFlags() gopurs_runtime.Value {
	once_noFlags.Do(func() {
		cache_noFlags = gopurs_runtime.RecordDict([]string{"global", "ignoreCase", "multiline", "dotAll", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_noFlags
}

var cache_newtypeRegexFlags gopurs_runtime.Value
var once_newtypeRegexFlags sync.Once
func Get_newtypeRegexFlags() gopurs_runtime.Value {
	once_newtypeRegexFlags.Do(func() {
		cache_newtypeRegexFlags = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeRegexFlags
}

var cache_multiline gopurs_runtime.Value
var once_multiline sync.Once
func Get_multiline() gopurs_runtime.Value {
	once_multiline.Do(func() {
		cache_multiline = gopurs_runtime.RecordDict([]string{"global", "ignoreCase", "multiline", "dotAll", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_multiline
}

var cache_monoidRegexFlags gopurs_runtime.Value
var once_monoidRegexFlags sync.Once
func Get_monoidRegexFlags() gopurs_runtime.Value {
	once_monoidRegexFlags.Do(func() {
		cache_monoidRegexFlags = gopurs_runtime.RecordDict2("mempty", "Semigroup0", Get_noFlags(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_semigroupRegexFlags()
}))
	})
	return cache_monoidRegexFlags
}

var cache_ignoreCase gopurs_runtime.Value
var once_ignoreCase sync.Once
func Get_ignoreCase() gopurs_runtime.Value {
	once_ignoreCase.Do(func() {
		cache_ignoreCase = gopurs_runtime.RecordDict([]string{"global", "ignoreCase", "multiline", "dotAll", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_ignoreCase
}

var cache_global gopurs_runtime.Value
var once_global sync.Once
func Get_global() gopurs_runtime.Value {
	once_global.Do(func() {
		cache_global = gopurs_runtime.RecordDict([]string{"global", "ignoreCase", "multiline", "dotAll", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_global
}

var cache_eqRegexFlags gopurs_runtime.Value
var once_eqRegexFlags sync.Once
func Get_eqRegexFlags() gopurs_runtime.Value {
	once_eqRegexFlags.Do(func() {
		cache_eqRegexFlags = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply4(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply4(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply4(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply4(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply4(pkg_Data_Eq.Get_eqRowCons(), gopurs_runtime.Apply4(pkg_Data_Eq.Get_eqRowCons(), pkg_Data_Eq.Get_eqRowNil(), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("unicode")
})), pkg_Data_Eq.Get_eqBoolean()), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("sticky")
})), pkg_Data_Eq.Get_eqBoolean()), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("multiline")
})), pkg_Data_Eq.Get_eqBoolean()), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("ignoreCase")
})), pkg_Data_Eq.Get_eqBoolean()), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("global")
})), pkg_Data_Eq.Get_eqBoolean()), gopurs_runtime.Value{}, gopurs_runtime.RecordDict1("reflectSymbol", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("dotAll")
})), pkg_Data_Eq.Get_eqBoolean()), "eqRecord"), gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: unsafe.Pointer(&pkg_Type_Proxy.Data_Type_Proxy_Proxy{})}))
	})
	return cache_eqRegexFlags
}

var cache_dotAll gopurs_runtime.Value
var once_dotAll sync.Once
func Get_dotAll() gopurs_runtime.Value {
	once_dotAll.Do(func() {
		cache_dotAll = gopurs_runtime.RecordDict([]string{"global", "ignoreCase", "multiline", "dotAll", "sticky", "unicode"}, []gopurs_runtime.Value{gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false), gopurs_runtime.Bool(true), gopurs_runtime.Bool(false), gopurs_runtime.Bool(false)})
	})
	return cache_dotAll
}





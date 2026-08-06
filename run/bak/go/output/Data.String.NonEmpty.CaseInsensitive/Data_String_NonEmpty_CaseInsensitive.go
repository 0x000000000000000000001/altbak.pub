package Data_String_NonEmpty_CaseInsensitive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_String_NonEmpty_Internal "gopurs/output/Data.String.NonEmpty.Internal"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var cache_CaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_CaseInsensitiveNonEmptyString sync.Once
func Get_CaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_CaseInsensitiveNonEmptyString.Do(func() {
		cache_CaseInsensitiveNonEmptyString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_CaseInsensitiveNonEmptyString(x_0_box)
})
	})
	return cache_CaseInsensitiveNonEmptyString
}

var cache_showCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_showCaseInsensitiveNonEmptyString sync.Once
func Get_showCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_showCaseInsensitiveNonEmptyString.Do(func() {
		cache_showCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(CaseInsensitiveNonEmptyString "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_String_NonEmpty_Internal.Get_showNonEmptyString(), "show"), v_0), gopurs_runtime.Str(")"))).StrVal())
}))
	})
	return cache_showCaseInsensitiveNonEmptyString
}

var cache_newtypeCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_newtypeCaseInsensitiveNonEmptyString sync.Once
func Get_newtypeCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_newtypeCaseInsensitiveNonEmptyString.Do(func() {
		cache_newtypeCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCaseInsensitiveNonEmptyString
}

var cache_eqCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_eqCaseInsensitiveNonEmptyString sync.Once
func Get_eqCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_eqCaseInsensitiveNonEmptyString.Do(func() {
		cache_eqCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq"), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1)).IntVal) != (0))
})
}))
	})
	return cache_eqCaseInsensitiveNonEmptyString
}

var cache_ordCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_ordCaseInsensitiveNonEmptyString sync.Once
func Get_ordCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_ordCaseInsensitiveNonEmptyString.Do(func() {
		cache_ordCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqCaseInsensitiveNonEmptyString()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordString(), "compare"), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1))
})
}))
	})
	return cache_ordCaseInsensitiveNonEmptyString
}

func Call_CaseInsensitiveNonEmptyString(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}



package Data_String_NonEmpty_CaseInsensitive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var CaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_CaseInsensitiveNonEmptyString sync.Once
func Get_CaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_CaseInsensitiveNonEmptyString.Do(func() {
		CaseInsensitiveNonEmptyString = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return CaseInsensitiveNonEmptyString
}

var showCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_showCaseInsensitiveNonEmptyString sync.Once
func Get_showCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_showCaseInsensitiveNonEmptyString.Do(func() {
		showCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(CaseInsensitiveNonEmptyString (NonEmptyString.unsafeFromString " + gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0).StrVal + "))")
}))
	})
	return showCaseInsensitiveNonEmptyString
}

var newtypeCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_newtypeCaseInsensitiveNonEmptyString sync.Once
func Get_newtypeCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_newtypeCaseInsensitiveNonEmptyString.Do(func() {
		newtypeCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeCaseInsensitiveNonEmptyString
}

var eqCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_eqCaseInsensitiveNonEmptyString sync.Once
func Get_eqCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_eqCaseInsensitiveNonEmptyString.Do(func() {
		eqCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Boolean(gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0).StrVal == gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1).StrVal)
}))
	})
	return eqCaseInsensitiveNonEmptyString
}

var ordCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_ordCaseInsensitiveNonEmptyString sync.Once
func Get_ordCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_ordCaseInsensitiveNonEmptyString.Do(func() {
		ordCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordString(), "compare"), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqCaseInsensitiveNonEmptyString()
}))
	})
	return ordCaseInsensitiveNonEmptyString
}





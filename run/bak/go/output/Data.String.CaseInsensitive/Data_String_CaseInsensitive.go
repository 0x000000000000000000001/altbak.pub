package Data_String_CaseInsensitive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var CaseInsensitiveString gopurs_runtime.Value
var once_CaseInsensitiveString sync.Once
func Get_CaseInsensitiveString() gopurs_runtime.Value {
	once_CaseInsensitiveString.Do(func() {
		CaseInsensitiveString = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return CaseInsensitiveString
}

var showCaseInsensitiveString gopurs_runtime.Value
var once_showCaseInsensitiveString sync.Once
func Get_showCaseInsensitiveString() gopurs_runtime.Value {
	once_showCaseInsensitiveString.Do(func() {
		showCaseInsensitiveString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str("(CaseInsensitiveString " + gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0).StrVal + ")")
}))
	})
	return showCaseInsensitiveString
}

var newtypeCaseInsensitiveString gopurs_runtime.Value
var once_newtypeCaseInsensitiveString sync.Once
func Get_newtypeCaseInsensitiveString() gopurs_runtime.Value {
	once_newtypeCaseInsensitiveString.Do(func() {
		newtypeCaseInsensitiveString = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return newtypeCaseInsensitiveString
}

var eqCaseInsensitiveString gopurs_runtime.Value
var once_eqCaseInsensitiveString sync.Once
func Get_eqCaseInsensitiveString() gopurs_runtime.Value {
	once_eqCaseInsensitiveString.Do(func() {
		eqCaseInsensitiveString = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0).StrVal == gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1).StrVal)
}))
	})
	return eqCaseInsensitiveString
}

var ordCaseInsensitiveString gopurs_runtime.Value
var once_ordCaseInsensitiveString sync.Once
func Get_ordCaseInsensitiveString() gopurs_runtime.Value {
	once_ordCaseInsensitiveString.Do(func() {
		ordCaseInsensitiveString = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordString(), "compare"), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqCaseInsensitiveString()
}))
	})
	return ordCaseInsensitiveString
}





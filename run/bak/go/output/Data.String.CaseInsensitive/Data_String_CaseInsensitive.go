package Data_String_CaseInsensitive

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Data_Ord "gopurs/output/Data.Ord"
)

var cache_CaseInsensitiveString gopurs_runtime.Value
var once_CaseInsensitiveString sync.Once
func Get_CaseInsensitiveString() gopurs_runtime.Value {
	once_CaseInsensitiveString.Do(func() {
		cache_CaseInsensitiveString = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}()
})
	})
	return cache_CaseInsensitiveString
}

var cache_showCaseInsensitiveString gopurs_runtime.Value
var once_showCaseInsensitiveString sync.Once
func Get_showCaseInsensitiveString() gopurs_runtime.Value {
	once_showCaseInsensitiveString.Do(func() {
		cache_showCaseInsensitiveString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(CaseInsensitiveString ") + (gopurs_runtime.Apply(pkg_Data_Show.Get_showStringImpl(), v_0).StrVal())) + (")"))
}))
	})
	return cache_showCaseInsensitiveString
}

var cache_newtypeCaseInsensitiveString gopurs_runtime.Value
var once_newtypeCaseInsensitiveString sync.Once
func Get_newtypeCaseInsensitiveString() gopurs_runtime.Value {
	once_newtypeCaseInsensitiveString.Do(func() {
		cache_newtypeCaseInsensitiveString = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeCaseInsensitiveString
}

var cache_eqCaseInsensitiveString gopurs_runtime.Value
var once_eqCaseInsensitiveString sync.Once
func Get_eqCaseInsensitiveString() gopurs_runtime.Value {
	once_eqCaseInsensitiveString.Do(func() {
		cache_eqCaseInsensitiveString = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0).StrVal()) == (gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1).StrVal()))
}))
	})
	return cache_eqCaseInsensitiveString
}

var cache_ordCaseInsensitiveString gopurs_runtime.Value
var once_ordCaseInsensitiveString sync.Once
func Get_ordCaseInsensitiveString() gopurs_runtime.Value {
	once_ordCaseInsensitiveString.Do(func() {
		cache_ordCaseInsensitiveString = gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ord.Get_ordString(), "compare"), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v_0), gopurs_runtime.Apply(pkg_Data_String_Common.Get_toLower(), v1_1))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eqCaseInsensitiveString()
}))
	})
	return cache_ordCaseInsensitiveString
}





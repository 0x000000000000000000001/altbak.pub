package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_String_CaseInsensitive_CaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_CaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_CaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_CaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_CaseInsensitiveString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_CaseInsensitive_CaseInsensitiveString(x_0_box)
})
	})
	return cache_Data_String_CaseInsensitive_CaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_showCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_showCaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_showCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_showCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_showCaseInsensitiveString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(CaseInsensitiveString ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_String_CaseInsensitive_showCaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_newtypeCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_newtypeCaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_newtypeCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_newtypeCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_newtypeCaseInsensitiveString = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_String_CaseInsensitive_newtypeCaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_eqCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_eqCaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_eqCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_eqCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_eqCaseInsensitiveString = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v_0.StrVal())).StrVal()) == (gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v1_1.StrVal())).StrVal()))
})
}))
	})
	return cache_Data_String_CaseInsensitive_eqCaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_ordCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_ordCaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_ordCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_ordCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_ordCaseInsensitiveString = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_String_CaseInsensitive_eqCaseInsensitiveString()
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v_0.StrVal())).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v1_1.StrVal())).StrVal())).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_Data_String_CaseInsensitive_ordCaseInsensitiveString
}

func Call_Data_String_CaseInsensitive_CaseInsensitiveString(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}



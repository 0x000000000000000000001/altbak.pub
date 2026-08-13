package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString(x_0_box)
})
	})
	return cache_Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString
}

var cache_Data_String_NonEmpty_CaseInsensitive_showCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CaseInsensitive_showCaseInsensitiveNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CaseInsensitive_showCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CaseInsensitive_showCaseInsensitiveNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CaseInsensitive_showCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(CaseInsensitiveNonEmptyString ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
}))
	})
	return cache_Data_String_NonEmpty_CaseInsensitive_showCaseInsensitiveNonEmptyString
}

var cache_Data_String_NonEmpty_CaseInsensitive_newtypeCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CaseInsensitive_newtypeCaseInsensitiveNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CaseInsensitive_newtypeCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CaseInsensitive_newtypeCaseInsensitiveNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CaseInsensitive_newtypeCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Data_String_NonEmpty_CaseInsensitive_newtypeCaseInsensitiveNonEmptyString
}

var cache_Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v_0.StrVal())).StrVal()) == (gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v1_1.StrVal())).StrVal()))
})
}))
	})
	return cache_Data_String_NonEmpty_CaseInsensitive_eqCaseInsensitiveNonEmptyString
}

var cache_Data_String_NonEmpty_CaseInsensitive_ordCaseInsensitiveNonEmptyString gopurs_runtime.Value
var once_Data_String_NonEmpty_CaseInsensitive_ordCaseInsensitiveNonEmptyString sync.Once
func Get_Data_String_NonEmpty_CaseInsensitive_ordCaseInsensitiveNonEmptyString() gopurs_runtime.Value {
	once_Data_String_NonEmpty_CaseInsensitive_ordCaseInsensitiveNonEmptyString.Do(func() {
		cache_Data_String_NonEmpty_CaseInsensitive_ordCaseInsensitiveNonEmptyString = gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v_1.StrVal())).StrVal()) == (gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v1_2.StrVal())).StrVal()))
})
}))
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v_0.StrVal())).StrVal()), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v1_1.StrVal())).StrVal())).IntVal)), UnsafePtr: nil}
})
}))
	})
	return cache_Data_String_NonEmpty_CaseInsensitive_ordCaseInsensitiveNonEmptyString
}

func Call_Data_String_NonEmpty_CaseInsensitive_CaseInsensitiveNonEmptyString(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}



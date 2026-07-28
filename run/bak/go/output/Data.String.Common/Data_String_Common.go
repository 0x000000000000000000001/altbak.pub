package Data_String_Common

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Eq "gopurs/output/Data.Eq"
)

var cache_null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		cache_null = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_null(s_0_box.StrVal()))
})
	})
	return cache_null
}

var cache_localeCompare gopurs_runtime.Value
var once_localeCompare sync.Once
func Get_localeCompare() gopurs_runtime.Value {
	once_localeCompare.Do(func() {
		cache_localeCompare = gopurs_runtime.Apply3(Get__localeCompare(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil})
	})
	return cache_localeCompare
}

func Call_null(s_0_loop string) bool {
var s_0 string = s_0_loop
_ = s_0
return (gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Eq.Get_eqString(), "eq"), gopurs_runtime.Str(s_0), gopurs_runtime.Str("")).IntVal) != (0)
}

func Get__localeCompare() gopurs_runtime.Value {
	return _Gopurs__LocaleCompare
}

func Get_joinWith() gopurs_runtime.Value {
	return _Gopurs_JoinWith
}

func Get_replace() gopurs_runtime.Value {
	return _Gopurs_Replace
}

func Get_replaceAll() gopurs_runtime.Value {
	return _Gopurs_ReplaceAll
}

func Get_split() gopurs_runtime.Value {
	return _Gopurs_Split
}

func Get_toLower() gopurs_runtime.Value {
	return _Gopurs_ToLower
}

func Get_toUpper() gopurs_runtime.Value {
	return _Gopurs_ToUpper
}

func Get_trim() gopurs_runtime.Value {
	return _Gopurs_Trim
}

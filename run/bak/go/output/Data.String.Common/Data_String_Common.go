package Data_String_Common

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	unsafe "unsafe"
)

var cache_null gopurs_runtime.Value
var once_null sync.Once
func Get_null() gopurs_runtime.Value {
	once_null.Do(func() {
		cache_null = gopurs_runtime.Func(func(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
return gopurs_runtime.Bool((s_0.StrVal()) == (""))
}()
})
	})
	return cache_null
}

var cache_localeCompare gopurs_runtime.Value
var once_localeCompare sync.Once
func Get_localeCompare() gopurs_runtime.Value {
	once_localeCompare.Do(func() {
		cache_localeCompare = gopurs_runtime.Apply3(Get__localeCompare(), gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_LT{})}, gopurs_runtime.Value{Type: 9, IntVal: 902936544, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_EQ{})}, gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: unsafe.Pointer(&pkg_Data_Ordering.Data_Data_Ordering_GT{})})
	})
	return cache_localeCompare
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

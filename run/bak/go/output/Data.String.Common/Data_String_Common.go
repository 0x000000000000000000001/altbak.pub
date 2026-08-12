package Data_String_Common

import (
	pkg_Data_Eq "gopurs/output/Data.Eq"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
		cache_localeCompare = gopurs_runtime.Apply3(Get__localeCompare(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})
	})
	return cache_localeCompare
}

var cache_eq__472317769 gopurs_runtime.Value
var once_eq__472317769 sync.Once
func Get_eq__472317769() gopurs_runtime.Value {
	once_eq__472317769.Do(func() {
		cache_eq__472317769 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__472317769(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_eq__472317769
}

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

func Call_null(s_0_loop string) bool {
var s_0 string = s_0_loop
_ = s_0
return (Call_eq__472317769(gopurs_runtime.Str(s_0), gopurs_runtime.Str("")).IntVal) != (0)
}

func Call_eq__472317769(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool((__eta0_0.StrVal()) == (__eta1_1.StrVal()))
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

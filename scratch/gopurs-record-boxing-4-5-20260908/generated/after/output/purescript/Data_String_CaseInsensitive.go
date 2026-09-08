package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_CaseInsensitive_CaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_CaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_CaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_CaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_CaseInsensitiveString = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_CaseInsensitive_CaseInsensitiveString(x_0_box.StrVal()))
})
	})
	return cache_Data_String_CaseInsensitive_CaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_showCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_showCaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_showCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_showCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_showCaseInsensitiveString = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(CaseInsensitiveString ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
})})))}
	})
	return cache_Data_String_CaseInsensitive_showCaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_newtypeCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_newtypeCaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_newtypeCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_newtypeCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_newtypeCaseInsensitiveString = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_2199435624_385277032((&Constructor_Data_Newtype_Newtype[string, string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
	})
	return cache_Data_String_CaseInsensitive_newtypeCaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_eqCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_eqCaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_eqCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_eqCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_eqCaseInsensitiveString = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_1140313009_3790796878((&Constructor_Data_Eq_Eq[string]{1, gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v_0.StrVal())).StrVal()) == (gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v1_1.StrVal())).StrVal()))
})})))}
	})
	return cache_Data_String_CaseInsensitive_eqCaseInsensitiveString
}

var cache_Data_String_CaseInsensitive_ordCaseInsensitiveString gopurs_runtime.Value
var once_Data_String_CaseInsensitive_ordCaseInsensitiveString sync.Once
func Get_Data_String_CaseInsensitive_ordCaseInsensitiveString() gopurs_runtime.Value {
	once_Data_String_CaseInsensitive_ordCaseInsensitiveString.Do(func() {
		cache_Data_String_CaseInsensitive_ordCaseInsensitiveString = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_CaseInsensitive_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_String_CaseInsensitive_eqCaseInsensitiveString())))}
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v_0.StrVal())), gopurs_runtime.Apply(Get_Data_String_Common_toLower(), gopurs_runtime.Str(v1_1.StrVal()))).IntVal)), UnsafePtr: nil}
})})))}
	})
	return cache_Data_String_CaseInsensitive_ordCaseInsensitiveString
}

func Call_Data_String_CaseInsensitive_CaseInsensitiveString(x_0_loop string) string {
var x_0 string = x_0_loop
_ = x_0
return x_0
}

func Rebox_Data_String_CaseInsensitive_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_CaseInsensitive_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_CaseInsensitive_2199435624_385277032(in *Constructor_Data_Newtype_Newtype[string, string]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_CaseInsensitive_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}



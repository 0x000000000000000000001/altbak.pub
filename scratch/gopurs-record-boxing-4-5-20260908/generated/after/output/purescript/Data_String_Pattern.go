package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Pattern_Replacement gopurs_runtime.Value
var once_Data_String_Pattern_Replacement sync.Once
func Get_Data_String_Pattern_Replacement() gopurs_runtime.Value {
	once_Data_String_Pattern_Replacement.Do(func() {
		cache_Data_String_Pattern_Replacement = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_Pattern_Replacement(x_0_box.StrVal()))
})
	})
	return cache_Data_String_Pattern_Replacement
}

var cache_Data_String_Pattern_Pattern gopurs_runtime.Value
var once_Data_String_Pattern_Pattern sync.Once
func Get_Data_String_Pattern_Pattern() gopurs_runtime.Value {
	once_Data_String_Pattern_Pattern.Do(func() {
		cache_Data_String_Pattern_Pattern = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_Data_String_Pattern_Pattern(x_0_box.StrVal()))
})
	})
	return cache_Data_String_Pattern_Pattern
}

var cache_Data_String_Pattern_showReplacement gopurs_runtime.Value
var once_Data_String_Pattern_showReplacement sync.Once
func Get_Data_String_Pattern_showReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_showReplacement.Do(func() {
		cache_Data_String_Pattern_showReplacement = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Replacement ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
})})))}
	})
	return cache_Data_String_Pattern_showReplacement
}

var cache_Data_String_Pattern_showPattern gopurs_runtime.Value
var once_Data_String_Pattern_showPattern sync.Once
func Get_Data_String_Pattern_showPattern() gopurs_runtime.Value {
	once_Data_String_Pattern_showPattern.Do(func() {
		cache_Data_String_Pattern_showPattern = gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1514099793_1386611502((&Constructor_Data_Show_Show[string]{1, gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(Pattern ") + (gopurs_runtime.Apply(Get_Data_Show_showStringImpl(), gopurs_runtime.Str(v_0.StrVal())).StrVal())) + (")"))
})})))}
	})
	return cache_Data_String_Pattern_showPattern
}

var cache_Data_String_Pattern_newtypeReplacement gopurs_runtime.Value
var once_Data_String_Pattern_newtypeReplacement sync.Once
func Get_Data_String_Pattern_newtypeReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_newtypeReplacement.Do(func() {
		cache_Data_String_Pattern_newtypeReplacement = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_2199435624_385277032((&Constructor_Data_Newtype_Newtype[string, string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
	})
	return cache_Data_String_Pattern_newtypeReplacement
}

var cache_Data_String_Pattern_newtypePattern gopurs_runtime.Value
var once_Data_String_Pattern_newtypePattern sync.Once
func Get_Data_String_Pattern_newtypePattern() gopurs_runtime.Value {
	once_Data_String_Pattern_newtypePattern.Do(func() {
		cache_Data_String_Pattern_newtypePattern = gopurs_runtime.Value{Type: 9, IntVal: 3322196858, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_2199435624_385277032((&Constructor_Data_Newtype_Newtype[string, string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
})})))}
	})
	return cache_Data_String_Pattern_newtypePattern
}

var cache_Data_String_Pattern_eqReplacement gopurs_runtime.Value
var once_Data_String_Pattern_eqReplacement sync.Once
func Get_Data_String_Pattern_eqReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_eqReplacement.Do(func() {
		cache_Data_String_Pattern_eqReplacement = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1140313009_3790796878((&Constructor_Data_Eq_Eq[string]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_0.StrVal()) == (y_1.StrVal()))
})})))}
	})
	return cache_Data_String_Pattern_eqReplacement
}

var cache_Data_String_Pattern_ordReplacement gopurs_runtime.Value
var once_Data_String_Pattern_ordReplacement sync.Once
func Get_Data_String_Pattern_ordReplacement() gopurs_runtime.Value {
	once_Data_String_Pattern_ordReplacement.Do(func() {
		cache_Data_String_Pattern_ordReplacement = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_String_Pattern_eqReplacement())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str(x_0.StrVal()), gopurs_runtime.Str(y_1.StrVal())).IntVal)), UnsafePtr: nil}
})})))}
	})
	return cache_Data_String_Pattern_ordReplacement
}

var cache_Data_String_Pattern_eqPattern gopurs_runtime.Value
var once_Data_String_Pattern_eqPattern sync.Once
func Get_Data_String_Pattern_eqPattern() gopurs_runtime.Value {
	once_Data_String_Pattern_eqPattern.Do(func() {
		cache_Data_String_Pattern_eqPattern = gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1140313009_3790796878((&Constructor_Data_Eq_Eq[string]{1, gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((x_0.StrVal()) == (y_1.StrVal()))
})})))}
	})
	return cache_Data_String_Pattern_eqPattern
}

var cache_Data_String_Pattern_ordPattern gopurs_runtime.Value
var once_Data_String_Pattern_ordPattern sync.Once
func Get_Data_String_Pattern_ordPattern() gopurs_runtime.Value {
	once_Data_String_Pattern_ordPattern.Do(func() {
		cache_Data_String_Pattern_ordPattern = gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_2406510097_4177771502((&Constructor_Data_Ord_Ord[string]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_String_Pattern_1140313009_3790796878(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[string]](Get_Data_String_Pattern_eqPattern())))}
}), gopurs_runtime.Func2(func(x_0 gopurs_runtime.Value, y_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordStringImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Str(x_0.StrVal()), gopurs_runtime.Str(y_1.StrVal())).IntVal)), UnsafePtr: nil}
})})))}
	})
	return cache_Data_String_Pattern_ordPattern
}

func Call_Data_String_Pattern_Replacement(x_0_loop string) string {
var x_0 string = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_Pattern_Pattern(x_0_loop string) string {
var x_0 string = x_0_loop
_ = x_0
return x_0
}

func Rebox_Data_String_Pattern_1140313009_3790796878(in *Constructor_Data_Eq_Eq[string]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_Pattern_1514099793_1386611502(in *Constructor_Data_Show_Show[string]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_Pattern_2199435624_385277032(in *Constructor_Data_Newtype_Newtype[string, string]) *Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Newtype_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_String_Pattern_2406510097_4177771502(in *Constructor_Data_Ord_Ord[string]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}



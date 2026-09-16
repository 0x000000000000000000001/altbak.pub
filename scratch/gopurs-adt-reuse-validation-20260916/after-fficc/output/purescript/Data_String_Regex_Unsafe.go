package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_String_Regex_Unsafe_identity gopurs_runtime.Value
var once_Data_String_Regex_Unsafe_identity sync.Once
func Get_Data_String_Regex_Unsafe_identity() gopurs_runtime.Value {
	once_Data_String_Regex_Unsafe_identity.Do(func() {
		cache_Data_String_Regex_Unsafe_identity = Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))})
	})
	return cache_Data_String_Regex_Unsafe_identity
}

var cache_Data_String_Regex_Unsafe_unsafeRegex gopurs_runtime.Value
var once_Data_String_Regex_Unsafe_unsafeRegex sync.Once
func Get_Data_String_Regex_Unsafe_unsafeRegex() gopurs_runtime.Value {
	once_Data_String_Regex_Unsafe_unsafeRegex.Do(func() {
		cache_Data_String_Regex_Unsafe_unsafeRegex = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Regex_Unsafe_unsafeRegex(s_0_box.StrVal(), f_1_box)
})
	})
	return cache_Data_String_Regex_Unsafe_unsafeRegex
}

func Call_Data_String_Regex_Unsafe_unsafeRegex(s_0_loop string, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return Call_Data_Either_either__3820936092(Get_Partial_Unsafe_unsafeCrashWith(), Call_Control_Category_identity(gopurs_runtime.Value{Type: 9, IntVal: 784524589, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()))}), func() gopurs_runtime.Value {
				_v := Call_Data_String_Regex_regex(s_0, f_1)
				if _v.V2 {
					return gopurs_runtime.Value{Type: 9, IntVal: 2465973597, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V1})}
				}
				return gopurs_runtime.Value{Type: 9, IntVal: 3711209382, UnsafePtr: unsafe.Pointer(&Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0})}
			}())
}



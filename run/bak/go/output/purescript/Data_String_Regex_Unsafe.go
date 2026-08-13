package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_String_Regex_Unsafe_identity gopurs_runtime.Value
var once_Data_String_Regex_Unsafe_identity sync.Once
func Get_Data_String_Regex_Unsafe_identity() gopurs_runtime.Value {
	once_Data_String_Regex_Unsafe_identity.Do(func() {
		cache_Data_String_Regex_Unsafe_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_String_Regex_Unsafe_identity(x_0_box)
})
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

func Call_Data_String_Regex_Unsafe_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_String_Regex_Unsafe_unsafeRegex(s_0_loop string, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply2(Get_Data_String_Regex_regex(), gopurs_runtime.Str(s_0), f_1)
_ = __local_var_2_0
var __t2 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3711209382) {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := (*Constructor_Data_Either_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0
_ = __local_var_3_1
__t2 = gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Partial__crashWith(), gopurs_runtime.Str(__local_var_3_1.StrVal()))
}))
goto end_branch_2
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 2465973597) {
__t2 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}



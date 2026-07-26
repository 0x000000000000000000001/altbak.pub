package Data_String_Regex_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_String_Regex "gopurs/output/Data.String.Regex"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Partial "gopurs/output/Partial"
)

var cache_unsafeRegex gopurs_runtime.Value
var once_unsafeRegex sync.Once
func Get_unsafeRegex() gopurs_runtime.Value {
	once_unsafeRegex.Do(func() {
		cache_unsafeRegex = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeRegex(s_0_box.StrVal(), f_1_box)
})
	})
	return cache_unsafeRegex
}

func Call_unsafeRegex(s_0_loop string, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 string = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
__local_var_2_0 := gopurs_runtime.Apply4(pkg_Data_String_Regex.Get_regexImpl(), pkg_Data_Either.Get_Left(), pkg_Data_Either.Get_Right(), gopurs_runtime.Str(s_0), gopurs_runtime.Apply(pkg_Data_String_Regex.Get_renderFlags(), f_1))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 3711209382) {
__t1 = gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), (*pkg_Data_Either.Constructor_Left[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 2465973597) {
__t1 = (*pkg_Data_Either.Constructor_Right[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_2_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}



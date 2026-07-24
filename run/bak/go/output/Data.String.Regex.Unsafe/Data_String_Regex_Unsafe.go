package Data_String_Regex_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_String_Regex "gopurs/output/Data.String.Regex"
	pkg_Partial "gopurs/output/Partial"
	pkg_Data_Either "gopurs/output/Data.Either"
)

var unsafeRegex gopurs_runtime.Value
var once_unsafeRegex sync.Once
func Get_unsafeRegex() gopurs_runtime.Value {
	once_unsafeRegex.Do(func() {
		unsafeRegex = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unsafeRegex(s_0_box, f_1_box)
})
	})
	return unsafeRegex
}

func Call_unsafeRegex(s_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_Regex.Get_regex(), s_0, f_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 590902115) {
__t1 = gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), (*pkg_Data_Either.Data_Data_Either_Left)(__local_var_2_0.UnsafePtr).V0)
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 4096564120) {
__t1 = (*pkg_Data_Either.Data_Data_Either_Right)(__local_var_2_0.UnsafePtr).V0
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



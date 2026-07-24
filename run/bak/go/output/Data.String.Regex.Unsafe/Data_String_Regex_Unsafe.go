package Data_String_Regex_Unsafe

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_String_Regex "gopurs/output/Data.String.Regex"
	pkg_Partial "gopurs/output/Partial"
)

var unsafeRegex gopurs_runtime.Value
var once_unsafeRegex sync.Once
func Get_unsafeRegex() gopurs_runtime.Value {
	once_unsafeRegex.Do(func() {
		unsafeRegex = gopurs_runtime.Func2(Call_unsafeRegex)
	})
	return unsafeRegex
}

func Call_unsafeRegex(s_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
__local_var_2_0 := gopurs_runtime.Apply2(pkg_Data_String_Regex.Get_regex(), s_0_loop, f_1_loop)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if gopurs_runtime.Bool(__local_var_2_0.StrVal == "Left").IntVal != 0 {
__t1 = gopurs_runtime.Apply(pkg_Partial.Get__crashWith(), (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0])
goto end_branch_1
} else {

}
}
{
if gopurs_runtime.Bool(__local_var_2_0.StrVal == "Right").IntVal != 0 {
__t1 = (*[1024]gopurs_runtime.Value)(__local_var_2_0.UnsafePtr)[0]
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



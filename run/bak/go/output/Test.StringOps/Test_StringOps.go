package Test_StringOps

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_String_Regex "gopurs/output/Data.String.Regex"
	pkg_Data_String_Regex_Flags "gopurs/output/Data.String.Regex.Flags"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_String_CodePoints "gopurs/output/Data.String.CodePoints"
	pkg_Data_String_Common "gopurs/output/Data.String.Common"
	pkg_Effect_Console "gopurs/output/Effect.Console"
)

var regexPattern gopurs_runtime.Value
var once_regexPattern sync.Once
func Get_regexPattern() gopurs_runtime.Value {
	once_regexPattern.Do(func() {
		regexPattern = func() gopurs_runtime.Value {
v_0_0 := gopurs_runtime.Apply2(pkg_Data_String_Regex.Get_regex(), gopurs_runtime.Str("(hello|world)[0-9]+"), pkg_Data_String_Regex_Flags.Get_noFlags())
_ = v_0_0
var __t1 gopurs_runtime.Value
{
if (v_0_0.Type == 9 && v_0_0.IntVal == 4096564120) {
__t1 = (*pkg_Data_Either.Data_Data_Either_Right)(v_0_0.UnsafePtr).V0
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}()
	})
	return regexPattern
}

var runStringOps gopurs_runtime.Value
var once_runStringOps sync.Once
func Get_runStringOps() gopurs_runtime.Value {
	once_runStringOps.Do(func() {
		runStringOps = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
var loop_1_0 gopurs_runtime.Value
loop_1_0 = gopurs_runtime.Func(func(v_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v2_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
loop_1_0:
for {
if false { continue loop_1_0 }
var v_2 gopurs_runtime.Value = v_2_loop
_ = v_2
var v1_3 gopurs_runtime.Value = v1_3_loop
_ = v1_3
var v2_4 gopurs_runtime.Value = v2_4_loop
_ = v2_4
var __t2 gopurs_runtime.Value
{
if v_2.IntVal == 0 {
__t2 = v2_4
goto end_branch_2
} else {

}
}
{
concatted_5_1 := v1_3.StrVal() + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), v_2).StrVal() + "world" + gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Int(v_2.IntVal + 1)).StrVal()
_ = concatted_5_1
v_2_loop = gopurs_runtime.Int(v_2.IntVal - 1)
v1_3_loop = gopurs_runtime.Apply2(pkg_Data_String_CodePoints.Get_take(), gopurs_runtime.Int(10), gopurs_runtime.Str(concatted_5_1))
v2_4_loop = gopurs_runtime.Int(v2_4.IntVal + gopurs_runtime.Int(int64(len(gopurs_runtime.Apply2(pkg_Data_String_Common.Get_split(), gopurs_runtime.Str("e"), gopurs_runtime.Apply3(pkg_Data_String_Regex.Get_replace(), Get_regexPattern(), gopurs_runtime.Str("matched"), gopurs_runtime.Str(concatted_5_1))).PtrVal().([]gopurs_runtime.Value)))).IntVal)
continue loop_1_0
__t2 = gopurs_runtime.Value{}
}
end_branch_2:
return __t2
}
}()
})
})
})
return gopurs_runtime.Apply3(loop_1_0, n_0, gopurs_runtime.Str("hello"), gopurs_runtime.Int(0))
}()
})
	})
	return runStringOps
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("String Operations (1k Regex/Split):"))
	})
	return describe
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply(Get_runStringOps(), gopurs_runtime.Int(1000))))
	})
	return act
}





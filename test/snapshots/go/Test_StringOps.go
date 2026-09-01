package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_StringOps_regexPattern gopurs_runtime.Value
var once_Test_StringOps_regexPattern sync.Once

func Get_Test_StringOps_regexPattern() gopurs_runtime.Value {
	once_Test_StringOps_regexPattern.Do(func() {
		cache_Test_StringOps_regexPattern = func() gopurs_runtime.Value {
			// TAST (Let): v_0_0 shape=App(Var) bindingType=(ADT ["Data","Either","Either"] [String, (ADT ["Data","String","Regex","Regex"] [])])
			v_0_0 := gopurs_runtime.Apply2(Get_Data_String_Regex_regex(), gopurs_runtime.Str("(hello|world)[0-9]+"), Get_Data_String_Regex_Flags_noFlags())
			_ = v_0_0
			var __t1 gopurs_runtime.Value
			{
				if v_0_0.Type == 9 && v_0_0.IntVal == 2465973597 {
					__t1 = (*Constructor_Data_Either_Right[gopurs_runtime.Value, gopurs_runtime.Value])(v_0_0.UnsafePtr).V0
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
	return cache_Test_StringOps_regexPattern
}

var cache_Test_StringOps_runStringOps gopurs_runtime.Value
var once_Test_StringOps_runStringOps sync.Once

func Get_Test_StringOps_runStringOps() gopurs_runtime.Value {
	once_Test_StringOps_runStringOps.Do(func() {
		cache_Test_StringOps_runStringOps = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_StringOps_runStringOps(n_0_box.IntVal))
		})
	})
	return cache_Test_StringOps_runStringOps
}

var cache_Test_StringOps_describe gopurs_runtime.Value
var once_Test_StringOps_describe sync.Once

func Get_Test_StringOps_describe() gopurs_runtime.Value {
	once_Test_StringOps_describe.Do(func() {
		cache_Test_StringOps_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("String Operations (1k Regex/Split):"))
	})
	return cache_Test_StringOps_describe
}

var cache_Test_StringOps_act gopurs_runtime.Value
var once_Test_StringOps_act sync.Once

func Get_Test_StringOps_act() gopurs_runtime.Value {
	once_Test_StringOps_act.Do(func() {
		cache_Test_StringOps_act = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_StringOps_runStringOps(1000))).StrVal()))
	})
	return cache_Test_StringOps_act
}

func Call_Test_StringOps_runStringOps(n_0_loop int64) int64 {
	var n_0 int64 = n_0_loop
	_ = n_0
	var Call_local_Test_StringOps_loop_1_0_0 func(gopurs_runtime.Value, gopurs_runtime.Value, int64) int64
	_ = Call_local_Test_StringOps_loop_1_0_0
	var loop_1_0_0 gopurs_runtime.Value
	_ = loop_1_0_0
	Call_local_Test_StringOps_loop_1_0_0 = func(v_2_loop gopurs_runtime.Value, v1_3_loop gopurs_runtime.Value, v2_4_loop int64) int64 {
	loop_1_0_0:
		for {
			if false {
				continue loop_1_0_0
			}
			var v_2 gopurs_runtime.Value = v_2_loop
			_ = v_2
			var v1_3 gopurs_runtime.Value = v1_3_loop
			_ = v1_3
			var v2_4 int64 = v2_4_loop
			_ = v2_4
			var __t2 int64
			{
				if (v_2.IntVal) == (0) {
					__t2 = v2_4
					goto end_branch_2
				} else {

				}
			}
			{
				// TAST (Let): concatted_5_1 shape=Other bindingType=String
				concatted_5_1 := (((v1_3.StrVal()) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), v_2).StrVal())) + ("world")) + (gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int((v_2.IntVal)+(1))).StrVal())
				_ = concatted_5_1
				v_2_loop = gopurs_runtime.Int((v_2.IntVal) - (1))
				v1_3_loop = gopurs_runtime.Str(gopurs_runtime.Apply2(Get_Data_String_CodePoints_take(), gopurs_runtime.Int(10), gopurs_runtime.Str(concatted_5_1)).StrVal())
				v2_4_loop = (v2_4) + (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(gopurs_runtime.Apply2(Get_Data_String_Common_split(), gopurs_runtime.Str("e"), gopurs_runtime.Str(gopurs_runtime.Apply3(Get_Data_String_Regex_replace(), Get_Test_StringOps_regexPattern(), gopurs_runtime.Str("matched"), gopurs_runtime.Str(concatted_5_1)).StrVal()))))).IntVal)
				continue loop_1_0_0
				__t2 = gopurs_runtime.Value{}.IntVal
			}
		end_branch_2:
			return __t2
		}
	}
	loop_1_0_0 = gopurs_runtime.Func(func(v_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
		return gopurs_runtime.Func(func(v1_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Func(func(v2_4_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(Call_local_Test_StringOps_loop_1_0_0(v_2_loop_val, v1_3_loop_val, v2_4_loop_val.IntVal))
			})
		})
	})
	return Call_local_Test_StringOps_loop_1_0_0(gopurs_runtime.Int(n_0), gopurs_runtime.Str("hello"), 0)
}

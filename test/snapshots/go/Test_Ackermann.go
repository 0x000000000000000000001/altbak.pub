package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_Ackermann_logShow gopurs_runtime.Value
var once_Test_Ackermann_logShow sync.Once

func Get_Test_Ackermann_logShow() gopurs_runtime.Value {
	once_Test_Ackermann_logShow.Do(func() {
		cache_Test_Ackermann_logShow = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Ackermann_logShow(a_0_box.IntVal)
		})
	})
	return cache_Test_Ackermann_logShow
}

var cache_Test_Ackermann_describe gopurs_runtime.Value
var once_Test_Ackermann_describe sync.Once

func Get_Test_Ackermann_describe() gopurs_runtime.Value {
	once_Test_Ackermann_describe.Do(func() {
		cache_Test_Ackermann_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Ackermann (3, 4):"))
	})
	return cache_Test_Ackermann_describe
}

var cache_Test_Ackermann_ackermann gopurs_runtime.Value
var once_Test_Ackermann_ackermann sync.Once

func Get_Test_Ackermann_ackermann() gopurs_runtime.Value {
	once_Test_Ackermann_ackermann.Do(func() {
		cache_Test_Ackermann_ackermann = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_Ackermann_ackermann(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Test_Ackermann_ackermann
}

var cache_Test_Ackermann_act gopurs_runtime.Value
var once_Test_Ackermann_act sync.Once

func Get_Test_Ackermann_act() gopurs_runtime.Value {
	once_Test_Ackermann_act.Do(func() {
		cache_Test_Ackermann_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(3))
			_ = __local_var_0_0
			dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = dummy_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_Ackermann_ackermann(dummy_1_1.IntVal, 4))).StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Test_Ackermann_act
}

func Call_Test_Ackermann_logShow(a_0_loop int64) gopurs_runtime.Value {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(a_0)).StrVal()))
}

func Call_Test_Ackermann_ackermann(v_0_loop int64, v1_1_loop int64) int64 {
ackermann:
	for {
		if false {
			continue ackermann
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var __t0 int64
		{
			if (v_0) == (0) {
				__t0 = (v1_1) + (1)
				goto end_branch_0
			} else {

			}
		}
		{
			if (v1_1) == (0) {
				v_0_loop = (v_0) - (1)
				v1_1_loop = 1
				continue ackermann
				__t0 = gopurs_runtime.Value{}.IntVal
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (1)
			v1_1_loop = Call_Test_Ackermann_ackermann(v_0, (v1_1)-(1))
			continue ackermann
			__t0 = gopurs_runtime.Value{}.IntVal
		}
	end_branch_0:
		return __t0
	}
}

package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_TCO_logShow gopurs_runtime.Value
var once_Test_TCO_logShow sync.Once

func Get_Test_TCO_logShow() gopurs_runtime.Value {
	once_Test_TCO_logShow.Do(func() {
		cache_Test_TCO_logShow = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_TCO_logShow(a_0_box.IntVal)
		})
	})
	return cache_Test_TCO_logShow
}

var cache_Test_TCO_describe gopurs_runtime.Value
var once_Test_TCO_describe sync.Once

func Get_Test_TCO_describe() gopurs_runtime.Value {
	once_Test_TCO_describe.Do(func() {
		cache_Test_TCO_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Tail Call Optimization (100k calls):"))
	})
	return cache_Test_TCO_describe
}

var cache_Test_TCO_deepTailRec gopurs_runtime.Value
var once_Test_TCO_deepTailRec sync.Once

func Get_Test_TCO_deepTailRec() gopurs_runtime.Value {
	once_Test_TCO_deepTailRec.Do(func() {
		cache_Test_TCO_deepTailRec = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_TCO_deepTailRec(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Test_TCO_deepTailRec
}

var cache_Test_TCO_act gopurs_runtime.Value
var once_Test_TCO_act sync.Once

func Get_Test_TCO_act() gopurs_runtime.Value {
	once_Test_TCO_act.Do(func() {
		cache_Test_TCO_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(100000))
			_ = __local_var_0_0
			dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = dummy_1_1
			return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_TCO_deepTailRec(dummy_1_1.IntVal, 0))).StrVal())), gopurs_runtime.Value{})
		})
	})
	return cache_Test_TCO_act
}

func Call_Test_TCO_logShow(a_0_loop int64) gopurs_runtime.Value {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(a_0)).StrVal()))
}

func Call_Test_TCO_deepTailRec(v_0_loop int64, v1_1_loop int64) int64 {
deepTailRec:
	for {
		if false {
			continue deepTailRec
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 int64 = v1_1_loop
		_ = v1_1
		var __t0 int64
		{
			if (v_0) == (0) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (1)
			v1_1_loop = (v1_1) + (gopurs_runtime.Apply2(Get_Data_EuclideanRing_intMod(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(3)).IntVal)
			continue deepTailRec
			__t0 = gopurs_runtime.Value{}.IntVal
		}
	end_branch_0:
		return __t0
	}
}

package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_Fib_logShow gopurs_runtime.Value
var once_Test_Fib_logShow sync.Once

func Get_Test_Fib_logShow() gopurs_runtime.Value {
	once_Test_Fib_logShow.Do(func() {
		cache_Test_Fib_logShow = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_Fib_logShow(a_0_box.IntVal)
		})
	})
	return cache_Test_Fib_logShow
}

var cache_Test_Fib_fib gopurs_runtime.Value
var once_Test_Fib_fib sync.Once

func Get_Test_Fib_fib() gopurs_runtime.Value {
	once_Test_Fib_fib.Do(func() {
		cache_Test_Fib_fib = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_Fib_fib(v_0_box.IntVal))
		})
	})
	return cache_Test_Fib_fib
}

var cache_Test_Fib_describe gopurs_runtime.Value
var once_Test_Fib_describe sync.Once

func Get_Test_Fib_describe() gopurs_runtime.Value {
	once_Test_Fib_describe.Do(func() {
		cache_Test_Fib_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Fibonacci:"))
	})
	return cache_Test_Fib_describe
}

var cache_Test_Fib_act gopurs_runtime.Value
var once_Test_Fib_act sync.Once

func Get_Test_Fib_act() gopurs_runtime.Value {
	once_Test_Fib_act.Do(func() {
		cache_Test_Fib_act = func() gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(10))
			_ = __local_var_0_0
			return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
				dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
				_ = dummy_1_1
				return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_Fib_fib(dummy_1_1.IntVal))).StrVal())), gopurs_runtime.Value{})
			})
		}()
	})
	return cache_Test_Fib_act
}

func Call_Test_Fib_logShow(a_0_loop int64) gopurs_runtime.Value {
	var a_0 int64 = a_0_loop
	_ = a_0
	return gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(a_0)).StrVal()))
}

func Call_Test_Fib_fib(v_0_loop int64) int64 {
fib:
	for {
		if false {
			continue fib
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var __t0 int64
		{
			if (v_0) == (0) {
				__t0 = 0
				goto end_branch_0
			} else {

			}
		}
		{
			if (v_0) == (1) {
				__t0 = 1
				goto end_branch_0
			} else {

			}
		}
		{
			__t0 = (Call_Test_Fib_fib((v_0) - (1))) + (Call_Test_Fib_fib((v_0) - (2)))
		}
	end_branch_0:
		return __t0
	}
}

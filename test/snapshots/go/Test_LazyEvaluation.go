package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_LazyEvaluation_Lazy gopurs_runtime.Value
var once_Test_LazyEvaluation_Lazy sync.Once

func Get_Test_LazyEvaluation_Lazy() gopurs_runtime.Value {
	once_Test_LazyEvaluation_Lazy.Do(func() {
		cache_Test_LazyEvaluation_Lazy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_LazyEvaluation_Lazy(x_0_box)
		})
	})
	return cache_Test_LazyEvaluation_Lazy
}

var cache_Test_LazyEvaluation_force gopurs_runtime.Value
var once_Test_LazyEvaluation_force sync.Once

func Get_Test_LazyEvaluation_force() gopurs_runtime.Value {
	once_Test_LazyEvaluation_force.Do(func() {
		cache_Test_LazyEvaluation_force = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_LazyEvaluation_force(v_0_box)
		})
	})
	return cache_Test_LazyEvaluation_force
}

var cache_Test_LazyEvaluation_force__819661500 gopurs_runtime.Value
var once_Test_LazyEvaluation_force__819661500 sync.Once

func Get_Test_LazyEvaluation_force__819661500() gopurs_runtime.Value {
	once_Test_LazyEvaluation_force__819661500.Do(func() {
		cache_Test_LazyEvaluation_force__819661500 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_LazyEvaluation_force__819661500(v_0_box))
		})
	})
	return cache_Test_LazyEvaluation_force__819661500
}

var cache_Test_LazyEvaluation_describe gopurs_runtime.Value
var once_Test_LazyEvaluation_describe sync.Once

func Get_Test_LazyEvaluation_describe() gopurs_runtime.Value {
	once_Test_LazyEvaluation_describe.Do(func() {
		cache_Test_LazyEvaluation_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Lazy Evaluation (1M Thunks Forced, 1k Depth):"))
	})
	return cache_Test_LazyEvaluation_describe
}

var cache_Test_LazyEvaluation_go__defer gopurs_runtime.Value
var once_Test_LazyEvaluation_go__defer sync.Once

func Get_Test_LazyEvaluation_go__defer() gopurs_runtime.Value {
	once_Test_LazyEvaluation_go__defer.Do(func() {
		cache_Test_LazyEvaluation_go__defer = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_LazyEvaluation_go__defer(x_0_box)
		})
	})
	return cache_Test_LazyEvaluation_go__defer
}

var cache_Test_LazyEvaluation_defer__3628429125 gopurs_runtime.Value
var once_Test_LazyEvaluation_defer__3628429125 sync.Once

func Get_Test_LazyEvaluation_defer__3628429125() gopurs_runtime.Value {
	once_Test_LazyEvaluation_defer__3628429125.Do(func() {
		cache_Test_LazyEvaluation_defer__3628429125 = gopurs_runtime.Func(func(__eta_norm_0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_LazyEvaluation_defer__3628429125(__eta_norm_0_0_box)
		})
	})
	return cache_Test_LazyEvaluation_defer__3628429125
}

var cache_Test_LazyEvaluation_buildThunks gopurs_runtime.Value
var once_Test_LazyEvaluation_buildThunks sync.Once

func Get_Test_LazyEvaluation_buildThunks() gopurs_runtime.Value {
	once_Test_LazyEvaluation_buildThunks.Do(func() {
		cache_Test_LazyEvaluation_buildThunks = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_LazyEvaluation_buildThunks(v_0_box.IntVal, v1_1_box)
		})
	})
	return cache_Test_LazyEvaluation_buildThunks
}

var cache_Test_LazyEvaluation_runManyTimes gopurs_runtime.Value
var once_Test_LazyEvaluation_runManyTimes sync.Once

func Get_Test_LazyEvaluation_runManyTimes() gopurs_runtime.Value {
	once_Test_LazyEvaluation_runManyTimes.Do(func() {
		cache_Test_LazyEvaluation_runManyTimes = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
			return gopurs_runtime.Int(Call_Test_LazyEvaluation_runManyTimes(v_0_box.IntVal, v1_1_box.IntVal))
		})
	})
	return cache_Test_LazyEvaluation_runManyTimes
}

var cache_Test_LazyEvaluation_act gopurs_runtime.Value
var once_Test_LazyEvaluation_act sync.Once

func Get_Test_LazyEvaluation_act() gopurs_runtime.Value {
	once_Test_LazyEvaluation_act.Do(func() {
		cache_Test_LazyEvaluation_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(1000))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_LazyEvaluation_runManyTimes(__local_var_1_1.IntVal, 0))).StrVal())
		})
	})
	return cache_Test_LazyEvaluation_act
}

func Call_Test_LazyEvaluation_Lazy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Test_LazyEvaluation_force(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var v_0 gopurs_runtime.Value = v_0_loop
	_ = v_0
	return gopurs_runtime.Apply(v_0, Get_Data_Unit_unit())
}

func Call_Test_LazyEvaluation_force__819661500(v_0_loop gopurs_runtime.Value) int64 {
force__819661500:
	for {
		if false {
			continue force__819661500
		}
		var v_0 gopurs_runtime.Value = v_0_loop
		_ = v_0
		return gopurs_runtime.Apply(v_0, Get_Data_Unit_unit()).IntVal
	}
}

func Call_Test_LazyEvaluation_go__defer(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
	var x_0 gopurs_runtime.Value = x_0_loop
	_ = x_0
	return x_0
}

func Call_Test_LazyEvaluation_defer__3628429125(__eta_norm_0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
defer__3628429125:
	for {
		if false {
			continue defer__3628429125
		}
		var __eta_norm_0_0 gopurs_runtime.Value = __eta_norm_0_0_loop
		_ = __eta_norm_0_0
		return __eta_norm_0_0
	}
}

func Call_Test_LazyEvaluation_buildThunks(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
buildThunks:
	for {
		if false {
			continue buildThunks
		}
		var v_0 int64 = v_0_loop
		_ = v_0
		var v1_1 gopurs_runtime.Value = v1_1_loop
		_ = v1_1
		var __t0 gopurs_runtime.Value
		{
			if (v_0) == (0) {
				__t0 = v1_1
				goto end_branch_0
			} else {

			}
		}
		{
			v_0_loop = (v_0) - (1)
			v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int((gopurs_runtime.Apply(v1_1, Get_Data_Unit_unit()).IntVal) + (1))
			})
			continue buildThunks
			__t0 = gopurs_runtime.Value{}
		}
	end_branch_0:
		return __t0
	}
}

func Call_Test_LazyEvaluation_runManyTimes(v_0_loop int64, v1_1_loop int64) int64 {
runManyTimes:
	for {
		if false {
			continue runManyTimes
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
			v1_1_loop = (v1_1) + (gopurs_runtime.Apply(Call_Test_LazyEvaluation_buildThunks(1000, gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
				return gopurs_runtime.Int(0)
			})), Get_Data_Unit_unit()).IntVal)
			continue runManyTimes
			__t0 = gopurs_runtime.Value{}.IntVal
		}
	end_branch_0:
		return __t0
	}
}

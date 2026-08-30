package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_LazyEvaluationFFI_describe gopurs_runtime.Value
var once_Test_LazyEvaluationFFI_describe sync.Once

func Get_Test_LazyEvaluationFFI_describe() gopurs_runtime.Value {
	once_Test_LazyEvaluationFFI_describe.Do(func() {
		cache_Test_LazyEvaluationFFI_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Lazy Evaluation FFI (1M Thunks Forced, 1k Depth):"))
	})
	return cache_Test_LazyEvaluationFFI_describe
}

var cache_Test_LazyEvaluationFFI_act gopurs_runtime.Value
var once_Test_LazyEvaluationFFI_act sync.Once

func Get_Test_LazyEvaluationFFI_act() gopurs_runtime.Value {
	once_Test_LazyEvaluationFFI_act.Do(func() {
		cache_Test_LazyEvaluationFFI_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(1000))
			_ = __local_var_0_0
			dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = dummy_1_1
			return gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Test_LazyEvaluationFFI_runLazyEvaluationFFI(), gopurs_runtime.Int(dummy_1_1.IntVal)).IntVal))
		})
	})
	return cache_Test_LazyEvaluationFFI_act
}

func Get_Test_LazyEvaluationFFI_runLazyEvaluationFFI() gopurs_runtime.Value {
	return _Gopurs_Test_LazyEvaluationFFI_RunLazyEvaluationFFI
}

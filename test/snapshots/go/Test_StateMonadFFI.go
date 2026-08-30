package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_StateMonadFFI_describe gopurs_runtime.Value
var once_Test_StateMonadFFI_describe sync.Once

func Get_Test_StateMonadFFI_describe() gopurs_runtime.Value {
	once_Test_StateMonadFFI_describe.Do(func() {
		cache_Test_StateMonadFFI_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("State Monad FFI (1.2k Binds, 60 Stack Depth):"))
	})
	return cache_Test_StateMonadFFI_describe
}

var cache_Test_StateMonadFFI_act gopurs_runtime.Value
var once_Test_StateMonadFFI_act sync.Once

func Get_Test_StateMonadFFI_act() gopurs_runtime.Value {
	once_Test_StateMonadFFI_act.Do(func() {
		cache_Test_StateMonadFFI_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(60))
			_ = __local_var_0_0
			dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = dummy_1_1
			return gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(gopurs_runtime.Apply(Get_Test_StateMonadFFI_runStateMonadFFI(), gopurs_runtime.Int(dummy_1_1.IntVal)).IntVal))
		})
	})
	return cache_Test_StateMonadFFI_act
}

func Get_Test_StateMonadFFI_runStateMonadFFI() gopurs_runtime.Value {
	return _Gopurs_Test_StateMonadFFI_RunStateMonadFFI
}

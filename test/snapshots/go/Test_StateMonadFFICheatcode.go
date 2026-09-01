package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_StateMonadFFICheatcode_describe gopurs_runtime.Value
var once_Test_StateMonadFFICheatcode_describe sync.Once

func Get_Test_StateMonadFFICheatcode_describe() gopurs_runtime.Value {
	once_Test_StateMonadFFICheatcode_describe.Do(func() {
		cache_Test_StateMonadFFICheatcode_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("State Monad FFICheatcode (1.2k Binds, 60 Stack Depth):"))
	})
	return cache_Test_StateMonadFFICheatcode_describe
}

var cache_Test_StateMonadFFICheatcode_act gopurs_runtime.Value
var once_Test_StateMonadFFICheatcode_act sync.Once

func Get_Test_StateMonadFFICheatcode_act() gopurs_runtime.Value {
	once_Test_StateMonadFFICheatcode_act.Do(func() {
		cache_Test_StateMonadFFICheatcode_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(60))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Apply(Get_Test_StateMonadFFICheatcode_runStateMonadFFICheatcode(), gopurs_runtime.Int(__local_var_1_1.IntVal))).StrVal())
		})
	})
	return cache_Test_StateMonadFFICheatcode_act
}

func Get_Test_StateMonadFFICheatcode_runStateMonadFFICheatcode() gopurs_runtime.Value {
	return _Gopurs_Test_StateMonadFFICheatcode_RunStateMonadFFICheatcode
}

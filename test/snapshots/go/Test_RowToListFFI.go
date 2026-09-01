package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_RowToListFFI_describe gopurs_runtime.Value
var once_Test_RowToListFFI_describe sync.Once

func Get_Test_RowToListFFI_describe() gopurs_runtime.Value {
	once_Test_RowToListFFI_describe.Do(func() {
		cache_Test_RowToListFFI_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("RowToList FFI (Keys Count):"))
	})
	return cache_Test_RowToListFFI_describe
}

var cache_Test_RowToListFFI_act gopurs_runtime.Value
var once_Test_RowToListFFI_act sync.Once

func Get_Test_RowToListFFI_act() gopurs_runtime.Value {
	once_Test_RowToListFFI_act.Do(func() {
		cache_Test_RowToListFFI_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(0))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Apply(Get_Test_RowToListFFI_runRowToListFFI(), gopurs_runtime.Int(__local_var_1_1.IntVal))).StrVal())
		})
	})
	return cache_Test_RowToListFFI_act
}

func Get_Test_RowToListFFI_runRowToListFFI() gopurs_runtime.Value {
	return _Gopurs_Test_RowToListFFI_RunRowToListFFI
}

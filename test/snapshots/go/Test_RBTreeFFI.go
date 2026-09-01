package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_RBTreeFFI_describe gopurs_runtime.Value
var once_Test_RBTreeFFI_describe sync.Once

func Get_Test_RBTreeFFI_describe() gopurs_runtime.Value {
	once_Test_RBTreeFFI_describe.Do(func() {
		cache_Test_RBTreeFFI_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("Red-Black Tree FFI (100k Worst-Case Insertions):"))
	})
	return cache_Test_RBTreeFFI_describe
}

var cache_Test_RBTreeFFI_act gopurs_runtime.Value
var once_Test_RBTreeFFI_act sync.Once

func Get_Test_RBTreeFFI_act() gopurs_runtime.Value {
	once_Test_RBTreeFFI_act.Do(func() {
		cache_Test_RBTreeFFI_act = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
			// TAST (Let): __local_var_0_0 shape=App(Var) bindingType=Any
			__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(100000))
			_ = __local_var_0_0
			__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
			_ = __local_var_1_1
			return gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Apply(Get_Test_RBTreeFFI_runRBTreeFFI(), gopurs_runtime.Int(__local_var_1_1.IntVal))).StrVal())
		})
	})
	return cache_Test_RBTreeFFI_act
}

func Get_Test_RBTreeFFI_runRBTreeFFI() gopurs_runtime.Value {
	return _Gopurs_Test_RBTreeFFI_RunRBTreeFFI
}

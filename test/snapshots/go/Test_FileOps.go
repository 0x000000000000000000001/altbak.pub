package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_FileOps_loopIO gopurs_runtime.Value
var once_Test_FileOps_loopIO sync.Once

func Get_Test_FileOps_loopIO() gopurs_runtime.Value {
	once_Test_FileOps_loopIO.Do(func() {
		cache_Test_FileOps_loopIO = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
			return Call_Test_FileOps_loopIO(n_0_box.IntVal)
		})
	})
	return cache_Test_FileOps_loopIO
}

var cache_Test_FileOps_describe gopurs_runtime.Value
var once_Test_FileOps_describe sync.Once

func Get_Test_FileOps_describe() gopurs_runtime.Value {
	once_Test_FileOps_describe.Do(func() {
		cache_Test_FileOps_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("File I/O (10k writes/reads):"))
	})
	return cache_Test_FileOps_describe
}

var cache_Test_FileOps_act gopurs_runtime.Value
var once_Test_FileOps_act sync.Once

func Get_Test_FileOps_act() gopurs_runtime.Value {
	once_Test_FileOps_act.Do(func() {
		cache_Test_FileOps_act = Call_Test_FileOps_loopIO(10000)
	})
	return cache_Test_FileOps_act
}

func Call_Test_FileOps_loopIO(n_0_loop int64) gopurs_runtime.Value {
	var n_0 int64 = n_0_loop
	_ = n_0
	// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
	__local_var_1_0 := gopurs_runtime.Apply2(Get_Test_FileOps_writeFileSync(), gopurs_runtime.Str("var/iotest.txt"), gopurs_runtime.Str("Hello IO Benchmarks!"))
	_ = __local_var_1_0
	return gopurs_runtime.Apply2(Get_Test_FileOps_loopE(), gopurs_runtime.Int(n_0), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
		_dollar___unused_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
		_ = _dollar___unused_2_1
		_dollar___unused_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Test_FileOps_readFileSync(), gopurs_runtime.Str("var/iotest.txt")), gopurs_runtime.Value{})
		_ = _dollar___unused_3_2
		return Get_Data_Unit_unit()
	}))
}

func Get_Test_FileOps_loopE() gopurs_runtime.Value {
	return _Gopurs_Test_FileOps_LoopE
}

func Get_Test_FileOps_readFileSync() gopurs_runtime.Value {
	return _Gopurs_Test_FileOps_ReadFileSync
}

func Get_Test_FileOps_writeFileSync() gopurs_runtime.Value {
	return _Gopurs_Test_FileOps_WriteFileSync
}

package Test_FileOps

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect_Console "gopurs/output/Effect.Console"
)

var loopIO gopurs_runtime.Value
var once_loopIO sync.Once
func Get_loopIO() gopurs_runtime.Value {
	once_loopIO.Do(func() {
		loopIO = gopurs_runtime.Func(func(n_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var n_0 gopurs_runtime.Value = n_0_loop
_ = n_0
__local_var_1_0 := gopurs_runtime.Apply2(Get_writeFileSync(), gopurs_runtime.Str("var/iotest.txt"), gopurs_runtime.Str("Hello IO Benchmarks!"))
_ = __local_var_1_0
return gopurs_runtime.Apply2(Get_loopE(), n_0, gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_2_1 := gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Value{})
_ = _dollar__unused_2_1
_dollar__unused_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_readFileSync(), gopurs_runtime.Str("var/iotest.txt")), gopurs_runtime.Value{})
_ = _dollar__unused_3_2
return pkg_Data_Unit.Get_unit()
}))
}()
})
	})
	return loopIO
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("File I/O (10k writes/reads):"))
	})
	return describe
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = gopurs_runtime.Apply(Get_loopIO(), gopurs_runtime.Int(10000))
	})
	return act
}



func Get_loopE() gopurs_runtime.Value {
	return _Gopurs_LoopE
}

func Get_readFileSync() gopurs_runtime.Value {
	return _Gopurs_ReadFileSync
}

func Get_writeFileSync() gopurs_runtime.Value {
	return _Gopurs_WriteFileSync
}

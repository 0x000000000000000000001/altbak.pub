package Test_FileOps

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var cache_discard gopurs_runtime.Value
var once_discard sync.Once
func Get_discard() gopurs_runtime.Value {
	once_discard.Do(func() {
		cache_discard = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard
}

var cache_loopIO gopurs_runtime.Value
var once_loopIO sync.Once
func Get_loopIO() gopurs_runtime.Value {
	once_loopIO.Do(func() {
		cache_loopIO = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_loopIO(n_0_box.IntVal)
})
	})
	return cache_loopIO
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("File I/O (10k writes/reads):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = Call_loopIO(10000)
	})
	return cache_act
}

func Call_loopIO(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
return gopurs_runtime.Apply2(Get_loopE(), gopurs_runtime.Int(n_0), gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply2(Get_writeFileSync(), gopurs_runtime.Str("var/iotest.txt"), gopurs_runtime.Str("Hello IO Benchmarks!")), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(Get_readFileSync(), gopurs_runtime.Str("var/iotest.txt")), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), pkg_Data_Unit.Get_unit())
}))
})))
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

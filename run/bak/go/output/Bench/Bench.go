package Bench

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	pkg_Effect "gopurs/output/Effect"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_discard1 gopurs_runtime.Value
var once_discard1 sync.Once
func Get_discard1() gopurs_runtime.Value {
	once_discard1.Do(func() {
		cache_discard1 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect_Aff.Get_bindAff())
	})
	return cache_discard1
}

var cache_discard2 gopurs_runtime.Value
var once_discard2 sync.Once
func Get_discard2() gopurs_runtime.Value {
	once_discard2.Do(func() {
		cache_discard2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard2
}

var cache_runBenchAff gopurs_runtime.Value
var once_runBenchAff sync.Once
func Get_runBenchAff() gopurs_runtime.Value {
	once_runBenchAff.Do(func() {
		cache_runBenchAff = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runBenchAff(describe_0_box, act_1_box)
})
	})
	return cache_runBenchAff
}

var cache_runBench gopurs_runtime.Value
var once_runBench sync.Once
func Get_runBench() gopurs_runtime.Value {
	once_runBench.Do(func() {
		cache_runBench = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runBench(describe_0_box, act_1_box)
})
	})
	return cache_runBench
}

func Call_runBenchAff(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
return gopurs_runtime.Apply2(Get_discard1(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a"))), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard1(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), describe_0), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard1(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("\x0a(Output)\x0a"))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_bindAff(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), Get_benchNow()), gopurs_runtime.Func(func(t1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard1(), act_1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_bindAff(), "bind"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), Get_benchNow()), gopurs_runtime.Func(func(t2_7 gopurs_runtime.Value) gopurs_runtime.Value {
dt_8_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), t2_7, t1_5)
_ = dt_8_0
return gopurs_runtime.Apply2(Get_discard1(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_monadEffectAff(), "liftEffect"), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("\x0a(Execution time)\x0a\x0a"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(Get_formatNumber(), dt_8_0), gopurs_runtime.Str(" μs\x0a"))))), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_applicativeAff(), "pure"), dt_8_0)
}))
}))
}))
}))
}))
}))
}))
}

func Call_runBench(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
return gopurs_runtime.Apply2(Get_discard2(), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a")), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard2(), describe_0, gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard2(), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("\x0a(Output)\x0a")), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), Get_benchNow(), gopurs_runtime.Func(func(t1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard2(), act_1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), Get_benchNow(), gopurs_runtime.Func(func(t2_7 gopurs_runtime.Value) gopurs_runtime.Value {
dt_8_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), t2_7, t1_5)
_ = dt_8_0
return gopurs_runtime.Apply2(Get_discard2(), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("\x0a(Execution time)\x0a\x0a"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(Get_formatNumber(), dt_8_0), gopurs_runtime.Str(" μs\x0a")))), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), dt_8_0)
}))
}))
}))
}))
}))
}))
}))
}

func Get_benchNow() gopurs_runtime.Value {
	return _Gopurs_BenchNow
}

func Get_formatNumber() gopurs_runtime.Value {
	return _Gopurs_FormatNumber
}

func Get_opaque() gopurs_runtime.Value {
	return _Gopurs_Opaque
}

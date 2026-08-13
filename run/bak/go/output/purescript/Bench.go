package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Bench_liftEffect gopurs_runtime.Value
var once_Bench_liftEffect sync.Once
func Get_Bench_liftEffect() gopurs_runtime.Value {
	once_Bench_liftEffect.Do(func() {
		cache_Bench_liftEffect = Get_Effect_Aff__liftEffect()
	})
	return cache_Bench_liftEffect
}

var cache_Bench_runBenchAff gopurs_runtime.Value
var once_Bench_runBenchAff sync.Once
func Get_Bench_runBenchAff() gopurs_runtime.Value {
	once_Bench_runBenchAff.Do(func() {
		cache_Bench_runBenchAff = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Bench_runBenchAff(describe_0_box, act_1_box)
})
	})
	return cache_Bench_runBenchAff
}

var cache_Bench_runBench gopurs_runtime.Value
var once_Bench_runBench sync.Once
func Get_Bench_runBench() gopurs_runtime.Value {
	once_Bench_runBench.Do(func() {
		cache_Bench_runBench = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Bench_runBench(describe_0_box, act_1_box)
})
	})
	return cache_Bench_runBench
}

func Call_Bench_runBenchAff(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a"))), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), describe_0), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("\x0a(Output)\x0a"))), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), act_1, gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), Get_Bench_benchNow()), gopurs_runtime.Func(func(t2_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): dt_8_0 -> float64
dt_8_0 := (t2_7.FloatVal()) - (t1_5.FloatVal())
_ = dt_8_0
return gopurs_runtime.Apply2(Get_Effect_Aff__bind(), gopurs_runtime.Apply(Get_Effect_Aff__liftEffect(), gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("\x0a(Execution time)\x0a\x0a") + (gopurs_runtime.Apply(Get_Bench_formatNumber(), gopurs_runtime.Float(dt_8_0)).StrVal())) + (" μs\x0a")))), gopurs_runtime.Func(func(_dollar__unused_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Effect_Aff__pure(), gopurs_runtime.Float(dt_8_0))
}))
}))
}))
}))
}))
}))
}))
}

func Call_Bench_runBench(describe_0_loop gopurs_runtime.Value, act_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var describe_0 gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 gopurs_runtime.Value = act_1_loop
_ = act_1
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("--------------------------------------------------\x0a\x0a(Test)\x0a"))
_ = __local_var_2_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(__local_var_2_0, gopurs_runtime.Value{})
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(describe_0, gopurs_runtime.Value{})
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("\x0a(Output)\x0a")), gopurs_runtime.Value{})
_ = __local_var_5_3
__local_var_6_4 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_6_4
__local_var_7_5 := gopurs_runtime.Apply(act_1, gopurs_runtime.Value{})
_ = __local_var_7_5
__local_var_8_6 := gopurs_runtime.Apply(Get_Bench_benchNow(), gopurs_runtime.Value{})
_ = __local_var_8_6
// TAST (Let): dt_9_7 -> float64
dt_9_7 := (__local_var_8_6.FloatVal()) - (__local_var_6_4.FloatVal())
_ = dt_9_7
return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_10_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str((("\x0a(Execution time)\x0a\x0a") + (gopurs_runtime.Apply(Get_Bench_formatNumber(), gopurs_runtime.Float(dt_9_7)).StrVal())) + (" μs\x0a"))), gopurs_runtime.Value{})
_ = __local_var_10_8
return gopurs_runtime.Float(dt_9_7)
}), gopurs_runtime.Value{})
})
}

func Get_Bench_benchNow() gopurs_runtime.Value {
	return _Gopurs_Bench_BenchNow
}

func Get_Bench_formatNumber() gopurs_runtime.Value {
	return _Gopurs_Bench_FormatNumber
}

func Get_Bench_opaque() gopurs_runtime.Value {
	return _Gopurs_Bench_Opaque
}

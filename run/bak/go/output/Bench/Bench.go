package Bench

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
)

var runBench gopurs_runtime.Value
var once_runBench sync.Once
func Get_runBench() gopurs_runtime.Value {
	once_runBench.Do(func() {
		runBench = gopurs_runtime.Func2(func(describe_0 gopurs_runtime.Value, act_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_2_0 := gopurs_runtime.Apply(describe_0, gopurs_runtime.Value{})
_ = _dollar__unused_2_0
t1_3_1 := gopurs_runtime.Apply(Get_benchNow(), gopurs_runtime.Value{})
_ = t1_3_1
_dollar__unused_4_2 := gopurs_runtime.Apply(act_1, gopurs_runtime.Value{})
_ = _dollar__unused_4_2
t2_5_3 := gopurs_runtime.Apply(Get_benchNow(), gopurs_runtime.Value{})
_ = t2_5_3
dt_6_4 := gopurs_runtime.FloatSub(t2_5_3, t1_3_1)
_ = dt_6_4
return gopurs_runtime.Apply(gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
_dollar__unused_7_5 := gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str(gopurs_runtime.Str(gopurs_runtime.Str("\nExecution time: ").StrVal + gopurs_runtime.Apply(Get_formatNumber(), dt_6_4).StrVal).StrVal + gopurs_runtime.Str(" μs\n").StrVal)), gopurs_runtime.Value{})
_ = _dollar__unused_7_5
return dt_6_4
}), gopurs_runtime.Value{})
})
})
	})
	return runBench
}

func Get_benchNow() gopurs_runtime.Value {
	return BenchNow
}

func Get_formatNumber() gopurs_runtime.Value {
	return FormatNumber
}

func Get_opaque() gopurs_runtime.Value {
	return Opaque
}

package Bench

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Effect "gopurs/output/Effect"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_discard gopurs_runtime.Value
var once_discard sync.Once
func Get_discard() gopurs_runtime.Value {
	once_discard.Do(func() {
		cache_discard = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Control_Bind.Get_discardUnit(), "discard"), pkg_Effect.Get_bindEffect())
	})
	return cache_discard
}

var cache_runBench gopurs_runtime.Value
var once_runBench sync.Once
func Get_runBench() gopurs_runtime.Value {
	once_runBench.Do(func() {
		cache_runBench = gopurs_runtime.Func2(func(describe_0_box gopurs_runtime.Value, act_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Float(Call_runBench(func() gopurs_runtime.Value {
return gopurs_runtime.Apply(describe_0_box, nil)
}, func() gopurs_runtime.Value {
return gopurs_runtime.Apply(act_1_box, nil)
})())
})
})
	})
	return cache_runBench
}

var cache_benchNow gopurs_runtime.Value
var once_benchNow sync.Once
func Get_benchNow() gopurs_runtime.Value {
	once_benchNow.Do(func() {
		cache_benchNow = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Float(BenchNow())
})
	})
	return cache_benchNow
}

var cache_formatNumber gopurs_runtime.Value
var once_formatNumber sync.Once
func Get_formatNumber() gopurs_runtime.Value {
	once_formatNumber.Do(func() {
		cache_formatNumber = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(FormatNumber(arg0.FloatVal()))
})
	})
	return cache_formatNumber
}

var cache_opaque gopurs_runtime.Value
var once_opaque sync.Once
func Get_opaque() gopurs_runtime.Value {
	once_opaque.Do(func() {
		cache_opaque = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func0(func() gopurs_runtime.Value {
return gopurs_runtime.Any(Opaque(gopurs_runtime.UnboxAny(arg0))())
})
})
	})
	return cache_opaque
}

func Call_runBench(describe_0_loop func() gopurs_runtime.Value, act_1_loop func() gopurs_runtime.Value) func() float64 {
var describe_0 func() gopurs_runtime.Value = describe_0_loop
_ = describe_0
var act_1 func() gopurs_runtime.Value = act_1_loop
_ = act_1
return func() float64 {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return describe_0()
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), Get_benchNow(), gopurs_runtime.Func(func(t1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Func0(func() gopurs_runtime.Value {
return act_1()
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), Get_benchNow(), gopurs_runtime.Func(func(t2_5 gopurs_runtime.Value) gopurs_runtime.Value {
dt_6_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Ring.Get_ringNumber(), "sub"), t2_5, t1_3)
_ = dt_6_0
return gopurs_runtime.Apply2(Get_discard(), gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("\nExecution time: "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(Get_formatNumber(), dt_6_0), gopurs_runtime.Str(" μs\n")))), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect.Get_applicativeEffect(), "pure"), dt_6_0)
}))
}))
}))
}))
})), nil).FloatVal()
}
}

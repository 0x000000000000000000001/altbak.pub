package Test_Parallelism

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Effect_Aff "gopurs/output/Effect.Aff"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Either "gopurs/output/Data.Either"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Array "gopurs/output/Data.Array"
)

var cache_traverse gopurs_runtime.Value
var once_traverse sync.Once
func Get_traverse() gopurs_runtime.Value {
	once_traverse.Do(func() {
		cache_traverse = gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "traverse"), pkg_Effect_Aff.Get_applicativeAff())
	})
	return cache_traverse
}

var cache_fib gopurs_runtime.Value
var once_fib sync.Once
func Get_fib() gopurs_runtime.Value {
	once_fib.Do(func() {
		cache_fib = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_fib(v_0_box.IntVal))
})
	})
	return cache_fib
}

var cache_heavyTask gopurs_runtime.Value
var once_heavyTask sync.Once
func Get_heavyTask() gopurs_runtime.Value {
	once_heavyTask.Do(func() {
		cache_heavyTask = gopurs_runtime.Func(func(n_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_heavyTask(n_0_box.IntVal)
})
	})
	return cache_heavyTask
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Parallelism (40 x Fib 35):"))
	})
	return cache_describe
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply(pkg_Effect_Aff.Get_void(), gopurs_runtime.Apply(pkg_Effect_Aff.Get_launchAff(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_bindAff(), "bind"), gopurs_runtime.Apply2(Get_traverse(), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Aff.Get_forkAff(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_bindAff(), "bind"), gopurs_runtime.UncurriedApp2(pkg_Effect_Aff.Get__delay(), pkg_Data_Either.Get_Right(), gopurs_runtime.Float(0.0)), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_applicativeAff(), "pure"), pkg_Data_Unit.Get_unit())
})))
}), gopurs_runtime.UncurriedApp2(pkg_Data_Array.Get_replicateImpl(), gopurs_runtime.Int(200), pkg_Data_Unit.Get_unit())), gopurs_runtime.Func(func(fibers_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_bindAff(), "bind"), gopurs_runtime.Apply2(Get_traverse(), pkg_Effect_Aff.Get_joinFiber(), fibers_0), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_applicativeAff(), "pure"), pkg_Data_Unit.Get_unit())
}))
}))))
	})
	return cache_act
}

func Call_fib(v_0_loop int64) int64 {
fib:
for {
if false { continue fib }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Int(0)
goto end_branch_0
} else {

}
}
{
if (v_0) == (1) {
__t0 = gopurs_runtime.Int(1)
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Int((gopurs_runtime.Int(Call_fib((v_0) - (1))).IntVal) + (gopurs_runtime.Int(Call_fib((v_0) - (2))).IntVal))
}
end_branch_0:
return __t0.IntVal
}
}

func Call_heavyTask(n_0_loop int64) gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_bindAff(), "bind"), gopurs_runtime.UncurriedApp2(pkg_Effect_Aff.Get__delay(), pkg_Data_Either.Get_Right(), gopurs_runtime.Float(0.0)), gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Effect_Aff.Get_applicativeAff(), "pure"), pkg_Data_Unit.Get_unit())
}))
}



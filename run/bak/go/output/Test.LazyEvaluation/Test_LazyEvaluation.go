package Test_LazyEvaluation

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var Lazy gopurs_runtime.Value
var once_Lazy sync.Once
func Get_Lazy() gopurs_runtime.Value {
	once_Lazy.Do(func() {
		Lazy = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return Lazy
}

var force gopurs_runtime.Value
var once_force sync.Once
func Get_force() gopurs_runtime.Value {
	once_force.Do(func() {
		force = gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
})
	})
	return force
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Lazy Evaluation (1M Thunks Forced, 1k Depth):"))
	})
	return describe
}

var defer_ gopurs_runtime.Value
var once_defer_ sync.Once
func Get_defer_() gopurs_runtime.Value {
	once_defer_.Do(func() {
		defer_ = gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return f_0
})
	})
	return defer_
}

var buildThunks gopurs_runtime.Value
var once_buildThunks sync.Once
func Get_buildThunks() gopurs_runtime.Value {
	once_buildThunks.Do(func() {
		buildThunks = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
buildThunks:
for {
if false { continue buildThunks }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(Get_buildThunks(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(gopurs_runtime.Apply(v1_1, pkg_Data_Unit.Get_unit()).IntVal + gopurs_runtime.Int(1).IntVal)
}))
}
end_branch_0:
return __t0
}
}()
})
})
	})
	return buildThunks
}

var runManyTimes gopurs_runtime.Value
var once_runManyTimes sync.Once
func Get_runManyTimes() gopurs_runtime.Value {
	once_runManyTimes.Do(func() {
		runManyTimes = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
runManyTimes:
for {
if false { continue runManyTimes }
var v_0 = v_0_loop
_ = v_0
var v1_1 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(Get_runManyTimes(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.Int(v1_1.IntVal + gopurs_runtime.Apply3(Get_buildThunks(), gopurs_runtime.Int(1000), gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}), pkg_Data_Unit.Get_unit()).IntVal))
}
end_branch_0:
return __t0
}
}()
})
})
	})
	return runManyTimes
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(1000))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), gopurs_runtime.Apply2(Get_runManyTimes(), dummy_1_1, gopurs_runtime.Int(0)))), gopurs_runtime.Value{})
})
}()
	})
	return act
}



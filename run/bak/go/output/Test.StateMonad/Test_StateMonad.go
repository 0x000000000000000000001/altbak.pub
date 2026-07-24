package Test_StateMonad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var State gopurs_runtime.Value
var once_State sync.Once
func Get_State() gopurs_runtime.Value {
	once_State.Do(func() {
		State = gopurs_runtime.Func(func(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0_loop
}()
})
	})
	return State
}

var runState gopurs_runtime.Value
var once_runState sync.Once
func Get_runState() gopurs_runtime.Value {
	once_runState.Do(func() {
		runState = gopurs_runtime.Func2(Call_runState)
	})
	return runState
}

var put gopurs_runtime.Value
var once_put sync.Once
func Get_put() gopurs_runtime.Value {
	once_put.Do(func() {
		put = gopurs_runtime.Func2(Call_put)
	})
	return put
}

var pureState gopurs_runtime.Value
var once_pureState sync.Once
func Get_pureState() gopurs_runtime.Value {
	once_pureState.Do(func() {
		pureState = gopurs_runtime.Func2(Call_pureState)
	})
	return pureState
}

var get gopurs_runtime.Value
var once_get sync.Once
func Get_get() gopurs_runtime.Value {
	once_get.Do(func() {
		get = gopurs_runtime.Func(func(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict2("val", "state", s_0_loop, s_0_loop)
}()
})
	})
	return get
}

var describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("State Monad (1.2k Binds, 60 Stack Depth):"))
	})
	return describe
}

var bindState gopurs_runtime.Value
var once_bindState sync.Once
func Get_bindState() gopurs_runtime.Value {
	once_bindState.Do(func() {
		bindState = gopurs_runtime.Func3(Call_bindState)
	})
	return bindState
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func2(Call_modify)
	})
	return modify
}

var chainModifications gopurs_runtime.Value
var once_chainModifications sync.Once
func Get_chainModifications() gopurs_runtime.Value {
	once_chainModifications.Do(func() {
		chainModifications = gopurs_runtime.Func(func(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
chainModifications:
for {
if false { continue chainModifications }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if v_0_loop.IntVal == 0 {
__t0 = gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("val", "state", pkg_Data_Unit.Get_unit(), s_1)
})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_chainModifications(), gopurs_runtime.Int(v_0_loop.IntVal - 1), gopurs_runtime.Int(s_1.IntVal + 1))
})
}
end_branch_0:
return __t0
}
}()
})
	})
	return chainModifications
}

var runManyTimes gopurs_runtime.Value
var once_runManyTimes sync.Once
func Get_runManyTimes() gopurs_runtime.Value {
	once_runManyTimes.Do(func() {
		runManyTimes = gopurs_runtime.Func2(Call_runManyTimes)
	})
	return runManyTimes
}

var act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		act = func() gopurs_runtime.Value {
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(20))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
dummy_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = dummy_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(pkg_Data_Show.Get_showIntImpl(), Call_runManyTimes(dummy_1_1, gopurs_runtime.Int(0)))), gopurs_runtime.Value{})
})
}()
	})
	return act
}

func Call_runState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0_loop, s_1_loop)
}

func Call_put(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.RecordDict2("val", "state", pkg_Data_Unit.Get_unit(), s_0_loop)
}

func Call_pureState(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("val", "state", a_0_loop, s_1_loop)
}

func Call_bindState(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0_loop, s_2_loop)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1_loop, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_modify(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("val", "state", pkg_Data_Unit.Get_unit(), gopurs_runtime.Apply(f_0_loop, s_1_loop))
}

func Call_runManyTimes(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
runManyTimes:
for {
if false { continue runManyTimes }
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if v_0_loop.IntVal == 0 {
__t0 = v1_1_loop
goto end_branch_0
} else {

}
}
{
__t0 = Call_runManyTimes(gopurs_runtime.Int(v_0_loop.IntVal - 1), gopurs_runtime.Int(v1_1_loop.IntVal + gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_chainModifications(), gopurs_runtime.Int(60), gopurs_runtime.Int(0)), "state").IntVal))
}
end_branch_0:
return __t0
}
}



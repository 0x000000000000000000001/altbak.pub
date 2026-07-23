package Test_StateMonad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
)

var State gopurs_runtime.Value
var once_State sync.Once
func Get_State() gopurs_runtime.Value {
	once_State.Do(func() {
		State = gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
})
	})
	return State
}

var runState gopurs_runtime.Value
var once_runState sync.Once
func Get_runState() gopurs_runtime.Value {
	once_runState.Do(func() {
		runState = gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(v_0, s_1)
})
	})
	return runState
}

var put gopurs_runtime.Value
var once_put sync.Once
func Get_put() gopurs_runtime.Value {
	once_put.Do(func() {
		put = gopurs_runtime.Func2(func(s_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("val", "state", pkg_Data_Unit.Get_unit(), s_0)
})
	})
	return put
}

var pureState gopurs_runtime.Value
var once_pureState sync.Once
func Get_pureState() gopurs_runtime.Value {
	once_pureState.Do(func() {
		pureState = gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("val", "state", a_0, s_1)
})
	})
	return pureState
}

var get gopurs_runtime.Value
var once_get sync.Once
func Get_get() gopurs_runtime.Value {
	once_get.Do(func() {
		get = gopurs_runtime.Func(func(s_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("val", "state", s_0, s_0)
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
		bindState = gopurs_runtime.Func3(func(v_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, s_2 gopurs_runtime.Value) gopurs_runtime.Value {
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
})
	})
	return bindState
}

var modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		modify = gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("val", "state", pkg_Data_Unit.Get_unit(), gopurs_runtime.Apply(f_0, s_1))
})
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
var v_0 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Bool(v_0.IntVal == gopurs_runtime.Int(0).IntVal)).IntVal != 0 {
__t0 = gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("val", "state", pkg_Data_Unit.Get_unit(), s_1)
})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_chainModifications(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.Int(s_1.IntVal + gopurs_runtime.Int(1).IntVal))
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
__t0 = gopurs_runtime.Apply2(Get_runManyTimes(), gopurs_runtime.Int(v_0.IntVal - gopurs_runtime.Int(1).IntVal), gopurs_runtime.Int(v1_1.IntVal + gopurs_runtime.RecordGet(gopurs_runtime.Apply2(Get_chainModifications(), gopurs_runtime.Int(60), gopurs_runtime.Int(0)), "state").IntVal))
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
__local_var_0_0 := gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(20))
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



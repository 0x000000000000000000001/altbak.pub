package Test_StateMonad

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var cache_State gopurs_runtime.Value
var once_State sync.Once
func Get_State() gopurs_runtime.Value {
	once_State.Do(func() {
		cache_State = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_State(x_0_box)
})
	})
	return cache_State
}

var cache_runState gopurs_runtime.Value
var once_runState sync.Once
func Get_runState() gopurs_runtime.Value {
	once_runState.Do(func() {
		cache_runState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runState(v_0_box, s_1_box)
})
	})
	return cache_runState
}

var cache_runState__gopurs_runtime_Value_1201697018 gopurs_runtime.Value
var once_runState__gopurs_runtime_Value_1201697018 sync.Once
func Get_runState__gopurs_runtime_Value_1201697018() gopurs_runtime.Value {
	once_runState__gopurs_runtime_Value_1201697018.Do(func() {
		cache_runState__gopurs_runtime_Value_1201697018 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runState__gopurs_runtime_Value_1201697018(v_0_box, s_1_box.IntVal)
})
	})
	return cache_runState__gopurs_runtime_Value_1201697018
}

var cache_runState__gopurs_runtime_Value_458711162 gopurs_runtime.Value
var once_runState__gopurs_runtime_Value_458711162 sync.Once
func Get_runState__gopurs_runtime_Value_458711162() gopurs_runtime.Value {
	once_runState__gopurs_runtime_Value_458711162.Do(func() {
		cache_runState__gopurs_runtime_Value_458711162 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runState__gopurs_runtime_Value_458711162(v_0_box, s_1_box)
})
	})
	return cache_runState__gopurs_runtime_Value_458711162
}

var cache_put gopurs_runtime.Value
var once_put sync.Once
func Get_put() gopurs_runtime.Value {
	once_put.Do(func() {
		cache_put = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_put(s_0_box, v_1_box)
})
	})
	return cache_put
}

var cache_put__gopurs_runtime_Value_1496134642 gopurs_runtime.Value
var once_put__gopurs_runtime_Value_1496134642 sync.Once
func Get_put__gopurs_runtime_Value_1496134642() gopurs_runtime.Value {
	once_put__gopurs_runtime_Value_1496134642.Do(func() {
		cache_put__gopurs_runtime_Value_1496134642 = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_put__gopurs_runtime_Value_1496134642(s_0_box, v_1_box)
})
	})
	return cache_put__gopurs_runtime_Value_1496134642
}

var cache_pureState gopurs_runtime.Value
var once_pureState sync.Once
func Get_pureState() gopurs_runtime.Value {
	once_pureState.Do(func() {
		cache_pureState = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pureState(a_0_box, s_1_box)
})
	})
	return cache_pureState
}

var cache_pureState__gopurs_runtime_Value_1575991999 gopurs_runtime.Value
var once_pureState__gopurs_runtime_Value_1575991999 sync.Once
func Get_pureState__gopurs_runtime_Value_1575991999() gopurs_runtime.Value {
	once_pureState__gopurs_runtime_Value_1575991999.Do(func() {
		cache_pureState__gopurs_runtime_Value_1575991999 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pureState__gopurs_runtime_Value_1575991999(a_0_box, s_1_box.IntVal)
})
	})
	return cache_pureState__gopurs_runtime_Value_1575991999
}

var cache_pureState__gopurs_runtime_Value_1496134642 gopurs_runtime.Value
var once_pureState__gopurs_runtime_Value_1496134642 sync.Once
func Get_pureState__gopurs_runtime_Value_1496134642() gopurs_runtime.Value {
	once_pureState__gopurs_runtime_Value_1496134642.Do(func() {
		cache_pureState__gopurs_runtime_Value_1496134642 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pureState__gopurs_runtime_Value_1496134642(a_0_box, s_1_box)
})
	})
	return cache_pureState__gopurs_runtime_Value_1496134642
}

var cache_get gopurs_runtime.Value
var once_get sync.Once
func Get_get() gopurs_runtime.Value {
	once_get.Do(func() {
		cache_get = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get(s_0_box)
})
	})
	return cache_get
}

var cache_get__gopurs_runtime_Value_2001193531 gopurs_runtime.Value
var once_get__gopurs_runtime_Value_2001193531 sync.Once
func Get_get__gopurs_runtime_Value_2001193531() gopurs_runtime.Value {
	once_get__gopurs_runtime_Value_2001193531.Do(func() {
		cache_get__gopurs_runtime_Value_2001193531 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_get__gopurs_runtime_Value_2001193531(s_0_box)
})
	})
	return cache_get__gopurs_runtime_Value_2001193531
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("State Monad (1.2k Binds, 60 Stack Depth):"))
	})
	return cache_describe
}

var cache_bindState gopurs_runtime.Value
var once_bindState sync.Once
func Get_bindState() gopurs_runtime.Value {
	once_bindState.Do(func() {
		cache_bindState = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_bindState
}

var cache_bindState__gopurs_runtime_Value_2121462815 gopurs_runtime.Value
var once_bindState__gopurs_runtime_Value_2121462815 sync.Once
func Get_bindState__gopurs_runtime_Value_2121462815() gopurs_runtime.Value {
	once_bindState__gopurs_runtime_Value_2121462815.Do(func() {
		cache_bindState__gopurs_runtime_Value_2121462815 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState__gopurs_runtime_Value_2121462815(v_0_box, g_1_box, s_2_box.IntVal)
})
	})
	return cache_bindState__gopurs_runtime_Value_2121462815
}

var cache_bindState__gopurs_runtime_Value_2297329746 gopurs_runtime.Value
var once_bindState__gopurs_runtime_Value_2297329746 sync.Once
func Get_bindState__gopurs_runtime_Value_2297329746() gopurs_runtime.Value {
	once_bindState__gopurs_runtime_Value_2297329746.Do(func() {
		cache_bindState__gopurs_runtime_Value_2297329746 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bindState__gopurs_runtime_Value_2297329746(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_bindState__gopurs_runtime_Value_2297329746
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify(f_0_box, s_1_box)
})
	})
	return cache_modify
}

var cache_modify__gopurs_runtime_Value_1411425727 gopurs_runtime.Value
var once_modify__gopurs_runtime_Value_1411425727 sync.Once
func Get_modify__gopurs_runtime_Value_1411425727() gopurs_runtime.Value {
	once_modify__gopurs_runtime_Value_1411425727.Do(func() {
		cache_modify__gopurs_runtime_Value_1411425727 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__gopurs_runtime_Value_1411425727(f_0_box, s_1_box.IntVal)
})
	})
	return cache_modify__gopurs_runtime_Value_1411425727
}

var cache_modify__gopurs_runtime_Value_458711162 gopurs_runtime.Value
var once_modify__gopurs_runtime_Value_458711162 sync.Once
func Get_modify__gopurs_runtime_Value_458711162() gopurs_runtime.Value {
	once_modify__gopurs_runtime_Value_458711162.Do(func() {
		cache_modify__gopurs_runtime_Value_458711162 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_modify__gopurs_runtime_Value_458711162(f_0_box, s_1_box)
})
	})
	return cache_modify__gopurs_runtime_Value_458711162
}

var cache_chainModifications gopurs_runtime.Value
var once_chainModifications sync.Once
func Get_chainModifications() gopurs_runtime.Value {
	once_chainModifications.Do(func() {
		cache_chainModifications = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_chainModifications(v_0_box.IntVal)
})
	})
	return cache_chainModifications
}

var cache_runManyTimes gopurs_runtime.Value
var once_runManyTimes sync.Once
func Get_runManyTimes() gopurs_runtime.Value {
	once_runManyTimes.Do(func() {
		cache_runManyTimes = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_runManyTimes(v_0_box.IntVal, v1_1_box.IntVal))
})
	})
	return cache_runManyTimes
}

var cache_act gopurs_runtime.Value
var once_act sync.Once
func Get_act() gopurs_runtime.Value {
	once_act.Do(func() {
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(20)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_runManyTimes(dummy_0.IntVal, 0))))
}))
	})
	return cache_act
}

func Call_State(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_runState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}

func Call_runState__gopurs_runtime_Value_1201697018(v_0_loop gopurs_runtime.Value, s_1_loop int64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 int64 = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, gopurs_runtime.Int(s_1))
}

func Call_runState__gopurs_runtime_Value_458711162(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}

func Call_put(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.RecordDict2("state", "val", s_0, pkg_Data_Unit.Get_unit())
}

func Call_put__gopurs_runtime_Value_1496134642(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.RecordDict2("state", "val", s_0, pkg_Data_Unit.Get_unit())
}

func Call_pureState(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", s_1, a_0)
}

func Call_pureState__gopurs_runtime_Value_1575991999(a_0_loop gopurs_runtime.Value, s_1_loop int64) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 int64 = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Int(s_1), a_0)
}

func Call_pureState__gopurs_runtime_Value_1496134642(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", s_1, a_0)
}

func Call_get(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict2("state", "val", s_0, s_0)
}

func Call_get__gopurs_runtime_Value_2001193531(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict2("state", "val", s_0, s_0)
}

func Call_bindState(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_bindState__gopurs_runtime_Value_2121462815(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop int64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 int64 = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, gopurs_runtime.Int(s_2))
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_bindState__gopurs_runtime_Value_2297329746(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_modify(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, s_1), pkg_Data_Unit.Get_unit())
}

func Call_modify__gopurs_runtime_Value_1411425727(f_0_loop gopurs_runtime.Value, s_1_loop int64) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 int64 = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, gopurs_runtime.Int(s_1)), pkg_Data_Unit.Get_unit())
}

func Call_modify__gopurs_runtime_Value_458711162(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, s_1), pkg_Data_Unit.Get_unit())
}

func Call_chainModifications(v_0_loop int64) gopurs_runtime.Value {
chainModifications:
for {
if false { continue chainModifications }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("state", "val", s_1, pkg_Data_Unit.Get_unit())
})
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Func(func(s_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_chainModifications((v_0) - (1)), gopurs_runtime.Int((s_1.IntVal) + (1)))
})
}
end_branch_0:
return __t0
}
}

func Call_runManyTimes(v_0_loop int64, v1_1_loop int64) int64 {
runManyTimes:
for {
if false { continue runManyTimes }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Int(v1_1)
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
v1_1_loop = (v1_1) + (gopurs_runtime.RecordGet(gopurs_runtime.Apply(Call_chainModifications(60), gopurs_runtime.Int(0)), "state").IntVal)
continue runManyTimes
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0.IntVal
}
}



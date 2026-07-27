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
return gopurs_runtime.Any(Call_runState(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(s_1_box)))
})
	})
	return cache_runState
}

var cache_runState__func_func_int64__interface____int64__interface___1201697018 gopurs_runtime.Value
var once_runState__func_func_int64__interface____int64__interface___1201697018 sync.Once
func Get_runState__func_func_int64__interface____int64__interface___1201697018() gopurs_runtime.Value {
	once_runState__func_func_int64__interface____int64__interface___1201697018.Do(func() {
		cache_runState__func_func_int64__interface____int64__interface___1201697018 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_runState__func_func_int64__interface____int64__interface___1201697018(func(inner_arg0 int64) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Int(inner_arg0)))
}, s_1_box.IntVal))
})
	})
	return cache_runState__func_func_int64__interface____int64__interface___1201697018
}

var cache_runState__func_func_interface____interface____interface____interface___2078380730 gopurs_runtime.Value
var once_runState__func_func_interface____interface____interface____interface___2078380730 sync.Once
func Get_runState__func_func_interface____interface____interface____interface___2078380730() gopurs_runtime.Value {
	once_runState__func_func_interface____interface____interface____interface___2078380730.Do(func() {
		cache_runState__func_func_interface____interface____interface____interface___2078380730 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_runState__func_func_interface____interface____interface____interface___2078380730(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(s_1_box)))
})
	})
	return cache_runState__func_func_interface____interface____interface____interface___2078380730
}

var cache_put gopurs_runtime.Value
var once_put sync.Once
func Get_put() gopurs_runtime.Value {
	once_put.Do(func() {
		cache_put = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_put(gopurs_runtime.UnboxAny(s_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_put
}

var cache_put__func_interface____interface____interface___1617116210 gopurs_runtime.Value
var once_put__func_interface____interface____interface___1617116210 sync.Once
func Get_put__func_interface____interface____interface___1617116210() gopurs_runtime.Value {
	once_put__func_interface____interface____interface___1617116210.Do(func() {
		cache_put__func_interface____interface____interface___1617116210 = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_put__func_interface____interface____interface___1617116210(gopurs_runtime.UnboxAny(s_0_box), gopurs_runtime.UnboxAny(v_1_box)))
})
	})
	return cache_put__func_interface____interface____interface___1617116210
}

var cache_pureState gopurs_runtime.Value
var once_pureState sync.Once
func Get_pureState() gopurs_runtime.Value {
	once_pureState.Do(func() {
		cache_pureState = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_pureState(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(s_1_box)))
})
	})
	return cache_pureState
}

var cache_pureState__func_gopurs_runtime_Value__interface____interface___1513847429 gopurs_runtime.Value
var once_pureState__func_gopurs_runtime_Value__interface____interface___1513847429 sync.Once
func Get_pureState__func_gopurs_runtime_Value__interface____interface___1513847429() gopurs_runtime.Value {
	once_pureState__func_gopurs_runtime_Value__interface____interface___1513847429.Do(func() {
		cache_pureState__func_gopurs_runtime_Value__interface____interface___1513847429 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_pureState__func_gopurs_runtime_Value__interface____interface___1513847429(a_0_box, gopurs_runtime.UnboxAny(s_1_box)))
})
	})
	return cache_pureState__func_gopurs_runtime_Value__interface____interface___1513847429
}

var cache_pureState__func_interface____interface____interface___3883848288 gopurs_runtime.Value
var once_pureState__func_interface____interface____interface___3883848288 sync.Once
func Get_pureState__func_interface____interface____interface___3883848288() gopurs_runtime.Value {
	once_pureState__func_interface____interface____interface___3883848288.Do(func() {
		cache_pureState__func_interface____interface____interface___3883848288 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_pureState__func_interface____interface____interface___3883848288(gopurs_runtime.UnboxAny(a_0_box), gopurs_runtime.UnboxAny(s_1_box)))
})
	})
	return cache_pureState__func_interface____interface____interface___3883848288
}

var cache_get gopurs_runtime.Value
var once_get sync.Once
func Get_get() gopurs_runtime.Value {
	once_get.Do(func() {
		cache_get = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get(gopurs_runtime.UnboxAny(s_0_box)))
})
	})
	return cache_get
}

var cache_get__func_interface____interface___22483300 gopurs_runtime.Value
var once_get__func_interface____interface___22483300 sync.Once
func Get_get__func_interface____interface___22483300() gopurs_runtime.Value {
	once_get__func_interface____interface___22483300.Do(func() {
		cache_get__func_interface____interface___22483300 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_get__func_interface____interface___22483300(gopurs_runtime.UnboxAny(s_0_box)))
})
	})
	return cache_get__func_interface____interface___22483300
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("State Monad (1.2k Binds, 60 Stack Depth):")), nil)
}()
})
	})
	return cache_describe
}

var cache_bindState gopurs_runtime.Value
var once_bindState sync.Once
func Get_bindState() gopurs_runtime.Value {
	once_bindState.Do(func() {
		cache_bindState = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindState(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(g_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(s_2_box)))
})
	})
	return cache_bindState
}

var cache_bindState__func_func_int64__interface____func_gopurs_runtime_Value__int64__interface____int64__interface___2121462815 gopurs_runtime.Value
var once_bindState__func_func_int64__interface____func_gopurs_runtime_Value__int64__interface____int64__interface___2121462815 sync.Once
func Get_bindState__func_func_int64__interface____func_gopurs_runtime_Value__int64__interface____int64__interface___2121462815() gopurs_runtime.Value {
	once_bindState__func_func_int64__interface____func_gopurs_runtime_Value__int64__interface____int64__interface___2121462815.Do(func() {
		cache_bindState__func_func_int64__interface____func_gopurs_runtime_Value__int64__interface____int64__interface___2121462815 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindState__func_func_int64__interface____func_gopurs_runtime_Value__int64__interface____int64__interface___2121462815(func(inner_arg0 int64) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Int(inner_arg0)))
}, func(inner_arg0 gopurs_runtime.Value, inner_arg1 int64) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(g_1_box, inner_arg0, gopurs_runtime.Int(inner_arg1)))
}, s_2_box.IntVal))
})
	})
	return cache_bindState__func_func_int64__interface____func_gopurs_runtime_Value__int64__interface____int64__interface___2121462815
}

var cache_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3457345440 gopurs_runtime.Value
var once_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3457345440 sync.Once
func Get_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3457345440() gopurs_runtime.Value {
	once_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3457345440.Do(func() {
		cache_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3457345440 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3457345440(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(g_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(s_2_box)))
})
	})
	return cache_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3457345440
}

var cache_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3258626290 gopurs_runtime.Value
var once_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3258626290 sync.Once
func Get_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3258626290() gopurs_runtime.Value {
	once_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3258626290.Do(func() {
		cache_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3258626290 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3258626290(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(g_1_box, gopurs_runtime.Any(inner_arg0), gopurs_runtime.Any(inner_arg1)))
}, gopurs_runtime.UnboxAny(s_2_box)))
})
	})
	return cache_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3258626290
}

var cache_modify gopurs_runtime.Value
var once_modify sync.Once
func Get_modify() gopurs_runtime.Value {
	once_modify.Do(func() {
		cache_modify = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_modify(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(s_1_box)))
})
	})
	return cache_modify
}

var cache_modify__func_func_int64__int64__int64__interface___1411425727 gopurs_runtime.Value
var once_modify__func_func_int64__int64__int64__interface___1411425727 sync.Once
func Get_modify__func_func_int64__int64__int64__interface___1411425727() gopurs_runtime.Value {
	once_modify__func_func_int64__int64__int64__interface___1411425727.Do(func() {
		cache_modify__func_func_int64__int64__int64__interface___1411425727 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_modify__func_func_int64__int64__int64__interface___1411425727(func(inner_arg0 int64) int64 {
return gopurs_runtime.Apply(f_0_box, gopurs_runtime.Int(inner_arg0)).IntVal
}, s_1_box.IntVal))
})
	})
	return cache_modify__func_func_int64__int64__int64__interface___1411425727
}

var cache_modify__func_func_interface____interface____interface____interface___3630542661 gopurs_runtime.Value
var once_modify__func_func_interface____interface____interface____interface___3630542661 sync.Once
func Get_modify__func_func_interface____interface____interface____interface___3630542661() gopurs_runtime.Value {
	once_modify__func_func_interface____interface____interface____interface___3630542661.Do(func() {
		cache_modify__func_func_interface____interface____interface____interface___3630542661 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_modify__func_func_interface____interface____interface____interface___3630542661(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(s_1_box)))
})
	})
	return cache_modify__func_func_interface____interface____interface____interface___3630542661
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
		cache_act = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(20)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_runManyTimes(dummy_0.IntVal, 0))))
})), nil)
}()
})
	})
	return cache_act
}

func Call_State(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_runState(v_0_loop func(interface{}) interface{}, s_1_loop interface{}) interface{} {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
var s_1 interface{} = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(s_1)))
}

func Call_runState__func_func_int64__interface____int64__interface___1201697018(v_0_loop func(int64) interface{}, s_1_loop int64) interface{} {
var v_0 func(int64) interface{} = v_0_loop
_ = v_0
var s_1 int64 = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(s_1)))
}

func Call_runState__func_func_interface____interface____interface____interface___2078380730(v_0_loop func(interface{}) interface{}, s_1_loop interface{}) interface{} {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
var s_1 interface{} = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(s_1)))
}

func Call_put(s_0_loop interface{}, v_1_loop interface{}) interface{} {
var s_0 interface{} = s_0_loop
_ = s_0
var v_1 interface{} = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Any(s_0), pkg_Data_Unit.Get_unit()))
}

func Call_put__func_interface____interface____interface___1617116210(s_0_loop interface{}, v_1_loop interface{}) interface{} {
var s_0 interface{} = s_0_loop
_ = s_0
var v_1 interface{} = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Any(s_0), pkg_Data_Unit.Get_unit()))
}

func Call_pureState(a_0_loop interface{}, s_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var s_1 interface{} = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Any(s_1), gopurs_runtime.Any(a_0)))
}

func Call_pureState__func_gopurs_runtime_Value__interface____interface___1513847429(a_0_loop gopurs_runtime.Value, s_1_loop interface{}) interface{} {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 interface{} = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Any(s_1), a_0))
}

func Call_pureState__func_interface____interface____interface___3883848288(a_0_loop interface{}, s_1_loop interface{}) interface{} {
var a_0 interface{} = a_0_loop
_ = a_0
var s_1 interface{} = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Any(s_1), gopurs_runtime.Any(a_0)))
}

func Call_get(s_0_loop interface{}) interface{} {
var s_0 interface{} = s_0_loop
_ = s_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Any(s_0), gopurs_runtime.Any(s_0)))
}

func Call_get__func_interface____interface___22483300(s_0_loop interface{}) interface{} {
var s_0 interface{} = s_0_loop
_ = s_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Any(s_0), gopurs_runtime.Any(s_0)))
}

func Call_bindState(v_0_loop func(interface{}) interface{}, g_1_loop func(interface{}, interface{}) interface{}, s_2_loop interface{}) interface{} {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
var g_1 func(interface{}, interface{}) interface{} = g_1_loop
_ = g_1
var s_2 interface{} = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Any(v_0(s_2))
_ = r1_3_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(r1_3_0, "val")), gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(r1_3_0, "state")))))
}

func Call_bindState__func_func_int64__interface____func_gopurs_runtime_Value__int64__interface____int64__interface___2121462815(v_0_loop func(int64) interface{}, g_1_loop func(gopurs_runtime.Value, int64) interface{}, s_2_loop int64) interface{} {
var v_0 func(int64) interface{} = v_0_loop
_ = v_0
var g_1 func(gopurs_runtime.Value, int64) interface{} = g_1_loop
_ = g_1
var s_2 int64 = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Any(v_0(s_2))
_ = r1_3_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state").IntVal)))
}

func Call_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3457345440(v_0_loop func(interface{}) interface{}, g_1_loop func(interface{}, interface{}) interface{}, s_2_loop interface{}) interface{} {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
var g_1 func(interface{}, interface{}) interface{} = g_1_loop
_ = g_1
var s_2 interface{} = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Any(v_0(s_2))
_ = r1_3_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(r1_3_0, "val")), gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(r1_3_0, "state")))))
}

func Call_bindState__func_func_interface____interface____func_interface____interface____interface____interface____interface___3258626290(v_0_loop func(interface{}) interface{}, g_1_loop func(interface{}, interface{}) interface{}, s_2_loop interface{}) interface{} {
var v_0 func(interface{}) interface{} = v_0_loop
_ = v_0
var g_1 func(interface{}, interface{}) interface{} = g_1_loop
_ = g_1
var s_2 interface{} = s_2_loop
_ = s_2
r1_3_0 := gopurs_runtime.Any(v_0(s_2))
_ = r1_3_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(g_1(gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(r1_3_0, "val")), gopurs_runtime.UnboxAny(gopurs_runtime.RecordGet(r1_3_0, "state")))))
}

func Call_modify(f_0_loop func(interface{}) interface{}, s_1_loop interface{}) interface{} {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var s_1 interface{} = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Any(f_0(s_1)), pkg_Data_Unit.Get_unit()))
}

func Call_modify__func_func_int64__int64__int64__interface___1411425727(f_0_loop func(int64) int64, s_1_loop int64) interface{} {
var f_0 func(int64) int64 = f_0_loop
_ = f_0
var s_1 int64 = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Int(f_0(s_1)), pkg_Data_Unit.Get_unit()))
}

func Call_modify__func_func_interface____interface____interface____interface___3630542661(f_0_loop func(interface{}) interface{}, s_1_loop interface{}) interface{} {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var s_1 interface{} = s_1_loop
_ = s_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Any(f_0(s_1)), pkg_Data_Unit.Get_unit()))
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

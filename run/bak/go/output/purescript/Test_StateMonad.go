package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Test_StateMonad_logShow gopurs_runtime.Value
var once_Test_StateMonad_logShow sync.Once
func Get_Test_StateMonad_logShow() gopurs_runtime.Value {
	once_Test_StateMonad_logShow.Do(func() {
		cache_Test_StateMonad_logShow = gopurs_runtime.Apply(Get_Effect_Console_logShow(), Get_Data_Show_showInt())
	})
	return cache_Test_StateMonad_logShow
}

var cache_Test_StateMonad_State gopurs_runtime.Value
var once_Test_StateMonad_State sync.Once
func Get_Test_StateMonad_State() gopurs_runtime.Value {
	once_Test_StateMonad_State.Do(func() {
		cache_Test_StateMonad_State = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_State(x_0_box)
})
	})
	return cache_Test_StateMonad_State
}

var cache_Test_StateMonad_runState gopurs_runtime.Value
var once_Test_StateMonad_runState sync.Once
func Get_Test_StateMonad_runState() gopurs_runtime.Value {
	once_Test_StateMonad_runState.Do(func() {
		cache_Test_StateMonad_runState = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_runState(v_0_box, s_1_box)
})
	})
	return cache_Test_StateMonad_runState
}

var cache_Test_StateMonad_put gopurs_runtime.Value
var once_Test_StateMonad_put sync.Once
func Get_Test_StateMonad_put() gopurs_runtime.Value {
	once_Test_StateMonad_put.Do(func() {
		cache_Test_StateMonad_put = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_put(s_0_box, v_1_box)
})
	})
	return cache_Test_StateMonad_put
}

var cache_Test_StateMonad_pureState gopurs_runtime.Value
var once_Test_StateMonad_pureState sync.Once
func Get_Test_StateMonad_pureState() gopurs_runtime.Value {
	once_Test_StateMonad_pureState.Do(func() {
		cache_Test_StateMonad_pureState = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_pureState(a_0_box, s_1_box)
})
	})
	return cache_Test_StateMonad_pureState
}

var cache_Test_StateMonad_get gopurs_runtime.Value
var once_Test_StateMonad_get sync.Once
func Get_Test_StateMonad_get() gopurs_runtime.Value {
	once_Test_StateMonad_get.Do(func() {
		cache_Test_StateMonad_get = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_get(s_0_box)
})
	})
	return cache_Test_StateMonad_get
}

var cache_Test_StateMonad_describe gopurs_runtime.Value
var once_Test_StateMonad_describe sync.Once
func Get_Test_StateMonad_describe() gopurs_runtime.Value {
	once_Test_StateMonad_describe.Do(func() {
		cache_Test_StateMonad_describe = gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str("State Monad (1.2k Binds, 60 Stack Depth):"))
	})
	return cache_Test_StateMonad_describe
}

var cache_Test_StateMonad_bindState gopurs_runtime.Value
var once_Test_StateMonad_bindState sync.Once
func Get_Test_StateMonad_bindState() gopurs_runtime.Value {
	once_Test_StateMonad_bindState.Do(func() {
		cache_Test_StateMonad_bindState = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_bindState(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_Test_StateMonad_bindState
}

var cache_Test_StateMonad_modify gopurs_runtime.Value
var once_Test_StateMonad_modify sync.Once
func Get_Test_StateMonad_modify() gopurs_runtime.Value {
	once_Test_StateMonad_modify.Do(func() {
		cache_Test_StateMonad_modify = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_modify(f_0_box, s_1_box)
})
	})
	return cache_Test_StateMonad_modify
}

var cache_Test_StateMonad_chainModifications gopurs_runtime.Value
var once_Test_StateMonad_chainModifications sync.Once
func Get_Test_StateMonad_chainModifications() gopurs_runtime.Value {
	once_Test_StateMonad_chainModifications.Do(func() {
		cache_Test_StateMonad_chainModifications = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_chainModifications(v_0_box.IntVal)
})
	})
	return cache_Test_StateMonad_chainModifications
}

var cache_Test_StateMonad_runManyTimes gopurs_runtime.Value
var once_Test_StateMonad_runManyTimes sync.Once
func Get_Test_StateMonad_runManyTimes() gopurs_runtime.Value {
	once_Test_StateMonad_runManyTimes.Do(func() {
		cache_Test_StateMonad_runManyTimes = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Test_StateMonad_runManyTimes(v_0_box.IntVal, v1_1_box.IntVal))
})
	})
	return cache_Test_StateMonad_runManyTimes
}

var cache_Test_StateMonad_act gopurs_runtime.Value
var once_Test_StateMonad_act sync.Once
func Get_Test_StateMonad_act() gopurs_runtime.Value {
	once_Test_StateMonad_act.Do(func() {
		cache_Test_StateMonad_act = func() gopurs_runtime.Value {
// TAST (Let): __local_var_0_0 -> gopurs_runtime.Value
__local_var_0_0 := gopurs_runtime.Apply(Get_Bench_opaque(), gopurs_runtime.Int(20))
_ = __local_var_0_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_1 := gopurs_runtime.Apply(__local_var_0_0, gopurs_runtime.Value{})
_ = __local_var_1_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Effect_Console_log(), gopurs_runtime.Str(gopurs_runtime.Apply(Get_Data_Show_showIntImpl(), gopurs_runtime.Int(Call_Test_StateMonad_runManyTimes(__local_var_1_1.IntVal, 0))).StrVal())), gopurs_runtime.Value{})
})
}()
	})
	return cache_Test_StateMonad_act
}

var cache_Test_StateMonad_bindState__3889441427 gopurs_runtime.Value
var once_Test_StateMonad_bindState__3889441427 sync.Once
func Get_Test_StateMonad_bindState__3889441427() gopurs_runtime.Value {
	once_Test_StateMonad_bindState__3889441427.Do(func() {
		cache_Test_StateMonad_bindState__3889441427 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_bindState__3889441427(v_0_box, g_1_box, s_2_box.IntVal)
})
	})
	return cache_Test_StateMonad_bindState__3889441427
}

var cache_Test_StateMonad_bindState__567439955 gopurs_runtime.Value
var once_Test_StateMonad_bindState__567439955 sync.Once
func Get_Test_StateMonad_bindState__567439955() gopurs_runtime.Value {
	once_Test_StateMonad_bindState__567439955.Do(func() {
		cache_Test_StateMonad_bindState__567439955 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_bindState__567439955(v_0_box, g_1_box, s_2_box.IntVal)
})
	})
	return cache_Test_StateMonad_bindState__567439955
}

var cache_Test_StateMonad_bindState__1042354259 gopurs_runtime.Value
var once_Test_StateMonad_bindState__1042354259 sync.Once
func Get_Test_StateMonad_bindState__1042354259() gopurs_runtime.Value {
	once_Test_StateMonad_bindState__1042354259.Do(func() {
		cache_Test_StateMonad_bindState__1042354259 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_bindState__1042354259(v_0_box, g_1_box, s_2_box.IntVal)
})
	})
	return cache_Test_StateMonad_bindState__1042354259
}

var cache_Test_StateMonad_bindState__2171045075 gopurs_runtime.Value
var once_Test_StateMonad_bindState__2171045075 sync.Once
func Get_Test_StateMonad_bindState__2171045075() gopurs_runtime.Value {
	once_Test_StateMonad_bindState__2171045075.Do(func() {
		cache_Test_StateMonad_bindState__2171045075 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_bindState__2171045075(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_Test_StateMonad_bindState__2171045075
}

var cache_Test_StateMonad_bindState__3267751411 gopurs_runtime.Value
var once_Test_StateMonad_bindState__3267751411 sync.Once
func Get_Test_StateMonad_bindState__3267751411() gopurs_runtime.Value {
	once_Test_StateMonad_bindState__3267751411.Do(func() {
		cache_Test_StateMonad_bindState__3267751411 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, s_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_bindState__3267751411(v_0_box, g_1_box, s_2_box)
})
	})
	return cache_Test_StateMonad_bindState__3267751411
}

var cache_Test_StateMonad_get__314728309 gopurs_runtime.Value
var once_Test_StateMonad_get__314728309 sync.Once
func Get_Test_StateMonad_get__314728309() gopurs_runtime.Value {
	once_Test_StateMonad_get__314728309.Do(func() {
		cache_Test_StateMonad_get__314728309 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_get__314728309(s_0_box.IntVal)
})
	})
	return cache_Test_StateMonad_get__314728309
}

var cache_Test_StateMonad_get__676984528 gopurs_runtime.Value
var once_Test_StateMonad_get__676984528 sync.Once
func Get_Test_StateMonad_get__676984528() gopurs_runtime.Value {
	once_Test_StateMonad_get__676984528.Do(func() {
		cache_Test_StateMonad_get__676984528 = gopurs_runtime.Func(func(s_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_get__676984528(s_0_box)
})
	})
	return cache_Test_StateMonad_get__676984528
}

var cache_Test_StateMonad_modify__1175978184 gopurs_runtime.Value
var once_Test_StateMonad_modify__1175978184 sync.Once
func Get_Test_StateMonad_modify__1175978184() gopurs_runtime.Value {
	once_Test_StateMonad_modify__1175978184.Do(func() {
		cache_Test_StateMonad_modify__1175978184 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_modify__1175978184(f_0_box, s_1_box.IntVal)
})
	})
	return cache_Test_StateMonad_modify__1175978184
}

var cache_Test_StateMonad_modify__3050914184 gopurs_runtime.Value
var once_Test_StateMonad_modify__3050914184 sync.Once
func Get_Test_StateMonad_modify__3050914184() gopurs_runtime.Value {
	once_Test_StateMonad_modify__3050914184.Do(func() {
		cache_Test_StateMonad_modify__3050914184 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_modify__3050914184(f_0_box, s_1_box)
})
	})
	return cache_Test_StateMonad_modify__3050914184
}

var cache_Test_StateMonad_pureState__608762702 gopurs_runtime.Value
var once_Test_StateMonad_pureState__608762702 sync.Once
func Get_Test_StateMonad_pureState__608762702() gopurs_runtime.Value {
	once_Test_StateMonad_pureState__608762702.Do(func() {
		cache_Test_StateMonad_pureState__608762702 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_pureState__608762702(a_0_box, s_1_box.IntVal)
})
	})
	return cache_Test_StateMonad_pureState__608762702
}

var cache_Test_StateMonad_pureState__1329830318 gopurs_runtime.Value
var once_Test_StateMonad_pureState__1329830318 sync.Once
func Get_Test_StateMonad_pureState__1329830318() gopurs_runtime.Value {
	once_Test_StateMonad_pureState__1329830318.Do(func() {
		cache_Test_StateMonad_pureState__1329830318 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_pureState__1329830318(a_0_box, s_1_box)
})
	})
	return cache_Test_StateMonad_pureState__1329830318
}

var cache_Test_StateMonad_put__1769569765 gopurs_runtime.Value
var once_Test_StateMonad_put__1769569765 sync.Once
func Get_Test_StateMonad_put__1769569765() gopurs_runtime.Value {
	once_Test_StateMonad_put__1769569765.Do(func() {
		cache_Test_StateMonad_put__1769569765 = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_put__1769569765(s_0_box.IntVal, v_1_box.IntVal)
})
	})
	return cache_Test_StateMonad_put__1769569765
}

var cache_Test_StateMonad_put__3685210848 gopurs_runtime.Value
var once_Test_StateMonad_put__3685210848 sync.Once
func Get_Test_StateMonad_put__3685210848() gopurs_runtime.Value {
	once_Test_StateMonad_put__3685210848.Do(func() {
		cache_Test_StateMonad_put__3685210848 = gopurs_runtime.Func2(func(s_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_put__3685210848(s_0_box, v_1_box)
})
	})
	return cache_Test_StateMonad_put__3685210848
}

var cache_Test_StateMonad_runState__2373419117 gopurs_runtime.Value
var once_Test_StateMonad_runState__2373419117 sync.Once
func Get_Test_StateMonad_runState__2373419117() gopurs_runtime.Value {
	once_Test_StateMonad_runState__2373419117.Do(func() {
		cache_Test_StateMonad_runState__2373419117 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_runState__2373419117(v_0_box, s_1_box)
})
	})
	return cache_Test_StateMonad_runState__2373419117
}

var cache_Test_StateMonad_runState__3059282509 gopurs_runtime.Value
var once_Test_StateMonad_runState__3059282509 sync.Once
func Get_Test_StateMonad_runState__3059282509() gopurs_runtime.Value {
	once_Test_StateMonad_runState__3059282509.Do(func() {
		cache_Test_StateMonad_runState__3059282509 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, s_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_runState__3059282509(v_0_box, s_1_box)
})
	})
	return cache_Test_StateMonad_runState__3059282509
}

func Call_Test_StateMonad_State(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Test_StateMonad_runState(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}

func Call_Test_StateMonad_put(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.RecordDict2("state", "val", s_0, Get_Data_Unit_unit())
}

func Call_Test_StateMonad_pureState(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", s_1, a_0)
}

func Call_Test_StateMonad_get(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict2("state", "val", s_0, s_0)
}

func Call_Test_StateMonad_bindState(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
// TAST (Let): r1_3_0 -> gopurs_runtime.Value
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_Test_StateMonad_modify(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, s_1), Get_Data_Unit_unit())
}

func Call_Test_StateMonad_chainModifications(v_0_loop int64) gopurs_runtime.Value {
chainModifications:
for {
if false { continue chainModifications }
var v_0 int64 = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Apply(Get_Test_StateMonad_pureState__608762702(), Get_Data_Unit_unit())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Apply2(Get_Test_StateMonad_bindState__567439955(), gopurs_runtime.Apply(Get_Test_StateMonad_modify__1175978184(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((x_1.IntVal) + (1))
})), gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Test_StateMonad_chainModifications((v_0) - (1))
}))
}
end_branch_0:
return __t0
}
}

func Call_Test_StateMonad_runManyTimes(v_0_loop int64, v1_1_loop int64) int64 {
runManyTimes:
for {
if false { continue runManyTimes }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 int64 = v1_1_loop
_ = v1_1
var __t0 int64
{
if (v_0) == (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
v1_1_loop = (v1_1) + (gopurs_runtime.RecordGet(Call_Test_StateMonad_runState__2373419117(Call_Test_StateMonad_chainModifications(60), gopurs_runtime.Int(0)), "state").IntVal)
continue runManyTimes
__t0 = gopurs_runtime.Value{}.IntVal
}
end_branch_0:
return __t0
}
}

func Call_Test_StateMonad_bindState__3889441427(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop int64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 int64 = s_2_loop
_ = s_2
// TAST (Let): r1_3_0 -> gopurs_runtime.Value
r1_3_0 := gopurs_runtime.Apply(v_0, gopurs_runtime.Int(s_2))
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.Int(gopurs_runtime.RecordGet(r1_3_0, "val").IntVal), gopurs_runtime.Int(gopurs_runtime.RecordGet(r1_3_0, "state").IntVal))
}

func Call_Test_StateMonad_bindState__567439955(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop int64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 int64 = s_2_loop
_ = s_2
// TAST (Let): r1_3_0 -> gopurs_runtime.Value
r1_3_0 := gopurs_runtime.Apply(v_0, gopurs_runtime.Int(s_2))
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.Int(gopurs_runtime.RecordGet(r1_3_0, "state").IntVal))
}

func Call_Test_StateMonad_bindState__1042354259(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop int64) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 int64 = s_2_loop
_ = s_2
// TAST (Let): r1_3_0 -> gopurs_runtime.Value
r1_3_0 := gopurs_runtime.Apply(v_0, gopurs_runtime.Int(s_2))
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.Int(gopurs_runtime.RecordGet(r1_3_0, "state").IntVal))
}

func Call_Test_StateMonad_bindState__2171045075(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
// TAST (Let): r1_3_0 -> gopurs_runtime.Value
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_Test_StateMonad_bindState__3267751411(v_0_loop gopurs_runtime.Value, g_1_loop gopurs_runtime.Value, s_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var g_1 gopurs_runtime.Value = g_1_loop
_ = g_1
var s_2 gopurs_runtime.Value = s_2_loop
_ = s_2
// TAST (Let): r1_3_0 -> gopurs_runtime.Value
r1_3_0 := gopurs_runtime.Apply(v_0, s_2)
_ = r1_3_0
return gopurs_runtime.Apply2(g_1, gopurs_runtime.RecordGet(r1_3_0, "val"), gopurs_runtime.RecordGet(r1_3_0, "state"))
}

func Call_Test_StateMonad_get__314728309(s_0_loop int64) gopurs_runtime.Value {
var s_0 int64 = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Int(s_0), gopurs_runtime.Int(s_0))
}

func Call_Test_StateMonad_get__676984528(s_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
return gopurs_runtime.RecordDict2("state", "val", s_0, s_0)
}

func Call_Test_StateMonad_modify__1175978184(f_0_loop gopurs_runtime.Value, s_1_loop int64) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 int64 = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Int(gopurs_runtime.Apply(f_0, gopurs_runtime.Int(gopurs_runtime.Int(s_1).IntVal)).IntVal), Get_Data_Unit_unit())
}

func Call_Test_StateMonad_modify__3050914184(f_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Apply(f_0, s_1), Get_Data_Unit_unit())
}

func Call_Test_StateMonad_pureState__608762702(a_0_loop gopurs_runtime.Value, s_1_loop int64) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 int64 = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Int(s_1), a_0)
}

func Call_Test_StateMonad_pureState__1329830318(a_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.RecordDict2("state", "val", s_1, a_0)
}

func Call_Test_StateMonad_put__1769569765(s_0_loop int64, v_1_loop int64) gopurs_runtime.Value {
var s_0 int64 = s_0_loop
_ = s_0
var v_1 int64 = v_1_loop
_ = v_1
return gopurs_runtime.RecordDict2("state", "val", gopurs_runtime.Int(s_0), Get_Data_Unit_unit())
}

func Call_Test_StateMonad_put__3685210848(s_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var s_0 gopurs_runtime.Value = s_0_loop
_ = s_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.RecordDict2("state", "val", s_0, Get_Data_Unit_unit())
}

func Call_Test_StateMonad_runState__2373419117(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}

func Call_Test_StateMonad_runState__3059282509(v_0_loop gopurs_runtime.Value, s_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var s_1 gopurs_runtime.Value = s_1_loop
_ = s_1
return gopurs_runtime.Apply(v_0, s_1)
}



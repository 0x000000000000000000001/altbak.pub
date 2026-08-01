package Test_LazyEvaluation

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Effect_Console "gopurs/output/Effect.Console"
	pkg_Effect "gopurs/output/Effect"
	pkg_Bench "gopurs/output/Bench"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var cache_Lazy gopurs_runtime.Value
var once_Lazy sync.Once
func Get_Lazy() gopurs_runtime.Value {
	once_Lazy.Do(func() {
		cache_Lazy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Lazy(x_0_box)
})
	})
	return cache_Lazy
}

var cache_force gopurs_runtime.Value
var once_force sync.Once
func Get_force() gopurs_runtime.Value {
	once_force.Do(func() {
		cache_force = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_force(v_0_box)
})
	})
	return cache_force
}

var cache_force__gopurs_runtime_Value_2469426587 gopurs_runtime.Value
var once_force__gopurs_runtime_Value_2469426587 sync.Once
func Get_force__gopurs_runtime_Value_2469426587() gopurs_runtime.Value {
	once_force__gopurs_runtime_Value_2469426587.Do(func() {
		cache_force__gopurs_runtime_Value_2469426587 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_force__gopurs_runtime_Value_2469426587(v_0_box))
})
	})
	return cache_force__gopurs_runtime_Value_2469426587
}

var cache_force__gopurs_runtime_Value_1612086811 gopurs_runtime.Value
var once_force__gopurs_runtime_Value_1612086811 sync.Once
func Get_force__gopurs_runtime_Value_1612086811() gopurs_runtime.Value {
	once_force__gopurs_runtime_Value_1612086811.Do(func() {
		cache_force__gopurs_runtime_Value_1612086811 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_force__gopurs_runtime_Value_1612086811(v_0_box)
})
	})
	return cache_force__gopurs_runtime_Value_1612086811
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Lazy Evaluation (1M Thunks Forced, 1k Depth):"))
	})
	return cache_describe
}

var cache_go__defer gopurs_runtime.Value
var once_go__defer sync.Once
func Get_go__defer() gopurs_runtime.Value {
	once_go__defer.Do(func() {
		cache_go__defer = Get_Lazy()
	})
	return cache_go__defer
}

var cache_defer__gopurs_runtime_Value_487812186 gopurs_runtime.Value
var once_defer__gopurs_runtime_Value_487812186 sync.Once
func Get_defer__gopurs_runtime_Value_487812186() gopurs_runtime.Value {
	once_defer__gopurs_runtime_Value_487812186.Do(func() {
		cache_defer__gopurs_runtime_Value_487812186 = Get_Lazy()
	})
	return cache_defer__gopurs_runtime_Value_487812186
}

var cache_defer__gopurs_runtime_Value_3386315898 gopurs_runtime.Value
var once_defer__gopurs_runtime_Value_3386315898 sync.Once
func Get_defer__gopurs_runtime_Value_3386315898() gopurs_runtime.Value {
	once_defer__gopurs_runtime_Value_3386315898.Do(func() {
		cache_defer__gopurs_runtime_Value_3386315898 = Get_Lazy()
	})
	return cache_defer__gopurs_runtime_Value_3386315898
}

var cache_buildThunks gopurs_runtime.Value
var once_buildThunks sync.Once
func Get_buildThunks() gopurs_runtime.Value {
	once_buildThunks.Do(func() {
		cache_buildThunks = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_buildThunks(v_0_box.IntVal, v1_1_box)
})
	})
	return cache_buildThunks
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
		cache_act = gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(1000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_runManyTimes(dummy_0.IntVal, 0))))
}))
	})
	return cache_act
}

func Call_Lazy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_force(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}

func Call_force__gopurs_runtime_Value_2469426587(v_0_loop gopurs_runtime.Value) int64 {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit()).IntVal
}

func Call_force__gopurs_runtime_Value_1612086811(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return gopurs_runtime.Apply(v_0, pkg_Data_Unit.Get_unit())
}

func Call_buildThunks(v_0_loop int64, v1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
buildThunks:
for {
if false { continue buildThunks }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = v1_1
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
v1_1_loop = gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((gopurs_runtime.Apply(v1_1, pkg_Data_Unit.Get_unit()).IntVal) + (1))
})
continue buildThunks
__t0 = gopurs_runtime.Value{}
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
v1_1_loop = (v1_1) + (gopurs_runtime.Apply(Call_buildThunks(1000, gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
})), pkg_Data_Unit.Get_unit()).IntVal)
continue runManyTimes
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0.IntVal
}
}



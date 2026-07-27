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
return gopurs_runtime.Any(Call_force(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, inner_arg0))
}))
})
	})
	return cache_force
}

var cache_force__func_func_gopurs_runtime_Value__int64__int64_2469426587 gopurs_runtime.Value
var once_force__func_func_gopurs_runtime_Value__int64__int64_2469426587 sync.Once
func Get_force__func_func_gopurs_runtime_Value__int64__int64_2469426587() gopurs_runtime.Value {
	once_force__func_func_gopurs_runtime_Value__int64__int64_2469426587.Do(func() {
		cache_force__func_func_gopurs_runtime_Value__int64__int64_2469426587 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_force__func_func_gopurs_runtime_Value__int64__int64_2469426587(func(inner_arg0 gopurs_runtime.Value) int64 {
return gopurs_runtime.Apply(v_0_box, inner_arg0).IntVal
}))
})
	})
	return cache_force__func_func_gopurs_runtime_Value__int64__int64_2469426587
}

var cache_force__func_func_gopurs_runtime_Value__interface____interface___2139985499 gopurs_runtime.Value
var once_force__func_func_gopurs_runtime_Value__interface____interface___2139985499 sync.Once
func Get_force__func_func_gopurs_runtime_Value__interface____interface___2139985499() gopurs_runtime.Value {
	once_force__func_func_gopurs_runtime_Value__interface____interface___2139985499.Do(func() {
		cache_force__func_func_gopurs_runtime_Value__interface____interface___2139985499 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_force__func_func_gopurs_runtime_Value__interface____interface___2139985499(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, inner_arg0))
}))
})
	})
	return cache_force__func_func_gopurs_runtime_Value__interface____interface___2139985499
}

var cache_describe gopurs_runtime.Value
var once_describe sync.Once
func Get_describe() gopurs_runtime.Value {
	once_describe.Do(func() {
		cache_describe = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Str("Lazy Evaluation (1M Thunks Forced, 1k Depth):")), nil)
}()
})
	})
	return cache_describe
}

var cache_defer_ gopurs_runtime.Value
var once_defer_ sync.Once
func Get_defer_() gopurs_runtime.Value {
	once_defer_.Do(func() {
		cache_defer_ = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(gopurs_runtime.Value) interface{}, inner_arg1 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(Get_Lazy(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(arg0))
}), inner_arg1))
}(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, inner_arg0))
}, arg1))
})
	})
	return cache_defer_
}

var cache_defer__func_func_gopurs_runtime_Value__int64__gopurs_runtime_Value__int64_487812186 gopurs_runtime.Value
var once_defer__func_func_gopurs_runtime_Value__int64__gopurs_runtime_Value__int64_487812186 sync.Once
func Get_defer__func_func_gopurs_runtime_Value__int64__gopurs_runtime_Value__int64_487812186() gopurs_runtime.Value {
	once_defer__func_func_gopurs_runtime_Value__int64__gopurs_runtime_Value__int64_487812186.Do(func() {
		cache_defer__func_func_gopurs_runtime_Value__int64__gopurs_runtime_Value__int64_487812186 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(func(inner_arg0 func(gopurs_runtime.Value) int64, inner_arg1 gopurs_runtime.Value) int64 {
return gopurs_runtime.Apply2(Get_Lazy(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(inner_arg0(arg0))
}), inner_arg1).IntVal
}(func(inner_arg0 gopurs_runtime.Value) int64 {
return gopurs_runtime.Apply(arg0, inner_arg0).IntVal
}, arg1))
})
	})
	return cache_defer__func_func_gopurs_runtime_Value__int64__gopurs_runtime_Value__int64_487812186
}

var cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538 gopurs_runtime.Value
var once_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538 sync.Once
func Get_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538() gopurs_runtime.Value {
	once_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538.Do(func() {
		cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538 = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(func(inner_arg0 func(gopurs_runtime.Value) interface{}, inner_arg1 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(Get_Lazy(), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(arg0))
}), inner_arg1))
}(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, inner_arg0))
}, arg1))
})
	})
	return cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538
}

var cache_buildThunks gopurs_runtime.Value
var once_buildThunks sync.Once
func Get_buildThunks() gopurs_runtime.Value {
	once_buildThunks.Do(func() {
		cache_buildThunks = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_buildThunks(v_0_box.IntVal, func(inner_arg0 gopurs_runtime.Value) int64 {
return gopurs_runtime.Apply(v1_1_box, inner_arg0).IntVal
})
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
		cache_act = gopurs_runtime.Func0(func() gopurs_runtime.Value {
return func() gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Effect.Get_bindEffect(), "bind"), gopurs_runtime.Apply(pkg_Bench.Get_opaque(), gopurs_runtime.Int(1000)), gopurs_runtime.Func(func(dummy_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Effect_Console.Get_log(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Show.Get_showInt(), "show"), gopurs_runtime.Int(Call_runManyTimes(dummy_0.IntVal, 0))))
})), nil)
}()
})
	})
	return cache_act
}

func Call_Lazy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_force(v_0_loop func(gopurs_runtime.Value) interface{}) interface{} {
var v_0 func(gopurs_runtime.Value) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(pkg_Data_Unit.Get_unit())))
}

func Call_force__func_func_gopurs_runtime_Value__int64__int64_2469426587(v_0_loop func(gopurs_runtime.Value) int64) int64 {
var v_0 func(gopurs_runtime.Value) int64 = v_0_loop
_ = v_0
return gopurs_runtime.Int(v_0(pkg_Data_Unit.Get_unit())).IntVal
}

func Call_force__func_func_gopurs_runtime_Value__interface____interface___2139985499(v_0_loop func(gopurs_runtime.Value) interface{}) interface{} {
var v_0 func(gopurs_runtime.Value) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(pkg_Data_Unit.Get_unit())))
}

func Call_buildThunks(v_0_loop int64, v1_1_loop func(gopurs_runtime.Value) int64) gopurs_runtime.Value {
buildThunks:
for {
if false { continue buildThunks }
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 func(gopurs_runtime.Value) int64 = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (v_0) == (0) {
__t0 = gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(v1_1(arg0))
})
goto end_branch_0
} else {

}
}
{
v_0_loop = (v_0) - (1)
v1_1_loop = func(inner_arg0 gopurs_runtime.Value) int64 {
return gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int((gopurs_runtime.Int(v1_1(pkg_Data_Unit.Get_unit())).IntVal) + (1))
}), inner_arg0).IntVal
}
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
v1_1_loop = (v1_1) + (gopurs_runtime.Apply(Call_buildThunks(1000, func(inner_arg0 gopurs_runtime.Value) int64 {
return gopurs_runtime.Apply(gopurs_runtime.Func(func(v2_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(0)
}), inner_arg0).IntVal
}), pkg_Data_Unit.Get_unit()).IntVal)
continue runManyTimes
__t0 = gopurs_runtime.Value{}
}
end_branch_0:
return __t0.IntVal
}
}

package Control_Semigroupoid

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_semigroupoidFn gopurs_runtime.Value
var once_semigroupoidFn sync.Once
func Get_semigroupoidFn() gopurs_runtime.Value {
	once_semigroupoidFn.Do(func() {
		cache_semigroupoidFn = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
}))))
	})
	return cache_semigroupoidFn
}

var cache_compose gopurs_runtime.Value
var once_compose sync.Once
func Get_compose() gopurs_runtime.Value {
	once_compose.Do(func() {
		cache_compose = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose(dict_0_box)
})
	})
	return cache_compose
}

var cache_compose__func_gopurs_runtime_Value__func_float64__float64__func_float64__float64__func_float64__float64_447380895 gopurs_runtime.Value
var once_compose__func_gopurs_runtime_Value__func_float64__float64__func_float64__float64__func_float64__float64_447380895 sync.Once
func Get_compose__func_gopurs_runtime_Value__func_float64__float64__func_float64__float64__func_float64__float64_447380895() gopurs_runtime.Value {
	once_compose__func_gopurs_runtime_Value__func_float64__float64__func_float64__float64__func_float64__float64_447380895.Do(func() {
		cache_compose__func_gopurs_runtime_Value__func_float64__float64__func_float64__float64__func_float64__float64_447380895 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__func_gopurs_runtime_Value__func_float64__float64__func_float64__float64__func_float64__float64_447380895(dict_0_box)
})
	})
	return cache_compose__func_gopurs_runtime_Value__func_float64__float64__func_float64__float64__func_float64__float64_447380895
}

var cache_compose__func_gopurs_runtime_Value__interface____interface____interface___153747967 gopurs_runtime.Value
var once_compose__func_gopurs_runtime_Value__interface____interface____interface___153747967 sync.Once
func Get_compose__func_gopurs_runtime_Value__interface____interface____interface___153747967() gopurs_runtime.Value {
	once_compose__func_gopurs_runtime_Value__interface____interface____interface___153747967.Do(func() {
		cache_compose__func_gopurs_runtime_Value__interface____interface____interface___153747967 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__func_gopurs_runtime_Value__interface____interface____interface___153747967(dict_0_box)
})
	})
	return cache_compose__func_gopurs_runtime_Value__interface____interface____interface___153747967
}

var cache_composeFlipped gopurs_runtime.Value
var once_composeFlipped sync.Once
func Get_composeFlipped() gopurs_runtime.Value {
	once_composeFlipped.Do(func() {
		cache_composeFlipped = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_composeFlipped(dictSemigroupoid_0_box, gopurs_runtime.UnboxAny(f_1_box), gopurs_runtime.UnboxAny(g_2_box)))
})
	})
	return cache_composeFlipped
}

var cache_composeFlipped__func_gopurs_runtime_Value__func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface___2956969919 gopurs_runtime.Value
var once_composeFlipped__func_gopurs_runtime_Value__func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface___2956969919 sync.Once
func Get_composeFlipped__func_gopurs_runtime_Value__func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface___2956969919() gopurs_runtime.Value {
	once_composeFlipped__func_gopurs_runtime_Value__func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface___2956969919.Do(func() {
		cache_composeFlipped__func_gopurs_runtime_Value__func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface___2956969919 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__func_gopurs_runtime_Value__func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface___2956969919(dictSemigroupoid_0_box, func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(f_1_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
}, func(inner_arg0 func(interface{}) interface{}, inner_arg1 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(g_2_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg1)))
})
})
	})
	return cache_composeFlipped__func_gopurs_runtime_Value__func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface___2956969919
}

var cache_composeFlipped__func_gopurs_runtime_Value__interface____interface____interface___153747967 gopurs_runtime.Value
var once_composeFlipped__func_gopurs_runtime_Value__interface____interface____interface___153747967 sync.Once
func Get_composeFlipped__func_gopurs_runtime_Value__interface____interface____interface___153747967() gopurs_runtime.Value {
	once_composeFlipped__func_gopurs_runtime_Value__interface____interface____interface___153747967.Do(func() {
		cache_composeFlipped__func_gopurs_runtime_Value__interface____interface____interface___153747967 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_composeFlipped__func_gopurs_runtime_Value__interface____interface____interface___153747967(dictSemigroupoid_0_box, gopurs_runtime.UnboxAny(f_1_box), gopurs_runtime.UnboxAny(g_2_box)))
})
	})
	return cache_composeFlipped__func_gopurs_runtime_Value__interface____interface____interface___153747967
}

func Call_compose(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compose")
}

func Call_compose__func_gopurs_runtime_Value__func_float64__float64__func_float64__float64__func_float64__float64_447380895(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compose")
}

func Call_compose__func_gopurs_runtime_Value__interface____interface____interface___153747967(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "compose")
}

func Call_composeFlipped(dictSemigroupoid_0_loop gopurs_runtime.Value, f_1_loop interface{}, g_2_loop interface{}) interface{} {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 interface{} = f_1_loop
_ = f_1
var g_2 interface{} = g_2_loop
_ = g_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), gopurs_runtime.Any(g_2), gopurs_runtime.Any(f_1)))
}

func Call_composeFlipped__func_gopurs_runtime_Value__func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface____func_func_interface____interface____interface____interface___2956969919(dictSemigroupoid_0_loop gopurs_runtime.Value, f_1_loop func(func(interface{}) interface{}, interface{}) interface{}, g_2_loop func(func(interface{}) interface{}, interface{}) interface{}) gopurs_runtime.Value {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 func(func(interface{}) interface{}, interface{}) interface{} = f_1_loop
_ = f_1
var g_2 func(func(interface{}) interface{}, interface{}) interface{} = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(g_2(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(arg1)))
}), gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_1(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(arg1)))
}))
}

func Call_composeFlipped__func_gopurs_runtime_Value__interface____interface____interface___153747967(dictSemigroupoid_0_loop gopurs_runtime.Value, f_1_loop interface{}, g_2_loop interface{}) interface{} {
var dictSemigroupoid_0 gopurs_runtime.Value = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 interface{} = f_1_loop
_ = f_1
var g_2 interface{} = g_2_loop
_ = g_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroupoid_0, "compose"), gopurs_runtime.Any(g_2), gopurs_runtime.Any(f_1)))
}

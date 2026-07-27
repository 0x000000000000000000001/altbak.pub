package Data_Functor

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var cache_map_ gopurs_runtime.Value
var once_map_ sync.Once
func Get_map_() gopurs_runtime.Value {
	once_map_.Do(func() {
		cache_map_ = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map_(dict_0_box)
})
	})
	return cache_map_
}

var cache_map__func_gopurs_runtime_Value__func_interface____interface____interface____interface___4251085963 gopurs_runtime.Value
var once_map__func_gopurs_runtime_Value__func_interface____interface____interface____interface___4251085963 sync.Once
func Get_map__func_gopurs_runtime_Value__func_interface____interface____interface____interface___4251085963() gopurs_runtime.Value {
	once_map__func_gopurs_runtime_Value__func_interface____interface____interface____interface___4251085963.Do(func() {
		cache_map__func_gopurs_runtime_Value__func_interface____interface____interface____interface___4251085963 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__func_gopurs_runtime_Value__func_interface____interface____interface____interface___4251085963(dict_0_box)
})
	})
	return cache_map__func_gopurs_runtime_Value__func_interface____interface____interface____interface___4251085963
}

var cache_mapFlipped gopurs_runtime.Value
var once_mapFlipped sync.Once
func Get_mapFlipped() gopurs_runtime.Value {
	once_mapFlipped.Do(func() {
		cache_mapFlipped = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_mapFlipped(dictFunctor_0_box, gopurs_runtime.UnboxAny(fa_1_box), func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_2_box, gopurs_runtime.Any(inner_arg0)))
}))
})
	})
	return cache_mapFlipped
}

var cache_mapFlipped__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2935680491 gopurs_runtime.Value
var once_mapFlipped__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2935680491 sync.Once
func Get_mapFlipped__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2935680491() gopurs_runtime.Value {
	once_mapFlipped__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2935680491.Do(func() {
		cache_mapFlipped__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2935680491 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_mapFlipped__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2935680491(dictFunctor_0_box, gopurs_runtime.UnboxAny(fa_1_box), func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_2_box, gopurs_runtime.Any(inner_arg0)))
}))
})
	})
	return cache_mapFlipped__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2935680491
}

var cache_void gopurs_runtime.Value
var once_void sync.Once
func Get_void() gopurs_runtime.Value {
	once_void.Do(func() {
		cache_void = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_void(dictFunctor_0_box)
})
	})
	return cache_void
}

var cache_void__func_gopurs_runtime_Value__interface____interface___2088361225 gopurs_runtime.Value
var once_void__func_gopurs_runtime_Value__interface____interface___2088361225 sync.Once
func Get_void__func_gopurs_runtime_Value__interface____interface___2088361225() gopurs_runtime.Value {
	once_void__func_gopurs_runtime_Value__interface____interface___2088361225.Do(func() {
		cache_void__func_gopurs_runtime_Value__interface____interface___2088361225 = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_void__func_gopurs_runtime_Value__interface____interface___2088361225(dictFunctor_0_box)
})
	})
	return cache_void__func_gopurs_runtime_Value__interface____interface___2088361225
}

var cache_voidLeft gopurs_runtime.Value
var once_voidLeft sync.Once
func Get_voidLeft() gopurs_runtime.Value {
	once_voidLeft.Do(func() {
		cache_voidLeft = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_voidLeft(dictFunctor_0_box, gopurs_runtime.UnboxAny(f_1_box), gopurs_runtime.UnboxAny(x_2_box)))
})
	})
	return cache_voidLeft
}

var cache_voidLeft__func_gopurs_runtime_Value__interface____interface____interface___216704462 gopurs_runtime.Value
var once_voidLeft__func_gopurs_runtime_Value__interface____interface____interface___216704462 sync.Once
func Get_voidLeft__func_gopurs_runtime_Value__interface____interface____interface___216704462() gopurs_runtime.Value {
	once_voidLeft__func_gopurs_runtime_Value__interface____interface____interface___216704462.Do(func() {
		cache_voidLeft__func_gopurs_runtime_Value__interface____interface____interface___216704462 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_voidLeft__func_gopurs_runtime_Value__interface____interface____interface___216704462(dictFunctor_0_box, gopurs_runtime.UnboxAny(f_1_box), gopurs_runtime.UnboxAny(x_2_box)))
})
	})
	return cache_voidLeft__func_gopurs_runtime_Value__interface____interface____interface___216704462
}

var cache_voidRight gopurs_runtime.Value
var once_voidRight sync.Once
func Get_voidRight() gopurs_runtime.Value {
	once_voidRight.Do(func() {
		cache_voidRight = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidRight(dictFunctor_0_box, gopurs_runtime.UnboxAny(x_1_box))
})
	})
	return cache_voidRight
}

var cache_voidRight__func_gopurs_runtime_Value__interface____interface____interface___3066442509 gopurs_runtime.Value
var once_voidRight__func_gopurs_runtime_Value__interface____interface____interface___3066442509 sync.Once
func Get_voidRight__func_gopurs_runtime_Value__interface____interface____interface___3066442509() gopurs_runtime.Value {
	once_voidRight__func_gopurs_runtime_Value__interface____interface____interface___3066442509.Do(func() {
		cache_voidRight__func_gopurs_runtime_Value__interface____interface____interface___3066442509 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidRight__func_gopurs_runtime_Value__interface____interface____interface___3066442509(dictFunctor_0_box, gopurs_runtime.UnboxAny(x_1_box))
})
	})
	return cache_voidRight__func_gopurs_runtime_Value__interface____interface____interface___3066442509
}

var cache_functorProxy gopurs_runtime.Value
var once_functorProxy sync.Once
func Get_functorProxy() gopurs_runtime.Value {
	once_functorProxy.Do(func() {
		cache_functorProxy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil})
}))))
	})
	return cache_functorProxy
}

var cache_functorFn gopurs_runtime.Value
var once_functorFn sync.Once
func Get_functorFn() gopurs_runtime.Value {
	once_functorFn.Do(func() {
		cache_functorFn = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))))
	})
	return cache_functorFn
}

var cache_functorArray gopurs_runtime.Value
var once_functorArray sync.Once
func Get_functorArray() gopurs_runtime.Value {
	once_functorArray.Do(func() {
		cache_functorArray = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("map", Get_arrayMap())))
	})
	return cache_functorArray
}

var cache_flap gopurs_runtime.Value
var once_flap sync.Once
func Get_flap() gopurs_runtime.Value {
	once_flap.Do(func() {
		cache_flap = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, ff_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_flap(dictFunctor_0_box, gopurs_runtime.UnboxAny(ff_1_box), gopurs_runtime.UnboxAny(x_2_box)))
})
	})
	return cache_flap
}

var cache_arrayMap gopurs_runtime.Value
var once_arrayMap sync.Once
func Get_arrayMap() gopurs_runtime.Value {
	once_arrayMap.Do(func() {
		cache_arrayMap = gopurs_runtime.Func2(func(arg0 gopurs_runtime.Value, arg1 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := ArrayMap(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(arg0, gopurs_runtime.Any(inner_arg0)))
}, func() []interface{} {
					arr := *(*[]gopurs_runtime.Value)(arg1.UnsafePtr)
					unboxed := make([]interface{}, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.UnboxAny(v) }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Any(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_arrayMap
}

func Call_map_(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "map")
}

func Call_map__func_gopurs_runtime_Value__func_interface____interface____interface____interface___4251085963(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "map")
}

func Call_mapFlipped(dictFunctor_0_loop gopurs_runtime.Value, fa_1_loop interface{}, f_2_loop func(interface{}) interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 interface{} = fa_1_loop
_ = fa_1
var f_2 func(interface{}) interface{} = f_2_loop
_ = f_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_2(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(fa_1)))
}

func Call_mapFlipped__func_gopurs_runtime_Value__interface____func_interface____interface____interface___2935680491(dictFunctor_0_loop gopurs_runtime.Value, fa_1_loop interface{}, f_2_loop func(interface{}) interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 interface{} = fa_1_loop
_ = fa_1
var f_2 func(interface{}) interface{} = f_2_loop
_ = f_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_2(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(fa_1)))
}

func Call_void(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
}

func Call_void__func_gopurs_runtime_Value__interface____interface___2088361225(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
}

func Call_voidLeft(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop interface{}, x_2_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 interface{} = f_1_loop
_ = f_1
var x_2 interface{} = x_2_loop
_ = x_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(x_2)
}), gopurs_runtime.Any(f_1)))
}

func Call_voidLeft__func_gopurs_runtime_Value__interface____interface____interface___216704462(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop interface{}, x_2_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 interface{} = f_1_loop
_ = f_1
var x_2 interface{} = x_2_loop
_ = x_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(x_2)
}), gopurs_runtime.Any(f_1)))
}

func Call_voidRight(dictFunctor_0_loop gopurs_runtime.Value, x_1_loop interface{}) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 interface{} = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(x_1)
}))
}

func Call_voidRight__func_gopurs_runtime_Value__interface____interface____interface___3066442509(dictFunctor_0_loop gopurs_runtime.Value, x_1_loop interface{}) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 interface{} = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(x_1)
}))
}

func Call_flap(dictFunctor_0_loop gopurs_runtime.Value, ff_1_loop interface{}, x_2_loop interface{}) interface{} {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var ff_1 interface{} = ff_1_loop
_ = ff_1
var x_2 interface{} = x_2_loop
_ = x_2
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Any(x_2))
}), gopurs_runtime.Any(ff_1)))
}

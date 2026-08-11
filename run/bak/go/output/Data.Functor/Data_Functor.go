package Data_Functor

import (
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_go__map gopurs_runtime.Value
var once_go__map sync.Once
func Get_go__map() gopurs_runtime.Value {
	once_go__map.Do(func() {
		cache_go__map = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_go__map(dict_0_box)
})
	})
	return cache_go__map
}

var cache_map__gopurs_runtime_Value_2199395572 gopurs_runtime.Value
var once_map__gopurs_runtime_Value_2199395572 sync.Once
func Get_map__gopurs_runtime_Value_2199395572() gopurs_runtime.Value {
	once_map__gopurs_runtime_Value_2199395572.Do(func() {
		cache_map__gopurs_runtime_Value_2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__gopurs_runtime_Value_2199395572(dict_0_box)
})
	})
	return cache_map__gopurs_runtime_Value_2199395572
}

var cache_mapFlipped gopurs_runtime.Value
var once_mapFlipped sync.Once
func Get_mapFlipped() gopurs_runtime.Value {
	once_mapFlipped.Do(func() {
		cache_mapFlipped = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped(dictFunctor_0_box, fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped
}

var cache_mapFlipped__gopurs_runtime_Value_4215217780 gopurs_runtime.Value
var once_mapFlipped__gopurs_runtime_Value_4215217780 sync.Once
func Get_mapFlipped__gopurs_runtime_Value_4215217780() gopurs_runtime.Value {
	once_mapFlipped__gopurs_runtime_Value_4215217780.Do(func() {
		cache_mapFlipped__gopurs_runtime_Value_4215217780 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, fa_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapFlipped__gopurs_runtime_Value_4215217780(dictFunctor_0_box, fa_1_box, f_2_box)
})
	})
	return cache_mapFlipped__gopurs_runtime_Value_4215217780
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

var cache_void__gopurs_runtime_Value_3020373336 gopurs_runtime.Value
var once_void__gopurs_runtime_Value_3020373336 sync.Once
func Get_void__gopurs_runtime_Value_3020373336() gopurs_runtime.Value {
	once_void__gopurs_runtime_Value_3020373336.Do(func() {
		cache_void__gopurs_runtime_Value_3020373336 = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_void__gopurs_runtime_Value_3020373336(dictFunctor_0_box)
})
	})
	return cache_void__gopurs_runtime_Value_3020373336
}

var cache_voidLeft gopurs_runtime.Value
var once_voidLeft sync.Once
func Get_voidLeft() gopurs_runtime.Value {
	once_voidLeft.Do(func() {
		cache_voidLeft = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidLeft(dictFunctor_0_box, f_1_box, x_2_box)
})
	})
	return cache_voidLeft
}

var cache_voidLeft__gopurs_runtime_Value_32301756 gopurs_runtime.Value
var once_voidLeft__gopurs_runtime_Value_32301756 sync.Once
func Get_voidLeft__gopurs_runtime_Value_32301756() gopurs_runtime.Value {
	once_voidLeft__gopurs_runtime_Value_32301756.Do(func() {
		cache_voidLeft__gopurs_runtime_Value_32301756 = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidLeft__gopurs_runtime_Value_32301756(dictFunctor_0_box, f_1_box, x_2_box)
})
	})
	return cache_voidLeft__gopurs_runtime_Value_32301756
}

var cache_voidRight gopurs_runtime.Value
var once_voidRight sync.Once
func Get_voidRight() gopurs_runtime.Value {
	once_voidRight.Do(func() {
		cache_voidRight = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidRight(dictFunctor_0_box, x_1_box)
})
	})
	return cache_voidRight
}

var cache_voidRight__gopurs_runtime_Value_1142845180 gopurs_runtime.Value
var once_voidRight__gopurs_runtime_Value_1142845180 sync.Once
func Get_voidRight__gopurs_runtime_Value_1142845180() gopurs_runtime.Value {
	once_voidRight__gopurs_runtime_Value_1142845180.Do(func() {
		cache_voidRight__gopurs_runtime_Value_1142845180 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidRight__gopurs_runtime_Value_1142845180(dictFunctor_0_box, x_1_box)
})
	})
	return cache_voidRight__gopurs_runtime_Value_1142845180
}

var cache_functorProxy gopurs_runtime.Value
var once_functorProxy sync.Once
func Get_functorProxy() gopurs_runtime.Value {
	once_functorProxy.Do(func() {
		cache_functorProxy = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}
})
}))
	})
	return cache_functorProxy
}

var cache_functorProxy__gopurs_runtime_Value_711768561 gopurs_runtime.Value
var once_functorProxy__gopurs_runtime_Value_711768561 sync.Once
func Get_functorProxy__gopurs_runtime_Value_711768561() gopurs_runtime.Value {
	once_functorProxy__gopurs_runtime_Value_711768561.Do(func() {
		cache_functorProxy__gopurs_runtime_Value_711768561 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 513803634, UnsafePtr: nil}
})
}))
	})
	return cache_functorProxy__gopurs_runtime_Value_711768561
}

var cache_functorFn gopurs_runtime.Value
var once_functorFn sync.Once
func Get_functorFn() gopurs_runtime.Value {
	once_functorFn.Do(func() {
		cache_functorFn = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_functorFn
}

var cache_functorFn__gopurs_runtime_Value_20325936 gopurs_runtime.Value
var once_functorFn__gopurs_runtime_Value_20325936 sync.Once
func Get_functorFn__gopurs_runtime_Value_20325936() gopurs_runtime.Value {
	once_functorFn__gopurs_runtime_Value_20325936.Do(func() {
		cache_functorFn__gopurs_runtime_Value_20325936 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_functorFn__gopurs_runtime_Value_20325936
}

var cache_functorArray gopurs_runtime.Value
var once_functorArray sync.Once
func Get_functorArray() gopurs_runtime.Value {
	once_functorArray.Do(func() {
		cache_functorArray = gopurs_runtime.RecordDict1("map", Get_arrayMap())
	})
	return cache_functorArray
}

var cache_functorArray__gopurs_runtime_Value_361387505 gopurs_runtime.Value
var once_functorArray__gopurs_runtime_Value_361387505 sync.Once
func Get_functorArray__gopurs_runtime_Value_361387505() gopurs_runtime.Value {
	once_functorArray__gopurs_runtime_Value_361387505.Do(func() {
		cache_functorArray__gopurs_runtime_Value_361387505 = gopurs_runtime.RecordDict1("map", Get_arrayMap())
	})
	return cache_functorArray__gopurs_runtime_Value_361387505
}

var cache_flap gopurs_runtime.Value
var once_flap sync.Once
func Get_flap() gopurs_runtime.Value {
	once_flap.Do(func() {
		cache_flap = gopurs_runtime.Func3(func(dictFunctor_0_box gopurs_runtime.Value, ff_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flap(dictFunctor_0_box, ff_1_box, x_2_box)
})
	})
	return cache_flap
}

func Call_go__map(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "map")
}

func Call_map__gopurs_runtime_Value_2199395572(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "map")
}

func Call_mapFlipped(dictFunctor_0_loop gopurs_runtime.Value, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, fa_1)
}

func Call_mapFlipped__gopurs_runtime_Value_4215217780(dictFunctor_0_loop gopurs_runtime.Value, fa_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var fa_1 gopurs_runtime.Value = fa_1_loop
_ = fa_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_2, fa_1)
}

func Call_void(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
}

func Call_void__gopurs_runtime_Value_3020373336(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}))
}

func Call_voidLeft(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), f_1)
}

func Call_voidLeft__gopurs_runtime_Value_32301756(dictFunctor_0_loop gopurs_runtime.Value, f_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), f_1)
}

func Call_voidRight(dictFunctor_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_voidRight__gopurs_runtime_Value_1142845180(dictFunctor_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_flap(dictFunctor_0_loop gopurs_runtime.Value, ff_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
var ff_1 gopurs_runtime.Value = ff_1_loop
_ = ff_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, x_2)
}), ff_1)
}

func Get_arrayMap() gopurs_runtime.Value {
	return _Gopurs_ArrayMap
}

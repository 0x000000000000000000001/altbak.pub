package Data_Profunctor_Split

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_identity(gopurs_runtime.UnboxAny(x_0_box)))
})
	})
	return cache_identity
}

var cache_SplitF gopurs_runtime.Value
var once_SplitF sync.Once
func Get_SplitF() gopurs_runtime.Value {
	once_SplitF.Do(func() {
		cache_SplitF = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(value0, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(value1, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(value2)})}
})
})
})
	})
	return cache_SplitF
}

var cache_unSplit gopurs_runtime.Value
var once_unSplit sync.Once
func Get_unSplit() gopurs_runtime.Value {
	once_unSplit.Do(func() {
		cache_unSplit = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unSplit(func(inner_arg0 func(interface{}) interface{}, inner_arg1 func(interface{}) interface{}, inner_arg2 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply3(f_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg2)))
}, v_1_box))
})
	})
	return cache_unSplit
}

var cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_867043384 gopurs_runtime.Value
var once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_867043384 sync.Once
func Get_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_867043384() gopurs_runtime.Value {
	once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_867043384.Do(func() {
		cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_867043384 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_867043384(func(inner_arg0 func(interface{}) interface{}, inner_arg1 func(interface{}) interface{}, inner_arg2 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg2))
}, v_1_box)
})
	})
	return cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_867043384
}

var cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_820818299 gopurs_runtime.Value
var once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_820818299 sync.Once
func Get_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_820818299() gopurs_runtime.Value {
	once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_820818299.Do(func() {
		cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_820818299 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_820818299(func(inner_arg0 func(interface{}) interface{}, inner_arg1 func(interface{}) interface{}, inner_arg2 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg2))
}, v_1_box)
})
	})
	return cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_820818299
}

var cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___3521066075 gopurs_runtime.Value
var once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___3521066075 sync.Once
func Get_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___3521066075() gopurs_runtime.Value {
	once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___3521066075.Do(func() {
		cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___3521066075 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___3521066075(func(inner_arg0 func(interface{}) interface{}, inner_arg1 func(interface{}) interface{}, inner_arg2 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply3(f_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg2)))
}, v_1_box))
})
	})
	return cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___3521066075
}

var cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___401747963 gopurs_runtime.Value
var once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___401747963 sync.Once
func Get_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___401747963() gopurs_runtime.Value {
	once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___401747963.Do(func() {
		cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___401747963 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___401747963(func(inner_arg0 func(interface{}) interface{}, inner_arg1 func(interface{}) interface{}, inner_arg2 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply3(f_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg2)))
}, v_1_box))
})
	})
	return cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___401747963
}

var cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_3498664473 gopurs_runtime.Value
var once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_3498664473 sync.Once
func Get_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_3498664473() gopurs_runtime.Value {
	once_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_3498664473.Do(func() {
		cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_3498664473 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_3498664473(func(inner_arg0 func(interface{}) interface{}, inner_arg1 func(interface{}) interface{}, inner_arg2 interface{}) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0_box, gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(inner_arg1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg2))
}, v_1_box)
})
	})
	return cache_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_3498664473
}

var cache_split gopurs_runtime.Value
var once_split sync.Once
func Get_split() gopurs_runtime.Value {
	once_split.Do(func() {
		cache_split = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, fx_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_split(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(fx_2_box))
})
	})
	return cache_split
}

var cache_split__func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value_1066366586 gopurs_runtime.Value
var once_split__func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value_1066366586 sync.Once
func Get_split__func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value_1066366586() gopurs_runtime.Value {
	once_split__func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value_1066366586.Do(func() {
		cache_split__func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value_1066366586 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, g_1_box gopurs_runtime.Value, fx_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_split__func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value_1066366586(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(g_1_box, gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(fx_2_box))
})
	})
	return cache_split__func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value_1066366586
}

var cache_profunctorSplit gopurs_runtime.Value
var once_profunctorSplit sync.Once
func Get_profunctorSplit() gopurs_runtime.Value {
	once_profunctorSplit.Do(func() {
		cache_profunctorSplit = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("dimap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, g_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Apply(f_0, x_3))
}), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(g_1, gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), x_3))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V2))})})
}))))
	})
	return cache_profunctorSplit
}

var cache_lowerSplit gopurs_runtime.Value
var once_lowerSplit sync.Once
func Get_lowerSplit() gopurs_runtime.Value {
	once_lowerSplit.Do(func() {
		cache_lowerSplit = gopurs_runtime.Func2(func(dictInvariant_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_lowerSplit(dictInvariant_0_box, v_1_box))
})
	})
	return cache_lowerSplit
}

var cache_liftSplit gopurs_runtime.Value
var once_liftSplit sync.Once
func Get_liftSplit() gopurs_runtime.Value {
	once_liftSplit.Do(func() {
		cache_liftSplit = gopurs_runtime.Func(func(fx_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_liftSplit(gopurs_runtime.UnboxAny(fx_0_box))
})
	})
	return cache_liftSplit
}

var cache_hoistSplit gopurs_runtime.Value
var once_hoistSplit sync.Once
func Get_hoistSplit() gopurs_runtime.Value {
	once_hoistSplit.Do(func() {
		cache_hoistSplit = gopurs_runtime.Func2(func(nat_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_hoistSplit(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(nat_0_box, gopurs_runtime.Any(inner_arg0)))
}, v_1_box)
})
	})
	return cache_hoistSplit
}

var cache_functorSplit gopurs_runtime.Value
var once_functorSplit sync.Once
func Get_functorSplit() gopurs_runtime.Value {
	once_functorSplit.Do(func() {
		cache_functorSplit = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), x_2))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2))})})
}))))
	})
	return cache_functorSplit
}

type Constructor_SplitF[T_f any, T_a any, T_b any, T_x any] struct {
	V0 func(interface{}) interface{}
	V1 func(interface{}) interface{}
	V2 T_f
}


func Call_identity(x_0_loop interface{}) interface{} {
var x_0 interface{} = x_0_loop
_ = x_0
return x_0
}

func Call_unSplit(f_0_loop func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_0 func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)))))
}

func Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_867043384(f_0_loop func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return f_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)))
}

func Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_820818299(f_0_loop func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return f_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)))
}

func Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___3521066075(f_0_loop func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_0 func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)))))
}

func Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____interface____gopurs_runtime_Value__interface___401747963(f_0_loop func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) interface{}, v_1_loop gopurs_runtime.Value) interface{} {
var f_0 func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) interface{} = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(f_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)))))
}

func Call_unSplit__func_func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value__gopurs_runtime_Value__gopurs_runtime_Value_3498664473(f_0_loop func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 func(func(interface{}) interface{}, func(interface{}) interface{}, interface{}) gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return f_0(func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)))
}

func Call_split(f_0_loop func(interface{}) interface{}, g_1_loop func(interface{}) interface{}, fx_2_loop interface{}) gopurs_runtime.Value {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var g_1 func(interface{}) interface{} = g_1_loop
_ = g_1
var fx_2 interface{} = fx_2_loop
_ = fx_2
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{f_0, g_1, fx_2})})
}

func Call_split__func_func_interface____interface____func_interface____interface____interface____gopurs_runtime_Value_1066366586(f_0_loop func(interface{}) interface{}, g_1_loop func(interface{}) interface{}, fx_2_loop interface{}) gopurs_runtime.Value {
var f_0 func(interface{}) interface{} = f_0_loop
_ = f_0
var g_1 func(interface{}) interface{} = g_1_loop
_ = g_1
var fx_2 interface{} = fx_2_loop
_ = fx_2
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{f_0, g_1, fx_2})})
}

func Call_lowerSplit(dictInvariant_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) interface{} {
var dictInvariant_0 gopurs_runtime.Value = dictInvariant_0_loop
_ = dictInvariant_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictInvariant_0, "imap"), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)))
}

func Call_liftSplit(fx_0_loop interface{}) gopurs_runtime.Value {
var fx_0 interface{} = fx_0_loop
_ = fx_0
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(Get_identity(), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(Get_identity(), gopurs_runtime.Any(inner_arg0)))
}, fx_0})})
}

func Call_hoistSplit(nat_0_loop func(interface{}) interface{}, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var nat_0 func(interface{}) interface{} = nat_0_loop
_ = nat_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Any(gopurs_runtime.Value{Type: 9, IntVal: 1995432569, UnsafePtr: unsafe.Pointer(&Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value]{func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, func(inner_arg0 interface{}) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1(gopurs_runtime.UnboxAny(arg0)))
}), gopurs_runtime.Any(inner_arg0)))
}, gopurs_runtime.UnboxAny(gopurs_runtime.Any(nat_0(gopurs_runtime.UnboxAny(gopurs_runtime.Any((*Constructor_SplitF[gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V2)))))})})
}

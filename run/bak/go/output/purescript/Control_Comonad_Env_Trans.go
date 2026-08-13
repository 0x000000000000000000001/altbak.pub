package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Control_Comonad_Env_Trans_EnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_EnvT sync.Once
func Get_Control_Comonad_Env_Trans_EnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_EnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_EnvT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Trans_EnvT(x_0_box)
})
	})
	return cache_Control_Comonad_Env_Trans_EnvT
}

var cache_Control_Comonad_Env_Trans_withEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_withEnvT sync.Once
func Get_Control_Comonad_Env_Trans_withEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_withEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_withEnvT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Env_Trans_withEnvT(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box)))}
})
	})
	return cache_Control_Comonad_Env_Trans_withEnvT
}

var cache_Control_Comonad_Env_Trans_runEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_runEnvT sync.Once
func Get_Control_Comonad_Env_Trans_runEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_runEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_runEnvT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Env_Trans_runEnvT(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_0_box)))}
})
	})
	return cache_Control_Comonad_Env_Trans_runEnvT
}

var cache_Control_Comonad_Env_Trans_newtypeEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_newtypeEnvT sync.Once
func Get_Control_Comonad_Env_Trans_newtypeEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_newtypeEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_newtypeEnvT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_Control_Comonad_Env_Trans_newtypeEnvT
}

var cache_Control_Comonad_Env_Trans_mapEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_mapEnvT sync.Once
func Get_Control_Comonad_Env_Trans_mapEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_mapEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_mapEnvT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Env_Trans_mapEnvT(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box)))}
})
	})
	return cache_Control_Comonad_Env_Trans_mapEnvT
}

var cache_Control_Comonad_Env_Trans_functorEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_functorEnvT sync.Once
func Get_Control_Comonad_Env_Trans_functorEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_functorEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_functorEnvT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Trans_functorEnvT(dictFunctor_0_box)
})
	})
	return cache_Control_Comonad_Env_Trans_functorEnvT
}

var cache_Control_Comonad_Env_Trans_functorWithIndexEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_functorWithIndexEnvT sync.Once
func Get_Control_Comonad_Env_Trans_functorWithIndexEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_functorWithIndexEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_functorWithIndexEnvT = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Trans_functorWithIndexEnvT(dictFunctorWithIndex_0_box)
})
	})
	return cache_Control_Comonad_Env_Trans_functorWithIndexEnvT
}

var cache_Control_Comonad_Env_Trans_foldableEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_foldableEnvT sync.Once
func Get_Control_Comonad_Env_Trans_foldableEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_foldableEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_foldableEnvT = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Trans_foldableEnvT(dictFoldable_0_box)
})
	})
	return cache_Control_Comonad_Env_Trans_foldableEnvT
}

var cache_Control_Comonad_Env_Trans_foldableWithIndexEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_foldableWithIndexEnvT sync.Once
func Get_Control_Comonad_Env_Trans_foldableWithIndexEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_foldableWithIndexEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_foldableWithIndexEnvT = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Trans_foldableWithIndexEnvT(dictFoldableWithIndex_0_box)
})
	})
	return cache_Control_Comonad_Env_Trans_foldableWithIndexEnvT
}

var cache_Control_Comonad_Env_Trans_traversableEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_traversableEnvT sync.Once
func Get_Control_Comonad_Env_Trans_traversableEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_traversableEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_traversableEnvT = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Trans_traversableEnvT(dictTraversable_0_box)
})
	})
	return cache_Control_Comonad_Env_Trans_traversableEnvT
}

var cache_Control_Comonad_Env_Trans_traversableWithIndexEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_traversableWithIndexEnvT sync.Once
func Get_Control_Comonad_Env_Trans_traversableWithIndexEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_traversableWithIndexEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_traversableWithIndexEnvT = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Trans_traversableWithIndexEnvT(dictTraversableWithIndex_0_box)
})
	})
	return cache_Control_Comonad_Env_Trans_traversableWithIndexEnvT
}

var cache_Control_Comonad_Env_Trans_extendEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_extendEnvT sync.Once
func Get_Control_Comonad_Env_Trans_extendEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_extendEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_extendEnvT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Trans_extendEnvT(dictExtend_0_box)
})
	})
	return cache_Control_Comonad_Env_Trans_extendEnvT
}

var cache_Control_Comonad_Env_Trans_comonadTransEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_comonadTransEnvT sync.Once
func Get_Control_Comonad_Env_Trans_comonadTransEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_comonadTransEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_comonadTransEnvT = gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(v_1.UnsafePtr).V1
})
}))
	})
	return cache_Control_Comonad_Env_Trans_comonadTransEnvT
}

var cache_Control_Comonad_Env_Trans_comonadEnvT gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_comonadEnvT sync.Once
func Get_Control_Comonad_Env_Trans_comonadEnvT() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_comonadEnvT.Do(func() {
		cache_Control_Comonad_Env_Trans_comonadEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Control_Comonad_Env_Trans_comonadEnvT(dictComonad_0_box)
})
	})
	return cache_Control_Comonad_Env_Trans_comonadEnvT
}

var cache_Control_Comonad_Env_Trans_withEnvT__1680291482 gopurs_runtime.Value
var once_Control_Comonad_Env_Trans_withEnvT__1680291482 sync.Once
func Get_Control_Comonad_Env_Trans_withEnvT__1680291482() gopurs_runtime.Value {
	once_Control_Comonad_Env_Trans_withEnvT__1680291482.Do(func() {
		cache_Control_Comonad_Env_Trans_withEnvT__1680291482 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Control_Comonad_Env_Trans_withEnvT__1680291482(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v_1_box)))}
})
	})
	return cache_Control_Comonad_Env_Trans_withEnvT__1680291482
}

func Call_Control_Comonad_Env_Trans_EnvT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Control_Comonad_Env_Trans_withEnvT(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_0, (v_1).V0), (v_1).V1}
}

func Call_Control_Comonad_Env_Trans_runEnvT(v_0_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var v_0 *Constructor_Data_Tuple_Tuple = v_0_loop
_ = v_0
return v_0
}

func Call_Control_Comonad_Env_Trans_mapEnvT(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return &Constructor_Data_Tuple_Tuple{1, (v_1).V0, gopurs_runtime.Apply(f_0, (v_1).V1)}
}

func Call_Control_Comonad_Env_Trans_functorEnvT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)})}
})
}))
}

func Call_Control_Comonad_Env_Trans_functorWithIndexEnvT(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorEnvT1_1_0 -> gopurs_runtime.Value
functorEnvT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)})}
})
}))
}

func Call_Control_Comonad_Env_Trans_foldableEnvT(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fn_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_1))}, fn_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), fn_1, a_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), fn_1, a_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)
})
})
}))
}

func Call_Control_Comonad_Env_Trans_foldableWithIndexEnvT(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): foldableEnvT1_1_0 -> gopurs_runtime.Value
foldableEnvT1_1_0 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fn_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, fn_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldl"), fn_2, a_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldr"), fn_2, a_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)
})
})
}))
_ = foldableEnvT1_1_0
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_1_0
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_2))}, f_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), f_2, a_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), f_2, a_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)
})
})
}))
}

func Call_Control_Comonad_Env_Trans_traversableEnvT(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): functorEnvT1_1_0 -> gopurs_runtime.Value
functorEnvT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_Data_Tuple_Tuple)(v_3.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_1_0
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): foldableEnvT1_2_2 -> gopurs_runtime.Value
foldableEnvT1_2_2 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fn_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_3))}, fn_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), fn_3, a_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), fn_3, a_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1)
})
})
}))
_ = foldableEnvT1_2_2
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_4 -> *Constructor_Data_Functor_Functor
Functor0_4_4 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_4.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_4_5 -> *Constructor_Data_Functor_Functor
Functor0_4_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_4_5.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return x_7
}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_3))}, f_5, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1))
})
})
}))
}

func Call_Control_Comonad_Env_Trans_traversableWithIndexEnvT(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): __local_var_2_3 -> gopurs_runtime.Value
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
// TAST (Let): functorEnvT1_2_2 -> gopurs_runtime.Value
functorEnvT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_2_2
// TAST (Let): functorWithIndexEnvT1_1_0 -> gopurs_runtime.Value
functorWithIndexEnvT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_2
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mapWithIndex"), f_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)})}
})
}))
_ = functorWithIndexEnvT1_1_0
// TAST (Let): __local_var_2_5 -> gopurs_runtime.Value
__local_var_2_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{})
_ = __local_var_2_5
// TAST (Let): __local_var_3_7 -> gopurs_runtime.Value
__local_var_3_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_2_5, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_3_7
// TAST (Let): foldableEnvT1_3_6 -> gopurs_runtime.Value
foldableEnvT1_3_6 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fn_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_7, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_4))}, fn_5, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_7, "foldl"), fn_4, a_5, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_7, "foldr"), fn_4, a_5, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1)
})
})
}))
_ = foldableEnvT1_3_6
// TAST (Let): foldableWithIndexEnvT1_2_4 -> gopurs_runtime.Value
foldableWithIndexEnvT1_2_4 := gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_3_6
}), gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_5, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_4))}, f_5, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_5, "foldlWithIndex"), f_4, a_5, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_5, "foldrWithIndex"), f_4, a_5, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1)
})
})
}))
_ = foldableWithIndexEnvT1_2_4
// TAST (Let): __local_var_3_9 -> gopurs_runtime.Value
__local_var_3_9 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{})
_ = __local_var_3_9
// TAST (Let): __local_var_4_11 -> gopurs_runtime.Value
__local_var_4_11 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_9, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_4_11
// TAST (Let): functorEnvT1_4_10 -> gopurs_runtime.Value
functorEnvT1_4_10 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_4_11, "map"), f_5, (*Constructor_Data_Tuple_Tuple)(v_6.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_4_10
// TAST (Let): __local_var_5_13 -> gopurs_runtime.Value
__local_var_5_13 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_9, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_5_13
// TAST (Let): foldableEnvT1_5_12 -> gopurs_runtime.Value
foldableEnvT1_5_12 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fn_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_13, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_6))}, fn_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_13, "foldl"), fn_6, a_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_5_13, "foldr"), fn_6, a_7, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1)
})
})
}))
_ = foldableEnvT1_5_12
// TAST (Let): traversableEnvT1_3_8 -> gopurs_runtime.Value
traversableEnvT1_3_8 := gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_5_12
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_4_10
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_14 -> *Constructor_Data_Functor_Functor
Functor0_7_14 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_14
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_14.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return x_9
}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_9, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, (*Constructor_Data_Tuple_Tuple)(v_8.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_7_15 -> *Constructor_Data_Functor_Functor
Functor0_7_15 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_7_15
return gopurs_runtime.Func(func(f_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_7_15.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Func(func(x_10 gopurs_runtime.Value) gopurs_runtime.Value {
return x_10
}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_9, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_6))}, f_8, (*Constructor_Data_Tuple_Tuple)(v_9.UnsafePtr).V1))
})
})
}))
_ = traversableEnvT1_3_8
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexEnvT1_2_4
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexEnvT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableEnvT1_3_8
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_5_16 -> *Constructor_Data_Functor_Functor
Functor0_5_16 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_16
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_5_16.V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorFn(), "map"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return x_8
}), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_4))}, f_6, (*Constructor_Data_Tuple_Tuple)(v_7.UnsafePtr).V1))
})
})
}))
}

func Call_Control_Comonad_Env_Trans_extendEnvT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
// TAST (Let): __local_var_2_2 -> gopurs_runtime.Value
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
// TAST (Let): functorEnvT1_2_1 -> gopurs_runtime.Value
functorEnvT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), f_3, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_2_1
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_3 -> gopurs_runtime.Value
__local_var_5_3 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0)
_ = __local_var_5_3
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), f_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
}), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1))})}
})
}))
}

func Call_Control_Comonad_Env_Trans_comonadEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
// TAST (Let): __local_var_1_1 -> gopurs_runtime.Value
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{})
_ = __local_var_1_1
// TAST (Let): Functor0_2_2 -> *Constructor_Data_Functor_Functor
Functor0_2_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_2
// TAST (Let): __local_var_3_4 -> gopurs_runtime.Value
__local_var_3_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_3_4
// TAST (Let): functorEnvT1_3_3 -> gopurs_runtime.Value
functorEnvT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_3_4, "map"), f_4, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_3_3
// TAST (Let): extendEnvT1_1_0 -> gopurs_runtime.Value
extendEnvT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_3_3
}), gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> gopurs_runtime.Value
__local_var_6_5 := gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0)
_ = __local_var_6_5
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_2.V0), f_4, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "extend"), gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_5, x_7)
}), (*Constructor_Data_Tuple_Tuple)(v_5.UnsafePtr).V1))})}
})
}))
_ = extendEnvT1_1_0
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return extendEnvT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*Constructor_Data_Tuple_Tuple)(v_2.UnsafePtr).V1)
}))
}

func Call_Control_Comonad_Env_Trans_withEnvT__1680291482(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_Data_Tuple_Tuple = v_1_loop
_ = v_1
return &Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(f_0, (v_1).V0), (v_1).V1}
}



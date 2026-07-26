package Control_Comonad_Env_Trans

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	unsafe "unsafe"
)

var cache_EnvT gopurs_runtime.Value
var once_EnvT sync.Once
func Get_EnvT() gopurs_runtime.Value {
	once_EnvT.Do(func() {
		cache_EnvT = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_EnvT(x_0_box)
})
	})
	return cache_EnvT
}

var cache_withEnvT gopurs_runtime.Value
var once_withEnvT sync.Once
func Get_withEnvT() gopurs_runtime.Value {
	once_withEnvT.Do(func() {
		cache_withEnvT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_withEnvT(f_0_box, v_1_box)
})
	})
	return cache_withEnvT
}

var cache_runEnvT gopurs_runtime.Value
var once_runEnvT sync.Once
func Get_runEnvT() gopurs_runtime.Value {
	once_runEnvT.Do(func() {
		cache_runEnvT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runEnvT(v_0_box)
})
	})
	return cache_runEnvT
}

var cache_newtypeEnvT gopurs_runtime.Value
var once_newtypeEnvT sync.Once
func Get_newtypeEnvT() gopurs_runtime.Value {
	once_newtypeEnvT.Do(func() {
		cache_newtypeEnvT = gopurs_runtime.RecordDict1("Coercible0", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{}
}))
	})
	return cache_newtypeEnvT
}

var cache_mapEnvT gopurs_runtime.Value
var once_mapEnvT sync.Once
func Get_mapEnvT() gopurs_runtime.Value {
	once_mapEnvT.Do(func() {
		cache_mapEnvT = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapEnvT(f_0_box, v_1_box)
})
	})
	return cache_mapEnvT
}

var cache_functorEnvT gopurs_runtime.Value
var once_functorEnvT sync.Once
func Get_functorEnvT() gopurs_runtime.Value {
	once_functorEnvT.Do(func() {
		cache_functorEnvT = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorEnvT(dictFunctor_0_box)
})
	})
	return cache_functorEnvT
}

var cache_functorWithIndexEnvT gopurs_runtime.Value
var once_functorWithIndexEnvT sync.Once
func Get_functorWithIndexEnvT() gopurs_runtime.Value {
	once_functorWithIndexEnvT.Do(func() {
		cache_functorWithIndexEnvT = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorWithIndexEnvT(dictFunctorWithIndex_0_box)
})
	})
	return cache_functorWithIndexEnvT
}

var cache_foldableEnvT gopurs_runtime.Value
var once_foldableEnvT sync.Once
func Get_foldableEnvT() gopurs_runtime.Value {
	once_foldableEnvT.Do(func() {
		cache_foldableEnvT = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableEnvT(dictFoldable_0_box)
})
	})
	return cache_foldableEnvT
}

var cache_foldableWithIndexEnvT gopurs_runtime.Value
var once_foldableWithIndexEnvT sync.Once
func Get_foldableWithIndexEnvT() gopurs_runtime.Value {
	once_foldableWithIndexEnvT.Do(func() {
		cache_foldableWithIndexEnvT = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableWithIndexEnvT(dictFoldableWithIndex_0_box)
})
	})
	return cache_foldableWithIndexEnvT
}

var cache_traversableEnvT gopurs_runtime.Value
var once_traversableEnvT sync.Once
func Get_traversableEnvT() gopurs_runtime.Value {
	once_traversableEnvT.Do(func() {
		cache_traversableEnvT = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableEnvT(dictTraversable_0_box)
})
	})
	return cache_traversableEnvT
}

var cache_traversableWithIndexEnvT gopurs_runtime.Value
var once_traversableWithIndexEnvT sync.Once
func Get_traversableWithIndexEnvT() gopurs_runtime.Value {
	once_traversableWithIndexEnvT.Do(func() {
		cache_traversableWithIndexEnvT = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableWithIndexEnvT(dictTraversableWithIndex_0_box)
})
	})
	return cache_traversableWithIndexEnvT
}

var cache_extendEnvT gopurs_runtime.Value
var once_extendEnvT sync.Once
func Get_extendEnvT() gopurs_runtime.Value {
	once_extendEnvT.Do(func() {
		cache_extendEnvT = gopurs_runtime.Func(func(dictExtend_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extendEnvT(dictExtend_0_box)
})
	})
	return cache_extendEnvT
}

var cache_comonadTransEnvT gopurs_runtime.Value
var once_comonadTransEnvT sync.Once
func Get_comonadTransEnvT() gopurs_runtime.Value {
	once_comonadTransEnvT.Do(func() {
		cache_comonadTransEnvT = gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func2(func(dictComonad_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
}))
	})
	return cache_comonadTransEnvT
}

var cache_comonadEnvT gopurs_runtime.Value
var once_comonadEnvT sync.Once
func Get_comonadEnvT() gopurs_runtime.Value {
	once_comonadEnvT.Do(func() {
		cache_comonadEnvT = gopurs_runtime.Func(func(dictComonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_comonadEnvT(dictComonad_0_box)
})
	})
	return cache_comonadEnvT
}

func Call_EnvT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_withEnvT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1})}
}

func Call_runEnvT(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_mapEnvT(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)})}
}

func Call_functorEnvT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
}))
}

func Call_functorWithIndexEnvT(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorEnvT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
}))
_ = functorEnvT1_2_1
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_1
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
}))
}

func Call_foldableEnvT(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1)
_ = foldMap1_2_0
return gopurs_runtime.Func2(func(fn_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_2_0, fn_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
}), gopurs_runtime.Func3(func(fn_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), fn_1, a_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(fn_1 gopurs_runtime.Value, a_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), fn_1, a_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
}))
}

func Call_foldableWithIndexEnvT(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_0
foldableEnvT1_2_1 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "foldMap"), dictMonoid_2)
_ = foldMap1_3_2
return gopurs_runtime.Func2(func(fn_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_3_2, fn_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
})
}), gopurs_runtime.Func3(func(fn_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldl"), fn_2, a_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(fn_2 gopurs_runtime.Value, a_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldr"), fn_2, a_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
}))
_ = foldableEnvT1_2_1
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_2_1
}), gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), dictMonoid_3)
_ = foldMapWithIndex1_4_3
return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMapWithIndex1_4_3, f_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
})
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), f_3, a_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), f_3, a_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
}))
}

func Call_traversableEnvT(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorEnvT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
}))
_ = functorEnvT1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_3_2
foldableEnvT1_4_3 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "foldMap"), dictMonoid_4)
_ = foldMap1_5_4
return gopurs_runtime.Func2(func(fn_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(foldMap1_5_4, fn_6, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1)
})
}), gopurs_runtime.Func3(func(fn_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldl"), fn_4, a_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(fn_4 gopurs_runtime.Value, a_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldr"), fn_4, a_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
}))
_ = foldableEnvT1_4_3
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_4_3
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_1
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
sequence1_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_5)
_ = sequence1_6_5
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorFn(), "map"), Get_EnvT(), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0)), gopurs_runtime.Apply(sequence1_6_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_6_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_5)
_ = traverse1_6_6
return gopurs_runtime.Func2(func(f_7 gopurs_runtime.Value, v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorFn(), "map"), Get_EnvT(), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)), gopurs_runtime.Apply2(traverse1_6_6, f_7, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
})
}))
}

func Call_traversableWithIndexEnvT(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorEnvT1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
}))
_ = functorEnvT1_3_3
functorWithIndexEnvT1_3_2 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_3_3
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mapWithIndex"), f_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)})}
}))
_ = functorWithIndexEnvT1_3_2
foldableWithIndexEnvT1_4_4 := Call_foldableWithIndexEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexEnvT1_4_4
traversableEnvT1_5_5 := Call_traversableEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableEnvT1_5_5
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexEnvT1_4_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexEnvT1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableEnvT1_5_5
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
traverseWithIndex1_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_6)
_ = traverseWithIndex1_7_6
return gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorFn(), "map"), Get_EnvT(), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0)), gopurs_runtime.Apply2(traverseWithIndex1_7_6, f_8, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1))
})
}))
}

func Call_extendEnvT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
Functor0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = Functor0_1_0
functorEnvT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
}))
_ = functorEnvT1_2_1
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_1
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Functor0_1_0, "map"), f_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))})}
}))
}

func Call_comonadEnvT(dictComonad_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictComonad_0 gopurs_runtime.Value = dictComonad_0_loop
_ = dictComonad_0
extendEnvT1_1_0 := Call_extendEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "Extend0"), gopurs_runtime.Value{}))
_ = extendEnvT1_1_0
return gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return extendEnvT1_1_0
}), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictComonad_0, "extract"), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
}))
}



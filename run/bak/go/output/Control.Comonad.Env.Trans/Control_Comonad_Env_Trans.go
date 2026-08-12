package Control_Comonad_Env_Trans

import (
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Comonad "gopurs/output/Control.Comonad"
	pkg_Control_Extend "gopurs/output/Control.Extend"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_TraversableWithIndex "gopurs/output/Data.TraversableWithIndex"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_withEnvT(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)))}
})
	})
	return cache_withEnvT
}

var cache_runEnvT gopurs_runtime.Value
var once_runEnvT sync.Once
func Get_runEnvT() gopurs_runtime.Value {
	once_runEnvT.Do(func() {
		cache_runEnvT = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_runEnvT(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box)))}
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
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_mapEnvT(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box)))}
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
		cache_comonadTransEnvT = gopurs_runtime.RecordDict1("lower", gopurs_runtime.Func(func(dictComonad_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1
})
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

var cache_extract__1031647521 gopurs_runtime.Value
var once_extract__1031647521 sync.Once
func Get_extract__1031647521() gopurs_runtime.Value {
	once_extract__1031647521.Do(func() {
		cache_extract__1031647521 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extract__1031647521(gopurs_runtime.CoerceToStruct[pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extract__1031647521
}

var cache_extend__1264481661 gopurs_runtime.Value
var once_extend__1264481661 sync.Once
func Get_extend__1264481661() gopurs_runtime.Value {
	once_extend__1264481661.Do(func() {
		cache_extend__1264481661 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend__1264481661(gopurs_runtime.CoerceToStruct[pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extend__1264481661
}

var cache_extend__267444733 gopurs_runtime.Value
var once_extend__267444733 sync.Once
func Get_extend__267444733() gopurs_runtime.Value {
	once_extend__267444733.Do(func() {
		cache_extend__267444733 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_extend__267444733(gopurs_runtime.CoerceToStruct[pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_extend__267444733
}

var cache_compose__858342840 gopurs_runtime.Value
var once_compose__858342840 sync.Once
func Get_compose__858342840() gopurs_runtime.Value {
	once_compose__858342840.Do(func() {
		cache_compose__858342840 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compose__858342840(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compose__858342840
}

var cache_composeFlipped__2583068543 gopurs_runtime.Value
var once_composeFlipped__2583068543 sync.Once
func Get_composeFlipped__2583068543() gopurs_runtime.Value {
	once_composeFlipped__2583068543.Do(func() {
		cache_composeFlipped__2583068543 = gopurs_runtime.Func3(func(dictSemigroupoid_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, g_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_composeFlipped__2583068543(gopurs_runtime.CoerceToStruct[pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]](dictSemigroupoid_0_box), f_1_box, g_2_box)
})
	})
	return cache_composeFlipped__2583068543
}

var cache_semigroupoidFn__2387483462 gopurs_runtime.Value
var once_semigroupoidFn__2387483462 sync.Once
func Get_semigroupoidFn__2387483462() gopurs_runtime.Value {
	once_semigroupoidFn__2387483462.Do(func() {
		cache_semigroupoidFn__2387483462 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__2387483462
}

var cache_foldMap__4098395794 gopurs_runtime.Value
var once_foldMap__4098395794 sync.Once
func Get_foldMap__4098395794() gopurs_runtime.Value {
	once_foldMap__4098395794.Do(func() {
		cache_foldMap__4098395794 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap__4098395794(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap__4098395794
}

var cache_foldl__2151204251 gopurs_runtime.Value
var once_foldl__2151204251 sync.Once
func Get_foldl__2151204251() gopurs_runtime.Value {
	once_foldl__2151204251.Do(func() {
		cache_foldl__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl__2151204251
}

var cache_foldr__2151204251 gopurs_runtime.Value
var once_foldr__2151204251 sync.Once
func Get_foldr__2151204251() gopurs_runtime.Value {
	once_foldr__2151204251.Do(func() {
		cache_foldr__2151204251 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__2151204251(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__2151204251
}

var cache_foldMapWithIndex__2292551140 gopurs_runtime.Value
var once_foldMapWithIndex__2292551140 sync.Once
func Get_foldMapWithIndex__2292551140() gopurs_runtime.Value {
	once_foldMapWithIndex__2292551140.Do(func() {
		cache_foldMapWithIndex__2292551140 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMapWithIndex__2292551140(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMapWithIndex__2292551140
}

var cache_foldlWithIndex__2986161357 gopurs_runtime.Value
var once_foldlWithIndex__2986161357 sync.Once
func Get_foldlWithIndex__2986161357() gopurs_runtime.Value {
	once_foldlWithIndex__2986161357.Do(func() {
		cache_foldlWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldlWithIndex__2986161357(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldlWithIndex__2986161357
}

var cache_foldrWithIndex__2986161357 gopurs_runtime.Value
var once_foldrWithIndex__2986161357 sync.Once
func Get_foldrWithIndex__2986161357() gopurs_runtime.Value {
	once_foldrWithIndex__2986161357.Do(func() {
		cache_foldrWithIndex__2986161357 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldrWithIndex__2986161357(gopurs_runtime.CoerceToStruct[pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldrWithIndex__2986161357
}

var cache_functorFn__2451697751 gopurs_runtime.Value
var once_functorFn__2451697751 sync.Once
func Get_functorFn__2451697751() gopurs_runtime.Value {
	once_functorFn__2451697751.Do(func() {
		cache_functorFn__2451697751 = gopurs_runtime.RecordDict1("map", gopurs_runtime.RecordGet(pkg_Control_Semigroupoid.Get_semigroupoidFn(), "compose"))
	})
	return cache_functorFn__2451697751
}

var cache_map__2199395572 gopurs_runtime.Value
var once_map__2199395572 sync.Once
func Get_map__2199395572() gopurs_runtime.Value {
	once_map__2199395572.Do(func() {
		cache_map__2199395572 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2199395572(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2199395572
}

var cache_map__3436901780 gopurs_runtime.Value
var once_map__3436901780 sync.Once
func Get_map__3436901780() gopurs_runtime.Value {
	once_map__3436901780.Do(func() {
		cache_map__3436901780 = gopurs_runtime.Func3(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value, x_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3436901780(__eta0_0_box, __eta1_1_box, x_2_box)
})
	})
	return cache_map__3436901780
}

var cache_map__2549863700 gopurs_runtime.Value
var once_map__2549863700 sync.Once
func Get_map__2549863700() gopurs_runtime.Value {
	once_map__2549863700.Do(func() {
		cache_map__2549863700 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2549863700(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2549863700
}

var cache_map__2198784724 gopurs_runtime.Value
var once_map__2198784724 sync.Once
func Get_map__2198784724() gopurs_runtime.Value {
	once_map__2198784724.Do(func() {
		cache_map__2198784724 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__2198784724(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__2198784724
}

var cache_mapWithIndex__55256674 gopurs_runtime.Value
var once_mapWithIndex__55256674 sync.Once
func Get_mapWithIndex__55256674() gopurs_runtime.Value {
	once_mapWithIndex__55256674.Do(func() {
		cache_mapWithIndex__55256674 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mapWithIndex__55256674(gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mapWithIndex__55256674
}

var cache_sequence__1886310617 gopurs_runtime.Value
var once_sequence__1886310617 sync.Once
func Get_sequence__1886310617() gopurs_runtime.Value {
	once_sequence__1886310617.Do(func() {
		cache_sequence__1886310617 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence__1886310617(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sequence__1886310617
}

var cache_traverse__314957093 gopurs_runtime.Value
var once_traverse__314957093 sync.Once
func Get_traverse__314957093() gopurs_runtime.Value {
	once_traverse__314957093.Do(func() {
		cache_traverse__314957093 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse__314957093(gopurs_runtime.CoerceToStruct[pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverse__314957093
}

var cache_traverseWithIndex__2726076659 gopurs_runtime.Value
var once_traverseWithIndex__2726076659 sync.Once
func Get_traverseWithIndex__2726076659() gopurs_runtime.Value {
	once_traverseWithIndex__2726076659.Do(func() {
		cache_traverseWithIndex__2726076659 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverseWithIndex__2726076659(gopurs_runtime.CoerceToStruct[pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_traverseWithIndex__2726076659
}

func Call_EnvT(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_withEnvT(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1})})
}

func Call_runEnvT(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return v_0
}

func Call_mapEnvT(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)})})
}

func Call_functorEnvT(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)})}
})
}))
}

func Call_functorWithIndexEnvT(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorEnvT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_1_0
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
})
}))
}

func Call_foldableEnvT(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fn_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, fn_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), fn_1, a_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), fn_1, a_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
})
}))
}

func Call_foldableWithIndexEnvT(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_1
foldableEnvT1_1_0 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fn_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, fn_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldl"), fn_2, a_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_1, "foldr"), fn_2, a_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}))
_ = foldableEnvT1_1_0
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_1_0
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), f_2, a_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), f_2, a_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}))
}

func Call_traversableEnvT(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorEnvT1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_1_0
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_2_3
foldableEnvT1_2_2 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(fn_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_3))}, fn_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldl"), fn_3, a_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(fn_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_2_3, "foldr"), fn_3, a_4, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
})
})
}))
_ = foldableEnvT1_2_2
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableEnvT1_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_4
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_4.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorFn(), "map"), Get_EnvT(), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0)), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_4_5 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_4_5
return gopurs_runtime.Func(func(f_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_4_5.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorFn(), "map"), Get_EnvT(), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_5, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
})
})
}))
}

func Call_traversableWithIndexEnvT(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorEnvT1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_2_2
functorWithIndexEnvT1_1_0 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_2
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mapWithIndex"), f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
})
}))
_ = functorWithIndexEnvT1_1_0
foldableWithIndexEnvT1_2_4 := Call_foldableWithIndexEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexEnvT1_2_4
traversableEnvT1_3_5 := Call_traversableEnvT(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableEnvT1_3_5
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexEnvT1_2_4
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndexEnvT1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableEnvT1_3_5
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_5_6 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_6
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_5_6.V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorFn(), "map"), Get_EnvT(), gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_4))}, f_6, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
})
})
}))
}

func Call_extendEnvT(dictExtend_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictExtend_0 gopurs_runtime.Value = dictExtend_0_loop
_ = dictExtend_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictExtend_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_2
functorEnvT1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_2, "map"), f_3, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
})
}))
_ = functorEnvT1_2_1
return gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorEnvT1_2_1
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(pkg_Data_Tuple.Get_Tuple(), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0)
_ = __local_var_5_3
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply2(Functor0_1_0.V0, f_3, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictExtend_0, "extend"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_5_3, x_6)
}), (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))})}
})
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

func Call_extract__1031647521(dict_0_loop *pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Comonad.Constructor_Comonad[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_extend__1264481661(dict_0_loop *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_extend__267444733(dict_0_loop *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Extend.Constructor_Extend[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_composeFlipped__2583068543(dictSemigroupoid_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, g_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroupoid_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dictSemigroupoid_0_loop
_ = dictSemigroupoid_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var g_2 gopurs_runtime.Value = g_2_loop
_ = g_2
return gopurs_runtime.Apply2(dictSemigroupoid_0.V0, g_2, f_1)
}

func Call_foldMap__4098395794(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_foldl__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldr__2151204251(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldMapWithIndex__2292551140(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldlWithIndex__2986161357(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldrWithIndex__2986161357(dict_0_loop *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FoldableWithIndex.Constructor_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__3436901780(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value, x_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
var x_2 gopurs_runtime.Value = x_2_loop
_ = x_2
return gopurs_runtime.Apply(__eta0_0, gopurs_runtime.Apply(__eta1_1, x_2))
}

func Call_map__2549863700(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__2198784724(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mapWithIndex__55256674(dict_0_loop *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sequence__1886310617(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_traverse__314957093(dict_0_loop *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Traversable.Constructor_Traversable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_traverseWithIndex__2726076659(dict_0_loop *pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_TraversableWithIndex.Constructor_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}



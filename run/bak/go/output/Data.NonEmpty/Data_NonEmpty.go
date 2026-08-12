package Data_NonEmpty

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Alternative "gopurs/output/Control.Alternative"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Plus "gopurs/output/Control.Plus"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Monoid "gopurs/output/Data.Monoid"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_TraversableWithIndex "gopurs/output/Data.TraversableWithIndex"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable "gopurs/output/Data.Unfoldable"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_NonEmpty gopurs_runtime.Value
var once_NonEmpty sync.Once
func Get_NonEmpty() gopurs_runtime.Value {
	once_NonEmpty.Do(func() {
		cache_NonEmpty = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_NonEmpty
}

var cache_unfoldable1NonEmpty gopurs_runtime.Value
var once_unfoldable1NonEmpty sync.Once
func Get_unfoldable1NonEmpty() gopurs_runtime.Value {
	once_unfoldable1NonEmpty.Do(func() {
		cache_unfoldable1NonEmpty = gopurs_runtime.Func(func(dictUnfoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldable1NonEmpty(dictUnfoldable_0_box)
})
	})
	return cache_unfoldable1NonEmpty
}

var cache_tail gopurs_runtime.Value
var once_tail sync.Once
func Get_tail() gopurs_runtime.Value {
	once_tail.Do(func() {
		cache_tail = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tail(gopurs_runtime.CoerceToStruct[Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_tail
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(gopurs_runtime.CoerceToStruct[pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]](dictPlus_0_box))
})
	})
	return cache_singleton
}

var cache_showNonEmpty gopurs_runtime.Value
var once_showNonEmpty sync.Once
func Get_showNonEmpty() gopurs_runtime.Value {
	once_showNonEmpty.Do(func() {
		cache_showNonEmpty = gopurs_runtime.Func2(func(dictShow_0_box gopurs_runtime.Value, dictShow1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showNonEmpty(dictShow_0_box, dictShow1_1_box)
})
	})
	return cache_showNonEmpty
}

var cache_semigroupNonEmpty gopurs_runtime.Value
var once_semigroupNonEmpty sync.Once
func Get_semigroupNonEmpty() gopurs_runtime.Value {
	once_semigroupNonEmpty.Do(func() {
		cache_semigroupNonEmpty = gopurs_runtime.Func2(func(dictApplicative_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupNonEmpty(dictApplicative_0_box, dictSemigroup_1_box)
})
	})
	return cache_semigroupNonEmpty
}

var cache_oneOf gopurs_runtime.Value
var once_oneOf sync.Once
func Get_oneOf() gopurs_runtime.Value {
	once_oneOf.Do(func() {
		cache_oneOf = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOf(gopurs_runtime.CoerceToStruct[pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_oneOf
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(gopurs_runtime.CoerceToStruct[Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_head
}

var cache_functorNonEmpty gopurs_runtime.Value
var once_functorNonEmpty sync.Once
func Get_functorNonEmpty() gopurs_runtime.Value {
	once_functorNonEmpty.Do(func() {
		cache_functorNonEmpty = gopurs_runtime.Func(func(dictFunctor_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorNonEmpty(dictFunctor_0_box)
})
	})
	return cache_functorNonEmpty
}

var cache_functorWithIndex gopurs_runtime.Value
var once_functorWithIndex sync.Once
func Get_functorWithIndex() gopurs_runtime.Value {
	once_functorWithIndex.Do(func() {
		cache_functorWithIndex = gopurs_runtime.Func(func(dictFunctorWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_functorWithIndex(dictFunctorWithIndex_0_box)
})
	})
	return cache_functorWithIndex
}

var cache_fromNonEmpty gopurs_runtime.Value
var once_fromNonEmpty sync.Once
func Get_fromNonEmpty() gopurs_runtime.Value {
	once_fromNonEmpty.Do(func() {
		cache_fromNonEmpty = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromNonEmpty(f_0_box, gopurs_runtime.CoerceToStruct[Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromNonEmpty
}

var cache_foldableNonEmpty gopurs_runtime.Value
var once_foldableNonEmpty sync.Once
func Get_foldableNonEmpty() gopurs_runtime.Value {
	once_foldableNonEmpty.Do(func() {
		cache_foldableNonEmpty = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableNonEmpty(dictFoldable_0_box)
})
	})
	return cache_foldableNonEmpty
}

var cache_foldableWithIndexNonEmpty gopurs_runtime.Value
var once_foldableWithIndexNonEmpty sync.Once
func Get_foldableWithIndexNonEmpty() gopurs_runtime.Value {
	once_foldableWithIndexNonEmpty.Do(func() {
		cache_foldableWithIndexNonEmpty = gopurs_runtime.Func(func(dictFoldableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldableWithIndexNonEmpty(dictFoldableWithIndex_0_box)
})
	})
	return cache_foldableWithIndexNonEmpty
}

var cache_traversableNonEmpty gopurs_runtime.Value
var once_traversableNonEmpty sync.Once
func Get_traversableNonEmpty() gopurs_runtime.Value {
	once_traversableNonEmpty.Do(func() {
		cache_traversableNonEmpty = gopurs_runtime.Func(func(dictTraversable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableNonEmpty(dictTraversable_0_box)
})
	})
	return cache_traversableNonEmpty
}

var cache_traversableWithIndexNonEmpty gopurs_runtime.Value
var once_traversableWithIndexNonEmpty sync.Once
func Get_traversableWithIndexNonEmpty() gopurs_runtime.Value {
	once_traversableWithIndexNonEmpty.Do(func() {
		cache_traversableWithIndexNonEmpty = gopurs_runtime.Func(func(dictTraversableWithIndex_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traversableWithIndexNonEmpty(dictTraversableWithIndex_0_box)
})
	})
	return cache_traversableWithIndexNonEmpty
}

var cache_foldable1NonEmpty gopurs_runtime.Value
var once_foldable1NonEmpty sync.Once
func Get_foldable1NonEmpty() gopurs_runtime.Value {
	once_foldable1NonEmpty.Do(func() {
		cache_foldable1NonEmpty = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldable1NonEmpty(dictFoldable_0_box)
})
	})
	return cache_foldable1NonEmpty
}

var cache_foldl1 gopurs_runtime.Value
var once_foldl1 sync.Once
func Get_foldl1() gopurs_runtime.Value {
	once_foldl1.Do(func() {
		cache_foldl1 = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_foldl1
}

var cache_eqNonEmpty gopurs_runtime.Value
var once_eqNonEmpty sync.Once
func Get_eqNonEmpty() gopurs_runtime.Value {
	once_eqNonEmpty.Do(func() {
		cache_eqNonEmpty = gopurs_runtime.Func2(func(dictEq1_0_box gopurs_runtime.Value, dictEq_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqNonEmpty(dictEq1_0_box, dictEq_1_box)
})
	})
	return cache_eqNonEmpty
}

var cache_ordNonEmpty gopurs_runtime.Value
var once_ordNonEmpty sync.Once
func Get_ordNonEmpty() gopurs_runtime.Value {
	once_ordNonEmpty.Do(func() {
		cache_ordNonEmpty = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordNonEmpty(dictOrd1_0_box)
})
	})
	return cache_ordNonEmpty
}

var cache_eq1NonEmpty gopurs_runtime.Value
var once_eq1NonEmpty sync.Once
func Get_eq1NonEmpty() gopurs_runtime.Value {
	once_eq1NonEmpty.Do(func() {
		cache_eq1NonEmpty = gopurs_runtime.Func(func(dictEq1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1NonEmpty(dictEq1_0_box)
})
	})
	return cache_eq1NonEmpty
}

var cache_ord1NonEmpty gopurs_runtime.Value
var once_ord1NonEmpty sync.Once
func Get_ord1NonEmpty() gopurs_runtime.Value {
	once_ord1NonEmpty.Do(func() {
		cache_ord1NonEmpty = gopurs_runtime.Func(func(dictOrd1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ord1NonEmpty(dictOrd1_0_box)
})
	})
	return cache_ord1NonEmpty
}

var cache_alt__267341625 gopurs_runtime.Value
var once_alt__267341625 sync.Once
func Get_alt__267341625() gopurs_runtime.Value {
	once_alt__267341625.Do(func() {
		cache_alt__267341625 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alt__267341625(gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_alt__267341625
}

var cache_pure__3215807376 gopurs_runtime.Value
var once_pure__3215807376 sync.Once
func Get_pure__3215807376() gopurs_runtime.Value {
	once_pure__3215807376.Do(func() {
		cache_pure__3215807376 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_pure__3215807376(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_pure__3215807376
}

var cache_apply__353515660 gopurs_runtime.Value
var once_apply__353515660 sync.Once
func Get_apply__353515660() gopurs_runtime.Value {
	once_apply__353515660.Do(func() {
		cache_apply__353515660 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__353515660(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__353515660
}

var cache_apply__2285104812 gopurs_runtime.Value
var once_apply__2285104812 sync.Once
func Get_apply__2285104812() gopurs_runtime.Value {
	once_apply__2285104812.Do(func() {
		cache_apply__2285104812 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_apply__2285104812(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_apply__2285104812
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

var cache_eq__2384498378 gopurs_runtime.Value
var once_eq__2384498378 sync.Once
func Get_eq__2384498378() gopurs_runtime.Value {
	once_eq__2384498378.Do(func() {
		cache_eq__2384498378 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq__2384498378(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq__2384498378
}

var cache_eq1__1773593252 gopurs_runtime.Value
var once_eq1__1773593252 sync.Once
func Get_eq1__1773593252() gopurs_runtime.Value {
	once_eq1__1773593252.Do(func() {
		cache_eq1__1773593252 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eq1__1773593252(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_eq1__1773593252
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

var cache_foldr__4105571355 gopurs_runtime.Value
var once_foldr__4105571355 sync.Once
func Get_foldr__4105571355() gopurs_runtime.Value {
	once_foldr__4105571355.Do(func() {
		cache_foldr__4105571355 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr__4105571355(gopurs_runtime.CoerceToStruct[pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr__4105571355
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

var cache_map__3098878004 gopurs_runtime.Value
var once_map__3098878004 sync.Once
func Get_map__3098878004() gopurs_runtime.Value {
	once_map__3098878004.Do(func() {
		cache_map__3098878004 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__3098878004(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__3098878004
}

var cache_map__1208952924 gopurs_runtime.Value
var once_map__1208952924 sync.Once
func Get_map__1208952924() gopurs_runtime.Value {
	once_map__1208952924.Do(func() {
		cache_map__1208952924 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_map__1208952924(v_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v1_1_box)))}
})
	})
	return cache_map__1208952924
}

var cache_map__2311960860 gopurs_runtime.Value
var once_map__2311960860 sync.Once
func Get_map__2311960860() gopurs_runtime.Value {
	once_map__2311960860.Do(func() {
		cache_map__2311960860 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, m_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_map__2311960860(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](m_1_box)))}
})
	})
	return cache_map__2311960860
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

var cache_conj__3676519832 gopurs_runtime.Value
var once_conj__3676519832 sync.Once
func Get_conj__3676519832() gopurs_runtime.Value {
	once_conj__3676519832.Do(func() {
		cache_conj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_conj__3676519832
}

var cache_conj__3472268504 gopurs_runtime.Value
var once_conj__3472268504 sync.Once
func Get_conj__3472268504() gopurs_runtime.Value {
	once_conj__3472268504.Do(func() {
		cache_conj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_conj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_conj__3472268504
}

var cache_disj__3676519832 gopurs_runtime.Value
var once_disj__3676519832 sync.Once
func Get_disj__3676519832() gopurs_runtime.Value {
	once_disj__3676519832.Do(func() {
		cache_disj__3676519832 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3676519832(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_disj__3676519832
}

var cache_disj__3472268504 gopurs_runtime.Value
var once_disj__3472268504 sync.Once
func Get_disj__3472268504() gopurs_runtime.Value {
	once_disj__3472268504.Do(func() {
		cache_disj__3472268504 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_disj__3472268504(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_disj__3472268504
}

var cache_not__3201284355 gopurs_runtime.Value
var once_not__3201284355 sync.Once
func Get_not__3201284355() gopurs_runtime.Value {
	once_not__3201284355.Do(func() {
		cache_not__3201284355 = gopurs_runtime.Func(func(__eta0_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__3201284355(__eta0_0_box)
})
	})
	return cache_not__3201284355
}

var cache_not__1505204753 gopurs_runtime.Value
var once_not__1505204753 sync.Once
func Get_not__1505204753() gopurs_runtime.Value {
	once_not__1505204753.Do(func() {
		cache_not__1505204753 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_not__1505204753(gopurs_runtime.CoerceToStruct[pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_not__1505204753
}

var cache_functorMaybe__2569569018 gopurs_runtime.Value
var once_functorMaybe__2569569018 sync.Once
func Get_functorMaybe__2569569018() gopurs_runtime.Value {
	once_functorMaybe__2569569018.Do(func() {
		cache_functorMaybe__2569569018 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v1_1.Type == 9 && v1_1.IntVal == 930809136 && v1_1.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v1_1.UnsafePtr).V0)})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](__t0))}
})
}))
	})
	return cache_functorMaybe__2569569018
}

var cache_maybe__3658316244 gopurs_runtime.Value
var once_maybe__3658316244 sync.Once
func Get_maybe__3658316244() gopurs_runtime.Value {
	once_maybe__3658316244.Do(func() {
		cache_maybe__3658316244 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__3658316244(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__3658316244
}

var cache_compare__821463600 gopurs_runtime.Value
var once_compare__821463600 sync.Once
func Get_compare__821463600() gopurs_runtime.Value {
	once_compare__821463600.Do(func() {
		cache_compare__821463600 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare__821463600(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare__821463600
}

var cache_compare1__650153534 gopurs_runtime.Value
var once_compare1__650153534 sync.Once
func Get_compare1__650153534() gopurs_runtime.Value {
	once_compare1__650153534.Do(func() {
		cache_compare1__650153534 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_compare1__650153534(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_compare1__650153534
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.Func2(func(__eta0_0_box gopurs_runtime.Value, __eta1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__493084344(__eta0_0_box, __eta1_1_box)
})
	})
	return cache_append__493084344
}

var cache_append__1230318264 gopurs_runtime.Value
var once_append__1230318264 sync.Once
func Get_append__1230318264() gopurs_runtime.Value {
	once_append__1230318264.Do(func() {
		cache_append__1230318264 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__1230318264(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__1230318264
}

var cache_append__4016378200 gopurs_runtime.Value
var once_append__4016378200 sync.Once
func Get_append__4016378200() gopurs_runtime.Value {
	once_append__4016378200.Do(func() {
		cache_append__4016378200 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_append__4016378200(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_append__4016378200
}

var cache_show__2742601362 gopurs_runtime.Value
var once_show__2742601362 sync.Once
func Get_show__2742601362() gopurs_runtime.Value {
	once_show__2742601362.Do(func() {
		cache_show__2742601362 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__2742601362(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__2742601362
}

var cache_show__3978978930 gopurs_runtime.Value
var once_show__3978978930 sync.Once
func Get_show__3978978930() gopurs_runtime.Value {
	once_show__3978978930.Do(func() {
		cache_show__3978978930 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__3978978930(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__3978978930
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

var cache_functorTuple__2544689875 gopurs_runtime.Value
var once_functorTuple__2544689875 sync.Once
func Get_functorTuple__2544689875() gopurs_runtime.Value {
	once_functorTuple__2544689875.Do(func() {
		cache_functorTuple__2544689875 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(m_1.UnsafePtr).V1)})}
})
}))
	})
	return cache_functorTuple__2544689875
}

var cache_uncurry__3533477633 gopurs_runtime.Value
var once_uncurry__3533477633 sync.Once
func Get_uncurry__3533477633() gopurs_runtime.Value {
	once_uncurry__3533477633.Do(func() {
		cache_uncurry__3533477633 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncurry__3533477633(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_uncurry__3533477633
}

var cache_uncurry__601611969 gopurs_runtime.Value
var once_uncurry__601611969 sync.Once
func Get_uncurry__601611969() gopurs_runtime.Value {
	once_uncurry__601611969.Do(func() {
		cache_uncurry__601611969 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_uncurry__601611969(f_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_uncurry__601611969
}

var cache_unfoldr__1128708256 gopurs_runtime.Value
var once_unfoldr__1128708256 sync.Once
func Get_unfoldr__1128708256() gopurs_runtime.Value {
	once_unfoldr__1128708256.Do(func() {
		cache_unfoldr__1128708256 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__1128708256(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__1128708256
}

var cache_unfoldr__3132297377 gopurs_runtime.Value
var once_unfoldr__3132297377 sync.Once
func Get_unfoldr__3132297377() gopurs_runtime.Value {
	once_unfoldr__3132297377.Do(func() {
		cache_unfoldr__3132297377 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_unfoldr__3132297377(gopurs_runtime.CoerceToStruct[pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_unfoldr__3132297377
}

type Constructor_NonEmpty[T_f any, T_a any] struct {
	Rc uint32
	V0 T_a
	V1 gopurs_runtime.Value
}


func Call_unfoldable1NonEmpty(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), f_1)), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(f_1, b_2)))}))
_ = __local_var_3_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(__local_var_3_0)}.UnsafePtr).V1})}
})
}))
}

func Call_tail(v_0_loop *Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Call_singleton(dictPlus_0_loop *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictPlus_0 *pkg_Control_Plus.Constructor_Plus[gopurs_runtime.Value] = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := dictPlus_0.V1
_ = empty_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, a_2, empty_1_0})}
})
}

func Call_showNonEmpty(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str("(NonEmpty "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0).StrVal()), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(" "), gopurs_runtime.Str(Call_append__493084344(gopurs_runtime.Str(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())).StrVal())).StrVal())
}))
}

func Call_semigroupNonEmpty(dictApplicative_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1))})}
})
}))
}

func Call_oneOf(dictAlternative_0_loop *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *pkg_Control_Alternative.Constructor_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
Alt0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictAlternative_0.V1, gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(dictAlternative_0.V0, gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Alt0_1_0.V1, gopurs_runtime.Apply(Applicative0_2_1.V1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}

func Call_head(v_0_loop *Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_functorNonEmpty(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)})}
})
}))
}

func Call_functorWithIndex(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorNonEmpty1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_1_0
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_1_0
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, x_4})})
}), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)})}
})
}))
}

func Call_fromNonEmpty(f_0_loop gopurs_runtime.Value, v_1_loop *Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}

func Call_foldableNonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_2_0
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_2_0.V0, gopurs_runtime.Apply(f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_1))}, f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, b_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
})
})
}))
}

func Call_foldableWithIndexNonEmpty(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
foldableNonEmpty1_1_0 := Call_foldableNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{}))
_ = foldableNonEmpty1_1_0
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_1_0
}), gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_3_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_3_1.V0, gopurs_runtime.Apply2(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Monoid.Constructor_Monoid[gopurs_runtime.Value]](dictMonoid_2))}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, x_6})})
}), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, x_5})})
}), gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, b_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_2, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, x_5})})
}), b_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
})
})
}))
}

func Call_traversableNonEmpty(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_1
functorNonEmpty1_1_0 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "map"), f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_1_0
foldableNonEmpty1_2_2 := Call_foldableNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{}))
_ = foldableNonEmpty1_2_2
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_2_2
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_1_0
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_4_3 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_3
Functor0_5_4 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_4
return gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_4_3.V1, gopurs_runtime.Apply2(Functor0_5_4.V0, Get_NonEmpty(), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_3 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_4_5 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_4_5
Functor0_5_6 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_3, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_5_6
return gopurs_runtime.Func(func(f_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_4_5.V1, gopurs_runtime.Apply2(Functor0_5_6.V0, Get_NonEmpty(), gopurs_runtime.Apply(f_6, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_3))}, f_6, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
})
})
}))
}

func Call_traversableWithIndexNonEmpty(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
__local_var_1_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_1
__local_var_2_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_1, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_3
functorNonEmpty1_2_2 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply(f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_3, "map"), f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)})}
})
}))
_ = functorNonEmpty1_2_2
functorWithIndex1_1_0 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_2_2
}), gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Apply2(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_1, "mapWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, x_5})})
}), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
})
}))
_ = functorWithIndex1_1_0
foldableWithIndexNonEmpty1_2_4 := Call_foldableWithIndexNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexNonEmpty1_2_4
traversableNonEmpty1_3_5 := Call_traversableNonEmpty(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableNonEmpty1_3_5
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexNonEmpty1_2_4
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndex1_1_0
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableNonEmpty1_3_5
}), gopurs_runtime.Func(func(dictApplicative_4 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_5_6 := gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_5_6
Functor0_6_7 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_4, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_6_7
return gopurs_runtime.Func(func(f_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Apply0_5_6.V1, gopurs_runtime.Apply2(Functor0_6_7.V0, Get_NonEmpty(), gopurs_runtime.Apply2(f_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0)), gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_4))}, gopurs_runtime.Func(func(x_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_7, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, x_9})})
}), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
})
})
}))
}

func Call_foldable1NonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
foldableNonEmpty1_1_0 := Call_foldableNonEmpty(dictFoldable_0)
_ = foldableNonEmpty1_1_0
return gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_1_0
}), gopurs_runtime.Func(func(dictSemigroup_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func(func(s_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), s_5, gopurs_runtime.Apply(f_3, a1_6))
})
}), gopurs_runtime.Apply(f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}), gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_1 := gopurs_runtime.Apply(f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)
_ = __local_var_4_1
__local_var_5_2 := gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply(f_2, a1_5)
_ = __local_var_6_4
__local_var_6_3 := gopurs_runtime.Func(func(v2_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr == nil) {
__t5 = a1_5
goto end_branch_5
} else {

}
}
{
if (v2_7.Type == 9 && v2_7.IntVal == 930809136 && v2_7.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(__local_var_6_4, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v2_7.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})
_ = __local_var_6_3
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Apply(__local_var_6_3, x_7)})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
_ = __local_var_5_2
var __t6 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.UnsafePtr == nil) {
__t6 = (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
goto end_branch_6
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.UnsafePtr != nil) {
__t6 = gopurs_runtime.Apply(__local_var_4_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__local_var_5_2)}.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
})
}))
}

func Call_foldl1(dictFoldable_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordGet(Call_foldable1NonEmpty(gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}), "foldl1")
}

func Call_eqNonEmpty(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
_ = dictEq_1
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((Call_conj__3676519832(gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V0).IntVal) != (0)), gopurs_runtime.Bool((gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]](dictEq_1))}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal) != (0))).IntVal) != (0))
})
}))
}

func Call_ordNonEmpty(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
eqNonEmpty1_1_0 := gopurs_runtime.Apply(Get_eqNonEmpty(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
eqNonEmpty2_3_1 := gopurs_runtime.Apply(eqNonEmpty1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{}))
_ = eqNonEmpty2_3_1
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eqNonEmpty2_3_1
}), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
v_6_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_2, "compare"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V0)
_ = v_6_2
var __t3 uint32
{
if (uint32(v_6_2.IntVal) == 1527465420) {
__t3 = 1527465420
goto end_branch_3
} else {

}
}
{
if (uint32(v_6_2.IntVal) == 380165415) {
__t3 = 380165415
goto end_branch_3
} else {

}
}
{
__t3 = uint32(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_2))}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_4.UnsafePtr).V1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_5.UnsafePtr).V1).IntVal)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: int64(__t3), UnsafePtr: nil}
})
}))
})
}

func Call_eq1NonEmpty(dictEq1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
return gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqNonEmpty(dictEq1_0, dictEq_1), "eq")
}))
}

func Call_ord1NonEmpty(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
ordNonEmpty1_1_0 := Call_ordNonEmpty(dictOrd1_0)
_ = ordNonEmpty1_1_0
__local_var_2_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_2
eq1NonEmpty1_2_1 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqNonEmpty(__local_var_2_2, dictEq_3), "eq")
}))
_ = eq1NonEmpty1_2_1
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1NonEmpty1_2_1
}), gopurs_runtime.Func(func(dictOrd_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordNonEmpty1_1_0, dictOrd_3), "compare")
}))
}

func Call_alt__267341625(dict_0_loop *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Alt.Constructor_Alt[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_pure__3215807376(dict_0_loop *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_apply__2285104812(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_eq1__1773593252(dict_0_loop *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_foldr__4105571355(dict_0_loop *pkg_Data_Foldable.Constructor_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_map__3098878004(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1208952924(v_0_loop gopurs_runtime.Value, v1_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) *pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]] {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v1_1_loop
_ = v1_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]](gopurs_runtime.Apply(v_0, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v1_1)}.UnsafePtr).V0)))}})}
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(nil))}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]]]](__t0)
}

func Call_map__2311960860(f_0_loop gopurs_runtime.Value, m_1_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var m_1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = m_1_loop
_ = m_1
return gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(m_1)}.UnsafePtr).V0, gopurs_runtime.Apply(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(m_1)}.UnsafePtr).V1)})})
}

func Call_mapWithIndex__55256674(dict_0_loop *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_conj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) && ((__eta1_1.IntVal) != (0)))
}

func Call_conj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_disj__3676519832(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) || ((__eta1_1.IntVal) != (0)))
}

func Call_disj__3472268504(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_not__3201284355(__eta0_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
return gopurs_runtime.Bool(((__eta0_0.IntVal) != (0)) != (true))
}

func Call_not__1505204753(dict_0_loop *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_HeytingAlgebra.Constructor_HeytingAlgebra[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V4
}

func Call_maybe__3658316244(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_2_loop
_ = v2_2
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr == nil) {
__t0 = v_0
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Apply(v1_1, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_2)}.UnsafePtr).V0)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compare1__650153534(dict_0_loop *pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__493084344(__eta0_0_loop gopurs_runtime.Value, __eta1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __eta0_0 gopurs_runtime.Value = __eta0_0_loop
_ = __eta0_0
var __eta1_1 gopurs_runtime.Value = __eta1_1_loop
_ = __eta1_1
return gopurs_runtime.Str((__eta0_0.StrVal()) + (__eta1_1.StrVal()))
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_append__4016378200(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__3978978930(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
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

func Call_uncurry__3533477633(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}

func Call_uncurry__601611969(f_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V1)
}

func Call_unfoldr__1128708256(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_unfoldr__3132297377(dict_0_loop *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Unfoldable.Constructor_Unfoldable[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}



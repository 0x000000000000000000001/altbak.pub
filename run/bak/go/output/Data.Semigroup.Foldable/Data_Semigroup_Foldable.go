package Data_Semigroup_Foldable

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Function "gopurs/output/Data.Function"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Newtype "gopurs/output/Data.Newtype"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ordering "gopurs/output/Data.Ordering"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Unsafe_Coerce "gopurs/output/Unsafe.Coerce"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_identity gopurs_runtime.Value
var once_identity sync.Once
func Get_identity() gopurs_runtime.Value {
	once_identity.Do(func() {
		cache_identity = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity(x_0_box)
})
	})
	return cache_identity
}

var cache_identity1 gopurs_runtime.Value
var once_identity1 sync.Once
func Get_identity1() gopurs_runtime.Value {
	once_identity1.Do(func() {
		cache_identity1 = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity1(x_0_box)
})
	})
	return cache_identity1
}

var cache_FoldRight1 gopurs_runtime.Value
var once_FoldRight1 sync.Once
func Get_FoldRight1() gopurs_runtime.Value {
	once_FoldRight1.Do(func() {
		cache_FoldRight1 = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_FoldRight1[gopurs_runtime.Value]{1, value0, value1})}
})
})
	})
	return cache_FoldRight1
}

var cache_mkFoldRight1 gopurs_runtime.Value
var once_mkFoldRight1 sync.Once
func Get_mkFoldRight1() gopurs_runtime.Value {
	once_mkFoldRight1.Do(func() {
		cache_mkFoldRight1 = gopurs_runtime.Apply(Get_FoldRight1(), pkg_Data_Function.Get_go__const())
	})
	return cache_mkFoldRight1
}

var cache_foldr1 gopurs_runtime.Value
var once_foldr1 sync.Once
func Get_foldr1() gopurs_runtime.Value {
	once_foldr1.Do(func() {
		cache_foldr1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr1(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr1
}

var cache_foldl1 gopurs_runtime.Value
var once_foldl1 sync.Once
func Get_foldl1() gopurs_runtime.Value {
	once_foldl1.Do(func() {
		cache_foldl1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl1
}

var cache_maximumBy gopurs_runtime.Value
var once_maximumBy sync.Once
func Get_maximumBy() gopurs_runtime.Value {
	once_maximumBy.Do(func() {
		cache_maximumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximumBy(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), cmp_1_box)
})
	})
	return cache_maximumBy
}

var cache_minimumBy gopurs_runtime.Value
var once_minimumBy sync.Once
func Get_minimumBy() gopurs_runtime.Value {
	once_minimumBy.Do(func() {
		cache_minimumBy = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, cmp_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimumBy(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), cmp_1_box)
})
	})
	return cache_minimumBy
}

var cache_foldableTuple gopurs_runtime.Value
var once_foldableTuple sync.Once
func Get_foldableTuple() gopurs_runtime.Value {
	once_foldableTuple.Do(func() {
		cache_foldableTuple = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableTuple()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v1_1.UnsafePtr).V1
})
}))
	})
	return cache_foldableTuple
}

var cache_foldableMultiplicative gopurs_runtime.Value
var once_foldableMultiplicative sync.Once
func Get_foldableMultiplicative() gopurs_runtime.Value {
	once_foldableMultiplicative.Do(func() {
		cache_foldableMultiplicative = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableMultiplicative()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_foldableMultiplicative
}

var cache_foldableIdentity gopurs_runtime.Value
var once_foldableIdentity sync.Once
func Get_foldableIdentity() gopurs_runtime.Value {
	once_foldableIdentity.Do(func() {
		cache_foldableIdentity = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableIdentity()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_foldableIdentity
}

var cache_foldableDual gopurs_runtime.Value
var once_foldableDual sync.Once
func Get_foldableDual() gopurs_runtime.Value {
	once_foldableDual.Do(func() {
		cache_foldableDual = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableDual()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return v1_1
})
}))
	})
	return cache_foldableDual
}

var cache_foldRight1Semigroup gopurs_runtime.Value
var once_foldRight1Semigroup sync.Once
func Get_foldRight1Semigroup() gopurs_runtime.Value {
	once_foldRight1Semigroup.Do(func() {
		cache_foldRight1Semigroup = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := (*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_FoldRight1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, a_3, f_4)), f_4)
})
}), (*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V1})}
})
}))
	})
	return cache_foldRight1Semigroup
}

var cache_semigroupDual gopurs_runtime.Value
var once_semigroupDual sync.Once
func Get_semigroupDual() gopurs_runtime.Value {
	once_semigroupDual.Do(func() {
		cache_semigroupDual = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldRight1Semigroup(), "append"), v1_1, v_0)
})
}))
	})
	return cache_semigroupDual
}

var cache_foldMap1DefaultR gopurs_runtime.Value
var once_foldMap1DefaultR sync.Once
func Get_foldMap1DefaultR() gopurs_runtime.Value {
	once_foldMap1DefaultR.Do(func() {
		cache_foldMap1DefaultR = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1DefaultR(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_2_box))
})
	})
	return cache_foldMap1DefaultR
}

var cache_foldMap1DefaultL gopurs_runtime.Value
var once_foldMap1DefaultL sync.Once
func Get_foldMap1DefaultL() gopurs_runtime.Value {
	once_foldMap1DefaultL.Do(func() {
		cache_foldMap1DefaultL = gopurs_runtime.Func3(func(dictFoldable1_0_box gopurs_runtime.Value, dictFunctor_1_box gopurs_runtime.Value, dictSemigroup_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1DefaultL(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_2_box))
})
	})
	return cache_foldMap1DefaultL
}

var cache_foldMap1 gopurs_runtime.Value
var once_foldMap1 sync.Once
func Get_foldMap1() gopurs_runtime.Value {
	once_foldMap1.Do(func() {
		cache_foldMap1 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1
}

var cache_foldl1Default gopurs_runtime.Value
var once_foldl1Default sync.Once
func Get_foldl1Default() gopurs_runtime.Value {
	once_foldl1Default.Do(func() {
		cache_foldl1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1Default(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_foldl1Default
}

var cache_foldr1Default gopurs_runtime.Value
var once_foldr1Default sync.Once
func Get_foldr1Default() gopurs_runtime.Value {
	once_foldr1Default.Do(func() {
		cache_foldr1Default = gopurs_runtime.Func(func(dictFoldable1_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr1Default(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box))
})
	})
	return cache_foldr1Default
}

var cache_intercalateMap gopurs_runtime.Value
var once_intercalateMap sync.Once
func Get_intercalateMap() gopurs_runtime.Value {
	once_intercalateMap.Do(func() {
		cache_intercalateMap = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalateMap(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box))
})
	})
	return cache_intercalateMap
}

var cache_intercalate gopurs_runtime.Value
var once_intercalate sync.Once
func Get_intercalate() gopurs_runtime.Value {
	once_intercalate.Do(func() {
		cache_intercalate = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box))
})
	})
	return cache_intercalate
}

var cache_maximum gopurs_runtime.Value
var once_maximum sync.Once
func Get_maximum() gopurs_runtime.Value {
	once_maximum.Do(func() {
		cache_maximum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maximum(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_maximum
}

var cache_minimum gopurs_runtime.Value
var once_minimum sync.Once
func Get_minimum() gopurs_runtime.Value {
	once_minimum.Do(func() {
		cache_minimum = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_minimum(gopurs_runtime.CoerceToStruct[pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_minimum
}

var cache_traverse1_ gopurs_runtime.Value
var once_traverse1_ sync.Once
func Get_traverse1_() gopurs_runtime.Value {
	once_traverse1_.Do(func() {
		cache_traverse1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1_(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_traverse1_
}

var cache_for1_ gopurs_runtime.Value
var once_for1_ sync.Once
func Get_for1_() gopurs_runtime.Value {
	once_for1_.Do(func() {
		cache_for1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_for1_(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_for1_
}

var cache_sequence1_ gopurs_runtime.Value
var once_sequence1_ sync.Once
func Get_sequence1_() gopurs_runtime.Value {
	once_sequence1_.Do(func() {
		cache_sequence1_ = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sequence1_(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_sequence1_
}

var cache_fold1 gopurs_runtime.Value
var once_fold1 sync.Once
func Get_fold1() gopurs_runtime.Value {
	once_fold1.Do(func() {
		cache_fold1 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictSemigroup_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fold1(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](dictSemigroup_1_box))
})
	})
	return cache_fold1
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

var cache_applySecond__1627424644 gopurs_runtime.Value
var once_applySecond__1627424644 sync.Once
func Get_applySecond__1627424644() gopurs_runtime.Value {
	once_applySecond__1627424644.Do(func() {
		cache_applySecond__1627424644 = gopurs_runtime.Func(func(dictApply_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_applySecond__1627424644(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0_box))
})
	})
	return cache_applySecond__1627424644
}

var cache_categoryFn__3492036198 gopurs_runtime.Value
var once_categoryFn__3492036198 sync.Once
func Get_categoryFn__3492036198() gopurs_runtime.Value {
	once_categoryFn__3492036198.Do(func() {
		cache_categoryFn__3492036198 = gopurs_runtime.RecordDict2("Semigroupoid0", "identity", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Semigroupoid.Get_semigroupoidFn()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_categoryFn__3492036198
}

var cache_identity__2527656589 gopurs_runtime.Value
var once_identity__2527656589 sync.Once
func Get_identity__2527656589() gopurs_runtime.Value {
	once_identity__2527656589.Do(func() {
		cache_identity__2527656589 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_identity__2527656589(dict_0_box)
})
	})
	return cache_identity__2527656589
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

var cache_semigroupoidFn__3002128382 gopurs_runtime.Value
var once_semigroupoidFn__3002128382 sync.Once
func Get_semigroupoidFn__3002128382() gopurs_runtime.Value {
	once_semigroupoidFn__3002128382.Do(func() {
		cache_semigroupoidFn__3002128382 = gopurs_runtime.RecordDict1("compose", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(g_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(g_1, x_2))
})
})
}))
	})
	return cache_semigroupoidFn__3002128382
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

var cache_eq__1272715810 gopurs_runtime.Value
var once_eq__1272715810 sync.Once
func Get_eq__1272715810() gopurs_runtime.Value {
	once_eq__1272715810.Do(func() {
		cache_eq__1272715810 = gopurs_runtime.RecordGet(pkg_Data_Ordering.Get_eqOrdering(), "eq")
	})
	return cache_eq__1272715810
}

var cache_foldableDual__1841171440 gopurs_runtime.Value
var once_foldableDual__1841171440 sync.Once
func Get_foldableDual__1841171440() gopurs_runtime.Value {
	once_foldableDual__1841171440.Do(func() {
		cache_foldableDual__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_foldableDual__1841171440
}

var cache_foldableIdentity__1841171440 gopurs_runtime.Value
var once_foldableIdentity__1841171440 sync.Once
func Get_foldableIdentity__1841171440() gopurs_runtime.Value {
	once_foldableIdentity__1841171440.Do(func() {
		cache_foldableIdentity__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_foldableIdentity__1841171440
}

var cache_foldableMultiplicative__1841171440 gopurs_runtime.Value
var once_foldableMultiplicative__1841171440 sync.Once
func Get_foldableMultiplicative__1841171440() gopurs_runtime.Value {
	once_foldableMultiplicative__1841171440.Do(func() {
		cache_foldableMultiplicative__1841171440 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, v_2)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, v_2, z_1)
})
})
}))
	})
	return cache_foldableMultiplicative__1841171440
}

var cache_foldableTuple__1455669080 gopurs_runtime.Value
var once_foldableTuple__1455669080 sync.Once
func Get_foldableTuple__1455669080() gopurs_runtime.Value {
	once_foldableTuple__1455669080.Do(func() {
		cache_foldableTuple__1455669080 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1)
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, z_1)
})
})
}))
	})
	return cache_foldableTuple__1455669080
}

var cache_const__1496134642 gopurs_runtime.Value
var once_const__1496134642 sync.Once
func Get_const__1496134642() gopurs_runtime.Value {
	once_const__1496134642.Do(func() {
		cache_const__1496134642 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1496134642(a_0_box, v_1_box)
})
	})
	return cache_const__1496134642
}

var cache_const__1832354163 gopurs_runtime.Value
var once_const__1832354163 sync.Once
func Get_const__1832354163() gopurs_runtime.Value {
	once_const__1832354163.Do(func() {
		cache_const__1832354163 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__1832354163(a_0_box, v_1_box)
})
	})
	return cache_const__1832354163
}

var cache_const__641934996 gopurs_runtime.Value
var once_const__641934996 sync.Once
func Get_const__641934996() gopurs_runtime.Value {
	once_const__641934996.Do(func() {
		cache_const__641934996 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__641934996(a_0_box, v_1_box)
})
	})
	return cache_const__641934996
}

var cache_flip__2673533882 gopurs_runtime.Value
var once_flip__2673533882 sync.Once
func Get_flip__2673533882() gopurs_runtime.Value {
	once_flip__2673533882.Do(func() {
		cache_flip__2673533882 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__2673533882(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__2673533882
}

var cache_flip__3525081280 gopurs_runtime.Value
var once_flip__3525081280 sync.Once
func Get_flip__3525081280() gopurs_runtime.Value {
	once_flip__3525081280.Do(func() {
		cache_flip__3525081280 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3525081280(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3525081280
}

var cache_flip__535340480 gopurs_runtime.Value
var once_flip__535340480 sync.Once
func Get_flip__535340480() gopurs_runtime.Value {
	once_flip__535340480.Do(func() {
		cache_flip__535340480 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__535340480(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__535340480
}

var cache_flip__3659058336 gopurs_runtime.Value
var once_flip__3659058336 sync.Once
func Get_flip__3659058336() gopurs_runtime.Value {
	once_flip__3659058336.Do(func() {
		cache_flip__3659058336 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__3659058336(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__3659058336
}

var cache_flip__1074184704 gopurs_runtime.Value
var once_flip__1074184704 sync.Once
func Get_flip__1074184704() gopurs_runtime.Value {
	once_flip__1074184704.Do(func() {
		cache_flip__1074184704 = gopurs_runtime.Func3(func(f_0_box gopurs_runtime.Value, b_1_box gopurs_runtime.Value, a_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_flip__1074184704(f_0_box, b_1_box, a_2_box)
})
	})
	return cache_flip__1074184704
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

var cache_map__1483545076 gopurs_runtime.Value
var once_map__1483545076 sync.Once
func Get_map__1483545076() gopurs_runtime.Value {
	once_map__1483545076.Do(func() {
		cache_map__1483545076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1483545076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1483545076
}

var cache_voidRight__3258033404 gopurs_runtime.Value
var once_voidRight__3258033404 sync.Once
func Get_voidRight__3258033404() gopurs_runtime.Value {
	once_voidRight__3258033404.Do(func() {
		cache_voidRight__3258033404 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidRight__3258033404(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), x_1_box)
})
	})
	return cache_voidRight__3258033404
}

var cache_voidRight__1142845180 gopurs_runtime.Value
var once_voidRight__1142845180 sync.Once
func Get_voidRight__1142845180() gopurs_runtime.Value {
	once_voidRight__1142845180.Do(func() {
		cache_voidRight__1142845180 = gopurs_runtime.Func2(func(dictFunctor_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_voidRight__1142845180(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dictFunctor_0_box), x_1_box)
})
	})
	return cache_voidRight__1142845180
}

var cache_ala__2430830187 gopurs_runtime.Value
var once_ala__2430830187 sync.Once
func Get_ala__2430830187() gopurs_runtime.Value {
	once_ala__2430830187.Do(func() {
		cache_ala__2430830187 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, v_3_box gopurs_runtime.Value, f_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ala__2430830187(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_1_box), gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_2_box), v_3_box, f_4_box)
})
	})
	return cache_ala__2430830187
}

var cache_alaF__4085337484 gopurs_runtime.Value
var once_alaF__4085337484 sync.Once
func Get_alaF__4085337484() gopurs_runtime.Value {
	once_alaF__4085337484.Do(func() {
		cache_alaF__4085337484 = gopurs_runtime.Func5(func(_dollar__unused_0_box gopurs_runtime.Value, _dollar__unused_1_box gopurs_runtime.Value, _dollar__unused_2_box gopurs_runtime.Value, _dollar__unused_3_box gopurs_runtime.Value, v_4_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF__4085337484(_dollar__unused_0_box, _dollar__unused_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_2_box), gopurs_runtime.CoerceToStruct[pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value]](_dollar__unused_3_box), v_4_box)
})
	})
	return cache_alaF__4085337484
}

var cache_alaF__2035907084 gopurs_runtime.Value
var once_alaF__2035907084 sync.Once
func Get_alaF__2035907084() gopurs_runtime.Value {
	once_alaF__2035907084.Do(func() {
		cache_alaF__2035907084 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_alaF__2035907084(v_0_box)
})
	})
	return cache_alaF__2035907084
}

var cache_foldMap1__3342855683 gopurs_runtime.Value
var once_foldMap1__3342855683 sync.Once
func Get_foldMap1__3342855683() gopurs_runtime.Value {
	once_foldMap1__3342855683.Do(func() {
		cache_foldMap1__3342855683 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3342855683(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3342855683
}

var cache_foldMap1__3881804011 gopurs_runtime.Value
var once_foldMap1__3881804011 sync.Once
func Get_foldMap1__3881804011() gopurs_runtime.Value {
	once_foldMap1__3881804011.Do(func() {
		cache_foldMap1__3881804011 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3881804011(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3881804011
}

var cache_foldMap1__1749181674 gopurs_runtime.Value
var once_foldMap1__1749181674 sync.Once
func Get_foldMap1__1749181674() gopurs_runtime.Value {
	once_foldMap1__1749181674.Do(func() {
		cache_foldMap1__1749181674 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__1749181674(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__1749181674
}

var cache_foldMap1__3028551755 gopurs_runtime.Value
var once_foldMap1__3028551755 sync.Once
func Get_foldMap1__3028551755() gopurs_runtime.Value {
	once_foldMap1__3028551755.Do(func() {
		cache_foldMap1__3028551755 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldMap1__3028551755(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldMap1__3028551755
}

var cache_foldRight1Semigroup__1201419834 gopurs_runtime.Value
var once_foldRight1Semigroup__1201419834 sync.Once
func Get_foldRight1Semigroup__1201419834() gopurs_runtime.Value {
	once_foldRight1Semigroup__1201419834.Do(func() {
		cache_foldRight1Semigroup__1201419834 = gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_0 := (*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V1
_ = __local_var_2_0
return gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(&Constructor_FoldRight1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v_0.UnsafePtr).V0, gopurs_runtime.Apply2(f_4, __local_var_2_0, gopurs_runtime.Apply2((*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V0, a_3, f_4)), f_4)
})
}), (*Constructor_FoldRight1[gopurs_runtime.Value])(v1_1.UnsafePtr).V1})}
})
}))
	})
	return cache_foldRight1Semigroup__1201419834
}

var cache_foldl1__3059734942 gopurs_runtime.Value
var once_foldl1__3059734942 sync.Once
func Get_foldl1__3059734942() gopurs_runtime.Value {
	once_foldl1__3059734942.Do(func() {
		cache_foldl1__3059734942 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldl1__3059734942(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldl1__3059734942
}

var cache_foldr1__3059734942 gopurs_runtime.Value
var once_foldr1__3059734942 sync.Once
func Get_foldr1__3059734942() gopurs_runtime.Value {
	once_foldr1__3059734942.Do(func() {
		cache_foldr1__3059734942 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_foldr1__3059734942(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_foldr1__3059734942
}

var cache_getAct__3579306522 gopurs_runtime.Value
var once_getAct__3579306522 sync.Once
func Get_getAct__3579306522() gopurs_runtime.Value {
	once_getAct__3579306522.Do(func() {
		cache_getAct__3579306522 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_getAct__3579306522(v_0_box)
})
	})
	return cache_getAct__3579306522
}

var cache_joinee__1857126433 gopurs_runtime.Value
var once_joinee__1857126433 sync.Once
func Get_joinee__1857126433() gopurs_runtime.Value {
	once_joinee__1857126433.Do(func() {
		cache_joinee__1857126433 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_joinee__1857126433(v_0_box)
})
	})
	return cache_joinee__1857126433
}

var cache_mkFoldRight1__364767315 gopurs_runtime.Value
var once_mkFoldRight1__364767315 sync.Once
func Get_mkFoldRight1__364767315() gopurs_runtime.Value {
	once_mkFoldRight1__364767315.Do(func() {
		cache_mkFoldRight1__364767315 = gopurs_runtime.Apply(Get_FoldRight1(), pkg_Data_Function.Get_go__const())
	})
	return cache_mkFoldRight1__364767315
}

var cache_runFoldRight1__3719153019 gopurs_runtime.Value
var once_runFoldRight1__3719153019 sync.Once
func Get_runFoldRight1__3719153019() gopurs_runtime.Value {
	once_runFoldRight1__3719153019.Do(func() {
		cache_runFoldRight1__3719153019 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_runFoldRight1__3719153019(gopurs_runtime.CoerceToStruct[Constructor_FoldRight1[gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_runFoldRight1__3719153019
}

var cache_traverse1___3055398386 gopurs_runtime.Value
var once_traverse1___3055398386 sync.Once
func Get_traverse1___3055398386() gopurs_runtime.Value {
	once_traverse1___3055398386.Do(func() {
		cache_traverse1___3055398386 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1___3055398386(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_traverse1___3055398386
}

var cache_traverse1___1509071858 gopurs_runtime.Value
var once_traverse1___1509071858 sync.Once
func Get_traverse1___1509071858() gopurs_runtime.Value {
	once_traverse1___1509071858.Do(func() {
		cache_traverse1___1509071858 = gopurs_runtime.Func2(func(dictFoldable1_0_box gopurs_runtime.Value, dictApply_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_traverse1___1509071858(gopurs_runtime.CoerceToStruct[Constructor_Foldable1[gopurs_runtime.Value]](dictFoldable1_0_box), gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_1_box))
})
	})
	return cache_traverse1___1509071858
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

var cache_coerce__529298068 gopurs_runtime.Value
var once_coerce__529298068 sync.Once
func Get_coerce__529298068() gopurs_runtime.Value {
	once_coerce__529298068.Do(func() {
		cache_coerce__529298068 = gopurs_runtime.Func(func(_dollar__unused_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_coerce__529298068(_dollar__unused_0_box)
})
	})
	return cache_coerce__529298068
}

var cache_coerce__3644730900 gopurs_runtime.Value
var once_coerce__3644730900 sync.Once
func Get_coerce__3644730900() gopurs_runtime.Value {
	once_coerce__3644730900.Do(func() {
		cache_coerce__3644730900 = pkg_Unsafe_Coerce.Get_unsafeCoerce()
	})
	return cache_coerce__3644730900
}

type Constructor_FoldRight1[T_a any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 T_a
}


type Constructor_Foldable1[T_t any] struct {
	Rc uint32
	V0 gopurs_runtime.Value
	V1 gopurs_runtime.Value
	V2 gopurs_runtime.Value
	V3 gopurs_runtime.Value
}


func init() {
	gopurs_runtime.StructGetters[2465059545] = func(ptr unsafe.Pointer, key string) gopurs_runtime.Value {
		c := (*Constructor_Foldable1[gopurs_runtime.Value])(ptr)
		switch key {
		case "Foldable0": return c.V0
		case "foldMap1": return c.V1
		case "foldl1": return c.V2
		case "foldr1": return c.V3
		default: panic("Key not found in dictionary Constructor_Foldable1: " + key)
		}
	}
}


func Call_identity(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_identity1(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_foldr1(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_foldl1(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_maximumBy(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(dictFoldable1_0.V2, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__1272715810(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_1, x_2, y_3).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}).IntVal) != (0) {
__t0 = x_2
goto end_branch_0
} else {

}
}
{
__t0 = y_3
}
end_branch_0:
return __t0
})
}))
}

func Call_minimumBy(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], cmp_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var cmp_1 gopurs_runtime.Value = cmp_1_loop
_ = cmp_1
return gopurs_runtime.Apply(dictFoldable1_0.V2, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Apply2(Get_eq__1272715810(), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(cmp_1, x_2, y_3).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}).IntVal) != (0) {
__t0 = x_2
goto end_branch_0
} else {

}
}
{
__t0 = y_3
}
end_branch_0:
return __t0
})
}))
}

func Call_foldMap1DefaultR(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictFunctor_1_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], dictSemigroup_2_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_2_loop
_ = dictSemigroup_2
append_3_0 := dictSemigroup_2.V0
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(dictFunctor_1.V0, f_4)
_ = __local_var_5_1
__local_var_6_2 := gopurs_runtime.Apply(dictFoldable1_0.V3, append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_foldMap1DefaultL(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictFunctor_1_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], dictSemigroup_2_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictFunctor_1 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_1_loop
_ = dictFunctor_1
var dictSemigroup_2 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_2_loop
_ = dictSemigroup_2
append_3_0 := dictSemigroup_2.V0
_ = append_3_0
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_1 := gopurs_runtime.Apply(dictFunctor_1.V0, f_4)
_ = __local_var_5_1
__local_var_6_2 := gopurs_runtime.Apply(dictFoldable1_0.V2, append_3_0)
_ = __local_var_6_2
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_6_2, gopurs_runtime.Apply(__local_var_5_1, x_7))
})
})
}

func Call_foldMap1(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl1Default(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_2 := gopurs_runtime.Apply2(dictFoldable1_0.V1, Get_semigroupDual(), Get_mkFoldRight1())
_ = __local_var_1_2
__local_var_1_1 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_3 := gopurs_runtime.Apply(__local_var_1_2, x_2)
_ = __local_var_3_3
return gopurs_runtime.Apply((*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_3_3.UnsafePtr).V0, (*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_3_3.UnsafePtr).V1)
})
_ = __local_var_1_1
__local_var_1_0 := gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_1, a_3, b_2)
})
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(x_2, a_4, b_3)
})
}))
})
}

func Call_foldr1Default(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
__local_var_1_1 := gopurs_runtime.Apply2(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[*Constructor_FoldRight1[gopurs_runtime.Value]]](Get_foldRight1Semigroup()))}, Get_mkFoldRight1())
_ = __local_var_1_1
__local_var_1_0 := gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_2 := gopurs_runtime.Apply(__local_var_1_1, x_2)
_ = __local_var_3_2
return gopurs_runtime.Apply((*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_3_2.UnsafePtr).V0, (*Constructor_FoldRight1[gopurs_runtime.Value])(__local_var_3_2.UnsafePtr).V1)
})
_ = __local_var_1_0
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_1_0, a_3, b_2)
})
})
}

func Call_intercalateMap(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictSemigroup_1_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
semigroupJoinWith1_2_0 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemigroup_1.V0, gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(dictSemigroup_1.V0, j_4, gopurs_runtime.Apply(v1_3, j_4)))
})
})
})}
_ = semigroupJoinWith1_2_0
return gopurs_runtime.Func(func(j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_2_0)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_1 := gopurs_runtime.Apply(f_4, x_6)
_ = __local_var_7_1
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_1
})
}), foldable_5, j_3)
})
})
})
}

func Call_intercalate(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictSemigroup_1_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
semigroupJoinWith1_2_1 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(j_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictSemigroup_1.V0, gopurs_runtime.Apply(v_2, j_4), gopurs_runtime.Apply2(dictSemigroup_1.V0, j_4, gopurs_runtime.Apply(v1_3, j_4)))
})
})
})}
_ = semigroupJoinWith1_2_1
__local_var_2_0 := gopurs_runtime.Func(func(j_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(foldable_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupJoinWith1_2_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_2 := gopurs_runtime.Apply(f_4, x_6)
_ = __local_var_7_2
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return __local_var_7_2
})
}), foldable_5, j_3)
})
})
})
_ = __local_var_2_0
return gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_3, Get_identity())
})
}

func Call_maximum(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
semigroupMax_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply2(dictOrd_0.V1, v_1, v1_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_3_1.IntVal) == 1527465420) {
__t2 = v1_2
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 902936544) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 380165415) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
}))
_ = semigroupMax_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), semigroupMax_1_0, pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
}

func Call_minimum(dictOrd_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
semigroupMin_1_0 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
v_3_1 := gopurs_runtime.Apply2(dictOrd_0.V1, v_1, v1_2)
_ = v_3_1
var __t2 gopurs_runtime.Value
{
if (uint32(v_3_1.IntVal) == 1527465420) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 902936544) {
__t2 = v_1
goto end_branch_2
} else {

}
}
{
if (uint32(v_3_1.IntVal) == 380165415) {
__t2 = v1_2
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
})
}))
_ = semigroupMin_1_0
return gopurs_runtime.Func(func(dictFoldable1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFoldable1_2, "foldMap1"), semigroupMin_1_0, pkg_Unsafe_Coerce.Get_unsafeCoerce())
})
}

func Call_traverse1_(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
Functor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}))
_ = Functor0_2_0
semigroupAct1_3_1 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_1.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_0.V0, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Apply3(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
})
}

func Call_for1_(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
__local_var_2_0 := Call_traverse1_(dictFoldable1_0, dictApply_1)
_ = __local_var_2_0
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(__local_var_2_0, a_4, b_3)
})
})
}

func Call_sequence1_(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
return gopurs_runtime.Apply(Call_traverse1_(dictFoldable1_0, dictApply_1), Get_identity1())
}

func Call_fold1(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictSemigroup_1_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictSemigroup_1 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dictSemigroup_1_loop
_ = dictSemigroup_1
return gopurs_runtime.Apply2(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(dictSemigroup_1)}, Get_identity())
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_applySecond__1627424644(dictApply_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApply_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_0_loop
_ = dictApply_0
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_0.V0, gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_0.V1, gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_5
})
}), a_2), b_3)
})
})
}

func Call_identity__2527656589(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "identity")
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

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_const__1496134642(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__1832354163(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_flip__2673533882(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3525081280(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__535340480(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__3659058336(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_flip__1074184704(f_0_loop gopurs_runtime.Value, b_1_loop gopurs_runtime.Value, a_2_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var b_1 gopurs_runtime.Value = b_1_loop
_ = b_1
var a_2 gopurs_runtime.Value = a_2_loop
_ = a_2
return gopurs_runtime.Apply2(f_0, a_2, b_1)
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1483545076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_voidRight__3258033404(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_voidRight__1142845180(dictFunctor_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dictFunctor_0_loop
_ = dictFunctor_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(dictFunctor_0.V0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}

func Call_ala__2430830187(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], _dollar__unused_2_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_3_loop gopurs_runtime.Value, f_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_2_loop
_ = _dollar__unused_2
var v_3 gopurs_runtime.Value = v_3_loop
_ = v_3
var f_4 gopurs_runtime.Value = f_4_loop
_ = f_4
return gopurs_runtime.Apply(f_4, pkg_Unsafe_Coerce.Get_unsafeCoerce())
}

func Call_alaF__4085337484(_dollar__unused_0_loop gopurs_runtime.Value, _dollar__unused_1_loop gopurs_runtime.Value, _dollar__unused_2_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], _dollar__unused_3_loop *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value], v_4_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var _dollar__unused_1 gopurs_runtime.Value = _dollar__unused_1_loop
_ = _dollar__unused_1
var _dollar__unused_2 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_2_loop
_ = _dollar__unused_2
var _dollar__unused_3 *pkg_Data_Newtype.Constructor_Newtype[gopurs_runtime.Value, gopurs_runtime.Value] = _dollar__unused_3_loop
_ = _dollar__unused_3
var v_4 gopurs_runtime.Value = v_4_loop
_ = v_4
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_alaF__2035907084(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}

func Call_foldMap1__3342855683(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMap1__3881804011(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return gopurs_runtime.Apply(dict_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[*Constructor_FoldRight1[gopurs_runtime.Value]]](Get_foldRight1Semigroup()))})
}

func Call_foldMap1__1749181674(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldMap1__3028551755(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_foldl1__3059734942(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_foldr1__3059734942(dict_0_loop *Constructor_Foldable1[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *Constructor_Foldable1[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_getAct__3579306522(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_joinee__1857126433(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return v_0
}

func Call_runFoldRight1__3719153019(v_0_loop *Constructor_FoldRight1[gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *Constructor_FoldRight1[gopurs_runtime.Value] = v_0_loop
_ = v_0
return gopurs_runtime.Apply((*Constructor_FoldRight1[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0, (*Constructor_FoldRight1[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 3805997843, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1)
}

func Call_traverse1___3055398386(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
Functor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}))
_ = Functor0_2_0
semigroupAct1_3_1 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_1.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_0.V0, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Apply3(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
})
}

func Call_traverse1___1509071858(dictFoldable1_0_loop *Constructor_Foldable1[gopurs_runtime.Value], dictApply_1_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable1_0 *Constructor_Foldable1[gopurs_runtime.Value] = dictFoldable1_0_loop
_ = dictFoldable1_0
var dictApply_1 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dictApply_1_loop
_ = dictApply_1
Functor0_2_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}))
_ = Functor0_2_0
semigroupAct1_3_1 := &pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(dictApply_1.V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(dictApply_1.V0, gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return x_6
})
}), v_3), v1_4)
})
})}
_ = semigroupAct1_3_1
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(t_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_2_0.V0, gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Unit.Get_unit()
}), gopurs_runtime.Apply3(dictFoldable1_0.V1, gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupAct1_3_1)}, gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, x_6)
}), t_5))
})
})
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_coerce__529298068(_dollar__unused_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
return pkg_Unsafe_Coerce.Get_unsafeCoerce()
}



package Data_NonEmpty

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_HeytingAlgebra "gopurs/output/Data.HeytingAlgebra"
	unsafe "unsafe"
)

var cache_NonEmpty gopurs_runtime.Value
var once_NonEmpty sync.Once
func Get_NonEmpty() gopurs_runtime.Value {
	once_NonEmpty.Do(func() {
		cache_NonEmpty = gopurs_runtime.Func(func(value0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(value1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{value0, value1})}
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
return Call_tail(v_0_box)
})
	})
	return cache_tail
}

var cache_singleton gopurs_runtime.Value
var once_singleton sync.Once
func Get_singleton() gopurs_runtime.Value {
	once_singleton.Do(func() {
		cache_singleton = gopurs_runtime.Func(func(dictPlus_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_singleton(dictPlus_0_box)
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
		cache_oneOf = gopurs_runtime.Func2(func(dictAlternative_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_oneOf(dictAlternative_0_box, v_1_box)
})
	})
	return cache_oneOf
}

var cache_head gopurs_runtime.Value
var once_head sync.Once
func Get_head() gopurs_runtime.Value {
	once_head.Do(func() {
		cache_head = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_head(v_0_box)
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
return Call_fromNonEmpty(f_0_box, v_1_box)
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
return Call_foldl1(dictFoldable_0_box)
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

type Constructor_NonEmpty[T_f any, T_a any] struct {
	V0 T_a
	V1 T_f
}


func Call_unfoldable1NonEmpty(dictUnfoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 gopurs_runtime.Value = dictUnfoldable_0_loop
_ = dictUnfoldable_0
return gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_0 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Tuple.Get_functorTuple(), "map"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictUnfoldable_0, "unfoldr"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Maybe.Get_functorMaybe(), "map"), f_1)), gopurs_runtime.Apply(f_1, b_2))
_ = __local_var_3_0
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{(*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V0, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(__local_var_3_0.UnsafePtr).V1})}
}))
}

func Call_tail(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V1
}

func Call_singleton(dictPlus_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictPlus_0 gopurs_runtime.Value = dictPlus_0_loop
_ = dictPlus_0
empty_1_0 := gopurs_runtime.RecordGet(dictPlus_0, "empty")
_ = empty_1_0
return gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{a_2, empty_1_0})}
})
}

func Call_showNonEmpty(dictShow_0_loop gopurs_runtime.Value, dictShow1_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
var dictShow1_1 gopurs_runtime.Value = dictShow1_1_loop
_ = dictShow1_1
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(NonEmpty "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str(" "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow1_1, "show"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1), gopurs_runtime.Str(")")))))
}))
}

func Call_semigroupNonEmpty(dictApplicative_0_loop gopurs_runtime.Value, dictSemigroup_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 gopurs_runtime.Value = dictApplicative_0_loop
_ = dictApplicative_0
var dictSemigroup_1 gopurs_runtime.Value = dictSemigroup_1_loop
_ = dictSemigroup_1
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(v_2 gopurs_runtime.Value, v1_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{(*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1, gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_1, "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "pure"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v1_3.UnsafePtr).V1))})}
}))
}

func Call_oneOf(dictAlternative_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictAlternative_0 gopurs_runtime.Value = dictAlternative_0_loop
_ = dictAlternative_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Plus1"), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "alt"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictAlternative_0, "Applicative0"), gopurs_runtime.Value{}), "pure"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)
}

func Call_head(v_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
return (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_0.UnsafePtr).V0
}

func Call_functorNonEmpty(dictFunctor_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctor_0 gopurs_runtime.Value = dictFunctor_0_loop
_ = dictFunctor_0
return gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_1 gopurs_runtime.Value, m_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctor_0, "map"), f_1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_2.UnsafePtr).V1)})}
}))
}

func Call_functorWithIndex(dictFunctorWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFunctorWithIndex_0 gopurs_runtime.Value = dictFunctorWithIndex_0_loop
_ = dictFunctorWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorNonEmpty1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)})}
}))
_ = functorNonEmpty1_2_1
return gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_2_1
}), gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictFunctorWithIndex_0, "mapWithIndex"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{x_5})})
}), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)})}
}))
}

func Call_fromNonEmpty(f_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return gopurs_runtime.Apply2(f_0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_1.UnsafePtr).V1)
}

func Call_foldableNonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1)
_ = foldMap1_2_0
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(foldMap1_2_0, f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, b_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
}))
}

func Call_foldableWithIndexNonEmpty(dictFoldableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldableWithIndex_0 gopurs_runtime.Value = dictFoldableWithIndex_0_loop
_ = dictFoldableWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "Foldable0"), gopurs_runtime.Value{})
_ = __local_var_1_0
foldableNonEmpty1_2_1 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_2 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "foldMap"), dictMonoid_2)
_ = foldMap1_3_2
return gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_2, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_4, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(foldMap1_3_2, f_4, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldl"), f_2, gopurs_runtime.Apply2(f_2, b_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_1_0, "foldr"), f_2, b_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
}))
_ = foldableNonEmpty1_2_1
return gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_2_1
}), gopurs_runtime.Func(func(dictMonoid_3 gopurs_runtime.Value) gopurs_runtime.Value {
foldMapWithIndex1_4_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldMapWithIndex"), dictMonoid_3)
_ = foldMapWithIndex1_4_3
return gopurs_runtime.Func2(func(f_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_3, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply2(f_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), gopurs_runtime.Apply2(foldMapWithIndex1_4_3, gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_5, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{x_7})})
}), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldlWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{x_6})})
}), gopurs_runtime.Apply3(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, b_4, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldableWithIndex_0, "foldrWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_3, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{x_6})})
}), b_4, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1))
}))
}

func Call_traversableNonEmpty(dictTraversable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversable_0 gopurs_runtime.Value = dictTraversable_0_loop
_ = dictTraversable_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_1_0
functorNonEmpty1_2_1 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, m_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "map"), f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_3.UnsafePtr).V1)})}
}))
_ = functorNonEmpty1_2_1
__local_var_3_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "Foldable1"), gopurs_runtime.Value{})
_ = __local_var_3_2
foldableNonEmpty1_4_3 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_4 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_5_4 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_3_2, "foldMap"), dictMonoid_4)
_ = foldMap1_5_4
return gopurs_runtime.Func2(func(f_6 gopurs_runtime.Value, v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_4, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_6, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V0), gopurs_runtime.Apply2(foldMap1_5_4, f_6, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_7.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldl"), f_4, gopurs_runtime.Apply2(f_4, b_5, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_4, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(__local_var_3_2, "foldr"), f_4, b_5, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_6.UnsafePtr).V1))
}))
_ = foldableNonEmpty1_4_3
return gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_4_3
}), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_2_1
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_6_5 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_5
sequence1_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "sequence"), dictApplicative_5)
_ = sequence1_7_6
return gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_6_5, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_5, "Functor0"), gopurs_runtime.Value{}), "map"), Get_NonEmpty(), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V0), gopurs_runtime.Apply(sequence1_7_6, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_8.UnsafePtr).V1))
})
}), gopurs_runtime.Func(func(dictApplicative_5 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_6_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_5, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_6_7
traverse1_7_8 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversable_0, "traverse"), dictApplicative_5)
_ = traverse1_7_8
return gopurs_runtime.Func2(func(f_8 gopurs_runtime.Value, v_9 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_6_7, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_6_7, "Functor0"), gopurs_runtime.Value{}), "map"), Get_NonEmpty(), gopurs_runtime.Apply(f_8, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V0)), gopurs_runtime.Apply2(traverse1_7_8, f_8, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_9.UnsafePtr).V1))
})
}))
}

func Call_traversableWithIndexNonEmpty(dictTraversableWithIndex_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictTraversableWithIndex_0 gopurs_runtime.Value = dictTraversableWithIndex_0_loop
_ = dictTraversableWithIndex_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FunctorWithIndex0"), gopurs_runtime.Value{})
_ = __local_var_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(__local_var_1_0, "Functor0"), gopurs_runtime.Value{})
_ = __local_var_2_1
functorNonEmpty1_3_3 := gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, m_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply(f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "map"), f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(m_4.UnsafePtr).V1)})}
}))
_ = functorNonEmpty1_3_3
functorWithIndex1_3_2 := gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return functorNonEmpty1_3_3
}), gopurs_runtime.Func2(func(f_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3111306138, UnsafePtr: unsafe.Pointer(&Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value]{gopurs_runtime.Apply2(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V0), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "mapWithIndex"), gopurs_runtime.Func(func(x_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_4, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{x_6})})
}), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_5.UnsafePtr).V1)})}
}))
_ = functorWithIndex1_3_2
foldableWithIndexNonEmpty1_4_4 := gopurs_runtime.Apply(Get_foldableWithIndexNonEmpty(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "FoldableWithIndex1"), gopurs_runtime.Value{}))
_ = foldableWithIndexNonEmpty1_4_4
traversableNonEmpty1_5_5 := gopurs_runtime.Apply(Get_traversableNonEmpty(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "Traversable2"), gopurs_runtime.Value{}))
_ = traversableNonEmpty1_5_5
return gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableWithIndexNonEmpty1_4_4
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return functorWithIndex1_3_2
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return traversableNonEmpty1_5_5
}), gopurs_runtime.Func(func(dictApplicative_6 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_7_6 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_6, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_7_6
traverseWithIndex1_8_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictTraversableWithIndex_0, "traverseWithIndex"), dictApplicative_6)
_ = traverseWithIndex1_8_7
return gopurs_runtime.Func2(func(f_9 gopurs_runtime.Value, v_10 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Apply0_7_6, "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_7_6, "Functor0"), gopurs_runtime.Value{}), "map"), Get_NonEmpty(), gopurs_runtime.Apply2(f_9, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V0)), gopurs_runtime.Apply2(traverseWithIndex1_8_7, gopurs_runtime.Func(func(x_11 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_9, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{x_11})})
}), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_10.UnsafePtr).V1))
})
}))
}

func Call_foldable1NonEmpty(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
foldableNonEmpty1_1_0 := gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_1 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictFoldable_0, "foldMap"), dictMonoid_1)
_ = foldMap1_2_1
return gopurs_runtime.Func2(func(f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_1, "Semigroup0"), gopurs_runtime.Value{}), "append"), gopurs_runtime.Apply(f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), gopurs_runtime.Apply2(foldMap1_2_1, f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1))
})
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_1, gopurs_runtime.Apply2(f_1, b_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
}), gopurs_runtime.Func3(func(f_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), f_1, b_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1))
}))
_ = foldableNonEmpty1_1_0
return gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return foldableNonEmpty1_1_0
}), gopurs_runtime.Func3(func(dictSemigroup_2 gopurs_runtime.Value, f_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), gopurs_runtime.Func2(func(s_5 gopurs_runtime.Value, a1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_2, "append"), s_5, gopurs_runtime.Apply(f_3, a1_6))
}), gopurs_runtime.Apply(f_3, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1)
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldl"), f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
}), gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(f_2, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0)
_ = __local_var_4_2
__local_var_5_3 := gopurs_runtime.Apply3(gopurs_runtime.RecordGet(dictFoldable_0, "foldr"), gopurs_runtime.Func(func(a1_5 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_6_4 := gopurs_runtime.Apply(f_2, a1_5)
_ = __local_var_6_4
return gopurs_runtime.Func(func(x_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t5 gopurs_runtime.Value
{
if (x_7.Type == 9 && x_7.IntVal == 930809136 && x_7.UnsafePtr == nil) {
__t5 = a1_5
goto end_branch_5
} else {

}
}
{
if (x_7.Type == 9 && x_7.IntVal == 930809136 && x_7.UnsafePtr != nil) {
__t5 = gopurs_runtime.Apply(__local_var_6_4, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(x_7.UnsafePtr).V0)
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]{__t5})}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: nil}, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
_ = __local_var_5_3
var __t6 gopurs_runtime.Value
{
if (__local_var_5_3.Type == 9 && __local_var_5_3.IntVal == 930809136 && __local_var_5_3.UnsafePtr == nil) {
__t6 = (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0
goto end_branch_6
} else {

}
}
{
if (__local_var_5_3.Type == 9 && __local_var_5_3.IntVal == 930809136 && __local_var_5_3.UnsafePtr != nil) {
__t6 = gopurs_runtime.Apply(__local_var_4_2, (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(__local_var_5_3.UnsafePtr).V0)
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return __t6
}))
}

func Call_foldl1(dictFoldable_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictFoldable_0 gopurs_runtime.Value = dictFoldable_0_loop
_ = dictFoldable_0
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_foldable1NonEmpty(), dictFoldable_0), "foldl1")
}

func Call_eqNonEmpty(dictEq1_0_loop gopurs_runtime.Value, dictEq_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq1_0 gopurs_runtime.Value = dictEq1_0_loop
_ = dictEq1_0
var dictEq_1 gopurs_runtime.Value = dictEq_1_loop
_ = dictEq_1
eq11_2_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEq1_0, "eq1"), dictEq_1)
_ = eq11_2_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_HeytingAlgebra.Get_heytingAlgebraBoolean(), "conj"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_1, "eq"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V0), gopurs_runtime.Apply2(eq11_2_0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_3.UnsafePtr).V1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_4.UnsafePtr).V1))
}))
}

func Call_ordNonEmpty(dictOrd1_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd1_0 gopurs_runtime.Value = dictOrd1_0_loop
_ = dictOrd1_0
eqNonEmpty1_1_0 := gopurs_runtime.Apply(Get_eqNonEmpty(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{}))
_ = eqNonEmpty1_1_0
return gopurs_runtime.Func(func(dictOrd_2 gopurs_runtime.Value) gopurs_runtime.Value {
compare11_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "compare1"), dictOrd_2)
_ = compare11_3_1
eqNonEmpty2_4_2 := gopurs_runtime.Apply(eqNonEmpty1_1_0, gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_2, "Eq0"), gopurs_runtime.Value{}))
_ = eqNonEmpty2_4_2
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return eqNonEmpty2_4_2
}), gopurs_runtime.Func2(func(x_5 gopurs_runtime.Value, y_6 gopurs_runtime.Value) gopurs_runtime.Value {
v_7_3 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_2, "compare"), (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V0, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V0)
_ = v_7_3
var __t4 gopurs_runtime.Value
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 1527465420) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 1527465420, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
if (v_7_3.Type == 9 && v_7_3.IntVal == 380165415) {
__t4 = gopurs_runtime.Value{Type: 9, IntVal: 380165415, UnsafePtr: nil}
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(compare11_3_1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(x_5.UnsafePtr).V1, (*Constructor_NonEmpty[gopurs_runtime.Value, gopurs_runtime.Value])(y_6.UnsafePtr).V1)
}
end_branch_4:
return __t4
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
ordNonEmpty1_1_0 := gopurs_runtime.Apply(Get_ordNonEmpty(), dictOrd1_0)
_ = ordNonEmpty1_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd1_0, "Eq10"), gopurs_runtime.Value{})
_ = __local_var_2_1
eq1NonEmpty1_3_2 := gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqNonEmpty(__local_var_2_1, dictEq_3), "eq")
}))
_ = eq1NonEmpty1_3_2
return gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return eq1NonEmpty1_3_2
}), gopurs_runtime.Func(func(dictOrd_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(ordNonEmpty1_1_0, dictOrd_4), "compare")
}))
}



package Data_Lazy

import (
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_EuclideanRing "gopurs/output/Data.EuclideanRing"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Ring "gopurs/output/Data.Ring"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semiring "gopurs/output/Data.Semiring"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_showLazy gopurs_runtime.Value
var once_showLazy sync.Once
func Get_showLazy() gopurs_runtime.Value {
	once_showLazy.Do(func() {
		cache_showLazy = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showLazy(dictShow_0_box)
})
	})
	return cache_showLazy
}

var cache_semiringLazy gopurs_runtime.Value
var once_semiringLazy sync.Once
func Get_semiringLazy() gopurs_runtime.Value {
	once_semiringLazy.Do(func() {
		cache_semiringLazy = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semiringLazy(dictSemiring_0_box)
})
	})
	return cache_semiringLazy
}

var cache_semigroupLazy gopurs_runtime.Value
var once_semigroupLazy sync.Once
func Get_semigroupLazy() gopurs_runtime.Value {
	once_semigroupLazy.Do(func() {
		cache_semigroupLazy = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_semigroupLazy(dictSemigroup_0_box)
})
	})
	return cache_semigroupLazy
}

var cache_ringLazy gopurs_runtime.Value
var once_ringLazy sync.Once
func Get_ringLazy() gopurs_runtime.Value {
	once_ringLazy.Do(func() {
		cache_ringLazy = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ringLazy(dictRing_0_box)
})
	})
	return cache_ringLazy
}

var cache_monoidLazy gopurs_runtime.Value
var once_monoidLazy sync.Once
func Get_monoidLazy() gopurs_runtime.Value {
	once_monoidLazy.Do(func() {
		cache_monoidLazy = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_monoidLazy(dictMonoid_0_box)
})
	})
	return cache_monoidLazy
}

var cache_lazyLazy gopurs_runtime.Value
var once_lazyLazy sync.Once
func Get_lazyLazy() gopurs_runtime.Value {
	once_lazyLazy.Do(func() {
		cache_lazyLazy = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
}))
	})
	return cache_lazyLazy
}

var cache_functorLazy gopurs_runtime.Value
var once_functorLazy sync.Once
func Get_functorLazy() gopurs_runtime.Value {
	once_functorLazy.Do(func() {
		cache_functorLazy = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_force(), l_1))
}))
})
}))
	})
	return cache_functorLazy
}

var cache_functorWithIndexLazy gopurs_runtime.Value
var once_functorWithIndexLazy sync.Once
func Get_functorWithIndexLazy() gopurs_runtime.Value {
	once_functorWithIndexLazy.Do(func() {
		cache_functorWithIndexLazy = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexLazy
}

var cache_invariantLazy gopurs_runtime.Value
var once_invariantLazy sync.Once
func Get_invariantLazy() gopurs_runtime.Value {
	once_invariantLazy.Do(func() {
		cache_invariantLazy = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), f_0)
})
}))
	})
	return cache_invariantLazy
}

var cache_foldableLazy gopurs_runtime.Value
var once_foldableLazy sync.Once
func Get_foldableLazy() gopurs_runtime.Value {
	once_foldableLazy.Do(func() {
		cache_foldableLazy = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.Apply(Get_force(), l_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(Get_force(), l_2), z_1)
})
})
}))
	})
	return cache_foldableLazy
}

var cache_foldableWithIndexLazy gopurs_runtime.Value
var once_foldableWithIndexLazy sync.Once
func Get_foldableWithIndexLazy() gopurs_runtime.Value {
	once_foldableWithIndexLazy.Do(func() {
		cache_foldableWithIndexLazy = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableLazy(), "foldMap"), dictMonoid_0)
_ = foldMap1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap1_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableLazy(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableLazy(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexLazy
}

var cache_traversableLazy gopurs_runtime.Value
var once_traversableLazy sync.Once
func Get_traversableLazy() gopurs_runtime.Value {
	once_traversableLazy.Do(func() {
		cache_traversableLazy = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
}), gopurs_runtime.Apply(Get_force(), l_2))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}))
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(Get_force(), l_3)))
})
})
}))
	})
	return cache_traversableLazy
}

var cache_traversableWithIndexLazy gopurs_runtime.Value
var once_traversableWithIndexLazy sync.Once
func Get_traversableWithIndexLazy() gopurs_runtime.Value {
	once_traversableWithIndexLazy.Do(func() {
		cache_traversableWithIndexLazy = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableLazy()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
traverse1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_traversableLazy(), "traverse"), dictApplicative_0)
_ = traverse1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse1_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}))
	})
	return cache_traversableWithIndexLazy
}

var cache_foldable1Lazy gopurs_runtime.Value
var once_foldable1Lazy sync.Once
func Get_foldable1Lazy() gopurs_runtime.Value {
	once_foldable1Lazy.Do(func() {
		cache_foldable1Lazy = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2))
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), l_1)
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), l_1)
})
}))
	})
	return cache_foldable1Lazy
}

var cache_traversable1Lazy gopurs_runtime.Value
var once_traversable1Lazy sync.Once
func Get_traversable1Lazy() gopurs_runtime.Value {
	once_traversable1Lazy.Do(func() {
		cache_traversable1Lazy = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1Lazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableLazy()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
}), gopurs_runtime.Apply(Get_force(), l_2))
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}))
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(Get_force(), l_3)))
})
})
}))
	})
	return cache_traversable1Lazy
}

var cache_extendLazy gopurs_runtime.Value
var once_extendLazy sync.Once
func Get_extendLazy() gopurs_runtime.Value {
	once_extendLazy.Do(func() {
		cache_extendLazy = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_1)
}))
})
}))
	})
	return cache_extendLazy
}

var cache_eqLazy gopurs_runtime.Value
var once_eqLazy sync.Once
func Get_eqLazy() gopurs_runtime.Value {
	once_eqLazy.Do(func() {
		cache_eqLazy = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqLazy(dictEq_0_box)
})
	})
	return cache_eqLazy
}

var cache_ordLazy gopurs_runtime.Value
var once_ordLazy sync.Once
func Get_ordLazy() gopurs_runtime.Value {
	once_ordLazy.Do(func() {
		cache_ordLazy = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordLazy(dictOrd_0_box)
})
	})
	return cache_ordLazy
}

var cache_eq1Lazy gopurs_runtime.Value
var once_eq1Lazy sync.Once
func Get_eq1Lazy() gopurs_runtime.Value {
	once_eq1Lazy.Do(func() {
		cache_eq1Lazy = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqLazy(dictEq_0), "eq")
}))
	})
	return cache_eq1Lazy
}

var cache_ord1Lazy gopurs_runtime.Value
var once_ord1Lazy sync.Once
func Get_ord1Lazy() gopurs_runtime.Value {
	once_ord1Lazy.Do(func() {
		cache_ord1Lazy = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Lazy()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_ordLazy(dictOrd_0), "compare")
}))
	})
	return cache_ord1Lazy
}

var cache_comonadLazy gopurs_runtime.Value
var once_comonadLazy sync.Once
func Get_comonadLazy() gopurs_runtime.Value {
	once_comonadLazy.Do(func() {
		cache_comonadLazy = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendLazy()
}), Get_force())
	})
	return cache_comonadLazy
}

var cache_commutativeRingLazy gopurs_runtime.Value
var once_commutativeRingLazy sync.Once
func Get_commutativeRingLazy() gopurs_runtime.Value {
	once_commutativeRingLazy.Do(func() {
		cache_commutativeRingLazy = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_commutativeRingLazy(dictCommutativeRing_0_box)
})
	})
	return cache_commutativeRingLazy
}

var cache_euclideanRingLazy gopurs_runtime.Value
var once_euclideanRingLazy sync.Once
func Get_euclideanRingLazy() gopurs_runtime.Value {
	once_euclideanRingLazy.Do(func() {
		cache_euclideanRingLazy = gopurs_runtime.Func(func(dictEuclideanRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_euclideanRingLazy(dictEuclideanRing_0_box)
})
	})
	return cache_euclideanRingLazy
}

var cache_boundedLazy gopurs_runtime.Value
var once_boundedLazy sync.Once
func Get_boundedLazy() gopurs_runtime.Value {
	once_boundedLazy.Do(func() {
		cache_boundedLazy = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_boundedLazy(dictBounded_0_box)
})
	})
	return cache_boundedLazy
}

var cache_applyLazy gopurs_runtime.Value
var once_applyLazy sync.Once
func Get_applyLazy() gopurs_runtime.Value {
	once_applyLazy.Do(func() {
		cache_applyLazy = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_force(), f_0, gopurs_runtime.Apply(Get_force(), x_1))
}))
})
}))
	})
	return cache_applyLazy
}

var cache_bindLazy gopurs_runtime.Value
var once_bindLazy sync.Once
func Get_bindLazy() gopurs_runtime.Value {
	once_bindLazy.Do(func() {
		cache_bindLazy = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
}), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_0)))
}))
})
}))
	})
	return cache_bindLazy
}

var cache_heytingAlgebraLazy gopurs_runtime.Value
var once_heytingAlgebraLazy sync.Once
func Get_heytingAlgebraLazy() gopurs_runtime.Value {
	once_heytingAlgebraLazy.Do(func() {
		cache_heytingAlgebraLazy = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_heytingAlgebraLazy(dictHeytingAlgebra_0_box)
})
	})
	return cache_heytingAlgebraLazy
}

var cache_booleanAlgebraLazy gopurs_runtime.Value
var once_booleanAlgebraLazy sync.Once
func Get_booleanAlgebraLazy() gopurs_runtime.Value {
	once_booleanAlgebraLazy.Do(func() {
		cache_booleanAlgebraLazy = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_booleanAlgebraLazy(dictBooleanAlgebra_0_box)
})
	})
	return cache_booleanAlgebraLazy
}

var cache_applicativeLazy gopurs_runtime.Value
var once_applicativeLazy sync.Once
func Get_applicativeLazy() gopurs_runtime.Value {
	once_applicativeLazy.Do(func() {
		cache_applicativeLazy = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}))
	})
	return cache_applicativeLazy
}

var cache_monadLazy gopurs_runtime.Value
var once_monadLazy sync.Once
func Get_monadLazy() gopurs_runtime.Value {
	once_monadLazy.Do(func() {
		cache_monadLazy = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindLazy()
}))
	})
	return cache_monadLazy
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

var cache_apply__2140510474 gopurs_runtime.Value
var once_apply__2140510474 sync.Once
func Get_apply__2140510474() gopurs_runtime.Value {
	once_apply__2140510474.Do(func() {
		cache_apply__2140510474 = gopurs_runtime.RecordGet(Get_applyLazy(), "apply")
	})
	return cache_apply__2140510474
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

var cache_bottom__338427193 gopurs_runtime.Value
var once_bottom__338427193 sync.Once
func Get_bottom__338427193() gopurs_runtime.Value {
	once_bottom__338427193.Do(func() {
		cache_bottom__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_bottom__338427193(dict_0_box)
})
	})
	return cache_bottom__338427193
}

var cache_top__338427193 gopurs_runtime.Value
var once_top__338427193 sync.Once
func Get_top__338427193() gopurs_runtime.Value {
	once_top__338427193.Do(func() {
		cache_top__338427193 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_top__338427193(dict_0_box)
})
	})
	return cache_top__338427193
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

var cache_div__2579358968 gopurs_runtime.Value
var once_div__2579358968 sync.Once
func Get_div__2579358968() gopurs_runtime.Value {
	once_div__2579358968.Do(func() {
		cache_div__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_div__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_div__2579358968
}

var cache_mod__2579358968 gopurs_runtime.Value
var once_mod__2579358968 sync.Once
func Get_mod__2579358968() gopurs_runtime.Value {
	once_mod__2579358968.Do(func() {
		cache_mod__2579358968 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mod__2579358968(gopurs_runtime.CoerceToStruct[pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mod__2579358968
}

var cache_const__3524684546 gopurs_runtime.Value
var once_const__3524684546 sync.Once
func Get_const__3524684546() gopurs_runtime.Value {
	once_const__3524684546.Do(func() {
		cache_const__3524684546 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_const__3524684546(a_0_box, v_1_box)
})
	})
	return cache_const__3524684546
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

var cache_map__1132949076 gopurs_runtime.Value
var once_map__1132949076 sync.Once
func Get_map__1132949076() gopurs_runtime.Value {
	once_map__1132949076.Do(func() {
		cache_map__1132949076 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_map__1132949076(gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_map__1132949076
}

var cache_map__1510739772 gopurs_runtime.Value
var once_map__1510739772 sync.Once
func Get_map__1510739772() gopurs_runtime.Value {
	once_map__1510739772.Do(func() {
		cache_map__1510739772 = gopurs_runtime.RecordGet(Get_functorLazy(), "map")
	})
	return cache_map__1510739772
}

var cache_map__3467322428 gopurs_runtime.Value
var once_map__3467322428 sync.Once
func Get_map__3467322428() gopurs_runtime.Value {
	once_map__3467322428.Do(func() {
		cache_map__3467322428 = gopurs_runtime.RecordGet(Get_functorLazy(), "map")
	})
	return cache_map__3467322428
}

var cache_ff__2527024921 gopurs_runtime.Value
var once_ff__2527024921 sync.Once
func Get_ff__2527024921() gopurs_runtime.Value {
	once_ff__2527024921.Do(func() {
		cache_ff__2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ff__2527024921(dict_0_box)
})
	})
	return cache_ff__2527024921
}

var cache_tt__2527024921 gopurs_runtime.Value
var once_tt__2527024921 sync.Once
func Get_tt__2527024921() gopurs_runtime.Value {
	once_tt__2527024921.Do(func() {
		cache_tt__2527024921 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_tt__2527024921(dict_0_box)
})
	})
	return cache_tt__2527024921
}

var cache_applicativeLazy__3467920360 gopurs_runtime.Value
var once_applicativeLazy__3467920360 sync.Once
func Get_applicativeLazy__3467920360() gopurs_runtime.Value {
	once_applicativeLazy__3467920360.Do(func() {
		cache_applicativeLazy__3467920360 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}))
	})
	return cache_applicativeLazy__3467920360
}

var cache_applyLazy__879424557 gopurs_runtime.Value
var once_applyLazy__879424557 sync.Once
func Get_applyLazy__879424557() gopurs_runtime.Value {
	once_applyLazy__879424557.Do(func() {
		cache_applyLazy__879424557 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_force(), f_0, gopurs_runtime.Apply(Get_force(), x_1))
}))
})
}))
	})
	return cache_applyLazy__879424557
}

var cache_applyLazy__225241115 gopurs_runtime.Value
var once_applyLazy__225241115 sync.Once
func Get_applyLazy__225241115() gopurs_runtime.Value {
	once_applyLazy__225241115.Do(func() {
		cache_applyLazy__225241115 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_force(), f_0, gopurs_runtime.Apply(Get_force(), x_1))
}))
})
}))
	})
	return cache_applyLazy__225241115
}

var cache_bindLazy__1994192487 gopurs_runtime.Value
var once_bindLazy__1994192487 sync.Once
func Get_bindLazy__1994192487() gopurs_runtime.Value {
	once_bindLazy__1994192487.Do(func() {
		cache_bindLazy__1994192487 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
}), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_0)))
}))
})
}))
	})
	return cache_bindLazy__1994192487
}

var cache_eq1Lazy__251633054 gopurs_runtime.Value
var once_eq1Lazy__251633054 sync.Once
func Get_eq1Lazy__251633054() gopurs_runtime.Value {
	once_eq1Lazy__251633054.Do(func() {
		cache_eq1Lazy__251633054 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_eqLazy(dictEq_0), "eq")
}))
	})
	return cache_eq1Lazy__251633054
}

var cache_extendLazy__2169161609 gopurs_runtime.Value
var once_extendLazy__2169161609 sync.Once
func Get_extendLazy__2169161609() gopurs_runtime.Value {
	once_extendLazy__2169161609.Do(func() {
		cache_extendLazy__2169161609 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_1)
}))
})
}))
	})
	return cache_extendLazy__2169161609
}

var cache_foldable1Lazy__1238235135 gopurs_runtime.Value
var once_foldable1Lazy__1238235135 sync.Once
func Get_foldable1Lazy__1238235135() gopurs_runtime.Value {
	once_foldable1Lazy__1238235135.Do(func() {
		cache_foldable1Lazy__1238235135 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2))
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), l_1)
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), l_1)
})
}))
	})
	return cache_foldable1Lazy__1238235135
}

var cache_foldableLazy__3814277777 gopurs_runtime.Value
var once_foldableLazy__3814277777 sync.Once
func Get_foldableLazy__3814277777() gopurs_runtime.Value {
	once_foldableLazy__3814277777.Do(func() {
		cache_foldableLazy__3814277777 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.Apply(Get_force(), l_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(Get_force(), l_2), z_1)
})
})
}))
	})
	return cache_foldableLazy__3814277777
}

var cache_foldableWithIndexLazy__2458991819 gopurs_runtime.Value
var once_foldableWithIndexLazy__2458991819 sync.Once
func Get_foldableWithIndexLazy__2458991819() gopurs_runtime.Value {
	once_foldableWithIndexLazy__2458991819.Do(func() {
		cache_foldableWithIndexLazy__2458991819 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableLazy(), "foldMap"), dictMonoid_0)
_ = foldMap1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap1_1_0, gopurs_runtime.Apply(f_2, pkg_Data_Unit.Get_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableLazy(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableLazy(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_foldableWithIndexLazy__2458991819
}

var cache_functorLazy__491347738 gopurs_runtime.Value
var once_functorLazy__491347738 sync.Once
func Get_functorLazy__491347738() gopurs_runtime.Value {
	once_functorLazy__491347738.Do(func() {
		cache_functorLazy__491347738 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_force(), l_1))
}))
})
}))
	})
	return cache_functorLazy__491347738
}

var cache_functorLazy__3988504945 gopurs_runtime.Value
var once_functorLazy__3988504945 sync.Once
func Get_functorLazy__3988504945() gopurs_runtime.Value {
	once_functorLazy__3988504945.Do(func() {
		cache_functorLazy__3988504945 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_force(), l_1))
}))
})
}))
	})
	return cache_functorLazy__3988504945
}

var cache_functorWithIndexLazy__3312587351 gopurs_runtime.Value
var once_functorWithIndexLazy__3312587351 sync.Once
func Get_functorWithIndexLazy__3312587351() gopurs_runtime.Value {
	once_functorWithIndexLazy__3312587351.Do(func() {
		cache_functorWithIndexLazy__3312587351 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))
	})
	return cache_functorWithIndexLazy__3312587351
}

var cache_traversableLazy__1395024506 gopurs_runtime.Value
var once_traversableLazy__1395024506 sync.Once
func Get_traversableLazy__1395024506() gopurs_runtime.Value {
	once_traversableLazy__1395024506.Do(func() {
		cache_traversableLazy__1395024506 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
}), gopurs_runtime.Apply(Get_force(), l_2))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Functor0_1_1 := gopurs_runtime.CoerceToStruct[pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}))
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(Get_force(), l_3)))
})
})
}))
	})
	return cache_traversableLazy__1395024506
}

var cache_mempty__2312420373 gopurs_runtime.Value
var once_mempty__2312420373 sync.Once
func Get_mempty__2312420373() gopurs_runtime.Value {
	once_mempty__2312420373.Do(func() {
		cache_mempty__2312420373 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mempty__2312420373(dict_0_box)
})
	})
	return cache_mempty__2312420373
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

var cache_sub__3675938712 gopurs_runtime.Value
var once_sub__3675938712 sync.Once
func Get_sub__3675938712() gopurs_runtime.Value {
	once_sub__3675938712.Do(func() {
		cache_sub__3675938712 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_sub__3675938712(gopurs_runtime.CoerceToStruct[pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_sub__3675938712
}

var cache_append__493084344 gopurs_runtime.Value
var once_append__493084344 sync.Once
func Get_append__493084344() gopurs_runtime.Value {
	once_append__493084344.Do(func() {
		cache_append__493084344 = gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append")
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

var cache_add__1614463960 gopurs_runtime.Value
var once_add__1614463960 sync.Once
func Get_add__1614463960() gopurs_runtime.Value {
	once_add__1614463960.Do(func() {
		cache_add__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_add__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_add__1614463960
}

var cache_mul__1614463960 gopurs_runtime.Value
var once_mul__1614463960 sync.Once
func Get_mul__1614463960() gopurs_runtime.Value {
	once_mul__1614463960.Do(func() {
		cache_mul__1614463960 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_mul__1614463960(gopurs_runtime.CoerceToStruct[pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_mul__1614463960
}

var cache_one__1204848985 gopurs_runtime.Value
var once_one__1204848985 sync.Once
func Get_one__1204848985() gopurs_runtime.Value {
	once_one__1204848985.Do(func() {
		cache_one__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_one__1204848985(dict_0_box)
})
	})
	return cache_one__1204848985
}

var cache_zero__1204848985 gopurs_runtime.Value
var once_zero__1204848985 sync.Once
func Get_zero__1204848985() gopurs_runtime.Value {
	once_zero__1204848985.Do(func() {
		cache_zero__1204848985 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_zero__1204848985(dict_0_box)
})
	})
	return cache_zero__1204848985
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

func Call_showLazy(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(defer \\_ -> "), gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Apply(Get_force(), x_1)), gopurs_runtime.Str(")"))).StrVal())
}))
}

func Call_semiringLazy(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), gopurs_runtime.Apply(Get_force(), a_1), gopurs_runtime.Apply(Get_force(), b_2))
}))
})
}), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), gopurs_runtime.Apply(Get_force(), a_1), gopurs_runtime.Apply(Get_force(), b_2))
}))
})
}), gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictSemiring_0, "one")
})), gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictSemiring_0, "zero")
})))
}

func Call_semigroupLazy(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(Get_force(), a_1), gopurs_runtime.Apply(Get_force(), b_2))
}))
})
}))
}

func Call_ringLazy(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
semiringLazy1_1_0 := Call_semiringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{}))
_ = semiringLazy1_1_0
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringLazy1_1_0
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(Get_force(), a_2), gopurs_runtime.Apply(Get_force(), b_3))
}))
})
}))
}

func Call_monoidLazy(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
semigroupLazy1_1_0 := Call_semigroupLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupLazy1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupLazy1_1_0
}), gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
})))
}

func Call_eqLazy(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Apply(Get_force(), x_1), gopurs_runtime.Apply(Get_force(), y_2)).IntVal) != (0))
})
}))
}

func Call_ordLazy(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
eqLazy1_1_0 := Call_eqLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqLazy1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqLazy1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(Get_force(), x_2), gopurs_runtime.Apply(Get_force(), y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_commutativeRingLazy(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
ringLazy1_1_0 := Call_ringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{}))
_ = ringLazy1_1_0
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_0
}))
}

func Call_euclideanRingLazy(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
ringLazy1_1_1 := Call_ringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_0, "CommutativeRing0"), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}))
_ = ringLazy1_1_1
commutativeRingLazy1_1_0 := gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_1
}))
_ = commutativeRingLazy1_1_0
return gopurs_runtime.RecordDict4("CommutativeRing0", "degree", "div", "mod", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return commutativeRingLazy1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_0, "degree"), gopurs_runtime.Apply(Get_force(), x_2))
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "div"), gopurs_runtime.Apply(Get_force(), a_2), gopurs_runtime.Apply(Get_force(), b_3))
}))
})
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "mod"), gopurs_runtime.Apply(Get_force(), a_2), gopurs_runtime.Apply(Get_force(), b_3))
}))
})
}))
}

func Call_boundedLazy(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
ordLazy1_1_0 := Call_ordLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordLazy1_1_0
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ordLazy1_1_0
}), gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBounded_0, "bottom")
})), gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBounded_0, "top")
})))
}

func Call_heytingAlgebraLazy(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
implies_1_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies")
_ = implies_1_0
conj_2_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj")
_ = conj_2_1
disj_3_2 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj")
_ = disj_3_2
not_4_3 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not")
_ = not_4_3
return gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), conj_2_1, a_5), b_6)
})
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), disj_3_2, a_5), b_6)
})
}), gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
})), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), implies_1_0, a_5), b_6)
})
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), not_4_3, a_5)
}), gopurs_runtime.Apply(Get_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
}))})
}

func Call_booleanAlgebraLazy(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
heytingAlgebraLazy1_1_0 := Call_heytingAlgebraLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraLazy1_1_0
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraLazy1_1_0
}))
}

func Call_apply__353515660(dict_0_loop *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_bottom__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "bottom")
}

func Call_top__338427193(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "top")
}

func Call_eq__2384498378(dict_0_loop *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Eq.Constructor_Eq[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_div__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V2
}

func Call_mod__2579358968(dict_0_loop *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_EuclideanRing.Constructor_EuclideanRing[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V3
}

func Call_const__3524684546(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_map__2199395572(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_map__1132949076(dict_0_loop *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Functor.Constructor_Functor[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_ff__2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "ff")
}

func Call_tt__2527024921(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "tt")
}

func Call_mempty__2312420373(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "mempty")
}

func Call_compare__821463600(dict_0_loop *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ord.Constructor_Ord[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_sub__3675938712(dict_0_loop *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Ring.Constructor_Ring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_add__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_mul__1614463960(dict_0_loop *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semiring.Constructor_Semiring[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_one__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "one")
}

func Call_zero__1204848985(dict_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dict_0 gopurs_runtime.Value = dict_0_loop
_ = dict_0
return gopurs_runtime.RecordGet(dict_0, "zero")
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Get_go__defer() gopurs_runtime.Value {
	return _Gopurs_Go__defer
}

func Get_force() gopurs_runtime.Value {
	return _Gopurs_Force
}

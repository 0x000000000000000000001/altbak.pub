package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
)

var cache_Data_Lazy_showLazy gopurs_runtime.Value
var once_Data_Lazy_showLazy sync.Once
func Get_Data_Lazy_showLazy() gopurs_runtime.Value {
	once_Data_Lazy_showLazy.Do(func() {
		cache_Data_Lazy_showLazy = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_showLazy(dictShow_0_box)
})
	})
	return cache_Data_Lazy_showLazy
}

var cache_Data_Lazy_semiringLazy gopurs_runtime.Value
var once_Data_Lazy_semiringLazy sync.Once
func Get_Data_Lazy_semiringLazy() gopurs_runtime.Value {
	once_Data_Lazy_semiringLazy.Do(func() {
		cache_Data_Lazy_semiringLazy = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_semiringLazy(dictSemiring_0_box)
})
	})
	return cache_Data_Lazy_semiringLazy
}

var cache_Data_Lazy_semigroupLazy gopurs_runtime.Value
var once_Data_Lazy_semigroupLazy sync.Once
func Get_Data_Lazy_semigroupLazy() gopurs_runtime.Value {
	once_Data_Lazy_semigroupLazy.Do(func() {
		cache_Data_Lazy_semigroupLazy = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_semigroupLazy(dictSemigroup_0_box)
})
	})
	return cache_Data_Lazy_semigroupLazy
}

var cache_Data_Lazy_ringLazy gopurs_runtime.Value
var once_Data_Lazy_ringLazy sync.Once
func Get_Data_Lazy_ringLazy() gopurs_runtime.Value {
	once_Data_Lazy_ringLazy.Do(func() {
		cache_Data_Lazy_ringLazy = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_ringLazy(dictRing_0_box)
})
	})
	return cache_Data_Lazy_ringLazy
}

var cache_Data_Lazy_monoidLazy gopurs_runtime.Value
var once_Data_Lazy_monoidLazy sync.Once
func Get_Data_Lazy_monoidLazy() gopurs_runtime.Value {
	once_Data_Lazy_monoidLazy.Do(func() {
		cache_Data_Lazy_monoidLazy = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_monoidLazy(dictMonoid_0_box)
})
	})
	return cache_Data_Lazy_monoidLazy
}

var cache_Data_Lazy_lazyLazy gopurs_runtime.Value
var once_Data_Lazy_lazyLazy sync.Once
func Get_Data_Lazy_lazyLazy() gopurs_runtime.Value {
	once_Data_Lazy_lazyLazy.Do(func() {
		cache_Data_Lazy_lazyLazy = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
}))
	})
	return cache_Data_Lazy_lazyLazy
}

var cache_Data_Lazy_functorLazy gopurs_runtime.Value
var once_Data_Lazy_functorLazy sync.Once
func Get_Data_Lazy_functorLazy() gopurs_runtime.Value {
	once_Data_Lazy_functorLazy.Do(func() {
		cache_Data_Lazy_functorLazy = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
})
}))
	})
	return cache_Data_Lazy_functorLazy
}

var cache_Data_Lazy_go__map gopurs_runtime.Value
var once_Data_Lazy_go__map sync.Once
func Get_Data_Lazy_go__map() gopurs_runtime.Value {
	once_Data_Lazy_go__map.Do(func() {
		cache_Data_Lazy_go__map = gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map")
	})
	return cache_Data_Lazy_go__map
}

var cache_Data_Lazy_functorWithIndexLazy gopurs_runtime.Value
var once_Data_Lazy_functorWithIndexLazy sync.Once
func Get_Data_Lazy_functorWithIndexLazy() gopurs_runtime.Value {
	once_Data_Lazy_functorWithIndexLazy.Do(func() {
		cache_Data_Lazy_functorWithIndexLazy = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_Lazy_functorWithIndexLazy
}

var cache_Data_Lazy_invariantLazy gopurs_runtime.Value
var once_Data_Lazy_invariantLazy sync.Once
func Get_Data_Lazy_invariantLazy() gopurs_runtime.Value {
	once_Data_Lazy_invariantLazy.Do(func() {
		cache_Data_Lazy_invariantLazy = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), f_0)
})
}))
	})
	return cache_Data_Lazy_invariantLazy
}

var cache_Data_Lazy_foldableLazy gopurs_runtime.Value
var once_Data_Lazy_foldableLazy sync.Once
func Get_Data_Lazy_foldableLazy() gopurs_runtime.Value {
	once_Data_Lazy_foldableLazy.Do(func() {
		cache_Data_Lazy_foldableLazy = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2), z_1)
})
})
}))
	})
	return cache_Data_Lazy_foldableLazy
}

var cache_Data_Lazy_foldr gopurs_runtime.Value
var once_Data_Lazy_foldr sync.Once
func Get_Data_Lazy_foldr() gopurs_runtime.Value {
	once_Data_Lazy_foldr.Do(func() {
		cache_Data_Lazy_foldr = gopurs_runtime.RecordGet(Get_Data_Lazy_foldableLazy(), "foldr")
	})
	return cache_Data_Lazy_foldr
}

var cache_Data_Lazy_foldl gopurs_runtime.Value
var once_Data_Lazy_foldl sync.Once
func Get_Data_Lazy_foldl() gopurs_runtime.Value {
	once_Data_Lazy_foldl.Do(func() {
		cache_Data_Lazy_foldl = gopurs_runtime.RecordGet(Get_Data_Lazy_foldableLazy(), "foldl")
	})
	return cache_Data_Lazy_foldl
}

var cache_Data_Lazy_foldMap gopurs_runtime.Value
var once_Data_Lazy_foldMap sync.Once
func Get_Data_Lazy_foldMap() gopurs_runtime.Value {
	once_Data_Lazy_foldMap.Do(func() {
		cache_Data_Lazy_foldMap = gopurs_runtime.RecordGet(Get_Data_Lazy_foldableLazy(), "foldMap")
	})
	return cache_Data_Lazy_foldMap
}

var cache_Data_Lazy_foldableWithIndexLazy gopurs_runtime.Value
var once_Data_Lazy_foldableWithIndexLazy sync.Once
func Get_Data_Lazy_foldableWithIndexLazy() gopurs_runtime.Value {
	once_Data_Lazy_foldableWithIndexLazy.Do(func() {
		cache_Data_Lazy_foldableWithIndexLazy = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_foldableLazy()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap1_1_0 -> gopurs_runtime.Value
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_foldableLazy(), "foldMap"), dictMonoid_0)
_ = foldMap1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap1_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_foldableLazy(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_foldableLazy(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_Lazy_foldableWithIndexLazy
}

var cache_Data_Lazy_traversableLazy gopurs_runtime.Value
var once_Data_Lazy_traversableLazy sync.Once
func Get_Data_Lazy_traversableLazy() gopurs_runtime.Value {
	once_Data_Lazy_traversableLazy.Do(func() {
		cache_Data_Lazy_traversableLazy = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_foldableLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorLazy()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}))
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_3)))
})
})
}))
	})
	return cache_Data_Lazy_traversableLazy
}

var cache_Data_Lazy_traverse gopurs_runtime.Value
var once_Data_Lazy_traverse sync.Once
func Get_Data_Lazy_traverse() gopurs_runtime.Value {
	once_Data_Lazy_traverse.Do(func() {
		cache_Data_Lazy_traverse = gopurs_runtime.RecordGet(Get_Data_Lazy_traversableLazy(), "traverse")
	})
	return cache_Data_Lazy_traverse
}

var cache_Data_Lazy_traversableWithIndexLazy gopurs_runtime.Value
var once_Data_Lazy_traversableWithIndexLazy sync.Once
func Get_Data_Lazy_traversableWithIndexLazy() gopurs_runtime.Value {
	once_Data_Lazy_traversableWithIndexLazy.Do(func() {
		cache_Data_Lazy_traversableWithIndexLazy = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_foldableWithIndexLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorWithIndexLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_traversableLazy()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse1_1_0 -> gopurs_runtime.Value
traverse1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_traversableLazy(), "traverse"), dictApplicative_0)
_ = traverse1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse1_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}))
	})
	return cache_Data_Lazy_traversableWithIndexLazy
}

var cache_Data_Lazy_foldable1Lazy gopurs_runtime.Value
var once_Data_Lazy_foldable1Lazy sync.Once
func Get_Data_Lazy_foldable1Lazy() gopurs_runtime.Value {
	once_Data_Lazy_foldable1Lazy.Do(func() {
		cache_Data_Lazy_foldable1Lazy = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_foldableLazy()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)
})
}))
	})
	return cache_Data_Lazy_foldable1Lazy
}

var cache_Data_Lazy_traversable1Lazy gopurs_runtime.Value
var once_Data_Lazy_traversable1Lazy sync.Once
func Get_Data_Lazy_traversable1Lazy() gopurs_runtime.Value {
	once_Data_Lazy_traversable1Lazy.Do(func() {
		cache_Data_Lazy_traversable1Lazy = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_foldable1Lazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_traversableLazy()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}))
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_3)))
})
})
}))
	})
	return cache_Data_Lazy_traversable1Lazy
}

var cache_Data_Lazy_extendLazy gopurs_runtime.Value
var once_Data_Lazy_extendLazy sync.Once
func Get_Data_Lazy_extendLazy() gopurs_runtime.Value {
	once_Data_Lazy_extendLazy.Do(func() {
		cache_Data_Lazy_extendLazy = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_1)
}))
})
}))
	})
	return cache_Data_Lazy_extendLazy
}

var cache_Data_Lazy_eqLazy gopurs_runtime.Value
var once_Data_Lazy_eqLazy sync.Once
func Get_Data_Lazy_eqLazy() gopurs_runtime.Value {
	once_Data_Lazy_eqLazy.Do(func() {
		cache_Data_Lazy_eqLazy = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_eqLazy(dictEq_0_box)
})
	})
	return cache_Data_Lazy_eqLazy
}

var cache_Data_Lazy_ordLazy gopurs_runtime.Value
var once_Data_Lazy_ordLazy sync.Once
func Get_Data_Lazy_ordLazy() gopurs_runtime.Value {
	once_Data_Lazy_ordLazy.Do(func() {
		cache_Data_Lazy_ordLazy = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_ordLazy(dictOrd_0_box)
})
	})
	return cache_Data_Lazy_ordLazy
}

var cache_Data_Lazy_eq1Lazy gopurs_runtime.Value
var once_Data_Lazy_eq1Lazy sync.Once
func Get_Data_Lazy_eq1Lazy() gopurs_runtime.Value {
	once_Data_Lazy_eq1Lazy.Do(func() {
		cache_Data_Lazy_eq1Lazy = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_Data_Lazy_eqLazy(dictEq_0), "eq")
}))
	})
	return cache_Data_Lazy_eq1Lazy
}

var cache_Data_Lazy_ord1Lazy gopurs_runtime.Value
var once_Data_Lazy_ord1Lazy sync.Once
func Get_Data_Lazy_ord1Lazy() gopurs_runtime.Value {
	once_Data_Lazy_ord1Lazy.Do(func() {
		cache_Data_Lazy_ord1Lazy = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_eq1Lazy()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_Data_Lazy_ordLazy(dictOrd_0), "compare")
}))
	})
	return cache_Data_Lazy_ord1Lazy
}

var cache_Data_Lazy_comonadLazy gopurs_runtime.Value
var once_Data_Lazy_comonadLazy sync.Once
func Get_Data_Lazy_comonadLazy() gopurs_runtime.Value {
	once_Data_Lazy_comonadLazy.Do(func() {
		cache_Data_Lazy_comonadLazy = gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_extendLazy()
}), Get_Data_Lazy_force())
	})
	return cache_Data_Lazy_comonadLazy
}

var cache_Data_Lazy_commutativeRingLazy gopurs_runtime.Value
var once_Data_Lazy_commutativeRingLazy sync.Once
func Get_Data_Lazy_commutativeRingLazy() gopurs_runtime.Value {
	once_Data_Lazy_commutativeRingLazy.Do(func() {
		cache_Data_Lazy_commutativeRingLazy = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_commutativeRingLazy(dictCommutativeRing_0_box)
})
	})
	return cache_Data_Lazy_commutativeRingLazy
}

var cache_Data_Lazy_euclideanRingLazy gopurs_runtime.Value
var once_Data_Lazy_euclideanRingLazy sync.Once
func Get_Data_Lazy_euclideanRingLazy() gopurs_runtime.Value {
	once_Data_Lazy_euclideanRingLazy.Do(func() {
		cache_Data_Lazy_euclideanRingLazy = gopurs_runtime.Func(func(dictEuclideanRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_euclideanRingLazy(dictEuclideanRing_0_box)
})
	})
	return cache_Data_Lazy_euclideanRingLazy
}

var cache_Data_Lazy_boundedLazy gopurs_runtime.Value
var once_Data_Lazy_boundedLazy sync.Once
func Get_Data_Lazy_boundedLazy() gopurs_runtime.Value {
	once_Data_Lazy_boundedLazy.Do(func() {
		cache_Data_Lazy_boundedLazy = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_boundedLazy(dictBounded_0_box)
})
	})
	return cache_Data_Lazy_boundedLazy
}

var cache_Data_Lazy_applyLazy gopurs_runtime.Value
var once_Data_Lazy_applyLazy sync.Once
func Get_Data_Lazy_applyLazy() gopurs_runtime.Value {
	once_Data_Lazy_applyLazy.Do(func() {
		cache_Data_Lazy_applyLazy = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
}))
})
}))
	})
	return cache_Data_Lazy_applyLazy
}

var cache_Data_Lazy_bindLazy gopurs_runtime.Value
var once_Data_Lazy_bindLazy sync.Once
func Get_Data_Lazy_bindLazy() gopurs_runtime.Value {
	once_Data_Lazy_bindLazy.Do(func() {
		cache_Data_Lazy_bindLazy = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_applyLazy()
}), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_0)))
}))
})
}))
	})
	return cache_Data_Lazy_bindLazy
}

var cache_Data_Lazy_heytingAlgebraLazy gopurs_runtime.Value
var once_Data_Lazy_heytingAlgebraLazy sync.Once
func Get_Data_Lazy_heytingAlgebraLazy() gopurs_runtime.Value {
	once_Data_Lazy_heytingAlgebraLazy.Do(func() {
		cache_Data_Lazy_heytingAlgebraLazy = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_heytingAlgebraLazy(dictHeytingAlgebra_0_box)
})
	})
	return cache_Data_Lazy_heytingAlgebraLazy
}

var cache_Data_Lazy_booleanAlgebraLazy gopurs_runtime.Value
var once_Data_Lazy_booleanAlgebraLazy sync.Once
func Get_Data_Lazy_booleanAlgebraLazy() gopurs_runtime.Value {
	once_Data_Lazy_booleanAlgebraLazy.Do(func() {
		cache_Data_Lazy_booleanAlgebraLazy = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Lazy_booleanAlgebraLazy(dictBooleanAlgebra_0_box)
})
	})
	return cache_Data_Lazy_booleanAlgebraLazy
}

var cache_Data_Lazy_applicativeLazy gopurs_runtime.Value
var once_Data_Lazy_applicativeLazy sync.Once
func Get_Data_Lazy_applicativeLazy() gopurs_runtime.Value {
	once_Data_Lazy_applicativeLazy.Do(func() {
		cache_Data_Lazy_applicativeLazy = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_applyLazy()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}))
	})
	return cache_Data_Lazy_applicativeLazy
}

var cache_Data_Lazy_monadLazy gopurs_runtime.Value
var once_Data_Lazy_monadLazy sync.Once
func Get_Data_Lazy_monadLazy() gopurs_runtime.Value {
	once_Data_Lazy_monadLazy.Do(func() {
		cache_Data_Lazy_monadLazy = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_applicativeLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_bindLazy()
}))
	})
	return cache_Data_Lazy_monadLazy
}

var cache_Data_Lazy_applicativeLazy__690919725 gopurs_runtime.Value
var once_Data_Lazy_applicativeLazy__690919725 sync.Once
func Get_Data_Lazy_applicativeLazy__690919725() gopurs_runtime.Value {
	once_Data_Lazy_applicativeLazy__690919725.Do(func() {
		cache_Data_Lazy_applicativeLazy__690919725 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_applyLazy()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}))
	})
	return cache_Data_Lazy_applicativeLazy__690919725
}

var cache_Data_Lazy_applicativeLazy__3467920360 gopurs_runtime.Value
var once_Data_Lazy_applicativeLazy__3467920360 sync.Once
func Get_Data_Lazy_applicativeLazy__3467920360() gopurs_runtime.Value {
	once_Data_Lazy_applicativeLazy__3467920360.Do(func() {
		cache_Data_Lazy_applicativeLazy__3467920360 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_applyLazy()
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}))
	})
	return cache_Data_Lazy_applicativeLazy__3467920360
}

var cache_Data_Lazy_applyLazy__879424557 gopurs_runtime.Value
var once_Data_Lazy_applyLazy__879424557 sync.Once
func Get_Data_Lazy_applyLazy__879424557() gopurs_runtime.Value {
	once_Data_Lazy_applyLazy__879424557.Do(func() {
		cache_Data_Lazy_applyLazy__879424557 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
}))
})
}))
	})
	return cache_Data_Lazy_applyLazy__879424557
}

var cache_Data_Lazy_applyLazy__225241115 gopurs_runtime.Value
var once_Data_Lazy_applyLazy__225241115 sync.Once
func Get_Data_Lazy_applyLazy__225241115() gopurs_runtime.Value {
	once_Data_Lazy_applyLazy__225241115.Do(func() {
		cache_Data_Lazy_applyLazy__225241115 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
}))
})
}))
	})
	return cache_Data_Lazy_applyLazy__225241115
}

var cache_Data_Lazy_bindLazy__1994192487 gopurs_runtime.Value
var once_Data_Lazy_bindLazy__1994192487 sync.Once
func Get_Data_Lazy_bindLazy__1994192487() gopurs_runtime.Value {
	once_Data_Lazy_bindLazy__1994192487.Do(func() {
		cache_Data_Lazy_bindLazy__1994192487 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_applyLazy()
}), gopurs_runtime.Func(func(l_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_0)))
}))
})
}))
	})
	return cache_Data_Lazy_bindLazy__1994192487
}

var cache_Data_Lazy_eq1Lazy__3789574347 gopurs_runtime.Value
var once_Data_Lazy_eq1Lazy__3789574347 sync.Once
func Get_Data_Lazy_eq1Lazy__3789574347() gopurs_runtime.Value {
	once_Data_Lazy_eq1Lazy__3789574347.Do(func() {
		cache_Data_Lazy_eq1Lazy__3789574347 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_Data_Lazy_eqLazy(dictEq_0), "eq")
}))
	})
	return cache_Data_Lazy_eq1Lazy__3789574347
}

var cache_Data_Lazy_eq1Lazy__251633054 gopurs_runtime.Value
var once_Data_Lazy_eq1Lazy__251633054 sync.Once
func Get_Data_Lazy_eq1Lazy__251633054() gopurs_runtime.Value {
	once_Data_Lazy_eq1Lazy__251633054.Do(func() {
		cache_Data_Lazy_eq1Lazy__251633054 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_Data_Lazy_eqLazy(dictEq_0), "eq")
}))
	})
	return cache_Data_Lazy_eq1Lazy__251633054
}

var cache_Data_Lazy_extendLazy__2169161609 gopurs_runtime.Value
var once_Data_Lazy_extendLazy__2169161609 sync.Once
func Get_Data_Lazy_extendLazy__2169161609() gopurs_runtime.Value {
	once_Data_Lazy_extendLazy__2169161609.Do(func() {
		cache_Data_Lazy_extendLazy__2169161609 = gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_1)
}))
})
}))
	})
	return cache_Data_Lazy_extendLazy__2169161609
}

var cache_Data_Lazy_foldable1Lazy__1238235135 gopurs_runtime.Value
var once_Data_Lazy_foldable1Lazy__1238235135 sync.Once
func Get_Data_Lazy_foldable1Lazy__1238235135() gopurs_runtime.Value {
	once_Data_Lazy_foldable1Lazy__1238235135.Do(func() {
		cache_Data_Lazy_foldable1Lazy__1238235135 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_foldableLazy()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)
})
}), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)
})
}))
	})
	return cache_Data_Lazy_foldable1Lazy__1238235135
}

var cache_Data_Lazy_foldableLazy__3814277777 gopurs_runtime.Value
var once_Data_Lazy_foldableLazy__3814277777 sync.Once
func Get_Data_Lazy_foldableLazy__3814277777() gopurs_runtime.Value {
	once_Data_Lazy_foldableLazy__3814277777.Do(func() {
		cache_Data_Lazy_foldableLazy__3814277777 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2), z_1)
})
})
}))
	})
	return cache_Data_Lazy_foldableLazy__3814277777
}

var cache_Data_Lazy_foldableWithIndexLazy__2458991819 gopurs_runtime.Value
var once_Data_Lazy_foldableWithIndexLazy__2458991819 sync.Once
func Get_Data_Lazy_foldableWithIndexLazy__2458991819() gopurs_runtime.Value {
	once_Data_Lazy_foldableWithIndexLazy__2458991819.Do(func() {
		cache_Data_Lazy_foldableWithIndexLazy__2458991819 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_foldableLazy()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap1_1_0 -> gopurs_runtime.Value
foldMap1_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_foldableLazy(), "foldMap"), dictMonoid_0)
_ = foldMap1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap1_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_foldableLazy(), "foldl"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_foldableLazy(), "foldr"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_Lazy_foldableWithIndexLazy__2458991819
}

var cache_Data_Lazy_functorLazy__491347738 gopurs_runtime.Value
var once_Data_Lazy_functorLazy__491347738 sync.Once
func Get_Data_Lazy_functorLazy__491347738() gopurs_runtime.Value {
	once_Data_Lazy_functorLazy__491347738.Do(func() {
		cache_Data_Lazy_functorLazy__491347738 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
})
}))
	})
	return cache_Data_Lazy_functorLazy__491347738
}

var cache_Data_Lazy_functorLazy__3988504945 gopurs_runtime.Value
var once_Data_Lazy_functorLazy__3988504945 sync.Once
func Get_Data_Lazy_functorLazy__3988504945() gopurs_runtime.Value {
	once_Data_Lazy_functorLazy__3988504945.Do(func() {
		cache_Data_Lazy_functorLazy__3988504945 = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
})
}))
	})
	return cache_Data_Lazy_functorLazy__3988504945
}

var cache_Data_Lazy_functorWithIndexLazy__3312587351 gopurs_runtime.Value
var once_Data_Lazy_functorWithIndexLazy__3312587351 sync.Once
func Get_Data_Lazy_functorWithIndexLazy__3312587351() gopurs_runtime.Value {
	once_Data_Lazy_functorWithIndexLazy__3312587351.Do(func() {
		cache_Data_Lazy_functorWithIndexLazy__3312587351 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
	})
	return cache_Data_Lazy_functorWithIndexLazy__3312587351
}

var cache_Data_Lazy_ord1Lazy__2079329387 gopurs_runtime.Value
var once_Data_Lazy_ord1Lazy__2079329387 sync.Once
func Get_Data_Lazy_ord1Lazy__2079329387() gopurs_runtime.Value {
	once_Data_Lazy_ord1Lazy__2079329387.Do(func() {
		cache_Data_Lazy_ord1Lazy__2079329387 = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_eq1Lazy()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(Call_Data_Lazy_ordLazy(dictOrd_0), "compare")
}))
	})
	return cache_Data_Lazy_ord1Lazy__2079329387
}

var cache_Data_Lazy_traversableLazy__1395024506 gopurs_runtime.Value
var once_Data_Lazy_traversableLazy__1395024506 sync.Once
func Get_Data_Lazy_traversableLazy__1395024506() gopurs_runtime.Value {
	once_Data_Lazy_traversableLazy__1395024506.Do(func() {
		cache_Data_Lazy_traversableLazy__1395024506 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_foldableLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Lazy_functorLazy()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
}), gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 -> *Constructor_Data_Functor_Functor[gopurs_runtime.Value]
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_1.V0), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return x_4
}))
}), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_3)))
})
})
}))
	})
	return cache_Data_Lazy_traversableLazy__1395024506
}

func Call_Data_Lazy_showLazy(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(defer \\_ -> ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)).StrVal())) + (")"))
}))
}

func Call_Data_Lazy_semiringLazy(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_2))
}))
})
}), gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_2))
}))
})
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictSemiring_0, "one")
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictSemiring_0, "zero")
})))
}

func Call_Data_Lazy_semigroupLazy(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func(func(a_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_2))
}))
})
}))
}

func Call_Data_Lazy_ringLazy(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
// TAST (Let): semiringLazy1_1_0 -> gopurs_runtime.Value
semiringLazy1_1_0 := Call_Data_Lazy_semiringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{}))
_ = semiringLazy1_1_0
return gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringLazy1_1_0
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_3))
}))
})
}))
}

func Call_Data_Lazy_monoidLazy(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): semigroupLazy1_1_0 -> gopurs_runtime.Value
semigroupLazy1_1_0 := Call_Data_Lazy_semigroupLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupLazy1_1_0
return gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupLazy1_1_0
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
})))
}

func Call_Data_Lazy_eqLazy(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1), gopurs_runtime.Apply(Get_Data_Lazy_force(), y_2)).IntVal) != (0))
})
}))
}

func Call_Data_Lazy_ordLazy(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqLazy1_1_0 -> gopurs_runtime.Value
eqLazy1_1_0 := Call_Data_Lazy_eqLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}))
_ = eqLazy1_1_0
return gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return eqLazy1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_Lazy_commutativeRingLazy(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
// TAST (Let): ringLazy1_1_0 -> gopurs_runtime.Value
ringLazy1_1_0 := Call_Data_Lazy_ringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{}))
_ = ringLazy1_1_0
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_0
}))
}

func Call_Data_Lazy_euclideanRingLazy(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
// TAST (Let): ringLazy1_1_1 -> gopurs_runtime.Value
ringLazy1_1_1 := Call_Data_Lazy_ringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_0, "CommutativeRing0"), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}))
_ = ringLazy1_1_1
// TAST (Let): commutativeRingLazy1_1_0 -> gopurs_runtime.Value
commutativeRingLazy1_1_0 := gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_1
}))
_ = commutativeRingLazy1_1_0
return gopurs_runtime.RecordDict4("CommutativeRing0", "degree", "div", "mod", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return commutativeRingLazy1_1_0
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_0, "degree"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2))
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "div"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_3))
}))
})
}), gopurs_runtime.Func(func(a_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "mod"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_3))
}))
})
}))
}

func Call_Data_Lazy_boundedLazy(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): ordLazy1_1_0 -> gopurs_runtime.Value
ordLazy1_1_0 := Call_Data_Lazy_ordLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordLazy1_1_0
return gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ordLazy1_1_0
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBounded_0, "bottom")
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBounded_0, "top")
})))
}

func Call_Data_Lazy_heytingAlgebraLazy(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
// TAST (Let): implies_1_0 -> gopurs_runtime.Value
implies_1_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies")
_ = implies_1_0
// TAST (Let): conj_2_1 -> gopurs_runtime.Value
conj_2_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj")
_ = conj_2_1
// TAST (Let): disj_3_2 -> gopurs_runtime.Value
disj_3_2 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj")
_ = disj_3_2
// TAST (Let): not_4_3 -> gopurs_runtime.Value
not_4_3 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not")
_ = not_4_3
return gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), conj_2_1, a_5), b_6)
})
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), disj_3_2, a_5), b_6)
})
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
})), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), implies_1_0, a_5), b_6)
})
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Lazy_functorLazy(), "map"), not_4_3, a_5)
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
}))})
}

func Call_Data_Lazy_booleanAlgebraLazy(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
// TAST (Let): heytingAlgebraLazy1_1_0 -> gopurs_runtime.Value
heytingAlgebraLazy1_1_0 := Call_Data_Lazy_heytingAlgebraLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraLazy1_1_0
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraLazy1_1_0
}))
}

func Get_Data_Lazy_go__defer() gopurs_runtime.Value {
	return _Gopurs_Data_Lazy_Go__defer
}

func Get_Data_Lazy_force() gopurs_runtime.Value {
	return _Gopurs_Data_Lazy_Force
}

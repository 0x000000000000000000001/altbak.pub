package Data_Lazy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
)

var cache_showLazy gopurs_runtime.Value
var once_showLazy sync.Once
func Get_showLazy() gopurs_runtime.Value {
	once_showLazy.Do(func() {
		cache_showLazy = gopurs_runtime.Func(func(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(defer \\_ -> ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Apply(Get_force(), x_1)).StrVal())) + (")"))
}))
}()
})
	})
	return cache_showLazy
}

var cache_semiringLazy gopurs_runtime.Value
var once_semiringLazy sync.Once
func Get_semiringLazy() gopurs_runtime.Value {
	once_semiringLazy.Do(func() {
		cache_semiringLazy = gopurs_runtime.Func(func(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
zero_1_0 := gopurs_runtime.RecordGet(dictSemiring_0, "zero")
_ = zero_1_0
one_2_1 := gopurs_runtime.RecordGet(dictSemiring_0, "one")
_ = one_2_1
return gopurs_runtime.RecordDict4("add", "zero", "mul", "one", gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), gopurs_runtime.Apply(Get_force(), a_3), gopurs_runtime.Apply(Get_force(), b_4))
}))
}), gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return zero_1_0
})), gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), gopurs_runtime.Apply(Get_force(), a_3), gopurs_runtime.Apply(Get_force(), b_4))
}))
}), gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return one_2_1
})))
}()
})
	})
	return cache_semiringLazy
}

var cache_semigroupLazy gopurs_runtime.Value
var once_semigroupLazy sync.Once
func Get_semigroupLazy() gopurs_runtime.Value {
	once_semigroupLazy.Do(func() {
		cache_semigroupLazy = gopurs_runtime.Func(func(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.RecordDict1("append", gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(Get_force(), a_1), gopurs_runtime.Apply(Get_force(), b_2))
}))
}))
}()
})
	})
	return cache_semigroupLazy
}

var cache_ringLazy gopurs_runtime.Value
var once_ringLazy sync.Once
func Get_ringLazy() gopurs_runtime.Value {
	once_ringLazy.Do(func() {
		cache_ringLazy = gopurs_runtime.Func(func(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
semiringLazy1_1_0 := gopurs_runtime.Apply(Get_semiringLazy(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{}))
_ = semiringLazy1_1_0
return gopurs_runtime.RecordDict2("sub", "Semiring0", gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(Get_force(), a_2), gopurs_runtime.Apply(Get_force(), b_3))
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringLazy1_1_0
}))
}()
})
	})
	return cache_ringLazy
}

var cache_monoidLazy gopurs_runtime.Value
var once_monoidLazy sync.Once
func Get_monoidLazy() gopurs_runtime.Value {
	once_monoidLazy.Do(func() {
		cache_monoidLazy = gopurs_runtime.Func(func(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
semigroupLazy1_2_1 := gopurs_runtime.Apply(Get_semigroupLazy(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = semigroupLazy1_2_1
return gopurs_runtime.RecordDict2("mempty", "Semigroup0", gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupLazy1_2_1
}))
}()
})
	})
	return cache_monoidLazy
}

var cache_lazyLazy gopurs_runtime.Value
var once_lazyLazy sync.Once
func Get_lazyLazy() gopurs_runtime.Value {
	once_lazyLazy.Do(func() {
		cache_lazyLazy = gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
		cache_functorLazy = gopurs_runtime.RecordDict1("map", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_force(), l_1))
}))
}))
	})
	return cache_functorLazy
}

var cache_functorWithIndexLazy gopurs_runtime.Value
var once_functorWithIndexLazy sync.Once
func Get_functorWithIndexLazy() gopurs_runtime.Value {
	once_functorWithIndexLazy.Do(func() {
		cache_functorWithIndexLazy = gopurs_runtime.RecordDict2("mapWithIndex", "Functor0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}))
	})
	return cache_functorWithIndexLazy
}

var cache_invariantLazy gopurs_runtime.Value
var once_invariantLazy sync.Once
func Get_invariantLazy() gopurs_runtime.Value {
	once_invariantLazy.Do(func() {
		cache_invariantLazy = gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), f_0)
}))
	})
	return cache_invariantLazy
}

var cache_foldableLazy gopurs_runtime.Value
var once_foldableLazy sync.Once
func Get_foldableLazy() gopurs_runtime.Value {
	once_foldableLazy.Do(func() {
		cache_foldableLazy = gopurs_runtime.RecordDict3("foldr", "foldl", "foldMap", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(Get_force(), l_2), z_1)
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.Apply(Get_force(), l_2))
}), gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2))
}))
	})
	return cache_foldableLazy
}

var cache_foldableWithIndexLazy gopurs_runtime.Value
var once_foldableWithIndexLazy sync.Once
func Get_foldableWithIndexLazy() gopurs_runtime.Value {
	once_foldableWithIndexLazy.Do(func() {
		cache_foldableWithIndexLazy = gopurs_runtime.RecordDict4("foldrWithIndex", "foldlWithIndex", "foldMapWithIndex", "Foldable0", gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableLazy(), "foldr"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldableLazy(), "foldl"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func2(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_foldableLazy(), "foldMap"), dictMonoid_0, gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}))
	})
	return cache_foldableWithIndexLazy
}

var cache_traversableLazy gopurs_runtime.Value
var once_traversableLazy sync.Once
func Get_traversableLazy() gopurs_runtime.Value {
	once_traversableLazy.Do(func() {
		cache_traversableLazy = gopurs_runtime.RecordDict4("traverse", "sequence", "Functor0", "Foldable1", gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
}), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2)))
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}), gopurs_runtime.Apply(Get_force(), l_1))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}))
	})
	return cache_traversableLazy
}

var cache_traversableWithIndexLazy gopurs_runtime.Value
var once_traversableWithIndexLazy sync.Once
func Get_traversableWithIndexLazy() gopurs_runtime.Value {
	once_traversableWithIndexLazy.Do(func() {
		cache_traversableWithIndexLazy = gopurs_runtime.RecordDict4("traverseWithIndex", "FunctorWithIndex0", "FoldableWithIndex1", "Traversable2", gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversableLazy(), "traverse"), dictApplicative_0, gopurs_runtime.Apply(f_1, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorWithIndexLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableWithIndexLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableLazy()
}))
	})
	return cache_traversableWithIndexLazy
}

var cache_foldable1Lazy gopurs_runtime.Value
var once_foldable1Lazy sync.Once
func Get_foldable1Lazy() gopurs_runtime.Value {
	once_foldable1Lazy.Do(func() {
		cache_foldable1Lazy = gopurs_runtime.RecordDict4("foldMap1", "foldr1", "foldl1", "Foldable0", gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2))
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), l_1)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), l_1)
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}))
	})
	return cache_foldable1Lazy
}

var cache_traversable1Lazy gopurs_runtime.Value
var once_traversable1Lazy sync.Once
func Get_traversable1Lazy() gopurs_runtime.Value {
	once_traversable1Lazy.Do(func() {
		cache_traversable1Lazy = gopurs_runtime.RecordDict4("traverse1", "sequence1", "Foldable10", "Traversable1", gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}))
}), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_2)))
}), gopurs_runtime.Func2(func(dictApply_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}))
}), gopurs_runtime.Apply(Get_force(), l_1))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1Lazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableLazy()
}))
	})
	return cache_traversable1Lazy
}

var cache_extendLazy gopurs_runtime.Value
var once_extendLazy sync.Once
func Get_extendLazy() gopurs_runtime.Value {
	once_extendLazy.Do(func() {
		cache_extendLazy = gopurs_runtime.RecordDict2("extend", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_1)
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}))
	})
	return cache_extendLazy
}

var cache_eqLazy gopurs_runtime.Value
var once_eqLazy sync.Once
func Get_eqLazy() gopurs_runtime.Value {
	once_eqLazy.Do(func() {
		cache_eqLazy = gopurs_runtime.Func(func(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Apply(Get_force(), x_1), gopurs_runtime.Apply(Get_force(), y_2))
}))
}()
})
	})
	return cache_eqLazy
}

var cache_ordLazy gopurs_runtime.Value
var once_ordLazy sync.Once
func Get_ordLazy() gopurs_runtime.Value {
	once_ordLazy.Do(func() {
		cache_ordLazy = gopurs_runtime.Func(func(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqLazy1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), gopurs_runtime.Apply(Get_force(), x_2), gopurs_runtime.Apply(Get_force(), y_3))
}))
_ = eqLazy1_2_1
return gopurs_runtime.RecordDict2("compare", "Eq0", gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(Get_force(), x_3), gopurs_runtime.Apply(Get_force(), y_4))
}), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqLazy1_2_1
}))
}()
})
	})
	return cache_ordLazy
}

var cache_eq1Lazy gopurs_runtime.Value
var once_eq1Lazy sync.Once
func Get_eq1Lazy() gopurs_runtime.Value {
	once_eq1Lazy.Do(func() {
		cache_eq1Lazy = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Apply(Get_force(), x_1), gopurs_runtime.Apply(Get_force(), y_2))
}))
	})
	return cache_eq1Lazy
}

var cache_ord1Lazy gopurs_runtime.Value
var once_ord1Lazy sync.Once
func Get_ord1Lazy() gopurs_runtime.Value {
	once_ord1Lazy.Do(func() {
		cache_ord1Lazy = gopurs_runtime.RecordDict2("compare1", "Eq10", gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(Get_ordLazy(), dictOrd_0), "compare")
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Lazy()
}))
	})
	return cache_ord1Lazy
}

var cache_comonadLazy gopurs_runtime.Value
var once_comonadLazy sync.Once
func Get_comonadLazy() gopurs_runtime.Value {
	once_comonadLazy.Do(func() {
		cache_comonadLazy = gopurs_runtime.RecordDict2("extract", "Extend0", Get_force(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendLazy()
}))
	})
	return cache_comonadLazy
}

var cache_commutativeRingLazy gopurs_runtime.Value
var once_commutativeRingLazy sync.Once
func Get_commutativeRingLazy() gopurs_runtime.Value {
	once_commutativeRingLazy.Do(func() {
		cache_commutativeRingLazy = gopurs_runtime.Func(func(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
ringLazy1_1_0 := gopurs_runtime.Apply(Get_ringLazy(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{}))
_ = ringLazy1_1_0
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_0
}))
}()
})
	})
	return cache_commutativeRingLazy
}

var cache_euclideanRingLazy gopurs_runtime.Value
var once_euclideanRingLazy sync.Once
func Get_euclideanRingLazy() gopurs_runtime.Value {
	once_euclideanRingLazy.Do(func() {
		cache_euclideanRingLazy = gopurs_runtime.Func(func(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
ringLazy1_1_0 := gopurs_runtime.Apply(Get_ringLazy(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_0, "CommutativeRing0"), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{}))
_ = ringLazy1_1_0
return gopurs_runtime.RecordDict4("degree", "div", "mod", "CommutativeRing0", gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_0, "degree"), gopurs_runtime.Apply(Get_force(), x_2))
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "div"), gopurs_runtime.Apply(Get_force(), a_2), gopurs_runtime.Apply(Get_force(), b_3))
}))
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "mod"), gopurs_runtime.Apply(Get_force(), a_2), gopurs_runtime.Apply(Get_force(), b_3))
}))
}), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_0
}))
}))
}()
})
	})
	return cache_euclideanRingLazy
}

var cache_boundedLazy gopurs_runtime.Value
var once_boundedLazy sync.Once
func Get_boundedLazy() gopurs_runtime.Value {
	once_boundedLazy.Do(func() {
		cache_boundedLazy = gopurs_runtime.Func(func(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
top_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "top")
_ = top_1_0
bottom_2_1 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
_ = bottom_2_1
ordLazy1_3_2 := gopurs_runtime.Apply(Get_ordLazy(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{}))
_ = ordLazy1_3_2
return gopurs_runtime.RecordDict3("top", "bottom", "Ord0", gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return top_1_0
})), gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bottom_2_1
})), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordLazy1_3_2
}))
}()
})
	})
	return cache_boundedLazy
}

var cache_applyLazy gopurs_runtime.Value
var once_applyLazy sync.Once
func Get_applyLazy() gopurs_runtime.Value {
	once_applyLazy.Do(func() {
		cache_applyLazy = gopurs_runtime.RecordDict2("apply", "Functor0", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_force(), f_0, gopurs_runtime.Apply(Get_force(), x_1))
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}))
	})
	return cache_applyLazy
}

var cache_bindLazy gopurs_runtime.Value
var once_bindLazy sync.Once
func Get_bindLazy() gopurs_runtime.Value {
	once_bindLazy.Do(func() {
		cache_bindLazy = gopurs_runtime.RecordDict2("bind", "Apply0", gopurs_runtime.Func2(func(l_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_force(), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_force(), l_0)))
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
}))
	})
	return cache_bindLazy
}

var cache_heytingAlgebraLazy gopurs_runtime.Value
var once_heytingAlgebraLazy sync.Once
func Get_heytingAlgebraLazy() gopurs_runtime.Value {
	once_heytingAlgebraLazy.Do(func() {
		cache_heytingAlgebraLazy = gopurs_runtime.Func(func(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
ff_1_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
_ = ff_1_0
tt_2_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
_ = tt_2_1
return gopurs_runtime.RecordDict([]string{"ff", "tt", "implies", "conj", "disj", "not"}, []gopurs_runtime.Value{gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ff_1_0
})), gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return tt_2_1
})), gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_2 := gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies"), gopurs_runtime.Apply(Get_force(), a_3))
}))
_ = __local_var_5_2
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_force(), __local_var_5_2, gopurs_runtime.Apply(Get_force(), b_4))
}))
}), gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_3 := gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj"), gopurs_runtime.Apply(Get_force(), a_3))
}))
_ = __local_var_5_3
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_force(), __local_var_5_3, gopurs_runtime.Apply(Get_force(), b_4))
}))
}), gopurs_runtime.Func2(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_5_4 := gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj"), gopurs_runtime.Apply(Get_force(), a_3))
}))
_ = __local_var_5_4
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_force(), __local_var_5_4, gopurs_runtime.Apply(Get_force(), b_4))
}))
}), gopurs_runtime.Func(func(a_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not"), gopurs_runtime.Apply(Get_force(), a_3))
}))
})})
}()
})
	})
	return cache_heytingAlgebraLazy
}

var cache_booleanAlgebraLazy gopurs_runtime.Value
var once_booleanAlgebraLazy sync.Once
func Get_booleanAlgebraLazy() gopurs_runtime.Value {
	once_booleanAlgebraLazy.Do(func() {
		cache_booleanAlgebraLazy = gopurs_runtime.Func(func(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
heytingAlgebraLazy1_1_0 := gopurs_runtime.Apply(Get_heytingAlgebraLazy(), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{}))
_ = heytingAlgebraLazy1_1_0
return gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraLazy1_1_0
}))
}()
})
	})
	return cache_booleanAlgebraLazy
}

var cache_applicativeLazy gopurs_runtime.Value
var once_applicativeLazy sync.Once
func Get_applicativeLazy() gopurs_runtime.Value {
	once_applicativeLazy.Do(func() {
		cache_applicativeLazy = gopurs_runtime.RecordDict2("pure", "Apply0", gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_defer_(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
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



func Get_defer_() gopurs_runtime.Value {
	return _Gopurs_Defer_
}

func Get_force() gopurs_runtime.Value {
	return _Gopurs_Force
}

package Data_Lazy

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unit "gopurs/output/Data.Unit"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
)

var cache_Lazy gopurs_runtime.Value
var once_Lazy sync.Once
func Get_Lazy() gopurs_runtime.Value {
	once_Lazy.Do(func() {
		cache_Lazy = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Lazy(x_0_box)
})
	})
	return cache_Lazy
}

var cache_force gopurs_runtime.Value
var once_force sync.Once
func Get_force() gopurs_runtime.Value {
	once_force.Do(func() {
		cache_force = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_force(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, inner_arg0))
}))
})
	})
	return cache_force
}

var cache_force__func_func_gopurs_runtime_Value__interface____interface___2139985499 gopurs_runtime.Value
var once_force__func_func_gopurs_runtime_Value__interface____interface___2139985499 sync.Once
func Get_force__func_func_gopurs_runtime_Value__interface____interface___2139985499() gopurs_runtime.Value {
	once_force__func_func_gopurs_runtime_Value__interface____interface___2139985499.Do(func() {
		cache_force__func_func_gopurs_runtime_Value__interface____interface___2139985499 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_force__func_func_gopurs_runtime_Value__interface____interface___2139985499(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, inner_arg0))
}))
})
	})
	return cache_force__func_func_gopurs_runtime_Value__interface____interface___2139985499
}

var cache_force__func_func_gopurs_runtime_Value__interface____interface___3614634907 gopurs_runtime.Value
var once_force__func_func_gopurs_runtime_Value__interface____interface___3614634907 sync.Once
func Get_force__func_func_gopurs_runtime_Value__interface____interface___3614634907() gopurs_runtime.Value {
	once_force__func_func_gopurs_runtime_Value__interface____interface___3614634907.Do(func() {
		cache_force__func_func_gopurs_runtime_Value__interface____interface___3614634907 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_force__func_func_gopurs_runtime_Value__interface____interface___3614634907(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, inner_arg0))
}))
})
	})
	return cache_force__func_func_gopurs_runtime_Value__interface____interface___3614634907
}

var cache_force__func_func_gopurs_runtime_Value__interface____interface___3409790107 gopurs_runtime.Value
var once_force__func_func_gopurs_runtime_Value__interface____interface___3409790107 sync.Once
func Get_force__func_func_gopurs_runtime_Value__interface____interface___3409790107() gopurs_runtime.Value {
	once_force__func_func_gopurs_runtime_Value__interface____interface___3409790107.Do(func() {
		cache_force__func_func_gopurs_runtime_Value__interface____interface___3409790107 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_force__func_func_gopurs_runtime_Value__interface____interface___3409790107(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, inner_arg0))
}))
})
	})
	return cache_force__func_func_gopurs_runtime_Value__interface____interface___3409790107
}

var cache_force__func_func_gopurs_runtime_Value__interface____interface___2059076955 gopurs_runtime.Value
var once_force__func_func_gopurs_runtime_Value__interface____interface___2059076955 sync.Once
func Get_force__func_func_gopurs_runtime_Value__interface____interface___2059076955() gopurs_runtime.Value {
	once_force__func_func_gopurs_runtime_Value__interface____interface___2059076955.Do(func() {
		cache_force__func_func_gopurs_runtime_Value__interface____interface___2059076955 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_force__func_func_gopurs_runtime_Value__interface____interface___2059076955(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(v_0_box, inner_arg0))
}))
})
	})
	return cache_force__func_func_gopurs_runtime_Value__interface____interface___2059076955
}

var cache_showLazy gopurs_runtime.Value
var once_showLazy sync.Once
func Get_showLazy() gopurs_runtime.Value {
	once_showLazy.Do(func() {
		cache_showLazy = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_showLazy(dictShow_0_box))
})
	})
	return cache_showLazy
}

var cache_foldableLazy gopurs_runtime.Value
var once_foldableLazy sync.Once
func Get_foldableLazy() gopurs_runtime.Value {
	once_foldableLazy.Do(func() {
		cache_foldableLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(l_2, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.Apply(l_2, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(l_2, pkg_Data_Unit.Get_unit()), z_1)
}))))
	})
	return cache_foldableLazy
}

var cache_foldableWithIndexLazy gopurs_runtime.Value
var once_foldableWithIndexLazy sync.Once
func Get_foldableWithIndexLazy() gopurs_runtime.Value {
	once_foldableWithIndexLazy.Do(func() {
		cache_foldableWithIndexLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))))
	})
	return cache_foldableWithIndexLazy
}

var cache_foldable1Lazy gopurs_runtime.Value
var once_foldable1Lazy sync.Once
func Get_foldable1Lazy() gopurs_runtime.Value {
	once_foldable1Lazy.Do(func() {
		cache_foldable1Lazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(l_2, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(l_1, pkg_Data_Unit.Get_unit())
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(l_1, pkg_Data_Unit.Get_unit())
}))))
	})
	return cache_foldable1Lazy
}

var cache_eqLazy gopurs_runtime.Value
var once_eqLazy sync.Once
func Get_eqLazy() gopurs_runtime.Value {
	once_eqLazy.Do(func() {
		cache_eqLazy = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_eqLazy(dictEq_0_box))
})
	})
	return cache_eqLazy
}

var cache_ordLazy gopurs_runtime.Value
var once_ordLazy sync.Once
func Get_ordLazy() gopurs_runtime.Value {
	once_ordLazy.Do(func() {
		cache_ordLazy = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_ordLazy(dictOrd_0_box))
})
	})
	return cache_ordLazy
}

var cache_eq1Lazy gopurs_runtime.Value
var once_eq1Lazy sync.Once
func Get_eq1Lazy() gopurs_runtime.Value {
	once_eq1Lazy.Do(func() {
		cache_eq1Lazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func3(func(dictEq_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Apply(x_1, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(y_2, pkg_Data_Unit.Get_unit()))
}))))
	})
	return cache_eq1Lazy
}

var cache_ord1Lazy gopurs_runtime.Value
var once_ord1Lazy sync.Once
func Get_ord1Lazy() gopurs_runtime.Value {
	once_ord1Lazy.Do(func() {
		cache_ord1Lazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_eq1Lazy()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Any(Call_ordLazy(dictOrd_0)), "compare")
}))))
	})
	return cache_ord1Lazy
}

var cache_defer_ gopurs_runtime.Value
var once_defer_ sync.Once
func Get_defer_() gopurs_runtime.Value {
	once_defer_.Do(func() {
		cache_defer_ = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer_(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, inner_arg0))
})
})
	})
	return cache_defer_
}

var cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538 gopurs_runtime.Value
var once_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538 sync.Once
func Get_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538() gopurs_runtime.Value {
	once_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538.Do(func() {
		cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, inner_arg0))
})
})
	})
	return cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538
}

var cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___2907862714 gopurs_runtime.Value
var once_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___2907862714 sync.Once
func Get_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___2907862714() gopurs_runtime.Value {
	once_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___2907862714.Do(func() {
		cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___2907862714 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___2907862714(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, inner_arg0))
})
})
	})
	return cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___2907862714
}

var cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___1377601082 gopurs_runtime.Value
var once_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___1377601082 sync.Once
func Get_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___1377601082() gopurs_runtime.Value {
	once_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___1377601082.Do(func() {
		cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___1377601082 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___1377601082(func(inner_arg0 gopurs_runtime.Value) interface{} {
return gopurs_runtime.UnboxAny(gopurs_runtime.Apply(f_0_box, inner_arg0))
})
})
	})
	return cache_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___1377601082
}

var cache_functorLazy gopurs_runtime.Value
var once_functorLazy sync.Once
func Get_functorLazy() gopurs_runtime.Value {
	once_functorLazy.Do(func() {
		cache_functorLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("map", gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(l_1, pkg_Data_Unit.Get_unit()))
}))))
	})
	return cache_functorLazy
}

var cache_extendLazy gopurs_runtime.Value
var once_extendLazy sync.Once
func Get_extendLazy() gopurs_runtime.Value {
	once_extendLazy.Do(func() {
		cache_extendLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "extend", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_1)
}))))
	})
	return cache_extendLazy
}

var cache_functorWithIndexLazy gopurs_runtime.Value
var once_functorWithIndexLazy sync.Once
func Get_functorWithIndexLazy() gopurs_runtime.Value {
	once_functorWithIndexLazy.Do(func() {
		cache_functorWithIndexLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), gopurs_runtime.Apply(f_0, pkg_Data_Unit.Get_unit()))
}))))
	})
	return cache_functorWithIndexLazy
}

var cache_invariantLazy gopurs_runtime.Value
var once_invariantLazy sync.Once
func Get_invariantLazy() gopurs_runtime.Value {
	once_invariantLazy.Do(func() {
		cache_invariantLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("imap", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), f_0)
}))))
	})
	return cache_invariantLazy
}

var cache_lazyLazy gopurs_runtime.Value
var once_lazyLazy sync.Once
func Get_lazyLazy() gopurs_runtime.Value {
	once_lazyLazy.Do(func() {
		cache_lazyLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("defer", gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), pkg_Data_Unit.Get_unit())
}))))
	})
	return cache_lazyLazy
}

var cache_semigroupLazy gopurs_runtime.Value
var once_semigroupLazy sync.Once
func Get_semigroupLazy() gopurs_runtime.Value {
	once_semigroupLazy.Do(func() {
		cache_semigroupLazy = gopurs_runtime.Func(func(dictSemigroup_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_semigroupLazy(dictSemigroup_0_box))
})
	})
	return cache_semigroupLazy
}

var cache_monoidLazy gopurs_runtime.Value
var once_monoidLazy sync.Once
func Get_monoidLazy() gopurs_runtime.Value {
	once_monoidLazy.Do(func() {
		cache_monoidLazy = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_monoidLazy(dictMonoid_0_box))
})
	})
	return cache_monoidLazy
}

var cache_semiringLazy gopurs_runtime.Value
var once_semiringLazy sync.Once
func Get_semiringLazy() gopurs_runtime.Value {
	once_semiringLazy.Do(func() {
		cache_semiringLazy = gopurs_runtime.Func(func(dictSemiring_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_semiringLazy(dictSemiring_0_box))
})
	})
	return cache_semiringLazy
}

var cache_ringLazy gopurs_runtime.Value
var once_ringLazy sync.Once
func Get_ringLazy() gopurs_runtime.Value {
	once_ringLazy.Do(func() {
		cache_ringLazy = gopurs_runtime.Func(func(dictRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_ringLazy(dictRing_0_box))
})
	})
	return cache_ringLazy
}

var cache_traversableLazy gopurs_runtime.Value
var once_traversableLazy sync.Once
func Get_traversableLazy() gopurs_runtime.Value {
	once_traversableLazy.Do(func() {
		cache_traversableLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldableLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func2(func(dictApplicative_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), gopurs_runtime.Apply(l_1, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func3(func(dictApplicative_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(l_2, pkg_Data_Unit.Get_unit())))
}))))
	})
	return cache_traversableLazy
}

var cache_traversable1Lazy gopurs_runtime.Value
var once_traversable1Lazy sync.Once
func Get_traversable1Lazy() gopurs_runtime.Value {
	once_traversable1Lazy.Do(func() {
		cache_traversable1Lazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1Lazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_traversableLazy()
}), gopurs_runtime.Func2(func(dictApply_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return x_2
}), gopurs_runtime.Apply(l_1, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func3(func(dictApply_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return x_3
}), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(l_2, pkg_Data_Unit.Get_unit())))
}))))
	})
	return cache_traversable1Lazy
}

var cache_traversableWithIndexLazy gopurs_runtime.Value
var once_traversableWithIndexLazy sync.Once
func Get_traversableWithIndexLazy() gopurs_runtime.Value {
	once_traversableWithIndexLazy.Do(func() {
		cache_traversableWithIndexLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))))
	})
	return cache_traversableWithIndexLazy
}

var cache_comonadLazy gopurs_runtime.Value
var once_comonadLazy sync.Once
func Get_comonadLazy() gopurs_runtime.Value {
	once_comonadLazy.Do(func() {
		cache_comonadLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Extend0", "extract", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_extendLazy()
}), Get_force())))
	})
	return cache_comonadLazy
}

var cache_commutativeRingLazy gopurs_runtime.Value
var once_commutativeRingLazy sync.Once
func Get_commutativeRingLazy() gopurs_runtime.Value {
	once_commutativeRingLazy.Do(func() {
		cache_commutativeRingLazy = gopurs_runtime.Func(func(dictCommutativeRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_commutativeRingLazy(dictCommutativeRing_0_box))
})
	})
	return cache_commutativeRingLazy
}

var cache_euclideanRingLazy gopurs_runtime.Value
var once_euclideanRingLazy sync.Once
func Get_euclideanRingLazy() gopurs_runtime.Value {
	once_euclideanRingLazy.Do(func() {
		cache_euclideanRingLazy = gopurs_runtime.Func(func(dictEuclideanRing_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_euclideanRingLazy(dictEuclideanRing_0_box))
})
	})
	return cache_euclideanRingLazy
}

var cache_boundedLazy gopurs_runtime.Value
var once_boundedLazy sync.Once
func Get_boundedLazy() gopurs_runtime.Value {
	once_boundedLazy.Do(func() {
		cache_boundedLazy = gopurs_runtime.Func(func(dictBounded_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_boundedLazy(dictBounded_0_box))
})
	})
	return cache_boundedLazy
}

var cache_applyLazy gopurs_runtime.Value
var once_applyLazy sync.Once
func Get_applyLazy() gopurs_runtime.Value {
	once_applyLazy.Do(func() {
		cache_applyLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_functorLazy()
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, pkg_Data_Unit.Get_unit(), gopurs_runtime.Apply(x_1, pkg_Data_Unit.Get_unit()))
}))))
	})
	return cache_applyLazy
}

var cache_bindLazy gopurs_runtime.Value
var once_bindLazy sync.Once
func Get_bindLazy() gopurs_runtime.Value {
	once_bindLazy.Do(func() {
		cache_bindLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
}), gopurs_runtime.Func3(func(l_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_1, gopurs_runtime.Apply(l_0, pkg_Data_Unit.Get_unit()), pkg_Data_Unit.Get_unit())
}))))
	})
	return cache_bindLazy
}

var cache_heytingAlgebraLazy gopurs_runtime.Value
var once_heytingAlgebraLazy sync.Once
func Get_heytingAlgebraLazy() gopurs_runtime.Value {
	once_heytingAlgebraLazy.Do(func() {
		cache_heytingAlgebraLazy = gopurs_runtime.Func(func(dictHeytingAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_heytingAlgebraLazy(dictHeytingAlgebra_0_box))
})
	})
	return cache_heytingAlgebraLazy
}

var cache_booleanAlgebraLazy gopurs_runtime.Value
var once_booleanAlgebraLazy sync.Once
func Get_booleanAlgebraLazy() gopurs_runtime.Value {
	once_booleanAlgebraLazy.Do(func() {
		cache_booleanAlgebraLazy = gopurs_runtime.Func(func(dictBooleanAlgebra_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(Call_booleanAlgebraLazy(dictBooleanAlgebra_0_box))
})
	})
	return cache_booleanAlgebraLazy
}

var cache_applicativeLazy gopurs_runtime.Value
var once_applicativeLazy sync.Once
func Get_applicativeLazy() gopurs_runtime.Value {
	once_applicativeLazy.Do(func() {
		cache_applicativeLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applyLazy()
}), gopurs_runtime.Func2(func(a_0 gopurs_runtime.Value, v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))))
	})
	return cache_applicativeLazy
}

var cache_monadLazy gopurs_runtime.Value
var once_monadLazy sync.Once
func Get_monadLazy() gopurs_runtime.Value {
	once_monadLazy.Do(func() {
		cache_monadLazy = gopurs_runtime.Any(gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_applicativeLazy()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_bindLazy()
}))))
	})
	return cache_monadLazy
}

func Call_Lazy(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_force(v_0_loop func(gopurs_runtime.Value) interface{}) interface{} {
var v_0 func(gopurs_runtime.Value) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(pkg_Data_Unit.Get_unit())))
}

func Call_force__func_func_gopurs_runtime_Value__interface____interface___2139985499(v_0_loop func(gopurs_runtime.Value) interface{}) interface{} {
var v_0 func(gopurs_runtime.Value) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(pkg_Data_Unit.Get_unit())))
}

func Call_force__func_func_gopurs_runtime_Value__interface____interface___3614634907(v_0_loop func(gopurs_runtime.Value) interface{}) interface{} {
var v_0 func(gopurs_runtime.Value) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(pkg_Data_Unit.Get_unit())))
}

func Call_force__func_func_gopurs_runtime_Value__interface____interface___3409790107(v_0_loop func(gopurs_runtime.Value) interface{}) interface{} {
var v_0 func(gopurs_runtime.Value) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(pkg_Data_Unit.Get_unit())))
}

func Call_force__func_func_gopurs_runtime_Value__interface____interface___2059076955(v_0_loop func(gopurs_runtime.Value) interface{}) interface{} {
var v_0 func(gopurs_runtime.Value) interface{} = v_0_loop
_ = v_0
return gopurs_runtime.UnboxAny(gopurs_runtime.Any(v_0(pkg_Data_Unit.Get_unit())))
}

func Call_showLazy(dictShow_0_loop gopurs_runtime.Value) interface{} {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(defer \\_ -> "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Apply(x_1, pkg_Data_Unit.Get_unit())), gopurs_runtime.Str(")")))
})))
}

func Call_eqLazy(dictEq_0_loop gopurs_runtime.Value) interface{} {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Apply(x_1, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(y_2, pkg_Data_Unit.Get_unit()))
})))
}

func Call_ordLazy(dictOrd_0_loop gopurs_runtime.Value) interface{} {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
__local_var_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})
_ = __local_var_1_0
eqLazy1_2_1 := gopurs_runtime.RecordDict1("eq", gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_1_0, "eq"), gopurs_runtime.Apply(x_2, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(y_3, pkg_Data_Unit.Get_unit()))
}))
_ = eqLazy1_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Eq0", "compare", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return eqLazy1_2_1
}), gopurs_runtime.Func2(func(x_3 gopurs_runtime.Value, y_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(x_3, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(y_4, pkg_Data_Unit.Get_unit()))
})))
}

func Call_defer_(f_0_loop func(gopurs_runtime.Value) interface{}) gopurs_runtime.Value {
var f_0 func(gopurs_runtime.Value) interface{} = f_0_loop
_ = f_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(arg0))
})
}

func Call_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___503512538(f_0_loop func(gopurs_runtime.Value) interface{}) gopurs_runtime.Value {
var f_0 func(gopurs_runtime.Value) interface{} = f_0_loop
_ = f_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(arg0))
})
}

func Call_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___2907862714(f_0_loop func(gopurs_runtime.Value) interface{}) gopurs_runtime.Value {
var f_0 func(gopurs_runtime.Value) interface{} = f_0_loop
_ = f_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(arg0))
})
}

func Call_defer__func_func_gopurs_runtime_Value__interface____gopurs_runtime_Value__interface___1377601082(f_0_loop func(gopurs_runtime.Value) interface{}) gopurs_runtime.Value {
var f_0 func(gopurs_runtime.Value) interface{} = f_0_loop
_ = f_0
return gopurs_runtime.Func(func(arg0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Any(f_0(arg0))
})
}

func Call_semigroupLazy(dictSemigroup_0_loop gopurs_runtime.Value) interface{} {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value, v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(a_1, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(b_2, pkg_Data_Unit.Get_unit()))
})))
}

func Call_monoidLazy(dictMonoid_0_loop gopurs_runtime.Value) interface{} {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
mempty_1_0 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_1_0
__local_var_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})
_ = __local_var_2_1
semigroupLazy1_3_2 := gopurs_runtime.RecordDict1("append", gopurs_runtime.Func3(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(__local_var_2_1, "append"), gopurs_runtime.Apply(a_3, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(b_4, pkg_Data_Unit.Get_unit()))
}))
_ = semigroupLazy1_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Semigroup0", "mempty", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return semigroupLazy1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return mempty_1_0
})))
}

func Call_semiringLazy(dictSemiring_0_loop gopurs_runtime.Value) interface{} {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
zero_1_0 := gopurs_runtime.RecordGet(dictSemiring_0, "zero")
_ = zero_1_0
one_2_1 := gopurs_runtime.RecordGet(dictSemiring_0, "one")
_ = one_2_1
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict4("add", "mul", "one", "zero", gopurs_runtime.Func3(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), gopurs_runtime.Apply(a_3, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(b_4, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func3(func(a_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), gopurs_runtime.Apply(a_3, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(b_4, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return one_2_1
}), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return zero_1_0
})))
}

func Call_ringLazy(dictRing_0_loop gopurs_runtime.Value) interface{} {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
semiringLazy1_1_0 := gopurs_runtime.Any(Call_semiringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})))
_ = semiringLazy1_1_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict2("Semiring0", "sub", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return semiringLazy1_1_0
}), gopurs_runtime.Func3(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(a_2, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(b_3, pkg_Data_Unit.Get_unit()))
})))
}

func Call_commutativeRingLazy(dictCommutativeRing_0_loop gopurs_runtime.Value) interface{} {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
ringLazy1_1_0 := gopurs_runtime.Any(Call_ringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{})))
_ = ringLazy1_1_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_0
})))
}

func Call_euclideanRingLazy(dictEuclideanRing_0_loop gopurs_runtime.Value) interface{} {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
ringLazy1_1_0 := gopurs_runtime.Any(Call_ringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_0, "CommutativeRing0"), gopurs_runtime.Value{}), "Ring0"), gopurs_runtime.Value{})))
_ = ringLazy1_1_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict4("CommutativeRing0", "degree", "div", "mod", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("Ring0", gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return ringLazy1_1_0
}))
}), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_0, "degree"), gopurs_runtime.Apply(x_2, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func3(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "div"), gopurs_runtime.Apply(a_2, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(b_3, pkg_Data_Unit.Get_unit()))
}), gopurs_runtime.Func3(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "mod"), gopurs_runtime.Apply(a_2, pkg_Data_Unit.Get_unit()), gopurs_runtime.Apply(b_3, pkg_Data_Unit.Get_unit()))
})))
}

func Call_boundedLazy(dictBounded_0_loop gopurs_runtime.Value) interface{} {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
top_1_0 := gopurs_runtime.RecordGet(dictBounded_0, "top")
_ = top_1_0
bottom_2_1 := gopurs_runtime.RecordGet(dictBounded_0, "bottom")
_ = bottom_2_1
ordLazy1_3_2 := gopurs_runtime.Any(Call_ordLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})))
_ = ordLazy1_3_2
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict3("Ord0", "bottom", "top", gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return ordLazy1_3_2
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return bottom_2_1
}), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return top_1_0
})))
}

func Call_heytingAlgebraLazy(dictHeytingAlgebra_0_loop gopurs_runtime.Value) interface{} {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
ff_1_0 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
_ = ff_1_0
tt_2_1 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
_ = tt_2_1
implies_3_2 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "implies")
_ = implies_3_2
conj_4_3 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "conj")
_ = conj_4_3
disj_5_4 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "disj")
_ = disj_5_4
not_6_5 := gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "not")
_ = not_6_5
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict([]string{"conj", "disj", "ff", "implies", "not", "tt"}, []gopurs_runtime.Value{gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), conj_4_3, a_7), b_8)
}), gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), disj_5_4, a_7), b_8)
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return ff_1_0
}), gopurs_runtime.Func2(func(a_7 gopurs_runtime.Value, b_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_applyLazy(), "apply"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), implies_3_2, a_7), b_8)
}), gopurs_runtime.Func(func(a_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_functorLazy(), "map"), not_6_5, a_7)
}), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return tt_2_1
})}))
}

func Call_booleanAlgebraLazy(dictBooleanAlgebra_0_loop gopurs_runtime.Value) interface{} {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
heytingAlgebraLazy1_1_0 := gopurs_runtime.Any(Call_heytingAlgebraLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{})))
_ = heytingAlgebraLazy1_1_0
return gopurs_runtime.UnboxAny(gopurs_runtime.RecordDict1("HeytingAlgebra0", gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return heytingAlgebraLazy1_1_0
})))
}

package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
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
		cache_Data_Lazy_lazyLazy = gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer((&Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}))
})}))}
	})
	return cache_Data_Lazy_lazyLazy
}

var cache_Data_Lazy_functorLazy gopurs_runtime.Value
var once_Data_Lazy_functorLazy sync.Once
func Get_Data_Lazy_functorLazy() gopurs_runtime.Value {
	once_Data_Lazy_functorLazy.Do(func() {
		cache_Data_Lazy_functorLazy = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Functor[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1))
}))
})}))}
	})
	return cache_Data_Lazy_functorLazy
}

var cache_Data_Lazy_go__map gopurs_runtime.Value
var once_Data_Lazy_go__map sync.Once
func Get_Data_Lazy_go__map() gopurs_runtime.Value {
	once_Data_Lazy_go__map.Do(func() {
		cache_Data_Lazy_go__map = Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()))
	})
	return cache_Data_Lazy_go__map
}

var cache_Data_Lazy_functorWithIndexLazy gopurs_runtime.Value
var once_Data_Lazy_functorWithIndexLazy sync.Once
func Get_Data_Lazy_functorWithIndexLazy() gopurs_runtime.Value {
	once_Data_Lazy_functorWithIndexLazy.Do(func() {
		cache_Data_Lazy_functorWithIndexLazy = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer((&Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()))}
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Functor_go__map(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})}))}
	})
	return cache_Data_Lazy_functorWithIndexLazy
}

var cache_Data_Lazy_invariantLazy gopurs_runtime.Value
var once_Data_Lazy_invariantLazy sync.Once
func Get_Data_Lazy_invariantLazy() gopurs_runtime.Value {
	once_Data_Lazy_invariantLazy.Do(func() {
		cache_Data_Lazy_invariantLazy = gopurs_runtime.Value{Type: 9, IntVal: 2396985522, UnsafePtr: unsafe.Pointer((&Constructor_Data_Functor_Invariant_Invariant[gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Functor_Invariant_imapF(), gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()))})}))}
	})
	return cache_Data_Lazy_invariantLazy
}

var cache_Data_Lazy_foldableLazy gopurs_runtime.Value
var once_Data_Lazy_foldableLazy sync.Once
func Get_Data_Lazy_foldableLazy() gopurs_runtime.Value {
	once_Data_Lazy_foldableLazy.Do(func() {
		cache_Data_Lazy_foldableLazy = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer((&Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]{1, gopurs_runtime.Func3(func(dictMonoid_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, z_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
}), gopurs_runtime.Func3(func(f_0 gopurs_runtime.Value, z_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2), z_1)
})}))}
	})
	return cache_Data_Lazy_foldableLazy
}

var cache_Data_Lazy_foldr gopurs_runtime.Value
var once_Data_Lazy_foldr sync.Once
func Get_Data_Lazy_foldr() gopurs_runtime.Value {
	once_Data_Lazy_foldr.Do(func() {
		cache_Data_Lazy_foldr = Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Lazy_foldableLazy()))
	})
	return cache_Data_Lazy_foldr
}

var cache_Data_Lazy_foldl gopurs_runtime.Value
var once_Data_Lazy_foldl sync.Once
func Get_Data_Lazy_foldl() gopurs_runtime.Value {
	once_Data_Lazy_foldl.Do(func() {
		cache_Data_Lazy_foldl = Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Lazy_foldableLazy()))
	})
	return cache_Data_Lazy_foldl
}

var cache_Data_Lazy_foldMap gopurs_runtime.Value
var once_Data_Lazy_foldMap sync.Once
func Get_Data_Lazy_foldMap() gopurs_runtime.Value {
	once_Data_Lazy_foldMap.Do(func() {
		cache_Data_Lazy_foldMap = Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Lazy_foldableLazy()))
	})
	return cache_Data_Lazy_foldMap
}

var cache_Data_Lazy_foldableWithIndexLazy gopurs_runtime.Value
var once_Data_Lazy_foldableWithIndexLazy sync.Once
func Get_Data_Lazy_foldableWithIndexLazy() gopurs_runtime.Value {
	once_Data_Lazy_foldableWithIndexLazy.Do(func() {
		cache_Data_Lazy_foldableWithIndexLazy = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer((&Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Lazy_foldableLazy()))}
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): foldMap1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope122)] (TypeVar m$scope123)), (ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope122)])] (TypeVar m$scope123))
foldMap1_1_0 := gopurs_runtime.Apply(Call_Data_Foldable_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Lazy_foldableLazy())), dictMonoid_0)
_ = foldMap1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(foldMap1_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldl(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Lazy_foldableLazy())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Call_Data_Foldable_foldr(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Lazy_foldableLazy())), gopurs_runtime.Apply(f_0, Get_Data_Unit_unit()))
})}))}
	})
	return cache_Data_Lazy_foldableWithIndexLazy
}

var cache_Data_Lazy_traversableLazy gopurs_runtime.Value
var once_Data_Lazy_traversableLazy sync.Once
func Get_Data_Lazy_traversableLazy() gopurs_runtime.Value {
	once_Data_Lazy_traversableLazy.Do(func() {
		cache_Data_Lazy_traversableLazy = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer((&Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Lazy_foldableLazy()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope21)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Lazy_go__defer(), Get_Data_Function_go__const()), gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar m$scope16)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Lazy_go__defer(), Get_Data_Function_go__const()), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_3)))
})
})}))}
	})
	return cache_Data_Lazy_traversableLazy
}

var cache_Data_Lazy_traverse gopurs_runtime.Value
var once_Data_Lazy_traverse sync.Once
func Get_Data_Lazy_traverse() gopurs_runtime.Value {
	once_Data_Lazy_traverse.Do(func() {
		cache_Data_Lazy_traverse = Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Lazy_traversableLazy()))
	})
	return cache_Data_Lazy_traverse
}

var cache_Data_Lazy_traversableWithIndexLazy gopurs_runtime.Value
var once_Data_Lazy_traversableWithIndexLazy sync.Once
func Get_Data_Lazy_traversableWithIndexLazy() gopurs_runtime.Value {
	once_Data_Lazy_traversableWithIndexLazy.Do(func() {
		cache_Data_Lazy_traversableWithIndexLazy = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer((&Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Lazy_foldableWithIndexLazy()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](Get_Data_Lazy_functorWithIndexLazy()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Lazy_traversableLazy()))}
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): traverse1_1_0 shape=App(Var) bindingType=(Func [(Func [(TypeVar a$scope5)] (TypeApp (TypeVar m$scope7) [(TypeVar b$scope6)])), (ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope5)])] (TypeApp (TypeVar m$scope7) [(ADT ["Data","Lazy","Lazy"] [(TypeVar b$scope6)])]))
traverse1_1_0 := gopurs_runtime.Apply(Call_Data_Traversable_traverse(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Lazy_traversableLazy())), dictApplicative_0)
_ = traverse1_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(traverse1_1_0, gopurs_runtime.Apply(f_2, Get_Data_Unit_unit()))
})
})}))}
	})
	return cache_Data_Lazy_traversableWithIndexLazy
}

var cache_Data_Lazy_foldable1Lazy gopurs_runtime.Value
var once_Data_Lazy_foldable1Lazy sync.Once
func Get_Data_Lazy_foldable1Lazy() gopurs_runtime.Value {
	once_Data_Lazy_foldable1Lazy.Do(func() {
		cache_Data_Lazy_foldable1Lazy = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Lazy_foldableLazy()))}
}), gopurs_runtime.Func3(func(dictSemigroup_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value, l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)
}), gopurs_runtime.Func2(func(v_0 gopurs_runtime.Value, l_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), l_1)
})}))}
	})
	return cache_Data_Lazy_foldable1Lazy
}

var cache_Data_Lazy_traversable1Lazy gopurs_runtime.Value
var once_Data_Lazy_traversable1Lazy sync.Once
func Get_Data_Lazy_traversable1Lazy() gopurs_runtime.Value {
	once_Data_Lazy_traversable1Lazy.Do(func() {
		cache_Data_Lazy_traversable1Lazy = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]](Get_Data_Lazy_foldable1Lazy()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Lazy_traversableLazy()))}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope38)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(l_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_0.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Lazy_go__defer(), Get_Data_Function_go__const()), gopurs_runtime.Apply(Get_Data_Lazy_force(), l_2))
})
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): Functor0_1_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f$scope33)])
Functor0_1_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_1
return gopurs_runtime.Func2(func(f_2 gopurs_runtime.Value, l_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Functor0_1_1.V0, gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Get_Data_Lazy_go__defer(), Get_Data_Function_go__const()), gopurs_runtime.Apply(f_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_3)))
})
})}))}
	})
	return cache_Data_Lazy_traversable1Lazy
}

var cache_Data_Lazy_extendLazy gopurs_runtime.Value
var once_Data_Lazy_extendLazy sync.Once
func Get_Data_Lazy_extendLazy() gopurs_runtime.Value {
	once_Data_Lazy_extendLazy.Do(func() {
		cache_Data_Lazy_extendLazy = gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer((&Constructor_Control_Extend_Extend[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(f_0, x_1)
}))
})}))}
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
		cache_Data_Lazy_eq1Lazy = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Eq_eq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Lazy_eqLazy(dictEq_0)))
})}))}
	})
	return cache_Data_Lazy_eq1Lazy
}

var cache_Data_Lazy_ord1Lazy gopurs_runtime.Value
var once_Data_Lazy_ord1Lazy sync.Once
func Get_Data_Lazy_ord1Lazy() gopurs_runtime.Value {
	once_Data_Lazy_ord1Lazy.Do(func() {
		cache_Data_Lazy_ord1Lazy = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord1[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Lazy_eq1Lazy()))}
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Ord_compare(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Lazy_ordLazy(dictOrd_0)))
})}))}
	})
	return cache_Data_Lazy_ord1Lazy
}

var cache_Data_Lazy_comonadLazy gopurs_runtime.Value
var once_Data_Lazy_comonadLazy sync.Once
func Get_Data_Lazy_comonadLazy() gopurs_runtime.Value {
	once_Data_Lazy_comonadLazy.Do(func() {
		cache_Data_Lazy_comonadLazy = gopurs_runtime.Value{Type: 9, IntVal: 2886863693, UnsafePtr: unsafe.Pointer((&Constructor_Control_Comonad_Comonad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3028639021, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Extend_Extend[gopurs_runtime.Value]](Get_Data_Lazy_extendLazy()))}
}), Get_Data_Lazy_force()}))}
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
		cache_Data_Lazy_applyLazy = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer((&Constructor_Control_Apply_Apply[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Lazy_functorLazy()))}
}), gopurs_runtime.Func2(func(f_0 gopurs_runtime.Value, x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), f_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1))
}))
})}))}
	})
	return cache_Data_Lazy_applyLazy
}

var cache_Data_Lazy_bindLazy gopurs_runtime.Value
var once_Data_Lazy_bindLazy sync.Once
func Get_Data_Lazy_bindLazy() gopurs_runtime.Value {
	once_Data_Lazy_bindLazy.Do(func() {
		cache_Data_Lazy_bindLazy = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer((&Constructor_Control_Bind_Bind[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Lazy_applyLazy()))}
}), gopurs_runtime.Func2(func(l_0 gopurs_runtime.Value, f_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_force(), gopurs_runtime.Apply(f_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), l_0)))
}))
})}))}
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
		cache_Data_Lazy_applicativeLazy = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer((&Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Data_Lazy_applyLazy()))}
}), gopurs_runtime.Func(func(a_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return a_0
}))
})}))}
	})
	return cache_Data_Lazy_applicativeLazy
}

var cache_Data_Lazy_monadLazy gopurs_runtime.Value
var once_Data_Lazy_monadLazy sync.Once
func Get_Data_Lazy_monadLazy() gopurs_runtime.Value {
	once_Data_Lazy_monadLazy.Do(func() {
		cache_Data_Lazy_monadLazy = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer((&Constructor_Control_Monad_Monad[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Data_Lazy_applicativeLazy()))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Data_Lazy_bindLazy()))}
})}))}
	})
	return cache_Data_Lazy_monadLazy
}

func Call_Data_Lazy_showLazy(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer((&Constructor_Data_Show_Show[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(defer \\_ -> ") + (gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictShow_0, "show"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1)).StrVal())) + (")"))
})}))}
}

func Call_Data_Lazy_semiringLazy(dictSemiring_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemiring_0 gopurs_runtime.Value = dictSemiring_0_loop
_ = dictSemiring_0
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "add"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_2))
}))
}), gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemiring_0, "mul"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_2))
}))
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictSemiring_0, "one")
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictSemiring_0, "zero")
}))}))}
}

func Call_Data_Lazy_semigroupLazy(dictSemigroup_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictSemigroup_0 gopurs_runtime.Value = dictSemigroup_0_loop
_ = dictSemigroup_0
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer((&Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_1 gopurs_runtime.Value, b_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictSemigroup_0, "append"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_1), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_2))
}))
})}))}
}

func Call_Data_Lazy_ringLazy(dictRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictRing_0 gopurs_runtime.Value = dictRing_0_loop
_ = dictRing_0
// TAST (Let): semiringLazy1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semiring","Semiring"] [(ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope64)])])
semiringLazy1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semiring_Semiring[gopurs_runtime.Value]](Call_Data_Lazy_semiringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictRing_0, "Semiring0"), gopurs_runtime.Value{})))
_ = semiringLazy1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ring_Ring[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 134961754, UnsafePtr: unsafe.Pointer(semiringLazy1_1_0)}
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictRing_0, "sub"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_3))
}))
})}))}
}

func Call_Data_Lazy_monoidLazy(dictMonoid_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictMonoid_0 gopurs_runtime.Value = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): semigroupLazy1_1_0 shape=App(Var) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope78)])])
semigroupLazy1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Call_Data_Lazy_semigroupLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{})))
_ = semigroupLazy1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer((&Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(semigroupLazy1_1_0)}
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
}))}))}
}

func Call_Data_Lazy_eqLazy(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer((&Constructor_Data_Eq_Eq[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEq_0, "eq"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_1), gopurs_runtime.Apply(Get_Data_Lazy_force(), y_2)).IntVal) != (0))
})}))}
}

func Call_Data_Lazy_ordLazy(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqLazy1_1_0 shape=App(Var) bindingType=(ADT ["Data","Eq","Eq"] [(ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope70)])])
eqLazy1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](Call_Data_Lazy_eqLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{})))
_ = eqLazy1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer((&Constructor_Data_Ord_Ord[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqLazy1_1_0)}
}), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), gopurs_runtime.Apply(Get_Data_Lazy_force(), x_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), y_3)).IntVal)), UnsafePtr: nil}
})}))}
}

func Call_Data_Lazy_commutativeRingLazy(dictCommutativeRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictCommutativeRing_0 gopurs_runtime.Value = dictCommutativeRing_0_loop
_ = dictCommutativeRing_0
// TAST (Let): ringLazy1_1_0 shape=App(Var) bindingType=(ADT ["Data","Ring","Ring"] [(ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope182)])])
ringLazy1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ring_Ring[gopurs_runtime.Value]](Call_Data_Lazy_ringLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictCommutativeRing_0, "Ring0"), gopurs_runtime.Value{})))
_ = ringLazy1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer((&Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3955491866, UnsafePtr: unsafe.Pointer(ringLazy1_1_0)}
})}))}
}

func Call_Data_Lazy_euclideanRingLazy(dictEuclideanRing_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEuclideanRing_0 gopurs_runtime.Value = dictEuclideanRing_0_loop
_ = dictEuclideanRing_0
// TAST (Let): commutativeRingLazy1_1_0 shape=App(Var) bindingType=(ADT ["Data","CommutativeRing","CommutativeRing"] [(ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope159)])])
commutativeRingLazy1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_CommutativeRing_CommutativeRing[gopurs_runtime.Value]](Call_Data_Lazy_commutativeRingLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictEuclideanRing_0, "CommutativeRing0"), gopurs_runtime.Value{})))
_ = commutativeRingLazy1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3214993658, UnsafePtr: unsafe.Pointer((&Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1775085946, UnsafePtr: unsafe.Pointer(commutativeRingLazy1_1_0)}
}), gopurs_runtime.Apply2(Call_Control_Semigroupoid_compose(gopurs_runtime.CoerceToStruct[Constructor_Control_Semigroupoid_Semigroupoid[gopurs_runtime.Value]](Get_Control_Semigroupoid_semigroupoidFn())), Call_Data_EuclideanRing_degree(gopurs_runtime.CoerceToStruct[Constructor_Data_EuclideanRing_EuclideanRing[gopurs_runtime.Value]](dictEuclideanRing_0)), Get_Data_Lazy_force()), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "div"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_3))
}))
}), gopurs_runtime.Func2(func(a_2 gopurs_runtime.Value, b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictEuclideanRing_0, "mod"), gopurs_runtime.Apply(Get_Data_Lazy_force(), a_2), gopurs_runtime.Apply(Get_Data_Lazy_force(), b_3))
}))
})}))}
}

func Call_Data_Lazy_boundedLazy(dictBounded_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBounded_0 gopurs_runtime.Value = dictBounded_0_loop
_ = dictBounded_0
// TAST (Let): ordLazy1_1_0 shape=App(Var) bindingType=(ADT ["Data","Ord","Ord"] [(ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope185)])])
ordLazy1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](Call_Data_Lazy_ordLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBounded_0, "Ord0"), gopurs_runtime.Value{})))
_ = ordLazy1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3510799738, UnsafePtr: unsafe.Pointer((&Constructor_Data_Bounded_Bounded[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(ordLazy1_1_0)}
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBounded_0, "bottom")
})), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictBounded_0, "top")
}))}))}
}

func Call_Data_Lazy_heytingAlgebraLazy(dictHeytingAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictHeytingAlgebra_0 gopurs_runtime.Value = dictHeytingAlgebra_0_loop
_ = dictHeytingAlgebra_0
// TAST (Let): implies_1_0 shape=App(Var) bindingType=(Func [(TypeVar a$scope93), (TypeVar a$scope93)] (TypeVar a$scope93))
implies_1_0 := Call_Data_HeytingAlgebra_implies(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_0))
_ = implies_1_0
// TAST (Let): conj_2_1 shape=App(Var) bindingType=(Func [(TypeVar a$scope93), (TypeVar a$scope93)] (TypeVar a$scope93))
conj_2_1 := Call_Data_HeytingAlgebra_conj(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_0))
_ = conj_2_1
// TAST (Let): disj_3_2 shape=App(Var) bindingType=(Func [(TypeVar a$scope93), (TypeVar a$scope93)] (TypeVar a$scope93))
disj_3_2 := Call_Data_HeytingAlgebra_disj(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_0))
_ = disj_3_2
// TAST (Let): not_4_3 shape=App(Var) bindingType=(Func [(TypeVar a$scope93)] (TypeVar a$scope93))
not_4_3 := Call_Data_HeytingAlgebra_not(gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](dictHeytingAlgebra_0))
_ = not_4_3
return gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer((&Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]{1, gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_4 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(Func [(TypeVar a$scope93)] (TypeVar a$scope93))])
__local_var_7_4 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(conj_2_1, gopurs_runtime.Apply(Get_Data_Lazy_force(), a_5))
}))
_ = __local_var_7_4
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_7_4, gopurs_runtime.Apply(Get_Data_Lazy_force(), b_6))
}))
}), gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(Func [(TypeVar a$scope93)] (TypeVar a$scope93))])
__local_var_7_5 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(disj_3_2, gopurs_runtime.Apply(Get_Data_Lazy_force(), a_5))
}))
_ = __local_var_7_5
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_7_5, gopurs_runtime.Apply(Get_Data_Lazy_force(), b_6))
}))
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "ff")
})), gopurs_runtime.Func2(func(a_5 gopurs_runtime.Value, b_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_6 shape=App(Var) bindingType=(ADT ["Data","Lazy","Lazy"] [(Func [(TypeVar a$scope93)] (TypeVar a$scope93))])
__local_var_7_6 := gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(implies_1_0, gopurs_runtime.Apply(Get_Data_Lazy_force(), a_5))
}))
_ = __local_var_7_6
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Get_Data_Lazy_force(), __local_var_7_6, gopurs_runtime.Apply(Get_Data_Lazy_force(), b_6))
}))
}), gopurs_runtime.Func(func(a_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(not_4_3, gopurs_runtime.Apply(Get_Data_Lazy_force(), a_5))
}))
}), gopurs_runtime.Apply(Get_Data_Lazy_go__defer(), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(dictHeytingAlgebra_0, "tt")
}))}))}
}

func Call_Data_Lazy_booleanAlgebraLazy(dictBooleanAlgebra_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictBooleanAlgebra_0 gopurs_runtime.Value = dictBooleanAlgebra_0_loop
_ = dictBooleanAlgebra_0
// TAST (Let): heytingAlgebraLazy1_1_0 shape=App(Var) bindingType=(ADT ["Data","HeytingAlgebra","HeytingAlgebra"] [(ADT ["Data","Lazy","Lazy"] [(TypeVar a$scope190)])])
heytingAlgebraLazy1_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_HeytingAlgebra_HeytingAlgebra[gopurs_runtime.Value]](Call_Data_Lazy_heytingAlgebraLazy(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictBooleanAlgebra_0, "HeytingAlgebra0"), gopurs_runtime.Value{})))
_ = heytingAlgebraLazy1_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 3257204378, UnsafePtr: unsafe.Pointer((&Constructor_Data_BooleanAlgebra_BooleanAlgebra[gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 926771738, UnsafePtr: unsafe.Pointer(heytingAlgebraLazy1_1_0)}
})}))}
}

func Get_Data_Lazy_go__defer() gopurs_runtime.Value {
	return _Gopurs_Data_Lazy_Go__defer
}

func Get_Data_Lazy_force() gopurs_runtime.Value {
	return _Gopurs_Data_Lazy_Force
}

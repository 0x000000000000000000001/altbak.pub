package Data_Array_NonEmpty_Internal

import (
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Control_Semigroupoid "gopurs/output/Control.Semigroupoid"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Maybe "gopurs/output/Data.Maybe"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Semigroup_Traversable "gopurs/output/Data.Semigroup.Traversable"
	pkg_Data_Show "gopurs/output/Data.Show"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_TraversableWithIndex "gopurs/output/Data.TraversableWithIndex"
	pkg_Data_Tuple "gopurs/output/Data.Tuple"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Partial_Unsafe "gopurs/output/Partial.Unsafe"
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_NonEmptyArray gopurs_runtime.Value
var once_NonEmptyArray sync.Once
func Get_NonEmptyArray() gopurs_runtime.Value {
	once_NonEmptyArray.Do(func() {
		cache_NonEmptyArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_NonEmptyArray(x_0_box)
})
	})
	return cache_NonEmptyArray
}

var cache_unfoldable1NonEmptyArray gopurs_runtime.Value
var once_unfoldable1NonEmptyArray sync.Once
func Get_unfoldable1NonEmptyArray() gopurs_runtime.Value {
	once_unfoldable1NonEmptyArray.Do(func() {
		cache_unfoldable1NonEmptyArray = pkg_Data_Unfoldable1.Get_unfoldable1Array()
	})
	return cache_unfoldable1NonEmptyArray
}

var cache_traversableWithIndexNonEmptyArray gopurs_runtime.Value
var once_traversableWithIndexNonEmptyArray sync.Once
func Get_traversableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_traversableWithIndexNonEmptyArray.Do(func() {
		cache_traversableWithIndexNonEmptyArray = pkg_Data_TraversableWithIndex.Get_traversableWithIndexArray()
	})
	return cache_traversableWithIndexNonEmptyArray
}

var cache_traversableNonEmptyArray gopurs_runtime.Value
var once_traversableNonEmptyArray sync.Once
func Get_traversableNonEmptyArray() gopurs_runtime.Value {
	once_traversableNonEmptyArray.Do(func() {
		cache_traversableNonEmptyArray = pkg_Data_Traversable.Get_traversableArray()
	})
	return cache_traversableNonEmptyArray
}

var cache_showNonEmptyArray gopurs_runtime.Value
var once_showNonEmptyArray sync.Once
func Get_showNonEmptyArray() gopurs_runtime.Value {
	once_showNonEmptyArray.Do(func() {
		cache_showNonEmptyArray = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_showNonEmptyArray(dictShow_0_box)
})
	})
	return cache_showNonEmptyArray
}

var cache_semigroupNonEmptyArray gopurs_runtime.Value
var once_semigroupNonEmptyArray sync.Once
func Get_semigroupNonEmptyArray() gopurs_runtime.Value {
	once_semigroupNonEmptyArray.Do(func() {
		cache_semigroupNonEmptyArray = pkg_Data_Semigroup.Get_semigroupArray()
	})
	return cache_semigroupNonEmptyArray
}

var cache_ordNonEmptyArray gopurs_runtime.Value
var once_ordNonEmptyArray sync.Once
func Get_ordNonEmptyArray() gopurs_runtime.Value {
	once_ordNonEmptyArray.Do(func() {
		cache_ordNonEmptyArray = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_ordNonEmptyArray(dictOrd_0_box)
})
	})
	return cache_ordNonEmptyArray
}

var cache_ord1NonEmptyArray gopurs_runtime.Value
var once_ord1NonEmptyArray sync.Once
func Get_ord1NonEmptyArray() gopurs_runtime.Value {
	once_ord1NonEmptyArray.Do(func() {
		cache_ord1NonEmptyArray = pkg_Data_Ord.Get_ord1Array()
	})
	return cache_ord1NonEmptyArray
}

var cache_monadNonEmptyArray gopurs_runtime.Value
var once_monadNonEmptyArray sync.Once
func Get_monadNonEmptyArray() gopurs_runtime.Value {
	once_monadNonEmptyArray.Do(func() {
		cache_monadNonEmptyArray = pkg_Control_Monad.Get_monadArray()
	})
	return cache_monadNonEmptyArray
}

var cache_functorWithIndexNonEmptyArray gopurs_runtime.Value
var once_functorWithIndexNonEmptyArray sync.Once
func Get_functorWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_functorWithIndexNonEmptyArray.Do(func() {
		cache_functorWithIndexNonEmptyArray = pkg_Data_FunctorWithIndex.Get_functorWithIndexArray()
	})
	return cache_functorWithIndexNonEmptyArray
}

var cache_functorNonEmptyArray gopurs_runtime.Value
var once_functorNonEmptyArray sync.Once
func Get_functorNonEmptyArray() gopurs_runtime.Value {
	once_functorNonEmptyArray.Do(func() {
		cache_functorNonEmptyArray = pkg_Data_Functor.Get_functorArray()
	})
	return cache_functorNonEmptyArray
}

var cache_foldableWithIndexNonEmptyArray gopurs_runtime.Value
var once_foldableWithIndexNonEmptyArray sync.Once
func Get_foldableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_foldableWithIndexNonEmptyArray.Do(func() {
		cache_foldableWithIndexNonEmptyArray = pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray()
	})
	return cache_foldableWithIndexNonEmptyArray
}

var cache_foldableNonEmptyArray gopurs_runtime.Value
var once_foldableNonEmptyArray sync.Once
func Get_foldableNonEmptyArray() gopurs_runtime.Value {
	once_foldableNonEmptyArray.Do(func() {
		cache_foldableNonEmptyArray = pkg_Data_Foldable.Get_foldableArray()
	})
	return cache_foldableNonEmptyArray
}

var cache_foldable1NonEmptyArray gopurs_runtime.Value
var once_foldable1NonEmptyArray sync.Once
func Get_foldable1NonEmptyArray() gopurs_runtime.Value {
	once_foldable1NonEmptyArray.Do(func() {
		cache_foldable1NonEmptyArray = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), f_2)
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldable1NonEmptyArray(), "foldl1"), append_1_0)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
})
})
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_foldl1Impl(), __local_var_0, __local_var_1)
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_foldr1Impl(), __local_var_0, __local_var_1)
}))
	})
	return cache_foldable1NonEmptyArray
}

var cache_traversable1NonEmptyArray gopurs_runtime.Value
var once_traversable1NonEmptyArray sync.Once
func Get_traversable1NonEmptyArray() gopurs_runtime.Value {
	once_traversable1NonEmptyArray.Do(func() {
		cache_traversable1NonEmptyArray = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_foldable1NonEmptyArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableArray()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversable1NonEmptyArray(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Apply.Constructor_Apply[gopurs_runtime.Value]](dictApply_0))}, pkg_Data_Semigroup_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
apply_1_0 := gopurs_runtime.RecordGet(dictApply_0, "apply")
_ = apply_1_0
go__map_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map")
_ = go__map_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_traverse1Impl(), apply_1_0, go__map_2_1, f_3)
})
}))
	})
	return cache_traversable1NonEmptyArray
}

var cache_eqNonEmptyArray gopurs_runtime.Value
var once_eqNonEmptyArray sync.Once
func Get_eqNonEmptyArray() gopurs_runtime.Value {
	once_eqNonEmptyArray.Do(func() {
		cache_eqNonEmptyArray = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_eqNonEmptyArray(dictEq_0_box)
})
	})
	return cache_eqNonEmptyArray
}

var cache_eq1NonEmptyArray gopurs_runtime.Value
var once_eq1NonEmptyArray sync.Once
func Get_eq1NonEmptyArray() gopurs_runtime.Value {
	once_eq1NonEmptyArray.Do(func() {
		cache_eq1NonEmptyArray = pkg_Data_Eq.Get_eq1Array()
	})
	return cache_eq1NonEmptyArray
}

var cache_bindNonEmptyArray gopurs_runtime.Value
var once_bindNonEmptyArray sync.Once
func Get_bindNonEmptyArray() gopurs_runtime.Value {
	once_bindNonEmptyArray.Do(func() {
		cache_bindNonEmptyArray = pkg_Control_Bind.Get_bindArray()
	})
	return cache_bindNonEmptyArray
}

var cache_applyNonEmptyArray gopurs_runtime.Value
var once_applyNonEmptyArray sync.Once
func Get_applyNonEmptyArray() gopurs_runtime.Value {
	once_applyNonEmptyArray.Do(func() {
		cache_applyNonEmptyArray = pkg_Control_Apply.Get_applyArray()
	})
	return cache_applyNonEmptyArray
}

var cache_applicativeNonEmptyArray gopurs_runtime.Value
var once_applicativeNonEmptyArray sync.Once
func Get_applicativeNonEmptyArray() gopurs_runtime.Value {
	once_applicativeNonEmptyArray.Do(func() {
		cache_applicativeNonEmptyArray = pkg_Control_Applicative.Get_applicativeArray()
	})
	return cache_applicativeNonEmptyArray
}

var cache_altNonEmptyArray gopurs_runtime.Value
var once_altNonEmptyArray sync.Once
func Get_altNonEmptyArray() gopurs_runtime.Value {
	once_altNonEmptyArray.Do(func() {
		cache_altNonEmptyArray = pkg_Control_Alt.Get_altArray()
	})
	return cache_altNonEmptyArray
}

var cache_altArray__2010533188 gopurs_runtime.Value
var once_altArray__2010533188 sync.Once
func Get_altArray__2010533188() gopurs_runtime.Value {
	once_altArray__2010533188.Do(func() {
		cache_altArray__2010533188 = gopurs_runtime.RecordDict2("Functor0", "alt", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"))
	})
	return cache_altArray__2010533188
}

var cache_applicativeArray__1604836744 gopurs_runtime.Value
var once_applicativeArray__1604836744 sync.Once
func Get_applicativeArray__1604836744() gopurs_runtime.Value {
	once_applicativeArray__1604836744.Do(func() {
		cache_applicativeArray__1604836744 = gopurs_runtime.RecordDict2("Apply0", "pure", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
}))
	})
	return cache_applicativeArray__1604836744
}

var cache_applyArray__2998472828 gopurs_runtime.Value
var once_applyArray__2998472828 sync.Once
func Get_applyArray__2998472828() gopurs_runtime.Value {
	once_applyArray__2998472828.Do(func() {
		cache_applyArray__2998472828 = gopurs_runtime.RecordDict2("Functor0", "apply", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Control_Apply.Get_arrayApply())
	})
	return cache_applyArray__2998472828
}

var cache_bindArray__1650562023 gopurs_runtime.Value
var once_bindArray__1650562023 sync.Once
func Get_bindArray__1650562023() gopurs_runtime.Value {
	once_bindArray__1650562023.Do(func() {
		cache_bindArray__1650562023 = gopurs_runtime.RecordDict2("Apply0", "bind", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Apply.Get_applyArray()
}), pkg_Control_Bind.Get_arrayBind())
	})
	return cache_bindArray__1650562023
}

var cache_monadArray__2289780851 gopurs_runtime.Value
var once_monadArray__2289780851 sync.Once
func Get_monadArray__2289780851() gopurs_runtime.Value {
	once_monadArray__2289780851.Do(func() {
		cache_monadArray__2289780851 = gopurs_runtime.RecordDict2("Applicative0", "Bind1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Applicative.Get_applicativeArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Control_Bind.Get_bindArray()
}))
	})
	return cache_monadArray__2289780851
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

var cache_foldable1NonEmptyArray__4102232191 gopurs_runtime.Value
var once_foldable1NonEmptyArray__4102232191 sync.Once
func Get_foldable1NonEmptyArray__4102232191() gopurs_runtime.Value {
	once_foldable1NonEmptyArray__4102232191.Do(func() {
		cache_foldable1NonEmptyArray__4102232191 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_Functor.Get_functorArray(), "map"), f_2)
_ = __local_var_3_1
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_foldable1NonEmptyArray(), "foldl1"), append_1_0)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
})
})
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_foldl1Impl(), __local_var_0, __local_var_1)
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_foldr1Impl(), __local_var_0, __local_var_1)
}))
	})
	return cache_foldable1NonEmptyArray__4102232191
}

var cache_foldableNonEmptyArray__3859409398 gopurs_runtime.Value
var once_foldableNonEmptyArray__3859409398 sync.Once
func Get_foldableNonEmptyArray__3859409398() gopurs_runtime.Value {
	once_foldableNonEmptyArray__3859409398.Do(func() {
		cache_foldableNonEmptyArray__3859409398 = pkg_Data_Foldable.Get_foldableArray()
	})
	return cache_foldableNonEmptyArray__3859409398
}

var cache_traversableNonEmptyArray__2643873085 gopurs_runtime.Value
var once_traversableNonEmptyArray__2643873085 sync.Once
func Get_traversableNonEmptyArray__2643873085() gopurs_runtime.Value {
	once_traversableNonEmptyArray__2643873085.Do(func() {
		cache_traversableNonEmptyArray__2643873085 = pkg_Data_Traversable.Get_traversableArray()
	})
	return cache_traversableNonEmptyArray__2643873085
}

var cache_eq1Array__2389734302 gopurs_runtime.Value
var once_eq1Array__2389734302 sync.Once
func Get_eq1Array__2389734302() gopurs_runtime.Value {
	once_eq1Array__2389734302.Do(func() {
		cache_eq1Array__2389734302 = gopurs_runtime.RecordDict1("eq1", gopurs_runtime.Func(func(dictEq_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))
}))
	})
	return cache_eq1Array__2389734302
}

var cache_foldableArray__2950015754 gopurs_runtime.Value
var once_foldableArray__2950015754 sync.Once
func Get_foldableArray__2950015754() gopurs_runtime.Value {
	once_foldableArray__2950015754.Do(func() {
		cache_foldableArray__2950015754 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), pkg_Data_Foldable.Get_foldlArray(), pkg_Data_Foldable.Get_foldrArray())
	})
	return cache_foldableArray__2950015754
}

var cache_foldableArray__3859409398 gopurs_runtime.Value
var once_foldableArray__3859409398 sync.Once
func Get_foldableArray__3859409398() gopurs_runtime.Value {
	once_foldableArray__3859409398.Do(func() {
		cache_foldableArray__3859409398 = gopurs_runtime.RecordDict3("foldMap", "foldl", "foldr", gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply(f_3, x_4), acc_5)
})
}), mempty_2_1)
})
}), pkg_Data_Foldable.Get_foldlArray(), pkg_Data_Foldable.Get_foldrArray())
	})
	return cache_foldableArray__3859409398
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

var cache_foldableWithIndexArray__740253118 gopurs_runtime.Value
var once_foldableWithIndexArray__740253118 sync.Once
func Get_foldableWithIndexArray__740253118() gopurs_runtime.Value {
	once_foldableWithIndexArray__740253118.Do(func() {
		cache_foldableWithIndexArray__740253118 = gopurs_runtime.RecordDict4("Foldable0", "foldMapWithIndex", "foldlWithIndex", "foldrWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}), gopurs_runtime.Func(func(dictMonoid_0 gopurs_runtime.Value) gopurs_runtime.Value {
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictMonoid_0, "Semigroup0"), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
mempty_2_1 := gopurs_runtime.RecordGet(dictMonoid_0, "mempty")
_ = mempty_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray(), "foldrWithIndex"), gopurs_runtime.Func(func(i_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(acc_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Semigroup0_1_0.V0, gopurs_runtime.Apply2(f_3, i_4, x_5), acc_6)
})
})
}), mempty_2_1)
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_2 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldl"), gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int((*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V0.IntVal), y_2, (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_3.UnsafePtr).V1)
})
}), z_1)
_ = __local_var_2_2
__local_var_3_3 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), pkg_Data_Tuple.Get_Tuple())
_ = __local_var_3_3
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_2, gopurs_runtime.Apply(__local_var_3_3, x_4))
})
})
}), gopurs_runtime.Func(func(f_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(z_1 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_2_4 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Foldable.Get_foldableArray(), "foldr"), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_3_5 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V0
_ = __local_var_3_5
__local_var_4_6 := (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_2.UnsafePtr).V1
_ = __local_var_4_6
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(f_0, gopurs_runtime.Int(__local_var_3_5.IntVal), __local_var_4_6, y_5)
})
}), z_1)
_ = __local_var_2_4
__local_var_3_7 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex"), pkg_Data_Tuple.Get_Tuple())
_ = __local_var_3_7
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_2_4, gopurs_runtime.Apply(__local_var_3_7, x_4))
})
})
}))
	})
	return cache_foldableWithIndexArray__740253118
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

var cache_functorArray__361387505 gopurs_runtime.Value
var once_functorArray__361387505 sync.Once
func Get_functorArray__361387505() gopurs_runtime.Value {
	once_functorArray__361387505.Do(func() {
		cache_functorArray__361387505 = gopurs_runtime.RecordDict1("map", pkg_Data_Functor.Get_arrayMap())
	})
	return cache_functorArray__361387505
}

var cache_functorWithIndexArray__3811533158 gopurs_runtime.Value
var once_functorWithIndexArray__3811533158 sync.Once
func Get_functorWithIndexArray__3811533158() gopurs_runtime.Value {
	once_functorWithIndexArray__3811533158.Do(func() {
		cache_functorWithIndexArray__3811533158 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Data_FunctorWithIndex.Get_mapWithIndexArray())
	})
	return cache_functorWithIndexArray__3811533158
}

var cache_functorWithIndexArray__490015842 gopurs_runtime.Value
var once_functorWithIndexArray__490015842 sync.Once
func Get_functorWithIndexArray__490015842() gopurs_runtime.Value {
	once_functorWithIndexArray__490015842.Do(func() {
		cache_functorWithIndexArray__490015842 = gopurs_runtime.RecordDict2("Functor0", "mapWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), pkg_Data_FunctorWithIndex.Get_mapWithIndexArray())
	})
	return cache_functorWithIndexArray__490015842
}

var cache_mapWithIndex__835054498 gopurs_runtime.Value
var once_mapWithIndex__835054498 sync.Once
func Get_mapWithIndex__835054498() gopurs_runtime.Value {
	once_mapWithIndex__835054498.Do(func() {
		cache_mapWithIndex__835054498 = gopurs_runtime.RecordGet(pkg_Data_FunctorWithIndex.Get_functorWithIndexArray(), "mapWithIndex")
	})
	return cache_mapWithIndex__835054498
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

var cache_fromJust__1791383420 gopurs_runtime.Value
var once_fromJust__1791383420 sync.Once
func Get_fromJust__1791383420() gopurs_runtime.Value {
	once_fromJust__1791383420.Do(func() {
		cache_fromJust__1791383420 = gopurs_runtime.Func2(func(_dollar__unused_0_box gopurs_runtime.Value, v_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fromJust__1791383420(_dollar__unused_0_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v_1_box))
})
	})
	return cache_fromJust__1791383420
}

var cache_isNothing__2591355336 gopurs_runtime.Value
var once_isNothing__2591355336 sync.Once
func Get_isNothing__2591355336() gopurs_runtime.Value {
	once_isNothing__2591355336.Do(func() {
		cache_isNothing__2591355336 = gopurs_runtime.Func(func(v2_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_isNothing__2591355336(gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_0_box)))
})
	})
	return cache_isNothing__2591355336
}

var cache_maybe__1594528518 gopurs_runtime.Value
var once_maybe__1594528518 sync.Once
func Get_maybe__1594528518() gopurs_runtime.Value {
	once_maybe__1594528518.Do(func() {
		cache_maybe__1594528518 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_maybe__1594528518(v_0_box, v1_1_box, gopurs_runtime.CoerceToStruct[pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]](v2_2_box))
})
	})
	return cache_maybe__1594528518
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

var cache_ord1Array__2954870633 gopurs_runtime.Value
var once_ord1Array__2954870633 sync.Once
func Get_ord1Array__2954870633() gopurs_runtime.Value {
	once_ord1Array__2954870633.Do(func() {
		cache_ord1Array__2954870633 = gopurs_runtime.RecordDict2("Eq10", "compare1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Eq.Get_eq1Array()
}), gopurs_runtime.Func(func(dictOrd_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply(pkg_Data_Ord.Get_ordArray(), dictOrd_0), "compare")
}))
	})
	return cache_ord1Array__2954870633
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

var cache_semigroupArray__1728406699 gopurs_runtime.Value
var once_semigroupArray__1728406699 sync.Once
func Get_semigroupArray__1728406699() gopurs_runtime.Value {
	once_semigroupArray__1728406699.Do(func() {
		cache_semigroupArray__1728406699 = gopurs_runtime.RecordDict1("append", pkg_Data_Semigroup.Get_concatArray())
	})
	return cache_semigroupArray__1728406699
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

var cache_show__4031350738 gopurs_runtime.Value
var once_show__4031350738 sync.Once
func Get_show__4031350738() gopurs_runtime.Value {
	once_show__4031350738.Do(func() {
		cache_show__4031350738 = gopurs_runtime.Func(func(dict_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_show__4031350738(gopurs_runtime.CoerceToStruct[pkg_Data_Show.Constructor_Show[[]gopurs_runtime.Value]](dict_0_box))
})
	})
	return cache_show__4031350738
}

var cache_traversableArray__2643873085 gopurs_runtime.Value
var once_traversableArray__2643873085 sync.Once
func Get_traversableArray__2643873085() gopurs_runtime.Value {
	once_traversableArray__2643873085.Do(func() {
		cache_traversableArray__2643873085 = gopurs_runtime.RecordDict4("Foldable1", "Functor0", "sequence", "traverse", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Foldable.Get_foldableArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Functor.Get_functorArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Traversable.Get_traversableArray(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[pkg_Control_Applicative.Constructor_Applicative[gopurs_runtime.Value]](dictApplicative_0))}, pkg_Data_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
Apply0_1_0 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApplicative_0, "Apply0"), gopurs_runtime.Value{})
_ = Apply0_1_0
return gopurs_runtime.Apply4(pkg_Data_Traversable.Get_traverseArrayImpl(), gopurs_runtime.RecordGet(Apply0_1_0, "apply"), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(Apply0_1_0, "Functor0"), gopurs_runtime.Value{}), "map"), gopurs_runtime.RecordGet(dictApplicative_0, "pure"), gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupArray(), "append"))
}))
	})
	return cache_traversableArray__2643873085
}

var cache_traversableWithIndexArray__1681559805 gopurs_runtime.Value
var once_traversableWithIndexArray__1681559805 sync.Once
func Get_traversableWithIndexArray__1681559805() gopurs_runtime.Value {
	once_traversableWithIndexArray__1681559805.Do(func() {
		cache_traversableWithIndexArray__1681559805 = gopurs_runtime.RecordDict4("FoldableWithIndex1", "FunctorWithIndex0", "Traversable2", "traverseWithIndex", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FoldableWithIndex.Get_foldableWithIndexArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_FunctorWithIndex.Get_functorWithIndexArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return pkg_Data_Traversable.Get_traversableArray()
}), gopurs_runtime.Func(func(dictApplicative_0 gopurs_runtime.Value) gopurs_runtime.Value {
FunctorWithIndex0_1_0 := gopurs_runtime.CoerceToStruct[pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_TraversableWithIndex.Get_traversableWithIndexArray(), "FunctorWithIndex0"), gopurs_runtime.Value{}))
_ = FunctorWithIndex0_1_0
sequence1_2_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(pkg_Data_TraversableWithIndex.Get_traversableWithIndexArray(), "Traversable2"), gopurs_runtime.Value{}), "sequence"), dictApplicative_0)
_ = sequence1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_4_2 := gopurs_runtime.Apply(FunctorWithIndex0_1_0.V1, f_3)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(sequence1_2_1, gopurs_runtime.Apply(__local_var_4_2, x_5))
})
})
}))
	})
	return cache_traversableWithIndexArray__1681559805
}

var cache_fst__20422131 gopurs_runtime.Value
var once_fst__20422131 sync.Once
func Get_fst__20422131() gopurs_runtime.Value {
	once_fst__20422131.Do(func() {
		cache_fst__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_fst__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_fst__20422131
}

var cache_snd__20422131 gopurs_runtime.Value
var once_snd__20422131 sync.Once
func Get_snd__20422131() gopurs_runtime.Value {
	once_snd__20422131.Do(func() {
		cache_snd__20422131 = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_snd__20422131(gopurs_runtime.CoerceToStruct[pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v_0_box))
})
	})
	return cache_snd__20422131
}

var cache_unfoldable1Array__4196906331 gopurs_runtime.Value
var once_unfoldable1Array__4196906331 sync.Once
func Get_unfoldable1Array__4196906331() gopurs_runtime.Value {
	once_unfoldable1Array__4196906331.Do(func() {
		cache_unfoldable1Array__4196906331 = gopurs_runtime.RecordDict1("unfoldr1", gopurs_runtime.Apply4(pkg_Data_Unfoldable1.Get_unfoldr1ArrayImpl(), pkg_Data_Maybe.Get_isNothing(), gopurs_runtime.Apply(pkg_Partial_Unsafe.Get__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 gopurs_runtime.Value
{
if (v_1.Type == 9 && v_1.IntVal == 930809136 && v_1.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(v_1.UnsafePtr).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return __t0
})
})), pkg_Data_Tuple.Get_fst(), pkg_Data_Tuple.Get_snd()))
	})
	return cache_unfoldable1Array__4196906331
}

var cache_unsafePartial__1306634845 gopurs_runtime.Value
var once_unsafePartial__1306634845 sync.Once
func Get_unsafePartial__1306634845() gopurs_runtime.Value {
	once_unsafePartial__1306634845.Do(func() {
		cache_unsafePartial__1306634845 = pkg_Partial_Unsafe.Get__unsafePartial()
	})
	return cache_unsafePartial__1306634845
}

func Call_NonEmptyArray(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showNonEmptyArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
showArray_1_0 := &pkg_Data_Show.Constructor_Show[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(pkg_Data_Show.Get_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))}
_ = showArray_1_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str("(NonEmptyArray "), gopurs_runtime.Str(gopurs_runtime.Apply2(Get_append__493084344(), gopurs_runtime.Str(gopurs_runtime.Apply(showArray_1_0.V0, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal()), gopurs_runtime.Str(")")).StrVal())).StrVal())
}))
}

func Call_ordNonEmptyArray(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(pkg_Data_Ord.Get_ordArray(), dictOrd_0)
}

func Call_eqNonEmptyArray(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq")))
}

func Call_compose__858342840(dict_0_loop *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Control_Semigroupoid.Constructor_Semigroupoid[gopurs_runtime.Value] = dict_0_loop
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

func Call_const__641934996(a_0_loop gopurs_runtime.Value, v_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var v_1 gopurs_runtime.Value = v_1_loop
_ = v_1
return a_0
}

func Call_mapWithIndex__55256674(dict_0_loop *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_FunctorWithIndex.Constructor_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V1
}

func Call_fromJust__1791383420(_dollar__unused_0_loop gopurs_runtime.Value, v_1_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var v_1 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v_1_loop
_ = v_1
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr != nil) {
__t0 = (*pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v_1)}.UnsafePtr).V0
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

func Call_isNothing__2591355336(v2_0_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) bool {
var v2_0 *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value] = v2_0_loop
_ = v2_0
var __t0 gopurs_runtime.Value
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr == nil) {
__t0 = gopurs_runtime.Bool(true)
goto end_branch_0
} else {

}
}
{
if (gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.Type == 9 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.IntVal == 930809136 && gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v2_0)}.UnsafePtr != nil) {
__t0 = gopurs_runtime.Bool(false)
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return (__t0.IntVal) != (0)
}

func Call_maybe__1594528518(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop *pkg_Data_Maybe.Constructor_Just[gopurs_runtime.Value]) gopurs_runtime.Value {
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

func Call_append__1230318264(dict_0_loop *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Semigroup.Constructor_Semigroup[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__2742601362(dict_0_loop *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_show__4031350738(dict_0_loop *pkg_Data_Show.Constructor_Show[[]gopurs_runtime.Value]) gopurs_runtime.Value {
var dict_0 *pkg_Data_Show.Constructor_Show[[]gopurs_runtime.Value] = dict_0_loop
_ = dict_0
return dict_0.V0
}

func Call_fst__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V0
}

func Call_snd__20422131(v_0_loop *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) gopurs_runtime.Value {
var v_0 *pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = v_0_loop
_ = v_0
return (*pkg_Data_Tuple.Constructor_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v_0)}.UnsafePtr).V1
}

func Get_foldl1Impl() gopurs_runtime.Value {
	return _Gopurs_Foldl1Impl
}

func Get_foldr1Impl() gopurs_runtime.Value {
	return _Gopurs_Foldr1Impl
}

func Get_traverse1Impl() gopurs_runtime.Value {
	return _Gopurs_Traverse1Impl
}

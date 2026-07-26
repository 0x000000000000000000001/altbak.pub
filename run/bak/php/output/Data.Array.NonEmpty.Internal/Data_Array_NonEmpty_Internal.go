package Data_Array_NonEmpty_Internal

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	pkg_Data_Unfoldable1 "gopurs/output/Data.Unfoldable1"
	pkg_Data_TraversableWithIndex "gopurs/output/Data.TraversableWithIndex"
	pkg_Data_Traversable "gopurs/output/Data.Traversable"
	pkg_Data_Semigroup "gopurs/output/Data.Semigroup"
	pkg_Data_Ord "gopurs/output/Data.Ord"
	pkg_Control_Monad "gopurs/output/Control.Monad"
	pkg_Data_FunctorWithIndex "gopurs/output/Data.FunctorWithIndex"
	pkg_Data_Functor "gopurs/output/Data.Functor"
	pkg_Data_FoldableWithIndex "gopurs/output/Data.FoldableWithIndex"
	pkg_Data_Foldable "gopurs/output/Data.Foldable"
	pkg_Data_Semigroup_Foldable "gopurs/output/Data.Semigroup.Foldable"
	pkg_Data_Semigroup_Traversable "gopurs/output/Data.Semigroup.Traversable"
	pkg_Data_Eq "gopurs/output/Data.Eq"
	pkg_Control_Bind "gopurs/output/Control.Bind"
	pkg_Control_Apply "gopurs/output/Control.Apply"
	pkg_Control_Applicative "gopurs/output/Control.Applicative"
	pkg_Control_Alt "gopurs/output/Control.Alt"
	pkg_Data_Show "gopurs/output/Data.Show"
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
return gopurs_runtime.Apply3(pkg_Data_Semigroup_Foldable.Get_foldMap1DefaultL(), Get_foldable1NonEmptyArray(), pkg_Data_Functor.Get_functorArray(), dictSemigroup_0)
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
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_traversable1NonEmptyArray(), "traverse1"), dictApply_0, pkg_Data_Semigroup_Traversable.Get_identity())
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
apply_1_0 := gopurs_runtime.RecordGet(dictApply_0, "apply")
_ = apply_1_0
map__2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map")
_ = map__2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_traverse1Impl(), apply_1_0, map__2_1, f_3)
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

func Call_NonEmptyArray(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_showNonEmptyArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Str("(NonEmptyArray "), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(pkg_Data_Semigroup.Get_semigroupString(), "append"), gopurs_runtime.Apply2(pkg_Data_Show.Get_showArrayImpl(), ((*gopurs_runtime.RecordData1)(dictShow_0.UnsafePtr)).V0, v_1), gopurs_runtime.Str(")")))
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
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(pkg_Data_Eq.Get_eqArrayImpl(), ((*gopurs_runtime.RecordData1)(dictEq_0.UnsafePtr)).V0))
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

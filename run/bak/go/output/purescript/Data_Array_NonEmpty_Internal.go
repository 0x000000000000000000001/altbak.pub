package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Array_NonEmpty_Internal_NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_NonEmptyArray = gopurs_runtime.Func(func(x_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_Internal_NonEmptyArray(x_0_box)
})
	})
	return cache_Data_Array_NonEmpty_Internal_NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_Unfoldable1_unfoldable1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex](Get_Data_TraversableWithIndex_traversableWithIndexArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_traversableNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_Traversable_traversableArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_showNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_showNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_showNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_showNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_showNonEmptyArray = gopurs_runtime.Func(func(dictShow_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_Internal_showNonEmptyArray(dictShow_0_box)
})
	})
	return cache_Data_Array_NonEmpty_Internal_showNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](Get_Data_Semigroup_semigroupArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_ordNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_ordNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_ordNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_ordNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_ordNonEmptyArray = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_Internal_ordNonEmptyArray(dictOrd_0_box)
})
	})
	return cache_Data_Array_NonEmpty_Internal_ordNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_ord1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_ord1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_ord1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1](Get_Data_Ord_ord1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_monadNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_monadNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_monadNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](Get_Control_Monad_monadArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex](Get_Data_FunctorWithIndex_functorWithIndexArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_functorNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Functor_functorArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex](Get_Data_FoldableWithIndex_foldableWithIndexArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_foldableNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Foldable_foldableArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Foldable_foldableArray()))}
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): append_1_0 -> gopurs_runtime.Value
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_2)
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()).V2), append_1_0)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
})
})
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldl1Impl(), __local_var_0, __local_var_1)
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldr1Impl(), __local_var_0, __local_var_1)
})})}
	})
	return cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Traversable_Traversable1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()))}
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_Traversable_traversableArray()))}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1](Get_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray()).V3), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return x_1
}))
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): apply_1_0 -> gopurs_runtime.Value
apply_1_0 := gopurs_runtime.RecordGet(dictApply_0, "apply")
_ = apply_1_0
// TAST (Let): go__map_2_1 -> gopurs_runtime.Value
go__map_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map")
_ = go__map_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Data_Array_NonEmpty_Internal_traverse1Impl(), apply_1_0, go__map_2_1, f_3)
})
})})}
	})
	return cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_eqNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_eqNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_eqNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_eqNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_eqNonEmptyArray = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_NonEmpty_Internal_eqNonEmptyArray(dictEq_0_box)
})
	})
	return cache_Data_Array_NonEmpty_Internal_eqNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_eq1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_eq1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_eq1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1](Get_Data_Eq_eq1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_bindNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_bindNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_bindNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Bind_bindArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_applyNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_applyNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_applyNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](Get_Control_Apply_applyArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Applicative_applicativeArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_altNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_altNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_altNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_altNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_altNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](Get_Control_Alt_altArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_altNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__2055237815 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__2055237815 sync.Once
func Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__2055237815() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__2055237815.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__2055237815 = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(&Constructor_Data_Semigroup_Foldable_Foldable1{1, gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Foldable_foldableArray()))}
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): append_1_0 -> gopurs_runtime.Value
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldl1Impl(), append_1_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldl1Impl(), __local_var_0, __local_var_1)
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldr1Impl(), __local_var_0, __local_var_1)
})})}
	})
	return cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__2055237815
}

var cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__1758553428 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__1758553428 sync.Once
func Get_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__1758553428() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__1758553428.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__1758553428 = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](Get_Data_Foldable_foldableArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__1758553428
}

var cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796 sync.Once
func Get_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796 = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](Get_Data_Functor_functorArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796
}

var cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__1589012692 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__1589012692 sync.Once
func Get_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__1589012692() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__1589012692.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__1589012692 = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable](Get_Data_Traversable_traversableArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__1589012692
}

var cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500 sync.Once
func Get_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500.Do(func() {
		cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500 = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1](Get_Data_Unfoldable1_unfoldable1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500
}

func Call_Data_Array_NonEmpty_Internal_NonEmptyArray(x_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Array_NonEmpty_Internal_showNonEmptyArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showArray_1_0 -> *Constructor_Data_Show_Show
showArray_1_0 := &Constructor_Data_Show_Show{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))}
_ = showArray_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(&Constructor_Data_Show_Show{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyArray ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_1_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
})})}
}

func Call_Data_Array_NonEmpty_Internal_ordNonEmptyArray(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): eqArray_1_0 -> *Constructor_Data_Eq_Eq
eqArray_1_0 := &Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictOrd_0, "Eq0"), gopurs_runtime.Value{}), "eq"))}
_ = eqArray_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(eqArray_1_0)}
}), gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(ys_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply5(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil}, gopurs_runtime.Int(0), gopurs_runtime.Int(gopurs_runtime.Apply3(Get_Data_Ord_ordArrayImpl(), gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): v_6_1 -> gopurs_runtime.Value
v_6_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(dictOrd_0, "compare"), x_4, y_5)
_ = v_6_1
var __t2 int64
{
if (uint32(v_6_1.IntVal) == 902936544) {
__t2 = 0
goto end_branch_2
} else {

}
}
{
if (uint32(v_6_1.IntVal) == 1527465420) {
__t2 = 1
goto end_branch_2
} else {

}
}
{
if (uint32(v_6_1.IntVal) == 380165415) {
__t2 = -1
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_2:
return gopurs_runtime.Int(__t2)
})
}), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).IntVal)).IntVal)), UnsafePtr: nil}
})
})})}
}

func Call_Data_Array_NonEmpty_Internal_eqNonEmptyArray(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(&Constructor_Data_Eq_Eq{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))})}
}

func Get_Data_Array_NonEmpty_Internal_foldl1Impl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_NonEmpty_Internal_Foldl1Impl
}

func Get_Data_Array_NonEmpty_Internal_foldr1Impl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_NonEmpty_Internal_Foldr1Impl
}

func Get_Data_Array_NonEmpty_Internal_traverse1Impl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_NonEmpty_Internal_Traverse1Impl
}

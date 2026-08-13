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
		cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray = Get_Data_Unfoldable1_unfoldable1Array()
	})
	return cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray = Get_Data_TraversableWithIndex_traversableWithIndexArray()
	})
	return cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_traversableNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray = Get_Data_Traversable_traversableArray()
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
		cache_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray = Get_Data_Semigroup_semigroupArray()
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
		cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray = Get_Data_Ord_ord1Array()
	})
	return cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_monadNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_monadNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_monadNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray = Get_Control_Monad_monadArray()
	})
	return cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray = Get_Data_FunctorWithIndex_functorWithIndexArray()
	})
	return cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_functorNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray = Get_Data_Functor_functorArray()
	})
	return cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray = Get_Data_FoldableWithIndex_foldableWithIndexArray()
	})
	return cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_foldableNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray = Get_Data_Foldable_foldableArray()
	})
	return cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableArray()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): append_1_0 -> gopurs_runtime.Value
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), f_2)
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldl1"), append_1_0)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
})
})
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldl1Impl(), __local_var_0, __local_var_1)
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldr1Impl(), __local_var_0, __local_var_1)
}))
	})
	return cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray = gopurs_runtime.RecordDict4("Foldable10", "Traversable1", "sequence1", "traverse1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()
}), gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Traversable_traversableArray()
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray(), "traverse1"), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](dictApply_0))}, gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
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
}))
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
		cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray = Get_Data_Eq_eq1Array()
	})
	return cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_bindNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_bindNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_bindNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray = Get_Control_Bind_bindArray()
	})
	return cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_applyNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_applyNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_applyNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray = Get_Control_Apply_applyArray()
	})
	return cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray = Get_Control_Applicative_applicativeArray()
	})
	return cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_altNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_altNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_altNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_altNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_altNonEmptyArray = Get_Control_Alt_altArray()
	})
	return cache_Data_Array_NonEmpty_Internal_altNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__4102232191 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__4102232191 sync.Once
func Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__4102232191() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__4102232191.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__4102232191 = gopurs_runtime.RecordDict4("Foldable0", "foldMap1", "foldl1", "foldr1", gopurs_runtime.Func(func(_dollar__unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Foldable_foldableArray()
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): append_1_0 -> gopurs_runtime.Value
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), f_2)
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray(), "foldl1"), append_1_0)
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
})
})
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldl1Impl(), __local_var_0, __local_var_1)
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldr1Impl(), __local_var_0, __local_var_1)
}))
	})
	return cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray__4102232191
}

var cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__3859409398 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__3859409398 sync.Once
func Get_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__3859409398() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__3859409398.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__3859409398 = Get_Data_Foldable_foldableArray()
	})
	return cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray__3859409398
}

var cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796 sync.Once
func Get_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796 = Get_Data_Functor_functorArray()
	})
	return cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray__2527715796
}

var cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__2643873085 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__2643873085 sync.Once
func Get_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__2643873085() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__2643873085.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__2643873085 = Get_Data_Traversable_traversableArray()
	})
	return cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray__2643873085
}

var cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500 gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500 sync.Once
func Get_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500.Do(func() {
		cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray__3769668500 = Get_Data_Unfoldable1_unfoldable1Array()
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
return gopurs_runtime.RecordDict1("show", gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyArray ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_1_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
}))
}

func Call_Data_Array_NonEmpty_Internal_ordNonEmptyArray(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Ord_ordArray(), dictOrd_0)
}

func Call_Data_Array_NonEmpty_Internal_eqNonEmptyArray(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.RecordDict1("eq", gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq")))
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

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
return gopurs_runtime.Array(Call_Data_Array_NonEmpty_Internal_NonEmptyArray(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_NonEmpty_Internal_NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 3553002490, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable1_Unfoldable1[gopurs_runtime.Value]](Get_Data_Unfoldable1_unfoldable1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_unfoldable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 2078610234, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_2955889203_1812164904(gopurs_runtime.CoerceToStruct[Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]](Get_Data_TraversableWithIndex_traversableWithIndexArray())))}
	})
	return cache_Data_Array_NonEmpty_Internal_traversableWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_traversableNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversableNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversableNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}
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
		cache_Data_Array_NonEmpty_Internal_semigroupNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 2053112122, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_986430664_4179793454(Rebox_Data_Array_NonEmpty_Internal_4179793454_986430664(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](Get_Data_Semigroup_semigroupArray()))))}
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
		cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1632188299, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord1[gopurs_runtime.Value]](Get_Data_Ord_ord1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_ord1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_monadNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_monadNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_monadNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 778916621, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](Get_Control_Monad_monadArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_monadNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4077743418, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_2773701683_2412140840(gopurs_runtime.CoerceToStruct[Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]](Get_Data_FunctorWithIndex_functorWithIndexArray())))}
	})
	return cache_Data_Array_NonEmpty_Internal_functorWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_functorNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_functorNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_functorNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 929368378, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](Get_Data_Functor_functorArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_functorNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 74250362, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_2491554675_3725484264(gopurs_runtime.CoerceToStruct[Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]](Get_Data_FoldableWithIndex_foldableWithIndexArray())))}
	})
	return cache_Data_Array_NonEmpty_Internal_foldableWithIndexNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_foldableNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldableNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_foldableNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_3016437835_4151366573((&Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))}
}), gopurs_runtime.Func(func(dictSemigroup_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): append_1_0 shape=Other bindingType=(Func [(TypeVar m), (TypeVar m)] (TypeVar m))
append_1_0 := gopurs_runtime.RecordGet(dictSemigroup_0, "append")
_ = append_1_0
return gopurs_runtime.Func(func(f_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Var) bindingType=(Func [(TypeApp (TypeVar t) [(TypeVar a)])] (TypeApp (TypeVar t) [(TypeVar m)]))
__local_var_3_1 := gopurs_runtime.Apply(Get_Data_Functor_arrayMap(), f_2)
_ = __local_var_3_1
return gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray()).V2), append_1_0, gopurs_runtime.Apply(__local_var_3_1, x_4))
})
})
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldl1Impl(), __local_var_0, __local_var_1)
}), gopurs_runtime.Func2(func(__local_var_0 gopurs_runtime.Value, __local_var_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_NonEmpty_Internal_foldr1Impl(), __local_var_0, __local_var_1)
})})))}
	})
	return cache_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1596088409, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_3516332491_306175789((&Constructor_Data_Semigroup_Traversable_Traversable1[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2465059545, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_3016437835_4151366573(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_foldable1NonEmptyArray())))}
}), gopurs_runtime.Func(func(_dollar___unused_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 3941073978, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()))}
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Traversable_Traversable1[[]gopurs_runtime.Value]](Get_Data_Array_NonEmpty_Internal_traversable1NonEmptyArray()).V3), gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](dictApply_0))}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}), gopurs_runtime.Func(func(dictApply_0 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): apply_1_0 shape=Other bindingType=(Func [(TypeApp (TypeVar f) [(Func [(TypeVar a')] (TypeVar b'))]), (TypeApp (TypeVar f) [(TypeVar a')])] (TypeApp (TypeVar f) [(TypeVar b')]))
apply_1_0 := gopurs_runtime.RecordGet(dictApply_0, "apply")
_ = apply_1_0
// TAST (Let): go__map_2_1 shape=Other bindingType=(Func [(Func [(TypeVar a')] (TypeVar b')), (TypeApp (TypeVar f) [(TypeVar a')])] (TypeApp (TypeVar f) [(TypeVar b')]))
go__map_2_1 := gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictApply_0, "Functor0"), gopurs_runtime.Value{}), "map")
_ = go__map_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Data_Array_NonEmpty_Internal_traverse1Impl(), apply_1_0, go__map_2_1, f_3)
})
})})))}
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
		cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1715248107, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq1[gopurs_runtime.Value]](Get_Data_Eq_eq1Array()))}
	})
	return cache_Data_Array_NonEmpty_Internal_eq1NonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_bindNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_bindNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_bindNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](Get_Control_Bind_bindArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_bindNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_applyNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_applyNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_applyNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 3032403085, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](Get_Control_Apply_applyArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_applyNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Applicative_applicativeArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_applicativeNonEmptyArray
}

var cache_Data_Array_NonEmpty_Internal_altNonEmptyArray gopurs_runtime.Value
var once_Data_Array_NonEmpty_Internal_altNonEmptyArray sync.Once
func Get_Data_Array_NonEmpty_Internal_altNonEmptyArray() gopurs_runtime.Value {
	once_Data_Array_NonEmpty_Internal_altNonEmptyArray.Do(func() {
		cache_Data_Array_NonEmpty_Internal_altNonEmptyArray = gopurs_runtime.Value{Type: 9, IntVal: 4060500237, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](Get_Control_Alt_altArray()))}
	})
	return cache_Data_Array_NonEmpty_Internal_altNonEmptyArray
}

func Call_Data_Array_NonEmpty_Internal_NonEmptyArray(x_0_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 []gopurs_runtime.Value = x_0_loop
_ = x_0
return x_0
}

func Call_Data_Array_NonEmpty_Internal_showNonEmptyArray(dictShow_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictShow_0 gopurs_runtime.Value = dictShow_0_loop
_ = dictShow_0
// TAST (Let): showArray_1_0 shape=LitRecord bindingType=(ADT ["Data","Show","Show"] [(Array (TypeVar a))])
showArray_1_0 := (&Constructor_Data_Show_Show[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Show_showArrayImpl(), gopurs_runtime.RecordGet(dictShow_0, "show"))})
_ = showArray_1_0
return gopurs_runtime.Value{Type: 9, IntVal: 1835580986, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_1356436936_1386611502((&Constructor_Data_Show_Show[[]gopurs_runtime.Value]{1, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Str((("(NonEmptyArray ") + (gopurs_runtime.Apply(gopurs_runtime.Box(showArray_1_0.V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).StrVal())) + (")"))
})})))}
}

func Call_Data_Array_NonEmpty_Internal_ordNonEmptyArray(dictOrd_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 gopurs_runtime.Value = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_4088624520_4177771502(Rebox_Data_Array_NonEmpty_Internal_4177771502_4088624520(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Data_Ord_ordArray(), dictOrd_0)))))}
}

func Call_Data_Array_NonEmpty_Internal_eqNonEmptyArray(dictEq_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 gopurs_runtime.Value = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_NonEmpty_Internal_1939691112_3790796878((&Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]{1, gopurs_runtime.Apply(Get_Data_Eq_eqArrayImpl(), gopurs_runtime.RecordGet(dictEq_0, "eq"))})))}
}

func Rebox_Data_Array_NonEmpty_Internal_1356436936_1386611502(in *Constructor_Data_Show_Show[[]gopurs_runtime.Value]) *Constructor_Data_Show_Show[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Show_Show[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_1939691112_3790796878(in *Constructor_Data_Eq_Eq[[]gopurs_runtime.Value]) *Constructor_Data_Eq_Eq[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Eq_Eq[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_2491554675_3725484264(in *Constructor_Data_FoldableWithIndex_FoldableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FoldableWithIndex_FoldableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_2773701683_2412140840(in *Constructor_Data_FunctorWithIndex_FunctorWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_FunctorWithIndex_FunctorWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_2955889203_1812164904(in *Constructor_Data_TraversableWithIndex_TraversableWithIndex[int64, gopurs_runtime.Value]) *Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_TraversableWithIndex_TraversableWithIndex[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_3016437835_4151366573(in *Constructor_Data_Semigroup_Foldable_Foldable1[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Foldable_Foldable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_3516332491_306175789(in *Constructor_Data_Semigroup_Traversable_Traversable1[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Traversable_Traversable1[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
		out.V2 = in.V2
		out.V3 = in.V3
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_4088624520_4177771502(in *Constructor_Data_Ord_Ord[[]gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_4177771502_4088624520(in *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) *Constructor_Data_Ord_Ord[[]gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[[]gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_4179793454_986430664(in *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Array_NonEmpty_Internal_986430664_4179793454(in *Constructor_Data_Semigroup_Semigroup[[]gopurs_runtime.Value]) *Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
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

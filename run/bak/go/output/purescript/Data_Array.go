package purescript

import (
	gopurs_runtime "gopurs/output/gopurs_runtime"
	sync "sync"
	unsafe "unsafe"
)

var cache_Data_Array_intercalate1 gopurs_runtime.Value
var once_Data_Array_intercalate1 sync.Once
func Get_Data_Array_intercalate1() gopurs_runtime.Value {
	once_Data_Array_intercalate1.Do(func() {
		cache_Data_Array_intercalate1 = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_intercalate1(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0_box))
})
	})
	return cache_Data_Array_intercalate1
}

var cache_Data_Array_zero gopurs_runtime.Value
var once_Data_Array_zero sync.Once
func Get_Data_Array_zero() gopurs_runtime.Value {
	once_Data_Array_zero.Do(func() {
		cache_Data_Array_zero = gopurs_runtime.Int(0)
	})
	return cache_Data_Array_zero
}

var cache_Data_Array_one gopurs_runtime.Value
var once_Data_Array_one sync.Once
func Get_Data_Array_one() gopurs_runtime.Value {
	once_Data_Array_one.Do(func() {
		cache_Data_Array_one = gopurs_runtime.Int(1)
	})
	return cache_Data_Array_one
}

var cache_Data_Array_void gopurs_runtime.Value
var once_Data_Array_void sync.Once
func Get_Data_Array_void() gopurs_runtime.Value {
	once_Data_Array_void.Do(func() {
		cache_Data_Array_void = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_0 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}))
	})
	return cache_Data_Array_void
}

var cache_Data_Array_pure gopurs_runtime.Value
var once_Data_Array_pure sync.Once
func Get_Data_Array_pure() gopurs_runtime.Value {
	once_Data_Array_pure.Do(func() {
		cache_Data_Array_pure = gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure")
	})
	return cache_Data_Array_pure
}

var cache_Data_Array_fromJust gopurs_runtime.Value
var once_Data_Array_fromJust sync.Once
func Get_Data_Array_fromJust() gopurs_runtime.Value {
	once_Data_Array_fromJust.Do(func() {
		cache_Data_Array_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Array_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v_0_box)))}
})
	})
	return cache_Data_Array_fromJust
}

var cache_Data_Array_foldMap1 gopurs_runtime.Value
var once_Data_Array_foldMap1 sync.Once
func Get_Data_Array_foldMap1() gopurs_runtime.Value {
	once_Data_Array_foldMap1.Do(func() {
		cache_Data_Array_foldMap1 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldMap")
	})
	return cache_Data_Array_foldMap1
}

var cache_Data_Array_fold1 gopurs_runtime.Value
var once_Data_Array_fold1 sync.Once
func Get_Data_Array_fold1() gopurs_runtime.Value {
	once_Data_Array_fold1.Do(func() {
		cache_Data_Array_fold1 = gopurs_runtime.Apply(Get_Data_Foldable_fold(), Get_Data_Foldable_foldableArray())
	})
	return cache_Data_Array_fold1
}

var cache_Data_Array_not gopurs_runtime.Value
var once_Data_Array_not sync.Once
func Get_Data_Array_not() gopurs_runtime.Value {
	once_Data_Array_not.Do(func() {
		cache_Data_Array_not = gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not")
	})
	return cache_Data_Array_not
}

var cache_Data_Array_zipWith gopurs_runtime.Value
var once_Data_Array_zipWith sync.Once
func Get_Data_Array_zipWith() gopurs_runtime.Value {
	once_Data_Array_zipWith.Do(func() {
		cache_Data_Array_zipWith = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_zipWith(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_zipWith
}

var cache_Data_Array_zipWithA gopurs_runtime.Value
var once_Data_Array_zipWithA sync.Once
func Get_Data_Array_zipWithA() gopurs_runtime.Value {
	once_Data_Array_zipWithA.Do(func() {
		cache_Data_Array_zipWithA = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_3_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_zipWithA
}

var cache_Data_Array_zip gopurs_runtime.Value
var once_Data_Array_zip sync.Once
func Get_Data_Array_zip() gopurs_runtime.Value {
	once_Data_Array_zip.Do(func() {
		cache_Data_Array_zip = gopurs_runtime.Apply(Get_Data_Array_zipWith(), Get_Data_Tuple_Tuple())
	})
	return cache_Data_Array_zip
}

var cache_Data_Array_updateAtIndices gopurs_runtime.Value
var once_Data_Array_updateAtIndices sync.Once
func Get_Data_Array_updateAtIndices() gopurs_runtime.Value {
	once_Data_Array_updateAtIndices.Do(func() {
		cache_Data_Array_updateAtIndices = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, us_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_updateAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), us_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_updateAtIndices
}

var cache_Data_Array_updateAt gopurs_runtime.Value
var once_Data_Array_updateAt sync.Once
func Get_Data_Array_updateAt() gopurs_runtime.Value {
	once_Data_Array_updateAt.Do(func() {
		cache_Data_Array_updateAt = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_updateAt(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_updateAt
}

var cache_Data_Array_unsafeIndex gopurs_runtime.Value
var once_Data_Array_unsafeIndex sync.Once
func Get_Data_Array_unsafeIndex() gopurs_runtime.Value {
	once_Data_Array_unsafeIndex.Do(func() {
		cache_Data_Array_unsafeIndex = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_unsafeIndex(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_2_box.IntVal)
})
	})
	return cache_Data_Array_unsafeIndex
}

var cache_Data_Array_uncons gopurs_runtime.Value
var once_Data_Array_uncons sync.Once
func Get_Data_Array_uncons() gopurs_runtime.Value {
	once_Data_Array_uncons.Do(func() {
		cache_Data_Array_uncons = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_uncons(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_uncons
}

var cache_Data_Array_toUnfoldable gopurs_runtime.Value
var once_Data_Array_toUnfoldable sync.Once
func Get_Data_Array_toUnfoldable() gopurs_runtime.Value {
	once_Data_Array_toUnfoldable.Do(func() {
		cache_Data_Array_toUnfoldable = gopurs_runtime.Func2(func(dictUnfoldable_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable](dictUnfoldable_0_box), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_toUnfoldable
}

var cache_Data_Array_tail gopurs_runtime.Value
var once_Data_Array_tail sync.Once
func Get_Data_Array_tail() gopurs_runtime.Value {
	once_Data_Array_tail.Do(func() {
		cache_Data_Array_tail = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_tail(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_tail
}

var cache_Data_Array_sortBy gopurs_runtime.Value
var once_Data_Array_sortBy sync.Once
func Get_Data_Array_sortBy() gopurs_runtime.Value {
	once_Data_Array_sortBy.Do(func() {
		cache_Data_Array_sortBy = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_sortBy(comp_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_sortBy
}

var cache_Data_Array_sortWith gopurs_runtime.Value
var once_Data_Array_sortWith sync.Once
func Get_Data_Array_sortWith() gopurs_runtime.Value {
	once_Data_Array_sortWith.Do(func() {
		cache_Data_Array_sortWith = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Array_sortWith
}

var cache_Data_Array_sort gopurs_runtime.Value
var once_Data_Array_sort sync.Once
func Get_Data_Array_sort() gopurs_runtime.Value {
	once_Data_Array_sort.Do(func() {
		cache_Data_Array_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Array_sort
}

var cache_Data_Array_snoc gopurs_runtime.Value
var once_Data_Array_snoc sync.Once
func Get_Data_Array_snoc() gopurs_runtime.Value {
	once_Data_Array_snoc.Do(func() {
		cache_Data_Array_snoc = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_snoc(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_Data_Array_snoc
}

var cache_Data_Array_slice gopurs_runtime.Value
var once_Data_Array_slice sync.Once
func Get_Data_Array_slice() gopurs_runtime.Value {
	once_Data_Array_slice.Do(func() {
		cache_Data_Array_slice = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_slice(__local_var_0_box.IntVal, __local_var_1_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_slice
}

var cache_Data_Array_splitAt gopurs_runtime.Value
var once_Data_Array_splitAt sync.Once
func Get_Data_Array_splitAt() gopurs_runtime.Value {
	once_Data_Array_splitAt.Do(func() {
		cache_Data_Array_splitAt = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_splitAt(v_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v1_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_splitAt
}

var cache_Data_Array_take gopurs_runtime.Value
var once_Data_Array_take sync.Once
func Get_Data_Array_take() gopurs_runtime.Value {
	once_Data_Array_take.Do(func() {
		cache_Data_Array_take = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_take(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_take
}

var cache_Data_Array_singleton gopurs_runtime.Value
var once_Data_Array_singleton sync.Once
func Get_Data_Array_singleton() gopurs_runtime.Value {
	once_Data_Array_singleton.Do(func() {
		cache_Data_Array_singleton = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_singleton(a_0_box))
})
	})
	return cache_Data_Array_singleton
}

var cache_Data_Array_scanr gopurs_runtime.Value
var once_Data_Array_scanr sync.Once
func Get_Data_Array_scanr() gopurs_runtime.Value {
	once_Data_Array_scanr.Do(func() {
		cache_Data_Array_scanr = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_scanr(__local_var_0_box, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_scanr
}

var cache_Data_Array_scanl gopurs_runtime.Value
var once_Data_Array_scanl sync.Once
func Get_Data_Array_scanl() gopurs_runtime.Value {
	once_Data_Array_scanl.Do(func() {
		cache_Data_Array_scanl = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_scanl(__local_var_0_box, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_scanl
}

var cache_Data_Array_replicate gopurs_runtime.Value
var once_Data_Array_replicate sync.Once
func Get_Data_Array_replicate() gopurs_runtime.Value {
	once_Data_Array_replicate.Do(func() {
		cache_Data_Array_replicate = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_replicate(__local_var_0_box.IntVal, __local_var_1_box))
})
	})
	return cache_Data_Array_replicate
}

var cache_Data_Array_go__range gopurs_runtime.Value
var once_Data_Array_go__range sync.Once
func Get_Data_Array_go__range() gopurs_runtime.Value {
	once_Data_Array_go__range.Do(func() {
		cache_Data_Array_go__range = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_go__range(__local_var_0_box.IntVal, __local_var_1_box.IntVal)
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_go__range
}

var cache_Data_Array_partition gopurs_runtime.Value
var once_Data_Array_partition sync.Once
func Get_Data_Array_partition() gopurs_runtime.Value {
	once_Data_Array_partition.Do(func() {
		cache_Data_Array_partition = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_partition(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_partition
}

var cache_Data_Array_null gopurs_runtime.Value
var once_Data_Array_null sync.Once
func Get_Data_Array_null() gopurs_runtime.Value {
	once_Data_Array_null.Do(func() {
		cache_Data_Array_null = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_null(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_null
}

var cache_Data_Array_modifyAtIndices gopurs_runtime.Value
var once_Data_Array_modifyAtIndices sync.Once
func Get_Data_Array_modifyAtIndices() gopurs_runtime.Value {
	once_Data_Array_modifyAtIndices.Do(func() {
		cache_Data_Array_modifyAtIndices = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, is_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_modifyAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), is_1_box, f_2_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_3_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_modifyAtIndices
}

var cache_Data_Array_mapWithIndex gopurs_runtime.Value
var once_Data_Array_mapWithIndex sync.Once
func Get_Data_Array_mapWithIndex() gopurs_runtime.Value {
	once_Data_Array_mapWithIndex.Do(func() {
		cache_Data_Array_mapWithIndex = gopurs_runtime.RecordGet(Get_Data_FunctorWithIndex_functorWithIndexArray(), "mapWithIndex")
	})
	return cache_Data_Array_mapWithIndex
}

var cache_Data_Array_intersperse gopurs_runtime.Value
var once_Data_Array_intersperse sync.Once
func Get_Data_Array_intersperse() gopurs_runtime.Value {
	once_Data_Array_intersperse.Do(func() {
		cache_Data_Array_intersperse = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_intersperse(a_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_intersperse
}

var cache_Data_Array_intercalate gopurs_runtime.Value
var once_Data_Array_intercalate sync.Once
func Get_Data_Array_intercalate() gopurs_runtime.Value {
	once_Data_Array_intercalate.Do(func() {
		cache_Data_Array_intercalate = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0_box))
})
	})
	return cache_Data_Array_intercalate
}

var cache_Data_Array_insertAt gopurs_runtime.Value
var once_Data_Array_insertAt sync.Once
func Get_Data_Array_insertAt() gopurs_runtime.Value {
	once_Data_Array_insertAt.Do(func() {
		cache_Data_Array_insertAt = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_insertAt(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_insertAt
}

var cache_Data_Array_init gopurs_runtime.Value
var once_Data_Array_init sync.Once
func Get_Data_Array_init() gopurs_runtime.Value {
	once_Data_Array_init.Do(func() {
		cache_Data_Array_init = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_init(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_init
}

var cache_Data_Array_index gopurs_runtime.Value
var once_Data_Array_index sync.Once
func Get_Data_Array_index() gopurs_runtime.Value {
	once_Data_Array_index.Do(func() {
		cache_Data_Array_index = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_index(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_1_box.IntVal))}
})
	})
	return cache_Data_Array_index
}

var cache_Data_Array_last gopurs_runtime.Value
var once_Data_Array_last sync.Once
func Get_Data_Array_last() gopurs_runtime.Value {
	once_Data_Array_last.Do(func() {
		cache_Data_Array_last = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_last(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_last
}

var cache_Data_Array_unsnoc gopurs_runtime.Value
var once_Data_Array_unsnoc sync.Once
func Get_Data_Array_unsnoc() gopurs_runtime.Value {
	once_Data_Array_unsnoc.Do(func() {
		cache_Data_Array_unsnoc = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_unsnoc(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_unsnoc
}

var cache_Data_Array_modifyAt gopurs_runtime.Value
var once_Data_Array_modifyAt sync.Once
func Get_Data_Array_modifyAt() gopurs_runtime.Value {
	once_Data_Array_modifyAt.Do(func() {
		cache_Data_Array_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_modifyAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_modifyAt
}

var cache_Data_Array_unzip gopurs_runtime.Value
var once_Data_Array_unzip sync.Once
func Get_Data_Array_unzip() gopurs_runtime.Value {
	once_Data_Array_unzip.Do(func() {
		cache_Data_Array_unzip = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Array_unzip(func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_unzip
}

var cache_Data_Array_head gopurs_runtime.Value
var once_Data_Array_head sync.Once
func Get_Data_Array_head() gopurs_runtime.Value {
	once_Data_Array_head.Do(func() {
		cache_Data_Array_head = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_head(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_head
}

var cache_Data_Array_nubBy gopurs_runtime.Value
var once_Data_Array_nubBy sync.Once
func Get_Data_Array_nubBy() gopurs_runtime.Value {
	once_Data_Array_nubBy.Do(func() {
		cache_Data_Array_nubBy = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_nubBy(comp_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_nubBy
}

var cache_Data_Array_nub gopurs_runtime.Value
var once_Data_Array_nub sync.Once
func Get_Data_Array_nub() gopurs_runtime.Value {
	once_Data_Array_nub.Do(func() {
		cache_Data_Array_nub = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Array_nub
}

var cache_Data_Array_groupBy gopurs_runtime.Value
var once_Data_Array_groupBy sync.Once
func Get_Data_Array_groupBy() gopurs_runtime.Value {
	once_Data_Array_groupBy.Do(func() {
		cache_Data_Array_groupBy = gopurs_runtime.Func2(func(op_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_groupBy(op_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_groupBy
}

var cache_Data_Array_groupAllBy gopurs_runtime.Value
var once_Data_Array_groupAllBy sync.Once
func Get_Data_Array_groupAllBy() gopurs_runtime.Value {
	once_Data_Array_groupAllBy.Do(func() {
		cache_Data_Array_groupAllBy = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_groupAllBy(cmp_0_box)
})
	})
	return cache_Data_Array_groupAllBy
}

var cache_Data_Array_groupAll gopurs_runtime.Value
var once_Data_Array_groupAll sync.Once
func Get_Data_Array_groupAll() gopurs_runtime.Value {
	once_Data_Array_groupAll.Do(func() {
		cache_Data_Array_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Array_groupAll
}

var cache_Data_Array_group gopurs_runtime.Value
var once_Data_Array_group sync.Once
func Get_Data_Array_group() gopurs_runtime.Value {
	once_Data_Array_group.Do(func() {
		cache_Data_Array_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_group
}

var cache_Data_Array_fromFoldable gopurs_runtime.Value
var once_Data_Array_fromFoldable sync.Once
func Get_Data_Array_fromFoldable() gopurs_runtime.Value {
	once_Data_Array_fromFoldable.Do(func() {
		cache_Data_Array_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box))
})
	})
	return cache_Data_Array_fromFoldable
}

var cache_Data_Array_foldr gopurs_runtime.Value
var once_Data_Array_foldr sync.Once
func Get_Data_Array_foldr() gopurs_runtime.Value {
	once_Data_Array_foldr.Do(func() {
		cache_Data_Array_foldr = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldr")
	})
	return cache_Data_Array_foldr
}

var cache_Data_Array_foldl gopurs_runtime.Value
var once_Data_Array_foldl sync.Once
func Get_Data_Array_foldl() gopurs_runtime.Value {
	once_Data_Array_foldl.Do(func() {
		cache_Data_Array_foldl = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl")
	})
	return cache_Data_Array_foldl
}

var cache_Data_Array_transpose gopurs_runtime.Value
var once_Data_Array_transpose sync.Once
func Get_Data_Array_transpose() gopurs_runtime.Value {
	once_Data_Array_transpose.Do(func() {
		cache_Data_Array_transpose = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_transpose(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_transpose
}

var cache_Data_Array_foldRecM gopurs_runtime.Value
var once_Data_Array_foldRecM sync.Once
func Get_Data_Array_foldRecM() gopurs_runtime.Value {
	once_Data_Array_foldRecM.Do(func() {
		cache_Data_Array_foldRecM = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_foldRecM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Data_Array_foldRecM
}

var cache_Data_Array_foldMap gopurs_runtime.Value
var once_Data_Array_foldMap sync.Once
func Get_Data_Array_foldMap() gopurs_runtime.Value {
	once_Data_Array_foldMap.Do(func() {
		cache_Data_Array_foldMap = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0_box))
})
	})
	return cache_Data_Array_foldMap
}

var cache_Data_Array_foldM gopurs_runtime.Value
var once_Data_Array_foldM sync.Once
func Get_Data_Array_foldM() gopurs_runtime.Value {
	once_Data_Array_foldM.Do(func() {
		cache_Data_Array_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_Array_foldM
}

var cache_Data_Array_fold gopurs_runtime.Value
var once_Data_Array_fold sync.Once
func Get_Data_Array_fold() gopurs_runtime.Value {
	once_Data_Array_fold.Do(func() {
		cache_Data_Array_fold = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_fold(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid](dictMonoid_0_box))
})
	})
	return cache_Data_Array_fold
}

var cache_Data_Array_findMap gopurs_runtime.Value
var once_Data_Array_findMap sync.Once
func Get_Data_Array_findMap() gopurs_runtime.Value {
	once_Data_Array_findMap.Do(func() {
		cache_Data_Array_findMap = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_findMap(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_findMap
}

var cache_Data_Array_findLastIndex gopurs_runtime.Value
var once_Data_Array_findLastIndex sync.Once
func Get_Data_Array_findLastIndex() gopurs_runtime.Value {
	once_Data_Array_findLastIndex.Do(func() {
		cache_Data_Array_findLastIndex = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_findLastIndex(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_findLastIndex
}

var cache_Data_Array_insertBy gopurs_runtime.Value
var once_Data_Array_insertBy sync.Once
func Get_Data_Array_insertBy() gopurs_runtime.Value {
	once_Data_Array_insertBy.Do(func() {
		cache_Data_Array_insertBy = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_insertBy(cmp_0_box, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_insertBy
}

var cache_Data_Array_insert gopurs_runtime.Value
var once_Data_Array_insert sync.Once
func Get_Data_Array_insert() gopurs_runtime.Value {
	once_Data_Array_insert.Do(func() {
		cache_Data_Array_insert = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Array_insert
}

var cache_Data_Array_findIndex gopurs_runtime.Value
var once_Data_Array_findIndex sync.Once
func Get_Data_Array_findIndex() gopurs_runtime.Value {
	once_Data_Array_findIndex.Do(func() {
		cache_Data_Array_findIndex = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_findIndex(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_findIndex
}

var cache_Data_Array_span gopurs_runtime.Value
var once_Data_Array_span sync.Once
func Get_Data_Array_span() gopurs_runtime.Value {
	once_Data_Array_span.Do(func() {
		cache_Data_Array_span = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_span(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_span
}

var cache_Data_Array_takeWhile gopurs_runtime.Value
var once_Data_Array_takeWhile sync.Once
func Get_Data_Array_takeWhile() gopurs_runtime.Value {
	once_Data_Array_takeWhile.Do(func() {
		cache_Data_Array_takeWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_takeWhile(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_takeWhile
}

var cache_Data_Array_find gopurs_runtime.Value
var once_Data_Array_find sync.Once
func Get_Data_Array_find() gopurs_runtime.Value {
	once_Data_Array_find.Do(func() {
		cache_Data_Array_find = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_find(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_find
}

var cache_Data_Array_filter gopurs_runtime.Value
var once_Data_Array_filter sync.Once
func Get_Data_Array_filter() gopurs_runtime.Value {
	once_Data_Array_filter.Do(func() {
		cache_Data_Array_filter = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_filter(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_filter
}

var cache_Data_Array_intersectBy gopurs_runtime.Value
var once_Data_Array_intersectBy sync.Once
func Get_Data_Array_intersectBy() gopurs_runtime.Value {
	once_Data_Array_intersectBy.Do(func() {
		cache_Data_Array_intersectBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_intersectBy(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_intersectBy
}

var cache_Data_Array_intersect gopurs_runtime.Value
var once_Data_Array_intersect sync.Once
func Get_Data_Array_intersect() gopurs_runtime.Value {
	once_Data_Array_intersect.Do(func() {
		cache_Data_Array_intersect = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_intersect
}

var cache_Data_Array_elemLastIndex gopurs_runtime.Value
var once_Data_Array_elemLastIndex sync.Once
func Get_Data_Array_elemLastIndex() gopurs_runtime.Value {
	once_Data_Array_elemLastIndex.Do(func() {
		cache_Data_Array_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_Array_elemLastIndex
}

var cache_Data_Array_elemIndex gopurs_runtime.Value
var once_Data_Array_elemIndex sync.Once
func Get_Data_Array_elemIndex() gopurs_runtime.Value {
	once_Data_Array_elemIndex.Do(func() {
		cache_Data_Array_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_Array_elemIndex
}

var cache_Data_Array_notElem gopurs_runtime.Value
var once_Data_Array_notElem sync.Once
func Get_Data_Array_notElem() gopurs_runtime.Value {
	once_Data_Array_notElem.Do(func() {
		cache_Data_Array_notElem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_notElem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), a_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_notElem
}

var cache_Data_Array_elem gopurs_runtime.Value
var once_Data_Array_elem sync.Once
func Get_Data_Array_elem() gopurs_runtime.Value {
	once_Data_Array_elem.Do(func() {
		cache_Data_Array_elem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_elem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), a_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_elem
}

var cache_Data_Array_dropWhile gopurs_runtime.Value
var once_Data_Array_dropWhile sync.Once
func Get_Data_Array_dropWhile() gopurs_runtime.Value {
	once_Data_Array_dropWhile.Do(func() {
		cache_Data_Array_dropWhile = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_dropWhile(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_dropWhile
}

var cache_Data_Array_dropEnd gopurs_runtime.Value
var once_Data_Array_dropEnd sync.Once
func Get_Data_Array_dropEnd() gopurs_runtime.Value {
	once_Data_Array_dropEnd.Do(func() {
		cache_Data_Array_dropEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_dropEnd(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_dropEnd
}

var cache_Data_Array_drop gopurs_runtime.Value
var once_Data_Array_drop sync.Once
func Get_Data_Array_drop() gopurs_runtime.Value {
	once_Data_Array_drop.Do(func() {
		cache_Data_Array_drop = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_drop(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_drop
}

var cache_Data_Array_takeEnd gopurs_runtime.Value
var once_Data_Array_takeEnd sync.Once
func Get_Data_Array_takeEnd() gopurs_runtime.Value {
	once_Data_Array_takeEnd.Do(func() {
		cache_Data_Array_takeEnd = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_takeEnd(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_takeEnd
}

var cache_Data_Array_deleteAt gopurs_runtime.Value
var once_Data_Array_deleteAt sync.Once
func Get_Data_Array_deleteAt() gopurs_runtime.Value {
	once_Data_Array_deleteAt.Do(func() {
		cache_Data_Array_deleteAt = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_deleteAt(__local_var_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_deleteAt
}

var cache_Data_Array_deleteBy gopurs_runtime.Value
var once_Data_Array_deleteBy sync.Once
func Get_Data_Array_deleteBy() gopurs_runtime.Value {
	once_Data_Array_deleteBy.Do(func() {
		cache_Data_Array_deleteBy = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_deleteBy(v_0_box, v1_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v2_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_deleteBy
}

var cache_Data_Array_delete gopurs_runtime.Value
var once_Data_Array_delete sync.Once
func Get_Data_Array_delete() gopurs_runtime.Value {
	once_Data_Array_delete.Do(func() {
		cache_Data_Array_delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_delete
}

var cache_Data_Array_difference gopurs_runtime.Value
var once_Data_Array_difference sync.Once
func Get_Data_Array_difference() gopurs_runtime.Value {
	once_Data_Array_difference.Do(func() {
		cache_Data_Array_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_difference
}

var cache_Data_Array_cons gopurs_runtime.Value
var once_Data_Array_cons sync.Once
func Get_Data_Array_cons() gopurs_runtime.Value {
	once_Data_Array_cons.Do(func() {
		cache_Data_Array_cons = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_cons(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_cons
}

var cache_Data_Array_some gopurs_runtime.Value
var once_Data_Array_some sync.Once
func Get_Data_Array_some() gopurs_runtime.Value {
	once_Data_Array_some.Do(func() {
		cache_Data_Array_some = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_Array_some
}

var cache_Data_Array_many gopurs_runtime.Value
var once_Data_Array_many sync.Once
func Get_Data_Array_many() gopurs_runtime.Value {
	once_Data_Array_many.Do(func() {
		cache_Data_Array_many = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_many(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_Array_many
}

var cache_Data_Array_concatMap gopurs_runtime.Value
var once_Data_Array_concatMap sync.Once
func Get_Data_Array_concatMap() gopurs_runtime.Value {
	once_Data_Array_concatMap.Do(func() {
		cache_Data_Array_concatMap = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_concatMap(b_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_concatMap
}

var cache_Data_Array_mapMaybe gopurs_runtime.Value
var once_Data_Array_mapMaybe sync.Once
func Get_Data_Array_mapMaybe() gopurs_runtime.Value {
	once_Data_Array_mapMaybe.Do(func() {
		cache_Data_Array_mapMaybe = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_mapMaybe(f_0_box)
})
	})
	return cache_Data_Array_mapMaybe
}

var cache_Data_Array_filterA gopurs_runtime.Value
var once_Data_Array_filterA sync.Once
func Get_Data_Array_filterA() gopurs_runtime.Value {
	once_Data_Array_filterA.Do(func() {
		cache_Data_Array_filterA = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_filterA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Data_Array_filterA
}

var cache_Data_Array_catMaybes gopurs_runtime.Value
var once_Data_Array_catMaybes sync.Once
func Get_Data_Array_catMaybes() gopurs_runtime.Value {
	once_Data_Array_catMaybes.Do(func() {
		cache_Data_Array_catMaybes = Call_Data_Array_mapMaybe(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Array_catMaybes
}

var cache_Data_Array_any gopurs_runtime.Value
var once_Data_Array_any sync.Once
func Get_Data_Array_any() gopurs_runtime.Value {
	once_Data_Array_any.Do(func() {
		cache_Data_Array_any = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_any(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_any
}

var cache_Data_Array_nubByEq gopurs_runtime.Value
var once_Data_Array_nubByEq sync.Once
func Get_Data_Array_nubByEq() gopurs_runtime.Value {
	once_Data_Array_nubByEq.Do(func() {
		cache_Data_Array_nubByEq = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_nubByEq(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_nubByEq
}

var cache_Data_Array_nubEq gopurs_runtime.Value
var once_Data_Array_nubEq sync.Once
func Get_Data_Array_nubEq() gopurs_runtime.Value {
	once_Data_Array_nubEq.Do(func() {
		cache_Data_Array_nubEq = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_nubEq
}

var cache_Data_Array_unionBy gopurs_runtime.Value
var once_Data_Array_unionBy sync.Once
func Get_Data_Array_unionBy() gopurs_runtime.Value {
	once_Data_Array_unionBy.Do(func() {
		cache_Data_Array_unionBy = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_unionBy(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_unionBy
}

var cache_Data_Array_union gopurs_runtime.Value
var once_Data_Array_union sync.Once
func Get_Data_Array_union() gopurs_runtime.Value {
	once_Data_Array_union.Do(func() {
		cache_Data_Array_union = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_union
}

var cache_Data_Array_alterAt gopurs_runtime.Value
var once_Data_Array_alterAt sync.Once
func Get_Data_Array_alterAt() gopurs_runtime.Value {
	once_Data_Array_alterAt.Do(func() {
		cache_Data_Array_alterAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_alterAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_alterAt
}

var cache_Data_Array_all gopurs_runtime.Value
var once_Data_Array_all sync.Once
func Get_Data_Array_all() gopurs_runtime.Value {
	once_Data_Array_all.Do(func() {
		cache_Data_Array_all = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_all(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_all
}

var cache_Data_Array_all__3571149479 gopurs_runtime.Value
var once_Data_Array_all__3571149479 sync.Once
func Get_Data_Array_all__3571149479() gopurs_runtime.Value {
	once_Data_Array_all__3571149479.Do(func() {
		cache_Data_Array_all__3571149479 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_all__3571149479(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_all__3571149479
}

var cache_Data_Array_alterAt__2287604653 gopurs_runtime.Value
var once_Data_Array_alterAt__2287604653 sync.Once
func Get_Data_Array_alterAt__2287604653() gopurs_runtime.Value {
	once_Data_Array_alterAt__2287604653.Do(func() {
		cache_Data_Array_alterAt__2287604653 = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_alterAt__2287604653(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_alterAt__2287604653
}

var cache_Data_Array_any__3571149479 gopurs_runtime.Value
var once_Data_Array_any__3571149479 sync.Once
func Get_Data_Array_any__3571149479() gopurs_runtime.Value {
	once_Data_Array_any__3571149479.Do(func() {
		cache_Data_Array_any__3571149479 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_any__3571149479(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_any__3571149479
}

var cache_Data_Array_catMaybes__3882324634 gopurs_runtime.Value
var once_Data_Array_catMaybes__3882324634 sync.Once
func Get_Data_Array_catMaybes__3882324634() gopurs_runtime.Value {
	once_Data_Array_catMaybes__3882324634.Do(func() {
		cache_Data_Array_catMaybes__3882324634 = Call_Data_Array_mapMaybe(gopurs_runtime.Func(func(x_0 gopurs_runtime.Value) gopurs_runtime.Value {
return x_0
}))
	})
	return cache_Data_Array_catMaybes__3882324634
}

var cache_Data_Array_concatMap__435921434 gopurs_runtime.Value
var once_Data_Array_concatMap__435921434 sync.Once
func Get_Data_Array_concatMap__435921434() gopurs_runtime.Value {
	once_Data_Array_concatMap__435921434.Do(func() {
		cache_Data_Array_concatMap__435921434 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_concatMap__435921434(b_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(a_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_concatMap__435921434
}

var cache_Data_Array_concatMap__2727787066 gopurs_runtime.Value
var once_Data_Array_concatMap__2727787066 sync.Once
func Get_Data_Array_concatMap__2727787066() gopurs_runtime.Value {
	once_Data_Array_concatMap__2727787066.Do(func() {
		cache_Data_Array_concatMap__2727787066 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_concatMap__2727787066(b_0_box, func() []*Constructor_Data_Maybe_Just {
					arr := *(*[]gopurs_runtime.Value)(a_1_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Maybe_Just, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](v) }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_concatMap__2727787066
}

var cache_Data_Array_concatMap__2827656794 gopurs_runtime.Value
var once_Data_Array_concatMap__2827656794 sync.Once
func Get_Data_Array_concatMap__2827656794() gopurs_runtime.Value {
	once_Data_Array_concatMap__2827656794.Do(func() {
		cache_Data_Array_concatMap__2827656794 = gopurs_runtime.Func2(func(b_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_concatMap__2827656794(b_0_box, func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(a_1_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_concatMap__2827656794
}

var cache_Data_Array_cons__3485573810 gopurs_runtime.Value
var once_Data_Array_cons__3485573810 sync.Once
func Get_Data_Array_cons__3485573810() gopurs_runtime.Value {
	once_Data_Array_cons__3485573810.Do(func() {
		cache_Data_Array_cons__3485573810 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_cons__3485573810(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_cons__3485573810
}

var cache_Data_Array_cons__4002752745 gopurs_runtime.Value
var once_Data_Array_cons__4002752745 sync.Once
func Get_Data_Array_cons__4002752745() gopurs_runtime.Value {
	once_Data_Array_cons__4002752745.Do(func() {
		cache_Data_Array_cons__4002752745 = gopurs_runtime.Func2(func(x_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_cons__4002752745(x_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_cons__4002752745
}

var cache_Data_Array_delete__525954648 gopurs_runtime.Value
var once_Data_Array_delete__525954648 sync.Once
func Get_Data_Array_delete__525954648() gopurs_runtime.Value {
	once_Data_Array_delete__525954648.Do(func() {
		cache_Data_Array_delete__525954648 = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_delete__525954648(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box))
})
	})
	return cache_Data_Array_delete__525954648
}

var cache_Data_Array_deleteAt__454851725 gopurs_runtime.Value
var once_Data_Array_deleteAt__454851725 sync.Once
func Get_Data_Array_deleteAt__454851725() gopurs_runtime.Value {
	once_Data_Array_deleteAt__454851725.Do(func() {
		cache_Data_Array_deleteAt__454851725 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_deleteAt__454851725(__local_var_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_deleteAt__454851725
}

var cache_Data_Array_deleteBy__519303411 gopurs_runtime.Value
var once_Data_Array_deleteBy__519303411 sync.Once
func Get_Data_Array_deleteBy__519303411() gopurs_runtime.Value {
	once_Data_Array_deleteBy__519303411.Do(func() {
		cache_Data_Array_deleteBy__519303411 = gopurs_runtime.Func3(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value, v2_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_deleteBy__519303411(v_0_box, v1_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v2_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_deleteBy__519303411
}

var cache_Data_Array_drop__1426757676 gopurs_runtime.Value
var once_Data_Array_drop__1426757676 sync.Once
func Get_Data_Array_drop__1426757676() gopurs_runtime.Value {
	once_Data_Array_drop__1426757676.Do(func() {
		cache_Data_Array_drop__1426757676 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_drop__1426757676(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_drop__1426757676
}

var cache_Data_Array_dropEnd__1426757676 gopurs_runtime.Value
var once_Data_Array_dropEnd__1426757676 sync.Once
func Get_Data_Array_dropEnd__1426757676() gopurs_runtime.Value {
	once_Data_Array_dropEnd__1426757676.Do(func() {
		cache_Data_Array_dropEnd__1426757676 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_dropEnd__1426757676(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_dropEnd__1426757676
}

var cache_Data_Array_dropWhile__377906483 gopurs_runtime.Value
var once_Data_Array_dropWhile__377906483 sync.Once
func Get_Data_Array_dropWhile__377906483() gopurs_runtime.Value {
	once_Data_Array_dropWhile__377906483.Do(func() {
		cache_Data_Array_dropWhile__377906483 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_dropWhile__377906483(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_dropWhile__377906483
}

var cache_Data_Array_elem__1106871500 gopurs_runtime.Value
var once_Data_Array_elem__1106871500 sync.Once
func Get_Data_Array_elem__1106871500() gopurs_runtime.Value {
	once_Data_Array_elem__1106871500.Do(func() {
		cache_Data_Array_elem__1106871500 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_elem__1106871500(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), a_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_elem__1106871500
}

var cache_Data_Array_elemIndex__33401498 gopurs_runtime.Value
var once_Data_Array_elemIndex__33401498 sync.Once
func Get_Data_Array_elemIndex__33401498() gopurs_runtime.Value {
	once_Data_Array_elemIndex__33401498.Do(func() {
		cache_Data_Array_elemIndex__33401498 = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_elemIndex__33401498(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_Array_elemIndex__33401498
}

var cache_Data_Array_elemLastIndex__33401498 gopurs_runtime.Value
var once_Data_Array_elemLastIndex__33401498 sync.Once
func Get_Data_Array_elemLastIndex__33401498() gopurs_runtime.Value {
	once_Data_Array_elemLastIndex__33401498.Do(func() {
		cache_Data_Array_elemLastIndex__33401498 = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_elemLastIndex__33401498(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), x_1_box)
})
	})
	return cache_Data_Array_elemLastIndex__33401498
}

var cache_Data_Array_filter__4047711382 gopurs_runtime.Value
var once_Data_Array_filter__4047711382 sync.Once
func Get_Data_Array_filter__4047711382() gopurs_runtime.Value {
	once_Data_Array_filter__4047711382.Do(func() {
		cache_Data_Array_filter__4047711382 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_filter__4047711382(__local_var_0_box, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_filter__4047711382
}

var cache_Data_Array_filter__377906483 gopurs_runtime.Value
var once_Data_Array_filter__377906483 sync.Once
func Get_Data_Array_filter__377906483() gopurs_runtime.Value {
	once_Data_Array_filter__377906483.Do(func() {
		cache_Data_Array_filter__377906483 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_filter__377906483(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_filter__377906483
}

var cache_Data_Array_filterA__2723385228 gopurs_runtime.Value
var once_Data_Array_filterA__2723385228 sync.Once
func Get_Data_Array_filterA__2723385228() gopurs_runtime.Value {
	once_Data_Array_filterA__2723385228.Do(func() {
		cache_Data_Array_filterA__2723385228 = gopurs_runtime.Func(func(dictApplicative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_filterA__2723385228(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box))
})
	})
	return cache_Data_Array_filterA__2723385228
}

var cache_Data_Array_find__2560752692 gopurs_runtime.Value
var once_Data_Array_find__2560752692 sync.Once
func Get_Data_Array_find__2560752692() gopurs_runtime.Value {
	once_Data_Array_find__2560752692.Do(func() {
		cache_Data_Array_find__2560752692 = gopurs_runtime.Func2(func(f_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_find__2560752692(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_find__2560752692
}

var cache_Data_Array_findIndex__139581937 gopurs_runtime.Value
var once_Data_Array_findIndex__139581937 sync.Once
func Get_Data_Array_findIndex__139581937() gopurs_runtime.Value {
	once_Data_Array_findIndex__139581937.Do(func() {
		cache_Data_Array_findIndex__139581937 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_findIndex__139581937(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_findIndex__139581937
}

var cache_Data_Array_findLastIndex__139581937 gopurs_runtime.Value
var once_Data_Array_findLastIndex__139581937 sync.Once
func Get_Data_Array_findLastIndex__139581937() gopurs_runtime.Value {
	once_Data_Array_findLastIndex__139581937.Do(func() {
		cache_Data_Array_findLastIndex__139581937 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_findLastIndex__139581937(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_findLastIndex__139581937
}

var cache_Data_Array_findMap__3943035258 gopurs_runtime.Value
var once_Data_Array_findMap__3943035258 sync.Once
func Get_Data_Array_findMap__3943035258() gopurs_runtime.Value {
	once_Data_Array_findMap__3943035258.Do(func() {
		cache_Data_Array_findMap__3943035258 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_findMap__3943035258(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_findMap__3943035258
}

var cache_Data_Array_foldM__2595407950 gopurs_runtime.Value
var once_Data_Array_foldM__2595407950 sync.Once
func Get_Data_Array_foldM__2595407950() gopurs_runtime.Value {
	once_Data_Array_foldM__2595407950.Do(func() {
		cache_Data_Array_foldM__2595407950 = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_foldM__2595407950(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad](dictMonad_0_box))
})
	})
	return cache_Data_Array_foldM__2595407950
}

var cache_Data_Array_foldRecM__306774880 gopurs_runtime.Value
var once_Data_Array_foldRecM__306774880 sync.Once
func Get_Data_Array_foldRecM__306774880() gopurs_runtime.Value {
	once_Data_Array_foldRecM__306774880.Do(func() {
		cache_Data_Array_foldRecM__306774880 = gopurs_runtime.Func(func(dictMonadRec_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_foldRecM__306774880(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec](dictMonadRec_0_box))
})
	})
	return cache_Data_Array_foldRecM__306774880
}

var cache_Data_Array_foldl__2208423996 gopurs_runtime.Value
var once_Data_Array_foldl__2208423996 sync.Once
func Get_Data_Array_foldl__2208423996() gopurs_runtime.Value {
	once_Data_Array_foldl__2208423996.Do(func() {
		cache_Data_Array_foldl__2208423996 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl")
	})
	return cache_Data_Array_foldl__2208423996
}

var cache_Data_Array_foldl__849397914 gopurs_runtime.Value
var once_Data_Array_foldl__849397914 sync.Once
func Get_Data_Array_foldl__849397914() gopurs_runtime.Value {
	once_Data_Array_foldl__849397914.Do(func() {
		cache_Data_Array_foldl__849397914 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl")
	})
	return cache_Data_Array_foldl__849397914
}

var cache_Data_Array_foldl__1469296346 gopurs_runtime.Value
var once_Data_Array_foldl__1469296346 sync.Once
func Get_Data_Array_foldl__1469296346() gopurs_runtime.Value {
	once_Data_Array_foldl__1469296346.Do(func() {
		cache_Data_Array_foldl__1469296346 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl")
	})
	return cache_Data_Array_foldl__1469296346
}

var cache_Data_Array_foldl__1522453594 gopurs_runtime.Value
var once_Data_Array_foldl__1522453594 sync.Once
func Get_Data_Array_foldl__1522453594() gopurs_runtime.Value {
	once_Data_Array_foldl__1522453594.Do(func() {
		cache_Data_Array_foldl__1522453594 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl")
	})
	return cache_Data_Array_foldl__1522453594
}

var cache_Data_Array_foldr__1469296346 gopurs_runtime.Value
var once_Data_Array_foldr__1469296346 sync.Once
func Get_Data_Array_foldr__1469296346() gopurs_runtime.Value {
	once_Data_Array_foldr__1469296346.Do(func() {
		cache_Data_Array_foldr__1469296346 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldr")
	})
	return cache_Data_Array_foldr__1469296346
}

var cache_Data_Array_foldr__1916116122 gopurs_runtime.Value
var once_Data_Array_foldr__1916116122 sync.Once
func Get_Data_Array_foldr__1916116122() gopurs_runtime.Value {
	once_Data_Array_foldr__1916116122.Do(func() {
		cache_Data_Array_foldr__1916116122 = gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldr")
	})
	return cache_Data_Array_foldr__1916116122
}

var cache_Data_Array_groupAllBy__1923945894 gopurs_runtime.Value
var once_Data_Array_groupAllBy__1923945894 sync.Once
func Get_Data_Array_groupAllBy__1923945894() gopurs_runtime.Value {
	once_Data_Array_groupAllBy__1923945894.Do(func() {
		cache_Data_Array_groupAllBy__1923945894 = gopurs_runtime.Func(func(cmp_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_groupAllBy__1923945894(cmp_0_box)
})
	})
	return cache_Data_Array_groupAllBy__1923945894
}

var cache_Data_Array_groupBy__693635452 gopurs_runtime.Value
var once_Data_Array_groupBy__693635452 sync.Once
func Get_Data_Array_groupBy__693635452() gopurs_runtime.Value {
	once_Data_Array_groupBy__693635452.Do(func() {
		cache_Data_Array_groupBy__693635452 = gopurs_runtime.Func2(func(op_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_groupBy__693635452(op_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_groupBy__693635452
}

var cache_Data_Array_head__2042355260 gopurs_runtime.Value
var once_Data_Array_head__2042355260 sync.Once
func Get_Data_Array_head__2042355260() gopurs_runtime.Value {
	once_Data_Array_head__2042355260.Do(func() {
		cache_Data_Array_head__2042355260 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_head__2042355260(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_head__2042355260
}

var cache_Data_Array_head__2056156327 gopurs_runtime.Value
var once_Data_Array_head__2056156327 sync.Once
func Get_Data_Array_head__2056156327() gopurs_runtime.Value {
	once_Data_Array_head__2056156327.Do(func() {
		cache_Data_Array_head__2056156327 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_head__2056156327(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_head__2056156327
}

var cache_Data_Array_head__1956412839 gopurs_runtime.Value
var once_Data_Array_head__1956412839 sync.Once
func Get_Data_Array_head__1956412839() gopurs_runtime.Value {
	once_Data_Array_head__1956412839.Do(func() {
		cache_Data_Array_head__1956412839 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_head__1956412839(func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_head__1956412839
}

var cache_Data_Array_index__4267297680 gopurs_runtime.Value
var once_Data_Array_index__4267297680 sync.Once
func Get_Data_Array_index__4267297680() gopurs_runtime.Value {
	once_Data_Array_index__4267297680.Do(func() {
		cache_Data_Array_index__4267297680 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_index__4267297680(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_1_box.IntVal))}
})
	})
	return cache_Data_Array_index__4267297680
}

var cache_Data_Array_index__2196477387 gopurs_runtime.Value
var once_Data_Array_index__2196477387 sync.Once
func Get_Data_Array_index__2196477387() gopurs_runtime.Value {
	once_Data_Array_index__2196477387.Do(func() {
		cache_Data_Array_index__2196477387 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_index__2196477387(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_1_box.IntVal))}
})
	})
	return cache_Data_Array_index__2196477387
}

var cache_Data_Array_index__2659850123 gopurs_runtime.Value
var once_Data_Array_index__2659850123 sync.Once
func Get_Data_Array_index__2659850123() gopurs_runtime.Value {
	once_Data_Array_index__2659850123.Do(func() {
		cache_Data_Array_index__2659850123 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_index__2659850123(func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}(), __local_var_1_box.IntVal))}
})
	})
	return cache_Data_Array_index__2659850123
}

var cache_Data_Array_index__3264285643 gopurs_runtime.Value
var once_Data_Array_index__3264285643 sync.Once
func Get_Data_Array_index__3264285643() gopurs_runtime.Value {
	once_Data_Array_index__3264285643.Do(func() {
		cache_Data_Array_index__3264285643 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_index__3264285643(func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}(), __local_var_1_box.IntVal))}
})
	})
	return cache_Data_Array_index__3264285643
}

var cache_Data_Array_init__7061562 gopurs_runtime.Value
var once_Data_Array_init__7061562 sync.Once
func Get_Data_Array_init__7061562() gopurs_runtime.Value {
	once_Data_Array_init__7061562.Do(func() {
		cache_Data_Array_init__7061562 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_init__7061562(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_init__7061562
}

var cache_Data_Array_init__976795489 gopurs_runtime.Value
var once_Data_Array_init__976795489 sync.Once
func Get_Data_Array_init__976795489() gopurs_runtime.Value {
	once_Data_Array_init__976795489.Do(func() {
		cache_Data_Array_init__976795489 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_init__976795489(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_init__976795489
}

var cache_Data_Array_insert__1514035128 gopurs_runtime.Value
var once_Data_Array_insert__1514035128 sync.Once
func Get_Data_Array_insert__1514035128() gopurs_runtime.Value {
	once_Data_Array_insert__1514035128.Do(func() {
		cache_Data_Array_insert__1514035128 = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_insert__1514035128(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box))
})
	})
	return cache_Data_Array_insert__1514035128
}

var cache_Data_Array_insertAt__388410084 gopurs_runtime.Value
var once_Data_Array_insertAt__388410084 sync.Once
func Get_Data_Array_insertAt__388410084() gopurs_runtime.Value {
	once_Data_Array_insertAt__388410084.Do(func() {
		cache_Data_Array_insertAt__388410084 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_insertAt__388410084(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_insertAt__388410084
}

var cache_Data_Array_insertBy__1563432905 gopurs_runtime.Value
var once_Data_Array_insertBy__1563432905 sync.Once
func Get_Data_Array_insertBy__1563432905() gopurs_runtime.Value {
	once_Data_Array_insertBy__1563432905.Do(func() {
		cache_Data_Array_insertBy__1563432905 = gopurs_runtime.Func3(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_insertBy__1563432905(cmp_0_box, x_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_insertBy__1563432905
}

var cache_Data_Array_intersectBy__145374773 gopurs_runtime.Value
var once_Data_Array_intersectBy__145374773 sync.Once
func Get_Data_Array_intersectBy__145374773() gopurs_runtime.Value {
	once_Data_Array_intersectBy__145374773.Do(func() {
		cache_Data_Array_intersectBy__145374773 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_intersectBy__145374773(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_intersectBy__145374773
}

var cache_Data_Array_intersperse__4002752745 gopurs_runtime.Value
var once_Data_Array_intersperse__4002752745 sync.Once
func Get_Data_Array_intersperse__4002752745() gopurs_runtime.Value {
	once_Data_Array_intersperse__4002752745.Do(func() {
		cache_Data_Array_intersperse__4002752745 = gopurs_runtime.Func2(func(a_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_intersperse__4002752745(a_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_intersperse__4002752745
}

var cache_Data_Array_last__2042355260 gopurs_runtime.Value
var once_Data_Array_last__2042355260 sync.Once
func Get_Data_Array_last__2042355260() gopurs_runtime.Value {
	once_Data_Array_last__2042355260.Do(func() {
		cache_Data_Array_last__2042355260 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_last__2042355260(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_last__2042355260
}

var cache_Data_Array_last__2345912124 gopurs_runtime.Value
var once_Data_Array_last__2345912124 sync.Once
func Get_Data_Array_last__2345912124() gopurs_runtime.Value {
	once_Data_Array_last__2345912124.Do(func() {
		cache_Data_Array_last__2345912124 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_last__2345912124(func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_last__2345912124
}

var cache_Data_Array_last__2056156327 gopurs_runtime.Value
var once_Data_Array_last__2056156327 sync.Once
func Get_Data_Array_last__2056156327() gopurs_runtime.Value {
	once_Data_Array_last__2056156327.Do(func() {
		cache_Data_Array_last__2056156327 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_last__2056156327(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_last__2056156327
}

var cache_Data_Array_many__1839052385 gopurs_runtime.Value
var once_Data_Array_many__1839052385 sync.Once
func Get_Data_Array_many__1839052385() gopurs_runtime.Value {
	once_Data_Array_many__1839052385.Do(func() {
		cache_Data_Array_many__1839052385 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_many__1839052385(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_Array_many__1839052385
}

var cache_Data_Array_mapMaybe__1271145181 gopurs_runtime.Value
var once_Data_Array_mapMaybe__1271145181 sync.Once
func Get_Data_Array_mapMaybe__1271145181() gopurs_runtime.Value {
	once_Data_Array_mapMaybe__1271145181.Do(func() {
		cache_Data_Array_mapMaybe__1271145181 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_mapMaybe__1271145181(f_0_box)
})
	})
	return cache_Data_Array_mapMaybe__1271145181
}

var cache_Data_Array_mapMaybe__3137412285 gopurs_runtime.Value
var once_Data_Array_mapMaybe__3137412285 sync.Once
func Get_Data_Array_mapMaybe__3137412285() gopurs_runtime.Value {
	once_Data_Array_mapMaybe__3137412285.Do(func() {
		cache_Data_Array_mapMaybe__3137412285 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_mapMaybe__3137412285(f_0_box)
})
	})
	return cache_Data_Array_mapMaybe__3137412285
}

var cache_Data_Array_mapMaybe__2261006141 gopurs_runtime.Value
var once_Data_Array_mapMaybe__2261006141 sync.Once
func Get_Data_Array_mapMaybe__2261006141() gopurs_runtime.Value {
	once_Data_Array_mapMaybe__2261006141.Do(func() {
		cache_Data_Array_mapMaybe__2261006141 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_mapMaybe__2261006141(f_0_box)
})
	})
	return cache_Data_Array_mapMaybe__2261006141
}

var cache_Data_Array_mapWithIndex__3745622640 gopurs_runtime.Value
var once_Data_Array_mapWithIndex__3745622640 sync.Once
func Get_Data_Array_mapWithIndex__3745622640() gopurs_runtime.Value {
	once_Data_Array_mapWithIndex__3745622640.Do(func() {
		cache_Data_Array_mapWithIndex__3745622640 = gopurs_runtime.RecordGet(Get_Data_FunctorWithIndex_functorWithIndexArray(), "mapWithIndex")
	})
	return cache_Data_Array_mapWithIndex__3745622640
}

var cache_Data_Array_mapWithIndex__1705728720 gopurs_runtime.Value
var once_Data_Array_mapWithIndex__1705728720 sync.Once
func Get_Data_Array_mapWithIndex__1705728720() gopurs_runtime.Value {
	once_Data_Array_mapWithIndex__1705728720.Do(func() {
		cache_Data_Array_mapWithIndex__1705728720 = gopurs_runtime.RecordGet(Get_Data_FunctorWithIndex_functorWithIndexArray(), "mapWithIndex")
	})
	return cache_Data_Array_mapWithIndex__1705728720
}

var cache_Data_Array_modifyAt__3384125836 gopurs_runtime.Value
var once_Data_Array_modifyAt__3384125836 sync.Once
func Get_Data_Array_modifyAt__3384125836() gopurs_runtime.Value {
	once_Data_Array_modifyAt__3384125836.Do(func() {
		cache_Data_Array_modifyAt__3384125836 = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_modifyAt__3384125836(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_modifyAt__3384125836
}

var cache_Data_Array_modifyAtIndices__536948024 gopurs_runtime.Value
var once_Data_Array_modifyAtIndices__536948024 sync.Once
func Get_Data_Array_modifyAtIndices__536948024() gopurs_runtime.Value {
	once_Data_Array_modifyAtIndices__536948024.Do(func() {
		cache_Data_Array_modifyAtIndices__536948024 = gopurs_runtime.Func4(func(dictFoldable_0_box gopurs_runtime.Value, is_1_box gopurs_runtime.Value, f_2_box gopurs_runtime.Value, xs_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_modifyAtIndices__536948024(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), is_1_box, f_2_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_3_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_modifyAtIndices__536948024
}

var cache_Data_Array_notElem__1106871500 gopurs_runtime.Value
var once_Data_Array_notElem__1106871500 sync.Once
func Get_Data_Array_notElem__1106871500() gopurs_runtime.Value {
	once_Data_Array_notElem__1106871500.Do(func() {
		cache_Data_Array_notElem__1106871500 = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_notElem__1106871500(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq](dictEq_0_box), a_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_notElem__1106871500
}

var cache_Data_Array_nubBy__3347533344 gopurs_runtime.Value
var once_Data_Array_nubBy__3347533344 sync.Once
func Get_Data_Array_nubBy__3347533344() gopurs_runtime.Value {
	once_Data_Array_nubBy__3347533344.Do(func() {
		cache_Data_Array_nubBy__3347533344 = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_nubBy__3347533344(comp_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_nubBy__3347533344
}

var cache_Data_Array_nubByEq__3443670074 gopurs_runtime.Value
var once_Data_Array_nubByEq__3443670074 sync.Once
func Get_Data_Array_nubByEq__3443670074() gopurs_runtime.Value {
	once_Data_Array_nubByEq__3443670074.Do(func() {
		cache_Data_Array_nubByEq__3443670074 = gopurs_runtime.Func2(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_nubByEq__3443670074(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_nubByEq__3443670074
}

var cache_Data_Array_null__1123412116 gopurs_runtime.Value
var once_Data_Array_null__1123412116 sync.Once
func Get_Data_Array_null__1123412116() gopurs_runtime.Value {
	once_Data_Array_null__1123412116.Do(func() {
		cache_Data_Array_null__1123412116 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_null__1123412116(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_null__1123412116
}

var cache_Data_Array_partition__3230398268 gopurs_runtime.Value
var once_Data_Array_partition__3230398268 sync.Once
func Get_Data_Array_partition__3230398268() gopurs_runtime.Value {
	once_Data_Array_partition__3230398268.Do(func() {
		cache_Data_Array_partition__3230398268 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_partition__3230398268(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_partition__3230398268
}

var cache_Data_Array_replicate__1064418410 gopurs_runtime.Value
var once_Data_Array_replicate__1064418410 sync.Once
func Get_Data_Array_replicate__1064418410() gopurs_runtime.Value {
	once_Data_Array_replicate__1064418410.Do(func() {
		cache_Data_Array_replicate__1064418410 = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_replicate__1064418410(__local_var_0_box.IntVal, __local_var_1_box))
})
	})
	return cache_Data_Array_replicate__1064418410
}

var cache_Data_Array_scanl__3156262044 gopurs_runtime.Value
var once_Data_Array_scanl__3156262044 sync.Once
func Get_Data_Array_scanl__3156262044() gopurs_runtime.Value {
	once_Data_Array_scanl__3156262044.Do(func() {
		cache_Data_Array_scanl__3156262044 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_scanl__3156262044(__local_var_0_box, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_scanl__3156262044
}

var cache_Data_Array_scanr__3156262044 gopurs_runtime.Value
var once_Data_Array_scanr__3156262044 sync.Once
func Get_Data_Array_scanr__3156262044() gopurs_runtime.Value {
	once_Data_Array_scanr__3156262044.Do(func() {
		cache_Data_Array_scanr__3156262044 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_scanr__3156262044(__local_var_0_box, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_scanr__3156262044
}

var cache_Data_Array_singleton__193199869 gopurs_runtime.Value
var once_Data_Array_singleton__193199869 sync.Once
func Get_Data_Array_singleton__193199869() gopurs_runtime.Value {
	once_Data_Array_singleton__193199869.Do(func() {
		cache_Data_Array_singleton__193199869 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_singleton__193199869(a_0_box))
})
	})
	return cache_Data_Array_singleton__193199869
}

var cache_Data_Array_singleton__2286220742 gopurs_runtime.Value
var once_Data_Array_singleton__2286220742 sync.Once
func Get_Data_Array_singleton__2286220742() gopurs_runtime.Value {
	once_Data_Array_singleton__2286220742.Do(func() {
		cache_Data_Array_singleton__2286220742 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_singleton__2286220742(a_0_box))
})
	})
	return cache_Data_Array_singleton__2286220742
}

var cache_Data_Array_singleton__2535196422 gopurs_runtime.Value
var once_Data_Array_singleton__2535196422 sync.Once
func Get_Data_Array_singleton__2535196422() gopurs_runtime.Value {
	once_Data_Array_singleton__2535196422.Do(func() {
		cache_Data_Array_singleton__2535196422 = gopurs_runtime.Func(func(a_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_singleton__2535196422(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](a_0_box))
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_singleton__2535196422
}

var cache_Data_Array_slice__3011328576 gopurs_runtime.Value
var once_Data_Array_slice__3011328576 sync.Once
func Get_Data_Array_slice__3011328576() gopurs_runtime.Value {
	once_Data_Array_slice__3011328576.Do(func() {
		cache_Data_Array_slice__3011328576 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_slice__3011328576(__local_var_0_box.IntVal, __local_var_1_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_slice__3011328576
}

var cache_Data_Array_snoc__3419689714 gopurs_runtime.Value
var once_Data_Array_snoc__3419689714 sync.Once
func Get_Data_Array_snoc__3419689714() gopurs_runtime.Value {
	once_Data_Array_snoc__3419689714.Do(func() {
		cache_Data_Array_snoc__3419689714 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_snoc__3419689714(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_Data_Array_snoc__3419689714
}

var cache_Data_Array_snoc__3911647657 gopurs_runtime.Value
var once_Data_Array_snoc__3911647657 sync.Once
func Get_Data_Array_snoc__3911647657() gopurs_runtime.Value {
	once_Data_Array_snoc__3911647657.Do(func() {
		cache_Data_Array_snoc__3911647657 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_snoc__3911647657(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), x_1_box))
})
	})
	return cache_Data_Array_snoc__3911647657
}

var cache_Data_Array_snoc__1505998191 gopurs_runtime.Value
var once_Data_Array_snoc__1505998191 sync.Once
func Get_Data_Array_snoc__1505998191() gopurs_runtime.Value {
	once_Data_Array_snoc__1505998191.Do(func() {
		cache_Data_Array_snoc__1505998191 = gopurs_runtime.Func2(func(xs_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_snoc__1505998191(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_snoc__1505998191
}

var cache_Data_Array_some__1839052385 gopurs_runtime.Value
var once_Data_Array_some__1839052385 sync.Once
func Get_Data_Array_some__1839052385() gopurs_runtime.Value {
	once_Data_Array_some__1839052385.Do(func() {
		cache_Data_Array_some__1839052385 = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_some__1839052385(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative](dictAlternative_0_box))
})
	})
	return cache_Data_Array_some__1839052385
}

var cache_Data_Array_sortBy__3347533344 gopurs_runtime.Value
var once_Data_Array_sortBy__3347533344 sync.Once
func Get_Data_Array_sortBy__3347533344() gopurs_runtime.Value {
	once_Data_Array_sortBy__3347533344.Do(func() {
		cache_Data_Array_sortBy__3347533344 = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_sortBy__3347533344(comp_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_sortBy__3347533344
}

var cache_Data_Array_sortBy__1848798496 gopurs_runtime.Value
var once_Data_Array_sortBy__1848798496 sync.Once
func Get_Data_Array_sortBy__1848798496() gopurs_runtime.Value {
	once_Data_Array_sortBy__1848798496.Do(func() {
		cache_Data_Array_sortBy__1848798496 = gopurs_runtime.Func2(func(comp_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_sortBy__1848798496(comp_0_box, func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_sortBy__1848798496
}

var cache_Data_Array_sortWith__478414925 gopurs_runtime.Value
var once_Data_Array_sortWith__478414925 sync.Once
func Get_Data_Array_sortWith__478414925() gopurs_runtime.Value {
	once_Data_Array_sortWith__478414925.Do(func() {
		cache_Data_Array_sortWith__478414925 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_sortWith__478414925(f_0_box)
})
	})
	return cache_Data_Array_sortWith__478414925
}

var cache_Data_Array_sortWith__1917042304 gopurs_runtime.Value
var once_Data_Array_sortWith__1917042304 sync.Once
func Get_Data_Array_sortWith__1917042304() gopurs_runtime.Value {
	once_Data_Array_sortWith__1917042304.Do(func() {
		cache_Data_Array_sortWith__1917042304 = gopurs_runtime.Func(func(f_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_sortWith__1917042304(f_0_box)
})
	})
	return cache_Data_Array_sortWith__1917042304
}

var cache_Data_Array_sortWith__3115909389 gopurs_runtime.Value
var once_Data_Array_sortWith__3115909389 sync.Once
func Get_Data_Array_sortWith__3115909389() gopurs_runtime.Value {
	once_Data_Array_sortWith__3115909389.Do(func() {
		cache_Data_Array_sortWith__3115909389 = gopurs_runtime.Func2(func(dictOrd_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_sortWith__3115909389(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Array_sortWith__3115909389
}

var cache_Data_Array_span__174751768 gopurs_runtime.Value
var once_Data_Array_span__174751768 sync.Once
func Get_Data_Array_span__174751768() gopurs_runtime.Value {
	once_Data_Array_span__174751768.Do(func() {
		cache_Data_Array_span__174751768 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_span__174751768(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_span__174751768
}

var cache_Data_Array_splitAt__3301820464 gopurs_runtime.Value
var once_Data_Array_splitAt__3301820464 sync.Once
func Get_Data_Array_splitAt__3301820464() gopurs_runtime.Value {
	once_Data_Array_splitAt__3301820464.Do(func() {
		cache_Data_Array_splitAt__3301820464 = gopurs_runtime.Func2(func(v_0_box gopurs_runtime.Value, v1_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_splitAt__3301820464(v_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v1_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_splitAt__3301820464
}

var cache_Data_Array_tail__7061562 gopurs_runtime.Value
var once_Data_Array_tail__7061562 sync.Once
func Get_Data_Array_tail__7061562() gopurs_runtime.Value {
	once_Data_Array_tail__7061562.Do(func() {
		cache_Data_Array_tail__7061562 = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_tail__7061562(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_tail__7061562
}

var cache_Data_Array_take__1426757676 gopurs_runtime.Value
var once_Data_Array_take__1426757676 sync.Once
func Get_Data_Array_take__1426757676() gopurs_runtime.Value {
	once_Data_Array_take__1426757676.Do(func() {
		cache_Data_Array_take__1426757676 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_take__1426757676(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_take__1426757676
}

var cache_Data_Array_takeEnd__1426757676 gopurs_runtime.Value
var once_Data_Array_takeEnd__1426757676 sync.Once
func Get_Data_Array_takeEnd__1426757676() gopurs_runtime.Value {
	once_Data_Array_takeEnd__1426757676.Do(func() {
		cache_Data_Array_takeEnd__1426757676 = gopurs_runtime.Func2(func(n_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_takeEnd__1426757676(n_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_takeEnd__1426757676
}

var cache_Data_Array_takeWhile__377906483 gopurs_runtime.Value
var once_Data_Array_takeWhile__377906483 sync.Once
func Get_Data_Array_takeWhile__377906483() gopurs_runtime.Value {
	once_Data_Array_takeWhile__377906483.Do(func() {
		cache_Data_Array_takeWhile__377906483 = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_takeWhile__377906483(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_takeWhile__377906483
}

var cache_Data_Array_transpose__4194748859 gopurs_runtime.Value
var once_Data_Array_transpose__4194748859 sync.Once
func Get_Data_Array_transpose__4194748859() gopurs_runtime.Value {
	once_Data_Array_transpose__4194748859.Do(func() {
		cache_Data_Array_transpose__4194748859 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_transpose__4194748859(func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_transpose__4194748859
}

var cache_Data_Array_uncons__2020799173 gopurs_runtime.Value
var once_Data_Array_uncons__2020799173 sync.Once
func Get_Data_Array_uncons__2020799173() gopurs_runtime.Value {
	once_Data_Array_uncons__2020799173.Do(func() {
		cache_Data_Array_uncons__2020799173 = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_uncons__2020799173(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_uncons__2020799173
}

var cache_Data_Array_unionBy__145374773 gopurs_runtime.Value
var once_Data_Array_unionBy__145374773 sync.Once
func Get_Data_Array_unionBy__145374773() gopurs_runtime.Value {
	once_Data_Array_unionBy__145374773.Do(func() {
		cache_Data_Array_unionBy__145374773 = gopurs_runtime.Func3(func(eq_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value, ys_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_unionBy__145374773(eq_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_unionBy__145374773
}

var cache_Data_Array_unsafeIndex__2808089623 gopurs_runtime.Value
var once_Data_Array_unsafeIndex__2808089623 sync.Once
func Get_Data_Array_unsafeIndex__2808089623() gopurs_runtime.Value {
	once_Data_Array_unsafeIndex__2808089623.Do(func() {
		cache_Data_Array_unsafeIndex__2808089623 = gopurs_runtime.Func3(func(_dollar__unused_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_unsafeIndex__2808089623(_dollar__unused_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_2_box.IntVal)
})
	})
	return cache_Data_Array_unsafeIndex__2808089623
}

var cache_Data_Array_unsnoc__2531125997 gopurs_runtime.Value
var once_Data_Array_unsnoc__2531125997 sync.Once
func Get_Data_Array_unsnoc__2531125997() gopurs_runtime.Value {
	once_Data_Array_unsnoc__2531125997.Do(func() {
		cache_Data_Array_unsnoc__2531125997 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_unsnoc__2531125997(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_unsnoc__2531125997
}

var cache_Data_Array_unzip__1480671261 gopurs_runtime.Value
var once_Data_Array_unzip__1480671261 sync.Once
func Get_Data_Array_unzip__1480671261() gopurs_runtime.Value {
	once_Data_Array_unzip__1480671261.Do(func() {
		cache_Data_Array_unzip__1480671261 = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Call_Data_Array_unzip__1480671261(func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_unzip__1480671261
}

var cache_Data_Array_updateAt__388410084 gopurs_runtime.Value
var once_Data_Array_updateAt__388410084 sync.Once
func Get_Data_Array_updateAt__388410084() gopurs_runtime.Value {
	once_Data_Array_updateAt__388410084.Do(func() {
		cache_Data_Array_updateAt__388410084 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Call_Data_Array_updateAt__388410084(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))}
})
	})
	return cache_Data_Array_updateAt__388410084
}

var cache_Data_Array_updateAtIndices__889079281 gopurs_runtime.Value
var once_Data_Array_updateAtIndices__889079281 sync.Once
func Get_Data_Array_updateAtIndices__889079281() gopurs_runtime.Value {
	once_Data_Array_updateAtIndices__889079281.Do(func() {
		cache_Data_Array_updateAtIndices__889079281 = gopurs_runtime.Func3(func(dictFoldable_0_box gopurs_runtime.Value, us_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_updateAtIndices__889079281(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable](dictFoldable_0_box), us_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_updateAtIndices__889079281
}

var cache_Data_Array_zip__1254936954 gopurs_runtime.Value
var once_Data_Array_zip__1254936954 sync.Once
func Get_Data_Array_zip__1254936954() gopurs_runtime.Value {
	once_Data_Array_zip__1254936954.Do(func() {
		cache_Data_Array_zip__1254936954 = gopurs_runtime.Apply(Get_Data_Array_zipWith(), Get_Data_Tuple_Tuple())
	})
	return cache_Data_Array_zip__1254936954
}

var cache_Data_Array_zipWith__1350747206 gopurs_runtime.Value
var once_Data_Array_zipWith__1350747206 sync.Once
func Get_Data_Array_zipWith__1350747206() gopurs_runtime.Value {
	once_Data_Array_zipWith__1350747206.Do(func() {
		cache_Data_Array_zipWith__1350747206 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_zipWith__1350747206(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_zipWith__1350747206
}

var cache_Data_Array_zipWith__1220584870 gopurs_runtime.Value
var once_Data_Array_zipWith__1220584870 sync.Once
func Get_Data_Array_zipWith__1220584870() gopurs_runtime.Value {
	once_Data_Array_zipWith__1220584870.Do(func() {
		cache_Data_Array_zipWith__1220584870 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_zipWith__1220584870(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()
})
	})
	return cache_Data_Array_zipWith__1220584870
}

var cache_Data_Array_zipWith__2000246342 gopurs_runtime.Value
var once_Data_Array_zipWith__2000246342 sync.Once
func Get_Data_Array_zipWith__2000246342() gopurs_runtime.Value {
	once_Data_Array_zipWith__2000246342.Do(func() {
		cache_Data_Array_zipWith__2000246342 = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_zipWith__2000246342(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_zipWith__2000246342
}

var cache_Data_Array_zipWithA__3208598546 gopurs_runtime.Value
var once_Data_Array_zipWithA__3208598546 sync.Once
func Get_Data_Array_zipWithA__3208598546() gopurs_runtime.Value {
	once_Data_Array_zipWithA__3208598546.Do(func() {
		cache_Data_Array_zipWithA__3208598546 = gopurs_runtime.Func4(func(dictApplicative_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value, ys_3_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_zipWithA__3208598546(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](dictApplicative_0_box), f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(ys_3_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})
	})
	return cache_Data_Array_zipWithA__3208598546
}

func Call_Data_Array_intercalate1(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 -> *Constructor_Data_Semigroup_Semigroup
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_0.V0), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 -> gopurs_runtime.Value
mempty_2_1 := gopurs_runtime.Box(dictMonoid_0.V1)
_ = mempty_2_1
return gopurs_runtime.Func(func(sep_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(v_5, "init").IntVal) != (0) {
__t2 = gopurs_runtime.RecordDict2("acc", "init", v1_6, gopurs_runtime.Bool(false))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.RecordDict2("acc", "init", gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.RecordGet(v_5, "acc"), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), sep_3, v1_6)), gopurs_runtime.Bool(false))
}
end_branch_2:
return __t2
})
}), gopurs_runtime.RecordDict2("acc", "init", mempty_2_1, gopurs_runtime.Bool(true)), xs_4), "acc")
})
})
}

func Call_Data_Array_fromJust(v_0_loop *Constructor_Data_Maybe_Just) *Constructor_Data_Tuple_Tuple {
var v_0 *Constructor_Data_Maybe_Just = v_0_loop
_ = v_0
var __t0 gopurs_runtime.Value
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](__t0)
}

func Call_Data_Array_zipWith(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value, ys_3_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
var ys_3 []gopurs_runtime.Value = ys_3_loop
_ = ys_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableArray(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), f_1, gopurs_runtime.Array(xs_2), gopurs_runtime.Array(ys_3)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_updateAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, us_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var us_1 gopurs_runtime.Value = us_1_loop
_ = us_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Func(func(res_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_Foldable_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Monad_ST_Internal_applicativeST()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Data_Array_ST_pokeImpl(), gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, res_3)
})
}), us_1)
}), gopurs_runtime.Array(xs_2))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_updateAt(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)))
}

func Call_Data_Array_unsafeIndex(_dollar__unused_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return __local_var_1[__local_var_2]
}

func Call_Data_Array_uncons(__local_var_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))})}
})
}), gopurs_runtime.Array(__local_var_0)))
}

func Call_Data_Array_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable, xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): len_2_0 -> int64
len_2_0 := gopurs_runtime.Int(int64(len(xs_1))).IntVal
_ = len_2_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 *Constructor_Data_Maybe_Just
{
var __t1 bool
{
if (i_3.IntVal) < (len_2_0) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return xs_1[i_3.IntVal]
})), gopurs_runtime.Int((i_3.IntVal) + (1))})}}
goto end_branch_2
} else {

}
}
{
__t2 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_2:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
}), gopurs_runtime.Int(0))
}

func Call_Data_Array_tail(__local_var_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})}
})
}), gopurs_runtime.Array(__local_var_0)))
}

func Call_Data_Array_sortBy(comp_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if (uint32(v_2.IntVal) == 380165415) {
__t0 = 1
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 902936544) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 1527465420) {
__t0 = -1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return gopurs_runtime.Int(__t0)
}), gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_sortWith(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(Get_Data_Array_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_Array_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 -> gopurs_runtime.Value
compare_1_0 := gopurs_runtime.Box(dictOrd_0.V1)
_ = compare_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_sortBy(compare_1_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}

func Call_Data_Array_snoc(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), x_1), gopurs_runtime.Array(xs_0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_slice(__local_var_0_loop int64, __local_var_1_loop int64, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(__local_var_0), gopurs_runtime.Int(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_splitAt(v_0_loop int64, v1_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 []gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (v_0) > (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(v1_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(v1_1))).IntVal), gopurs_runtime.Array(v1_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(v_0), gopurs_runtime.Array(v1_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}
end_branch_1:
return __t1
}

func Call_Data_Array_take(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t1 []gopurs_runtime.Value
{
var __t0 bool
{
if (n_0) < (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_1
} else {

}
}
{
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(n_0), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_1:
return __t1
}

func Call_Data_Array_singleton(a_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_scanr(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanrImpl(), __local_var_0, __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_scanl(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanlImpl(), __local_var_0, __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_replicate(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_replicateImpl(), gopurs_runtime.Int(__local_var_0), __local_var_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_go__range(__local_var_0_loop int64, __local_var_1_loop int64) []int64 {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_rangeImpl(), gopurs_runtime.Int(__local_var_0), gopurs_runtime.Int(__local_var_1)).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}

func Call_Data_Array_partition(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.UncurriedApp2(Get_Data_Array_partitionImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1))
}

func Call_Data_Array_null(xs_0_loop []gopurs_runtime.Value) bool {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (0)
}

func Call_Data_Array_modifyAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, is_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, xs_3_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var is_1 gopurs_runtime.Value = is_1_loop
_ = is_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var xs_3 []gopurs_runtime.Value = xs_3_loop
_ = xs_3
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Func(func(res_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_Foldable_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Monad_ST_Internal_applicativeST()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Data_Array_ST_modify(), gopurs_runtime.Int(i_5.IntVal), f_2, res_4)
}), is_1)
}), gopurs_runtime.Array(xs_3))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_intersperse(a_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
// TAST (Let): v_2_0 -> int64
v_2_0 := gopurs_runtime.Int(int64(len(arr_1))).IntVal
_ = v_2_0
var __t4 []gopurs_runtime.Value
{
var __t3 bool
{
if (v_2_0) < (2) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
if __t3 {
__t4 = arr_1
goto end_branch_4
} else {

}
}
{
__t4 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(out_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return arr_1[0]
}))
_ = __local_var_4_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_4_1, out_3)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply3(Get_Control_Monad_ST_Internal_forImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(v_2_0), gopurs_runtime.Func(func(idx_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), a_0, out_3)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_2 -> gopurs_runtime.Value
__local_var_7_2 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return arr_1[idx_5.IntVal]
}))
_ = __local_var_7_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_7_2, out_3)
}))
}))
})), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), out_3)
}))
}))
})), Get_Data_Array_ST_unsafeFreeze())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_4:
return __t4
}

func Call_Data_Array_intercalate(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid = dictMonoid_0_loop
_ = dictMonoid_0
return Call_Data_Array_intercalate1(dictMonoid_0)
}

func Call_Data_Array_insertAt(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)))
}

func Call_Data_Array_init(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (0) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1)), gopurs_runtime.Array(xs_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_Array_index(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) *Constructor_Data_Maybe_Just {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(__local_var_0), gopurs_runtime.Int(__local_var_1)))
}

func Call_Data_Array_last(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1))))
}

func Call_Data_Array_unsnoc(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (0) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1)), gopurs_runtime.Array(xs_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_applyMaybe(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), v1_2)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1)))))}))
}

func Call_Data_Array_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0)))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_3_0 == nil) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), gopurs_runtime.Apply(f_1, (__local_var_3_0).V0), gopurs_runtime.Array(xs_2))))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_Array_unzip(xs_0_loop []*Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var xs_0 []*Constructor_Data_Tuple_Tuple = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(fsts_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(snds_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(v_3.IntVal))))}
})), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(0))), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply2(Get_Data_Array_ST_Iterator_iterate(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_3))}, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_0 -> gopurs_runtime.Value
__local_var_5_0 := (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1
_ = __local_var_5_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply(Get_Data_Array_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, fsts_1)
})), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Array_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_5_0, snds_2)
}))
}))
})), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), fsts_1)
}), gopurs_runtime.Func(func(fsts_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), snds_2)
}), gopurs_runtime.Func(func(snds_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(fsts_prime_5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(snds_prime_6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
}))
}))
}))
}))
}))
}))))
}

func Call_Data_Array_head(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(0)))
}

func Call_Data_Array_nubBy(comp_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): indexedAndSorted_2_0 -> []*Constructor_Data_Tuple_Tuple
indexedAndSorted_2_0 := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_sortBy(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(comp_0, (*Constructor_Data_Tuple_Tuple)(x_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_3.UnsafePtr).V1).IntVal)), UnsafePtr: nil}
})
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_FunctorWithIndex_functorWithIndexArray(), "mapWithIndex"), Get_Data_Tuple_Tuple(), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
_ = indexedAndSorted_2_0
// TAST (Let): v_3_1 -> *Constructor_Data_Maybe_Just
v_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := indexedAndSorted_2_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Int(0)))
_ = v_3_1
var __t8 []gopurs_runtime.Value
{
if (v_3_1 == nil) {
__t8 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_8
} else {

}
}
{
if (v_3_1 != nil) {
__t8 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), Get_Data_Tuple_snd(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_Data_Array_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqIntImpl())
}), gopurs_runtime.Apply3(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})}), Get_Data_Tuple_fst()), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeThawImpl(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := []*Constructor_Data_Tuple_Tuple{gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((v_3_1).V0)}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}())
}), gopurs_runtime.Func(func(result_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_foreach(), func() gopurs_runtime.Value {
					arr := indexedAndSorted_2_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_4 -> gopurs_runtime.Value
var __local_var_9_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal) - (1)))))}
var __t5 gopurs_runtime.Value
{
if (__local_var_9_4.Type == 9 && __local_var_9_4.IntVal == 930809136 && __local_var_9_4.UnsafePtr != nil) {
__t5 = (*Constructor_Data_Maybe_Just)(__local_var_9_4.UnsafePtr).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})
}))
_ = __local_var_7_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(__local_var_7_3, x_8).UnsafePtr).V1
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_4)
})), gopurs_runtime.Func(func(lst_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := gopurs_runtime.Apply(Get_Data_Array_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v1_5))}, result_4)
}))
_ = __local_var_8_6
var __t7 gopurs_runtime.Value
{
if ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Ordering_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(comp_0, lst_7, __local_var_6_2).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}).IntVal) != (0)) != (true) {
__t7 = __local_var_8_6
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), Get_Data_Unit_unit())
}
end_branch_7:
return __t7
}))
})), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_4)
})
}))
}))).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_8
} else {

}
}
{
__t8 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_8:
return __t8
}

func Call_Data_Array_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Array_nubBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_Array_groupBy(op_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) [][]gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_1), gopurs_runtime.Int(v_3.IntVal))))}
})), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(0))), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply2(Get_Data_Array_ST_Iterator_iterate(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_3))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(sub_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), x_4, sub_5)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply3(Get_Data_Array_ST_Iterator_pushWhile(), gopurs_runtime.Apply(op_0, x_4), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_3))}, sub_5), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), sub_5)
}), gopurs_runtime.Func(func(grp_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(grp_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), result_2)
})
}))
}))
}))
})))
})), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2)
})
}))
}))
}))).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
}

func Call_Data_Array_groupAllBy(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> uint32
__local_var_3_1 := uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_2).IntVal)
_ = __local_var_3_1
var __t2 bool
{
if (__local_var_3_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1 == 380165415) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1 == 902936544) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
})
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Array(Call_Data_Array_sortBy(cmp_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
})
}

func Call_Data_Array_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return Call_Data_Array_groupAllBy(gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_Array_group(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eq_1_0 -> gopurs_runtime.Value
eq_1_0 := gopurs_runtime.Box(dictEq_0.V0)
_ = eq_1_0
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_groupBy(eq_1_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
})
}

func Call_Data_Array_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Box(dictFoldable_0.V2)
_ = __local_var_1_0
return gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_fromFoldableImpl(), __local_var_1_0, __local_var_2)
})
}

func Call_Data_Array_transpose(xs_0_loop [][]gopurs_runtime.Value) [][]gopurs_runtime.Value {
var xs_0 [][]gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_0 gopurs_runtime.Value
go__go_1_0_0 = gopurs_runtime.Func(func(idx_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(allArrays_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var idx_2_loop int64 = idx_2_loop_val.IntVal
var allArrays_3_loop [][]gopurs_runtime.Value = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(allArrays_3_loop_val.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var idx_2 int64 = idx_2_loop
_ = idx_2
var allArrays_3 [][]gopurs_runtime.Value = allArrays_3_loop
_ = allArrays_3
// TAST (Let): v_4_1 -> *Constructor_Data_Maybe_Just
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(nextArr_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> *Constructor_Data_Maybe_Just
__local_var_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(nextArr_5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Int(idx_2)))
_ = __local_var_6_2
var __t6 *Constructor_Data_Maybe_Just
{
if (__local_var_6_2 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](acc_4)
goto end_branch_6
} else {

}
}
{
if (__local_var_6_2 != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](acc_4)
if (__t_tag_3 == nil) {
__t5 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(__local_var_6_2).V0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](acc_4)
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), (__local_var_6_2).V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((*Constructor_Data_Maybe_Just)(acc_4.UnsafePtr).V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()))
_ = v_4_1
var __t7 [][]gopurs_runtime.Value
{
if (v_4_1 == nil) {
__t7 = allArrays_3
goto end_branch_7
} else {

}
}
{
if (v_4_1 != nil) {
idx_2_loop = (idx_2) + (1)
allArrays_3_loop = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((v_4_1).V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := allArrays_3
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
continue go__go_1_0_0
__t7 = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Value{}.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
goto end_branch_7
} else {

}
}
{
__t7 = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
}
end_branch_7:
return func() gopurs_runtime.Value {
					arr := __t7
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
}
}()
})
})
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(go__go_1_0_0, gopurs_runtime.Int(0), func() gopurs_runtime.Value {
					arr := func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
}

func Call_Data_Array_foldRecM(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(array_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t3 bool
{
if (gopurs_runtime.RecordGet(o_7, "b").IntVal) < (gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(array_6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
if __t3 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.RecordGet(o_7, "a")})})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply2(f_4, gopurs_runtime.RecordGet(o_7, "a"), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(array_6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()[gopurs_runtime.RecordGet(o_7, "b").IntVal]
}))), gopurs_runtime.Func(func(res_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict2("a", "b", res_prime_8, gopurs_runtime.Int((gopurs_runtime.RecordGet(o_7, "b").IntVal) + (1)))})})
}))
}
end_branch_4:
return __t4
}), gopurs_runtime.RecordDict2("a", "b", b_5, gopurs_runtime.Int(0)))
})
})
})
}

func Call_Data_Array_foldMap(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_0)})
}

func Call_Data_Array_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_4)
}), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_4, a_6), gopurs_runtime.Func(func(b_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_Data_Array_foldM(dictMonad_0), f_3, b_prime_8, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(as_7.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}))
})
}), __local_var_5)
})
})
})
}
}

func Call_Data_Array_fold(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldMap"), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_0)}, Get_Data_Foldable_identity1())
}

func Call_Data_Array_findMap(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findMapImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, Get_Data_Maybe_isJust(), __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_Data_Array_findLastIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_Data_Array_insertBy(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> uint32
__local_var_4_1 := uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_3).IntVal)
_ = __local_var_4_1
var __t2 bool
{
if (__local_var_4_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 == 380165415) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}), gopurs_runtime.Array(ys_2)))
_ = __local_var_3_0
var __t4 int64
{
if (__local_var_3_0 == nil) {
__t4 = 0
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0 != nil) {
__t4 = ((__local_var_3_0).V0.IntVal) + (1)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_4:
// TAST (Let): i_4_3 -> int64
i_4_3 := __t4
_ = i_4_3
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> *Constructor_Data_Maybe_Just
__local_var_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_4_3), x_1, gopurs_runtime.Array(ys_2)))
_ = __local_var_6_5
var __t6 gopurs_runtime.Value
{
if (__local_var_6_5 != nil) {
__t6 = (__local_var_6_5).V0
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Array_insertBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_Array_findIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_Data_Array_span(p_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
// TAST (Let): breakIndex_2_0 -> gopurs_runtime.Value
breakIndex_2_0 := gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(p_0, x_2).IntVal) != (0)) != (true))
}), gopurs_runtime.Array(arr_1))
_ = breakIndex_2_0
var __t2 gopurs_runtime.Value
{
if (breakIndex_2_0.Type == 9 && breakIndex_2_0.IntVal == 930809136 && breakIndex_2_0.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if ((*Constructor_Data_Maybe_Just)(breakIndex_2_0.UnsafePtr).V0.IntVal) == (0) {
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(arr_1))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((*Constructor_Data_Maybe_Just)(breakIndex_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Array(arr_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int((*Constructor_Data_Maybe_Just)(breakIndex_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(arr_1))).IntVal), gopurs_runtime.Array(arr_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (breakIndex_2_0.Type == 9 && breakIndex_2_0.IntVal == 930809136 && breakIndex_2_0.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array(arr_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}

func Call_Data_Array_takeWhile(p_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Call_Data_Array_span(p_0, xs_1), "init").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_find(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
return xs_1[__local_var_3.IntVal]
})
})), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, f_0, gopurs_runtime.Array(xs_1))))}))
}

func Call_Data_Array_filter(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
arr_val_filterImpl0 := __local_var_1
_ = arr_val_filterImpl0
_ = arr_val_filterImpl0
arr_go_filterImpl0 := arr_val_filterImpl0
_ = arr_go_filterImpl0
res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
_ = res_go_filterImpl0
for _, v_filterImpl0 := range arr_go_filterImpl0 {
if gopurs_runtime.Apply(__local_var_0, v_filterImpl0).BoolVal() {
res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
} else {

}
}
return res_go_filterImpl0
}()
}

func Call_Data_Array_intersectBy(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
arr_val_filterImpl0 := xs_1
_ = arr_val_filterImpl0
_ = arr_val_filterImpl0
arr_go_filterImpl0 := arr_val_filterImpl0
_ = arr_go_filterImpl0
res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
_ = res_go_filterImpl0
for _, v_filterImpl0 := range arr_go_filterImpl0 {
if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> *Constructor_Data_Maybe_Just
__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply(eq_0, x_3), gopurs_runtime.Array(ys_2)))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0 == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 != nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
}), v_filterImpl0).BoolVal() {
res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
} else {

}
}
return res_go_filterImpl0
}()
}

func Call_Data_Array_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_intersectBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Array_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_Array_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Array_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_Array_notElem(dictEq_0_loop *Constructor_Data_Eq_Eq, a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, a_1).IntVal) != (0))
}), gopurs_runtime.Array(arr_2)))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0 == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_Data_Array_elem(dictEq_0_loop *Constructor_Data_Eq_Eq, a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, a_1).IntVal) != (0))
}), gopurs_runtime.Array(arr_2)))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0 == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_Data_Array_dropWhile(p_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Call_Data_Array_span(p_0, xs_1), "rest").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_dropEnd(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 -> int64
__local_var_2_0 := (gopurs_runtime.Int(int64(len(xs_1))).IntVal) - (n_0)
_ = __local_var_2_0
var __t2 []gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_2_0) < (1) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_2
} else {

}
}
{
__t2 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_2:
return __t2
}

func Call_Data_Array_drop(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t1 []gopurs_runtime.Value
{
var __t0 bool
{
if (n_0) < (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t1 = xs_1
goto end_branch_1
} else {

}
}
{
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_1:
return __t1
}

func Call_Data_Array_takeEnd(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 -> int64
__local_var_2_0 := (gopurs_runtime.Int(int64(len(xs_1))).IntVal) - (n_0)
_ = __local_var_2_0
var __t2 []gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_2_0) < (1) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = xs_1
goto end_branch_2
} else {

}
}
{
__t2 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_2:
return __t2
}

func Call_Data_Array_deleteAt(__local_var_0_loop int64, __local_var_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(__local_var_0), gopurs_runtime.Array(__local_var_1)))
}

func Call_Data_Array_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 []gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t5 []gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(len(v2_2))).IntVal) == (0) {
__t5 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_5
} else {

}
}
{
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Array(v2_2)))
_ = __local_var_3_0
var __t4 gopurs_runtime.Value
{
if (__local_var_3_0 == nil) {
__t4 = gopurs_runtime.Array(v2_2)
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0 != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (__local_var_3_0).V0
_ = __local_var_4_1
__t4 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> *Constructor_Data_Maybe_Just
__local_var_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(__local_var_4_1.IntVal), gopurs_runtime.Array(v2_2)))
_ = __local_var_6_2
var __t3 gopurs_runtime.Value
{
if (__local_var_6_2 != nil) {
__t3 = (__local_var_6_2).V0
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t4.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_5:
return __t5
}

func Call_Data_Array_delete(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_deleteBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_difference(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldr"), Call_Data_Array_delete(dictEq_0))
}

func Call_Data_Array_cons(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Array_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_Array_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4)
})))
})
})
}

func Call_Data_Array_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 -> *Constructor_Control_Alt_Alt
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply2(Call_Data_Array_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
})
})
}

func Call_Data_Array_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 []gopurs_runtime.Value = a_1_loop
_ = a_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Bind_bindArray(), "bind"), gopurs_runtime.Array(a_1), b_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Data_Array_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(f_0, x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
}

func Call_Data_Array_filterA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(p_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableArray(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), x_3), gopurs_runtime.Apply(p_2, x_3))
}))
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), Call_Data_Array_mapMaybe(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if ((*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1.IntVal) != (0) {
__t3 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})))
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
})
})
}

func Call_Data_Array_any(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (gopurs_runtime.UncurriedApp2(Get_Data_Array_anyImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Call_Data_Array_nubByEq(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_foreach(), gopurs_runtime.Array(xs_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Array_any(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eq_0, v_4, x_3).IntVal) != (0))
}))
_ = __local_var_4_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(__local_var_4_0, x_5))
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), arr_2)
})), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Array_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), x_3, arr_2)
}))
_ = __local_var_5_1
var __t2 gopurs_runtime.Value
{
if (e_4.IntVal) != (0) {
__t2 = __local_var_5_1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), Get_Data_Unit_unit())
}
end_branch_2:
return __t2
}))
})), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), arr_2)
})
}))
}))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_nubByEq(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_unionBy(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append"), gopurs_runtime.Array(xs_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_deleteBy(eq_0, a_4, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(b_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}), gopurs_runtime.Array(Call_Data_Array_nubByEq(eq_0, ys_2)), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_union(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_unionBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_alterAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0)))
_ = __local_var_3_0
var __t3 *Constructor_Data_Maybe_Just
{
if (__local_var_3_0 == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_3_0 != nil) {
// TAST (Let): v_4_1 -> *Constructor_Data_Maybe_Just
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, (__local_var_3_0).V0))
_ = v_4_1
var __t2 *Constructor_Data_Maybe_Just
{
if (v_4_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), gopurs_runtime.Array(xs_2)))
goto end_branch_2
} else {

}
}
{
if (v_4_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), (v_4_1).V0, gopurs_runtime.Array(xs_2)))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)})
}

func Call_Data_Array_all(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (gopurs_runtime.UncurriedApp2(Get_Data_Array_allImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Call_Data_Array_all__3571149479(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (gopurs_runtime.UncurriedApp2(Get_Data_Array_allImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Call_Data_Array_alterAt__2287604653(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0)))
_ = __local_var_3_0
var __t3 *Constructor_Data_Maybe_Just
{
if (__local_var_3_0 == nil) {
__t3 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_3
} else {

}
}
{
if (__local_var_3_0 != nil) {
// TAST (Let): v_4_1 -> *Constructor_Data_Maybe_Just
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply(f_1, (__local_var_3_0).V0))
_ = v_4_1
var __t2 *Constructor_Data_Maybe_Just
{
if (v_4_1 == nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), gopurs_runtime.Array(xs_2)))
goto end_branch_2
} else {

}
}
{
if (v_4_1 != nil) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), (v_4_1).V0, gopurs_runtime.Array(xs_2)))
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_2:
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_3:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)})
}

func Call_Data_Array_any__3571149479(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (gopurs_runtime.UncurriedApp2(Get_Data_Array_anyImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Call_Data_Array_concatMap__435921434(b_0_loop gopurs_runtime.Value, a_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 []gopurs_runtime.Value = a_1_loop
_ = a_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Bind_bindArray(), "bind"), gopurs_runtime.Array(a_1), b_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_concatMap__2727787066(b_0_loop gopurs_runtime.Value, a_1_loop []*Constructor_Data_Maybe_Just) []gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 []*Constructor_Data_Maybe_Just = a_1_loop
_ = a_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Bind_bindArray(), "bind"), func() gopurs_runtime.Value {
					arr := a_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), b_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_concatMap__2827656794(b_0_loop gopurs_runtime.Value, a_1_loop []*Constructor_Data_Tuple_Tuple) []gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 []*Constructor_Data_Tuple_Tuple = a_1_loop
_ = a_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Bind_bindArray(), "bind"), func() gopurs_runtime.Value {
					arr := a_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), b_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_cons__3485573810(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_cons__4002752745(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append"), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{x_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_delete__525954648(dictEq_0_loop *Constructor_Data_Eq_Eq) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_deleteBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_deleteAt__454851725(__local_var_0_loop int64, __local_var_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(__local_var_0), gopurs_runtime.Array(__local_var_1)))
}

func Call_Data_Array_deleteBy__519303411(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 []gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t5 []gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(len(v2_2))).IntVal) == (0) {
__t5 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_5
} else {

}
}
{
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Array(v2_2)))
_ = __local_var_3_0
var __t4 gopurs_runtime.Value
{
if (__local_var_3_0 == nil) {
__t4 = gopurs_runtime.Array(v2_2)
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0 != nil) {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := (__local_var_3_0).V0
_ = __local_var_4_1
__t4 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> *Constructor_Data_Maybe_Just
__local_var_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(__local_var_4_1.IntVal), gopurs_runtime.Array(v2_2)))
_ = __local_var_6_2
var __t3 gopurs_runtime.Value
{
if (__local_var_6_2 != nil) {
__t3 = (__local_var_6_2).V0
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_4:
__t5 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t4.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_5:
return __t5
}

func Call_Data_Array_drop__1426757676(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t1 []gopurs_runtime.Value
{
var __t0 bool
{
if (n_0) < (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t1 = xs_1
goto end_branch_1
} else {

}
}
{
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_1:
return __t1
}

func Call_Data_Array_dropEnd__1426757676(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 -> int64
__local_var_2_0 := (gopurs_runtime.Int(int64(len(xs_1))).IntVal) - (n_0)
_ = __local_var_2_0
var __t2 []gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_2_0) < (1) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_2
} else {

}
}
{
__t2 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_2:
return __t2
}

func Call_Data_Array_dropWhile__377906483(p_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Call_Data_Array_span(p_0, xs_1), "rest").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_elem__1106871500(dictEq_0_loop *Constructor_Data_Eq_Eq, a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, a_1).IntVal) != (0))
}), gopurs_runtime.Array(arr_2)))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0 == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_Data_Array_elemIndex__33401498(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Array_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_Array_elemLastIndex__33401498(dictEq_0_loop *Constructor_Data_Eq_Eq, x_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return gopurs_runtime.Apply(Get_Data_Array_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
}

func Call_Data_Array_filter__4047711382(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []int64) []int64 {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []int64 = __local_var_1_loop
_ = __local_var_1
return func() []int64 {
arr_val_filterImpl0 := __local_var_1
_ = arr_val_filterImpl0
_ = arr_val_filterImpl0
arr_go_filterImpl0 := arr_val_filterImpl0
_ = arr_go_filterImpl0
res_go_filterImpl0 := make([]int64, 0)
_ = res_go_filterImpl0
for _, v_filterImpl0 := range arr_go_filterImpl0 {
if gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Int(v_filterImpl0)).BoolVal() {
res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
} else {

}
}
return res_go_filterImpl0
}()
}

func Call_Data_Array_filter__377906483(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
arr_val_filterImpl0 := __local_var_1
_ = arr_val_filterImpl0
_ = arr_val_filterImpl0
arr_go_filterImpl0 := arr_val_filterImpl0
_ = arr_go_filterImpl0
res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
_ = res_go_filterImpl0
for _, v_filterImpl0 := range arr_go_filterImpl0 {
if gopurs_runtime.Apply(__local_var_0, v_filterImpl0).BoolVal() {
res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
} else {

}
}
return res_go_filterImpl0
}()
}

func Call_Data_Array_filterA__2723385228(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 -> *Constructor_Data_Functor_Functor
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(p_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> gopurs_runtime.Value
__local_var_3_1 := gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableArray(), "traverse"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), x_3), gopurs_runtime.Apply(p_2, x_3))
}))
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 -> gopurs_runtime.Value
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), Call_Data_Array_mapMaybe(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just
{
if ((*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1.IntVal) != (0) {
__t3 = &Constructor_Data_Maybe_Just{1, (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0}
goto end_branch_3
} else {

}
}
{
__t3 = (*Constructor_Data_Maybe_Just)(nil)
}
end_branch_3:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t3)}
})))
_ = __local_var_4_2
return gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Apply(__local_var_3_1, x_5))
})
})
}

func Call_Data_Array_find__2560752692(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_3 gopurs_runtime.Value) gopurs_runtime.Value {
return xs_1[__local_var_3.IntVal]
})
})), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, f_0, gopurs_runtime.Array(xs_1))))}))
}

func Call_Data_Array_findIndex__139581937(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_Data_Array_findLastIndex__139581937(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_Data_Array_findMap__3943035258(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findMapImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, Get_Data_Maybe_isJust(), __local_var_0, gopurs_runtime.Array(__local_var_1)))
}

func Call_Data_Array_foldM__2595407950(dictMonad_0_loop *Constructor_Control_Monad_Monad) gopurs_runtime.Value {
var dictMonad_0 *Constructor_Control_Monad_Monad = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 -> *Constructor_Control_Applicative_Applicative
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 -> *Constructor_Control_Bind_Bind
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func(func(f_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(__local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_4)
}), gopurs_runtime.Func(func(a_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(as_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_4, a_6), gopurs_runtime.Func(func(b_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_Data_Array_foldM(dictMonad_0), f_3, b_prime_8, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(as_7.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}))
})
}), __local_var_5)
})
})
})
}

func Call_Data_Array_foldRecM__306774880(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 -> gopurs_runtime.Value
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 -> *Constructor_Control_Bind_Bind
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
return gopurs_runtime.Func(func(f_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(b_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(array_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t4 gopurs_runtime.Value
{
var __t3 bool
{
if (gopurs_runtime.RecordGet(o_7, "b").IntVal) < (gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(array_6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal) {
__t3 = false
goto end_branch_3
} else {

}
}
{
__t3 = true
}
end_branch_3:
if __t3 {
__t4 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Done{1, gopurs_runtime.RecordGet(o_7, "a")})})
goto end_branch_4
} else {

}
}
{
__t4 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply2(f_4, gopurs_runtime.RecordGet(o_7, "a"), gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_8 gopurs_runtime.Value) gopurs_runtime.Value {
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(array_6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()[gopurs_runtime.RecordGet(o_7, "b").IntVal]
}))), gopurs_runtime.Func(func(res_prime_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(&Constructor_Control_Monad_Rec_Class_Loop{1, gopurs_runtime.RecordDict2("a", "b", res_prime_8, gopurs_runtime.Int((gopurs_runtime.RecordGet(o_7, "b").IntVal) + (1)))})})
}))
}
end_branch_4:
return __t4
}), gopurs_runtime.RecordDict2("a", "b", b_5, gopurs_runtime.Int(0)))
})
})
})
}

func Call_Data_Array_groupAllBy__1923945894(cmp_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
// TAST (Let): __local_var_1_0 -> gopurs_runtime.Value
__local_var_1_0 := gopurs_runtime.Apply(Get_Data_Array_groupBy(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 -> uint32
__local_var_3_1 := uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_2).IntVal)
_ = __local_var_3_1
var __t2 bool
{
if (__local_var_3_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1 == 380165415) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1 == 902936544) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
})
}))
_ = __local_var_1_0
return gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(__local_var_1_0, gopurs_runtime.Array(Call_Data_Array_sortBy(cmp_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
})
}

func Call_Data_Array_groupBy__693635452(op_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) [][]gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(result_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_1), gopurs_runtime.Int(v_3.IntVal))))}
})), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(0))), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply2(Get_Data_Array_ST_Iterator_iterate(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_3))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_5 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(sub_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), x_4, sub_5)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply3(Get_Data_Array_ST_Iterator_pushWhile(), gopurs_runtime.Apply(op_0, x_4), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_3))}, sub_5), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), sub_5)
}), gopurs_runtime.Func(func(grp_8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(grp_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), result_2)
})
}))
}))
}))
})))
})), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_2)
})
}))
}))
}))).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
}

func Call_Data_Array_head__2042355260(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(0)))
}

func Call_Data_Array_head__2056156327(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(0)))
}

func Call_Data_Array_head__1956412839(xs_0_loop []*Constructor_Data_Tuple_Tuple) *Constructor_Data_Maybe_Just {
var xs_0 []*Constructor_Data_Tuple_Tuple = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(0)))
}

func Call_Data_Array_index__4267297680(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) *Constructor_Data_Maybe_Just {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(__local_var_0), gopurs_runtime.Int(__local_var_1)))
}

func Call_Data_Array_index__2196477387(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) *Constructor_Data_Maybe_Just {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(__local_var_0), gopurs_runtime.Int(__local_var_1)))
}

func Call_Data_Array_index__2659850123(__local_var_0_loop []*Constructor_Data_Tuple_Tuple, __local_var_1_loop int64) *Constructor_Data_Maybe_Just {
var __local_var_0 []*Constructor_Data_Tuple_Tuple = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, func() gopurs_runtime.Value {
					arr := __local_var_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(__local_var_1)))
}

func Call_Data_Array_index__3264285643(__local_var_0_loop []*Constructor_Data_Tuple_Tuple, __local_var_1_loop int64) *Constructor_Data_Maybe_Just {
var __local_var_0 []*Constructor_Data_Tuple_Tuple = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, func() gopurs_runtime.Value {
					arr := __local_var_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(__local_var_1)))
}

func Call_Data_Array_init__7061562(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (0) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1)), gopurs_runtime.Array(xs_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_Array_init__976795489(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (0) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1)), gopurs_runtime.Array(xs_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})
}

func Call_Data_Array_insert__1514035128(dictOrd_0_loop *Constructor_Data_Ord_Ord) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Array_insertBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_Array_insertAt__388410084(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)))
}

func Call_Data_Array_insertBy__1563432905(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> uint32
__local_var_4_1 := uint32(gopurs_runtime.Apply2(cmp_0, x_1, y_3).IntVal)
_ = __local_var_4_1
var __t2 bool
{
if (__local_var_4_1 == 1527465420) {
__t2 = false
goto end_branch_2
} else {

}
}
{
if (__local_var_4_1 == 380165415) {
__t2 = true
goto end_branch_2
} else {

}
}
{
__t2 = false
}
end_branch_2:
return gopurs_runtime.Bool(__t2)
}), gopurs_runtime.Array(ys_2)))
_ = __local_var_3_0
var __t4 int64
{
if (__local_var_3_0 == nil) {
__t4 = 0
goto end_branch_4
} else {

}
}
{
if (__local_var_3_0 != nil) {
__t4 = ((__local_var_3_0).V0.IntVal) + (1)
goto end_branch_4
} else {

}
}
{
__t4 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_4:
// TAST (Let): i_4_3 -> int64
i_4_3 := __t4
_ = i_4_3
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 -> *Constructor_Data_Maybe_Just
__local_var_6_5 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_4_3), x_1, gopurs_runtime.Array(ys_2)))
_ = __local_var_6_5
var __t6 gopurs_runtime.Value
{
if (__local_var_6_5 != nil) {
__t6 = (__local_var_6_5).V0
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
})).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_intersectBy__145374773(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
arr_val_filterImpl0 := xs_1
_ = arr_val_filterImpl0
_ = arr_val_filterImpl0
arr_go_filterImpl0 := arr_val_filterImpl0
_ = arr_go_filterImpl0
res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
_ = res_go_filterImpl0
for _, v_filterImpl0 := range arr_go_filterImpl0 {
if gopurs_runtime.Apply(gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> *Constructor_Data_Maybe_Just
__local_var_4_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Apply(eq_0, x_3), gopurs_runtime.Array(ys_2)))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0 == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0 != nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return gopurs_runtime.Bool((__t1.IntVal) != (0))
}), v_filterImpl0).BoolVal() {
res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
} else {

}
}
return res_go_filterImpl0
}()
}

func Call_Data_Array_intersperse__4002752745(a_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
// TAST (Let): v_2_0 -> int64
v_2_0 := gopurs_runtime.Int(int64(len(arr_1))).IntVal
_ = v_2_0
var __t4 []gopurs_runtime.Value
{
var __t3 bool
{
if (v_2_0) < (2) {
__t3 = true
goto end_branch_3
} else {

}
}
{
__t3 = false
}
end_branch_3:
if __t3 {
__t4 = arr_1
goto end_branch_4
} else {

}
}
{
__t4 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(out_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_1 -> gopurs_runtime.Value
__local_var_4_1 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return arr_1[0]
}))
_ = __local_var_4_1
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_4_1, out_3)
}), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply3(Get_Control_Monad_ST_Internal_forImpl(), gopurs_runtime.Int(1), gopurs_runtime.Int(v_2_0), gopurs_runtime.Func(func(idx_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), a_0, out_3)
}), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_2 -> gopurs_runtime.Value
__local_var_7_2 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return arr_1[idx_5.IntVal]
}))
_ = __local_var_7_2
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(v_7 gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_7_2, out_3)
}))
}))
})), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), out_3)
}))
}))
})), Get_Data_Array_ST_unsafeFreeze())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_4:
return __t4
}

func Call_Data_Array_last__2042355260(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1))))
}

func Call_Data_Array_last__2345912124(xs_0_loop []*Constructor_Data_Tuple_Tuple) *Constructor_Data_Maybe_Just {
var xs_0 []*Constructor_Data_Tuple_Tuple = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1))))
}

func Call_Data_Array_last__2056156327(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1))))
}

func Call_Data_Array_many__1839052385(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 -> *Constructor_Control_Alt_Alt
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 -> *Constructor_Control_Applicative_Applicative
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply2(Call_Data_Array_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
})
})
}

func Call_Data_Array_mapMaybe__1271145181(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Data_Array_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(f_0, x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
}

func Call_Data_Array_mapMaybe__3137412285(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Data_Array_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(f_0, x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
}

func Call_Data_Array_mapMaybe__2261006141(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Data_Array_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 -> gopurs_runtime.Value
__local_var_2_0 := gopurs_runtime.Apply(f_0, x_1)
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0.Type == 9 && __local_var_2_0.IntVal == 930809136 && __local_var_2_0.UnsafePtr != nil) {
__t1 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(*Constructor_Data_Maybe_Just)(__local_var_2_0.UnsafePtr).V0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return __t1
}))
}

func Call_Data_Array_modifyAt__3384125836(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0)))
_ = __local_var_3_0
var __t1 *Constructor_Data_Maybe_Just
{
if (__local_var_3_0 == nil) {
__t1 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(i_0), gopurs_runtime.Apply(f_1, (__local_var_3_0).V0), gopurs_runtime.Array(xs_2))))})
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_1:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)})
}

func Call_Data_Array_modifyAtIndices__536948024(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, is_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, xs_3_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var is_1 gopurs_runtime.Value = is_1_loop
_ = is_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var xs_3 []gopurs_runtime.Value = xs_3_loop
_ = xs_3
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Func(func(res_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_Foldable_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Monad_ST_Internal_applicativeST()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Data_Array_ST_modify(), gopurs_runtime.Int(i_5.IntVal), f_2, res_4)
}), is_1)
}), gopurs_runtime.Array(xs_3))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_notElem__1106871500(dictEq_0_loop *Constructor_Data_Eq_Eq, a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
// TAST (Let): __local_var_3_0 -> *Constructor_Data_Maybe_Just
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, a_1).IntVal) != (0))
}), gopurs_runtime.Array(arr_2)))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0 == nil) {
__t1 = gopurs_runtime.Bool(true)
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return (__t1.IntVal) != (0)
}

func Call_Data_Array_nubBy__3347533344(comp_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): indexedAndSorted_2_0 -> []*Constructor_Data_Tuple_Tuple
indexedAndSorted_2_0 := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_sortBy(gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(comp_0, (*Constructor_Data_Tuple_Tuple)(x_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple)(y_3.UnsafePtr).V1).IntVal)), UnsafePtr: nil}
})
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_FunctorWithIndex_functorWithIndexArray(), "mapWithIndex"), Get_Data_Tuple_Tuple(), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
_ = indexedAndSorted_2_0
// TAST (Let): v_3_1 -> *Constructor_Data_Maybe_Just
v_3_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := indexedAndSorted_2_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Int(0)))
_ = v_3_1
var __t8 []gopurs_runtime.Value
{
if (v_3_1 == nil) {
__t8 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_8
} else {

}
}
{
if (v_3_1 != nil) {
__t8 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Functor_functorArray(), "map"), Get_Data_Tuple_snd(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_Data_Array_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord](gopurs_runtime.Value{Type: 9, IntVal: 1435789946, UnsafePtr: unsafe.Pointer(&Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqIntImpl())
}), gopurs_runtime.Apply3(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})})}), Get_Data_Tuple_fst()), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeThawImpl(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := []*Constructor_Data_Tuple_Tuple{gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple]((v_3_1).V0)}
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}())
}), gopurs_runtime.Func(func(result_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_foreach(), func() gopurs_runtime.Value {
					arr := indexedAndSorted_2_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> gopurs_runtime.Value
__local_var_6_2 := (*Constructor_Data_Tuple_Tuple)(v1_5.UnsafePtr).V1
_ = __local_var_6_2
// TAST (Let): __local_var_7_3 -> gopurs_runtime.Value
__local_var_7_3 := gopurs_runtime.Apply(Get_Partial_Unsafe__unsafePartial(), gopurs_runtime.Func(func(_dollar__unused_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_9_4 -> gopurs_runtime.Value
var __local_var_9_4 gopurs_runtime.Value = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(x_8.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).IntVal) - (1)))))}
var __t5 gopurs_runtime.Value
{
if (__local_var_9_4.Type == 9 && __local_var_9_4.IntVal == 930809136 && __local_var_9_4.UnsafePtr != nil) {
__t5 = (*Constructor_Data_Maybe_Just)(__local_var_9_4.UnsafePtr).V0
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
return __t5
})
}))
_ = __local_var_7_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(x_8 gopurs_runtime.Value) gopurs_runtime.Value {
return (*Constructor_Data_Tuple_Tuple)(gopurs_runtime.Apply(__local_var_7_3, x_8).UnsafePtr).V1
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_4)
})), gopurs_runtime.Func(func(lst_7 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_8_6 -> gopurs_runtime.Value
__local_var_8_6 := gopurs_runtime.Apply(Get_Data_Array_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v1_5))}, result_4)
}))
_ = __local_var_8_6
var __t7 gopurs_runtime.Value
{
if ((gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Ordering_eqOrdering(), "eq"), gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(comp_0, lst_7, __local_var_6_2).IntVal)), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}).IntVal) != (0)) != (true) {
__t7 = __local_var_8_6
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), Get_Data_Unit_unit())
}
end_branch_7:
return __t7
}))
})), gopurs_runtime.Func(func(_dollar__unused_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), result_4)
})
}))
}))).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_8
} else {

}
}
{
__t8 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_8:
return __t8
}

func Call_Data_Array_nubByEq__3443670074(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(arr_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_foreach(), gopurs_runtime.Array(xs_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_0 -> gopurs_runtime.Value
__local_var_4_0 := gopurs_runtime.Apply(Get_Data_Array_any(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eq_0, v_4, x_3).IntVal) != (0))
}))
_ = __local_var_4_0
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Func(func(x_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Data_HeytingAlgebra_heytingAlgebraBoolean(), "not"), gopurs_runtime.Apply(__local_var_4_0, x_5))
}), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), arr_2)
})), gopurs_runtime.Func(func(e_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_1 -> gopurs_runtime.Value
__local_var_5_1 := gopurs_runtime.Apply(Get_Data_Array_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), x_3, arr_2)
}))
_ = __local_var_5_1
var __t2 gopurs_runtime.Value
{
if (e_4.IntVal) != (0) {
__t2 = __local_var_5_1
goto end_branch_2
} else {

}
}
{
__t2 = gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), Get_Data_Unit_unit())
}
end_branch_2:
return __t2
}))
})), gopurs_runtime.Func(func(_dollar__unused_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), arr_2)
})
}))
}))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_null__1123412116(xs_0_loop []gopurs_runtime.Value) bool {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (0)
}

func Call_Data_Array_partition__3230398268(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return gopurs_runtime.UncurriedApp2(Get_Data_Array_partitionImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1))
}

func Call_Data_Array_replicate__1064418410(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp2(Get_Data_Array_replicateImpl(), gopurs_runtime.Int(__local_var_0), __local_var_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_scanl__3156262044(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanlImpl(), __local_var_0, __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_scanr__3156262044(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_scanrImpl(), __local_var_0, __local_var_1, gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_singleton__193199869(a_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_singleton__2286220742(a_0_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var a_0 gopurs_runtime.Value = a_0_loop
_ = a_0
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{a_0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_singleton__2535196422(a_0_loop *Constructor_Data_Tuple_Tuple) []*Constructor_Data_Tuple_Tuple {
var a_0 *Constructor_Data_Tuple_Tuple = a_0_loop
_ = a_0
return []*Constructor_Data_Tuple_Tuple{a_0}
}

func Call_Data_Array_slice__3011328576(__local_var_0_loop int64, __local_var_1_loop int64, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(__local_var_0), gopurs_runtime.Int(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_snoc__3419689714(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), x_1), gopurs_runtime.Array(xs_0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_snoc__3911647657(xs_0_loop []gopurs_runtime.Value, x_1_loop gopurs_runtime.Value) []gopurs_runtime.Value {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), x_1), gopurs_runtime.Array(xs_0))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_snoc__1505998191(xs_0_loop [][]gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) [][]gopurs_runtime.Value {
var xs_0 [][]gopurs_runtime.Value = xs_0_loop
_ = xs_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), gopurs_runtime.Array(x_1)), func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}())).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
}

func Call_Data_Array_some__1839052385(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Apply0_1_0 -> *Constructor_Control_Apply_Apply
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 -> *Constructor_Data_Functor_Functor
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func(func(dictLazy_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Array_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_Array_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy](dictLazy_3))}, v_4)
})))
})
})
}

func Call_Data_Array_sortBy__3347533344(comp_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if (uint32(v_2.IntVal) == 380165415) {
__t0 = 1
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 902936544) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 1527465420) {
__t0 = -1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return gopurs_runtime.Int(__t0)
}), gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_sortBy__1848798496(comp_0_loop gopurs_runtime.Value, __local_var_1_loop []*Constructor_Data_Tuple_Tuple) []*Constructor_Data_Tuple_Tuple {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 []*Constructor_Data_Tuple_Tuple = __local_var_1_loop
_ = __local_var_1
return func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t0 int64
{
if (uint32(v_2.IntVal) == 380165415) {
__t0 = 1
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 902936544) {
__t0 = 0
goto end_branch_0
} else {

}
}
{
if (uint32(v_2.IntVal) == 1527465420) {
__t0 = -1
goto end_branch_0
} else {

}
}
{
__t0 = func() gopurs_runtime.Value { panic("Failed pattern match") }().IntVal
}
end_branch_0:
return gopurs_runtime.Int(__t0)
}), func() gopurs_runtime.Value {
					arr := __local_var_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
}

func Call_Data_Array_sortWith__478414925(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Ord_Ord
__local_var_1_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqIntImpl())
}), gopurs_runtime.Apply3(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})}
_ = __local_var_1_0
return gopurs_runtime.Apply(Get_Data_Array_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_1_0.V1), gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_Array_sortWith__1917042304(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
// TAST (Let): __local_var_1_0 -> *Constructor_Data_Ord_Ord
__local_var_1_0 := &Constructor_Data_Ord_Ord{1, gopurs_runtime.Func(func(_dollar__unused_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict1("eq", Get_Data_Eq_eqIntImpl())
}), gopurs_runtime.Apply3(Get_Data_Ord_ordIntImpl(), gopurs_runtime.Value{Type: 9, IntVal: int64(1527465420), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(902936544), UnsafePtr: nil}, gopurs_runtime.Value{Type: 9, IntVal: int64(380165415), UnsafePtr: nil})}
_ = __local_var_1_0
return gopurs_runtime.Apply(Get_Data_Array_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(__local_var_1_0.V1), gopurs_runtime.Apply(f_0, x_2), gopurs_runtime.Apply(f_0, y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_Array_sortWith__3115909389(dictOrd_0_loop *Constructor_Data_Ord_Ord, f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(Get_Data_Array_sortBy(), gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)), UnsafePtr: nil}
})
}))
}

func Call_Data_Array_span__174751768(p_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
// TAST (Let): breakIndex_2_0 -> gopurs_runtime.Value
breakIndex_2_0 := gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(p_0, x_2).IntVal) != (0)) != (true))
}), gopurs_runtime.Array(arr_1))
_ = breakIndex_2_0
var __t2 gopurs_runtime.Value
{
if (breakIndex_2_0.Type == 9 && breakIndex_2_0.IntVal == 930809136 && breakIndex_2_0.UnsafePtr != nil) {
var __t1 gopurs_runtime.Value
{
if ((*Constructor_Data_Maybe_Just)(breakIndex_2_0.UnsafePtr).V0.IntVal) == (0) {
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(arr_1))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((*Constructor_Data_Maybe_Just)(breakIndex_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Array(arr_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int((*Constructor_Data_Maybe_Just)(breakIndex_2_0.UnsafePtr).V0.IntVal), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(arr_1))).IntVal), gopurs_runtime.Array(arr_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (breakIndex_2_0.Type == 9 && breakIndex_2_0.IntVal == 930809136 && breakIndex_2_0.UnsafePtr == nil) {
__t2 = gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array(arr_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}

func Call_Data_Array_splitAt__3301820464(v_0_loop int64, v1_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 []gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t1 gopurs_runtime.Value
{
var __t0 bool
{
if (v_0) > (0) {
__t0 = false
goto end_branch_0
} else {

}
}
{
__t0 = true
}
end_branch_0:
if __t0 {
__t1 = gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(v1_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(v1_1))).IntVal), gopurs_runtime.Array(v1_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(v_0), gopurs_runtime.Array(v1_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}
end_branch_1:
return __t1
}

func Call_Data_Array_tail__7061562(__local_var_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})}
})
}), gopurs_runtime.Array(__local_var_0)))
}

func Call_Data_Array_take__1426757676(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t1 []gopurs_runtime.Value
{
var __t0 bool
{
if (n_0) < (1) {
__t0 = true
goto end_branch_0
} else {

}
}
{
__t0 = false
}
end_branch_0:
if __t0 {
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_1
} else {

}
}
{
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int(n_0), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_1:
return __t1
}

func Call_Data_Array_takeEnd__1426757676(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 -> int64
__local_var_2_0 := (gopurs_runtime.Int(int64(len(xs_1))).IntVal) - (n_0)
_ = __local_var_2_0
var __t2 []gopurs_runtime.Value
{
var __t1 bool
{
if (__local_var_2_0) < (1) {
__t1 = true
goto end_branch_1
} else {

}
}
{
__t1 = false
}
end_branch_1:
if __t1 {
__t2 = xs_1
goto end_branch_2
} else {

}
}
{
__t2 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(gopurs_runtime.Int(int64(len(xs_1))).IntVal), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_2:
return __t2
}

func Call_Data_Array_takeWhile__377906483(p_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(Call_Data_Array_span(p_0, xs_1), "init").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_transpose__4194748859(xs_0_loop [][]gopurs_runtime.Value) [][]gopurs_runtime.Value {
var xs_0 [][]gopurs_runtime.Value = xs_0_loop
_ = xs_0
var go__go_1_0_1 gopurs_runtime.Value
go__go_1_0_1 = gopurs_runtime.Func(func(idx_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(allArrays_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
var idx_2_loop int64 = idx_2_loop_val.IntVal
var allArrays_3_loop [][]gopurs_runtime.Value = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(allArrays_3_loop_val.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
go__go_1_0_1:
for {
if false { continue go__go_1_0_1 }
var idx_2 int64 = idx_2_loop
_ = idx_2
var allArrays_3 [][]gopurs_runtime.Value = allArrays_3_loop
_ = allArrays_3
// TAST (Let): v_4_1 -> *Constructor_Data_Maybe_Just
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl"), gopurs_runtime.Func(func(acc_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(nextArr_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 -> *Constructor_Data_Maybe_Just
__local_var_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(nextArr_5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Int(idx_2)))
_ = __local_var_6_2
var __t6 *Constructor_Data_Maybe_Just
{
if (__local_var_6_2 == nil) {
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](acc_4)
goto end_branch_6
} else {

}
}
{
if (__local_var_6_2 != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](acc_4)
if (__t_tag_3 == nil) {
__t5 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(__local_var_6_2).V0}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](acc_4)
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), (__local_var_6_2).V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((*Constructor_Data_Maybe_Just)(acc_4.UnsafePtr).V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_5
} else {

}
}
{
__t5 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_5:
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
goto end_branch_6
} else {

}
}
{
__t6 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](func() gopurs_runtime.Value { panic("Failed pattern match") }())
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t6)}
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()))
_ = v_4_1
var __t7 [][]gopurs_runtime.Value
{
if (v_4_1 == nil) {
__t7 = allArrays_3
goto end_branch_7
} else {

}
}
{
if (v_4_1 != nil) {
idx_2_loop = (idx_2) + (1)
allArrays_3_loop = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((v_4_1).V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := allArrays_3
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
continue go__go_1_0_1
__t7 = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Value{}.UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
goto end_branch_7
} else {

}
}
{
__t7 = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value { panic("Failed pattern match") }().UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
}
end_branch_7:
return func() gopurs_runtime.Value {
					arr := __t7
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
}
}()
})
})
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(go__go_1_0_1, gopurs_runtime.Int(0), func() gopurs_runtime.Value {
					arr := func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([][]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}() }
					return unboxed
				}()
}

func Call_Data_Array_uncons__2020799173(__local_var_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}
}), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(&Constructor_Data_Maybe_Just{1, gopurs_runtime.RecordDict2("head", "tail", x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))})}
})
}), gopurs_runtime.Array(__local_var_0)))
}

func Call_Data_Array_unionBy__145374773(eq_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var eq_0 gopurs_runtime.Value = eq_0_loop
_ = eq_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Semigroup_semigroupArray(), "append"), gopurs_runtime.Array(xs_1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Data_Foldable_foldableArray(), "foldl"), gopurs_runtime.Func(func(b_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_deleteBy(eq_0, a_4, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(b_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
})
}), gopurs_runtime.Array(Call_Data_Array_nubByEq(eq_0, ys_2)), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_unsafeIndex__2808089623(_dollar__unused_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar__unused_0 gopurs_runtime.Value = _dollar__unused_0_loop
_ = _dollar__unused_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return __local_var_1[__local_var_2]
}

func Call_Data_Array_unsnoc__2531125997(xs_0_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 *Constructor_Data_Maybe_Just
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (0) {
__t0 = (*Constructor_Data_Maybe_Just)(nil)
goto end_branch_0
} else {

}
}
{
__t0 = &Constructor_Data_Maybe_Just{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1)), gopurs_runtime.Array(xs_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}
}
end_branch_0:
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_applyMaybe(), "apply"), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Maybe_functorMaybe(), "map"), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(v1_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v_1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), v1_2)
})
}), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t0)})))}, gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (1)))))}))
}

func Call_Data_Array_unzip__1480671261(xs_0_loop []*Constructor_Data_Tuple_Tuple) *Constructor_Data_Tuple_Tuple {
var xs_0 []*Constructor_Data_Tuple_Tuple = xs_0_loop
_ = xs_0
return gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(fsts_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), Get_Data_Array_ST_newImpl(), gopurs_runtime.Func(func(snds_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_functorST(), "map"), gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_Iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(v_3.IntVal))))}
})), gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_newImpl(), gopurs_runtime.Int(0))), gopurs_runtime.Func(func(iter_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply2(Get_Data_Array_ST_Iterator_iterate(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator](iter_3))}, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_0 -> gopurs_runtime.Value
__local_var_5_0 := (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1
_ = __local_var_5_0
return gopurs_runtime.Apply3(gopurs_runtime.RecordGet(Get_Control_Bind_discardUnit(), "discard"), gopurs_runtime.Value{Type: 9, IntVal: 4032919565, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind](Get_Control_Monad_ST_Internal_bindST()))}, gopurs_runtime.Apply(Get_Data_Array_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0, fsts_1)
})), gopurs_runtime.Func(func(_dollar__unused_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(Get_Data_Array_void(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_ST_pushImpl(), __local_var_5_0, snds_2)
}))
}))
})), gopurs_runtime.Func(func(_dollar__unused_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), fsts_1)
}), gopurs_runtime.Func(func(fsts_prime_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_bindST(), "bind"), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp(Get_Data_Array_ST_unsafeFreezeImpl(), snds_2)
}), gopurs_runtime.Func(func(snds_prime_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.RecordGet(Get_Control_Monad_ST_Internal_applicativeST(), "pure"), gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(&Constructor_Data_Tuple_Tuple{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(fsts_prime_5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(snds_prime_6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())})})
}))
}))
}))
}))
}))
}))))
}

func Call_Data_Array_updateAt__388410084(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) *Constructor_Data_Maybe_Just {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just)(nil))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2)))
}

func Call_Data_Array_updateAtIndices__889079281(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable, us_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable = dictFoldable_0_loop
_ = dictFoldable_0
var us_1 gopurs_runtime.Value = us_1_loop
_ = us_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Func(func(res_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_Foldable_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative](Get_Control_Monad_ST_Internal_applicativeST()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Data_Array_ST_pokeImpl(), gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple)(v_4.UnsafePtr).V1, res_3)
})
}), us_1)
}), gopurs_runtime.Array(xs_2))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_zipWith__1350747206(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_zipWith__1220584870(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []*Constructor_Data_Tuple_Tuple {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []*Constructor_Data_Tuple_Tuple {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple, len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple](v) }
					return unboxed
				}()
}

func Call_Data_Array_zipWith__2000246342(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1), gopurs_runtime.Array(__local_var_2)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_zipWithA__3208598546(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value, ys_3_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
var ys_3 []gopurs_runtime.Value = ys_3_loop
_ = ys_3
return gopurs_runtime.Apply2(gopurs_runtime.RecordGet(Get_Data_Traversable_traversableArray(), "sequence"), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), f_1, gopurs_runtime.Array(xs_2), gopurs_runtime.Array(ys_3)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Get_Data_Array__deleteAt() gopurs_runtime.Value {
	return _Gopurs_Data_Array__DeleteAt
}

func Get_Data_Array__insertAt() gopurs_runtime.Value {
	return _Gopurs_Data_Array__InsertAt
}

func Get_Data_Array__updateAt() gopurs_runtime.Value {
	return _Gopurs_Data_Array__UpdateAt
}

func Get_Data_Array_allImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_AllImpl
}

func Get_Data_Array_anyImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_AnyImpl
}

func Get_Data_Array_concat() gopurs_runtime.Value {
	return _Gopurs_Data_Array_Concat
}

func Get_Data_Array_filterImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FilterImpl
}

func Get_Data_Array_findIndexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FindIndexImpl
}

func Get_Data_Array_findLastIndexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FindLastIndexImpl
}

func Get_Data_Array_findMapImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FindMapImpl
}

func Get_Data_Array_fromFoldableImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_FromFoldableImpl
}

func Get_Data_Array_indexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_IndexImpl
}

func Get_Data_Array_length() gopurs_runtime.Value {
	return _Gopurs_Data_Array_Length
}

func Get_Data_Array_partitionImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_PartitionImpl
}

func Get_Data_Array_rangeImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_RangeImpl
}

func Get_Data_Array_replicateImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ReplicateImpl
}

func Get_Data_Array_reverse() gopurs_runtime.Value {
	return _Gopurs_Data_Array_Reverse
}

func Get_Data_Array_scanlImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ScanlImpl
}

func Get_Data_Array_scanrImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ScanrImpl
}

func Get_Data_Array_sliceImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_SliceImpl
}

func Get_Data_Array_sortByImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_SortByImpl
}

func Get_Data_Array_unconsImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_UnconsImpl
}

func Get_Data_Array_unsafeIndexImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_UnsafeIndexImpl
}

func Get_Data_Array_zipWithImpl() gopurs_runtime.Value {
	return _Gopurs_Data_Array_ZipWithImpl
}

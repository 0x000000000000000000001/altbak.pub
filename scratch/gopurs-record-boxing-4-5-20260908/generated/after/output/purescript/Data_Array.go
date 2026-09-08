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
return Call_Data_Array_intercalate1(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_Data_Array_intercalate1
}

var cache_Data_Array_zero gopurs_runtime.Value
var once_Data_Array_zero sync.Once
func Get_Data_Array_zero() gopurs_runtime.Value {
	once_Data_Array_zero.Do(func() {
		cache_Data_Array_zero = gopurs_runtime.Int(int64(0))
	})
	return cache_Data_Array_zero
}

var cache_Data_Array_one gopurs_runtime.Value
var once_Data_Array_one sync.Once
func Get_Data_Array_one() gopurs_runtime.Value {
	once_Data_Array_one.Do(func() {
		cache_Data_Array_one = gopurs_runtime.Int(int64(1))
	})
	return cache_Data_Array_one
}

var cache_Data_Array_void gopurs_runtime.Value
var once_Data_Array_void sync.Once
func Get_Data_Array_void() gopurs_runtime.Value {
	once_Data_Array_void.Do(func() {
		cache_Data_Array_void = gopurs_runtime.Func(func(__local_var_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_void(__local_var_0_box)
})
	})
	return cache_Data_Array_void
}

var cache_Data_Array_pure gopurs_runtime.Value
var once_Data_Array_pure sync.Once
func Get_Data_Array_pure() gopurs_runtime.Value {
	once_Data_Array_pure.Do(func() {
		cache_Data_Array_pure = Get_Control_Monad_ST_Internal_pure_()
	})
	return cache_Data_Array_pure
}

var cache_Data_Array_fromJust gopurs_runtime.Value
var once_Data_Array_fromJust sync.Once
func Get_Data_Array_fromJust() gopurs_runtime.Value {
	once_Data_Array_fromJust.Do(func() {
		cache_Data_Array_fromJust = gopurs_runtime.Func(func(v_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_fromJust(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]]](v_0_box))
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Data_Array_fromJust
}

var cache_Data_Array_foldMap1 gopurs_runtime.Value
var once_Data_Array_foldMap1 sync.Once
func Get_Data_Array_foldMap1() gopurs_runtime.Value {
	once_Data_Array_foldMap1.Do(func() {
		cache_Data_Array_foldMap1 = gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()).V0)
	})
	return cache_Data_Array_foldMap1
}

var cache_Data_Array_fold1 gopurs_runtime.Value
var once_Data_Array_fold1 sync.Once
func Get_Data_Array_fold1() gopurs_runtime.Value {
	once_Data_Array_fold1.Do(func() {
		cache_Data_Array_fold1 = gopurs_runtime.Apply(Get_Data_Foldable_fold(), gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()))})
	})
	return cache_Data_Array_fold1
}

var cache_Data_Array_not gopurs_runtime.Value
var once_Data_Array_not sync.Once
func Get_Data_Array_not() gopurs_runtime.Value {
	once_Data_Array_not.Do(func() {
		cache_Data_Array_not = Get_Data_HeytingAlgebra_boolNot()
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
return Call_Data_Array_zipWithA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box), f_1_box, func() []gopurs_runtime.Value {
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
return gopurs_runtime.Array(Call_Data_Array_updateAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), us_1_box, func() []gopurs_runtime.Value {
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
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_updateAt(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_updateAt
}

var cache_Data_Array_unsafeIndex gopurs_runtime.Value
var once_Data_Array_unsafeIndex sync.Once
func Get_Data_Array_unsafeIndex() gopurs_runtime.Value {
	once_Data_Array_unsafeIndex.Do(func() {
		cache_Data_Array_unsafeIndex = gopurs_runtime.Func3(func(_dollar___unused_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_unsafeIndex(_dollar___unused_0_box, func() []gopurs_runtime.Value {
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
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_uncons(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_uncons
}

var cache_Data_Array_toUnfoldable gopurs_runtime.Value
var once_Data_Array_toUnfoldable sync.Once
func Get_Data_Array_toUnfoldable() gopurs_runtime.Value {
	once_Data_Array_toUnfoldable.Do(func() {
		cache_Data_Array_toUnfoldable = gopurs_runtime.Func2(func(dictUnfoldable_0_box gopurs_runtime.Value, xs_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_toUnfoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value]](dictUnfoldable_0_box), func() []gopurs_runtime.Value {
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
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_tail(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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
return Call_Data_Array_sortWith(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box), f_1_box)
})
	})
	return cache_Data_Array_sortWith
}

var cache_Data_Array_sort gopurs_runtime.Value
var once_Data_Array_sort sync.Once
func Get_Data_Array_sort() gopurs_runtime.Value {
	once_Data_Array_sort.Do(func() {
		cache_Data_Array_sort = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_sort(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
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
return func() gopurs_runtime.Value {
				orig := Call_Data_Array_splitAt(v_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(v1_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("after", "before", gopurs_runtime.Array(orig.after), gopurs_runtime.Array(orig.before))
				}()
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
return func() gopurs_runtime.Value {
				orig := Call_Data_Array_partition(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("no", "yes", gopurs_runtime.Array(orig.no), gopurs_runtime.Array(orig.yes))
				}()
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
return gopurs_runtime.Array(Call_Data_Array_modifyAtIndices(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box), is_1_box, f_2_box, func() []gopurs_runtime.Value {
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
		cache_Data_Array_mapWithIndex = Get_Data_FunctorWithIndex_mapWithIndexArray()
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
return Call_Data_Array_intercalate(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_Data_Array_intercalate
}

var cache_Data_Array_insertAt gopurs_runtime.Value
var once_Data_Array_insertAt sync.Once
func Get_Data_Array_insertAt() gopurs_runtime.Value {
	once_Data_Array_insertAt.Do(func() {
		cache_Data_Array_insertAt = gopurs_runtime.Func3(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value, __local_var_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_insertAt(__local_var_0_box.IntVal, __local_var_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_insertAt
}

var cache_Data_Array_go__init gopurs_runtime.Value
var once_Data_Array_go__init sync.Once
func Get_Data_Array_go__init() gopurs_runtime.Value {
	once_Data_Array_go__init.Do(func() {
		cache_Data_Array_go__init = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_go__init(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_go__init
}

var cache_Data_Array_index gopurs_runtime.Value
var once_Data_Array_index sync.Once
func Get_Data_Array_index() gopurs_runtime.Value {
	once_Data_Array_index.Do(func() {
		cache_Data_Array_index = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_index(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), __local_var_1_box.IntVal)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_index
}

var cache_Data_Array_last gopurs_runtime.Value
var once_Data_Array_last sync.Once
func Get_Data_Array_last() gopurs_runtime.Value {
	once_Data_Array_last.Do(func() {
		cache_Data_Array_last = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_last(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_last
}

var cache_Data_Array_unsnoc gopurs_runtime.Value
var once_Data_Array_unsnoc sync.Once
func Get_Data_Array_unsnoc() gopurs_runtime.Value {
	once_Data_Array_unsnoc.Do(func() {
		cache_Data_Array_unsnoc = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_unsnoc(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_unsnoc
}

var cache_Data_Array_modifyAt gopurs_runtime.Value
var once_Data_Array_modifyAt sync.Once
func Get_Data_Array_modifyAt() gopurs_runtime.Value {
	once_Data_Array_modifyAt.Do(func() {
		cache_Data_Array_modifyAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_modifyAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_modifyAt
}

var cache_Data_Array_unzip gopurs_runtime.Value
var once_Data_Array_unzip sync.Once
func Get_Data_Array_unzip() gopurs_runtime.Value {
	once_Data_Array_unzip.Do(func() {
		cache_Data_Array_unzip = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_unzip(func() []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](v) }
					return unboxed
				}())
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}()
})
	})
	return cache_Data_Array_unzip
}

var cache_Data_Array_head gopurs_runtime.Value
var once_Data_Array_head sync.Once
func Get_Data_Array_head() gopurs_runtime.Value {
	once_Data_Array_head.Do(func() {
		cache_Data_Array_head = gopurs_runtime.Func(func(xs_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_head(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_0_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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
return Call_Data_Array_nub(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
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
		cache_Data_Array_groupAllBy = gopurs_runtime.Func2(func(cmp_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_groupAllBy(cmp_0_box, func() []gopurs_runtime.Value {
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
	return cache_Data_Array_groupAllBy
}

var cache_Data_Array_groupAll gopurs_runtime.Value
var once_Data_Array_groupAll sync.Once
func Get_Data_Array_groupAll() gopurs_runtime.Value {
	once_Data_Array_groupAll.Do(func() {
		cache_Data_Array_groupAll = gopurs_runtime.Func(func(dictOrd_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_groupAll(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Array_groupAll
}

var cache_Data_Array_group gopurs_runtime.Value
var once_Data_Array_group sync.Once
func Get_Data_Array_group() gopurs_runtime.Value {
	once_Data_Array_group.Do(func() {
		cache_Data_Array_group = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_group(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_group
}

var cache_Data_Array_fromFoldable gopurs_runtime.Value
var once_Data_Array_fromFoldable sync.Once
func Get_Data_Array_fromFoldable() gopurs_runtime.Value {
	once_Data_Array_fromFoldable.Do(func() {
		cache_Data_Array_fromFoldable = gopurs_runtime.Func(func(dictFoldable_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_fromFoldable(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](dictFoldable_0_box))
})
	})
	return cache_Data_Array_fromFoldable
}

var cache_Data_Array_foldr gopurs_runtime.Value
var once_Data_Array_foldr sync.Once
func Get_Data_Array_foldr() gopurs_runtime.Value {
	once_Data_Array_foldr.Do(func() {
		cache_Data_Array_foldr = Get_Data_Foldable_foldrArray()
	})
	return cache_Data_Array_foldr
}

var cache_Data_Array_foldl gopurs_runtime.Value
var once_Data_Array_foldl sync.Once
func Get_Data_Array_foldl() gopurs_runtime.Value {
	once_Data_Array_foldl.Do(func() {
		cache_Data_Array_foldl = Get_Data_Foldable_foldlArray()
	})
	return cache_Data_Array_foldl
}

var cache_Data_Array_foldl__1894943466 gopurs_runtime.Value
var once_Data_Array_foldl__1894943466 sync.Once
func Get_Data_Array_foldl__1894943466() gopurs_runtime.Value {
	once_Data_Array_foldl__1894943466.Do(func() {
		cache_Data_Array_foldl__1894943466 = gopurs_runtime.Func3(func(__eta_norm_2_unused_0_box gopurs_runtime.Value, __eta_norm_1_1_box gopurs_runtime.Value, __eta_norm_0_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Int(Call_Data_Array_foldl__1894943466(__eta_norm_2_unused_0_box, __eta_norm_1_1_box.IntVal, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_2_box.UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()))
})
	})
	return cache_Data_Array_foldl__1894943466
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
return Call_Data_Array_foldRecM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]](dictMonadRec_0_box))
})
	})
	return cache_Data_Array_foldRecM
}

var cache_Data_Array_foldMap gopurs_runtime.Value
var once_Data_Array_foldMap sync.Once
func Get_Data_Array_foldMap() gopurs_runtime.Value {
	once_Data_Array_foldMap.Do(func() {
		cache_Data_Array_foldMap = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_foldMap(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_Data_Array_foldMap
}

var cache_Data_Array_foldM gopurs_runtime.Value
var once_Data_Array_foldM sync.Once
func Get_Data_Array_foldM() gopurs_runtime.Value {
	once_Data_Array_foldM.Do(func() {
		cache_Data_Array_foldM = gopurs_runtime.Func(func(dictMonad_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_foldM(gopurs_runtime.CoerceToStruct[Constructor_Control_Monad_Monad[gopurs_runtime.Value]](dictMonad_0_box))
})
	})
	return cache_Data_Array_foldM
}

var cache_Data_Array_fold gopurs_runtime.Value
var once_Data_Array_fold sync.Once
func Get_Data_Array_fold() gopurs_runtime.Value {
	once_Data_Array_fold.Do(func() {
		cache_Data_Array_fold = gopurs_runtime.Func(func(dictMonoid_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_fold(gopurs_runtime.CoerceToStruct[Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]](dictMonoid_0_box))
})
	})
	return cache_Data_Array_fold
}

var cache_Data_Array_findMap gopurs_runtime.Value
var once_Data_Array_findMap sync.Once
func Get_Data_Array_findMap() gopurs_runtime.Value {
	once_Data_Array_findMap.Do(func() {
		cache_Data_Array_findMap = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_findMap(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_findMap
}

var cache_Data_Array_findLastIndex gopurs_runtime.Value
var once_Data_Array_findLastIndex sync.Once
func Get_Data_Array_findLastIndex() gopurs_runtime.Value {
	once_Data_Array_findLastIndex.Do(func() {
		cache_Data_Array_findLastIndex = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_findLastIndex(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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
return Call_Data_Array_insert(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[gopurs_runtime.Value]](dictOrd_0_box))
})
	})
	return cache_Data_Array_insert
}

var cache_Data_Array_findIndex gopurs_runtime.Value
var once_Data_Array_findIndex sync.Once
func Get_Data_Array_findIndex() gopurs_runtime.Value {
	once_Data_Array_findIndex.Do(func() {
		cache_Data_Array_findIndex = gopurs_runtime.Func2(func(__local_var_0_box gopurs_runtime.Value, __local_var_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_findIndex(__local_var_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_findIndex
}

var cache_Data_Array_span gopurs_runtime.Value
var once_Data_Array_span sync.Once
func Get_Data_Array_span() gopurs_runtime.Value {
	once_Data_Array_span.Do(func() {
		cache_Data_Array_span = gopurs_runtime.Func2(func(p_0_box gopurs_runtime.Value, arr_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				orig := Call_Data_Array_span(p_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(arr_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				_ = orig
				return gopurs_runtime.RecordDict2("init", "rest", gopurs_runtime.Array(orig.go__init), gopurs_runtime.Array(orig.rest))
				}()
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
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_find(f_0_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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

var cache_Data_Array_filter__1859503602 gopurs_runtime.Value
var once_Data_Array_filter__1859503602 sync.Once
func Get_Data_Array_filter__1859503602() gopurs_runtime.Value {
	once_Data_Array_filter__1859503602.Do(func() {
		cache_Data_Array_filter__1859503602 = gopurs_runtime.Func2(func(__eta_norm_1_0_box gopurs_runtime.Value, __eta_norm_0_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_Data_Array_filter__1859503602(__eta_norm_1_0_box, func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(__eta_norm_0_1_box.UnsafePtr)
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
	return cache_Data_Array_filter__1859503602
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
return Call_Data_Array_intersect(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_intersect
}

var cache_Data_Array_elemLastIndex gopurs_runtime.Value
var once_Data_Array_elemLastIndex sync.Once
func Get_Data_Array_elemLastIndex() gopurs_runtime.Value {
	once_Data_Array_elemLastIndex.Do(func() {
		cache_Data_Array_elemLastIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_elemLastIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_elemLastIndex
}

var cache_Data_Array_elemIndex gopurs_runtime.Value
var once_Data_Array_elemIndex sync.Once
func Get_Data_Array_elemIndex() gopurs_runtime.Value {
	once_Data_Array_elemIndex.Do(func() {
		cache_Data_Array_elemIndex = gopurs_runtime.Func2(func(dictEq_0_box gopurs_runtime.Value, x_1_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_elemIndex(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), x_1_box)
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
})
	})
	return cache_Data_Array_elemIndex
}

var cache_Data_Array_notElem gopurs_runtime.Value
var once_Data_Array_notElem sync.Once
func Get_Data_Array_notElem() gopurs_runtime.Value {
	once_Data_Array_notElem.Do(func() {
		cache_Data_Array_notElem = gopurs_runtime.Func3(func(dictEq_0_box gopurs_runtime.Value, a_1_box gopurs_runtime.Value, arr_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(Call_Data_Array_notElem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), a_1_box, func() []gopurs_runtime.Value {
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
return gopurs_runtime.Bool(Call_Data_Array_elem(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box), a_1_box, func() []gopurs_runtime.Value {
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
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_deleteAt(__local_var_0_box.IntVal, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__local_var_1_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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

var cache_Data_Array_go__delete gopurs_runtime.Value
var once_Data_Array_go__delete sync.Once
func Get_Data_Array_go__delete() gopurs_runtime.Value {
	once_Data_Array_go__delete.Do(func() {
		cache_Data_Array_go__delete = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_go__delete(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_go__delete
}

var cache_Data_Array_difference gopurs_runtime.Value
var once_Data_Array_difference sync.Once
func Get_Data_Array_difference() gopurs_runtime.Value {
	once_Data_Array_difference.Do(func() {
		cache_Data_Array_difference = gopurs_runtime.Func(func(dictEq_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_difference(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
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
return Call_Data_Array_some(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
})
	})
	return cache_Data_Array_some
}

var cache_Data_Array_many gopurs_runtime.Value
var once_Data_Array_many sync.Once
func Get_Data_Array_many() gopurs_runtime.Value {
	once_Data_Array_many.Do(func() {
		cache_Data_Array_many = gopurs_runtime.Func(func(dictAlternative_0_box gopurs_runtime.Value) gopurs_runtime.Value {
return Call_Data_Array_many(gopurs_runtime.CoerceToStruct[Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]](dictAlternative_0_box))
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
return Call_Data_Array_filterA(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](dictApplicative_0_box))
})
	})
	return cache_Data_Array_filterA
}

var cache_Data_Array_catMaybes gopurs_runtime.Value
var once_Data_Array_catMaybes sync.Once
func Get_Data_Array_catMaybes() gopurs_runtime.Value {
	once_Data_Array_catMaybes.Do(func() {
		cache_Data_Array_catMaybes = Call_Data_Array_mapMaybe(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
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
return Call_Data_Array_nubEq(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
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
return Call_Data_Array_union(gopurs_runtime.CoerceToStruct[Constructor_Data_Eq_Eq[gopurs_runtime.Value]](dictEq_0_box))
})
	})
	return cache_Data_Array_union
}

var cache_Data_Array_alterAt gopurs_runtime.Value
var once_Data_Array_alterAt sync.Once
func Get_Data_Array_alterAt() gopurs_runtime.Value {
	once_Data_Array_alterAt.Do(func() {
		cache_Data_Array_alterAt = gopurs_runtime.Func3(func(i_0_box gopurs_runtime.Value, f_1_box gopurs_runtime.Value, xs_2_box gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
				_v := Call_Data_Array_alterAt(i_0_box.IntVal, f_1_box, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2_box.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()
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

func Call_Data_Array_intercalate1(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_0.V0), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.Box(dictMonoid_0.V1)
_ = mempty_2_1
return gopurs_runtime.Func2(func(sep_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(func() gopurs_runtime.Value {
arr_val_foldlArray4 := xs_4
_ = arr_val_foldlArray4
res_go_foldlArray4 := func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	go__init bool
}{mempty_2_1, true}
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "init", orig.acc, gopurs_runtime.Bool(orig.go__init))
				}()
_ = res_go_foldlArray4
arr_go_foldlArray4 := (*[]gopurs_runtime.Value)(arr_val_foldlArray4.UnsafePtr)
_ = arr_go_foldlArray4
for _, v_foldlArray4 := range *arr_go_foldlArray4 {
res_go_foldlArray4 = gopurs_runtime.Apply2(gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 struct{
	acc gopurs_runtime.Value
	go__init bool
}
{
if (gopurs_runtime.RecordGet(v_5, "init").IntVal) != (0) {
__t2 = struct{
	acc gopurs_runtime.Value
	go__init bool
}{v1_6, false}
goto end_branch_2
} else {

}
}
{
__t2 = struct{
	acc gopurs_runtime.Value
	go__init bool
}{gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.RecordGet(v_5, "acc"), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), sep_3, v1_6)), false}
}
end_branch_2:
return func() gopurs_runtime.Value {
				orig := __t2
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "init", orig.acc, gopurs_runtime.Bool(orig.go__init))
				}()
}), res_go_foldlArray4, v_foldlArray4)
}
return res_go_foldlArray4
}(), "acc")
})
}

func Call_Data_Array_void(__local_var_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_1_0 := gopurs_runtime.Apply(__local_var_0, gopurs_runtime.Value{})
_ = __local_var_1_0
return Get_Data_Unit_unit()
})
}

func Call_Data_Array_fromJust(v_0_loop *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var v_0 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]] = v_0_loop
_ = v_0
var __t0 *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]
{
if (v_0 != nil) {
__t0 = (v_0).V0
goto end_branch_0
} else {

}
}
{
__t0 = Rebox_Data_Array_138441832_1728839155(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](func() gopurs_runtime.Value { panic("Failed pattern match") }()))
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(__t0))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
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

func Call_Data_Array_zipWithA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value, ys_3_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
var ys_3 []gopurs_runtime.Value = ys_3_loop
_ = ys_3
return gopurs_runtime.Apply3(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_zipWithImpl(), f_1, gopurs_runtime.Array(xs_2), gopurs_runtime.Array(ys_3)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}

func Call_Data_Array_updateAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], us_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var us_1 gopurs_runtime.Value = us_1_loop
_ = us_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Func(func(res_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_Foldable_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Get_Data_Array_ST_poke(), gopurs_runtime.Int((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0.IntVal), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1, res_3)
}), us_1)
}), gopurs_runtime.Array(xs_2))).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_updateAt(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_unsafeIndex(_dollar___unused_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value, __local_var_2_loop int64) gopurs_runtime.Value {
var _dollar___unused_0 gopurs_runtime.Value = _dollar___unused_0_loop
_ = _dollar___unused_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 int64 = __local_var_2_loop
_ = __local_var_2
return __local_var_1[__local_var_2]
}

func Call_Data_Array_uncons(__local_var_0_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
}), gopurs_runtime.Func2(func(x_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2762242827_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail []gopurs_runtime.Value
}]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.RecordDict2("head", "tail", x_1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(xs_2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}), gopurs_runtime.Array(__local_var_0))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_toUnfoldable(dictUnfoldable_0_loop *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value], xs_1_loop []gopurs_runtime.Value) gopurs_runtime.Value {
var dictUnfoldable_0 *Constructor_Data_Unfoldable_Unfoldable[gopurs_runtime.Value] = dictUnfoldable_0_loop
_ = dictUnfoldable_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): len_2_0 shape=Other bindingType=Int
var len_2_0 int64 = gopurs_runtime.Int(int64(len(xs_1))).IntVal
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictUnfoldable_0.V1), gopurs_runtime.Func(func(i_3 gopurs_runtime.Value) gopurs_runtime.Value {
var __t1 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]
{
if (i_3.IntVal) < (len_2_0) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer((&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{1, xs_1[i_3.IntVal], gopurs_runtime.Int((i_3.IntVal) + (int64(1)))}))}, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = Rebox_Data_Array_3094389156_1413047506(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))
}
end_branch_1:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1413047506_3094389156(__t1))}
}), gopurs_runtime.Int(int64(0)))
}

func Call_Data_Array_tail(__local_var_0_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_1 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
}), gopurs_runtime.Func2(func(v_1 gopurs_runtime.Value, xs_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{xs_2, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())))}
}), gopurs_runtime.Array(__local_var_0))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_sortBy(comp_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sortByImpl(), comp_0, gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 int64
{
var __t_tag_0 uint32 = uint32(v_2.IntVal)
if (uint32(__t_tag_0) == 380165415) {
__t3 = int64(1)
goto end_branch_3
} else {

}
}
{
var __t_tag_1 uint32 = uint32(v_2.IntVal)
if (uint32(__t_tag_1) == 902936544) {
__t3 = int64(0)
goto end_branch_3
} else {

}
}
{
var __t_tag_2 uint32 = uint32(v_2.IntVal)
if (uint32(__t_tag_2) == 1527465420) {
__t3 = int64(-1)
goto end_branch_3
} else {

}
}
{
__t3 = func() int64 { panic("Failed pattern match") }()
}
end_branch_3:
return gopurs_runtime.Int(__t3)
}), gopurs_runtime.Array(__local_var_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_sortWith(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value], f_1_loop gopurs_runtime.Value) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
return gopurs_runtime.Apply(Get_Data_Array_sortBy(), gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(gopurs_runtime.Box(dictOrd_0.V1), gopurs_runtime.Apply(f_1, x_2), gopurs_runtime.Apply(f_1, y_3)).IntVal)), UnsafePtr: nil}
}))
}

func Call_Data_Array_sort(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
// TAST (Let): compare_1_0 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] (ADT ["Data","Ordering","Ordering"] []))
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

func Call_Data_Array_splitAt(v_0_loop int64, v1_1_loop []gopurs_runtime.Value) struct{
	after []gopurs_runtime.Value
	before []gopurs_runtime.Value
} {
var v_0 int64 = v_0_loop
_ = v_0
var v1_1 []gopurs_runtime.Value = v1_1_loop
_ = v1_1
var __t0 struct{
	after []gopurs_runtime.Value
	before []gopurs_runtime.Value
}
{
if (v_0) <= (int64(0)) {
__t0 = struct{
	after []gopurs_runtime.Value
	before []gopurs_runtime.Value
}{v1_1, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()}
goto end_branch_0
} else {

}
}
{
__t0 = struct{
	after []gopurs_runtime.Value
	before []gopurs_runtime.Value
}{func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(v_0), gopurs_runtime.Int(int64(len(v1_1))), gopurs_runtime.Array(v1_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(v_0), gopurs_runtime.Array(v1_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()}
}
end_branch_0:
return __t0
}

func Call_Data_Array_take(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t0 []gopurs_runtime.Value
{
if (n_0) < (int64(1)) {
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(n_0), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
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

func Call_Data_Array_partition(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) struct{
	no []gopurs_runtime.Value
	yes []gopurs_runtime.Value
} {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() struct{
	no []gopurs_runtime.Value
	yes []gopurs_runtime.Value
} {
					orig := gopurs_runtime.UncurriedApp2(Get_Data_Array_partitionImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1))
					_ = orig
					clone := struct{
	no []gopurs_runtime.Value
	yes []gopurs_runtime.Value
}{}
					clone.no = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "no").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					clone.yes = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.RecordGet(orig, "yes").UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
					return clone
				}()
}

func Call_Data_Array_null(xs_0_loop []gopurs_runtime.Value) bool {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (int64(0))
}

func Call_Data_Array_modifyAtIndices(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value], is_1_loop gopurs_runtime.Value, f_2_loop gopurs_runtime.Value, xs_3_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
var is_1 gopurs_runtime.Value = is_1_loop
_ = is_1
var f_2 gopurs_runtime.Value = f_2_loop
_ = f_2
var xs_3 []gopurs_runtime.Value = xs_3_loop
_ = xs_3
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Func(func(res_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply4(Get_Data_Foldable_traverse_(), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](Get_Control_Monad_ST_Internal_applicativeST()))}, gopurs_runtime.Value{Type: 9, IntVal: 4280266298, UnsafePtr: unsafe.Pointer(dictFoldable_0)}, gopurs_runtime.Func(func(i_5 gopurs_runtime.Value) gopurs_runtime.Value {
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
// TAST (Let): v_2_0 shape=Other bindingType=Int
v_2_0 := gopurs_runtime.Int(int64(len(arr_1))).IntVal
_ = v_2_0
var __t8 []gopurs_runtime.Value
{
if (v_2_0) < (int64(2)) {
__t8 = arr_1
goto end_branch_8
} else {

}
}
{
__t8 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
out_3_1 := gopurs_runtime.Apply(Get_Data_Array_ST_newImpl(), gopurs_runtime.Value{})
_ = out_3_1
_dollar___unused_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_push(), arr_1[int64(0)], out_3_1), gopurs_runtime.Value{})
_ = _dollar___unused_4_3
_dollar___unused_5_4 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Get_Control_Monad_ST_Internal_forImpl(), gopurs_runtime.Int(int64(1)), gopurs_runtime.Int(v_2_0), gopurs_runtime.Func(func(idx_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_5 shape=App(Var) bindingType=Any
__local_var_6_5 := gopurs_runtime.Apply2(Get_Data_Array_ST_push(), a_0, out_3_1)
_ = __local_var_6_5
_dollar___unused_7_6 := gopurs_runtime.Apply(__local_var_6_5, gopurs_runtime.Value{})
_ = _dollar___unused_7_6
__local_var_8_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_push(), arr_1[idx_5.IntVal], out_3_1), gopurs_runtime.Value{})
_ = __local_var_8_7
return Get_Data_Unit_unit()
})
})), gopurs_runtime.Value{})
_ = _dollar___unused_5_4
__local_var_4_2 := out_3_1
_ = __local_var_4_2
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Array_ST_unsafeFreeze(), __local_var_4_2), gopurs_runtime.Value{})
})).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_8:
return __t8
}

func Call_Data_Array_intercalate(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
// TAST (Let): Semigroup0_1_0 shape=App(Other) bindingType=(ADT ["Data","Semigroup","Semigroup"] [(TypeVar m)])
Semigroup0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Semigroup_Semigroup[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonoid_0.V0), gopurs_runtime.Value{}))
_ = Semigroup0_1_0
// TAST (Let): mempty_2_1 shape=Other bindingType=(TypeVar m)
mempty_2_1 := gopurs_runtime.Box(dictMonoid_0.V1)
_ = mempty_2_1
return gopurs_runtime.Func2(func(sep_3 gopurs_runtime.Value, xs_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.RecordGet(func() gopurs_runtime.Value {
arr_val_foldlArray4 := xs_4
_ = arr_val_foldlArray4
res_go_foldlArray4 := func() gopurs_runtime.Value {
				orig := struct{
	acc gopurs_runtime.Value
	go__init bool
}{mempty_2_1, true}
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "init", orig.acc, gopurs_runtime.Bool(orig.go__init))
				}()
_ = res_go_foldlArray4
arr_go_foldlArray4 := (*[]gopurs_runtime.Value)(arr_val_foldlArray4.UnsafePtr)
_ = arr_go_foldlArray4
for _, v_foldlArray4 := range *arr_go_foldlArray4 {
res_go_foldlArray4 = gopurs_runtime.Apply2(gopurs_runtime.Func2(func(v_5 gopurs_runtime.Value, v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
var __t2 struct{
	acc gopurs_runtime.Value
	go__init bool
}
{
if (gopurs_runtime.RecordGet(v_5, "init").IntVal) != (0) {
__t2 = struct{
	acc gopurs_runtime.Value
	go__init bool
}{v1_6, false}
goto end_branch_2
} else {

}
}
{
__t2 = struct{
	acc gopurs_runtime.Value
	go__init bool
}{gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), gopurs_runtime.RecordGet(v_5, "acc"), gopurs_runtime.Apply2(gopurs_runtime.Box(Semigroup0_1_0.V0), sep_3, v1_6)), false}
}
end_branch_2:
return func() gopurs_runtime.Value {
				orig := __t2
				_ = orig
				return gopurs_runtime.RecordDict2("acc", "init", orig.acc, gopurs_runtime.Bool(orig.go__init))
				}()
}), res_go_foldlArray4, v_foldlArray4)
}
return res_go_foldlArray4
}(), "acc")
})
}

func Call_Data_Array_insertAt(__local_var_0_loop int64, __local_var_1_loop gopurs_runtime.Value, __local_var_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
var __local_var_2 []gopurs_runtime.Value = __local_var_2_loop
_ = __local_var_2
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(__local_var_0), __local_var_1, gopurs_runtime.Array(__local_var_2))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_go__init(xs_0_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t0 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (int64(0)) {
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_0
} else {

}
}
{
__t0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (int64(1))), gopurs_runtime.Array(xs_0)), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_0:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(__t0))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_index(__local_var_0_loop []gopurs_runtime.Value, __local_var_1_loop int64) struct{V0 gopurs_runtime.Value; V1 bool} {
var __local_var_0 []gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 int64 = __local_var_1_loop
_ = __local_var_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Array(__local_var_0), gopurs_runtime.Int(__local_var_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_last(xs_0_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (int64(1))))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_unsnoc(xs_0_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
var __t2 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (gopurs_runtime.Int(int64(len(xs_0))).IntVal) == (int64(0)) {
__t2 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_2
} else {

}
}
{
// TAST (Let): __local_var_1_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (int64(1)))))
_ = __local_var_1_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_1_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{func() gopurs_runtime.Value {
				orig := struct{
	go__init []gopurs_runtime.Value
	last gopurs_runtime.Value
}{func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int((gopurs_runtime.Int(int64(len(xs_0))).IntVal) - (int64(1))), gopurs_runtime.Array(xs_0)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), (__local_var_1_0).V0}
				_ = orig
				return gopurs_runtime.RecordDict2("init", "last", gopurs_runtime.Array(orig.go__init), orig.last)
				}(), true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
__t2 = __t1
}
end_branch_2:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t2)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_modifyAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0)))
_ = __local_var_3_0
var __t1 gopurs_runtime.Value
{
if (__local_var_3_0 == nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_1
} else {

}
}
{
if (__local_var_3_0 != nil) {
__t1 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(i_0), gopurs_runtime.Apply(f_1, (__local_var_3_0).V0), gopurs_runtime.Array(xs_2))))))}
goto end_branch_1
} else {

}
}
{
__t1 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t1))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_unzip(xs_0_loop []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
var xs_0 []*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] = xs_0_loop
_ = xs_0
return func() struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1807945000_138441832(Rebox_Data_Array_138441832_1807945000(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
fsts_1_0 := gopurs_runtime.Apply(Get_Data_Array_ST_newImpl(), gopurs_runtime.Value{})
_ = fsts_1_0
snds_2_1 := gopurs_runtime.Apply(Get_Data_Array_ST_newImpl(), gopurs_runtime.Value{})
_ = snds_2_1
iter_3_2 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_4010058633_3094389156(Rebox_Data_Array_3094389156_4010058633(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(v)} }
					return gopurs_runtime.Array(boxed)
				}(), v_3)))))}
})), gopurs_runtime.Value{})
_ = iter_3_2
_dollar___unused_4_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_Iterator_iterate(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_3988677819_394862806(Rebox_Data_Array_394862806_3988677819(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3_2))))}, gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_5_4 shape=Other bindingType=Any
__local_var_5_4 := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1
_ = __local_var_5_4
__local_var_6_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_push(), (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, fsts_1_0), gopurs_runtime.Value{})
_ = __local_var_6_6
_dollar___unused_6_5 := Get_Data_Unit_unit()
_ = _dollar___unused_6_5
__local_var_7_7 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_push(), __local_var_5_4, snds_2_1), gopurs_runtime.Value{})
_ = __local_var_7_7
return Get_Data_Unit_unit()
})
})), gopurs_runtime.Value{})
_ = _dollar___unused_4_3
fsts_prime__5_8 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Array_ST_unsafeFreeze(), fsts_1_0), gopurs_runtime.Value{})
_ = fsts_prime__5_8
snds_prime__6_9 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Array_ST_unsafeFreeze(), snds_2_1), gopurs_runtime.Value{})
_ = snds_prime__6_9
return gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1807945000_138441832(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, []gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{fsts_prime__5_8, snds_prime__6_9}
				return gopurs_runtime.Box(&Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{V0: _v.V0, V1: _v.V1})
			}())))}
}))))))}
				_p := (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(_v.UnsafePtr)
				return struct{V0 gopurs_runtime.Value; V1 gopurs_runtime.Value}{V0: _p.V0, V1: _p.V1}
			}()
}

func Call_Data_Array_head(xs_0_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var xs_0 []gopurs_runtime.Value = xs_0_loop
_ = xs_0
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Array(xs_0), gopurs_runtime.Int(int64(0)))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_nubBy(comp_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var comp_0 gopurs_runtime.Value = comp_0_loop
_ = comp_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): indexedAndSorted_2_0 shape=App(Var) bindingType=(Array (ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a)]))
indexedAndSorted_2_0 := func() []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_sortBy(gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: int64(uint32(gopurs_runtime.Apply2(comp_0, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(x_2.UnsafePtr).V1, (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(y_3.UnsafePtr).V1).IntVal)), UnsafePtr: nil}
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_FunctorWithIndex_mapWithIndexArray(), Get_Data_Tuple_Tuple(), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))} }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
_ = indexedAndSorted_2_0
// TAST (Let): v_3_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a)])])
var v_3_1 *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, func() gopurs_runtime.Value {
					arr := indexedAndSorted_2_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Int(int64(0)))))})
var __t10 []gopurs_runtime.Value
{
if (v_3_1 == nil) {
__t10 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_10
} else {

}
}
{
if (v_3_1 != nil) {
__t10 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
arr_val_arrayMap3 := func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Call_Data_Array_sortWith(Rebox_Data_Array_3308271157_4177771502(gopurs_runtime.CoerceToStruct[Constructor_Data_Ord_Ord[int64]](Get_Data_Ord_ordInt())), Get_Data_Tuple_fst()), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=(TypeVar b)
__local_var_4_2 := gopurs_runtime.Apply(Get_Data_Array_ST_unsafeThaw(), func() gopurs_runtime.Value {
					arr := func() []*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832((v_3_1).V0))}}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))} }
					return gopurs_runtime.Array(boxed)
				}())
_ = __local_var_4_2
result_5_3 := gopurs_runtime.Apply(__local_var_4_2, gopurs_runtime.Value{})
_ = result_5_3
_dollar___unused_6_4 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_foreach(), func() gopurs_runtime.Value {
					arr := indexedAndSorted_2_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))} }
					return gopurs_runtime.Array(boxed)
				}(), gopurs_runtime.Func(func(v1_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_7_5 shape=App(Var) bindingType=(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (Array (ADT ["Data","Tuple","Tuple"] [Int, (TypeVar a)]))])
__local_var_7_5 := gopurs_runtime.Apply(Get_Data_Array_ST_unsafeFreeze(), result_5_3)
_ = __local_var_7_5
__local_var_8_7 := gopurs_runtime.Apply(__local_var_7_5, gopurs_runtime.Value{})
_ = __local_var_8_7
// TAST (Let): __local_var_9_8 shape=UncurriedApp(Var) bindingType=(TypeVar c)
__local_var_9_8 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, __local_var_8_7, gopurs_runtime.Int((gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(__local_var_8_7))).IntVal) - (int64(1)))))
_ = __local_var_9_8
var __t9 gopurs_runtime.Value
{
if (__local_var_9_8 != nil) {
__t9 = (*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])((__local_var_9_8).V0.UnsafePtr).V1
goto end_branch_9
} else {

}
}
{
__t9 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_9:
lst_8_6 := __t9
_ = lst_8_6
return Get_Data_Unit_unit()
})
})), gopurs_runtime.Value{})
_ = _dollar___unused_6_4
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Array_ST_unsafeFreeze(), result_5_3), gopurs_runtime.Value{})
})).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))} }
					return gopurs_runtime.Array(boxed)
				}()).UnsafePtr)
					unboxed := make([]*Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value], len(arr))
					for i, v := range arr { unboxed[i] = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]](v) }
					return unboxed
				}()
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_1728839155_138441832(v))} }
					return gopurs_runtime.Array(boxed)
				}()
_ = arr_val_arrayMap3
arr_go_arrayMap3 := (*[]gopurs_runtime.Value)(arr_val_arrayMap3.UnsafePtr)
_ = arr_go_arrayMap3
res_go_arrayMap3 := make([]gopurs_runtime.Value, len(*arr_go_arrayMap3))
_ = res_go_arrayMap3
for i_arrayMap3, v_arrayMap3 := range *arr_go_arrayMap3 {
res_go_arrayMap3[i_arrayMap3] = gopurs_runtime.Apply(Get_Data_Tuple_snd(), v_arrayMap3)
}
return gopurs_runtime.Array(res_go_arrayMap3)
}().UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_10
} else {

}
}
{
__t10 = func() []gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_10:
return __t10
}

func Call_Data_Array_nub(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Array_nubBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_Array_groupBy(op_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) [][]gopurs_runtime.Value {
var op_0 gopurs_runtime.Value = op_0_loop
_ = op_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
result_2_0 := gopurs_runtime.Apply(Get_Data_Array_ST_newImpl(), gopurs_runtime.Value{})
_ = result_2_0
iter_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Array_ST_Iterator_iterator(), gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Array(xs_1), v_3)))}
})), gopurs_runtime.Value{})
_ = iter_3_1
_dollar___unused_4_2 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_Iterator_iterate(), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3_1))}, gopurs_runtime.Func(func(x_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
sub_5_4 := gopurs_runtime.Apply(Get_Data_Array_ST_newImpl(), gopurs_runtime.Value{})
_ = sub_5_4
_dollar___unused_6_5 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_push(), x_4, sub_5_4), gopurs_runtime.Value{})
_ = _dollar___unused_6_5
_dollar___unused_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply3(Get_Data_Array_ST_Iterator_pushWhile(), gopurs_runtime.Apply(op_0, x_4), gopurs_runtime.Value{Type: 9, IntVal: 3127013252, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]](iter_3_1))}, sub_5_4), gopurs_runtime.Value{})
_ = _dollar___unused_7_6
grp_8_7 := gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Array_ST_unsafeFreeze(), sub_5_4), gopurs_runtime.Value{})
_ = grp_8_7
__local_var_5_3 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_push(), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(grp_8_7.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()), result_2_0), gopurs_runtime.Value{})
_ = __local_var_5_3
return Get_Data_Unit_unit()
})
})), gopurs_runtime.Value{})
_ = _dollar___unused_4_2
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Array_ST_unsafeFreeze(), result_2_0), gopurs_runtime.Value{})
})).UnsafePtr)
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

func Call_Data_Array_groupAllBy(cmp_0_loop gopurs_runtime.Value, x_1_loop []gopurs_runtime.Value) [][]gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 []gopurs_runtime.Value = x_1_loop
_ = x_1
return func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(func() gopurs_runtime.Value {
					arr := Call_Data_Array_groupBy(gopurs_runtime.Func2(func(x_2 gopurs_runtime.Value, y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(Call_Data_Array_sortBy(cmp_0, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(x_1).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}().UnsafePtr)
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

func Call_Data_Array_groupAll(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Array_groupAllBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_Array_group(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
// TAST (Let): eq_1_0 shape=Other bindingType=(Func [(TypeVar a), (TypeVar a)] Boolean)
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

func Call_Data_Array_fromFoldable(dictFoldable_0_loop *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictFoldable_0 *Constructor_Data_Foldable_Foldable[gopurs_runtime.Value] = dictFoldable_0_loop
_ = dictFoldable_0
// TAST (Let): __local_var_1_0 shape=Other bindingType=(Func [(Func [(TypeVar a), (TypeVar b)] (TypeVar b)), (TypeVar b), (TypeApp (TypeVar f) [(TypeVar a)])] (TypeVar b))
__local_var_1_0 := gopurs_runtime.Box(dictFoldable_0.V2)
_ = __local_var_1_0
return gopurs_runtime.Func(func(__local_var_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp2(Get_Data_Array_fromFoldableImpl(), __local_var_1_0, __local_var_2)
})
}

func Call_Data_Array_foldl__1894943466(__eta_norm_2_unused_0_loop gopurs_runtime.Value, __eta_norm_1_1_loop int64, __eta_norm_0_2_loop []int64) int64 {
foldl__1894943466:
for {
if false { continue foldl__1894943466 }
var __eta_norm_2_unused_0 gopurs_runtime.Value = __eta_norm_2_unused_0_loop
_ = __eta_norm_2_unused_0
var __eta_norm_1_1 int64 = __eta_norm_1_1_loop
_ = __eta_norm_1_1
var __eta_norm_0_2 []int64 = __eta_norm_0_2_loop
_ = __eta_norm_0_2
return func() gopurs_runtime.Value {
arr_val_foldlArray0 := func() gopurs_runtime.Value {
					arr := __eta_norm_0_2
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
_ = arr_val_foldlArray0
res_go_foldlArray0 := gopurs_runtime.Int(__eta_norm_1_1)
_ = res_go_foldlArray0
arr_go_foldlArray0 := (*[]gopurs_runtime.Value)(arr_val_foldlArray0.UnsafePtr)
_ = arr_go_foldlArray0
for _, v_foldlArray0 := range *arr_go_foldlArray0 {
res_go_foldlArray0 = gopurs_runtime.Apply2(Get_Data_Semiring_intAdd(), res_go_foldlArray0, v_foldlArray0)
}
return res_go_foldlArray0
}().IntVal
}
}

func Call_Data_Array_transpose(xs_0_loop [][]gopurs_runtime.Value) [][]gopurs_runtime.Value {
var xs_0 [][]gopurs_runtime.Value = xs_0_loop
_ = xs_0
var Call_local_Data_Array_go__go_1_0_0 func(int64, [][]gopurs_runtime.Value) [][]gopurs_runtime.Value
_ = Call_local_Data_Array_go__go_1_0_0
var go__go_1_0_0 gopurs_runtime.Value
_ = go__go_1_0_0
Call_local_Data_Array_go__go_1_0_0 = func(idx_2_loop int64, allArrays_3_loop [][]gopurs_runtime.Value) [][]gopurs_runtime.Value {
go__go_1_0_0:
for {
if false { continue go__go_1_0_0 }
var idx_2 int64 = idx_2_loop
_ = idx_2
var allArrays_3 [][]gopurs_runtime.Value = allArrays_3_loop
_ = allArrays_3
// TAST (Let): v_4_1 shape=App(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Array (TypeVar a))])
var v_4_1 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]](func() gopurs_runtime.Value {
arr_val_foldlArray2 := func() gopurs_runtime.Value {
					arr := xs_0
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Array(v) }
					return gopurs_runtime.Array(boxed)
				}()
_ = arr_val_foldlArray2
res_go_foldlArray2 := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}))))}
_ = res_go_foldlArray2
arr_go_foldlArray2 := (*[]gopurs_runtime.Value)(arr_val_foldlArray2.UnsafePtr)
_ = arr_go_foldlArray2
for _, v_foldlArray2 := range *arr_go_foldlArray2 {
res_go_foldlArray2 = gopurs_runtime.Apply2(gopurs_runtime.Func2(func(acc_4 gopurs_runtime.Value, nextArr_5 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_6_2 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_6_2 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, nextArr_5, gopurs_runtime.Int(idx_2)))
_ = __local_var_6_2
var __t6 gopurs_runtime.Value
{
if (__local_var_6_2 == nil) {
__t6 = acc_4
goto end_branch_6
} else {

}
}
{
if (__local_var_6_2 != nil) {
var __t5 gopurs_runtime.Value
{
var __t_tag_3 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] = Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](acc_4))
if (__t_tag_3 == nil) {
__t5 = gopurs_runtime.Array([]gopurs_runtime.Value{(__local_var_6_2).V0})
goto end_branch_5
} else {

}
}
{
var __t_tag_4 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] = Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](acc_4))
if (__t_tag_4 != nil) {
__t5 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), (__local_var_6_2).V0), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(acc_4.UnsafePtr).V0.UnsafePtr)
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
__t6 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{1, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t5.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())}))}
goto end_branch_6
} else {

}
}
{
__t6 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_6:
return gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t6))))}
}), res_go_foldlArray2, v_foldlArray2)
}
return res_go_foldlArray2
}())
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
idx_2_loop = (idx_2) + (int64(1))
allArrays_3_loop = func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Apply2(Get_Data_Array_ST_withArray(), gopurs_runtime.Apply(Get_Data_Array_ST_push(), gopurs_runtime.Array((v_4_1).V0)), gopurs_runtime.Array(func() []gopurs_runtime.Value {
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
__t7 = func() [][]gopurs_runtime.Value { panic("unreachable") }()
goto end_branch_7
} else {

}
}
{
__t7 = func() [][]gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_7:
return __t7
}
}
go__go_1_0_0 = gopurs_runtime.Func(func(idx_2_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(allArrays_3_loop_val gopurs_runtime.Value) gopurs_runtime.Value {
return func() gopurs_runtime.Value {
					arr := Call_local_Data_Array_go__go_1_0_0(idx_2_loop_val.IntVal, func() [][]gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(allArrays_3_loop_val.UnsafePtr)
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
return Call_local_Data_Array_go__go_1_0_0(int64(0), [][]gopurs_runtime.Value{})
}

func Call_Data_Array_foldRecM(dictMonadRec_0_loop *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonadRec_0 *Constructor_Control_Monad_Rec_Class_MonadRec[gopurs_runtime.Value] = dictMonadRec_0_loop
_ = dictMonadRec_0
// TAST (Let): Monad0_1_0 shape=App(Other) bindingType=Any
Monad0_1_0 := gopurs_runtime.Apply(gopurs_runtime.Box(dictMonadRec_0.V0), gopurs_runtime.Value{})
_ = Monad0_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Applicative0"), gopurs_runtime.Value{}))
_ = Applicative0_2_1
// TAST (Let): Bind1_3_2 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_3_2 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(Monad0_1_0, "Bind1"), gopurs_runtime.Value{}))
_ = Bind1_3_2
return gopurs_runtime.Func3(func(f_4 gopurs_runtime.Value, b_5 gopurs_runtime.Value, array_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(dictMonadRec_0.V1), gopurs_runtime.Func(func(o_7 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 gopurs_runtime.Value
{
if (gopurs_runtime.RecordGet(o_7, "b").IntVal) >= (gopurs_runtime.Int(int64(gopurs_runtime.ArrayLength(array_6))).IntVal) {
__t3 = gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 60402430, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2206451173_3603546092((&Constructor_Control_Monad_Rec_Class_Done[struct{
	a gopurs_runtime.Value
	b int64
}, gopurs_runtime.Value]{1, gopurs_runtime.RecordGet(o_7, "a")})))})
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_3_2.V1), gopurs_runtime.Apply2(f_4, gopurs_runtime.RecordGet(o_7, "a"), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(array_6.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()[gopurs_runtime.RecordGet(o_7, "b").IntVal]), gopurs_runtime.Func(func(res_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Value{Type: 9, IntVal: 525585346, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2660862937_4008603408((&Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b int64
}, gopurs_runtime.Value]{1, func() struct{
	a gopurs_runtime.Value
	b int64
} {
					orig := gopurs_runtime.RecordDict2("a", "b", res_prime__8, gopurs_runtime.Int((gopurs_runtime.RecordGet(o_7, "b").IntVal) + (int64(1))))
					_ = orig
					clone := struct{
	a gopurs_runtime.Value
	b int64
}{}
					clone.a = gopurs_runtime.RecordGet(orig, "a")
					clone.b = gopurs_runtime.RecordGet(orig, "b").IntVal
					return clone
				}()})))})
}))
}
end_branch_3:
return __t3
}), func() gopurs_runtime.Value {
				orig := struct{
	a gopurs_runtime.Value
	b gopurs_runtime.Value
}{b_5, gopurs_runtime.Int(int64(0))}
				_ = orig
				return gopurs_runtime.RecordDict2("a", "b", orig.a, orig.b)
				}())
})
}

func Call_Data_Array_foldMap(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Apply(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_0)})
}

func Call_Data_Array_foldM(dictMonad_0_loop *Constructor_Control_Monad_Monad[gopurs_runtime.Value]) gopurs_runtime.Value {
foldM:
for {
if false { continue foldM }
var dictMonad_0 *Constructor_Control_Monad_Monad[gopurs_runtime.Value] = dictMonad_0_loop
_ = dictMonad_0
// TAST (Let): Applicative0_1_0 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar m)])
Applicative0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_1_0
// TAST (Let): Bind1_2_1 shape=App(Other) bindingType=(ADT ["Control","Bind","Bind"] [(TypeVar m)])
Bind1_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Bind_Bind[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictMonad_0.V1), gopurs_runtime.Value{}))
_ = Bind1_2_1
return gopurs_runtime.Func3(func(f_3 gopurs_runtime.Value, b_4 gopurs_runtime.Value, __local_var_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.UncurriedApp3(Get_Data_Array_unconsImpl(), gopurs_runtime.Func(func(v_6 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_1_0.V1), b_4)
}), gopurs_runtime.Func2(func(a_6 gopurs_runtime.Value, as_7 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Bind1_2_1.V1), gopurs_runtime.Apply2(f_3, b_4, a_6), gopurs_runtime.Func(func(b_prime__8 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply3(Call_Data_Array_foldM(dictMonad_0), f_3, b_prime__8, gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(as_7.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}))
}), __local_var_5)
})
}
}

func Call_Data_Array_fold(dictMonoid_0_loop *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictMonoid_0 *Constructor_Data_Monoid_Monoid[gopurs_runtime.Value] = dictMonoid_0_loop
_ = dictMonoid_0
return gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Foldable_Foldable[gopurs_runtime.Value]](Get_Data_Foldable_foldableArray()).V0), gopurs_runtime.Value{Type: 9, IntVal: 1722653594, UnsafePtr: unsafe.Pointer(dictMonoid_0)}, gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Control_Category_Category[gopurs_runtime.Value]](Get_Control_Category_categoryFn()).V1))
}

func Call_Data_Array_findMap(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp4(Get_Data_Array_findMapImpl(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, Get_Data_Maybe_isJust(), __local_var_0, gopurs_runtime.Array(__local_var_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_findLastIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, __local_var_0, gopurs_runtime.Array(__local_var_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_insertBy(cmp_0_loop gopurs_runtime.Value, x_1_loop gopurs_runtime.Value, ys_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var cmp_0 gopurs_runtime.Value = cmp_0_loop
_ = cmp_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
var ys_2 []gopurs_runtime.Value = ys_2_loop
_ = ys_2
// TAST (Let): __local_var_3_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_3_1 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findLastIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Func(func(y_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(true)
}), gopurs_runtime.Array(ys_2))))
_ = __local_var_3_1
var __t2 int64
{
if (__local_var_3_1 == nil) {
__t2 = int64(0)
goto end_branch_2
} else {

}
}
{
if (__local_var_3_1 != nil) {
__t2 = ((__local_var_3_1).V0) + (int64(1))
goto end_branch_2
} else {

}
}
{
__t2 = func() int64 { panic("Failed pattern match") }()
}
end_branch_2:
// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Array (TypeVar a))])
__local_var_3_0 := Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get_Data_Array__insertAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(__t2), x_1, gopurs_runtime.Array(ys_2))))
_ = __local_var_3_0
var __t3 gopurs_runtime.Value
{
if (__local_var_3_0 != nil) {
__t3 = gopurs_runtime.Array((__local_var_3_0).V0)
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_insert(dictOrd_0_loop *Constructor_Data_Ord_Ord[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictOrd_0 *Constructor_Data_Ord_Ord[gopurs_runtime.Value] = dictOrd_0_loop
_ = dictOrd_0
return gopurs_runtime.Apply(Get_Data_Array_insertBy(), gopurs_runtime.Box(dictOrd_0.V1))
}

func Call_Data_Array_findIndex(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, __local_var_0, gopurs_runtime.Array(__local_var_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_span(p_0_loop gopurs_runtime.Value, arr_1_loop []gopurs_runtime.Value) struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
} {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var arr_1 []gopurs_runtime.Value = arr_1_loop
_ = arr_1
// TAST (Let): breakIndex_2_0 shape=UncurriedApp(Var) bindingType=Any
breakIndex_2_0 := gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Func(func(x_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool(((gopurs_runtime.Apply(p_0, x_2).IntVal) != (0)) != (true))
}), gopurs_runtime.Array(arr_1))
_ = breakIndex_2_0
var __t2 struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
}
{
if (breakIndex_2_0.Type == 9 && breakIndex_2_0.IntVal == 930809136 && breakIndex_2_0.UnsafePtr != nil) {
var __t1 struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
}
{
if ((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(breakIndex_2_0.UnsafePtr).V0.IntVal) == (int64(0)) {
__t1 = struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
}{func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), arr_1}
goto end_branch_1
} else {

}
}
{
__t1 = struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
}{func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(breakIndex_2_0.UnsafePtr).V0, gopurs_runtime.Array(arr_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}(), func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(breakIndex_2_0.UnsafePtr).V0, gopurs_runtime.Int(int64(len(arr_1))), gopurs_runtime.Array(arr_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()}
}
end_branch_1:
__t2 = __t1
goto end_branch_2
} else {

}
}
{
if (breakIndex_2_0.Type == 9 && breakIndex_2_0.IntVal == 930809136 && breakIndex_2_0.UnsafePtr == nil) {
__t2 = struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
}{arr_1, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()}
goto end_branch_2
} else {

}
}
{
__t2 = func() struct{
	go__init []gopurs_runtime.Value
	rest []gopurs_runtime.Value
} { panic("Failed pattern match") }()
}
end_branch_2:
return __t2
}

func Call_Data_Array_takeWhile(p_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var p_0 gopurs_runtime.Value = p_0_loop
_ = p_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return Call_Data_Array_span(p_0, xs_1).go__init
}

func Call_Data_Array_find(f_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_2_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, f_0, gopurs_runtime.Array(xs_1))))
_ = __local_var_2_0
var __t1 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{xs_1[gopurs_runtime.Int((__local_var_2_0).V0).IntVal], true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_1
} else {

}
}
{
__t1 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
}
end_branch_1:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(__t1)}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
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

func Call_Data_Array_filter__1859503602(__eta_norm_1_0_loop gopurs_runtime.Value, __eta_norm_0_1_loop []int64) []int64 {
filter__1859503602:
for {
if false { continue filter__1859503602 }
var __eta_norm_1_0 gopurs_runtime.Value = __eta_norm_1_0_loop
_ = __eta_norm_1_0
var __eta_norm_0_1 []int64 = __eta_norm_0_1_loop
_ = __eta_norm_0_1
return func() []int64 {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array(func() []gopurs_runtime.Value {
arr_val_filterImpl0 := func() gopurs_runtime.Value {
					arr := __eta_norm_0_1
					boxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { boxed[i] = gopurs_runtime.Int(v) }
					return gopurs_runtime.Array(boxed)
				}()
_ = arr_val_filterImpl0
_ = arr_val_filterImpl0
arr_go_filterImpl0 := (*[]gopurs_runtime.Value)(arr_val_filterImpl0.UnsafePtr)
_ = arr_go_filterImpl0
res_go_filterImpl0 := make([]gopurs_runtime.Value, 0)
_ = res_go_filterImpl0
for _, v_filterImpl0 := range *arr_go_filterImpl0 {
if gopurs_runtime.Apply(__eta_norm_1_0, v_filterImpl0).BoolVal() {
res_go_filterImpl0 = append(res_go_filterImpl0, v_filterImpl0)
} else {

}
}
return res_go_filterImpl0
}()).UnsafePtr)
					unboxed := make([]int64, len(arr))
					for i, v := range arr { unboxed[i] = v.IntVal }
					return unboxed
				}()
}
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
// TAST (Let): __local_var_4_0 shape=UncurriedApp(Var) bindingType=Any
__local_var_4_0 := gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Apply(eq_0, x_3), gopurs_runtime.Array(ys_2))
_ = __local_var_4_0
var __t1 gopurs_runtime.Value
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 930809136 && __local_var_4_0.UnsafePtr == nil) {
__t1 = gopurs_runtime.Bool(false)
goto end_branch_1
} else {

}
}
{
if (__local_var_4_0.Type == 9 && __local_var_4_0.IntVal == 930809136 && __local_var_4_0.UnsafePtr != nil) {
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

func Call_Data_Array_intersect(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_intersectBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_elemLastIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Get_Data_Array_findLastIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_elemIndex(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], x_1_loop gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var x_1 gopurs_runtime.Value = x_1_loop
_ = x_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Apply(Get_Data_Array_findIndex(), gopurs_runtime.Func(func(v_2 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_2, x_1).IntVal) != (0))
}))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_notElem(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_3_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, a_1).IntVal) != (0))
}), gopurs_runtime.Array(arr_2))))
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

func Call_Data_Array_elem(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value], a_1_loop gopurs_runtime.Value, arr_2_loop []gopurs_runtime.Value) bool {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
var a_1 gopurs_runtime.Value = a_1_loop
_ = a_1
var arr_2 []gopurs_runtime.Value = arr_2_loop
_ = arr_2
// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_3_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Func(func(v_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(gopurs_runtime.Box(dictEq_0.V0), v_3, a_1).IntVal) != (0))
}), gopurs_runtime.Array(arr_2))))
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
return Call_Data_Array_span(p_0, xs_1).rest
}

func Call_Data_Array_dropEnd(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 shape=Other bindingType=Int
__local_var_2_0 := (gopurs_runtime.Int(int64(len(xs_1))).IntVal) - (n_0)
_ = __local_var_2_0
var __t1 []gopurs_runtime.Value
{
if (__local_var_2_0) < (int64(1)) {
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(int64(0)), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_1:
return __t1
}

func Call_Data_Array_drop(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
var __t0 []gopurs_runtime.Value
{
if (n_0) < (int64(1)) {
__t0 = xs_1
goto end_branch_0
} else {

}
}
{
__t0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(n_0), gopurs_runtime.Int(int64(len(xs_1))), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_0:
return __t0
}

func Call_Data_Array_takeEnd(n_0_loop int64, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var n_0 int64 = n_0_loop
_ = n_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
// TAST (Let): __local_var_2_0 shape=Other bindingType=Int
__local_var_2_0 := (gopurs_runtime.Int(int64(len(xs_1))).IntVal) - (n_0)
_ = __local_var_2_0
var __t1 []gopurs_runtime.Value
{
if (__local_var_2_0) < (int64(1)) {
__t1 = xs_1
goto end_branch_1
} else {

}
}
{
__t1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.UncurriedApp3(Get_Data_Array_sliceImpl(), gopurs_runtime.Int(__local_var_2_0), gopurs_runtime.Int(int64(len(xs_1))), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_1:
return __t1
}

func Call_Data_Array_deleteAt(__local_var_0_loop int64, __local_var_1_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var __local_var_0 int64 = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(__local_var_0), gopurs_runtime.Array(__local_var_1))
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_deleteBy(v_0_loop gopurs_runtime.Value, v1_1_loop gopurs_runtime.Value, v2_2_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var v_0 gopurs_runtime.Value = v_0_loop
_ = v_0
var v1_1 gopurs_runtime.Value = v1_1_loop
_ = v1_1
var v2_2 []gopurs_runtime.Value = v2_2_loop
_ = v2_2
var __t4 []gopurs_runtime.Value
{
if (gopurs_runtime.Int(int64(len(v2_2))).IntVal) == (int64(0)) {
__t4 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
goto end_branch_4
} else {

}
}
{
// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [Int])
__local_var_3_0 := Rebox_Data_Array_3094389156_1170268447(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_findIndexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Apply(v_0, v1_1), gopurs_runtime.Array(v2_2))))
_ = __local_var_3_0
var __t3 gopurs_runtime.Value
{
if (__local_var_3_0 == nil) {
__t3 = gopurs_runtime.Array(v2_2)
goto end_branch_3
} else {

}
}
{
if (__local_var_3_0 != nil) {
// TAST (Let): __local_var_4_1 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(Array (TypeVar a))])
__local_var_4_1 := Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int((__local_var_3_0).V0), gopurs_runtime.Array(v2_2))))
_ = __local_var_4_1
var __t2 gopurs_runtime.Value
{
if (__local_var_4_1 != nil) {
__t2 = gopurs_runtime.Array((__local_var_4_1).V0)
goto end_branch_2
} else {

}
}
{
__t2 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t2.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
__t4 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(__t3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}
end_branch_4:
return __t4
}

func Call_Data_Array_go__delete(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_deleteBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_difference(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Foldable_foldrArray(), gopurs_runtime.Apply(Get_Data_Array_go__delete(), gopurs_runtime.Value{Type: 9, IntVal: 1012063514, UnsafePtr: unsafe.Pointer(dictEq_0)}))
}

func Call_Data_Array_cons(x_0_loop gopurs_runtime.Value, xs_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var x_0 gopurs_runtime.Value = x_0_loop
_ = x_0
var xs_1 []gopurs_runtime.Value = xs_1_loop
_ = xs_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array([]gopurs_runtime.Value{x_0}), gopurs_runtime.Array(xs_1)).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_some(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Apply0_1_0 shape=App(Other) bindingType=(ADT ["Control","Apply","Apply"] [(TypeVar f)])
Apply0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Apply_Apply[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}), "Apply0"), gopurs_runtime.Value{}))
_ = Apply0_1_0
// TAST (Let): Functor0_2_1 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_2_1
return gopurs_runtime.Func2(func(dictLazy_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Apply0_1_0.V1), gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_2_1.V0), Get_Data_Array_cons(), v_4), gopurs_runtime.Apply(gopurs_runtime.RecordGet(dictLazy_3, "defer"), gopurs_runtime.Func(func(v1_5 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(Call_Data_Array_many(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4)
})))
})
}

func Call_Data_Array_many(dictAlternative_0_loop *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictAlternative_0 *Constructor_Control_Alternative_Alternative[gopurs_runtime.Value] = dictAlternative_0_loop
_ = dictAlternative_0
// TAST (Let): Alt0_1_0 shape=App(Other) bindingType=(ADT ["Control","Alt","Alt"] [(TypeVar f)])
Alt0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Control_Alt_Alt[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V1), gopurs_runtime.Value{}), "Alt0"), gopurs_runtime.Value{}))
_ = Alt0_1_0
// TAST (Let): Applicative0_2_1 shape=App(Other) bindingType=(ADT ["Control","Applicative","Applicative"] [(TypeVar f)])
Applicative0_2_1 := gopurs_runtime.CoerceToStruct[Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.Box(dictAlternative_0.V0), gopurs_runtime.Value{}))
_ = Applicative0_2_1
return gopurs_runtime.Func2(func(dictLazy_3 gopurs_runtime.Value, v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Alt0_1_0.V1), gopurs_runtime.Apply2(Call_Data_Array_some(dictAlternative_0), gopurs_runtime.Value{Type: 9, IntVal: 1860244333, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Control_Lazy_Lazy[gopurs_runtime.Value]](dictLazy_3))}, v_4), gopurs_runtime.Apply(gopurs_runtime.Box(Applicative0_2_1.V1), gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{}).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}())))
})
}

func Call_Data_Array_concatMap(b_0_loop gopurs_runtime.Value, a_1_loop []gopurs_runtime.Value) []gopurs_runtime.Value {
var b_0 gopurs_runtime.Value = b_0_loop
_ = b_0
var a_1 []gopurs_runtime.Value = a_1_loop
_ = a_1
return func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Control_Bind_arrayBind(), gopurs_runtime.Array(a_1), b_0).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_mapMaybe(f_0_loop gopurs_runtime.Value) gopurs_runtime.Value {
var f_0 gopurs_runtime.Value = f_0_loop
_ = f_0
return gopurs_runtime.Apply(Get_Data_Array_concatMap(), gopurs_runtime.Func(func(x_1 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_2_0 shape=App(Other) bindingType=(TypeVar c)
__local_var_2_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_0, x_1))
_ = __local_var_2_0
var __t1 gopurs_runtime.Value
{
if (__local_var_2_0 == nil) {
__t1 = gopurs_runtime.Array([]gopurs_runtime.Value{})
goto end_branch_1
} else {

}
}
{
if (__local_var_2_0 != nil) {
__t1 = gopurs_runtime.Array(func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Array([]gopurs_runtime.Value{(__local_var_2_0).V0}).UnsafePtr)
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

func Call_Data_Array_filterA(dictApplicative_0_loop *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictApplicative_0 *Constructor_Control_Applicative_Applicative[gopurs_runtime.Value] = dictApplicative_0_loop
_ = dictApplicative_0
// TAST (Let): Functor0_1_0 shape=App(Other) bindingType=(ADT ["Data","Functor","Functor"] [(TypeVar f)])
Functor0_1_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Functor_Functor[gopurs_runtime.Value]](gopurs_runtime.Apply(gopurs_runtime.RecordGet(gopurs_runtime.Apply(gopurs_runtime.Box(dictApplicative_0.V0), gopurs_runtime.Value{}), "Functor0"), gopurs_runtime.Value{}))
_ = Functor0_1_0
return gopurs_runtime.Func(func(p_2 gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_3_1 shape=App(Other) bindingType=(Func [(Array (TypeVar a))] (TypeApp (TypeVar f) [(Array (ADT ["Data","Tuple","Tuple"] [(TypeVar a), Boolean]))]))
__local_var_3_1 := gopurs_runtime.Apply2(gopurs_runtime.Box(gopurs_runtime.CoerceToStruct[Constructor_Data_Traversable_Traversable[gopurs_runtime.Value]](Get_Data_Traversable_traversableArray()).V3), gopurs_runtime.Value{Type: 9, IntVal: 1459134221, UnsafePtr: unsafe.Pointer(dictApplicative_0)}, gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Apply2(gopurs_runtime.Box(Functor0_1_0.V0), gopurs_runtime.Apply(Get_Data_Tuple_Tuple(), x_3), gopurs_runtime.Apply(p_2, x_3))
}))
_ = __local_var_3_1
// TAST (Let): __local_var_4_2 shape=App(Other) bindingType=(Func [(TypeApp (TypeVar f) [(Array (ADT ["Data","Tuple","Tuple"] [(TypeVar a), Boolean]))])] (TypeApp (TypeVar f) [(Array (TypeVar a))]))
__local_var_4_2 := gopurs_runtime.Apply(gopurs_runtime.Box(Functor0_1_0.V0), Call_Data_Array_mapMaybe(gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
var __t3 *Constructor_Data_Maybe_Just[gopurs_runtime.Value]
{
if ((*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V1.IntVal) != (0) {
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{(*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value])(v_4.UnsafePtr).V0, true}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
goto end_branch_3
} else {

}
}
{
__t3 = gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}())
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply(Get_Control_Monad_ST_Internal_run(), gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
arr_2_0 := gopurs_runtime.Apply(Get_Data_Array_ST_newImpl(), gopurs_runtime.Value{})
_ = arr_2_0
_dollar___unused_3_1 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Control_Monad_ST_Internal_foreach(), gopurs_runtime.Array(xs_1), gopurs_runtime.Func(func(x_3 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
// TAST (Let): __local_var_4_2 shape=App(Var) bindingType=(Func [(Array (TypeVar a))] Boolean)
__local_var_4_2 := gopurs_runtime.Apply(Get_Data_Array_any(), gopurs_runtime.Func(func(v_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Bool((gopurs_runtime.Apply2(eq_0, v_4, x_3).IntVal) != (0))
}))
_ = __local_var_4_2
// TAST (Let): __local_var_5_3 shape=App(Var) bindingType=(ADT ["Control","Monad","ST","Internal","ST"] [(TypeVar r), (Array (TypeVar a))])
__local_var_5_3 := gopurs_runtime.Apply(Get_Data_Array_ST_unsafeFreeze(), arr_2_0)
_ = __local_var_5_3
__local_var_6_5 := gopurs_runtime.Apply(__local_var_5_3, gopurs_runtime.Value{})
_ = __local_var_6_5
e_6_4 := gopurs_runtime.Bool(((gopurs_runtime.Apply(__local_var_4_2, __local_var_6_5).IntVal) != (0)) != (true))
_ = e_6_4
var __t7 gopurs_runtime.Value
{
if (e_6_4.IntVal) != (0) {
__t7 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
__local_var_7_6 := gopurs_runtime.Apply(gopurs_runtime.Apply2(Get_Data_Array_ST_push(), x_3, arr_2_0), gopurs_runtime.Value{})
_ = __local_var_7_6
return Get_Data_Unit_unit()
})
goto end_branch_7
} else {

}
}
{
__t7 = gopurs_runtime.Func(func(_ gopurs_runtime.Value) gopurs_runtime.Value {
return Get_Data_Unit_unit()
})
}
end_branch_7:
return gopurs_runtime.Apply(__t7, gopurs_runtime.Value{})
})
})), gopurs_runtime.Value{})
_ = _dollar___unused_3_1
return gopurs_runtime.Apply(gopurs_runtime.Apply(Get_Data_Array_ST_unsafeFreeze(), arr_2_0), gopurs_runtime.Value{})
})).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_nubEq(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
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
					arr := *(*[]gopurs_runtime.Value)(gopurs_runtime.Apply2(Get_Data_Semigroup_concatArray(), gopurs_runtime.Array(xs_1), func() gopurs_runtime.Value {
arr_val_foldlArray1 := gopurs_runtime.Array(xs_1)
_ = arr_val_foldlArray1
res_go_foldlArray1 := gopurs_runtime.Array(Call_Data_Array_nubByEq(eq_0, ys_2))
_ = res_go_foldlArray1
arr_go_foldlArray1 := (*[]gopurs_runtime.Value)(arr_val_foldlArray1.UnsafePtr)
_ = arr_go_foldlArray1
for _, v_foldlArray1 := range *arr_go_foldlArray1 {
res_go_foldlArray1 = gopurs_runtime.Apply2(gopurs_runtime.Func2(func(b_3 gopurs_runtime.Value, a_4 gopurs_runtime.Value) gopurs_runtime.Value {
return gopurs_runtime.Array(Call_Data_Array_deleteBy(eq_0, a_4, func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(b_3.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()))
}), res_go_foldlArray1, v_foldlArray1)
}
return res_go_foldlArray1
}()).UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
}

func Call_Data_Array_union(dictEq_0_loop *Constructor_Data_Eq_Eq[gopurs_runtime.Value]) gopurs_runtime.Value {
var dictEq_0 *Constructor_Data_Eq_Eq[gopurs_runtime.Value] = dictEq_0_loop
_ = dictEq_0
return gopurs_runtime.Apply(Get_Data_Array_unionBy(), gopurs_runtime.Box(dictEq_0.V0))
}

func Call_Data_Array_alterAt(i_0_loop int64, f_1_loop gopurs_runtime.Value, xs_2_loop []gopurs_runtime.Value) struct{V0 gopurs_runtime.Value; V1 bool} {
var i_0 int64 = i_0_loop
_ = i_0
var f_1 gopurs_runtime.Value = f_1_loop
_ = f_1
var xs_2 []gopurs_runtime.Value = xs_2_loop
_ = xs_2
// TAST (Let): __local_var_3_0 shape=UncurriedApp(Var) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
__local_var_3_0 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array_indexImpl(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Array(xs_2), gopurs_runtime.Int(i_0)))
_ = __local_var_3_0
var __t3 gopurs_runtime.Value
{
if (__local_var_3_0 == nil) {
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer((*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(nil))}
goto end_branch_3
} else {

}
}
{
if (__local_var_3_0 != nil) {
// TAST (Let): v_4_1 shape=App(Other) bindingType=(ADT ["Data","Maybe","Maybe"] [(TypeVar a)])
v_4_1 := gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.Apply(f_1, (__local_var_3_0).V0))
_ = v_4_1
var __t2 *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]
{
if (v_4_1 == nil) {
__t2 = Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp4(Get_Data_Array__deleteAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(i_0), gopurs_runtime.Array(xs_2))))
goto end_branch_2
} else {

}
}
{
if (v_4_1 != nil) {
__t2 = Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](gopurs_runtime.UncurriedApp5(Get_Data_Array__updateAt(), Get_Data_Maybe_Just(), gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](func() gopurs_runtime.Value {
				_v := struct{V0 gopurs_runtime.Value; V1 bool}{gopurs_runtime.Value{}, false}
				if _v.V1 {
					return gopurs_runtime.Box(&Constructor_Data_Maybe_Just[gopurs_runtime.Value]{V0: _v.V0})
				}
				return gopurs_runtime.Box(&Constructor_Data_Maybe_Nothing[gopurs_runtime.Value]{})
			}()))}, gopurs_runtime.Int(i_0), (v_4_1).V0, gopurs_runtime.Array(xs_2))))
goto end_branch_2
} else {

}
}
{
__t2 = func() *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] { panic("Failed pattern match") }()
}
end_branch_2:
__t3 = gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(__t2))}
goto end_branch_3
} else {

}
}
{
__t3 = func() gopurs_runtime.Value { panic("Failed pattern match") }()
}
end_branch_3:
return func() struct{V0 gopurs_runtime.Value; V1 bool} {
				_v := gopurs_runtime.Value{Type: 9, IntVal: 930809136, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_2807397954_3094389156(Rebox_Data_Array_3094389156_2807397954(gopurs_runtime.CoerceToStruct[Constructor_Data_Maybe_Just[gopurs_runtime.Value]](__t3))))}
				if _v.Type == 9 && _v.IntVal == 3562950408 && _v.UnsafePtr != nil {
					return struct{V0 gopurs_runtime.Value; V1 bool}{V0: (*Constructor_Data_Maybe_Just[gopurs_runtime.Value])(_v.UnsafePtr).V0, V1: true}
				}
				return struct{V0 gopurs_runtime.Value; V1 bool}{V0: gopurs_runtime.Value{}, V1: false}
			}()
}

func Call_Data_Array_all(__local_var_0_loop gopurs_runtime.Value, __local_var_1_loop []gopurs_runtime.Value) bool {
var __local_var_0 gopurs_runtime.Value = __local_var_0_loop
_ = __local_var_0
var __local_var_1 []gopurs_runtime.Value = __local_var_1_loop
_ = __local_var_1
return (gopurs_runtime.UncurriedApp2(Get_Data_Array_allImpl(), __local_var_0, gopurs_runtime.Array(__local_var_1)).IntVal) != (0)
}

func Rebox_Data_Array_138441832_1728839155(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]{}
		out.V0 = in.V0.IntVal
		out.V1 = in.V1
	return out
}

func Rebox_Data_Array_138441832_1807945000(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, []gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, []gopurs_runtime.Value]{}
		out.V0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(in.V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
		out.V1 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(in.V1.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
	return out
}

func Rebox_Data_Array_138441832_3415943795(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]{}
		out.V0 = in.V0
		out.V1 = in.V1.IntVal
	return out
}

func Rebox_Data_Array_1413047506_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(Rebox_Data_Array_3415943795_138441832(in.V0))}
	return out
}

func Rebox_Data_Array_1728839155_138441832(in *Constructor_Data_Tuple_Tuple[int64, gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Int(in.V0)
		out.V1 = in.V1
	return out
}

func Rebox_Data_Array_1807945000_138441832(in *Constructor_Data_Tuple_Tuple[[]gopurs_runtime.Value, []gopurs_runtime.Value]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Array(in.V0)
		out.V1 = gopurs_runtime.Array(in.V1)
	return out
}

func Rebox_Data_Array_2206451173_3603546092(in *Constructor_Control_Monad_Rec_Class_Done[struct{
	a gopurs_runtime.Value
	b int64
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Done[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
	return out
}

func Rebox_Data_Array_2660862937_4008603408(in *Constructor_Control_Monad_Rec_Class_Loop[struct{
	a gopurs_runtime.Value
	b int64
}, gopurs_runtime.Value]) *Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Control_Monad_Rec_Class_Loop[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict2("a", "b", orig.a, gopurs_runtime.Int(orig.b))
				}()
	return out
}

func Rebox_Data_Array_2762242827_3094389156(in *Constructor_Data_Maybe_Just[struct{
	head gopurs_runtime.Value
	tail []gopurs_runtime.Value
}]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = func() gopurs_runtime.Value {
				orig := in.V0
				_ = orig
				return gopurs_runtime.RecordDict2("head", "tail", orig.head, gopurs_runtime.Array(orig.tail))
				}()
	return out
}

func Rebox_Data_Array_2807397954_3094389156(in *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Array(in.V0)
	return out
}

func Rebox_Data_Array_3094389156_1170268447(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[int64] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[int64]{}
		out.V0 = in.V0.IntVal
	return out
}

func Rebox_Data_Array_3094389156_1413047506(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]]{}
		out.V0 = Rebox_Data_Array_138441832_3415943795(gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0))
	return out
}

func Rebox_Data_Array_3094389156_2807397954(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[[]gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[[]gopurs_runtime.Value]{}
		out.V0 = func() []gopurs_runtime.Value {
					arr := *(*[]gopurs_runtime.Value)(in.V0.UnsafePtr)
					unboxed := make([]gopurs_runtime.Value, len(arr))
					for i, v := range arr { unboxed[i] = v }
					return unboxed
				}()
	return out
}

func Rebox_Data_Array_3094389156_4010058633(in *Constructor_Data_Maybe_Just[gopurs_runtime.Value]) *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = gopurs_runtime.CoerceToStruct[Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]](in.V0)
	return out
}

func Rebox_Data_Array_3308271157_4177771502(in *Constructor_Data_Ord_Ord[int64]) *Constructor_Data_Ord_Ord[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Ord_Ord[gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Array_3415943795_138441832(in *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, int64]) *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = gopurs_runtime.Int(in.V1)
	return out
}

func Rebox_Data_Array_394862806_3988677819(in *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]) *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]] {
	if in == nil { return nil }
	out := &Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Array_3988677819_394862806(in *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, *Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Array_ST_Iterator_Iterator[gopurs_runtime.Value, gopurs_runtime.Value]{}
		out.V0 = in.V0
		out.V1 = in.V1
	return out
}

func Rebox_Data_Array_4010058633_3094389156(in *Constructor_Data_Maybe_Just[*Constructor_Data_Tuple_Tuple[gopurs_runtime.Value, gopurs_runtime.Value]]) *Constructor_Data_Maybe_Just[gopurs_runtime.Value] {
	if in == nil { return nil }
	out := &Constructor_Data_Maybe_Just[gopurs_runtime.Value]{}
		out.V0 = gopurs_runtime.Value{Type: 9, IntVal: 2339352186, UnsafePtr: unsafe.Pointer(in.V0)}
	return out
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
